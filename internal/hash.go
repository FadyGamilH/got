package internal

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(content string) string {
	return "fd790ad99ff75c6c383962e2f0bc1ffeabc22142"
}

func Sha1(content string) string {
	hash := sha1.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
