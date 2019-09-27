package model

import (
	"database/sql"

	"github.com/thinhlvv/resource-management/pkg"
)

// App represents independent libraries
type App struct {
	DB               *sql.DB
	RequestValidator pkg.RequestValidator
	JWTSigner        pkg.Signer
	Hasher           pkg.Hasher
}
