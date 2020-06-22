package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db    *sqlx.DB
	cache = make(map[int]string)
)

func init() {
	const driverName string = "postgres"
	var err error
	db, err = sqlx.Connect(driverName, dataSourceName())
	if err != nil {
		panicOnError(err)
	}
}

func insert(sql string, arg map[string]interface{}) error {
	_, err := db.NamedExec(sql, arg)
	return err
}

func insertWithLastId(sql string, arg map[string]interface{}) (int64, error) {
	rs, _ := db.NamedExec(sql, arg)
	return rs.LastInsertId()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func CloseDBConnection() {
	if db != nil {
		db.Close()
	}
}
