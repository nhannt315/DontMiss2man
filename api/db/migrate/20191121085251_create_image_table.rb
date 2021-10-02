class CreateImageTable < ActiveRecord::Migration[6.0]
  def change
    create_table :images do |t|
      t.string :url
      t.string :description
      t.integer :room_id
      t.timestamps
    end
  end
end
