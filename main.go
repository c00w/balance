package main

import (
	"fmt"
)

func main() {
	us_stock := 0.0
	us_bond := 0.0
	int_stock := 0.0
	int_bond := 0.0
	gain := 0.0

	fmt.Print("What is your current US Stock Amount? :")
	fmt.Scan(&us_stock)
	fmt.Print("International Stock? :")
	fmt.Scan(&int_stock)
	fmt.Print("US Bond? :")
	fmt.Scan(&us_bond)
	fmt.Print("International Bond? :")
	fmt.Scan(&int_bond)
	fmt.Print("How much do you want to purchase? :")
	fmt.Scan(&gain)

	dom_to_int := 0.7
	stock_to_bond := 0.9

	total := gain + us_stock + us_bond + int_stock + int_bond

	us_stock_desired := dom_to_int * stock_to_bond * total
	fmt.Printf("You should have %f in US Stock, have %f delta %f\n",
		us_stock_desired, us_stock, us_stock_desired-us_stock)

	int_stock_desired := (1 - dom_to_int) * stock_to_bond * total
	fmt.Printf("You should have %f in International Stock, have %f delta %f\n",
		int_stock_desired, int_stock, int_stock_desired-int_stock)

	us_bond_desired := dom_to_int * (1 - stock_to_bond) * total
	fmt.Printf("You should have %f in US Bonds, have %f delta %f\n",
		us_bond_desired, us_bond, us_bond_desired-us_bond)

	int_bond_desired := (1 - dom_to_int) * (1 - stock_to_bond) * total
	fmt.Printf("You should have %f in International Bonds, have %f delta %f\n",
		int_bond_desired, int_bond, int_bond_desired-int_bond)
}
