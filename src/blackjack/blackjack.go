package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	userScore := 0
	compScore := 0

	fmt.Println("Blackjack! Computer: "+strconv.Itoa(compScore), "You: "+strconv.Itoa(userScore))
	time.Sleep(2 * time.Second)

}
