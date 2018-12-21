package asset

// Minimum and maximum scale a valid Asset must have.
const (
	MinScale = 0
	MaxScale = 18
)

func validateScale(scale int) error {
	if scale < MinScale || scale > MaxScale {
		return ErrInvalidScale.WithMessage(
			"validateScale(): scale is %d but must be in [%d, %d]",
			scale, MinScale, MaxScale,
		)
	}
	return nil
}
