package traversal

type Node struct {
	Val int
	Next *Node
}

// DFS uses DFS to return the number of connected components
type ComponentResult struct {
	Components []int
	Count int
}
func DFS(nodes []*Node) ComponentResult {
	components := make([]int, len(nodes))
	visited := make([]bool, len(nodes))
	var count int

	// initialize components
	for i := 0; i < len(components); i++ {
		components[i] = i
	}

	for nodeVal, _ := range nodes {
		if !visited[nodeVal] {
			count++
			components, visited = dft(count, nodeVal, nodes, components, visited)
		}
	}

	return ComponentResult{
		Components: components,
		Count: count,
	}
}

func dft(component, at int, nodes []*Node, components []int, visited []bool) ([]int, []bool) {
	if visited[at] {
		return components, visited
	}

	visited[at] = true
	components[at] = component

	for node := nodes[at]; node != nil; node = node.Next {
		components, visited = dft(component, node.Val, nodes, components, visited)
	}

	return components, visited
}