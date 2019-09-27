package pkg

// REFERENCES: https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	// ErrInvalidPasswordLength ...
	ErrInvalidPasswordLength = errors.New("invalid password length")
	// ErrInvalidHash ...
	ErrInvalidHash = errors.New("invalid hash format")
	// ErrIncompatibleVersion ...
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

// Config ...
type Config struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// NewConfig return new config of passport.
func NewConfig() Config {
	return Config{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

// Hasher is interface of hashing package.
type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, encodedHash string) (bool, error)
}

// Manager is implementation of hasher.
type Manager struct {
	config Config
}

// NewHasher return new implementation of hasher.
func NewHasher(cfg Config) Hasher {
	return &Manager{
		config: cfg,
	}
}

// Hash is hashing action to hash password.
func (m *Manager) Hash(password string) (string, error) {
	if trimmed := strings.TrimSpace(password); len(trimmed) < 8 {
		return "", ErrInvalidPasswordLength
	}

	salt, err := generateRandomBytes(m.config.saltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		m.config.iterations,
		m.config.memory,
		m.config.parallelism,
		m.config.keyLength,
	)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		m.config.memory,
		m.config.iterations,
		m.config.parallelism,
		b64Salt,
		b64Hash,
	)
	return encodedHash, nil
}

// Compare is action to compare two hash.
func (m *Manager) Compare(password, encodedHash string) (bool, error) {
	cfg, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.IDKey([]byte(password), salt, cfg.iterations, cfg.memory, cfg.parallelism, cfg.keyLength)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (cfg *Config, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}
	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}
	cfg = &Config{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &cfg.memory, &cfg.iterations, &cfg.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	cfg.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	cfg.keyLength = uint32(len(hash))
	return cfg, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
