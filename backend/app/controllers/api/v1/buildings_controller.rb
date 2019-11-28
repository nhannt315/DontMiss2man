# frozen_string_literal: true

class Api::V1::BuildingsController < ApplicationController
  before_action :pagination_params, only: [:index]

  def index
    @building_list = Building.all.page(@page).per(@per_page).eager_load :rooms
  end


  private
  def pagination_params
    @page = params[:page] || Settings.building_list.default_page
    @per_page = params[:per_page] || Settings.building_list.default_items_per_page
    @sort = params[:sort]
  end
end
