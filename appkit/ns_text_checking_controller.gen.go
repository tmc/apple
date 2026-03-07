// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextCheckingController] class.
var (
	_NSTextCheckingControllerClass     NSTextCheckingControllerClass
	_NSTextCheckingControllerClassOnce sync.Once
)

func getNSTextCheckingControllerClass() NSTextCheckingControllerClass {
	_NSTextCheckingControllerClassOnce.Do(func() {
		_NSTextCheckingControllerClass = NSTextCheckingControllerClass{class: objc.GetClass("NSTextCheckingController")}
	})
	return _NSTextCheckingControllerClass
}

// GetNSTextCheckingControllerClass returns the class object for NSTextCheckingController.
func GetNSTextCheckingControllerClass() NSTextCheckingControllerClass {
	return getNSTextCheckingControllerClass()
}

type NSTextCheckingControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextCheckingControllerClass) Alloc() NSTextCheckingController {
	rv := objc.Send[NSTextCheckingController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







//
// # Initializers
//
//   - [NSTextCheckingController.InitWithClient]
//
// # Instance Properties
//
//   - [NSTextCheckingController.Client]
//   - [NSTextCheckingController.SpellCheckerDocumentTag]
//   - [NSTextCheckingController.SetSpellCheckerDocumentTag]
//
// # Instance Methods
//
//   - [NSTextCheckingController.ChangeSpelling]
//   - [NSTextCheckingController.CheckSpelling]
//   - [NSTextCheckingController.CheckTextInRangeTypesOptions]
//   - [NSTextCheckingController.CheckTextInDocument]
//   - [NSTextCheckingController.CheckTextInSelection]
//   - [NSTextCheckingController.ConsiderTextCheckingForRange]
//   - [NSTextCheckingController.DidChangeSelectedRange]
//   - [NSTextCheckingController.DidChangeTextInRange]
//   - [NSTextCheckingController.IgnoreSpelling]
//   - [NSTextCheckingController.InsertedTextInRange]
//   - [NSTextCheckingController.Invalidate]
//   - [NSTextCheckingController.MenuAtIndexClickedOnSelectionEffectiveRange]
//   - [NSTextCheckingController.OrderFrontSubstitutionsPanel]
//   - [NSTextCheckingController.ShowGuessPanel]
//   - [NSTextCheckingController.UpdateCandidates]
//   - [NSTextCheckingController.ValidAnnotations]
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController
type NSTextCheckingController struct {
	objectivec.Object
}

// NSTextCheckingControllerFromID constructs a [NSTextCheckingController] from an objc.ID.
func NSTextCheckingControllerFromID(id objc.ID) NSTextCheckingController {
	return NSTextCheckingController{objectivec.Object{id}}
}
// NOTE: NSTextCheckingController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextCheckingController] class.
//
// # Initializers
//
//   - [INSTextCheckingController.InitWithClient]
//
// # Instance Properties
//
//   - [INSTextCheckingController.Client]
//   - [INSTextCheckingController.SpellCheckerDocumentTag]
//   - [INSTextCheckingController.SetSpellCheckerDocumentTag]
//
// # Instance Methods
//
//   - [INSTextCheckingController.ChangeSpelling]
//   - [INSTextCheckingController.CheckSpelling]
//   - [INSTextCheckingController.CheckTextInRangeTypesOptions]
//   - [INSTextCheckingController.CheckTextInDocument]
//   - [INSTextCheckingController.CheckTextInSelection]
//   - [INSTextCheckingController.ConsiderTextCheckingForRange]
//   - [INSTextCheckingController.DidChangeSelectedRange]
//   - [INSTextCheckingController.DidChangeTextInRange]
//   - [INSTextCheckingController.IgnoreSpelling]
//   - [INSTextCheckingController.InsertedTextInRange]
//   - [INSTextCheckingController.Invalidate]
//   - [INSTextCheckingController.MenuAtIndexClickedOnSelectionEffectiveRange]
//   - [INSTextCheckingController.OrderFrontSubstitutionsPanel]
//   - [INSTextCheckingController.ShowGuessPanel]
//   - [INSTextCheckingController.UpdateCandidates]
//   - [INSTextCheckingController.ValidAnnotations]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController
type INSTextCheckingController interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithClient(client NSTextCheckingClient) NSTextCheckingController

	// Topic: Instance Properties

	Client() NSTextCheckingClient
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)

	// Topic: Instance Methods

	ChangeSpelling(sender objectivec.IObject)
	CheckSpelling(sender objectivec.IObject)
	CheckTextInRangeTypesOptions(range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary)
	CheckTextInDocument(sender objectivec.IObject)
	CheckTextInSelection(sender objectivec.IObject)
	ConsiderTextCheckingForRange(range_ foundation.NSRange)
	DidChangeSelectedRange()
	DidChangeTextInRange(range_ foundation.NSRange)
	IgnoreSpelling(sender objectivec.IObject)
	InsertedTextInRange(range_ foundation.NSRange)
	Invalidate()
	MenuAtIndexClickedOnSelectionEffectiveRange(location uint, clickedOnSelection bool, effectiveRange foundation.NSRange) INSMenu
	OrderFrontSubstitutionsPanel(sender objectivec.IObject)
	ShowGuessPanel(sender objectivec.IObject)
	UpdateCandidates()
	ValidAnnotations() []string
}





