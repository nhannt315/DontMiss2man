# frozen_string_literal: true

class DeviseMailer < Devise::Mailer
  def confirmation_instructions(record, token, opts = {})
    @token = token
    initialize_from_record record
    mail headers_for(:confirmation_instructions, opts)
  end
end
