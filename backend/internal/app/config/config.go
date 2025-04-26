package config

import (
	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	HTTPServer HTTPServerConfig `env:"HTTP_SERVER"`
	DB         DBConfig         `env:"DB"`
	JWT		   JWTConfig        `env:"JWT"`
}

type HTTPServerConfig struct {
	Port string `env:"PORTHTTP" envDefault:"8080"`
}

type DBConfig struct {
	Host     string `env:"HOSTdb" envDefault:"localhost"`
	Port     string `env:"PORTdb" envDefault:"5432"`
	User     string `env:"USERdb" envDefault:"postgres"`
	Password string `env:"PASSWORDdb" envDefault:"postgres"`
	DBName   string `env:"DBNamedb" envDefault:"worldmap"`
	SSLMode  string `env:"SSL_MODEdb" envDefault:"disable"`
}

type JWTConfig struct {
	Secret              string `env:"SECRETTEST" envDefault:"supersecretkey"`            // access token
	ExpiredHours        int    `env:"EXPIRED_HOURSTEST" envDefault:"1"`                // access token lifetime
	// RefreshSecret       string `env:"REFRESH_SECRET" envDefault:"superrefreshsecret"`   // refresh token
	// RefreshExpiresHours int    `env:"REFRESH_EXPIRED_HOURS" envDefault:"168"`           // refresh token lifetime (7 дней)
}

func NewConfig(filenames ...string) (*Config, error) {
	
	log.Println("Config file loaded: 1", filenames)

	if len(filenames) > 0 && filenames[0] != "" {
		log.Println("Loading config file:", filenames[0])
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}
	
	log.Println("Config file loaded: 2", filenames)

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully")

	return cfg, nil
}

// for connections.go
func (c *DBConfig) GetDBConnString() string {
	return "host=" + c.Host +
		" port=" + c.Port +
		" user=" + c.User +
		" password=" + c.Password +
		" dbname=" + c.DBName +
		" sslmode=" + c.SSLMode
}