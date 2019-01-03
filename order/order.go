package order

import (
	"log"
	"time"

	"../market"
	"../numeric"
)

type Order struct {
	id           int
	market       *market.Market
	member       *member.Member
	side         Side
	orderType    Type
	state        State
	price        numeric.Numeric
	amount       numeric.Numeric
	originLocked numeric.Numeric
	locked       numeric.Numeric
	received     numeric.Numeric
	given        numeric.Numeric
	tradeCount   int
	createdAt    time.Time
	updatedAt    time.Time
}

func New(
	market *market.Market, member *member.Member,
	side Side, orderType Type,
	price, amount numeric.Numeric,
) (*Order, error) {
	if market == nil {
		return nil, ErrInvalidMarket
	}
	if member == nil {
		return nil, ErrInvalidMember
	}
	switch side {
	case SideBuy, SideSell: // OK
	default:
		return nil, ErrInvalidSide
	}
	switch orderType {
	case TypeLimit, TypeMarket, TypeStop: // OK
	default:
		return nil, ErrInvalidType
	}

	now := time.Now()
	// TODO: Calculate OriginLocked according to Market.

	return nil, &Order{
		market:       market,
		member:       member,
		side:         side,
		orderType:    orderType,
		state:        StateSubmitted,
		source:       source,
		price:        price,
		amount:       amount,
		originLocked: originLocked,
		locked:       numeric.Copy(originLocked),
		received:     numeric.NewZero(),
		given:        numeric.NewZero(),
		tradeCount:   0,
		createdAt:    now,
		updatedAt:    now,
	}
}

func (o *Order) Market() *market.Market {
	if o == nil {
		log.Println(ErrNil.WithMessage("Market(): o is nil"))
		return nil
	}
	return o.market
}

func (o *Order) BidAsset() *asset.Asset {
	if o == nil {
		log.Println(ErrNil.WithMessage("BidAsset(): o is nil"))
		return nil
	}
	return o.market.quoteAsset
}

func (o *Order) AskAsset() *asset.Asset {
	if o == nil {
		log.Println(ErrNil.WithMessage("AskAsset(): o is nil"))
		return nil
	}
	return o.market.baseAsset
}

func (o *Order) Side() Side {
	if o == nil {
		log.Println(ErrNil.WithMessage("Side(): o is nil"))
		return 0
	}
	return o.side
}

func (o *Order) Type() Type {
	if o == nil {
		// TODO log
		return 0
	}
	return o.orderType
}

func (o *Order) State() State {
	if o == nil {
		log.Println(ErrNil.WithMessage("State(): o is nil"))
		return 0
	}
	return o.state
}

func (o *Order) Status() Status {
	if o == nil {
		log.Println(ErrNil.WithMessage("Status(): o is nil"))
		return 0
	}
	if o.receive.IsZero() {
		return StatusNew
	}
	if o.received.IsPositive() {
		if o.amount.IsZero() {
			return StatusDone
		}
		if o.amount.IsPositive() {
			return StatusPartial
		}
	}
	log.Println(ErrInvalidStatus.WithMessage("Status(): o has invalid status"))
	return StatusNone
}

func (o *Order) Member() *member.Member {
	if o == nil {
		log.Println(ErrNil.WithMessage("Member(): o is nil"))
		return 0
	}
	return o.member
}

func (o *Order) Price() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("Price(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.price)
}

func (o *Order) Amount() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("Amount(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.amount)
}

func (o *Order) OriginLocked() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("OriginLocked(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.originLocked)
}

func (o *Order) Locked() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("Locked(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.locked)
}

func (o *Order) Received() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("Received(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.received)
}

func (o *Order) Given() numeric.Numeric {
	if o == nil {
		log.Println(ErrNil.WithMessage("Given(): o is nil"))
		return numeric.NewZero()
	}
	return numeric.Copy(o.given)
}

func (o *Order) TradeCount() int {
	if o == nil {
		log.Println(ErrNil.WithMessage("TradeCount(): o is nil"))
		return 0
	}
	return o.tradeCount
}

func (o *Order) CreatedAt() time.Time {
	if o == nil {
		log.Println(ErrNil.WithMessage("CreatedAt(): o is nil"))
		return time.Now()
	}
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	if o == nil {
		log.Println(ErrNil.WithMessage("UpdatedAt(): o is nil"))
		return time.Now()
	}
	return o.updatedAt
}

func (o *Order) SubAmount(value numeric.Numeric) {
	if o == nil {
		log.Println(ErrNil.WithMessage("SubAmount(): o is nil"))
		return
	}
	o.amount.Sub(value)
}

func (o *Order) SubLocked(value numeric.Numeric) {
	if o == nil {
		log.Println(ErrNil.WithMessage("SubLocked(): o is nil"))
		return
	}
	o.locked.Sub(value)
}

func (o *Order) AddReceived(value numeric.Numeric) {
	if o == nil {
		log.Println(ErrNil.WithMessage("AddReceived(): o is nil"))
		return
	}
	o.received.Add(value)
}

func (o *Order) AddGiven(value numeric.Numeric) {
	if o == nil {
		log.Println(ErrNil.WithMessage("AddGiven(): o is nil"))
		return
	}
	o.given.Add(value)
}

func (o *Order) IncTradeCount() {
	if o == nil {
		log.Println(ErrNil.WithMessage("IncTradeCount(): o is nil"))
		return
	}
	o.tradeCount++
}
