FactoryBot.define do
  factory :room do
    suumo_id { rand(100000..999999) }
    trait :with_agent do
      after :build do |room|
        room.agent = FactoryBot.build :agent
      end
    end
    trait :with_building do
      after :build do |room|
        room.building = FactoryBot.build :building
      end
    end
    trait :with_images do
      after :build do |room|
        room.images = FactoryBot.build_list(:image, rand(2..10), :with_room)
      end
    end
    rent_fee { rand(10000..500000) }
    reikin { rand(10000..500000) }
    shikikin { rand(10000..500000) }
    management_cost { rand(10000..500000) }
    caution_fee { rand(10000..500000) }
    layout_image_url { Faker::Internet.url }
    size { rand(7.0..300.0) }
    direction { Faker::Lorem.word }
    facilities { Faker::Lorem.paragraph }
    floor { rand(1..50) }
    car_park { Faker::Lorem.sentence }
    condition { Faker::Lorem.sentence }
    note { Faker::Lorem.paragraph }
    layout { Faker::Lorem.word }
    layout_detail { Faker::Lorem.word }
    deal_type { "仲介" }
    move_in_time { Faker::Time.between(from: Time.now, to: 1.month.from_now) }
    move_in { Faker::Lorem.sentence }
    damage_insurance { Faker::Lorem.sentence }
    guarantor { Faker::Lorem.sentence }
    other_fees { Faker::Lorem.sentence }
    other_initial_fees { Faker::Lorem.sentence }
    last_update { Faker::Time.between(from: 1.month.ago, to: Time.now) }
    suumo_link { Faker::Internet.url }
  end
end
