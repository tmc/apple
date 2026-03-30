// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerInterstitialEvent] class.
var (
	_AVPlayerInterstitialEventClass     AVPlayerInterstitialEventClass
	_AVPlayerInterstitialEventClassOnce sync.Once
)

func getAVPlayerInterstitialEventClass() AVPlayerInterstitialEventClass {
	_AVPlayerInterstitialEventClassOnce.Do(func() {
		_AVPlayerInterstitialEventClass = AVPlayerInterstitialEventClass{class: objc.GetClass("AVPlayerInterstitialEvent")}
	})
	return _AVPlayerInterstitialEventClass
}

// GetAVPlayerInterstitialEventClass returns the class object for AVPlayerInterstitialEvent.
func GetAVPlayerInterstitialEventClass() AVPlayerInterstitialEventClass {
	return getAVPlayerInterstitialEventClass()
}

type AVPlayerInterstitialEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerInterstitialEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerInterstitialEventClass) Alloc() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides instructions for how a player presents interstitial
// content.
//
// # Overview
//
// An interstitial event defines a [AVPlayerInterstitialEvent.Date] or [AVPlayerInterstitialEvent.Time], on the timeline of its
// [AVPlayerInterstitialEvent.PrimaryItem], at which playback of interstitial content begins. It
// specifies the alternative interstitial content to play as an array of one
// or more template player items. The system uses the configuration of the
// event’s [AVPlayerInterstitialEvent.TemplateItems] to build new player item instances to present the
// interstitial content.
//
// Use [AVPlayerInterstitialEventMonitor] to observe the scheduling and
// progress of interstitial events. If your app requires specifying the
// schedule of interstitial events, use [AVPlayerInterstitialEventController]
// instead.
//
// # Identifying events
//
//   - [AVPlayerInterstitialEvent.Identifier]: An identifier for the event.
//
// # Accessing player items
//
//   - [AVPlayerInterstitialEvent.PrimaryItem]: The player item that represents the primary content.
//   - [AVPlayerInterstitialEvent.TemplateItems]: An array of player item configurations to use as templates for player items that play interstitial content.
//
// # Configuring cues
//
//   - [AVPlayerInterstitialEvent.Cue]: A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// # Inspecting timing
//
//   - [AVPlayerInterstitialEvent.Time]: A time within the timeline of the primary content that playback of interstitial content begins.
//   - [AVPlayerInterstitialEvent.Date]: A date within the date range of the primary content that playback of interstitial content begins.
//   - [AVPlayerInterstitialEvent.WillPlayOnce]: A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//   - [AVPlayerInterstitialEvent.ResumptionOffset]: A time offset at which playback of primary content resumes after interstitial content finishes.
//   - [AVPlayerInterstitialEvent.PlayoutLimit]: The time offset at which playback of the interstitial ends.
//   - [AVPlayerInterstitialEvent.AlignsStartWithPrimarySegmentBoundary]: A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//   - [AVPlayerInterstitialEvent.AlignsResumptionWithPrimarySegmentBoundary]: A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// # Managing restrictions
//
//   - [AVPlayerInterstitialEvent.Restrictions]: The restrictions the event imposes on the playback of interstitial content.
//
// # Accessing asset lists
//
//   - [AVPlayerInterstitialEvent.AssetListResponse]: The asset list JSON response as a dictionary.
//
// # Accessing attributes
//
//   - [AVPlayerInterstitialEvent.UserDefinedAttributes]: Attributes of the event that the vendor or app defines.
//
// # Inspecting timeline occupancy
//
//   - [AVPlayerInterstitialEvent.TimelineOccupancy]: An event’s occupancy on the integrated timeline.
//   - [AVPlayerInterstitialEvent.SupplementsPrimaryContent]: A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//   - [AVPlayerInterstitialEvent.ContentMayVary]: A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//   - [AVPlayerInterstitialEvent.PlannedDuration]: The planned duration of the event.
//   - [AVPlayerInterstitialEvent.SetPlannedDuration]
//
// # Managing skipping behavior
//
//   - [AVPlayerInterstitialEvent.SkipControlLocalizedLabelBundleKey]: The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//   - [AVPlayerInterstitialEvent.SkipControlTimeRange]: The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// # Instance Properties
//
//   - [AVPlayerInterstitialEvent.ScheduleIdentifier]: The identifier of the daterange-schedule that produced this event. nil if the event was not a product of a daterange-schedule.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent
type AVPlayerInterstitialEvent struct {
	objectivec.Object
}

