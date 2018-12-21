package asset

import "../symbol"

var data []*Asset
var nextID = 1

func init() {
	// Initialize data skipping the zero index.
	data = append(data, nil)

	// Load data from Redis if any.
	// TODO
}

func NextID() int {
	return nextID
}

func Insert(a *Asset) {
	data = append(data, a)
	a.id = len(data) - 1
}

func FindByID(id int) (*Asset, error) {
	// TODO
	return nil, nil
}

func FindBySymbol(sym symbol.Symbol) (*Asset, error) {
	// TODO
	return nil, nil
}

func Remove(a *Asset) error {
	// TODO
	return nil
}

func RemoveByID(id int) error {
	// TODO
	return nil
}
