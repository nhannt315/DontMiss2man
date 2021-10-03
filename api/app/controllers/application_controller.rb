# frozen_string_literal: true

class ApplicationController < ActionController::API
  include DeviseTokenAuth::Concerns::SetUserByToken
  include ExceptionHandler
  include BuildingsHelper

  before_action :set_locale

  protected

  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end

  def authorize_request
    header = request.headers["Authorization"]
    header = header.split(' ').last if header
    begin
      @decoded = JsonWebToken.decode(header)
      @current_user = User.find(@decoded[:user_id])
    rescue ActiveRecord::RecordNotFound => e
      render_error :unauthorized, e.message
    rescue JWT::DecodeError => e
      render_error :unauthorized, e.message
    end
  end

  def render_error status_code, message
    render json: {
      errors: message,
      code: status_code,
    }, status: status_code
  end
end
