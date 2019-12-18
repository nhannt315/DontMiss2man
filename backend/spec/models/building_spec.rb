require "rails_helper"

RSpec.describe Building, type: :model do
  let!(:building) { create :building }
  describe "associations" do
    it { should have_many :rooms }
  end

  describe "validations" do
    it { validate_presence_of :name }
    it { validate_presence_of :address }
    it { validate_presence_of :longitude }
    it { validate_presence_of :latitude }
    it { validate_presence_of :year_built }
    it { validate_uniqueness_of :name }
  end

  describe "scope" do
    it "newly_built scope will return collections by year_built descending" do
      expect(Building.newly_built.to_sql).to eq Building.all.order(year_built: :desc).to_sql
    end

    it "cheapest scope will return collections by average_fee ascending" do
      expect(Building.cheapest.to_sql).to eq Building.all.order(average_fee: :asc).to_sql
    end

    it "most_expensive scope will return collections by average_fee descending" do
      expect(Building.most_expensive.to_sql).to eq Building.all.order(average_fee: :desc).to_sql
    end

    it "largest scope will return collections by average_size descending" do
      expect(Building.largest.to_sql).to eq Building.all.order(average_size: :desc).to_sql
    end
  end
end
