// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPredicateEditor] class.
var (
	_NSPredicateEditorClass     NSPredicateEditorClass
	_NSPredicateEditorClassOnce sync.Once
)

func getNSPredicateEditorClass() NSPredicateEditorClass {
	_NSPredicateEditorClassOnce.Do(func() {
		_NSPredicateEditorClass = NSPredicateEditorClass{class: objc.GetClass("NSPredicateEditor")}
	})
	return _NSPredicateEditorClass
}

// GetNSPredicateEditorClass returns the class object for NSPredicateEditor.
func GetNSPredicateEditorClass() NSPredicateEditorClass {
	return getNSPredicateEditorClass()
}

type NSPredicateEditorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPredicateEditorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPredicateEditorClass) Alloc() NSPredicateEditor {
	rv := objc.Send[NSPredicateEditor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A defined set of rules that allows the editing of predicate objects.
//
// # Overview
// 
// [NSPredicateEditor] provides an [NSPredicate] property—[NSPredicateEditor.ObjectValue]
// (inherited from [NSControl])—that you can get and set directly, and that
// you can bind using Cocoa bindings (you typically configure a predicate
// editor in Interface Builder). [NSPredicateEditor] depends on another class,
// [NSPredicateEditorRowTemplate], that describes the available predicates and
// how to display them.
// 
// Unlike [NSRuleEditor], [NSPredicateEditor] does not depend on its delegate
// to populate its rows (and ). Instead, its rows are populated from its
// `objectValue` property (an instance of [NSPredicate]). [NSPredicateEditor]
// relies on instances [NSPredicateEditorRowTemplate], which are responsible
// for mapping back and forth between the displayed view values and various
// predicates.
// 
// [NSPredicateEditor] exposes one property, [NSPredicateEditor.RowTemplates], which is an array
// of [NSPredicateEditorRowTemplate] objects.
//
// [NSPredicate]: https://developer.apple.com/documentation/Foundation/NSPredicate
//
// # Managing Row Templates
//
//   - [NSPredicateEditor.RowTemplates]: The row templates for the receiver.
//   - [NSPredicateEditor.SetRowTemplates]
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditor
type NSPredicateEditor struct {
	NSRuleEditor
}

// NSPredicateEditorFromID constructs a [NSPredicateEditor] from an objc.ID.
//
// A defined set of rules that allows the editing of predicate objects.
func NSPredicateEditorFromID(id objc.ID) NSPredicateEditor {
	return NSPredicateEditor{NSRuleEditor: NSRuleEditorFromID(id)}
}
// NOTE: NSPredicateEditor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPredicateEditor] class.
//
// # Managing Row Templates
//
//   - [INSPredicateEditor.RowTemplates]: The row templates for the receiver.
//   - [INSPredicateEditor.SetRowTemplates]
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditor
type INSPredicateEditor interface {
	INSRuleEditor

	// Topic: Managing Row Templates

	// The row templates for the receiver.
	RowTemplates() []NSPredicateEditorRowTemplate
	SetRowTemplates(value []NSPredicateEditorRowTemplate)
}

// Init initializes the instance.
func (p NSPredicateEditor) Init() NSPredicateEditor {
	rv := objc.Send[NSPredicateEditor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPredicateEditor) Autorelease() NSPredicateEditor {
	rv := objc.Send[NSPredicateEditor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPredicateEditor creates a new NSPredicateEditor instance.
func NewNSPredicateEditor() NSPredicateEditor {
	class := getNSPredicateEditorClass()
	rv := objc.Send[NSPredicateEditor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewPredicateEditorWithCoder(coder foundation.INSCoder) NSPredicateEditor {
	instance := getNSPredicateEditorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPredicateEditorFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewPredicateEditorWithFrame(frameRect corefoundation.CGRect) NSPredicateEditor {
	instance := getNSPredicateEditorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSPredicateEditorFromID(rv)
}

// The row templates for the receiver.
//
// # Discussion
// 
// The default value is a single compound [NSPredicateEditorRowTemplate]
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPredicateEditor/rowTemplates
func (p NSPredicateEditor) RowTemplates() []NSPredicateEditorRowTemplate {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("rowTemplates"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPredicateEditorRowTemplate {
		return NSPredicateEditorRowTemplateFromID(id)
	})
}
func (p NSPredicateEditor) SetRowTemplates(value []NSPredicateEditorRowTemplate) {
	objc.Send[struct{}](p.ID, objc.Sel("setRowTemplates:"), objectivec.IObjectSliceToNSArray(value))
}

