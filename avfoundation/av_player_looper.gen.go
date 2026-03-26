// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerLooper] class.
var (
	_AVPlayerLooperClass     AVPlayerLooperClass
	_AVPlayerLooperClassOnce sync.Once
)

func getAVPlayerLooperClass() AVPlayerLooperClass {
	_AVPlayerLooperClassOnce.Do(func() {
		_AVPlayerLooperClass = AVPlayerLooperClass{class: objc.GetClass("AVPlayerLooper")}
	})
	return _AVPlayerLooperClass
}

// GetAVPlayerLooperClass returns the class object for AVPlayerLooper.
func GetAVPlayerLooperClass() AVPlayerLooperClass {
	return getAVPlayerLooperClass()
}

type AVPlayerLooperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerLooperClass) Alloc() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that loops media content using a queue player.
//
// # Overview
// 
// You can manually implement looping playback in your app using
// [AVQueuePlayer], but [AVPlayerLooper] provides a much simpler interface to
// loop a single [AVPlayerItem]. You create a player looper by passing it a
// reference to your [AVQueuePlayer] and a template [AVPlayerItem] and the
// looper automatically manages the looping playback of this content (see
// example).
//
// # Creating a player looper
//
//   - [AVPlayerLooper.InitWithPlayerTemplateItemTimeRangeExistingItemsOrdering]: Creates a player looper that continuously plays the full duration of a player item while adhering to the specified ordering of existing items in the queue.
//   - [AVPlayerLooper.InitWithPlayerTemplateItemTimeRange]: Creates a player looper that continuously plays the specified time range of a player item.
//
// # Configuring looping
//
//   - [AVPlayerLooper.LoopingPlayerItems]: An array containing replicas of the template player item used to accomplish the looping.
//   - [AVPlayerLooper.DisableLooping]: Disables looping for the player queue.
//
// # Observing looping state
//
//   - [AVPlayerLooper.LoopCount]: The number of times the object played the media.
//   - [AVPlayerLooper.Status]: A status that indicates the object’s ability to loop playback.
//
// # Monitoring errors
//
//   - [AVPlayerLooper.Error]: An error that describes the reason looping failed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper
type AVPlayerLooper struct {
	objectivec.Object
}

// AVPlayerLooperFromID constructs a [AVPlayerLooper] from an objc.ID.
//
// An object that loops media content using a queue player.
func AVPlayerLooperFromID(id objc.ID) AVPlayerLooper {
	return AVPlayerLooper{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerLooper adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerLooper] class.
//
// # Creating a player looper
//
//   - [IAVPlayerLooper.InitWithPlayerTemplateItemTimeRangeExistingItemsOrdering]: Creates a player looper that continuously plays the full duration of a player item while adhering to the specified ordering of existing items in the queue.
//   - [IAVPlayerLooper.InitWithPlayerTemplateItemTimeRange]: Creates a player looper that continuously plays the specified time range of a player item.
//
// # Configuring looping
//
//   - [IAVPlayerLooper.LoopingPlayerItems]: An array containing replicas of the template player item used to accomplish the looping.
//   - [IAVPlayerLooper.DisableLooping]: Disables looping for the player queue.
//
// # Observing looping state
//
//   - [IAVPlayerLooper.LoopCount]: The number of times the object played the media.
//   - [IAVPlayerLooper.Status]: A status that indicates the object’s ability to loop playback.
//
// # Monitoring errors
//
//   - [IAVPlayerLooper.Error]: An error that describes the reason looping failed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper
type IAVPlayerLooper interface {
	objectivec.IObject

	// Topic: Creating a player looper

	// Creates a player looper that continuously plays the full duration of a player item while adhering to the specified ordering of existing items in the queue.
	InitWithPlayerTemplateItemTimeRangeExistingItemsOrdering(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject, itemOrdering AVPlayerLooperItemOrdering) AVPlayerLooper
	// Creates a player looper that continuously plays the specified time range of a player item.
	InitWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) AVPlayerLooper

	// Topic: Configuring looping

	// An array containing replicas of the template player item used to accomplish the looping.
	LoopingPlayerItems() []AVPlayerItem
	// Disables looping for the player queue.
	DisableLooping()

	// Topic: Observing looping state

	// The number of times the object played the media.
	LoopCount() int
	// A status that indicates the object’s ability to loop playback.
	Status() AVPlayerLooperStatus

	// Topic: Monitoring errors

	// An error that describes the reason looping failed.
	Error() foundation.INSError
}

