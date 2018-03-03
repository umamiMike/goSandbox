// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// AntigenVialSize represents a row from 'AllergyNew.antigen_vial_sizes'.
type AntigenVialSize struct {
	ID        uint           `json:"id"`         // id
	Size      sql.NullString `json:"size"`       // size
	Order     sql.NullInt64  `json:"order"`      // order
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AntigenVialSize exists in the database.
func (avs *AntigenVialSize) Exists() bool {
	return avs._exists
}

// Deleted provides information if the AntigenVialSize has been deleted from the database.
func (avs *AntigenVialSize) Deleted() bool {
	return avs._deleted
}

// Insert inserts the AntigenVialSize to the database.
func (avs *AntigenVialSize) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if avs._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.antigen_vial_sizes (` +
		`size, order, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, avs.Size, avs.Order, avs.UpdatedBy, avs.Updated)
	res, err := db.Exec(sqlstr, avs.Size, avs.Order, avs.UpdatedBy, avs.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	avs.ID = uint(id)
	avs._exists = true

	return nil
}

// Update updates the AntigenVialSize in the database.
func (avs *AntigenVialSize) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !avs._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if avs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.antigen_vial_sizes SET ` +
		`size = ?, order = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, avs.Size, avs.Order, avs.UpdatedBy, avs.Updated, avs.ID)
	_, err = db.Exec(sqlstr, avs.Size, avs.Order, avs.UpdatedBy, avs.Updated, avs.ID)
	return err
}

// Save saves the AntigenVialSize to the database.
func (avs *AntigenVialSize) Save(db XODB) error {
	if avs.Exists() {
		return avs.Update(db)
	}

	return avs.Insert(db)
}

// Delete deletes the AntigenVialSize from the database.
func (avs *AntigenVialSize) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !avs._exists {
		return nil
	}

	// if deleted, bail
	if avs._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.antigen_vial_sizes WHERE id = ?`

	// run query
	XOLog(sqlstr, avs.ID)
	_, err = db.Exec(sqlstr, avs.ID)
	if err != nil {
		return err
	}

	// set deleted
	avs._deleted = true

	return nil
}

// User returns the User associated with the AntigenVialSize's UpdatedBy (updated_by).
//
// Generated from foreign key 'antigen_vial_sizes_updated'.
func (avs *AntigenVialSize) User(db XODB) (*User, error) {
	return UserByID(db, uint(avs.UpdatedBy.Int64))
}

// AntigenVialSizeByID retrieves a row from 'AllergyNew.antigen_vial_sizes' as a AntigenVialSize.
//
// Generated from index 'antigen_vial_sizes_id_pkey'.
func AntigenVialSizeByID(db XODB, id uint) (*AntigenVialSize, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, size, order, updated_by, updated ` +
		`FROM AllergyNew.antigen_vial_sizes ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	avs := AntigenVialSize{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&avs.ID, &avs.Size, &avs.Order, &avs.UpdatedBy, &avs.Updated)
	if err != nil {
		return nil, err
	}

	return &avs, nil
}

// AntigenVialSizesByUpdatedBy retrieves a row from 'AllergyNew.antigen_vial_sizes' as a AntigenVialSize.
//
// Generated from index 'antigen_vial_sizes_updated'.
func AntigenVialSizesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*AntigenVialSize, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, size, order, updated_by, updated ` +
		`FROM AllergyNew.antigen_vial_sizes ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenVialSize{}
	for q.Next() {
		avs := AntigenVialSize{
			_exists: true,
		}

		// scan
		err = q.Scan(&avs.ID, &avs.Size, &avs.Order, &avs.UpdatedBy, &avs.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &avs)
	}

	return res, nil
}
