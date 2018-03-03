// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// VialTestsLocalReactionsSymptom represents a row from 'AllergyNew.vial_tests_local_reactions_symptoms'.
type VialTestsLocalReactionsSymptom struct {
	ID         uint          `json:"id"`          // id
	ReactionID sql.NullInt64 `json:"reaction_id"` // reaction_id
	SymptomID  sql.NullInt64 `json:"symptom_id"`  // symptom_id
	UpdatedBy  sql.NullInt64 `json:"updated_by"`  // updated_by
	Updated    time.Time     `json:"updated"`     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the VialTestsLocalReactionsSymptom exists in the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Exists() bool {
	return vtlrs._exists
}

// Deleted provides information if the VialTestsLocalReactionsSymptom has been deleted from the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Deleted() bool {
	return vtlrs._deleted
}

// Insert inserts the VialTestsLocalReactionsSymptom to the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if vtlrs._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.vial_tests_local_reactions_symptoms (` +
		`reaction_id, symptom_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, vtlrs.ReactionID, vtlrs.SymptomID, vtlrs.UpdatedBy, vtlrs.Updated)
	res, err := db.Exec(sqlstr, vtlrs.ReactionID, vtlrs.SymptomID, vtlrs.UpdatedBy, vtlrs.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	vtlrs.ID = uint(id)
	vtlrs._exists = true

	return nil
}

// Update updates the VialTestsLocalReactionsSymptom in the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vtlrs._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if vtlrs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.vial_tests_local_reactions_symptoms SET ` +
		`reaction_id = ?, symptom_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, vtlrs.ReactionID, vtlrs.SymptomID, vtlrs.UpdatedBy, vtlrs.Updated, vtlrs.ID)
	_, err = db.Exec(sqlstr, vtlrs.ReactionID, vtlrs.SymptomID, vtlrs.UpdatedBy, vtlrs.Updated, vtlrs.ID)
	return err
}

// Save saves the VialTestsLocalReactionsSymptom to the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Save(db XODB) error {
	if vtlrs.Exists() {
		return vtlrs.Update(db)
	}

	return vtlrs.Insert(db)
}

// Delete deletes the VialTestsLocalReactionsSymptom from the database.
func (vtlrs *VialTestsLocalReactionsSymptom) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !vtlrs._exists {
		return nil
	}

	// if deleted, bail
	if vtlrs._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.vial_tests_local_reactions_symptoms WHERE id = ?`

	// run query
	XOLog(sqlstr, vtlrs.ID)
	_, err = db.Exec(sqlstr, vtlrs.ID)
	if err != nil {
		return err
	}

	// set deleted
	vtlrs._deleted = true

	return nil
}

// VialTestsLocalReaction returns the VialTestsLocalReaction associated with the VialTestsLocalReactionsSymptom's ReactionID (reaction_id).
//
// Generated from foreign key 'vial_tests_local_reactions_symptoms_ibfk_1'.
func (vtlrs *VialTestsLocalReactionsSymptom) VialTestsLocalReaction(db XODB) (*VialTestsLocalReaction, error) {
	return VialTestsLocalReactionByID(db, uint(vtlrs.ReactionID.Int64))
}

// LocalReactionsSymptom returns the LocalReactionsSymptom associated with the VialTestsLocalReactionsSymptom's SymptomID (symptom_id).
//
// Generated from foreign key 'vial_tests_local_reactions_symptoms_ibfk_2'.
func (vtlrs *VialTestsLocalReactionsSymptom) LocalReactionsSymptom(db XODB) (*LocalReactionsSymptom, error) {
	return LocalReactionsSymptomByID(db, uint(vtlrs.SymptomID.Int64))
}

// User returns the User associated with the VialTestsLocalReactionsSymptom's UpdatedBy (updated_by).
//
// Generated from foreign key 'vial_tests_local_reactions_symptoms_ibfk_3'.
func (vtlrs *VialTestsLocalReactionsSymptom) User(db XODB) (*User, error) {
	return UserByID(db, uint(vtlrs.UpdatedBy.Int64))
}

