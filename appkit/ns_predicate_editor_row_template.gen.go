// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPredicateEditorRowTemplate] class.
var (
	_NSPredicateEditorRowTemplateClass     NSPredicateEditorRowTemplateClass
	_NSPredicateEditorRowTemplateClassOnce sync.Once
)

func getNSPredicateEditorRowTemplateClass() NSPredicateEditorRowTemplateClass {
	_NSPredicateEditorRowTemplateClassOnce.Do(func() {
		_NSPredicateEditorRowTemplateClass = NSPredicateEditorRowTemplateClass{class: objc.GetClass("NSPredicateEditorRowTemplate")}
	})
	return _NSPredicateEditorRowTemplateClass
}

// GetNSPredicateEditorRowTemplateClass returns the class object for NSPredicateEditorRowTemplate.
func GetNSPredicateEditorRowTemplateClass() NSPredicateEditorRowTemplateClass {
	return getNSPredicateEditorRowTemplateClass()
}

type NSPredicateEditorRowTemplateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPredicateEditorRowTemplateClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPredicateEditorRowTemplateClass) Alloc() NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A template that describes available predicates and how to display them.
//
// # Overview
//
// You can create instances of [NSPredicateEditorRowTemplate] programmatically
// or in Interface Builder. By default, a noncompound row template has three
// views: a popup (or static text field) on the left, a popup or static text
// field for operators, and either a popup or other view on the right. You can
// subclass [NSPredicateEditorRowTemplate] to create a row template with
// different numbers or types of views.
//
// [NSPredicateEditorRowTemplate] is a concrete class, but it has five
// primitive methods that are called by [NSPredicateEditor]: [NSPredicateEditorRowTemplate.TemplateViews],
// [NSPredicateEditorRowTemplate.MatchForPredicate], [NSPredicateEditorRowTemplate.SetPredicate], [NSPredicateEditorRowTemplate.DisplayableSubpredicatesOfPredicate],
// and [NSPredicateEditorRowTemplate.PredicateWithSubpredicates]. [NSPredicateEditorRowTemplate] implements
// all of these methods, but you can override them for custom templates. The
// primitive methods are used by an instance of [NSPredicateEditor] as
// follows.
//
// First, an instance of [NSPredicateEditor] is created, and some row
// templates are set on it—either through a nib file or programmatically.
// The first thing predicate editor does is ask each of the templates for
// their views, using [NSPredicateEditorRowTemplate.TemplateViews].
//
// After setting up the predicate editor, you typically send it a
// [NSPredicateEditorRowTemplate.ObjectValue] message to restore a saved predicate. [NSPredicateEditor]
// needs to determine which of its templates should display each predicate in
// the predicate tree. It does this by sending each of its row templates a
// [NSPredicateEditorRowTemplate.MatchForPredicate] message and choosing the one that returns the highest
// value.
//
// After finding the best match for a predicate, [NSPredicateEditor] copies
// that template to get fresh views, inserts them into the proper row, and
// then sets the predicate on the template using [NSPredicateEditorRowTemplate.SetPredicate]. Within that
// method, the [NSPredicateEditorRowTemplate] object must set its views’
// values to represent that predicate.
//
// [NSPredicateEditorRowTemplate] next asks the template for the
// “displayable sub-predicates” of the predicate by sending a
// [NSPredicateEditorRowTemplate.DisplayableSubpredicatesOfPredicate] message. If a template represents a
// predicate in its entirety, or if the predicate has no subpredicates, it can
// return `nil` for this. Otherwise, it should return a list of predicates to
// be made into sub-rows of that template’s row. The whole process repeats
// for each sub-predicate.
//
// At this point, the user sees the predicate that was saved. If the user then
// makes some changes to the views of the templates, this causes
// [NSPredicateEditor] to recompute its predicate by asking each of the
// templates to return the predicate represented by the new view values,
// passing in the subpredicates represented by the sub-rows (an empty array if
// there are none, or `nil` if they aren’t supported by that predicate
// type):
//
// [NSPredicateEditorRowTemplate.PredicateWithSubpredicates]
//
// # Initializing a Template
//
//   - [NSPredicateEditorRowTemplate.InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions]: Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
//   - [NSPredicateEditorRowTemplate.InitWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions]: Initializes and returns a “pop-up-pop-up-view”–style row template.
//   - [NSPredicateEditorRowTemplate.InitWithCompoundTypes]: Initializes and returns a row template suitable for displaying compound predicates.
//
// # Primitive Methods
//
//   - [NSPredicateEditorRowTemplate.MatchForPredicate]: Returns a positive number if the receiver can represent a given predicate, and `0` if it cannot.
//   - [NSPredicateEditorRowTemplate.TemplateViews]: Returns the views that display this template’s predicate.
//   - [NSPredicateEditorRowTemplate.SetPredicate]: Sets the value of the views according to the given predicate.
//   - [NSPredicateEditorRowTemplate.DisplayableSubpredicatesOfPredicate]: Returns the subpredicates that should be made sub-rows of a given predicate.
//   - [NSPredicateEditorRowTemplate.PredicateWithSubpredicates]: Returns the predicate represented by the receiver’s views’ values and the given sub-predicates.
//
// # Information About a Row Template
//
//   - [NSPredicateEditorRowTemplate.LeftExpressions]: Returns the left hand expressions for the receiver.
//   - [NSPredicateEditorRowTemplate.RightExpressions]: Returns the right hand expressions for the receiver.
//   - [NSPredicateEditorRowTemplate.CompoundTypes]: Returns the compound predicate types.
//   - [NSPredicateEditorRowTemplate.Modifier]: Returns the comparison predicate modifier for the receiver.
//   - [NSPredicateEditorRowTemplate.Operators]: Returns the array of comparison predicate operators.
//   - [NSPredicateEditorRowTemplate.Options]: Returns the comparison predicate options.
//   - [NSPredicateEditorRowTemplate.RightExpressionAttributeType]: Returns the attribute type of the receiver’s right expression.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate
type NSPredicateEditorRowTemplate struct {
	objectivec.Object
}

