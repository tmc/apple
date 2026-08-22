// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CWChannel] class.
var (
	_CWChannelClass     CWChannelClass
	_CWChannelClassOnce sync.Once
)

func getCWChannelClass() CWChannelClass {
	_CWChannelClassOnce.Do(func() {
		_CWChannelClass = CWChannelClass{class: objc.GetClass("CWChannel")}
	})
	return _CWChannelClass
}

// GetCWChannelClass returns the class object for CWChannel.
func GetCWChannelClass() CWChannelClass {
	return getCWChannelClass()
}

type CWChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWChannelClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWChannelClass) Alloc() CWChannel {
	rv := objc.Send[CWChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates an IEEE 802.11 channel.
//
// # Comparing channels
//
//   - [CWChannel.IsEqualToChannel]: Determine CWChannel object equality.
//
// # Instance Properties
//
//   - [CWChannel.ChannelBand]: The channel band.
//   - [CWChannel.ChannelNumber]: The channel number.
//   - [CWChannel.ChannelWidth]: The channel width.
//
// # Initializers
//
//   - [CWChannel.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel
type CWChannel struct {
	objectivec.Object
}

// CWChannelFromID constructs a [CWChannel] from an objc.ID.
//
// Encapsulates an IEEE 802.11 channel.
func CWChannelFromID(id objc.ID) CWChannel {
	return CWChannel{objectivec.Object{ID: id}}
}

// NOTE: CWChannel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWChannel] class.
//
// # Comparing channels
//
//   - [ICWChannel.IsEqualToChannel]: Determine CWChannel object equality.
//
// # Instance Properties
//
//   - [ICWChannel.ChannelBand]: The channel band.
//   - [ICWChannel.ChannelNumber]: The channel number.
//   - [ICWChannel.ChannelWidth]: The channel width.
//
// # Initializers
//
//   - [ICWChannel.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel
type ICWChannel interface {
	objectivec.IObject

	// Topic: Comparing channels

	// Determine CWChannel object equality.
	IsEqualToChannel(channel ICWChannel) bool

	// Topic: Instance Properties

	// The channel band.
	ChannelBand() CWChannelBand
	// The channel number.
	ChannelNumber() int
	// The channel width.
	ChannelWidth() CWChannelWidth

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CWChannel

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CWChannel) Init() CWChannel {
	rv := objc.Send[CWChannel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWChannel) Autorelease() CWChannel {
	rv := objc.Send[CWChannel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWChannel creates a new CWChannel instance.
func NewCWChannel() CWChannel {
	class := getCWChannelClass()
	rv := objc.Send[CWChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/init(coder:)
func NewCWChannelWithCoder(coder foundation.INSCoder) CWChannel {
	instance := getCWChannelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWChannelFromID(rv)
}

// Determine CWChannel object equality.
//
// channel: The CWChannel object with which to compare the receiver.
//
// # Return Value
//
// YES if the objects are equal.
//
// # Discussion
//
// CWChannel objects are considered equal if all their corresponding
// properties are equal.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/isEqual(to:)
func (c CWChannel) IsEqualToChannel(channel ICWChannel) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEqualToChannel:"), channel)
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/init(coder:)
func (c CWChannel) InitWithCoder(coder foundation.INSCoder) CWChannel {
	rv := objc.Send[CWChannel](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CWChannel) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The channel band.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelBand
func (c CWChannel) ChannelBand() CWChannelBand {
	rv := objc.Send[CWChannelBand](c.ID, objc.Sel("channelBand"))
	return CWChannelBand(rv)
}

// The channel number.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelNumber
func (c CWChannel) ChannelNumber() int {
	rv := objc.Send[int](c.ID, objc.Sel("channelNumber"))
	return rv
}

// The channel width.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWChannel/channelWidth
func (c CWChannel) ChannelWidth() CWChannelWidth {
	rv := objc.Send[CWChannelWidth](c.ID, objc.Sel("channelWidth"))
	return CWChannelWidth(rv)
}
