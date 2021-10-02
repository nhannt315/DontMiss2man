require "rails_helper"

RSpec.describe "Building API" do
  # Initialize test data
  let!(:buildings) { create_list :building, 100 }
  let(:headers) { valid_headers }
  let!(:per_page) { 10 }

  describe "GET /api/v1/buildings" do
    before { get "/api/v1/buildings", params: {per_page: per_page}, headers: headers }

    it "return status code 200" do
      expect(response).to have_http_status 200
    end

    it "return buildings list" do
      expect(json).not_to be_empty
    end

    it "Return the same number of items as declared in the parameter" do
      expect(json["data"]["list"].size).to eq per_page
    end
  end
end
