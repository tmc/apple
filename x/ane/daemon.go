//go:build darwin

package ane

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

var ErrDaemonClosed = fmt.Errorf("ane: daemon client closed")

// DaemonClient wraps the ANE daemon XPC protocol for model compilation,
// loading, and cache management via the aned system service.
//
// All methods are synchronous — they block until the daemon's async
// reply fires.
type DaemonClient struct {
	conn objectivec.Object // XPC connection proxy conforming to _ANEDaemonProtocol

	mu     sync.Mutex
	closed bool
}

// DaemonOptions configures daemon requests.
type DaemonOptions struct {
	QoS uint32 // scheduling priority (0 uses default of 21)
}

func (o *DaemonOptions) qos() uint32 {
	if o == nil || o.QoS == 0 {
		return 21
	}
	return o.QoS
}

// ConnectDaemon establishes a connection to the ANE daemon (aned)
// via the shared ANEClient XPC proxy.
func ConnectDaemon() (*DaemonClient, error) {
	ac, err := acquireClient()
	if err != nil {
		return nil, fmt.Errorf("ane daemon: %w", err)
	}
	c := &DaemonClient{
		conn: objectivec.Object{ID: ac.ID},
	}
	runtime.SetFinalizer(c, (*DaemonClient).Close)
	return c, nil
}

func (d *DaemonClient) checkClosed() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return ErrDaemonClosed
	}
	return nil
}

// voidReply creates a VoidHandler block and waits for it to fire.
func voidReply() (appleneuralengine.VoidHandler, <-chan struct{}) {
	done := make(chan struct{}, 1)
	return func() { done <- struct{}{} }, done
}

// CompileModel requests the daemon to compile a model descriptor.
func (d *DaemonClient) CompileModel(model objectivec.IObject, opts *DaemonOptions) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("compileModel:sandboxExtension:options:qos:withReply:"),
		model, nil, nil, opts.qos(), block)
	<-done
	return nil
}

// LoadModel requests the daemon to load a compiled model.
func (d *DaemonClient) LoadModel(model objectivec.IObject, opts *DaemonOptions) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("loadModel:sandboxExtension:options:qos:withReply:"),
		model, nil, nil, opts.qos(), block)
	<-done
	return nil
}

// CompiledModelExists checks whether a compiled model exists in the daemon cache.
func (d *DaemonClient) CompiledModelExists(identifier objectivec.IObject) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("compiledModelExistsFor:withReply:"),
		identifier, block)
	<-done
	return nil
}

// CompiledModelExistsHash checks the cache by content hash string.
func (d *DaemonClient) CompiledModelExistsHash(hash string) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	nsHash := objectivec.Object{ID: objc.String(hash)}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("compiledModelExistsMatchingHash:withReply:"),
		nsHash, block)
	<-done
	return nil
}

// PurgeModel removes a compiled model from the daemon cache.
func (d *DaemonClient) PurgeModel(model objectivec.IObject) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("purgeCompiledModel:withReply:"),
		model, block)
	<-done
	return nil
}

// PurgeModelHash removes a compiled model by content hash.
func (d *DaemonClient) PurgeModelHash(hash string) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	nsHash := objectivec.Object{ID: objc.String(hash)}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("purgeCompiledModelMatchingHash:withReply:"),
		nsHash, block)
	<-done
	return nil
}

// UnloadModel requests the daemon to unload a model from memory.
func (d *DaemonClient) UnloadModel(model objectivec.IObject, opts *DaemonOptions) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("unloadModel:options:qos:withReply:"),
		model, nil, opts.qos(), block)
	<-done
	return nil
}

// PrepareChainingReq sets up IOSurface chaining between models for
// zero-copy pipeline execution on the ANE.
func (d *DaemonClient) PrepareChainingReq(model, chainingReq objectivec.IObject, opts *DaemonOptions) error {
	if err := d.checkClosed(); err != nil {
		return err
	}
	reply, done := voidReply()
	block, cleanup := appleneuralengine.NewVoidBlock(reply)
	defer cleanup()
	objc.Send[struct{}](d.conn.ID, objc.Sel("prepareChainingWithModel:options:chainingReq:qos:withReply:"),
		model, nil, chainingReq, opts.qos(), block)
	<-done
	return nil
}

// Close disconnects from the daemon.
func (d *DaemonClient) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return nil
	}
	d.closed = true
	runtime.SetFinalizer(d, nil)
	return nil
}
