require "rails_helper"

RSpec.describe "Rooms API" do
  # Initialize test data
  let!(:room) { create :room, :with_images, :with_building, :with_agent }
  let(:room_id) { room.id }
  let(:headers) { valid_headers }

  describe 'GET /rooms/:id' do
    before { get "/api/v1/rooms/#{room_id}", params: {}, headers: headers }

    context "when the record exists" do
      it "returns status code 200" do
        expect(response).to have_http_status 200
      end

      it "return the room info" do
        expect(json["data"]["id"]).to eq room_id
      end
    end

    context "when the record do not exist" do
      let(:room_id) { -1 }
      it "returns status code 404" do
        expect(response).to have_http_status 404
      end
    end
  end
end
