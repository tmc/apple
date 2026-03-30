// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSpecifierTest] class.
var (
	_NSSpecifierTestClass     NSSpecifierTestClass
	_NSSpecifierTestClassOnce sync.Once
)

func getNSSpecifierTestClass() NSSpecifierTestClass {
	_NSSpecifierTestClassOnce.Do(func() {
		_NSSpecifierTestClass = NSSpecifierTestClass{class: objc.GetClass("NSSpecifierTest")}
	})
	return _NSSpecifierTestClass
}

// GetNSSpecifierTestClass returns the class object for NSSpecifierTest.
func GetNSSpecifierTestClass() NSSpecifierTestClass {
	return getNSSpecifierTestClass()
}

type NSSpecifierTestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSpecifierTestClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSpecifierTestClass) Alloc() NSSpecifierTest {
	rv := objc.Send[NSSpecifierTest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A comparison between an object specifier and a test object.
//
// # Overview
//
// Instances of this class represent a Boolean expression; they evaluate an
// object specifier and compare the resulting object to another object using a
// given comparison method. For more information on [NSSpecifierTest], see the
// method description for its sole public method, its initializer,
// [NSSpecifierTest.InitWithObjectSpecifierComparisonOperatorTestObject].
//
// When an [NSSpecifierTest] object is properly initialized, it holds two
// objects:
//
// - A “value” or “test” object used as the basis of the comparison;
// this object can be a regular object or object specifier (such as “blue”
// in “words whose color is blue”). - An object specifier evaluating to
// the container (“words”).
//
// The instance also encapsulates a selector identifying the method performing
// this comparison. The informal protocol [NSComparisonMethods] defines a set
// of comparison methods useful for this purpose, while
// [NSScriptingComparisonMethods] describes additional methods you may need to
// use for scripting.
//
// The test object is compared, using the selector, against each object in the
// container. Specifiers in these tests usually have
// [NSSpecifierTest.ContainerIsObjectBeingTested] invoked on their topmost container.
//
// You should rarely need to subclass [NSSpecifierTest].
//
// # Initializing a specifier test
//
//   - [NSSpecifierTest.InitWithObjectSpecifierComparisonOperatorTestObject]: Returns a specifier test initialized to evaluate a test object against an object specified by an object specifier using a given comparison operation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest
//
// [NSComparisonMethods]: https://developer.apple.com/documentation/Foundation/nscomparisonmethods
// [NSScriptingComparisonMethods]: https://developer.apple.com/documentation/ObjectiveC/nsscriptingcomparisonmethods
type NSSpecifierTest struct {
	NSScriptWhoseTest
}

// NSSpecifierTestFromID constructs a [NSSpecifierTest] from an objc.ID.
//
// A comparison between an object specifier and a test object.
func NSSpecifierTestFromID(id objc.ID) NSSpecifierTest {
	return NSSpecifierTest{NSScriptWhoseTest: NSScriptWhoseTestFromID(id)}
}

// NOTE: NSSpecifierTest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSpecifierTest] class.
//
// # Initializing a specifier test
//
//   - [INSSpecifierTest.InitWithObjectSpecifierComparisonOperatorTestObject]: Returns a specifier test initialized to evaluate a test object against an object specified by an object specifier using a given comparison operation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest
type INSSpecifierTest interface {
	INSScriptWhoseTest

	// Topic: Initializing a specifier test

	// Returns a specifier test initialized to evaluate a test object against an object specified by an object specifier using a given comparison operation.
	InitWithObjectSpecifierComparisonOperatorTestObject(obj1 INSScriptObjectSpecifier, compOp NSTestComparisonOperation, obj2 objectivec.IObject) NSSpecifierTest

	// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
}

// Init initializes the instance.
func (s NSSpecifierTest) Init() NSSpecifierTest {
	rv := objc.Send[NSSpecifierTest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSpecifierTest) Autorelease() NSSpecifierTest {
	rv := objc.Send[NSSpecifierTest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSpecifierTest creates a new NSSpecifierTest instance.
func NewNSSpecifierTest() NSSpecifierTest {
	class := getNSSpecifierTestClass()
	rv := objc.Send[NSSpecifierTest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/init(coder:)
func NewSpecifierTestWithCoder(inCoder INSCoder) NSSpecifierTest {
	instance := getNSSpecifierTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSSpecifierTestFromID(rv)
}

// Returns a specifier test initialized to evaluate a test object against an
// object specified by an object specifier using a given comparison operation.
//
// obj1: An object specifier.
//
// compOp: The comparison operation.
//
// obj2: The object against which to evaluate the object specified by `obj1`.
//
// # Return Value
//
// A specifier test initialized to evaluate (`obj2`) against an object
// specified by `obj1` using the comparison operation `compOp`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/init(objectSpecifier:comparisonOperator:test:)
func NewSpecifierTestWithObjectSpecifierComparisonOperatorTestObject(obj1 INSScriptObjectSpecifier, compOp NSTestComparisonOperation, obj2 objectivec.IObject) NSSpecifierTest {
	instance := getNSSpecifierTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectSpecifier:comparisonOperator:testObject:"), obj1, compOp, obj2)
	return NSSpecifierTestFromID(rv)
}

// Returns a specifier test initialized to evaluate a test object against an
// object specified by an object specifier using a given comparison operation.
//
// obj1: An object specifier.
//
// compOp: The comparison operation.
//
// obj2: The object against which to evaluate the object specified by `obj1`.
//
// # Return Value
//
// A specifier test initialized to evaluate (`obj2`) against an object
// specified by `obj1` using the comparison operation `compOp`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/init(objectSpecifier:comparisonOperator:test:)
func (s NSSpecifierTest) InitWithObjectSpecifierComparisonOperatorTestObject(obj1 INSScriptObjectSpecifier, compOp NSTestComparisonOperation, obj2 objectivec.IObject) NSSpecifierTest {
	rv := objc.Send[NSSpecifierTest](s.ID, objc.Sel("initWithObjectSpecifier:comparisonOperator:testObject:"), obj1, compOp, obj2)
	return rv
}

// Sets whether the receiver’s container should be an object involved in a
// filter reference or the top-level object.
//
// See: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisobjectbeingtested
func (s NSSpecifierTest) ContainerIsObjectBeingTested() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("containerIsObjectBeingTested"))
	return rv
}
func (s NSSpecifierTest) SetContainerIsObjectBeingTested(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setContainerIsObjectBeingTested:"), value)
}
