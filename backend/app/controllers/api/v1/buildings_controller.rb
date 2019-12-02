# frozen_string_literal: true

class Api::V1::BuildingsController < ApplicationController
  before_action :pagination_params, only: [:index]

  def index
    case @sort
    when "newly_built"
      @building_list = Building.newly_built
    when "cheapest"
      @building_list = Building.cheapest
    when "most_expensive"
      @building_list = Building.most_expensive
    when "largest"
      @building_list = Building.largest
    else
      @building_list = Building.all
      puts "ALLLLL"
    end
    @building_list = @building_list.page(@page).per(@per_page).eager_load :rooms
  end


  private
  def pagination_params
    @page = params[:page] || Settings.building_list.default_page
    @per_page = params[:per_page] || Settings.building_list.default_items_per_page
    @sort = params[:sort]
  end
end
