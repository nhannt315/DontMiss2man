class Building < ApplicationRecord
  has_many :rooms

  serialize :access, Array
end
