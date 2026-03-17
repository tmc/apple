package hiservices

// AXUIElementGetWindowSymbol returns the raw symbol address for the private
// _AXUIElementGetWindow entry point.
//
// The discovered metadata does not include a typed C signature, so callers must
// wrap the returned address manually.
func AXUIElementGetWindowSymbol() uintptr {
	return _AXUIElementGetWindowSymbol()
}
