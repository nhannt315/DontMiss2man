require "rails_helper"

RSpec.describe Api::V1::BuildingsController, type: :controller do
  describe "GET #index" do
    before do
      get :index
    end

    it "returns http success" do
      expect(response).to have_http_status(:success)
    end

    it "response with JSON body containing expected Attributes" do

    end
  end
end
