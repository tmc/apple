// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVQueuePlayer] class.
var (
	_AVQueuePlayerClass     AVQueuePlayerClass
	_AVQueuePlayerClassOnce sync.Once
)

func getAVQueuePlayerClass() AVQueuePlayerClass {
	_AVQueuePlayerClassOnce.Do(func() {
		_AVQueuePlayerClass = AVQueuePlayerClass{class: objc.GetClass("AVQueuePlayer")}
	})
	return _AVQueuePlayerClass
}

// GetAVQueuePlayerClass returns the class object for AVQueuePlayer.
func GetAVQueuePlayerClass() AVQueuePlayerClass {
	return getAVQueuePlayerClass()
}

type AVQueuePlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVQueuePlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVQueuePlayerClass) Alloc() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that plays a sequence of player items.
//
// # Overview
// 
// Use an instance of this class to manage a queue of player items.
//
// # Creating a queue player
//
//   - [AVQueuePlayer.InitWithItems]: Creates an object that plays a queue of items.
//
// # Managing the player queue
//
//   - [AVQueuePlayer.Items]: Returns an array of the currently enqueued items.
//   - [AVQueuePlayer.AdvanceToNextItem]: Ends playback of the current item and starts playback of the next item in the player’s queue.
//   - [AVQueuePlayer.CanInsertItemAfterItem]: Returns a Boolean value that indicates whether you can insert a player item into the player’s queue.
//   - [AVQueuePlayer.InsertItemAfterItem]: Inserts a player item after another player item in the queue.
//   - [AVQueuePlayer.RemoveItem]: Removes a given player item from the queue.
//   - [AVQueuePlayer.RemoveAllItems]: Removes all player items from the queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer
type AVQueuePlayer struct {
	AVPlayer
}

// AVQueuePlayerFromID constructs a [AVQueuePlayer] from an objc.ID.
//
// An object that plays a sequence of player items.
func AVQueuePlayerFromID(id objc.ID) AVQueuePlayer {
	return AVQueuePlayer{AVPlayer: AVPlayerFromID(id)}
}
// NOTE: AVQueuePlayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVQueuePlayer] class.
//
// # Creating a queue player
//
//   - [IAVQueuePlayer.InitWithItems]: Creates an object that plays a queue of items.
//
// # Managing the player queue
//
//   - [IAVQueuePlayer.Items]: Returns an array of the currently enqueued items.
//   - [IAVQueuePlayer.AdvanceToNextItem]: Ends playback of the current item and starts playback of the next item in the player’s queue.
//   - [IAVQueuePlayer.CanInsertItemAfterItem]: Returns a Boolean value that indicates whether you can insert a player item into the player’s queue.
//   - [IAVQueuePlayer.InsertItemAfterItem]: Inserts a player item after another player item in the queue.
//   - [IAVQueuePlayer.RemoveItem]: Removes a given player item from the queue.
//   - [IAVQueuePlayer.RemoveAllItems]: Removes all player items from the queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer
type IAVQueuePlayer interface {
	IAVPlayer

	// Topic: Creating a queue player

	// Creates an object that plays a queue of items.
	InitWithItems(items []AVPlayerItem) AVQueuePlayer

	// Topic: Managing the player queue

	// Returns an array of the currently enqueued items.
	Items() []AVPlayerItem
	// Ends playback of the current item and starts playback of the next item in the player’s queue.
	AdvanceToNextItem()
	// Returns a Boolean value that indicates whether you can insert a player item into the player’s queue.
	CanInsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) bool
	// Inserts a player item after another player item in the queue.
	InsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem)
	// Removes a given player item from the queue.
	RemoveItem(item IAVPlayerItem)
	// Removes all player items from the queue.
	RemoveAllItems()
}

