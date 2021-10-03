# frozen_string_literal: true

class User < ApplicationRecord
  has_secure_password
  validates :email, presence: true, uniqueness: true
  validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }
  validates :password,
            length: { minimum: 6 },
            if: -> { new_record? || !password.nil? }

  has_many :user_rooms, dependent: :delete_all
  has_many :rooms, through: :user_rooms

  def as_json(*)
    super.tap do |hash|
      hash["favorites"] = rooms.map(&:id).compact
    end
  end
end
