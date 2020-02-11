package main

// Board structure
type Board struct {
	x int
	y int
}

// Coord on Board
type Coord struct {
	x int
	y int
}

func newBoard(x int, y int) Board {
	return Board{x, y}
}
