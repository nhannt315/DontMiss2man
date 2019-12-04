# frozen_string_literal: true

class Api::V1::BuildingsController < ApplicationController
  before_action :pagination_params, only: [:index]

  def index
    case @sort
    when "newly_built"
      @building_list = Building.eager_load(:rooms).newly_built
    when "cheapest"
      @building_list = Building.eager_load(:rooms).cheapest
    when "most_expensive"
      @building_list = Building.eager_load(:rooms).most_expensive
    when "largest"
      @building_list = Building.eager_load(:rooms).largest
    else
      @building_list = Building.eager_load(:rooms).all
    end
    if @is_advanced_search
      upper_fee = if params[:upper_fee]
                    params[:upper_fee].to_i
                  else
                    nil
                  end
      upper_size = if params[:upper_size]
                     params[:upper_size].to_i
                   else
                     nil
                   end
      condition = {
          rooms: {rent_fee: params[:lower_fee].to_i..upper_fee, size: params[:lower_size].to_i..upper_size}
      }
      condition[:rooms][:layout] = params[:layout_types] if params[:layout_types] && params[:layout_types].length > 0
      condition[:type] = params[:building_type] if params[:building_type] && params[:building_type].length > 0
      condition[:rooms][:management_cost] = 0 if params[:no_management_fee] && ActiveModel::Type::Boolean.new.cast(params[:no_management_fee])
      condition[:rooms][:reikin] = 0 if params[:no_reikin] && ActiveModel::Type::Boolean.new.cast(params[:no_reikin])
      condition[:rooms][:shikikin] = 0 if params[:no_shikikin] && ActiveModel::Type::Boolean.new.cast(params[:no_shikikin])
      @building_list = @building_list.where(condition)
      @building_list = @building_list.where('year_built > ?', Date.strptime(params[:years_built].to_s, "%Y")) if params[:years_built]
      @building_list = @building_list.page(@page).per @per_page
    else
      @building_list = @building_list.page(@page).per(@per_page)
    end
  end


  private

  def pagination_params
    @page = params[:page] || Settings.building_list.default_page
    @per_page = params[:per_page] || Settings.building_list.default_items_per_page
    @sort = params[:sort]
    @is_advanced_search = params.has_key?(:upper_fee) ||
        params.has_key?(:lower_fee) || params.has_key?(:no_management_fee) ||
        params.has_key?(:no_reikin) || params.has_key?(:no_shikikin) ||
        params.has_key?(:layout_types) || params.has_key?(:building_type) ||
        params.has_key?(:upper_size) || params.has_key?(:lower_size) ||
        params.has_key?(:years_built)
  end
end
