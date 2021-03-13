package path

import (
	"testing"
	"github.com/stevequadros/go-graph-theory/graph"
	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	g := graph.New(11, &graph.Options{Directed: true})
	g.Insert(0,1)
	g.Insert(0,2)
	g.Insert(0,6)
	g.Insert(1,5)
	g.Insert(1,4)
	g.Insert(2,3)
	g.Insert(4,7)
	g.Insert(4,9)
	g.Insert(5,8)
	g.Insert(6,7)
	g.Insert(7,10)
	g.Insert(9,10)
	assert.Equal(t, []int{0,6,7,10}, ShortestPathBFS(g, 0, 10))
	assert.Equal(t, []int{}, ShortestPathBFS(g, 1, 6))
}

func TestSSSP(t *testing.T) {
	g := graph.New(8, &graph.Options{Directed: true, Weighted: true})
	g.InsertWeighted(0,2,3)
	g.InsertWeighted(0,1,6)
	g.InsertWeighted(1,4,8)
	g.InsertWeighted(1,5,11)
	g.InsertWeighted(2,1,4)
	g.InsertWeighted(2,4,4)
	g.InsertWeighted(2,3,11)
	g.InsertWeighted(3,6,9)
	g.InsertWeighted(4,3,-4)
	g.InsertWeighted(4,5,2)
	g.InsertWeighted(4,6,1)
	g.InsertWeighted(4,7,5)
	g.InsertWeighted(5,6,2)
	g.InsertWeighted(7,6,1)

	dist := DAGSP(g, 0) // 0,6,3,3,7,9,11,12
	assert.Equal(t, 7, dist[4])
	assert.Equal(t, 6, dist[1])
	assert.Equal(t, 12, dist[7])
}

func TestDjikstra(t *testing.T) {
	g := graph.New(5, &graph.Options{Directed: true, Weighted: true})
	g.InsertWeighted(0,1,4)
	g.InsertWeighted(0,2,1)
	g.InsertWeighted(1,3,1)
	g.InsertWeighted(2,1,2)
	g.InsertWeighted(2,3,5)
	g.InsertWeighted(3,4,3)

	t.Run("Acyclic", func(t *testing.T) {
		dist, path := Djikstra(g, 0, 4)
		// dist
		assert.Equal(t, dist[0], 0, "dist wrong")
		assert.Equal(t, dist[1], 3, "dist wrong")
		assert.Equal(t, dist[2], 1, "dist wrong")	
		assert.Equal(t, dist[3], 4, "dist wrong")	
		assert.Equal(t, dist[4], 7, "dist wrong")	

		// path
		assert.Equal(t, path[0], 4, "path wrong")
		assert.Equal(t, path[1], 3, "path wrong")
		assert.Equal(t, path[2], 1, "path wrong")
		assert.Equal(t, path[3], 2, "path wrong")
		assert.Equal(t, path[4], 0, "path wrong")
	})

	t.Run("not last node", func(t *testing.T) {
		dist, path := Djikstra(g, 0, 3)
		assert.Equal(t, dist[3], 4, "dist wrong")
		assert.Equal(t, path[0], 3, "path wrong")
		assert.Equal(t, path[1], 1, "path wrong")
		assert.Equal(t, path[2], 2, "path wrong")
		assert.Equal(t, path[3], 0, "path wrong")
	})
	
	t.Run("Has Cycle", func(t *testing.T) {
		// add a cycle between 2 <-> 1
		g.InsertWeighted(1,2,3)
		g.InsertWeighted(2,0,3)
		g.InsertWeighted(1,0,3)
		
		dist, _ := Djikstra(g, 0, 4)
		assert.Equal(t, dist[0], 0)
		assert.Equal(t, dist[1], 3)
		assert.Equal(t, dist[2], 1)
		assert.Equal(t, dist[3], 4)
		assert.Equal(t, dist[4], 7)
	})

	t.Run("unconnected", func(t *testing.T) {
		t.Skip("requires a little changing of graph, but would be useful in future")
		g := graph.New(3, &graph.Options{Directed: true, Weighted: true})
		dist, path := Djikstra(g,0,1)
		assert.Equal(t, 127, dist[0])
		assert.Equal(t, 127, dist[1])
		assert.Equal(t, -1, path[0])
		assert.Equal(t, -1, path[1])
	})
}