package crypto_lib

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"golang.org/x/crypto/argon2"
	"strings"
)

func HashString(str string, cfg *config.Config) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := MakeHash(str, salt, cfg)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return cfg.Secret + "$" + b64Salt + "$" + b64Hash, nil
}

func VerifyString(str, encodedHash string, cfg *config.Config) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 3 || parts[0] != cfg.Secret {
		return false, fmt.Errorf("invalid encoded hash")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		return false, err
	}

	newHash := MakeHash(str, salt, cfg)
	return subtle.ConstantTimeCompare(newHash, storedHash) == 1, nil
}

func MakeHash(str string, salt []byte, cfg *config.Config) []byte {
	time := uint32(cfg.Argon2Time)
	memory := uint32(cfg.Argon2Memory)
	threads := cfg.Argon2Threads
	keyLen := uint32(cfg.Argon2KeyLen)

	hash := argon2.IDKey(
		[]byte(str),
		salt,
		time,
		memory,
		threads,
		keyLen,
	)
	return hash
}
