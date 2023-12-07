package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardsToType(t *testing.T) {
	typ := CardsToType([]int{3, 2, 10, 3, 13})
	assert.Equal(t, 2, typ)

	typ = CardsToType([]int{13, 10, 11, 11, 10})
	assert.Equal(t, 3, typ)

	typ = CardsToType([]int{14, 14, 14, 14, 7})
	assert.Equal(t, 6, typ)
}
