//go:build darwin

package ane

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/tmc/apple/private/appleneuralengine"
)

// Client manages a connection to the ANE hardware.
type Client struct {
	aneClient appleneuralengine.ANEClient
	info      DeviceInfo
	mu        sync.Mutex
	closed    bool
	compiles  atomic.Int64
}

// Open establishes a connection to the ANE hardware.
// Returns ErrNoANE if no ANE is available.
func Open() (*Client, error) {
	info, err := Probe()
	if err != nil {
		return nil, err
	}

	ac, err := acquireClient()
	if err != nil {
		return nil, fmt.Errorf("ane open: %w", err)
	}

	c := &Client{
		aneClient: ac,
		info:      info,
	}
	runtime.SetFinalizer(c, (*Client).Close)
	return c, nil
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

// Close releases the client resources.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	runtime.SetFinalizer(c, nil)
	return nil
}

// Info returns the device information for this client.
func (c *Client) Info() DeviceInfo { return c.info }

// CompileCount returns the number of compilations performed.
func (c *Client) CompileCount() int64 { return c.compiles.Load() }

// ClientObjcID returns the ObjC object pointer for the underlying ANEClient.
func (c *Client) ClientObjcID() uintptr { return uintptr(c.aneClient.ID) }
