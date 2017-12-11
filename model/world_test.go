package model

import (
	"testing"
)

func TestNewWorldEmpty(t *testing.T) {
	world := NewWorld([][]int64{})
	if len(world.liveCells) != 0 {
		t.Errorf("Initializing an empty world failed")
	}
}

func TestNewWorldWithData(t *testing.T) {
	points := [][]int64{
		{0, 0},
		{1, 1},
		{3, 3}}
	world := NewWorld(points)

	if len(world.liveCells) != len(points) {
		t.Errorf("World had incorrect number of initial points, expected %d, got %d", len(points), len(world.liveCells))
	}
}

func TestIsAliveReturnsFalse(t *testing.T) {
	points := [][]int64{
		{0, 0},
		{1, 1},
		{3, 3}}
	world := NewWorld(points)

	if world.IsAlive([]int64{2, 2}) {
		t.Errorf("Expected IsAlive to return false, got true")
	}
}

func TestIsAliveReturnsTrue(t *testing.T) {
	points := [][]int64{
		{0, 0},
		{1, 1},
		{3, 3}}
	world := NewWorld(points)

	for _, cell := range points {
		if !world.IsAlive(cell) {
			t.Errorf("Expected IsAlive to return true, got false for %d", cell)
		}
	}
}

func TestGenerateNextIteration(t *testing.T) {
	points := [][]int64{{7, 0}, {7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 7}, {7, 8}, {7, 9}}
	expectedPoints := [][]int64{
		{6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}, {6, 8},
		{7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 7}, {7, 8},
		{8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}, {8, 6}, {8, 7}, {6, 8}}

	world := NewWorld(points).GenerateNextIteration()

	for _, cell := range expectedPoints {
		if !world.IsAlive(cell) {
			t.Errorf("Expected IsAlive to return true, got false for %d", cell)
		}
	}
}
