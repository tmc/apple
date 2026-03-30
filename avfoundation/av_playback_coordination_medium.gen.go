// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlaybackCoordinationMedium] class.
var (
	_AVPlaybackCoordinationMediumClass     AVPlaybackCoordinationMediumClass
	_AVPlaybackCoordinationMediumClassOnce sync.Once
)

func getAVPlaybackCoordinationMediumClass() AVPlaybackCoordinationMediumClass {
	_AVPlaybackCoordinationMediumClassOnce.Do(func() {
		_AVPlaybackCoordinationMediumClass = AVPlaybackCoordinationMediumClass{class: objc.GetClass("AVPlaybackCoordinationMedium")}
	})
	return _AVPlaybackCoordinationMediumClass
}

// GetAVPlaybackCoordinationMediumClass returns the class object for AVPlaybackCoordinationMedium.
func GetAVPlaybackCoordinationMediumClass() AVPlaybackCoordinationMediumClass {
	return getAVPlaybackCoordinationMediumClass()
}

type AVPlaybackCoordinationMediumClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlaybackCoordinationMediumClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlaybackCoordinationMediumClass) Alloc() AVPlaybackCoordinationMedium {
	rv := objc.Send[AVPlaybackCoordinationMedium](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Managing playback coordinators
//
//   - [AVPlaybackCoordinationMedium.ConnectedPlaybackCoordinators]: All playback coordinators that are connected to the coordination medium.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium
type AVPlaybackCoordinationMedium struct {
	objectivec.Object
}

// AVPlaybackCoordinationMediumFromID constructs a [AVPlaybackCoordinationMedium] from an objc.ID.
func AVPlaybackCoordinationMediumFromID(id objc.ID) AVPlaybackCoordinationMedium {
	return AVPlaybackCoordinationMedium{objectivec.Object{ID: id}}
}

// NOTE: AVPlaybackCoordinationMedium adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlaybackCoordinationMedium] class.
//
// # Managing playback coordinators
//
//   - [IAVPlaybackCoordinationMedium.ConnectedPlaybackCoordinators]: All playback coordinators that are connected to the coordination medium.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium
type IAVPlaybackCoordinationMedium interface {
	objectivec.IObject

	// Topic: Managing playback coordinators

	// All playback coordinators that are connected to the coordination medium.
	ConnectedPlaybackCoordinators() []AVPlayerPlaybackCoordinator
}

// Init initializes the instance.
func (p AVPlaybackCoordinationMedium) Init() AVPlaybackCoordinationMedium {
	rv := objc.Send[AVPlaybackCoordinationMedium](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlaybackCoordinationMedium) Autorelease() AVPlaybackCoordinationMedium {
	rv := objc.Send[AVPlaybackCoordinationMedium](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlaybackCoordinationMedium creates a new AVPlaybackCoordinationMedium instance.
func NewAVPlaybackCoordinationMedium() AVPlaybackCoordinationMedium {
	class := getAVPlaybackCoordinationMediumClass()
	rv := objc.Send[AVPlaybackCoordinationMedium](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// All playback coordinators that are connected to the coordination medium.
//
// # Discussion
//
// Returns an array of all the AVPlayerPlaybackCoordinators that are connected
// to the coordination medium. This coordination is specifically for
// AVPlayerPlaybackCoordinators, and we exclude
// AVDelegatingPlaybackCoordinators. AVPlaybackCoordinator properties and
// methods are individually configurable for each playback coordinator. To
// ensure correct synchronized behavior across all local playback
// coordinators, any common AVPlaybackCoordinator properties and methods
// should be set and called on all playback coordinators in the coordination
// medium. The properties and methods `otherParticipants`, “, and “ refer
// specifically to remote participants that are coordinated through a group
// session rather than through the playback coordination medium.
// `otherParticipants` only returns participants connected to the same group
// session. `setParticipantLimit` and
// `participantLimitForWaitingOutSuspensionsWithReason` affect only policies
// and behavior with the group session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinationMedium/connectedPlaybackCoordinators
func (p AVPlaybackCoordinationMedium) ConnectedPlaybackCoordinators() []AVPlayerPlaybackCoordinator {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("connectedPlaybackCoordinators"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerPlaybackCoordinator {
		return AVPlayerPlaybackCoordinatorFromID(id)
	})
}
