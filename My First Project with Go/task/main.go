package main

import "fmt"

func main() {
	// Write your code here
	var (
		bubbleGum     = 202
		toffee        = 118
		iceCream      = 2250
		milkChocolate = 1680
		doughnut      = 1075
		pancake       = 80
	)

	income := float32(bubbleGum + toffee + iceCream + milkChocolate + doughnut + pancake)

	fmt.Printf(`Earned amount:
Bubblegum: $%d
Toffee: $%d
Ice cream: $%d
Milk chocolate: $%d
Doughnut: $%d
Pancake: $%d

Income: $%.1f`, bubbleGum, toffee, iceCream, milkChocolate, doughnut, pancake, income)

	var (
		staffExpenses int
		otherExpenses int
	)

	fmt.Println("\nStaff expenses: ")
	fmt.Scanln(&staffExpenses)

	fmt.Println("Other expenses: ")
	fmt.Scanln(&otherExpenses)

	netIncome := income - float32(staffExpenses+otherExpenses)

	fmt.Printf("Net Income: $%.0f", netIncome)

}
