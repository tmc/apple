// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCSessionFactory] class.
var (
	_AVVCSessionFactoryClass     AVVCSessionFactoryClass
	_AVVCSessionFactoryClassOnce sync.Once
)

func getAVVCSessionFactoryClass() AVVCSessionFactoryClass {
	_AVVCSessionFactoryClassOnce.Do(func() {
		_AVVCSessionFactoryClass = AVVCSessionFactoryClass{class: objc.GetClass("AVVCSessionFactory")}
	})
	return _AVVCSessionFactoryClass
}

// GetAVVCSessionFactoryClass returns the class object for AVVCSessionFactory.
func GetAVVCSessionFactoryClass() AVVCSessionFactoryClass {
	return getAVVCSessionFactoryClass()
}

type AVVCSessionFactoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCSessionFactoryClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCSessionFactoryClass) Alloc() AVVCSessionFactory {
	rv := objc.SendIfResponds[AVVCSessionFactory](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCSessionFactory._wqCreatePrimarySessionManagerIfNeededClientTypeError]
//   - [AVVCSessionFactory._wqSessionAndManagerForContextClientTypeSessionManagerError]
//   - [AVVCSessionFactory.AuxSessionManagers]
//   - [AVVCSessionFactory.CleanupContext]
//   - [AVVCSessionFactory.PrimarySessionManager]
//   - [AVVCSessionFactory.SetPrimarySessionManager]
//   - [AVVCSessionFactory.ReleasePrimarySessionManager]
//   - [AVVCSessionFactory.SessionForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SessionForContextClientTypeError]
//   - [AVVCSessionFactory.SessionForContextCompletion]
//   - [AVVCSessionFactory.SessionForContextError]
//   - [AVVCSessionFactory.SessionManagerForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SessionManagerForContextClientTypeError]
//   - [AVVCSessionFactory.SessionManagerMap]
//   - [AVVCSessionFactory.SetSessionManagerMap]
//   - [AVVCSessionFactory.SetSessionWasCreatedBlock]
//   - [AVVCSessionFactory.SetSessionWillBeDestroyedBlock]
//   - [AVVCSessionFactory.WorkQueue]
//   - [AVVCSessionFactory.SetWorkQueue]
type AVVCSessionFactory struct {
	objectivec.Object
}

// AVVCSessionFactoryFromID constructs a [AVVCSessionFactory] from an objc.ID.
func AVVCSessionFactoryFromID(id objc.ID) AVVCSessionFactory {
	return AVVCSessionFactory{objectivec.Object{ID: id}}
}

// Ensure AVVCSessionFactory implements IAVVCSessionFactory.
var _ IAVVCSessionFactory = AVVCSessionFactory{}

// An interface definition for the [AVVCSessionFactory] class.
//
// # Methods
//
//   - [IAVVCSessionFactory._wqCreatePrimarySessionManagerIfNeededClientTypeError]
//   - [IAVVCSessionFactory._wqSessionAndManagerForContextClientTypeSessionManagerError]
//   - [IAVVCSessionFactory.AuxSessionManagers]
//   - [IAVVCSessionFactory.CleanupContext]
//   - [IAVVCSessionFactory.PrimarySessionManager]
//   - [IAVVCSessionFactory.SetPrimarySessionManager]
//   - [IAVVCSessionFactory.ReleasePrimarySessionManager]
//   - [IAVVCSessionFactory.SessionForContextClientTypeCompletion]
//   - [IAVVCSessionFactory.SessionForContextClientTypeError]
//   - [IAVVCSessionFactory.SessionForContextCompletion]
//   - [IAVVCSessionFactory.SessionForContextError]
//   - [IAVVCSessionFactory.SessionManagerForContextClientTypeCompletion]
//   - [IAVVCSessionFactory.SessionManagerForContextClientTypeError]
//   - [IAVVCSessionFactory.SessionManagerMap]
//   - [IAVVCSessionFactory.SetSessionManagerMap]
//   - [IAVVCSessionFactory.SetSessionWasCreatedBlock]
//   - [IAVVCSessionFactory.SetSessionWillBeDestroyedBlock]
//   - [IAVVCSessionFactory.WorkQueue]
//   - [IAVVCSessionFactory.SetWorkQueue]
type IAVVCSessionFactory interface {
	objectivec.IObject

	// Topic: Methods

	_wqCreatePrimarySessionManagerIfNeededClientTypeError(needed objectivec.IObject, type_ int64) (objectivec.IObject, error)
	_wqSessionAndManagerForContextClientTypeSessionManagerError(context objectivec.IObject, type_ int64, session []objectivec.IObject, manager []objectivec.IObject) error
	AuxSessionManagers() objectivec.IObject
	CleanupContext(context objectivec.IObject)
	PrimarySessionManager() IAVVCSessionManager
	SetPrimarySessionManager(value IAVVCSessionManager)
	ReleasePrimarySessionManager()
	SessionForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler)
	SessionForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error)
	SessionForContextCompletion(context objectivec.IObject, completion VoidHandler)
	SessionForContextError(context objectivec.IObject) (objectivec.IObject, error)
	SessionManagerForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler)
	SessionManagerForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error)
	SessionManagerMap() foundation.INSDictionary
	SetSessionManagerMap(value foundation.INSDictionary)
	SetSessionWasCreatedBlock(block VoidHandler)
	SetSessionWillBeDestroyedBlock(block VoidHandler)
	WorkQueue() objectivec.Object
	SetWorkQueue(value objectivec.Object)
}

