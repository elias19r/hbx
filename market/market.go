package market

import (
	"time"

	"../asset"
	"../numeric"
	"../symbol"
)

type Market struct {
	id          int
	symbol      symbol.Symbol // Unique. e.g.: "XRP/BTC"
	baseAsset   *asset.Asset
	quoteAsset  *asset.Asset
	priceScale  int
	amountScale int
	minPrice    numeric.Numeric
	maxPrice    numeric.Numeric
	minAmount   numeric.Numeric
	maxAmount   numeric.Numeric
	enabled     bool
	visible     bool
	displayPos  int
	createdAt   time.Time
	updatedAt   time.Time
}

func New(sym symbol.Symbol, baseAsset, quoteAsset *asset.Asset) (*Market, error) {
	if err := sym.Validate(); err != nil {
		return nil, err
	}
	if err := quoteAsset.Validate(); err != nil {
		return nil, err
	}
	if err := baseAsset.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()

	return &Market{
		symbol:      sym,
		baseAsset:   baseAsset,
		quoteAsset:  quoteAsset,
		priceScale:  2,
		amountScale: 8,
		enabled:     true,
		visible:     true,
		displayPos:  0,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

func (m *Market) Symbol() symbol.Symbol {
	return m.symbol
}

func (m *Market) BaseAsset() *asset.Asset {
	return m.baseAsset
}

func (m *Market) QuoteAsset() *asset.Asset {
	return m.quoteAsset
}

func (m *Market) DisplayPos() int {
	return m.displayPos
}

func (m *Market) SetDisplayPos(pos int) {
	m.displayPos = pos
}

func (m *Market) Visible() bool {
	return m.visible
}

func (m *Market) SetVisible(visible bool) {
	m.visible = visible
}

func (m *Market) Enabled() bool {
	return m.enabled
}

func (m *Market) SetEnabled(enabled bool) {
	m.enabled = enabled
}

func (m *Market) PriceScale() int {
	return m.priceScale
}

func (m *Market) SetPriceScale(scale int) {
	m.priceScale = scale
}

func (m *Market) AmountScale() int {
	return m.amountScale
}

func (m *Market) SetAmountScale(scale int) {
	m.amountScale = scale
}

func (m *Market) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Market) UpdatedAt() time.Time {
	return m.updatedAt
}
