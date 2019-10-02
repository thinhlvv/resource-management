package user

import (
	"database/sql"

	"github.com/thinhlvv/resource-management/model"
)

type (
	// Repository is interface to interact with outside world.
	Repository interface {
		GetUserByEmail(email string) (*model.User, error)
		CreateUser(user model.User) (int, error)
		UpdateUser(model.User) error
		SoftDeleteUser(int) error
		GetAllUser() ([]model.User, error)
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

// Update user's info (email, quota).
func (repo *RepoImpl) UpdateUser(u model.User) error {
	return repo.update(u)
}

// SoftDelete updates deleted_at.
func (repo *RepoImpl) SoftDeleteUser(id int) error {
	return repo.softDelete(id)
}

// GetAll returns all users.
func (repo *RepoImpl) GetAllUser() ([]model.User, error) {
	return repo.getAll()
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

func (repo *RepoImpl) softDelete(id int) error {
	stmt, err := repo.db.Prepare(`
		UPDATE user
		SET deleted_at = now() 
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return err
}

func (repo *RepoImpl) update(u model.User) error {
	stmt, err := repo.db.Prepare(`
		UPDATE user
		SET email = ?, quota = ? 
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Quota, u.ID)
	if err != nil {
		return err
	}

	return err
}

func (repo *RepoImpl) getAll() ([]model.User, error) {
	stmt := `
		SELECT		id, email, quota, role
		FROM			user
		WHERE			deleted_at IS NULL
	`
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.User
	for rows.Next() {
		var u model.User
		err = rows.Scan(
			&u.ID,
			&u.Email,
			&u.Quota,
			&u.Role,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	err = rows.Err()

	return result, err
}
