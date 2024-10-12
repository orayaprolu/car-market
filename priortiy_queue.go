package orderbook

type Limit_Min_PQ []Limit

func (h Limit_Min_PQ) Len() int           { return len(h) }
func (h Limit_Min_PQ) Less(i, j int) bool { return h[i].LimitPrice < h[j].LimitPrice }
func (h Limit_Min_PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Limit_Min_PQ) Push(x any) {
	*h = append(*h, x.(Limit))
}

func (h *Limit_Min_PQ) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Limit_Max_PQ []Limit

func (h Limit_Max_PQ) Len() int           { return len(h) }
func (h Limit_Max_PQ) Less(i, j int) bool { return h[i].LimitPrice > h[j].LimitPrice }
func (h Limit_Max_PQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Limit_Max_PQ) Push(x any) {
	*h = append(*h, x.(Limit))
}

func (h *Limit_Max_PQ) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
