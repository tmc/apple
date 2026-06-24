package fskitbridge

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"syscall"

	"github.com/tmc/apple/objc"
)

// An Extension hosts a [Server] inside an ExtensionFoundation extension built
// as a c-archive. It owns the lifecycle a hosted extension needs around the
// Server: lazy, retryable initialization; last-error reporting; a reply
// fallback for calls that arrive before the Server is ready; and panic
// recovery for the exported entry points.
//
// A c-archive cannot re-export Go functions defined in an imported package, so
// the file system's principal class invokes per-extension //export wrappers.
// Each wrapper is a one-line call into the matching Extension method, which
// keeps the lifecycle in one place rather than duplicated across extensions.
//
// The zero Extension is not usable; create one with [NewExtension].
type Extension struct {
	newServer func() (*Server, error)
	fallback  *ReplyBlocks

	mu     sync.Mutex
	server *Server
	err    error
}

// NewExtension returns an Extension that builds its Server with newServer on
// first use. shims names the linked reply block shims so the fallback can
// answer calls that arrive before the Server is ready (see [ReplyBlockShims]).
func NewExtension(shims ReplyBlockShims, newServer func() (*Server, error)) *Extension {
	return &Extension{
		newServer: newServer,
		fallback:  NewReplyBlocksWithShims(shims),
	}
}

// Init builds the Server if it is not already built and returns any error.
// A failure is not sticky: the next call retries, so an early call that races
// extension startup (before the principal class registers) does not poison the
// process.
func (e *Extension) Init() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.server != nil {
		return nil
	}
	srv, err := e.newServer()
	if err != nil {
		e.err = err
		return err
	}
	e.server, e.err = srv, nil
	return nil
}

// Server returns the built Server, or nil before [Extension.Init] succeeds.
func (e *Extension) Server() *Server {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.server
}

// LastError returns the most recent initialization error, or nil.
func (e *Extension) LastError() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.err
}

// NewFileSystem returns a new instance of the file system class, or 0 if the
// Server cannot be initialized.
func (e *Extension) NewFileSystem() objc.ID {
	if e.Init() != nil {
		return 0
	}
	return e.Server().NewFileSystem()
}

// ProbeResource routes probeResource:replyHandler: to the Server, replying
// with EINVAL through the fallback if the Server is not ready.
func (e *Extension) ProbeResource(self, resource, reply objc.ID) {
	defer e.recover("probeResource")
	if e.Init() != nil {
		_ = e.fallback.ObjectError(reply, 0, POSIXError(syscall.EINVAL))
		return
	}
	e.Server().ProbeResource(self, resource, reply)
}

// LoadResource routes loadResource:options:replyHandler: to the Server,
// replying with EINVAL through the fallback if the Server is not ready.
func (e *Extension) LoadResource(self, resource, options, reply objc.ID) {
	defer e.recover("loadResource")
	if e.Init() != nil {
		_ = e.fallback.ObjectError(reply, 0, POSIXError(syscall.EINVAL))
		return
	}
	e.Server().LoadResource(self, resource, options, reply)
}

// UnloadResource routes unloadResource:options:replyHandler: to the Server,
// replying with EINVAL through the fallback if the Server is not ready.
func (e *Extension) UnloadResource(self, resource, options, reply objc.ID) {
	defer e.recover("unloadResource")
	if e.Init() != nil {
		_ = e.fallback.Error(reply, POSIXError(syscall.EINVAL))
		return
	}
	e.Server().UnloadResource(self, resource, options, reply)
}

// recover turns a panic in an exported entry point into a logged process exit.
// A panic must not unwind across the cgo boundary back into FSKit; exiting
// lets the extension host restart cleanly.
func (e *Extension) recover(name string) {
	if r := recover(); r != nil {
		if srv := e.Server(); srv != nil {
			srv.logf("%s panic: %v\n%s", name, r, debug.Stack())
		}
		fmt.Fprintf(os.Stderr, "fskitbridge: %s panic: %v\n%s\n", name, r, debug.Stack())
		os.Exit(2)
	}
}
