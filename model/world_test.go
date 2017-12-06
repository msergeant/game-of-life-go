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
