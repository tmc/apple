// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IMKCandidates] class.
var (
	_IMKCandidatesClass     IMKCandidatesClass
	_IMKCandidatesClassOnce sync.Once
)

func getIMKCandidatesClass() IMKCandidatesClass {
	_IMKCandidatesClassOnce.Do(func() {
		_IMKCandidatesClass = IMKCandidatesClass{class: objc.GetClass("IMKCandidates")}
	})
	return _IMKCandidatesClass
}

// GetIMKCandidatesClass returns the class object for IMKCandidates.
func GetIMKCandidatesClass() IMKCandidatesClass {
	return getIMKCandidatesClass()
}

type IMKCandidatesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IMKCandidatesClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IMKCandidatesClass) Alloc() IMKCandidates {
	rv := objc.Send[IMKCandidates](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The [IMKCandidates] class presents candidates to users and notifies the
// appropriate [IMKInputController] object when the user selects a candidate.
// Candidates are alternate characters for a given input sequence. The
// [IMKCandidates] class supports using a candidates window in your input
// method; using [IMKCandidates] is optional. Not all input methods require
// them.
//
// # Overview
//
// When you create an [IMKCandidates] object, you attach it to the [IMKServer]
// object for your input method. You then need to override the
// [IMKInputController] methods “ and “ as well as implement a candidates
// method in your delegate object. The [IMKInputController] subclass supplies
// candidates to the [IMKCandidates] object by implementing the candidates
// method. When you are ready to display a candidates window, call the
// candidates method to update candidates and to show the candidates window.
//
// # Initializing a Candidates Window
//
//   - [IMKCandidates.InitWithServerPanelType]: Returns the initialized [IMKCandidates] object.
//
// # Managing Selection Keys
//
//   - [IMKCandidates.SetSelectionKeys]: Sets the selection keys for the candidates.
//   - [IMKCandidates.SelectionKeys]: Returns an array of [NSNumber] objects where each [NSNumber] object represents a virtual key code.
//   - [IMKCandidates.SetSelectionKeysKeylayout]: Sets the key layout that is used to map virtual key codes to characters.
//   - [IMKCandidates.SelectionKeysKeylayout]: Returns the key layout that maps virtual key codes to selection keys.
//
// # Managing Window Visibility and Behavior
//
//   - [IMKCandidates.Show]: Shows the candidates window.
//   - [IMKCandidates.Hide]: Hides a candidates window, if it is visible.
//   - [IMKCandidates.IsVisible]: Returns whether or not the candidates window is visible.
//   - [IMKCandidates.SetDismissesAutomatically]: Sets the state of the flag that determines whether the candidates window dismisses automatically.
//   - [IMKCandidates.DismissesAutomatically]: Returns the state of the flag that determines whether the candidates window dismisses automatically.
//   - [IMKCandidates.UpdateCandidates]: Updates the candidates that are displayed in the candidates window.
//
// # Managing Window Type and Text Attributes
//
//   - [IMKCandidates.PanelType]: Returns the style of the candidates window.
//   - [IMKCandidates.SetPanelType]: Sets the style of the candidates window.
//   - [IMKCandidates.SetAttributes]: Sets the style attributes for the candidates window.
//   - [IMKCandidates.Attributes]: Returns a dictionary of the style attributes used for the candidates window..
//
// # Showing an Annotation Window
//
//   - [IMKCandidates.ShowAnnotation]: Displays an annotation string in an annotation window.
//
// # Initializers
//
//   - [IMKCandidates.InitWithServerPanelTypeStyleType]
//
// # Instance Methods
//
//   - [IMKCandidates.AttachChildToCandidateType]
//   - [IMKCandidates.CandidateFrame]
//   - [IMKCandidates.CandidateIdentifierAtLineNumber]
//   - [IMKCandidates.CandidateStringIdentifier]
//   - [IMKCandidates.ClearSelection]
//   - [IMKCandidates.DetachChild]
//   - [IMKCandidates.HideChild]
//   - [IMKCandidates.LineNumberForCandidateWithIdentifier]
//   - [IMKCandidates.SelectCandidate]
//   - [IMKCandidates.SelectCandidateWithIdentifier]
//   - [IMKCandidates.SelectedCandidate]
//   - [IMKCandidates.SelectedCandidateString]
//   - [IMKCandidates.SetCandidateData]
//   - [IMKCandidates.SetCandidateFrameTopLeft]
//   - [IMKCandidates.ShowCandidates]
//   - [IMKCandidates.ShowChild]
//   - [IMKCandidates.ShowSublistSubListDelegate]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates
type IMKCandidates struct {
	appkit.NSResponder
}

// IMKCandidatesFromID constructs a [IMKCandidates] from an objc.ID.
//
// The [IMKCandidates] class presents candidates to users and notifies the
// appropriate [IMKInputController] object when the user selects a candidate.
// Candidates are alternate characters for a given input sequence. The
// [IMKCandidates] class supports using a candidates window in your input
// method; using [IMKCandidates] is optional. Not all input methods require
// them.
func IMKCandidatesFromID(id objc.ID) IMKCandidates {
	return IMKCandidates{NSResponder: appkit.NSResponderFromID(id)}
}

// NOTE: IMKCandidates adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IMKCandidates] class.
//
// # Initializing a Candidates Window
//
//   - [IIMKCandidates.InitWithServerPanelType]: Returns the initialized [IMKCandidates] object.
//
// # Managing Selection Keys
//
//   - [IIMKCandidates.SetSelectionKeys]: Sets the selection keys for the candidates.
//   - [IIMKCandidates.SelectionKeys]: Returns an array of [NSNumber] objects where each [NSNumber] object represents a virtual key code.
//   - [IIMKCandidates.SetSelectionKeysKeylayout]: Sets the key layout that is used to map virtual key codes to characters.
//   - [IIMKCandidates.SelectionKeysKeylayout]: Returns the key layout that maps virtual key codes to selection keys.
//
// # Managing Window Visibility and Behavior
//
//   - [IIMKCandidates.Show]: Shows the candidates window.
//   - [IIMKCandidates.Hide]: Hides a candidates window, if it is visible.
//   - [IIMKCandidates.IsVisible]: Returns whether or not the candidates window is visible.
//   - [IIMKCandidates.SetDismissesAutomatically]: Sets the state of the flag that determines whether the candidates window dismisses automatically.
//   - [IIMKCandidates.DismissesAutomatically]: Returns the state of the flag that determines whether the candidates window dismisses automatically.
//   - [IIMKCandidates.UpdateCandidates]: Updates the candidates that are displayed in the candidates window.
//
// # Managing Window Type and Text Attributes
//
//   - [IIMKCandidates.PanelType]: Returns the style of the candidates window.
//   - [IIMKCandidates.SetPanelType]: Sets the style of the candidates window.
//   - [IIMKCandidates.SetAttributes]: Sets the style attributes for the candidates window.
//   - [IIMKCandidates.Attributes]: Returns a dictionary of the style attributes used for the candidates window..
//
// # Showing an Annotation Window
//
//   - [IIMKCandidates.ShowAnnotation]: Displays an annotation string in an annotation window.
//
// # Initializers
//
//   - [IIMKCandidates.InitWithServerPanelTypeStyleType]
//
// # Instance Methods
//
//   - [IIMKCandidates.AttachChildToCandidateType]
//   - [IIMKCandidates.CandidateFrame]
//   - [IIMKCandidates.CandidateIdentifierAtLineNumber]
//   - [IIMKCandidates.CandidateStringIdentifier]
//   - [IIMKCandidates.ClearSelection]
//   - [IIMKCandidates.DetachChild]
//   - [IIMKCandidates.HideChild]
//   - [IIMKCandidates.LineNumberForCandidateWithIdentifier]
//   - [IIMKCandidates.SelectCandidate]
//   - [IIMKCandidates.SelectCandidateWithIdentifier]
//   - [IIMKCandidates.SelectedCandidate]
//   - [IIMKCandidates.SelectedCandidateString]
//   - [IIMKCandidates.SetCandidateData]
//   - [IIMKCandidates.SetCandidateFrameTopLeft]
//   - [IIMKCandidates.ShowCandidates]
//   - [IIMKCandidates.ShowChild]
//   - [IIMKCandidates.ShowSublistSubListDelegate]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates
type IIMKCandidates interface {
	appkit.INSResponder

	// Topic: Initializing a Candidates Window

	// Returns the initialized [IMKCandidates] object.
	InitWithServerPanelType(server IIMKServer, panelType IMKCandidatePanelType) IMKCandidates

	// Topic: Managing Selection Keys

	// Sets the selection keys for the candidates.
	SetSelectionKeys(keyCodes foundation.INSArray)
	// Returns an array of [NSNumber] objects where each [NSNumber] object represents a virtual key code.
	SelectionKeys() foundation.INSArray
	// Sets the key layout that is used to map virtual key codes to characters.
	SetSelectionKeysKeylayout(layout uintptr)
	// Returns the key layout that maps virtual key codes to selection keys.
	SelectionKeysKeylayout() uintptr

	// Topic: Managing Window Visibility and Behavior

	// Shows the candidates window.
	Show(locationHint IMKCandidatesLocationHint)
	// Hides a candidates window, if it is visible.
	Hide()
	// Returns whether or not the candidates window is visible.
	IsVisible() bool
	// Sets the state of the flag that determines whether the candidates window dismisses automatically.
	SetDismissesAutomatically(flag bool)
	// Returns the state of the flag that determines whether the candidates window dismisses automatically.
	DismissesAutomatically() bool
	// Updates the candidates that are displayed in the candidates window.
	UpdateCandidates()

	// Topic: Managing Window Type and Text Attributes

	// Returns the style of the candidates window.
	PanelType() IMKCandidatePanelType
	// Sets the style of the candidates window.
	SetPanelType(panelType IMKCandidatePanelType)
	// Sets the style attributes for the candidates window.
	SetAttributes(attributes foundation.INSDictionary)
	// Returns a dictionary of the style attributes used for the candidates window..
	Attributes() foundation.INSDictionary

	// Topic: Showing an Annotation Window

	// Displays an annotation string in an annotation window.
	ShowAnnotation(annotationString foundation.NSAttributedString)

	// Topic: Initializers

	InitWithServerPanelTypeStyleType(server IIMKServer, panelType IMKCandidatePanelType, style IMKStyleType) IMKCandidates

	// Topic: Instance Methods

	AttachChildToCandidateType(child IIMKCandidates, candidateIdentifier int, theType IMKStyleType)
	CandidateFrame() corefoundation.CGRect
	CandidateIdentifierAtLineNumber(lineNumber int) int
	CandidateStringIdentifier(candidateString objectivec.IObject) int
	ClearSelection()
	DetachChild(candidateIdentifier int)
	HideChild()
	LineNumberForCandidateWithIdentifier(candidateIdentifier int) int
	SelectCandidate(candidateIdentifier int)
	SelectCandidateWithIdentifier(candidateIdentifier int) bool
	SelectedCandidate() int
	SelectedCandidateString() foundation.NSAttributedString
	SetCandidateData(candidatesArray foundation.INSArray)
	SetCandidateFrameTopLeft(point corefoundation.CGPoint)
	ShowCandidates()
	ShowChild()
	ShowSublistSubListDelegate(candidates foundation.INSArray, delegate objectivec.IObject)
}

// Init initializes the instance.
func (i IMKCandidates) Init() IMKCandidates {
	rv := objc.Send[IMKCandidates](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IMKCandidates) Autorelease() IMKCandidates {
	rv := objc.Send[IMKCandidates](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKCandidates creates a new IMKCandidates instance.
func NewIMKCandidates() IMKCandidates {
	class := getIMKCandidatesClass()
	rv := objc.Send[IMKCandidates](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the initialized [IMKCandidates] object.
//
// server: The [IMKServer] object that manages the candidate and the panel type.
//
// panelType: A panel type for the candidate window.
//
// # Return Value
//
// The initialized [IMKCandidates] object.
//
// # Discussion
//
// When an input method allocates an [IMKCandidates] object it should
// initialize that object by calling this method.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:)
func NewIMKCandidatesWithServerPanelType(server IIMKServer, panelType IMKCandidatePanelType) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServer:panelType:"), server, panelType)
	return IMKCandidatesFromID(rv)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:styleType:)
func NewIMKCandidatesWithServerPanelTypeStyleType(server IIMKServer, panelType IMKCandidatePanelType, style IMKStyleType) IMKCandidates {
	instance := getIMKCandidatesClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServer:panelType:styleType:"), server, panelType, style)
	return IMKCandidatesFromID(rv)
}

// Returns the initialized [IMKCandidates] object.
//
// server: The [IMKServer] object that manages the candidate and the panel type.
//
// panelType: A panel type for the candidate window.
//
// # Return Value
//
// The initialized [IMKCandidates] object.
//
// # Discussion
//
// When an input method allocates an [IMKCandidates] object it should
// initialize that object by calling this method.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:)
func (i IMKCandidates) InitWithServerPanelType(server IIMKServer, panelType IMKCandidatePanelType) IMKCandidates {
	rv := objc.Send[IMKCandidates](i.ID, objc.Sel("initWithServer:panelType:"), server, panelType)
	return rv
}

// Sets the selection keys for the candidates.
//
// keyCodes: An array of [NSNumber] objects where each [NSNumber] object represents a
// virtual key code. The input controller maps these key codes to characters
// that are displayed either across the top of the candidates, if the
// candidates are laid out horizontally, or along the left edge of the
// candidates, if they are aligned vertically.
//
// # Discussion
//
// Selection keys are keys that can be used to select one of the candidates.
// They are displayed next to the candidate that will be selected when the
// user types that key.
//
// The number of selection keys determines how many candidates are displayed
// per page. For example, if you pass an array of four key codes, four
// candidates are displayed per page. If you pass eleven key codes, eleven
// candidates are displayed. By default, the key codes are mapped using the
// keyboard layout whose source id is
// `com.AppleXCUIElementTypeKeylayout().US`. You can replace the default
// layout by calling [IMKCandidates.SetSelectionKeysKeylayout]. The default
// selection keys are the digits 1 through 9 or, in terms of key codes, 18,
// 19, 20, 21, 23, 22, 26, 28, and 25.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeys(_:)
func (i IMKCandidates) SetSelectionKeys(keyCodes foundation.INSArray) {
	objc.Send[objc.ID](i.ID, objc.Sel("setSelectionKeys:"), keyCodes)
}

// Returns an array of [NSNumber] objects where each [NSNumber] object
// represents a virtual key code.
//
// # Return Value
//
// The array of [NSNumber] objects.
//
// # Discussion
//
// Selection keys are keys that can be used to select one of the candidates.
// They are displayed next to the candidate that will be selected when the
// user types that key.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeys()
func (i IMKCandidates) SelectionKeys() foundation.INSArray {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("selectionKeys"))
	return foundation.NSArrayFromID(rv)
}

// Sets the key layout that is used to map virtual key codes to characters.
//
// layout: The key layout to use.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setSelectionKeysKeylayout(_:)
func (i IMKCandidates) SetSelectionKeysKeylayout(layout uintptr) {
	objc.Send[objc.ID](i.ID, objc.Sel("setSelectionKeysKeylayout:"), layout)
}

// Returns the key layout that maps virtual key codes to selection keys.
//
// # Return Value
//
// The key layout in use. By default this is the key layout whose source id is
// `com.AppleXCUIElementTypeKeylayout().US`.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectionKeysKeylayout()
func (i IMKCandidates) SelectionKeysKeylayout() uintptr {
	rv := objc.Send[uintptr](i.ID, objc.Sel("selectionKeysKeylayout"))
	return rv
}

// Shows the candidates window.
//
// locationHint: A [IMKCandidatesLocationHint] constant that specifies the desired position
// of the candidates window. The Input Method Kit uses the hint to place the
// candidates window in a location that is in the vicinity of the hint
// location and ensures that the candidates window is fully visible.
//
// # Discussion
//
// Your input method calls this method when it is appropriate during text
// conversion to display a list of candidates.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show(_:)
func (i IMKCandidates) Show(locationHint IMKCandidatesLocationHint) {
	objc.Send[objc.ID](i.ID, objc.Sel("show:"), locationHint)
}

// Hides a candidates window, if it is visible.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hide()
func (i IMKCandidates) Hide() {
	objc.Send[objc.ID](i.ID, objc.Sel("hide"))
}

// Returns whether or not the candidates window is visible.
//
// # Return Value
//
// true if the candidates window is visible; otherwise false.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/isVisible()
func (i IMKCandidates) IsVisible() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isVisible"))
	return rv
}

// Sets the state of the flag that determines whether the candidates window
// dismisses automatically.
//
// flag: true to have the candidates window dismiss automatically; otherwise false.
//
// # Discussion
//
// By default, if the user presses the Return or Enter keys, the candidates
// are dismissed and a “ message is sent to the input controller. You can
// call the “ method, passing false as the `flag` parameter to change the
// default dismissal behavior. The input controller still receives a “
// message.
//
// When you set the flag to false, an input method processes text input while
// dynamically updating the content of the candidates as the user inputs text.
// When a session deactivates, candidate window is hidden regardless of the
// state of the flag.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setDismissesAutomatically(_:)
func (i IMKCandidates) SetDismissesAutomatically(flag bool) {
	objc.Send[objc.ID](i.ID, objc.Sel("setDismissesAutomatically:"), flag)
}

// Returns the state of the flag that determines whether the candidates window
// dismisses automatically.
//
// # Return Value
//
// true if the candidates window dismisses automatically; otherwise false.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/dismissesAutomatically()
func (i IMKCandidates) DismissesAutomatically() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("dismissesAutomatically"))
	return rv
}

// Updates the candidates that are displayed in the candidates window.
//
// # Discussion
//
// When you call this method, the Input Method Kit calls the candidates method
// of the [IMKInputController] class. Note that the candidates list is
// updated, but the visible state of the window does not change. In other
// words, if the window is hidden, it remains hidden. If the window is
// visible, it remains visible.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/update()
func (i IMKCandidates) UpdateCandidates() {
	objc.Send[objc.ID](i.ID, objc.Sel("updateCandidates"))
}

// Returns the style of the candidates window.
//
// # Return Value
//
// A [IMKCandidatePanelType] constant that represents the style of the
// candidates window.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/panelType()
func (i IMKCandidates) PanelType() IMKCandidatePanelType {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("panelType"))
	return IMKCandidatePanelType(rv)
}

// Sets the style of the candidates window.
//
// panelType: A [IMKCandidatePanelType] constant that represents the style of the
// candidates window.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setPanelType(_:)
func (i IMKCandidates) SetPanelType(panelType IMKCandidatePanelType) {
	objc.Send[objc.ID](i.ID, objc.Sel("setPanelType:"), panelType)
}

// Sets the style attributes for the candidates window.
//
// attributes: A dictionary that contains keys and values for the styles to use. You can
// supply the keys and values listed in the following table:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setAttributes(_:)
func (i IMKCandidates) SetAttributes(attributes foundation.INSDictionary) {
	objc.Send[objc.ID](i.ID, objc.Sel("setAttributes:"), attributes)
}

// Returns a dictionary of the style attributes used for the candidates
// window..
//
// # Return Value
//
// The dictionary that contains the keys and values for the styles.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attributes()
func (i IMKCandidates) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(rv)
}

