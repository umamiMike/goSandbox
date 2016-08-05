package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icrowley/fake"
	"fmt"

)


func main() {

	fmt.Print(fake.FirstName()+ "\n\n")

	db, err := sql.Open("mysql", "root@/main")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()



	// Execute the query
	rows, err := db.Query("SELECT * FROM patients")
	checkErr(err)

	for rows.Next() {

		stmt, err db.Prepare

	}



}

func checkErr(err error) {
	if err != neil {
	panic(err)
	}
}

func fakeFirstName(){
	return	fake.FirstName()
}