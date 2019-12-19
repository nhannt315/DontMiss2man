# frozen_string_literal: true

class Room < ApplicationRecord
  has_many :images, dependent: :destroy
  belongs_to :building
  belongs_to :agent
  has_many :user_rooms
  has_many :users, through: :user_rooms
end
