package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostegresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostegresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func DefaultPostgresConfig() PostegresConfig {
	return PostegresConfig{
		Host:     "localhost",
		Port:     "5433",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

func Open(config PostegresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
