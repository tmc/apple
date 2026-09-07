// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// BKIOHIDServiceMatcherDataProviding protocol.
type BKIOHIDServiceMatcherDataProviding interface {
	objectivec.IObject

	// IOHIDServicesMatching protocol.
	IOHIDServicesMatching(matching objectivec.IObject) objectivec.IObject

	// HoistOnThread protocol.
	HoistOnThread(thread VoidHandler)

	// HoistOnThreadForSessionID protocol.
	HoistOnThreadForSessionID(thread VoidHandler, id uint32)

	// RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon protocol.
	RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer)

	// UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon protocol.
	UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer)
}

// BKIOHIDServiceMatcherDataProvidingObject wraps an existing Objective-C object that conforms to the BKIOHIDServiceMatcherDataProviding protocol.
type BKIOHIDServiceMatcherDataProvidingObject struct {
	objectivec.Object
}

func (o BKIOHIDServiceMatcherDataProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKIOHIDServiceMatcherDataProvidingObjectFromID constructs a [BKIOHIDServiceMatcherDataProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKIOHIDServiceMatcherDataProvidingObjectFromID(id objc.ID) BKIOHIDServiceMatcherDataProvidingObject {
	return BKIOHIDServiceMatcherDataProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKIOHIDServiceMatcherDataProvidingObject) IOHIDServicesMatching(matching objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("IOHIDServicesMatching:"), matching)
	return objectivec.Object{ID: rv}
}
func (o BKIOHIDServiceMatcherDataProvidingObject) HoistOnThread(thread VoidHandler) {
	_block0, _cleanup0 := NewVoidBlock(thread)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("hoistOnThread:"), objc.ID(_block0))
}
func (o BKIOHIDServiceMatcherDataProvidingObject) HoistOnThreadForSessionID(thread VoidHandler, id uint32) {
	_block0, _cleanup0 := NewVoidBlock(thread)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("hoistOnThread:forSessionID:"), objc.ID(_block0), id)
}
func (o BKIOHIDServiceMatcherDataProvidingObject) RegisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer) {
	_block0, _cleanup0 := NewVoidBlock(callback)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("registerIOHIDServicesCallback:matchingDictionary:target:refCon:"), objc.ID(_block0), dictionary, target, con)
}
func (o BKIOHIDServiceMatcherDataProvidingObject) UnregisterIOHIDServicesCallbackMatchingDictionaryTargetRefCon(callback VoidHandler, dictionary objectivec.IObject, target unsafe.Pointer, con unsafe.Pointer) {
	_block0, _cleanup0 := NewVoidBlock(callback)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("unregisterIOHIDServicesCallback:matchingDictionary:target:refCon:"), objc.ID(_block0), dictionary, target, con)
}
