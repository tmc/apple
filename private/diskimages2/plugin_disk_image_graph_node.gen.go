// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PluginDiskImageGraphNode] class.
var (
	_PluginDiskImageGraphNodeClass     PluginDiskImageGraphNodeClass
	_PluginDiskImageGraphNodeClassOnce sync.Once
)

func getPluginDiskImageGraphNodeClass() PluginDiskImageGraphNodeClass {
	_PluginDiskImageGraphNodeClassOnce.Do(func() {
		_PluginDiskImageGraphNodeClass = PluginDiskImageGraphNodeClass{class: objc.GetClass("PluginDiskImageGraphNode")}
	})
	return _PluginDiskImageGraphNodeClass
}

// GetPluginDiskImageGraphNodeClass returns the class object for PluginDiskImageGraphNode.
func GetPluginDiskImageGraphNodeClass() PluginDiskImageGraphNodeClass {
	return getPluginDiskImageGraphNodeClass()
}

type PluginDiskImageGraphNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PluginDiskImageGraphNodeClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PluginDiskImageGraphNodeClass) Alloc() PluginDiskImageGraphNode {
	rv := objc.Send[PluginDiskImageGraphNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PluginDiskImageGraphNode.PluginName]
//   - [PluginDiskImageGraphNode.PluginParams]
//   - [PluginDiskImageGraphNode.InitWithPluginNamePluginParamsTagUUIDParentNodeMetadataIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode
type PluginDiskImageGraphNode struct {
	DiskImageGraphNode
}

// PluginDiskImageGraphNodeFromID constructs a [PluginDiskImageGraphNode] from an objc.ID.
func PluginDiskImageGraphNodeFromID(id objc.ID) PluginDiskImageGraphNode {
	return PluginDiskImageGraphNode{DiskImageGraphNode: DiskImageGraphNodeFromID(id)}
}

// Ensure PluginDiskImageGraphNode implements IPluginDiskImageGraphNode.
var _ IPluginDiskImageGraphNode = PluginDiskImageGraphNode{}

// An interface definition for the [PluginDiskImageGraphNode] class.
//
// # Methods
//
//   - [IPluginDiskImageGraphNode.PluginName]
//   - [IPluginDiskImageGraphNode.PluginParams]
//   - [IPluginDiskImageGraphNode.InitWithPluginNamePluginParamsTagUUIDParentNodeMetadataIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode
type IPluginDiskImageGraphNode interface {
	IDiskImageGraphNode

	// Topic: Methods

	PluginName() string
	PluginParams() foundation.INSDictionary
	InitWithPluginNamePluginParamsTagUUIDParentNodeMetadataIsCache(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) PluginDiskImageGraphNode
}

// Init initializes the instance.
func (p PluginDiskImageGraphNode) Init() PluginDiskImageGraphNode {
	rv := objc.Send[PluginDiskImageGraphNode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PluginDiskImageGraphNode) Autorelease() PluginDiskImageGraphNode {
	rv := objc.Send[PluginDiskImageGraphNode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPluginDiskImageGraphNode creates a new PluginDiskImageGraphNode instance.
func NewPluginDiskImageGraphNode() PluginDiskImageGraphNode {
	class := getPluginDiskImageGraphNodeClass()
	rv := objc.Send[PluginDiskImageGraphNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode/initWithDictionary:workDir:error:
func NewPluginDiskImageGraphNodeWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (PluginDiskImageGraphNode, error) {
	var errorPtr objc.ID
	instance := getPluginDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:workDir:error:"), dictionary, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return PluginDiskImageGraphNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return PluginDiskImageGraphNodeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode/initWithPluginName:pluginParams:tag:UUID:parentNode:metadata:isCache:
func NewPluginDiskImageGraphNodeWithPluginNamePluginParamsTagUUIDParentNodeMetadataIsCache(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) PluginDiskImageGraphNode {
	instance := getPluginDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginParams:tag:UUID:parentNode:metadata:isCache:"), name, params, tag, uid, node, metadata, cache)
	return PluginDiskImageGraphNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithTag:UUID:parentNode:metadata:isCache:
func NewPluginDiskImageGraphNodeWithTagUUIDParentNodeMetadataIsCache(tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) PluginDiskImageGraphNode {
	instance := getPluginDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:UUID:parentNode:metadata:isCache:"), tag, uid, node, metadata, cache)
	return PluginDiskImageGraphNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode/initWithPluginName:pluginParams:tag:UUID:parentNode:metadata:isCache:
func (p PluginDiskImageGraphNode) InitWithPluginNamePluginParamsTagUUIDParentNodeMetadataIsCache(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) PluginDiskImageGraphNode {
	rv := objc.Send[PluginDiskImageGraphNode](p.ID, objc.Sel("initWithPluginName:pluginParams:tag:UUID:parentNode:metadata:isCache:"), name, params, tag, uid, node, metadata, cache)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode/pluginName
func (p PluginDiskImageGraphNode) PluginName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pluginName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/PluginDiskImageGraphNode/pluginParams
func (p PluginDiskImageGraphNode) PluginParams() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pluginParams"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
