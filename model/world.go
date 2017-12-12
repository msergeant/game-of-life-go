package model

import (
	"fmt"
	"reflect"
)

type World struct {
	LiveCells [][]int64 `json:"live_cells"`
}

// NewWorld returns initialized World
func NewWorld(cells [][]int64) World {
	return World{cells}
}

// IsAlive returns true if the given cell is currently alive
func (w World) IsAlive(cell []int64) bool {
	for _, liveCell := range w.LiveCells {
		if reflect.DeepEqual(cell, liveCell) {
			return true
		}
	}
	return false
}

// GenerateNextIteration returns the next generation of the World
func (w World) GenerateNextIteration() World {
	liveCells := [][]int64{}
	for _, cell := range w.cellsToTest() {
		liveNeighbors := w.liveNeighborCount(cell)

		if (w.IsAlive(cell) && liveNeighbors == 2) || liveNeighbors == 3 {
			liveCells = append(liveCells, cell)
		}
	}
	return NewWorld(liveCells)
}

func (w World) cellsToTest() (output [][]int64) {
	setToTest := make(map[string][]int64)

	for _, liveCell := range w.LiveCells {
		neighbors := neighborCells(liveCell)
		for _, cell := range neighbors {
			setToTest[fmt.Sprintf("%d,%d", cell[0], cell[1])] = cell
		}
	}

	for _, value := range setToTest {
		output = append(output, value)
	}

	return
}

func (w World) liveNeighborCount(cell []int64) (output uint8) {
	for _, neighbor := range neighborCells(cell) {
		if w.IsAlive(neighbor) {
			output += 1
		}
	}

	return
}

func neighborCells(cell []int64) [][]int64 {
	x, y := cell[0], cell[1]
	return [][]int64{
		{x - 1, y - 1}, {x - 0, y - 1}, {x + 1, y - 1},
		{x - 1, y - 0}, {x + 1, y - 0},
		{x - 1, y + 1}, {x - 0, y + 1}, {x + 1, y + 1}}
}
