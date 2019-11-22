namespace :suumo_crawl do
  desc "TODO"
  task mansion: :environment do
    ROOT_URL = "https://suumo.jp/jj/chintai/ichiran/FR301FC001/?shkr1=03&cb=0.0&shkr3=03&shkr2=03&mt=9999999&sngz=&sc=13103&ar=030&bs=040&shkr4=03&ct=9999999&cn=9999999&mb=0&ta=13&et=9999999&page=".freeze
    current_page = 1
    response = RestClient.get "#{ROOT_URL}#{current_page}"
    root_page = Nokogiri.HTML response
    buildings = root_page.css("#js-bukkenList > ul > li > div")
    buildings.each do |building_node|
      building = Building.new
      building.title = building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-title").text
      building.address = building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-body > ul > li.cassetteitem_detail-col1").text
      access = []
      building_node.css("div.cassetteitem-detail > div.cassetteitem-detail-body > div > div.cassetteitem_content-body > ul > li.cassetteitem_detail-col2 > div").each do |element|
        access.push element.text
      end
      building.access = access
      

    end
  end



  API_KEY = "AIzaSyDgkKg98IqQL1oHJ7LUYc_fTO-wKq-YUnY"

  def get_long_lat_from_address address
    url = "https://maps.googleapis.com/maps/api/geocode/json?address=#{URI.encode(address)}&key=#{API_KEY}"
    response = RestClient.get url
    lat = JSON.parse(response)["results"][0].dig("geometry", "location", "lat")
    lng = JSON.parse(response)["results"][0].dig("geometry", "location", "lng")
    return [lat, lng]
  end

end
