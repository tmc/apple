// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVFragmentedMovieMinder] class.
var (
	_AVFragmentedMovieMinderClass     AVFragmentedMovieMinderClass
	_AVFragmentedMovieMinderClassOnce sync.Once
)

func getAVFragmentedMovieMinderClass() AVFragmentedMovieMinderClass {
	_AVFragmentedMovieMinderClassOnce.Do(func() {
		_AVFragmentedMovieMinderClass = AVFragmentedMovieMinderClass{class: objc.GetClass("AVFragmentedMovieMinder")}
	})
	return _AVFragmentedMovieMinderClass
}

// GetAVFragmentedMovieMinderClass returns the class object for AVFragmentedMovieMinder.
func GetAVFragmentedMovieMinderClass() AVFragmentedMovieMinderClass {
	return getAVFragmentedMovieMinderClass()
}

type AVFragmentedMovieMinderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFragmentedMovieMinderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedMovieMinderClass) Alloc() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that checks whether a fragmented movie appends additional movie
// fragments.
//
// # Overview
// 
// This class is identical to [AVFragmentedAssetMinder] except that it’s
// capable of minding only assets of type [AVFragmentedMovie].
//
// # Creating a movie minder
//
//   - [AVFragmentedMovieMinder.InitWithMovieMindingInterval]: Creates a movie minder and adds a movie with a minding interval.
//
// # Adding and removing movies
//
//   - [AVFragmentedMovieMinder.Movies]: An array containing the fragmented movie objects being minded.
//   - [AVFragmentedMovieMinder.AddFragmentedMovie]: Adds a fragmented movie to the array of movies being minded.
//   - [AVFragmentedMovieMinder.RemoveFragmentedMovie]: Removes a fragmented movie from the array of movies being minded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder
type AVFragmentedMovieMinder struct {
	AVFragmentedAssetMinder
}

// AVFragmentedMovieMinderFromID constructs a [AVFragmentedMovieMinder] from an objc.ID.
//
// An object that checks whether a fragmented movie appends additional movie
// fragments.
func AVFragmentedMovieMinderFromID(id objc.ID) AVFragmentedMovieMinder {
	return AVFragmentedMovieMinder{AVFragmentedAssetMinder: AVFragmentedAssetMinderFromID(id)}
}
// NOTE: AVFragmentedMovieMinder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedMovieMinder] class.
//
// # Creating a movie minder
//
//   - [IAVFragmentedMovieMinder.InitWithMovieMindingInterval]: Creates a movie minder and adds a movie with a minding interval.
//
// # Adding and removing movies
//
//   - [IAVFragmentedMovieMinder.Movies]: An array containing the fragmented movie objects being minded.
//   - [IAVFragmentedMovieMinder.AddFragmentedMovie]: Adds a fragmented movie to the array of movies being minded.
//   - [IAVFragmentedMovieMinder.RemoveFragmentedMovie]: Removes a fragmented movie from the array of movies being minded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder
type IAVFragmentedMovieMinder interface {
	IAVFragmentedAssetMinder

	// Topic: Creating a movie minder

	// Creates a movie minder and adds a movie with a minding interval.
	InitWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) AVFragmentedMovieMinder

	// Topic: Adding and removing movies

	// An array containing the fragmented movie objects being minded.
	Movies() []AVFragmentedMovie
	// Adds a fragmented movie to the array of movies being minded.
	AddFragmentedMovie(movie IAVFragmentedMovie)
	// Removes a fragmented movie from the array of movies being minded.
	RemoveFragmentedMovie(movie IAVFragmentedMovie)
}

// Init initializes the instance.
func (f AVFragmentedMovieMinder) Init() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedMovieMinder) Autorelease() AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedMovieMinder creates a new AVFragmentedMovieMinder instance.
func NewAVFragmentedMovieMinder() AVFragmentedMovieMinder {
	class := getAVFragmentedMovieMinderClass()
	rv := objc.Send[AVFragmentedMovieMinder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fragmented asset minder that monitors the specified asset at the
// indicated minding interval.
//
// asset: The fragmented asset added to the fragmented asset minder.
//
// mindingInterval: The amount of time between checking to see if the system appended
// additional fragments to the minded asset.
//
// # Return Value
// 
// The new fragmented asset minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/init(asset:mindingInterval:)
func NewFragmentedMovieMinderWithAssetMindingInterval(asset IAVAsset, mindingInterval float64) AVFragmentedMovieMinder {
	instance := getAVFragmentedMovieMinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:mindingInterval:"), asset, mindingInterval)
	return AVFragmentedMovieMinderFromID(rv)
}

// Creates a movie minder and adds a movie with a minding interval.
//
// movie: The fragmented movie object added to the movie minder.
//
// mindingInterval: The initial minding interval for the movie minder.
//
// # Return Value
// 
// A new [AVFragmentedMovieMinder] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/init(movie:mindingInterval:)
func NewFragmentedMovieMinderWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) AVFragmentedMovieMinder {
	instance := getAVFragmentedMovieMinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMovie:mindingInterval:"), movie, mindingInterval)
	return AVFragmentedMovieMinderFromID(rv)
}

// Creates a movie minder and adds a movie with a minding interval.
//
// movie: The fragmented movie object added to the movie minder.
//
// mindingInterval: The initial minding interval for the movie minder.
//
// # Return Value
// 
// A new [AVFragmentedMovieMinder] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/init(movie:mindingInterval:)
func (f AVFragmentedMovieMinder) InitWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) AVFragmentedMovieMinder {
	rv := objc.Send[AVFragmentedMovieMinder](f.ID, objc.Sel("initWithMovie:mindingInterval:"), movie, mindingInterval)
	return rv
}
// Adds a fragmented movie to the array of movies being minded.
//
// movie: The fragmented movie added to the minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/add(_:)
func (f AVFragmentedMovieMinder) AddFragmentedMovie(movie IAVFragmentedMovie) {
	objc.Send[objc.ID](f.ID, objc.Sel("addFragmentedMovie:"), movie)
}
// Removes a fragmented movie from the array of movies being minded.
//
// movie: The fragmented movie removed from the minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/remove(_:)
func (f AVFragmentedMovieMinder) RemoveFragmentedMovie(movie IAVFragmentedMovie) {
	objc.Send[objc.ID](f.ID, objc.Sel("removeFragmentedMovie:"), movie)
}

// Creates a movie minder and adds a movie with a minding interval.
//
// movie: The fragmented movie object added to the movie minder.
//
// mindingInterval: The initial minding interval for the movie minder.
//
// # Return Value
// 
// A new [AVFragmentedMovieMinder] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/fragmentedMovieMinderWithMovie:mindingInterval:
func (_AVFragmentedMovieMinderClass AVFragmentedMovieMinderClass) FragmentedMovieMinderWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) AVFragmentedMovieMinder {
	rv := objc.Send[objc.ID](objc.ID(_AVFragmentedMovieMinderClass.class), objc.Sel("fragmentedMovieMinderWithMovie:mindingInterval:"), movie, mindingInterval)
	return AVFragmentedMovieMinderFromID(rv)
}

// An array containing the fragmented movie objects being minded.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/movies
func (f AVFragmentedMovieMinder) Movies() []AVFragmentedMovie {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("movies"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVFragmentedMovie {
		return AVFragmentedMovieFromID(id)
	})
}

