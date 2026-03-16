// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingANEcompilerAnalytics] class.
var (
	_EspressoProfilingANEcompilerAnalyticsClass     EspressoProfilingANEcompilerAnalyticsClass
	_EspressoProfilingANEcompilerAnalyticsClassOnce sync.Once
)

func getEspressoProfilingANEcompilerAnalyticsClass() EspressoProfilingANEcompilerAnalyticsClass {
	_EspressoProfilingANEcompilerAnalyticsClassOnce.Do(func() {
		_EspressoProfilingANEcompilerAnalyticsClass = EspressoProfilingANEcompilerAnalyticsClass{class: objc.GetClass("EspressoProfilingANEcompilerAnalytics")}
	})
	return _EspressoProfilingANEcompilerAnalyticsClass
}

// GetEspressoProfilingANEcompilerAnalyticsClass returns the class object for EspressoProfilingANEcompilerAnalytics.
func GetEspressoProfilingANEcompilerAnalyticsClass() EspressoProfilingANEcompilerAnalyticsClass {
	return getEspressoProfilingANEcompilerAnalyticsClass()
}

type EspressoProfilingANEcompilerAnalyticsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingANEcompilerAnalyticsClass) Alloc() EspressoProfilingANEcompilerAnalytics {
	rv := objc.Send[EspressoProfilingANEcompilerAnalytics](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoProfilingANEcompilerAnalytics.Compiler_analytics_file_names]
//   - [EspressoProfilingANEcompilerAnalytics.SetCompiler_analytics_file_names]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingANEcompilerAnalytics
type EspressoProfilingANEcompilerAnalytics struct {
	objectivec.Object
}

// EspressoProfilingANEcompilerAnalyticsFromID constructs a [EspressoProfilingANEcompilerAnalytics] from an objc.ID.
func EspressoProfilingANEcompilerAnalyticsFromID(id objc.ID) EspressoProfilingANEcompilerAnalytics {
	return EspressoProfilingANEcompilerAnalytics{objectivec.Object{id}}
}
// Ensure EspressoProfilingANEcompilerAnalytics implements IEspressoProfilingANEcompilerAnalytics.
var _ IEspressoProfilingANEcompilerAnalytics = EspressoProfilingANEcompilerAnalytics{}

// An interface definition for the [EspressoProfilingANEcompilerAnalytics] class.
//
// # Methods
//
//   - [IEspressoProfilingANEcompilerAnalytics.Compiler_analytics_file_names]
//   - [IEspressoProfilingANEcompilerAnalytics.SetCompiler_analytics_file_names]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingANEcompilerAnalytics
type IEspressoProfilingANEcompilerAnalytics interface {
	objectivec.IObject

	// Topic: Methods

	Compiler_analytics_file_names() foundation.INSArray
	SetCompiler_analytics_file_names(value foundation.INSArray)
}

// Init initializes the instance.
func (e EspressoProfilingANEcompilerAnalytics) Init() EspressoProfilingANEcompilerAnalytics {
	rv := objc.Send[EspressoProfilingANEcompilerAnalytics](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingANEcompilerAnalytics) Autorelease() EspressoProfilingANEcompilerAnalytics {
	rv := objc.Send[EspressoProfilingANEcompilerAnalytics](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingANEcompilerAnalytics creates a new EspressoProfilingANEcompilerAnalytics instance.
func NewEspressoProfilingANEcompilerAnalytics() EspressoProfilingANEcompilerAnalytics {
	class := getEspressoProfilingANEcompilerAnalyticsClass()
	rv := objc.Send[EspressoProfilingANEcompilerAnalytics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingANEcompilerAnalytics/compiler_analytics_file_names
func (e EspressoProfilingANEcompilerAnalytics) Compiler_analytics_file_names() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("compiler_analytics_file_names"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoProfilingANEcompilerAnalytics) SetCompiler_analytics_file_names(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setCompiler_analytics_file_names:"), value)
}

