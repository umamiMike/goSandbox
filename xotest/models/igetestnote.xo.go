// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// IgeTestNote represents a row from 'AllergyNew.ige_test_notes'.
type IgeTestNote struct {
	ID        uint           `json:"id"`         // id
	TestID    sql.NullInt64  `json:"test_id"`    // test_id
	PatientID sql.NullInt64  `json:"patient_id"` // patient_id
	Note      sql.NullString `json:"note"`       // note
	DateAdded mysql.NullTime `json:"date_added"` // date_added
	AddedBy   sql.NullInt64  `json:"added_by"`   // added_by
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the IgeTestNote exists in the database.
func (itn *IgeTestNote) Exists() bool {
	return itn._exists
}

// Deleted provides information if the IgeTestNote has been deleted from the database.
func (itn *IgeTestNote) Deleted() bool {
	return itn._deleted
}

// Insert inserts the IgeTestNote to the database.
func (itn *IgeTestNote) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if itn._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.ige_test_notes (` +
		`test_id, patient_id, note, date_added, added_by, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, itn.TestID, itn.PatientID, itn.Note, itn.DateAdded, itn.AddedBy, itn.UpdatedBy, itn.Updated)
	res, err := db.Exec(sqlstr, itn.TestID, itn.PatientID, itn.Note, itn.DateAdded, itn.AddedBy, itn.UpdatedBy, itn.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	itn.ID = uint(id)
	itn._exists = true

	return nil
}

// Update updates the IgeTestNote in the database.
func (itn *IgeTestNote) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itn._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if itn._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.ige_test_notes SET ` +
		`test_id = ?, patient_id = ?, note = ?, date_added = ?, added_by = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, itn.TestID, itn.PatientID, itn.Note, itn.DateAdded, itn.AddedBy, itn.UpdatedBy, itn.Updated, itn.ID)
	_, err = db.Exec(sqlstr, itn.TestID, itn.PatientID, itn.Note, itn.DateAdded, itn.AddedBy, itn.UpdatedBy, itn.Updated, itn.ID)
	return err
}

// Save saves the IgeTestNote to the database.
func (itn *IgeTestNote) Save(db XODB) error {
	if itn.Exists() {
		return itn.Update(db)
	}

	return itn.Insert(db)
}

// Delete deletes the IgeTestNote from the database.
func (itn *IgeTestNote) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itn._exists {
		return nil
	}

	// if deleted, bail
	if itn._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.ige_test_notes WHERE id = ?`

	// run query
	XOLog(sqlstr, itn.ID)
	_, err = db.Exec(sqlstr, itn.ID)
	if err != nil {
		return err
	}

	// set deleted
	itn._deleted = true

	return nil
}

// UserByAddedBy returns the User associated with the IgeTestNote's AddedBy (added_by).
//
// Generated from foreign key 'ige_test_notes_added_by'.
func (itn *IgeTestNote) UserByAddedBy(db XODB) (*User, error) {
	return UserByID(db, uint(itn.AddedBy.Int64))
}

// Patient returns the Patient associated with the IgeTestNote's PatientID (patient_id).
//
// Generated from foreign key 'ige_test_notes_patient'.
func (itn *IgeTestNote) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(itn.PatientID.Int64))
}

// IgeTest returns the IgeTest associated with the IgeTestNote's TestID (test_id).
//
// Generated from foreign key 'ige_test_notes_test'.
func (itn *IgeTestNote) IgeTest(db XODB) (*IgeTest, error) {
	return IgeTestByID(db, uint(itn.TestID.Int64))
}

// UserByUpdatedBy returns the User associated with the IgeTestNote's UpdatedBy (updated_by).
//
// Generated from foreign key 'ige_test_notes_updated'.
func (itn *IgeTestNote) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(itn.UpdatedBy.Int64))
}

// IgeTestNotesByTestID retrieves a row from 'AllergyNew.ige_test_notes' as a IgeTestNote.
//
// Generated from index 'IgE_Tests'.
func IgeTestNotesByTestID(db XODB, testID sql.NullInt64) ([]*IgeTestNote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, patient_id, note, date_added, added_by, updated_by, updated ` +
		`FROM AllergyNew.ige_test_notes ` +
		`WHERE test_id = ?`

	// run query
	XOLog(sqlstr, testID)
	q, err := db.Query(sqlstr, testID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IgeTestNote{}
	for q.Next() {
		itn := IgeTestNote{
			_exists: true,
		}

		// scan
		err = q.Scan(&itn.ID, &itn.TestID, &itn.PatientID, &itn.Note, &itn.DateAdded, &itn.AddedBy, &itn.UpdatedBy, &itn.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itn)
	}

	return res, nil
}

