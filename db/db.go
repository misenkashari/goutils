package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config holds the settings for connecting to a database.
type Config struct {
	Driver   string // e.g., "postgres", "mysql"
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // for Postgres: "disable"/"require"; for MySQL: tls parameter
}

// Builder provides a fluent API for constructing DBConfig and opening the connection.
type Builder struct {
	cfg Config
}

func MySQL() *Builder {
	return &Builder{
		cfg: Config{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			User:     "",
			Password: "",
			DBName:   "",
			SSLMode:  "false",
		},
	}
}

func Postgres() *Builder {
	return &Builder{
		cfg: Config{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     5432,
			User:     "",
			Password: "",
			DBName:   "",
			SSLMode:  "disable",
		},
	}
}

func (b *Builder) WithHost(host string) *Builder {
	b.cfg.Host = host
	return b
}

func (b *Builder) WithPort(port int) *Builder {
	b.cfg.Port = port
	return b
}

func (b *Builder) WithUser(user string) *Builder {
	b.cfg.User = user
	return b
}

func (b *Builder) WithPassword(pw string) *Builder {
	b.cfg.Password = pw
	return b
}

func (b *Builder) WithDatabase(name string) *Builder {
	b.cfg.DBName = name
	return b
}

func (b *Builder) WithSSLMode(mode string) *Builder {
	b.cfg.SSLMode = mode
	return b
}

// Open constructs the DSN based on the driver and opens a GORM connection.
func (b *Builder) Open() (*gorm.DB, error) {
	var dsn string
	switch b.cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			b.cfg.Host, b.cfg.Port, b.cfg.User, b.cfg.Password, b.cfg.DBName, b.cfg.SSLMode,
		)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "mysql":
		// e.g. user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local&tls=custom
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
			b.cfg.User, b.cfg.Password, b.cfg.Host, b.cfg.Port, b.cfg.DBName, b.cfg.SSLMode,
		)
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})

	default:
		return nil, fmt.Errorf("unsupported driver: %s", b.cfg.Driver)
	}
}
