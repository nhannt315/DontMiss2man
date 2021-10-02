class ChangeLatitudeLongitudeTypeBuildings < ActiveRecord::Migration[6.0]
  def change
    change_column :buildings, :latitude, :decimal, precision: 20, scale: 13
    change_column :buildings, :longitude, :decimal, precision: 20, scale: 13
    change_column :buildings, :distance, :decimal, precision: 20, scale: 13
  end
end
