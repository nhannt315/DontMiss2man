class AddAverageSizePrice < ActiveRecord::Migration[6.0]
  def change
    add_column :buildings, :average_size, :float
    add_column :buildings, :average_fee, :float
  end
end
