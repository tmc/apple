// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [BlockOperation] class.
var (
	_BlockOperationClass     BlockOperationClass
	_BlockOperationClassOnce sync.Once
)

func getBlockOperationClass() BlockOperationClass {
	_BlockOperationClassOnce.Do(func() {
		_BlockOperationClass = BlockOperationClass{class: objc.GetClass("NSBlockOperation")}
	})
	return _BlockOperationClass
}

// GetBlockOperationClass returns the class object for NSBlockOperation.
func GetBlockOperationClass() BlockOperationClass {
	return getBlockOperationClass()
}

type BlockOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BlockOperationClass) Alloc() BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}







// An operation that manages the concurrent execution of one or more blocks.
//
// # Overview
// 
// The [NSBlockOperation] class is a concrete subclass of [NSOperation] that
// manages the concurrent execution of one or more blocks. You can use this
// object to execute several blocks at once without having to create separate
// operation objects for each. When executing more than one block, the
// operation itself is considered finished only when all blocks have finished
// executing.
// 
// Blocks added to a block operation are dispatched with default priority to
// an appropriate work queue. The blocks themselves should not make any
// assumptions about the configuration of their execution environment.
// 
// For more information about blocks, see [Blocks Programming Topics].
//
// [Blocks Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Blocks/Articles/00_Introduction.html#//apple_ref/doc/uid/TP40007502
//
// # Managing the Blocks in the Operation
//
//   - [BlockOperation.AddExecutionBlock]: Adds the specified block to the receiver’s list of blocks to perform.
//   - [BlockOperation.ExecutionBlocks]: The blocks associated with the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/BlockOperation
type BlockOperation struct {
	NSOperation
}

// BlockOperationFromID constructs a [BlockOperation] from an objc.ID.
//
// An operation that manages the concurrent execution of one or more blocks.
func BlockOperationFromID(id objc.ID) BlockOperation {
	return NSBlockOperation{NSOperation: NSOperationFromID(id)}
}

// NSBlockOperationFromID is an alias for [BlockOperationFromID] for cross-framework compatibility.
func NSBlockOperationFromID(id objc.ID) BlockOperation { return BlockOperationFromID(id) }
// NOTE: BlockOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [BlockOperation] class.
//
// # Managing the Blocks in the Operation
//
//   - [IBlockOperation.AddExecutionBlock]: Adds the specified block to the receiver’s list of blocks to perform.
//   - [IBlockOperation.ExecutionBlocks]: The blocks associated with the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/BlockOperation
type IBlockOperation interface {
	INSOperation

	// Topic: Managing the Blocks in the Operation

	// Adds the specified block to the receiver’s list of blocks to perform.
	AddExecutionBlock(block VoidHandler)
	// The blocks associated with the receiver.
	ExecutionBlocks() VoidHandler
}





// Init initializes the instance.
func (b BlockOperation) Init() BlockOperation {
	rv := objc.Send[BlockOperation](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BlockOperation) Autorelease() BlockOperation {
	rv := objc.Send[BlockOperation](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlockOperation creates a new BlockOperation instance.
func NewBlockOperation() BlockOperation {
	class := getBlockOperationClass()
	rv := objc.Send[BlockOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}











// Adds the specified block to the receiver’s list of blocks to perform.
//
// block: The block to add to the receiver’s list. The block should take no
// parameters and have no return value.
//
// # Discussion
// 
// The specified block should not make any assumptions about its execution
// environment.
// 
// Calling this method while the receiver is executing or has already finished
// causes an [NSInvalidArgumentException] exception to be thrown.
//
// See: https://developer.apple.com/documentation/Foundation/BlockOperation/addExecutionBlock(_:)
func (b BlockOperation) AddExecutionBlock(block VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
		objc.Send[objc.ID](b.ID, objc.Sel("addExecutionBlock:"), _block0)
}





// Creates and returns an [NSBlockOperation] object and adds the specified
// block to it.
//
// block: The block to add to the new block operation object’s list. The block
// should take no parameters and have no return value.
//
// # Return Value
// 
// A new block operation object.
//
// See: https://developer.apple.com/documentation/Foundation/BlockOperation/init(block:)
func (_BlockOperationClass BlockOperationClass) BlockOperationWithBlock(block VoidHandler) BlockOperation {
		_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
		rv := objc.Send[objc.ID](objc.ID(_BlockOperationClass.class), objc.Sel("blockOperationWithBlock:"), _block0)
	return NSBlockOperationFromID(rv)
}








// The blocks associated with the receiver.
//
// # Discussion
// 
// The blocks in this array are copies of those originally added using the
// [AddExecutionBlock] method.
//
// See: https://developer.apple.com/documentation/Foundation/BlockOperation/executionBlocks
func (b BlockOperation) ExecutionBlocks() VoidHandler {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("executionBlocks"))
	_ = rv
	return nil
}




















// BlockOperationWithBlockSync is a synchronous wrapper around [BlockOperation.BlockOperationWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (bc BlockOperationClass) BlockOperationWithBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	bc.BlockOperationWithBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AddExecutionBlockSync is a synchronous wrapper around [BlockOperation.AddExecutionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (b BlockOperation) AddExecutionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.AddExecutionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






