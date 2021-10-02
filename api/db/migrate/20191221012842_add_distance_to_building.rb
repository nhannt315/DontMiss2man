class AddDistanceToBuilding < ActiveRecord::Migration[6.0]
  def change
    add_column :buildings, :distance, :float
    add_column :buildings, :condition_type, :integer, default: 0
  end
end
