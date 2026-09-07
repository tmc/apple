// Package usbmux connects to iOS device services through the usbmuxd plist
// protocol. It does not discover raw DFU devices or implement device pairing.
// The zero Client connects to the macOS system daemon. Each operation opens its
// own daemon connection; a successful Dial returns a stream owned by the caller.
package usbmux
