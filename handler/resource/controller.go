package resource

import (
	"net/http"
	"strconv"

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
		service:   svc,
		signer:    app.JWTSigner,
		validator: app.RequestValidator,
	}
}

// GetList ...
func (ctrl Controller) GetList(c echo.Context) error {
	userID := model.UserIDFromContext(c)
	iUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}
	role := model.RoleFromContext(c)

	resources, err := ctrl.service.GetResourcesOfUser(iUserID, role)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.NewErrorResponse(c, err))
	}

	return c.JSON(http.StatusOK, resources)
}

type (
	// CreateReq ...
	CreateReq struct {
		Name   string `json:"name" validate:"required,max=100"`
		UserID int    `json:"user_id" validate:"required"`
	}
	// CreateRes returns info of resource.
	CreateRes struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

// Create ...
func (ctrl Controller) Create(c echo.Context) error {
	req := CreateReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	userID := model.UserIDFromContext(c)
	iUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}
	req.UserID = iUserID

	if err := ctrl.validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	}

	resource, err := ctrl.service.CreateResource(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.NewErrorResponse(c, err))
	}
	return c.JSON(http.StatusOK, CreateRes{
		ID:   resource.ID,
		Name: resource.Name,
	})
}

// Delete ...
func (ctrl Controller) Delete(c echo.Context) error {
	// resourceID := c.Param("id")
	// iResourceID , err := strconv.Atoi(resourceID)
	// if err != nil {
	// return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
	//}
	// role := model.RoleFromContext(c)

	// if err := ctrl.service.SoftDeleteResource(c, iResourceID); err != nil {
	// return c.JSON(http.StatusUnprocessableEntity, model.NewErrorResponse(c, err))
	// }
	return c.JSON(http.StatusAccepted, nil)
}

// RegisterHTTPRouter registers HTTP endpoints.
func (ctrl Controller) RegisterHTTPRouter(e *echo.Echo) {
	auth := middleware.NewAuthorizer(ctrl.signer)

	userRout := e.Group("/resource")

	userRout.GET("", ctrl.GetList, auth.Authenticate())
	userRout.POST("", ctrl.Create, auth.Authenticate())
	userRout.DELETE("/:id", ctrl.Delete, auth.Authenticate())
}
