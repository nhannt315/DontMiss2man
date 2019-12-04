# frozen_string_literal: true

class Api::V1::BuildingsController < ApplicationController
  before_action :get_params, only: [:index]

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
                     else
                       Building.eager_load(:rooms).all
                     end
    if @is_advanced_search
      @building_list = @building_list.where(advance_search_condition)
      if params[:years_built]
        @building_list = @building_list.where("year_built > ?", Date.strptime(params[:years_built].to_s, "%Y"))
      end
    end
    @building_list = @building_list.page(@page).per @per_page
  end

  private

  def get_params
    @page = params[:page] || Settings.building_list.default_page
    @per_page = params[:per_page] || Settings.building_list.default_items_per_page
    @sort = params[:sort]
    @is_advanced_search = params.key?(:upper_fee) || params.key?(:lower_fee) ||
                          params.key?(:no_management_fee) || params.key?(:no_reikin) ||
                          params.key?(:no_shikikin) || params.key?(:years_built) ||
                          params.key?(:layout_types) || params.key?(:building_type) ||
                          params.key?(:upper_size) || params.key?(:lower_size)
  end

  def advance_search_condition
    upper_fee = (params[:upper_fee]&.to_i)
    upper_size = (params[:upper_size]&.to_i)
    condition = {
      rooms: {rent_fee: params[:lower_fee].to_i..upper_fee, size: params[:lower_size].to_i..upper_size}
    }
    condition[:rooms][:layout] = params[:layout_types] if params[:layout_types].present?
    condition[:type] = params[:building_type] if params[:building_type].present?
    if params[:no_management_fee] && ActiveModel::Type::Boolean.new.cast(params[:no_management_fee])
      condition[:rooms][:management_cost] = 0
    end
    condition[:rooms][:reikin] = 0 if params[:no_reikin] && ActiveModel::Type::Boolean.new.cast(params[:no_reikin])
    if params[:no_shikikin] && ActiveModel::Type::Boolean.new.cast(params[:no_shikikin])
      condition[:rooms][:shikikin] = 0
    end
    condition
  end
end
