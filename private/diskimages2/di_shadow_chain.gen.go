// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIShadowChain] class.
var (
	_DIShadowChainClass     DIShadowChainClass
	_DIShadowChainClassOnce sync.Once
)

func getDIShadowChainClass() DIShadowChainClass {
	_DIShadowChainClassOnce.Do(func() {
		_DIShadowChainClass = DIShadowChainClass{class: objc.GetClass("DIShadowChain")}
	})
	return _DIShadowChainClass
}

// GetDIShadowChainClass returns the class object for DIShadowChain.
func GetDIShadowChainClass() DIShadowChainClass {
	return getDIShadowChainClass()
}

type DIShadowChainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIShadowChainClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIShadowChainClass) Alloc() DIShadowChain {
	rv := objc.SendIfResponds[DIShadowChain](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIShadowChain.ActiveShadowURL]
//   - [DIShadowChain.AddShadowNodesError]
//   - [DIShadowChain.AddShadowNodesWrapReadOnlyError]
//   - [DIShadowChain.AddShadowURLsError]
//   - [DIShadowChain.EncodeWithCoder]
//   - [DIShadowChain.HasBaseImageCache]
//   - [DIShadowChain.IsEmpty]
//   - [DIShadowChain.MountPoints]
//   - [DIShadowChain.Nodes]
//   - [DIShadowChain.NonCacheNodes]
//   - [DIShadowChain.OpenWritableCreateNonExisting]
//   - [DIShadowChain.ShadowStats]
//   - [DIShadowChain.ShouldValidate]
//   - [DIShadowChain.SetShouldValidate]
//   - [DIShadowChain.StatWithError]
//   - [DIShadowChain.TopDiskImageNumBlocks]
//   - [DIShadowChain.VerifyNodesError]
//   - [DIShadowChain.InitWithCoder]
type DIShadowChain struct {
	objectivec.Object
}

// DIShadowChainFromID constructs a [DIShadowChain] from an objc.ID.
func DIShadowChainFromID(id objc.ID) DIShadowChain {
	return DIShadowChain{objectivec.Object{ID: id}}
}

// Ensure DIShadowChain implements IDIShadowChain.
var _ IDIShadowChain = DIShadowChain{}

// An interface definition for the [DIShadowChain] class.
//
// # Methods
//
//   - [IDIShadowChain.ActiveShadowURL]
//   - [IDIShadowChain.AddShadowNodesError]
//   - [IDIShadowChain.AddShadowNodesWrapReadOnlyError]
//   - [IDIShadowChain.AddShadowURLsError]
//   - [IDIShadowChain.EncodeWithCoder]
//   - [IDIShadowChain.HasBaseImageCache]
//   - [IDIShadowChain.IsEmpty]
//   - [IDIShadowChain.MountPoints]
//   - [IDIShadowChain.Nodes]
//   - [IDIShadowChain.NonCacheNodes]
//   - [IDIShadowChain.OpenWritableCreateNonExisting]
//   - [IDIShadowChain.ShadowStats]
//   - [IDIShadowChain.ShouldValidate]
//   - [IDIShadowChain.SetShouldValidate]
//   - [IDIShadowChain.StatWithError]
//   - [IDIShadowChain.TopDiskImageNumBlocks]
//   - [IDIShadowChain.VerifyNodesError]
//   - [IDIShadowChain.InitWithCoder]
type IDIShadowChain interface {
	objectivec.IObject

	// Topic: Methods

	ActiveShadowURL() foundation.NSURL
	AddShadowNodesError(nodes objectivec.IObject) (bool, error)
	AddShadowNodesWrapReadOnlyError(nodes objectivec.IObject, only bool) (bool, error)
	AddShadowURLsError(uRLs objectivec.IObject) (bool, error)
	EncodeWithCoder(coder foundation.INSCoder)
	HasBaseImageCache() bool
	IsEmpty() bool
	MountPoints() foundation.INSArray
	Nodes() foundation.INSArray
	NonCacheNodes() foundation.INSArray
	OpenWritableCreateNonExisting(writable bool, existing bool)
	ShadowStats() foundation.INSArray
	ShouldValidate() bool
	SetShouldValidate(value bool)
	StatWithError() (objectivec.IObject, error)
	TopDiskImageNumBlocks() int64
	VerifyNodesError(nodes objectivec.IObject) (bool, error)
	InitWithCoder(coder foundation.INSCoder) DIShadowChain
}

// Init initializes the instance.
func (d DIShadowChain) Init() DIShadowChain {
	rv := objc.SendIfResponds[DIShadowChain](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIShadowChain) Autorelease() DIShadowChain {
	rv := objc.SendIfResponds[DIShadowChain](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIShadowChain creates a new DIShadowChain instance.
func NewDIShadowChain() DIShadowChain {
	class := getDIShadowChainClass()
	rv := objc.SendIfResponds[DIShadowChain](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIShadowChainWithCoder(coder objectivec.IObject) DIShadowChain {
	instance := getDIShadowChainClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIShadowChainFromID(rv)
}

func (d DIShadowChain) AddShadowNodesError(nodes objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addShadowNodes:error:"), nodes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addShadowNodes:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIShadowChain) AddShadowNodesWrapReadOnlyError(nodes objectivec.IObject, only bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addShadowNodes:wrapReadOnly:error:"), nodes, only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addShadowNodes:wrapReadOnly:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIShadowChain) AddShadowURLsError(uRLs objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addShadowURLs:error:"), uRLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addShadowURLs:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIShadowChain) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (d DIShadowChain) OpenWritableCreateNonExisting(writable bool, existing bool) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("openWritable:createNonExisting:"), writable, existing)
}
func (d DIShadowChain) StatWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("statWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIShadowChain) VerifyNodesError(nodes objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("verifyNodes:error:"), nodes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyNodes:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIShadowChain) InitWithCoder(coder foundation.INSCoder) DIShadowChain {
	rv := objc.SendIfResponds[DIShadowChain](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_DIShadowChainClass DIShadowChainClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DIShadowChainClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (d DIShadowChain) ActiveShadowURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("activeShadowURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d DIShadowChain) HasBaseImageCache() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("hasBaseImageCache"))
	return rv
}
func (d DIShadowChain) IsEmpty() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("isEmpty"))
	return rv
}
func (d DIShadowChain) MountPoints() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("mountPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIShadowChain) Nodes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("nodes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIShadowChain) NonCacheNodes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("nonCacheNodes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIShadowChain) ShadowStats() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("shadowStats"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DIShadowChain) ShouldValidate() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("shouldValidate"))
	return rv
}
func (d DIShadowChain) SetShouldValidate(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setShouldValidate:"), value)
}
func (d DIShadowChain) TopDiskImageNumBlocks() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("topDiskImageNumBlocks"))
	return rv
}
