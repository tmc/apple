// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SerializedDiskImageGraph] class.
var (
	_SerializedDiskImageGraphClass     SerializedDiskImageGraphClass
	_SerializedDiskImageGraphClassOnce sync.Once
)

func getSerializedDiskImageGraphClass() SerializedDiskImageGraphClass {
	_SerializedDiskImageGraphClassOnce.Do(func() {
		_SerializedDiskImageGraphClass = SerializedDiskImageGraphClass{class: objc.GetClass("SerializedDiskImageGraph")}
	})
	return _SerializedDiskImageGraphClass
}

// GetSerializedDiskImageGraphClass returns the class object for SerializedDiskImageGraph.
func GetSerializedDiskImageGraphClass() SerializedDiskImageGraphClass {
	return getSerializedDiskImageGraphClass()
}

type SerializedDiskImageGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SerializedDiskImageGraphClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SerializedDiskImageGraphClass) Alloc() SerializedDiskImageGraph {
	rv := objc.SendIfResponds[SerializedDiskImageGraph](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SerializedDiskImageGraph.PstackURL]
//   - [SerializedDiskImageGraph.InitWithBaseImageURLPstackURLTagError]
//   - [SerializedDiskImageGraph.InitWithGraphDBPstackURLError]
//   - [SerializedDiskImageGraph.InitWithPluginNamePluginParamsPstackURLTagError]
type SerializedDiskImageGraph struct {
	DiskImageGraph
}

// SerializedDiskImageGraphFromID constructs a [SerializedDiskImageGraph] from an objc.ID.
func SerializedDiskImageGraphFromID(id objc.ID) SerializedDiskImageGraph {
	return SerializedDiskImageGraph{DiskImageGraph: DiskImageGraphFromID(id)}
}

// Ensure SerializedDiskImageGraph implements ISerializedDiskImageGraph.
var _ ISerializedDiskImageGraph = SerializedDiskImageGraph{}

// An interface definition for the [SerializedDiskImageGraph] class.
//
// # Methods
//
//   - [ISerializedDiskImageGraph.PstackURL]
//   - [ISerializedDiskImageGraph.InitWithBaseImageURLPstackURLTagError]
//   - [ISerializedDiskImageGraph.InitWithGraphDBPstackURLError]
//   - [ISerializedDiskImageGraph.InitWithPluginNamePluginParamsPstackURLTagError]
type ISerializedDiskImageGraph interface {
	IDiskImageGraph

	// Topic: Methods

	PstackURL() foundation.NSURL
	InitWithBaseImageURLPstackURLTagError(url foundation.NSURL, url2 foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error)
	InitWithGraphDBPstackURLError(db objectivec.IObject, url foundation.NSURL) (SerializedDiskImageGraph, error)
	InitWithPluginNamePluginParamsPstackURLTagError(name objectivec.IObject, params objectivec.IObject, url foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error)
}

// Init initializes the instance.
func (s SerializedDiskImageGraph) Init() SerializedDiskImageGraph {
	rv := objc.SendIfResponds[SerializedDiskImageGraph](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SerializedDiskImageGraph) Autorelease() SerializedDiskImageGraph {
	rv := objc.SendIfResponds[SerializedDiskImageGraph](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSerializedDiskImageGraph creates a new SerializedDiskImageGraph instance.
func NewSerializedDiskImageGraph() SerializedDiskImageGraph {
	class := getSerializedDiskImageGraphClass()
	rv := objc.SendIfResponds[SerializedDiskImageGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSerializedDiskImageGraphWithBaseImageURLNewPstackURLTagError(url foundation.NSURL, url2 foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBaseImageURL:newPstackURL:tag:error:"), url, url2, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithBaseImageURLPstackURLTagError(url foundation.NSURL, url2 foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBaseImageURL:pstackURL:tag:error:"), url, url2, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithBaseImageURLTagError(url foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBaseImageURL:tag:error:"), url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithGraphDBError(db objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGraphDB:error:"), db, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithGraphDBPstackURLError(db objectivec.IObject, url foundation.NSURL) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGraphDB:pstackURL:error:"), db, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithGraphDBWorkDirError(db objectivec.IObject, dir objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGraphDB:workDir:error:"), db, dir, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithPluginNamePluginParamsPstackURLTagError(name objectivec.IObject, params objectivec.IObject, url foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginParams:pstackURL:tag:error:"), name, params, url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithPluginNamePluginParamsTagError(name objectivec.IObject, params objectivec.IObject, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPluginName:pluginParams:tag:error:"), name, params, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func NewSerializedDiskImageGraphWithPstackURLError(url foundation.NSURL) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	instance := getSerializedDiskImageGraphClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPstackURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SerializedDiskImageGraph{}, objc.ErrInitFailed
	}
	return SerializedDiskImageGraphFromID(rv), nil
}

func (s SerializedDiskImageGraph) InitWithBaseImageURLPstackURLTagError(url foundation.NSURL, url2 foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithBaseImageURL:pstackURL:tag:error:"), url, url2, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return SerializedDiskImageGraphFromID(rv), nil

}
func (s SerializedDiskImageGraph) InitWithGraphDBPstackURLError(db objectivec.IObject, url foundation.NSURL) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithGraphDB:pstackURL:error:"), db, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return SerializedDiskImageGraphFromID(rv), nil

}
func (s SerializedDiskImageGraph) InitWithPluginNamePluginParamsPstackURLTagError(name objectivec.IObject, params objectivec.IObject, url foundation.NSURL, tag objectivec.IObject) (SerializedDiskImageGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithPluginName:pluginParams:pstackURL:tag:error:"), name, params, url, tag, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SerializedDiskImageGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return SerializedDiskImageGraphFromID(rv), nil

}

func (_SerializedDiskImageGraphClass SerializedDiskImageGraphClass) GetRelativeIfContainedWithChildURLParentURL(url foundation.NSURL, url2 foundation.NSURL) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SerializedDiskImageGraphClass.class), objc.Sel("getRelativeIfContainedWithChildURL:parentURL:"), url, url2)
	return objectivec.Object{ID: rv}
}

func (s SerializedDiskImageGraph) PstackURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("pstackURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
