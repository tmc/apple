// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PBRequest] class.
var (
	_PBRequestClass     PBRequestClass
	_PBRequestClassOnce sync.Once
)

func getPBRequestClass() PBRequestClass {
	_PBRequestClassOnce.Do(func() {
		_PBRequestClass = PBRequestClass{class: objc.GetClass("PBRequest")}
	})
	return _PBRequestClass
}

// GetPBRequestClass returns the class object for PBRequest.
func GetPBRequestClass() PBRequestClass {
	return getPBRequestClass()
}

type PBRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PBRequestClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PBRequestClass) Alloc() PBRequest {
	rv := objc.Send[PBRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other CoreML classes. [Full Topic]
type PBRequest struct {
	objectivec.Object
}

// PBRequestFromID constructs a [PBRequest] from an objc.ID.
//
// A parent class referenced by other CoreML classes.
func PBRequestFromID(id objc.ID) PBRequest {
	return PBRequest{objectivec.Object{ID: id}}
}
// Ensure PBRequest implements IPBRequest.
var _ IPBRequest = PBRequest{}

// An interface definition for the [PBRequest] class.
type IPBRequest interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p PBRequest) Init() PBRequest {
	rv := objc.Send[PBRequest](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PBRequest) Autorelease() PBRequest {
	rv := objc.Send[PBRequest](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPBRequest creates a new PBRequest instance.
func NewPBRequest() PBRequest {
	class := getPBRequestClass()
	rv := objc.Send[PBRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

