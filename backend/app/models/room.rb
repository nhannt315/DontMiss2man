# frozen_string_literal: true

class Room < ApplicationRecord
  has_many :images, dependent: :destroy
  belongs_to :building
  belongs_to :agent
end
