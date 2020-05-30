package store

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	DbUserNameKey = "database.user"
	DbPwdKey      = "database.pwd"
	DbHostKey     = "database.host"
)

type DatabaseConfig struct {
	User   string
	Pwd    string
	Host   string
}

var dbConfig DatabaseConfig

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		//we cant continue if the resource is missing as we
		//wont be able to connect to the database
		panic(err.Error())
	}

	dbConfig.User = viper.GetString(DbUserNameKey)
	dbConfig.Pwd = viper.GetString(DbPwdKey)
	dbConfig.Host = viper.GetString(DbHostKey)
}

func DataSourceName() string {
	return dbConfig.String()
}

func (dc *DatabaseConfig) String() string {
	//postgres://postgres:root@localhost/amakhosi_pumps?sslmode=disable
	return fmt.Sprintf("postgres://%s:%s/%s/amakhosi_pumps?sslmode=disable", dc.User, dc.Pwd, dc.Host)
}

