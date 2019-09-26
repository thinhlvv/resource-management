package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// import init mysql
	_ "github.com/go-sql-driver/mysql"
)

// MustInitDB returns DB pointer.
func MustInitDB(cfg *Config) *sql.DB {
	var doOnce sync.Once
	var db *sql.DB
	var err error

	doOnce.Do(func() {
		db, err = sql.Open("mysql", conStr(cfg))
		if err != nil {
			log.Fatal(err)
		}

		// Ping to check connection
		var connErr error
		for i := 0; i < 3; i++ {
			connErr = db.Ping()
			if connErr != nil {
				log.Fatal("Can not init database:", connErr)
			}
			time.Sleep(1 * time.Second)
		}
	})

	return db
}

func conStr(cfg *Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Name)
}
