package knapsack

import (
	"math"
)

//Knapsack blah
type Knapsack struct {
	items    []Item
	capacity int
}

//NewKnapsack blah
func NewKnapsack(items []Item, capacity int) Knapsack {
	return Knapsack{
		items:    items,
		capacity: capacity,
	}
}

//SolveKnapsack blah
func (k *Knapsack) SolveKnapsack() *Solution {
	numItems := len(k.items)
	capacity := k.capacity

	matrix := make([][]int, numItems+1)
	for i := range matrix {
		matrix[i] = make([]int, capacity+1)
	}

	//initialize index 0 for all rows to 0
	for i := 0; i < capacity; i++ {
		matrix[0][i] = 0
	}

	//fill in the matrix and figure out the weights
	for i := 1; i < numItems; i++ {
		for j := 0; j < capacity; j++ {
			if k.items[i-1].weight > j {
				matrix[i][j] = matrix[i-1][j]
			} else {
				matrix[i][j] = int(
					math.Max(float64(matrix[i-1][j]),
						float64(matrix[i-1][j-k.items[i-1].weight]+k.items[i-1].value)))
			}
		}
	}

	//Determine values via the matrix
	res := matrix[numItems-1][capacity-1]
	w := capacity
	var sol = make([]Item, 0)

	for i := numItems; i > 0 && res > 0; i-- {
		if res != matrix[i-1][w] {
			//remove an item and the weight of it
			sol = append(sol, k.items[i-1])
			res -= k.items[i-1].value
			w -= k.items[i-1].weight
		}
	}

	return NewSolution(sol, matrix[numItems-1][capacity-1])
}
