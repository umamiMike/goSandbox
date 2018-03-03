// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// ImmunotherapyVial represents a row from 'AllergyNew.immunotherapy_vials'.
type ImmunotherapyVial struct {
	ID                      uint           `json:"id"`                        // id
	ImmunotherapyScheduleID uint           `json:"immunotherapy_schedule_id"` // immunotherapy_schedule_id
	VialName                sql.NullString `json:"vial_name"`                 // vial_name
	VialColorName           sql.NullString `json:"vial_color_name"`           // vial_color_name
	VialColorValue          sql.NullString `json:"vial_color_value"`          // vial_color_value
	VialDilution            sql.NullInt64  `json:"vial_dilution"`             // vial_dilution
	UpdatedBy               sql.NullInt64  `json:"updated_by"`                // updated_by
	Updated                 time.Time      `json:"updated"`                   // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ImmunotherapyVial exists in the database.
func (iv *ImmunotherapyVial) Exists() bool {
	return iv._exists
}

// Deleted provides information if the ImmunotherapyVial has been deleted from the database.
func (iv *ImmunotherapyVial) Deleted() bool {
	return iv._deleted
}

// Insert inserts the ImmunotherapyVial to the database.
func (iv *ImmunotherapyVial) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if iv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.immunotherapy_vials (` +
		`immunotherapy_schedule_id, vial_name, vial_color_name, vial_color_value, vial_dilution, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, iv.ImmunotherapyScheduleID, iv.VialName, iv.VialColorName, iv.VialColorValue, iv.VialDilution, iv.UpdatedBy, iv.Updated)
	res, err := db.Exec(sqlstr, iv.ImmunotherapyScheduleID, iv.VialName, iv.VialColorName, iv.VialColorValue, iv.VialDilution, iv.UpdatedBy, iv.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	iv.ID = uint(id)
	iv._exists = true

	return nil
}

// Update updates the ImmunotherapyVial in the database.
func (iv *ImmunotherapyVial) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if iv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.immunotherapy_vials SET ` +
		`immunotherapy_schedule_id = ?, vial_name = ?, vial_color_name = ?, vial_color_value = ?, vial_dilution = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, iv.ImmunotherapyScheduleID, iv.VialName, iv.VialColorName, iv.VialColorValue, iv.VialDilution, iv.UpdatedBy, iv.Updated, iv.ID)
	_, err = db.Exec(sqlstr, iv.ImmunotherapyScheduleID, iv.VialName, iv.VialColorName, iv.VialColorValue, iv.VialDilution, iv.UpdatedBy, iv.Updated, iv.ID)
	return err
}

// Save saves the ImmunotherapyVial to the database.
func (iv *ImmunotherapyVial) Save(db XODB) error {
	if iv.Exists() {
		return iv.Update(db)
	}

	return iv.Insert(db)
}

// Delete deletes the ImmunotherapyVial from the database.
func (iv *ImmunotherapyVial) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !iv._exists {
		return nil
	}

	// if deleted, bail
	if iv._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.immunotherapy_vials WHERE id = ?`

	// run query
	XOLog(sqlstr, iv.ID)
	_, err = db.Exec(sqlstr, iv.ID)
	if err != nil {
		return err
	}

	// set deleted
	iv._deleted = true

	return nil
}

// ImmunotherapySchedule returns the ImmunotherapySchedule associated with the ImmunotherapyVial's ImmunotherapyScheduleID (immunotherapy_schedule_id).
//
// Generated from foreign key 'fk_vial_schedule_id'.
func (iv *ImmunotherapyVial) ImmunotherapySchedule(db XODB) (*ImmunotherapySchedule, error) {
	return ImmunotherapyScheduleByID(db, iv.ImmunotherapyScheduleID)
}

// ImmunotherapyVialsByImmunotherapyScheduleID retrieves a row from 'AllergyNew.immunotherapy_vials' as a ImmunotherapyVial.
//
// Generated from index 'fk_vial_schedule_id'.
func ImmunotherapyVialsByImmunotherapyScheduleID(db XODB, immunotherapyScheduleID uint) ([]*ImmunotherapyVial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, immunotherapy_schedule_id, vial_name, vial_color_name, vial_color_value, vial_dilution, updated_by, updated ` +
		`FROM AllergyNew.immunotherapy_vials ` +
		`WHERE immunotherapy_schedule_id = ?`

	// run query
	XOLog(sqlstr, immunotherapyScheduleID)
	q, err := db.Query(sqlstr, immunotherapyScheduleID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ImmunotherapyVial{}
	for q.Next() {
		iv := ImmunotherapyVial{
			_exists: true,
		}

		// scan
		err = q.Scan(&iv.ID, &iv.ImmunotherapyScheduleID, &iv.VialName, &iv.VialColorName, &iv.VialColorValue, &iv.VialDilution, &iv.UpdatedBy, &iv.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &iv)
	}

	return res, nil
}

// ImmunotherapyVialByID retrieves a row from 'AllergyNew.immunotherapy_vials' as a ImmunotherapyVial.
//
// Generated from index 'immunotherapy_vials_id_pkey'.
func ImmunotherapyVialByID(db XODB, id uint) (*ImmunotherapyVial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, immunotherapy_schedule_id, vial_name, vial_color_name, vial_color_value, vial_dilution, updated_by, updated ` +
		`FROM AllergyNew.immunotherapy_vials ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	iv := ImmunotherapyVial{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&iv.ID, &iv.ImmunotherapyScheduleID, &iv.VialName, &iv.VialColorName, &iv.VialColorValue, &iv.VialDilution, &iv.UpdatedBy, &iv.Updated)
	if err != nil {
		return nil, err
	}

	return &iv, nil
}
