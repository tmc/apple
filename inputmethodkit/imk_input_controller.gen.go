// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IMKInputController] class.
var (
	_IMKInputControllerClass     IMKInputControllerClass
	_IMKInputControllerClassOnce sync.Once
)

func getIMKInputControllerClass() IMKInputControllerClass {
	_IMKInputControllerClassOnce.Do(func() {
		_IMKInputControllerClass = IMKInputControllerClass{class: objc.GetClass("IMKInputController")}
	})
	return _IMKInputControllerClass
}

// GetIMKInputControllerClass returns the class object for IMKInputController.
func GetIMKInputControllerClass() IMKInputControllerClass {
	return getIMKInputControllerClass()
}

type IMKInputControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IMKInputControllerClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IMKInputControllerClass) Alloc() IMKInputController {
	rv := objc.Send[IMKInputController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The [IMKInputController] class provides a base class for custom input
// controller classes. The [IMKServer] class, which is allocated in the main
// function of an input method, creates an input controller object for each
// input session created by a client application. For every input session
// there is a corresponding [IMKInputController] object.
//
// # Overview
//
// An [IMKInputController] object controls text input on the input method
// side. It manages events and text from the applications and converted text
// from the input method engine. [IMKInputController] implements fully the
// [IMKStateSetting] and [IMKMouseHandling] protocols. Typically you do not
// need to override this class, but you do need to provide a delegate object
// that implements the methods that your are interested in. The
// [IMKInputController] versions of the protocol methods check whether the
// delegate object implements a method, and calls the delegate version if it
// exists.
//
// # Initializing an Input Controller
//
//   - [IMKInputController.InitWithServerDelegateClient]: Initializes the input control by setting the delegate.
//
// # Working with Ranges
//
//   - [IMKInputController.CompositionAttributesAtRange]: Returns a dictionary of text attributes.
//   - [IMKInputController.SelectionRange]: Returns where the range of the selection that should be placed inside marked text.
//   - [IMKInputController.ReplacementRange]: Returns the range in the client document that the text should replace.
//   - [IMKInputController.MarkForStyleAtRange]: Returns a dictionary of text attributes that can mark a range of an attributed string to send to a client.
//
// # Managing the Delegate
//
//   - [IMKInputController.Delegate]: Returns the delegate for input controller  object.
//   - [IMKInputController.SetDelegate]: Sets the delegate for input controller  object.
//
// # Getting the Client and Server Objects
//
//   - [IMKInputController.Server]: Returns the server object that manages the input controller.
//   - [IMKInputController.Client]: Returns the client object associated with the input controller.
//
// # Tracking Selections
//
//   - [IMKInputController.AnnotationSelectedForCandidate]: Sends the selected candidate string and annotation string to the input controller.
//   - [IMKInputController.CandidateSelectionChanged]: Informs an input controller that the current candidate selection in the candidate window has changed.
//   - [IMKInputController.CandidateSelected]: Informs an input controller that a new candidate is selected.
//
// # Managing Composition
//
//   - [IMKInputController.UpdateComposition]: Informs the input controller that the composition has changed.
//   - [IMKInputController.CancelComposition]: Stops the current composition and replaces marked text with the original text.
//
// # Hiding the User Interface
//
//   - [IMKInputController.HidePalettes]: Informs an input method that it should  close any visible user interface.
//
// # Working with Custom Commands
//
//   - [IMKInputController.DoCommandBySelectorCommandDictionary]: Passes commands that are not generated as part of the text input process.
//   - [IMKInputController.Menu]: Returns a menu of commands that are specific to an input method.
//
// # Instance Methods
//
//   - [IMKInputController.InputControllerWillClose]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController
type IMKInputController struct {
	objectivec.Object
}

// IMKInputControllerFromID constructs a [IMKInputController] from an objc.ID.
//
// The [IMKInputController] class provides a base class for custom input
// controller classes. The [IMKServer] class, which is allocated in the main
// function of an input method, creates an input controller object for each
// input session created by a client application. For every input session
// there is a corresponding [IMKInputController] object.
func IMKInputControllerFromID(id objc.ID) IMKInputController {
	return IMKInputController{objectivec.Object{ID: id}}
}

// NOTE: IMKInputController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IMKInputController] class.
//
// # Initializing an Input Controller
//
//   - [IIMKInputController.InitWithServerDelegateClient]: Initializes the input control by setting the delegate.
//
// # Working with Ranges
//
//   - [IIMKInputController.CompositionAttributesAtRange]: Returns a dictionary of text attributes.
//   - [IIMKInputController.SelectionRange]: Returns where the range of the selection that should be placed inside marked text.
//   - [IIMKInputController.ReplacementRange]: Returns the range in the client document that the text should replace.
//   - [IIMKInputController.MarkForStyleAtRange]: Returns a dictionary of text attributes that can mark a range of an attributed string to send to a client.
//
// # Managing the Delegate
//
//   - [IIMKInputController.Delegate]: Returns the delegate for input controller  object.
//   - [IIMKInputController.SetDelegate]: Sets the delegate for input controller  object.
//
// # Getting the Client and Server Objects
//
//   - [IIMKInputController.Server]: Returns the server object that manages the input controller.
//   - [IIMKInputController.Client]: Returns the client object associated with the input controller.
//
// # Tracking Selections
//
//   - [IIMKInputController.AnnotationSelectedForCandidate]: Sends the selected candidate string and annotation string to the input controller.
//   - [IIMKInputController.CandidateSelectionChanged]: Informs an input controller that the current candidate selection in the candidate window has changed.
//   - [IIMKInputController.CandidateSelected]: Informs an input controller that a new candidate is selected.
//
// # Managing Composition
//
//   - [IIMKInputController.UpdateComposition]: Informs the input controller that the composition has changed.
//   - [IIMKInputController.CancelComposition]: Stops the current composition and replaces marked text with the original text.
//
// # Hiding the User Interface
//
//   - [IIMKInputController.HidePalettes]: Informs an input method that it should  close any visible user interface.
//
// # Working with Custom Commands
//
//   - [IIMKInputController.DoCommandBySelectorCommandDictionary]: Passes commands that are not generated as part of the text input process.
//   - [IIMKInputController.Menu]: Returns a menu of commands that are specific to an input method.
//
// # Instance Methods
//
//   - [IIMKInputController.InputControllerWillClose]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController
type IIMKInputController interface {
	objectivec.IObject

	// Topic: Initializing an Input Controller

	// Initializes the input control by setting the delegate.
	InitWithServerDelegateClient(server IIMKServer, delegate objectivec.IObject, inputClient objectivec.IObject) IMKInputController

	// Topic: Working with Ranges

	// Returns a dictionary of text attributes.
	CompositionAttributesAtRange(range_ foundation.NSRange) foundation.INSDictionary
	// Returns where the range of the selection that should be placed inside marked text.
	SelectionRange() foundation.NSRange
	// Returns the range in the client document that the text should replace.
	ReplacementRange() foundation.NSRange
	// Returns a dictionary of text attributes that can mark a range of an attributed string to send to a client.
	MarkForStyleAtRange(style int, range_ foundation.NSRange) foundation.INSDictionary

	// Topic: Managing the Delegate

	// Returns the delegate for input controller  object.
	Delegate() objectivec.IObject
	// Sets the delegate for input controller  object.
	SetDelegate(newDelegate objectivec.IObject)

	// Topic: Getting the Client and Server Objects

	// Returns the server object that manages the input controller.
	Server() IIMKServer
	// Returns the client object associated with the input controller.
	Client() objectivec.IObject

	// Topic: Tracking Selections

	// Sends the selected candidate string and annotation string to the input controller.
	AnnotationSelectedForCandidate(annotationString foundation.NSAttributedString, candidateString foundation.NSAttributedString)
	// Informs an input controller that the current candidate selection in the candidate window has changed.
	CandidateSelectionChanged(candidateString foundation.NSAttributedString)
	// Informs an input controller that a new candidate is selected.
	CandidateSelected(candidateString foundation.NSAttributedString)

	// Topic: Managing Composition

	// Informs the input controller that the composition has changed.
	UpdateComposition()
	// Stops the current composition and replaces marked text with the original text.
	CancelComposition()

	// Topic: Hiding the User Interface

	// Informs an input method that it should  close any visible user interface.
	HidePalettes()

	// Topic: Working with Custom Commands

	// Passes commands that are not generated as part of the text input process.
	DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary foundation.INSDictionary)
	// Returns a menu of commands that are specific to an input method.
	Menu() appkit.NSMenu

	// Topic: Instance Methods

	InputControllerWillClose()

	// Activates the input method server.
	ActivateServer(sender objectivec.IObject)
	// Deactivates the input method server.
	DeactivateServer(sender objectivec.IObject)
	// Returns the modes dictionary associated with the input method.
	Modes(sender objectivec.IObject) foundation.INSDictionary
	// Handles mouse-down event send to an input method.
	MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) (bool, bool)
	// Handles a mouse-moved event sent to an input method.
	MouseMovedOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool
	// Handles a mouse-up event sent to an input method.
	MouseUpOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool
	// Returns an unsigned integer that contains a union of event masks
	RecognizedEvents(sender objectivec.IObject) uint
	// Set the value for the provided key.
	SetValueForTagClient(value objectivec.IObject, tag int, sender objectivec.IObject)
	// Displays a preferences window.
	ShowPreferences(sender objectivec.IObject)
	// Returns a value object whose key is the provided tag.
	ValueForTagClient(tag int, sender objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (i IMKInputController) Init() IMKInputController {
	rv := objc.Send[IMKInputController](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IMKInputController) Autorelease() IMKInputController {
	rv := objc.Send[IMKInputController](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKInputController creates a new IMKInputController instance.
func NewIMKInputController() IMKInputController {
	class := getIMKInputControllerClass()
	rv := objc.Send[IMKInputController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the input control by setting the delegate.
//
// server: The server object for the controller.
//
// delegate: The delegate object.
//
// inputClient: The client object that will send messages to the controller using the
// server object. The client object must confirm to the [IMKTextInput]
// protocol.
//
// # Return Value
//
// The initialized input controller object.
//
// # Discussion
//
// Methods in the [IMKStateSetting] and [IMKMouseHandling] protocols that are
// implemented by the delegate object always include a client parameter.
// Methods in the [IMKInputController] class do not need to take a client
// because the “ method stores the client object you supply as an ivar when
// it initializes the [IMKInputController] object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/init(server:delegate:client:)
func NewIMKInputControllerWithServerDelegateClient(server IIMKServer, delegate objectivec.IObject, inputClient objectivec.IObject) IMKInputController {
	instance := getIMKInputControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServer:delegate:client:"), server, delegate, inputClient)
	return IMKInputControllerFromID(rv)
}

// Initializes the input control by setting the delegate.
//
// server: The server object for the controller.
//
// delegate: The delegate object.
//
// inputClient: The client object that will send messages to the controller using the
// server object. The client object must confirm to the [IMKTextInput]
// protocol.
//
// # Return Value
//
// The initialized input controller object.
//
// # Discussion
//
// Methods in the [IMKStateSetting] and [IMKMouseHandling] protocols that are
// implemented by the delegate object always include a client parameter.
// Methods in the [IMKInputController] class do not need to take a client
// because the “ method stores the client object you supply as an ivar when
// it initializes the [IMKInputController] object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/init(server:delegate:client:)
func (i IMKInputController) InitWithServerDelegateClient(server IIMKServer, delegate objectivec.IObject, inputClient objectivec.IObject) IMKInputController {
	rv := objc.Send[IMKInputController](i.ID, objc.Sel("initWithServer:delegate:client:"), server, delegate, inputClient)
	return rv
}

// Returns a dictionary of text attributes.
//
// range: The range of text whose attributes you want to obtain.
//
// # Return Value
//
// The dictionary of text attributes. The default implementation returns an
// empty dictionary.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/compositionAttributes(at:)
func (i IMKInputController) CompositionAttributesAtRange(range_ foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("compositionAttributesAtRange:"), range_)
	return foundation.NSDictionaryFromID(rv)
}

// Returns where the range of the selection that should be placed inside
// marked text.
//
// # Return Value
//
// The range of the selection.
//
// # Discussion
//
// This method is called by [IMKInputController.UpdateComposition] to obtain
// the selection range for marked text. The default implementation sets the
// selection range at the end of the marked text. You should override this
// method if your input method provides font or glyph information.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/selectionRange()
func (i IMKInputController) SelectionRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](i.ID, objc.Sel("selectionRange"))
	return foundation.NSRange(rv)
}

// Returns the range in the client document that the text should replace.
//
// # Return Value
//
// The range to replace.
//
// # Discussion
//
// This method is called by [IMKInputController.UpdateComposition] to obtain
// the range in the client document where marked text should be placed. The
// default implementation returns an [NSRange] object whose location and
// length are [NSNotFound]. That indicates that the marked text should be
// placed at the current insertion point. Input methods that insert marked
// text somewhere other than at the current insertion point should override
// this method.
//
// An example of an input method that might override this method would be one
// replaces words with synonyms. That input method would watch for certain
// words and when it detects such a word it would replaced the word by marked
// text that was a synonym of the word.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/replacementRange()
func (i IMKInputController) ReplacementRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](i.ID, objc.Sel("replacementRange"))
	return foundation.NSRange(rv)
}

// Returns a dictionary of text attributes that can mark a range of an
// attributed string to send to a client.
//
// style: A style, which should be one of the following values:
// [kTSMHiliteSelectedRawText], [kTSMHiliteConvertedText], or
// [kTSMHiliteSelectedConvertedText]. See the `AERegistry.H()` header file for
// the definition of these values.
//
// range: The range (that is, a clause) to mark.
//
// # Return Value
//
// The dictionary of text attributes.
//
// # Discussion
//
// This utility function can be called by input methods to mark each range
// (i.e. clause ) of marked text. T
//
// The default implementation first calls the method
// [IMKInputController.CompositionAttributesAtRange] to obtain the additional
// attributes that an input method wants to include, such as font or glyph
// information. Then, it adds the appropriate underline and underline color
// information to the attributes dictionary for the style parameter. Finally
// it adds the style value as the dictionary value. The key for the style
// value is [markedClauseSegment].
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/mark(forStyle:at:)
//
// [kTSMHiliteConvertedText]: https://developer.apple.com/documentation/coreservices/ktsmhiliteconvertedtext
// [kTSMHiliteSelectedConvertedText]: https://developer.apple.com/documentation/coreservices/ktsmhiliteselectedconvertedtext
// [kTSMHiliteSelectedRawText]: https://developer.apple.com/documentation/coreservices/ktsmhiliteselectedrawtext
// [markedClauseSegment]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/markedClauseSegment
func (i IMKInputController) MarkForStyleAtRange(style int, range_ foundation.NSRange) foundation.INSDictionary {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("markForStyle:atRange:"), style, range_)
	return foundation.NSDictionaryFromID(rv)
}

// Returns the delegate for input controller object.
//
// # Return Value
//
// The delegate object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/delegate()
func (i IMKInputController) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}

