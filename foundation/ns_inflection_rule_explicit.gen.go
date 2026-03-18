// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSInflectionRuleExplicit] class.
var (
	_NSInflectionRuleExplicitClass     NSInflectionRuleExplicitClass
	_NSInflectionRuleExplicitClassOnce sync.Once
)

func getNSInflectionRuleExplicitClass() NSInflectionRuleExplicitClass {
	_NSInflectionRuleExplicitClassOnce.Do(func() {
		_NSInflectionRuleExplicitClass = NSInflectionRuleExplicitClass{class: objc.GetClass("NSInflectionRuleExplicit")}
	})
	return _NSInflectionRuleExplicitClass
}

// GetNSInflectionRuleExplicitClass returns the class object for NSInflectionRuleExplicit.
func GetNSInflectionRuleExplicitClass() NSInflectionRuleExplicitClass {
	return getNSInflectionRuleExplicitClass()
}

type NSInflectionRuleExplicitClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSInflectionRuleExplicitClass) Alloc() NSInflectionRuleExplicit {
	rv := objc.Send[NSInflectionRuleExplicit](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An inflection rule that uses a morphology instance to determine how to
// inflect attribued strings.
//
// # Creating an Explicit Inflection Rule
//
//   - [NSInflectionRuleExplicit.InitWithMorphology]: Creates an inflection rule with the given morphology.
//
// # Accessing Rule Properties
//
//   - [NSInflectionRuleExplicit.Morphology]: The morphology used by this inflection rule.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit
type NSInflectionRuleExplicit struct {
	NSInflectionRule
}

// NSInflectionRuleExplicitFromID constructs a [NSInflectionRuleExplicit] from an objc.ID.
//
// An inflection rule that uses a morphology instance to determine how to
// inflect attribued strings.
func NSInflectionRuleExplicitFromID(id objc.ID) NSInflectionRuleExplicit {
	return NSInflectionRuleExplicit{NSInflectionRule: NSInflectionRuleFromID(id)}
}
// Ensure NSInflectionRuleExplicit implements INSInflectionRuleExplicit.
var _ INSInflectionRuleExplicit = NSInflectionRuleExplicit{}

// An interface definition for the [NSInflectionRuleExplicit] class.
//
// # Creating an Explicit Inflection Rule
//
//   - [INSInflectionRuleExplicit.InitWithMorphology]: Creates an inflection rule with the given morphology.
//
// # Accessing Rule Properties
//
//   - [INSInflectionRuleExplicit.Morphology]: The morphology used by this inflection rule.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit
type INSInflectionRuleExplicit interface {
	INSInflectionRule

	// Topic: Creating an Explicit Inflection Rule

	// Creates an inflection rule with the given morphology.
	InitWithMorphology(morphology INSMorphology) NSInflectionRuleExplicit

	// Topic: Accessing Rule Properties

	// The morphology used by this inflection rule.
	Morphology() INSMorphology
}

// Init initializes the instance.
func (i NSInflectionRuleExplicit) Init() NSInflectionRuleExplicit {
	rv := objc.Send[NSInflectionRuleExplicit](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSInflectionRuleExplicit) Autorelease() NSInflectionRuleExplicit {
	rv := objc.Send[NSInflectionRuleExplicit](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSInflectionRuleExplicit creates a new NSInflectionRuleExplicit instance.
func NewNSInflectionRuleExplicit() NSInflectionRuleExplicit {
	class := getNSInflectionRuleExplicitClass()
	rv := objc.Send[NSInflectionRuleExplicit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an inflection rule with the given morphology.
//
// morphology: The morphology this rule applies when inflecting.
//
// # Return Value
// 
// An inflection rule that uses the given morphology.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit/initWithMorphology:
func NewInflectionRuleExplicitWithMorphology(morphology INSMorphology) NSInflectionRuleExplicit {
	instance := getNSInflectionRuleExplicitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMorphology:"), morphology)
	return NSInflectionRuleExplicitFromID(rv)
}

// Creates an inflection rule with the given morphology.
//
// morphology: The morphology this rule applies when inflecting.
//
// # Return Value
// 
// An inflection rule that uses the given morphology.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit/initWithMorphology:
func (i NSInflectionRuleExplicit) InitWithMorphology(morphology INSMorphology) NSInflectionRuleExplicit {
	rv := objc.Send[NSInflectionRuleExplicit](i.ID, objc.Sel("initWithMorphology:"), morphology)
	return rv
}

// The morphology used by this inflection rule.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit/morphology
func (i NSInflectionRuleExplicit) Morphology() INSMorphology {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("morphology"))
	return NSMorphologyFromID(objc.ID(rv))
}

