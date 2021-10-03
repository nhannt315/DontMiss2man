# frozen_string_literal: true

class Api::V1::AuthenticationsController < ApplicationController
  def login
    @user = User.find_by_email(login_params[:email])
    if @user&.authenticate(login_params[:password])
      token = JsonWebToken.encode(user_id: @user.id)
      time = Time.now + 24.hours.to_i
      render json: { token: token, exp: time.strftime("%m-%d-%Y %H:%M"), email: @user.email }, status: :ok
    else
      render_error :unauthorized, {error: "unauthorized"}
    end

  end

  def register
    @user = User.new(registration_params)
    if @user.save
      access_token = JsonWebToken.encode(user_id: @user.id)
      time = Time.now + 24.hours.to_i
      render json: {token: access_token, exp: time.strftime("%m-%d-%Y %H:%M"), email: @user.email}, status: :ok
    else
      render_error :unprocessable_entity, @user.errors.full_messages
    end
  end

  private

  def login_params
    params.permit(:email, :password)
  end

  def registration_params
    params.permit(:email, :password, :password_confirmation)
  end
end
