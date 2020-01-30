package main

import (
	"testing"
)

func TestFuel(t *testing.T) {
	TestCases := []struct {
		name   string
		fuel   int
		result int
		rate   int
		want   int
	}{
		{name: "Fuel more", fuel: 24, rate: 16, want: 150},
		{name: "Fuel less", fuel: 14, rate: 16, want: 87},
		{name: "Fuel equal", fuel: 15, rate: 15, want: 100},
	}

	for _, testCases := range TestCases {
		got := fuelConsumption(testCases.rate, testCases.fuel)
		if got != testCases.want {
			t.Error("Test error")
		}
	}
}
