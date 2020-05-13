package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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
	rand.Seed(time.Now().Unix())

	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	mySuit := randFromSliceString(suits)
	myVal := randFromSliceInt(vals)

	return mySuit, myVal
}

func getDeck() (string, int) {
	rand.Seed(time.Now().Unix())

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
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("NOTE: This program is not endorsing gambling in any way. This is just a fun Golang project. Be respoonsible.")
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
	// cDeckV, cDeckT := getDeck()

	uDeckVP := "Your deck: " + uDeckV
	uDeckTP := "Deck value: " + strconv.Itoa(uDeckT)

	fmt.Println(uDeckVP)
	fmt.Println(uDeckTP)

	for {

		var move string

		fmt.Print("(H)it, (S)tand, or (F)old: ")
		fmt.Scanln(&move)
		fmt.Println()

		if strings.TrimRight(move, "\n") == "H" {
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

			if uDeckT > 21 {
				var restart string
				fmt.Print("Oh no! You busted :(. Wanna play again? (y/n): ")
				fmt.Scanln(&restart)

				if strings.TrimRight(restart, "\n") == "y" {
					fmt.Println()
					compScore++
					goto Init
				} else {
					os.Exit(0)
				}
			}
		}
	}
}