// AVPlayerInterstitialEventFromID constructs a [AVPlayerInterstitialEvent] from an objc.ID.
//
// An object that provides instructions for how a player presents interstitial
// content.
func AVPlayerInterstitialEventFromID(id objc.ID) AVPlayerInterstitialEvent {
	return AVPlayerInterstitialEvent{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerInterstitialEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerInterstitialEvent] class.
//
// # Identifying events
//
//   - [IAVPlayerInterstitialEvent.Identifier]: An identifier for the event.
//
// # Accessing player items
//
//   - [IAVPlayerInterstitialEvent.PrimaryItem]: The player item that represents the primary content.
//   - [IAVPlayerInterstitialEvent.TemplateItems]: An array of player item configurations to use as templates for player items that play interstitial content.
//
// # Configuring cues
//
//   - [IAVPlayerInterstitialEvent.Cue]: A cue to schedule interstitial event playback at a predefined position during primary playback.
//
// # Inspecting timing
//
//   - [IAVPlayerInterstitialEvent.Time]: A time within the timeline of the primary content that playback of interstitial content begins.
//   - [IAVPlayerInterstitialEvent.Date]: A date within the date range of the primary content that playback of interstitial content begins.
//   - [IAVPlayerInterstitialEvent.WillPlayOnce]: A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
//   - [IAVPlayerInterstitialEvent.ResumptionOffset]: A time offset at which playback of primary content resumes after interstitial content finishes.
//   - [IAVPlayerInterstitialEvent.PlayoutLimit]: The time offset at which playback of the interstitial ends.
//   - [IAVPlayerInterstitialEvent.AlignsStartWithPrimarySegmentBoundary]: A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
//   - [IAVPlayerInterstitialEvent.AlignsResumptionWithPrimarySegmentBoundary]: A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
//
// # Managing restrictions
//
//   - [IAVPlayerInterstitialEvent.Restrictions]: The restrictions the event imposes on the playback of interstitial content.
//
// # Accessing asset lists
//
//   - [IAVPlayerInterstitialEvent.AssetListResponse]: The asset list JSON response as a dictionary.
//
// # Accessing attributes
//
//   - [IAVPlayerInterstitialEvent.UserDefinedAttributes]: Attributes of the event that the vendor or app defines.
//
// # Inspecting timeline occupancy
//
//   - [IAVPlayerInterstitialEvent.TimelineOccupancy]: An event’s occupancy on the integrated timeline.
//   - [IAVPlayerInterstitialEvent.SupplementsPrimaryContent]: A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
//   - [IAVPlayerInterstitialEvent.ContentMayVary]: A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
//   - [IAVPlayerInterstitialEvent.PlannedDuration]: The planned duration of the event.
//   - [IAVPlayerInterstitialEvent.SetPlannedDuration]
//
// # Managing skipping behavior
//
//   - [IAVPlayerInterstitialEvent.SkipControlLocalizedLabelBundleKey]: The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
//   - [IAVPlayerInterstitialEvent.SkipControlTimeRange]: The time range within the duration of the interstitial event for which a skip button should be displayed.
//
// # Instance Properties
//
//   - [IAVPlayerInterstitialEvent.ScheduleIdentifier]: The identifier of the daterange-schedule that produced this event. nil if the event was not a product of a daterange-schedule.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent
type IAVPlayerInterstitialEvent interface {
	objectivec.IObject

	// Topic: Identifying events

	// An identifier for the event.
	Identifier() string

	// Topic: Accessing player items

	// The player item that represents the primary content.
	PrimaryItem() IAVPlayerItem
	// An array of player item configurations to use as templates for player items that play interstitial content.
	TemplateItems() []AVPlayerItem

	// Topic: Configuring cues

	// A cue to schedule interstitial event playback at a predefined position during primary playback.
	Cue() AVPlayerInterstitialEventCue

	// Topic: Inspecting timing

	// A time within the timeline of the primary content that playback of interstitial content begins.
	Time() coremedia.CMTime
	// A date within the date range of the primary content that playback of interstitial content begins.
	Date() foundation.INSDate
	// A Boolean value that indicates whether to schedule this event one time only and suppress subsequent replay.
	WillPlayOnce() bool
	// A time offset at which playback of primary content resumes after interstitial content finishes.
	ResumptionOffset() coremedia.CMTime
	// The time offset at which playback of the interstitial ends.
	PlayoutLimit() coremedia.CMTime
	// A Boolean value that indicates whether the start time of interstitial playback should snap to a segment boundary of the primary asset.
	AlignsStartWithPrimarySegmentBoundary() bool
	// A Boolean value that indicates whether the resumption time of primary playback should snap to a segment boundary of the primary asset.
	AlignsResumptionWithPrimarySegmentBoundary() bool

	// Topic: Managing restrictions

	// The restrictions the event imposes on the playback of interstitial content.
	Restrictions() AVPlayerInterstitialEventRestrictions

	// Topic: Accessing asset lists

	// The asset list JSON response as a dictionary.
	AssetListResponse() foundation.INSDictionary

	// Topic: Accessing attributes

	// Attributes of the event that the vendor or app defines.
	UserDefinedAttributes() foundation.INSDictionary

	// Topic: Inspecting timeline occupancy

	// An event’s occupancy on the integrated timeline.
	TimelineOccupancy() AVPlayerInterstitialEventTimelineOccupancy
	// A Boolean value that indicates whether an event supplements the primary content and should present with the primary item.
	SupplementsPrimaryContent() bool
	// A Boolean value that indicates whether an event’s content is dynamic and the server may respond with different interstitial assets for other participants in a coordinated playback session.
	ContentMayVary() bool
	// The planned duration of the event.
	PlannedDuration() coremedia.CMTime
	SetPlannedDuration(value coremedia.CMTime)

	// Topic: Managing skipping behavior

	// The key defined in the AVPlayerInterstitialEventController’s localizedStringsBundle that points to the localized label for the skip button.
	SkipControlLocalizedLabelBundleKey() string
	// The time range within the duration of the interstitial event for which a skip button should be displayed.
	SkipControlTimeRange() coremedia.CMTimeRange

	// Topic: Instance Properties

	// The identifier of the daterange-schedule that produced this event. nil if the event was not a product of a daterange-schedule.
	ScheduleIdentifier() string
}

// Init initializes the instance.
func (p AVPlayerInterstitialEvent) Init() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerInterstitialEvent) Autorelease() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEvent creates a new AVPlayerInterstitialEvent instance.
func NewAVPlayerInterstitialEvent() AVPlayerInterstitialEvent {
	class := getAVPlayerInterstitialEventClass()
	rv := objc.Send[AVPlayerInterstitialEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an interstitial event for the specified date.
//
// primaryItem: A player item that provides the primary playback content. It defines the
// timeline during which an interstitial event occurs. The item must have an
// asset that provides an intrinsic mapping from its timeline to real-time
// dates.
//
// date: A date within the timeline of the primary item at which to temporarily
// suspend playback of primary content, and play interstitial content instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:date:)
func NewPlayerInterstitialEventWithPrimaryItemDate(primaryItem IAVPlayerItem, date foundation.INSDate) AVPlayerInterstitialEvent {
	rv := objc.Send[objc.ID](objc.ID(getAVPlayerInterstitialEventClass().class), objc.Sel("interstitialEventWithPrimaryItem:date:"), primaryItem, date)
	return AVPlayerInterstitialEventFromID(rv)
}

// Creates an interstitial event for the specified time.
//
// primaryItem: A player item that provides the primary playback content. It defines the
// timeline during which an interstitial event occurs. The item must have an
// asset that provides an intrinsic mapping from its timeline to real-time
// dates.
//
// time: A time within the timeline of the primary item at which to temporarily
// suspend playback of primary content, and play interstitial content instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/init(primaryItem:time:)
func NewPlayerInterstitialEventWithPrimaryItemTime(primaryItem IAVPlayerItem, time coremedia.CMTime) AVPlayerInterstitialEvent {
	rv := objc.Send[objc.ID](objc.ID(getAVPlayerInterstitialEventClass().class), objc.Sel("interstitialEventWithPrimaryItem:time:"), primaryItem, time)
	return AVPlayerInterstitialEventFromID(rv)
}

// An identifier for the event.
//
// # Discussion
//
// Setting an event on the interstitial event controller replaces any existing
// event with the same identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/identifier
func (p AVPlayerInterstitialEvent) Identifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The player item that represents the primary content.
//
// # Discussion
//
// The item must contain an [AVAsset] that provides intrinsic mappings from
// its timeline to realtime dates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/primaryItem
func (p AVPlayerInterstitialEvent) PrimaryItem() IAVPlayerItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("primaryItem"))
	return AVPlayerItemFromID(objc.ID(rv))
}

