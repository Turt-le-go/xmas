package main

import (
	"math/rand"
	"sort"
	"time"
	"xmas/src/canvas"
	"xmas/src/colors"
)

func main() {
	xtree()
}

func pong(){
	s := "█"
	canvas.Clear()
	var x,y int
	var dx,dy int = 1,1
	w,h := canvas.GetSize()
	for{
		if x == w { dx = -1}
		if y == h { dy = -1}
		if x == 1 { dx = 1 }
		if y == 1 { dy = 1 }
		
		//canvas.DrawSymbol(x,y,"  ")

		x += dx
		y += dy
		ss := colors.Random(s)
		canvas.DrawSymbol(x,y,ss)

		if x == w && y == h {break}
		if x == w && y == 0 {break}
		if x == 0 && y == h {break}
		if x == 0 && y == 0 {break}

		time.Sleep(1*time.Millisecond)
	}
}

func arr_to_coordinates(arr [][]int, dx,dy int) [][2]int {
	var coordinates [][2]int
	for y,row := range arr {
		for x,v := range row {
			if v == 1 { 
				coordinates = append(coordinates, [2]int{x+1+dx,y+1+dy})
			}
		}
	}
	return coordinates
}

func xtree(){
	s := "█"
	xtree_symbol := colors.Green(s)
	xtree_stump_symbol := colors.Yellow(s)
	canvas.Clear()
	w,h := canvas.GetSize()
	var xtree_arr = [][]int{
		{0,0,0,0, 0,0,0, 0,0,1,0,0, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,0, 0,1,1,1,0, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,0, 1,1,1,1,1, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,1, 1,1,1,1,1, 1,0,0, 0,0,0,0},
                                                 
		{0,0,0,0, 0,1,1, 1,1,1,1,1, 1,1,0, 0,0,0,0},
		{0,0,0,0, 1,1,1, 1,1,1,1,1, 1,1,1, 0,0,0,0},
		{0,0,0,0, 0,0,0, 1,1,1,1,1, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,1, 1,1,1,1,1, 1,0,0, 0,0,0,0},
		                                         
		{0,0,0,0, 0,1,1, 1,1,1,1,1, 1,1,0, 0,0,0,0},
		{0,0,0,0, 1,1,1, 1,1,1,1,1, 1,1,1, 0,0,0,0},
		{0,0,0,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,0,0,0},
		{0,0,1,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,1,0,0},
                                                 
		{0,0,0,0, 0,0,1, 1,1,1,1,1, 1,0,0, 0,0,0,0},
		{0,0,0,0, 0,1,1, 1,1,1,1,1, 1,1,0, 0,0,0,0},
		{0,0,0,0, 1,1,1, 1,1,1,1,1, 1,1,1, 0,0,0,0},
                                                 
		{0,0,0,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,0,0,0},
		{0,0,1,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,1,0,0},
		{0,1,1,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,1,1,0},
		{1,1,1,1, 1,1,1, 1,1,1,1,1, 1,1,1, 1,1,1,1},
	}
	var xtree_stump_arr = [][]int{
		{0,0,0,0, 0,0,0, 1,1,1,1,1, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,0, 1,1,1,1,1, 0,0,0, 0,0,0,0},
		{0,0,0,0, 0,0,0, 1,1,1,1,1, 0,0,0, 0,0,0,0},
	}
	xtree_x := w/2 - len(xtree_arr[0])/2
	xtree_y := h/2 - len(xtree_arr)/2
	var xtree_coordinates [][2]int
	xtree_coordinates = arr_to_coordinates(xtree_arr, xtree_x, xtree_y)
	var xtree_stump_coordinates [][2]int
	xtree_stump_coordinates = arr_to_coordinates(xtree_stump_arr, xtree_x, xtree_y + len(xtree_arr))
	for _,xy := range xtree_coordinates {
		x := xy[0]
		y := xy[1]
		canvas.DrawSymbol(x,y,xtree_symbol)
	}
	for _,xy := range xtree_stump_coordinates {
		x := xy[0]
		y := xy[1]
		canvas.DrawSymbol(x,y,xtree_stump_symbol)
	}
	//var coordinatesXmasTreeDecorations [50][2]int
	coordinatesXmasTreeDecorations := make([][2]int, 50)
	for i := range coordinatesXmasTreeDecorations{
		rand.Seed(time.Now().UnixMicro())
		coordinatesXmasTreeDecorations[i] = xtree_coordinates[rand.Intn(len(xtree_coordinates))]
	}
	sort.Slice(coordinatesXmasTreeDecorations, func(i,j int) bool{
		if (coordinatesXmasTreeDecorations[i][1] == coordinatesXmasTreeDecorations[j][1]){
			return coordinatesXmasTreeDecorations[i][0] < coordinatesXmasTreeDecorations[j][0]
		}
		return coordinatesXmasTreeDecorations[i][1] < coordinatesXmasTreeDecorations[j][1]
	})
	for{
		for _,xy := range coordinatesXmasTreeDecorations{
			canvas.DrawSymbol(xy[0], xy[1], colors.Random(s))
			canvas.DrawSymbol(w, h, " ")
			time.Sleep(1*time.Millisecond)
		}	
	}
}
