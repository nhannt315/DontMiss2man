FactoryBot.define do
  factory :image do
    url { Faker::Internet.url }
    description { Faker::Lorem.word }
    association(:room)
  end
end
