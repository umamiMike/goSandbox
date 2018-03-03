// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// AntigenInventory represents a row from 'AllergyNew.antigen_inventory'.
type AntigenInventory struct {
	ID                       uint           `json:"id"`                          // id
	AntigenID                sql.NullInt64  `json:"antigen_id"`                  // antigen_id
	SupplierID               sql.NullInt64  `json:"supplier_id"`                 // supplier_id
	PracticeID               sql.NullInt64  `json:"practice_id"`                 // practice_id
	LocationID               sql.NullInt64  `json:"location_id"`                 // location_id
	StatusID                 sql.NullInt64  `json:"status_id"`                   // status_id
	VialSizeID               sql.NullInt64  `json:"vial_size_id"`                // vial_size_id
	ReceivedBy               sql.NullString `json:"received_by"`                 // received_by
	ReceiverID               sql.NullInt64  `json:"receiver_id"`                 // receiver_id
	ReceivedDate             mysql.NullTime `json:"received_date"`               // received_date
	ExpirationDate           mysql.NullTime `json:"expiration_date"`             // expiration_date
	Lot                      sql.NullString `json:"lot"`                         // lot
	IsMix                    sql.NullInt64  `json:"is_mix"`                      // is_mix
	MixAntigenID             sql.NullInt64  `json:"mix_antigen_id"`              // mix_antigen_id
	ParentAntigenInventoryID sql.NullInt64  `json:"parent_antigen_inventory_id"` // parent_antigen_inventory_id
	UpdatedBy                sql.NullInt64  `json:"updated_by"`                  // updated_by
	Updated                  time.Time      `json:"updated"`                     // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AntigenInventory exists in the database.
func (ai *AntigenInventory) Exists() bool {
	return ai._exists
}

// Deleted provides information if the AntigenInventory has been deleted from the database.
func (ai *AntigenInventory) Deleted() bool {
	return ai._deleted
}

// Insert inserts the AntigenInventory to the database.
func (ai *AntigenInventory) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ai._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.antigen_inventory (` +
		`antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ai.AntigenID, ai.SupplierID, ai.PracticeID, ai.LocationID, ai.StatusID, ai.VialSizeID, ai.ReceivedBy, ai.ReceiverID, ai.ReceivedDate, ai.ExpirationDate, ai.Lot, ai.IsMix, ai.MixAntigenID, ai.ParentAntigenInventoryID, ai.UpdatedBy, ai.Updated)
	res, err := db.Exec(sqlstr, ai.AntigenID, ai.SupplierID, ai.PracticeID, ai.LocationID, ai.StatusID, ai.VialSizeID, ai.ReceivedBy, ai.ReceiverID, ai.ReceivedDate, ai.ExpirationDate, ai.Lot, ai.IsMix, ai.MixAntigenID, ai.ParentAntigenInventoryID, ai.UpdatedBy, ai.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ai.ID = uint(id)
	ai._exists = true

	return nil
}

// Update updates the AntigenInventory in the database.
func (ai *AntigenInventory) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ai._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ai._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.antigen_inventory SET ` +
		`antigen_id = ?, supplier_id = ?, practice_id = ?, location_id = ?, status_id = ?, vial_size_id = ?, received_by = ?, receiver_id = ?, received_date = ?, expiration_date = ?, lot = ?, is_mix = ?, mix_antigen_id = ?, parent_antigen_inventory_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ai.AntigenID, ai.SupplierID, ai.PracticeID, ai.LocationID, ai.StatusID, ai.VialSizeID, ai.ReceivedBy, ai.ReceiverID, ai.ReceivedDate, ai.ExpirationDate, ai.Lot, ai.IsMix, ai.MixAntigenID, ai.ParentAntigenInventoryID, ai.UpdatedBy, ai.Updated, ai.ID)
	_, err = db.Exec(sqlstr, ai.AntigenID, ai.SupplierID, ai.PracticeID, ai.LocationID, ai.StatusID, ai.VialSizeID, ai.ReceivedBy, ai.ReceiverID, ai.ReceivedDate, ai.ExpirationDate, ai.Lot, ai.IsMix, ai.MixAntigenID, ai.ParentAntigenInventoryID, ai.UpdatedBy, ai.Updated, ai.ID)
	return err
}

// Save saves the AntigenInventory to the database.
func (ai *AntigenInventory) Save(db XODB) error {
	if ai.Exists() {
		return ai.Update(db)
	}

	return ai.Insert(db)
}

