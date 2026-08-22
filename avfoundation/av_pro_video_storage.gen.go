// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVProVideoStorage] class.
var (
	_AVProVideoStorageClass     AVProVideoStorageClass
	_AVProVideoStorageClassOnce sync.Once
)

func getAVProVideoStorageClass() AVProVideoStorageClass {
	_AVProVideoStorageClassOnce.Do(func() {
		_AVProVideoStorageClass = AVProVideoStorageClass{class: objc.GetClass("AVProVideoStorage")}
	})
	return _AVProVideoStorageClass
}

// GetAVProVideoStorageClass returns the class object for AVProVideoStorage.
func GetAVProVideoStorageClass() AVProVideoStorageClass {
	return getAVProVideoStorageClass()
}

type AVProVideoStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVProVideoStorageClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVProVideoStorageClass) Alloc() AVProVideoStorage {
	rv := objc.Send[AVProVideoStorage](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class to track and manage pre-allocated storage for high data rate video
// capture.
//
// # Overview
//
// [AVProVideoStorage] is a singleton that manages system-wide pre-allocated
// storage used during high data rate video capture to ensure I/O determinism
// and sustain high bandwidth captures (e.g. ProRes).
//
// # Inspecting capacity
//
//   - [AVProVideoStorage.InitialCapacity]: Initial size of Pro Video Storage in bytes.
//   - [AVProVideoStorage.RemainingCapacity]: Current size of Pro Video Storage in bytes.
//   - [AVProVideoStorage.ReplenishCapacityWithCompletionHandler]: Performs a best-effort attempt to restore Pro Video Storage to the initial capacity specified by the user in Settings app.
//
// # Determining whether storage is busy
//
//   - [AVProVideoStorage.BusyReasons]: Whether Pro Video Storage is busy and the associated reasons.
//
// # Presenting the settings interface
//
//   - [AVProVideoStorage.OpenSettings]: Opens the Pro Video Storage UI in Settings app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage
type AVProVideoStorage struct {
	objectivec.Object
}

// AVProVideoStorageFromID constructs a [AVProVideoStorage] from an objc.ID.
//
// A class to track and manage pre-allocated storage for high data rate video
// capture.
func AVProVideoStorageFromID(id objc.ID) AVProVideoStorage {
	return AVProVideoStorage{objectivec.Object{ID: id}}
}

// NOTE: AVProVideoStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVProVideoStorage] class.
//
// # Inspecting capacity
//
//   - [IAVProVideoStorage.InitialCapacity]: Initial size of Pro Video Storage in bytes.
//   - [IAVProVideoStorage.RemainingCapacity]: Current size of Pro Video Storage in bytes.
//   - [IAVProVideoStorage.ReplenishCapacityWithCompletionHandler]: Performs a best-effort attempt to restore Pro Video Storage to the initial capacity specified by the user in Settings app.
//
// # Determining whether storage is busy
//
//   - [IAVProVideoStorage.BusyReasons]: Whether Pro Video Storage is busy and the associated reasons.
//
// # Presenting the settings interface
//
//   - [IAVProVideoStorage.OpenSettings]: Opens the Pro Video Storage UI in Settings app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage
type IAVProVideoStorage interface {
	objectivec.IObject

	// Topic: Inspecting capacity

	// Initial size of Pro Video Storage in bytes.
	InitialCapacity() int
	// Current size of Pro Video Storage in bytes.
	RemainingCapacity() int
	// Performs a best-effort attempt to restore Pro Video Storage to the initial capacity specified by the user in Settings app.
	ReplenishCapacityWithCompletionHandler(completionHandler IntErrorHandler)

	// Topic: Determining whether storage is busy

	// Whether Pro Video Storage is busy and the associated reasons.
	BusyReasons() foundation.INSSet

	// Topic: Presenting the settings interface

	// Opens the Pro Video Storage UI in Settings app.
	OpenSettings()
}

// Init initializes the instance.
func (p AVProVideoStorage) Init() AVProVideoStorage {
	rv := objc.Send[AVProVideoStorage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVProVideoStorage) Autorelease() AVProVideoStorage {
	rv := objc.Send[AVProVideoStorage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVProVideoStorage creates a new AVProVideoStorage instance.
func NewAVProVideoStorage() AVProVideoStorage {
	class := getAVProVideoStorageClass()
	rv := objc.Send[AVProVideoStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Performs a best-effort attempt to restore Pro Video Storage to the initial
// capacity specified by the user in Settings app.
//
// completionHandler: The completion handler is called on an arbitrary dispatch queue when the
// replenish operation finishes. The `remainingCapacity` parameter reflects
// the new size in bytes, which may be less than
// [AVProVideoStorage.InitialCapacity]. If the operation fails, the `error`
// parameter is set and `remainingCapacity` is unchanged or -1 if there was a
// failure retrieving the value.
//
// # Discussion
//
// If there is enough readily available free space on the file system, Pro
// Video Storage will be resized to [AVProVideoStorage.InitialCapacity].
// Otherwise, this method will attempt to resize it near that value.
//
// Pro Video Storage is busy when the replenish operation starts and is no
// longer busy when the completion handler is called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/replenishCapacity(completionHandler:)
func (p AVProVideoStorage) ReplenishCapacityWithCompletionHandler(completionHandler IntErrorHandler) {
	_block0, _ := NewIntErrorBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("replenishCapacityWithCompletionHandler:"), _block0)
}

// Opens the Pro Video Storage UI in Settings app.
//
// # Discussion
//
// Presents system UI that allows the user to adjust pre-allocated storage
// capacity.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/openSettings()
func (p AVProVideoStorage) OpenSettings() {
	objc.Send[objc.ID](p.ID, objc.Sel("openSettings"))
}

// Initial size of Pro Video Storage in bytes.
//
// # Return Value
//
// 0 if Pro Video Storage is not configured or -1 if there was a failure while
// extracting information from it.
//
// # Discussion
//
// The initial capacity is defined by the user via the Settings app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/initialCapacity
func (p AVProVideoStorage) InitialCapacity() int {
	rv := objc.Send[int](p.ID, objc.Sel("initialCapacity"))
	return rv
}

// Current size of Pro Video Storage in bytes.
//
// # Return Value
//
// 0 if Pro Video Storage is not configured or -1 if there was a failure while
// extracting information from it.
//
// # Discussion
//
// The remaining capacity decreases as recordings are captured.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/remainingCapacity
func (p AVProVideoStorage) RemainingCapacity() int {
	rv := objc.Send[int](p.ID, objc.Sel("remainingCapacity"))
	return rv
}

// Whether Pro Video Storage is busy and the associated reasons.
//
// # Discussion
//
// A non-empty set indicates that Pro Video Storage is currently being
// modified. While this is non-empty, starting a video capture will fail with
// an error. This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/busyReasons
func (p AVProVideoStorage) BusyReasons() foundation.INSSet {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("busyReasons"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// Returns the singleton instance for Pro Video Storage.
//
// # Return Value
//
// An instance of the Pro Video Storage class if supported; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/shared
func (_AVProVideoStorageClass AVProVideoStorageClass) SharedStorage() AVProVideoStorage {
	rv := objc.Send[objc.ID](objc.ID(_AVProVideoStorageClass.class), objc.Sel("sharedStorage"))
	return AVProVideoStorageFromID(objc.ID(rv))
}

// Whether Pro Video Storage is supported in its current configuration.
//
// # Return Value
//
// [YES] if the device and OS support Pro Video Storage functionality;
// otherwise, [NO].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVProVideoStorage/isSupported
func (_AVProVideoStorageClass AVProVideoStorageClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(_AVProVideoStorageClass.class), objc.Sel("isSupported"))
	return rv
}

// ReplenishCapacity is a synchronous wrapper around [AVProVideoStorage.ReplenishCapacityWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVProVideoStorage) ReplenishCapacity(ctx context.Context) (int, error) {
	type result struct {
		val int
		err error
	}
	done := make(chan result, 1)
	p.ReplenishCapacityWithCompletionHandler(func(val int, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
