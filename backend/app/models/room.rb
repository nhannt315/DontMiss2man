class Room < ApplicationRecord
  has_many :images
  belongs_to :building
  belongs_to :agent
end
