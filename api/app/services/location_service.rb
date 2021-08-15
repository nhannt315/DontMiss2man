# frozen_string_literal: true

class LocationService
    API_KEY = ENV.fetch('google_api_key')
    BASE_URL = 'https://suumo.jp'

    # @param String address
    # return [Float, Float] [latitude, longitude]
    def get_lat_lng_from_address(address)
      url = "https://maps.googleapis.com/maps/api/geocode/json?address=#{CGI.escape(address)}&key=#{API_KEY}"
      response = RestClient.get url
      lat = JSON.parse(response)["results"][0].dig("geometry", "location", "lat")
      lng = JSON.parse(response)["results"][0].dig("geometry", "location", "lng")
      [lat, lng]
    end

    def get_position_from_suumo_url(room_url, name)
      room_url = "#{BASE_URL}#{room_url}"
      detail_url = room_url.gsub(/\?bc=/, "kankyo/\\0")
      begin
        response = RestClient.get detail_url
        root_page_node = Nokogiri.HTML response
        form_url = root_page_node.css("#js-timesForm")[0]["action"]
        latitude = form_url[/(?<=ido=)(.*?)(?=&keido)/].to_f
        longitude = form_url[/(?<=&keido=)[+-]?([0-9]*[.])?[0-9]+/].to_f
        DontMiss2Man.debugging_logger.info(
          room_url: room_url,
          message: "Get lat, lng from detail page #{name}, lat: #{latitude}, lng: #{longitude}",
          )
        [latitude, longitude]
      rescue RestClient::NotFound, NoMethodError, RestClient::InternalServerError
        [nil, nil]
      end
    end

    def get_walking_time(src_lat, src_lng, des_lat, des_lng)
      url = "https://maps.googleapis.com/maps/api/distancematrix/json?units=metric"\
            "&origins=#{src_lat},#{src_lng}&destinations=#{des_lat},#{des_lng}"\
            "&key=#{API_KEY}&mode=walking"
      response = RestClient.get url
      travel_time_in_sec = JSON.parse(response)["rows"][0]["elements"][0]["duration"]["value"]
      travel_time_in_sec.to_f / 60
    end

  end
