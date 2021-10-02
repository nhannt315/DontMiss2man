class CreateUserFavoriteRoom < ActiveRecord::Migration[6.0]
  def change
    create_table :user_rooms do |t|
      t.integer :user_id
      t.integer :room_id

      t.timestamps
    end
  end
end
