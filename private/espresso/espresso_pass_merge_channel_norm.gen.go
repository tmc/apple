// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMergeChannelNorm] class.
var (
	_EspressoPassMergeChannelNormClass     EspressoPassMergeChannelNormClass
	_EspressoPassMergeChannelNormClassOnce sync.Once
)

func getEspressoPassMergeChannelNormClass() EspressoPassMergeChannelNormClass {
	_EspressoPassMergeChannelNormClassOnce.Do(func() {
		_EspressoPassMergeChannelNormClass = EspressoPassMergeChannelNormClass{class: objc.GetClass("EspressoPass_merge_channel_norm")}
	})
	return _EspressoPassMergeChannelNormClass
}

// GetEspressoPassMergeChannelNormClass returns the class object for EspressoPass_merge_channel_norm.
func GetEspressoPassMergeChannelNormClass() EspressoPassMergeChannelNormClass {
	return getEspressoPassMergeChannelNormClass()
}

type EspressoPassMergeChannelNormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMergeChannelNormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMergeChannelNormClass) Alloc() EspressoPassMergeChannelNorm {
	rv := objc.SendIfResponds[EspressoPassMergeChannelNorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassMergeChannelNorm struct {
	EspressoCustomPass
}

// EspressoPassMergeChannelNormFromID constructs a [EspressoPassMergeChannelNorm] from an objc.ID.
func EspressoPassMergeChannelNormFromID(id objc.ID) EspressoPassMergeChannelNorm {
	return EspressoPassMergeChannelNorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_merge_channel_normFromID is an alias for [EspressoPassMergeChannelNormFromID] for cross-framework compatibility.
func EspressoPass_merge_channel_normFromID(id objc.ID) EspressoPassMergeChannelNorm {
	return EspressoPassMergeChannelNormFromID(id)
}

// Ensure EspressoPassMergeChannelNorm implements IEspressoPassMergeChannelNorm.
var _ IEspressoPassMergeChannelNorm = EspressoPassMergeChannelNorm{}

// An interface definition for the [EspressoPassMergeChannelNorm] class.
type IEspressoPassMergeChannelNorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMergeChannelNorm) Init() EspressoPassMergeChannelNorm {
	rv := objc.SendIfResponds[EspressoPassMergeChannelNorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMergeChannelNorm) Autorelease() EspressoPassMergeChannelNorm {
	rv := objc.SendIfResponds[EspressoPassMergeChannelNorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMergeChannelNorm creates a new EspressoPassMergeChannelNorm instance.
func NewEspressoPassMergeChannelNorm() EspressoPassMergeChannelNorm {
	class := getEspressoPassMergeChannelNormClass()
	rv := objc.SendIfResponds[EspressoPassMergeChannelNorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}
