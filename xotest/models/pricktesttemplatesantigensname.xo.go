// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// PrickTestTemplatesAntigensName represents a row from 'AllergyNew.prick_test_templates_antigens_names'.
type PrickTestTemplatesAntigensName struct {
	ID        uint           `json:"id"`         // id
	AntigenID sql.NullInt64  `json:"antigen_id"` // antigen_id
	Name      sql.NullString `json:"name"`       // name
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PrickTestTemplatesAntigensName exists in the database.
func (pttan *PrickTestTemplatesAntigensName) Exists() bool {
	return pttan._exists
}

// Deleted provides information if the PrickTestTemplatesAntigensName has been deleted from the database.
func (pttan *PrickTestTemplatesAntigensName) Deleted() bool {
	return pttan._deleted
}

// Insert inserts the PrickTestTemplatesAntigensName to the database.
func (pttan *PrickTestTemplatesAntigensName) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pttan._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.prick_test_templates_antigens_names (` +
		`antigen_id, name, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, pttan.AntigenID, pttan.Name, pttan.UpdatedBy, pttan.Updated)
	res, err := db.Exec(sqlstr, pttan.AntigenID, pttan.Name, pttan.UpdatedBy, pttan.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	pttan.ID = uint(id)
	pttan._exists = true

	return nil
}

// Update updates the PrickTestTemplatesAntigensName in the database.
func (pttan *PrickTestTemplatesAntigensName) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pttan._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pttan._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.prick_test_templates_antigens_names SET ` +
		`antigen_id = ?, name = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, pttan.AntigenID, pttan.Name, pttan.UpdatedBy, pttan.Updated, pttan.ID)
	_, err = db.Exec(sqlstr, pttan.AntigenID, pttan.Name, pttan.UpdatedBy, pttan.Updated, pttan.ID)
	return err
}

// Save saves the PrickTestTemplatesAntigensName to the database.
func (pttan *PrickTestTemplatesAntigensName) Save(db XODB) error {
	if pttan.Exists() {
		return pttan.Update(db)
	}

	return pttan.Insert(db)
}

// Delete deletes the PrickTestTemplatesAntigensName from the database.
func (pttan *PrickTestTemplatesAntigensName) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pttan._exists {
		return nil
	}

	// if deleted, bail
	if pttan._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.prick_test_templates_antigens_names WHERE id = ?`

	// run query
	XOLog(sqlstr, pttan.ID)
	_, err = db.Exec(sqlstr, pttan.ID)
	if err != nil {
		return err
	}

	// set deleted
	pttan._deleted = true

	return nil
}

// PrickTestTemplatesAntigensNameByID retrieves a row from 'AllergyNew.prick_test_templates_antigens_names' as a PrickTestTemplatesAntigensName.
//
// Generated from index 'prick_test_templates_antigens_names_id_pkey'.
func PrickTestTemplatesAntigensNameByID(db XODB, id uint) (*PrickTestTemplatesAntigensName, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, name, updated_by, updated ` +
		`FROM AllergyNew.prick_test_templates_antigens_names ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	pttan := PrickTestTemplatesAntigensName{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pttan.ID, &pttan.AntigenID, &pttan.Name, &pttan.UpdatedBy, &pttan.Updated)
	if err != nil {
		return nil, err
	}

	return &pttan, nil
}