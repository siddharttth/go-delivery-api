package utils

var ProductWeight = 0.5 // kg per product

var WarehouseProducts = map[string][]string{
	"C1": {"A", "B", "D", "E"},
	"C2": {"C", "F", "G"},
	"C3": {"H", "I"},
}

var Distances = map[string]map[string]int{
	"C1": {"L1": 10, "C2": 8, "C3": 15},
	"C2": {"L1": 12, "C1": 8, "C3": 7},
	"C3": {"L1": 10, "C1": 15, "C2": 7},
	"L1": {"C1": 10, "C2": 12, "C3": 10},
}

var CostPerKmPerKg = 2
