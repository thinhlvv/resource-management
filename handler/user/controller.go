package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/middleware"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
)

// Controller returns endpoint handlers.
type Controller struct {
	service   Service
	signer    pkg.Signer
	validator pkg.RequestValidator
}

// NewController returns new Controller.
func NewController(svc Service, app model.App) *Controller {
	return &Controller{
		signer:    app.JWTSigner,
		service:   svc,
		validator: app.RequestValidator,
	}
}

type (
	// LoginRequest represents the request of a admin logging in.
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email" conform:"trim"`
		Password string `json:"password" validate:"required,min=8" conform:"trim"`
	}
	// LoginResp represents the response of a admin logging in.
	LoginResp struct {
		AccessToken string `json:"access_token"`
	}
)

// Login returns token for user to access platform.
func (ctrl Controller) Login(c echo.Context) error {
	req := LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	if err := ctrl.validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	accessToken, err := ctrl.service.Login(c, req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.NewErrorResponse(c, err))
	}
	return c.JSON(http.StatusOK, &LoginResp{
		AccessToken: accessToken,
	})
}

type (
	// SignupReq ...
	SignupReq struct {
		Email    string `json:"email" validate:"required,email" conform:"trim"`
		Password string `json:"password" validate:"required,min=8" conform:"trim"`
		Role     int    `json:"role" validate:"required,min=1,max=2"`
	}
	// SignupResp ...
	SignupResp struct {
		AccessToken string `json:"access_token"`
	}
)

// Signup returns token for user to access platform.
func (ctrl Controller) Signup(c echo.Context) error {
	req := SignupReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	if err := ctrl.validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	accessToken, err := ctrl.service.Signup(c, req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.NewErrorResponse(c, err))
	}

	return c.JSON(http.StatusOK, &SignupResp{
		AccessToken: accessToken,
	})
}

// GetList returns token for user to access platform.
func (ctrl Controller) GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Delete returns token for user to access platform.
func (ctrl Controller) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Create returns token for user to access platform.
func (ctrl Controller) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Update returns token for user to access platform.
func (ctrl Controller) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

// RegisterHTTPRouter registers HTTP endpoints.
func (ctrl Controller) RegisterHTTPRouter(e *echo.Echo) {
	auth := middleware.NewAuthorizer(ctrl.signer)

	userRout := e.Group("/user")

	userRout.POST("/login", ctrl.Login, auth.Authenticate())
	userRout.POST("/signup", ctrl.Signup)

	// CRUD
	userRout.GET("", ctrl.GetList, auth.ValidateWithRoles([]int{model.RoleAdmin.Int()}))
	userRout.DELETE("", ctrl.Delete, auth.ValidateWithRoles([]int{model.RoleAdmin.Int()}))
	userRout.POST("", ctrl.Create, auth.ValidateWithRoles([]int{model.RoleAdmin.Int()}))
	userRout.PUT("/:id", ctrl.Update, auth.ValidateWithRoles([]int{model.RoleAdmin.Int()}))
}
