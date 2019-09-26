package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/config"
)

func main() {
	cfg := config.New()

	db := config.MustInitDB(cfg)
	defer db.Close()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	e := echo.New()

	s := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
