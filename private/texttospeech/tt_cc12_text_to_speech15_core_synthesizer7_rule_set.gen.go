// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RuleSet] class.
var (
	_RuleSetClass     RuleSetClass
	_RuleSetClassOnce sync.Once
)

func getRuleSetClass() RuleSetClass {
	_RuleSetClassOnce.Do(func() {
		_RuleSetClass = RuleSetClass{class: objc.GetClass("_TtCC12TextToSpeech15CoreSynthesizer7RuleSet")}
	})
	return _RuleSetClass
}

// GetRuleSetClass returns the class object for _TtCC12TextToSpeech15CoreSynthesizer7RuleSet.
func GetRuleSetClass() RuleSetClass {
	return getRuleSetClass()
}

type RuleSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RuleSetClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RuleSetClass) Alloc() RuleSet {
	rv := objc.SendIfResponds[RuleSet](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

type RuleSet struct {
	objectivec.Object
}

// RuleSetFromID constructs a [RuleSet] from an objc.ID.
func RuleSetFromID(id objc.ID) RuleSet {
	return RuleSet{objectivec.Object{ID: id}}
}

// Ensure RuleSet implements IRuleSet.
var _ IRuleSet = RuleSet{}

// An interface definition for the [RuleSet] class.
type IRuleSet interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r RuleSet) Init() RuleSet {
	rv := objc.SendIfResponds[RuleSet](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RuleSet) Autorelease() RuleSet {
	rv := objc.SendIfResponds[RuleSet](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRuleSet creates a new RuleSet instance.
func NewRuleSet() RuleSet {
	class := getRuleSetClass()
	rv := objc.SendIfResponds[RuleSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}
