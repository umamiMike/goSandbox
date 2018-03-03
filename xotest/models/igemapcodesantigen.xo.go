// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// IgeMapCodesAntigen represents a row from 'AllergyNew.ige_map_codes_antigens'.
type IgeMapCodesAntigen struct {
	ID         uint          `json:"id"`          // id
	PracticeID uint          `json:"practice_id"` // practice_id
	Code       int           `json:"code"`        // code
	AntigenID  int           `json:"antigen_id"`  // antigen_id
	UpdatedBy  sql.NullInt64 `json:"updated_by"`  // updated_by
	Updated    time.Time     `json:"updated"`     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the IgeMapCodesAntigen exists in the database.
func (imca *IgeMapCodesAntigen) Exists() bool {
	return imca._exists
}

// Deleted provides information if the IgeMapCodesAntigen has been deleted from the database.
func (imca *IgeMapCodesAntigen) Deleted() bool {
	return imca._deleted
}

// Insert inserts the IgeMapCodesAntigen to the database.
func (imca *IgeMapCodesAntigen) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if imca._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.ige_map_codes_antigens (` +
		`practice_id, code, antigen_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, imca.PracticeID, imca.Code, imca.AntigenID, imca.UpdatedBy, imca.Updated)
	res, err := db.Exec(sqlstr, imca.PracticeID, imca.Code, imca.AntigenID, imca.UpdatedBy, imca.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	imca.ID = uint(id)
	imca._exists = true

	return nil
}

// Update updates the IgeMapCodesAntigen in the database.
func (imca *IgeMapCodesAntigen) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !imca._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if imca._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.ige_map_codes_antigens SET ` +
		`practice_id = ?, code = ?, antigen_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, imca.PracticeID, imca.Code, imca.AntigenID, imca.UpdatedBy, imca.Updated, imca.ID)
	_, err = db.Exec(sqlstr, imca.PracticeID, imca.Code, imca.AntigenID, imca.UpdatedBy, imca.Updated, imca.ID)
	return err
}

// Save saves the IgeMapCodesAntigen to the database.
func (imca *IgeMapCodesAntigen) Save(db XODB) error {
	if imca.Exists() {
		return imca.Update(db)
	}

	return imca.Insert(db)
}

// Delete deletes the IgeMapCodesAntigen from the database.
func (imca *IgeMapCodesAntigen) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !imca._exists {
		return nil
	}

	// if deleted, bail
	if imca._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.ige_map_codes_antigens WHERE id = ?`

	// run query
	XOLog(sqlstr, imca.ID)
	_, err = db.Exec(sqlstr, imca.ID)
	if err != nil {
		return err
	}

	// set deleted
	imca._deleted = true

	return nil
}

// IgeMapCodesAntigenByID retrieves a row from 'AllergyNew.ige_map_codes_antigens' as a IgeMapCodesAntigen.
//
// Generated from index 'ige_map_codes_antigens_id_pkey'.
func IgeMapCodesAntigenByID(db XODB, id uint) (*IgeMapCodesAntigen, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, code, antigen_id, updated_by, updated ` +
		`FROM AllergyNew.ige_map_codes_antigens ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	imca := IgeMapCodesAntigen{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&imca.ID, &imca.PracticeID, &imca.Code, &imca.AntigenID, &imca.UpdatedBy, &imca.Updated)
	if err != nil {
		return nil, err
	}

	return &imca, nil
}