// Init initializes the instance.
func (a AVVCSessionFactory) Init() AVVCSessionFactory {
	rv := objc.SendIfResponds[AVVCSessionFactory](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCSessionFactory) Autorelease() AVVCSessionFactory {
	rv := objc.SendIfResponds[AVVCSessionFactory](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCSessionFactory creates a new AVVCSessionFactory instance.
func NewAVVCSessionFactory() AVVCSessionFactory {
	class := getAVVCSessionFactoryClass()
	rv := objc.SendIfResponds[AVVCSessionFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCSessionFactory) _wqCreatePrimarySessionManagerIfNeededClientTypeError(needed objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_wqCreatePrimarySessionManagerIfNeeded:clientType:error:"), needed, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// WqCreatePrimarySessionManagerIfNeededClientTypeError is an exported wrapper for the private method _wqCreatePrimarySessionManagerIfNeededClientTypeError.
func (a AVVCSessionFactory) WqCreatePrimarySessionManagerIfNeededClientTypeError(needed objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_wqCreatePrimarySessionManagerIfNeeded:clientType:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_wqCreatePrimarySessionManagerIfNeeded:clientType:error:"}
		return nil, err
	}
	return a._wqCreatePrimarySessionManagerIfNeededClientTypeError(needed, type_)
}

// CanWqCreatePrimarySessionManagerIfNeededClientTypeError reports whether the receiver responds to the private selector _wqCreatePrimarySessionManagerIfNeeded:clientType:error:.
func (a AVVCSessionFactory) CanWqCreatePrimarySessionManagerIfNeededClientTypeError() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_wqCreatePrimarySessionManagerIfNeeded:clientType:error:"))
}
func (a AVVCSessionFactory) _wqSessionAndManagerForContextClientTypeSessionManagerError(context objectivec.IObject, type_ int64, session []objectivec.IObject, manager []objectivec.IObject) error {
	var errorPtr objc.ID
	objc.Send[struct{}](a.ID, objc.Sel("_wqSessionAndManagerForContext:clientType:session:manager:error:"), context, type_, objectivec.IObjectSliceToNSArray(session), objectivec.IObjectSliceToNSArray(manager), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// WqSessionAndManagerForContextClientTypeSessionManagerError is an exported wrapper for the private method _wqSessionAndManagerForContextClientTypeSessionManagerError.
func (a AVVCSessionFactory) WqSessionAndManagerForContextClientTypeSessionManagerError(context objectivec.IObject, type_ int64, session []objectivec.IObject, manager []objectivec.IObject) error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_wqSessionAndManagerForContext:clientType:session:manager:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_wqSessionAndManagerForContext:clientType:session:manager:error:"}
		return err
	}
	return a._wqSessionAndManagerForContextClientTypeSessionManagerError(context, type_, session, manager)
}

// CanWqSessionAndManagerForContextClientTypeSessionManagerError reports whether the receiver responds to the private selector _wqSessionAndManagerForContext:clientType:session:manager:error:.
func (a AVVCSessionFactory) CanWqSessionAndManagerForContextClientTypeSessionManagerError() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_wqSessionAndManagerForContext:clientType:session:manager:error:"))
}
func (a AVVCSessionFactory) AuxSessionManagers() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("auxSessionManagers"))
	return objectivec.Object{ID: rv}
}
func (a AVVCSessionFactory) CleanupContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("cleanupContext:"), context)
}
func (a AVVCSessionFactory) ReleasePrimarySessionManager() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("releasePrimarySessionManager"))
}
func (a AVVCSessionFactory) SessionForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sessionForContext:clientType:completion:"), context, type_, _block2)
}
func (a AVVCSessionFactory) SessionForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sessionForContext:clientType:error:"), context, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (a AVVCSessionFactory) SessionForContextCompletion(context objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sessionForContext:completion:"), context, _block1)
}
func (a AVVCSessionFactory) SessionForContextError(context objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sessionForContext:error:"), context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (a AVVCSessionFactory) SessionManagerForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sessionManagerForContext:clientType:completion:"), context, type_, _block2)
}
func (a AVVCSessionFactory) SessionManagerForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sessionManagerForContext:clientType:error:"), context, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

