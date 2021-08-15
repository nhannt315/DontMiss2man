FactoryBot.define do
  factory :image do
    url { Faker::Internet.url }
    description { Faker::Lorem.word }
    trait :with_room do
      after :build do |image|
        image.room = FactoryBot.build :room
      end
    end
  end
end
