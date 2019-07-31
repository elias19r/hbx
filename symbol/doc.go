package symbol

// Symbol identifies an asset or a pair of different assets,
// e.g.: "BTC", "USD", "BTC/USD"
//
// In case of an asset pair, for example "BTC/USD", "BTC" is representing
// the base asset, and "USD" is representing the quote asset.
//
// Symbols must be created using the New() function and errors must be checked.
