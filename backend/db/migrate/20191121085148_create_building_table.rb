class CreateBuildingTable < ActiveRecord::Migration[6.0]
  def change
    create_table :buildings do |t|
      t.string :name
      t.string :address
      t.string :access
      t.date :year_built
      t.string :type
      t.string :structure
      t.string :storeys
      t.string :photo_url
      t.float :longitude
      t.float :latitude
    end
  end
end
