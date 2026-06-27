// Package portfwd relays TCP connections from a host listener to a backend
// reached through a caller-supplied dial function.
//
// It is the convergent core of host→guest port forwarding: a listener accepts
// connections and each is spliced bidirectionally to a freshly dialed backend.
// The backend transport is deliberately abstracted behind a dial closure so the
// same relay serves a guest TCP address, a vsock port, or anything else that
// yields a [net.Conn] — the relay never names a transport.
//
// Parsing of forwarding specifications stays with the caller: a "guest TCP
// port" and a "vsock CID port" are different things and do not share a parser.
// This package forwards bytes; it does not interpret addresses.
package portfwd
