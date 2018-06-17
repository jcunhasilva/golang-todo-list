package config

import (
	"strings"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

var (
	dbConfig *viper.Viper
	// DBConnection database connection details
	DBConnection *databaseConnection
)

func init() {
	replacer := strings.NewReplacer(".", "_")
	dbConfig = viper.New()
	dbConfig.SetEnvKeyReplacer(replacer)
	dbConfig.SetEnvPrefix("TODO_DATABASE")
	dbConfig.BindEnv("host")
	dbConfig.BindEnv("user")
	dbConfig.BindEnv("password")
	dbConfig.BindEnv("name")

	DBConnection = newConnection()
}

// Database represents general api configuration
type Database struct {
	Host     string
	User     string
	Password string
	Name     string
}

type databaseConnection struct {
	config Database
	Client *pg.DB
}

// LoadConfig loads the database configuration from env variables
func LoadConfig() (Database, error) {
	var config Database
	if err := dbConfig.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func newConnection() *databaseConnection {
	config, _ := LoadConfig()
	client := pg.Connect(&pg.Options{
		Addr:     config.Host,
		User:     config.User,
		Password: config.Password,
		Database: config.Name,
	})

	return &databaseConnection{config: config, Client: client}

}