// Sets the delegate for input controller object.
//
// newDelegate: The delegate object to set.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/setDelegate(_:)
func (i IMKInputController) SetDelegate(newDelegate objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("setDelegate:"), newDelegate)
}

// Returns the server object that manages the input controller.
//
// # Return Value
//
// The server object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/server()
func (i IMKInputController) Server() IIMKServer {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("server"))
	return IMKServerFromID(rv)
}

// Returns the client object associated with the input controller.
//
// # Return Value
//
// The client object.
//
// # Discussion
//
// The client object conforms to the [IMKTextInput] protocol.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/client()
func (i IMKInputController) Client() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("client"))
	return objectivec.Object{ID: rv}
}

// Sends the selected candidate string and annotation string to the input
// controller.
//
// annotationString: The annotation string associated with the candidate.
//
// candidateString: The candidate string that the user moved to.
//
// # Discussion
//
// This method is called when the user moves to a candidate.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/annotationSelected(_:forCandidate:)
func (i IMKInputController) AnnotationSelectedForCandidate(annotationString foundation.NSAttributedString, candidateString foundation.NSAttributedString) {
	objc.Send[objc.ID](i.ID, objc.Sel("annotationSelected:forCandidate:"), annotationString, candidateString)
}

// Informs an input controller that the current candidate selection in the
// candidate window has changed.
//
// candidateString: The changed candidate string.
//
// # Discussion
//
// Note this method is called to indicate user activity in the candidate
// window. The candidate object might not be the user’s final selection.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelectionChanged(_:)
func (i IMKInputController) CandidateSelectionChanged(candidateString foundation.NSAttributedString) {
	objc.Send[objc.ID](i.ID, objc.Sel("candidateSelectionChanged:"), candidateString)
}

