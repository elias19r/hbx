package asset

import (
	"log"
	"time"

	"../symbol"
	"./api"
)

type Asset struct {
	id           int
	symbol       symbol.Symbol            // Unique. e.g.: "BTC"
	otherSymbols map[string]symbol.Symbol // e.g.: "XBT"
	fiat         bool
	scale        int
	api          *api.API
	enabled      bool
	visible      bool
	displayPos   int
	createdAt    time.Time
	updatedAt    time.Time
}

func New(sym symbol.Symbol, fiat bool, scale int) (*Asset, error) {
	if err := validateScale(scale); err != nil {
		return nil, ErrInvalid.WithMessage("New()").WithPrevError(err)
	}

	now := time.Now()
	return &Asset{
		symbol:       sym,
		otherSymbols: make(map[string]symbol.Symbol),
		fiat:         fiat,
		scale:        scale,
		api:          nil,
		enabled:      true,
		visible:      true,
		displayPos:   0,
		createdAt:    now,
		updatedAt:    now,
	}, nil
}

func (a *Asset) ID() int {
	if a == nil {
		log.Println(ErrNil.WithMessage("ID(): a is nil"))
		return 0
	}
	return a.id
}

func (a *Asset) Symbol() symbol.Symbol {
	if a == nil {
		log.Println(ErrNil.WithMessage("Symbol(): a is nil"))
		return symbol.Symbol{}
	}
	return a.symbol
}

func (a *Asset) Fiat() bool {
	if a == nil {
		log.Println(ErrNil.WithMessage("Fiat(): a is nil"))
		return false
	}
	return a.fiat
}

func (a *Asset) Scale() int {
	if a == nil {
		log.Println(ErrNil.WithMessage("Scale(): a is nil"))
		return 0
	}
	return a.scale
}

func (a *Asset) API() *api.API {
	if a == nil {
		log.Println(ErrNil.WithMessage("API(): a is nil"))
		return nil
	}
	return a.api
}

func (a *Asset) SetAPI(api *api.API) {
	if a == nil {
		log.Println(ErrNil.WithMessage("SetAPI(): a is nil"))
		return
	}
	a.api = api
	a.updatedAt = time.Now()
}

func (a *Asset) Enabled() bool {
	if a == nil {
		log.Println(ErrNil.WithMessage("Enabled(): a is nil"))
		return false
	}
	return a.enabled
}

func (a *Asset) SetEnabled(enabled bool) {
	if a == nil {
		log.Println(ErrNil.WithMessage("SetEnabled(): a is nil"))
		return
	}
	a.enabled = enabled
	a.updatedAt = time.Now()
}

func (a *Asset) Visible() bool {
	if a == nil {
		log.Println(ErrNil.WithMessage("Visible(): a is nil"))
		return false
	}
	return a.visible
}

func (a *Asset) SetVisible(visible bool) {
	if a == nil {
		log.Println(ErrNil.WithMessage("SetVisible(): a is nil"))
		return
	}
	a.visible = visible
	a.updatedAt = time.Now()
}

func (a *Asset) DisplayPos() int {
	if a == nil {
		log.Println(ErrNil.WithMessage("DisplayPos(): a is nil"))
		return 0
	}
	return a.displayPos
}

func (a *Asset) SetDisplayPos(pos int) {
	if a == nil {
		log.Println(ErrNil.WithMessage("SetDisplayPos(): a is nil"))
		return
	}
	a.displayPos = pos
	a.updatedAt = time.Now()
}
