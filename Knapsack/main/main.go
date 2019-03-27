package main

import "Knapsack/knapsack"

func main() {
	i1 := knapsack.NewItem("item1", 1, 3)
	i2 := knapsack.NewItem("item2", 2, 5)
	i3 := knapsack.NewItem("item3", 3, 8)
	i4 := knapsack.NewItem("item4", 4, 1)
	i5 := knapsack.NewItem("item5", 5, 6)

	items := []knapsack.Item{i1, i2, i3, i4, i5}
	sack := knapsack.NewKnapsack(items, 5)
	solution := sack.SolveKnapsack()
	solution.Display()
}
