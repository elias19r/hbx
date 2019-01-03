package order

func (o *Order) Validation() error {
	if o == nil {
		// TODO log
		return ErrNilOrder
	}

	// Market
	if !o.market.Enabled() {
		return ErrInvalidMarket
	}

	// Type
	switch o.orderType {
	case TypeLimit, TypeMarket, TypeStop: // OK, valid types.
	default:
		return ErrInvalidType
	}

	// Side
	switch o.side {
	case SideBuy, SideSell: // OK, valid sides.
	default:
		return ErrInvalidSide
	}

	// State
	switch o.state {
	case StateActive, StateCancelled, StateRejected, StateDone: // OK, valid states.
	default:
		return ErrInvalidState
	}

	// Status
	switch o.Status() {
	case StatusNew, StatusPartial, StatusFilled: // OK, valid status.
	default:
		return ErrInvalidStatus // TODO mensage unkonw status
	}

	// State and Status
	switch o.state {
	case StateActive:
		switch o.Status() {
		case StatusNew, StatusPartial: // OK, valid status for StateActive.
		default:
			return ErrInvalidStatus // TODO message cannot be active and done
		}
	case StateCancelled:
		switch o.Status() {
		case StatusNew, StatusPartial: // OK, valid status for StateCancelled.
		default:
			return ErrInvalidStatus // TODO message cannot be cancelled and done
		}
	case StateRejected:
		switch o.Status() {
		case StatusNew: // OK, valid status for StateRejected.
		default:
			return ErrInvalidStatus // TODO message cannot be rejected and partial or filled
		}
	case StateDone:
		switch o.Status() {
		case StatusFilled: // OK, valid status for StateDone.
		default:
			return ErrInvalidStatus // TODO message cannot be StateDone and StatusNew or StatusFilled
		}
	default:
		return ErrInvalidState
	}

	// Member
	if !o.member.Enabled() {
		return ErrInvalidMember
	}
	if err := o.member.Validation(); err != nil {
		return err
	}

	// Member and Account
	if !o.member.Account(o.market.BaseAsset()).Enabled() {
		return ErrInvalidMember
	}
	if !o.member.Account(o.market.QuoteAsset()).Enabled() {
		return ErrInvalidMember
	}

	// Member and Market
	if !o.member.IsMarketAllowed(o.market) {
		return ErrInvalidMember // TODO custom message
	}

	// Price
	if !o.price.IsPositive() {
		return ErrInvalidPrice
	}
	if o.price.LT(o.market.MinPrice()) {
		return ErrInvalidPrice
	}
	if o.price.GT(o.market.MaxPrice()) {
		return ErrInvalidPrice
	}

	// Amount
	if !o.amount.IsPositive() {
		return ErrInvalidAmount
	}
	if o.amount.LT(o.market.MinAmount()) {
		return ErrInvalidAmount
	}
	if o.amount.GT(o.market.MaxAmount()) {
		return ErrInvalidAmount
	}

	// OriginLocked
	if !o.originLocked.IsPositive() {
		return ErrInvalidOriginLocked
	}
	if o.side == SideBuy && o.originLocked.LT(numeric.Mul(o.price, o.amount)) {
		return ErrInvalidOriginLocked
	}
	if o.side == SideSell && o.originLocked.LT(o.amount) {
		return ErrInvalidOriginLocked
	}

	// Locked
	if o.locked.IsNegative() {
		return ErrInvalidLocked
	}

	// Received
	if o.received.IsNegative() {
		return ErrInvalidReceived
	}
	if o.side == SideBuy && o.received.GT(o.amount) {
		return ErrInvalidReceived
	}

	// Given
	if o.given.IsNegative() {
		return ErrInvalidGiven
	}
	if o.side == SideSell && o.given.GT(o.amount) {
		return ErrInvalidGiven
	}

	// TradeCount
	if o.tradeCount < 0 {
		return ErrInvalidTradeCount
	}

	// CreatedAt and UpdatedAt
	if o.updatedAt.Sub(o.createdAt) < 0 {
		return ErrInvalidTimestamp // TODO message updated cannot be lesser than created
	}

	return nil
}
