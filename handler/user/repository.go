package user

import (
	"database/sql"

	"github.com/thinhlvv/resource-management/model"
)

type (
	// Repository is interface to interact with outside world.
	Repository interface {
		// user
		GetUserByEmail(email string) (*model.User, error)
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

// GetUserByEmail find user with email.
func (repo *RepoImpl) GetUserByEmail(email string) (*model.User, error) {
	return repo.getUserByEmail(email)
}

func (repo *RepoImpl) getUserByEmail(email string) (*model.User, error) {
	query := repo.db.QueryRow(`
		SELECT	id, email, hashed_password, quota, role
		FROM		user
		WHERE		email = ?
		LIMIT		1
	`, email)

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
