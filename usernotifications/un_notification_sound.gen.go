// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationSound] class.
var (
	_UNNotificationSoundClass     UNNotificationSoundClass
	_UNNotificationSoundClassOnce sync.Once
)

func getUNNotificationSoundClass() UNNotificationSoundClass {
	_UNNotificationSoundClassOnce.Do(func() {
		_UNNotificationSoundClass = UNNotificationSoundClass{class: objc.GetClass("UNNotificationSound")}
	})
	return _UNNotificationSoundClass
}

// GetUNNotificationSoundClass returns the class object for UNNotificationSound.
func GetUNNotificationSoundClass() UNNotificationSoundClass {
	return getUNNotificationSoundClass()
}

type UNNotificationSoundClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationSoundClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationSoundClass) Alloc() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The sound played upon delivery of a notification.
//
// # Overview
// 
// Create a [UNNotificationSound] object when you want the system to play a
// specific sound when it delivers with your notification. To play the default
// system sound, create your sound object using the [UNNotificationSound.DefaultSound] method. If
// you want to play a custom sound, create a new sound object and specify the
// name of the audio file that you want to play.
// 
// For local notifications, assign the sound object to the [UNNotificationSound.Sound] property of
// your [UNMutableNotificationContent] object. For a remote notification,
// assign the name of your sound file to the `sound` key in the `aps`
// dictionary. You can also use a notification service app extension to add a
// sound file to a notification shortly before delivery. In your extension,
// create a [UNNotificationSound] object and add it to your notification
// content in the same way that you’d for a local notification.
// 
// Audio files must already be on the user’s device before the system can
// play them. If you use a predefined set of sounds for your notifications,
// include the audio files in your app’s bundle. For all other sounds, the
// [UNNotificationSound] object looks only in the following locations:
// 
// - The `/Library/Sounds` directory of the app’s container directory. - The
// `/Library/Sounds` directory of one of the app’s shared group container
// directories. - The main bundle of the current executable.
// 
// # Prepare Sound Resources
// 
// The system sound facility plays custom alert sounds, so they must be in one
// of the following audio data formats:
// 
// - Linear PCM
// - MA4 (IMA/ADPCM)
// - µLaw
// - aLaw
// 
// You can package the audio data in an `aiff`, `wav`, or `caf` file. Sound
// files must be less than 30 seconds in length. If the sound file is longer
// than 30 seconds, the system plays the default sound instead.
// 
// You can use the `afconvert` command-line tool to convert sounds. For
// example, to convert the system sound `Submarine.Aiff()` to IMA4 audio in a
// CAF file, use the following command in Terminal:
// 
// `afconvert /System/Library/Sounds/Submarine.Aiff() ~/Desktop/sub.Caf() -d
// ima4 -f caff -v`
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound
type UNNotificationSound struct {
	objectivec.Object
}

