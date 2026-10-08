package conf

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type AppConfig interface {
	GetDbDSN() string
	GetLogLevel() string
}
type cfg struct {
	DbDSN    string `envconfig:"DB_DSN" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

var _ AppConfig = (*cfg)(nil)

func NewAppConfig() (AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "godotenv.Load() error")
	}

	var conf = new(cfg)
	if err := envconfig.Process("", conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *cfg) GetDbDSN() string {
	return c.DbDSN
}

func (c *cfg) GetLogLevel() string {
	return strings.ToUpper(c.LogLevel)
}
