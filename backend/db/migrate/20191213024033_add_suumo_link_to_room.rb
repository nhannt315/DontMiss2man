class AddSuumoLinkToRoom < ActiveRecord::Migration[6.0]
  def change
    add_column :rooms, :suumo_link, :string
  end
end
