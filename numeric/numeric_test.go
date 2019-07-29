package numeric

import (
	"log"
	"math/big"
	"testing"
)

// buildNumeric is a helper function to create valid Numerics to use in tests.
func buildNumeric(s string) Numeric {
	n, err := New(s)
	if err != nil {
		log.Fatalf("Must test with valid Numeric\n")
	}

	return n
}

// buildBigInt is a helper function to create valid big.Ints to use in tests.
func buildBigInt(s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		log.Fatalf("Must test with valid big.Int\n")
	}

	return i
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("When given valid formats", func(t *testing.T) {
		t.Parallel()

		type validFormatsTest struct {
			n, bigIntValue string
		}
		validFormats := func() []validFormatsTest {
			return []validFormatsTest{
				{"0.0", "0"},
				{"0000.0", "0"},
				{"0.1", "100000000000000000000000000000000000000000"},
				{"-0.1", "-100000000000000000000000000000000000000000"},
				{"1.0", "1000000000000000000000000000000000000000000"},
				{"-1.0", "-1000000000000000000000000000000000000000000"},
				{"12.1234", "12123400000000000000000000000000000000000000"},
				{"-12.1234", "-12123400000000000000000000000000000000000000"},
				{"12.00000004", "12000000040000000000000000000000000000000000"},
				{"-12.00000004", "-12000000040000000000000000000000000000000000"},
				{"9999999999.003871739120213112", "9999999999003871739120213112000000000000000000000000"},
				{"-9999999999.003871739120213112", "-9999999999003871739120213112000000000000000000000000"},
			}
		}

		t.Run("It creates the underlying big.Int accordingly", func(t *testing.T) {
			t.Parallel()

			for _, test := range validFormats() {
				n := buildNumeric(test.n)
				bigIntValue := buildBigInt(test.bigIntValue)

				if n.value.Cmp(bigIntValue) != 0 {
					t.Errorf("\nWant: %s\nHave: %s\n", bigIntValue, n.value)
				}
			}
		})
	})

	t.Run("When given invalid formats", func(t *testing.T) {
		t.Parallel()

		invalidFormats := func() []string {
			return []string{
				"", "abcd",
				"0", "-0", "+0",
				"0.", "-0.", "+0.",
				".0", "-.0", "+.0",
				"1", "-1", "+1",
				"1.", "-1.", "+1.",
				".1", "-.1", "+.1",
				"10", "-10", "+10",
				"10.", "-10.", "+10.",
				".10", "-.10", "+.10",
				"--1.0", "+1.0",
				"12..12345", ".12.12345", "12.12345.",
				"10e2", "-10e2",
			}
		}

		t.Run("It returns ErrInvalidFormat", func(t *testing.T) {
			t.Parallel()

			for _, test := range invalidFormats() {
				_, err := New(test)

				_, ok := err.(ErrInvalidFormat)
				if !ok {
					t.Errorf("\nWant: %s returns ErrInvalidFormat\n"+
						"Have: %s does not return ErrInvalidFormat\n",
						test, test)
				}
			}
		})
	})
}

func TestCopy(t *testing.T) {
	t.Parallel()

	t.Run("When the given Numeric has an underlying nil big.Int", func(t *testing.T) {
		t.Parallel()

		cp := Copy(Numeric{})

		t.Run("It sets cp to zero", func(t *testing.T) {
			t.Parallel()

			if cp.value.Cmp(zero) != 0 {
				t.Errorf("\nWant: %s\nHave: %s\n", zero, cp.value)
			}
		})
	})

	t.Run("It makes cp's underlying big.Int value equals to n's", func(t *testing.T) {
		t.Parallel()

		n := buildNumeric("123.456789")
		cp := Copy(n)

		if cp.value.Cmp(n.value) != 0 {
			t.Errorf("\nWant: %s\nHave: %s\n", n.value, cp.value)
		}
	})

	t.Run("It creates a new big.Int for c", func(t *testing.T) {
		t.Parallel()

		n := buildNumeric("123.456789")
		cp := Copy(n)

		if cp.value == n.value {
			t.Errorf("\nWant: %p != %p"+
				"\nHave: %p == %p\n",
				cp.value, n.value,
				cp.value, n.value,
			)
		}
	})
}
