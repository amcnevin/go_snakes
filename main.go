package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	printBanner()
	printChoices()

	reader := bufio.NewReader(os.Stdin)

	var difficulty int = -1
	for {

		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)
		if input == "1" {
			difficulty = 1
			break
		} else if input == "2" {
			difficulty = 2
			break
		} else if input == "3" {
			difficulty = 3
			break
		}

	}

	var game = initGame(difficulty)
	haltChannel := make(chan bool)
	go getUserInput(&game, haltChannel)

	for !isGameOver(game) {
		tick(&game)
		print(game)
		time.Sleep(1 * time.Second / 3)
	}

	fmt.Println("GAME OVER")
	haltChannel <- true
	close(haltChannel)

}

func printBanner() {
	fmt.Println("GO SNAKES")
	fmt.Println("Select Difficulty")
}

func printChoices() {
	fmt.Println("1 - Easy")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Hard")
}
