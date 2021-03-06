package path

import (
	"math"
	"github.com/stevequadros/go-graph-theory/graph"
	"github.com/stevequadros/go-graph-theory/topsort"
	"github.com/stevequadros/go-graph-theory/pq"
)

var Infinity = int(math.Inf(1))

func ShortestPathBFS(g *graph.Graph, start, end int) []int {
	path := make([]int, len(g.Nodes))
	for i := 0; i < len(path); i++ {
		path[i] = -1
	}
	q := []int{start}

	for len(q) != 0 {
		at := q[0]
		q = q[1:]
		for node := g.Nodes[at]; node != nil; node = node.Next {
			if path[node.Val] == -1 {
				path[node.Val] = at
				q = append(q, node.Val)
				if node.Val == end {
					break
				}
			}
		}
	}

	prev := []int{}
	if path[end] == -1 {
		return prev
	}

	cur := end
	for {
		prev = append(prev, cur)
		cur = path[cur]
		if cur == -1 {
			break
		}
	}

	return reverse(prev)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// The user experience of returning a slice of pointer to ints is no good
// I would normally either
// - wrap the output in something that would provide a `get` method for a vertex, eg: res.get(4) would return the path from source->4
// - or potentially limit the weight of my nodes to be less than math.MaxInt
// - use []*int and do a bunch of dereferencing :(
// Time: O(V + E)
// Space 
func DAGSP(g *graph.Graph, start int) []int {
	topOrder := topsort.Kahn(g)
	distance := make([]int, len(g.Nodes))
	for i, _ := range distance {
		distance[i] = math.MaxInt8
	}
	distance[start] = 0

	for i := 0; i < len(g.Nodes); i++ {
		at := topOrder[i]
		if distance[at] == math.MaxInt8 {
			continue
		}
		for node := g.Nodes[at]; node != nil; node = node.Next {
			newDistance := distance[at] + node.Weight()
			if newDistance < distance[node.Val] {
				distance[node.Val] = newDistance
			}
		}
	}
	return distance
}

// func LonestPath(g *graph.Graph, source, dest int) int {
// 	return int(math.Inf(-1))
// }

// Djikstra's algorithm can return the shortest path from 
// source to dest given non negative weights
// it runs in 
func Djikstra(g *graph.Graph, start, end int) ([]int, []int) {
	distance := make([]int, len(g.Nodes))
	for i := 0; i < len(distance); i++ {
		if i == start {
			distance[i] = 0
		} else {
			distance[i] = math.MaxInt8
		}
	}

	paths := make([]int, len(g.Nodes))
	for i := 0; i < len(paths); i++ {
		paths[i] = -1
	}

	pq := pq.NewMinPQ()
	pq.Add(start, 0)
	
	for !pq.Empty() {
		at, weight := pq.RemoveMin()
		if at == end {
			break
		}

		for node := g.Nodes[at]; node != nil; node = node.Next {
			newDist := weight + node.Weight()
			if newDist < distance[node.Val] {
				distance[node.Val] = newDist
				pq.Add(node.Val, newDist)
				paths[node.Val] = at
			}
		}
	}

	sp := []int{}
	for {
		sp = append(sp, end)
		if end == start || paths[end] == -1 {
			break
		}
		end = paths[end]
	}
	return distance, sp
}

func BellmanFord(g *graph.Graph, start int) ([]float64, []int) {
	dist, path := make([]float64, len(g.Nodes)), []int{}
	for i := 0; i < len(dist); i++ {
		dist[i] = math.Inf(1)
	}
	dist[start] = 0 

	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes); j++ {
			for node := g.Nodes[j]; node != nil; node = node.Next {
				if dist[j] + float64(node.Weight()) < dist[node.Val] {
					dist[node.Val] = dist[j] + float64(node.Weight())
				}
			}	
		}
	}

	for i := 0; i < len(g.Nodes); i++ {
		for j := 0; j < len(g.Nodes); j++ {
			for node := g.Nodes[j]; node != nil; node = node.Next {
				if dist[j] + float64(node.Weight()) < dist[node.Val] {
					dist[node.Val] = math.Inf(-1)
				}
			}	
		}
	}
	return dist, path
}
