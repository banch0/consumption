package main

func main() {
	rate := 16
	fuel := 24
	fuelConsumption(rate, fuel)
}

func fuelConsumption(rate, fuel int) int {
	distance := 100
	result := fuel * distance / rate
	return result
}
