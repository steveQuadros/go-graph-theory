package pq

// import ("fmt")

// this PQ is specifically for djikstras algorithm, and will store an edge at [0], and a weight at [1]
type PQ interface {
	Add(int,int)
	Poll() (int,int)
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

func (q *MinPQ) Add(to, weight int) {
	q.items = append(q.items, [2]int{to, weight})
	idx := q.Size()-1
	pIdx := parentIdx(idx)
	for pIdx >= 0 && q.items[idx][1] < q.items[pIdx][1] {
		q.swap(idx, pIdx)
		idx = pIdx
		pIdx = parentIdx(idx)
	}
}

func (q *MinPQ) Poll() (to, weight int) {
	to, weight = q.items[0][0], q.items[0][1]

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
	return to, weight
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