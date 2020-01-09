package main

import (
	"testing"
)

func TestFuel(t *testing.T) {
	rate := 22
	fuel := 53
	result := fuelConsumption(rate, fuel)
	if result != 240 {
		t.Error("Test error")
	}
}
