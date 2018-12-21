package api

import (
	"../../address"
	"../../numeric"
	"../../transaction"
)

type API interface {
	NewAddress(text string) (*address.Address, error)

	ValidateAddress(address *address.Address) error

	SendToAddress(
		address *address.Address,
		amount numeric.Numeric,
		speed float64,
	) (*transaction.Transaction, error)

	SendToMany(
		addresses map[*address.Address]numeric.Numeric,
		speed float64,
	) (*transaction.Transaction, error)

	FindTransaction(TxID string) (*transaction.Transaction, error)

	TotalBalance() (numeric.Numeric, error)
}
