// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLogicalTest] class.
var (
	_NSLogicalTestClass     NSLogicalTestClass
	_NSLogicalTestClassOnce sync.Once
)

func getNSLogicalTestClass() NSLogicalTestClass {
	_NSLogicalTestClassOnce.Do(func() {
		_NSLogicalTestClass = NSLogicalTestClass{class: objc.GetClass("NSLogicalTest")}
	})
	return _NSLogicalTestClass
}

// GetNSLogicalTestClass returns the class object for NSLogicalTest.
func GetNSLogicalTestClass() NSLogicalTestClass {
	return getNSLogicalTestClass()
}

type NSLogicalTestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLogicalTestClass) Alloc() NSLogicalTest {
	rv := objc.Send[NSLogicalTest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The logical combination of one or more specifier tests.
//
// # Overview
// 
// Instances of this class perform logical operations of [AND], [OR], and
// [NOT] on Boolean expressions represented by [NSSpecifierTest] objects.
// These operators are equivalent to “`&&`”, “`||`”, and “`!`” in
// the C language.
// 
// For [AND] and [OR] operations, an [NSLogicalTest] object is typically
// initialized with an array containing two or more [NSSpecifierTest] objects.
// [IsTrue]—inherited from [NSScriptWhoseTest]—evaluates the array in a
// manner appropriate to the logical operation. For [NOT] operations, an
// [NSLogicalTest] object is initialized with only one [NSSpecifierTest]
// object; it simply reverses the Boolean outcome of the [IsTrue] method.
// 
// You don’t normally subclass [NSLogicalTest].
//
// # Initializing a logical test
//
//   - [NSLogicalTest.InitAndTestWithTests]: Returns an [NSLogicalTest] object initialized to perform an [AND] operation with the [NSSpecifierTest] objects in a given array.
//   - [NSLogicalTest.InitNotTestWithTest]: Returns an [NSLogicalTest] object initialized to perform a [NOT] operation on the given [NSScriptWhoseTest] object.
//   - [NSLogicalTest.InitOrTestWithTests]: Returns an [NSLogicalTest] object initialized to perform an [OR] operation with the [NSSpecifierTest] objects in a given array.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest
type NSLogicalTest struct {
	NSScriptWhoseTest
}

// NSLogicalTestFromID constructs a [NSLogicalTest] from an objc.ID.
//
// The logical combination of one or more specifier tests.
func NSLogicalTestFromID(id objc.ID) NSLogicalTest {
	return NSLogicalTest{NSScriptWhoseTest: NSScriptWhoseTestFromID(id)}
}
// NOTE: NSLogicalTest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLogicalTest] class.
//
// # Initializing a logical test
//
//   - [INSLogicalTest.InitAndTestWithTests]: Returns an [NSLogicalTest] object initialized to perform an [AND] operation with the [NSSpecifierTest] objects in a given array.
//   - [INSLogicalTest.InitNotTestWithTest]: Returns an [NSLogicalTest] object initialized to perform a [NOT] operation on the given [NSScriptWhoseTest] object.
//   - [INSLogicalTest.InitOrTestWithTests]: Returns an [NSLogicalTest] object initialized to perform an [OR] operation with the [NSSpecifierTest] objects in a given array.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest
type INSLogicalTest interface {
	INSScriptWhoseTest

	// Topic: Initializing a logical test

	// Returns an [NSLogicalTest] object initialized to perform an [AND] operation with the [NSSpecifierTest] objects in a given array.
	InitAndTestWithTests(subTests []NSSpecifierTest) NSLogicalTest
	// Returns an [NSLogicalTest] object initialized to perform a [NOT] operation on the given [NSScriptWhoseTest] object.
	InitNotTestWithTest(subTest INSScriptWhoseTest) NSLogicalTest
	// Returns an [NSLogicalTest] object initialized to perform an [OR] operation with the [NSSpecifierTest] objects in a given array.
	InitOrTestWithTests(subTests []NSSpecifierTest) NSLogicalTest
}

// Init initializes the instance.
func (l NSLogicalTest) Init() NSLogicalTest {
	rv := objc.Send[NSLogicalTest](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLogicalTest) Autorelease() NSLogicalTest {
	rv := objc.Send[NSLogicalTest](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLogicalTest creates a new NSLogicalTest instance.
func NewNSLogicalTest() NSLogicalTest {
	class := getNSLogicalTestClass()
	rv := objc.Send[NSLogicalTest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSLogicalTest] object initialized to perform an [AND] operation
// with the [NSSpecifierTest] objects in a given array.
//
// subTests: An array of [NSSpecifierTest] objects representing Boolean expressions.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform an [AND] operation with
// the [NSSpecifierTest] objects in `subTests`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(andTestWith:)
func NewLogicalTestAndTestWithTests(subTests []NSSpecifierTest) NSLogicalTest {
	instance := getNSLogicalTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initAndTestWithTests:"), objectivec.IObjectSliceToNSArray(subTests))
	return NSLogicalTestFromID(rv)
}

// Returns an [NSLogicalTest] object initialized to perform a [NOT] operation
// on the given [NSScriptWhoseTest] object.
//
// subTest: The [NSScriptWhoseTest] object to invert.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform a [NOT] operation on
// `subTest`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(notTestWith:)
func NewLogicalTestNotTestWithTest(subTest INSScriptWhoseTest) NSLogicalTest {
	instance := getNSLogicalTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initNotTestWithTest:"), subTest)
	return NSLogicalTestFromID(rv)
}

// Returns an [NSLogicalTest] object initialized to perform an [OR] operation
// with the [NSSpecifierTest] objects in a given array.
//
// subTests: An array of [NSSpecifierTest] objects representing Boolean expressions.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform an [OR] operation with the
// [NSSpecifierTest] objects in `subTests`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(orTestWith:)
func NewLogicalTestOrTestWithTests(subTests []NSSpecifierTest) NSLogicalTest {
	instance := getNSLogicalTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initOrTestWithTests:"), objectivec.IObjectSliceToNSArray(subTests))
	return NSLogicalTestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSScriptWhoseTest/init(coder:)
func NewLogicalTestWithCoder(inCoder INSCoder) NSLogicalTest {
	instance := getNSLogicalTestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSLogicalTestFromID(rv)
}

// Returns an [NSLogicalTest] object initialized to perform an [AND] operation
// with the [NSSpecifierTest] objects in a given array.
//
// subTests: An array of [NSSpecifierTest] objects representing Boolean expressions.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform an [AND] operation with
// the [NSSpecifierTest] objects in `subTests`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(andTestWith:)
func (l NSLogicalTest) InitAndTestWithTests(subTests []NSSpecifierTest) NSLogicalTest {
	rv := objc.Send[NSLogicalTest](l.ID, objc.Sel("initAndTestWithTests:"), objectivec.IObjectSliceToNSArray(subTests))
	return rv
}

// Returns an [NSLogicalTest] object initialized to perform a [NOT] operation
// on the given [NSScriptWhoseTest] object.
//
// subTest: The [NSScriptWhoseTest] object to invert.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform a [NOT] operation on
// `subTest`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(notTestWith:)
func (l NSLogicalTest) InitNotTestWithTest(subTest INSScriptWhoseTest) NSLogicalTest {
	rv := objc.Send[NSLogicalTest](l.ID, objc.Sel("initNotTestWithTest:"), subTest)
	return rv
}

// Returns an [NSLogicalTest] object initialized to perform an [OR] operation
// with the [NSSpecifierTest] objects in a given array.
//
// subTests: An array of [NSSpecifierTest] objects representing Boolean expressions.
//
// # Return Value
// 
// An [NSLogicalTest] object initialized to perform an [OR] operation with the
// [NSSpecifierTest] objects in `subTests`.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogicalTest/init(orTestWith:)
func (l NSLogicalTest) InitOrTestWithTests(subTests []NSSpecifierTest) NSLogicalTest {
	rv := objc.Send[NSLogicalTest](l.ID, objc.Sel("initOrTestWithTests:"), objectivec.IObjectSliceToNSArray(subTests))
	return rv
}

