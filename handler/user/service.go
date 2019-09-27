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

// Login logs the admin into the dashboard.
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
