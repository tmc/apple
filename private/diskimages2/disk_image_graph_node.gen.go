// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageGraphNode] class.
var (
	_DiskImageGraphNodeClass     DiskImageGraphNodeClass
	_DiskImageGraphNodeClassOnce sync.Once
)

func getDiskImageGraphNodeClass() DiskImageGraphNodeClass {
	_DiskImageGraphNodeClassOnce.Do(func() {
		_DiskImageGraphNodeClass = DiskImageGraphNodeClass{class: objc.GetClass("DiskImageGraphNode")}
	})
	return _DiskImageGraphNodeClass
}

// GetDiskImageGraphNodeClass returns the class object for DiskImageGraphNode.
func GetDiskImageGraphNodeClass() DiskImageGraphNodeClass {
	return getDiskImageGraphNodeClass()
}

type DiskImageGraphNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageGraphNodeClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageGraphNodeClass) Alloc() DiskImageGraphNode {
	rv := objc.Send[DiskImageGraphNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DiskImageGraphNode.URL]
//   - [DiskImageGraphNode.UUID]
//   - [DiskImageGraphNode.SetUUID]
//   - [DiskImageGraphNode.AddDecendantsToArray]
//   - [DiskImageGraphNode.Children]
//   - [DiskImageGraphNode.ChildrenInfoWithExtraError]
//   - [DiskImageGraphNode.DeleteImage]
//   - [DiskImageGraphNode.GetChildren]
//   - [DiskImageGraphNode.GetDescendants]
//   - [DiskImageGraphNode.InfoWithExtraError]
//   - [DiskImageGraphNode.IsCache]
//   - [DiskImageGraphNode.Metadata]
//   - [DiskImageGraphNode.SetMetadata]
//   - [DiskImageGraphNode.MutableChildren]
//   - [DiskImageGraphNode.SetMutableChildren]
//   - [DiskImageGraphNode.Parent]
//   - [DiskImageGraphNode.SetParent]
//   - [DiskImageGraphNode.ParentUUID]
//   - [DiskImageGraphNode.PstackDict]
//   - [DiskImageGraphNode.SetPstackDict]
//   - [DiskImageGraphNode.RecursiveInfoWithExtraError]
//   - [DiskImageGraphNode.Tag]
//   - [DiskImageGraphNode.SetTag]
//   - [DiskImageGraphNode.ToDIShadowNode]
//   - [DiskImageGraphNode.ToDictionary]
//   - [DiskImageGraphNode.ValidateAppendedImageWithInfoError]
//   - [DiskImageGraphNode.InitWithDictionaryWorkDirError]
//   - [DiskImageGraphNode.InitWithTagUUIDParentNodeMetadataIsCache]
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode
type DiskImageGraphNode struct {
	objectivec.Object
}

// DiskImageGraphNodeFromID constructs a [DiskImageGraphNode] from an objc.ID.
func DiskImageGraphNodeFromID(id objc.ID) DiskImageGraphNode {
	return DiskImageGraphNode{objectivec.Object{ID: id}}
}
// Ensure DiskImageGraphNode implements IDiskImageGraphNode.
var _ IDiskImageGraphNode = DiskImageGraphNode{}

// An interface definition for the [DiskImageGraphNode] class.
//
// # Methods
//
//   - [IDiskImageGraphNode.URL]
//   - [IDiskImageGraphNode.UUID]
//   - [IDiskImageGraphNode.SetUUID]
//   - [IDiskImageGraphNode.AddDecendantsToArray]
//   - [IDiskImageGraphNode.Children]
//   - [IDiskImageGraphNode.ChildrenInfoWithExtraError]
//   - [IDiskImageGraphNode.DeleteImage]
//   - [IDiskImageGraphNode.GetChildren]
//   - [IDiskImageGraphNode.GetDescendants]
//   - [IDiskImageGraphNode.InfoWithExtraError]
//   - [IDiskImageGraphNode.IsCache]
//   - [IDiskImageGraphNode.Metadata]
//   - [IDiskImageGraphNode.SetMetadata]
//   - [IDiskImageGraphNode.MutableChildren]
//   - [IDiskImageGraphNode.SetMutableChildren]
//   - [IDiskImageGraphNode.Parent]
//   - [IDiskImageGraphNode.SetParent]
//   - [IDiskImageGraphNode.ParentUUID]
//   - [IDiskImageGraphNode.PstackDict]
//   - [IDiskImageGraphNode.SetPstackDict]
//   - [IDiskImageGraphNode.RecursiveInfoWithExtraError]
//   - [IDiskImageGraphNode.Tag]
//   - [IDiskImageGraphNode.SetTag]
//   - [IDiskImageGraphNode.ToDIShadowNode]
//   - [IDiskImageGraphNode.ToDictionary]
//   - [IDiskImageGraphNode.ValidateAppendedImageWithInfoError]
//   - [IDiskImageGraphNode.InitWithDictionaryWorkDirError]
//   - [IDiskImageGraphNode.InitWithTagUUIDParentNodeMetadataIsCache]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode
type IDiskImageGraphNode interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.INSURL
	UUID() foundation.NSUUID
	SetUUID(value foundation.NSUUID)
	AddDecendantsToArray(array objectivec.IObject)
	Children() foundation.INSArray
	ChildrenInfoWithExtraError(extra bool) (objectivec.IObject, error)
	DeleteImage() bool
	GetChildren() objectivec.IObject
	GetDescendants() objectivec.IObject
	InfoWithExtraError(extra bool) (objectivec.IObject, error)
	IsCache() bool
	Metadata() foundation.INSDictionary
	SetMetadata(value foundation.INSDictionary)
	MutableChildren() foundation.INSArray
	SetMutableChildren(value foundation.INSArray)
	Parent() IDiskImageGraphNode
	SetParent(value IDiskImageGraphNode)
	ParentUUID() foundation.NSUUID
	PstackDict() foundation.INSDictionary
	SetPstackDict(value foundation.INSDictionary)
	RecursiveInfoWithExtraError(extra bool) (objectivec.IObject, error)
	Tag() string
	SetTag(value string)
	ToDIShadowNode() objectivec.IObject
	ToDictionary() objectivec.IObject
	ValidateAppendedImageWithInfoError(info objectivec.IObject) (bool, error)
	InitWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (DiskImageGraphNode, error)
	InitWithTagUUIDParentNodeMetadataIsCache(tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) DiskImageGraphNode
}

