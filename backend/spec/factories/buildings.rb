FactoryBot.define do
  factory :building do
    TYPES = ["マンション", "アパート"]
    STRUCTURES = ["木造", "鉄構造", "鉄筋コンクリート"]
    name { Faker::Lorem.word }
    address { Faker::Address.full_address }
    access { Faker::Types.rb_array(len: rand(1..10)) }
    year_built { Faker::Time.between(from: 50.years.ago, to: Time.now) }
    building_type { TYPES.sample }
    structure { STRUCTURES.sample }
    storeys { rand(2..40) }
    underground_storeys { rand(2..4) }
    photo_url { Faker::Lorem.sentence }
    latitude { Faker::Address.latitude }
    longitude { Faker::Address.longitude }
    average_fee { rand(5.0..400.1) }
    average_size { rand(10.0..400.1) }
  end
end
