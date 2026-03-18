// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLocalizedNumberFormatRule] class.
var (
	_NSLocalizedNumberFormatRuleClass     NSLocalizedNumberFormatRuleClass
	_NSLocalizedNumberFormatRuleClassOnce sync.Once
)

func getNSLocalizedNumberFormatRuleClass() NSLocalizedNumberFormatRuleClass {
	_NSLocalizedNumberFormatRuleClassOnce.Do(func() {
		_NSLocalizedNumberFormatRuleClass = NSLocalizedNumberFormatRuleClass{class: objc.GetClass("NSLocalizedNumberFormatRule")}
	})
	return _NSLocalizedNumberFormatRuleClass
}

// GetNSLocalizedNumberFormatRuleClass returns the class object for NSLocalizedNumberFormatRule.
func GetNSLocalizedNumberFormatRuleClass() NSLocalizedNumberFormatRuleClass {
	return getNSLocalizedNumberFormatRuleClass()
}

type NSLocalizedNumberFormatRuleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLocalizedNumberFormatRuleClass) Alloc() NSLocalizedNumberFormatRule {
	rv := objc.Send[NSLocalizedNumberFormatRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule
type NSLocalizedNumberFormatRule struct {
	objectivec.Object
}

// NSLocalizedNumberFormatRuleFromID constructs a [NSLocalizedNumberFormatRule] from an objc.ID.
func NSLocalizedNumberFormatRuleFromID(id objc.ID) NSLocalizedNumberFormatRule {
	return NSLocalizedNumberFormatRule{objectivec.Object{ID: id}}
}
// NOTE: NSLocalizedNumberFormatRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLocalizedNumberFormatRule] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule
type INSLocalizedNumberFormatRule interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	InitWithCoder(coder INSCoder) NSLocalizedNumberFormatRule
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (l NSLocalizedNumberFormatRule) Init() NSLocalizedNumberFormatRule {
	rv := objc.Send[NSLocalizedNumberFormatRule](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLocalizedNumberFormatRule) Autorelease() NSLocalizedNumberFormatRule {
	rv := objc.Send[NSLocalizedNumberFormatRule](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLocalizedNumberFormatRule creates a new NSLocalizedNumberFormatRule instance.
func NewNSLocalizedNumberFormatRule() NSLocalizedNumberFormatRule {
	class := getNSLocalizedNumberFormatRuleClass()
	rv := objc.Send[NSLocalizedNumberFormatRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (l NSLocalizedNumberFormatRule) InitWithCoder(coder INSCoder) NSLocalizedNumberFormatRule {
	rv := objc.Send[NSLocalizedNumberFormatRule](l.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (l NSLocalizedNumberFormatRule) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule/automatic
func (_NSLocalizedNumberFormatRuleClass NSLocalizedNumberFormatRuleClass) Automatic() NSLocalizedNumberFormatRule {
	rv := objc.Send[objc.ID](objc.ID(_NSLocalizedNumberFormatRuleClass.class), objc.Sel("automatic"))
	return NSLocalizedNumberFormatRuleFromID(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

