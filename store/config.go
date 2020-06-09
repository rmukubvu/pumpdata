package store

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	DbUserNameKey = "system.data.postgres.user"
	DbPwdKey      = "system.data.postgres.pwd"
	DbHostKey     = "system.data.postgres.host"
	DbPortKey     = "system.data.postgres.port"
	DbNameKey     = "system.data.postgres.database"
	DbUrlKey      = "system.data.postgres.url"
	ServerPort    = "webserver.port"
)

type DatabaseConfig struct {
	User      string
	Pwd       string
	Host      string
	Port      int
	DbName    string
	UrlFormat string
}

type WebServerConfig struct {
	Port int
}

var dbConfig DatabaseConfig
var serverConfig WebServerConfig

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	dbConfig.User = viper.GetString(DbUserNameKey)
	dbConfig.Pwd = viper.GetString(DbPwdKey)
	dbConfig.Host = viper.GetString(DbHostKey)
	dbConfig.Port = viper.GetInt(DbPortKey)
	dbConfig.DbName = viper.GetString(DbNameKey)
	dbConfig.UrlFormat = viper.GetString(DbUrlKey)

	serverConfig.Port = viper.GetInt(ServerPort)
}

func dataSourceName() string {
	return dbConfig.String()
}

func WebServerPort() int {
	return serverConfig.Port
}

func (dc *DatabaseConfig) String() string {
	return fmt.Sprintf(dc.UrlFormat, dc.User, dc.Pwd, dc.Host, dc.Port, dc.DbName)
}
