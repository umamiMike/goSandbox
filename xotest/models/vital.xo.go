// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Vital represents a row from 'AllergyNew.vitals'.
type Vital struct {
	ID           uint            `json:"id"`           // id
	PracticeID   sql.NullInt64   `json:"practice_id"`  // practice_id
	LocationID   sql.NullInt64   `json:"location_id"`  // location_id
	PatientID    sql.NullInt64   `json:"patient_id"`   // patient_id
	Systolic     sql.NullInt64   `json:"systolic"`     // systolic
	Diastolic    sql.NullInt64   `json:"diastolic"`    // diastolic
	Pulse        sql.NullInt64   `json:"pulse"`        // pulse
	Respirations sql.NullInt64   `json:"respirations"` // respirations
	O2           sql.NullFloat64 `json:"O2"`           // O2
	Temperature  sql.NullFloat64 `json:"temperature"`  // temperature
	Weight       sql.NullInt64   `json:"weight"`       // weight
	Height       sql.NullInt64   `json:"height"`       // height
	PeakFlow     sql.NullFloat64 `json:"peak_flow"`    // peak_flow
	Notes        sql.NullString  `json:"notes"`        // notes
	CreatorID    sql.NullInt64   `json:"creator_id"`   // creator_id
	CreateDate   mysql.NullTime  `json:"create_date"`  // create_date
	Enabled      int8            `json:"enabled"`      // enabled
	UpdatedBy    sql.NullInt64   `json:"updated_by"`   // updated_by
	Updated      time.Time       `json:"updated"`      // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Vital exists in the database.
func (v *Vital) Exists() bool {
	return v._exists
}

// Deleted provides information if the Vital has been deleted from the database.
func (v *Vital) Deleted() bool {
	return v._deleted
}

// Insert inserts the Vital to the database.
func (v *Vital) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if v._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.vitals (` +
		`practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, v.PracticeID, v.LocationID, v.PatientID, v.Systolic, v.Diastolic, v.Pulse, v.Respirations, v.O2, v.Temperature, v.Weight, v.Height, v.PeakFlow, v.Notes, v.CreatorID, v.CreateDate, v.Enabled, v.UpdatedBy, v.Updated)
	res, err := db.Exec(sqlstr, v.PracticeID, v.LocationID, v.PatientID, v.Systolic, v.Diastolic, v.Pulse, v.Respirations, v.O2, v.Temperature, v.Weight, v.Height, v.PeakFlow, v.Notes, v.CreatorID, v.CreateDate, v.Enabled, v.UpdatedBy, v.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	v.ID = uint(id)
	v._exists = true

	return nil
}

// Update updates the Vital in the database.
func (v *Vital) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if v._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.vitals SET ` +
		`practice_id = ?, location_id = ?, patient_id = ?, systolic = ?, diastolic = ?, pulse = ?, respirations = ?, O2 = ?, temperature = ?, weight = ?, height = ?, peak_flow = ?, notes = ?, creator_id = ?, create_date = ?, enabled = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, v.PracticeID, v.LocationID, v.PatientID, v.Systolic, v.Diastolic, v.Pulse, v.Respirations, v.O2, v.Temperature, v.Weight, v.Height, v.PeakFlow, v.Notes, v.CreatorID, v.CreateDate, v.Enabled, v.UpdatedBy, v.Updated, v.ID)
	_, err = db.Exec(sqlstr, v.PracticeID, v.LocationID, v.PatientID, v.Systolic, v.Diastolic, v.Pulse, v.Respirations, v.O2, v.Temperature, v.Weight, v.Height, v.PeakFlow, v.Notes, v.CreatorID, v.CreateDate, v.Enabled, v.UpdatedBy, v.Updated, v.ID)
	return err
}

// Save saves the Vital to the database.
func (v *Vital) Save(db XODB) error {
	if v.Exists() {
		return v.Update(db)
	}

	return v.Insert(db)
}

// Delete deletes the Vital from the database.
func (v *Vital) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return nil
	}

	// if deleted, bail
	if v._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.vitals WHERE id = ?`

	// run query
	XOLog(sqlstr, v.ID)
	_, err = db.Exec(sqlstr, v.ID)
	if err != nil {
		return err
	}

	// set deleted
	v._deleted = true

	return nil
}

// UserByCreatorID returns the User associated with the Vital's CreatorID (creator_id).
//
// Generated from foreign key 'vitals_creator'.
func (v *Vital) UserByCreatorID(db XODB) (*User, error) {
	return UserByID(db, uint(v.CreatorID.Int64))
}

// PracticeLocation returns the PracticeLocation associated with the Vital's LocationID (location_id).
//
// Generated from foreign key 'vitals_location'.
func (v *Vital) PracticeLocation(db XODB) (*PracticeLocation, error) {
	return PracticeLocationByID(db, uint(v.LocationID.Int64))
}

// Patient returns the Patient associated with the Vital's PatientID (patient_id).
//
// Generated from foreign key 'vitals_patients'.
func (v *Vital) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(v.PatientID.Int64))
}

// Practice returns the Practice associated with the Vital's PracticeID (practice_id).
//
// Generated from foreign key 'vitals_practice'.
func (v *Vital) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(v.PracticeID.Int64))
}

// UserByUpdatedBy returns the User associated with the Vital's UpdatedBy (updated_by).
//
// Generated from foreign key 'vitals_updated'.
func (v *Vital) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(v.UpdatedBy.Int64))
}

// VitalsByCreatorID retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'Creators'.
func VitalsByCreatorID(db XODB, creatorID sql.NullInt64) ([]*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE creator_id = ?`

	// run query
	XOLog(sqlstr, creatorID)
	q, err := db.Query(sqlstr, creatorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vital{}
	for q.Next() {
		v := Vital{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VitalsByLocationID retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'Locations'.
func VitalsByLocationID(db XODB, locationID sql.NullInt64) ([]*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE location_id = ?`

	// run query
	XOLog(sqlstr, locationID)
	q, err := db.Query(sqlstr, locationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vital{}
	for q.Next() {
		v := Vital{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VitalsByPatientID retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'Patients'.
func VitalsByPatientID(db XODB, patientID sql.NullInt64) ([]*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vital{}
	for q.Next() {
		v := Vital{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VitalsByPracticeID retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'Practices'.
func VitalsByPracticeID(db XODB, practiceID sql.NullInt64) ([]*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vital{}
	for q.Next() {
		v := Vital{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VitalByID retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'vitals_id_pkey'.
func VitalByID(db XODB, id uint) (*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	v := Vital{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// VitalsByUpdatedBy retrieves a row from 'AllergyNew.vitals' as a Vital.
//
// Generated from index 'vitals_updated'.
func VitalsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*Vital, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, location_id, patient_id, systolic, diastolic, pulse, respirations, O2, temperature, weight, height, peak_flow, notes, creator_id, create_date, enabled, updated_by, updated ` +
		`FROM AllergyNew.vitals ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vital{}
	for q.Next() {
		v := Vital{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.PracticeID, &v.LocationID, &v.PatientID, &v.Systolic, &v.Diastolic, &v.Pulse, &v.Respirations, &v.O2, &v.Temperature, &v.Weight, &v.Height, &v.PeakFlow, &v.Notes, &v.CreatorID, &v.CreateDate, &v.Enabled, &v.UpdatedBy, &v.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}