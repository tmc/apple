//go:build darwin

// Package nwpacket provides a small Network.framework-backed net.PacketConn.
//
// It is intentionally narrow: callers choose the local address, interface
// policy, and tracing hook. The package opens clear UDP Network.framework
// listeners and outbound connections, then exposes them through net.PacketConn
// for demos that need to plug into Go or Pion surfaces. Returned connections
// also implement PathReporter, which reports the observed Network.framework
// path for an established peer. Use ListenPacketContext when listener startup
// must be canceled by a caller-owned context. Config.ConnectRetries can recreate
// an outbound peer connection after a readiness timeout on transient links.
package nwpacket
