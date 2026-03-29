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

// The class instance for the [DiskImageGraph] class.
var (
	_DiskImageGraphClass     DiskImageGraphClass
	_DiskImageGraphClassOnce sync.Once
)

func getDiskImageGraphClass() DiskImageGraphClass {
	_DiskImageGraphClassOnce.Do(func() {
		_DiskImageGraphClass = DiskImageGraphClass{class: objc.GetClass("DiskImageGraph")}
	})
	return _DiskImageGraphClass
}

// GetDiskImageGraphClass returns the class object for DiskImageGraph.
func GetDiskImageGraphClass() DiskImageGraphClass {
	return getDiskImageGraphClass()
}

type DiskImageGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageGraphClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageGraphClass) Alloc() DiskImageGraph {
	rv := objc.Send[DiskImageGraph](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DiskImageGraph.URLRelativeToPstackParentWithURL]
//   - [DiskImageGraph.ActiveInfoWithExtraError]
//   - [DiskImageGraph.ActiveNode]
//   - [DiskImageGraph.SetActiveNode]
//   - [DiskImageGraph.AppendCacheWithURLTagError]
//   - [DiskImageGraph.AppendOverlayWithURLTagError]
//   - [DiskImageGraph.AppendOverlayWithURLTagNumBlocksError]
//   - [DiskImageGraph.AppendWithURLIsCacheTagNumBlocksSetNewActiveError]
//   - [DiskImageGraph.AppendWithURLIsCacheTagNumBlocksToNodeSetNewActiveError]
//   - [DiskImageGraph.AppendWithURLTagError]
//   - [DiskImageGraph.BaseNode]
//   - [DiskImageGraph.CheckStackValidityWithError]
//   - [DiskImageGraph.CloneToURLError]
//   - [DiskImageGraph.GetImageWithTagError]
//   - [DiskImageGraph.GetImageWithUUIDError]
//   - [DiskImageGraph.GraphDB]
//   - [DiskImageGraph.SetGraphDB]
//   - [DiskImageGraph.ImagesDictsArray]
//   - [DiskImageGraph.SetImagesDictsArray]
//   - [DiskImageGraph.InfoWithExtraError]
//   - [DiskImageGraph.Nodes]
//   - [DiskImageGraph.SetNodes]
//   - [DiskImageGraph.RemoveNodeWithTagRecursiveError]
//   - [DiskImageGraph.RemoveNodeWithUUIDRecursiveError]
//   - [DiskImageGraph.RemoveWithNodeRecursiveError]
//   - [DiskImageGraph.RootNode]
//   - [DiskImageGraph.SavePstackWithError]
//   - [DiskImageGraph.SavePstackWithURLError]
//   - [DiskImageGraph.SetActiveNodeWithTagError]
//   - [DiskImageGraph.SetActiveNodeWithUUIDError]
//   - [DiskImageGraph.ValidateAppendedImageWithURLParentNodeIsCacheError]
//   - [DiskImageGraph.InitWithBaseImageURLNewPstackURLTagError]
//   - [DiskImageGraph.InitWithBaseImageURLTagError]
//   - [DiskImageGraph.InitWithGraphDBError]
//   - [DiskImageGraph.InitWithGraphDBWorkDirError]
//   - [DiskImageGraph.InitWithPluginNamePluginParamsTagError]
//   - [DiskImageGraph.InitWithPstackURLError]
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph
type DiskImageGraph struct {
	objectivec.Object
}

// DiskImageGraphFromID constructs a [DiskImageGraph] from an objc.ID.
func DiskImageGraphFromID(id objc.ID) DiskImageGraph {
	return DiskImageGraph{objectivec.Object{ID: id}}
}
// Ensure DiskImageGraph implements IDiskImageGraph.
var _ IDiskImageGraph = DiskImageGraph{}

// An interface definition for the [DiskImageGraph] class.
//
// # Methods
//
//   - [IDiskImageGraph.URLRelativeToPstackParentWithURL]
//   - [IDiskImageGraph.ActiveInfoWithExtraError]
//   - [IDiskImageGraph.ActiveNode]
//   - [IDiskImageGraph.SetActiveNode]
//   - [IDiskImageGraph.AppendCacheWithURLTagError]
//   - [IDiskImageGraph.AppendOverlayWithURLTagError]
//   - [IDiskImageGraph.AppendOverlayWithURLTagNumBlocksError]
//   - [IDiskImageGraph.AppendWithURLIsCacheTagNumBlocksSetNewActiveError]
//   - [IDiskImageGraph.AppendWithURLIsCacheTagNumBlocksToNodeSetNewActiveError]
//   - [IDiskImageGraph.AppendWithURLTagError]
//   - [IDiskImageGraph.BaseNode]
//   - [IDiskImageGraph.CheckStackValidityWithError]
//   - [IDiskImageGraph.CloneToURLError]
//   - [IDiskImageGraph.GetImageWithTagError]
//   - [IDiskImageGraph.GetImageWithUUIDError]
//   - [IDiskImageGraph.GraphDB]
//   - [IDiskImageGraph.SetGraphDB]
//   - [IDiskImageGraph.ImagesDictsArray]
//   - [IDiskImageGraph.SetImagesDictsArray]
//   - [IDiskImageGraph.InfoWithExtraError]
//   - [IDiskImageGraph.Nodes]
//   - [IDiskImageGraph.SetNodes]
//   - [IDiskImageGraph.RemoveNodeWithTagRecursiveError]
//   - [IDiskImageGraph.RemoveNodeWithUUIDRecursiveError]
//   - [IDiskImageGraph.RemoveWithNodeRecursiveError]
//   - [IDiskImageGraph.RootNode]
//   - [IDiskImageGraph.SavePstackWithError]
//   - [IDiskImageGraph.SavePstackWithURLError]
//   - [IDiskImageGraph.SetActiveNodeWithTagError]
//   - [IDiskImageGraph.SetActiveNodeWithUUIDError]
//   - [IDiskImageGraph.ValidateAppendedImageWithURLParentNodeIsCacheError]
//   - [IDiskImageGraph.InitWithBaseImageURLNewPstackURLTagError]
//   - [IDiskImageGraph.InitWithBaseImageURLTagError]
//   - [IDiskImageGraph.InitWithGraphDBError]
//   - [IDiskImageGraph.InitWithGraphDBWorkDirError]
//   - [IDiskImageGraph.InitWithPluginNamePluginParamsTagError]
//   - [IDiskImageGraph.InitWithPstackURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph
type IDiskImageGraph interface {
	objectivec.IObject

	// Topic: Methods

	URLRelativeToPstackParentWithURL(url foundation.INSURL) objectivec.IObject
	ActiveInfoWithExtraError(extra bool) (objectivec.IObject, error)
	ActiveNode() IDiskImageGraphNode
	SetActiveNode(value IDiskImageGraphNode)
	AppendCacheWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error)
	AppendOverlayWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error)
	AppendOverlayWithURLTagNumBlocksError(url foundation.INSURL, tag objectivec.IObject, blocks uint64) (bool, error)
	AppendWithURLIsCacheTagNumBlocksSetNewActiveError(url foundation.INSURL, cache bool, tag objectivec.IObject, blocks uint64, active bool) (bool, error)
	AppendWithURLIsCacheTagNumBlocksToNodeSetNewActiveError(url foundation.INSURL, cache bool, tag objectivec.IObject, blocks uint64, node objectivec.IObject, active bool) (bool, error)
	AppendWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error)
	BaseNode() objectivec.IObject
	CheckStackValidityWithError() (bool, error)
	CloneToURLError(url foundation.INSURL) (objectivec.IObject, error)
	GetImageWithTagError(tag objectivec.IObject) (objectivec.IObject, error)
	GetImageWithUUIDError(uuid objectivec.IObject) (objectivec.IObject, error)
	GraphDB() foundation.INSDictionary
	SetGraphDB(value foundation.INSDictionary)
	ImagesDictsArray() foundation.INSArray
	SetImagesDictsArray(value foundation.INSArray)
	InfoWithExtraError(extra bool) (objectivec.IObject, error)
	Nodes() foundation.INSDictionary
	SetNodes(value foundation.INSDictionary)
	RemoveNodeWithTagRecursiveError(tag objectivec.IObject, recursive bool) (bool, error)
	RemoveNodeWithUUIDRecursiveError(uuid objectivec.IObject, recursive bool) (bool, error)
	RemoveWithNodeRecursiveError(node objectivec.IObject, recursive bool) (bool, error)
	RootNode() IDiskImageGraphNode
	SavePstackWithError() (bool, error)
	SavePstackWithURLError(url foundation.INSURL) (bool, error)
	SetActiveNodeWithTagError(tag objectivec.IObject) (bool, error)
	SetActiveNodeWithUUIDError(uuid objectivec.IObject) (bool, error)
	ValidateAppendedImageWithURLParentNodeIsCacheError(url foundation.INSURL, node objectivec.IObject, cache bool) (bool, error)
	InitWithBaseImageURLNewPstackURLTagError(url foundation.INSURL, url2 foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error)
	InitWithBaseImageURLTagError(url foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error)
	InitWithGraphDBError(db objectivec.IObject) (DiskImageGraph, error)
	InitWithGraphDBWorkDirError(db objectivec.IObject, dir objectivec.IObject) (DiskImageGraph, error)
	InitWithPluginNamePluginParamsTagError(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject) (DiskImageGraph, error)
	InitWithPstackURLError(url foundation.INSURL) (DiskImageGraph, error)
}