// An array of player item configurations to use as templates for player items
// that play interstitial content.
//
// # Discussion
//
// If you require the system to create new player items using the same asset
// instance as the template item, create the asset with an
// [AVURLAssetPrimarySessionIdentifierKey] value equal to
// [HttpSessionIdentifier] of the primary item’s [Asset]. Creating assets
// this way simplifies cases where you require loading their data with a
// custom [AVAssetResourceLoader] delegate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/templateItems
//
// [AVURLAssetPrimarySessionIdentifierKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPrimarySessionIdentifierKey
func (p AVPlayerInterstitialEvent) TemplateItems() []AVPlayerItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("templateItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItem {
		return AVPlayerItemFromID(id)
	})
}

// A cue to schedule interstitial event playback at a predefined position
// during primary playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/cue-swift.property
func (p AVPlayerInterstitialEvent) Cue() AVPlayerInterstitialEventCue {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cue"))
	return AVPlayerInterstitialEventCue(foundation.NSStringFromID(rv).String())
}

// A time within the timeline of the primary content that playback of
// interstitial content begins.
//
// # Discussion
//
// This property value is [invalid] if you initialized the event with a date
// instead of a time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/time
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (p AVPlayerInterstitialEvent) Time() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("time"))
	return coremedia.CMTime(rv)
}

