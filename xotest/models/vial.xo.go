// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Vial represents a row from 'AllergyNew.vials'.
type Vial struct {
	ID                    uint            `json:"id"`                      // id
	VialTypeID            sql.NullInt64   `json:"vial_type_id"`            // vial_type_id
	PatientID             sql.NullInt64   `json:"patient_id"`              // patient_id
	Barcode               sql.NullString  `json:"barcode"`                 // barcode
	StorageCode           sql.NullString  `json:"storage_code"`            // storage_code
	FormulaVialID         sql.NullInt64   `json:"formula_vial_id"`         // formula_vial_id
	ExpirationDate        mysql.NullTime  `json:"expiration_date"`         // expiration_date
	Volume                sql.NullInt64   `json:"volume"`                  // volume
	MixMethod             MixMethod       `json:"mix_method"`              // mix_method
	Maintenance           sql.NullInt64   `json:"maintenance"`             // maintenance
	Concentrated          int8            `json:"concentrated"`            // concentrated
	ConcentrateRatio      sql.NullInt64   `json:"concentrate_ratio"`       // concentrate_ratio
	Ratio                 sql.NullInt64   `json:"ratio"`                   // ratio
	LastDosage            sql.NullFloat64 `json:"last_dosage"`             // last_dosage
	CreateDate            mysql.NullTime  `json:"create_date"`             // create_date
	CreatedBy             sql.NullInt64   `json:"created_by"`              // created_by
	RequestID             sql.NullInt64   `json:"request_id"`              // request_id
	Mixed                 int8            `json:"mixed"`                   // mixed
	MixerID               sql.NullInt64   `json:"mixer_id"`                // mixer_id
	MixDate               mysql.NullTime  `json:"mix_date"`                // mix_date
	MixPracticeID         sql.NullInt64   `json:"mix_practice_id"`         // mix_practice_id
	MixLocationID         sql.NullInt64   `json:"mix_location_id"`         // mix_location_id
	Reviewed              int8            `json:"reviewed"`                // reviewed
	ReviewerID            sql.NullInt64   `json:"reviewer_id"`             // reviewer_id
	ReviewDate            mysql.NullTime  `json:"review_date"`             // review_date
	ReviewPracticeID      sql.NullInt64   `json:"review_practice_id"`      // review_practice_id
	ReviewLocationID      sql.NullInt64   `json:"review_location_id"`      // review_location_id
	Weakened              int8            `json:"weakened"`                // weakened
	WeakenedAmt           sql.NullInt64   `json:"weakened_amt"`            // weakened_amt
	ParentVialID          sql.NullInt64   `json:"parent_vial_id"`          // parent_vial_id
	OriginalVialID        sql.NullInt64   `json:"original_vial_id"`        // original_vial_id
	TestRequired          sql.NullInt64   `json:"test_required"`           // test_required
	TestID                sql.NullInt64   `json:"test_id"`                 // test_id
	StatusID              sql.NullInt64   `json:"status_id"`               // status_id
	AvailableForTreatment sql.NullInt64   `json:"available_for_treatment"` // available_for_treatment
	Archived              int8            `json:"archived"`                // archived
	Titrated              int8            `json:"titrated"`                // titrated
	IsChild               sql.NullInt64   `json:"is_child"`                // is_child
	HistoryParentID       sql.NullInt64   `json:"history_parent_id"`       // history_parent_id
	Method                Method          `json:"method"`                  // method
	CurrentLocationID     sql.NullInt64   `json:"current_location_id"`     // current_location_id
	BillableEventID       sql.NullInt64   `json:"billable_event_id"`       // billable_event_id
	UpdatedBy             sql.NullInt64   `json:"updated_by"`              // updated_by
	Updated               time.Time       `json:"updated"`                 // updated
	FromTitrated          sql.NullInt64   `json:"from_titrated"`           // from_titrated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Vial exists in the database.
func (v *Vial) Exists() bool {
	return v._exists
}

// Deleted provides information if the Vial has been deleted from the database.
func (v *Vial) Deleted() bool {
	return v._deleted
}

// Insert inserts the Vial to the database.
func (v *Vial) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if v._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.vials (` +
		`vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, v.VialTypeID, v.PatientID, v.Barcode, v.StorageCode, v.FormulaVialID, v.ExpirationDate, v.Volume, v.MixMethod, v.Maintenance, v.Concentrated, v.ConcentrateRatio, v.Ratio, v.LastDosage, v.CreateDate, v.CreatedBy, v.RequestID, v.Mixed, v.MixerID, v.MixDate, v.MixPracticeID, v.MixLocationID, v.Reviewed, v.ReviewerID, v.ReviewDate, v.ReviewPracticeID, v.ReviewLocationID, v.Weakened, v.WeakenedAmt, v.ParentVialID, v.OriginalVialID, v.TestRequired, v.TestID, v.StatusID, v.AvailableForTreatment, v.Archived, v.Titrated, v.IsChild, v.HistoryParentID, v.Method, v.CurrentLocationID, v.BillableEventID, v.UpdatedBy, v.Updated, v.FromTitrated)
	res, err := db.Exec(sqlstr, v.VialTypeID, v.PatientID, v.Barcode, v.StorageCode, v.FormulaVialID, v.ExpirationDate, v.Volume, v.MixMethod, v.Maintenance, v.Concentrated, v.ConcentrateRatio, v.Ratio, v.LastDosage, v.CreateDate, v.CreatedBy, v.RequestID, v.Mixed, v.MixerID, v.MixDate, v.MixPracticeID, v.MixLocationID, v.Reviewed, v.ReviewerID, v.ReviewDate, v.ReviewPracticeID, v.ReviewLocationID, v.Weakened, v.WeakenedAmt, v.ParentVialID, v.OriginalVialID, v.TestRequired, v.TestID, v.StatusID, v.AvailableForTreatment, v.Archived, v.Titrated, v.IsChild, v.HistoryParentID, v.Method, v.CurrentLocationID, v.BillableEventID, v.UpdatedBy, v.Updated, v.FromTitrated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	v.ID = uint(id)
	v._exists = true

	return nil
}

// Update updates the Vial in the database.
func (v *Vial) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if v._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.vials SET ` +
		`vial_type_id = ?, patient_id = ?, barcode = ?, storage_code = ?, formula_vial_id = ?, expiration_date = ?, volume = ?, mix_method = ?, maintenance = ?, concentrated = ?, concentrate_ratio = ?, ratio = ?, last_dosage = ?, create_date = ?, created_by = ?, request_id = ?, mixed = ?, mixer_id = ?, mix_date = ?, mix_practice_id = ?, mix_location_id = ?, reviewed = ?, reviewer_id = ?, review_date = ?, review_practice_id = ?, review_location_id = ?, weakened = ?, weakened_amt = ?, parent_vial_id = ?, original_vial_id = ?, test_required = ?, test_id = ?, status_id = ?, available_for_treatment = ?, archived = ?, titrated = ?, is_child = ?, history_parent_id = ?, method = ?, current_location_id = ?, billable_event_id = ?, updated_by = ?, updated = ?, from_titrated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, v.VialTypeID, v.PatientID, v.Barcode, v.StorageCode, v.FormulaVialID, v.ExpirationDate, v.Volume, v.MixMethod, v.Maintenance, v.Concentrated, v.ConcentrateRatio, v.Ratio, v.LastDosage, v.CreateDate, v.CreatedBy, v.RequestID, v.Mixed, v.MixerID, v.MixDate, v.MixPracticeID, v.MixLocationID, v.Reviewed, v.ReviewerID, v.ReviewDate, v.ReviewPracticeID, v.ReviewLocationID, v.Weakened, v.WeakenedAmt, v.ParentVialID, v.OriginalVialID, v.TestRequired, v.TestID, v.StatusID, v.AvailableForTreatment, v.Archived, v.Titrated, v.IsChild, v.HistoryParentID, v.Method, v.CurrentLocationID, v.BillableEventID, v.UpdatedBy, v.Updated, v.FromTitrated, v.ID)
	_, err = db.Exec(sqlstr, v.VialTypeID, v.PatientID, v.Barcode, v.StorageCode, v.FormulaVialID, v.ExpirationDate, v.Volume, v.MixMethod, v.Maintenance, v.Concentrated, v.ConcentrateRatio, v.Ratio, v.LastDosage, v.CreateDate, v.CreatedBy, v.RequestID, v.Mixed, v.MixerID, v.MixDate, v.MixPracticeID, v.MixLocationID, v.Reviewed, v.ReviewerID, v.ReviewDate, v.ReviewPracticeID, v.ReviewLocationID, v.Weakened, v.WeakenedAmt, v.ParentVialID, v.OriginalVialID, v.TestRequired, v.TestID, v.StatusID, v.AvailableForTreatment, v.Archived, v.Titrated, v.IsChild, v.HistoryParentID, v.Method, v.CurrentLocationID, v.BillableEventID, v.UpdatedBy, v.Updated, v.FromTitrated, v.ID)
	return err
}

