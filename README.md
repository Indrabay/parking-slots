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

customer | slots | parking_slots
------------- | ----- | --------
slot_id            | parking_slot_id | id
code         | code | price
checkin_time          | available | name
checkout_time
total_price

---
## Prerequisite
- go
- postgres
- etc

## Setup

- step 1
- step 2
- step 3
- step 4