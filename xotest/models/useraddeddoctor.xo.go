// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// UserAddedDoctor represents a row from 'AllergyNew.user_added_doctors'.
type UserAddedDoctor struct {
	ID                uint           `json:"id"`                   // id
	PracticeName      sql.NullString `json:"practice_name"`        // practice_name
	DoctorName        sql.NullString `json:"doctor_name"`          // doctor_name
	Phone             sql.NullString `json:"phone"`                // phone
	Address1          sql.NullString `json:"address1"`             // address1
	Address2          sql.NullString `json:"address2"`             // address2
	City              sql.NullString `json:"city"`                 // city
	State             sql.NullString `json:"state"`                // state
	Zip               sql.NullString `json:"zip"`                  // zip
	AddedByPracticeID sql.NullInt64  `json:"added_by_practice_id"` // added_by_practice_id
	AddedByUserID     sql.NullInt64  `json:"added_by_user_id"`     // added_by_user_id
	Comments          sql.NullString `json:"comments"`             // comments
	Active            int8           `json:"active"`               // active
	UpdatedBy         sql.NullInt64  `json:"updated_by"`           // updated_by
	Updated           time.Time      `json:"updated"`              // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserAddedDoctor exists in the database.
func (uad *UserAddedDoctor) Exists() bool {
	return uad._exists
}

// Deleted provides information if the UserAddedDoctor has been deleted from the database.
func (uad *UserAddedDoctor) Deleted() bool {
	return uad._deleted
}

// Insert inserts the UserAddedDoctor to the database.
func (uad *UserAddedDoctor) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if uad._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.user_added_doctors (` +
		`practice_name, doctor_name, phone, address1, address2, city, state, zip, added_by_practice_id, added_by_user_id, comments, active, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, uad.PracticeName, uad.DoctorName, uad.Phone, uad.Address1, uad.Address2, uad.City, uad.State, uad.Zip, uad.AddedByPracticeID, uad.AddedByUserID, uad.Comments, uad.Active, uad.UpdatedBy, uad.Updated)
	res, err := db.Exec(sqlstr, uad.PracticeName, uad.DoctorName, uad.Phone, uad.Address1, uad.Address2, uad.City, uad.State, uad.Zip, uad.AddedByPracticeID, uad.AddedByUserID, uad.Comments, uad.Active, uad.UpdatedBy, uad.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	uad.ID = uint(id)
	uad._exists = true

	return nil
}

// Update updates the UserAddedDoctor in the database.
func (uad *UserAddedDoctor) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uad._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if uad._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.user_added_doctors SET ` +
		`practice_name = ?, doctor_name = ?, phone = ?, address1 = ?, address2 = ?, city = ?, state = ?, zip = ?, added_by_practice_id = ?, added_by_user_id = ?, comments = ?, active = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, uad.PracticeName, uad.DoctorName, uad.Phone, uad.Address1, uad.Address2, uad.City, uad.State, uad.Zip, uad.AddedByPracticeID, uad.AddedByUserID, uad.Comments, uad.Active, uad.UpdatedBy, uad.Updated, uad.ID)
	_, err = db.Exec(sqlstr, uad.PracticeName, uad.DoctorName, uad.Phone, uad.Address1, uad.Address2, uad.City, uad.State, uad.Zip, uad.AddedByPracticeID, uad.AddedByUserID, uad.Comments, uad.Active, uad.UpdatedBy, uad.Updated, uad.ID)
	return err
}

// Save saves the UserAddedDoctor to the database.
func (uad *UserAddedDoctor) Save(db XODB) error {
	if uad.Exists() {
		return uad.Update(db)
	}

	return uad.Insert(db)
}

