package user

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

// SignIn returns token for user to access platform.
func (ctrl Controller) SignIn(e echo.Context) error {
	return nil
}

// RegisterHTTPRouter registers HTTP endpoints.
func (ctrl Controller) RegisterHTTPRouter(e *echo.Echo) {
	auth := middleware.NewAuthorizer(ctrl.signer)

	userRout := e.Group("/user")

	userRout.POST("/login", ctrl.SignIn, auth.Authenticate())
}
