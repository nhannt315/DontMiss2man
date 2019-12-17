require "rails_helper"

RSpec.describe Image, type: :model do
  let!(:image) { create :image }
  describe "association" do
    it { should belong_to :room }
  end

  describe "validations" do
    it { validate_presence_of :url }
  end
end
