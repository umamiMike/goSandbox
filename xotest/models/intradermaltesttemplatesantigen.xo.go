// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// IntradermalTestTemplatesAntigen represents a row from 'AllergyNew.intradermal_test_templates_antigens'.
type IntradermalTestTemplatesAntigen struct {
	ID                       uint          `json:"id"`                          // id
	TemplateID               sql.NullInt64 `json:"template_id"`                 // template_id
	AntigenID                sql.NullInt64 `json:"antigen_id"`                  // antigen_id
	DefaultEndpointWhealSize sql.NullInt64 `json:"default_endpoint_wheal_size"` // default_endpoint_wheal_size
	Order                    sql.NullInt64 `json:"order"`                       // order
	IsBreak                  int8          `json:"is_break"`                    // is_break
	UpdatedBy                sql.NullInt64 `json:"updated_by"`                  // updated_by
	Updated                  time.Time     `json:"updated"`                     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the IntradermalTestTemplatesAntigen exists in the database.
func (itta *IntradermalTestTemplatesAntigen) Exists() bool {
	return itta._exists
}

// Deleted provides information if the IntradermalTestTemplatesAntigen has been deleted from the database.
func (itta *IntradermalTestTemplatesAntigen) Deleted() bool {
	return itta._deleted
}

// Insert inserts the IntradermalTestTemplatesAntigen to the database.
func (itta *IntradermalTestTemplatesAntigen) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if itta._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.intradermal_test_templates_antigens (` +
		`template_id, antigen_id, default_endpoint_wheal_size, order, is_break, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, itta.TemplateID, itta.AntigenID, itta.DefaultEndpointWhealSize, itta.Order, itta.IsBreak, itta.UpdatedBy, itta.Updated)
	res, err := db.Exec(sqlstr, itta.TemplateID, itta.AntigenID, itta.DefaultEndpointWhealSize, itta.Order, itta.IsBreak, itta.UpdatedBy, itta.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	itta.ID = uint(id)
	itta._exists = true

	return nil
}

// Update updates the IntradermalTestTemplatesAntigen in the database.
func (itta *IntradermalTestTemplatesAntigen) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itta._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if itta._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.intradermal_test_templates_antigens SET ` +
		`template_id = ?, antigen_id = ?, default_endpoint_wheal_size = ?, order = ?, is_break = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, itta.TemplateID, itta.AntigenID, itta.DefaultEndpointWhealSize, itta.Order, itta.IsBreak, itta.UpdatedBy, itta.Updated, itta.ID)
	_, err = db.Exec(sqlstr, itta.TemplateID, itta.AntigenID, itta.DefaultEndpointWhealSize, itta.Order, itta.IsBreak, itta.UpdatedBy, itta.Updated, itta.ID)
	return err
}

// Save saves the IntradermalTestTemplatesAntigen to the database.
func (itta *IntradermalTestTemplatesAntigen) Save(db XODB) error {
	if itta.Exists() {
		return itta.Update(db)
	}

	return itta.Insert(db)
}