// UNNotificationSoundFromID constructs a [UNNotificationSound] from an objc.ID.
//
// The sound played upon delivery of a notification.
func UNNotificationSoundFromID(id objc.ID) UNNotificationSound {
	return UNNotificationSound{objectivec.Object{ID: id}}
}
// NOTE: UNNotificationSound adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationSound] class.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound
type IUNNotificationSound interface {
	objectivec.IObject

	// The sound that plays when the system delivers the notification.
	Sound() IUNNotificationSound
	SetSound(value IUNNotificationSound)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationSound) Init() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationSound) Autorelease() UNNotificationSound {
	rv := objc.Send[UNNotificationSound](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationSound creates a new UNNotificationSound instance.
func NewUNNotificationSound() UNNotificationSound {
	class := getUNNotificationSoundClass()
	rv := objc.Send[UNNotificationSound](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a sound object that represents a custom sound file.
//
// name: The name of the sound file to play. This parameter must not be `nil`.
//
// # Return Value
// 
// A sound object representing the custom sound.
//
// # Discussion
// 
// This method searches for sound files in the following locations, in order:
// 
// - The `/Library/Sounds` directory, where is the app’s container
// directory. - The `/Library/Sounds` directory, where is one of the app’s
// shared group container directories. For information about how to configure
// group containers for your app, see [Configure app groups]. - The main
// bundle of the current executable.
// 
// The method chooses the first file it finds with the specified name.
//
// [Configure app groups]: https://help.apple.com/xcode/mac/current/#/dev8dd3880fe
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/init(named:)
func NewUNNotificationSoundNamed(name UNNotificationSoundName) UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationSoundClass().class), objc.Sel("soundNamed:"), objc.String(string(name)))
	return UNNotificationSoundFromID(rv)
}

func (u UNNotificationSound) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a sound object that plays the default critical alert sound at the
// volume you specify.
//
// volume: The volume must be a value between 0.0 and 1.0.
//
// # Return Value
// 
// A sound object representing the default critical alert sound at the
// specified volume.
//
// # Discussion
// 
// Critical alerts ignore the mute switch and Do Not Disturb. They require a
// special entitlement issued by Apple.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultCriticalSound(withAudioVolume:)
func (_UNNotificationSoundClass UNNotificationSoundClass) DefaultCriticalSoundWithAudioVolume(volume float32) UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("defaultCriticalSoundWithAudioVolume:"), volume)
	return UNNotificationSoundFromID(rv)
}
// Creates a custom sound object for critical alerts.
//
// name: The name of the sound file to play. This file must be located in the
// current executable’s main bundle or in the `Library/Sounds` directory of
// the current app container directory. If files exist at both locations, the
// system uses the file from the `Library/Sounds` directory. This parameter
// must not be `nil`.
//
// # Return Value
// 
// A sound object representing a custom critical alert sound.
//
// # Discussion
// 
// Critical alerts ignore the mute switch and Do Not Disturb. They require a
// special entitlement issued by Apple.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/criticalSoundNamed(_:)
func (_UNNotificationSoundClass UNNotificationSoundClass) CriticalSoundNamed(name UNNotificationSoundName) UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("criticalSoundNamed:"), objc.String(string(name)))
	return UNNotificationSoundFromID(rv)
}
// Creates a custom sound object for critical alerts with the volume you
// specify.
//
// name: The name of the sound file to play. This file must be located in the
// current executable’s main bundle or in the `Library/Sounds` directory of
// the current app container directory. If files exist at both locations, the
// system uses the file from the `Library/Sounds` directory. This parameter
// must not be `nil`.
//
// volume: The volume must be a value between 0.0 and 1.0.
//
// # Return Value
// 
// A sound object representing a custom critical alert sound at the specified
// volume.
//
// # Discussion
// 
// Critical alerts ignore the mute switch and Do Not Disturb. They require a
// special entitlement issued by Apple.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/criticalSoundNamed(_:withAudioVolume:)
func (_UNNotificationSoundClass UNNotificationSoundClass) CriticalSoundNamedWithAudioVolume(name UNNotificationSoundName, volume float32) UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("criticalSoundNamed:withAudioVolume:"), objc.String(string(name)), volume)
	return UNNotificationSoundFromID(rv)
}

// The sound that plays when the system delivers the notification.
//
// See: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/sound
func (u UNNotificationSound) Sound() IUNNotificationSound {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sound"))
	return UNNotificationSoundFromID(objc.ID(rv))
}
func (u UNNotificationSound) SetSound(value IUNNotificationSound) {
	objc.Send[struct{}](u.ID, objc.Sel("setSound:"), value)
}

// Returns an object representing the default sound for notifications.
//
// # Return Value
// 
// A sound object that represents the default notification sound.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/default
func (_UNNotificationSoundClass UNNotificationSoundClass) DefaultSound() UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("defaultSound"))
	return UNNotificationSoundFromID(objc.ID(rv))
}
// The default sound used for critical alerts.
//
// # Discussion
// 
// Critical alerts ingore the mute switch and Do Not Disturb. They require a
// special entitlement issued by Apple.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultCritical
func (_UNNotificationSoundClass UNNotificationSoundClass) DefaultCriticalSound() UNNotificationSound {
	rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("defaultCriticalSound"))
	return UNNotificationSoundFromID(objc.ID(rv))
}

