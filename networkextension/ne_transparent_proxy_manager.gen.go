// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NETransparentProxyManager] class.
var (
	_NETransparentProxyManagerClass     NETransparentProxyManagerClass
	_NETransparentProxyManagerClassOnce sync.Once
)

func getNETransparentProxyManagerClass() NETransparentProxyManagerClass {
	_NETransparentProxyManagerClassOnce.Do(func() {
		_NETransparentProxyManagerClass = NETransparentProxyManagerClass{class: objc.GetClass("NETransparentProxyManager")}
	})
	return _NETransparentProxyManagerClass
}

// GetNETransparentProxyManagerClass returns the class object for NETransparentProxyManager.
func GetNETransparentProxyManagerClass() NETransparentProxyManagerClass {
	return getNETransparentProxyManagerClass()
}

type NETransparentProxyManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETransparentProxyManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETransparentProxyManagerClass) Alloc() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures and controls transparent proxies.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager
type NETransparentProxyManager struct {
	NEVPNManager
}

// NETransparentProxyManagerFromID constructs a [NETransparentProxyManager] from an objc.ID.
//
// An object that configures and controls transparent proxies.
func NETransparentProxyManagerFromID(id objc.ID) NETransparentProxyManager {
	return NETransparentProxyManager{NEVPNManager: NEVPNManagerFromID(id)}
}

// NOTE: NETransparentProxyManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETransparentProxyManager] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager
type INETransparentProxyManager interface {
	INEVPNManager
}

// Init initializes the instance.
func (t NETransparentProxyManager) Init() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETransparentProxyManager) Autorelease() NETransparentProxyManager {
	rv := objc.Send[NETransparentProxyManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyManager creates a new NETransparentProxyManager instance.
func NewNETransparentProxyManager() NETransparentProxyManager {
	class := getNETransparentProxyManagerClass()
	rv := objc.Send[NETransparentProxyManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads all previously-saved transparent proxy configurations.
//
// completionHandler: A Swift closure or an ObjectiveC block that receives as parameters an array
// of transparent proxy manager instances loaded from disk and an error. If
// the error is `nil`, no error occurred.
//
// # Discussion
//
// This method asychronously reads all previously-saved transparent proxy
// configurations associated with the calling app.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyManager/loadAllFromPreferences(completionHandler:)
func (_NETransparentProxyManagerClass NETransparentProxyManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler NETransparentProxyManagerArrayErrorHandler) {
	_block0, _ := NewNETransparentProxyManagerArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NETransparentProxyManagerClass.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), _block0)
}

// LoadAllFromPreferences is a synchronous wrapper around [NETransparentProxyManager.LoadAllFromPreferencesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (tc NETransparentProxyManagerClass) LoadAllFromPreferences(ctx context.Context) ([]NETransparentProxyManager, error) {
	type result struct {
		val []NETransparentProxyManager
		err error
	}
	done := make(chan result, 1)
	tc.LoadAllFromPreferencesWithCompletionHandler(func(val *[]NETransparentProxyManager, err error) {
		var out []NETransparentProxyManager
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
