class AddIndexForNames < ActiveRecord::Migration[6.0]
  def change
    add_index :buildings, :name, unique: true
    add_index :agents, :name, unique: true
  end
end
