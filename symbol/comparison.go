package symbol

// EQ compares two symbols s == sym
func (s Symbol) EQ(sym Symbol) bool {
	return s.value == sym.value
}
