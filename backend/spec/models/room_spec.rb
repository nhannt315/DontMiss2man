require "rails_helper"

RSpec.describe Room, type: :model do
  let!(:room) { create :room }
  describe "associations" do
    it { should belong_to :building }
    it { should belong_to :agent }
    it { should have_many :images }
  end

  describe "validations" do
    it {validate_presence_of :rent_fee}
    it {validate_presence_of :management_cost}
    it {validate_presence_of :reikin}
    it {validate_presence_of :shikikin}
    it {validate_presence_of :size}
    it {validate_presence_of :direction}
    it {validate_presence_of :layout}
    it {validate_presence_of :move_in}
    it {validate_presence_of :building}
    it {validate_presence_of :agent}
  end
end
