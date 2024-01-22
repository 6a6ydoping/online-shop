package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"time"
)

type Config struct {
	HTTP   ServerConfig `yaml:"http"`
	DB     DBConfig     `yaml:"db"`
	Token  TokenConfig  `yaml:"token"`
	Router RouterConfig `yaml:"router"`
}

type TokenConfig struct {
	SecretKey  string        `env:"TOKEN_SECRET_KEY" env-required:"true"`
	TimeToLive time.Duration `yaml:"time_to_live"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

//type DBConfig struct {
//	Host             string `yaml:"host"`
//	Port             string `yaml:"port"`
//	DBName           string `yaml:"db_name"`
//	Username         string `yaml:"username"`
//	MigrationPath    string `yaml:"migration_path"`
//	MigrationVersion uint   `yaml:"migration_version"`
//	Password         string `env:"DB_PASSWORD" env-required:"true"`
//}

type DBConfig struct {
	ConnectionURI string `yaml:"connection_string"`
}

// RouterConfig TODO: full config + prod/debug modes for server
type RouterConfig struct {
	AppName       string `yaml:"app_name"`
	CaseSensitive bool   `yaml:"case_sensitive"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = cleanenv.ReadEnv(&cfg.DB)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
