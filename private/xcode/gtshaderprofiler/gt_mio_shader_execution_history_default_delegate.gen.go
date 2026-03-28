// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryDefaultDelegate] class.
var (
	_GTMioShaderExecutionHistoryDefaultDelegateClass     GTMioShaderExecutionHistoryDefaultDelegateClass
	_GTMioShaderExecutionHistoryDefaultDelegateClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryDefaultDelegateClass() GTMioShaderExecutionHistoryDefaultDelegateClass {
	_GTMioShaderExecutionHistoryDefaultDelegateClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryDefaultDelegateClass = GTMioShaderExecutionHistoryDefaultDelegateClass{class: objc.GetClass("GTMioShaderExecutionHistoryDefaultDelegate")}
	})
	return _GTMioShaderExecutionHistoryDefaultDelegateClass
}

// GetGTMioShaderExecutionHistoryDefaultDelegateClass returns the class object for GTMioShaderExecutionHistoryDefaultDelegate.
func GetGTMioShaderExecutionHistoryDefaultDelegateClass() GTMioShaderExecutionHistoryDefaultDelegateClass {
	return getGTMioShaderExecutionHistoryDefaultDelegateClass()
}

type GTMioShaderExecutionHistoryDefaultDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryDefaultDelegateClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryDefaultDelegateClass) Alloc() GTMioShaderExecutionHistoryDefaultDelegate {
	rv := objc.Send[GTMioShaderExecutionHistoryDefaultDelegate](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderExecutionHistoryDefaultDelegate.ExecutionHistoryProcessCliqueTotal]
//   - [GTMioShaderExecutionHistoryDefaultDelegate.UniqueIdentifierForFileDebugFunctionNameLineColumn]
//   - [GTMioShaderExecutionHistoryDefaultDelegate.DebugDescription]
//   - [GTMioShaderExecutionHistoryDefaultDelegate.Description]
//   - [GTMioShaderExecutionHistoryDefaultDelegate.Hash]
//   - [GTMioShaderExecutionHistoryDefaultDelegate.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate
type GTMioShaderExecutionHistoryDefaultDelegate struct {
	objectivec.Object
}

// GTMioShaderExecutionHistoryDefaultDelegateFromID constructs a [GTMioShaderExecutionHistoryDefaultDelegate] from an objc.ID.
func GTMioShaderExecutionHistoryDefaultDelegateFromID(id objc.ID) GTMioShaderExecutionHistoryDefaultDelegate {
	return GTMioShaderExecutionHistoryDefaultDelegate{objectivec.Object{ID: id}}
}
// Ensure GTMioShaderExecutionHistoryDefaultDelegate implements IGTMioShaderExecutionHistoryDefaultDelegate.
var _ IGTMioShaderExecutionHistoryDefaultDelegate = GTMioShaderExecutionHistoryDefaultDelegate{}

// An interface definition for the [GTMioShaderExecutionHistoryDefaultDelegate] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.ExecutionHistoryProcessCliqueTotal]
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.UniqueIdentifierForFileDebugFunctionNameLineColumn]
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.DebugDescription]
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.Description]
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.Hash]
//   - [IGTMioShaderExecutionHistoryDefaultDelegate.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate
type IGTMioShaderExecutionHistoryDefaultDelegate interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionHistoryProcessCliqueTotal(clique unsafe.Pointer, total uint32)
	UniqueIdentifierForFileDebugFunctionNameLineColumn(file objectivec.IObject, name objectivec.IObject, line uint32, column uint32) uint64
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryDefaultDelegate) Init() GTMioShaderExecutionHistoryDefaultDelegate {
	rv := objc.Send[GTMioShaderExecutionHistoryDefaultDelegate](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryDefaultDelegate) Autorelease() GTMioShaderExecutionHistoryDefaultDelegate {
	rv := objc.Send[GTMioShaderExecutionHistoryDefaultDelegate](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryDefaultDelegate creates a new GTMioShaderExecutionHistoryDefaultDelegate instance.
func NewGTMioShaderExecutionHistoryDefaultDelegate() GTMioShaderExecutionHistoryDefaultDelegate {
	class := getGTMioShaderExecutionHistoryDefaultDelegateClass()
	rv := objc.Send[GTMioShaderExecutionHistoryDefaultDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/executionHistoryProcessClique:total:
func (g GTMioShaderExecutionHistoryDefaultDelegate) ExecutionHistoryProcessCliqueTotal(clique unsafe.Pointer, total uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("executionHistoryProcessClique:total:"), clique, total)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/uniqueIdentifierForFile:debugFunctionName:line:column:
func (g GTMioShaderExecutionHistoryDefaultDelegate) UniqueIdentifierForFileDebugFunctionNameLineColumn(file objectivec.IObject, name objectivec.IObject, line uint32, column uint32) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("uniqueIdentifierForFile:debugFunctionName:line:column:"), file, name, line, column)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/shared
func (_GTMioShaderExecutionHistoryDefaultDelegateClass GTMioShaderExecutionHistoryDefaultDelegateClass) Shared() GTMioShaderExecutionHistoryDefaultDelegate {
	rv := objc.Send[objc.ID](objc.ID(_GTMioShaderExecutionHistoryDefaultDelegateClass.class), objc.Sel("shared"))
	return GTMioShaderExecutionHistoryDefaultDelegateFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/debugDescription
func (g GTMioShaderExecutionHistoryDefaultDelegate) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/description
func (g GTMioShaderExecutionHistoryDefaultDelegate) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/hash
func (g GTMioShaderExecutionHistoryDefaultDelegate) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryDefaultDelegate/superclass
func (g GTMioShaderExecutionHistoryDefaultDelegate) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

