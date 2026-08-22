// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveSeqwWseqTransposes] class.
var (
	_EspressoPassRemoveSeqwWseqTransposesClass     EspressoPassRemoveSeqwWseqTransposesClass
	_EspressoPassRemoveSeqwWseqTransposesClassOnce sync.Once
)

func getEspressoPassRemoveSeqwWseqTransposesClass() EspressoPassRemoveSeqwWseqTransposesClass {
	_EspressoPassRemoveSeqwWseqTransposesClassOnce.Do(func() {
		_EspressoPassRemoveSeqwWseqTransposesClass = EspressoPassRemoveSeqwWseqTransposesClass{class: objc.GetClass("EspressoPass_remove_seqw_wseq_transposes")}
	})
	return _EspressoPassRemoveSeqwWseqTransposesClass
}

// GetEspressoPassRemoveSeqwWseqTransposesClass returns the class object for EspressoPass_remove_seqw_wseq_transposes.
func GetEspressoPassRemoveSeqwWseqTransposesClass() EspressoPassRemoveSeqwWseqTransposesClass {
	return getEspressoPassRemoveSeqwWseqTransposesClass()
}

type EspressoPassRemoveSeqwWseqTransposesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveSeqwWseqTransposesClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveSeqwWseqTransposesClass) Alloc() EspressoPassRemoveSeqwWseqTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveSeqwWseqTransposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassRemoveSeqwWseqTransposes struct {
	EspressoCustomPass
}

// EspressoPassRemoveSeqwWseqTransposesFromID constructs a [EspressoPassRemoveSeqwWseqTransposes] from an objc.ID.
func EspressoPassRemoveSeqwWseqTransposesFromID(id objc.ID) EspressoPassRemoveSeqwWseqTransposes {
	return EspressoPassRemoveSeqwWseqTransposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_seqw_wseq_transposesFromID is an alias for [EspressoPassRemoveSeqwWseqTransposesFromID] for cross-framework compatibility.
func EspressoPass_remove_seqw_wseq_transposesFromID(id objc.ID) EspressoPassRemoveSeqwWseqTransposes {
	return EspressoPassRemoveSeqwWseqTransposesFromID(id)
}

// Ensure EspressoPassRemoveSeqwWseqTransposes implements IEspressoPassRemoveSeqwWseqTransposes.
var _ IEspressoPassRemoveSeqwWseqTransposes = EspressoPassRemoveSeqwWseqTransposes{}

// An interface definition for the [EspressoPassRemoveSeqwWseqTransposes] class.
type IEspressoPassRemoveSeqwWseqTransposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveSeqwWseqTransposes) Init() EspressoPassRemoveSeqwWseqTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveSeqwWseqTransposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveSeqwWseqTransposes) Autorelease() EspressoPassRemoveSeqwWseqTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveSeqwWseqTransposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveSeqwWseqTransposes creates a new EspressoPassRemoveSeqwWseqTransposes instance.
func NewEspressoPassRemoveSeqwWseqTransposes() EspressoPassRemoveSeqwWseqTransposes {
	class := getEspressoPassRemoveSeqwWseqTransposesClass()
	rv := objc.SendIfResponds[EspressoPassRemoveSeqwWseqTransposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}
