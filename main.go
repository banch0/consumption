package main

func main() {
	rate := 16
	fuel := 24
	fuelConsumption(rate, fuel)
}

func fuelConsumption(rate, fuel int) int {
	result := fuel * 100 / rate
	return result
}