// NSPredicateEditorRowTemplateFromID constructs a [NSPredicateEditorRowTemplate] from an objc.ID.
//
// A template that describes available predicates and how to display them.
func NSPredicateEditorRowTemplateFromID(id objc.ID) NSPredicateEditorRowTemplate {
	return NSPredicateEditorRowTemplate{objectivec.Object{ID: id}}
}

// NOTE: NSPredicateEditorRowTemplate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPredicateEditorRowTemplate] class.
//
// # Initializing a Template
//
//   - [INSPredicateEditorRowTemplate.InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions]: Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
//   - [INSPredicateEditorRowTemplate.InitWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions]: Initializes and returns a “pop-up-pop-up-view”–style row template.
//   - [INSPredicateEditorRowTemplate.InitWithCompoundTypes]: Initializes and returns a row template suitable for displaying compound predicates.
//
// # Primitive Methods
//
//   - [INSPredicateEditorRowTemplate.MatchForPredicate]: Returns a positive number if the receiver can represent a given predicate, and `0` if it cannot.
//   - [INSPredicateEditorRowTemplate.TemplateViews]: Returns the views that display this template’s predicate.
//   - [INSPredicateEditorRowTemplate.SetPredicate]: Sets the value of the views according to the given predicate.
//   - [INSPredicateEditorRowTemplate.DisplayableSubpredicatesOfPredicate]: Returns the subpredicates that should be made sub-rows of a given predicate.
//   - [INSPredicateEditorRowTemplate.PredicateWithSubpredicates]: Returns the predicate represented by the receiver’s views’ values and the given sub-predicates.
//
// # Information About a Row Template
//
//   - [INSPredicateEditorRowTemplate.LeftExpressions]: Returns the left hand expressions for the receiver.
//   - [INSPredicateEditorRowTemplate.RightExpressions]: Returns the right hand expressions for the receiver.
//   - [INSPredicateEditorRowTemplate.CompoundTypes]: Returns the compound predicate types.
//   - [INSPredicateEditorRowTemplate.Modifier]: Returns the comparison predicate modifier for the receiver.
//   - [INSPredicateEditorRowTemplate.Operators]: Returns the array of comparison predicate operators.
//   - [INSPredicateEditorRowTemplate.Options]: Returns the comparison predicate options.
//   - [INSPredicateEditorRowTemplate.RightExpressionAttributeType]: Returns the attribute type of the receiver’s right expression.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate
type INSPredicateEditorRowTemplate interface {
	objectivec.IObject

	// Topic: Initializing a Template

	// Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
	InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.NSExpression, rightExpressions []foundation.NSExpression, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate
	// Initializes and returns a “pop-up-pop-up-view”–style row template.
	InitWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions(leftExpressions []foundation.NSExpression, attributeType objectivec.IObject, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate
	// Initializes and returns a row template suitable for displaying compound predicates.
	InitWithCompoundTypes(compoundTypes []foundation.NSNumber) NSPredicateEditorRowTemplate

	// Topic: Primitive Methods

	// Returns a positive number if the receiver can represent a given predicate, and `0` if it cannot.
	MatchForPredicate(predicate foundation.INSPredicate) float64
	// Returns the views that display this template’s predicate.
	TemplateViews() []NSView
	// Sets the value of the views according to the given predicate.
	SetPredicate(predicate foundation.INSPredicate)
	// Returns the subpredicates that should be made sub-rows of a given predicate.
	DisplayableSubpredicatesOfPredicate(predicate foundation.INSPredicate) []foundation.NSPredicate
	// Returns the predicate represented by the receiver’s views’ values and the given sub-predicates.
	PredicateWithSubpredicates(subpredicates []foundation.NSPredicate) foundation.INSPredicate

	// Topic: Information About a Row Template

	// Returns the left hand expressions for the receiver.
	LeftExpressions() []foundation.NSExpression
	// Returns the right hand expressions for the receiver.
	RightExpressions() []foundation.NSExpression
	// Returns the compound predicate types.
	CompoundTypes() []foundation.NSNumber
	// Returns the comparison predicate modifier for the receiver.
	Modifier() foundation.NSComparisonPredicateModifier
	// Returns the array of comparison predicate operators.
	Operators() []foundation.NSNumber
	// Returns the comparison predicate options.
	Options() uint
	// Returns the attribute type of the receiver’s right expression.
	RightExpressionAttributeType() objectivec.IObject

	// The value of the receiver’s cell as an Objective-C object.
	ObjectValue() objectivec.IObject
	SetObjectValue(value objectivec.IObject)
	// The row templates for the receiver.
	RowTemplates() INSPredicateEditorRowTemplate
	SetRowTemplates(value INSPredicateEditorRowTemplate)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPredicateEditorRowTemplate) Init() NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPredicateEditorRowTemplate) Autorelease() NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPredicateEditorRowTemplate creates a new NSPredicateEditorRowTemplate instance.
