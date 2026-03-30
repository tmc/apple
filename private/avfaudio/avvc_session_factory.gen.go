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
	rv := objc.Send[AVVCSessionFactory](objc.ID(ac.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory
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
func (v AVVCSessionFactory) Init() AVVCSessionFactory {
	rv := objc.Send[AVVCSessionFactory](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCSessionFactory) Autorelease() AVVCSessionFactory {
	rv := objc.Send[AVVCSessionFactory](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCSessionFactory creates a new AVVCSessionFactory instance.
func NewAVVCSessionFactory() AVVCSessionFactory {
	class := getAVVCSessionFactoryClass()
	rv := objc.Send[AVVCSessionFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/_wqCreatePrimarySessionManagerIfNeeded:clientType:error:
func (v AVVCSessionFactory) _wqCreatePrimarySessionManagerIfNeededClientTypeError(needed objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_wqCreatePrimarySessionManagerIfNeeded:clientType:error:"), needed, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// WqCreatePrimarySessionManagerIfNeededClientTypeError is an exported wrapper for the private method _wqCreatePrimarySessionManagerIfNeededClientTypeError.
func (v AVVCSessionFactory) WqCreatePrimarySessionManagerIfNeededClientTypeError(needed objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	return v._wqCreatePrimarySessionManagerIfNeededClientTypeError(needed, type_)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/_wqSessionAndManagerForContext:clientType:session:manager:error:
func (v AVVCSessionFactory) _wqSessionAndManagerForContextClientTypeSessionManagerError(context objectivec.IObject, type_ int64, session []objectivec.IObject, manager []objectivec.IObject) error {
	var errorPtr objc.ID
	objc.Send[struct{}](v.ID, objc.Sel("_wqSessionAndManagerForContext:clientType:session:manager:error:"), context, type_, objectivec.IObjectSliceToNSArray(session), objectivec.IObjectSliceToNSArray(manager), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// WqSessionAndManagerForContextClientTypeSessionManagerError is an exported wrapper for the private method _wqSessionAndManagerForContextClientTypeSessionManagerError.
func (v AVVCSessionFactory) WqSessionAndManagerForContextClientTypeSessionManagerError(context objectivec.IObject, type_ int64, session []objectivec.IObject, manager []objectivec.IObject) error {
	return v._wqSessionAndManagerForContextClientTypeSessionManagerError(context, type_, session, manager)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/auxSessionManagers
func (v AVVCSessionFactory) AuxSessionManagers() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("auxSessionManagers"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/cleanupContext:
func (v AVVCSessionFactory) CleanupContext(context objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("cleanupContext:"), context)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/releasePrimarySessionManager
func (v AVVCSessionFactory) ReleasePrimarySessionManager() {
	objc.Send[objc.ID](v.ID, objc.Sel("releasePrimarySessionManager"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionForContext:clientType:completion:
func (v AVVCSessionFactory) SessionForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("sessionForContext:clientType:completion:"), context, type_, _block2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionForContext:clientType:error:
func (v AVVCSessionFactory) SessionForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sessionForContext:clientType:error:"), context, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionForContext:completion:
func (v AVVCSessionFactory) SessionForContextCompletion(context objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("sessionForContext:completion:"), context, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionForContext:error:
func (v AVVCSessionFactory) SessionForContextError(context objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sessionForContext:error:"), context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionManagerForContext:clientType:completion:
func (v AVVCSessionFactory) SessionManagerForContextClientTypeCompletion(context objectivec.IObject, type_ int64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("sessionManagerForContext:clientType:completion:"), context, type_, _block2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionManagerForContext:clientType:error:
func (v AVVCSessionFactory) SessionManagerForContextClientTypeError(context objectivec.IObject, type_ int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sessionManagerForContext:clientType:error:"), context, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/setSessionWasCreatedBlock:
func (v AVVCSessionFactory) SetSessionWasCreatedBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setSessionWasCreatedBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/setSessionWillBeDestroyedBlock:
func (v AVVCSessionFactory) SetSessionWillBeDestroyedBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setSessionWillBeDestroyedBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sharedInstance
func (_AVVCSessionFactoryClass AVVCSessionFactoryClass) SharedInstance() AVVCSessionFactory {
	rv := objc.Send[objc.ID](objc.ID(_AVVCSessionFactoryClass.class), objc.Sel("sharedInstance"))
	return AVVCSessionFactoryFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/primarySessionManager
func (v AVVCSessionFactory) PrimarySessionManager() IAVVCSessionManager {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("primarySessionManager"))
	return AVVCSessionManagerFromID(objc.ID(rv))
}
func (v AVVCSessionFactory) SetPrimarySessionManager(value IAVVCSessionManager) {
	objc.Send[struct{}](v.ID, objc.Sel("setPrimarySessionManager:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/sessionManagerMap
func (v AVVCSessionFactory) SessionManagerMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sessionManagerMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v AVVCSessionFactory) SetSessionManagerMap(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setSessionManagerMap:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionFactory/workQueue
func (v AVVCSessionFactory) WorkQueue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("workQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v AVVCSessionFactory) SetWorkQueue(value objectivec.Object) {
	objc.Send[struct{}](v.ID, objc.Sel("setWorkQueue:"), value)
}

// SessionForContextClientTypeCompletionSync is a synchronous wrapper around [AVVCSessionFactory.SessionForContextClientTypeCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVCSessionFactory) SessionForContextClientTypeCompletionSync(ctx context.Context, context objectivec.IObject, type_ int64) error {
	done := make(chan struct{}, 1)
	v.SessionForContextClientTypeCompletion(context, type_, func() {
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
func (v AVVCSessionFactory) SessionForContextCompletionSync(ctx context.Context, context objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.SessionForContextCompletion(context, func() {
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
func (v AVVCSessionFactory) SessionManagerForContextClientTypeCompletionSync(ctx context.Context, context objectivec.IObject, type_ int64) error {
	done := make(chan struct{}, 1)
	v.SessionManagerForContextClientTypeCompletion(context, type_, func() {
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
func (v AVVCSessionFactory) SetSessionWasCreatedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetSessionWasCreatedBlock(func() {
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
func (v AVVCSessionFactory) SetSessionWillBeDestroyedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetSessionWillBeDestroyedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
