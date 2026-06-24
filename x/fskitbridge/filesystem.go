package fskitbridge

import "github.com/tmc/apple/fskit"

// A UnaryFileSystem implements an FSKit unary file system: a file system
// that manages a single resource as a single volume. The Server adapts it
// to the FSUnaryFileSystemOperations protocol.
//
// Errors report errnos to FSKit as described for [Volume].
type UnaryFileSystem interface {
	// Probe inspects a resource and names the file system that would
	// mount it.
	Probe(resource fskit.FSResource) (ProbeResult, error)

	// Load prepares the resource and returns its volume.
	Load(resource fskit.FSResource) (Volume, error)

	// Unload releases the resource. It may be called more than once.
	Unload() error
}

// A ProbeResult names a successfully probed resource.
type ProbeResult struct {
	Name string
}
