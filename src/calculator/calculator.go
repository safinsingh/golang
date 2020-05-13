package main

import (
	"fmt"
	"strings"
)

func main() {
	var done bool
	done = false
	for !done {
		var text string

		fmt.Print("Enter calculation (A/S/M/D): ")
		fmt.Scanln(&text)

		if strings.TrimRight(text, "\n") == "A" {
			var num1 float32
			var num2 float32

			fmt.Print("Nice! Let's do some addition.\n")
			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			sum := num1 + num2
			sumString := fmt.Sprintf("%f", sum)
			fmt.Print("Your sum is: " + sumString + "\n")
		} else if strings.TrimRight(text, "\n") == "D" {
			var num1 float32
			var num2 float32

			fmt.Print("Nice! Let's do some division.\n")
			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			quotient := num1 / num2
			quotientString := fmt.Sprintf("%f", quotient)
			fmt.Print("Your quotient is: " + quotientString + "\n")
		} else if strings.TrimRight(text, "\n") == "M" {
			var num1 float32
			var num2 float32

			fmt.Print("Nice! Let's do some multiplication.\n")
			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			product := num1 * num2
			productString := fmt.Sprintf("%f", product)
			fmt.Print("Your product is: " + productString + "\n")
		} else if strings.TrimRight(text, "\n") == "S" {
			var num1 float32
			var num2 float32

			fmt.Print("Nice! Let's do some subtraction.\n")
			fmt.Print("Enter your first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter your second number: ")
			fmt.Scanln(&num2)

			diff := num1 - num2
			diffString := fmt.Sprintf("%f", diff)
			fmt.Print("Your product is: " + diffString + "\n")
		} else {
			fmt.Print("Hmm... I didn't recognize that, let's try again!\n")
			main()
		}

		var doneStr string

		fmt.Print("Done? (y/n): ")
		fmt.Scanln(&doneStr)

		if strings.TrimRight(doneStr, "\n") == "y" {
			done = true
		}
	}
}
