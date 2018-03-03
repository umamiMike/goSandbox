// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// CardsLocallyPrinted represents a row from 'AllergyNew.cards_locally_printed'.
type CardsLocallyPrinted struct {
	ID         uint           `json:"id"`          // id
	PracticeID sql.NullInt64  `json:"practice_id"` // practice_id
	LocationID sql.NullInt64  `json:"location_id"` // location_id
	PatientID  sql.NullInt64  `json:"patient_id"`  // patient_id
	PrintDate  mysql.NullTime `json:"print_date"`  // print_date
	PrintedID  sql.NullInt64  `json:"printed_id"`  // printed_id
	UpdatedBy  sql.NullInt64  `json:"updated_by"`  // updated_by
	Updated    time.Time      `json:"updated"`     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the CardsLocallyPrinted exists in the database.
func (clp *CardsLocallyPrinted) Exists() bool {
	return clp._exists
}

// Deleted provides information if the CardsLocallyPrinted has been deleted from the database.
func (clp *CardsLocallyPrinted) Deleted() bool {
	return clp._deleted
}

// Insert inserts the CardsLocallyPrinted to the database.
func (clp *CardsLocallyPrinted) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if clp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.cards_locally_printed (` +
		`practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, clp.PracticeID, clp.LocationID, clp.PatientID, clp.PrintDate, clp.PrintedID, clp.UpdatedBy, clp.Updated)
	res, err := db.Exec(sqlstr, clp.PracticeID, clp.LocationID, clp.PatientID, clp.PrintDate, clp.PrintedID, clp.UpdatedBy, clp.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	clp.ID = uint(id)
	clp._exists = true

	return nil
}

// Update updates the CardsLocallyPrinted in the database.
func (clp *CardsLocallyPrinted) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !clp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if clp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.cards_locally_printed SET ` +
		`practice_id = ?, location_id = ?, patient_id = ?, print_date = ?, printed_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, clp.PracticeID, clp.LocationID, clp.PatientID, clp.PrintDate, clp.PrintedID, clp.UpdatedBy, clp.Updated, clp.ID)
	_, err = db.Exec(sqlstr, clp.PracticeID, clp.LocationID, clp.PatientID, clp.PrintDate, clp.PrintedID, clp.UpdatedBy, clp.Updated, clp.ID)
	return err
}

// Save saves the CardsLocallyPrinted to the database.
func (clp *CardsLocallyPrinted) Save(db XODB) error {
	if clp.Exists() {
		return clp.Update(db)
	}

	return clp.Insert(db)
}

// Delete deletes the CardsLocallyPrinted from the database.
func (clp *CardsLocallyPrinted) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !clp._exists {
		return nil
	}

	// if deleted, bail
	if clp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.cards_locally_printed WHERE id = ?`

	// run query
	XOLog(sqlstr, clp.ID)
	_, err = db.Exec(sqlstr, clp.ID)
	if err != nil {
		return err
	}

	// set deleted
	clp._deleted = true

	return nil
}

// PracticeLocation returns the PracticeLocation associated with the CardsLocallyPrinted's LocationID (location_id).
//
// Generated from foreign key 'cards_locally_printed_location'.
func (clp *CardsLocallyPrinted) PracticeLocation(db XODB) (*PracticeLocation, error) {
	return PracticeLocationByID(db, uint(clp.LocationID.Int64))
}

// Patient returns the Patient associated with the CardsLocallyPrinted's PatientID (patient_id).
//
// Generated from foreign key 'cards_locally_printed_patient'.
func (clp *CardsLocallyPrinted) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(clp.PatientID.Int64))
}

// Practice returns the Practice associated with the CardsLocallyPrinted's PracticeID (practice_id).
//
// Generated from foreign key 'cards_locally_printed_practice'.
func (clp *CardsLocallyPrinted) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(clp.PracticeID.Int64))
}

// UserByPrintedID returns the User associated with the CardsLocallyPrinted's PrintedID (printed_id).
//
// Generated from foreign key 'cards_locally_printed_printed'.
func (clp *CardsLocallyPrinted) UserByPrintedID(db XODB) (*User, error) {
	return UserByID(db, uint(clp.PrintedID.Int64))
}

// UserByUpdatedBy returns the User associated with the CardsLocallyPrinted's UpdatedBy (updated_by).
//
// Generated from foreign key 'cards_locally_printed_updated'.
func (clp *CardsLocallyPrinted) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(clp.UpdatedBy.Int64))
}

// CardsLocallyPrintedsByLocationID retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'Locations'.
func CardsLocallyPrintedsByLocationID(db XODB, locationID sql.NullInt64) ([]*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE location_id = ?`

	// run query
	XOLog(sqlstr, locationID)
	q, err := db.Query(sqlstr, locationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CardsLocallyPrinted{}
	for q.Next() {
		clp := CardsLocallyPrinted{
			_exists: true,
		}

		// scan
		err = q.Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &clp)
	}

	return res, nil
}

// CardsLocallyPrintedsByPatientID retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'Patients'.
func CardsLocallyPrintedsByPatientID(db XODB, patientID sql.NullInt64) ([]*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CardsLocallyPrinted{}
	for q.Next() {
		clp := CardsLocallyPrinted{
			_exists: true,
		}

		// scan
		err = q.Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &clp)
	}

	return res, nil
}

// CardsLocallyPrintedsByPracticeID retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'Practices'.
func CardsLocallyPrintedsByPracticeID(db XODB, practiceID sql.NullInt64) ([]*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CardsLocallyPrinted{}
	for q.Next() {
		clp := CardsLocallyPrinted{
			_exists: true,
		}

		// scan
		err = q.Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &clp)
	}

	return res, nil
}

// CardsLocallyPrintedByID retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'cards_locally_printed_id_pkey'.
func CardsLocallyPrintedByID(db XODB, id uint) (*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	clp := CardsLocallyPrinted{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
	if err != nil {
		return nil, err
	}

	return &clp, nil
}

// CardsLocallyPrintedsByPrintedID retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'cards_locally_printed_printed'.
func CardsLocallyPrintedsByPrintedID(db XODB, printedID sql.NullInt64) ([]*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE printed_id = ?`

	// run query
	XOLog(sqlstr, printedID)
	q, err := db.Query(sqlstr, printedID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CardsLocallyPrinted{}
	for q.Next() {
		clp := CardsLocallyPrinted{
			_exists: true,
		}

		// scan
		err = q.Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &clp)
	}

	return res, nil
}

// CardsLocallyPrintedsByUpdatedBy retrieves a row from 'AllergyNew.cards_locally_printed' as a CardsLocallyPrinted.
//
// Generated from index 'cards_locally_printed_updated'.
func CardsLocallyPrintedsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*CardsLocallyPrinted, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, print_date, printed_id, updated_by, updated ` +
		`FROM AllergyNew.cards_locally_printed ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*CardsLocallyPrinted{}
	for q.Next() {
		clp := CardsLocallyPrinted{
			_exists: true,
		}

		// scan
		err = q.Scan(&clp.ID, &clp.PracticeID, &clp.LocationID, &clp.PatientID, &clp.PrintDate, &clp.PrintedID, &clp.UpdatedBy, &clp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &clp)
	}

	return res, nil
}
