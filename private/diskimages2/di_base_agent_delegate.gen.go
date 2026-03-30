// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [DIBaseAgentDelegate] class.
var (
	_DIBaseAgentDelegateClass     DIBaseAgentDelegateClass
	_DIBaseAgentDelegateClassOnce sync.Once
)

func getDIBaseAgentDelegateClass() DIBaseAgentDelegateClass {
	_DIBaseAgentDelegateClassOnce.Do(func() {
		_DIBaseAgentDelegateClass = DIBaseAgentDelegateClass{class: objc.GetClass("DIBaseAgentDelegate")}
	})
	return _DIBaseAgentDelegateClass
}

// GetDIBaseAgentDelegateClass returns the class object for DIBaseAgentDelegate.
func GetDIBaseAgentDelegateClass() DIBaseAgentDelegateClass {
	return getDIBaseAgentDelegateClass()
}

type DIBaseAgentDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBaseAgentDelegateClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseAgentDelegateClass) Alloc() DIBaseAgentDelegate {
	rv := objc.Send[DIBaseAgentDelegate](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseAgentDelegate
type DIBaseAgentDelegate struct {
	DIBaseServiceDelegate
}

// DIBaseAgentDelegateFromID constructs a [DIBaseAgentDelegate] from an objc.ID.
func DIBaseAgentDelegateFromID(id objc.ID) DIBaseAgentDelegate {
	return DIBaseAgentDelegate{DIBaseServiceDelegate: DIBaseServiceDelegateFromID(id)}
}

// Ensure DIBaseAgentDelegate implements IDIBaseAgentDelegate.
var _ IDIBaseAgentDelegate = DIBaseAgentDelegate{}

// An interface definition for the [DIBaseAgentDelegate] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseAgentDelegate
type IDIBaseAgentDelegate interface {
	IDIBaseServiceDelegate
}

// Init initializes the instance.
func (d DIBaseAgentDelegate) Init() DIBaseAgentDelegate {
	rv := objc.Send[DIBaseAgentDelegate](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseAgentDelegate) Autorelease() DIBaseAgentDelegate {
	rv := objc.Send[DIBaseAgentDelegate](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseAgentDelegate creates a new DIBaseAgentDelegate instance.
func NewDIBaseAgentDelegate() DIBaseAgentDelegate {
	class := getDIBaseAgentDelegateClass()
	rv := objc.Send[DIBaseAgentDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}
