// Package configcodec provides virtual machine configuration encoding helpers.
//
// It is intended to wrap the private configuration encoder, decoder, start
// options, and save options APIs so callers can serialize VM configurations,
// restore them later, and opt into private save or start behaviors such as
// compression, encryption, DFU, or iBoot stop points.
//
// The package should cover:
//
//	Encoding and decoding configurations
//	Builders for save and start option values
//	Helpers that keep path handling explicit and predictable
//
// This package complements snapshot and lifecycle helpers in the parent vzkit
// package.
package configcodec
