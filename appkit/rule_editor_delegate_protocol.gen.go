// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The [NSRuleEditorDelegate] protocol defines the optional methods implemented by delegates of [NSRuleEditor](<doc://com.apple.appkit/documentation/AppKit/NSRuleEditor>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate
type NSRuleEditorDelegate interface {
	objectivec.IObject
}

// NSRuleEditorDelegateObject wraps an existing Objective-C object that conforms to the NSRuleEditorDelegate protocol.
type NSRuleEditorDelegateObject struct {
	objectivec.Object
}
func (o NSRuleEditorDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSRuleEditorDelegateObjectFromID constructs a [NSRuleEditorDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSRuleEditorDelegateObjectFromID(id objc.ID) NSRuleEditorDelegateObject {
	return NSRuleEditorDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the child of a given item at a given index.
//
// editor: The rule editor that sent the message.
//
// index: The index of the requested child criterion. This value must be in the range
// from `0` up to (but not including) the number of children, as reported by
// the delegate in [RuleEditorNumberOfChildrenForCriterionWithRowType].
//
// criterion: The parent of the requested child, or `nil` if the rule editor is
// requesting a root criterion.
//
// rowType: The type of the row.
//
// # Return Value
// 
// An object representing the requested child (or root) criterion. This object
// is used by the delegate to represent that position in the tree, and is
// passed as a parameter in subsequent calls to the delegate.
//
// # Discussion
// 
// The delegate must implement this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate/ruleEditor(_:child:forCriterion:with:)

func (o NSRuleEditorDelegateObject) RuleEditorChildForCriterionWithRowType(editor INSRuleEditor, index int, criterion objectivec.IObject, rowType NSRuleEditorRowType) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("ruleEditor:child:forCriterion:withRowType:"), editor, index, criterion, rowType)
	return objectivec.Object{ID: rv}
	}

// Returns the value for a given criterion.
//
// editor: The rule editor that sent the message.
//
// criterion: The criterion for which the value is required.
//
// row: The row number of `criterion`.
//
// # Return Value
// 
// The value for `criterion`.
//
// # Discussion
// 
// The value should be an instance of [NSString], [NSView], or [NSMenuItem].
// If the value is an [NSView] or [NSMenuItem], you must ensure it is unique
// for every invocation of this method; that is, do not return a particular
// instance of [NSView] or [NSMenuItem] more than once.
// 
// # Special Considerations
// 
// The delegate must implement this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate/ruleEditor(_:displayValueForCriterion:inRow:)

func (o NSRuleEditorDelegateObject) RuleEditorDisplayValueForCriterionInRow(editor INSRuleEditor, criterion objectivec.IObject, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("ruleEditor:displayValueForCriterion:inRow:"), editor, criterion, row)
	return objectivec.Object{ID: rv}
	}

// Returns the number of child items of a given criterion or row type.
//
// editor: The rule editor that sent the message.
//
// criterion: The criterion for which the number of children is required.
//
// rowType: The type of row of `criterion`.
//
// # Return Value
// 
// The number of child items of `criterion`. If `criterion` is `nil`, return
// the number of root criteria for the row type `rowType`.
//
// # Discussion
// 
// The delegate must implement this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate/ruleEditor(_:numberOfChildrenForCriterion:with:)

func (o NSRuleEditorDelegateObject) RuleEditorNumberOfChildrenForCriterionWithRowType(editor INSRuleEditor, criterion objectivec.IObject, rowType NSRuleEditorRowType) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"), editor, criterion, rowType)
	return rv
	}

// Returns a dictionary representing the parts of the predicate determined by
// the given criterion and value.
//
// editor: The rule editor that sent the message.
//
// criterion: The criterion for which the predicate parts are required.
//
// value: The display value.
//
// row: The row number of `criterion`.
//
// # Return Value
// 
// A dictionary representing the parts of the predicate determined by the
// given criterion and value. The keys of the dictionary should be the string
// constants specified in Predicate Part Keys with corresponding appropriate
// values.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate/ruleEditor(_:predicatePartsForCriterion:withDisplayValue:inRow:)

func (o NSRuleEditorDelegateObject) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor INSRuleEditor, criterion objectivec.IObject, value objectivec.IObject, row int) foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), editor, criterion, value, row)
	return foundation.NSDictionaryFromID(rv)
	}

// Notifies the receiver that a rule editor’s rows changed.
//
// notification: A notification named[rowsDidChangeNotification].
// //
// [rowsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowsDidChangeNotification
//
// # Discussion
// 
// If the delegate implements this method, [NSRuleEditor] automatically
// registers its delegate to receive [rowsDidChangeNotification]
// notifications.
//
// [rowsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowsDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditorDelegate/ruleEditorRowsDidChange(_:)

func (o NSRuleEditorDelegateObject) RuleEditorRowsDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("ruleEditorRowsDidChange:"), notification)
	}

// NSRuleEditorDelegateConfig holds optional typed callbacks for [NSRuleEditorDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate
type NSRuleEditorDelegateConfig struct {

	// Monitoring Row Changes
	// RuleEditorRowsDidChange — Notifies the receiver that a rule editor’s rows changed.
	RuleEditorRowsDidChange func(notification foundation.NSNotification)
}

// NewNSRuleEditorDelegate creates an Objective-C object implementing the [NSRuleEditorDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSRuleEditorDelegateObject] satisfies the [NSRuleEditorDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate
func NewNSRuleEditorDelegate(config NSRuleEditorDelegateConfig) NSRuleEditorDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSRuleEditorDelegate_%d", n)

	var methods []objc.MethodDef

	if config.RuleEditorRowsDidChange != nil {
		fn := config.RuleEditorRowsDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("ruleEditorRowsDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSRuleEditorDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSRuleEditorDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSRuleEditorDelegateObjectFromID(instance)
}

