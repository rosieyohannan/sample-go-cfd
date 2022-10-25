package openapi

import (
	"fmt"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (Cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Cfg.Host, Cfg.Port, Cfg.User, Cfg.Password, Cfg.Database, Cfg.SSLMode)

}

func OpenDb() error {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "cfd",
		SSLMode:  "disable",
	}

	var err error
	DB, err = gorm.Open(postgres.Open(cfg.String()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	}

	return nil
}
