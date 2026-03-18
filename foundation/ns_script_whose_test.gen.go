// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptWhoseTest] class.
var (
	_NSScriptWhoseTestClass     NSScriptWhoseTestClass
	_NSScriptWhoseTestClassOnce sync.Once
)

func getNSScriptWhoseTestClass() NSScriptWhoseTestClass {
	_NSScriptWhoseTestClassOnce.Do(func() {
		_NSScriptWhoseTestClass = NSScriptWhoseTestClass{class: objc.GetClass("NSScriptWhoseTest")}
	})
	return _NSScriptWhoseTestClass
}

// GetNSScriptWhoseTestClass returns the class object for NSScriptWhoseTest.
func GetNSScriptWhoseTestClass() NSScriptWhoseTestClass {
	return getNSScriptWhoseTestClass()
}

type NSScriptWhoseTestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptWhoseTestClass) Alloc() NSScriptWhoseTest {
	rv := objc.Send[NSScriptWhoseTest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that provides the basis for testing specifiers one at a
// time or in groups.
//
// # Overview
// 
// [NSScriptWhoseTest] is an abstract class whose sole method is [NSScriptWhoseTest.IsTrue]. Two
// concrete subclasses of [NSScriptWhoseTest] generate objects representing
// Boolean expressions comparing one object with another and objects
// representing multiple Boolean expressions connected by logical operators
// ([OR], [AND], [NOT]). These classes are, respectively, [NSSpecifierTest]
// and [NSLogicalTest]. In evaluating itself, an [NSWhoseSpecifier] invokes
// the [NSScriptWhoseTest.IsTrue] method of its “test” object.
// 
// You shouldn’t need to subclass [NSScriptWhoseTest], and you should rarely
// need to subclass one of its subclasses.
//
// # Evaluating a test
//
//   - [NSScriptWhoseTest.IsTrue]: Returns a Boolean value that indicates whether the test represented by the receiver evaluates to true.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest
type NSScriptWhoseTest struct {
	objectivec.Object
}

// NSScriptWhoseTestFromID constructs a [NSScriptWhoseTest] from an objc.ID.
//
// An abstract class that provides the basis for testing specifiers one at a
// time or in groups.
func NSScriptWhoseTestFromID(id objc.ID) NSScriptWhoseTest {
	return NSScriptWhoseTest{objectivec.Object{ID: id}}
}
// NOTE: NSScriptWhoseTest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptWhoseTest] class.
//
// # Evaluating a test
//
//   - [INSScriptWhoseTest.IsTrue]: Returns a Boolean value that indicates whether the test represented by the receiver evaluates to true.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest
type INSScriptWhoseTest interface {
	objectivec.IObject
	NSCoding

	// Topic: Evaluating a test

	// Returns a Boolean value that indicates whether the test represented by the receiver evaluates to true.
	IsTrue() bool
}

// Init initializes the instance.
func (s NSScriptWhoseTest) Init() NSScriptWhoseTest {
	rv := objc.Send[NSScriptWhoseTest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptWhoseTest) Autorelease() NSScriptWhoseTest {
	rv := objc.Send[NSScriptWhoseTest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptWhoseTest creates a new NSScriptWhoseTest instance.
func NewNSScriptWhoseTest() NSScriptWhoseTest {
	class := getNSScriptWhoseTestClass()
	rv := objc.Send[NSScriptWhoseTest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest/init(coder:)
func NewScriptWhoseTestWithCoder(inCoder INSCoder) NSScriptWhoseTest {
	instance := getNSScriptWhoseTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSScriptWhoseTestFromID(rv)
}

// Returns a Boolean value that indicates whether the test represented by the
// receiver evaluates to true.
//
// # Return Value
// 
// [true] if the test represented by the receiver evaluates to [true],
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest/isTrue()
func (s NSScriptWhoseTest) IsTrue() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isTrue"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest/init(coder:)
func (s NSScriptWhoseTest) InitWithCoder(inCoder INSCoder) NSScriptWhoseTest {
	rv := objc.Send[NSScriptWhoseTest](s.ID, objc.Sel("initWithCoder:"), inCoder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSScriptWhoseTest) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

			// Protocol methods for NSCoding
			

