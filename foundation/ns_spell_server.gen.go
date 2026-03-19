// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSpellServer] class.
var (
	_NSSpellServerClass     NSSpellServerClass
	_NSSpellServerClassOnce sync.Once
)

func getNSSpellServerClass() NSSpellServerClass {
	_NSSpellServerClassOnce.Do(func() {
		_NSSpellServerClass = NSSpellServerClass{class: objc.GetClass("NSSpellServer")}
	})
	return _NSSpellServerClass
}

// GetNSSpellServerClass returns the class object for NSSpellServer.
func GetNSSpellServerClass() NSSpellServerClass {
	return getNSSpellServerClass()
}

type NSSpellServerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSpellServerClass) Alloc() NSSpellServer {
	rv := objc.Send[NSSpellServer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A server that your app uses to provide a spell checker service to other
// apps running in the system.
//
// # Overview
// 
// A is an application that declares its availability in a standard way, so
// that any other applications that wish to use it can do so. If you build a
// spelling checker that makes use of the [NSSpellServer] class and list it as
// an available service, then users of any application that makes use of
// [NSSpellChecker] or includes a Services menu will see your spelling checker
// as one of the available dictionaries.
//
// [NSSpellChecker]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
//
// # Configuring Spelling Servers
//
//   - [NSSpellServer.Delegate]: Returns the receiver’s delegate.
//   - [NSSpellServer.SetDelegate]
//
// # Providing Spelling Services
//
//   - [NSSpellServer.RegisterLanguageByVendor]: Notifies the receiver of a language your spelling checker can check.
//   - [NSSpellServer.Run]: Causes the receiver to start listening for spell-checking requests.
//
// # Managing the Spell-Checking Process
//
//   - [NSSpellServer.IsWordInUserDictionariesCaseSensitive]: Indicates whether a given word is in the user’s list of learned words or the document’s list of words to ignore.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer
type NSSpellServer struct {
	objectivec.Object
}

// NSSpellServerFromID constructs a [NSSpellServer] from an objc.ID.
//
// A server that your app uses to provide a spell checker service to other
// apps running in the system.
func NSSpellServerFromID(id objc.ID) NSSpellServer {
	return NSSpellServer{objectivec.Object{ID: id}}
}
// NOTE: NSSpellServer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSpellServer] class.
//
// # Configuring Spelling Servers
//
//   - [INSSpellServer.Delegate]: Returns the receiver’s delegate.
//   - [INSSpellServer.SetDelegate]
//
// # Providing Spelling Services
//
//   - [INSSpellServer.RegisterLanguageByVendor]: Notifies the receiver of a language your spelling checker can check.
//   - [INSSpellServer.Run]: Causes the receiver to start listening for spell-checking requests.
//
// # Managing the Spell-Checking Process
//
//   - [INSSpellServer.IsWordInUserDictionariesCaseSensitive]: Indicates whether a given word is in the user’s list of learned words or the document’s list of words to ignore.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer
type INSSpellServer interface {
	objectivec.IObject

	// Topic: Configuring Spelling Servers

	// Returns the receiver’s delegate.
	Delegate() NSSpellServerDelegate
	SetDelegate(value NSSpellServerDelegate)

	// Topic: Providing Spelling Services

	// Notifies the receiver of a language your spelling checker can check.
	RegisterLanguageByVendor(language string, vendor string) bool
	// Causes the receiver to start listening for spell-checking requests.
	Run()

	// Topic: Managing the Spell-Checking Process

	// Indicates whether a given word is in the user’s list of learned words or the document’s list of words to ignore.
	IsWordInUserDictionariesCaseSensitive(word string, flag bool) bool
}

// Init initializes the instance.
func (s NSSpellServer) Init() NSSpellServer {
	rv := objc.Send[NSSpellServer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSpellServer) Autorelease() NSSpellServer {
	rv := objc.Send[NSSpellServer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSpellServer creates a new NSSpellServer instance.
func NewNSSpellServer() NSSpellServer {
	class := getNSSpellServerClass()
	rv := objc.Send[NSSpellServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Notifies the receiver of a language your spelling checker can check.
//
// language: A string specifying the English name of a language on Apple’s list of
// languages.
//
// vendor: A string that identifies the vendor (to distinguish your spelling checker
// from those that others may offer for the same language).
//
// # Return Value
// 
// Returns [true] if the language is registered, [false] if for some reason it
// can’t be registered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If your spelling checker supports more than one language, it should invoke
// this method once for each language. Registering a language-vendor
// combination causes it to appear in the Spelling panel’s pop-up menu of
// spelling checkers.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer/registerLanguage(_:byVendor:)
func (s NSSpellServer) RegisterLanguageByVendor(language string, vendor string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("registerLanguage:byVendor:"), objc.String(language), objc.String(vendor))
	return rv
}
// Causes the receiver to start listening for spell-checking requests.
//
// # Discussion
// 
// This method starts a loop that never returns; you need to set the
// [NSSpellServer] object’s delegate before sending this message.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer/run()
func (s NSSpellServer) Run() {
	objc.Send[objc.ID](s.ID, objc.Sel("run"))
}
// Indicates whether a given word is in the user’s list of learned words or
// the document’s list of words to ignore.
//
// word: The word to compare with those in the user dictionaries.
//
// flag: Specifies whether the comparison is case sensitive.
//
// # Return Value
// 
// A Boolean value indicating whether the word is in the user dictionaries. If
// [true], the word is acceptable to the user.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer/isWord(inUserDictionaries:caseSensitive:)
func (s NSSpellServer) IsWordInUserDictionariesCaseSensitive(word string, flag bool) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isWordInUserDictionaries:caseSensitive:"), objc.String(word), flag)
	return rv
}

// Returns the receiver’s delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServer/delegate
func (s NSSpellServer) Delegate() NSSpellServerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSpellServerDelegateObjectFromID(rv)
}
func (s NSSpellServer) SetDelegate(value NSSpellServerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

