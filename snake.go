package main

// Bearing of the snake
type Bearing int

// bearings
const (
	North Bearing = iota
	East
	South
	West
	Null
)

// Snake class
type Snake struct {
	xPos    int
	yPos    int
	length  int
	alive   bool
	body    []Coord
	bearing Bearing
}

// build new Snake
func newSnake() Snake {
	snake := Snake{}
	snake.bearing = East
	snake.xPos = 5
	snake.yPos = 5
	snake.alive = true
	return snake
}

// Move the Snake according to its bearing
func move(snake *Snake) {

	if isAlive(*snake) {

		switch snake.bearing {
		case North:
			snake.xPos--
		case South:
			snake.xPos++
		case East:
			snake.yPos++
		case West:
			snake.yPos--
		default:

		}
		// grow the snake if needed		
		growSnake(snake, snake.xPos, snake.yPos)
		// pop off the body
		if len(snake.body) > 1 {
			snake.body = snake.body[1:]
		}
	}
}

// is the Snake at a particular position
func isSnakeAt(snake Snake, x int, y int) bool {
	if snake.xPos == x && snake.yPos == y {
		return true
	}

	for bodyPos := range snake.body {
		if snake.body[bodyPos].x == x && snake.body[bodyPos].y == y {
			return true
		}
	}

	return false
}

// did the snake occupy the same coordinate at the cookie
func didSnakeEatCookie(snake *Snake, cookie *Cookie, score *Score) {
	// are snake and cookie at same spot?
	if snake.xPos == cookie.xPos && snake.yPos == cookie.yPos {
		growSnake(snake, cookie.xPos, cookie.yPos)
		generateNewCookie(cookie)
		incrScore(score)
		snake.length++
	}
}

// grow the snake at a particular coordinate
func growSnake(snake *Snake, x int, y int) {
	coord := Coord{x, y}
	snake.body = append(snake.body, coord)

}

// is the snake not dead yet
func isAlive(snake Snake) bool {
	return snake.alive
}

// change the snakes bearing
func changeBearing(snake *Snake, bearing Bearing) {
	currentBearing := snake.bearing
	// don't allow the snake to travel in opposite direction
	if currentBearing == North && bearing == South {
		return
	}
	if currentBearing == South && bearing == North {
		return
	}
	if currentBearing == East && bearing == West {
		return
	}
	if currentBearing == West && bearing == East {
		return
	}
	snake.bearing = bearing
}

// check to determine if the snake has died
func didSnakeDie(board *Board, snake *Snake) bool {
	// did snake hit wall?
	if board.x <= snake.xPos || 0 > snake.xPos {
		snake.alive = false
		return true
	}

	if board.y <= snake.yPos || 0 > snake.yPos {
		snake.alive = false
		return true
	}

	// TODO did snake hit its body?
	// snake has an x and y for its head,
	// perhaps need to check body -1

	return false
}
