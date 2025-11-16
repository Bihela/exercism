package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    productionRateFloat := float64(productionRate) / 60
	floatPerMinute := productionRateFloat * successRate / 100
    return int(floatPerMinute)  
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    carsCountMod:= carsCount % 10
    carCountDivide := carsCount / 10
    return uint(carCountDivide * 95000 + carsCountMod * 10000)
}