func NewNSPredicateEditorRowTemplate() NSPredicateEditorRowTemplate {
	class := getNSPredicateEditorRowTemplateClass()
	rv := objc.Send[NSPredicateEditorRowTemplate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a row template suitable for displaying compound
// predicates.
//
// compoundTypes: An array of [NSNumber] objects specifying compound predicate types. See
// [NSCompoundPredicate.LogicalType] for possible values.
//
// # Return Value
//
// A row template initialized for displaying compound predicates of the types
// specified by `compoundTypes`.
//
// # Discussion
//
// [NSPredicateEditor] contains such a template by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(compoundTypes:)
//
// [NSCompoundPredicate.LogicalType]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func NewPredicateEditorRowTemplateWithCompoundTypes(compoundTypes []foundation.NSNumber) NSPredicateEditorRowTemplate {
	instance := getNSPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompoundTypes:"), objectivec.IObjectSliceToNSArray(compoundTypes))
	return NSPredicateEditorRowTemplateFromID(rv)
}

// Initializes and returns a “pop-up-pop-up-view”–style row template.
//
// leftExpressions: An array of [NSExpression] objects that represent the left side of a
// predicate.
//
// attributeType: An attribute type for the right side of a predicate. This value dictates
// the type of view created, and how the control’s object value is coerced
// before putting it into a predicate.
//
// modifier: A modifier for the predicate (see [NSComparisonPredicate.Modifier] for
// possible values).
//
// operators: An array of [NSNumber] objects specifying the operator type (see
// [NSComparisonPredicate.Operator] for possible values).
//
// options: Options for the predicate (see [NSComparisonPredicate.Options] for possible
// values).
//
// # Return Value
//
// A row template initialized using the specified arguments.
//
// # Discussion
//
// The type of `attributeType` dictates the type of view created. For example,
// [NSAttributeType.dateAttributeType] creates an [NSDatePicker] object,
// [NSAttributeType.integer64AttributeType] creates a short text field, and
// [NSAttributeType.stringAttributeType] produces a longer text field. You can
// resize the views as you want.
//
// Predicates do not automatically coerce types for you. For example,
// comparing a number to a string will raise an exception. Therefore, the
// attribute type is also needed to determine how the control’s object value
// must be coerced before putting it into a predicate.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressionAttributeType:modifier:operators:options:)
// attributeType is a [coredata.NSAttributeType].
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
// [NSComparisonPredicate.Modifier]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
// [NSComparisonPredicate.Operator]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
// [NSAttributeType.dateAttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/dateAttributeType
// [NSAttributeType.integer64AttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/integer64AttributeType
// [NSAttributeType.stringAttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/stringAttributeType
func NewPredicateEditorRowTemplateWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions(leftExpressions []foundation.NSExpression, attributeType objectivec.IObject, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate {
	instance := getNSPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftExpressions:rightExpressionAttributeType:modifier:operators:options:"), objectivec.IObjectSliceToNSArray(leftExpressions), attributeType, modifier, objectivec.IObjectSliceToNSArray(operators), options)
	return NSPredicateEditorRowTemplateFromID(rv)
}

// Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
//
// leftExpressions: An array of [NSExpression] objects that represent the left side of a
// predicate.
//
// rightExpressions: An array of [NSExpression] objects that represent the right side of a
// predicate.
//
// modifier: A modifier for the predicate (see [NSComparisonPredicate.Modifier] for
// possible values).
//
// operators: An array of [NSNumber] objects specifying the operator type (see
// [NSComparisonPredicate.Operator] for possible values).
//
// options: Options for the predicate (see [NSComparisonPredicate.Options] for possible
// values).
//
// # Return Value
//
// A row template of the “pop-up-pop-up-pop-up” form, with the left and
// right pop-ups representing the left and right expression arrays
// `leftExpressions` and `rightExpressions`, and the center pop-up
// representing the operators.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressions:modifier:operators:options:)
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
// [NSComparisonPredicate.Modifier]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
// [NSComparisonPredicate.Operator]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
func NewPredicateEditorRowTemplateWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.NSExpression, rightExpressions []foundation.NSExpression, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate {
	instance := getNSPredicateEditorRowTemplateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLeftExpressions:rightExpressions:modifier:operators:options:"), objectivec.IObjectSliceToNSArray(leftExpressions), objectivec.IObjectSliceToNSArray(rightExpressions), modifier, objectivec.IObjectSliceToNSArray(operators), options)
	return NSPredicateEditorRowTemplateFromID(rv)
}

// Initializes and returns a “pop-up-pop-up-pop-up”–style row template.
//
// leftExpressions: An array of [NSExpression] objects that represent the left side of a
// predicate.
//
// rightExpressions: An array of [NSExpression] objects that represent the right side of a
// predicate.
//
// modifier: A modifier for the predicate (see [NSComparisonPredicate.Modifier] for
// possible values).
//
// operators: An array of [NSNumber] objects specifying the operator type (see
// [NSComparisonPredicate.Operator] for possible values).
//
// options: Options for the predicate (see [NSComparisonPredicate.Options] for possible
// values).
//
// # Return Value
//
// A row template of the “pop-up-pop-up-pop-up” form, with the left and
// right pop-ups representing the left and right expression arrays
// `leftExpressions` and `rightExpressions`, and the center pop-up
// representing the operators.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressions:modifier:operators:options:)
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
// [NSComparisonPredicate.Modifier]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
// [NSComparisonPredicate.Operator]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
func (p NSPredicateEditorRowTemplate) InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions(leftExpressions []foundation.NSExpression, rightExpressions []foundation.NSExpression, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p.ID, objc.Sel("initWithLeftExpressions:rightExpressions:modifier:operators:options:"), objectivec.IObjectSliceToNSArray(leftExpressions), objectivec.IObjectSliceToNSArray(rightExpressions), modifier, objectivec.IObjectSliceToNSArray(operators), options)
	return rv
}

// Initializes and returns a “pop-up-pop-up-view”–style row template.
//
// leftExpressions: An array of [NSExpression] objects that represent the left side of a
// predicate.
//
// attributeType: An attribute type for the right side of a predicate. This value dictates
// the type of view created, and how the control’s object value is coerced
// before putting it into a predicate.
//
// modifier: A modifier for the predicate (see [NSComparisonPredicate.Modifier] for
// possible values).
//
// operators: An array of [NSNumber] objects specifying the operator type (see
// [NSComparisonPredicate.Operator] for possible values).
//
// options: Options for the predicate (see [NSComparisonPredicate.Options] for possible
// values).
//
// attributeType is a [coredata.NSAttributeType].
//
// # Return Value
//
// A row template initialized using the specified arguments.
//
// # Discussion
//
// The type of `attributeType` dictates the type of view created. For example,
// [NSAttributeType.dateAttributeType] creates an [NSDatePicker] object,
// [NSAttributeType.integer64AttributeType] creates a short text field, and
// [NSAttributeType.stringAttributeType] produces a longer text field. You can
// resize the views as you want.
//
// Predicates do not automatically coerce types for you. For example,
// comparing a number to a string will raise an exception. Therefore, the
// attribute type is also needed to determine how the control’s object value
// must be coerced before putting it into a predicate.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(leftExpressions:rightExpressionAttributeType:modifier:operators:options:)
// attributeType is a [coredata.NSAttributeType].
//
// [NSExpression]: https://developer.apple.com/documentation/Foundation/NSExpression
// [NSComparisonPredicate.Modifier]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
// [NSComparisonPredicate.Operator]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
// [NSAttributeType.dateAttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/dateAttributeType
// [NSAttributeType.integer64AttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/integer64AttributeType
// [NSAttributeType.stringAttributeType]: https://developer.apple.com/documentation/CoreData/NSAttributeType/stringAttributeType
func (p NSPredicateEditorRowTemplate) InitWithLeftExpressionsRightExpressionAttributeTypeModifierOperatorsOptions(leftExpressions []foundation.NSExpression, attributeType objectivec.IObject, modifier foundation.NSComparisonPredicateModifier, operators []foundation.NSNumber, options uint) NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p.ID, objc.Sel("initWithLeftExpressions:rightExpressionAttributeType:modifier:operators:options:"), objectivec.IObjectSliceToNSArray(leftExpressions), attributeType, modifier, objectivec.IObjectSliceToNSArray(operators), options)
	return rv
}

