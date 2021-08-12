class AddOfficeIdBuilding < ActiveRecord::Migration[6.0]
  def change
    add_column :buildings, :office_id, :integer
  end
end
