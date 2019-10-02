package user

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/leebenson/conform"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
)

// ErrInvalidEmailOrPassword ...
var ErrInvalidEmailOrPassword = errors.New("invalid email or password")

type (
	// Service is interface of user service.
	Service interface {
		Login(ctx echo.Context, req LoginRequest) (string, error)
		Signup(ctx echo.Context, req SignupReq) (string, error)

		// user
		Create(CreateUserReq) (int, error)
		Update(UpdateUserReq) error
		Delete(int) error
		GetAll() ([]model.User, error)
	}
	// ServiceImpl represents service implementation of service.
	ServiceImpl struct {
		repo   Repository
		hasher pkg.Hasher
		signer pkg.Signer
	}
)

// NewService returns service implementation of user service.
func NewService(repo Repository, app model.App) Service {
	return &ServiceImpl{
		repo:   repo,
		hasher: app.Hasher,
		signer: app.JWTSigner,
	}
}

// Signup registers new user.
func (s *ServiceImpl) Signup(ctx echo.Context, req SignupReq) (string, error) {
	conform.Strings(&req)

	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return "", err
	}

	newUser := model.User{
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Role:           model.UserRole(req.Role),
		Quota:          model.UnlimitedQuota.Int(),
	}
	id, err := s.repo.CreateUser(newUser)
	if err != nil {
		return "", err
	}

	accessToken, err := s.signer.SignWithRole(strconv.Itoa(id), newUser.Role.Int())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// Login is logic of login use case.
func (s *ServiceImpl) Login(ctx echo.Context, req LoginRequest) (string, error) {
	conform.Strings(&req)

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	// Must be admin or owner, and must be verified.
	if _, existing := model.Roles[user.Role]; !existing {
		return "", errors.New("unauthorized role")
	}

	// Compare the password with the hashed password.
	match, err := s.hasher.Compare(req.Password, user.HashedPassword)
	if err != nil || !match {
		return "", ErrInvalidEmailOrPassword
	}

	accessToken, err := s.signer.SignWithRole(strconv.Itoa(user.ID), user.Role.Int())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// Create new user.
func (s *ServiceImpl) Create(req CreateUserReq) (int, error) {
	quota := model.UnlimitedQuota.Int()
	if req.Quota != 0 {
		quota = req.Quota
	}

	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return 0, err
	}

	u := model.User{
		Email:          req.Email,
		Quota:          quota,
		HashedPassword: hashedPassword,
	}
	return s.repo.CreateUser(u)
}

// Update user's information.
func (s *ServiceImpl) Update(req UpdateUserReq) error {
	u := model.User{
		ID:    req.ID,
		Email: req.Email,
		Quota: req.Quota,
	}
	return s.repo.UpdateUser(u)
}

// Delete user by ID ...
func (s *ServiceImpl) Delete(id int) error {
	return s.repo.SoftDeleteUser(id)
}

// GetAll users.
func (s *ServiceImpl) GetAll() ([]model.User, error) {
	return s.repo.GetAllUser()
}
