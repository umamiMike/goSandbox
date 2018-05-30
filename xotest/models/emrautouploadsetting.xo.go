// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// EmrAutouploadSetting represents a row from 'AllergyNew.emr_autoupload_settings'.
type EmrAutouploadSetting struct {
	ID           uint           `json:"id"`            // id
	PracticeID   sql.NullInt64  `json:"practice_id"`   // practice_id
	DocumentType sql.NullString `json:"document_type"` // document_type
	AutoSubmit   sql.NullInt64  `json:"auto_submit"`   // auto_submit
	SubmitEvent  sql.NullString `json:"submit_event"`  // submit_event
	UpdatedBy    sql.NullInt64  `json:"updated_by"`    // updated_by
	Updated      time.Time      `json:"updated"`       // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the EmrAutouploadSetting exists in the database.
func (eas *EmrAutouploadSetting) Exists() bool {
	return eas._exists
}

// Deleted provides information if the EmrAutouploadSetting has been deleted from the database.
func (eas *EmrAutouploadSetting) Deleted() bool {
	return eas._deleted
}

// Insert inserts the EmrAutouploadSetting to the database.
func (eas *EmrAutouploadSetting) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if eas._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.emr_autoupload_settings (` +
		`practice_id, document_type, auto_submit, submit_event, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, eas.PracticeID, eas.DocumentType, eas.AutoSubmit, eas.SubmitEvent, eas.UpdatedBy, eas.Updated)
	res, err := db.Exec(sqlstr, eas.PracticeID, eas.DocumentType, eas.AutoSubmit, eas.SubmitEvent, eas.UpdatedBy, eas.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	eas.ID = uint(id)
	eas._exists = true

	return nil
}

// Update updates the EmrAutouploadSetting in the database.
func (eas *EmrAutouploadSetting) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eas._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if eas._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.emr_autoupload_settings SET ` +
		`practice_id = ?, document_type = ?, auto_submit = ?, submit_event = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, eas.PracticeID, eas.DocumentType, eas.AutoSubmit, eas.SubmitEvent, eas.UpdatedBy, eas.Updated, eas.ID)
	_, err = db.Exec(sqlstr, eas.PracticeID, eas.DocumentType, eas.AutoSubmit, eas.SubmitEvent, eas.UpdatedBy, eas.Updated, eas.ID)
	return err
}

// Save saves the EmrAutouploadSetting to the database.
func (eas *EmrAutouploadSetting) Save(db XODB) error {
	if eas.Exists() {
		return eas.Update(db)
	}

	return eas.Insert(db)
}

// Delete deletes the EmrAutouploadSetting from the database.
func (eas *EmrAutouploadSetting) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !eas._exists {
		return nil
	}

	// if deleted, bail
	if eas._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.emr_autoupload_settings WHERE id = ?`

	// run query
	XOLog(sqlstr, eas.ID)
	_, err = db.Exec(sqlstr, eas.ID)
	if err != nil {
		return err
	}

	// set deleted
	eas._deleted = true

	return nil
}

// EmrAutouploadSettingByID retrieves a row from 'AllergyNew.emr_autoupload_settings' as a EmrAutouploadSetting.
//
// Generated from index 'emr_autoupload_settings_id_pkey'.
func EmrAutouploadSettingByID(db XODB, id uint) (*EmrAutouploadSetting, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, document_type, auto_submit, submit_event, updated_by, updated ` +
		`FROM AllergyNew.emr_autoupload_settings ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	eas := EmrAutouploadSetting{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&eas.ID, &eas.PracticeID, &eas.DocumentType, &eas.AutoSubmit, &eas.SubmitEvent, &eas.UpdatedBy, &eas.Updated)
	if err != nil {
		return nil, err
	}

	return &eas, nil
}