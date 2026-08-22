// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIStatsParams] class.
var (
	_DIStatsParamsClass     DIStatsParamsClass
	_DIStatsParamsClassOnce sync.Once
)

func getDIStatsParamsClass() DIStatsParamsClass {
	_DIStatsParamsClassOnce.Do(func() {
		_DIStatsParamsClass = DIStatsParamsClass{class: objc.GetClass("DIStatsParams")}
	})
	return _DIStatsParamsClass
}

// GetDIStatsParamsClass returns the class object for DIStatsParams.
func GetDIStatsParamsClass() DIStatsParamsClass {
	return getDIStatsParamsClass()
}

type DIStatsParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIStatsParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIStatsParamsClass) Alloc() DIStatsParams {
	rv := objc.SendIfResponds[DIStatsParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIStatsParams.StatInstanceID]
//   - [DIStatsParams.StatsWithError]
//   - [DIStatsParams.InitWithURLInstanceIdError]
type DIStatsParams struct {
	DIBaseParams
}

// DIStatsParamsFromID constructs a [DIStatsParams] from an objc.ID.
func DIStatsParamsFromID(id objc.ID) DIStatsParams {
	return DIStatsParams{DIBaseParams: DIBaseParamsFromID(id)}
}

// Ensure DIStatsParams implements IDIStatsParams.
var _ IDIStatsParams = DIStatsParams{}

// An interface definition for the [DIStatsParams] class.
//
// # Methods
//
//   - [IDIStatsParams.StatInstanceID]
//   - [IDIStatsParams.StatsWithError]
//   - [IDIStatsParams.InitWithURLInstanceIdError]
type IDIStatsParams interface {
	IDIBaseParams

	// Topic: Methods

	StatInstanceID() foundation.NSUUID
	StatsWithError() (objectivec.IObject, error)
	InitWithURLInstanceIdError(url foundation.NSURL, id objectivec.IObject) (DIStatsParams, error)
}

// Init initializes the instance.
func (d DIStatsParams) Init() DIStatsParams {
	rv := objc.SendIfResponds[DIStatsParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIStatsParams) Autorelease() DIStatsParams {
	rv := objc.SendIfResponds[DIStatsParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIStatsParams creates a new DIStatsParams instance.
func NewDIStatsParams() DIStatsParams {
	class := getDIStatsParamsClass()
	rv := objc.SendIfResponds[DIStatsParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIStatsParamsWithCoder(coder objectivec.IObject) DIStatsParams {
	instance := getDIStatsParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIStatsParamsFromID(rv)
}

func NewDIStatsParamsWithURLError(url foundation.NSURL) (DIStatsParams, error) {
	var errorPtr objc.ID
	instance := getDIStatsParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStatsParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIStatsParams{}, objc.ErrInitFailed
	}
	return DIStatsParamsFromID(rv), nil
}

func NewDIStatsParamsWithURLInstanceIdError(url foundation.NSURL, id objectivec.IObject) (DIStatsParams, error) {
	var errorPtr objc.ID
	instance := getDIStatsParamsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:instanceId:error:"), url, id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStatsParams{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIStatsParams{}, objc.ErrInitFailed
	}
	return DIStatsParamsFromID(rv), nil
}

func (d DIStatsParams) StatsWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("statsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIStatsParams) InitWithURLInstanceIdError(url foundation.NSURL, id objectivec.IObject) (DIStatsParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:instanceId:error:"), url, id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStatsParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIStatsParamsFromID(rv), nil

}

func (d DIStatsParams) StatInstanceID() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("statInstanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
