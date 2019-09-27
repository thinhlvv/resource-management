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
		CreateUser(user model.User) (int, error)
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

// CreateUser creates new user.
func (repo *RepoImpl) CreateUser(user model.User) (int, error) {
	return repo.createUser(user)
}

// GetUserByEmail find user with email.
func (repo *RepoImpl) GetUserByEmail(email string) (*model.User, error) {
	return repo.getUserByEmail(email)
}

func (repo *RepoImpl) createUser(user model.User) (int, error) {
	stmt, err := repo.db.Prepare(`
		INSERT INTO	user(email, role, hashed_password, quota)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Email, user.Role.Int(), user.HashedPassword, user.Quota)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
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
