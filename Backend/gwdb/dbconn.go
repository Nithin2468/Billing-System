package gwdb

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

// Config structs
type Config struct {
	Database DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"password"`
	Name string `toml:"name"`
}

var GDbconn *sql.DB

func LoadConfig() DatabaseConfig {
	configPath := filepath.Join("config", "config.toml")

	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	return cfg.Database
}

func DbConnection() (*sql.DB, error) {
	cfg := LoadConfig()
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)

	lDb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open MySQL connection:", err)
		return lDb, err
	}

	if err := lDb.Ping(); err != nil {
		log.Fatal("Failed to ping MySQL:", err)
	}

	log.Println("MySQL connected successfully!")
	return lDb, nil
}
