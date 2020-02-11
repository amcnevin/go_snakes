package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

// Game Container
type Game struct {
	board      Board
	snake      Snake
	cookie     Cookie
	score      Score
	difficulty int
}

// initiate the game
func initGame(d int) Game {
	var x int
	var y int

	// difficulty config
	if d == 1 {
		x = 20
		y = 20
	} else if d == 2 {
		x = 15
		y = 15
	} else if d == 3 {
		x = 10
		y = 10
	}

	snake := newSnake()

	return Game{
		newBoard(x, y),
		snake,
		generateInitCookie(x, y),
		initScore(),
		int(d)}
}

// Clock tick on the game
func tick(game *Game) {
	if isAlive(game.snake) {
		// move the snake
		move(&game.snake)
		// did snake eat cookie?
		didSnakeEatCookie(&game.snake, &game.cookie, &game.score)
		// did the snake die?
		didSnakeDie(&game.board, &game.snake)
	}
}

// obtain the direction from the user
func getUserInput(game *Game, halt chan bool) {

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		select {
		case <-halt:
			return
		default:
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			} else if key == keyboard.KeyEsc {
				break
			}

			var bearing Bearing = Null

			if char == 'a' {
				bearing = West
			} else if char == 'w' {
				bearing = North
			} else if char == 's' {
				bearing = South
			} else if char == 'd' {
				bearing = East
			}

			if bearing != Null {
				changeBearing(&game.snake, bearing)
			}
		}
	}
}

// Game Over Man, Game Over
func isGameOver(game Game) bool {
	return !isAlive(game.snake)
}

// Print the Game Board
func print(game Game) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println()
	fmt.Println()
	board := game.board
	snake := game.snake
	cookie := game.cookie
	green := color.New(color.FgGreen).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()

	for i := 0; i < board.x; i++ {
		for j := 0; j < board.y; j++ {
			if isCookieAt(cookie, i, j) {
				yellow(" C ")
			} else if isSnakeAt(snake, i, j) {
				if isAlive(snake) {
					green(" S ")
				} else {
					red(" S ")
				}
			} else {
				fmt.Print("[ ]")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println(fmt.Sprintf("Score: %d", game.score.score))
	fmt.Println(fmt.Sprintf("Snake Body: %d", snake.length))
	fmt.Println()
}