// Initializes and returns a row template suitable for displaying compound
// predicates.
//
// compoundTypes: An array of [NSNumber] objects specifying compound predicate types. See
// [NSCompoundPredicate.LogicalType] for possible values.
//
// # Return Value
//
// A row template initialized for displaying compound predicates of the types
// specified by `compoundTypes`.
//
// # Discussion
//
// [NSPredicateEditor] contains such a template by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/init(compoundTypes:)
//
// [NSCompoundPredicate.LogicalType]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (p NSPredicateEditorRowTemplate) InitWithCompoundTypes(compoundTypes []foundation.NSNumber) NSPredicateEditorRowTemplate {
	rv := objc.Send[NSPredicateEditorRowTemplate](p.ID, objc.Sel("initWithCompoundTypes:"), objectivec.IObjectSliceToNSArray(compoundTypes))
	return rv
}

// Returns a positive number if the receiver can represent a given predicate,
// and `0` if it cannot.
//
// # Return Value
//
// A positive number if the template can represent `predicate`, and `0` if it
// cannot.
//
// # Discussion
//
// By default, returns values in the range `0` to `1`.
//
// The highest match among all the templates determines which template is
// responsible for displaying the predicate. You can override this to
// determine which predicates your custom template handles.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/match(for:)
func (p NSPredicateEditorRowTemplate) MatchForPredicate(predicate foundation.INSPredicate) float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("matchForPredicate:"), predicate)
	return rv
}

