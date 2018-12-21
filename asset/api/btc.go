package api

import (
	"../../address"
	"../../numeric"
	"../../transaction"
)

func init() {
	// TODO type assert to enforce API interface implementation.
}

type BTC struct {
}

func (asset *BTC) NewAddress(text string) (*address.Address, error) {
	// TODO
	return nil, nil
}

func (asset *BTC) ValidateAddress(address *address.Address) error {
	// TODO
	return nil
}

func (asset *BTC) SendToAddress(
	address *address.Address,
	amount numeric.Numeric,
	speed float64,
) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *BTC) SendToMany(
	addresses map[*address.Address]numeric.Numeric,
	speed float64,
) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *BTC) FindTransaction(TxID string) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *BTC) TotalBalance() (numeric.Numeric, error) {
	// TODO
	return numeric.Numeric{}, nil
}
