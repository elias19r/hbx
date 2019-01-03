package order

func (o *Order) ChangeState(newState State) error {
	if o == nil {
		// TODO log
		return ErrNilOrder
	}
	switch o.state {
	case StateSubmitted:
	case StateActive:
		switch newState {
		case StateCancelled:
			// TOOD unlock funds
			o.state = newState
		default:
			return ErrInvalidChangeState
		}
	case StateCancelled:
	case StateRejected:
	case StateDone:
	default:
		return ErrUnknownState
	}
	return nil
}
