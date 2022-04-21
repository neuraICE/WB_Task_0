package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "manager"
	DB_PASSWORD = "password"
	DB_NAME     = "testdb"
)

var connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
var M = make(map[string][]byte)

func InsertInDB(id string, b []byte) error {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO testtable (order_uid, data) VALUES ($1, $2);", id, b)
	if err != nil {
		return err
	}

	//fmt.Println(result.RowsAffected())

	return nil
}

func SelectInDB() error {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("select * from testtable;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data []byte
		err := rows.Scan(&id, &data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		M[id] = data
	}
	return nil
}
