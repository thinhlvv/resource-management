package resource

import (
	"github.com/thinhlvv/resource-management/model"
)

type (
	// Service is interface of user service.
	Service interface {
		CreateResource(req CreateReq) (*model.Resource, error)
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
