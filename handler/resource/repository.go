package resource

import (
	"database/sql"

	"github.com/thinhlvv/resource-management/model"
)

type (
	// Repository is interface to interact with outside world.
	Repository interface {
		// resource
		CreateResource(model.Resource) (int, error)
	}
	// RepoImpl is Repository's Implementation.
	RepoImpl struct {
		db *sql.DB
	}
)

// NewRepo returns repository implementation.
func NewRepo(db *sql.DB) Repository {
	return &RepoImpl{db: db}
}

// CreateResource creates new resource.
func (repo *RepoImpl) CreateResource(resource model.Resource) (int, error) {
	return repo.createResource(resource)
}

func (repo *RepoImpl) createResource(resource model.Resource) (int, error) {
	stmt, err := repo.db.Prepare(`
		INSERT INTO resource(name, user_id)
		VALUES (?, ?)
	`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(resource.Name, resource.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}
