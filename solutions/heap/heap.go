package heap

//made this to practice heaps

//https://www.youtube.com/watch?v=t0Cq6tVNRBA&t=554s&ab_channel=HackerRank

type Heap struct {
	data []int
	size int
	// comp func(int) bool
}

func New(data []int) *Heap { return &Heap{data: data} }

func Left(i int) int   { return 2*i + 1 }
func Right(i int) int  { return 2*i + 2 }
func Parent(i int) int { return (i - 1) / 2 }

func (h *Heap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Size() int     { return h.size }
func (h *Heap) Data() []int   { return h.data }

func (h *Heap) HeapifyDown(i int) {
	l, r := Left(i), Right(i)
	s := l

	if l > h.size-1 {
		return
	}

	if r < h.size && h.data[r] < h.data[l] {
		s = r
	}

	if h.data[s] < h.data[i] {
		h.Swap(i, s)
		h.HeapifyDown(s)
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
	h.size++
}

func (h *Heap) Pop() int {
	v := h.data[0]
	h.size--
	h.Swap(0, h.size)
	h.data = h.data[:h.size]

	h.HeapifyDown(0)
	return v
}