// VialTestsLocalReactionsSymptomsByReactionID retrieves a row from 'AllergyNew.vial_tests_local_reactions_symptoms' as a VialTestsLocalReactionsSymptom.
//
// Generated from index 'ReactionID'.
func VialTestsLocalReactionsSymptomsByReactionID(db XODB, reactionID sql.NullInt64) ([]*VialTestsLocalReactionsSymptom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, reaction_id, symptom_id, updated_by, updated ` +
		`FROM AllergyNew.vial_tests_local_reactions_symptoms ` +
		`WHERE reaction_id = ?`

	// run query
	XOLog(sqlstr, reactionID)
	q, err := db.Query(sqlstr, reactionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialTestsLocalReactionsSymptom{}
	for q.Next() {
		vtlrs := VialTestsLocalReactionsSymptom{
			_exists: true,
		}

		// scan
		err = q.Scan(&vtlrs.ID, &vtlrs.ReactionID, &vtlrs.SymptomID, &vtlrs.UpdatedBy, &vtlrs.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vtlrs)
	}

	return res, nil
}

// VialTestsLocalReactionsSymptomsBySymptomID retrieves a row from 'AllergyNew.vial_tests_local_reactions_symptoms' as a VialTestsLocalReactionsSymptom.
//
// Generated from index 'SymptomID'.
func VialTestsLocalReactionsSymptomsBySymptomID(db XODB, symptomID sql.NullInt64) ([]*VialTestsLocalReactionsSymptom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, reaction_id, symptom_id, updated_by, updated ` +
		`FROM AllergyNew.vial_tests_local_reactions_symptoms ` +
		`WHERE symptom_id = ?`

	// run query
	XOLog(sqlstr, symptomID)
	q, err := db.Query(sqlstr, symptomID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialTestsLocalReactionsSymptom{}
	for q.Next() {
		vtlrs := VialTestsLocalReactionsSymptom{
			_exists: true,
		}

		// scan
		err = q.Scan(&vtlrs.ID, &vtlrs.ReactionID, &vtlrs.SymptomID, &vtlrs.UpdatedBy, &vtlrs.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vtlrs)
	}

	return res, nil
}

// VialTestsLocalReactionsSymptomsByUpdatedBy retrieves a row from 'AllergyNew.vial_tests_local_reactions_symptoms' as a VialTestsLocalReactionsSymptom.
//
// Generated from index 'UpdatedBy'.
func VialTestsLocalReactionsSymptomsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*VialTestsLocalReactionsSymptom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, reaction_id, symptom_id, updated_by, updated ` +
		`FROM AllergyNew.vial_tests_local_reactions_symptoms ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*VialTestsLocalReactionsSymptom{}
	for q.Next() {
		vtlrs := VialTestsLocalReactionsSymptom{
			_exists: true,
		}

		// scan
		err = q.Scan(&vtlrs.ID, &vtlrs.ReactionID, &vtlrs.SymptomID, &vtlrs.UpdatedBy, &vtlrs.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &vtlrs)
	}

	return res, nil
}

// VialTestsLocalReactionsSymptomByID retrieves a row from 'AllergyNew.vial_tests_local_reactions_symptoms' as a VialTestsLocalReactionsSymptom.
//
// Generated from index 'vial_tests_local_reactions_symptoms_id_pkey'.
func VialTestsLocalReactionsSymptomByID(db XODB, id uint) (*VialTestsLocalReactionsSymptom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, reaction_id, symptom_id, updated_by, updated ` +
		`FROM AllergyNew.vial_tests_local_reactions_symptoms ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	vtlrs := VialTestsLocalReactionsSymptom{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&vtlrs.ID, &vtlrs.ReactionID, &vtlrs.SymptomID, &vtlrs.UpdatedBy, &vtlrs.Updated)
	if err != nil {
		return nil, err
	}

	return &vtlrs, nil
}
