// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// ImportQueue represents a row from 'AllergyNew.import_queue'.
type ImportQueue struct {
	ID          int            `json:"id"`           // id
	PracticeID  int            `json:"practice_id"`  // practice_id
	Filename    sql.NullString `json:"filename"`     // filename
	Type        string         `json:"type"`         // type
	Outer       string         `json:"outer"`        // outer
	XML         []byte         `json:"xml"`          // xml
	Synchronous bool           `json:"synchronous"`  // synchronous
	BatchID     string         `json:"batch_id"`     // batch_id
	ToggleAsync bool           `json:"toggle_async"` // toggle_async
	Processing  bool           `json:"processing"`   // processing

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ImportQueue exists in the database.
func (iq *ImportQueue) Exists() bool {
	return iq._exists
}

// Deleted provides information if the ImportQueue has been deleted from the database.
func (iq *ImportQueue) Deleted() bool {
	return iq._deleted
}

// Insert inserts the ImportQueue to the database.
func (iq *ImportQueue) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if iq._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.import_queue (` +
		`practice_id, filename, type, outer, xml, synchronous, batch_id, toggle_async, processing` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, iq.PracticeID, iq.Filename, iq.Type, iq.Outer, iq.XML, iq.Synchronous, iq.BatchID, iq.ToggleAsync, iq.Processing)
	res, err := db.Exec(sqlstr, iq.PracticeID, iq.Filename, iq.Type, iq.Outer, iq.XML, iq.Synchronous, iq.BatchID, iq.ToggleAsync, iq.Processing)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	iq.ID = int(id)
	iq._exists = true

	return nil
}

// Update updates the ImportQueue in the database.
func (iq *ImportQueue) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iq._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if iq._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.import_queue SET ` +
		`practice_id = ?, filename = ?, type = ?, outer = ?, xml = ?, synchronous = ?, batch_id = ?, toggle_async = ?, processing = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, iq.PracticeID, iq.Filename, iq.Type, iq.Outer, iq.XML, iq.Synchronous, iq.BatchID, iq.ToggleAsync, iq.Processing, iq.ID)
	_, err = db.Exec(sqlstr, iq.PracticeID, iq.Filename, iq.Type, iq.Outer, iq.XML, iq.Synchronous, iq.BatchID, iq.ToggleAsync, iq.Processing, iq.ID)
	return err
}

// Save saves the ImportQueue to the database.
func (iq *ImportQueue) Save(db XODB) error {
	if iq.Exists() {
		return iq.Update(db)
	}

	return iq.Insert(db)
}

// Delete deletes the ImportQueue from the database.
func (iq *ImportQueue) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iq._exists {
		return nil
	}

	// if deleted, bail
	if iq._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.import_queue WHERE id = ?`

	// run query
	XOLog(sqlstr, iq.ID)
	_, err = db.Exec(sqlstr, iq.ID)
	if err != nil {
		return err
	}

	// set deleted
	iq._deleted = true

	return nil
}

// Importqueue represents a row from 'AllergyNew.importQueue'.
type Importqueue struct {
	ID          int    `json:"id"`           // id
	PracticeID  int    `json:"practice_id"`  // practice_id
	Filename    string `json:"filename"`     // filename
	Type        string `json:"type"`         // type
	Outer       string `json:"outer"`        // outer
	Data        []byte `json:"data"`         // data
	Synchronous bool   `json:"synchronous"`  // synchronous
	BatchID     string `json:"batch_id"`     // batch_id
	ToggleAsync bool   `json:"toggle_async"` // toggle_async
	Processing  bool   `json:"processing"`   // processing

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Importqueue exists in the database.
func (i *Importqueue) Exists() bool {
	return i._exists
}

// Deleted provides information if the Importqueue has been deleted from the database.
func (i *Importqueue) Deleted() bool {
	return i._deleted
}

// Insert inserts the Importqueue to the database.
func (i *Importqueue) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if i._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO AllergyNew.importQueue (` +
		`id, practice_id, filename, type, outer, data, synchronous, batch_id, toggle_async, processing` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, i.ID, i.PracticeID, i.Filename, i.Type, i.Outer, i.Data, i.Synchronous, i.BatchID, i.ToggleAsync, i.Processing)
	_, err = db.Exec(sqlstr, i.ID, i.PracticeID, i.Filename, i.Type, i.Outer, i.Data, i.Synchronous, i.BatchID, i.ToggleAsync, i.Processing)
	if err != nil {
		return err
	}

	// set existence
	i._exists = true

	return nil
}

// Update updates the Importqueue in the database.
func (i *Importqueue) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if i._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.importQueue SET ` +
		`practice_id = ?, filename = ?, type = ?, outer = ?, data = ?, synchronous = ?, batch_id = ?, toggle_async = ?, processing = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, i.PracticeID, i.Filename, i.Type, i.Outer, i.Data, i.Synchronous, i.BatchID, i.ToggleAsync, i.Processing, i.ID)
	_, err = db.Exec(sqlstr, i.PracticeID, i.Filename, i.Type, i.Outer, i.Data, i.Synchronous, i.BatchID, i.ToggleAsync, i.Processing, i.ID)
	return err
}

// Save saves the Importqueue to the database.
func (i *Importqueue) Save(db XODB) error {
	if i.Exists() {
		return i.Update(db)
	}

	return i.Insert(db)
}

// Delete deletes the Importqueue from the database.
func (i *Importqueue) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return nil
	}

	// if deleted, bail
	if i._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.importQueue WHERE id = ?`

	// run query
	XOLog(sqlstr, i.ID)
	_, err = db.Exec(sqlstr, i.ID)
	if err != nil {
		return err
	}

	// set deleted
	i._deleted = true

	return nil
}

// ImportQueuesByFilename retrieves a row from 'AllergyNew.import_queue' as a ImportQueue.
//
// Generated from index 'XMLFilenames'.
func ImportQueuesByFilename(db XODB, filename sql.NullString) ([]*ImportQueue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, type, outer, xml, synchronous, batch_id, toggle_async, processing ` +
		`FROM AllergyNew.import_queue ` +
		`WHERE filename = ?`

	// run query
	XOLog(sqlstr, filename)
	q, err := db.Query(sqlstr, filename)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ImportQueue{}
	for q.Next() {
		iq := ImportQueue{
			_exists: true,
		}

		// scan
		err = q.Scan(&iq.ID, &iq.PracticeID, &iq.Filename, &iq.Type, &iq.Outer, &iq.XML, &iq.Synchronous, &iq.BatchID, &iq.ToggleAsync, &iq.Processing)
		if err != nil {
			return nil, err
		}

		res = append(res, &iq)
	}

	return res, nil
}

// ImportQueueByID retrieves a row from 'AllergyNew.import_queue' as a ImportQueue.
//
// Generated from index 'import_queue_id_pkey'.
func ImportQueueByID(db XODB, id int) (*ImportQueue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, type, outer, xml, synchronous, batch_id, toggle_async, processing ` +
		`FROM AllergyNew.import_queue ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	iq := ImportQueue{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&iq.ID, &iq.PracticeID, &iq.Filename, &iq.Type, &iq.Outer, &iq.XML, &iq.Synchronous, &iq.BatchID, &iq.ToggleAsync, &iq.Processing)
	if err != nil {
		return nil, err
	}

	return &iq, nil
}

// ImportqueuesByFilename retrieves a row from 'AllergyNew.importQueue' as a Importqueue.
//
// Generated from index 'XMLFilenames'.
func ImportqueuesByFilename(db XODB, filename string) ([]*Importqueue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, type, outer, data, synchronous, batch_id, toggle_async, processing ` +
		`FROM AllergyNew.importQueue ` +
		`WHERE filename = ?`

	// run query
	XOLog(sqlstr, filename)
	q, err := db.Query(sqlstr, filename)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Importqueue{}
	for q.Next() {
		i := Importqueue{
			_exists: true,
		}

		// scan
		err = q.Scan(&i.ID, &i.PracticeID, &i.Filename, &i.Type, &i.Outer, &i.Data, &i.Synchronous, &i.BatchID, &i.ToggleAsync, &i.Processing)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}

// ImportqueueByID retrieves a row from 'AllergyNew.importQueue' as a Importqueue.
//
// Generated from index 'importQueue_id_pkey'.
func ImportqueueByID(db XODB, id int) (*Importqueue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, practice_id, filename, type, outer, data, synchronous, batch_id, toggle_async, processing ` +
		`FROM AllergyNew.importQueue ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	i := Importqueue{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&i.ID, &i.PracticeID, &i.Filename, &i.Type, &i.Outer, &i.Data, &i.Synchronous, &i.BatchID, &i.ToggleAsync, &i.Processing)
	if err != nil {
		return nil, err
	}

	return &i, nil
}