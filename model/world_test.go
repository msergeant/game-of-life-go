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
