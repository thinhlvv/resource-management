package pkg

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrInvalidToken ...
	ErrInvalidToken = errors.New("invalid token")
)

// Claims represent the JWT Claims.
type Claims struct {
	jwt.StandardClaims
	Role int `json:"role"`
}

// NewClaims returns a new Claim with an expiry of 1 week.
func NewClaims(id string) Claims {
	return Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
			Subject:   id,
			IssuedAt:  time.Now().Unix(),
		},
	}
}

type (
	// Signer represents the JWT token operation.
	Signer interface {
		Sign(id string) (string, error)
		SignWithRole(id string, role int) (string, error)
		Verify(token string) (*Claims, error)
	}

	// Signer implements the Signer interface.
	signer struct {
		ttl    time.Duration
		secret []byte
	}
)

// NewSigner returns a new Signer with the given secret.
func NewSigner(secret string, ttl time.Duration) Signer {
	return &signer{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

// Sign takes a user id and returns a signed token or error.
func (s *signer) Sign(id string) (string, error) {
	claims := NewClaims(id)
	claims.StandardClaims.ExpiresAt = time.Now().Add(s.ttl).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *signer) SignWithRole(id string, role int) (string, error) {
	claims := NewClaims(id)
	claims.StandardClaims.ExpiresAt = time.Now().Add(s.ttl).Unix()
	claims.Role = role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// Verify takes in a raw token and attempts to parse it.
func (s *signer) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	// Passing in authorization header with "Bearer undefined" cause this
	// issue.
	if token == nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
