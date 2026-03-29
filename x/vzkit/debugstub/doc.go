// Package debugstub provides debug stub configuration helpers.
//
// It is intended to wrap the private debug stub configuration types used for
// GDB-style remote debugging, forwarding debug stubs, and VM- or
// coprocessor-specific debug endpoints.
//
// The package should expose:
//
//	Portable configuration builders
//	Helpers to attach debug stubs to VM configurations
//	Introspection helpers for running VMs
//
// This package is focused on configuration and inspection, not on speaking the
// remote debugger protocol itself.
package debugstub
