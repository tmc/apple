// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NativeDiskImageGraphNode] class.
var (
	_NativeDiskImageGraphNodeClass     NativeDiskImageGraphNodeClass
	_NativeDiskImageGraphNodeClassOnce sync.Once
)

func getNativeDiskImageGraphNodeClass() NativeDiskImageGraphNodeClass {
	_NativeDiskImageGraphNodeClassOnce.Do(func() {
		_NativeDiskImageGraphNodeClass = NativeDiskImageGraphNodeClass{class: objc.GetClass("NativeDiskImageGraphNode")}
	})
	return _NativeDiskImageGraphNodeClass
}

// GetNativeDiskImageGraphNodeClass returns the class object for NativeDiskImageGraphNode.
func GetNativeDiskImageGraphNodeClass() NativeDiskImageGraphNodeClass {
	return getNativeDiskImageGraphNodeClass()
}

type NativeDiskImageGraphNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NativeDiskImageGraphNodeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NativeDiskImageGraphNodeClass) Alloc() NativeDiskImageGraphNode {
	rv := objc.Send[NativeDiskImageGraphNode](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [NativeDiskImageGraphNode.FilePath]
//   - [NativeDiskImageGraphNode.InitWithURLTagUUIDParentNodeMetadataIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode
type NativeDiskImageGraphNode struct {
	DiskImageGraphNode
}

// NativeDiskImageGraphNodeFromID constructs a [NativeDiskImageGraphNode] from an objc.ID.
func NativeDiskImageGraphNodeFromID(id objc.ID) NativeDiskImageGraphNode {
	return NativeDiskImageGraphNode{DiskImageGraphNode: DiskImageGraphNodeFromID(id)}
}

// Ensure NativeDiskImageGraphNode implements INativeDiskImageGraphNode.
var _ INativeDiskImageGraphNode = NativeDiskImageGraphNode{}

// An interface definition for the [NativeDiskImageGraphNode] class.
//
// # Methods
//
//   - [INativeDiskImageGraphNode.FilePath]
//   - [INativeDiskImageGraphNode.InitWithURLTagUUIDParentNodeMetadataIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode
type INativeDiskImageGraphNode interface {
	IDiskImageGraphNode

	// Topic: Methods

	FilePath() foundation.INSURL
	InitWithURLTagUUIDParentNodeMetadataIsCache(url foundation.INSURL, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) NativeDiskImageGraphNode
}

// Init initializes the instance.
func (n NativeDiskImageGraphNode) Init() NativeDiskImageGraphNode {
	rv := objc.Send[NativeDiskImageGraphNode](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NativeDiskImageGraphNode) Autorelease() NativeDiskImageGraphNode {
	rv := objc.Send[NativeDiskImageGraphNode](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNativeDiskImageGraphNode creates a new NativeDiskImageGraphNode instance.
func NewNativeDiskImageGraphNode() NativeDiskImageGraphNode {
	class := getNativeDiskImageGraphNodeClass()
	rv := objc.Send[NativeDiskImageGraphNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode/initWithDictionary:workDir:error:
func NewNativeDiskImageGraphNodeWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (NativeDiskImageGraphNode, error) {
	var errorPtr objc.ID
	instance := getNativeDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:workDir:error:"), dictionary, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NativeDiskImageGraphNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return NativeDiskImageGraphNodeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithTag:UUID:parentNode:metadata:isCache:
func NewNativeDiskImageGraphNodeWithTagUUIDParentNodeMetadataIsCache(tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) NativeDiskImageGraphNode {
	instance := getNativeDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:UUID:parentNode:metadata:isCache:"), tag, uid, node, metadata, cache)
	return NativeDiskImageGraphNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode/initWithURL:tag:UUID:parentNode:metadata:isCache:
func NewNativeDiskImageGraphNodeWithURLTagUUIDParentNodeMetadataIsCache(url foundation.INSURL, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) NativeDiskImageGraphNode {
	instance := getNativeDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:tag:UUID:parentNode:metadata:isCache:"), url, tag, uid, node, metadata, cache)
	return NativeDiskImageGraphNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode/initWithURL:tag:UUID:parentNode:metadata:isCache:
func (n NativeDiskImageGraphNode) InitWithURLTagUUIDParentNodeMetadataIsCache(url foundation.INSURL, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) NativeDiskImageGraphNode {
	rv := objc.Send[NativeDiskImageGraphNode](n.ID, objc.Sel("initWithURL:tag:UUID:parentNode:metadata:isCache:"), url, tag, uid, node, metadata, cache)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/NativeDiskImageGraphNode/filePath
func (n NativeDiskImageGraphNode) FilePath() foundation.INSURL {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("filePath"))
	return foundation.NSURLFromID(objc.ID(rv))
}
