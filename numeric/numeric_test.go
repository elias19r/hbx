package numeric

import (
	"math/big"
	"testing"
)

func TestSCALE(t *testing.T) {
	if SCALE != 36 {
		t.Errorf("\nwant: %d\nhave: %d\n", 36, SCALE)
	}
}

func TestNew(t *testing.T) {
	// Helper function to create bigints for tests.
	var createBigInt = func(s string) *big.Int {
		bigint, ok := new(big.Int).SetString(s, 10)
		if !ok {
			t.Fatalf("createBigInt(): fail to create bigint, test is wrongly written: %s\n", s)
		}
		return bigint
	}

	// Test valid formats.
	validFormats := []struct {
		str   string
		value *big.Int
	}{
		{"0.0", createBigInt("0")},
		{"0.1", createBigInt("100000000000000000000000000000000000")},
		{"-0.1", createBigInt("-100000000000000000000000000000000000")},
		{"1.0", createBigInt("1000000000000000000000000000000000000")},
		{"-1.0", createBigInt("-1000000000000000000000000000000000000")},
		{"12.1234", createBigInt("12123400000000000000000000000000000000")},
		{"12.00000004", createBigInt("12000000040000000000000000000000000000")},
		{"9999999999.003871739120213112", createBigInt("9999999999003871739120213112000000000000000000")},
	}

	for _, test := range validFormats {
		n, err := New(test.str)
		if err != nil {
			t.Fatalf("valid format must create Numeric without error: %s\n", test.str)
		}
		if n.value == nil {
			t.Fatalf("valid format must have non-nil bigint value: %s\n", test.str)
		}

		want := test.value
		have := n.value

		if want.Cmp(have) != 0 {
			t.Errorf("\nwant: %s\nhave: %s\n", want, have)
		}
	}

	// Test invalid formats.
	invalidFormats := []string{
		"0.",
		".1",
		"-.1",
		"10",
		"--1.0",
		"+1.0",
		"12..1234",
		"12.0000.0004",
		"12.00000004.",
	}

	for _, test := range invalidFormats {
		_, err := New(test)
		if err == nil {
			t.Fatalf("invalid format must create Numeric WITH error: %s\n", test)
		}

		want := ErrInvalidFormat
		have, ok := err.(*Error)
		if !ok {
			t.Fatalf("invalid format error must typecast to numeric.Error\n")
		}

		if want.Code() != have.Code() {
			t.Errorf("\nwant: %s\nhave: %s\n", want, have)
		}
	}
}

func TestNewZero(t *testing.T) {
	n := NewZero()
	zero := big.NewInt(0)

	if n.value == nil || n.value.Cmp(zero) != 0 {
		t.Errorf("\nwant: %d\nhave: %d\n", n.value, zero)
	}
}

func TestCopy(t *testing.T) {
	n, err := New("123.456789")
	if err != nil {
		t.Fatalf("must test with a valid Numeric\n")
	}
	cp := Copy(n)

	if n.value == nil || cp.value == nil ||
		n.value == cp.value || n.value.Cmp(cp.value) != 0 {
		t.Errorf("\nwant: %s (%p)\nhave: %s (%p)\n",
			n.value, n.value,
			cp.value, cp.value)
	}
}

func TestAdd(t *testing.T) {
	// TODO
}

func TestSub(t *testing.T) {
	// TODO
}

func TestMul(t *testing.T) {
	// TODO
}
func TestNumericAdd(t *testing.T) {
	// Add two positive Numeric values.
	a, err := New("22123.4456789")
	if err != nil {
		t.Fatalf("must test with valid Numeric: invalid a\n")
	}
	aValuePtr := a.value

	b, err := New("922187.112341")
	if err != nil {
		t.Fatalf("must test with valid Numeric: invalid b\n")
	}
	bigIntResult, ok := new(big.Int).SetString("944310558019900000000000000000000000000000", 10)
	if !ok {
		t.Fatalf("must create valid bigint for comparison: invalid bigIntResult\n")
	}

	a.Add(b)

	if a.value != aValuePtr {
		t.Errorf("\nwant: %p\nhave: %p\n", a.value, aValuePtr)
	}
	if a.value.Cmp(bigIntResult) != 0 {
		t.Errorf("\nwant: %s\nhave: %s\n", a.value, bigIntResult)
	}

	// Add positive and negative Numeric values.
	// TODO

	// Add two negative Numeric values.
	// TODO

	// Add negative and positive Numeric values.
	// TODO
}

func TestNumericSub(t *testing.T) {
	// TODO
}

func TestNumericMul(t *testing.T) {
	// TODO
}

func TestNumericTruncate(t *testing.T) {
	// TODO
}

func TestNumericRoundUp(t *testing.T) {
	// TODO
}

func TestNumericIsNegative(t *testing.T) {
	// TODO
}

func TestNumericIsPositive(t *testing.T) {
	// TODO
}

func TestNumericIsZero(t *testing.T) {
	// TODO
}

func TestNumericLT(t *testing.T) {
	// TODO
}

func TestNumericLTE(t *testing.T) {
	// TODO
}

func TestNumericEQ(t *testing.T) {
	// TODO
}

func TestNumericGT(t *testing.T) {
	// TODO
}

func TestNumericGTE(t *testing.T) {
	// TODO
}

func TestNumericstring(t *testing.T) {
	// TODO
}

func TestNumericString(t *testing.T) {
	// TODO
}

func TestNumericStringWithScale(t *testing.T) {
	// TODO
}
