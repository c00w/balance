package main

import (
	"fmt"
)

type class struct {
	human_string  string
	desired_ratio float64
}

func main() {

	dom_to_int := 0.7
	stock_to_bond := 0.9
	classes := []class{
		{"US Stock", dom_to_int * stock_to_bond},
		{"International Stock", (1 - dom_to_int) * stock_to_bond},
		{"US Bond", dom_to_int * (1 - stock_to_bond)},
		{"International Bond", (1 - dom_to_int) * (1 - stock_to_bond)},
	}

	balances := make(map[int]float64)
	for i, c := range classes {
		fmt.Printf("What is your current %s Amount? : ", c.human_string)
		balance := 0.0
		fmt.Scan(&balance)
		balances[i] = balance
	}

	gain := 0.0
	fmt.Print("How much do you want to purchase? : ")
	fmt.Scan(&gain)

	total := gain
	for _, b := range balances {
		total += b
	}

	deltas := make(map[int]float64)
	for i, c := range classes {
		desired := c.desired_ratio * total
		deltas[i] = desired - balances[i]
		fmt.Printf("You should have %f in %s, have %f delta %f\n",
			desired, c.human_string, balances[i], deltas[i])
	}

	purchase_orders := make(map[int]float64)
	purchase_sum := 0.0
	for i, c := range deltas {
		if c > 0.0 {
			purchase_orders[i] = c
			purchase_sum += c
		}
	}

	scale := gain / purchase_sum
	for i, c := range purchase_orders {
		purchase_orders[i] = c * scale
		fmt.Printf("You should buy %f %ss\n",
			purchase_orders[i], classes[i].human_string)
	}

}
