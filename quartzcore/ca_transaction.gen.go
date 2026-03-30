// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATransaction] class.
var (
	_CATransactionClass     CATransactionClass
	_CATransactionClassOnce sync.Once
)

func getCATransactionClass() CATransactionClass {
	_CATransactionClassOnce.Do(func() {
		_CATransactionClass = CATransactionClass{class: objc.GetClass("CATransaction")}
	})
	return _CATransactionClass
}

// GetCATransactionClass returns the class object for CATransaction.
func GetCATransactionClass() CATransactionClass {
	return getCATransactionClass()
}

type CATransactionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CATransactionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATransactionClass) Alloc() CATransaction {
	rv := objc.Send[CATransaction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A mechanism for grouping multiple layer-tree operations into atomic updates
// to the render tree.
//
// # Overview
//
// [CATransaction] is the Core Animation mechanism for batching multiple
// layer-tree operations into atomic updates to the render tree. Every
// modification to a layer tree must be part of a transaction. Nested
// transactions are supported.
//
// Core Animation supports two types of transactions: transactions and
// transactions. Implicit transactions are created automatically when the
// layer tree is modified by a thread without an active transaction and are
// committed automatically when the thread’s runloop next iterates. Explicit
// transactions occur when the the application sends the [CATransaction] class
// a [CATransaction.Begin] message before modifying the layer tree, and a [CATransaction.Commit] message
// afterwards.
//
// [CATransaction] allows you to override default animation properties that
// are set for animatable properties. You can customize duration, timing
// function, whether changes to properties trigger animations, and provide a
// handler that informs you when all animations from the transaction group are
// completed.
//
// During a transaction you can temporarily acquire a recursive spin lock for
// managing property atomicity.
//
// [CATransaction] supports nested transactions. The following code shows how
// you can fade out a layer (named `transitioningLayer`) over a 2 second
// duration while scaling it to three times its original size. The scale
// animation is within a nested transaction with its own duration of 1 second.
// After the outer transaction completes, a completion block removes
// `transitioningLayer` from its parent layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction
type CATransaction struct {
	objectivec.Object
}

// CATransactionFromID constructs a [CATransaction] from an objc.ID.
//
// A mechanism for grouping multiple layer-tree operations into atomic updates
// to the render tree.
func CATransactionFromID(id objc.ID) CATransaction {
	return CATransaction{objectivec.Object{ID: id}}
}

// NOTE: CATransaction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATransaction] class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction
type ICATransaction interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t CATransaction) Init() CATransaction {
	rv := objc.Send[CATransaction](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATransaction) Autorelease() CATransaction {
	rv := objc.Send[CATransaction](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATransaction creates a new CATransaction instance.
func NewCATransaction() CATransaction {
	class := getCATransactionClass()
	rv := objc.Send[CATransaction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Begin a new transaction for the current thread.
//
// # Discussion
//
// The transaction is nested within the thread’s current transaction, if
// there is one.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/begin()
func (_CATransactionClass CATransactionClass) Begin() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("begin"))
}

// Commit all changes made during the current transaction.
//
// # Discussion
//
// Raises an exception if no current transaction exists.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/commit()
func (_CATransactionClass CATransactionClass) Commit() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("commit"))
}

// Flushes any extant implicit transaction.
//
// # Discussion
//
// Delays the commit until any nested explicit transactions have completed.
//
// Flush is typically called automatically at the end of the current runloop,
// regardless of the runloop mode. If your application does not have a
// runloop, you must call this method explicitly.
//
// However, you should attempt to avoid calling `flush` explicitly. By
// allowing `flush` to execute during the runloop your application will
// achieve better performance, atomic screen updates will be preserved, and
// transactions and animations that work from transaction to transaction will
// continue to function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/flush()
func (_CATransactionClass CATransactionClass) Flush() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("flush"))
}

// Returns the animation duration used by all animations within this
// transaction group.
//
// # Return Value
//
// An interval of time used as the duration.
//
// # Discussion
//
// You can retrieve the animation duration for a specific transaction by
// calling the [ValueForKey] method of the transaction object and asking for
// the [kCATransactionAnimationDuration] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/animationDuration()
//
// [kCATransactionAnimationDuration]: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationDuration
func (_CATransactionClass CATransactionClass) AnimationDuration() float64 {
	rv := objc.Send[float64](objc.ID(_CATransactionClass.class), objc.Sel("animationDuration"))
	return rv
}