// Save saves the Vial to the database.
func (v *Vial) Save(db XODB) error {
	if v.Exists() {
		return v.Update(db)
	}

	return v.Insert(db)
}

// Delete deletes the Vial from the database.
func (v *Vial) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return nil
	}

	// if deleted, bail
	if v._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.vials WHERE id = ?`

	// run query
	XOLog(sqlstr, v.ID)
	_, err = db.Exec(sqlstr, v.ID)
	if err != nil {
		return err
	}

	// set deleted
	v._deleted = true

	return nil
}

// FormulaVial returns the FormulaVial associated with the Vial's FormulaVialID (formula_vial_id).
//
// Generated from foreign key 'vials_formula_vials'.
func (v *Vial) FormulaVial(db XODB) (*FormulaVial, error) {
	return FormulaVialByID(db, uint(v.FormulaVialID.Int64))
}

// VialLocation returns the VialLocation associated with the Vial's CurrentLocationID (current_location_id).
//
// Generated from foreign key 'vials_locations'.
func (v *Vial) VialLocation(db XODB) (*VialLocation, error) {
	return VialLocationByID(db, uint(v.CurrentLocationID.Int64))
}

// VialByOriginalVialID returns the Vial associated with the Vial's OriginalVialID (original_vial_id).
//
// Generated from foreign key 'vials_original_vial'.
func (v *Vial) VialByOriginalVialID(db XODB) (*Vial, error) {
	return VialByID(db, uint(v.OriginalVialID.Int64))
}

// VialByParentVialID returns the Vial associated with the Vial's ParentVialID (parent_vial_id).
//
// Generated from foreign key 'vials_parent_vial'.
func (v *Vial) VialByParentVialID(db XODB) (*Vial, error) {
	return VialByID(db, uint(v.ParentVialID.Int64))
}

// Patient returns the Patient associated with the Vial's PatientID (patient_id).
//
// Generated from foreign key 'vials_patients'.
func (v *Vial) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(v.PatientID.Int64))
}

// VialRequest returns the VialRequest associated with the Vial's RequestID (request_id).
//
// Generated from foreign key 'vials_requests'.
func (v *Vial) VialRequest(db XODB) (*VialRequest, error) {
	return VialRequestByID(db, uint(v.RequestID.Int64))
}

// VialStatus returns the VialStatus associated with the Vial's StatusID (status_id).
//
// Generated from foreign key 'vials_statuses'.
func (v *Vial) VialStatus(db XODB) (*VialStatus, error) {
	return VialStatusByID(db, uint(v.StatusID.Int64))
}

// VialTest returns the VialTest associated with the Vial's TestID (test_id).
//
// Generated from foreign key 'vials_test'.
func (v *Vial) VialTest(db XODB) (*VialTest, error) {
	return VialTestByID(db, uint(v.TestID.Int64))
}

// User returns the User associated with the Vial's UpdatedBy (updated_by).
//
// Generated from foreign key 'vials_updated'.
func (v *Vial) User(db XODB) (*User, error) {
	return UserByID(db, uint(v.UpdatedBy.Int64))
}

// VialType returns the VialType associated with the Vial's VialTypeID (vial_type_id).
//
// Generated from foreign key 'vials_vial_types'.
func (v *Vial) VialType(db XODB) (*VialType, error) {
	return VialTypeByID(db, uint(v.VialTypeID.Int64))
}

// VialsByMixerID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'Mixers'.
func VialsByMixerID(db XODB, mixerID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE mixer_id = ?`

	// run query
	XOLog(sqlstr, mixerID)
	q, err := db.Query(sqlstr, mixerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByPatientID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'Patients'.
func VialsByPatientID(db XODB, patientID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByRequestID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'Requests'.
func VialsByRequestID(db XODB, requestID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE request_id = ?`

	// run query
	XOLog(sqlstr, requestID)
	q, err := db.Query(sqlstr, requestID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByVialTypeID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'Vial_Types'.
func VialsByVialTypeID(db XODB, vialTypeID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE vial_type_id = ?`

	// run query
	XOLog(sqlstr, vialTypeID)
	q, err := db.Query(sqlstr, vialTypeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByMixPracticeID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vial_mix_practice_id'.
func VialsByMixPracticeID(db XODB, mixPracticeID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE mix_practice_id = ?`

	// run query
	XOLog(sqlstr, mixPracticeID)
	q, err := db.Query(sqlstr, mixPracticeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByFormulaVialID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_formula_vials'.
func VialsByFormulaVialID(db XODB, formulaVialID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE formula_vial_id = ?`

	// run query
	XOLog(sqlstr, formulaVialID)
	q, err := db.Query(sqlstr, formulaVialID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialByID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_id_pkey'.
func VialByID(db XODB, id uint) (*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	v := Vial{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// VialsByCurrentLocationID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_locations'.
func VialsByCurrentLocationID(db XODB, currentLocationID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE current_location_id = ?`

	// run query
	XOLog(sqlstr, currentLocationID)
	q, err := db.Query(sqlstr, currentLocationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByOriginalVialID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_original_vial'.
func VialsByOriginalVialID(db XODB, originalVialID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE original_vial_id = ?`

	// run query
	XOLog(sqlstr, originalVialID)
	q, err := db.Query(sqlstr, originalVialID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByParentVialID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_parent_vial'.
func VialsByParentVialID(db XODB, parentVialID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE parent_vial_id = ?`

	// run query
	XOLog(sqlstr, parentVialID)
	q, err := db.Query(sqlstr, parentVialID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByStatusID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_statuses'.
func VialsByStatusID(db XODB, statusID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE status_id = ?`

	// run query
	XOLog(sqlstr, statusID)
	q, err := db.Query(sqlstr, statusID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByTestID retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_test'.
func VialsByTestID(db XODB, testID sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE test_id = ?`

	// run query
	XOLog(sqlstr, testID)
	q, err := db.Query(sqlstr, testID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}

// VialsByUpdatedBy retrieves a row from 'AllergyNew.vials' as a Vial.
//
// Generated from index 'vials_updated'.
func VialsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*Vial, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, vial_type_id, patient_id, barcode, storage_code, formula_vial_id, expiration_date, volume, mix_method, maintenance, concentrated, concentrate_ratio, ratio, last_dosage, create_date, created_by, request_id, mixed, mixer_id, mix_date, mix_practice_id, mix_location_id, reviewed, reviewer_id, review_date, review_practice_id, review_location_id, weakened, weakened_amt, parent_vial_id, original_vial_id, test_required, test_id, status_id, available_for_treatment, archived, titrated, is_child, history_parent_id, method, current_location_id, billable_event_id, updated_by, updated, from_titrated ` +
		`FROM AllergyNew.vials ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Vial{}
	for q.Next() {
		v := Vial{
			_exists: true,
		}

		// scan
		err = q.Scan(&v.ID, &v.VialTypeID, &v.PatientID, &v.Barcode, &v.StorageCode, &v.FormulaVialID, &v.ExpirationDate, &v.Volume, &v.MixMethod, &v.Maintenance, &v.Concentrated, &v.ConcentrateRatio, &v.Ratio, &v.LastDosage, &v.CreateDate, &v.CreatedBy, &v.RequestID, &v.Mixed, &v.MixerID, &v.MixDate, &v.MixPracticeID, &v.MixLocationID, &v.Reviewed, &v.ReviewerID, &v.ReviewDate, &v.ReviewPracticeID, &v.ReviewLocationID, &v.Weakened, &v.WeakenedAmt, &v.ParentVialID, &v.OriginalVialID, &v.TestRequired, &v.TestID, &v.StatusID, &v.AvailableForTreatment, &v.Archived, &v.Titrated, &v.IsChild, &v.HistoryParentID, &v.Method, &v.CurrentLocationID, &v.BillableEventID, &v.UpdatedBy, &v.Updated, &v.FromTitrated)
		if err != nil {
			return nil, err
		}

		res = append(res, &v)
	}

	return res, nil
}
