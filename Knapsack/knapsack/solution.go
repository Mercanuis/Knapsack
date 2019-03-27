package knapsack

import "fmt"

//Solution represents a possible solution
type Solution struct {
	items []Item
	value int
}

//NewSolution returns a Solution
func NewSolution(items []Item, value int) *Solution {
	return &Solution{
		items: items,
		value: value,
	}
}

//Display prints out the solution
func (s *Solution) Display() {
	if s.items != nil && len(s.items) > 0 {
		fmt.Printf("Solution contains %d items\n", len(s.items))
		fmt.Printf("Items = %v", s.items)
		fmt.Printf("Total = %d", s.value)
	}
}