// Sets the value of the views according to the given predicate.
//
// predicate: The predicate value for the receiver.
//
// # Discussion
//
// This method is only called if [MatchForPredicate] returned a positive value
// for the receiver.
//
// You can override this to set the values of custom views.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/setPredicate(_:)
func (p NSPredicateEditorRowTemplate) SetPredicate(predicate foundation.INSPredicate) {
	objc.Send[objc.ID](p.ID, objc.Sel("setPredicate:"), predicate)
}

// Returns the subpredicates that should be made sub-rows of a given
// predicate.
//
// predicate: A predicate object.
//
// # Return Value
//
// The subpredicates that should be made sub-rows of `predicate`. For compound
// predicates (instances of [NSCompoundPredicate]), the array of
// subpredicates; for other types of predicate, returns `nil`. If a template
// represents a predicate in its entirety, or if the predicate has no
// subpredicates, returns `nil`.
//
// # Discussion
//
// You can override this method to create custom templates that handle
// complicated compound predicates.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/displayableSubpredicates(of:)
//
// [NSCompoundPredicate]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate
func (p NSPredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.INSPredicate) []foundation.NSPredicate {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("displayableSubpredicatesOfPredicate:"), predicate)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSPredicate {
		return foundation.NSPredicateFromID(id)
	})
}

