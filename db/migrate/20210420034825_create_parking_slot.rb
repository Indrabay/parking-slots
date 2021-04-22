class CreateParkingSlot < ActiveRecord::Migration[5.2]
  def change
    create_table :parking_slots do |t|
      t.integer :price
      t.string :name
    end
  end
end
