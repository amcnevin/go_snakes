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

func newSnake() Snake {
	snake := Snake{}
	snake.bearing = East
	snake.xPos = 5
	snake.yPos = 5
	snake.alive = true
	return snake
}

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
		growSnake(snake, snake.xPos, snake.yPos)
		// pop off the body
		if len(snake.body) > 1 {
			snake.body = snake.body[1:]
		}
	}
}

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

func didSnakeEatCookie(snake *Snake, cookie *Cookie, score *Score) {
	// are snake and cookie at same spot?
	if snake.xPos == cookie.xPos && snake.yPos == cookie.yPos {
		growSnake(snake, cookie.xPos, cookie.yPos)
		generateNewCookie(cookie)
		incrScore(score)
		snake.length++
	}
}

func growSnake(snake *Snake, x int, y int) {
	coord := Coord{x, y}
	snake.body = append(snake.body, coord)

}

func isAlive(snake Snake) bool {
	return snake.alive
}

func changeBearing(snake *Snake, bearing Bearing) {

	currentBearing := snake.bearing
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
