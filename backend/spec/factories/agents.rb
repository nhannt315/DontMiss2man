FactoryBot.define do
  factory :agent do
    name { Faker::Name.name }
    address { Faker::Address.full_address }
    working_time { Faker::Lorem.sentence }
    telephone_number { Faker::PhoneNumber.phone_number_with_country_code }
    email { Faker::Internet.email }
    photo_url { Faker::Internet.url }
    slogan { Faker::Lorem.sentence }
    access { Faker::Lorem.sentence }
  end
end
