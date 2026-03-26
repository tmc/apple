// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVFragmentedMovie] class.
var (
	_AVFragmentedMovieClass     AVFragmentedMovieClass
	_AVFragmentedMovieClassOnce sync.Once
)

func getAVFragmentedMovieClass() AVFragmentedMovieClass {
	_AVFragmentedMovieClassOnce.Do(func() {
		_AVFragmentedMovieClass = AVFragmentedMovieClass{class: objc.GetClass("AVFragmentedMovie")}
	})
	return _AVFragmentedMovieClass
}

// GetAVFragmentedMovieClass returns the class object for AVFragmentedMovie.
func GetAVFragmentedMovieClass() AVFragmentedMovieClass {
	return getAVFragmentedMovieClass()
}

type AVFragmentedMovieClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedMovieClass) Alloc() AVFragmentedMovie {
	rv := objc.Send[AVFragmentedMovie](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a fragmented movie file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie
type AVFragmentedMovie struct {
	AVMovie
}

// AVFragmentedMovieFromID constructs a [AVFragmentedMovie] from an objc.ID.
//
// An object that represents a fragmented movie file.
func AVFragmentedMovieFromID(id objc.ID) AVFragmentedMovie {
	return AVFragmentedMovie{AVMovie: AVMovieFromID(id)}
}
// NOTE: AVFragmentedMovie adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedMovie] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie
type IAVFragmentedMovie interface {
	IAVMovie
	AVFragmentMinding
}

// Init initializes the instance.
func (f AVFragmentedMovie) Init() AVFragmentedMovie {
	rv := objc.Send[AVFragmentedMovie](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedMovie) Autorelease() AVFragmentedMovie {
	rv := objc.Send[AVFragmentedMovie](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedMovie creates a new AVFragmentedMovie instance.
func NewAVFragmentedMovie() AVFragmentedMovie {
	class := getAVFragmentedMovieClass()
	rv := objc.Send[AVFragmentedMovie](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewFragmentedMovieAssetWithURL(URL foundation.INSURL) AVFragmentedMovie {
	rv := objc.Send[objc.ID](objc.ID(getAVFragmentedMovieClass().class), objc.Sel("assetWithURL:"), URL)
	return AVFragmentedMovieFromID(rv)
}

// Creates a movie object from a movie file’s data.
//
// data: A data object that contains a movie header.
//
// options: A dictionary of options to use to initialize the movie.
//
// # Discussion
// 
// Use this method to create movies from movie headers that aren’t stored in
// files, which can include movies that the pasteboard contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(data:options:)
func NewFragmentedMovieWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVFragmentedMovie {
	instance := getAVFragmentedMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:"), data, options)
	return AVFragmentedMovieFromID(rv)
}

// Creates a movie object from a movie header stored in a QuickTime movie file
// of ISO base media file.
//
// URL: A URL that points to a file containing a movie header.
//
// options: A dictionary of options to use to initialize the movie.
//
// # Discussion
// 
// Upon creation, the values of the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(url:options:)
func NewFragmentedMovieWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVFragmentedMovie {
	instance := getAVFragmentedMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVFragmentedMovieFromID(rv)
}

// A Boolean value that indicates whether an asset that supports fragment
// minding is currently associated with a fragment minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding/isAssociatedWithFragmentMinder
func (f AVFragmentedMovie) IsAssociatedWithFragmentMinder() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isAssociatedWithFragmentMinder"))
	return rv
}

// A Boolean value that indicates whether an asset that supports fragment
// minding is currently associated with a fragment minder.
//
// # Discussion
// 
// Only asset objects associated with a fragment minder post change
// notifications.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding/isAssociatedWithFragmentMinder
func (f AVFragmentedMovie) AssociatedWithFragmentMinder() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isAssociatedWithFragmentMinder"))
	return rv
}

			// Protocol methods for AVFragmentMinding
			

