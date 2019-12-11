# frozen_string_literal: true

class ApplicationController < ActionController::API
  include DeviseTokenAuth::Concerns::SetUserByToken
  include ExceptionHandler

  before_action :set_locale

  protected

  def set_locale
    # Remove inappropriate/unnecessary ones
    #I18n.locale = params[:locale] || # Request parameter
    #    session[:locale] || # Current session
    #    (current_user.preferred_locale if user_signed_in?) || # Model saved configuration
    #    extract_locale_from_accept_language_header || # Language header - browser config
    #    I18n.default_locale # Set in your config files, english by super-default
    I18n.locale = params[:locale] || I18n.default_locale
  end

  # Extract language from request header
  #def get_locale_from_accept_language_header
  #  if request.env['HTTP_ACCEPT_LANGUAGE']
  #    lg = request.env['HTTP_ACCEPT_LANGUAGE'].scan(/^[a-z]{2}/).first.to_sym
  #    lg.in?([:en, YOUR_AVAILABLE_LANGUAGES]) ? lg : nil
  #  end
  #end
end
