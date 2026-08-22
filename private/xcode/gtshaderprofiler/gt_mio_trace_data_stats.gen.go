// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceDataStats] class.
var (
	_GTMioTraceDataStatsClass     GTMioTraceDataStatsClass
	_GTMioTraceDataStatsClassOnce sync.Once
)

func getGTMioTraceDataStatsClass() GTMioTraceDataStatsClass {
	_GTMioTraceDataStatsClassOnce.Do(func() {
		_GTMioTraceDataStatsClass = GTMioTraceDataStatsClass{class: objc.GetClass("GTMioTraceDataStats")}
	})
	return _GTMioTraceDataStatsClass
}

// GetGTMioTraceDataStatsClass returns the class object for GTMioTraceDataStats.
func GetGTMioTraceDataStatsClass() GTMioTraceDataStatsClass {
	return getGTMioTraceDataStatsClass()
}

type GTMioTraceDataStatsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceDataStatsClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceDataStatsClass) Alloc() GTMioTraceDataStats {
	rv := objc.SendIfResponds[GTMioTraceDataStats](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioTraceDataStats.Build]
//   - [GTMioTraceDataStats.ShaderStatForShaderProgramType]
//   - [GTMioTraceDataStats.InitWithTraceData]
type GTMioTraceDataStats struct {
	objectivec.Object
}

// GTMioTraceDataStatsFromID constructs a [GTMioTraceDataStats] from an objc.ID.
func GTMioTraceDataStatsFromID(id objc.ID) GTMioTraceDataStats {
	return GTMioTraceDataStats{objectivec.Object{ID: id}}
}

// Ensure GTMioTraceDataStats implements IGTMioTraceDataStats.
var _ IGTMioTraceDataStats = GTMioTraceDataStats{}

// An interface definition for the [GTMioTraceDataStats] class.
//
// # Methods
//
//   - [IGTMioTraceDataStats.Build]
//   - [IGTMioTraceDataStats.ShaderStatForShaderProgramType]
//   - [IGTMioTraceDataStats.InitWithTraceData]
type IGTMioTraceDataStats interface {
	objectivec.IObject

	// Topic: Methods

	Build()
	ShaderStatForShaderProgramType(shader uint64, type_ uint16) objectivec.IObject
	InitWithTraceData(data objectivec.IObject) GTMioTraceDataStats
}

// Init initializes the instance.
func (g GTMioTraceDataStats) Init() GTMioTraceDataStats {
	rv := objc.SendIfResponds[GTMioTraceDataStats](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceDataStats) Autorelease() GTMioTraceDataStats {
	rv := objc.SendIfResponds[GTMioTraceDataStats](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceDataStats creates a new GTMioTraceDataStats instance.
func NewGTMioTraceDataStats() GTMioTraceDataStats {
	class := getGTMioTraceDataStatsClass()
	rv := objc.SendIfResponds[GTMioTraceDataStats](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioTraceDataStatsWithTraceData(data objectivec.IObject) GTMioTraceDataStats {
	instance := getGTMioTraceDataStatsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTraceData:"), data)
	return GTMioTraceDataStatsFromID(rv)
}

func (g GTMioTraceDataStats) Build() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("build"))
}
func (g GTMioTraceDataStats) ShaderStatForShaderProgramType(shader uint64, type_ uint16) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("shaderStatForShader:programType:"), shader, type_)
	return objectivec.Object{ID: rv}
}
func (g GTMioTraceDataStats) InitWithTraceData(data objectivec.IObject) GTMioTraceDataStats {
	rv := objc.SendIfResponds[GTMioTraceDataStats](g.ID, objc.Sel("initWithTraceData:"), data)
	return rv
}
