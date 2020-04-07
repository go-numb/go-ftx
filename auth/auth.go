package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Config struct {
	Key    string
	Secret string
}

func New(key, secret string) *Config {
	return &Config{
		Key:    key,
		Secret: secret,
	}
}

func (p *Config) Signture(body string) string {
	mac := hmac.New(sha256.New, []byte(p.Secret))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}
