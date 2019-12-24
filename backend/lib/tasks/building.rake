namespace :building do
  desc "TODO"
  task calculate_average: :environment do
    Building.find_each do |building|
      total_size = 0
      min_fee_room = Room.where(building_id: building.id).order('rent_fee + management_cost ASC')[0]
      min_fee = min_fee_room.rent_fee + min_fee_room.management_cost
      building.rooms.each do |room|
        total_size = total_size + room.size.to_i
      end
      building.average_fee = min_fee
      building.average_size = total_size.to_f / building.rooms.count
      building.save!
    end
  end

  task calculate_distance: :environment do
    Building.find_each do |building|
      building.distance = Formula.haversine_distance([building.latitude, building.longitude], [Settings.head_office_lat, Settings.head_office_lng])
      building.save!
    end
  end

  task clean: :environment do
    Image.delete_all
    Agent.delete_all
    Room.delete_all
    Building.delete_all
  end
end
