// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// PatientPicture represents a row from 'AllergyNew.patient_pictures'.
type PatientPicture struct {
	ID         uint           `json:"id"`          // id
	PatientID  sql.NullInt64  `json:"patient_id"`  // patient_id
	Picture    []byte         `json:"picture"`     // picture
	UploadDate mysql.NullTime `json:"upload_date"` // upload_date
	UpdatedBy  sql.NullInt64  `json:"updated_by"`  // updated_by
	Updated    time.Time      `json:"updated"`     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PatientPicture exists in the database.
func (pp *PatientPicture) Exists() bool {
	return pp._exists
}

// Deleted provides information if the PatientPicture has been deleted from the database.
func (pp *PatientPicture) Deleted() bool {
	return pp._deleted
}

// Insert inserts the PatientPicture to the database.
func (pp *PatientPicture) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.patient_pictures (` +
		`patient_id, picture, upload_date, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pp.PatientID, pp.Picture, pp.UploadDate, pp.UpdatedBy, pp.Updated)
	res, err := db.Exec(sqlstr, pp.PatientID, pp.Picture, pp.UploadDate, pp.UpdatedBy, pp.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pp.ID = uint(id)
	pp._exists = true

	return nil
}

// Update updates the PatientPicture in the database.
func (pp *PatientPicture) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.patient_pictures SET ` +
		`patient_id = ?, picture = ?, upload_date = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, pp.PatientID, pp.Picture, pp.UploadDate, pp.UpdatedBy, pp.Updated, pp.ID)
	_, err = db.Exec(sqlstr, pp.PatientID, pp.Picture, pp.UploadDate, pp.UpdatedBy, pp.Updated, pp.ID)
	return err
}

// Save saves the PatientPicture to the database.
func (pp *PatientPicture) Save(db XODB) error {
	if pp.Exists() {
		return pp.Update(db)
	}

	return pp.Insert(db)
}

// Delete deletes the PatientPicture from the database.
func (pp *PatientPicture) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pp._exists {
		return nil
	}

	// if deleted, bail
	if pp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.patient_pictures WHERE id = ?`

	// run query
	XOLog(sqlstr, pp.ID)
	_, err = db.Exec(sqlstr, pp.ID)
	if err != nil {
		return err
	}

	// set deleted
	pp._deleted = true

	return nil
}

// Patient returns the Patient associated with the PatientPicture's PatientID (patient_id).
//
// Generated from foreign key 'patient_pictures_patient'.
func (pp *PatientPicture) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(pp.PatientID.Int64))
}

// User returns the User associated with the PatientPicture's UpdatedBy (updated_by).
//
// Generated from foreign key 'patient_pictures_updated'.
func (pp *PatientPicture) User(db XODB) (*User, error) {
	return UserByID(db, uint(pp.UpdatedBy.Int64))
}

// PatientPicturesByPatientID retrieves a row from 'AllergyNew.patient_pictures' as a PatientPicture.
//
// Generated from index 'Patients'.
func PatientPicturesByPatientID(db XODB, patientID sql.NullInt64) ([]*PatientPicture, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, picture, upload_date, updated_by, updated ` +
		`FROM AllergyNew.patient_pictures ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientPicture{}
	for q.Next() {
		pp := PatientPicture{
			_exists: true,
		}

		// scan
		err = q.Scan(&pp.ID, &pp.PatientID, &pp.Picture, &pp.UploadDate, &pp.UpdatedBy, &pp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pp)
	}

	return res, nil
}

// PatientPictureByID retrieves a row from 'AllergyNew.patient_pictures' as a PatientPicture.
//
// Generated from index 'patient_pictures_id_pkey'.
func PatientPictureByID(db XODB, id uint) (*PatientPicture, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, picture, upload_date, updated_by, updated ` +
		`FROM AllergyNew.patient_pictures ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	pp := PatientPicture{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pp.ID, &pp.PatientID, &pp.Picture, &pp.UploadDate, &pp.UpdatedBy, &pp.Updated)
	if err != nil {
		return nil, err
	}

	return &pp, nil
}

// PatientPicturesByUpdatedBy retrieves a row from 'AllergyNew.patient_pictures' as a PatientPicture.
//
// Generated from index 'patient_pictures_updated'.
func PatientPicturesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*PatientPicture, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, picture, upload_date, updated_by, updated ` +
		`FROM AllergyNew.patient_pictures ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientPicture{}
	for q.Next() {
		pp := PatientPicture{
			_exists: true,
		}

		// scan
		err = q.Scan(&pp.ID, &pp.PatientID, &pp.Picture, &pp.UploadDate, &pp.UpdatedBy, &pp.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pp)
	}

	return res, nil
}