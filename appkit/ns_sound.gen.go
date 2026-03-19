// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSound] class.
var (
	_NSSoundClass     NSSoundClass
	_NSSoundClassOnce sync.Once
)

func getNSSoundClass() NSSoundClass {
	_NSSoundClassOnce.Do(func() {
		_NSSoundClass = NSSoundClass{class: objc.GetClass("NSSound")}
	})
	return _NSSoundClass
}

// GetNSSoundClass returns the class object for NSSound.
func GetNSSoundClass() NSSoundClass {
	return getNSSoundClass()
}

type NSSoundClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSoundClass) Alloc() NSSound {
	rv := objc.Send[NSSound](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A simple interface for loading and playing audio files.
//
// # Overview
// 
// You create a sound object with an audio file or data, which can be in any
// format that Core Audio supports. Customize the sound by configuring its
// properties, such as setting its playback volume and looping behavior. Call
// the sound’s [NSSound.Play] method to begin playback. The system executes this
// call asynchronously so that it doesn’t interrupt the functioning of your
// app.
// 
// If you want to play the system beep sound, use the [beep()] (Swift) or
// [NSBeep] (Objective-C) function.
//
// [beep()]: https://developer.apple.com/documentation/AppKit/NSSound/beep()
//
// # Detecting When a Sound Finishes Playing
//
//   - [NSSound.Delegate]: The sound’s delegate.
//   - [NSSound.SetDelegate]
//
// # Creating Sounds
//
//   - [NSSound.InitWithContentsOfFileByReference]: Initializes the receiver with the audio data located at a given filepath.
//   - [NSSound.InitWithContentsOfURLByReference]: Initializes the receiver with the audio data located at a given URL.
//   - [NSSound.InitWithData]: Initializes the receiver with a given audio data.
//   - [NSSound.InitWithPasteboard]: Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by [NSSound](<doc://com.apple.appkit/documentation/AppKit/NSSound>). [NSSound] expects the data to have a proper magic number, sound header, and data for the formats it supports.
//
// # Configuring Sounds
//
//   - [NSSound.Name]: The name assigned to the sound.
//   - [NSSound.Volume]: The volume of the sound.
//   - [NSSound.SetVolume]
//   - [NSSound.CurrentTime]: The sound’s playback progress, in seconds.
//   - [NSSound.SetCurrentTime]
//   - [NSSound.Loops]: A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//   - [NSSound.SetLoops]
//   - [NSSound.PlaybackDeviceIdentifier]: Identifies the sound’s output device
//   - [NSSound.SetPlaybackDeviceIdentifier]
//
// # Getting Sound Information
//
//   - [NSSound.Duration]: The duration of the sound, in seconds.
//
// # Playing Sounds
//
//   - [NSSound.Playing]: A Boolean that indicates whether the sound is playing its audio data.
//   - [NSSound.Pause]: Pauses audio playback.
//   - [NSSound.Play]: Initiates audio playback.
//   - [NSSound.Resume]: Resumes audio playback.
//   - [NSSound.Stop]: Concludes audio playback.
//
// # Writing Sounds
//
//   - [NSSound.WriteToPasteboard]: Writes the receiver’s data to a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound
type NSSound struct {
	objectivec.Object
}

// NSSoundFromID constructs a [NSSound] from an objc.ID.
//
// A simple interface for loading and playing audio files.
func NSSoundFromID(id objc.ID) NSSound {
	return NSSound{objectivec.Object{ID: id}}
}
// NOTE: NSSound adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSound] class.
//
// # Detecting When a Sound Finishes Playing
//
//   - [INSSound.Delegate]: The sound’s delegate.
//   - [INSSound.SetDelegate]
//
// # Creating Sounds
//
//   - [INSSound.InitWithContentsOfFileByReference]: Initializes the receiver with the audio data located at a given filepath.
//   - [INSSound.InitWithContentsOfURLByReference]: Initializes the receiver with the audio data located at a given URL.
//   - [INSSound.InitWithData]: Initializes the receiver with a given audio data.
//   - [INSSound.InitWithPasteboard]: Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by [NSSound](<doc://com.apple.appkit/documentation/AppKit/NSSound>). [NSSound] expects the data to have a proper magic number, sound header, and data for the formats it supports.
//
// # Configuring Sounds
//
//   - [INSSound.Name]: The name assigned to the sound.
//   - [INSSound.Volume]: The volume of the sound.
//   - [INSSound.SetVolume]
//   - [INSSound.CurrentTime]: The sound’s playback progress, in seconds.
//   - [INSSound.SetCurrentTime]
//   - [INSSound.Loops]: A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//   - [INSSound.SetLoops]
//   - [INSSound.PlaybackDeviceIdentifier]: Identifies the sound’s output device
//   - [INSSound.SetPlaybackDeviceIdentifier]
//
// # Getting Sound Information
//
//   - [INSSound.Duration]: The duration of the sound, in seconds.
//
// # Playing Sounds
//
//   - [INSSound.Playing]: A Boolean that indicates whether the sound is playing its audio data.
//   - [INSSound.Pause]: Pauses audio playback.
//   - [INSSound.Play]: Initiates audio playback.
//   - [INSSound.Resume]: Resumes audio playback.
//   - [INSSound.Stop]: Concludes audio playback.
//
// # Writing Sounds
//
//   - [INSSound.WriteToPasteboard]: Writes the receiver’s data to a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound
type INSSound interface {
	objectivec.IObject
	NSPasteboardWriting

	// Topic: Detecting When a Sound Finishes Playing

	// The sound’s delegate.
	Delegate() NSSoundDelegate
	SetDelegate(value NSSoundDelegate)

	// Topic: Creating Sounds

	// Initializes the receiver with the audio data located at a given filepath.
	InitWithContentsOfFileByReference(path string, byRef bool) NSSound
	// Initializes the receiver with the audio data located at a given URL.
	InitWithContentsOfURLByReference(url foundation.INSURL, byRef bool) NSSound
	// Initializes the receiver with a given audio data.
	InitWithData(data foundation.INSData) NSSound
	// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by [NSSound](<doc://com.apple.appkit/documentation/AppKit/NSSound>). [NSSound] expects the data to have a proper magic number, sound header, and data for the formats it supports.
	InitWithPasteboard(pasteboard INSPasteboard) NSSound

	// Topic: Configuring Sounds

	// The name assigned to the sound.
	Name() NSSoundName
	// The volume of the sound.
	Volume() float32
	SetVolume(value float32)
	// The sound’s playback progress, in seconds.
	CurrentTime() float64
	SetCurrentTime(value float64)
	// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
	Loops() bool
	SetLoops(value bool)
	// Identifies the sound’s output device
	PlaybackDeviceIdentifier() NSSoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value NSSoundPlaybackDeviceIdentifier)

	// Topic: Getting Sound Information

	// The duration of the sound, in seconds.
	Duration() float64

	// Topic: Playing Sounds

	// A Boolean that indicates whether the sound is playing its audio data.
	Playing() bool
	// Pauses audio playback.
	Pause() bool
	// Initiates audio playback.
	Play() bool
	// Resumes audio playback.
	Resume() bool
	// Concludes audio playback.
	Stop() bool

	// Topic: Writing Sounds

	// Writes the receiver’s data to a pasteboard.
	WriteToPasteboard(pasteboard INSPasteboard)

	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSSound
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSound) Init() NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSound) Autorelease() NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSound creates a new NSSound instance.
func NewNSSound() NSSound {
	class := getNSSoundClass()
	rv := objc.Send[NSSound](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the [NSSound] instance associated with a given name.
//
// name: Name that identifies sound data.
//
// # Return Value
// 
// [NSSound] instance initialized with the sound data identified by
// `soundName`.
//
// # Discussion
// 
// The returned object can be one of the following:
// 
// - One that’s been assigned a name with the [Name] property - One of the
// named system sounds provided by the Application Kit framework
// 
// If there’s no known [NSSound] object with `soundName`, this method tries
// to create one by searching for sound files in the application’s main
// bundle (see [Bundle] for a description of how the bundle’s contents are
// searched). If no sound file can be located in the application main bundle,
// the following directories are searched in order:
// 
// - `~/Library/Sounds` - `/Library/Sounds` - `/Network/Library/Sounds` -
// `/System/Library/Sounds`
// 
// If no data can be found for `soundName`, no object is created, and `nil` is
// returned.
// 
// The preferred way to locate a sound is to pass a name without the file
// extension. See the class description for a list of the supported sound file
// extensions.
//
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(named:)
func NewSoundNamed(name NSSoundName) NSSound {
	rv := objc.Send[objc.ID](objc.ID(getNSSoundClass().class), objc.Sel("soundNamed:"), objc.String(string(name)))
	return NSSoundFromID(rv)
}

// Initializes the receiver with the audio data located at a given filepath.
//
// path: Path to the sound file with which the receiver is to be initialized.
//
// byRef: When [true] only the name of the sound is stored with the [NSSound]
// instance when archived using [encode(with:)]; otherwise the audio data is
// archived along with the instance.
// //
// [encode(with:)]: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOfFile:byReference:)
func NewSoundWithContentsOfFileByReference(path string, byRef bool) NSSound {
	instance := getNSSoundClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:byReference:"), objc.String(path), byRef)
	return NSSoundFromID(rv)
}

// Initializes the receiver with the audio data located at a given URL.
//
// url: URL to the sound file with which the receiver is to be initialized.
//
// byRef: When [true] only the name of the sound is stored with the [NSSound]
// instance when archived using [encode(with:)]; otherwise the audio data is
// archived along with the instance.
// //
// [encode(with:)]: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOf:byReference:)
func NewSoundWithContentsOfURLByReference(url foundation.INSURL, byRef bool) NSSound {
	instance := getNSSoundClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:byReference:"), url, byRef)
	return NSSoundFromID(rv)
}

// Initializes the receiver with a given audio data.
//
// data: Audio data with which the receiver is to be initialized. The data must have
// a proper magic number, sound header, and data for the formats the [NSSound]
// class supports.
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(data:)
func NewSoundWithData(data foundation.INSData) NSSound {
	instance := getNSSoundClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSSoundFromID(rv)
}

// Initializes the receiver with data from a pasteboard. The pasteboard should
// contain a type returned by [NSSound]. [NSSound] expects the data to have a
// proper magic number, sound header, and data for the formats it supports.
//
// pasteboard: The pasteboard containing the audio data with which the receiver is to be
// initialized. The pasteboard must contain a type returned by [NSSound]. The
// contained data must have a proper magic number, sound header, and data for
// the formats the [NSSound] class supports.
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(pasteboard:)
func NewSoundWithPasteboard(pasteboard INSPasteboard) NSSound {
	instance := getNSSoundClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	return NSSoundFromID(rv)
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func NewSoundWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSSound {
	instance := getNSSoundClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSSoundFromID(rv)
}

// Initializes the receiver with the audio data located at a given filepath.
//
// path: Path to the sound file with which the receiver is to be initialized.
//
// byRef: When [true] only the name of the sound is stored with the [NSSound]
// instance when archived using [encode(with:)]; otherwise the audio data is
// archived along with the instance.
// //
// [encode(with:)]: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOfFile:byReference:)
func (s NSSound) InitWithContentsOfFileByReference(path string, byRef bool) NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("initWithContentsOfFile:byReference:"), objc.String(path), byRef)
	return rv
}
// Initializes the receiver with the audio data located at a given URL.
//
// url: URL to the sound file with which the receiver is to be initialized.
//
// byRef: When [true] only the name of the sound is stored with the [NSSound]
// instance when archived using [encode(with:)]; otherwise the audio data is
// archived along with the instance.
// //
// [encode(with:)]: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOf:byReference:)
func (s NSSound) InitWithContentsOfURLByReference(url foundation.INSURL, byRef bool) NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("initWithContentsOfURL:byReference:"), url, byRef)
	return rv
}
// Initializes the receiver with a given audio data.
//
// data: Audio data with which the receiver is to be initialized. The data must have
// a proper magic number, sound header, and data for the formats the [NSSound]
// class supports.
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(data:)
func (s NSSound) InitWithData(data foundation.INSData) NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("initWithData:"), data)
	return rv
}
// Initializes the receiver with data from a pasteboard. The pasteboard should
// contain a type returned by [NSSound]. [NSSound] expects the data to have a
// proper magic number, sound header, and data for the formats it supports.
//
// pasteboard: The pasteboard containing the audio data with which the receiver is to be
// initialized. The pasteboard must contain a type returned by [NSSound]. The
// contained data must have a proper magic number, sound header, and data for
// the formats the [NSSound] class supports.
//
// # Return Value
// 
// Initialized [NSSound] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/init(pasteboard:)
func (s NSSound) InitWithPasteboard(pasteboard INSPasteboard) NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/setName(_:)
func (s NSSound) SetName(string_ NSSoundName) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setName:"), objc.String(string(string_)))
	return rv
}
// Pauses audio playback.
//
// # Return Value
// 
// [true] when playback is paused successfully, [false] when playback is
// already paused or when an error occurred.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/pause()
func (s NSSound) Pause() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("pause"))
	return rv
}
// Initiates audio playback.
//
// # Return Value
// 
// [true] when playback is initiated, [false] when playback is already in
// progress or when an error occurred.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method initiates playback asynchronously and returns control to your
// application. Therefore, your application can continue doing work while the
// audio is playing.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/play()
func (s NSSound) Play() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("play"))
	return rv
}
// Resumes audio playback.
//
// # Return Value
// 
// [true] when playback is resumed, [false] when playback is in progress or
// when an error occurred.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Assumes the receiver has been previously paused by sending it [NSSound].
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/resume()
func (s NSSound) Resume() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("resume"))
	return rv
}
// Concludes audio playback.
//
// # Return Value
// 
// [true] when playback is concluded successfully or if it’s paused, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/stop()
func (s NSSound) Stop() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("stop"))
	return rv
}
// Writes the receiver’s data to a pasteboard.
//
// pasteboard: Pasteboard to which the receiver is to write its data.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/write(to:)
func (s NSSound) WriteToPasteboard(pasteboard INSPasteboard) {
	objc.Send[objc.ID](s.ID, objc.Sel("writeToPasteboard:"), pasteboard)
}
// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func (s NSSound) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSSound {
	rv := objc.Send[NSSound](s.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}
// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
// 
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
// 
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (s NSSound) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}
// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
// 
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
// 
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [PasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (s NSSound) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}
// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
//
// # Discussion
// 
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
func (s NSSound) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](s.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}
func (s NSSound) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Indicates whether the receiver can create an instance of itself from the
// data in a pasteboard.
//
// pasteboard: Pasteboard containing sound data.
//
// # Return Value
// 
// [true] when the receiver can handle the data represented by `pasteboard`;
// [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [NSSound] method is used to find out whether the class can handle the
// data in `pasteboard`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/canInit(with:)
func (_NSSoundClass NSSoundClass) CanInitWithPasteboard(pasteboard INSPasteboard) bool {
	rv := objc.Send[bool](objc.ID(_NSSoundClass.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}
// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
// 
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
// 
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
// 
// # Special Considerations
// 
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSSoundClass NSSoundClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSSoundClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}
// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
// 
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
//
// # Discussion
// 
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
func (_NSSoundClass NSSoundClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSSoundClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}

// The sound’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/delegate
func (s NSSound) Delegate() NSSoundDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSoundDelegateObjectFromID(rv)
}
func (s NSSound) SetDelegate(value NSSoundDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
// The name assigned to the sound.
//
// # Discussion
// 
// The value of this property is `nil` when no name has been assigned.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/name-swift.property
func (s NSSound) Name() NSSoundName {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return NSSoundName(foundation.NSStringFromID(rv).String())
}
// The volume of the sound.
//
// # Discussion
// 
// The valid range is between `0.0` and `1.0`.
// 
// The value of this property does not affect the systemwide volume.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/volume
func (s NSSound) Volume() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("volume"))
	return rv
}
func (s NSSound) SetVolume(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVolume:"), value)
}
// The sound’s playback progress, in seconds.
//
// # Discussion
// 
// Sounds start with `currentTime == 0` and end with `currentTime == ([
// duration] - 1)`.
// 
// This property is not archived, copied, or stored on the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/currentTime
func (s NSSound) CurrentTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("currentTime"))
	return rv
}
func (s NSSound) SetCurrentTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCurrentTime:"), value)
}
// A Boolean that indicates whether the sound restarts playback when it
// reaches the end of its content.
//
// # Discussion
// 
// When the value of this property is [true], the sounds restarts playback
// when it finishes and does not send [SoundDidFinishPlaying] to its delegate
// when it reaches the end of its content and restarts playback. The default
// value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/loops
func (s NSSound) Loops() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("loops"))
	return rv
}
func (s NSSound) SetLoops(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setLoops:"), value)
}
// Identifies the sound’s output device
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/playbackDeviceIdentifier-swift.property
func (s NSSound) PlaybackDeviceIdentifier() NSSoundPlaybackDeviceIdentifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("playbackDeviceIdentifier"))
	return NSSoundPlaybackDeviceIdentifier(foundation.NSStringFromID(rv).String())
}
func (s NSSound) SetPlaybackDeviceIdentifier(value NSSoundPlaybackDeviceIdentifier) {
	objc.Send[struct{}](s.ID, objc.Sel("setPlaybackDeviceIdentifier:"), objc.String(string(value)))
}
// The duration of the sound, in seconds.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/duration
func (s NSSound) Duration() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("duration"))
	return rv
}
// A Boolean that indicates whether the sound is playing its audio data.
//
// # Discussion
// 
// The value of this property is [true] when the receiver is playing its audio
// data.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/isPlaying
func (s NSSound) Playing() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPlaying"))
	return rv
}

// Provides the file types the [NSSound] class understands.
//
// # Return Value
// 
// Array of UTIs identifying the file types the [NSSound] class understands.
//
// See: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredTypes
func (_NSSoundClass NSSoundClass) SoundUnfilteredTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSSoundClass.class), objc.Sel("soundUnfilteredTypes"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSPasteboardWriting
			

