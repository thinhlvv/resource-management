package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/thinhlvv/resource-management/model"
	"github.com/thinhlvv/resource-management/pkg"
)

var (
	// ErrInvalidAuthorizationHeader ...
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
)

// Authorizer represents all middleware interface related to Authentication.
type Authorizer interface {
	Authenticate() echo.MiddlewareFunc
	ValidateWithRoles(role []int) echo.MiddlewareFunc
}

type authorizer struct {
	signer pkg.Signer
}

// NewAuthorizer returns authorizer implementation.
func NewAuthorizer(signer pkg.Signer) Authorizer {
	return &authorizer{
		signer: signer,
	}
}

func (a *authorizer) Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			values := strings.Split(auth, " ")
			if len(values) != 2 {
				return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, ErrInvalidAuthorizationHeader))
			}
			bearer := values[0]
			if !strings.EqualFold(bearer, "Bearer") {
				return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, ErrInvalidAuthorizationHeader))
			}
			token := values[1]
			claims, err := a.signer.Verify(token)
			if err != nil {
				return c.JSON(http.StatusBadRequest, model.NewErrorResponse(c, err))
			}

			// Populate the context and inject them back into the header.
			c = model.ContextWithUserID(c, claims.StandardClaims.Subject)
			c = model.ContextWithRole(c, claims.Role)
			return next(c)
		}
	}
}

func (a *authorizer) ValidateWithRoles(validRoles []int) echo.MiddlewareFunc {
	return func(echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return nil
		}
	}
}
