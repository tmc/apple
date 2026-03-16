// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_remove_seqw_wseq_transposes] class.
var (
	_EspressoPass_remove_seqw_wseq_transposesClass     EspressoPass_remove_seqw_wseq_transposesClass
	_EspressoPass_remove_seqw_wseq_transposesClassOnce sync.Once
)

func getEspressoPass_remove_seqw_wseq_transposesClass() EspressoPass_remove_seqw_wseq_transposesClass {
	_EspressoPass_remove_seqw_wseq_transposesClassOnce.Do(func() {
		_EspressoPass_remove_seqw_wseq_transposesClass = EspressoPass_remove_seqw_wseq_transposesClass{class: objc.GetClass("EspressoPass_remove_seqw_wseq_transposes")}
	})
	return _EspressoPass_remove_seqw_wseq_transposesClass
}

// GetEspressoPass_remove_seqw_wseq_transposesClass returns the class object for EspressoPass_remove_seqw_wseq_transposes.
func GetEspressoPass_remove_seqw_wseq_transposesClass() EspressoPass_remove_seqw_wseq_transposesClass {
	return getEspressoPass_remove_seqw_wseq_transposesClass()
}

type EspressoPass_remove_seqw_wseq_transposesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_remove_seqw_wseq_transposesClass) Alloc() EspressoPass_remove_seqw_wseq_transposes {
	rv := objc.Send[EspressoPass_remove_seqw_wseq_transposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_seqw_wseq_transposes
type EspressoPass_remove_seqw_wseq_transposes struct {
	EspressoCustomPass
}

// EspressoPass_remove_seqw_wseq_transposesFromID constructs a [EspressoPass_remove_seqw_wseq_transposes] from an objc.ID.
func EspressoPass_remove_seqw_wseq_transposesFromID(id objc.ID) EspressoPass_remove_seqw_wseq_transposes {
	return EspressoPass_remove_seqw_wseq_transposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_remove_seqw_wseq_transposes implements IEspressoPass_remove_seqw_wseq_transposes.
var _ IEspressoPass_remove_seqw_wseq_transposes = EspressoPass_remove_seqw_wseq_transposes{}

// An interface definition for the [EspressoPass_remove_seqw_wseq_transposes] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_seqw_wseq_transposes
type IEspressoPass_remove_seqw_wseq_transposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_remove_seqw_wseq_transposes) Init() EspressoPass_remove_seqw_wseq_transposes {
	rv := objc.Send[EspressoPass_remove_seqw_wseq_transposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_remove_seqw_wseq_transposes) Autorelease() EspressoPass_remove_seqw_wseq_transposes {
	rv := objc.Send[EspressoPass_remove_seqw_wseq_transposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_remove_seqw_wseq_transposes creates a new EspressoPass_remove_seqw_wseq_transposes instance.
func NewEspressoPass_remove_seqw_wseq_transposes() EspressoPass_remove_seqw_wseq_transposes {
	class := getEspressoPass_remove_seqw_wseq_transposesClass()
	rv := objc.Send[EspressoPass_remove_seqw_wseq_transposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

