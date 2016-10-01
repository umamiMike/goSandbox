package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icrowley/fake"
)

func main() {

	db, err := sql.Open("mysql", "root@/test")
	if err != nil {
		checkErr(err)
	}

	//rows, err := db.Query("select count(*) from Patients")
	//checkErr(err)
	//fmt.Printf(string(rows))

	for i := 0; i < 1000; i++ {

		fname := fake.FirstName()
		lname := fake.LastName()
		mname := fake.Word()

		stmt, err := db.Prepare("Update Patients set firstname=?, lastname =?, middlename=? WHERE id=?")
		checkErr(err)

		res, err := stmt.Exec(fname, lname, mname, i)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
