package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hash(t *testing.T) {
	asserts := assert.New(t)
	content := "git with go"
	hash := Hash(content)
	asserts.Equal("fd790ad99ff75c6c383962e2f0bc1ffeabc22142", hash)
}

func Test_Sha1(t *testing.T) {
	asserts := assert.New(t)
	content := "git with go"
	sha1Hash := Sha1(content)
	asserts.Equal("1cd141fb363aca49017c83d35f3609d1819eb3b0", sha1Hash)
}
