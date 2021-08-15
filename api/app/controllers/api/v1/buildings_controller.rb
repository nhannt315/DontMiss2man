# frozen_string_literal: true

class Api::V1::BuildingsController < ApplicationController
  before_action :check_params, only: [:index]

  def index
    @building_list = case @sort
                     when "newly_built"
                       Building.eager_load(:rooms).newly_built
                     when "cheapest"
                       Building.eager_load(:rooms).cheapest
                     when "most_expensive"
                       Building.eager_load(:rooms).most_expensive
                     when "largest"
                       Building.eager_load(:rooms).largest
                     when "nearest"
                       Building.eager_load(:rooms).nearest
                     else
                       Building.eager_load(:rooms).all
                     end
    if @is_advanced_search
      @building_list = @building_list.where(generate_condition_query(params))
      @building_list = @building_list.filter_by_year_built(Date.strptime(params[:years_built].to_s, "%Y")) if params[:years_built]
      if params[:with_furniture] && ActiveModel::Type::Boolean.new.cast(params[:with_furniture])
        @building_list = @building_list.where "rooms.facilities LIKE ? OR rooms.facilities LIKE ?", "%家具付%", "%家電付%"
      end
    end
    @building_list = @building_list.page(@page).per @per_page
  end

  private

  def check_params
    @page = params[:page] || Settings.building_list.default_page
    @per_page = params[:per_page] || Settings.building_list.default_items_per_page
    @sort = params[:sort]
    @is_advanced_search = params.key?(:upper_fee) || params.key?(:lower_fee) ||
                          params.key?(:no_management_fee) || params.key?(:no_reikin) ||
                          params.key?(:no_shikikin) || params.key?(:years_built) ||
                          params.key?(:layout_types) || params.key?(:building_type) ||
                          params.key?(:upper_size) || params.key?(:lower_size)
  end
end
