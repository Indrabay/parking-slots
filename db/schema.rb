# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2021_04_20_034847) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "customers", force: :cascade do |t|
    t.integer "slot_id"
    t.string "code"
    t.datetime "checkin_time"
    t.datetime "checkout_time"
    t.integer "price"
    t.index ["code"], name: "index_customers_on_code"
  end

  create_table "parking_slots", force: :cascade do |t|
    t.integer "price"
    t.string "name"
  end

  create_table "slots", force: :cascade do |t|
    t.integer "parking_slot_id"
    t.string "code"
    t.integer "available", limit: 2, default: 1
  end

end