// Init initializes the instance.
func (d DiskImageGraph) Init() DiskImageGraph {
	rv := objc.Send[DiskImageGraph](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageGraph) Autorelease() DiskImageGraph {
	rv := objc.Send[DiskImageGraph](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageGraph creates a new DiskImageGraph instance.
func NewDiskImageGraph() DiskImageGraph {
	class := getDiskImageGraphClass()
	rv := objc.Send[DiskImageGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithBaseImageURL:newPstackURL:tag:error:
func NewDiskImageGraphWithBaseImageURLNewPstackURLTagError(url foundation.INSURL, url2 foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseImageURL:newPstackURL:tag:error:"), url, url2, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithBaseImageURL:tag:error:
func NewDiskImageGraphWithBaseImageURLTagError(url foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseImageURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithGraphDB:error:
func NewDiskImageGraphWithGraphDBError(db objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGraphDB:error:"), db, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithGraphDB:workDir:error:
func NewDiskImageGraphWithGraphDBWorkDirError(db objectivec.IObject, dir objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGraphDB:workDir:error:"), db, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithPluginName:pluginParams:tag:error:
func NewDiskImageGraphWithPluginNamePluginParamsTagError(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginParams:tag:error:"), name, params, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithPstackURL:error:
func NewDiskImageGraphWithPstackURLError(url foundation.INSURL) (DiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getDiskImageGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPstackURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/URLRelativeToPstackParentWithURL:
func (d DiskImageGraph) URLRelativeToPstackParentWithURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("URLRelativeToPstackParentWithURL:"), url)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/activeInfoWithExtra:error:
func (d DiskImageGraph) ActiveInfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("activeInfoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendCacheWithURL:tag:error:
func (d DiskImageGraph) AppendCacheWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendCacheWithURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendCacheWithURL:tag:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendOverlayWithURL:tag:error:
func (d DiskImageGraph) AppendOverlayWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendOverlayWithURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendOverlayWithURL:tag:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendOverlayWithURL:tag:numBlocks:error:
func (d DiskImageGraph) AppendOverlayWithURLTagNumBlocksError(url foundation.INSURL, tag objectivec.IObject, blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendOverlayWithURL:tag:numBlocks:error:"), url, tag, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendOverlayWithURL:tag:numBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendWithURL:isCache:tag:numBlocks:setNewActive:error:
func (d DiskImageGraph) AppendWithURLIsCacheTagNumBlocksSetNewActiveError(url foundation.INSURL, cache bool, tag objectivec.IObject, blocks uint64, active bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendWithURL:isCache:tag:numBlocks:setNewActive:error:"), url, cache, tag, blocks, active, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendWithURL:isCache:tag:numBlocks:setNewActive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendWithURL:isCache:tag:numBlocks:toNode:setNewActive:error:
func (d DiskImageGraph) AppendWithURLIsCacheTagNumBlocksToNodeSetNewActiveError(url foundation.INSURL, cache bool, tag objectivec.IObject, blocks uint64, node objectivec.IObject, active bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendWithURL:isCache:tag:numBlocks:toNode:setNewActive:error:"), url, cache, tag, blocks, node, active, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendWithURL:isCache:tag:numBlocks:toNode:setNewActive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/appendWithURL:tag:error:
func (d DiskImageGraph) AppendWithURLTagError(url foundation.INSURL, tag objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("appendWithURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("appendWithURL:tag:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/baseNode
func (d DiskImageGraph) BaseNode() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("baseNode"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/checkStackValidityWithError:
func (d DiskImageGraph) CheckStackValidityWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("checkStackValidityWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkStackValidityWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/cloneToURL:error:
func (d DiskImageGraph) CloneToURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cloneToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getImageWithTag:error:
func (d DiskImageGraph) GetImageWithTagError(tag objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getImageWithTag:error:"), tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getImageWithUUID:error:
func (d DiskImageGraph) GetImageWithUUIDError(uuid objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getImageWithUUID:error:"), uuid, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/infoWithExtra:error:
func (d DiskImageGraph) InfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("infoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/removeNodeWithTag:recursive:error:
func (d DiskImageGraph) RemoveNodeWithTagRecursiveError(tag objectivec.IObject, recursive bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("removeNodeWithTag:recursive:error:"), tag, recursive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeNodeWithTag:recursive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/removeNodeWithUUID:recursive:error:
func (d DiskImageGraph) RemoveNodeWithUUIDRecursiveError(uuid objectivec.IObject, recursive bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("removeNodeWithUUID:recursive:error:"), uuid, recursive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeNodeWithUUID:recursive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/removeWithNode:recursive:error:
func (d DiskImageGraph) RemoveWithNodeRecursiveError(node objectivec.IObject, recursive bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("removeWithNode:recursive:error:"), node, recursive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeWithNode:recursive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/savePstackWithError:
func (d DiskImageGraph) SavePstackWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("savePstackWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("savePstackWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/savePstackWithURL:error:
func (d DiskImageGraph) SavePstackWithURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("savePstackWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("savePstackWithURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/setActiveNodeWithTag:error:
func (d DiskImageGraph) SetActiveNodeWithTagError(tag objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setActiveNodeWithTag:error:"), tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setActiveNodeWithTag:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/setActiveNodeWithUUID:error:
func (d DiskImageGraph) SetActiveNodeWithUUIDError(uuid objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setActiveNodeWithUUID:error:"), uuid, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setActiveNodeWithUUID:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/validateAppendedImageWithURL:parentNode:isCache:error:
func (d DiskImageGraph) ValidateAppendedImageWithURLParentNodeIsCacheError(url foundation.INSURL, node objectivec.IObject, cache bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("validateAppendedImageWithURL:parentNode:isCache:error:"), url, node, cache, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAppendedImageWithURL:parentNode:isCache:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithBaseImageURL:newPstackURL:tag:error:
func (d DiskImageGraph) InitWithBaseImageURLNewPstackURLTagError(url foundation.INSURL, url2 foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithBaseImageURL:newPstackURL:tag:error:"), url, url2, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithBaseImageURL:tag:error:
func (d DiskImageGraph) InitWithBaseImageURLTagError(url foundation.INSURL, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithBaseImageURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithGraphDB:error:
func (d DiskImageGraph) InitWithGraphDBError(db objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithGraphDB:error:"), db, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithGraphDB:workDir:error:
func (d DiskImageGraph) InitWithGraphDBWorkDirError(db objectivec.IObject, dir objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithGraphDB:workDir:error:"), db, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithPluginName:pluginParams:tag:error:
func (d DiskImageGraph) InitWithPluginNamePluginParamsTagError(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithPluginName:pluginParams:tag:error:"), name, params, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/initWithPstackURL:error:
func (d DiskImageGraph) InitWithPstackURLError(url foundation.INSURL) (DiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithPstackURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageGraphFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/copyDictNodesToFolder:dict:error:
func (_DiskImageGraphClass DiskImageGraphClass) CopyDictNodesToFolderDictError(folder objectivec.IObject, dict objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("copyDictNodesToFolder:dict:error:"), folder, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyDictNodesToFolder:dict:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/createGraphDictWithNode:
func (_DiskImageGraphClass DiskImageGraphClass) CreateGraphDictWithNode(node objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphClass.class), objc.Sel("createGraphDictWithNode:"), node)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/createNodesConnectivityWithNodesDict:error:
func (_DiskImageGraphClass DiskImageGraphClass) CreateNodesConnectivityWithNodesDictError(dict objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("createNodesConnectivityWithNodesDict:error:"), dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createNodesConnectivityWithNodesDict:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/failWithNoPstackError:
func (_DiskImageGraphClass DiskImageGraphClass) FailWithNoPstackError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("failWithNoPstackError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("failWithNoPstackError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getFirstNonCacheAncestorWithNode:error:
func (_DiskImageGraphClass DiskImageGraphClass) GetFirstNonCacheAncestorWithNodeError(node objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphClass.class), objc.Sel("getFirstNonCacheAncestorWithNode:error:"), node, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getImageInfoDictWithURL:error:
func (_DiskImageGraphClass DiskImageGraphClass) GetImageInfoDictWithURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphClass.class), objc.Sel("getImageInfoDictWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getImageUUIDStrWithIdentityInfo:stackableUUIDFallback:error:
func (_DiskImageGraphClass DiskImageGraphClass) GetImageUUIDStrWithIdentityInfoStackableUUIDFallbackError(info objectivec.IObject, uUIDFallback bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphClass.class), objc.Sel("getImageUUIDStrWithIdentityInfo:stackableUUIDFallback:error:"), info, uUIDFallback, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/getImageUUIDWithURL:allowMissingUUID:error:
func (_DiskImageGraphClass DiskImageGraphClass) GetImageUUIDWithURLAllowMissingUUIDError(url foundation.INSURL, uuid bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageGraphClass.class), objc.Sel("getImageUUIDWithURL:allowMissingUUID:error:"), url, uuid, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/loadPlistDictFromFileHandle:dict:error:
func (_DiskImageGraphClass DiskImageGraphClass) LoadPlistDictFromFileHandleDictError(handle objectivec.IObject, dict []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("loadPlistDictFromFileHandle:dict:error:"), handle, objectivec.IObjectSliceToNSArray(dict), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadPlistDictFromFileHandle:dict:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/populateNodesDictsWithArray:workDir:nodesDict:error:
func (_DiskImageGraphClass DiskImageGraphClass) PopulateNodesDictsWithArrayWorkDirNodesDictError(array objectivec.IObject, dir objectivec.IObject, dict objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("populateNodesDictsWithArray:workDir:nodesDict:error:"), array, dir, dict, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("populateNodesDictsWithArray:workDir:nodesDict:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/saveToPlistWithDictionary:URL:error:
func (_DiskImageGraphClass DiskImageGraphClass) SaveToPlistWithDictionaryURLError(dictionary objectivec.IObject, rl objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("saveToPlistWithDictionary:URL:error:"), dictionary, rl, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveToPlistWithDictionary:URL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/validateWithDictionary:error:
func (_DiskImageGraphClass DiskImageGraphClass) ValidateWithDictionaryError(dictionary objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageGraphClass.class), objc.Sel("validateWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithDictionary:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/activeNode
func (d DiskImageGraph) ActiveNode() IDiskImageGraphNode {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("activeNode"))
	return DiskImageGraphNodeFromID(objc.ID(rv))
}
func (d DiskImageGraph) SetActiveNode(value IDiskImageGraphNode) {
	objc.Send[struct{}](d.ID, objc.Sel("setActiveNode:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/graphDB
func (d DiskImageGraph) GraphDB() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("graphDB"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DiskImageGraph) SetGraphDB(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setGraphDB:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/imagesDictsArray
func (d DiskImageGraph) ImagesDictsArray() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("imagesDictsArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DiskImageGraph) SetImagesDictsArray(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setImagesDictsArray:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/nodes
func (d DiskImageGraph) Nodes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("nodes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DiskImageGraph) SetNodes(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setNodes:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageGraph/rootNode
func (d DiskImageGraph) RootNode() IDiskImageGraphNode {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("rootNode"))
	return DiskImageGraphNodeFromID(objc.ID(rv))
}

