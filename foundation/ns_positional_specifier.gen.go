// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPositionalSpecifier] class.
var (
	_NSPositionalSpecifierClass     NSPositionalSpecifierClass
	_NSPositionalSpecifierClassOnce sync.Once
)

func getNSPositionalSpecifierClass() NSPositionalSpecifierClass {
	_NSPositionalSpecifierClassOnce.Do(func() {
		_NSPositionalSpecifierClass = NSPositionalSpecifierClass{class: objc.GetClass("NSPositionalSpecifier")}
	})
	return _NSPositionalSpecifierClass
}

// GetNSPositionalSpecifierClass returns the class object for NSPositionalSpecifier.
func GetNSPositionalSpecifierClass() NSPositionalSpecifierClass {
	return getNSPositionalSpecifierClass()
}

type NSPositionalSpecifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPositionalSpecifierClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPositionalSpecifierClass) Alloc() NSPositionalSpecifier {
	rv := objc.Send[NSPositionalSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specifier for an insertion point in a container relative to another
// object in the container.
//
// # Overview
//
// Instances of [NSPositionalSpecifier] specify an insertion point in a
// container relative to another object in the container, for example, `before
// first word` or `after paragraph 4`. The container is specified by an
// instance of [NSScriptObjectSpecifier]. [NSPositionalSpecifier] objects
// commonly encapsulate object specifiers used as arguments to the `make`
// (`create`) and `move` commands and indicate where the created or moved
// object is to be inserted relative to the object represented by an object
// specifier.
//
// Invoking an accessor method to obtain information about an instance of
// [NSPositionalSpecifier] causes the object to be evaluated if it hasn’t
// been already.
//
// You don’t normally subclass [NSPositionalSpecifier].
//
// # Initializing a positional specifier
//
//   - [NSPositionalSpecifier.InitWithPositionObjectSpecifier]: Initializes a positional specifier with a given position relative to another given specifier.
//
// # Accessing information about a positional specifier
//
//   - [NSPositionalSpecifier.InsertionContainer]: Returns the container in which the new or copied object or objects should be placed.
//   - [NSPositionalSpecifier.InsertionIndex]: Returns an insertion index that indicates where the new or copied object or objects should be placed.
//   - [NSPositionalSpecifier.InsertionKey]: Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
//   - [NSPositionalSpecifier.InsertionReplaces]: Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
//   - [NSPositionalSpecifier.ObjectSpecifier]: Returns the object specifier specified at initialization time.
//   - [NSPositionalSpecifier.Position]: Returns the insertion position specified at initialization time.
//   - [NSPositionalSpecifier.SetInsertionClassDescription]: Sets the class description for the object or objects to be inserted.
//
// # Evaluating a positional specifier
//
//   - [NSPositionalSpecifier.Evaluate]: Causes the receiver to evaluate its position.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier
type NSPositionalSpecifier struct {
	objectivec.Object
}

// NSPositionalSpecifierFromID constructs a [NSPositionalSpecifier] from an objc.ID.
//
// A specifier for an insertion point in a container relative to another
// object in the container.
func NSPositionalSpecifierFromID(id objc.ID) NSPositionalSpecifier {
	return NSPositionalSpecifier{objectivec.Object{ID: id}}
}

// NOTE: NSPositionalSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPositionalSpecifier] class.
//
// # Initializing a positional specifier
//
//   - [INSPositionalSpecifier.InitWithPositionObjectSpecifier]: Initializes a positional specifier with a given position relative to another given specifier.
//
// # Accessing information about a positional specifier
//
//   - [INSPositionalSpecifier.InsertionContainer]: Returns the container in which the new or copied object or objects should be placed.
//   - [INSPositionalSpecifier.InsertionIndex]: Returns an insertion index that indicates where the new or copied object or objects should be placed.
//   - [INSPositionalSpecifier.InsertionKey]: Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
//   - [INSPositionalSpecifier.InsertionReplaces]: Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
//   - [INSPositionalSpecifier.ObjectSpecifier]: Returns the object specifier specified at initialization time.
//   - [INSPositionalSpecifier.Position]: Returns the insertion position specified at initialization time.
//   - [INSPositionalSpecifier.SetInsertionClassDescription]: Sets the class description for the object or objects to be inserted.
//
// # Evaluating a positional specifier
//
//   - [INSPositionalSpecifier.Evaluate]: Causes the receiver to evaluate its position.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier
type INSPositionalSpecifier interface {
	objectivec.IObject

	// Topic: Initializing a positional specifier

	// Initializes a positional specifier with a given position relative to another given specifier.
	InitWithPositionObjectSpecifier(position NSInsertionPosition, specifier INSScriptObjectSpecifier) NSPositionalSpecifier

	// Topic: Accessing information about a positional specifier

	// Returns the container in which the new or copied object or objects should be placed.
	InsertionContainer() objectivec.IObject
	// Returns an insertion index that indicates where the new or copied object or objects should be placed.
	InsertionIndex() int
	// Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
	InsertionKey() string
	// Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
	InsertionReplaces() bool
	// Returns the object specifier specified at initialization time.
	ObjectSpecifier() INSScriptObjectSpecifier
	// Returns the insertion position specified at initialization time.
	Position() NSInsertionPosition
	// Sets the class description for the object or objects to be inserted.
	SetInsertionClassDescription(classDescription INSScriptClassDescription)

	// Topic: Evaluating a positional specifier

	// Causes the receiver to evaluate its position.
	Evaluate()
}

// Init initializes the instance.
func (p NSPositionalSpecifier) Init() NSPositionalSpecifier {
	rv := objc.Send[NSPositionalSpecifier](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPositionalSpecifier) Autorelease() NSPositionalSpecifier {
	rv := objc.Send[NSPositionalSpecifier](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPositionalSpecifier creates a new NSPositionalSpecifier instance.
func NewNSPositionalSpecifier() NSPositionalSpecifier {
	class := getNSPositionalSpecifierClass()
	rv := objc.Send[NSPositionalSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a positional specifier with a given position relative to
// another given specifier.
//
// position: The position for the new specifier relative to `specifier`.
//
// specifier: The reference specifier.
//
// # Return Value
//
// An initialized positional specifier with the position specified by
// `position` relative to the object specified by `specifier`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/init(position:objectSpecifier:)
func NewPositionalSpecifierWithPositionObjectSpecifier(position NSInsertionPosition, specifier INSScriptObjectSpecifier) NSPositionalSpecifier {
	instance := getNSPositionalSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPosition:objectSpecifier:"), position, specifier)
	return NSPositionalSpecifierFromID(rv)
}

// Initializes a positional specifier with a given position relative to
// another given specifier.
//
// position: The position for the new specifier relative to `specifier`.
//
// specifier: The reference specifier.
//
// # Return Value
//
// An initialized positional specifier with the position specified by
// `position` relative to the object specified by `specifier`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/init(position:objectSpecifier:)
func (p NSPositionalSpecifier) InitWithPositionObjectSpecifier(position NSInsertionPosition, specifier INSScriptObjectSpecifier) NSPositionalSpecifier {
	rv := objc.Send[NSPositionalSpecifier](p.ID, objc.Sel("initWithPosition:objectSpecifier:"), position, specifier)
	return rv
}

// Sets the class description for the object or objects to be inserted.
//
// classDescription: The class description for the object or objects to be inserted.
//
// # Discussion
//
// This message can be sent at any time after object initialization, but must
// be sent before evaluation to have any effect.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/setInsertionClassDescription(_:)
func (p NSPositionalSpecifier) SetInsertionClassDescription(classDescription INSScriptClassDescription) {
	objc.Send[objc.ID](p.ID, objc.Sel("setInsertionClassDescription:"), classDescription)
}

// Causes the receiver to evaluate its position.
//
// # Discussion
//
// Calling [InsertionContainer], [InsertionKey], [InsertionIndex], or
// [InsertionReplaces] also causes the receiver to be evaluated, if it
// hasn’t already been evaluated.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/evaluate()
func (p NSPositionalSpecifier) Evaluate() {
	objc.Send[objc.ID](p.ID, objc.Sel("evaluate"))
}

// Returns the container in which the new or copied object or objects should
// be placed.
//
// # Return Value
//
// A container. Determined by evaluating the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionContainer
func (p NSPositionalSpecifier) InsertionContainer() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("insertionContainer"))
	return objectivec.Object{ID: rv}
}

// Returns an insertion index that indicates where the new or copied object or
// objects should be placed.
//
// # Return Value
//
// An insertion index. Determined by evaluating the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionIndex
func (p NSPositionalSpecifier) InsertionIndex() int {
	rv := objc.Send[int](p.ID, objc.Sel("insertionIndex"))
	return rv
}

// Returns the key that identifies the relationship into which the new or
// copied object or objects should be inserted.
//
// # Return Value
//
// A key. Determined by evaluating the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionKey
func (p NSPositionalSpecifier) InsertionKey() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("insertionKey"))
	return NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether evaluation has been
// successful and the object to be inserted should actually replace the keyed,
// indexed object in the insertion container.
//
// # Return Value
//
// true if evaluation has been successful and the object to be inserted should
// actually replace the keyed, indexed object in the insertion container,
// instead of being inserted before it; false otherwise.
//
// # Discussion
//
// If this object has never been evaluated, evaluation is attempted.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionReplaces
func (p NSPositionalSpecifier) InsertionReplaces() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("insertionReplaces"))
	return rv
}

// Returns the object specifier specified at initialization time.
//
// # Return Value
//
// An object specifier for a container.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/objectSpecifier
func (p NSPositionalSpecifier) ObjectSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("objectSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}

// Returns the insertion position specified at initialization time.
//
// # Return Value
//
// An insertion position.
//
// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/position
func (p NSPositionalSpecifier) Position() NSInsertionPosition {
	rv := objc.Send[NSInsertionPosition](p.ID, objc.Sel("position"))
	return NSInsertionPosition(rv)
}
