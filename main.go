package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/config"
	"github.com/thinhlvv/resource-management/handler/user"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
)

func main() {
	cfg := config.New()

	db := config.MustInitDB(cfg)
	defer db.Close()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// signer
	signer := config.MustInitJWTSigner(cfg)
	requestValidator := pkg.NewRequestValidator()

	// app
	app := model.App{
		DB:               db,
		RequestValidator: requestValidator,
		JWTSigner:        signer,
	}

	e := echo.New()

	// User service
	{
		userHandler := user.New(app)
		userHandler.RegisterHTTPRouter(e)
	}

	s := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}
