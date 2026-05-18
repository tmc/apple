//go:build darwin

// Package nwpacket provides a small Network.framework-backed net.PacketConn.
//
// It is intentionally narrow: callers choose the local address, interface
// policy, and tracing hook. The package opens clear UDP Network.framework
// listeners and outbound connections, then exposes them through net.PacketConn
// for demos that need to plug into Go or Pion surfaces. Use
// ListenPacketContext when listener startup must be canceled by a caller-owned
// context.
package nwpacket
