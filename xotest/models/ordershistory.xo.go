// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// OrdersHistory represents a row from 'AllergyNew.orders_history'.
type OrdersHistory struct {
	ID      uint           `json:"id"`       // id
	OrderID sql.NullInt64  `json:"order_id"` // order_id
	Note    sql.NullString `json:"note"`     // note
	Date    time.Time      `json:"date"`     // date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OrdersHistory exists in the database.
func (oh *OrdersHistory) Exists() bool {
	return oh._exists
}

// Deleted provides information if the OrdersHistory has been deleted from the database.
func (oh *OrdersHistory) Deleted() bool {
	return oh._deleted
}

// Insert inserts the OrdersHistory to the database.
func (oh *OrdersHistory) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if oh._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.orders_history (` +
		`order_id, note, date` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, oh.OrderID, oh.Note, oh.Date)
	res, err := db.Exec(sqlstr, oh.OrderID, oh.Note, oh.Date)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	oh.ID = uint(id)
	oh._exists = true

	return nil
}

// Update updates the OrdersHistory in the database.
func (oh *OrdersHistory) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oh._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if oh._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.orders_history SET ` +
		`order_id = ?, note = ?, date = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, oh.OrderID, oh.Note, oh.Date, oh.ID)
	_, err = db.Exec(sqlstr, oh.OrderID, oh.Note, oh.Date, oh.ID)
	return err
}

// Save saves the OrdersHistory to the database.
func (oh *OrdersHistory) Save(db XODB) error {
	if oh.Exists() {
		return oh.Update(db)
	}

	return oh.Insert(db)
}

// Delete deletes the OrdersHistory from the database.
func (oh *OrdersHistory) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oh._exists {
		return nil
	}

	// if deleted, bail
	if oh._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.orders_history WHERE id = ?`

	// run query
	XOLog(sqlstr, oh.ID)
	_, err = db.Exec(sqlstr, oh.ID)
	if err != nil {
		return err
	}

	// set deleted
	oh._deleted = true

	return nil
}

// Order returns the Order associated with the OrdersHistory's OrderID (order_id).
//
// Generated from foreign key 'orders_history_order'.
func (oh *OrdersHistory) Order(db XODB) (*Order, error) {
	return OrderByID(db, uint(oh.OrderID.Int64))
}

// OrdersHistoriesByOrderID retrieves a row from 'AllergyNew.orders_history' as a OrdersHistory.
//
// Generated from index 'Orders'.
func OrdersHistoriesByOrderID(db XODB, orderID sql.NullInt64) ([]*OrdersHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, order_id, note, date ` +
		`FROM AllergyNew.orders_history ` +
		`WHERE order_id = ?`

	// run query
	XOLog(sqlstr, orderID)
	q, err := db.Query(sqlstr, orderID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*OrdersHistory{}
	for q.Next() {
		oh := OrdersHistory{
			_exists: true,
		}

		// scan
		err = q.Scan(&oh.ID, &oh.OrderID, &oh.Note, &oh.Date)
		if err != nil {
			return nil, err
		}

		res = append(res, &oh)
	}

	return res, nil
}

// OrdersHistoryByID retrieves a row from 'AllergyNew.orders_history' as a OrdersHistory.
//
// Generated from index 'orders_history_id_pkey'.
func OrdersHistoryByID(db XODB, id uint) (*OrdersHistory, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, order_id, note, date ` +
		`FROM AllergyNew.orders_history ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	oh := OrdersHistory{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&oh.ID, &oh.OrderID, &oh.Note, &oh.Date)
	if err != nil {
		return nil, err
	}

	return &oh, nil
}