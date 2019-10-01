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

		// user
		GetUserByID(int) (*model.User, error)
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

// GetUserByID find user with email.
func (repo *RepoImpl) GetUserByID(id int) (*model.User, error) {
	return repo.getUserByID(id)
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

func (repo *RepoImpl) getUserByID(id int) (*model.User, error) {
	query := repo.db.QueryRow(`
		SELECT	id, email, hashed_password, quota, role
		FROM		user
		WHERE		id = ?
		LIMIT		1
	`, id)

	var u model.User
	err := query.Scan(
		&u.ID,
		&u.Email,
		&u.HashedPassword,
		&u.Quota,
		&u.Role,
	)

	return &u, err
}
