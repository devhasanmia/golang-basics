package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)
	fmt.Printf("My name is %v\n", fullName)

	// var fullName string
	// var age int
	// var gpa float32
	// Input: Scan, ScanIn Scanf
	// fmt.Print("Enter your name")
	// fmt.Scan(&fullName)
	// fmt.Printf("My name is %v", fullName)
	// variable naming rules
	// 1. letters, digits, _
	// 2. keywords are not allowed as variable
	// 3. variable can not have space
	// 4. variable name can to start with digit
	// var names string = "MD. HASAN MIA"
	// shortcat way
	// country := "bangladesh"
	// const
	// const shopName string = "HAFSA SMART SOLUTION"
	// fmt.Println(names)
	// fmt.Println(country)
	// fmt.Printf("shop name: %v", shopName)
}
