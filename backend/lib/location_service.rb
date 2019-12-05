# frozen_string_literal: true

class LocationService
  API_KEY = ENV["google_api_key"]
  BASE_URL = "https://suumo.jp"
  @@building_by_address = 0
  @@building_by_detail = 0

  def self.get_lat_lng_from_address(address)
    url = "https://maps.googleapis.com/maps/api/geocode/json?address=#{CGI.escape(address)}&key=#{API_KEY}"
    response = RestClient.get url
    lat = JSON.parse(response)["results"][0].dig("geometry", "location", "lat")
    lng = JSON.parse(response)["results"][0].dig("geometry", "location", "lng")
    [lat, lng]
  end

  def self.get_walking_time(src_lat, src_lng, des_lat, des_lng)
    url = "https://maps.googleapis.com/maps/api/distancematrix/json?units=metric"\
          "&origins=#{src_lat},#{src_lng}&destinations=#{des_lat},#{des_lng}"\
          "&key=#{API_KEY}&mode=walking"
    response = RestClient.get url
    travel_time_in_sec = JSON.parse(response)["rows"][0]["elements"][0]["duration"]["value"]
    travel_time_in_sec.to_f / 60
  end

  def self.get_lat_lng_from_room_url(room_url, address)
    room_url = "#{BASE_URL}#{room_url}"
    detail_url = room_url.gsub(/\?bc=/, "kankyo/\\0")
    begin
      response = RestClient.get detail_url
    rescue RestClient::NotFound
      @@building_by_address = @@building_by_address + 1
      Rails.logger.warn "Get lat,lng from address #{address}"
      puts "Get lat,lng from address #{address}"
      puts "Total by address #{@@building_by_address}"
      return get_lat_lng_from_address address
    end
    begin
      root_page_node = Nokogiri.HTML response
      form_url = root_page_node.css("#js-timesForm")[0]["action"]
      puts "Get lat, lng from detail page #{address}"
      @@building_by_detail = @@building_by_detail + 1
      puts "Total by detail #{@@building_by_detail}"
      return [form_url[/(?<=ido=)(.*?)(?=&keido)/].to_f, form_url[/(?<=&keido=)[+-]?([0-9]*[.])?[0-9]+/].to_f]
    rescue NoMethodError
      @@building_by_address = @@building_by_address + 1
      Rails.logger.warn "Get lat,lng from address #{address}"
      puts "Get lat,lng from address #{address}"
      puts "Total by address #{@@building_by_address}"
      return get_lat_lng_from_address address
    end
  end

  def self.test
  end
end
