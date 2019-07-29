package numeric

// Numeric represents a decimal number and is implemented using big.Int.
//
// Numerics must be created using the New() function and errors must be checked.
// However, Numeric operations are built to be "nil-safe", i.e. each function
// deals with nil arguments without returning errors, so we can use Numeric's
// functions and chain them the same way we do with basic integer and float types.
//
// A Numeric with a nil value is considered to represent zero.
//
// Division by zero throws a runtime error (panic).
