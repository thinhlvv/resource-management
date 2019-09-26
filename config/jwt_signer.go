package config

import (
	"sync"

	"github.com/thinhlvv/resource-management/pkg"
)

// MustInitJWTSigner returns DB pointer.
func MustInitJWTSigner(cfg *Config) pkg.Signer {
	var doOnce sync.Once
	var signer pkg.Signer

	doOnce.Do(func() {
		signer = pkg.NewSigner(cfg.JWT.Secret, cfg.JWT.Duration)
	})

	return signer
}
