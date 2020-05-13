package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func sumSlice(slice []int) int {
	total := 0
	for _, sliceT := range slice {
		total += sliceT
	}
	return total
}

func randFromSliceString(mySlice []string) string {
	randVal := mySlice[rand.Intn(len(mySlice))]
	return randVal
}

func randFromSliceInt(mySlice []int) int {
	randVal := mySlice[rand.Intn(len(mySlice))]
	return randVal
}

func drawCard() (string, int) {
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	mySuit := randFromSliceString(suits)
	myVal := randFromSliceInt(vals)

	return mySuit, myVal
}

func getDeck() (string, int) {

	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	mySuit1 := randFromSliceString(suits)
	myVal1 := randFromSliceInt(vals)

	mySuit2 := randFromSliceString(suits)
	myVal2 := randFromSliceInt(vals)

	myDeckT := sumSlice([]int{myVal1, myVal2})
	myDeckVis := strconv.Itoa(myVal1) + " of " + mySuit1 + ", " + strconv.Itoa(myVal2) + " of " + mySuit2

	return myDeckVis, myDeckT
}

func main() {
	rand.Seed(time.Now().Unix())

	screen.Clear()
	screen.MoveTopLeft()

	accent := color.New(color.FgCyan)

	accent.Printf("NOTE: ")
	fmt.Printf("This program is not endorsing gambling in any way. This is just a fun Golang project. Be respoonsible.\n")
	fmt.Println("Also, this is just how I've seen people play blackjack, not exactly sure if it's super accurate :)")
	fmt.Println()

	userScore := 0
	compScore := 0

Init:
	fmt.Println("Blackjack! Computer: "+strconv.Itoa(compScore), "You: "+strconv.Itoa(userScore))
	fmt.Println()

	fmt.Print("Press 'Enter' to start! ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println()

	uDeckV, uDeckT := getDeck()
	cDeckV, cDeckT := getDeck()

	uDeckVP := "Your deck: " + uDeckV
	uDeckTP := "Deck value: " + strconv.Itoa(uDeckT)

	cDeckVP := "Computer's deck: " + cDeckV
	// cDeckTP := "Deck value: " + strconv.Itoa(cDeckT)

	fmt.Println(uDeckVP)
	fmt.Println(uDeckTP)

	cStood := false

	if uDeckT > 21 {
		var restart string
		fmt.Println()
		fmt.Println("Oh no! You busted :(")
		fmt.Print("Wanna play again? (y/n): ")
		fmt.Scanln(&restart)

		if strings.TrimRight(restart, "\n") == "y" {
			fmt.Println()
			compScore++
			goto Init
		} else if strings.TrimRight(restart, "\n") == "n" {
			os.Exit(0)
		} else {
			fmt.Println()
			compScore++
			goto Init
		}
	}

	if cDeckT > 21 {
		var restart string
		fmt.Println()
		fmt.Println("Woohoo! The computer busted!")
		fmt.Print("Wanna play again? (y/n): ")
		fmt.Scanln(&restart)

		if strings.TrimRight(restart, "\n") == "y" {
			fmt.Println()
			userScore++
			goto Init
		} else if strings.TrimRight(restart, "\n") == "n" {
			os.Exit(0)
		} else {
			fmt.Println()
			userScore++
			goto Init
		}
	}

	for {

		var move string

		fmt.Print("(H)it, (S)tand, or (F)old: ")
		fmt.Scanln(&move)
		fmt.Println()

		switch move {
		case "H":
			draw, drawV := drawCard()
			myDraw := "You drew a " + strconv.Itoa(drawV) + " of " + draw + "."
			fmt.Println(myDraw)

			fmt.Println()

			uDeckVP += ", " + strconv.Itoa(drawV) + " of " + draw
			uDeckT += drawV
			uDeckTP = "Deck value: " + strconv.Itoa(uDeckT)
			fmt.Println(uDeckVP)
			fmt.Println(uDeckTP)
			fmt.Println()

			if uDeckT > 21 && !cStood {
				var restart string
				fmt.Println("Oh no! You busted :(")
				fmt.Print("Wanna play again? (y/n): ")
				fmt.Scanln(&restart)

				if strings.TrimRight(restart, "\n") == "y" {
					fmt.Println()
					compScore++
					goto Init
				} else if strings.TrimRight(restart, "\n") == "n" {
					os.Exit(0)
				} else {
					fmt.Println()
					compScore++
					goto Init
				}
			}

			if cDeckT > 21 {
				var restart string

				fmt.Println("Woohoo! The computer busted!")
				fmt.Println("Computer's deck: " + cDeckVP)
				fmt.Print("Wanna play again? (y/n): ")
				fmt.Scanln(&restart)

				if strings.TrimRight(restart, "\n") == "y" {
					fmt.Println()
					userScore++
					goto Init
				} else if strings.TrimRight(restart, "\n") == "n" {
					os.Exit(0)
				} else {
					fmt.Println()
					userScore++
					goto Init
				}
			} else if cDeckT < 17 {
				cDraw, cDrawV := drawCard()
				cDeckVP += ", " + strconv.Itoa(cDrawV) + " of " + cDraw
				cDeckT += cDrawV
			} else {
				fmt.Println("Computer stands.")
			}

		case "S":
			if cDeckT > uDeckT {
				var restart string

				fmt.Println("Oh no! The computer won :(")
				fmt.Println("Computer's deck: " + cDeckV)
				fmt.Print("Wanna play again? (y/n): ")
				fmt.Scanln(&restart)

				if strings.TrimRight(restart, "\n") == "y" {
					fmt.Println()
					compScore++
					goto Init
				} else if strings.TrimRight(restart, "\n") == "n" {
					os.Exit(0)
				} else {
					fmt.Println()
					compScore++
					goto Init
				}
			} else {
				var restart string

				fmt.Println("Woohoo! You won!")
				fmt.Println("Computer's deck: " + cDeckV)
				fmt.Print("Wanna play again? (y/n): ")
				fmt.Scanln(&restart)

				if strings.TrimRight(restart, "\n") == "y" {
					fmt.Println()
					userScore++
					goto Init
				} else {
					fmt.Println()
					userScore++
					goto Init
				}
			}
		case "F":
			var restart string

			fmt.Println("Oh no! The computer won :(")
			fmt.Println("Computer's deck: " + cDeckV)
			fmt.Print("Wanna play again? (y/n): ")
			fmt.Scanln(&restart)

			if strings.TrimRight(restart, "\n") == "y" {
				fmt.Println()
				compScore++
				goto Init
			} else if strings.TrimRight(restart, "\n") == "n" {
				os.Exit(0)
			} else {
				fmt.Println()
				compScore++
				goto Init
			}
		}
	}
}
