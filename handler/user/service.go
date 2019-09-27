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
