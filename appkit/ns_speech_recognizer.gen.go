// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSpeechRecognizer] class.
var (
	_NSSpeechRecognizerClass     NSSpeechRecognizerClass
	_NSSpeechRecognizerClassOnce sync.Once
)

func getNSSpeechRecognizerClass() NSSpeechRecognizerClass {
	_NSSpeechRecognizerClassOnce.Do(func() {
		_NSSpeechRecognizerClass = NSSpeechRecognizerClass{class: objc.GetClass("NSSpeechRecognizer")}
	})
	return _NSSpeechRecognizerClass
}

// GetNSSpeechRecognizerClass returns the class object for NSSpeechRecognizer.
func GetNSSpeechRecognizerClass() NSSpeechRecognizerClass {
	return getNSSpeechRecognizerClass()
}

type NSSpeechRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSpeechRecognizerClass) Alloc() NSSpeechRecognizer {
	rv := objc.Send[NSSpeechRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The Cocoa interface to speech recognition in macOS.
//
// # Overview
// 
// [NSSpeechRecognizer] provides a “command and control” style of voice
// recognition system, where the command phrases must be defined prior to
// listening, in contrast to a dictation system where the recognized text is
// unconstrained. Through an [NSSpeechRecognizer] instance, Cocoa apps can use
// the speech recognition engine built into macOS to recognize spoken
// commands. With speech recognition, users can accomplish complex tasks with
// spoken commands—for example, “Move pawn B2 to B4” and “Take back
// move.”
// 
// The [NSSpeechRecognizer] class has a property that lets you specify which
// spoken words should be recognized as commands ([NSSpeechRecognizer.Commands]) and methods that
// let you start and stop listening ([NSSpeechRecognizer.StartListening] and [NSSpeechRecognizer.StopListening]).
// When the speech recognition facility recognizes one of the designated
// commands, [NSSpeechRecognizer] invokes the delegation method
// [SpeechRecognizerDidRecognizeCommand], allowing the delegate to perform the
// command.
// 
// Speech recognition is just one of the macOS speech technologies. The speech
// synthesis technology allows applications to “pronounce” written text in
// U.S. English and over 25 other languages, with a number of different voices
// and dialects for each language ([NSSpeechSynthesizer] is the Cocoa
// interface to this technology). Both speech technologies provide benefits
// for all users, and are particularly useful to those users who have
// difficulties seeing the screen or using the mouse and keyboard. By
// incorporating speech into your application, you can provide a concurrent
// mode of interaction for your users: In macOS, your software can accept
// input and provide output without requiring users to change their working
// context.
//
// # Handling the Recognition of a Spoken Command
//
//   - [NSSpeechRecognizer.Delegate]: The delegate for the speech recognizer object.
//   - [NSSpeechRecognizer.SetDelegate]
//
// # Configuring Speech Recognizers
//
//   - [NSSpeechRecognizer.Commands]: An array of strings defining the commands for which the speech recognizer object should listen.
//   - [NSSpeechRecognizer.SetCommands]
//   - [NSSpeechRecognizer.DisplayedCommandsTitle]: The title of the commands section in the Speech Commands window or `nil` if there is no title.
//   - [NSSpeechRecognizer.SetDisplayedCommandsTitle]
//   - [NSSpeechRecognizer.ListensInForegroundOnly]: A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//   - [NSSpeechRecognizer.SetListensInForegroundOnly]
//   - [NSSpeechRecognizer.BlocksOtherRecognizers]: A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//   - [NSSpeechRecognizer.SetBlocksOtherRecognizers]
//
// # Listening
//
//   - [NSSpeechRecognizer.StartListening]: Tells the speech recognition engine to begin listening for commands.
//   - [NSSpeechRecognizer.StopListening]: Tells the speech recognition engine to suspend listening for commands.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer
type NSSpeechRecognizer struct {
	objectivec.Object
}

// NSSpeechRecognizerFromID constructs a [NSSpeechRecognizer] from an objc.ID.
//
// The Cocoa interface to speech recognition in macOS.
func NSSpeechRecognizerFromID(id objc.ID) NSSpeechRecognizer {
	return NSSpeechRecognizer{objectivec.Object{ID: id}}
}
// NOTE: NSSpeechRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSpeechRecognizer] class.
//
// # Handling the Recognition of a Spoken Command
//
//   - [INSSpeechRecognizer.Delegate]: The delegate for the speech recognizer object.
//   - [INSSpeechRecognizer.SetDelegate]
//
// # Configuring Speech Recognizers
//
//   - [INSSpeechRecognizer.Commands]: An array of strings defining the commands for which the speech recognizer object should listen.
//   - [INSSpeechRecognizer.SetCommands]
//   - [INSSpeechRecognizer.DisplayedCommandsTitle]: The title of the commands section in the Speech Commands window or `nil` if there is no title.
//   - [INSSpeechRecognizer.SetDisplayedCommandsTitle]
//   - [INSSpeechRecognizer.ListensInForegroundOnly]: A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
//   - [INSSpeechRecognizer.SetListensInForegroundOnly]
//   - [INSSpeechRecognizer.BlocksOtherRecognizers]: A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
//   - [INSSpeechRecognizer.SetBlocksOtherRecognizers]
//
// # Listening
//
//   - [INSSpeechRecognizer.StartListening]: Tells the speech recognition engine to begin listening for commands.
//   - [INSSpeechRecognizer.StopListening]: Tells the speech recognition engine to suspend listening for commands.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer
type INSSpeechRecognizer interface {
	objectivec.IObject

	// Topic: Handling the Recognition of a Spoken Command

	// The delegate for the speech recognizer object.
	Delegate() NSSpeechRecognizerDelegate
	SetDelegate(value NSSpeechRecognizerDelegate)

	// Topic: Configuring Speech Recognizers

	// An array of strings defining the commands for which the speech recognizer object should listen.
	Commands() []string
	SetCommands(value []string)
	// The title of the commands section in the Speech Commands window or `nil` if there is no title.
	DisplayedCommandsTitle() string
	SetDisplayedCommandsTitle(value string)
	// A Boolean value that indicates whether the speech recognizer object should only enable its commands when its application is the frontmost one.
	ListensInForegroundOnly() bool
	SetListensInForegroundOnly(value bool)
	// A Boolean value that indicates whether the speech recognizer object should block all other recognizers (that is, other applications attempting to understand spoken commands) when listening.
	BlocksOtherRecognizers() bool
	SetBlocksOtherRecognizers(value bool)

	// Topic: Listening

	// Tells the speech recognition engine to begin listening for commands.
	StartListening()
	// Tells the speech recognition engine to suspend listening for commands.
	StopListening()
}

// Init initializes the instance.
func (s NSSpeechRecognizer) Init() NSSpeechRecognizer {
	rv := objc.Send[NSSpeechRecognizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSpeechRecognizer) Autorelease() NSSpeechRecognizer {
	rv := objc.Send[NSSpeechRecognizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSpeechRecognizer creates a new NSSpeechRecognizer instance.
func NewNSSpeechRecognizer() NSSpeechRecognizer {
	class := getNSSpeechRecognizerClass()
	rv := objc.Send[NSSpeechRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the speech recognition engine to begin listening for commands.
//
// # Discussion
// 
// When a command is recognized the message
// [SpeechRecognizerDidRecognizeCommand] is sent to the delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/startListening()
func (s NSSpeechRecognizer) StartListening() {
	objc.Send[objc.ID](s.ID, objc.Sel("startListening"))
}

// Tells the speech recognition engine to suspend listening for commands.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/stopListening()
func (s NSSpeechRecognizer) StopListening() {
	objc.Send[objc.ID](s.ID, objc.Sel("stopListening"))
}

// The delegate for the speech recognizer object.
//
// # Discussion
// 
// The delegate must conform to the [NSSpeechRecognizerDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/delegate
func (s NSSpeechRecognizer) Delegate() NSSpeechRecognizerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSpeechRecognizerDelegateObjectFromID(rv)
}
func (s NSSpeechRecognizer) SetDelegate(value NSSpeechRecognizerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// An array of strings defining the commands for which the speech recognizer
// object should listen.
//
// # Discussion
// 
// Setting this property when the speech recognizer is already listening means
// that the current command list is updated and listening continues. The items
// in the array should be [NSString] objects. The command strings must match
// the current locale of the recognizer, which is selected in the Dictation
// pane of Accessibility system preferences.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/commands
func (s NSSpeechRecognizer) Commands() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("commands"))
	return objc.ConvertSliceToStrings(rv)
}
func (s NSSpeechRecognizer) SetCommands(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setCommands:"), objectivec.StringSliceToNSArray(value))
}

// The title of the commands section in the Speech Commands window or `nil` if
// there is no title.
//
// # Discussion
// 
// When this property is a non-empty string, commands are displayed in the
// Speech Commands window indented under a section with this title. If `title`
// is `nil` or an empty string, the commands are displayed at the top level of
// the Speech Commands window. This default is not to display the commands
// under a section title.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/displayedCommandsTitle
func (s NSSpeechRecognizer) DisplayedCommandsTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayedCommandsTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSpeechRecognizer) SetDisplayedCommandsTitle(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayedCommandsTitle:"), objc.String(value))
}

// A Boolean value that indicates whether the speech recognizer object should
// only enable its commands when its application is the frontmost one.
//
// # Discussion
// 
// When the value of this property is [true], the speech recognizer’s
// commands are only recognized when the speech recognizer’s application is
// the frontmost application—typically the application displaying the menu
// bar. If the value of the property is [false], the commands are recognized
// regardless of the visibility of the application, including agent
// applications (agent applications, which have the [LSUIElement] property
// set, do not appear in the Dock or Force Quit window). The default value of
// this property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/listensInForegroundOnly
func (s NSSpeechRecognizer) ListensInForegroundOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("listensInForegroundOnly"))
	return rv
}
func (s NSSpeechRecognizer) SetListensInForegroundOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setListensInForegroundOnly:"), value)
}

// A Boolean value that indicates whether the speech recognizer object should
// block all other recognizers (that is, other applications attempting to
// understand spoken commands) when listening.
//
// # Discussion
// 
// When the value of this property is [true], all other speech recognition
// commands on the system are disabled until the speech recognizer is
// released, listening is stopped, or the property is set to [false]. Setting
// this property to [true] effectively takes over the computer at the expense
// of other applications using speech recognition, so you should use it only
// in circumstances that warrant it, such as when listening for a response
// important to overall system operation or when an application is running in
// full-screen mode (such as games and presentation software). The default is
// value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizer/blocksOtherRecognizers
func (s NSSpeechRecognizer) BlocksOtherRecognizers() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("blocksOtherRecognizers"))
	return rv
}
func (s NSSpeechRecognizer) SetBlocksOtherRecognizers(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setBlocksOtherRecognizers:"), value)
}

