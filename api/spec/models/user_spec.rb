require "rails_helper"

RSpec.describe User, type: :model do
  let!(:user) { create :user }
  describe "validations" do
    it "email should not be nil" do
      expect(user.email).not_to be nil
    end
    it "email's format" do
      expect(user.email).to match /\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z/
    end
  end
end
