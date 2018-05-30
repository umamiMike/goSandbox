// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// ImportQueueStream represents a row from 'AllergyNew.import_queue_streams'.
type ImportQueueStream struct {
	ID         int    `json:"id"`          // id
	PracticeID int    `json:"practice_id"` // practice_id
	Filename   string `json:"filename"`    // filename

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ImportQueueStream exists in the database.
func (iqs *ImportQueueStream) Exists() bool {
	return iqs._exists
}

// Deleted provides information if the ImportQueueStream has been deleted from the database.
func (iqs *ImportQueueStream) Deleted() bool {
	return iqs._deleted
}

// Insert inserts the ImportQueueStream to the database.
func (iqs *ImportQueueStream) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if iqs._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.import_queue_streams (` +
		`practice_id, filename` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, iqs.PracticeID, iqs.Filename)
	res, err := db.Exec(sqlstr, iqs.PracticeID, iqs.Filename)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	iqs.ID = int(id)
	iqs._exists = true

	return nil
}

// Update updates the ImportQueueStream in the database.
func (iqs *ImportQueueStream) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iqs._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if iqs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.import_queue_streams SET ` +
		`practice_id = ?, filename = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, iqs.PracticeID, iqs.Filename, iqs.ID)
	_, err = db.Exec(sqlstr, iqs.PracticeID, iqs.Filename, iqs.ID)
	return err
}

// Save saves the ImportQueueStream to the database.
func (iqs *ImportQueueStream) Save(db XODB) error {
	if iqs.Exists() {
		return iqs.Update(db)
	}

	return iqs.Insert(db)
}

// Delete deletes the ImportQueueStream from the database.
func (iqs *ImportQueueStream) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iqs._exists {
		return nil
	}

	// if deleted, bail
	if iqs._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.import_queue_streams WHERE id = ?`

	// run query
	XOLog(sqlstr, iqs.ID)
	_, err = db.Exec(sqlstr, iqs.ID)
	if err != nil {
		return err
	}

	// set deleted
	iqs._deleted = true

	return nil
}

// ImportQueueStreamsByFilename retrieves a row from 'AllergyNew.import_queue_streams' as a ImportQueueStream.
//
// Generated from index 'XMLFilenames'.
func ImportQueueStreamsByFilename(db XODB, filename string) ([]*ImportQueueStream, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename ` +
		`FROM AllergyNew.import_queue_streams ` +
		`WHERE filename = ?`

	// run query
	XOLog(sqlstr, filename)
	q, err := db.Query(sqlstr, filename)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ImportQueueStream{}
	for q.Next() {
		iqs := ImportQueueStream{
			_exists: true,
		}

		// scan
		err = q.Scan(&iqs.ID, &iqs.PracticeID, &iqs.Filename)
		if err != nil {
			return nil, err
		}

		res = append(res, &iqs)
	}

	return res, nil
}

// ImportQueueStreamByID retrieves a row from 'AllergyNew.import_queue_streams' as a ImportQueueStream.
//
// Generated from index 'import_queue_streams_id_pkey'.
func ImportQueueStreamByID(db XODB, id int) (*ImportQueueStream, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename ` +
		`FROM AllergyNew.import_queue_streams ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	iqs := ImportQueueStream{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&iqs.ID, &iqs.PracticeID, &iqs.Filename)
	if err != nil {
		return nil, err
	}

	return &iqs, nil
}