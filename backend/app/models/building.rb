# frozen_string_literal: true

class Building < ApplicationRecord
  has_many :rooms, dependent: :destroy

  serialize :access, Array
end
