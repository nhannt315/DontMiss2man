namespace :suumo_crawl do
  desc "Crawl building and rooms info from suumo"
  BASE_URL = "https://suumo.jp"
  task mansion: :environment do
    ROOT_URL = "https://suumo.jp/jj/chintai/ichiran/FR301FC001/?shkr1=03&cb=0.0&shkr3=03&shkr2=03&mt=9999999&sngz=&sc=13103&ar=030&bs=040&shkr4=03&ct=9999999&cn=9999999&mb=0&ta=13&et=9999999&page=".freeze
    current_page = 1
    total_page = 1
    loop do
      Rails.logger.info "Starting scraping page #{current_page}"
      response = RestClient.get "#{ROOT_URL}#{current_page}"
      root_page = Nokogiri.HTML response
      buildings = root_page.css("#js-bukkenList > ul > li > div")
      total_page = root_page.css("#js-leftColumnForm > div.pagination_set > div.pagination.pagination_set-nav > ol > li:nth-child(11) > a").text.to_i if current_page == 1
      buildings.each do |building_node|
        building = Building.new
        building.name = building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-title").text
        check_and_delete_building building.name
        building.address = building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-body > ul > li.cassetteitem_detail-col1").text
        building.photo_url = building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-object > div > div > img")[0]["rel"]
        access = []
        building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-body > ul > li.cassetteitem_detail-col2 > div").each do |element|
          access.push element.text
        end
        building.access = access
        first_room_url = building_node.css("div.cassetteitem-item > table > tbody:nth-child(2) > tr > td.ui-text--midium.ui-text--bold > a")[0]["href"]
        building.latitude, building.longitude = LocationService.get_lat_lng_from_room_url first_room_url, building.name
        unless check_building_condition(building.latitude, building.longitude, building)
          Rails.logger.info "Do not match condition, skip! #{building.name}"
          next
        end
        Rails.logger.info building.name
        Rails.logger.info "#{building.name} Successfully created" if building.save!
        building_node.css("div.cassetteitem-item > table > tbody").each_with_index do |room_node, index|
          room_url = room_node.css("tr > td.ui-text--midium.ui-text--bold > a")[0]["href"]
          extract_room_information building, "#{BASE_URL}#{room_url}", index
        end
      end
      current_page = current_page + 1
      break if current_page > total_page
    end
    Rake::Task["building:calculate_average"].execute
    Rake::Task["building:calculate_distance"].execute
  end

  def check_and_delete_building building_name
    building = Building.find_by name: building_name
    Rails.logger.info "Deleting building #{building_name}" if building&.destroy
  end

  def check_building_condition lat, lng, building
    return false unless lat || lng
    distance_from_head_office = Formula.haversine_distance([lat, lng], [Settings.head_office_lat, Settings.head_office_lng])
    building.distance = distance_from_head_office
    if distance_from_head_office <= Settings.fixed_distance
      building.condition_type = 0
      return true
    end
    return false if distance_from_head_office > Settings.fixed_distance + 0.2
    if LocationService.get_walking_time(Settings.head_office_lat, Settings.head_office_lng, lat, lng) <= Settings.max_travel_time_in_mins
      building.condition_type = 1
      return true
    end
    false
  end

  def extract_room_information building, url, index
    begin
      response = RestClient.get url
    rescue RestClient::NotFound
      Rails.logger.error "Url not found \n #{url}"
      return
    end
    begin
      root_page = Nokogiri.HTML response
      room = Room.new
      room.rent_fee = Formula.convert_currency root_page.css("#js-view_gallery > div.property_view_note > div > div:nth-child(1) > span.property_view_note-emphasis").text
      room.management_cost = root_page.css("#js-view_gallery > div.property_view_note > div > div:nth-child(1) > span:nth-child(2)").text[/[0-9]+(?=円)/].to_i
      room.shikikin = Formula.convert_currency root_page.css("#js-view_gallery > div.property_view_note > div > div:nth-child(2) > span:nth-child(1)").text[4..-1]
      room.reikin = Formula.convert_currency root_page.css("#js-view_gallery > div.property_view_note > div > div:nth-child(2) > span:nth-child(1)").text[4..-1]
      room.caution_fee = Formula.convert_currency root_page.css("#js-view_gallery > div.property_view_note > div > div:nth-child(2) > span:nth-child(3)").text[5..-1]
      room.layout = root_page.css("#js-view_gallery > div.l-property_view_table > table > tr:nth-child(3) > td:nth-child(2)").text
      room.size = root_page.css("#js-view_gallery > div.l-property_view_table > table > tr:nth-child(3) > td:nth-child(4)").text.chomp("m2").to_f
      room.direction = root_page.css("#js-view_gallery > div.l-property_view_table > table > tr:nth-child(5) > td:nth-child(2)").text
      room.floor = root_page.css("#js-view_gallery > div.l-property_view_table > table > tr:nth-child(4) > td:nth-child(4)").text.chomp("階").to_i
      room.facilities = root_page.css("#bkdt-option > div > ul > li").text
      room_info_html = root_page.css("#contents > div > table").to_s.strip.squish
      room.layout_detail = room_info_html[/(?<=間取り詳細<\/th> <td>)(.*?)(?=<\/td>)/]
      room.car_park = room_info_html[/(?<=駐車場<\/th> <td>)(.*?)(?=<\/td>)/]
      room.condition = room_info_html[/(?<=条件<\/th> <td>)(.*?)(?=<\/td>)/]
      room.deal_type = room_info_html[/(?<=取引態様<\/th> <td>)(.*?)(?=<\/td>)/]
      room.move_in = room_info_html[/(?<=入居<\/th> <td>)(.*?)(?=<\/td>)/]
      room.suumo_id = room_info_html[/(?<=SUUMO<br>物件コード<\/th> <td>)(.*?)(?=<\/td>)/]
      room.note = room_info_html[/(?<=備考<\/th> <td colspan="3"> <ul class="inline_list"> <li>)(.*?)(?=<\/li>)/]
      room.guarantor = room_info_html[/(?<=保証人代行<\/th> <td colspan="3"> <ul class="inline_list"> <li>)(.*?)(?=<\/li>)/]
      room.other_fees = room_info_html[/(?<=ほか諸費用<\/th> <td colspan="3"> <ul class="inline_list"> <li>)(.*?)(?=<\/li>)/]
      room.last_update = Date.parse root_page.css("#contents > div.captiontext.l-space_medium").text[/\d{4}\/\d+\/\d+/]
      room.suumo_link = url

    rescue Exception
      byebug
    end

    room.agent_id = find_or_create_agent root_page

    if index.zero?
      building.structure = room_info_html[/(?<=構造<\/th> <td>)(.*?)(?=<\/td>)/]
      building.storeys = room_info_html[/(?<=階建<\/th> <td>)(.*?)(?=<\/td>)/][/[\s\d]+(?=階建)/].to_i
      building.underground_storeys = room_info_html[/(?<=階建<\/th> <td>)(.*?)(?=<\/td>)/][/(?<=地下)[\s\d]+/].to_i
      building.year_built = Date.strptime(room_info_html[/(?<=築年月<\/th> <td>)(.*?)(?=<\/td>)/], "%Y年")
      building.building_type = root_page.css("#js-view_gallery > div.l-property_view_table > table > tr:nth-child(5) > td:nth-child(4)").text
      building.save!
    end
    room.building_id = building.id
    Rails.logger.info "Room successfully created id: #{room.id}" if room.save!

    # Get room photolist
    get_room_photo_list room, root_page
  end


  def find_or_create_agent root_page_node
    agent_info_node = root_page_node.css("#contents > div.itemcassette.l-space_medium")
    name = agent_info_node.css("div.itemcassette-header > span.itemcassette-header-ttl").text
    agent = Agent.find_by name: name
    unless agent
      slogan = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-contents > div.itemcassette_desc").text
      address = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-contents > div.itemcassette_matrix > div.itemcassette_matrix-cell01").text
      working_time = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-contents > div.itemcassette_matrix > div.itemcassette_matrix-cell02").text
      access = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-contents > div.itemcassette_matrix > div.itemcassette_matrix-cell03").text
      telephone_number = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-contents > div.itemcassette_matrix > div.itemcassette_matrix-cell04 > span").text
      photo_url = agent_info_node.css("div.itemcassette-body > div.itemcassette-body-object > div > div.itemcassette_img-object > img")[0]["src"]
      agent = Agent.new(name: name, address: address, working_time: working_time,
                        telephone_number: telephone_number, photo_url: photo_url,
                        slogan: slogan, access: access)
      Rails.logger.info "Created agent #{name} successfully" if agent.save!
    end
    return agent.id
  end

  def get_room_photo_list room, root_page_node
    room_photo_list_node = root_page_node.css("#js-view_gallery-list > li > a > img")
    room_photo_list_node.each do |photo_node|
      image = Image.new(url: photo_node["data-src"], description: photo_node["alt"], room_id: room.id)
      image.save
      if image.description == "間取り図"
        room.layout_image_url = image.url
        room.save!
      end
    end
  end

end