// Informs an input controller that a new candidate is selected.
//
// candidateString: The changed candidate string.
//
// # Discussion
//
// The candidate object is the user’s final choice from the candidate
// window. The candidate window is closed before this method is called.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/candidateSelected(_:)
func (i IMKInputController) CandidateSelected(candidateString foundation.NSAttributedString) {
	objc.Send[objc.ID](i.ID, objc.Sel("candidateSelected:"), candidateString)
}

// Informs the input controller that the composition has changed.
//
// # Discussion
//
// This method calls the protocol method composedString: to obtain the current
// composition. The current composition is sent to the client by a call to the
// method `setMarkedText(_:)`.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/updateComposition()
func (i IMKInputController) UpdateComposition() {
	objc.Send[objc.ID](i.ID, objc.Sel("updateComposition"))
}

// Stops the current composition and replaces marked text with the original
// text.
//
// # Discussion
//
// This method calls the method originalString: to obtain the original text
// and sends that text to the client using a call to the [IMKTextInput]
// protocol method `insertText(_:)`
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/cancelComposition()
func (i IMKInputController) CancelComposition() {
	objc.Send[objc.ID](i.ID, objc.Sel("cancelComposition"))
}

// Informs an input method that it should close any visible user interface.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/hidePalettes()
func (i IMKInputController) HidePalettes() {
	objc.Send[objc.ID](i.ID, objc.Sel("hidePalettes"))
}

