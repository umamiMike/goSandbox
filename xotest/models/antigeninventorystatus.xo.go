// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// AntigenInventoryStatus represents a row from 'AllergyNew.antigen_inventory_statuses'.
type AntigenInventoryStatus struct {
	ID        uint           `json:"id"`         // id
	Status    sql.NullString `json:"status"`     // status
	Order     sql.NullInt64  `json:"order"`      // order
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AntigenInventoryStatus exists in the database.
func (ais *AntigenInventoryStatus) Exists() bool {
	return ais._exists
}

// Deleted provides information if the AntigenInventoryStatus has been deleted from the database.
func (ais *AntigenInventoryStatus) Deleted() bool {
	return ais._deleted
}

// Insert inserts the AntigenInventoryStatus to the database.
func (ais *AntigenInventoryStatus) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ais._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.antigen_inventory_statuses (` +
		`status, order, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ais.Status, ais.Order, ais.UpdatedBy, ais.Updated)
	res, err := db.Exec(sqlstr, ais.Status, ais.Order, ais.UpdatedBy, ais.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ais.ID = uint(id)
	ais._exists = true

	return nil
}

// Update updates the AntigenInventoryStatus in the database.
func (ais *AntigenInventoryStatus) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ais._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ais._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.antigen_inventory_statuses SET ` +
		`status = ?, order = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ais.Status, ais.Order, ais.UpdatedBy, ais.Updated, ais.ID)
	_, err = db.Exec(sqlstr, ais.Status, ais.Order, ais.UpdatedBy, ais.Updated, ais.ID)
	return err
}

// Save saves the AntigenInventoryStatus to the database.
func (ais *AntigenInventoryStatus) Save(db XODB) error {
	if ais.Exists() {
		return ais.Update(db)
	}

	return ais.Insert(db)
}

// Delete deletes the AntigenInventoryStatus from the database.
func (ais *AntigenInventoryStatus) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ais._exists {
		return nil
	}

	// if deleted, bail
	if ais._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.antigen_inventory_statuses WHERE id = ?`

	// run query
	XOLog(sqlstr, ais.ID)
	_, err = db.Exec(sqlstr, ais.ID)
	if err != nil {
		return err
	}

	// set deleted
	ais._deleted = true

	return nil
}

// User returns the User associated with the AntigenInventoryStatus's UpdatedBy (updated_by).
//
// Generated from foreign key 'antigen_inventory_statuses_updated'.
func (ais *AntigenInventoryStatus) User(db XODB) (*User, error) {
	return UserByID(db, uint(ais.UpdatedBy.Int64))
}

// AntigenInventoryStatusByID retrieves a row from 'AllergyNew.antigen_inventory_statuses' as a AntigenInventoryStatus.
//
// Generated from index 'antigen_inventory_statuses_id_pkey'.
func AntigenInventoryStatusByID(db XODB, id uint) (*AntigenInventoryStatus, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, status, order, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory_statuses ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ais := AntigenInventoryStatus{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ais.ID, &ais.Status, &ais.Order, &ais.UpdatedBy, &ais.Updated)
	if err != nil {
		return nil, err
	}

	return &ais, nil
}

// AntigenInventoryStatusesByUpdatedBy retrieves a row from 'AllergyNew.antigen_inventory_statuses' as a AntigenInventoryStatus.
//
// Generated from index 'antigen_inventory_statuses_updated'.
func AntigenInventoryStatusesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*AntigenInventoryStatus, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, status, order, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory_statuses ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventoryStatus{}
	for q.Next() {
		ais := AntigenInventoryStatus{
			_exists: true,
		}

		// scan
		err = q.Scan(&ais.ID, &ais.Status, &ais.Order, &ais.UpdatedBy, &ais.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ais)
	}

	return res, nil
}
