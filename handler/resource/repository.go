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
		GetResourcesByUserID(int) ([]model.Resource, error)
		GetAll() ([]model.Resource, error)
		SoftDelete(id int) error
		ResourceBelongsUser(resourceID, userID int) bool

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

// GetResourcesByUserID ...
func (repo *RepoImpl) GetResourcesByUserID(userID int) ([]model.Resource, error) {
	return repo.getResourcesByUserID(userID)
}

// GetAll ...
func (repo *RepoImpl) GetAll() ([]model.Resource, error) {
	return repo.getAll()
}

// SoftDelete ...
func (repo *RepoImpl) SoftDelete(id int) error {
	return repo.softDelete(id)
}

// ResourceBelongsUser ...
func (repo *RepoImpl) ResourceBelongsUser(resourceID, userID int) bool {
	return repo.resourceBelongsUser(resourceID, userID)
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

func (repo *RepoImpl) getResourcesByUserID(userID int) ([]model.Resource, error) {
	stmt := `
		SELECT		id, name
		FROM			resource
		WHERE			user_id = ? AND deleted_at IS NULL
	`
	rows, err := repo.db.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Resource
	for rows.Next() {
		var resource model.Resource
		err = rows.Scan(
			&resource.ID,
			&resource.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, resource)
	}
	err = rows.Err()

	return result, err
}

func (repo *RepoImpl) getAll() ([]model.Resource, error) {
	stmt := `
		SELECT	id, name
		FROM		resource
		WHERE		deleted_at IS NULL
	`
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Resource
	for rows.Next() {
		var resource model.Resource
		err = rows.Scan(
			&resource.ID,
			&resource.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, resource)
	}
	err = rows.Err()

	return result, err
}

func (repo *RepoImpl) softDelete(id int) error {
	stmt, err := repo.db.Prepare(`
		UPDATE resource
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

func (repo *RepoImpl) resourceBelongsUser(resourceID, userID int) bool {
	query := repo.db.QueryRow(`
		SELECT	id
		FROM		resource
		WHERE		id = ? AND user = ?
		LIMIT		1
	`, resourceID, userID)

	var r model.Resource
	if err := query.Scan(
		&r.ID,
	); err != nil {
		return false
	}

	return true
}
