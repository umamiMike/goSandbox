// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// FormulaArchive represents a row from 'AllergyNew.formula_archive'.
type FormulaArchive struct {
	ID        uint           `json:"id"`         // id
	FormulaID sql.NullInt64  `json:"formula_id"` // formula_id
	User      sql.NullInt64  `json:"user"`       // user
	Date      mysql.NullTime `json:"date"`       // date
	Action    Action         `json:"action"`     // action
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FormulaArchive exists in the database.
func (fa *FormulaArchive) Exists() bool {
	return fa._exists
}

// Deleted provides information if the FormulaArchive has been deleted from the database.
func (fa *FormulaArchive) Deleted() bool {
	return fa._deleted
}

// Insert inserts the FormulaArchive to the database.
func (fa *FormulaArchive) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if fa._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.formula_archive (` +
		`formula_id, user, date, action, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, fa.FormulaID, fa.User, fa.Date, fa.Action, fa.UpdatedBy, fa.Updated)
	res, err := db.Exec(sqlstr, fa.FormulaID, fa.User, fa.Date, fa.Action, fa.UpdatedBy, fa.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	fa.ID = uint(id)
	fa._exists = true

	return nil
}

// Update updates the FormulaArchive in the database.
func (fa *FormulaArchive) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !fa._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if fa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.formula_archive SET ` +
		`formula_id = ?, user = ?, date = ?, action = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, fa.FormulaID, fa.User, fa.Date, fa.Action, fa.UpdatedBy, fa.Updated, fa.ID)
	_, err = db.Exec(sqlstr, fa.FormulaID, fa.User, fa.Date, fa.Action, fa.UpdatedBy, fa.Updated, fa.ID)
	return err
}

// Save saves the FormulaArchive to the database.
func (fa *FormulaArchive) Save(db XODB) error {
	if fa.Exists() {
		return fa.Update(db)
	}

	return fa.Insert(db)
}

// Delete deletes the FormulaArchive from the database.
func (fa *FormulaArchive) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !fa._exists {
		return nil
	}

	// if deleted, bail
	if fa._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.formula_archive WHERE id = ?`

	// run query
	XOLog(sqlstr, fa.ID)
	_, err = db.Exec(sqlstr, fa.ID)
	if err != nil {
		return err
	}

	// set deleted
	fa._deleted = true

	return nil
}

// FormulaArcByFormulaID retrieves a row from 'AllergyNew.formula_archive' as a FormulaArchive.
//
// Generated from index 'Formulas'.
func FormulaArcByFormulaID(db XODB, formulaID sql.NullInt64) ([]*FormulaArchive, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, formula_id, user, date, action, updated_by, updated ` +
		`FROM AllergyNew.formula_archive ` +
		`WHERE formula_id = ?`

	// run query
	XOLog(sqlstr, formulaID)
	q, err := db.Query(sqlstr, formulaID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*FormulaArchive{}
	for q.Next() {
		fa := FormulaArchive{
			_exists: true,
		}

		// scan
		err = q.Scan(&fa.ID, &fa.FormulaID, &fa.User, &fa.Date, &fa.Action, &fa.UpdatedBy, &fa.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &fa)
	}

	return res, nil
}

// FormulaArcByUser retrieves a row from 'AllergyNew.formula_archive' as a FormulaArchive.
//
// Generated from index 'Users'.
func FormulaArcByUser(db XODB, user sql.NullInt64) ([]*FormulaArchive, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, formula_id, user, date, action, updated_by, updated ` +
		`FROM AllergyNew.formula_archive ` +
		`WHERE user = ?`

	// run query
	XOLog(sqlstr, user)
	q, err := db.Query(sqlstr, user)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*FormulaArchive{}
	for q.Next() {
		fa := FormulaArchive{
			_exists: true,
		}

		// scan
		err = q.Scan(&fa.ID, &fa.FormulaID, &fa.User, &fa.Date, &fa.Action, &fa.UpdatedBy, &fa.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &fa)
	}

	return res, nil
}

// FormulaArchiveByID retrieves a row from 'AllergyNew.formula_archive' as a FormulaArchive.
//
// Generated from index 'formula_archive_id_pkey'.
func FormulaArchiveByID(db XODB, id uint) (*FormulaArchive, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, formula_id, user, date, action, updated_by, updated ` +
		`FROM AllergyNew.formula_archive ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	fa := FormulaArchive{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&fa.ID, &fa.FormulaID, &fa.User, &fa.Date, &fa.Action, &fa.UpdatedBy, &fa.Updated)
	if err != nil {
		return nil, err
	}

	return &fa, nil
}
