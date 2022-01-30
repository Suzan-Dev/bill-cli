package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	return Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (bill *Bill) format() string {
	formattedStr := "Bill Breakdown:\n"
	total := 0.0

	for key, value := range bill.items {
		formattedStr += fmt.Sprintf("%-15s ...$%.2f\n", key+":", value)
		total += value
	}

	formattedStr += fmt.Sprintf("%-15v ...$%.2f\n", "Tip:", bill.tip)
	formattedStr += fmt.Sprintf("%-15v ...$%.2f", "Total:", total+bill.tip)

	return formattedStr
}

func (bill *Bill) addTip(tip float64) {
	bill.tip += tip
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) save() {
	data := []byte(bill.format())

	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved!")
}
