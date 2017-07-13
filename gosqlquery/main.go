package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	//	"net/http"
)

//type JsonResponse struct
var (
	DB = map[string]string{
		"db_name": "user:password@tcp(127.0.0.1:PORT)/DBNAME",
	}
)

func main() {
	//	http.HandleFunc("/", jsonResponseHandler)
	//	http.ListenAndServe(":8080", nil)
	if len(os.Args) > 1 {
		queries := os.Args[2:]
		for query := range queries {
			fmt.Println(queries[query])
			getJSON(queries[query], os.Args[1])
		}
	} else {
		fmt.Println("USAGE: aesqlquery 'Name of Database you want to connect to, 'sql query wrapped in double quotes, as many as you want.")
		fmt.Println("Databases avail are: db_name")
	}
}

//func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("jsonRoseponse handler reproting for duty")
//}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getJSON(sqlString string, dbConn string) (string, error) {
	// db, err := sql.Open("mysql", "vagrant:vagrant@tcp(192.168.56.22:3306)/AllergyNew")
	db, err := sql.Open("mysql", DB[dbConn])
	checkErr(err)
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.MarshalIndent(tableData, "", "  ")
	if err != nil {
		return "", err
	}
	db.Close()
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