// Delete deletes the UserAddedDoctor from the database.
func (uad *UserAddedDoctor) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uad._exists {
		return nil
	}

	// if deleted, bail
	if uad._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.user_added_doctors WHERE id = ?`

	// run query
	XOLog(sqlstr, uad.ID)
	_, err = db.Exec(sqlstr, uad.ID)
	if err != nil {
		return err
	}

	// set deleted
	uad._deleted = true

	return nil
}

// Practice returns the Practice associated with the UserAddedDoctor's AddedByPracticeID (added_by_practice_id).
//
// Generated from foreign key 'user_added_doctors_practice'.
func (uad *UserAddedDoctor) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(uad.AddedByPracticeID.Int64))
}

// UserByUpdatedBy returns the User associated with the UserAddedDoctor's UpdatedBy (updated_by).
//
// Generated from foreign key 'user_added_doctors_updated'.
func (uad *UserAddedDoctor) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(uad.UpdatedBy.Int64))
}

// UserByAddedByUserID returns the User associated with the UserAddedDoctor's AddedByUserID (added_by_user_id).
//
// Generated from foreign key 'user_added_doctors_user'.
func (uad *UserAddedDoctor) UserByAddedByUserID(db XODB) (*User, error) {
	return UserByID(db, uint(uad.AddedByUserID.Int64))
}

// UserAddedDoctorsByAddedByPracticeID retrieves a row from 'AllergyNew.user_added_doctors' as a UserAddedDoctor.
//
// Generated from index 'Practices'.
func UserAddedDoctorsByAddedByPracticeID(db XODB, addedByPracticeID sql.NullInt64) ([]*UserAddedDoctor, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_name, doctor_name, phone, address1, address2, city, state, zip, added_by_practice_id, added_by_user_id, comments, active, updated_by, updated ` +
		`FROM AllergyNew.user_added_doctors ` +
		`WHERE added_by_practice_id = ?`

	// run query
	XOLog(sqlstr, addedByPracticeID)
	q, err := db.Query(sqlstr, addedByPracticeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UserAddedDoctor{}
	for q.Next() {
		uad := UserAddedDoctor{
			_exists: true,
		}

		// scan
		err = q.Scan(&uad.ID, &uad.PracticeName, &uad.DoctorName, &uad.Phone, &uad.Address1, &uad.Address2, &uad.City, &uad.State, &uad.Zip, &uad.AddedByPracticeID, &uad.AddedByUserID, &uad.Comments, &uad.Active, &uad.UpdatedBy, &uad.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &uad)
	}

	return res, nil
}

// UserAddedDoctorsByAddedByUserID retrieves a row from 'AllergyNew.user_added_doctors' as a UserAddedDoctor.
//
// Generated from index 'Users'.
func UserAddedDoctorsByAddedByUserID(db XODB, addedByUserID sql.NullInt64) ([]*UserAddedDoctor, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_name, doctor_name, phone, address1, address2, city, state, zip, added_by_practice_id, added_by_user_id, comments, active, updated_by, updated ` +
		`FROM AllergyNew.user_added_doctors ` +
		`WHERE added_by_user_id = ?`

	// run query
	XOLog(sqlstr, addedByUserID)
	q, err := db.Query(sqlstr, addedByUserID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UserAddedDoctor{}
	for q.Next() {
		uad := UserAddedDoctor{
			_exists: true,
		}

		// scan
		err = q.Scan(&uad.ID, &uad.PracticeName, &uad.DoctorName, &uad.Phone, &uad.Address1, &uad.Address2, &uad.City, &uad.State, &uad.Zip, &uad.AddedByPracticeID, &uad.AddedByUserID, &uad.Comments, &uad.Active, &uad.UpdatedBy, &uad.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &uad)
	}

	return res, nil
}

// UserAddedDoctorByID retrieves a row from 'AllergyNew.user_added_doctors' as a UserAddedDoctor.
//
// Generated from index 'user_added_doctors_id_pkey'.
func UserAddedDoctorByID(db XODB, id uint) (*UserAddedDoctor, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_name, doctor_name, phone, address1, address2, city, state, zip, added_by_practice_id, added_by_user_id, comments, active, updated_by, updated ` +
		`FROM AllergyNew.user_added_doctors ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	uad := UserAddedDoctor{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&uad.ID, &uad.PracticeName, &uad.DoctorName, &uad.Phone, &uad.Address1, &uad.Address2, &uad.City, &uad.State, &uad.Zip, &uad.AddedByPracticeID, &uad.AddedByUserID, &uad.Comments, &uad.Active, &uad.UpdatedBy, &uad.Updated)
	if err != nil {
		return nil, err
	}

	return &uad, nil
}

// UserAddedDoctorsByUpdatedBy retrieves a row from 'AllergyNew.user_added_doctors' as a UserAddedDoctor.
//
// Generated from index 'user_added_doctors_updated'.
func UserAddedDoctorsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*UserAddedDoctor, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_name, doctor_name, phone, address1, address2, city, state, zip, added_by_practice_id, added_by_user_id, comments, active, updated_by, updated ` +
		`FROM AllergyNew.user_added_doctors ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UserAddedDoctor{}
	for q.Next() {
		uad := UserAddedDoctor{
			_exists: true,
		}

		// scan
		err = q.Scan(&uad.ID, &uad.PracticeName, &uad.DoctorName, &uad.Phone, &uad.Address1, &uad.Address2, &uad.City, &uad.State, &uad.Zip, &uad.AddedByPracticeID, &uad.AddedByUserID, &uad.Comments, &uad.Active, &uad.UpdatedBy, &uad.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &uad)
	}

	return res, nil
}