// Init initializes the instance.
func (t NSTextCheckingController) Init() NSTextCheckingController {
	rv := objc.Send[NSTextCheckingController](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextCheckingController) Autorelease() NSTextCheckingController {
	rv := objc.Send[NSTextCheckingController](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextCheckingController creates a new NSTextCheckingController instance.
func NewNSTextCheckingController() NSTextCheckingController {
	class := getNSTextCheckingControllerClass()
	rv := objc.Send[NSTextCheckingController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/init(client:)
func NewTextCheckingControllerWithClient(client NSTextCheckingClient) NSTextCheckingController {
	instance := getNSTextCheckingControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClient:"), client)
	return NSTextCheckingControllerFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/init(client:)
func (t NSTextCheckingController) InitWithClient(client NSTextCheckingClient) NSTextCheckingController {
	rv := objc.Send[NSTextCheckingController](t.ID, objc.Sel("initWithClient:"), client)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/changeSpelling(_:)
func (t NSTextCheckingController) ChangeSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeSpelling:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/checkSpelling(_:)
func (t NSTextCheckingController) CheckSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkSpelling:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/checkText(in:types:options:)
func (t NSTextCheckingController) CheckTextInRangeTypesOptions(range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/checkTextInDocument(_:)
func (t NSTextCheckingController) CheckTextInDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInDocument:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/checkTextInSelection(_:)
func (t NSTextCheckingController) CheckTextInSelection(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInSelection:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/considerTextChecking(for:)
func (t NSTextCheckingController) ConsiderTextCheckingForRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("considerTextCheckingForRange:"), range_)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/didChangeSelectedRange()
func (t NSTextCheckingController) DidChangeSelectedRange() {
	objc.Send[objc.ID](t.ID, objc.Sel("didChangeSelectedRange"))
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/didChangeText(in:)
func (t NSTextCheckingController) DidChangeTextInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("didChangeTextInRange:"), range_)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/ignoreSpelling(_:)
func (t NSTextCheckingController) IgnoreSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("ignoreSpelling:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/insertedText(in:)
func (t NSTextCheckingController) InsertedTextInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertedTextInRange:"), range_)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/invalidate()
func (t NSTextCheckingController) Invalidate() {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidate"))
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/menu(at:clickedOnSelection:effectiveRange:)
func (t NSTextCheckingController) MenuAtIndexClickedOnSelectionEffectiveRange(location uint, clickedOnSelection bool, effectiveRange foundation.NSRange) INSMenu {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("menuAtIndex:clickedOnSelection:effectiveRange:"), location, clickedOnSelection, effectiveRange)
	return NSMenuFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/orderFrontSubstitutionsPanel(_:)
func (t NSTextCheckingController) OrderFrontSubstitutionsPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/showGuessPanel(_:)
func (t NSTextCheckingController) ShowGuessPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("showGuessPanel:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/updateCandidates()
func (t NSTextCheckingController) UpdateCandidates() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateCandidates"))
}

// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/validAnnotations()
func (t NSTextCheckingController) ValidAnnotations() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("validAnnotations"))
	return objc.ConvertSliceToStrings(rv)
}












// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/client
func (t NSTextCheckingController) Client() NSTextCheckingClient {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("client"))
	return NSTextCheckingClientObjectFromID(rv)
}



// See: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/spellCheckerDocumentTag
func (t NSTextCheckingController) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}
func (t NSTextCheckingController) SetSpellCheckerDocumentTag(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}























