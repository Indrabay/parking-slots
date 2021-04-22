class CreateSlot < ActiveRecord::Migration[5.2]
  def change
    create_table :slots do |t|
      t.integer :parking_slot_id
      t.string :code
      t.integer :available, limit: 1, default: 1
    end
  end
end
