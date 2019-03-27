package knapsack

import (
	"math"
)

//Knapsack represents a knapsack of items
type Knapsack struct {
	items    []Item
	capacity int
}

//NewKnapsack returns a new Knapsack
func NewKnapsack(items []Item, capacity int) Knapsack {
	return Knapsack{
		items:    items,
		capacity: capacity,
	}
}

//SolveKnapsack determines the knapsack's best solution per its capacity
//and then returns a Solution to print out
func (k *Knapsack) SolveKnapsack() *Solution {
	numItems := len(k.items)
	capacity := k.capacity

	//use make to initialize the 2D array because Go doesn't like non-constant initializations
	matrix := make([][]int, numItems+1)
	for i := range matrix {
		matrix[i] = make([]int, capacity+1)
	}

	//fill in the matrix and figure out the weights
	for i := 1; i <= numItems; i++ {
		//iterate each capacity
		for j := 0; j <= capacity; j++ {
			if k.items[i-1].weight > j {
				matrix[i][j] = matrix[i-1][j]
			} else {
				//for debugging purposes
				jvalue := j - k.items[i-1].weight
				kvalue := k.items[i-1].value
				matrix[i][j] = int(math.Max(float64(matrix[i-1][j]), float64(matrix[i-1][jvalue]+kvalue)))
			}
		}
	}

	//Determine values via the matrix
	res := matrix[numItems][capacity]
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

	return NewSolution(sol, matrix[numItems][capacity])
}
