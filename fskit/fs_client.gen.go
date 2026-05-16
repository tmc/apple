// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSClient] class.
var (
	_FSClientClass     FSClientClass
	_FSClientClassOnce sync.Once
)

func getFSClientClass() FSClientClass {
	_FSClientClassOnce.Do(func() {
		_FSClientClass = FSClientClass{class: objc.GetClass("FSClient")}
	})
	return _FSClientClass
}

// GetFSClientClass returns the class object for FSClient.
func GetFSClientClass() FSClientClass {
	return getFSClientClass()
}

type FSClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSClientClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSClientClass) Alloc() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An interface for apps and daemons to interact with FSKit.
//
// # Overview
//
// FSClient is the primary management interface for FSKit. Use this class to
// discover FSKit extensions installed on the system, including your own.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient
type FSClient struct {
	objectivec.Object
}

// FSClientFromID constructs a [FSClient] from an objc.ID.
//
// An interface for apps and daemons to interact with FSKit.
func FSClientFromID(id objc.ID) FSClient {
	return FSClient{objectivec.Object{ID: id}}
}

// NOTE: FSClient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSClient] class.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient
type IFSClient interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c FSClient) Init() FSClient {
	rv := objc.Send[FSClient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c FSClient) Autorelease() FSClient {
	rv := objc.Send[FSClient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSClient creates a new FSClient instance.
func NewFSClient() FSClient {
	class := getFSClientClass()
	rv := objc.Send[FSClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The shared instance of the FSKit client class.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient/shared
func (_FSClientClass FSClientClass) SharedInstance() FSClient {
	rv := objc.Send[objc.ID](objc.ID(_FSClientClass.class), objc.Sel("sharedInstance"))
	return FSClientFromID(objc.ID(rv))
}
