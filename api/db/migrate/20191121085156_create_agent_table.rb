class CreateAgentTable < ActiveRecord::Migration[6.0]
  def change
    create_table :agents do |t|
      t.string :name
      t.string :address
      t.string :working_time
      t.string :telephone_number
      t.string :email
      t.string :photo_url
      t.string :slogan
      t.text :access
      t.timestamps
    end
  end
end
