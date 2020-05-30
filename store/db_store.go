package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/square/squalor"
)

var (
	db    *squalor.DB
	cache = make(map[string]bool)
)

func init() {
	const driver string = "postgres"
	dataSourceName := DataSourceName()
	sdb, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panicOnError(err)
	}
	db, _ = squalor.NewDB(sdb)
}

func Insert(bind string, model interface{}) error {
	//cannot bind twice .. check if bound already , else skip this
	if val := cache[bind]; !val {
		_, err := db.BindModel(bind, model)
		if err != nil {
			return err
		}
		cache[bind] = true
	}
	err := db.Insert(model)
	return err
}

func Fetch(query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := db.Query(query, args...) // Note: Ignoring errors for brevity
	if err != nil {
		return nil, err
	}
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		return m, nil
	}
	return nil, nil
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
