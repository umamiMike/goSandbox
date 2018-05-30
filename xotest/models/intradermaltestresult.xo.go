// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// IntradermalTestResult represents a row from 'AllergyNew.intradermal_test_results'.
type IntradermalTestResult struct {
	ID            uint           `json:"id"`             // id
	TestID        sql.NullInt64  `json:"test_id"`        // test_id
	AntigenID     sql.NullInt64  `json:"antigen_id"`     // antigen_id
	PatientID     sql.NullInt64  `json:"patient_id"`     // patient_id
	Dilution      sql.NullInt64  `json:"dilution"`       // dilution
	Presumed      int8           `json:"presumed"`       // presumed
	FlarePresumed sql.NullInt64  `json:"flare_presumed"` // flare_presumed
	WhealSize     sql.NullString `json:"wheal_size"`     // wheal_size
	FlareSize     sql.NullString `json:"flare_size"`     // flare_size
	Enabled       int8           `json:"enabled"`        // enabled
	UpdatedBy     sql.NullInt64  `json:"updated_by"`     // updated_by
	Updated       time.Time      `json:"updated"`        // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the IntradermalTestResult exists in the database.
func (itr *IntradermalTestResult) Exists() bool {
	return itr._exists
}

// Deleted provides information if the IntradermalTestResult has been deleted from the database.
func (itr *IntradermalTestResult) Deleted() bool {
	return itr._deleted
}

// Insert inserts the IntradermalTestResult to the database.
func (itr *IntradermalTestResult) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if itr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.intradermal_test_results (` +
		`test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, itr.TestID, itr.AntigenID, itr.PatientID, itr.Dilution, itr.Presumed, itr.FlarePresumed, itr.WhealSize, itr.FlareSize, itr.Enabled, itr.UpdatedBy, itr.Updated)
	res, err := db.Exec(sqlstr, itr.TestID, itr.AntigenID, itr.PatientID, itr.Dilution, itr.Presumed, itr.FlarePresumed, itr.WhealSize, itr.FlareSize, itr.Enabled, itr.UpdatedBy, itr.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	itr.ID = uint(id)
	itr._exists = true

	return nil
}

// Update updates the IntradermalTestResult in the database.
func (itr *IntradermalTestResult) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if itr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.intradermal_test_results SET ` +
		`test_id = ?, antigen_id = ?, patient_id = ?, dilution = ?, presumed = ?, flare_presumed = ?, wheal_size = ?, flare_size = ?, enabled = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, itr.TestID, itr.AntigenID, itr.PatientID, itr.Dilution, itr.Presumed, itr.FlarePresumed, itr.WhealSize, itr.FlareSize, itr.Enabled, itr.UpdatedBy, itr.Updated, itr.ID)
	_, err = db.Exec(sqlstr, itr.TestID, itr.AntigenID, itr.PatientID, itr.Dilution, itr.Presumed, itr.FlarePresumed, itr.WhealSize, itr.FlareSize, itr.Enabled, itr.UpdatedBy, itr.Updated, itr.ID)
	return err
}

// Save saves the IntradermalTestResult to the database.
func (itr *IntradermalTestResult) Save(db XODB) error {
	if itr.Exists() {
		return itr.Update(db)
	}

	return itr.Insert(db)
}

