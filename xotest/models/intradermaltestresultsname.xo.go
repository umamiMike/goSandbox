// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// IntradermalTestResultsName represents a row from 'AllergyNew.intradermal_test_results_names'.
type IntradermalTestResultsName struct {
	ID        uint           `json:"id"`         // id
	ResultsID sql.NullInt64  `json:"results_id"` // results_id
	Name      sql.NullString `json:"name"`       // name
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the IntradermalTestResultsName exists in the database.
func (itrn *IntradermalTestResultsName) Exists() bool {
	return itrn._exists
}

// Deleted provides information if the IntradermalTestResultsName has been deleted from the database.
func (itrn *IntradermalTestResultsName) Deleted() bool {
	return itrn._deleted
}

// Insert inserts the IntradermalTestResultsName to the database.
func (itrn *IntradermalTestResultsName) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if itrn._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.intradermal_test_results_names (` +
		`results_id, name, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, itrn.ResultsID, itrn.Name, itrn.UpdatedBy, itrn.Updated)
	res, err := db.Exec(sqlstr, itrn.ResultsID, itrn.Name, itrn.UpdatedBy, itrn.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	itrn.ID = uint(id)
	itrn._exists = true

	return nil
}

// Update updates the IntradermalTestResultsName in the database.
func (itrn *IntradermalTestResultsName) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itrn._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if itrn._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.intradermal_test_results_names SET ` +
		`results_id = ?, name = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, itrn.ResultsID, itrn.Name, itrn.UpdatedBy, itrn.Updated, itrn.ID)
	_, err = db.Exec(sqlstr, itrn.ResultsID, itrn.Name, itrn.UpdatedBy, itrn.Updated, itrn.ID)
	return err
}

// Save saves the IntradermalTestResultsName to the database.
func (itrn *IntradermalTestResultsName) Save(db XODB) error {
	if itrn.Exists() {
		return itrn.Update(db)
	}

	return itrn.Insert(db)
}

// Delete deletes the IntradermalTestResultsName from the database.
func (itrn *IntradermalTestResultsName) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itrn._exists {
		return nil
	}

	// if deleted, bail
	if itrn._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.intradermal_test_results_names WHERE id = ?`

	// run query
	XOLog(sqlstr, itrn.ID)
	_, err = db.Exec(sqlstr, itrn.ID)
	if err != nil {
		return err
	}

	// set deleted
	itrn._deleted = true

	return nil
}

// IntradermalTestResultsNamesByResultsID retrieves a row from 'AllergyNew.intradermal_test_results_names' as a IntradermalTestResultsName.
//
// Generated from index 'Results'.
func IntradermalTestResultsNamesByResultsID(db XODB, resultsID sql.NullInt64) ([]*IntradermalTestResultsName, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, results_id, name, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results_names ` +
		`WHERE results_id = ?`

	// run query
	XOLog(sqlstr, resultsID)
	q, err := db.Query(sqlstr, resultsID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestResultsName{}
	for q.Next() {
		itrn := IntradermalTestResultsName{
			_exists: true,
		}

		// scan
		err = q.Scan(&itrn.ID, &itrn.ResultsID, &itrn.Name, &itrn.UpdatedBy, &itrn.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itrn)
	}

	return res, nil
}

// IntradermalTestResultsNameByID retrieves a row from 'AllergyNew.intradermal_test_results_names' as a IntradermalTestResultsName.
//
// Generated from index 'intradermal_test_results_names_id_pkey'.
func IntradermalTestResultsNameByID(db XODB, id uint) (*IntradermalTestResultsName, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, results_id, name, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_results_names ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	itrn := IntradermalTestResultsName{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&itrn.ID, &itrn.ResultsID, &itrn.Name, &itrn.UpdatedBy, &itrn.Updated)
	if err != nil {
		return nil, err
	}

	return &itrn, nil
}