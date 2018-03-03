// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// MqTest represents a row from 'AllergyNew.mq_tests'.
type MqTest struct {
	ID                 uint           `json:"id"`                  // id
	PatientID          sql.NullInt64  `json:"patient_id"`          // patient_id
	PracticeID         sql.NullInt64  `json:"practice_id"`         // practice_id
	LocationID         sql.NullInt64  `json:"location_id"`         // location_id
	OrderID            sql.NullInt64  `json:"order_id"`            // order_id
	TemplateID         sql.NullInt64  `json:"template_id"`         // template_id
	ControlPos         sql.NullString `json:"control_pos"`         // control_pos
	ControlNeg         sql.NullString `json:"control_neg"`         // control_neg
	ControlSaline      sql.NullString `json:"control_saline"`      // control_saline
	Dilution1Strength  sql.NullInt64  `json:"dilution_1_strength"` // dilution_1_strength
	Dilution2Strength  sql.NullInt64  `json:"dilution_2_strength"` // dilution_2_strength
	Pregnant           sql.NullInt64  `json:"pregnant"`            // pregnant
	BetaBlocker        sql.NullInt64  `json:"beta_blocker"`        // beta_blocker
	SeasonalAllergies  sql.NullInt64  `json:"seasonal_allergies"`  // seasonal_allergies
	Notes              sql.NullString `json:"notes"`               // notes
	TestDate           mysql.NullTime `json:"test_date"`           // test_date
	TesterID           sql.NullInt64  `json:"tester_id"`           // tester_id
	CompleteDate       mysql.NullTime `json:"complete_date"`       // complete_date
	VitalsID           sql.NullInt64  `json:"vitals_id"`           // vitals_id
	Status             Status         `json:"status"`              // status
	RecommendTreatment sql.NullInt64  `json:"recommend_treatment"` // recommend_treatment
	RecommendID        sql.NullInt64  `json:"recommend_id"`        // recommend_id
	BillableEventID    sql.NullInt64  `json:"billable_event_id"`   // billable_event_id
	Signed             int8           `json:"signed"`              // signed
	SignedDate         mysql.NullTime `json:"signed_date"`         // signed_date
	SignedBy           sql.NullInt64  `json:"signed_by"`           // signed_by
	EmrOrdersID        sql.NullInt64  `json:"emr_orders_id"`       // emr_orders_id
	UpdatedBy          sql.NullInt64  `json:"updated_by"`          // updated_by
	Updated            time.Time      `json:"updated"`             // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MqTest exists in the database.
func (mt *MqTest) Exists() bool {
	return mt._exists
}

// Deleted provides information if the MqTest has been deleted from the database.
func (mt *MqTest) Deleted() bool {
	return mt._deleted
}

// Insert inserts the MqTest to the database.
func (mt *MqTest) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if mt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO AllergyNew.mq_tests (` +
		`patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, mt.PatientID, mt.PracticeID, mt.LocationID, mt.OrderID, mt.TemplateID, mt.ControlPos, mt.ControlNeg, mt.ControlSaline, mt.Dilution1Strength, mt.Dilution2Strength, mt.Pregnant, mt.BetaBlocker, mt.SeasonalAllergies, mt.Notes, mt.TestDate, mt.TesterID, mt.CompleteDate, mt.VitalsID, mt.Status, mt.RecommendTreatment, mt.RecommendID, mt.BillableEventID, mt.Signed, mt.SignedDate, mt.SignedBy, mt.EmrOrdersID, mt.UpdatedBy, mt.Updated)
	res, err := db.Exec(sqlstr, mt.PatientID, mt.PracticeID, mt.LocationID, mt.OrderID, mt.TemplateID, mt.ControlPos, mt.ControlNeg, mt.ControlSaline, mt.Dilution1Strength, mt.Dilution2Strength, mt.Pregnant, mt.BetaBlocker, mt.SeasonalAllergies, mt.Notes, mt.TestDate, mt.TesterID, mt.CompleteDate, mt.VitalsID, mt.Status, mt.RecommendTreatment, mt.RecommendID, mt.BillableEventID, mt.Signed, mt.SignedDate, mt.SignedBy, mt.EmrOrdersID, mt.UpdatedBy, mt.Updated)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	mt.ID = uint(id)
	mt._exists = true

	return nil
}