// Returns the predicate represented by the receiver’s views’ values and
// the given sub-predicates.
//
// subpredicates: An array of predicates.
//
// # Return Value
//
// The predicate represented by the values of the template’s views and the
// given subpredicates. You can override this method to return the predicate
// represented by your custom views.
//
// # Discussion
//
// This method is only called if [MatchForPredicate] returned a positive value
// for the receiver.
//
// You can override this method to return the predicate represented by a
// custom view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/predicate(withSubpredicates:)
func (p NSPredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.NSPredicate) foundation.INSPredicate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("predicateWithSubpredicates:"), objectivec.IObjectSliceToNSArray(subpredicates))
	return foundation.NSPredicateFromID(rv)
}
func (p NSPredicateEditorRowTemplate) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns an array of predicate templates for the given attribute key paths
// for a given entity.
//
// keyPaths: An array of attribute key paths originating at `entityDescription`. The key
// paths may cross relationships but must terminate in attributes.
//
// entityDescription: A Core Data entity description.
//
// entityDescription is a [coredata.NSEntityDescription].
//
// # Return Value
//
// An array of predicate templates for `keyPaths` originating at
// `entityDescription`.
//
// # Discussion
//
// This method determines which key paths in the entity description can use
// the same views (that is, share the same attribute type). For each of these
// groups, it instantiates individual templates via
// [InitWithLeftExpressionsRightExpressionsModifierOperatorsOptions].
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/templates(withAttributeKeyPaths:in:)
// entityDescription is a [coredata.NSEntityDescription].
func (_NSPredicateEditorRowTemplateClass NSPredicateEditorRowTemplateClass) TemplatesWithAttributeKeyPathsInEntityDescription(keyPaths []string, entityDescription objectivec.IObject) []NSPredicateEditorRowTemplate {
	rv := objc.Send[[]objc.ID](objc.ID(_NSPredicateEditorRowTemplateClass.class), objc.Sel("templatesWithAttributeKeyPaths:inEntityDescription:"), objectivec.StringSliceToNSArray(keyPaths), entityDescription)
	return objc.ConvertSlice(rv, func(id objc.ID) NSPredicateEditorRowTemplate {
		return NSPredicateEditorRowTemplateFromID(id)
	})
}