// Delete deletes the IntradermalTestTemplatesAntigen from the database.
func (itta *IntradermalTestTemplatesAntigen) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !itta._exists {
		return nil
	}

	// if deleted, bail
	if itta._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.intradermal_test_templates_antigens WHERE id = ?`

	// run query
	XOLog(sqlstr, itta.ID)
	_, err = db.Exec(sqlstr, itta.ID)
	if err != nil {
		return err
	}

	// set deleted
	itta._deleted = true

	return nil
}

// Antigen returns the Antigen associated with the IntradermalTestTemplatesAntigen's AntigenID (antigen_id).
//
// Generated from foreign key 'intradermal_test_templates_antigens_antigen'.
func (itta *IntradermalTestTemplatesAntigen) Antigen(db XODB) (*Antigen, error) {
	return AntigenByID(db, uint(itta.AntigenID.Int64))
}

// IntradermalTestTemplate returns the IntradermalTestTemplate associated with the IntradermalTestTemplatesAntigen's TemplateID (template_id).
//
// Generated from foreign key 'intradermal_test_templates_antigens_template'.
func (itta *IntradermalTestTemplatesAntigen) IntradermalTestTemplate(db XODB) (*IntradermalTestTemplate, error) {
	return IntradermalTestTemplateByID(db, uint(itta.TemplateID.Int64))
}

// User returns the User associated with the IntradermalTestTemplatesAntigen's UpdatedBy (updated_by).
//
// Generated from foreign key 'intradermal_test_templates_antigens_updated'.
func (itta *IntradermalTestTemplatesAntigen) User(db XODB) (*User, error) {
	return UserByID(db, uint(itta.UpdatedBy.Int64))
}

// IntradermalTestTemplatesAntigensByTemplateID retrieves a row from 'AllergyNew.intradermal_test_templates_antigens' as a IntradermalTestTemplatesAntigen.
//
// Generated from index 'Templates'.
func IntradermalTestTemplatesAntigensByTemplateID(db XODB, templateID sql.NullInt64) ([]*IntradermalTestTemplatesAntigen, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_id, antigen_id, default_endpoint_wheal_size, order, is_break, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_templates_antigens ` +
		`WHERE template_id = ?`

	// run query
	XOLog(sqlstr, templateID)
	q, err := db.Query(sqlstr, templateID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestTemplatesAntigen{}
	for q.Next() {
		itta := IntradermalTestTemplatesAntigen{
			_exists: true,
		}

		// scan
		err = q.Scan(&itta.ID, &itta.TemplateID, &itta.AntigenID, &itta.DefaultEndpointWhealSize, &itta.Order, &itta.IsBreak, &itta.UpdatedBy, &itta.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itta)
	}

	return res, nil
}

// IntradermalTestTemplatesAntigensByAntigenID retrieves a row from 'AllergyNew.intradermal_test_templates_antigens' as a IntradermalTestTemplatesAntigen.
//
// Generated from index 'intradermal_test_templates_antigens_antigen'.
func IntradermalTestTemplatesAntigensByAntigenID(db XODB, antigenID sql.NullInt64) ([]*IntradermalTestTemplatesAntigen, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_id, antigen_id, default_endpoint_wheal_size, order, is_break, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_templates_antigens ` +
		`WHERE antigen_id = ?`

	// run query
	XOLog(sqlstr, antigenID)
	q, err := db.Query(sqlstr, antigenID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestTemplatesAntigen{}
	for q.Next() {
		itta := IntradermalTestTemplatesAntigen{
			_exists: true,
		}

		// scan
		err = q.Scan(&itta.ID, &itta.TemplateID, &itta.AntigenID, &itta.DefaultEndpointWhealSize, &itta.Order, &itta.IsBreak, &itta.UpdatedBy, &itta.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itta)
	}

	return res, nil
}

// IntradermalTestTemplatesAntigenByID retrieves a row from 'AllergyNew.intradermal_test_templates_antigens' as a IntradermalTestTemplatesAntigen.
//
// Generated from index 'intradermal_test_templates_antigens_id_pkey'.
func IntradermalTestTemplatesAntigenByID(db XODB, id uint) (*IntradermalTestTemplatesAntigen, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_id, antigen_id, default_endpoint_wheal_size, order, is_break, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_templates_antigens ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	itta := IntradermalTestTemplatesAntigen{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&itta.ID, &itta.TemplateID, &itta.AntigenID, &itta.DefaultEndpointWhealSize, &itta.Order, &itta.IsBreak, &itta.UpdatedBy, &itta.Updated)
	if err != nil {
		return nil, err
	}

	return &itta, nil
}

// IntradermalTestTemplatesAntigensByUpdatedBy retrieves a row from 'AllergyNew.intradermal_test_templates_antigens' as a IntradermalTestTemplatesAntigen.
//
// Generated from index 'intradermal_test_templates_antigens_updated'.
func IntradermalTestTemplatesAntigensByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*IntradermalTestTemplatesAntigen, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_id, antigen_id, default_endpoint_wheal_size, order, is_break, updated_by, updated ` +
		`FROM AllergyNew.intradermal_test_templates_antigens ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IntradermalTestTemplatesAntigen{}
	for q.Next() {
		itta := IntradermalTestTemplatesAntigen{
			_exists: true,
		}

		// scan
		err = q.Scan(&itta.ID, &itta.TemplateID, &itta.AntigenID, &itta.DefaultEndpointWhealSize, &itta.Order, &itta.IsBreak, &itta.UpdatedBy, &itta.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &itta)
	}

	return res, nil
}