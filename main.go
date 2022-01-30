package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter the bill name: ", reader)
	return newBill(name)
}

func promptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - Add item, s - Save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("You must enter a number.")
			promptOptions(bill)
		}

		bill.addItem(name, p)
		fmt.Println("Item added successfully.")
		promptOptions(bill)
	case "t":
		tip, _ := getInput("Enter a tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("You must enter a number.")
			promptOptions(bill)
		}

		bill.addTip(t)
		fmt.Println("Tip added successfully.")
		promptOptions(bill)
	case "s":
		bill.save()
	default:
		fmt.Println("Invalid option.")
		promptOptions(bill)
	}
}

func main() {
	bill := createBill()
	promptOptions(bill)
}
