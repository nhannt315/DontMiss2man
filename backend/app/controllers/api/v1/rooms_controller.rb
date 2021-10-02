# frozen_string_literal: true

class Api::V1::RoomsController < ApplicationController
  def show
    room = Room.find(params[:id])
    render json: {data: room}, include: [:images, :agent, :building]
  end
end
