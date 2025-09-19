package main

import "fmt"

func main() {
	// variable naming rules
	// 1. letters, digits, _
	// 2. keywords are not allowed as variable
	// 3. variable can not have space
	// 4. variable name can to start with digit
	var names string = "MD. HASAN MIA"
	// shortcat way
	country := "bangladesh"
	// const
	const shopName string = "HAFSA SMART SOLUTION"
	fmt.Println(names)
	fmt.Println(country)
	fmt.Printf("shop name: %v", shopName)
}
