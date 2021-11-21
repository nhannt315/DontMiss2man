package db

import (
	"encoding/json"
	"time"

	pkgstrings "github.com/nhannt315/real_estate_api/pkg/strings"
)

const (
	// driverMySQL is the value of MySQL driver
	// of Driver field in DBSettings
	driverMySQL = "mysql"
)

// Config contains all configs for DB layer
type Config struct {
	Driver            string                  `required:"true" default:"mysql" yaml:"driver" `
	Host              string                  `default:"127.0.0.1" yaml:"host" env:"DB_HOST"` // yaml for test only
	Port              uint                    `default:"3306" yaml:"port" env:"DB_PORT"`
	Location          string                  `required:"true" default:"auto" yaml:"location"`
	ParseTime         bool                    `default:"true" yaml:"parse_time"`
	Schema            string                  `required:"true" default:"real_estate" yaml:"schema"`
	TLS               string                  `default:"skip-verify" yaml:"tls"`
	CACertPath        string                  `yaml:"ca_cert_path" env:"DB_CA_CERT_PATH"`
	Collation         string                  `default:"utf8mb4_bin" yaml:"collation"`
	InterpolateParams bool                    `yaml:"interpolate_params"`
	Username          pkgstrings.MaskedString `yaml:"username" env:"DB_USERNAME"` // yaml for test only
	Password          pkgstrings.MaskedString `yaml:"password" env:"DB_PASSWORD"` // yaml for test only
	MaxOpenConn       uint64                  `default:"1000" yaml:"max_open_conn"`
	MaxIdleConn       uint64                  `default:"1000" yaml:"max_idle_conn"`
	MaxAllowedPacket  int                     `default:"33554432" yaml:"max_allowed_packet"` // bytes
	ConnMaxLifetime   time.Duration           `default:"60s" yaml:"conn_max_lifetime" `      // seconds 	// other params
	RetryInterval     time.Duration           `default:"5s" yaml:"retry_interval"`           // retry interval when getting connection to db
	MaxRetryCount     uint64                  `default:"1000" yaml:"max_retry_count"`        // max retry count when getting connection to db
	info              string                  // info caching
}

// Info returns db info without username and password (to avoid displaying in log, for example)
func (d *Config) Info() string {
	if d.info != "" { // info is already calculated before
		return d.info
	}
	bytes, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	d.info = string(bytes)
	return d.info
}

// Clone returns brand-new DB COnfig without cached data.
func (d *Config) Clone() *Config {
	c := *d
	c.info = ""
	return &c
}
