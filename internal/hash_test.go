package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	asserts := assert.New(t)
	content := "git with go"
	hash := Hash(content)
	asserts.Equal("fd790ad99ff75c6c383962e2f0bc1ffeabc22142", hash)
}
