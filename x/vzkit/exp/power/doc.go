// Package power provides experimental power source device helpers.
//
// It is intended to wrap private power source and battery device configuration
// APIs so tests can simulate AC power, battery-backed operation, and changing
// host power conditions inside a guest.
//
// The package should remain narrowly focused on configuration and simple state
// modeling until the runtime behavior is better characterized.
package power
