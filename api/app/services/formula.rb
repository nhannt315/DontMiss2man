# frozen_string_literal: true

module Formula
    ##
    # Haversine Distance Calculation
    #
    # Accepts two coordinates in the form
    # of a tuple. I.e.
    #   geo_a  Array(Num, Num)
    #   geo_b  Array(Num, Num)
    #   miles  Boolean
    #
    # Returns the distance between these two
    # points in either miles or kilometers
    def self.haversine_distance(geo_a, geo_b, miles = false)
      # Get latitude and longitude
      lat1, lon1 = geo_a
      lat2, lon2 = geo_b
  
      # Calculate radial arcs for latitude and longitude
      d_lat = (lat2 - lat1) * Math::PI / 180
      d_lng = (lon2 - lon1) * Math::PI / 180
  
      a = Math.sin(d_lat / 2) *
          Math.sin(d_lat / 2) +
          Math.cos(lat1 * Math::PI / 180) *
              Math.cos(lat2 * Math::PI / 180) *
              Math.sin(d_lng / 2) * Math.sin(d_lng / 2)
  
      c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
  
      6371 * c * (miles ? 1 / 1.6 : 1)
    end
  
    def self.convert_currency(src)
      return 0 if src.strip == "-"
      return src.chomp("万円").to_f * 10_000 if src[-2] == "万"
  
      src.chomp("円").to_f * 10_000
    end
  end
  