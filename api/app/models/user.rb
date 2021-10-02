# frozen_string_literal: true

class User < ApplicationRecord
  extend Devise::Models
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable, :confirmable
  include DeviseTokenAuth::Concerns::User
  has_many :user_rooms, dependent: :delete_all
  has_many :rooms, through: :user_rooms



  def as_json(*)
    super.tap do |hash|
      hash["favorites"] = rooms.map(&:id).compact
    end
  end
end