// A date within the date range of the primary content that playback of
// interstitial content begins.
//
// # Discussion
//
// This property value is `nil` if you initialized the event with a time
// instead of a date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/date
func (p AVPlayerInterstitialEvent) Date() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// A Boolean value that indicates whether to schedule this event one time only
// and suppress subsequent replay.
//
// # Discussion
//
// The “once” provision takes effect at the start of interstitial
// playback. The system doesn’t schedule playback again even if the first
// playback is canceled before completion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/willPlayOnce
func (p AVPlayerInterstitialEvent) WillPlayOnce() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("willPlayOnce"))
	return rv
}

// A time offset at which playback of primary content resumes after
// interstitial content finishes.
//
// # Discussion
//
// This property supports definite time values. Specify [indefinite] to
// indicate that the effective resumption time offset should align with the
// clock time elapsed during interstitial playback; this value is typically
// suitable for live broadcasts.
//
// The default value is [zero].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/resumptionOffset
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (p AVPlayerInterstitialEvent) ResumptionOffset() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("resumptionOffset"))
	return coremedia.CMTime(rv)
}

// The time offset at which playback of the interstitial ends.
//
// # Discussion
//
// This value can be any positive numeric value, or [invalid] (the default)
// which indicates no limit.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/playoutLimit
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (p AVPlayerInterstitialEvent) PlayoutLimit() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("playoutLimit"))
	return coremedia.CMTime(rv)
}

// A Boolean value that indicates whether the start time of interstitial
// playback should snap to a segment boundary of the primary asset.
//
// # Discussion
//
// If the value is true, the system adjusts the start time or date of the
// interstitial to the nearest segment boundary when the primary player is
// playing an HTTP Live Streaming asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/alignsStartWithPrimarySegmentBoundary
func (p AVPlayerInterstitialEvent) AlignsStartWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("alignsStartWithPrimarySegmentBoundary"))
	return rv
}

// A Boolean value that indicates whether the resumption time of primary
// playback should snap to a segment boundary of the primary asset.
//
// # Discussion
//
// If the value is true, the system adjusts the resumption time of primary
// playback following an interstitial to the nearest segment boundary when the
// primary player is playing an HTTP Live Streaming asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/alignsResumptionWithPrimarySegmentBoundary
func (p AVPlayerInterstitialEvent) AlignsResumptionWithPrimarySegmentBoundary() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("alignsResumptionWithPrimarySegmentBoundary"))
	return rv
}

