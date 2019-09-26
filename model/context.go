package model

import (
	"github.com/labstack/echo"
)

type contextKey string

var (
	requestIDKey contextKey = "req_id"
	userIDKey    contextKey = "user_id"
	localeKey    contextKey = "locale"
	roleKey      contextKey = "role"
)

// RequestIDFromContext returns the ID of a request of a context.
func RequestIDFromContext(ctx echo.Context) string {
	return ctx.Get(string(requestIDKey)).(string)
}

// ContextWithRequestID returns a context with the given request ID.
func ContextWithRequestID(ctx echo.Context, requestID string) echo.Context {
	ctx.Set(string(requestIDKey), requestID)
	return ctx
}

// UserIDFromContext returns the user ID of a context.
func UserIDFromContext(ctx echo.Context) string {
	return ctx.Get(string(userIDKey)).(string)
}

// ContextWithUserID returns a context with the given user ID.
func ContextWithUserID(ctx echo.Context, userID string) echo.Context {
	ctx.Set(string(userIDKey), userID)
	return ctx
}

// ContextWithRole sets role to context.
func ContextWithRole(ctx echo.Context, role int) echo.Context {
	ctx.Set(string(roleKey), role)
	return ctx
}

// RoleFromContext returns role from context.
func RoleFromContext(ctx echo.Context) int {
	return ctx.Get(string(roleKey)).(int)
}
