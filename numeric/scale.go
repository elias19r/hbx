package numeric

import "math/big"

const (
	// Scale determines the Numeric's scale. It is fixed to 42.
	// The maximum precision is not specified.
	Scale = 42
)

// zero is an always zero big.Int. It must be considered constant.
var zero = big.NewInt(0)

// pow10 stores powers of ten from 10^0 to 10^Scale. It is used throughout
// the package and we must considered it constant.
var pow10 [Scale + 1]*big.Int

func init() {
	// Populate pow10.
	pow10[0] = big.NewInt(1)      // 10^0
	pow10[1] = big.NewInt(10)     // 10^1
	for i := 2; i <= Scale; i++ { // 10^2 to 10^Scale
		pow10[i] = new(big.Int).Mul(pow10[1], pow10[i-1])
	}
}
