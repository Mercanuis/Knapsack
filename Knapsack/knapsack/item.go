package knapsack

//Item represents a knapsack item
type Item struct {
	name   string
	value  int
	weight int
}

//NewItem returns a new Item
func NewItem(n string, v, w int) Item {
	return Item{
		name:   n,
		value:  v,
		weight: w,
	}
}