// The restrictions the event imposes on the playback of interstitial content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/restrictions-swift.property
func (p AVPlayerInterstitialEvent) Restrictions() AVPlayerInterstitialEventRestrictions {
	rv := objc.Send[AVPlayerInterstitialEventRestrictions](p.ID, objc.Sel("restrictions"))
	return AVPlayerInterstitialEventRestrictions(rv)
}

// The asset list JSON response as a dictionary.
//
// # Discussion
//
// The value of this property is `nil` if there is no asset list loaded for
// the event. If this value is `nil` and the event’s [TemplateItems] is
// empty, then an asset list read is expected. If this value is `nil` and
// [TemplateItems] isn’t empty, an asset list read isn’t expected.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/assetListResponse
func (p AVPlayerInterstitialEvent) AssetListResponse() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("assetListResponse"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Attributes of the event that the vendor or app defines.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/userDefinedAttributes
func (p AVPlayerInterstitialEvent) UserDefinedAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userDefinedAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// An event’s occupancy on the integrated timeline.
//
// # Discussion
//
// The default value is
// [AVPlayerInterstitialEventTimelineOccupancySinglePoint].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/timelineOccupancy-swift.property
func (p AVPlayerInterstitialEvent) TimelineOccupancy() AVPlayerInterstitialEventTimelineOccupancy {
	rv := objc.Send[AVPlayerInterstitialEventTimelineOccupancy](p.ID, objc.Sel("timelineOccupancy"))
	return AVPlayerInterstitialEventTimelineOccupancy(rv)
}

// A Boolean value that indicates whether an event supplements the primary
// content and should present with the primary item.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/supplementsPrimaryContent
func (p AVPlayerInterstitialEvent) SupplementsPrimaryContent() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("supplementsPrimaryContent"))
	return rv
}

// A Boolean value that indicates whether an event’s content is dynamic and
// the server may respond with different interstitial assets for other
// participants in a coordinated playback session.
//
// # Discussion
//
// If the value is false, the primary asset participates in coordinated
// playback, this event does as well.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/contentMayVary
func (p AVPlayerInterstitialEvent) ContentMayVary() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("contentMayVary"))
	return rv
}

// The planned duration of the event.
//
// # Discussion
//
// The default value is [zero].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/plannedDuration
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (p AVPlayerInterstitialEvent) PlannedDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("plannedDuration"))
	return coremedia.CMTime(rv)
}
func (p AVPlayerInterstitialEvent) SetPlannedDuration(value coremedia.CMTime) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlannedDuration:"), value)
}

// The key defined in the AVPlayerInterstitialEventController’s
// localizedStringsBundle that points to the localized label for the skip
// button.
//
// # Discussion
//
// If the value of the property is nil, the skip button may contain a generic
// label depending on the implementation of the UI that’s in use. To ensure
// the best available user experience in various playback configurations,
// including external playback, set a value for this property that provides
// localized translations of skip control labels.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/skipControlLocalizedLabelBundleKey
func (p AVPlayerInterstitialEvent) SkipControlLocalizedLabelBundleKey() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("skipControlLocalizedLabelBundleKey"))
	return foundation.NSStringFromID(rv).String()
}

// The time range within the duration of the interstitial event for which a
// skip button should be displayed.
//
// # Discussion
//
// The start of the time range should indicate at which point the skip button
// should appear. The duration of the time range should indicate how long the
// skip button should be available. If this value is set to
// kCMTimePositiveInfinity, then the skip button will be available for the
// remainder of the interstitial’s duration after appearing. If either the
// start or duration of the time range is kCMTimeInvalid, then the
// interstitial will NOT be eligible to be skipped.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/skipControlTimeRange
func (p AVPlayerInterstitialEvent) SkipControlTimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](p.ID, objc.Sel("skipControlTimeRange"))
	return coremedia.CMTimeRange(rv)
}

// The identifier of the daterange-schedule that produced this event. nil if
// the event was not a product of a daterange-schedule.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/scheduleIdentifier
func (p AVPlayerInterstitialEvent) ScheduleIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("scheduleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
