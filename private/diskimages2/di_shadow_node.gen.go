// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIShadowNode] class.
var (
	_DIShadowNodeClass     DIShadowNodeClass
	_DIShadowNodeClassOnce sync.Once
)

func getDIShadowNodeClass() DIShadowNodeClass {
	_DIShadowNodeClassOnce.Do(func() {
		_DIShadowNodeClass = DIShadowNodeClass{class: objc.GetClass("DIShadowNode")}
	})
	return _DIShadowNodeClass
}

// GetDIShadowNodeClass returns the class object for DIShadowNode.
func GetDIShadowNodeClass() DIShadowNodeClass {
	return getDIShadowNodeClass()
}

type DIShadowNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIShadowNodeClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIShadowNodeClass) Alloc() DIShadowNode {
	rv := objc.Send[DIShadowNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIShadowNode.URL]
//   - [DIShadowNode.CreateBackendWithFlags]
//   - [DIShadowNode.EncodeWithCoder]
//   - [DIShadowNode.FileBackend]
//   - [DIShadowNode.SetFileBackend]
//   - [DIShadowNode.IsCache]
//   - [DIShadowNode.NumBlocks]
//   - [DIShadowNode.SetNumBlocks]
//   - [DIShadowNode.InitWithCoder]
//   - [DIShadowNode.InitWithURLIsCache]
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode
type DIShadowNode struct {
	objectivec.Object
}

// DIShadowNodeFromID constructs a [DIShadowNode] from an objc.ID.
func DIShadowNodeFromID(id objc.ID) DIShadowNode {
	return DIShadowNode{objectivec.Object{ID: id}}
}
// Ensure DIShadowNode implements IDIShadowNode.
var _ IDIShadowNode = DIShadowNode{}

// An interface definition for the [DIShadowNode] class.
//
// # Methods
//
//   - [IDIShadowNode.URL]
//   - [IDIShadowNode.CreateBackendWithFlags]
//   - [IDIShadowNode.EncodeWithCoder]
//   - [IDIShadowNode.FileBackend]
//   - [IDIShadowNode.SetFileBackend]
//   - [IDIShadowNode.IsCache]
//   - [IDIShadowNode.NumBlocks]
//   - [IDIShadowNode.SetNumBlocks]
//   - [IDIShadowNode.InitWithCoder]
//   - [IDIShadowNode.InitWithURLIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode
type IDIShadowNode interface {
	objectivec.IObject

	// Topic: Methods

	URL() IDIURL
	CreateBackendWithFlags(flags int)
	EncodeWithCoder(coder foundation.INSCoder)
	FileBackend() IFileLocalXPC
	SetFileBackend(value IFileLocalXPC)
	IsCache() bool
	NumBlocks() uint64
	SetNumBlocks(value uint64)
	InitWithCoder(coder foundation.INSCoder) DIShadowNode
	InitWithURLIsCache(url foundation.INSURL, cache bool) DIShadowNode
}

// Init initializes the instance.
func (d DIShadowNode) Init() DIShadowNode {
	rv := objc.Send[DIShadowNode](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIShadowNode) Autorelease() DIShadowNode {
	rv := objc.Send[DIShadowNode](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIShadowNode creates a new DIShadowNode instance.
func NewDIShadowNode() DIShadowNode {
	class := getDIShadowNodeClass()
	rv := objc.Send[DIShadowNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/initWithCoder:
func NewDIShadowNodeWithCoder(coder objectivec.IObject) DIShadowNode {
	instance := getDIShadowNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIShadowNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/initWithURL:isCache:
func NewDIShadowNodeWithURLIsCache(url foundation.INSURL, cache bool) DIShadowNode {
	instance := getDIShadowNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:isCache:"), url, cache)
	return DIShadowNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/createBackendWithFlags:
func (d DIShadowNode) CreateBackendWithFlags(flags int) {
	objc.Send[objc.ID](d.ID, objc.Sel("createBackendWithFlags:"), flags)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/encodeWithCoder:
func (d DIShadowNode) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/initWithCoder:
func (d DIShadowNode) InitWithCoder(coder foundation.INSCoder) DIShadowNode {
	rv := objc.Send[DIShadowNode](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/initWithURL:isCache:
func (d DIShadowNode) InitWithURLIsCache(url foundation.INSURL, cache bool) DIShadowNode {
	rv := objc.Send[DIShadowNode](d.ID, objc.Sel("initWithURL:isCache:"), url, cache)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/supportsSecureCoding
func (_DIShadowNodeClass DIShadowNodeClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIShadowNodeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/URL
func (d DIShadowNode) URL() IDIURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("URL"))
	return DIURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/fileBackend
func (d DIShadowNode) FileBackend() IFileLocalXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileBackend"))
	return FileLocalXPCFromID(objc.ID(rv))
}
func (d DIShadowNode) SetFileBackend(value IFileLocalXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setFileBackend:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/isCache
func (d DIShadowNode) IsCache() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isCache"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIShadowNode/numBlocks
func (d DIShadowNode) NumBlocks() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("numBlocks"))
	return rv
}
func (d DIShadowNode) SetNumBlocks(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setNumBlocks:"), value)
}

