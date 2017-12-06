package model

type Point struct {
	x, y int64
}

type World struct {
	liveCells []Point
}

func NewWorld(cells [][]int64) World {
	return World{}
}
