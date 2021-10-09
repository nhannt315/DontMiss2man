# frozen_string_literal: true

class Api::V1::UsersController < ApplicationController
  before_action :authorize_request, only: [:info]

  def info
    render json: {
      email: @current_user.email
    }
  end
end
