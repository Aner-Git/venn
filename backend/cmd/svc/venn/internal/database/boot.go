package database

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	cfgmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(config *MySQLConfig) (*GormDatabase, error) {
	db, err := connect(config)

	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

func connect(config *MySQLConfig) (*gorm.DB, error) {
	if config.User == "" || config.Host == "" || config.Name == "" {
		return nil, errors.New("missing one or more of user, host, or name for db config")
	}

	if config.Port == 0 {
		config.Port = 3306
	}

	cfg := cfgmysql.NewConfig()
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	cfg.DBName = config.Name
	cfg.Collation = "utf8mb4_unicode_ci"
	cfg.Loc = time.UTC
	cfg.ParseTime = true
	cfg.InterpolateParams = true
	cfg.Params = map[string]string{
		"sql_notes": "0",
		"charset":   "utf8mb4",
	}

	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()))
	return db, err
}

func ConfigMySQL(c *MySQLConfig) {

	c.Name = setFromEnv(c.Name, "DB_NAME")
	c.Host = setFromEnv(c.Host, "DB_HOST")
	c.Port = setFromEnvInt(c.Port, "DB_PORT")
	c.User = setFromEnv(c.User, "DB_USER")
	c.Password = setFromEnv(c.Password, "DB_PASSWORD")
	c.TLS = setFromEnv(c.TLS, "DB_TLS")
}

// MySQLConfig Configuration
type MySQLConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	TLS      string
}

func (c *MySQLConfig) SetDefaults(serviceName string) {
	c.Port = 3306
	c.TLS = "false"
	c.Name = serviceName
	c.Host = serviceName
	c.User = serviceName
	c.Password = serviceName
	c.Host = serviceName + "-mysql"
}

func setFromEnv(defaultValue, envKeyName string) string {
	if v, ok := os.LookupEnv(envKeyName); ok {
		return v
	}
	return defaultValue
}

func setFromEnvInt(defaultValue int, envKeyName string) int {

	val, err := strconv.Atoi(setFromEnv(strconv.Itoa(defaultValue), envKeyName))
	if err != nil {
		panic(fmt.Sprintf("%s is not an integer: %s", envKeyName, err))
	}
	return val
}

func setFromEnvBool(defaultValue bool, envKeyName string) bool {
	val := defaultValue
	yesno := setFromEnv(fmt.Sprintf("%t", defaultValue), envKeyName)

	switch yesno {
	case "1", "true", "TRUE":
		val = true
	case "0", "false", "FALSE":
		val = false
	default:
		panic(fmt.Sprintf("%s is not a bool: %s", envKeyName, yesno))

	}

	return val
}
