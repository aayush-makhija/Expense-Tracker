package main

import "fmt"

type Expenses struct {
	label    string
	amount   float64
	month    int
	category string
}

type Categories struct {
	name   string
	amount float64
}

type Months struct {
	month  int
	amount float64
}

var total_amount float64
var array [100]Expenses
var arrayCategory [20]Categories
var arrayMonth [12]Months
var number int

func addExpense() {
	var label, category string
	var amount float64
	var month, flag, flag1, i, j int
	fmt.Println("Label your expense: ")
	fmt.Scan(&label)
	fmt.Println("Enter amount: ")
	fmt.Scan(&amount)
	fmt.Println("Enter category: ")
	fmt.Scan(&category)
	fmt.Println("Enter month(1-12): ")
	fmt.Scan(&month)
	number++
	var expense Expenses
	expense.label = label
	expense.category = category
	expense.amount = amount
	expense.month = month
	a := 0
	for array[a].amount != 0 {
		a++
	}
	array[a] = expense

	for i = 0; i < 20; i++ {
		if category == arrayCategory[i].name {
			arrayCategory[i].amount += amount
			flag = 1
			break
		}
		if arrayCategory[i].amount == 0 {
			break
		}

	}
	if flag == 0 {
		arrayCategory[i].name = category
		arrayCategory[i].amount = amount
	}
	total_amount += amount

	for j = 0; j < 12; j++ {
		if month == arrayMonth[j].month {
			arrayMonth[j].amount += amount
			flag1 = 1
			break
		}
		if arrayMonth[j].amount == 0 {
			break
		}

	}
	if flag1 == 0 {
		arrayMonth[j].month = month
		arrayMonth[j].amount = amount
	}

}

func viewExpense() {
	var entries int
	fmt.Println("How many entries would you like to see?")
	fmt.Scan(&entries)
	for i := 0; i < entries; i++ {
		fmt.Println("Label :", array[number-i-1].label, "Amount: ", array[number-i-1].amount, "Category: ", array[number-i-1].category, "Month: ", array[number-i-1].month)
	}

}

func categorise() {
	for i := 0; i < 20; i++ {
		if arrayCategory[i].amount == 0 {
			break
		}
		fmt.Println(arrayCategory[i].name, arrayCategory[i].amount, (arrayCategory[i].amount/total_amount)*100, "percent of your expense")
	}

}
func show_monthly() {

	for i := 0; i < 12; i++ {
		if arrayMonth[i].month == 0 {
			break
		}
		fmt.Println(arrayMonth[i].month, arrayMonth[i].amount)
	}

}

func main() {
	var request int
	a := 1
	for a != 0 {
		fmt.Printf("Welcome to Allocate!\nThis expense tracker gives you the option to :\n1. Add new expenses\n2. View existing expenses\n3. Generate reports on the basis of category.\n4. Generate reports for a particular month.\n")
		fmt.Println("Please enter the serial number next to your desired request: ")
		fmt.Scan(&request)
		switch request {
		case 1:
			addExpense()
		case 2:
			viewExpense()
		case 3:
			categorise()
		case 4:
			show_monthly()
		default:
			fmt.Println("Invalid Input")

		}
		var more string
		fmt.Println("Would you like to make another request?(Yes/No)")
		fmt.Scan(&more)
		if more == "No" {
			a = 0
		}

	}
}
