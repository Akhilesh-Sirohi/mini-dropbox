package config

import (
	"github.com/mini-dropbox/app/providers"
	"os"

	"github.com/mini-dropbox/app/config/reader"
)

var cfg Config

type Config struct {
	App     App
	Redis   RedisConfig
	Session Session
	Db      providers.DbConfig
}

// App contains application-specific config values
type App struct {
	Env             string
	ServiceName     string
	Hostname        string
	Port            string
	ShutdownTimeout int
	ShutdownDelay   int
	GitCommitHash   string
}

type Session struct {
	AuthenticationKey string
	EncryptionKey     string
}

type RedisConfig struct {
	Host     string
	Port     int32
	Database int32
	Password string
}

func Init(env string) error {
	// Init config
	err := reader.NewDefaultConfig().Load(env, &cfg)
	if err != nil {
		return err
	}

	// Puts git commit hash into config.
	// This is not read automatically because env variable is not in expected format.
	cfg.App.GitCommitHash = os.Getenv("GIT_COMMIT_HASH")

	return nil
}

func GetConfig() Config {
	return cfg
}
