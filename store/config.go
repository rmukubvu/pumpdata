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
	AccountSid    = "twilio.sid"
	AuthToken     = "twilio.token"
	Number        = "twilio.number"
	MongoUrl      = "mongodatabase.uri"
	RabbitUrl     = "rabbit.url"
	SendGridApi   = "sendgrid.api"
	SendGridFrom  = "sendgrid.from"
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

type TwilioConfig struct {
	Sid    string
	Token  string
	Number string
}

type MongoConfig struct {
	Url string
}

type RabbitConfig struct {
	Url string
}

type SendGridConfig struct {
	ApiKey string
	From   string
}

var (
	dbConfig       DatabaseConfig
	serverConfig   WebServerConfig
	twilioConfig   TwilioConfig
	mongoConfig    MongoConfig
	rabbitConfig   RabbitConfig
	sendGridConfig SendGridConfig
)

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

	twilioConfig.Sid = viper.GetString(AccountSid)
	twilioConfig.Token = viper.GetString(AuthToken)
	twilioConfig.Number = viper.GetString(Number)

	mongoConfig.Url = viper.GetString(MongoUrl)
	rabbitConfig.Url = viper.GetString(RabbitUrl)
	sendGridConfig.ApiKey = viper.GetString(SendGridApi)
	sendGridConfig.From = viper.GetString(SendGridFrom)
}

func TwilioSmsConfig() *TwilioConfig {
	return &twilioConfig
}

func MongoUrlConfig() *MongoConfig {
	return &mongoConfig
}

func RabbitUrlConfig() *RabbitConfig {
	return &rabbitConfig
}

func SendGrid() *SendGridConfig {
	return &sendGridConfig
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
