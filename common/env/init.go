package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env string

// kek buat type string literal di typescript (buat 2 variabel dengan namo EnvProd dan EnvDev)
const (
	EnvProd Env = "PRODUCTION"
	EnvDev  Env = "DEVELOPMENT"
)

type Config struct {
	Env         Env    `mapstructure:"ENV"`
	AppHost     string `mapstructure:"APP_HOST"`
	AppPort     string `mapstructure:"APP_PORT"`
	PostgresUrl string `mapstructure:"POSTGRES_CONNECTION_URL"`

	// untuk cors siapo be yg boleh boleh akses (utk sekarang dak dipake) [cek be ck mano agek caro setting origin di gin gin-contrib/cors]
	AllowedOrigin string `mapstructure:"ORIGIN"`
}

func LoadConfig(envPath string) Config {
	var config Config

	viper.SetConfigFile(envPath)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error occured while .env reading: %v", err)
	}

	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("ENV", EnvDev)

	viper.Unmarshal(&config)

	return config
}
