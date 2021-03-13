package pq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinPQ(t *testing.T) {
	q := NewMinPQ()
	for i := 8; i > 0; i-- {
		q.Add(i,i)
	}

	for i := 1; i < 8; i++ {
		_, weight := q.RemoveMin()
		assert.Equal(t, i, weight)
	}
}