// Init initializes the instance.
func (p AVPlayerLooper) Init() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerLooper) Autorelease() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerLooper creates a new AVPlayerLooper instance.
func NewAVPlayerLooper() AVPlayerLooper {
	class := getAVPlayerLooperClass()
	rv := objc.Send[AVPlayerLooper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a player looper that continuously plays the full duration of a
// player item.
//
// player: The queue player to use for playback. The player must not be `nil`.
//
// itemToLoop: The player item to loop, which must not be `nil`.
//
// # Return Value
// 
// An new instance of [AVPlayerLooper].
//
// # Discussion
// 
// Creating an instance of this class using this method is equivalent to
// calling [InitWithPlayerTemplateItemTimeRange] and passing a value of
// [invalid] for the `timeRange` parameter.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:)
func NewPlayerLooperWithPlayerTemplateItem(player IAVQueuePlayer, itemToLoop IAVPlayerItem) AVPlayerLooper {
	rv := objc.Send[objc.ID](objc.ID(getAVPlayerLooperClass().class), objc.Sel("playerLooperWithPlayer:templateItem:"), player, itemToLoop)
	return AVPlayerLooperFromID(rv)
}

// Creates a player looper that continuously plays the specified time range of
// a player item.
//
// player: The queue player to use for playback. The player must not be `nil`.
//
// itemToLoop: The player item to loop, which must not be `nil`.
//
// loopRange: The player item time range to loop. Passing a value of [invalid] is
// equivalent to a time range of [0, player item’s duration].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// # Return Value
// 
// An new instance of [AVPlayerLooper].
//
// # Discussion
// 
// The player item you specify will be used as a to generate at least 3 player
// item replicas that will be inserted into the specified player’s queue to
// accomplish the looping playback. As the player item will only be used as a
// template, and not for actual playback, any changes you make to the player
// item after initialization will not be reflected in the replicas. If you
// need to explicitly configure the replica player items, such as adding
// instances of [AVPlayerItemOutput] to them, you can access them through the
// [LoopingPlayerItems] property.
// 
// You should not modify the player’s queue while [AVPlayerLooper] is
// performing looping playback. If you need to perform any additional
// configuration of the player prior to playback, you should set its [Rate] to
// `0.0`, perform the required configuration, and then begin playback once the
// configuration is complete.
// 
// The [CMTimeRange] you specify will limit each item’s loop iteration to
// playing within the specified time range. To loop the full duration of the
// asset, specify a time range value of [invalid] for the `timeRange`
// parameter. Time range looping will be accomplished by seeking to the time
// range’s start time and setting player item’s [ForwardPlaybackEndTime]
// property on the looping item replicas.
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:)
// loopRange is a [coremedia.CMTimeRange].
func NewPlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) AVPlayerLooper {
	instance := getAVPlayerLooperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	return AVPlayerLooperFromID(rv)
}

// Creates a player looper that continuously plays the full duration of a
// player item while adhering to the specified ordering of existing items in
// the queue.
//
// player: A queue player to control playback.
//
// itemToLoop: A player item to loop.
//
// loopRange: The player item time range to loop. Passing a value of [invalid] is
// equivalent to a time range of [0, player item’s duration].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// itemOrdering: A value that indicates whether the looper inserts replica items before or
// after existing items in the specified queue player.
//
// # Discussion
// 
// The player looper doesn’t use the player item you specify for playback,
// and instead uses it as a template to create at least three player item
// replicas that it uses for looping playback. Because the looper only uses
// the player item as a template, any changes that you make to it after
// initialization aren’t reflected in the looping playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:existingItemsOrdering:)
// loopRange is a [coremedia.CMTimeRange].
func NewPlayerLooperWithPlayerTemplateItemTimeRangeExistingItemsOrdering(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject, itemOrdering AVPlayerLooperItemOrdering) AVPlayerLooper {
	instance := getAVPlayerLooperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:existingItemsOrdering:"), player, itemToLoop, loopRange, itemOrdering)
	return AVPlayerLooperFromID(rv)
}

