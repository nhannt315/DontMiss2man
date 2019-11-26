# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `rails
# db:schema:load`. When creating a new database, `rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2019_11_21_085303) do

  create_table "agents", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "name"
    t.string "address"
    t.string "working_time"
    t.string "telephone_number"
    t.string "email"
    t.string "photo_url"
    t.string "slogan"
    t.text "access"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "buildings", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "name"
    t.string "address"
    t.string "access"
    t.date "year_built"
    t.string "type"
    t.string "structure"
    t.integer "storeys"
    t.integer "underground_storeys"
    t.string "photo_url"
    t.float "longitude"
    t.float "latitude"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "images", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "url"
    t.string "description"
    t.integer "room_id"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "rooms", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci", force: :cascade do |t|
    t.string "suumo_id"
    t.integer "building_id"
    t.integer "agent_id"
    t.decimal "rent_fee", precision: 10
    t.decimal "reikin", precision: 10
    t.decimal "shikikin", precision: 10
    t.decimal "management_cost", precision: 10
    t.decimal "caution_fee", precision: 10
    t.string "layout_image_url"
    t.integer "size"
    t.string "direction"
    t.text "facilities"
    t.integer "floor"
    t.string "car_park"
    t.string "condition"
    t.text "note"
    t.string "layout"
    t.string "layout_detail"
    t.string "deal_type"
    t.date "move_in_time"
    t.string "move_in"
    t.string "damage_insurance"
    t.string "guarantor"
    t.string "other_fees"
    t.string "other_initial_fees"
    t.date "last_update"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

end
