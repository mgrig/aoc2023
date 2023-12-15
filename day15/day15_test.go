package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	h := Hash("HASH")
	assert.Equal(t, uint32(52), h)
}
