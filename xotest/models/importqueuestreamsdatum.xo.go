// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// ImportQueueStreamsDatum represents a row from 'AllergyNew.import_queue_streams_data'.
type ImportQueueStreamsDatum struct {
	ID         int    `json:"id"`          // id
	PracticeID int    `json:"practice_id"` // practice_id
	Filename   string `json:"filename"`    // filename
	Data       []byte `json:"data"`        // data

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ImportQueueStreamsDatum exists in the database.
func (iqsd *ImportQueueStreamsDatum) Exists() bool {
	return iqsd._exists
}

// Deleted provides information if the ImportQueueStreamsDatum has been deleted from the database.
func (iqsd *ImportQueueStreamsDatum) Deleted() bool {
	return iqsd._deleted
}

// Insert inserts the ImportQueueStreamsDatum to the database.
func (iqsd *ImportQueueStreamsDatum) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if iqsd._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.import_queue_streams_data (` +
		`practice_id, filename, data` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, iqsd.PracticeID, iqsd.Filename, iqsd.Data)
	res, err := db.Exec(sqlstr, iqsd.PracticeID, iqsd.Filename, iqsd.Data)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	iqsd.ID = int(id)
	iqsd._exists = true

	return nil
}

// Update updates the ImportQueueStreamsDatum in the database.
func (iqsd *ImportQueueStreamsDatum) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iqsd._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if iqsd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.import_queue_streams_data SET ` +
		`practice_id = ?, filename = ?, data = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, iqsd.PracticeID, iqsd.Filename, iqsd.Data, iqsd.ID)
	_, err = db.Exec(sqlstr, iqsd.PracticeID, iqsd.Filename, iqsd.Data, iqsd.ID)
	return err
}

// Save saves the ImportQueueStreamsDatum to the database.
func (iqsd *ImportQueueStreamsDatum) Save(db XODB) error {
	if iqsd.Exists() {
		return iqsd.Update(db)
	}

	return iqsd.Insert(db)
}

// Delete deletes the ImportQueueStreamsDatum from the database.
func (iqsd *ImportQueueStreamsDatum) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iqsd._exists {
		return nil
	}

	// if deleted, bail
	if iqsd._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.import_queue_streams_data WHERE id = ?`

	// run query
	XOLog(sqlstr, iqsd.ID)
	_, err = db.Exec(sqlstr, iqsd.ID)
	if err != nil {
		return err
	}

	// set deleted
	iqsd._deleted = true

	return nil
}

// ImportQueueStreamsDataByFilename retrieves a row from 'AllergyNew.import_queue_streams_data' as a ImportQueueStreamsDatum.
//
// Generated from index 'XMLFilenames'.
func ImportQueueStreamsDataByFilename(db XODB, filename string) ([]*ImportQueueStreamsDatum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, data ` +
		`FROM AllergyNew.import_queue_streams_data ` +
		`WHERE filename = ?`

	// run query
	XOLog(sqlstr, filename)
	q, err := db.Query(sqlstr, filename)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ImportQueueStreamsDatum{}
	for q.Next() {
		iqsd := ImportQueueStreamsDatum{
			_exists: true,
		}

		// scan
		err = q.Scan(&iqsd.ID, &iqsd.PracticeID, &iqsd.Filename, &iqsd.Data)
		if err != nil {
			return nil, err
		}

		res = append(res, &iqsd)
	}

	return res, nil
}

// ImportQueueStreamsDatumByID retrieves a row from 'AllergyNew.import_queue_streams_data' as a ImportQueueStreamsDatum.
//
// Generated from index 'import_queue_streams_data_id_pkey'.
func ImportQueueStreamsDatumByID(db XODB, id int) (*ImportQueueStreamsDatum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, data ` +
		`FROM AllergyNew.import_queue_streams_data ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	iqsd := ImportQueueStreamsDatum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&iqsd.ID, &iqsd.PracticeID, &iqsd.Filename, &iqsd.Data)
	if err != nil {
		return nil, err
	}

	return &iqsd, nil
}