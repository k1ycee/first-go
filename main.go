package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("What is your name: ", reader)
	b := newBill(name)
	fmt.Println("Created Bill - ", b.name)
	return b
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price($): ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Print("The price must be a number")
			promptOptions(b)
		}
		b.updateItems(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
	case "t":
		tip, _ := getInput("How much do you want to tip($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Print("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("This is not a valid option")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
