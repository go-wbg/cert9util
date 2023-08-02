package main

import (
	"log"

	cert9util "github.com/eyedeekay/cert9util/lib"
)

func main() {
	//var dbtest *cert9util.SQLiteDB
	dbtest, err := cert9util.NewSQLiteDB("cert9.db")
	if err != nil {
		panic(err)
	}
	rows, err := dbtest.DB.Query("SELECT * FROM nssPublic")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var row string
		rows.Scan(&row)
		log.Println(row)
	}
	defer rows.Close()
	defer dbtest.Close()
}
