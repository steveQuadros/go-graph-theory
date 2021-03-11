package traversal

import (
	"github.com/stevequadros/go-graph-theory/graph"
)

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

func ShortestPathGrid() {
	
}