// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// VialRelocation represents a row from 'AllergyNew.vial_relocations'.
type VialRelocation struct {
	ID                   uint           `json:"id"`                     // id
	VialID               sql.NullInt64  `json:"vial_id"`                // vial_id
	PatientID            sql.NullInt64  `json:"patient_id"`             // patient_id
	CategoryID           sql.NullInt64  `json:"category_id"`            // category_id
	FromPracticeID       sql.NullInt64  `json:"from_practice_id"`       // from_practice_id
	FromLocationID       sql.NullInt64  `json:"from_location_id"`       // from_location_id
	FromLocationName     sql.NullString `json:"from_location_name"`     // from_location_name
	SentDate             mysql.NullTime `json:"sent_date"`              // sent_date
	ToPracticeID         sql.NullInt64  `json:"to_practice_id"`         // to_practice_id
	ToLocationID         sql.NullInt64  `json:"to_location_id"`         // to_location_id
	ToLocationName       sql.NullString `json:"to_location_name"`       // to_location_name
	SenderID             sql.NullInt64  `json:"sender_id"`              // sender_id
	SenderName           sql.NullString `json:"sender_name"`            // sender_name
	ReceiverID           sql.NullInt64  `json:"receiver_id"`            // receiver_id
	ReceiverName         sql.NullString `json:"receiver_name"`          // receiver_name
	ReceivedLocationID   sql.NullInt64  `json:"received_location_id"`   // received_location_id
	ReceivedLocationName sql.NullString `json:"received_location_name"` // received_location_name
	ReceivedDate         mysql.NullTime `json:"received_date"`          // received_date
	Enabled              int8           `json:"enabled"`                // enabled
	Notes                sql.NullString `json:"notes"`                  // notes
	UpdatedBy            sql.NullInt64  `json:"updated_by"`             // updated_by
	Updated              time.Time      `json:"updated"`                // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the VialRelocation exists in the database.
func (vr *VialRelocation) Exists() bool {
	return vr._exists
}

// Deleted provides information if the VialRelocation has been deleted from the database.
func (vr *VialRelocation) Deleted() bool {
	return vr._deleted
}

// Insert inserts the VialRelocation to the database.
func (vr *VialRelocation) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if vr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.vial_relocations (` +
		`vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, vr.VialID, vr.PatientID, vr.CategoryID, vr.FromPracticeID, vr.FromLocationID, vr.FromLocationName, vr.SentDate, vr.ToPracticeID, vr.ToLocationID, vr.ToLocationName, vr.SenderID, vr.SenderName, vr.ReceiverID, vr.ReceiverName, vr.ReceivedLocationID, vr.ReceivedLocationName, vr.ReceivedDate, vr.Enabled, vr.Notes, vr.UpdatedBy, vr.Updated)
	res, err := db.Exec(sqlstr, vr.VialID, vr.PatientID, vr.CategoryID, vr.FromPracticeID, vr.FromLocationID, vr.FromLocationName, vr.SentDate, vr.ToPracticeID, vr.ToLocationID, vr.ToLocationName, vr.SenderID, vr.SenderName, vr.ReceiverID, vr.ReceiverName, vr.ReceivedLocationID, vr.ReceivedLocationName, vr.ReceivedDate, vr.Enabled, vr.Notes, vr.UpdatedBy, vr.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	vr.ID = uint(id)
	vr._exists = true

	return nil
}

// Update updates the VialRelocation in the database.
func (vr *VialRelocation) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if vr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.vial_relocations SET ` +
		`vial_id = ?, patient_id = ?, category_id = ?, from_practice_id = ?, from_location_id = ?, from_location_name = ?, sent_date = ?, to_practice_id = ?, to_location_id = ?, to_location_name = ?, sender_id = ?, sender_name = ?, receiver_id = ?, receiver_name = ?, received_location_id = ?, received_location_name = ?, received_date = ?, enabled = ?, notes = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, vr.VialID, vr.PatientID, vr.CategoryID, vr.FromPracticeID, vr.FromLocationID, vr.FromLocationName, vr.SentDate, vr.ToPracticeID, vr.ToLocationID, vr.ToLocationName, vr.SenderID, vr.SenderName, vr.ReceiverID, vr.ReceiverName, vr.ReceivedLocationID, vr.ReceivedLocationName, vr.ReceivedDate, vr.Enabled, vr.Notes, vr.UpdatedBy, vr.Updated, vr.ID)
	_, err = db.Exec(sqlstr, vr.VialID, vr.PatientID, vr.CategoryID, vr.FromPracticeID, vr.FromLocationID, vr.FromLocationName, vr.SentDate, vr.ToPracticeID, vr.ToLocationID, vr.ToLocationName, vr.SenderID, vr.SenderName, vr.ReceiverID, vr.ReceiverName, vr.ReceivedLocationID, vr.ReceivedLocationName, vr.ReceivedDate, vr.Enabled, vr.Notes, vr.UpdatedBy, vr.Updated, vr.ID)
	return err
}