var _avvcsessionfactory_setsessionwascreatedblock_p0_key byte

func (a AVVCSessionFactory) SetSessionWasCreatedBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setSessionWasCreatedBlock:"), _block0)
}

var _avvcsessionfactory_setsessionwillbedestroyedblock_p0_key byte

func (a AVVCSessionFactory) SetSessionWillBeDestroyedBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setSessionWillBeDestroyedBlock:"), _block0)
}

func (_AVVCSessionFactoryClass AVVCSessionFactoryClass) SharedInstance() AVVCSessionFactory {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVVCSessionFactoryClass.class), objc.Sel("sharedInstance"))
	return AVVCSessionFactoryFromID(rv)
}

func (a AVVCSessionFactory) PrimarySessionManager() IAVVCSessionManager {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("primarySessionManager"))
	return AVVCSessionManagerFromID(objc.ID(rv))
}
func (a AVVCSessionFactory) SetPrimarySessionManager(value IAVVCSessionManager) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setPrimarySessionManager:"), value)
}
func (a AVVCSessionFactory) SessionManagerMap() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sessionManagerMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AVVCSessionFactory) SetSessionManagerMap(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setSessionManagerMap:"), value)
}
func (a AVVCSessionFactory) WorkQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("workQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a AVVCSessionFactory) SetWorkQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setWorkQueue:"), value)
}

// SessionForContextClientTypeCompletionSync is a synchronous wrapper around [AVVCSessionFactory.SessionForContextClientTypeCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVCSessionFactory) SessionForContextClientTypeCompletionSync(ctx context.Context, context objectivec.IObject, type_ int64) error {
	done := make(chan struct{}, 1)
	a.SessionForContextClientTypeCompletion(context, type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SessionForContextCompletionSync is a synchronous wrapper around [AVVCSessionFactory.SessionForContextCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVCSessionFactory) SessionForContextCompletionSync(ctx context.Context, context objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.SessionForContextCompletion(context, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SessionManagerForContextClientTypeCompletionSync is a synchronous wrapper around [AVVCSessionFactory.SessionManagerForContextClientTypeCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVCSessionFactory) SessionManagerForContextClientTypeCompletionSync(ctx context.Context, context objectivec.IObject, type_ int64) error {
	done := make(chan struct{}, 1)
	a.SessionManagerForContextClientTypeCompletion(context, type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSessionWasCreatedBlockSync is a synchronous wrapper around [AVVCSessionFactory.SetSessionWasCreatedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVCSessionFactory) SetSessionWasCreatedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetSessionWasCreatedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSessionWillBeDestroyedBlockSync is a synchronous wrapper around [AVVCSessionFactory.SetSessionWillBeDestroyedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVVCSessionFactory) SetSessionWillBeDestroyedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.SetSessionWillBeDestroyedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