// Update updates the MqTest in the database.
func (mt *MqTest) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mt._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if mt._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE AllergyNew.mq_tests SET ` +
		`patient_id = ?, practice_id = ?, location_id = ?, order_id = ?, template_id = ?, control_pos = ?, control_neg = ?, control_saline = ?, dilution_1_strength = ?, dilution_2_strength = ?, pregnant = ?, beta_blocker = ?, seasonal_allergies = ?, notes = ?, test_date = ?, tester_id = ?, complete_date = ?, vitals_id = ?, status = ?, recommend_treatment = ?, recommend_id = ?, billable_event_id = ?, signed = ?, signed_date = ?, signed_by = ?, emr_orders_id = ?, updated_by = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, mt.PatientID, mt.PracticeID, mt.LocationID, mt.OrderID, mt.TemplateID, mt.ControlPos, mt.ControlNeg, mt.ControlSaline, mt.Dilution1Strength, mt.Dilution2Strength, mt.Pregnant, mt.BetaBlocker, mt.SeasonalAllergies, mt.Notes, mt.TestDate, mt.TesterID, mt.CompleteDate, mt.VitalsID, mt.Status, mt.RecommendTreatment, mt.RecommendID, mt.BillableEventID, mt.Signed, mt.SignedDate, mt.SignedBy, mt.EmrOrdersID, mt.UpdatedBy, mt.Updated, mt.ID)
	_, err = db.Exec(sqlstr, mt.PatientID, mt.PracticeID, mt.LocationID, mt.OrderID, mt.TemplateID, mt.ControlPos, mt.ControlNeg, mt.ControlSaline, mt.Dilution1Strength, mt.Dilution2Strength, mt.Pregnant, mt.BetaBlocker, mt.SeasonalAllergies, mt.Notes, mt.TestDate, mt.TesterID, mt.CompleteDate, mt.VitalsID, mt.Status, mt.RecommendTreatment, mt.RecommendID, mt.BillableEventID, mt.Signed, mt.SignedDate, mt.SignedBy, mt.EmrOrdersID, mt.UpdatedBy, mt.Updated, mt.ID)
	return err
}

// Save saves the MqTest to the database.
func (mt *MqTest) Save(db XODB) error {
	if mt.Exists() {
		return mt.Update(db)
	}

	return mt.Insert(db)
}

