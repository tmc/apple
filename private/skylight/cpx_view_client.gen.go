// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXViewClient] class.
var (
	_CPXViewClientClass     CPXViewClientClass
	_CPXViewClientClassOnce sync.Once
)

func getCPXViewClientClass() CPXViewClientClass {
	_CPXViewClientClassOnce.Do(func() {
		_CPXViewClientClass = CPXViewClientClass{class: objc.GetClass("_CPXViewClient")}
	})
	return _CPXViewClientClass
}

// GetCPXViewClientClass returns the class object for _CPXViewClient.
func GetCPXViewClientClass() CPXViewClientClass {
	return getCPXViewClientClass()
}

type CPXViewClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXViewClientClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXViewClientClass) Alloc() CPXViewClient {
	rv := objc.Send[CPXViewClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXViewClient
type CPXViewClient struct {
	objectivec.Object
}

// CPXViewClientFromID constructs a [CPXViewClient] from an objc.ID.
func CPXViewClientFromID(id objc.ID) CPXViewClient {
	return CPXViewClient{objectivec.Object{ID: id}}
}

// Ensure CPXViewClient implements ICPXViewClient.
var _ ICPXViewClient = CPXViewClient{}

// An interface definition for the [CPXViewClient] class.
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXViewClient
type ICPXViewClient interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CPXViewClient) Init() CPXViewClient {
	rv := objc.Send[CPXViewClient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXViewClient) Autorelease() CPXViewClient {
	rv := objc.Send[CPXViewClient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXViewClient creates a new CPXViewClient instance.
func NewCPXViewClient() CPXViewClient {
	class := getCPXViewClientClass()
	rv := objc.Send[CPXViewClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}