// Init initializes the instance.
func (q AVQueuePlayer) Init() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q AVQueuePlayer) Autorelease() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVQueuePlayer creates a new AVQueuePlayer instance.
func NewAVQueuePlayer() AVQueuePlayer {
	class := getAVQueuePlayerClass()
	rv := objc.Send[AVQueuePlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that plays a queue of items.
//
// items: The array of player items with which to populate the queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/init(items:)
func NewQueuePlayerWithItems(items []AVPlayerItem) AVQueuePlayer {
	instance := getAVQueuePlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:"), objectivec.IObjectSliceToNSArray(items))
	return AVQueuePlayerFromID(rv)
}

// Creates a new player to play the specified player item.
//
// item: The player item to play.
//
// # Return Value
// 
// A new player initialized to play `item`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(playerItem:)
func NewQueuePlayerWithPlayerItem(item IAVPlayerItem) AVQueuePlayer {
	instance := getAVQueuePlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlayerItem:"), item)
	return AVQueuePlayerFromID(rv)
}

// Creates a new player to play a single audiovisual resource referenced by a
// given URL.
//
// URL: A URL that identifies an audiovisual resource.
//
// # Return Value
// 
// A new player instance initialized to play the audiovisual resource
// specified by [URL].
//
// # Discussion
// 
// This method implicitly creates an [AVPlayerItem] object. You can get the
// player item using [CurrentItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(url:)
func NewQueuePlayerWithURL(URL foundation.INSURL) AVQueuePlayer {
	instance := getAVQueuePlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return AVQueuePlayerFromID(rv)
}

// Creates an object that plays a queue of items.
//
// items: The array of player items with which to populate the queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/init(items:)
func (q AVQueuePlayer) InitWithItems(items []AVPlayerItem) AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](q.ID, objc.Sel("initWithItems:"), objectivec.IObjectSliceToNSArray(items))
	return rv
}
// Returns an array of the currently enqueued items.
//
// # Return Value
// 
// An array of the currently enqueued player items.
//
// # Discussion
// 
// The array contains [AVPlayerItem] objects currently in the player’s
// queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/items()
func (q AVQueuePlayer) Items() []AVPlayerItem {
	rv := objc.Send[[]objc.ID](q.ID, objc.Sel("items"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItem {
		return AVPlayerItemFromID(id)
	})
}
// Ends playback of the current item and starts playback of the next item in
// the player’s queue.
//
// # Discussion
// 
// Calling this method also removes the current item from the play queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/advanceToNextItem()
func (q AVQueuePlayer) AdvanceToNextItem() {
	objc.Send[objc.ID](q.ID, objc.Sel("advanceToNextItem"))
}
// Returns a Boolean value that indicates whether you can insert a player item
// into the player’s queue.
//
// item: The player item to insert.
//
// afterItem: The player item in the queue to follow. Pass `nil` to test if you can
// append the item to the queue.
//
// # Return Value
// 
// [true] if `item` can be appended to the queue, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Adding the same item to a player at more than one position in the queue
// isn’t supported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/canInsert(_:after:)
func (q AVQueuePlayer) CanInsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) bool {
	rv := objc.Send[bool](q.ID, objc.Sel("canInsertItem:afterItem:"), item, afterItem)
	return rv
}
// Inserts a player item after another player item in the queue.
//
// item: The item to insert into the queue.
//
// afterItem: The player item in the queue to follow. Pass `nil` to append the item to
// the queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/insert(_:after:)
func (q AVQueuePlayer) InsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) {
	objc.Send[objc.ID](q.ID, objc.Sel("insertItem:afterItem:"), item, afterItem)
}
// Removes a given player item from the queue.
//
// item: The player item to remove from the queue.
//
// # Discussion
// 
// If the item is currently playing, calling this method has the same effect
// as calling the [AdvanceToNextItem] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/remove(_:)
func (q AVQueuePlayer) RemoveItem(item IAVPlayerItem) {
	objc.Send[objc.ID](q.ID, objc.Sel("removeItem:"), item)
}
// Removes all player items from the queue.
//
// # Discussion
// 
// Calling this method removes all currently enqueued player items and stops
// playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/removeAllItems()
func (q AVQueuePlayer) RemoveAllItems() {
	objc.Send[objc.ID](q.ID, objc.Sel("removeAllItems"))
}

// Returns an object that plays a queue of items.
//
// items: An array of [AVPlayerItem] objects with which to initially populate the
// player’s queue.
//
// # Return Value
// 
// A new queue player.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/queuePlayerWithItems:
func (_AVQueuePlayerClass AVQueuePlayerClass) QueuePlayerWithItems(items []AVPlayerItem) AVQueuePlayer {
	rv := objc.Send[objc.ID](objc.ID(_AVQueuePlayerClass.class), objc.Sel("queuePlayerWithItems:"), objectivec.IObjectSliceToNSArray(items))
	return AVQueuePlayerFromID(rv)
}

