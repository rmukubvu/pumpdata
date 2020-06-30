package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
	"time"
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

func GetCreatedDate() int {
	now := time.Now()
	year, month, day := now.Date()
	res := fmt.Sprintf("%d%d%d", year, int(month), day)
	date, _ := strconv.Atoi(res)
	return date
}

func GetDateFromString(input string) int {
	//date format from service will be yyyy-MM-dd
	s := strings.ReplaceAll(input, "-", "")
	date, _ := strconv.Atoi(s)
	return date
}

func CloseDBConnection() {
	if db != nil {
		db.Close()
	}
}
