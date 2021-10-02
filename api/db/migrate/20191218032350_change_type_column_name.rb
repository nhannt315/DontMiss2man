class ChangeTypeColumnName < ActiveRecord::Migration[6.0]
  def change
    rename_column :buildings, :type, :building_type
  end
end
