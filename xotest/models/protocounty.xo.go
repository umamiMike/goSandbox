// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// ProtoCounty represents a row from 'AllergyNew.proto_counties'.
type ProtoCounty struct {
	ID          uint           `json:"id"`           // id
	PracticeID  sql.NullInt64  `json:"practice_id"`  // practice_id
	Name        sql.NullString `json:"name"`         // name
	DateAdded   mysql.NullTime `json:"date_added"`   // date_added
	AddedBy     sql.NullInt64  `json:"added_by"`     // added_by
	DateRemoved mysql.NullTime `json:"date_removed"` // date_removed
	RemovedBy   sql.NullInt64  `json:"removed_by"`   // removed_by
	Active      int8           `json:"active"`       // active
	UpdatedBy   sql.NullInt64  `json:"updated_by"`   // updated_by
	Updated     time.Time      `json:"updated"`      // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ProtoCounty exists in the database.
func (pc *ProtoCounty) Exists() bool {
	return pc._exists
}

// Deleted provides information if the ProtoCounty has been deleted from the database.
func (pc *ProtoCounty) Deleted() bool {
	return pc._deleted
}

// Insert inserts the ProtoCounty to the database.
func (pc *ProtoCounty) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pc._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.proto_counties (` +
		`practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pc.PracticeID, pc.Name, pc.DateAdded, pc.AddedBy, pc.DateRemoved, pc.RemovedBy, pc.Active, pc.UpdatedBy, pc.Updated)
	res, err := db.Exec(sqlstr, pc.PracticeID, pc.Name, pc.DateAdded, pc.AddedBy, pc.DateRemoved, pc.RemovedBy, pc.Active, pc.UpdatedBy, pc.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pc.ID = uint(id)
	pc._exists = true

	return nil
}

// Update updates the ProtoCounty in the database.
func (pc *ProtoCounty) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pc._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.proto_counties SET ` +
		`practice_id = ?, name = ?, date_added = ?, added_by = ?, date_removed = ?, removed_by = ?, active = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, pc.PracticeID, pc.Name, pc.DateAdded, pc.AddedBy, pc.DateRemoved, pc.RemovedBy, pc.Active, pc.UpdatedBy, pc.Updated, pc.ID)
	_, err = db.Exec(sqlstr, pc.PracticeID, pc.Name, pc.DateAdded, pc.AddedBy, pc.DateRemoved, pc.RemovedBy, pc.Active, pc.UpdatedBy, pc.Updated, pc.ID)
	return err
}

// Save saves the ProtoCounty to the database.
func (pc *ProtoCounty) Save(db XODB) error {
	if pc.Exists() {
		return pc.Update(db)
	}

	return pc.Insert(db)
}

