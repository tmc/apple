// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameStorageExecutor] class.
var (
	_EspressoDataFrameStorageExecutorClass     EspressoDataFrameStorageExecutorClass
	_EspressoDataFrameStorageExecutorClassOnce sync.Once
)

func getEspressoDataFrameStorageExecutorClass() EspressoDataFrameStorageExecutorClass {
	_EspressoDataFrameStorageExecutorClassOnce.Do(func() {
		_EspressoDataFrameStorageExecutorClass = EspressoDataFrameStorageExecutorClass{class: objc.GetClass("EspressoDataFrameStorageExecutor")}
	})
	return _EspressoDataFrameStorageExecutorClass
}

// GetEspressoDataFrameStorageExecutorClass returns the class object for EspressoDataFrameStorageExecutor.
func GetEspressoDataFrameStorageExecutorClass() EspressoDataFrameStorageExecutorClass {
	return getEspressoDataFrameStorageExecutorClass()
}

type EspressoDataFrameStorageExecutorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameStorageExecutorClass) Alloc() EspressoDataFrameStorageExecutor {
	rv := objc.Send[EspressoDataFrameStorageExecutor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlock]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex]
//   - [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex]
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutor
type EspressoDataFrameStorageExecutor struct {
	objectivec.Object
}

// EspressoDataFrameStorageExecutorFromID constructs a [EspressoDataFrameStorageExecutor] from an objc.ID.
func EspressoDataFrameStorageExecutorFromID(id objc.ID) EspressoDataFrameStorageExecutor {
	return EspressoDataFrameStorageExecutor{objectivec.Object{id}}
}
// Ensure EspressoDataFrameStorageExecutor implements IEspressoDataFrameStorageExecutor.
var _ IEspressoDataFrameStorageExecutor = EspressoDataFrameStorageExecutor{}

// An interface definition for the [EspressoDataFrameStorageExecutor] class.
//
// # Methods
//
//   - [IEspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlock]
//   - [IEspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex]
//   - [IEspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutor
type IEspressoDataFrameStorageExecutor interface {
	objectivec.IObject

	// Topic: Methods

	ExecuteDataFrameStorageWithNetworkBlock(storage objectivec.IObject, network objectivec.IObject, block VoidHandler)
	ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex(storage objectivec.IObject, network objectivec.IObject, block VoidHandler, index VoidHandler)
	ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex(storage objectivec.IObject, network objectivec.IObject, network2 objectivec.IObject, block VoidHandler, index VoidHandler)
}

// Init initializes the instance.
func (e EspressoDataFrameStorageExecutor) Init() EspressoDataFrameStorageExecutor {
	rv := objc.Send[EspressoDataFrameStorageExecutor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameStorageExecutor) Autorelease() EspressoDataFrameStorageExecutor {
	rv := objc.Send[EspressoDataFrameStorageExecutor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameStorageExecutor creates a new EspressoDataFrameStorageExecutor instance.
func NewEspressoDataFrameStorageExecutor() EspressoDataFrameStorageExecutor {
	class := getEspressoDataFrameStorageExecutorClass()
	rv := objc.Send[EspressoDataFrameStorageExecutor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutor/executeDataFrameStorage:withNetwork:block:
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkBlock(storage objectivec.IObject, network objectivec.IObject, block VoidHandler) {
_block2, _cleanup2 := NewVoidBlock(block)
	defer _cleanup2()
	objc.Send[objc.ID](e.ID, objc.Sel("executeDataFrameStorage:withNetwork:block:"), storage, network, _block2)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutor/executeDataFrameStorage:withNetwork:block:blockPrepareForIndex:
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex(storage objectivec.IObject, network objectivec.IObject, block VoidHandler, index VoidHandler) {
_block2, _cleanup2 := NewVoidBlock(block)
	defer _cleanup2()
	_block3, _cleanup3 := NewVoidBlock(index)
	defer _cleanup3()
	objc.Send[objc.ID](e.ID, objc.Sel("executeDataFrameStorage:withNetwork:block:blockPrepareForIndex:"), storage, network, _block2, _block3)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutor/executeDataFrameStorage:withNetwork:referenceNetwork:block:blockPrepareForIndex:
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex(storage objectivec.IObject, network objectivec.IObject, network2 objectivec.IObject, block VoidHandler, index VoidHandler) {
_block3, _cleanup3 := NewVoidBlock(block)
	defer _cleanup3()
	_block4, _cleanup4 := NewVoidBlock(index)
	defer _cleanup4()
	objc.Send[objc.ID](e.ID, objc.Sel("executeDataFrameStorage:withNetwork:referenceNetwork:block:blockPrepareForIndex:"), storage, network, network2, _block3, _block4)
}

// ExecuteDataFrameStorageWithNetworkBlockSync is a synchronous wrapper around [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkBlockSync(ctx context.Context, storage objectivec.IObject, network objectivec.IObject) error {
	done := make(chan struct{}, 1)
	e.ExecuteDataFrameStorageWithNetworkBlock(storage, network, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndexSync is a synchronous wrapper around [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex].
// It blocks until the completion handler fires or the context is cancelled.
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndexSync(ctx context.Context, storage objectivec.IObject, network objectivec.IObject, block VoidHandler) error {
	done := make(chan struct{}, 1)
	e.ExecuteDataFrameStorageWithNetworkBlockBlockPrepareForIndex(storage, network, block, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndexSync is a synchronous wrapper around [EspressoDataFrameStorageExecutor.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex].
// It blocks until the completion handler fires or the context is cancelled.
func (e EspressoDataFrameStorageExecutor) ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndexSync(ctx context.Context, storage objectivec.IObject, network objectivec.IObject, network2 objectivec.IObject, block VoidHandler) error {
	done := make(chan struct{}, 1)
	e.ExecuteDataFrameStorageWithNetworkReferenceNetworkBlockBlockPrepareForIndex(storage, network, network2, block, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