// Save saves the VialRelocation to the database.
func (vr *VialRelocation) Save(db XODB) error {
	if vr.Exists() {
		return vr.Update(db)
	}

	return vr.Insert(db)
}

// Delete deletes the VialRelocation from the database.
func (vr *VialRelocation) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vr._exists {
		return nil
	}

	// if deleted, bail
	if vr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.vial_relocations WHERE id = ?`

	// run query
	XOLog(sqlstr, vr.ID)
	_, err = db.Exec(sqlstr, vr.ID)
	if err != nil {
		return err
	}

	// set deleted
	vr._deleted = true

	return nil
}

// VialRelocationCategory returns the VialRelocationCategory associated with the VialRelocation's CategoryID (category_id).
//
// Generated from foreign key 'vial_relocations_category'.
func (vr *VialRelocation) VialRelocationCategory(db XODB) (*VialRelocationCategory, error) {
	return VialRelocationCategoryByID(db, uint(vr.CategoryID.Int64))
}

// Patient returns the Patient associated with the VialRelocation's PatientID (patient_id).
//
// Generated from foreign key 'vial_relocations_patients'.
func (vr *VialRelocation) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(vr.PatientID.Int64))
}

// User returns the User associated with the VialRelocation's UpdatedBy (updated_by).
//
// Generated from foreign key 'vial_relocations_updated'.
func (vr *VialRelocation) User(db XODB) (*User, error) {
	return UserByID(db, uint(vr.UpdatedBy.Int64))
}

// Vial returns the Vial associated with the VialRelocation's VialID (vial_id).
//
// Generated from foreign key 'vial_relocations_vials'.
func (vr *VialRelocation) Vial(db XODB) (*Vial, error) {
	return VialByID(db, uint(vr.VialID.Int64))
}

// VialRelocationsByCategoryID retrieves a row from 'AllergyNew.vial_relocations' as a VialRelocation.
//
// Generated from index 'Categories'.
func VialRelocationsByCategoryID(db XODB, categoryID sql.NullInt64) ([]*VialRelocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated ` +
		`FROM AllergyNew.vial_relocations ` +
		`WHERE category_id = ?`

	// run query
	XOLog(sqlstr, categoryID)
	q, err := db.Query(sqlstr, categoryID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialRelocation{}
	for q.Next() {
		vr := VialRelocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&vr.ID, &vr.VialID, &vr.PatientID, &vr.CategoryID, &vr.FromPracticeID, &vr.FromLocationID, &vr.FromLocationName, &vr.SentDate, &vr.ToPracticeID, &vr.ToLocationID, &vr.ToLocationName, &vr.SenderID, &vr.SenderName, &vr.ReceiverID, &vr.ReceiverName, &vr.ReceivedLocationID, &vr.ReceivedLocationName, &vr.ReceivedDate, &vr.Enabled, &vr.Notes, &vr.UpdatedBy, &vr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vr)
	}

	return res, nil
}

