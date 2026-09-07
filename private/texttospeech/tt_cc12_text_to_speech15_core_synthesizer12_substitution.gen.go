// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Substitution] class.
var (
	_SubstitutionClass     SubstitutionClass
	_SubstitutionClassOnce sync.Once
)

func getSubstitutionClass() SubstitutionClass {
	_SubstitutionClassOnce.Do(func() {
		_SubstitutionClass = SubstitutionClass{class: objc.GetClass("_TtCC12TextToSpeech15CoreSynthesizer12Substitution")}
	})
	return _SubstitutionClass
}

// GetSubstitutionClass returns the class object for _TtCC12TextToSpeech15CoreSynthesizer12Substitution.
func GetSubstitutionClass() SubstitutionClass {
	return getSubstitutionClass()
}

type SubstitutionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SubstitutionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SubstitutionClass) Alloc() Substitution {
	rv := objc.SendIfResponds[Substitution](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type Substitution struct {
	objectivec.Object
}

// SubstitutionFromID constructs a [Substitution] from an objc.ID.
func SubstitutionFromID(id objc.ID) Substitution {
	return Substitution{objectivec.Object{ID: id}}
}

// Ensure Substitution implements ISubstitution.
var _ ISubstitution = Substitution{}

// An interface definition for the [Substitution] class.
type ISubstitution interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s Substitution) Init() Substitution {
	rv := objc.SendIfResponds[Substitution](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Substitution) Autorelease() Substitution {
	rv := objc.SendIfResponds[Substitution](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubstitution creates a new Substitution instance.
func NewSubstitution() Substitution {
	class := getSubstitutionClass()
	rv := objc.SendIfResponds[Substitution](objc.ID(class.class), objc.Sel("new"))
	return rv
}
