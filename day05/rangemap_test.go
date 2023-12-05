package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	mapping := NewMapping()
	mapping.AddRange(*NewRangeMap(5, 0, 6))
	mapping.AddRange(*NewRangeMap(8, 0, 5))
	// should result in a single range [5, 12], aka RangeMap {5, ?, 8}
	assert.Equal(t, 1, len(mapping.rangeMaps))
	assert.Equal(t, 5, mapping.rangeMaps[0].sourceStart)
	assert.Equal(t, 12, mapping.rangeMaps[0].SourceEnd())
}

func TestNormalize2(t *testing.T) {
	mapping := NewMapping()
	mapping.AddRange(*NewRangeMap(0, 0, 3))  // [0, 2]
	mapping.AddRange(*NewRangeMap(4, 0, 5))  // [4, 8]
	mapping.AddRange(*NewRangeMap(12, 0, 9)) // [12, 20]
	mapping.AddRange(*NewRangeMap(7, 0, 6))  // [7, 12]

	// should result in 2 ranges
	assert.Equal(t, 2, len(mapping.rangeMaps))
}
