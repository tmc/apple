// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerDebugDump] class.
var (
	_GTShaderProfilerDebugDumpClass     GTShaderProfilerDebugDumpClass
	_GTShaderProfilerDebugDumpClassOnce sync.Once
)

func getGTShaderProfilerDebugDumpClass() GTShaderProfilerDebugDumpClass {
	_GTShaderProfilerDebugDumpClassOnce.Do(func() {
		_GTShaderProfilerDebugDumpClass = GTShaderProfilerDebugDumpClass{class: objc.GetClass("GTShaderProfilerDebugDump")}
	})
	return _GTShaderProfilerDebugDumpClass
}

// GetGTShaderProfilerDebugDumpClass returns the class object for GTShaderProfilerDebugDump.
func GetGTShaderProfilerDebugDumpClass() GTShaderProfilerDebugDumpClass {
	return getGTShaderProfilerDebugDumpClass()
}

type GTShaderProfilerDebugDumpClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerDebugDumpClass) Alloc() GTShaderProfilerDebugDump {
	rv := objc.Send[GTShaderProfilerDebugDump](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerDebugDump.FilePathFromFileName]
//   - [GTShaderProfilerDebugDump.SetDirectory]
//   - [GTShaderProfilerDebugDump.InitWithDirectory]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump
type GTShaderProfilerDebugDump struct {
	objectivec.Object
}

// GTShaderProfilerDebugDumpFromID constructs a [GTShaderProfilerDebugDump] from an objc.ID.
func GTShaderProfilerDebugDumpFromID(id objc.ID) GTShaderProfilerDebugDump {
	return GTShaderProfilerDebugDump{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerDebugDump implements IGTShaderProfilerDebugDump.
var _ IGTShaderProfilerDebugDump = GTShaderProfilerDebugDump{}

// An interface definition for the [GTShaderProfilerDebugDump] class.
//
// # Methods
//
//   - [IGTShaderProfilerDebugDump.FilePathFromFileName]
//   - [IGTShaderProfilerDebugDump.SetDirectory]
//   - [IGTShaderProfilerDebugDump.InitWithDirectory]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump
type IGTShaderProfilerDebugDump interface {
	objectivec.IObject

	// Topic: Methods

	FilePathFromFileName(name objectivec.IObject) objectivec.IObject
	SetDirectory(directory objectivec.IObject)
	InitWithDirectory(directory objectivec.IObject) GTShaderProfilerDebugDump
}

// Init initializes the instance.
func (g GTShaderProfilerDebugDump) Init() GTShaderProfilerDebugDump {
	rv := objc.Send[GTShaderProfilerDebugDump](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerDebugDump) Autorelease() GTShaderProfilerDebugDump {
	rv := objc.Send[GTShaderProfilerDebugDump](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerDebugDump creates a new GTShaderProfilerDebugDump instance.
func NewGTShaderProfilerDebugDump() GTShaderProfilerDebugDump {
	class := getGTShaderProfilerDebugDumpClass()
	rv := objc.Send[GTShaderProfilerDebugDump](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump/initWithDirectory:
func NewGTShaderProfilerDebugDumpWithDirectory(directory objectivec.IObject) GTShaderProfilerDebugDump {
	instance := getGTShaderProfilerDebugDumpClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDirectory:"), directory)
	return GTShaderProfilerDebugDumpFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump/filePathFromFileName:
func (g GTShaderProfilerDebugDump) FilePathFromFileName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("filePathFromFileName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump/setDirectory:
func (g GTShaderProfilerDebugDump) SetDirectory(directory objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("setDirectory:"), directory)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump/initWithDirectory:
func (g GTShaderProfilerDebugDump) InitWithDirectory(directory objectivec.IObject) GTShaderProfilerDebugDump {
	rv := objc.Send[GTShaderProfilerDebugDump](g.ID, objc.Sel("initWithDirectory:"), directory)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDebugDump/debugDump
func (_GTShaderProfilerDebugDumpClass GTShaderProfilerDebugDumpClass) DebugDump() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerDebugDumpClass.class), objc.Sel("debugDump"))
	return objectivec.Object{ID: rv}
}

