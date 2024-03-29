package resource

import "github.com/thinhlvv/resource-management/model"

// New return user controller services.
func New(app model.App) *Controller {
	repo := NewRepo(app.DB)
	svc := NewService(repo)
	c := NewController(svc, app)
	return c
}
