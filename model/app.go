package model

import (
	"database/sql"

	"github.com/thinhlvv/blog-system/pkg"
)

// App represents independent libraries
type App struct {
	DB               *sql.DB
	RequestValidator pkg.RequestValidator
}
