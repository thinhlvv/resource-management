package resource

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/middleware"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
)

// Controller returns endpoint handlers.
type Controller struct {
	signer pkg.Signer
}

// NewController returns new Controller.
func NewController(app model.App) *Controller {
	return &Controller{
		signer: app.JWTSigner,
	}
}

// GetList ...
func (ctrl Controller) GetList(e echo.Context) error {
	return nil
}

// Create ...
func (ctrl Controller) Create(e echo.Context) error {
	return nil
}

// Delete ...
func (ctrl Controller) Delete(e echo.Context) error {
	return nil
}

// RegisterHTTPRouter registers HTTP endpoints.
func (ctrl Controller) RegisterHTTPRouter(e *echo.Echo) {
	auth := middleware.NewAuthorizer(ctrl.signer)

	userRout := e.Group("/resource")

	userRout.GET("", ctrl.GetList, auth.Authenticate())
	userRout.POST("", ctrl.Create, auth.Authenticate())
	userRout.DELETE("", ctrl.Delete, auth.Authenticate())
}
