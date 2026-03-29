// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [DIAPFSPartition] class.
var (
	_DIAPFSPartitionClass     DIAPFSPartitionClass
	_DIAPFSPartitionClassOnce sync.Once
)

func getDIAPFSPartitionClass() DIAPFSPartitionClass {
	_DIAPFSPartitionClassOnce.Do(func() {
		_DIAPFSPartitionClass = DIAPFSPartitionClass{class: objc.GetClass("DIAPFSPartition")}
	})
	return _DIAPFSPartitionClass
}

// GetDIAPFSPartitionClass returns the class object for DIAPFSPartition.
func GetDIAPFSPartitionClass() DIAPFSPartitionClass {
	return getDIAPFSPartitionClass()
}

type DIAPFSPartitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIAPFSPartitionClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIAPFSPartitionClass) Alloc() DIAPFSPartition {
	rv := objc.Send[DIAPFSPartition](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIAPFSPartition
type DIAPFSPartition struct {
	DIDataPartition
}

// DIAPFSPartitionFromID constructs a [DIAPFSPartition] from an objc.ID.
func DIAPFSPartitionFromID(id objc.ID) DIAPFSPartition {
	return DIAPFSPartition{DIDataPartition: DIDataPartitionFromID(id)}
}
// Ensure DIAPFSPartition implements IDIAPFSPartition.
var _ IDIAPFSPartition = DIAPFSPartition{}

// An interface definition for the [DIAPFSPartition] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAPFSPartition
type IDIAPFSPartition interface {
	IDIDataPartition
}

// Init initializes the instance.
func (d DIAPFSPartition) Init() DIAPFSPartition {
	rv := objc.Send[DIAPFSPartition](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIAPFSPartition) Autorelease() DIAPFSPartition {
	rv := objc.Send[DIAPFSPartition](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIAPFSPartition creates a new DIAPFSPartition instance.
func NewDIAPFSPartition() DIAPFSPartition {
	class := getDIAPFSPartitionClass()
	rv := objc.Send[DIAPFSPartition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