// Passes commands that are not generated as part of the text input process.
//
// aSelector: A selector that represents a command from the text input menu.
//
// infoDictionary: A dictionary that contains two key-value pairs:
//
// - [kIMKCommandMenuItemName], whose value is an [NSMenuItem] object. That
// is, the item selected by the user. - [kIMKCommandClientName], whose value
// is the current client—`id`.
//
// # Discussion
//
// The default implementation checks if the input controller object (that is,
// self) responds to the selector. If so, it sends the message
// [perform(_:with:)] to the input controller class. The object parameter in
// that case is the `infoDictionary` parameter.
//
// This method is called when a user selects a command from the text input
// menu. To support this, an input method must provide actions for each menu
// item that is placed in the menu. For example, `(void)(id)sender`. Note that
// the sender in this instance is the info dictionary.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/doCommand(by:command:)
//
// [kIMKCommandClientName]: https://developer.apple.com/documentation/InputMethodKit/kIMKCommandClientName
// [kIMKCommandMenuItemName]: https://developer.apple.com/documentation/InputMethodKit/kIMKCommandMenuItemName
// [perform(_:with:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/perform(_:with:)
func (i IMKInputController) DoCommandBySelectorCommandDictionary(aSelector objc.SEL, infoDictionary foundation.INSDictionary) {
	objc.Send[objc.ID](i.ID, objc.Sel("doCommandBySelector:commandDictionary:"), aSelector, infoDictionary)
}

