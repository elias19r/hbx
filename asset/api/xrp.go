package api

import (
	"../../address"
	"../../numeric"
	"../../transaction"
)

func init() {
	// TODO type assert to enforce API interface implementation.
}

type XRP struct {
}

func (asset *XRP) NewAddress(text string) (*address.Address, error) {
	// TODO
	return nil, nil
}

func (asset *XRP) ValidateAddress(address *address.Address) error {
	// TODO
	return nil
}

func (asset *XRP) SendToAddress(
	address *address.Address,
	amount numeric.Numeric,
	speed float64,
) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *XRP) SendToMany(
	addresses map[*address.Address]numeric.Numeric,
	speed float64,
) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *XRP) FindTransaction(TxID string) (*transaction.Transaction, error) {
	// TODO
	return nil, nil
}

func (asset *XRP) TotalBalance() (numeric.Numeric, error) {
	// TODO
	return numeric.Numeric{}, nil
}
