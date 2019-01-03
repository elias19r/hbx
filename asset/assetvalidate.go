package asset

import "log"

// Minimum and maximum scale a valid Asset must have.
const (
	MinScale = 0
	MaxScale = 18
)

func (a *Asset) Validate() error {
	if a == nil {
		err := ErrNil.WithMessage("Validate(): a is nil")
		log.Println(err)
		return err
	}
	return validateScale(a.scale)
}

func validateScale(scale int) error {
	if scale < MinScale || scale > MaxScale {
		return ErrInvalidScale.WithMessage(
			"validateScale(): scale is %d but must be in [%d, %d]",
			scale, MinScale, MaxScale,
		)
	}
	return nil
}
