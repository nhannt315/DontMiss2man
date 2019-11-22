class CreateAgentTable < ActiveRecord::Migration[6.0]
  def change
    create_table :agents do |t|
      t.string :name
      t.string :address
      t.string :working_time
      t.string :telephone_number
      t.string :email
    end
  end
end
