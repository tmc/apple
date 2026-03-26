// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMovie] class.
var (
	_AVMovieClass     AVMovieClass
	_AVMovieClassOnce sync.Once
)

func getAVMovieClass() AVMovieClass {
	_AVMovieClassOnce.Do(func() {
		_AVMovieClass = AVMovieClass{class: objc.GetClass("AVMovie")}
	})
	return _AVMovieClass
}

// GetAVMovieClass returns the class object for AVMovie.
func GetAVMovieClass() AVMovieClass {
	return getAVMovieClass()
}

type AVMovieClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMovieClass) Alloc() AVMovie {
	rv := objc.Send[AVMovie](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an audiovisual container that conforms to the
// QuickTime movie file format or a related format like MPEG-4.
//
// # Overview
// 
// [AVMovie] supports operations involving the format-specific portions of the
// QuickTime movie model that [AVAsset] doesn’t support. For instance,
// retrieving the movie header from an existing QuickTime movie file. You can
// also use [AVMovie] to write a movie header into a new file, thereby
// creating a reference movie.
//
// # Creating a movie
//
//   - [AVMovie.InitWithURLOptions]: Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//   - [AVMovie.InitWithDataOptions]: Creates a movie object from a movie file’s data.
//
// # Creating and writing headers
//
//   - [AVMovie.IsCompatibleWithFileType]: Returns a Boolean value that indicates whether the system can create a movie header of the specified type.
//   - [AVMovie.MovieHeaderWithFileTypeError]: Creates a header for a movie for the specified file type.
//   - [AVMovie.WriteMovieHeaderToURLFileTypeOptionsError]: Writes the movie header to the specified URL.
//
// # Determining fragment support
//
//   - [AVMovie.CanContainMovieFragments]: A Boolean value that indicates whether fragments can extend the movie file.
//   - [AVMovie.ContainsMovieFragments]: A Boolean value that indicates whether at least one movie fragment extends the movie file.
//
// # Accessing movie information
//
//   - [AVMovie.URL]: A URL to a QuickTime or ISO base media file.
//   - [AVMovie.Data]: A data object that contains the movie file’s data.
//
// # Accessing tracks
//
//   - [AVMovie.Tracks]: The tracks that a movie contains.
//
// # Accessing data storage
//
//   - [AVMovie.DefaultMediaDataStorage]: The default storage container for media data added to a movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie
type AVMovie struct {
	AVAsset
}

// AVMovieFromID constructs a [AVMovie] from an objc.ID.
//
// An object that represents an audiovisual container that conforms to the
// QuickTime movie file format or a related format like MPEG-4.
func AVMovieFromID(id objc.ID) AVMovie {
	return AVMovie{AVAsset: AVAssetFromID(id)}
}
// NOTE: AVMovie adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMovie] class.
//
// # Creating a movie
//
//   - [IAVMovie.InitWithURLOptions]: Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//   - [IAVMovie.InitWithDataOptions]: Creates a movie object from a movie file’s data.
//
// # Creating and writing headers
//
//   - [IAVMovie.IsCompatibleWithFileType]: Returns a Boolean value that indicates whether the system can create a movie header of the specified type.
//   - [IAVMovie.MovieHeaderWithFileTypeError]: Creates a header for a movie for the specified file type.
//   - [IAVMovie.WriteMovieHeaderToURLFileTypeOptionsError]: Writes the movie header to the specified URL.
//
// # Determining fragment support
//
//   - [IAVMovie.CanContainMovieFragments]: A Boolean value that indicates whether fragments can extend the movie file.
//   - [IAVMovie.ContainsMovieFragments]: A Boolean value that indicates whether at least one movie fragment extends the movie file.
//
// # Accessing movie information
//
//   - [IAVMovie.URL]: A URL to a QuickTime or ISO base media file.
//   - [IAVMovie.Data]: A data object that contains the movie file’s data.
//
// # Accessing tracks
//
//   - [IAVMovie.Tracks]: The tracks that a movie contains.
//
// # Accessing data storage
//
//   - [IAVMovie.DefaultMediaDataStorage]: The default storage container for media data added to a movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie
type IAVMovie interface {
	IAVAsset

	// Topic: Creating a movie

	// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file.
	InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMovie
	// Creates a movie object from a movie file’s data.
	InitWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVMovie

	// Topic: Creating and writing headers

	// Returns a Boolean value that indicates whether the system can create a movie header of the specified type.
	IsCompatibleWithFileType(fileType AVFileType) bool
	// Creates a header for a movie for the specified file type.
	MovieHeaderWithFileTypeError(fileType AVFileType) (foundation.INSData, error)
	// Writes the movie header to the specified URL.
	WriteMovieHeaderToURLFileTypeOptionsError(URL foundation.INSURL, fileType AVFileType, options AVMovieWritingOptions) (bool, error)

	// Topic: Determining fragment support

	// A Boolean value that indicates whether fragments can extend the movie file.
	CanContainMovieFragments() bool
	// A Boolean value that indicates whether at least one movie fragment extends the movie file.
	ContainsMovieFragments() bool

	// Topic: Accessing movie information

	// A URL to a QuickTime or ISO base media file.
	URL() foundation.INSURL
	// A data object that contains the movie file’s data.
	Data() foundation.INSData

	// Topic: Accessing tracks

	// The tracks that a movie contains.
	Tracks() []AVMovieTrack

	// Topic: Accessing data storage

	// The default storage container for media data added to a movie.
	DefaultMediaDataStorage() IAVMediaDataStorage
}

// Init initializes the instance.
func (m AVMovie) Init() AVMovie {
	rv := objc.Send[AVMovie](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMovie) Autorelease() AVMovie {
	rv := objc.Send[AVMovie](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMovie creates a new AVMovie instance.
func NewAVMovie() AVMovie {
	class := getAVMovieClass()
	rv := objc.Send[AVMovie](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewMovieAssetWithURL(URL foundation.INSURL) AVMovie {
	rv := objc.Send[objc.ID](objc.ID(getAVMovieClass().class), objc.Sel("assetWithURL:"), URL)
	return AVMovieFromID(rv)
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
func NewMovieWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVMovie {
	instance := getAVMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:"), data, options)
	return AVMovieFromID(rv)
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
func NewMovieWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMovie {
	instance := getAVMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVMovieFromID(rv)
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
func (m AVMovie) InitWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMovie {
	rv := objc.Send[AVMovie](m.ID, objc.Sel("initWithURL:options:"), URL, options)
	return rv
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
func (m AVMovie) InitWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVMovie {
	rv := objc.Send[AVMovie](m.ID, objc.Sel("initWithData:options:"), data, options)
	return rv
}
// Returns a Boolean value that indicates whether the system can create a
// movie header of the specified type.
//
// fileType: A file type to test.
//
// # Return Value
// 
// [true] if the movie only contains tracks whose media types are allowed by
// the specified file type; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/is(compatibleWithFileType:)
func (m AVMovie) IsCompatibleWithFileType(fileType AVFileType) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isCompatibleWithFileType:"), objc.String(string(fileType)))
	return rv
}
// Creates a header for a movie for the specified file type.
//
// fileType: A UTI that indicates the specific file format for the movie header.
//
// # Return Value
// 
// An [NSData] object containing the movie header.
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// # Discussion
// 
// The created movie header is a pure reference movie, with no base URL,
// suitable for use on the pasteboard.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/makeMovieHeader(fileType:)
func (m AVMovie) MovieHeaderWithFileTypeError(fileType AVFileType) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("movieHeaderWithFileType:error:"), objc.String(string(fileType)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}
// Writes the movie header to the specified URL.
//
// URL: The URL indicating where to write the movie header.
//
// fileType: A UTI that indicates the specific file format for the movie header.
//
// options: The [AVMovieWritingOptions] constants whose bits specify the options for
// writing the movie header.
// //
// [AVMovieWritingOptions]: https://developer.apple.com/documentation/AVFoundation/AVMovieWritingOptions
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/writeHeader(to:fileType:options:)
func (m AVMovie) WriteMovieHeaderToURLFileTypeOptionsError(URL foundation.INSURL, fileType AVFileType, options AVMovieWritingOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeMovieHeaderToURL:fileType:options:error:"), URL, objc.String(string(fileType)), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeMovieHeaderToURL:fileType:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the file types that a movie supports.
//
// # Return Value
// 
// An array of supported file types.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieTypes()
func (_AVMovieClass AVMovieClass) MovieTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMovieClass.class), objc.Sel("movieTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns a new movie object from a movie file’s data.
//
// data: A data object that contains a movie header.
//
// options: A dictionary of options to use to initialize the movie.
//
// # Return Value
// 
// A movie object.
//
// # Discussion
// 
// Use this method to create movies from movie headers that aren’t stored in
// files, which can include movies that the pasteboard contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieWithData:options:
func (_AVMovieClass AVMovieClass) MovieWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVMovie {
	rv := objc.Send[objc.ID](objc.ID(_AVMovieClass.class), objc.Sel("movieWithData:options:"), data, options)
	return AVMovieFromID(rv)
}
// Returns a new movie object from a movie header stored in a QuickTime movie
// file of ISO base media file.
//
// URL: A URL that points to a file containing a movie header.
//
// options: A dictionary of initialization options with which to create the movie.
//
// # Return Value
// 
// A movie object.
//
// # Discussion
// 
// Upon creation, the values of the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieWithURL:options:
func (_AVMovieClass AVMovieClass) MovieWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMovie {
	rv := objc.Send[objc.ID](objc.ID(_AVMovieClass.class), objc.Sel("movieWithURL:options:"), URL, options)
	return AVMovieFromID(rv)
}

// A Boolean value that indicates whether fragments can extend the movie file.
//
// # Discussion
// 
// The value of this property is [YES] if an `mvex` box is present in the
// `moov` box. The `mvex` box is necessary to signal the possible presence of
// later `moof` boxes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/canContainMovieFragments
func (m AVMovie) CanContainMovieFragments() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("canContainMovieFragments"))
	return rv
}
// A Boolean value that indicates whether at least one movie fragment extends
// the movie file.
//
// # Discussion
// 
// This property is [YES] if [CanContainMovieFragments] is [YES] and at least
// one `moof` box is present after the `moov` box.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/containsMovieFragments
func (m AVMovie) ContainsMovieFragments() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("containsMovieFragments"))
	return rv
}
// A URL to a QuickTime or ISO base media file.
//
// # Discussion
// 
// The value is `nil` if you didn’t initialize the movie with a URL.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/url
func (m AVMovie) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// A data object that contains the movie file’s data.
//
// # Discussion
// 
// The value is `nil` if you didn’t initialize the movie with data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/data
func (m AVMovie) Data() foundation.INSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The tracks that a movie contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/tracks
func (m AVMovie) Tracks() []AVMovieTrack {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("tracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMovieTrack {
		return AVMovieTrackFromID(id)
	})
}
// The default storage container for media data added to a movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/defaultMediaDataStorage
func (m AVMovie) DefaultMediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultMediaDataStorage"))
	return AVMediaDataStorageFromID(objc.ID(rv))
}

