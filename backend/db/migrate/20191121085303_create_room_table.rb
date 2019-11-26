class CreateRoomTable < ActiveRecord::Migration[6.0]
  def change
    create_table :rooms do |t|
      t.string :suumo_id
      t.integer :building_id
      t.integer :agent_id
      t.decimal :rent_fee
      t.decimal :reikin
      t.decimal :shikikin
      t.decimal :management_cost
      t.decimal :caution_fee
      t.string :layout_image_url
      t.integer :size
      t.string :direction
      t.text :facilities
      t.integer :floor
      t.string :car_park
      t.string :condition
      t.text :note
      t.string :layout
      t.string :layout_detail
      t.string :deal_type
      t.date :move_in_time
      t.string :move_in
      t.string :damage_insurance
      t.string :guarantor
      t.string :other_fees
      t.string :other_initial_fees
      t.date :last_update
      t.timestamps
    end
  end
end
