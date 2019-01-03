package market

import (
	"../symbol"
)

var data []*Market
var indexSymbol map[symbol.Symbol]int

func Insert(m *Market) {
	data = append(data, m)
	id := len(data) - 1

	m.id = id
	indexSymbol[m.symbol] = id
}

func FindByID(id int) (*Market, error) {
	if id < 0 || id >= len(data) {
		return nil, ErrMarketNotFound
	}
	return data[id], nil
}

func FindBySymbol(s symbol.Symbol) (*Market, error) {
	id, ok := indexSymbol[s]
	if !ok {
		return nil, ErrMarketNotFound
	}
	return FindByID(id)
}
