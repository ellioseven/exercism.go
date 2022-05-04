package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := int(CalculateWorkingCarsPerHour(productionRate, successRate))
	return result / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	const GroupCost = 95_000
	const CarCost = 10_000
	groups := (carsCount / 10) * GroupCost
	remainder := carsCount % 10
	return uint(groups) + (uint(remainder) * CarCost)
}
