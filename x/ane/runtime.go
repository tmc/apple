//go:build darwin

package ane

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/tmc/apple/private/appleneuralengine"
)

// Runtime manages a connection to the ANE hardware.
type Runtime struct {
	client  appleneuralengine.ANEClient
	info    DeviceInfo
	mu      sync.Mutex
	closed  bool
	compiles atomic.Int64
}

// Open establishes a connection to the ANE hardware.
// Returns ErrNoANE if no ANE is available.
func Open() (*Runtime, error) {
	info, err := Probe()
	if err != nil {
		return nil, err
	}

	client, err := acquireClient()
	if err != nil {
		return nil, fmt.Errorf("ane open: %w", err)
	}

	return &Runtime{
		client: client,
		info:   info,
	}, nil
}

// acquireClient tries multiple strategies to obtain an ANE client connection.
func acquireClient() (appleneuralengine.ANEClient, error) {
	cls := appleneuralengine.GetANEClientClass()

	// Try shared private connection first.
	if c := cls.SharedPrivateConnection(); c != nil && c.GetID() != 0 {
		return *c, nil
	}

	// Try shared connection.
	if c := cls.SharedConnection(); c != nil && c.GetID() != 0 {
		return *c, nil
	}

	// Try creating with restricted access.
	c := appleneuralengine.NewANEClientWithRestrictedAccessAllowed(true)
	if c.ID != 0 {
		return c, nil
	}

	// Last resort: plain new.
	c = appleneuralengine.NewANEClient()
	if c.ID != 0 {
		return c, nil
	}

	return appleneuralengine.ANEClient{}, fmt.Errorf("ane: could not acquire client connection")
}

// Close releases the runtime resources.
func (rt *Runtime) Close() error {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.closed = true
	return nil
}

// Info returns the device information for this runtime.
func (rt *Runtime) Info() DeviceInfo { return rt.info }

// CompileCount returns the number of compilations performed.
func (rt *Runtime) CompileCount() int64 { return rt.compiles.Load() }

// ClientObjcID returns the ObjC object pointer for the underlying ANEClient.
func (rt *Runtime) ClientObjcID() uintptr { return uintptr(rt.client.ID) }
