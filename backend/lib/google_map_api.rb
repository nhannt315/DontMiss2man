# frozen_string_literal: true

class GoogleMapApi
  API_KEY = "AIzaSyDgkKg98IqQL1oHJ7LUYc_fTO-wKq-YUnY"
  def self.get_lat_lng_from_address(address)
    url = "https://maps.googleapis.com/maps/api/geocode/json?address=#{URI.encode(address)}&key=#{API_KEY}"
    response = RestClient.get url
    lat = JSON.parse(response)["results"][0].dig("geometry", "location", "lat")
    lng = JSON.parse(response)["results"][0].dig("geometry", "location", "lng")
    [lat, lng]
  end
end
