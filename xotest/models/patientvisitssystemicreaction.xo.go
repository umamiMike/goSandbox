// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// PatientVisitsSystemicReaction represents a row from 'AllergyNew.patient_visits_systemic_reactions'.
type PatientVisitsSystemicReaction struct {
	ID          uint           `json:"id"`            // id
	PatientID   sql.NullInt64  `json:"patient_id"`    // patient_id
	VisitID     sql.NullInt64  `json:"visit_id"`      // visit_id
	Duration    sql.NullString `json:"duration"`      // duration
	DateAdded   mysql.NullTime `json:"date_added"`    // date_added
	AddedByID   sql.NullInt64  `json:"added_by_id"`   // added_by_id
	DateRemoved mysql.NullTime `json:"date_removed"`  // date_removed
	RemovedByID sql.NullInt64  `json:"removed_by_id"` // removed_by_id
	UpdatedBy   sql.NullInt64  `json:"updated_by"`    // updated_by
	Updated     time.Time      `json:"updated"`       // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PatientVisitsSystemicReaction exists in the database.
func (pvsr *PatientVisitsSystemicReaction) Exists() bool {
	return pvsr._exists
}

// Deleted provides information if the PatientVisitsSystemicReaction has been deleted from the database.
func (pvsr *PatientVisitsSystemicReaction) Deleted() bool {
	return pvsr._deleted
}

// Insert inserts the PatientVisitsSystemicReaction to the database.
func (pvsr *PatientVisitsSystemicReaction) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pvsr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.patient_visits_systemic_reactions (` +
		`patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pvsr.PatientID, pvsr.VisitID, pvsr.Duration, pvsr.DateAdded, pvsr.AddedByID, pvsr.DateRemoved, pvsr.RemovedByID, pvsr.UpdatedBy, pvsr.Updated)
	res, err := db.Exec(sqlstr, pvsr.PatientID, pvsr.VisitID, pvsr.Duration, pvsr.DateAdded, pvsr.AddedByID, pvsr.DateRemoved, pvsr.RemovedByID, pvsr.UpdatedBy, pvsr.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pvsr.ID = uint(id)
	pvsr._exists = true

	return nil
}

// Update updates the PatientVisitsSystemicReaction in the database.
func (pvsr *PatientVisitsSystemicReaction) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pvsr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pvsr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.patient_visits_systemic_reactions SET ` +
		`patient_id = ?, visit_id = ?, duration = ?, date_added = ?, added_by_id = ?, date_removed = ?, removed_by_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, pvsr.PatientID, pvsr.VisitID, pvsr.Duration, pvsr.DateAdded, pvsr.AddedByID, pvsr.DateRemoved, pvsr.RemovedByID, pvsr.UpdatedBy, pvsr.Updated, pvsr.ID)
	_, err = db.Exec(sqlstr, pvsr.PatientID, pvsr.VisitID, pvsr.Duration, pvsr.DateAdded, pvsr.AddedByID, pvsr.DateRemoved, pvsr.RemovedByID, pvsr.UpdatedBy, pvsr.Updated, pvsr.ID)
	return err
}

// Save saves the PatientVisitsSystemicReaction to the database.
func (pvsr *PatientVisitsSystemicReaction) Save(db XODB) error {
	if pvsr.Exists() {
		return pvsr.Update(db)
	}

	return pvsr.Insert(db)
}

// Delete deletes the PatientVisitsSystemicReaction from the database.
func (pvsr *PatientVisitsSystemicReaction) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pvsr._exists {
		return nil
	}

	// if deleted, bail
	if pvsr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.patient_visits_systemic_reactions WHERE id = ?`

	// run query
	XOLog(sqlstr, pvsr.ID)
	_, err = db.Exec(sqlstr, pvsr.ID)
	if err != nil {
		return err
	}

	// set deleted
	pvsr._deleted = true

	return nil
}

// UserByAddedByID returns the User associated with the PatientVisitsSystemicReaction's AddedByID (added_by_id).
//
// Generated from foreign key 'patient_visits_systemic_reactions_creator'.
func (pvsr *PatientVisitsSystemicReaction) UserByAddedByID(db XODB) (*User, error) {
	return UserByID(db, uint(pvsr.AddedByID.Int64))
}

// Patient returns the Patient associated with the PatientVisitsSystemicReaction's PatientID (patient_id).
//
// Generated from foreign key 'patient_visits_systemic_reactions_patient'.
func (pvsr *PatientVisitsSystemicReaction) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(pvsr.PatientID.Int64))
}

// UserByUpdatedBy returns the User associated with the PatientVisitsSystemicReaction's UpdatedBy (updated_by).
//
// Generated from foreign key 'patient_visits_systemic_reactions_updated'.
func (pvsr *PatientVisitsSystemicReaction) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(pvsr.UpdatedBy.Int64))
}

// PatientVisit returns the PatientVisit associated with the PatientVisitsSystemicReaction's VisitID (visit_id).
//
// Generated from foreign key 'patient_visits_systemic_reactions_visits'.
func (pvsr *PatientVisitsSystemicReaction) PatientVisit(db XODB) (*PatientVisit, error) {
	return PatientVisitByID(db, uint(pvsr.VisitID.Int64))
}

// PatientVisitsSystemicReactionsByAddedByID retrieves a row from 'AllergyNew.patient_visits_systemic_reactions' as a PatientVisitsSystemicReaction.
//
// Generated from index 'Creator'.
func PatientVisitsSystemicReactionsByAddedByID(db XODB, addedByID sql.NullInt64) ([]*PatientVisitsSystemicReaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated ` +
		`FROM AllergyNew.patient_visits_systemic_reactions ` +
		`WHERE added_by_id = ?`

	// run query
	XOLog(sqlstr, addedByID)
	q, err := db.Query(sqlstr, addedByID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientVisitsSystemicReaction{}
	for q.Next() {
		pvsr := PatientVisitsSystemicReaction{
			_exists: true,
		}

		// scan
		err = q.Scan(&pvsr.ID, &pvsr.PatientID, &pvsr.VisitID, &pvsr.Duration, &pvsr.DateAdded, &pvsr.AddedByID, &pvsr.DateRemoved, &pvsr.RemovedByID, &pvsr.UpdatedBy, &pvsr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pvsr)
	}

	return res, nil
}

