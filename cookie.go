package main

import (
	"math/rand"
	"time"
)

// Cookie container
type Cookie struct {
	xPos   int
	yPos   int
	boardX int
	boardY int
}

// Is the cookie as a specific coordinate?
func isCookieAt(cookie Cookie, x int, y int) bool {
	return (cookie.xPos == x && cookie.yPos == y)
}

// Initiate a cookie thats bounded by dimensions
func generateInitCookie(maxX int, maxY int) Cookie {
	cookie := Cookie{}
	cookie.boardX = maxX
	cookie.boardY = maxY
	generateNewCookie(&cookie)
	return cookie
}

// generate a new cookie at a random coordinate
func generateNewCookie(cookie *Cookie) {
	rand.Seed(time.Now().UnixNano())
	cookie.xPos = rand.Intn(cookie.boardX)
	cookie.yPos = rand.Intn(cookie.boardY)
}
