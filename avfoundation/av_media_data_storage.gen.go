// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaDataStorage] class.
var (
	_AVMediaDataStorageClass     AVMediaDataStorageClass
	_AVMediaDataStorageClassOnce sync.Once
)

func getAVMediaDataStorageClass() AVMediaDataStorageClass {
	_AVMediaDataStorageClassOnce.Do(func() {
		_AVMediaDataStorageClass = AVMediaDataStorageClass{class: objc.GetClass("AVMediaDataStorage")}
	})
	return _AVMediaDataStorageClass
}

// GetAVMediaDataStorageClass returns the class object for AVMediaDataStorage.
func GetAVMediaDataStorageClass() AVMediaDataStorageClass {
	return getAVMediaDataStorageClass()
}

type AVMediaDataStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaDataStorageClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaDataStorageClass) Alloc() AVMediaDataStorage {
	rv := objc.Send[AVMediaDataStorage](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the media sample data storage file.
//
// # Creating media data storage
//
//   - [AVMediaDataStorage.InitWithURLOptions]: Creates a media data storage object associated with a file URL.
//
// # Accessing the URL
//
//   - [AVMediaDataStorage.URL]: Returns the URL used to initialize the receiver.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage
type AVMediaDataStorage struct {
	objectivec.Object
}

// AVMediaDataStorageFromID constructs a [AVMediaDataStorage] from an objc.ID.
//
// An object that represents the media sample data storage file.
func AVMediaDataStorageFromID(id objc.ID) AVMediaDataStorage {
	return AVMediaDataStorage{objectivec.Object{ID: id}}
}
// NOTE: AVMediaDataStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaDataStorage] class.
//
// # Creating media data storage
//
//   - [IAVMediaDataStorage.InitWithURLOptions]: Creates a media data storage object associated with a file URL.
//
// # Accessing the URL
//
//   - [IAVMediaDataStorage.URL]: Returns the URL used to initialize the receiver.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage
type IAVMediaDataStorage interface {
	objectivec.IObject

	// Topic: Creating media data storage

	// Creates a media data storage object associated with a file URL.
	InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMediaDataStorage

	// Topic: Accessing the URL

	// Returns the URL used to initialize the receiver.
	URL() foundation.INSURL
}

// Init initializes the instance.
func (m AVMediaDataStorage) Init() AVMediaDataStorage {
	rv := objc.Send[AVMediaDataStorage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaDataStorage) Autorelease() AVMediaDataStorage {
	rv := objc.Send[AVMediaDataStorage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaDataStorage creates a new AVMediaDataStorage instance.
func NewAVMediaDataStorage() AVMediaDataStorage {
	class := getAVMediaDataStorageClass()
	rv := objc.Send[AVMediaDataStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a media data storage object associated with a file URL.
//
// URL: The URL specifying where sample data added to a movie or track is written.
//
// options: A dictionary object containing keys for specifying initialization options.
// No keys are currently defined.
//
// # Return Value
// 
// An [AVMediaDataStorage] object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/init(url:options:)
func NewMediaDataStorageWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMediaDataStorage {
	instance := getAVMediaDataStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVMediaDataStorageFromID(rv)
}

// Creates a media data storage object associated with a file URL.
//
// URL: The URL specifying where sample data added to a movie or track is written.
//
// options: A dictionary object containing keys for specifying initialization options.
// No keys are currently defined.
//
// # Return Value
// 
// An [AVMediaDataStorage] object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/init(url:options:)
func (m AVMediaDataStorage) InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMediaDataStorage {
	rv := objc.Send[AVMediaDataStorage](m.ID, objc.Sel("initWithURL:options:"), URL, options)
	return rv
}
// Returns the URL used to initialize the receiver.
//
// # Return Value
// 
// The [NSURL] object used to initialize the receiver. This value can be
// `nil`.
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/url()
func (m AVMediaDataStorage) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(rv)
}