// Returns the views that display this template’s predicate.
//
// # Return Value
//
// The views for an [NSPredicateEditor] to display in a row that represents
// the predicate from [SetPredicate].
//
// # Discussion
//
// [NSPredicateEditor] treats instances of [NSPopUpButton] specially by
// merging their menu items into a single popup button, and by combining menu
// items with matching titles. In this way, the editor builds a single tree
// from the separate templates.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/templateViews
func (p NSPredicateEditorRowTemplate) TemplateViews() []NSView {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("templateViews"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// Returns the left hand expressions for the receiver.
//
// # Return Value
//
// # The left hand expressions for the receiver
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/leftExpressions
func (p NSPredicateEditorRowTemplate) LeftExpressions() []foundation.NSExpression {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("leftExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSExpression {
		return foundation.NSExpressionFromID(id)
	})
}

// Returns the right hand expressions for the receiver.
//
// # Return Value
//
// # The right hand expressions for the receiver
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/rightExpressions
func (p NSPredicateEditorRowTemplate) RightExpressions() []foundation.NSExpression {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("rightExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSExpression {
		return foundation.NSExpressionFromID(id)
	})
}

// Returns the compound predicate types.
//
// # Return Value
//
// An array of [NSNumber] objects specifying compound predicate types. See
// [NSCompoundPredicate.LogicalType] for possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/compoundTypes
//
// [NSCompoundPredicate.LogicalType]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (p NSPredicateEditorRowTemplate) CompoundTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("compoundTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the comparison predicate modifier for the receiver.
//
// # Return Value
//
// The comparison predicate modifier for the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/modifier
func (p NSPredicateEditorRowTemplate) Modifier() foundation.NSComparisonPredicateModifier {
	rv := objc.Send[foundation.NSComparisonPredicateModifier](p.ID, objc.Sel("modifier"))
	return foundation.NSComparisonPredicateModifier(rv)
}

// Returns the array of comparison predicate operators.
//
// # Return Value
//
// The array of [NSNumber] objects specifying the comparison predicate
// operators. See [NSComparisonPredicate.Operator] for possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/operators
//
// [NSComparisonPredicate.Operator]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (p NSPredicateEditorRowTemplate) Operators() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("operators"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the comparison predicate options.
//
// # Return Value
//
// The comparison predicate options for the receiver. See
// [NSComparisonPredicate.Options] for possible values. Returns `0` if this
// does not apply (for example, for a compound template initialized with
// [InitWithCompoundTypes]).
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/options
//
// [NSComparisonPredicate.Options]: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
func (p NSPredicateEditorRowTemplate) Options() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("options"))
	return rv
}

// Returns the attribute type of the receiver’s right expression.
//
// # Return Value
//
// The attribute type of the receiver’s right expression.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditorRowTemplate/rightExpressionAttributeType
func (p NSPredicateEditorRowTemplate) RightExpressionAttributeType() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("rightExpressionAttributeType"))
	return objectivec.Object{ID: rv}
}

// The value of the receiver’s cell as an Objective-C object.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/objectvalue
func (p NSPredicateEditorRowTemplate) ObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("objectValue"))
	return objectivec.Object{ID: rv}
}
func (p NSPredicateEditorRowTemplate) SetObjectValue(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setObjectValue:"), value)
}

// The row templates for the receiver.
//
// See: https://developer.apple.com/documentation/appkit/nspredicateeditor/rowtemplates
func (p NSPredicateEditorRowTemplate) RowTemplates() INSPredicateEditorRowTemplate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("rowTemplates"))
	return NSPredicateEditorRowTemplateFromID(objc.ID(rv))
}
func (p NSPredicateEditorRowTemplate) SetRowTemplates(value INSPredicateEditorRowTemplate) {
	objc.Send[struct{}](p.ID, objc.Sel("setRowTemplates:"), value)
}
