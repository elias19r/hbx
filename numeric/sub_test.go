package numeric

import "testing"

func TestSub(t *testing.T) {
	t.Parallel()

	t.Run("When a or b has an underlying nil big.Int", func(t *testing.T) {
		t.Parallel()

		type subtractionsWithNilTest struct {
			a, b         Numeric
			bigIntResult string
		}
		subtractionsWithNil := func() []subtractionsWithNilTest {
			return []subtractionsWithNilTest{
				{Numeric{}, Numeric{}, "0"},
				{Numeric{}, buildNumeric("123.456"), "-123456000000000000000000000000000000000000000"},
				{Numeric{}, buildNumeric("-123.456"), "123456000000000000000000000000000000000000000"},
				{buildNumeric("123.456"), Numeric{}, "123456000000000000000000000000000000000000000"},
				{buildNumeric("-123.456"), Numeric{}, "-123456000000000000000000000000000000000000000"},
			}
		}

		t.Run("It considers nil values as zeros and sets c's underlying big.Int to a - b accordingly", func(t *testing.T) {
			t.Parallel()

			for _, test := range subtractionsWithNil() {
				a := test.a
				b := test.b
				bigIntResult := buildBigInt(test.bigIntResult)

				c := Sub(a, b)

				if c.value.Cmp(bigIntResult) != 0 {
					t.Errorf("\nWant: %s\nHave: %s\n", bigIntResult, c.value)
				}
			}
		})
	})

	t.Run("When subtracting Numerics c = a - b", func(t *testing.T) {
		t.Parallel()

		t.Run("It returns a new big.Int for c", func(t *testing.T) {
			t.Parallel()

			a := buildNumeric("123.0")
			b := buildNumeric("456.0")

			c := Sub(a, b)

			if c.value == a.value {
				t.Errorf("\nWant: %p != %p"+
					"\nHave: %p == %p\n",
					c.value, a.value,
					c.value, a.value,
				)
			}
			if c.value == b.value {
				t.Errorf("\nWant: %p != %p"+
					"\nHave: %p == %p\n",
					c.value, b.value,
					c.value, b.value,
				)
			}
		})

		type subtractionsTest struct {
			a, b, bigIntResult string
		}
		subtractions := func() []subtractionsTest {
			return []subtractionsTest{
				{"0.0", "0.0", "0"},
				{"1.0", "0.0", "1000000000000000000000000000000000000000000"},
				{"0.0", "1.0", "-1000000000000000000000000000000000000000000"},
				{"1.0", "-1.0", "2000000000000000000000000000000000000000000"},
				{"22123.4456789", "922187.112341", "-900063666662100000000000000000000000000000000000"},
				{"22123.4456789", "-922187.112341", "944310558019900000000000000000000000000000000000"},
				{"-22123.4456789", "922187.112341", "-944310558019900000000000000000000000000000000000"},
				{"-22123.4456789", "-922187.112341", "900063666662100000000000000000000000000000000000"},
			}
		}

		t.Run("It sets c's underlying big.Int to a - b accordingly", func(t *testing.T) {
			t.Parallel()

			for _, test := range subtractions() {
				a := buildNumeric(test.a)
				b := buildNumeric(test.b)
				bigIntResult := buildBigInt(test.bigIntResult)

				c := Sub(a, b)

				if c.value.Cmp(bigIntResult) != 0 {
					t.Errorf("\nWant: %s\nHave: %s\n", bigIntResult, c.value)
				}
			}
		})
	})
}

func TestNumericSub(t *testing.T) {
	t.Parallel()

	t.Run("When a or a's value or b's value is nil", func(t *testing.T) {
		t.Parallel()

		type subtractionsWithNilTest struct {
			a, b         Numeric
			bigIntResult string
		}
		subtractionsWithNil := func() []subtractionsWithNilTest {
			return []subtractionsWithNilTest{
				{Numeric{}, Numeric{}, "0"},
				{Numeric{}, buildNumeric("123.456"), "-123456000000000000000000000000000000000000000"},
				{Numeric{}, buildNumeric("-123.456"), "123456000000000000000000000000000000000000000"},
				{buildNumeric("123.456"), Numeric{}, "123456000000000000000000000000000000000000000"},
				{buildNumeric("-123.456"), Numeric{}, "-123456000000000000000000000000000000000000000"},
			}
		}

		t.Run("It considers nil values as zeros and sets a's underlying big.Int to a - b accordingly", func(t *testing.T) {
			t.Parallel()

			for _, test := range subtractionsWithNil() {
				a := test.a
				b := test.b
				bigIntResult := buildBigInt(test.bigIntResult)

				a.Sub(b)

				if a.value.Cmp(bigIntResult) != 0 {
					t.Errorf("\nWant: %s\nHave: %s\n", bigIntResult, a.value)
				}
			}
		})
	})

	t.Run("When subtracting Numerics a = a - b", func(t *testing.T) {
		t.Parallel()

		t.Run("It overwrites a's value", func(t *testing.T) {
			t.Parallel()

			a := buildNumeric("123.0")
			b := buildNumeric("456.0")
			aValue := a.value

			a.Sub(b)

			if a.value != aValue {
				t.Errorf("\nWant: %p == %p"+
					"\nHave: %p != %p\n",
					a.value, aValue,
					a.value, aValue,
				)
			}
		})

		t.Run("It does not modify b's value", func(t *testing.T) {
			t.Parallel()

			a := buildNumeric("123.0")
			b := buildNumeric("456.0")
			bValue := b.value

			a.Sub(b)

			if b.value.Cmp(bValue) != 0 {
				t.Errorf("\nWant: %s == %s"+
					"\nHave: %s != %s\n",
					b.value, bValue,
					b.value, bValue,
				)
			}
		})

		type subtractionsTest struct {
			a, b, bigIntResult string
		}
		subtractions := func() []subtractionsTest {
			return []subtractionsTest{
				{"0.0", "0.0", "0"},
				{"1.0", "0.0", "1000000000000000000000000000000000000000000"},
				{"0.0", "1.0", "-1000000000000000000000000000000000000000000"},
				{"1.0", "-1.0", "2000000000000000000000000000000000000000000"},
				{"22123.4456789", "922187.112341", "-900063666662100000000000000000000000000000000000"},
				{"22123.4456789", "-922187.112341", "944310558019900000000000000000000000000000000000"},
				{"-22123.4456789", "922187.112341", "-944310558019900000000000000000000000000000000000"},
				{"-22123.4456789", "-922187.112341", "900063666662100000000000000000000000000000000000"},
			}
		}

		t.Run("It sets a's underlying big.Int to a - b accordingly", func(t *testing.T) {
			t.Parallel()

			for _, test := range subtractions() {
				a := buildNumeric(test.a)
				b := buildNumeric(test.b)
				bigIntResult := buildBigInt(test.bigIntResult)

				a.Sub(b)

				if a.value.Cmp(bigIntResult) != 0 {
					t.Errorf("\nWant: %s\nHave: %s\n", bigIntResult, a.value)
				}
			}
		})
	})
}
