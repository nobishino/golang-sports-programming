package fenwick

type BIT struct {
	Nodes []int
	Size  int
}

func NewBIT(size int) *BIT {
	bit := new(BIT)
	bit.Size = size
	bit.Nodes = make([]int, size+1)
	return bit
}

func (bit *BIT) Add(location, value int) {
	for i := location + 1; i <= bit.Size; i += (i & -i) {
		bit.Nodes[i] += value
	}
}

func (bit *BIT) Sum(start, end int) int {
	return bit.Accumulate(end) - bit.Accumulate(start-1)
}

func (bit *BIT) Accumulate(end int) (result int) {
	if end >= bit.Size {
		end = bit.Size - 1
	}
	for i := end + 1; i > 0; i -= i & -i {
		result += bit.Nodes[i]
	}
	return
}
