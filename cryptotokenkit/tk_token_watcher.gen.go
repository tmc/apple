// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenWatcher] class.
var (
	_TKTokenWatcherClass     TKTokenWatcherClass
	_TKTokenWatcherClassOnce sync.Once
)

func getTKTokenWatcherClass() TKTokenWatcherClass {
	_TKTokenWatcherClassOnce.Do(func() {
		_TKTokenWatcherClass = TKTokenWatcherClass{class: objc.GetClass("TKTokenWatcher")}
	})
	return _TKTokenWatcherClass
}

// GetTKTokenWatcherClass returns the class object for TKTokenWatcher.
func GetTKTokenWatcherClass() TKTokenWatcherClass {
	return getTKTokenWatcherClass()
}

type TKTokenWatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenWatcherClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenWatcherClass) Alloc() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An object that tracks the tokens available in the system.
//
// # Overview
//
// Create a token watcher and register an insertion handler to be notified
// when tokens are added to the system. You can also add removal handlers for
// specific tokens to be notified when those tokens are removed from the
// system.
//
// # Accessing Token Identifiers
//
//   - [TKTokenWatcher.TokenIDs]: The token IDs currently available in the system.
//
// # Configuring Handlers
//
//   - [TKTokenWatcher.AddRemovalHandlerForTokenID]: Adds a removal handler for the specified token ID.
//   - [TKTokenWatcher.SetInsertionHandler]: Sets an insertion handler closure to be called when a new token is inserted into the system.
//
// # Instance Methods
//
//   - [TKTokenWatcher.TokenInfoForTokenID]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher
type TKTokenWatcher struct {
	objectivec.Object
}

// TKTokenWatcherFromID constructs a [TKTokenWatcher] from an objc.ID.
//
// An object that tracks the tokens available in the system.
func TKTokenWatcherFromID(id objc.ID) TKTokenWatcher {
	return TKTokenWatcher{objectivec.Object{ID: id}}
}

// NOTE: TKTokenWatcher adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenWatcher] class.
//
// # Accessing Token Identifiers
//
//   - [ITKTokenWatcher.TokenIDs]: The token IDs currently available in the system.
//
// # Configuring Handlers
//
//   - [ITKTokenWatcher.AddRemovalHandlerForTokenID]: Adds a removal handler for the specified token ID.
//   - [ITKTokenWatcher.SetInsertionHandler]: Sets an insertion handler closure to be called when a new token is inserted into the system.
//
// # Instance Methods
//
//   - [ITKTokenWatcher.TokenInfoForTokenID]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher
type ITKTokenWatcher interface {
	objectivec.IObject

	// Topic: Accessing Token Identifiers

	// The token IDs currently available in the system.
	TokenIDs() []string

	// Topic: Configuring Handlers

	// Adds a removal handler for the specified token ID.
	AddRemovalHandlerForTokenID(removalHandler StringHandler, tokenID string)
	// Sets an insertion handler closure to be called when a new token is inserted into the system.
	SetInsertionHandler(insertionHandler StringHandler)

	// Topic: Instance Methods

	TokenInfoForTokenID(tokenID string) ITKTokenWatcherTokenInfo
}

// Init initializes the instance.
func (t TKTokenWatcher) Init() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenWatcher) Autorelease() TKTokenWatcher {
	rv := objc.Send[TKTokenWatcher](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenWatcher creates a new TKTokenWatcher instance.
func NewTKTokenWatcher() TKTokenWatcher {
	class := getTKTokenWatcherClass()
	rv := objc.Send[TKTokenWatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a removal handler for the specified token ID.
//
// removalHandler: A block to be called when the specified token is removed. This block takes
// a single argument:
//
// tokenID: The identifier of the removed token.
//
// tokenID: The identifier of the token to watch for removal.
//
// If [TKTokenWatcher.TokenIDs] doesn’t contain `tokenID`,
// `insertionHandler` is executed immediately.
//
// # Discussion
//
// You typically call this method in the `insertionHandler` passed to the
// token watcher initializer.
//
// Adding a removal handler will remove any existing removal handlers for the
// specified token ID.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/addRemovalHandler(_:forTokenID:)
func (t TKTokenWatcher) AddRemovalHandlerForTokenID(removalHandler StringHandler, tokenID string) {
	_block0, _ := NewStringBlock(removalHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("addRemovalHandler:forTokenID:"), _block0, objc.String(tokenID))
}

// Sets an insertion handler closure to be called when a new token is inserted
// into the system.
//
// insertionHandler: A closure to be called whenever a token is added to the system. The closure
// takes a single argument, the tokenID, that identifies the added token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/setInsertionHandler(_:)
func (t TKTokenWatcher) SetInsertionHandler(insertionHandler StringHandler) {
	_block0, _ := NewStringBlock(insertionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("setInsertionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/tokenInfo(forTokenID:)
func (t TKTokenWatcher) TokenInfoForTokenID(tokenID string) ITKTokenWatcherTokenInfo {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenInfoForTokenID:"), objc.String(tokenID))
	return TKTokenWatcherTokenInfoFromID(rv)
}

// The token IDs currently available in the system.
//
// # Discussion
//
// Each string in [TKTokenWatcher.TokenIDs] corresponds to the name of the
// token instance.
//
// You can observe this property to be notified of additions and removals to
// system tokens. See [Key-Value Observing Programming Guide] for more
// information.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/tokenIDs
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (t TKTokenWatcher) TokenIDs() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tokenIDs"))
	return objc.ConvertSliceToStrings(rv)
}

// SetInsertionHandlerSync is a synchronous wrapper around [TKTokenWatcher.SetInsertionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TKTokenWatcher) SetInsertionHandlerSync(ctx context.Context) (*string, error) {
	done := make(chan *string, 1)
	t.SetInsertionHandler(func(val *string) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
