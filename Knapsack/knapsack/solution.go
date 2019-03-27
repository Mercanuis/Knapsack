package knapsack

import "fmt"

//Solution blah
type Solution struct {
	items []Item
	value int
}

//NewSolution blah
func NewSolution(items []Item, value int) *Solution {
	return &Solution{
		items: items,
		value: value,
	}
}

//Display blah
func (s *Solution) Display() {
	if s.items != nil && len(s.items) > 0 {
		fmt.Printf("Solution contains %d items\n", len(s.items))
		fmt.Printf("Items = %v", s.items)
		fmt.Printf("Total = %d", s.value)
	}
}
