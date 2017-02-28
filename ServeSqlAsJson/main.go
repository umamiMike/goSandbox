package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//type JsonResponse struct

func main() {
	//sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	getJSON("select mix_date, patients.firstname, patients.lastname, barcode from vials inner join patients on patient_id  = patients.id where mix_date IS NOT null order by vials.mix_practice_id")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getJSON(sqlString string) (string, error) {
	db, err := sql.Open("mysql", "vagrant:vagrant@tcp(192.168.56.22:3306)/AllergyNew")
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
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	db.Close()
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
