package base_struct_one

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

type Item struct {
	// Service host
	Data string
}

type Bag struct {
	// Service host
	Name string

	Items []Item
}

func (b Bag) add(item Item) {
	items := b.Items
	b.Items = append(items, item)
}

func (b Bag) isEmpty() bool {
	return false
}

func (b Bag) size() int {
	return len(b.Items)
}
