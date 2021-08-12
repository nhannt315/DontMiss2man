# frozen_string_literal: true

class BuildingCalculator
    class << self
      def calculate_average_size!
        Building.find_each do |building|
          total_size = 0
          min_fee_room = Room.where(building_id: building.id).order('rent_fee + management_cost ASC')[0]
          min_fee = min_fee_room.rent_fee + min_fee_room.management_cost
          building.rooms.each do |room|
            total_size += room.size.to_i
          end
          building.average_fee = min_fee
          building.average_size = total_size.to_f / building.rooms.count
          building.save!
        end
      end
  
      def calculate_distance!
        Building.find_each do |building|
          office = building.office
          building.distance = Formula.haversine_distance(
              [building.latitude, building.longitude],
              [office.latitude, office.longitude],
            )
          building.save!
        end
      end
    end
  end
  