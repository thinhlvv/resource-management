package testhelper

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/thinhlvv/blog-system/config"
)

// NewDB returns db to test.
func NewDB() *sql.DB {
	cfg := config.New()
	db, err := sql.Open("mysql", connStr(cfg))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func connStr(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.MysqlTest.User,
		cfg.MysqlTest.Password,
		cfg.MysqlTest.Host,
		cfg.MysqlTest.Name)
}
