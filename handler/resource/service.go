package resource

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/model"
)

// ErrNotEnoughQuota ...
var ErrNotEnoughQuota = errors.New("You don't have enough quota to create new resource")

type (
	// Service is interface of user service.
	Service interface {
		CreateResource(req CreateReq) (*model.Resource, error)
		GetResourcesOfUser(userID, role int) ([]model.Resource, error)
		SoftDeleteResource(ctx echo.Context, resourceID int) error
	}
	// ServiceImpl represents service implementation of service.
	ServiceImpl struct {
		repo Repository
	}
)

// NewService returns service implementation of user service.
func NewService(repo Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

// CreateResource ...
func (svc *ServiceImpl) CreateResource(req CreateReq) (*model.Resource, error) {
	user, err := svc.repo.GetUserByID(req.UserID)
	if err != nil {
		return nil, err
	}

	if user.Quota == 0 {
		return nil, ErrNotEnoughQuota
	}

	resource := model.Resource{
		Name:   req.Name,
		UserID: req.UserID,
	}

	id, err := svc.repo.CreateResource(resource)
	if err != nil {
		return nil, err
	}
	resource.ID = id

	return &resource, nil
}

// GetResourcesOfUser ...
func (svc *ServiceImpl) GetResourcesOfUser(userID, role int) ([]model.Resource, error) {
	if role == model.RoleUser.Int() {
		resources, err := svc.repo.GetResourcesByUserID(userID)
		if err != nil {
			return nil, err
		}
		return resources, nil
	}

	// GetAll
	resources, err := svc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// SoftDeleteResource ...
func (svc *ServiceImpl) SoftDeleteResource(ctx echo.Context, resourceID int) error {
	role := model.RoleFromContext(ctx)
	userID := model.UserIDFromContext(ctx)
	iUserID, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	if role == model.RoleUser.Int() {
		if !svc.repo.ResourceBelongsUser(resourceID, iUserID) {
			return errors.New("you are not owner of this resource")
		}
	}

	return svc.repo.SoftDelete(resourceID)
}
