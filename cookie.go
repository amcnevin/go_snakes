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

func isCookieAt(cookie Cookie, x int, y int) bool {
	return (cookie.xPos == x && cookie.yPos == y)
}

func generateInitCookie(maxX int, maxY int) Cookie {
	cookie := Cookie{}
	cookie.boardX = maxX
	cookie.boardY = maxY
	generateNewCookie(&cookie)
	return cookie
}

func generateNewCookie(cookie *Cookie) {
	rand.Seed(time.Now().UnixNano())
	cookie.xPos = rand.Intn(cookie.boardX)
	cookie.yPos = rand.Intn(cookie.boardY)
}
