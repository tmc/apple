// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLModelSecurityServiceToClient] class.
var (
	_CoreMLModelSecurityServiceToClientClass     CoreMLModelSecurityServiceToClientClass
	_CoreMLModelSecurityServiceToClientClassOnce sync.Once
)

func getCoreMLModelSecurityServiceToClientClass() CoreMLModelSecurityServiceToClientClass {
	_CoreMLModelSecurityServiceToClientClassOnce.Do(func() {
		_CoreMLModelSecurityServiceToClientClass = CoreMLModelSecurityServiceToClientClass{class: objc.GetClass("CoreMLModelSecurityServiceToClient")}
	})
	return _CoreMLModelSecurityServiceToClientClass
}

// GetCoreMLModelSecurityServiceToClientClass returns the class object for CoreMLModelSecurityServiceToClient.
func GetCoreMLModelSecurityServiceToClientClass() CoreMLModelSecurityServiceToClientClass {
	return getCoreMLModelSecurityServiceToClientClass()
}

type CoreMLModelSecurityServiceToClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLModelSecurityServiceToClientClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLModelSecurityServiceToClientClass) Alloc() CoreMLModelSecurityServiceToClient {
	rv := objc.Send[CoreMLModelSecurityServiceToClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureNamesWithReply]
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
//   - [CoreMLModelSecurityServiceToClient.FeatureProviderCount]
//   - [CoreMLModelSecurityServiceToClient.SetFeatureProviderCount]
//   - [CoreMLModelSecurityServiceToClient.FeatureProviderMap]
//   - [CoreMLModelSecurityServiceToClient.SetFeatureProviderMap]
//   - [CoreMLModelSecurityServiceToClient.ServiceToClientQueue]
//   - [CoreMLModelSecurityServiceToClient.SetServiceToClientQueue]
//
// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient
type CoreMLModelSecurityServiceToClient struct {
	objectivec.Object
}

// CoreMLModelSecurityServiceToClientFromID constructs a [CoreMLModelSecurityServiceToClient] from an objc.ID.
func CoreMLModelSecurityServiceToClientFromID(id objc.ID) CoreMLModelSecurityServiceToClient {
	return CoreMLModelSecurityServiceToClient{objectivec.Object{ID: id}}
}

// Ensure CoreMLModelSecurityServiceToClient implements ICoreMLModelSecurityServiceToClient.
var _ ICoreMLModelSecurityServiceToClient = CoreMLModelSecurityServiceToClient{}

// An interface definition for the [CoreMLModelSecurityServiceToClient] class.
//
// # Methods
//
//   - [ICoreMLModelSecurityServiceToClient.ClientFeatureNamesWithReply]
//   - [ICoreMLModelSecurityServiceToClient.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
//   - [ICoreMLModelSecurityServiceToClient.FeatureProviderCount]
//   - [ICoreMLModelSecurityServiceToClient.SetFeatureProviderCount]
//   - [ICoreMLModelSecurityServiceToClient.FeatureProviderMap]
//   - [ICoreMLModelSecurityServiceToClient.SetFeatureProviderMap]
//   - [ICoreMLModelSecurityServiceToClient.ServiceToClientQueue]
//   - [ICoreMLModelSecurityServiceToClient.SetServiceToClientQueue]
//
// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient
type ICoreMLModelSecurityServiceToClient interface {
	objectivec.IObject

	// Topic: Methods

	ClientFeatureNamesWithReply(names objectivec.IObject, reply VoidHandler)
	ClientFeatureValueForNameUniqueKeyForProviderWithReply(name objectivec.IObject, provider objectivec.IObject, reply VoidHandler)
	FeatureProviderCount() foundation.INSCountedSet
	SetFeatureProviderCount(value foundation.INSCountedSet)
	FeatureProviderMap() foundation.INSDictionary
	SetFeatureProviderMap(value foundation.INSDictionary)
	ServiceToClientQueue() objectivec.Object
	SetServiceToClientQueue(value objectivec.Object)
}

// Init initializes the instance.
func (c CoreMLModelSecurityServiceToClient) Init() CoreMLModelSecurityServiceToClient {
	rv := objc.Send[CoreMLModelSecurityServiceToClient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLModelSecurityServiceToClient) Autorelease() CoreMLModelSecurityServiceToClient {
	rv := objc.Send[CoreMLModelSecurityServiceToClient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLModelSecurityServiceToClient creates a new CoreMLModelSecurityServiceToClient instance.
func NewCoreMLModelSecurityServiceToClient() CoreMLModelSecurityServiceToClient {
	class := getCoreMLModelSecurityServiceToClientClass()
	rv := objc.Send[CoreMLModelSecurityServiceToClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient/clientFeatureNames:withReply:
func (c CoreMLModelSecurityServiceToClient) ClientFeatureNamesWithReply(names objectivec.IObject, reply VoidHandler) {
	_block1, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("clientFeatureNames:withReply:"), names, _block1)
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient/clientFeatureValueForName:uniqueKeyForProvider:withReply:
func (c CoreMLModelSecurityServiceToClient) ClientFeatureValueForNameUniqueKeyForProviderWithReply(name objectivec.IObject, provider objectivec.IObject, reply VoidHandler) {
	_block2, _ := NewVoidBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("clientFeatureValueForName:uniqueKeyForProvider:withReply:"), name, provider, _block2)
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient/featureProviderCount
func (c CoreMLModelSecurityServiceToClient) FeatureProviderCount() foundation.INSCountedSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureProviderCount"))
	return foundation.NSCountedSetFromID(objc.ID(rv))
}
func (c CoreMLModelSecurityServiceToClient) SetFeatureProviderCount(value foundation.INSCountedSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setFeatureProviderCount:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient/featureProviderMap
func (c CoreMLModelSecurityServiceToClient) FeatureProviderMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureProviderMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c CoreMLModelSecurityServiceToClient) SetFeatureProviderMap(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setFeatureProviderMap:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClient/serviceToClientQueue
func (c CoreMLModelSecurityServiceToClient) ServiceToClientQueue() objectivec.Object {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("serviceToClientQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (c CoreMLModelSecurityServiceToClient) SetServiceToClientQueue(value objectivec.Object) {
	objc.Send[struct{}](c.ID, objc.Sel("setServiceToClientQueue:"), value)
}

// ClientFeatureNamesWithReplySync is a synchronous wrapper around [CoreMLModelSecurityServiceToClient.ClientFeatureNamesWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c CoreMLModelSecurityServiceToClient) ClientFeatureNamesWithReplySync(ctx context.Context, names objectivec.IObject) error {
	done := make(chan struct{}, 1)
	c.ClientFeatureNamesWithReply(names, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ClientFeatureValueForNameUniqueKeyForProviderWithReplySync is a synchronous wrapper around [CoreMLModelSecurityServiceToClient.ClientFeatureValueForNameUniqueKeyForProviderWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c CoreMLModelSecurityServiceToClient) ClientFeatureValueForNameUniqueKeyForProviderWithReplySync(ctx context.Context, name objectivec.IObject, provider objectivec.IObject) error {
	done := make(chan struct{}, 1)
	c.ClientFeatureValueForNameUniqueKeyForProviderWithReply(name, provider, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