// Returns a menu of commands that are specific to an input method.
//
// # Return Value
//
// The menu object.
//
// # Discussion
//
// This method is called whenever the menu needs to be drawn so that an input
// method can update the menu to reflect the current state.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/menu()
func (i IMKInputController) Menu() appkit.NSMenu {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("menu"))
	return appkit.NSMenuFromID(rv)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKInputController/inputControllerWillClose()
func (i IMKInputController) InputControllerWillClose() {
	objc.Send[objc.ID](i.ID, objc.Sel("inputControllerWillClose"))
}

// Activates the input method server.
//
// sender: The object sending the activation message.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/activateServer(_:)
func (i IMKInputController) ActivateServer(sender objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("activateServer:"), sender)
}

// Deactivates the input method server.
//
// sender: The object sending the deactivation message.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/deactivateServer(_:)
func (i IMKInputController) DeactivateServer(sender objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("deactivateServer:"), sender)
}

// Returns the modes dictionary associated with the input method.
//
// sender: The client object requesting the modes dictionary.
//
// # Return Value
//
// The modes dictionary associated with the input method.
//
// # Discussion
//
// Typically a client object calls this method to to build the text input
// menu. By calling the input method rather than reading the modes from the
// `Info.Plist()` file, the input method can dynamically modify the modes
// supported.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/modes(_:)
func (i IMKInputController) Modes(sender objectivec.IObject) foundation.INSDictionary {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("modes:"), sender)
	return foundation.NSDictionaryFromID(rv)
}

