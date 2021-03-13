package pq

// import ("fmt")

// this PQ is specifically for djikstras algorithm, and will store an edge at [0], and a weight at [1]
type PQ interface {
	Add(int,int)
	RemoveMin() (int,int)
}

func NewMinPQ() *MinPQ {
	return &MinPQ{
		items: [][2]int{},
	}
}

type MinPQ struct {
	items [][2]int	
}

var _ PQ = (*MinPQ)(nil)

func (p *MinPQ) Add(to, weight int) {
	p.items = append(p.items, [2]int{to, weight})
	idx := len(p.items)-1
	pidx := (idx-1)/2
	for idx >= 0 && p.items[idx][1] < p.items[pidx][1] {
		p.items[idx], p.items[pidx] = p.items[pidx], p.items[idx]
		idx = pidx
		pidx = (idx-1)/2
	}
}

func (p *MinPQ) RemoveMin() (to, weight int) {
	idx := 0
	to, weight = p.items[idx][0], p.items[idx][1]
	p.items[idx] = p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]

	maxIdx := len(p.items)-1
	for idx*2+1 <  maxIdx {
		ci := idx*2+1
		if  ci < maxIdx-1 && p.items[ci+1][1] < p.items[ci][1] {
			ci+=1
		}
		p.items[idx], p.items[ci] = p.items[ci], p.items[idx]
		idx = ci
	}
	return to, weight
}

func (p *MinPQ) Empty() bool {
	return len(p.items) == 0
}

func (p *MinPQ) Size() int {
	return len(p.items)
}

// @TODO - add indexed priority queue w/
// - Remove
// - Update