// Delete deletes the IntradermalTestResult from the database.
func (itr *IntradermalTestResult) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itr._exists {
		return nil
	}

	// if deleted, bail
	if itr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.intradermal_test_results WHERE id = ?`

	// run query
	XOLog(sqlstr, itr.ID)
	_, err = db.Exec(sqlstr, itr.ID)
	if err != nil {
		return err
	}

	// set deleted
	itr._deleted = true

	return nil
}

// IntradermalTestAntigen returns the IntradermalTestAntigen associated with the IntradermalTestResult's AntigenID (antigen_id).
//
// Generated from foreign key 'intradermal_test_results_antigen'.
func (itr *IntradermalTestResult) IntradermalTestAntigen(db XODB) (*IntradermalTestAntigen, error) {
	return IntradermalTestAntigenByID(db, uint(itr.AntigenID.Int64))
}

// Patient returns the Patient associated with the IntradermalTestResult's PatientID (patient_id).
//
// Generated from foreign key 'intradermal_test_results_patient'.
func (itr *IntradermalTestResult) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(itr.PatientID.Int64))
}

// IntradermalTest returns the IntradermalTest associated with the IntradermalTestResult's TestID (test_id).
//
// Generated from foreign key 'intradermal_test_results_tests'.
func (itr *IntradermalTestResult) IntradermalTest(db XODB) (*IntradermalTest, error) {
	return IntradermalTestByID(db, uint(itr.TestID.Int64))
}

// User returns the User associated with the IntradermalTestResult's UpdatedBy (updated_by).
//
// Generated from foreign key 'intradermal_test_results_updated'.
func (itr *IntradermalTestResult) User(db XODB) (*User, error) {
	return UserByID(db, uint(itr.UpdatedBy.Int64))
}

// IntradermalTestResultsByAntigenID retrieves a row from 'AllergyNew.intradermal_test_results' as a IntradermalTestResult.
//
// Generated from index 'Antigens'.
func IntradermalTestResultsByAntigenID(db XODB, antigenID sql.NullInt64) ([]*IntradermalTestResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results ` +
		`WHERE antigen_id = ?`

	// run query
	XOLog(sqlstr, antigenID)
	q, err := db.Query(sqlstr, antigenID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestResult{}
	for q.Next() {
		itr := IntradermalTestResult{
			_exists: true,
		}

		// scan
		err = q.Scan(&itr.ID, &itr.TestID, &itr.AntigenID, &itr.PatientID, &itr.Dilution, &itr.Presumed, &itr.FlarePresumed, &itr.WhealSize, &itr.FlareSize, &itr.Enabled, &itr.UpdatedBy, &itr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itr)
	}

	return res, nil
}

// IntradermalTestResultsByPatientID retrieves a row from 'AllergyNew.intradermal_test_results' as a IntradermalTestResult.
//
// Generated from index 'Patients'.
func IntradermalTestResultsByPatientID(db XODB, patientID sql.NullInt64) ([]*IntradermalTestResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestResult{}
	for q.Next() {
		itr := IntradermalTestResult{
			_exists: true,
		}

		// scan
		err = q.Scan(&itr.ID, &itr.TestID, &itr.AntigenID, &itr.PatientID, &itr.Dilution, &itr.Presumed, &itr.FlarePresumed, &itr.WhealSize, &itr.FlareSize, &itr.Enabled, &itr.UpdatedBy, &itr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itr)
	}

	return res, nil
}

// IntradermalTestResultsByTestID retrieves a row from 'AllergyNew.intradermal_test_results' as a IntradermalTestResult.
//
// Generated from index 'Tests'.
func IntradermalTestResultsByTestID(db XODB, testID sql.NullInt64) ([]*IntradermalTestResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results ` +
		`WHERE test_id = ?`

	// run query
	XOLog(sqlstr, testID)
	q, err := db.Query(sqlstr, testID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestResult{}
	for q.Next() {
		itr := IntradermalTestResult{
			_exists: true,
		}

		// scan
		err = q.Scan(&itr.ID, &itr.TestID, &itr.AntigenID, &itr.PatientID, &itr.Dilution, &itr.Presumed, &itr.FlarePresumed, &itr.WhealSize, &itr.FlareSize, &itr.Enabled, &itr.UpdatedBy, &itr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itr)
	}

	return res, nil
}

// IntradermalTestResultByID retrieves a row from 'AllergyNew.intradermal_test_results' as a IntradermalTestResult.
//
// Generated from index 'intradermal_test_results_id_pkey'.
func IntradermalTestResultByID(db XODB, id uint) (*IntradermalTestResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	itr := IntradermalTestResult{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&itr.ID, &itr.TestID, &itr.AntigenID, &itr.PatientID, &itr.Dilution, &itr.Presumed, &itr.FlarePresumed, &itr.WhealSize, &itr.FlareSize, &itr.Enabled, &itr.UpdatedBy, &itr.Updated)
	if err != nil {
		return nil, err
	}

	return &itr, nil
}

// IntradermalTestResultsByUpdatedBy retrieves a row from 'AllergyNew.intradermal_test_results' as a IntradermalTestResult.
//
// Generated from index 'intradermal_test_results_updated'.
func IntradermalTestResultsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*IntradermalTestResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, test_id, antigen_id, patient_id, dilution, presumed, flare_presumed, wheal_size, flare_size, enabled, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestResult{}
	for q.Next() {
		itr := IntradermalTestResult{
			_exists: true,
		}

		// scan
		err = q.Scan(&itr.ID, &itr.TestID, &itr.AntigenID, &itr.PatientID, &itr.Dilution, &itr.Presumed, &itr.FlarePresumed, &itr.WhealSize, &itr.FlareSize, &itr.Enabled, &itr.UpdatedBy, &itr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itr)
	}

	return res, nil
}