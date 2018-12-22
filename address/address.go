package address

import "log"

type Address struct {
	address string

	// Used by Ripple (XRP).
	destinationTag int32

	// Used by EOS.
	memo string

	// e.g.: bitcoin:17NJGu7kMncocFEKfLDwmGwvTSgPjMqpHF?label=HBX
	//       ripple:rGNqSsUN9Vko2WmskscMnoaY3natMVh1qD?dt=123456
	uri string
}

func (addr *Address) Address() string {
	if addr == nil {
		log.Println(ErrNil.WithMessage("Address(): addr is nil"))
		return ""
	}
	return addr.address
}

func (addr *Address) DestinationTag() int32 {
	if addr == nil {
		log.Println(ErrNil.WithMessage("DestinationTag(): addr is nil"))
		return 0
	}
	return addr.destinationTag
}

func (addr *Address) Memo() string {
	if addr == nil {
		log.Println(ErrNil.WithMessage("Memo(): addr is nil"))
		return ""
	}
	return addr.memo
}
