// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// PracticeLocation represents a row from 'AllergyNew.practice_locations'.
type PracticeLocation struct {
	ID                  uint           `json:"id"`                   // id
	PracticeID          sql.NullInt64  `json:"practice_id"`          // practice_id
	Name                sql.NullString `json:"name"`                 // name
	Address             sql.NullString `json:"address"`              // address
	Address2            sql.NullString `json:"address2"`             // address2
	City                sql.NullString `json:"city"`                 // city
	State               sql.NullString `json:"state"`                // state
	Zip                 sql.NullString `json:"zip"`                  // zip
	TimezoneOffset      sql.NullInt64  `json:"timezone_offset"`      // timezone_offset
	Dst                 sql.NullInt64  `json:"DST"`                  // DST
	Fastpass            sql.NullInt64  `json:"fastpass"`             // fastpass
	CustomOrder         sql.NullInt64  `json:"custom_order"`         // custom_order
	Notes               sql.NullString `json:"notes"`                // notes
	Active              int8           `json:"active"`               // active
	SupervisingProvider sql.NullInt64  `json:"supervising_provider"` // supervising_provider
	EmrReferenceID      sql.NullString `json:"emr_reference_id"`     // emr_reference_id
	UpdatedBy           sql.NullInt64  `json:"updated_by"`           // updated_by
	Updated             time.Time      `json:"updated"`              // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PracticeLocation exists in the database.
func (pl *PracticeLocation) Exists() bool {
	return pl._exists
}

// Deleted provides information if the PracticeLocation has been deleted from the database.
func (pl *PracticeLocation) Deleted() bool {
	return pl._deleted
}

// Insert inserts the PracticeLocation to the database.
func (pl *PracticeLocation) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.practice_locations (` +
		`practice_id, name, address, address2, city, state, zip, timezone_offset, DST, fastpass, custom_order, notes, active, supervising_provider, emr_reference_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pl.PracticeID, pl.Name, pl.Address, pl.Address2, pl.City, pl.State, pl.Zip, pl.TimezoneOffset, pl.Dst, pl.Fastpass, pl.CustomOrder, pl.Notes, pl.Active, pl.SupervisingProvider, pl.EmrReferenceID, pl.UpdatedBy, pl.Updated)
	res, err := db.Exec(sqlstr, pl.PracticeID, pl.Name, pl.Address, pl.Address2, pl.City, pl.State, pl.Zip, pl.TimezoneOffset, pl.Dst, pl.Fastpass, pl.CustomOrder, pl.Notes, pl.Active, pl.SupervisingProvider, pl.EmrReferenceID, pl.UpdatedBy, pl.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pl.ID = uint(id)
	pl._exists = true

	return nil
}

// Update updates the PracticeLocation in the database.
func (pl *PracticeLocation) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pl._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.practice_locations SET ` +
		`practice_id = ?, name = ?, address = ?, address2 = ?, city = ?, state = ?, zip = ?, timezone_offset = ?, DST = ?, fastpass = ?, custom_order = ?, notes = ?, active = ?, supervising_provider = ?, emr_reference_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, pl.PracticeID, pl.Name, pl.Address, pl.Address2, pl.City, pl.State, pl.Zip, pl.TimezoneOffset, pl.Dst, pl.Fastpass, pl.CustomOrder, pl.Notes, pl.Active, pl.SupervisingProvider, pl.EmrReferenceID, pl.UpdatedBy, pl.Updated, pl.ID)
	_, err = db.Exec(sqlstr, pl.PracticeID, pl.Name, pl.Address, pl.Address2, pl.City, pl.State, pl.Zip, pl.TimezoneOffset, pl.Dst, pl.Fastpass, pl.CustomOrder, pl.Notes, pl.Active, pl.SupervisingProvider, pl.EmrReferenceID, pl.UpdatedBy, pl.Updated, pl.ID)
	return err
}