// VialRelocationsByPatientID retrieves a row from 'AllergyNew.vial_relocations' as a VialRelocation.
//
// Generated from index 'Patients'.
func VialRelocationsByPatientID(db XODB, patientID sql.NullInt64) ([]*VialRelocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated ` +
		`FROM AllergyNew.vial_relocations ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialRelocation{}
	for q.Next() {
		vr := VialRelocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&vr.ID, &vr.VialID, &vr.PatientID, &vr.CategoryID, &vr.FromPracticeID, &vr.FromLocationID, &vr.FromLocationName, &vr.SentDate, &vr.ToPracticeID, &vr.ToLocationID, &vr.ToLocationName, &vr.SenderID, &vr.SenderName, &vr.ReceiverID, &vr.ReceiverName, &vr.ReceivedLocationID, &vr.ReceivedLocationName, &vr.ReceivedDate, &vr.Enabled, &vr.Notes, &vr.UpdatedBy, &vr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vr)
	}

	return res, nil
}

// VialRelocationsByVialID retrieves a row from 'AllergyNew.vial_relocations' as a VialRelocation.
//
// Generated from index 'Vials'.
func VialRelocationsByVialID(db XODB, vialID sql.NullInt64) ([]*VialRelocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated ` +
		`FROM AllergyNew.vial_relocations ` +
		`WHERE vial_id = ?`

	// run query
	XOLog(sqlstr, vialID)
	q, err := db.Query(sqlstr, vialID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialRelocation{}
	for q.Next() {
		vr := VialRelocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&vr.ID, &vr.VialID, &vr.PatientID, &vr.CategoryID, &vr.FromPracticeID, &vr.FromLocationID, &vr.FromLocationName, &vr.SentDate, &vr.ToPracticeID, &vr.ToLocationID, &vr.ToLocationName, &vr.SenderID, &vr.SenderName, &vr.ReceiverID, &vr.ReceiverName, &vr.ReceivedLocationID, &vr.ReceivedLocationName, &vr.ReceivedDate, &vr.Enabled, &vr.Notes, &vr.UpdatedBy, &vr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vr)
	}

	return res, nil
}

// VialRelocationByID retrieves a row from 'AllergyNew.vial_relocations' as a VialRelocation.
//
// Generated from index 'vial_relocations_id_pkey'.
func VialRelocationByID(db XODB, id uint) (*VialRelocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated ` +
		`FROM AllergyNew.vial_relocations ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	vr := VialRelocation{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&vr.ID, &vr.VialID, &vr.PatientID, &vr.CategoryID, &vr.FromPracticeID, &vr.FromLocationID, &vr.FromLocationName, &vr.SentDate, &vr.ToPracticeID, &vr.ToLocationID, &vr.ToLocationName, &vr.SenderID, &vr.SenderName, &vr.ReceiverID, &vr.ReceiverName, &vr.ReceivedLocationID, &vr.ReceivedLocationName, &vr.ReceivedDate, &vr.Enabled, &vr.Notes, &vr.UpdatedBy, &vr.Updated)
	if err != nil {
		return nil, err
	}

	return &vr, nil
}

// VialRelocationsByUpdatedBy retrieves a row from 'AllergyNew.vial_relocations' as a VialRelocation.
//
// Generated from index 'vial_relocations_updated'.
func VialRelocationsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*VialRelocation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_id, patient_id, category_id, from_practice_id, from_location_id, from_location_name, sent_date, to_practice_id, to_location_id, to_location_name, sender_id, sender_name, receiver_id, receiver_name, received_location_id, received_location_name, received_date, enabled, notes, updated_by, updated ` +
		`FROM AllergyNew.vial_relocations ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialRelocation{}
	for q.Next() {
		vr := VialRelocation{
			_exists: true,
		}

		// scan
		err = q.Scan(&vr.ID, &vr.VialID, &vr.PatientID, &vr.CategoryID, &vr.FromPracticeID, &vr.FromLocationID, &vr.FromLocationName, &vr.SentDate, &vr.ToPracticeID, &vr.ToLocationID, &vr.ToLocationName, &vr.SenderID, &vr.SenderName, &vr.ReceiverID, &vr.ReceiverName, &vr.ReceivedLocationID, &vr.ReceivedLocationName, &vr.ReceivedDate, &vr.Enabled, &vr.Notes, &vr.UpdatedBy, &vr.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vr)
	}

	return res, nil
}
