package orm

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	velora "github.com/aryan-salemababdi/Velora/app"
)

func InitDB(cfg *velora.Config) *gorm.DB {
	user := cfg.Get("DATABASE_USER")
	pass := url.QueryEscape(cfg.Get("DATABASE_PASSWORD"))
	host := cfg.Get("DATABASE_HOST")
	port := cfg.GetInt("DATABASE_PORT")
	dbname := cfg.Get("DATABASE_NAME")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, pass, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect database: %v", err)
	}

	fmt.Printf("✅ Connected to PostgreSQL at %s:%d, database: %s\n", host, port, dbname)
	return db
}
