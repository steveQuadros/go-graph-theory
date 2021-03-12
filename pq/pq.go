package pq

// import ("fmt")

// this PQ is specifically for djikstras algorithm, and will store an edge at [0], and a weight at [1]
type PQ interface {
	Add([2]int)
	Poll()[2]int
}

type MinPQ struct{
	items [][2]int
}

func NewMinPQ() *MinPQ {
	return &MinPQ{
		items: [][2]int{},
	}
}

var _ PQ = (*MinPQ)(nil)

func (q *MinPQ) Add(pair [2]int) {
	q.items = append(q.items, pair)
	idx := q.Size()-1
	pIdx := parentIdx(idx)
	for pIdx >= 0 && q.items[idx][1] < q.items[pIdx][1] {
		q.swap(idx, pIdx)
		idx = pIdx
		pIdx = parentIdx(idx)
	}
}

func (q *MinPQ) Poll() [2]int {
	res := q.items[0]
	idx, j := 0, q.Size()-1
	q.swap(idx,j)
	q.items = q.items[:q.Size()-1]

	l, r := childIdx(idx)
	var ci int
	for l < q.Size() {
		if r >= q.Size() || q.items[l][1] < q.items[r][1] {
			ci = l
		} else {
			ci = r
		}
		if q.items[idx][1] > q.items[ci][1] {
			q.swap(idx, ci)
		}
		idx = ci
		l, r = childIdx(idx)
	}
	return res
}

func (q *MinPQ) Size() int {
	return len(q.items)
}

func (q *MinPQ) Items() [][2]int {
	return q.items
}

func (q *MinPQ) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func parentIdx(i int) int {
	return (i-1)/2
}

func childIdx(i int) (l, r int) {
	l, r = i*2+1, i*2+2
	return
}