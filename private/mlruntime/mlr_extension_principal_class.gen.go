// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRExtensionPrincipalClass] class.
var (
	_MLRExtensionPrincipalClassClass     MLRExtensionPrincipalClassClass
	_MLRExtensionPrincipalClassClassOnce sync.Once
)

func getMLRExtensionPrincipalClassClass() MLRExtensionPrincipalClassClass {
	_MLRExtensionPrincipalClassClassOnce.Do(func() {
		_MLRExtensionPrincipalClassClass = MLRExtensionPrincipalClassClass{class: objc.GetClass("MLRExtensionPrincipalClass")}
	})
	return _MLRExtensionPrincipalClassClass
}

// GetMLRExtensionPrincipalClassClass returns the class object for MLRExtensionPrincipalClass.
func GetMLRExtensionPrincipalClassClass() MLRExtensionPrincipalClassClass {
	return getMLRExtensionPrincipalClassClass()
}

type MLRExtensionPrincipalClassClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRExtensionPrincipalClassClass) Alloc() MLRExtensionPrincipalClass {
	rv := objc.Send[MLRExtensionPrincipalClass](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRExtensionPrincipalClass.BeginRequestWithExtensionContext]
//   - [MLRExtensionPrincipalClass.DebugDescription]
//   - [MLRExtensionPrincipalClass.Description]
//   - [MLRExtensionPrincipalClass.Hash]
//   - [MLRExtensionPrincipalClass.Superclass]
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass
type MLRExtensionPrincipalClass struct {
	objectivec.Object
}

// MLRExtensionPrincipalClassFromID constructs a [MLRExtensionPrincipalClass] from an objc.ID.
func MLRExtensionPrincipalClassFromID(id objc.ID) MLRExtensionPrincipalClass {
	return MLRExtensionPrincipalClass{objectivec.Object{ID: id}}
}
// Ensure MLRExtensionPrincipalClass implements IMLRExtensionPrincipalClass.
var _ IMLRExtensionPrincipalClass = MLRExtensionPrincipalClass{}

// An interface definition for the [MLRExtensionPrincipalClass] class.
//
// # Methods
//
//   - [IMLRExtensionPrincipalClass.BeginRequestWithExtensionContext]
//   - [IMLRExtensionPrincipalClass.DebugDescription]
//   - [IMLRExtensionPrincipalClass.Description]
//   - [IMLRExtensionPrincipalClass.Hash]
//   - [IMLRExtensionPrincipalClass.Superclass]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass
type IMLRExtensionPrincipalClass interface {
	objectivec.IObject

	// Topic: Methods

	BeginRequestWithExtensionContext(context objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (r MLRExtensionPrincipalClass) Init() MLRExtensionPrincipalClass {
	rv := objc.Send[MLRExtensionPrincipalClass](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRExtensionPrincipalClass) Autorelease() MLRExtensionPrincipalClass {
	rv := objc.Send[MLRExtensionPrincipalClass](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRExtensionPrincipalClass creates a new MLRExtensionPrincipalClass instance.
func NewMLRExtensionPrincipalClass() MLRExtensionPrincipalClass {
	class := getMLRExtensionPrincipalClassClass()
	rv := objc.Send[MLRExtensionPrincipalClass](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass/beginRequestWithExtensionContext:
func (r MLRExtensionPrincipalClass) BeginRequestWithExtensionContext(context objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("beginRequestWithExtensionContext:"), context)
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass/debugDescription
func (r MLRExtensionPrincipalClass) DebugDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass/description
func (r MLRExtensionPrincipalClass) Description() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass/hash
func (r MLRExtensionPrincipalClass) Hash() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRExtensionPrincipalClass/superclass
func (r MLRExtensionPrincipalClass) Superclass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("superclass"))
	return rv
}

