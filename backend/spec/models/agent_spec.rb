require "rails_helper"

RSpec.describe Agent, type: :model do
  let!(:agent) { create :agent }
  describe "associations" do
    it { should have_many :rooms }
  end

  describe "validations" do
    it { validate_presence_of :name }
    it { validate_presence_of :address }
    it { validate_presence_of :telephone_number }
  end
end