// Delete deletes the AntigenInventory from the database.
func (ai *AntigenInventory) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ai._exists {
		return nil
	}

	// if deleted, bail
	if ai._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.antigen_inventory WHERE id = ?`

	// run query
	XOLog(sqlstr, ai.ID)
	_, err = db.Exec(sqlstr, ai.ID)
	if err != nil {
		return err
	}

	// set deleted
	ai._deleted = true

	return nil
}

// Antigen returns the Antigen associated with the AntigenInventory's AntigenID (antigen_id).
//
// Generated from foreign key 'antigen_inventory_antigen'.
func (ai *AntigenInventory) Antigen(db XODB) (*Antigen, error) {
	return AntigenByID(db, uint(ai.AntigenID.Int64))
}

// PracticeLocation returns the PracticeLocation associated with the AntigenInventory's LocationID (location_id).
//
// Generated from foreign key 'antigen_inventory_location'.
func (ai *AntigenInventory) PracticeLocation(db XODB) (*PracticeLocation, error) {
	return PracticeLocationByID(db, uint(ai.LocationID.Int64))
}

// Practice returns the Practice associated with the AntigenInventory's PracticeID (practice_id).
//
// Generated from foreign key 'antigen_inventory_practice'.
func (ai *AntigenInventory) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(ai.PracticeID.Int64))
}

// AntigenInventoryStatus returns the AntigenInventoryStatus associated with the AntigenInventory's StatusID (status_id).
//
// Generated from foreign key 'antigen_inventory_status'.
func (ai *AntigenInventory) AntigenInventoryStatus(db XODB) (*AntigenInventoryStatus, error) {
	return AntigenInventoryStatusByID(db, uint(ai.StatusID.Int64))
}

// Supplier returns the Supplier associated with the AntigenInventory's SupplierID (supplier_id).
//
// Generated from foreign key 'antigen_inventory_supplier'.
func (ai *AntigenInventory) Supplier(db XODB) (*Supplier, error) {
	return SupplierByID(db, uint(ai.SupplierID.Int64))
}

// User returns the User associated with the AntigenInventory's UpdatedBy (updated_by).
//
// Generated from foreign key 'antigen_inventory_updated'.
func (ai *AntigenInventory) User(db XODB) (*User, error) {
	return UserByID(db, uint(ai.UpdatedBy.Int64))
}

// AntigenVialSize returns the AntigenVialSize associated with the AntigenInventory's VialSizeID (vial_size_id).
//
// Generated from foreign key 'antigen_inventory_vial_size'.
func (ai *AntigenInventory) AntigenVialSize(db XODB) (*AntigenVialSize, error) {
	return AntigenVialSizeByID(db, uint(ai.VialSizeID.Int64))
}

// AntigenInventoriesByAntigenID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Antigens'.
func AntigenInventoriesByAntigenID(db XODB, antigenID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE antigen_id = ?`

	// run query
	XOLog(sqlstr, antigenID)
	q, err := db.Query(sqlstr, antigenID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByLocationID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Locations'.
func AntigenInventoriesByLocationID(db XODB, locationID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE location_id = ?`

	// run query
	XOLog(sqlstr, locationID)
	q, err := db.Query(sqlstr, locationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByLot retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Lot_Number'.
func AntigenInventoriesByLot(db XODB, lot sql.NullString) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE lot = ?`

	// run query
	XOLog(sqlstr, lot)
	q, err := db.Query(sqlstr, lot)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByPracticeID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Practices'.
func AntigenInventoriesByPracticeID(db XODB, practiceID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByReceiverID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Receivers'.
func AntigenInventoriesByReceiverID(db XODB, receiverID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE receiver_id = ?`

	// run query
	XOLog(sqlstr, receiverID)
	q, err := db.Query(sqlstr, receiverID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByStatusID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'Statuses'.
func AntigenInventoriesByStatusID(db XODB, statusID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE status_id = ?`

	// run query
	XOLog(sqlstr, statusID)
	q, err := db.Query(sqlstr, statusID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoryByID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'antigen_inventory_id_pkey'.
func AntigenInventoryByID(db XODB, id uint) (*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ai := AntigenInventory{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
	if err != nil {
		return nil, err
	}

	return &ai, nil
}

// AntigenInventoriesBySupplierID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'antigen_inventory_supplier'.
func AntigenInventoriesBySupplierID(db XODB, supplierID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE supplier_id = ?`

	// run query
	XOLog(sqlstr, supplierID)
	q, err := db.Query(sqlstr, supplierID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByUpdatedBy retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'antigen_inventory_updated'.
func AntigenInventoriesByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}

// AntigenInventoriesByVialSizeID retrieves a row from 'AllergyNew.antigen_inventory' as a AntigenInventory.
//
// Generated from index 'antigen_inventory_vial_size'.
func AntigenInventoriesByVialSizeID(db XODB, vialSizeID sql.NullInt64) ([]*AntigenInventory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, antigen_id, supplier_id, practice_id, location_id, status_id, vial_size_id, received_by, receiver_id, received_date, expiration_date, lot, is_mix, mix_antigen_id, parent_antigen_inventory_id, updated_by, updated ` +
		`FROM AllergyNew.antigen_inventory ` +
		`WHERE vial_size_id = ?`

	// run query
	XOLog(sqlstr, vialSizeID)
	q, err := db.Query(sqlstr, vialSizeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AntigenInventory{}
	for q.Next() {
		ai := AntigenInventory{
			_exists: true,
		}

		// scan
		err = q.Scan(&ai.ID, &ai.AntigenID, &ai.SupplierID, &ai.PracticeID, &ai.LocationID, &ai.StatusID, &ai.VialSizeID, &ai.ReceivedBy, &ai.ReceiverID, &ai.ReceivedDate, &ai.ExpirationDate, &ai.Lot, &ai.IsMix, &ai.MixAntigenID, &ai.ParentAntigenInventoryID, &ai.UpdatedBy, &ai.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &ai)
	}

	return res, nil
}
