// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// EmrOrdersDefaultsSetting represents a row from 'AllergyNew.emr_orders_defaults_settings'.
type EmrOrdersDefaultsSetting struct {
	ID                    uint           `json:"id"`                      // id
	PracticeID            sql.NullInt64  `json:"practice_id"`             // practice_id
	DocumentType          sql.NullString `json:"document_type"`           // document_type
	ProcedureCode         sql.NullString `json:"procedure_code"`          // procedure_code
	ProcedureCodeset      sql.NullString `json:"procedure_codeset"`       // procedure_codeset
	ProcedureDescription  sql.NullString `json:"procedure_description"`   // procedure_description
	OverrideProcedureInfo sql.NullInt64  `json:"override_procedure_info"` // override_procedure_info
	UpdatedBy             sql.NullInt64  `json:"updated_by"`              // updated_by
	Updated               time.Time      `json:"updated"`                 // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the EmrOrdersDefaultsSetting exists in the database.
func (eods *EmrOrdersDefaultsSetting) Exists() bool {
	return eods._exists
}

// Deleted provides information if the EmrOrdersDefaultsSetting has been deleted from the database.
func (eods *EmrOrdersDefaultsSetting) Deleted() bool {
	return eods._deleted
}

// Insert inserts the EmrOrdersDefaultsSetting to the database.
func (eods *EmrOrdersDefaultsSetting) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if eods._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.emr_orders_defaults_settings (` +
		`practice_id, document_type, procedure_code, procedure_codeset, procedure_description, override_procedure_info, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, eods.PracticeID, eods.DocumentType, eods.ProcedureCode, eods.ProcedureCodeset, eods.ProcedureDescription, eods.OverrideProcedureInfo, eods.UpdatedBy, eods.Updated)
	res, err := db.Exec(sqlstr, eods.PracticeID, eods.DocumentType, eods.ProcedureCode, eods.ProcedureCodeset, eods.ProcedureDescription, eods.OverrideProcedureInfo, eods.UpdatedBy, eods.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	eods.ID = uint(id)
	eods._exists = true

	return nil
}

// Update updates the EmrOrdersDefaultsSetting in the database.
func (eods *EmrOrdersDefaultsSetting) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eods._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if eods._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.emr_orders_defaults_settings SET ` +
		`practice_id = ?, document_type = ?, procedure_code = ?, procedure_codeset = ?, procedure_description = ?, override_procedure_info = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, eods.PracticeID, eods.DocumentType, eods.ProcedureCode, eods.ProcedureCodeset, eods.ProcedureDescription, eods.OverrideProcedureInfo, eods.UpdatedBy, eods.Updated, eods.ID)
	_, err = db.Exec(sqlstr, eods.PracticeID, eods.DocumentType, eods.ProcedureCode, eods.ProcedureCodeset, eods.ProcedureDescription, eods.OverrideProcedureInfo, eods.UpdatedBy, eods.Updated, eods.ID)
	return err
}

// Save saves the EmrOrdersDefaultsSetting to the database.
func (eods *EmrOrdersDefaultsSetting) Save(db XODB) error {
	if eods.Exists() {
		return eods.Update(db)
	}

	return eods.Insert(db)
}

// Delete deletes the EmrOrdersDefaultsSetting from the database.
func (eods *EmrOrdersDefaultsSetting) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eods._exists {
		return nil
	}

	// if deleted, bail
	if eods._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.emr_orders_defaults_settings WHERE id = ?`

	// run query
	XOLog(sqlstr, eods.ID)
	_, err = db.Exec(sqlstr, eods.ID)
	if err != nil {
		return err
	}

	// set deleted
	eods._deleted = true

	return nil
}

// EmrOrdersDefaultsSettingByID retrieves a row from 'AllergyNew.emr_orders_defaults_settings' as a EmrOrdersDefaultsSetting.
//
// Generated from index 'emr_orders_defaults_settings_id_pkey'.
func EmrOrdersDefaultsSettingByID(db XODB, id uint) (*EmrOrdersDefaultsSetting, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, document_type, procedure_code, procedure_codeset, procedure_description, override_procedure_info, updated_by, updated ` +
		`FROM AllergyNew.emr_orders_defaults_settings ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	eods := EmrOrdersDefaultsSetting{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&eods.ID, &eods.PracticeID, &eods.DocumentType, &eods.ProcedureCode, &eods.ProcedureCodeset, &eods.ProcedureDescription, &eods.OverrideProcedureInfo, &eods.UpdatedBy, &eods.Updated)
	if err != nil {
		return nil, err
	}

	return &eods, nil
}