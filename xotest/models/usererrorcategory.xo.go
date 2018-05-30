// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// UserErrorCategory represents a row from 'AllergyNew.user_error_categories'.
type UserErrorCategory struct {
	ID        uint           `json:"id"`         // id
	TypeID    sql.NullInt64  `json:"type_id"`    // type_id
	Category  sql.NullString `json:"category"`   // category
	UpdatedBy sql.NullInt64  `json:"updated_by"` // updated_by
	Updated   time.Time      `json:"updated"`    // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserErrorCategory exists in the database.
func (uec *UserErrorCategory) Exists() bool {
	return uec._exists
}

// Deleted provides information if the UserErrorCategory has been deleted from the database.
func (uec *UserErrorCategory) Deleted() bool {
	return uec._deleted
}

// Insert inserts the UserErrorCategory to the database.
func (uec *UserErrorCategory) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if uec._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.user_error_categories (` +
		`type_id, category, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, uec.TypeID, uec.Category, uec.UpdatedBy, uec.Updated)
	res, err := db.Exec(sqlstr, uec.TypeID, uec.Category, uec.UpdatedBy, uec.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	uec.ID = uint(id)
	uec._exists = true

	return nil
}

// Update updates the UserErrorCategory in the database.
func (uec *UserErrorCategory) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uec._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if uec._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.user_error_categories SET ` +
		`type_id = ?, category = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, uec.TypeID, uec.Category, uec.UpdatedBy, uec.Updated, uec.ID)
	_, err = db.Exec(sqlstr, uec.TypeID, uec.Category, uec.UpdatedBy, uec.Updated, uec.ID)
	return err
}

// Save saves the UserErrorCategory to the database.
func (uec *UserErrorCategory) Save(db XODB) error {
	if uec.Exists() {
		return uec.Update(db)
	}

	return uec.Insert(db)
}

// Delete deletes the UserErrorCategory from the database.
func (uec *UserErrorCategory) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uec._exists {
		return nil
	}

	// if deleted, bail
	if uec._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.user_error_categories WHERE id = ?`

	// run query
	XOLog(sqlstr, uec.ID)
	_, err = db.Exec(sqlstr, uec.ID)
	if err != nil {
		return err
	}

	// set deleted
	uec._deleted = true

	return nil
}

// UserErrorCategoryByID retrieves a row from 'AllergyNew.user_error_categories' as a UserErrorCategory.
//
// Generated from index 'user_error_categories_id_pkey'.
func UserErrorCategoryByID(db XODB, id uint) (*UserErrorCategory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, type_id, category, updated_by, updated ` +
		`FROM AllergyNew.user_error_categories ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	uec := UserErrorCategory{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&uec.ID, &uec.TypeID, &uec.Category, &uec.UpdatedBy, &uec.Updated)
	if err != nil {
		return nil, err
	}

	return &uec, nil
}