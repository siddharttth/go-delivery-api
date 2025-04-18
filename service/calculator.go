package service

import (
	"math"
)

// Calculate the minimum cost to deliver all products
func CalculateMinimumCost(order map[string]int) int {
	// Try starting from each center and get the minimum cost
	minCost := math.MaxInt32
	for _, startCenter := range []string{"C1", "C2", "C3"} {
		cost := calculatePath(startCenter, order)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

// Calculate the delivery cost for a specific path
func calculatePath(startCenter string, order map[string]int) int {
	// Define product weights
	productWeights := map[string]float64{
		"A": 3.0,  // C1
		"B": 2.0,  // C1
		"C": 8.0,  // C1
		"D": 12.0, // C2
		"E": 25.0, // C2
		"F": 15.0, // C2
		"G": 0.5,  // C3
		"H": 1.0,  // C3
		"I": 2.0,  // C3
	}

	// Define centers for each product
	productCenters := map[string]string{
		"A": "C1", "B": "C1", "C": "C1",
		"D": "C2", "E": "C2", "F": "C2",
		"G": "C3", "H": "C3", "I": "C3",
	}

	// Define distances between locations
	distances := map[string]map[string]float64{
		"C1": {"C2": 4.0, "L1": 3.0, "C3": 7.0}, // C3 distance derived from C1->C2->C3
		"C2": {"C1": 4.0, "L1": 2.5, "C3": 3.0},
		"C3": {"C2": 3.0, "L1": 2.0, "C1": 7.0}, // C1 distance derived from C3->C2->C1
		"L1": {"C1": 3.0, "C2": 2.5, "C3": 2.0},
	}

	// Find which centers have products for this order
	centersNeeded := make(map[string]bool)
	for product := range order {
		centersNeeded[productCenters[product]] = true
	}

	// If the start center isn't needed and has no products, choose a different start center
	if !centersNeeded[startCenter] {
		return math.MaxInt32 // This start center isn't viable
	}

	// Remove the start center from centers that need to be visited
	delete(centersNeeded, startCenter)

	// Generate all possible permutations of the remaining centers
	var permutations [][]string
	remainingCenters := make([]string, 0)
	for center := range centersNeeded {
		remainingCenters = append(remainingCenters, center)
	}
	permutations = generatePermutations(remainingCenters)

	// Calculate cost for each permutation
	minCost := math.MaxInt32
	for _, perm := range permutations {
		cost := simulateDelivery(startCenter, perm, order, productWeights, productCenters, distances)
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

// Generate all permutations of centers to visit
func generatePermutations(centers []string) [][]string {
	if len(centers) == 0 {
		return [][]string{{}}
	}

	var result [][]string
	for i, center := range centers {
		// Create a slice without the current center
		remaining := make([]string, 0)
		remaining = append(remaining, centers[:i]...)
		remaining = append(remaining, centers[i+1:]...)

		// Get permutations of the remaining centers
		subPermutations := generatePermutations(remaining)

		// Add the current center to the beginning of each sub-permutation
		for _, subPerm := range subPermutations {
			newPerm := make([]string, 0)
			newPerm = append(newPerm, center)
			newPerm = append(newPerm, subPerm...)
			result = append(result, newPerm)
		}
	}

	return result
}

// Simulate a delivery route and calculate the cost
func simulateDelivery(startCenter string, centerOrder []string, order map[string]int, productWeights map[string]float64, productCenters map[string]string, distances map[string]map[string]float64) int {
	// Calculate cost of the complete delivery route
	totalCost := 0
	currentLocation := startCenter
	currentWeight := 0.0

	// Pick up products from the start center
	for product, qty := range order {
		if productCenters[product] == startCenter {
			currentWeight += float64(qty) * productWeights[product]
		}
	}

	// Deliver to L1
	totalCost += calculateTransportCost(currentWeight, distances[currentLocation]["L1"])
	currentLocation = "L1"

	// Visit each remaining center and then deliver to L1
	for _, center := range centerOrder {
		// Move to the next center (empty vehicle)
		totalCost += calculateTransportCost(0, distances[currentLocation][center])
		currentLocation = center

		// Pick up products from this center
		pickupWeight := 0.0
		for product, qty := range order {
			if productCenters[product] == center {
				pickupWeight += float64(qty) * productWeights[product]
			}
		}

		// Deliver to L1
		totalCost += calculateTransportCost(pickupWeight, distances[currentLocation]["L1"])
		currentLocation = "L1"
	}

	return totalCost
}

// Calculate transport cost based on weight and distance
func calculateTransportCost(weight float64, distance float64) int {
	// For 0-5kg: 10 per unit distance
	// For each additional 5kg: 8 per unit distance
	baseCost := 10.0
	additionalCost := 8.0

	if weight <= 5.0 {
		return int(baseCost * distance)
	}

	// Calculate how many additional 5kg brackets we have
	additionalBrackets := math.Ceil((weight - 5.0) / 5.0)

	// Calculate total cost
	cost := (baseCost + additionalBrackets*additionalCost) * distance
	return int(cost)
}