// Sets the animation duration used by all animations within this transaction
// group.
//
// dur: An interval of time used as the duration.
//
// # Discussion
//
// You can also set the animation duration for a specific transaction object
// by calling the [SetValueForKey] method of that object and specifying the
// [kCATransactionAnimationDuration] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/setAnimationDuration(_:)
//
// [kCATransactionAnimationDuration]: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationDuration
func (_CATransactionClass CATransactionClass) SetAnimationDuration(dur float64) {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("setAnimationDuration:"), dur)
}

// Returns the timing function used for all animations within this transaction
// group.
//
// # Return Value
//
// An instance of [CAMediaTimingFunction].
//
// # Discussion
//
// This is a convenience method that returns the [CAMediaTimingFunction] for
// the [ValueForKey] value returned by the
// [kCATransactionAnimationTimingFunction] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/animationTimingFunction()
//
// [kCATransactionAnimationTimingFunction]: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationTimingFunction
func (_CATransactionClass CATransactionClass) AnimationTimingFunction() CAMediaTimingFunction {
	rv := objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("animationTimingFunction"))
	return CAMediaTimingFunctionFromID(rv)
}

// Sets the timing function used for all animations within this transaction
// group.
//
// function: An instance of [CAMediaTimingFunction].
//
// # Discussion
//
// This is a convenience method that sets the [CAMediaTimingFunction] for the
// [ValueForKey] value of the [kCATransactionAnimationTimingFunction] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/setAnimationTimingFunction(_:)
//
// [kCATransactionAnimationTimingFunction]: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationTimingFunction
func (_CATransactionClass CATransactionClass) SetAnimationTimingFunction(function ICAMediaTimingFunction) {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("setAnimationTimingFunction:"), function)
}

// Returns whether actions triggered as a result of property changes made
// within this transaction group are suppressed.
//
// # Return Value
//
// true if actions are disabled.
//
// # Discussion
//
// This is a convenience method that returns the `boolValue` for the
// [ValueForKey] value returned by the [kCATransactionDisableActions] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/disableActions()
//
// [kCATransactionDisableActions]: https://developer.apple.com/documentation/QuartzCore/kCATransactionDisableActions
func (_CATransactionClass CATransactionClass) DisableActions() bool {
	rv := objc.Send[bool](objc.ID(_CATransactionClass.class), objc.Sel("disableActions"))
	return rv
}

// Sets whether actions triggered as a result of property changes made within
// this transaction group are suppressed.
//
// flag: true, if actions should be disabled.
//
// # Discussion
//
// This is a convenience method that invokes [SetValueForKey] with an
// [NSNumber] containing a true for the [kCATransactionDisableActions] key.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/setDisableActions(_:)
//
// [kCATransactionDisableActions]: https://developer.apple.com/documentation/QuartzCore/kCATransactionDisableActions
func (_CATransactionClass CATransactionClass) SetDisableActions(flag bool) {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("setDisableActions:"), flag)
}

// Returns the completion block object.
//
// # Discussion
//
// See [SetCompletionBlock] for a description of the role of the completion
// block object.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/completionBlock()
func (_CATransactionClass CATransactionClass) CompletionBlock() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("completionBlock"))
}

// Sets the completion block object.
//
// block: A block object called when animations for this transaction group are
// completed.
//
// The block object takes no parameters and returns no value.
//
// # Discussion
//
// The completion block object that is guaranteed to be called (on the main
// thread) as soon as all animations subsequently added by this transaction
// group have completed (or have been removed.) If no animations are added
// before the current transaction group is committed (or the completion block
// is set to a different value,) the block will be invoked immediately.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/setCompletionBlock(_:)
func (_CATransactionClass CATransactionClass) SetCompletionBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("setCompletionBlock:"), _block0)
}

// Attempts to acquire a recursive spin-lock lock, ensuring that returned
// layer values are valid until unlocked.
//
// # Discussion
//
// Core Animation uses a data model that promises not to corrupt the internal
// data structures when called from multiple threads concurrently, but not
// that data returned is still valid if the property was valid on another
// thread. By locking during a transaction you can ensure data that is read,
// modified, and set is correctly managed.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/lock()
func (_CATransactionClass CATransactionClass) Lock() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("lock"))
}

// Relinquishes a previously acquired transaction lock.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransaction/unlock()
func (_CATransactionClass CATransactionClass) Unlock() {
	objc.Send[objc.ID](objc.ID(_CATransactionClass.class), objc.Sel("unlock"))
}

// SetCompletionBlockSync is a synchronous wrapper around [CATransaction.SetCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc CATransactionClass) SetCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	tc.SetCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
