package model

import (
	"reflect"
)

type World struct {
	liveCells [][]int64
}

func NewWorld(cells [][]int64) World {
	return World{cells}
}

func (w World) IsAlive(cell []int64) bool {
	for _, liveCell := range w.liveCells {
		if reflect.DeepEqual(cell, liveCell) {
			return true
		}
	}
	return false
}

func (w World) GenerateNextIteration() World {
}