// Handles mouse-down event send to an input method.
//
// index: The index within the sender’s text storage where the mouse-down event
// occurred.
//
// point: The point at which the mouse-down event occurred.
//
// flags: The modifier keys.
//
// keepTracking: Set this parameter to true if you want to receive subsequent mouse-moved
// and mouse -up events.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-down events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseDown(onCharacterIndex:coordinate:withModifier:continueTracking:client:)
func (i IMKInputController) MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) (bool, bool) {
	var keepTracking bool
	rv := objc.Send[bool](i.ID, objc.Sel("mouseDownOnCharacterIndex:coordinate:withModifier:continueTracking:client:"), index, point, flags, unsafe.Pointer(&keepTracking), sender)
	return keepTracking, rv
}

// Handles a mouse-moved event sent to an input method.
//
// index: The index within the sender’s text storage where the mouse-moved event
// occurred.
//
// point: The point at which the mouse-moved event occurred.
//
// flags: The modifier keys.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-moved events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseMoved(onCharacterIndex:coordinate:withModifier:client:)
func (i IMKInputController) MouseMovedOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("mouseMovedOnCharacterIndex:coordinate:withModifier:client:"), index, point, flags, sender)
	return rv
}

// Handles a mouse-up event sent to an input method.
//
// index: The index within the sender’s text storage where the mouse-up event
// occurred.
//
// point: The point at which the mouse-up event occurred.
//
// flags: The modifier keys.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-up events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseUp(onCharacterIndex:coordinate:withModifier:client:)
func (i IMKInputController) MouseUpOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("mouseUpOnCharacterIndex:coordinate:withModifier:client:"), index, point, flags, sender)
	return rv
}

// Returns an unsigned integer that contains a union of event masks
//
// sender: The client object requesting the supported events.
//
// # Return Value
//
// An unsigned integer that contains a union of event masks (See the
// `NSEvent.H()` header file.
//
// # Discussion
//
// A client calls this method to check whether an input method supports an
// event. The default implementation returns [NSKeyDownMask]. If your input
// method handles only key down events, the Input Method Kit provides the
// default mouse handling. The default mouse-down handling behavior is as
// follows: If there is an active composition area and the user clicks in the
// text but outside of the composition area, the Input Method Kit sends your
// input method a “ message. This happens only for input methods that return
// only the default value—[NSKeyDownMask].
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/recognizedEvents(_:)
func (i IMKInputController) RecognizedEvents(sender objectivec.IObject) uint {
	rv := objc.Send[uint](i.ID, objc.Sel("recognizedEvents:"), sender)
	return rv
}

// Set the value for the provided key.
//
// value: The value, specified as the appropriate object (such as [NSNumber]), to
// set.
//
// tag: The key whose value you want to set.
//
// sender: The client setting the value.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/setValue(_:forTag:client:)
func (i IMKInputController) SetValueForTagClient(value objectivec.IObject, tag int, sender objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("setValue:forTag:client:"), value, tag, sender)
}

// Displays a preferences window.
//
// sender: The object sending the message to show the preference window.
//
// # Discussion
//
// This method looks for a nib file that contains a window controller class
// and a preferences utility. If found, it displays the window. To use this
// method you must create a menu item in your input method menu whose action
// is “. When a user selects that item, the Input Method Kit invokes your “
// method. The default implementation looks for a nib file named
// `preferences.Nib()`. If found, it allocates a window controller class loads
// the nib file. You can provide a custom window controller class by naming
// the class in your input method `info.Plist()` file, providing a key-value
// pair. The key must be [InputMethodServerPreferencesWindowControllerClass]
// and the associated value must be the name of your custom class.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/showPreferences(_:)
func (i IMKInputController) ShowPreferences(sender objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("showPreferences:"), sender)
}

// Returns a value object whose key is the provided tag.
//
// tag: The key whose value you want to retrieve.
//
// sender: The client requesting the value.
//
// # Return Value
//
// The value object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/value(forTag:client:)
func (i IMKInputController) ValueForTagClient(tag int, sender objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("valueForTag:client:"), tag, sender)
	return objectivec.Object{ID: rv}
}

// Protocol methods for IMKMouseHandling

// Protocol methods for IMKStateSetting
