# frozen_string_literal: true

class Building < ApplicationRecord
  enum condition_type: {distance: 0, travel_time: 1}
  has_many :rooms, dependent: :destroy
  scope :newly_built, -> { order(year_built: :desc) }
  scope :cheapest, -> { order(average_fee: :asc) }
  scope :most_expensive, -> { order(average_fee: :desc) }
  scope :largest, -> { order(average_size: :desc) }
  scope :nearest, -> { where.not(distance: nil).order(distance: :asc) }
  scope :filter_by_year_built, ->(year) { where("year_built > ?", year) }

  serialize :access, Array

  def office
    Office.find(office_id)
  end
end