// PatientVisitsSystemicReactionsByPatientID retrieves a row from 'AllergyNew.patient_visits_systemic_reactions' as a PatientVisitsSystemicReaction.
//
// Generated from index 'Patients'.
func PatientVisitsSystemicReactionsByPatientID(db XODB, patientID sql.NullInt64) ([]*PatientVisitsSystemicReaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated ` +
		`FROM AllergyNew.patient_visits_systemic_reactions ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientVisitsSystemicReaction{}
	for q.Next() {
		pvsr := PatientVisitsSystemicReaction{
			_exists: true,
		}

		// scan
		err = q.Scan(&pvsr.ID, &pvsr.PatientID, &pvsr.VisitID, &pvsr.Duration, &pvsr.DateAdded, &pvsr.AddedByID, &pvsr.DateRemoved, &pvsr.RemovedByID, &pvsr.UpdatedBy, &pvsr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pvsr)
	}

	return res, nil
}

// PatientVisitsSystemicReactionsByVisitID retrieves a row from 'AllergyNew.patient_visits_systemic_reactions' as a PatientVisitsSystemicReaction.
//
// Generated from index 'Visits'.
func PatientVisitsSystemicReactionsByVisitID(db XODB, visitID sql.NullInt64) ([]*PatientVisitsSystemicReaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated ` +
		`FROM AllergyNew.patient_visits_systemic_reactions ` +
		`WHERE visit_id = ?`

	// run query
	XOLog(sqlstr, visitID)
	q, err := db.Query(sqlstr, visitID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientVisitsSystemicReaction{}
	for q.Next() {
		pvsr := PatientVisitsSystemicReaction{
			_exists: true,
		}

		// scan
		err = q.Scan(&pvsr.ID, &pvsr.PatientID, &pvsr.VisitID, &pvsr.Duration, &pvsr.DateAdded, &pvsr.AddedByID, &pvsr.DateRemoved, &pvsr.RemovedByID, &pvsr.UpdatedBy, &pvsr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pvsr)
	}

	return res, nil
}

// PatientVisitsSystemicReactionByID retrieves a row from 'AllergyNew.patient_visits_systemic_reactions' as a PatientVisitsSystemicReaction.
//
// Generated from index 'patient_visits_systemic_reactions_id_pkey'.
func PatientVisitsSystemicReactionByID(db XODB, id uint) (*PatientVisitsSystemicReaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated ` +
		`FROM AllergyNew.patient_visits_systemic_reactions ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	pvsr := PatientVisitsSystemicReaction{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pvsr.ID, &pvsr.PatientID, &pvsr.VisitID, &pvsr.Duration, &pvsr.DateAdded, &pvsr.AddedByID, &pvsr.DateRemoved, &pvsr.RemovedByID, &pvsr.UpdatedBy, &pvsr.Updated)
	if err != nil {
		return nil, err
	}

	return &pvsr, nil
}

// PatientVisitsSystemicReactionsByUpdatedBy retrieves a row from 'AllergyNew.patient_visits_systemic_reactions' as a PatientVisitsSystemicReaction.
//
// Generated from index 'patient_visits_systemic_reactions_updated'.
func PatientVisitsSystemicReactionsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*PatientVisitsSystemicReaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, visit_id, duration, date_added, added_by_id, date_removed, removed_by_id, updated_by, updated ` +
		`FROM AllergyNew.patient_visits_systemic_reactions ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PatientVisitsSystemicReaction{}
	for q.Next() {
		pvsr := PatientVisitsSystemicReaction{
			_exists: true,
		}

		// scan
		err = q.Scan(&pvsr.ID, &pvsr.PatientID, &pvsr.VisitID, &pvsr.Duration, &pvsr.DateAdded, &pvsr.AddedByID, &pvsr.DateRemoved, &pvsr.RemovedByID, &pvsr.UpdatedBy, &pvsr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pvsr)
	}

	return res, nil
}
