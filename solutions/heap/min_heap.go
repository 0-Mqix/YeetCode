package heap

//https://www.youtube.com/watch?v=t0Cq6tVNRBA&t=554s&ab_channel=HackerRank

type Heap struct {
	data []int
	size int
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*1 + 2
}

func Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) HeapifyDown(i int) {
	l, r := Left(i), Right(i)

	if h.data[l] > h.data[r] {

	} else if {

	}
}

func (h *Heap) HeapifyUp(i int) {
	p := Parent(i)

	for h.data[p] > h.data[i] {
		h.Swap(i, p)

		i = p
		p = Parent(i)
	}
}

func (h *Heap) Push(i int) {
	h.data = append(h.data, i)
	h.HeapifyUp(h.size)
}

func (h *Heap) Pop(i int) {
	v := h.data[0]
	h.Swap(0, h.size)
	h.HeapifyDown(h.size)
	h.size--

	return v
}