// Delete deletes the ProtoCounty from the database.
func (pc *ProtoCounty) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pc._exists {
		return nil
	}

	// if deleted, bail
	if pc._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.proto_counties WHERE id = ?`

	// run query
	XOLog(sqlstr, pc.ID)
	_, err = db.Exec(sqlstr, pc.ID)
	if err != nil {
		return err
	}

	// set deleted
	pc._deleted = true

	return nil
}

// UserByAddedBy returns the User associated with the ProtoCounty's AddedBy (added_by).
//
// Generated from foreign key 'proto_counties_added_by'.
func (pc *ProtoCounty) UserByAddedBy(db XODB) (*User, error) {
	return UserByID(db, uint(pc.AddedBy.Int64))
}

// Practice returns the Practice associated with the ProtoCounty's PracticeID (practice_id).
//
// Generated from foreign key 'proto_counties_practice'.
func (pc *ProtoCounty) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(pc.PracticeID.Int64))
}

// UserByRemovedBy returns the User associated with the ProtoCounty's RemovedBy (removed_by).
//
// Generated from foreign key 'proto_counties_removed_by'.
func (pc *ProtoCounty) UserByRemovedBy(db XODB) (*User, error) {
	return UserByID(db, uint(pc.RemovedBy.Int64))
}

// UserByUpdatedBy returns the User associated with the ProtoCounty's UpdatedBy (updated_by).
//
// Generated from foreign key 'proto_counties_updated'.
func (pc *ProtoCounty) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(pc.UpdatedBy.Int64))
}

// ProtoCountiesByPracticeID retrieves a row from 'AllergyNew.proto_counties' as a ProtoCounty.
//
// Generated from index 'Practices'.
func ProtoCountiesByPracticeID(db XODB, practiceID sql.NullInt64) ([]*ProtoCounty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated ` +
		`FROM AllergyNew.proto_counties ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ProtoCounty{}
	for q.Next() {
		pc := ProtoCounty{
			_exists: true,
		}

		// scan
		err = q.Scan(&pc.ID, &pc.PracticeID, &pc.Name, &pc.DateAdded, &pc.AddedBy, &pc.DateRemoved, &pc.RemovedBy, &pc.Active, &pc.UpdatedBy, &pc.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pc)
	}

	return res, nil
}

// ProtoCountiesByAddedBy retrieves a row from 'AllergyNew.proto_counties' as a ProtoCounty.
//
// Generated from index 'proto_counties_added_by'.
func ProtoCountiesByAddedBy(db XODB, addedBy sql.NullInt64) ([]*ProtoCounty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated ` +
		`FROM AllergyNew.proto_counties ` +
		`WHERE added_by = ?`

	// run query
	XOLog(sqlstr, addedBy)
	q, err := db.Query(sqlstr, addedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ProtoCounty{}
	for q.Next() {
		pc := ProtoCounty{
			_exists: true,
		}

		// scan
		err = q.Scan(&pc.ID, &pc.PracticeID, &pc.Name, &pc.DateAdded, &pc.AddedBy, &pc.DateRemoved, &pc.RemovedBy, &pc.Active, &pc.UpdatedBy, &pc.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pc)
	}

	return res, nil
}

// ProtoCountyByID retrieves a row from 'AllergyNew.proto_counties' as a ProtoCounty.
//
// Generated from index 'proto_counties_id_pkey'.
func ProtoCountyByID(db XODB, id uint) (*ProtoCounty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated ` +
		`FROM AllergyNew.proto_counties ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	pc := ProtoCounty{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pc.ID, &pc.PracticeID, &pc.Name, &pc.DateAdded, &pc.AddedBy, &pc.DateRemoved, &pc.RemovedBy, &pc.Active, &pc.UpdatedBy, &pc.Updated)
	if err != nil {
		return nil, err
	}

	return &pc, nil
}

// ProtoCountiesByRemovedBy retrieves a row from 'AllergyNew.proto_counties' as a ProtoCounty.
//
// Generated from index 'proto_counties_removed_by'.
func ProtoCountiesByRemovedBy(db XODB, removedBy sql.NullInt64) ([]*ProtoCounty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated ` +
		`FROM AllergyNew.proto_counties ` +
		`WHERE removed_by = ?`

	// run query
	XOLog(sqlstr, removedBy)
	q, err := db.Query(sqlstr, removedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ProtoCounty{}
	for q.Next() {
		pc := ProtoCounty{
			_exists: true,
		}

		// scan
		err = q.Scan(&pc.ID, &pc.PracticeID, &pc.Name, &pc.DateAdded, &pc.AddedBy, &pc.DateRemoved, &pc.RemovedBy, &pc.Active, &pc.UpdatedBy, &pc.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pc)
	}

	return res, nil
}

// ProtoCountiesByUpdatedBy retrieves a row from 'AllergyNew.proto_counties' as a ProtoCounty.
//
// Generated from index 'proto_counties_updated'.
func ProtoCountiesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*ProtoCounty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, name, date_added, added_by, date_removed, removed_by, active, updated_by, updated ` +
		`FROM AllergyNew.proto_counties ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ProtoCounty{}
	for q.Next() {
		pc := ProtoCounty{
			_exists: true,
		}

		// scan
		err = q.Scan(&pc.ID, &pc.PracticeID, &pc.Name, &pc.DateAdded, &pc.AddedBy, &pc.DateRemoved, &pc.RemovedBy, &pc.Active, &pc.UpdatedBy, &pc.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &pc)
	}

	return res, nil
}