// Init initializes the instance.
func (d DiskImageGraphNode) Init() DiskImageGraphNode {
	rv := objc.Send[DiskImageGraphNode](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageGraphNode) Autorelease() DiskImageGraphNode {
	rv := objc.Send[DiskImageGraphNode](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageGraphNode creates a new DiskImageGraphNode instance.
func NewDiskImageGraphNode() DiskImageGraphNode {
	class := getDiskImageGraphNodeClass()
	rv := objc.Send[DiskImageGraphNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithDictionary:workDir:error:
func NewDiskImageGraphNodeWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (DiskImageGraphNode, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:workDir:error:"), dictionary, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraphNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphNodeFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithTag:UUID:parentNode:metadata:isCache:
func NewDiskImageGraphNodeWithTagUUIDParentNodeMetadataIsCache(tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) DiskImageGraphNode {
	instance := getDiskImageGraphNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:UUID:parentNode:metadata:isCache:"), tag, uid, node, metadata, cache)
	return DiskImageGraphNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/addDecendantsToArray:
func (d DiskImageGraphNode) AddDecendantsToArray(array objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("addDecendantsToArray:"), array)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/childrenInfoWithExtra:error:
func (d DiskImageGraphNode) ChildrenInfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("childrenInfoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/deleteImage
func (d DiskImageGraphNode) DeleteImage() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("deleteImage"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/getChildren
func (d DiskImageGraphNode) GetChildren() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getChildren"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/getDescendants
func (d DiskImageGraphNode) GetDescendants() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getDescendants"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/infoWithExtra:error:
func (d DiskImageGraphNode) InfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("infoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/recursiveInfoWithExtra:error:
func (d DiskImageGraphNode) RecursiveInfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("recursiveInfoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/setIsCache:
func (d DiskImageGraphNode) SetIsCache(cache bool) {
	objc.Send[objc.ID](d.ID, objc.Sel("setIsCache:"), cache)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/toDIShadowNode
func (d DiskImageGraphNode) ToDIShadowNode() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("toDIShadowNode"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/toDictionary
func (d DiskImageGraphNode) ToDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("toDictionary"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/validateAppendedImageWithInfo:error:
func (d DiskImageGraphNode) ValidateAppendedImageWithInfoError(info objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("validateAppendedImageWithInfo:error:"), info, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAppendedImageWithInfo:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithDictionary:workDir:error:
func (d DiskImageGraphNode) InitWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (DiskImageGraphNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithDictionary:workDir:error:"), dictionary, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraphNode{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphNodeFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/initWithTag:UUID:parentNode:metadata:isCache:
func (d DiskImageGraphNode) InitWithTagUUIDParentNodeMetadataIsCache(tag objectivec.IObject, uid objectivec.IObject, node objectivec.IObject, metadata objectivec.IObject, cache bool) DiskImageGraphNode {
	rv := objc.Send[DiskImageGraphNode](d.ID, objc.Sel("initWithTag:UUID:parentNode:metadata:isCache:"), tag, uid, node, metadata, cache)
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/GraphNodeWithDictionary:workDir:error:
func (_DiskImageGraphNodeClass DiskImageGraphNodeClass) GraphNodeWithDictionaryWorkDirError(dictionary objectivec.IObject, dir objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphNodeClass.class), objc.Sel("GraphNodeWithDictionary:workDir:error:"), dictionary, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/validateWithDictionary:error:
func (_DiskImageGraphNodeClass DiskImageGraphNodeClass) ValidateWithDictionaryError(dictionary objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphNodeClass.class), objc.Sel("validateWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithDictionary:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/URL
func (d DiskImageGraphNode) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/UUID
func (d DiskImageGraphNode) UUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (d DiskImageGraphNode) SetUUID(value foundation.NSUUID) {
	objc.Send[struct{}](d.ID, objc.Sel("setUUID:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/children
func (d DiskImageGraphNode) Children() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("children"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/isCache
func (d DiskImageGraphNode) IsCache() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isCache"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/metadata
func (d DiskImageGraphNode) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DiskImageGraphNode) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setMetadata:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/mutableChildren
func (d DiskImageGraphNode) MutableChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mutableChildren"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DiskImageGraphNode) SetMutableChildren(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setMutableChildren:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/parent
func (d DiskImageGraphNode) Parent() IDiskImageGraphNode {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("parent"))
	return DiskImageGraphNodeFromID(objc.ID(rv))
}
func (d DiskImageGraphNode) SetParent(value IDiskImageGraphNode) {
	objc.Send[struct{}](d.ID, objc.Sel("setParent:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/parentUUID
func (d DiskImageGraphNode) ParentUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("parentUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/pstackDict
func (d DiskImageGraphNode) PstackDict() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("pstackDict"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DiskImageGraphNode) SetPstackDict(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setPstackDict:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraphNode/tag
func (d DiskImageGraphNode) Tag() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("tag"))
	return foundation.NSStringFromID(rv).String()
}
func (d DiskImageGraphNode) SetTag(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setTag:"), objc.String(value))
}