// Delete deletes the MqTest from the database.
func (mt *MqTest) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !mt._exists {
		return nil
	}

	// if deleted, bail
	if mt._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM AllergyNew.mq_tests WHERE id = ?`

	// run query
	XOLog(sqlstr, mt.ID)
	_, err = db.Exec(sqlstr, mt.ID)
	if err != nil {
		return err
	}

	// set deleted
	mt._deleted = true

	return nil
}

// PracticeLocation returns the PracticeLocation associated with the MqTest's LocationID (location_id).
//
// Generated from foreign key 'mq_tests_location'.
func (mt *MqTest) PracticeLocation(db XODB) (*PracticeLocation, error) {
	return PracticeLocationByID(db, uint(mt.LocationID.Int64))
}

// Order returns the Order associated with the MqTest's OrderID (order_id).
//
// Generated from foreign key 'mq_tests_order'.
func (mt *MqTest) Order(db XODB) (*Order, error) {
	return OrderByID(db, uint(mt.OrderID.Int64))
}

// Patient returns the Patient associated with the MqTest's PatientID (patient_id).
//
// Generated from foreign key 'mq_tests_patient'.
func (mt *MqTest) Patient(db XODB) (*Patient, error) {
	return PatientByID(db, uint(mt.PatientID.Int64))
}

// Practice returns the Practice associated with the MqTest's PracticeID (practice_id).
//
// Generated from foreign key 'mq_tests_practice'.
func (mt *MqTest) Practice(db XODB) (*Practice, error) {
	return PracticeByID(db, uint(mt.PracticeID.Int64))
}

// MqTestTemplate returns the MqTestTemplate associated with the MqTest's TemplateID (template_id).
//
// Generated from foreign key 'mq_tests_template'.
func (mt *MqTest) MqTestTemplate(db XODB) (*MqTestTemplate, error) {
	return MqTestTemplateByID(db, uint(mt.TemplateID.Int64))
}

// UserByTesterID returns the User associated with the MqTest's TesterID (tester_id).
//
// Generated from foreign key 'mq_tests_tester'.
func (mt *MqTest) UserByTesterID(db XODB) (*User, error) {
	return UserByID(db, uint(mt.TesterID.Int64))
}

// UserByUpdatedBy returns the User associated with the MqTest's UpdatedBy (updated_by).
//
// Generated from foreign key 'mq_tests_updated'.
func (mt *MqTest) UserByUpdatedBy(db XODB) (*User, error) {
	return UserByID(db, uint(mt.UpdatedBy.Int64))
}

// Vital returns the Vital associated with the MqTest's VitalsID (vitals_id).
//
// Generated from foreign key 'mq_tests_vitals'.
func (mt *MqTest) Vital(db XODB) (*Vital, error) {
	return VitalByID(db, uint(mt.VitalsID.Int64))
}

// MqTestsByOrderID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'Orders'.
func MqTestsByOrderID(db XODB, orderID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE order_id = ?`

	// run query
	XOLog(sqlstr, orderID)
	q, err := db.Query(sqlstr, orderID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByPatientID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'Patients'.
func MqTestsByPatientID(db XODB, patientID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE patient_id = ?`

	// run query
	XOLog(sqlstr, patientID)
	q, err := db.Query(sqlstr, patientID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByPracticeID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'Practices'.
func MqTestsByPracticeID(db XODB, practiceID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE practice_id = ?`

	// run query
	XOLog(sqlstr, practiceID)
	q, err := db.Query(sqlstr, practiceID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByTemplateID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'Templates'.
func MqTestsByTemplateID(db XODB, templateID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE template_id = ?`

	// run query
	XOLog(sqlstr, templateID)
	q, err := db.Query(sqlstr, templateID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByTestDate retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_test_date'.
func MqTestsByTestDate(db XODB, testDate mysql.NullTime) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE test_date = ?`

	// run query
	XOLog(sqlstr, testDate)
	q, err := db.Query(sqlstr, testDate)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByStatus retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_test_status'.
func MqTestsByStatus(db XODB, status Status) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE status = ?`

	// run query
	XOLog(sqlstr, status)
	q, err := db.Query(sqlstr, status)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestByID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_tests_id_pkey'.
func MqTestByID(db XODB, id uint) (*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	mt := MqTest{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
	if err != nil {
		return nil, err
	}

	return &mt, nil
}

// MqTestsByLocationID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_tests_location'.
func MqTestsByLocationID(db XODB, locationID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE location_id = ?`

	// run query
	XOLog(sqlstr, locationID)
	q, err := db.Query(sqlstr, locationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByTesterID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_tests_tester'.
func MqTestsByTesterID(db XODB, testerID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE tester_id = ?`

	// run query
	XOLog(sqlstr, testerID)
	q, err := db.Query(sqlstr, testerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByUpdatedBy retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_tests_updated'.
func MqTestsByUpdatedBy(db XODB, updatedBy sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}

// MqTestsByVitalsID retrieves a row from 'AllergyNew.mq_tests' as a MqTest.
//
// Generated from index 'mq_tests_vitals'.
func MqTestsByVitalsID(db XODB, vitalsID sql.NullInt64) ([]*MqTest, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, patient_id, practice_id, location_id, order_id, template_id, control_pos, control_neg, control_saline, dilution_1_strength, dilution_2_strength, pregnant, beta_blocker, seasonal_allergies, notes, test_date, tester_id, complete_date, vitals_id, status, recommend_treatment, recommend_id, billable_event_id, signed, signed_date, signed_by, emr_orders_id, updated_by, updated ` +
		`FROM AllergyNew.mq_tests ` +
		`WHERE vitals_id = ?`

	// run query
	XOLog(sqlstr, vitalsID)
	q, err := db.Query(sqlstr, vitalsID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MqTest{}
	for q.Next() {
		mt := MqTest{
			_exists: true,
		}

		// scan
		err = q.Scan(&mt.ID, &mt.PatientID, &mt.PracticeID, &mt.LocationID, &mt.OrderID, &mt.TemplateID, &mt.ControlPos, &mt.ControlNeg, &mt.ControlSaline, &mt.Dilution1Strength, &mt.Dilution2Strength, &mt.Pregnant, &mt.BetaBlocker, &mt.SeasonalAllergies, &mt.Notes, &mt.TestDate, &mt.TesterID, &mt.CompleteDate, &mt.VitalsID, &mt.Status, &mt.RecommendTreatment, &mt.RecommendID, &mt.BillableEventID, &mt.Signed, &mt.SignedDate, &mt.SignedBy, &mt.EmrOrdersID, &mt.UpdatedBy, &mt.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &mt)
	}

	return res, nil
}
