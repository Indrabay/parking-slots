class CreateCustomer < ActiveRecord::Migration[5.2]
  def change
    create_table :customers do |t|
      t.integer :slot_id
      t.string :code
      t.datetime :checkin_time
      t.datetime :checkout_time
      t.integer :price

      t.index :code
    end
  end
end
