# Parking Slots

This is exercise project to help me build skeleton for other project based on go language.

---
## List Endpoint

* POST /parking-slots
* PATCH /parking-slots/{id}
* GET /parking-slots
* GET /parking-slots/{id}
* POST /parking-slots/{id}/slots
* GET /parking-slots/{id}/slots
* PATCH /parking-slots/{parking_id}/slots/{slot_id}/check-in
* PATCH /parking-slots/{parking_id}/slots/{slot_id}/check-out

---
## Database Schema

customers     | slots           | parking_slots
------------- | --------------- | -------------
slot_id       | parking_slot_id | id
code          | code            | price
checkin_time  | available       | name
checkout_time
total_price

---
## Prerequisite
- go 1.13.5
- postgres 13
- ruby 2.5.3

## Setup

- clone this repo `git clone git@github.com:Indrabay/parking-slots.git`
- `cp env.sample .env`
- `gem install bundle`
- `bundle install`
- setup database `rake db:create & rake db:migrate`
- run this project `go run cmd/main.go`