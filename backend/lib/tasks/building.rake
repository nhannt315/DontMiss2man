namespace :building do
  desc "TODO"
  task calculate_average: :environment do
    Building.find_each do |building|
      total_fee = 0
      total_size = 0
      building.rooms.each do |room|
        total_fee = total_fee + room.rent_fee + room.management_cost.to_i
        total_size = total_size + room.size.to_i
      end
      building.average_fee = total_fee.to_f / building.rooms.count
      building.average_size = total_size.to_f / building.rooms.count
      building.save!
      puts building.name
    end
  end

  task test_job: :environment do
    Rails.logger.warn "So sleepy, wanna go home now.."
    puts "So sleepy, wanna go home now.."
  end
end
