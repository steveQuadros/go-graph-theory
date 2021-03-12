package pq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinPQ(t *testing.T) {
	q := NewMinPQ()
	for i := 8; i > 0; i-- {
		q.Add([2]int{0,i})
	}

	assert.Equal(t, 1, q.Poll()[1])
	assert.Equal(t, 2, q.Poll()[1])
	assert.Equal(t, 3, q.Poll()[1])
	assert.Equal(t, 4, q.Poll()[1])
	assert.Equal(t, 5, q.Poll()[1])
	assert.Equal(t, 6, q.Poll()[1])
	assert.Equal(t, 7, q.Poll()[1])	
}