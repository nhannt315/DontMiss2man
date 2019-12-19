class Api::V1::FavoritesController < ApplicationController
  before_action :authenticate_user!, only: [:create, :delete]
  before_action :get_param, only: [:create, :delete]

  def create
    return render json: {}, status: :conflict if current_user.room_ids.include? @room_id
    current_user.rooms << Room.find(@room_id)
    render json: {message: "Success"}, status: :ok
  end

  def delete
    return render json: {}, status: :not_found unless current_user.room_ids.include? @room_id
    current_user.rooms.delete @room_id
    render json: {message: "Success"}, status: :ok
  end

  private

  def get_param
    @room_id = params[:room_id]
    return render json: {}, status: :bad_request unless @room_id.present?
    render json: {}, status: :not_found if Room.where(id: @room_id).empty?
  end
end
