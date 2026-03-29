// Package mailbox provides experimental mailbox device helpers.
//
// It is intended to wrap private mailbox device configuration and attachment
// APIs for host-guest side channels that do not fit existing console or vsock
// abstractions.
//
// The package should begin with configuration and validation helpers before
// growing any higher-level messaging abstractions.
package mailbox