// IgeTestNotesByPatientID retrieves a row from 'AllergyNew.ige_test_notes' as a IgeTestNote.
//
// Generated from index 'Patients'.
func IgeTestNotesByPatientID(db XODB, patientID sql.NullInt64) ([]*IgeTestNote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, patient_id, note, date_added, added_by, updated_by, updated ` +
		`FROM AllergyNew.ige_test_notes ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IgeTestNote{}
	for q.Next() {
		itn := IgeTestNote{
			_exists: true,
		}

		// scan
		err = q.Scan(&itn.ID, &itn.TestID, &itn.PatientID, &itn.Note, &itn.DateAdded, &itn.AddedBy, &itn.UpdatedBy, &itn.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itn)
	}

	return res, nil
}

// IgeTestNotesByAddedBy retrieves a row from 'AllergyNew.ige_test_notes' as a IgeTestNote.
//
// Generated from index 'ige_test_notes_added_by'.
func IgeTestNotesByAddedBy(db XODB, addedBy sql.NullInt64) ([]*IgeTestNote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, patient_id, note, date_added, added_by, updated_by, updated ` +
		`FROM AllergyNew.ige_test_notes ` +
		`WHERE added_by = ?`

	// run query
	XOLog(sqlstr, addedBy)
	q, err := db.Query(sqlstr, addedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IgeTestNote{}
	for q.Next() {
		itn := IgeTestNote{
			_exists: true,
		}

		// scan
		err = q.Scan(&itn.ID, &itn.TestID, &itn.PatientID, &itn.Note, &itn.DateAdded, &itn.AddedBy, &itn.UpdatedBy, &itn.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itn)
	}

	return res, nil
}

// IgeTestNoteByID retrieves a row from 'AllergyNew.ige_test_notes' as a IgeTestNote.
//
// Generated from index 'ige_test_notes_id_pkey'.
func IgeTestNoteByID(db XODB, id uint) (*IgeTestNote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, patient_id, note, date_added, added_by, updated_by, updated ` +
		`FROM AllergyNew.ige_test_notes ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	itn := IgeTestNote{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&itn.ID, &itn.TestID, &itn.PatientID, &itn.Note, &itn.DateAdded, &itn.AddedBy, &itn.UpdatedBy, &itn.Updated)
	if err != nil {
		return nil, err
	}

	return &itn, nil
}

// IgeTestNotesByUpdatedBy retrieves a row from 'AllergyNew.ige_test_notes' as a IgeTestNote.
//
// Generated from index 'ige_test_notes_updated'.
func IgeTestNotesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*IgeTestNote, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, patient_id, note, date_added, added_by, updated_by, updated ` +
		`FROM AllergyNew.ige_test_notes ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IgeTestNote{}
	for q.Next() {
		itn := IgeTestNote{
			_exists: true,
		}

		// scan
		err = q.Scan(&itn.ID, &itn.TestID, &itn.PatientID, &itn.Note, &itn.DateAdded, &itn.AddedBy, &itn.UpdatedBy, &itn.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itn)
	}

	return res, nil
}