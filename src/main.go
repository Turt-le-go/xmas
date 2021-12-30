package main

import (
	"time"
	"xmas/src/canvas"
)

func main() {
	s := "@"
	canvas.Clear()
	var x,y int
	var dx,dy int = 1,1
	w,h := canvas.GetSize()
	for{
		if x == w { dx = -1}
		if y == h { dy = -1}
		if x == 1 { dx = 1 }
		if y == 1 { dy = 1 }
		
		canvas.DrawSymbol(x,y," ")

		x += dx
		y += dy

		canvas.DrawSymbol(x,y,s)

		if x == w && y == h {break}
		if x == w && y == 0 {break}
		if x == 0 && y == h {break}
		if x == 0 && y == 0 {break}

		time.Sleep(100*time.Millisecond)
	}
	//canvas.Clear()
}
