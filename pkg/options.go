package pkg

import (
	"github.com/spf13/viper"
	"strings"
)

type ExvalOptions struct {
	Address                  string
	Port                     int
	PostgresConnectionString string
	AuthUser                 string
	AuthPassword             string

	AutoMigration bool
}

func DefaultOptions() *ExvalOptions {
	return &ExvalOptions{
		Address:                  "0.0.0.0",
		Port:                     11139,
		PostgresConnectionString: "host=localhost user=postgres password=postgres port=5432",
		AuthUser:                 "admin",
		AuthPassword:             "admin",
		AutoMigration:            false,
	}
}

func SetupConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	options := DefaultOptions()

	viper.SetDefault("address", options.Address)
	viper.SetDefault("port", options.Port)
	viper.SetDefault("postgres-connection-string", options.PostgresConnectionString)
	viper.SetDefault("auto-migration", options.AutoMigration)
	viper.SetDefault("auth-user", options.AuthUser)
	viper.SetDefault("auth-password", options.AuthPassword)
}

func ParseOptions() (*ExvalOptions, error) {
	options := DefaultOptions()

	options.Address = viper.GetString("address")
	options.Port = viper.GetInt("port")
	options.PostgresConnectionString = viper.GetString("postgres-connection-string")
	options.AutoMigration = viper.GetBool("auto-migration")
	options.AuthUser = viper.GetString("auth-user")
	options.AuthPassword = viper.GetString("auth-password")

	return options, nil
}
