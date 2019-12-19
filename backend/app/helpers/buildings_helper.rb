# frozen_string_literal: true

module BuildingsHelper
  def generate_condition_query(params)
    upper_fee = (params[:upper_fee]&.to_i)
    upper_size = (params[:upper_size]&.to_i)
    condition = {
      rooms: {rent_fee: params[:lower_fee].to_i..upper_fee, size: params[:lower_size].to_i..upper_size}
    }
    condition[:rooms][:layout] = params[:layout_types] if params[:layout_types].present?
    condition[:building_type] = params[:building_type] if params[:building_type].present?
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
