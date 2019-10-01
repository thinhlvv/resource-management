package resource

import (
	"errors"

	"github.com/thinhlvv/resource-management/model"
)

// ErrNotEnoughQuota ...
var ErrNotEnoughQuota = errors.New("You don't have enough quota to create new resource")

type (
	// Service is interface of user service.
	Service interface {
		CreateResource(req CreateReq) (*model.Resource, error)
		GetResourcesOfUser(userID int) ([]model.Resource, error)
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

	// check quota
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
func (svc *ServiceImpl) GetResourcesOfUser(userID int) ([]model.Resource, error) {
	resources, err := svc.repo.GetResourcesByUserID(userID)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
