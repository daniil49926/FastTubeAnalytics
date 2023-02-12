package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InsertStatements(requestData []string, responseData []string) error {

	var sqlForInsert = `
	insert into analytics_data(request_method, request_url, request_process_time, response_status_code) 
	values (?, ?, ?, ?)
	`

	db, err := sql.Open("sqlite3", "analytics.db")

	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	db.QueryRow(sqlForInsert, requestData[0], requestData[1], requestData[2], responseData[0])

	return nil
}

func ReadAllStatements() ([][]string, error) {
	var sqlForSelect = `select * from analytics_data`

	db, err := sql.Open("sqlite3", "analytics.db")

	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	rows, err := db.Query(sqlForSelect)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var dataToRet [][]string
	for rows.Next() {
		var id int
		var requestMethod string
		var requestUrl string
		var requestProcessTime string
		var responseStatusCode string

		err = rows.Scan(&id, &requestMethod, &requestUrl, &requestProcessTime, &responseStatusCode)
		if err != nil {
			return nil, err
		}

		dataToRet = append(
			dataToRet,
			[]string{
				string(rune(id)),
				requestMethod,
				requestUrl,
				requestProcessTime,
				responseStatusCode,
			})
	}
	return dataToRet, nil
}
