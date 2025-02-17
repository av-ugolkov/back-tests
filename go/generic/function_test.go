package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	i := Sum(1, 2)
	assert.Equal(t, 3, i)

	f := Sum(1.2, 2.7)
	assert.InDelta(t, 3.9, f, 0.00001)
}
