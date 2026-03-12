package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Env struct {
	AppEnv   string `mapstructure:"APP_ENV"`
	Port     string `mapstructure:"PORT"`
	GrcpPort string `mapstructure:"GRCP_PORT"`

	PGHost string `mapstructure:"PG_HOST"`
	PGPort string `mapstructure:"PG_PORT"`
	PGUser string `mapstructure:"PG_USER"`
	PGPass string `mapstructure:"PG_PASS"`
	PGName string `mapstructure:"PG_NAME"`

	NatsHost string `mapstructure:"NATS_HOST"`
	NatsPort string `mapstructure:"NATS_PORT"`
	NatsName string `mapstructure:"NATS_NAME"`
	NatsChan string `mapstructure:"NATS_CHAN"`

	AppDomain string `mapstructure:"APP_DOMAIN"`
	FileUrl   string `mapstructure:"FILE_URL"`

	MigrationPath string `mapstructure:"MIGRATION_PATH"`

	JwtSecret string `mapstructure:"JWT_SECRET"`

	AuthServiceGrpcAddress string `mapstructure:"GRPC_AUTH_ADR"`

	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
}

func NewEnv() Env {
	env := Env{}

	_, err := os.Stat(".env")
	useEnvFile := !os.IsNotExist(err)

	if useEnvFile {
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't read the .env file: ", err)
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
	} else {
		env.bindEnv()
	}

	if env.AppEnv != "production" {
		log.Println("The App is running in development env")
	}

	return env
}

func (e *Env) bindEnv() {
	e.AppEnv = os.Getenv("APP_ENV")
	e.Port = os.Getenv("PORT")
	e.GrcpPort = os.Getenv("GRCP_PORT")

	e.PGHost = os.Getenv("PG_HOST")
	e.PGPort = os.Getenv("PG_PORT")
	e.PGUser = os.Getenv("PG_USER")
	e.PGPass = os.Getenv("PG_PASS")
	e.PGName = os.Getenv("PG_NAME")

	e.NatsHost = os.Getenv("NATS_HOST")
	e.NatsPort = os.Getenv("NATS_PORT")
	e.NatsName = os.Getenv("NATS_NAME")
	e.NatsChan = os.Getenv("NATS_CHAN")

	e.AppDomain = os.Getenv("APP_DOMAIN")
	e.FileUrl = os.Getenv("FILE_URL")

	e.JwtSecret = os.Getenv("JWT_SECRET")

	e.MigrationPath = os.Getenv("MIGRATION_PATH")

	e.AuthServiceGrpcAddress = os.Getenv("GRPC_AUTH_ADR")

	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {
		e.AllowedOrigins = strings.Split(val, ",")
	}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)
