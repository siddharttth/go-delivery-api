package service

import (
	"go-delivery-api/utils"
	"math"
)

func CalculateMinimumCost(order map[string]int) int {
	minCost := math.MaxInt32
	for _, startCenter := range []string{"C1", "C2", "C3"} {
		if cost := calculatePath(startCenter, order); cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func calculatePath(startCenter string, order map[string]int) int {
	centersNeeded := make(map[string]bool)
	for product := range order {
		centersNeeded[utils.ProductCenters[product]] = true
	}

	if !centersNeeded[startCenter] {
		return math.MaxInt32
	}
	delete(centersNeeded, startCenter)

	permutations := generatePermutationsFromMap(centersNeeded)
	minCost := math.MaxInt32
	for _, perm := range permutations {
		if cost := simulateDelivery(startCenter, perm, order); cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func simulateDelivery(startCenter string, centerOrder []string, order map[string]int) int {
	totalCost := 0
	currentLocation := startCenter
	currentWeight := 0.0

	for product, qty := range order {
		if utils.ProductCenters[product] == startCenter {
			currentWeight += float64(qty) * utils.ProductWeights[product]
		}
	}
	totalCost += calculateTransportCost(currentWeight, utils.Distances[currentLocation]["L1"])
	currentLocation = "L1"

	for _, center := range centerOrder {
		totalCost += calculateTransportCost(0, utils.Distances[currentLocation][center])
		currentLocation = center

		pickupWeight := 0.0
		for product, qty := range order {
			if utils.ProductCenters[product] == center {
				pickupWeight += float64(qty) * utils.ProductWeights[product]
			}
		}

		totalCost += calculateTransportCost(pickupWeight, utils.Distances[currentLocation]["L1"])
		currentLocation = "L1"
	}
	return totalCost
}

func generatePermutationsFromMap(centerMap map[string]bool) [][]string {
	centers := make([]string, 0, len(centerMap))
	for c := range centerMap {
		centers = append(centers, c)
	}
	return generatePermutations(centers)
}

func generatePermutations(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{{}}
	}
	var result [][]string
	for i, v := range arr {
		rest := append([]string{}, arr[:i]...)
		rest = append(rest, arr[i+1:]...)
		for _, perm := range generatePermutations(rest) {
			result = append(result, append([]string{v}, perm...))
		}
	}
	return result
}

func calculateTransportCost(weight, distance float64) int {
	if weight <= 5.0 {
		return int(utils.BaseCost * distance)
	}
	additionalBrackets := math.Ceil((weight - 5.0) / 5.0)
	cost := (utils.BaseCost + additionalBrackets*utils.AdditionalCost) * distance
	return int(cost)
}
