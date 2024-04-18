package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Hash(content string) string {
	return Sha1(fmt.Sprintf("blob %d\x00%s", len(content), content))
}

func Sha1(content string) string {
	hash := sha1.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