// Displays an annotation string in an annotation window.
//
// annotationString: The string to display.
//
// # Discussion
//
// An annotation string explains or comments on the candidate string in the
// candidates window. An annotation window is a small, borderless window that
// is aligned with the current candidates window. An input method calls “
// when the “ method of the [IMKInputController] class is called, and the
// candidate string has annotations.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showAnnotation(_:)
func (i IMKCandidates) ShowAnnotation(annotationString foundation.NSAttributedString) {
	objc.Send[objc.ID](i.ID, objc.Sel("showAnnotation:"), annotationString)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/init(server:panelType:styleType:)
func (i IMKCandidates) InitWithServerPanelTypeStyleType(server IIMKServer, panelType IMKCandidatePanelType, style IMKStyleType) IMKCandidates {
	rv := objc.Send[IMKCandidates](i.ID, objc.Sel("initWithServer:panelType:styleType:"), server, panelType, style)
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/attachChild(_:toCandidate:type:)
func (i IMKCandidates) AttachChildToCandidateType(child IIMKCandidates, candidateIdentifier int, theType IMKStyleType) {
	objc.Send[objc.ID](i.ID, objc.Sel("attachChild:toCandidate:type:"), child, candidateIdentifier, theType)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateFrame()
func (i IMKCandidates) CandidateFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("candidateFrame"))
	return corefoundation.CGRect(rv)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateIdentifier(atLineNumber:)
func (i IMKCandidates) CandidateIdentifierAtLineNumber(lineNumber int) int {
	rv := objc.Send[int](i.ID, objc.Sel("candidateIdentifierAtLineNumber:"), lineNumber)
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/candidateStringIdentifier(_:)
func (i IMKCandidates) CandidateStringIdentifier(candidateString objectivec.IObject) int {
	rv := objc.Send[int](i.ID, objc.Sel("candidateStringIdentifier:"), candidateString)
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/clearSelection()
func (i IMKCandidates) ClearSelection() {
	objc.Send[objc.ID](i.ID, objc.Sel("clearSelection"))
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/detachChild(_:)
func (i IMKCandidates) DetachChild(candidateIdentifier int) {
	objc.Send[objc.ID](i.ID, objc.Sel("detachChild:"), candidateIdentifier)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/hideChild()
func (i IMKCandidates) HideChild() {
	objc.Send[objc.ID](i.ID, objc.Sel("hideChild"))
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/lineNumberForCandidate(withIdentifier:)
func (i IMKCandidates) LineNumberForCandidateWithIdentifier(candidateIdentifier int) int {
	rv := objc.Send[int](i.ID, objc.Sel("lineNumberForCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(_:)
func (i IMKCandidates) SelectCandidate(candidateIdentifier int) {
	objc.Send[objc.ID](i.ID, objc.Sel("selectCandidate:"), candidateIdentifier)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectCandidate(withIdentifier:)
func (i IMKCandidates) SelectCandidateWithIdentifier(candidateIdentifier int) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("selectCandidateWithIdentifier:"), candidateIdentifier)
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidate()
func (i IMKCandidates) SelectedCandidate() int {
	rv := objc.Send[int](i.ID, objc.Sel("selectedCandidate"))
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/selectedCandidateString()
func (i IMKCandidates) SelectedCandidateString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("selectedCandidateString"))
	return foundation.NSAttributedStringFromID(rv)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateData(_:)
func (i IMKCandidates) SetCandidateData(candidatesArray foundation.INSArray) {
	objc.Send[objc.ID](i.ID, objc.Sel("setCandidateData:"), candidatesArray)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/setCandidateFrameTopLeft(_:)
func (i IMKCandidates) SetCandidateFrameTopLeft(point corefoundation.CGPoint) {
	objc.Send[objc.ID](i.ID, objc.Sel("setCandidateFrameTopLeft:"), point)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/show()
func (i IMKCandidates) ShowCandidates() {
	objc.Send[objc.ID](i.ID, objc.Sel("showCandidates"))
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showChild()
func (i IMKCandidates) ShowChild() {
	objc.Send[objc.ID](i.ID, objc.Sel("showChild"))
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKCandidates/showSublist(_:subListDelegate:)
func (i IMKCandidates) ShowSublistSubListDelegate(candidates foundation.INSArray, delegate objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("showSublist:subListDelegate:"), candidates, delegate)
}
