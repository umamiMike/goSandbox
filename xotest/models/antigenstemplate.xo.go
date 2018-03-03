// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// AntigensTemplate represents a row from 'AllergyNew.antigens_templates'.
type AntigensTemplate struct {
	ID        uint           `json:"id"`         // id
	Name      sql.NullString `json:"name"`       // name
	Order     int            `json:"order"`      // order
	Enabled   int8           `json:"enabled"`    // enabled
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AntigensTemplate exists in the database.
func (at *AntigensTemplate) Exists() bool {
	return at._exists
}

// Deleted provides information if the AntigensTemplate has been deleted from the database.
func (at *AntigensTemplate) Deleted() bool {
	return at._deleted
}

// Insert inserts the AntigensTemplate to the database.
func (at *AntigensTemplate) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if at._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.antigens_templates (` +
		`name, order, enabled, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, at.Name, at.Order, at.Enabled, at.UpdatedBy, at.Updated)
	res, err := db.Exec(sqlstr, at.Name, at.Order, at.Enabled, at.UpdatedBy, at.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	at.ID = uint(id)
	at._exists = true

	return nil
}

// Update updates the AntigensTemplate in the database.
func (at *AntigensTemplate) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !at._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if at._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.antigens_templates SET ` +
		`name = ?, order = ?, enabled = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, at.Name, at.Order, at.Enabled, at.UpdatedBy, at.Updated, at.ID)
	_, err = db.Exec(sqlstr, at.Name, at.Order, at.Enabled, at.UpdatedBy, at.Updated, at.ID)
	return err
}

// Save saves the AntigensTemplate to the database.
func (at *AntigensTemplate) Save(db XODB) error {
	if at.Exists() {
		return at.Update(db)
	}

	return at.Insert(db)
}

// Delete deletes the AntigensTemplate from the database.
func (at *AntigensTemplate) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !at._exists {
		return nil
	}

	// if deleted, bail
	if at._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.antigens_templates WHERE id = ?`

	// run query
	XOLog(sqlstr, at.ID)
	_, err = db.Exec(sqlstr, at.ID)
	if err != nil {
		return err
	}

	// set deleted
	at._deleted = true

	return nil
}

// AntigensTemplateByID retrieves a row from 'AllergyNew.antigens_templates' as a AntigensTemplate.
//
// Generated from index 'antigens_templates_id_pkey'.
func AntigensTemplateByID(db XODB, id uint) (*AntigensTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, order, enabled, updated_by, updated ` +
		`FROM AllergyNew.antigens_templates ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	at := AntigensTemplate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&at.ID, &at.Name, &at.Order, &at.Enabled, &at.UpdatedBy, &at.Updated)
	if err != nil {
		return nil, err
	}

	return &at, nil
}