// Creates a player looper that continuously plays the full duration of a
// player item while adhering to the specified ordering of existing items in
// the queue.
//
// player: A queue player to control playback.
//
// itemToLoop: A player item to loop.
//
// loopRange: The player item time range to loop. Passing a value of [invalid] is
// equivalent to a time range of [0, player item’s duration].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// itemOrdering: A value that indicates whether the looper inserts replica items before or
// after existing items in the specified queue player.
//
// loopRange is a [coremedia.CMTimeRange].
//
// # Discussion
// 
// The player looper doesn’t use the player item you specify for playback,
// and instead uses it as a template to create at least three player item
// replicas that it uses for looping playback. Because the looper only uses
// the player item as a template, any changes that you make to it after
// initialization aren’t reflected in the looping playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:existingItemsOrdering:)
func (p AVPlayerLooper) InitWithPlayerTemplateItemTimeRangeExistingItemsOrdering(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject, itemOrdering AVPlayerLooperItemOrdering) AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](p.ID, objc.Sel("initWithPlayer:templateItem:timeRange:existingItemsOrdering:"), player, itemToLoop, loopRange, itemOrdering)
	return rv
}
// Creates a player looper that continuously plays the specified time range of
// a player item.
//
// player: The queue player to use for playback. The player must not be `nil`.
//
// itemToLoop: The player item to loop, which must not be `nil`.
//
// loopRange: The player item time range to loop. Passing a value of [invalid] is
// equivalent to a time range of [0, player item’s duration].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// loopRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// An new instance of [AVPlayerLooper].
//
// # Discussion
// 
// The player item you specify will be used as a to generate at least 3 player
// item replicas that will be inserted into the specified player’s queue to
// accomplish the looping playback. As the player item will only be used as a
// template, and not for actual playback, any changes you make to the player
// item after initialization will not be reflected in the replicas. If you
// need to explicitly configure the replica player items, such as adding
// instances of [AVPlayerItemOutput] to them, you can access them through the
// [LoopingPlayerItems] property.
// 
// You should not modify the player’s queue while [AVPlayerLooper] is
// performing looping playback. If you need to perform any additional
// configuration of the player prior to playback, you should set its [Rate] to
// `0.0`, perform the required configuration, and then begin playback once the
// configuration is complete.
// 
// The [CMTimeRange] you specify will limit each item’s loop iteration to
// playing within the specified time range. To loop the full duration of the
// asset, specify a time range value of [invalid] for the `timeRange`
// parameter. Time range looping will be accomplished by seeking to the time
// range’s start time and setting player item’s [ForwardPlaybackEndTime]
// property on the looping item replicas.
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:)
func (p AVPlayerLooper) InitWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](p.ID, objc.Sel("initWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	return rv
}
// Disables looping for the player queue.
//
// # Discussion
// 
// The player looper will stop performing player queue operations for looping
// and let the current looping item replica play to the end. The player’s
// original [ActionAtItemEnd] property will be restored afterwards.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/disableLooping()
func (p AVPlayerLooper) DisableLooping() {
	objc.Send[objc.ID](p.ID, objc.Sel("disableLooping"))
}

// Returns player looper that continuously plays the specified time range of a
// player item.
//
// player: The queue player to use for playback. The player must not be `nil`.
//
// itemToLoop: The player item to loop, which must not be `nil`.
//
// loopRange: The player item time range to loop. Passing a value of [invalid] is
// equivalent to a time range of [0, player item’s duration].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// loopRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// An new instance of [AVPlayerLooper].
//
// # Discussion
// 
// The player item you specify will be used as a to generate at least 3 player
// item replicas that will be inserted into the specified player’s queue to
// accomplish the looping playback. As the player item will only be used as a
// template, and not for actual playback, any changes you make to the player
// item after initialization will not be reflected in the replicas. If you
// need to explicitly configure the replica player items, such as adding
// instances of [AVPlayerItemOutput] to them, you can access them through the
// [LoopingPlayerItems] property.
// 
// You should not modify the player’s queue while [AVPlayerLooper] is
// performing looping playback. If you need to perform any additional
// configuration of the player prior to playback, you should set its [Rate] to
// `0.0`, perform the required configuration, and then begin playback once the
// configuration is complete.
// 
// The [CMTimeRange] you specify will limit each item’s loop iteration to
// playing within the specified time range. To loop the full duration of the
// asset, specify a time range value of [invalid] for the `timeRange`
// parameter. Time range looping will be accomplished by seeking to the time
// range’s start time and setting player item’s [ForwardPlaybackEndTime]
// property on the looping item replicas.
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/playerLooperWithPlayer:templateItem:timeRange:
func (_AVPlayerLooperClass AVPlayerLooperClass) PlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) AVPlayerLooper {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerLooperClass.class), objc.Sel("playerLooperWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	return AVPlayerLooperFromID(rv)
}

// An array containing replicas of the template player item used to accomplish
// the looping.
//
// # Discussion
// 
// [AVPlayerLooper] creates replicas of the template [AVPlayerItem] using the
// [copyWithZone:] method and inserts them in the queue player’s queue to
// accomplish the looping. You can determine the number of replicas created
// and can listen for notifications and property changes from the replicas if
// desired.
// 
// Access to the [AVPlayerItem] replicas are for informational purposes and to
// allow you to apply any configuration that is not transferred from the
// template player item to the replicas. For instance, any instances of
// [AVPlayerItemOutput] and [AVPlayerItemMediaDataCollector] attached to the
// template player item are not transferred to the replicas so you should add
// them to each replica item if needed.
//
// [copyWithZone:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyWithZone:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopingPlayerItems
func (p AVPlayerLooper) LoopingPlayerItems() []AVPlayerItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("loopingPlayerItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItem {
		return AVPlayerItemFromID(id)
	})
}
// The number of times the object played the media.
//
// # Discussion
// 
// This value starts at 0 and increments as the player continues to loop the
// replica player items.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopCount
func (p AVPlayerLooper) LoopCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("loopCount"))
	return rv
}
// A status that indicates the object’s ability to loop playback.
//
// # Discussion
// 
// When the value of this property is [PlayerLooperStatusFailed] or
// [PlayerLooperStatusCancelled], you can no longer use the looper for
// playback. You need to create a new instance to begin looping again.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/status-swift.property
func (p AVPlayerLooper) Status() AVPlayerLooperStatus {
	rv := objc.Send[AVPlayerLooperStatus](p.ID, objc.Sel("status"))
	return AVPlayerLooperStatus(rv)
}
// An error that describes the reason looping failed.
//
// # Discussion
// 
// The value of this property is `nil` unless the looper’s [Status] changes
// to [PlayerLooperStatusFailed]. If this occurs, this property value contains
// an error object that provides the details of the error that prevented
// looping.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/error
func (p AVPlayerLooper) Error() foundation.INSError {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