// Save saves the PracticeLocation to the database.
func (pl *PracticeLocation) Save(db XODB) error {
	if pl.Exists() {
		return pl.Update(db)
	}

	return pl.Insert(db)
}

// Delete deletes the PracticeLocation from the database.
func (pl *PracticeLocation) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pl._exists {
		return nil
	}

	// if deleted, bail
	if pl._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.practice_locations WHERE id = ?`

	// run query
	XOLog(sqlstr, pl.ID)
	_, err = db.Exec(sqlstr, pl.ID)
	if err != nil {
		return err
	}

	// set deleted
	pl._deleted = true

	return nil
}

// Practice returns the Practice associated with the PracticeLocation's PracticeID (practice_id).
//
// Generated from foreign key 'practice_locations_practice'.
func (pl *PracticeLocation) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(pl.PracticeID.Int64))
}

// User returns the User associated with the PracticeLocation's UpdatedBy (updated_by).
//
// Generated from foreign key 'practice_locations_updated'.
func (pl *PracticeLocation) User(db XODB) (*User, error) {
	return UserByID(db, uint(pl.UpdatedBy.Int64))
}

// PracticeLocationsByPracticeID retrieves a row from 'AllergyNew.practice_locations' as a PracticeLocation.
//
// Generated from index 'Practices'.
func PracticeLocationsByPracticeID(db XODB, practiceID sql.NullInt64) ([]*PracticeLocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, address, address2, city, state, zip, timezone_offset, DST, fastpass, custom_order, notes, active, supervising_provider, emr_reference_id, updated_by, updated ` +
		`FROM AllergyNew.practice_locations ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PracticeLocation{}
	for q.Next() {
		pl := PracticeLocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&pl.ID, &pl.PracticeID, &pl.Name, &pl.Address, &pl.Address2, &pl.City, &pl.State, &pl.Zip, &pl.TimezoneOffset, &pl.Dst, &pl.Fastpass, &pl.CustomOrder, &pl.Notes, &pl.Active, &pl.SupervisingProvider, &pl.EmrReferenceID, &pl.UpdatedBy, &pl.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pl)
	}

	return res, nil
}

// PracticeLocationByID retrieves a row from 'AllergyNew.practice_locations' as a PracticeLocation.
//
// Generated from index 'practice_locations_id_pkey'.
func PracticeLocationByID(db XODB, id uint) (*PracticeLocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, address, address2, city, state, zip, timezone_offset, DST, fastpass, custom_order, notes, active, supervising_provider, emr_reference_id, updated_by, updated ` +
		`FROM AllergyNew.practice_locations ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	pl := PracticeLocation{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pl.ID, &pl.PracticeID, &pl.Name, &pl.Address, &pl.Address2, &pl.City, &pl.State, &pl.Zip, &pl.TimezoneOffset, &pl.Dst, &pl.Fastpass, &pl.CustomOrder, &pl.Notes, &pl.Active, &pl.SupervisingProvider, &pl.EmrReferenceID, &pl.UpdatedBy, &pl.Updated)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}

// PracticeLocationsByUpdatedBy retrieves a row from 'AllergyNew.practice_locations' as a PracticeLocation.
//
// Generated from index 'practice_locations_updated'.
func PracticeLocationsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*PracticeLocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, address, address2, city, state, zip, timezone_offset, DST, fastpass, custom_order, notes, active, supervising_provider, emr_reference_id, updated_by, updated ` +
		`FROM AllergyNew.practice_locations ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PracticeLocation{}
	for q.Next() {
		pl := PracticeLocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&pl.ID, &pl.PracticeID, &pl.Name, &pl.Address, &pl.Address2, &pl.City, &pl.State, &pl.Zip, &pl.TimezoneOffset, &pl.Dst, &pl.Fastpass, &pl.CustomOrder, &pl.Notes, &pl.Active, &pl.SupervisingProvider, &pl.EmrReferenceID, &pl.UpdatedBy, &pl.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pl)
	}

	return res, nil
}
