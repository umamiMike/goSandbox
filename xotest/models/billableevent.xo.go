// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// BillableEvent represents a row from 'AllergyNew.billable_events'.
type BillableEvent struct {
	ID          uint          `json:"id"`           // id
	PracticeID  uint          `json:"practice_id"`  // practice_id
	PatientID   uint          `json:"patient_id"`   // patient_id
	LocationID  uint          `json:"location_id"`  // location_id
	ServiceDate time.Time     `json:"service_date"` // service_date
	Type        Type          `json:"type"`         // type
	OwnerID     uint          `json:"owner_id"`     // owner_id
	Status      Status        `json:"status"`       // status
	UpdatedBy   sql.NullInt64 `json:"updated_by"`   // updated_by
	Updated     time.Time     `json:"updated"`      // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BillableEvent exists in the database.
func (be *BillableEvent) Exists() bool {
	return be._exists
}

// Deleted provides information if the BillableEvent has been deleted from the database.
func (be *BillableEvent) Deleted() bool {
	return be._deleted
}

// Insert inserts the BillableEvent to the database.
func (be *BillableEvent) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if be._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.billable_events (` +
		`practice_id, patient_id, location_id, service_date, type, owner_id, status, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, be.PracticeID, be.PatientID, be.LocationID, be.ServiceDate, be.Type, be.OwnerID, be.Status, be.UpdatedBy, be.Updated)
	res, err := db.Exec(sqlstr, be.PracticeID, be.PatientID, be.LocationID, be.ServiceDate, be.Type, be.OwnerID, be.Status, be.UpdatedBy, be.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	be.ID = uint(id)
	be._exists = true

	return nil
}

// Update updates the BillableEvent in the database.
func (be *BillableEvent) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !be._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if be._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.billable_events SET ` +
		`practice_id = ?, patient_id = ?, location_id = ?, service_date = ?, type = ?, owner_id = ?, status = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, be.PracticeID, be.PatientID, be.LocationID, be.ServiceDate, be.Type, be.OwnerID, be.Status, be.UpdatedBy, be.Updated, be.ID)
	_, err = db.Exec(sqlstr, be.PracticeID, be.PatientID, be.LocationID, be.ServiceDate, be.Type, be.OwnerID, be.Status, be.UpdatedBy, be.Updated, be.ID)
	return err
}

// Save saves the BillableEvent to the database.
func (be *BillableEvent) Save(db XODB) error {
	if be.Exists() {
		return be.Update(db)
	}

	return be.Insert(db)
}

// Delete deletes the BillableEvent from the database.
func (be *BillableEvent) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !be._exists {
		return nil
	}

	// if deleted, bail
	if be._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.billable_events WHERE id = ?`

	// run query
	XOLog(sqlstr, be.ID)
	_, err = db.Exec(sqlstr, be.ID)
	if err != nil {
		return err
	}

	// set deleted
	be._deleted = true

	return nil
}

// BillableEventByID retrieves a row from 'AllergyNew.billable_events' as a BillableEvent.
//
// Generated from index 'billable_events_id_pkey'.
func BillableEventByID(db XODB, id uint) (*BillableEvent, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, patient_id, location_id, service_date, type, owner_id, status, updated_by, updated ` +
		`FROM AllergyNew.billable_events ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	be := BillableEvent{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&be.ID, &be.PracticeID, &be.PatientID, &be.LocationID, &be.ServiceDate, &be.Type, &be.OwnerID, &be.Status, &be.UpdatedBy, &be.Updated)
	if err != nil {
		return nil, err
	}

	return &be, nil
}