// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetadataItem] class.
var (
	_AVMetadataItemClass     AVMetadataItemClass
	_AVMetadataItemClassOnce sync.Once
)

func getAVMetadataItemClass() AVMetadataItemClass {
	_AVMetadataItemClassOnce.Do(func() {
		_AVMetadataItemClass = AVMetadataItemClass{class: objc.GetClass("AVMetadataItem")}
	})
	return _AVMetadataItemClass
}

// GetAVMetadataItemClass returns the class object for AVMetadataItem.
func GetAVMetadataItemClass() AVMetadataItemClass {
	return getAVMetadataItemClass()
}

type AVMetadataItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetadataItemClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetadataItemClass) Alloc() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A metadata item for an audiovisual asset or one of its tracks.
//
// # Overview
//
// To effectively use [AVMetadataItem], you need to understand how
// [AVFoundation] organizes metadata. To simplify finding and filtering
// metadata items, the framework groups related metadata into key spaces:
//
// - The framework defines several format-specific key spaces. They roughly
// correlate to a particular container or file format, such as QuickTime
// (QuickTime metadata and user data) or MP3 (ID3). However, a single asset
// may contain metadata values across multiple key spaces. To retrieve an
// asset’s complete collection of format-specific metadata, you use its
// [AVMetadataItem.Metadata] property. - There are several common metadata values, such as a
// movie’s creation date or description, that can exist across multiple key
// spaces. To help normalize access to this common metadata, the framework
// provides a common key space that gives access to a limited set of metadata
// values common to several key spaces. This makes it easy to retrieve
// commonly used metadata without concern for the specific format. To retrieve
// an asset’s collection of common metadata, you use its [AVMetadataItem.CommonMetadata]
// property.
//
// Metadata items have keys that accord with the specification of the
// container format from which they’re drawn. Full details of the metadata
// formats, metadata keys, and metadata key spaces supported by AVFoundation
// are available in [AVMetadataKeySpace] and [AVMetadataKey].
//
// To load values of a metadata item when you access them for the first time,
// use the methods from the [AVAsynchronousKeyValueLoading] protocol. The
// [AVAsset] class and other classes in turn provide their metadata as needed
// so that you can obtain objects from those arrays without incurring overhead
// for items you don’t inspect.
//
// To filter arrays of metadata items, you use the methods of this class. For
// example, you can filter by key and key space, by locale, and by preferred
// language.
//
// # Identifying metadata items
//
//   - [AVMetadataItem.Identifier]: An identifier for a metadata item.
//
// # Loading values
//
//   - [AVMetadataItem.DataType]: The data type of the metadata item’s value.
//
// # Accessing keys and key spaces
//
//   - [AVMetadataItem.Key]: The key of the metadata item.
//   - [AVMetadataItem.CommonKey]: The common key of the metadata item.
//   - [AVMetadataItem.KeySpace]: The key space for the metadata item’s key.
//
// # Accessing timing
//
//   - [AVMetadataItem.Time]: The timestamp of the metadata item.
//   - [AVMetadataItem.StartDate]: The start date of the timed metadata.
//   - [AVMetadataItem.Duration]: The duration of the metadata item.
//
// # Accessing language support
//
//   - [AVMetadataItem.Locale]: The locale of the metadata item.
//   - [AVMetadataItem.ExtendedLanguageTag]: The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem
//
// [AVFoundation]: https://developer.apple.com/documentation/AVFoundation
type AVMetadataItem struct {
	objectivec.Object
}

// AVMetadataItemFromID constructs a [AVMetadataItem] from an objc.ID.
//
// A metadata item for an audiovisual asset or one of its tracks.
func AVMetadataItemFromID(id objc.ID) AVMetadataItem {
	return AVMetadataItem{objectivec.Object{ID: id}}
}

// NOTE: AVMetadataItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetadataItem] class.
//
// # Identifying metadata items
//
//   - [IAVMetadataItem.Identifier]: An identifier for a metadata item.
//
// # Loading values
//
//   - [IAVMetadataItem.DataType]: The data type of the metadata item’s value.
//
// # Accessing keys and key spaces
//
//   - [IAVMetadataItem.Key]: The key of the metadata item.
//   - [IAVMetadataItem.CommonKey]: The common key of the metadata item.
//   - [IAVMetadataItem.KeySpace]: The key space for the metadata item’s key.
//
// # Accessing timing
//
//   - [IAVMetadataItem.Time]: The timestamp of the metadata item.
//   - [IAVMetadataItem.StartDate]: The start date of the timed metadata.
//   - [IAVMetadataItem.Duration]: The duration of the metadata item.
//
// # Accessing language support
//
//   - [IAVMetadataItem.Locale]: The locale of the metadata item.
//   - [IAVMetadataItem.ExtendedLanguageTag]: The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem
type IAVMetadataItem interface {
	objectivec.IObject
	AVAsynchronousKeyValueLoading

	// Topic: Identifying metadata items

	// An identifier for a metadata item.
	Identifier() AVMetadataIdentifier

	// Topic: Loading values

	// The data type of the metadata item’s value.
	DataType() string

	// Topic: Accessing keys and key spaces

	// The key of the metadata item.
	Key() objectivec.IObject
	// The common key of the metadata item.
	CommonKey() AVMetadataKey
	// The key space for the metadata item’s key.
	KeySpace() AVMetadataKeySpace

	// Topic: Accessing timing

	// The timestamp of the metadata item.
	Time() coremedia.CMTime
	// The start date of the timed metadata.
	StartDate() foundation.INSDate
	// The duration of the metadata item.
	Duration() coremedia.CMTime

	// Topic: Accessing language support

	// The locale of the metadata item.
	Locale() foundation.NSLocale
	// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
	ExtendedLanguageTag() string

	// The metadata items an asset contains for common metadata identifiers that provide a value.
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	// An array of metadata items for all metadata identifiers for which a value is available.
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	// Tells the object to load the values of any of the specified keys that aren’t already loaded.
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler VoidHandler)
	// Reports whether the value for a given key is immediately available without blocking.
	StatusOfValueForKeyError(key string) (AVKeyValueStatus, error)
}

// Init initializes the instance.
func (m AVMetadataItem) Init() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetadataItem) Autorelease() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataItem creates a new AVMetadataItem instance.
func NewAVMetadataItem() AVMetadataItem {
	class := getAVMetadataItemClass()
	rv := objc.Send[AVMetadataItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the object to load the values of any of the specified keys that
// aren’t already loaded.
//
// keys: An array of [NSString] objects, each of which represents one of the
// required keys.
//
// handler: The block to invoke when loading succeeds, fails, or is canceled.
//
// # Discussion
//
// For full discussion, see [AVAsynchronousKeyValueLoading].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/loadValuesAsynchronouslyForKeys:completionHandler:
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (m AVMetadataItem) LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler VoidHandler) {
	_block1, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("loadValuesAsynchronouslyForKeys:completionHandler:"), keys, _block1)
}

// Reports whether the value for a given key is immediately available without
// blocking.
//
// key: The key whose status you want.
//
// outError: If the status of the value for the key is [AVKeyValueStatusFailed], upon
// return contains an [NSError] object that describes the failure that
// occurred.
//
// # Return Value
//
// The current loading status of the value for `key`.
//
// # Discussion
//
// For full discussion, see [AVAsynchronousKeyValueLoading].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/statusOfValueForKey:error:
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (m AVMetadataItem) StatusOfValueForKeyError(key string) (AVKeyValueStatus, error) {
	var errorPtr objc.ID
	rv := objc.Send[AVKeyValueStatus](m.ID, objc.Sel("statusOfValueForKey:error:"), objc.String(key), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(AVKeyValueStatus), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Returns metadata items for the specified identifier.
//
// metadataItems: The metadata items to filter.
//
// identifier: The identifier of the metadata items to retrieve.
//
// # Return Value
//
// An array of metadata items that match the specified identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredByIdentifier:)
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemsFromArrayFilteredByIdentifier(metadataItems []AVMetadataItem, identifier AVMetadataIdentifier) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemsFromArray:filteredByIdentifier:"), objectivec.IObjectSliceToNSArray(metadataItems), objc.String(string(identifier)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns metadata items that match a specified key or key space.
//
// metadataItems: The metadata items to filter.
//
// key: The key of the metadata items to retrieve, or `nil` if you don’t want to
// filter by key.
//
// keySpace: The key space of the metadata items to retrieve, or `nil` if you don’t
// want to filter by key space.
//
// # Return Value
//
// An array of metadata items that match the specified key and key space.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:withKey:keySpace:)
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemsFromArrayWithKeyKeySpace(metadataItems []AVMetadataItem, key objectivec.IObject, keySpace AVMetadataKeySpace) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemsFromArray:withKey:keySpace:"), objectivec.IObjectSliceToNSArray(metadataItems), key, objc.String(string(keySpace)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns metadata items that match a specified locale.
//
// metadataItems: The metadata items to filter.
//
// locale: The locale of the metadata items to retrieve.
//
// # Return Value
//
// An array of metadata items that match the specified key and key space.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:with:)
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemsFromArrayWithLocale(metadataItems []AVMetadataItem, locale foundation.NSLocale) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemsFromArray:withLocale:"), objectivec.IObjectSliceToNSArray(metadataItems), locale)
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns metadata items whose locales match one of the specified language
// identifiers.
//
// metadataItems: The metadata items to filter.
//
// preferredLanguages: An array of [NSString] objects, each of which contains a canonicalized IETF
// BCP 47 language identifier. The order of the identifiers in the array
// reflects the preferred language order, with the most preferred language
// being first in the array. Typically, you pass the user’s preferred
// languages by retrieving this array from the [preferredLanguages] class
// method of [NSLocale].
//
// # Return Value
//
// An array of metadata items that match the specified languages.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredAndSortedAccordingToPreferredLanguages:)
//
// [NSLocale]: https://developer.apple.com/documentation/Foundation/NSLocale
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [preferredLanguages]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages(metadataItems []AVMetadataItem, preferredLanguages []string) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemsFromArray:filteredAndSortedAccordingToPreferredLanguages:"), objectivec.IObjectSliceToNSArray(metadataItems), objectivec.StringSliceToNSArray(preferredLanguages))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns filtered metadata items.
//
// metadataItems: The metadata items to filter.
//
// metadataItemFilter: The metadata item filter to apply.
//
// # Return Value
//
// The filtered array of metadata items.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredBy:)
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemsFromArrayFilteredByMetadataItemFilter(metadataItems []AVMetadataItem, metadataItemFilter IAVMetadataItemFilter) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemsFromArray:filteredByMetadataItemFilter:"), objectivec.IObjectSliceToNSArray(metadataItems), metadataItemFilter)
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns a metadata identifier for the specified key and key space.
//
// key: A key to return an identifier for.
//
// keySpace: A key space to return an identifier for.
//
// # Return Value
//
// A metadata identifier, or `nil` if no equivalent identifier exists.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/identifier(forKey:keySpace:)
func (_AVMetadataItemClass AVMetadataItemClass) IdentifierForKeyKeySpace(key objectivec.IObject, keySpace AVMetadataKeySpace) AVMetadataIdentifier {
	rv := objc.Send[objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("identifierForKey:keySpace:"), key, objc.String(string(keySpace)))
	return AVMetadataIdentifier(foundation.NSStringFromID(rv).String())
}

// Returns a metadata key for the specified identifier.
//
// identifier: The metadata identifier.
//
// # Return Value
//
// A metadata key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/key(forIdentifier:)
func (_AVMetadataItemClass AVMetadataItemClass) KeyForIdentifier(identifier AVMetadataIdentifier) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("keyForIdentifier:"), objc.String(string(identifier)))
	return objectivec.Object{ID: rv}
}

// Returns a metadata key space for the specified identifier.
//
// identifier: The metadata identifier.
//
// # Return Value
//
// A metadata key space.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/keySpace(forIdentifier:)
func (_AVMetadataItemClass AVMetadataItemClass) KeySpaceForIdentifier(identifier AVMetadataIdentifier) AVMetadataKeySpace {
	rv := objc.Send[objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("keySpaceForIdentifier:"), objc.String(string(identifier)))
	return AVMetadataKeySpace(foundation.NSStringFromID(rv).String())
}

// metadataItem: An instance of AVMetadataItem with the identifier, extendedLanguageTag, and
// other property values that you want the newly created instance of
// AVMetadataItem to share. The value of metadataItem is ignored.
//
// handler: A block that loads the value of the metadata item.
//
// # Return Value
//
// An instance of AVMetadataItem.
//
// # Discussion
//
// Creates an instance of AVMutableMetadataItem with a value that you do not
// wish to load unless required, e.g. a large image value that needn’t be
// loaded into memory until another module wants to display it.
//
// This method is intended for the creation of metadata items for optional
// display purposes, when there is no immediate need to load specific metadata
// values. For example, see the interface for navigation markers as consumed
// by AVPlayerViewController. It’s not intended for the creation of metadata
// items with values that are required immediately, such as metadata items
// that are provided for impending serialization operations (e.g. via
// -[AVAssetExportSession setMetadata:] and other similar methods defined on
// AVAssetWriter and AVAssetWriterInput). When
// -loadValuesAsynchronouslyForKeys:completionHandler: is invoked on an
// AVMetadataItem created via
// +metadataItemWithPropertiesOfMetadataItem:valueLoadingHandler: and
// @“value” is among the keys for which loading is requested, the block
// you provide as the value loading handler will be executed on an arbitrary
// dispatch queue, off the main thread. The handler can perform I/O and other
// necessary operations to obtain the value. If loading of the value succeeds,
// provide the value by invoking -[AVMetadataItemValueRequest
// respondWithValue:]. If loading of the value fails, provide an instance of
// NSError that describes the failure by invoking -[AVMetadataItemValueRequest
// respondWithError:].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/init(propertiesOf:valueLoadingHandler:)
func (_AVMetadataItemClass AVMetadataItemClass) MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem IAVMetadataItem, handler AVMetadataItemValueRequestHandler) AVMetadataItem {
	_block1, _ := NewAVMetadataItemValueRequestBlock(handler)
	rv := objc.Send[objc.ID](objc.ID(_AVMetadataItemClass.class), objc.Sel("metadataItemWithPropertiesOfMetadataItem:valueLoadingHandler:"), metadataItem, _block1)
	return AVMetadataItemFromID(rv)
}

// An identifier for a metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/identifier
func (m AVMetadataItem) Identifier() AVMetadataIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return AVMetadataIdentifier(foundation.NSStringFromID(rv).String())
}

// The data type of the metadata item’s value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/dataType
func (m AVMetadataItem) DataType() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataType"))
	return foundation.NSStringFromID(rv).String()
}

// The key of the metadata item.
//
// # Discussion
//
// The key property contains the true key used to identify the contents of the
// metadata item. This value is specific to the key space of the metadata
// item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/key
func (m AVMetadataItem) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

// The common key of the metadata item.
//
// # Discussion
//
// This value contains the key that most closely corresponds to the [Key]
// property, but that belongs to the common key space. You can use this key to
// locate metadata items irrespective of the underlying media format.
//
// If the value of the [KeySpace] property is [common], this property value
// contains the same key as the [Key] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/commonKey
//
// [common]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKeySpace/common
func (m AVMetadataItem) CommonKey() AVMetadataKey {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("commonKey"))
	return AVMetadataKey(foundation.NSStringFromID(rv).String())
}

// The key space for the metadata item’s key.
//
// # Discussion
//
// The key space that this property value specifies is typically the default
// key space for the metadata container that stores the metadata item.
//
// AVFoundation uses key spaces to group related sets of keys. For example,
// the framework defines different key spaces for common keys, iTunes keys,
// ID3 keys, and QuickTime keys. Key spaces aid in filtering arrays of
// metadata items.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/keySpace
func (m AVMetadataItem) KeySpace() AVMetadataKeySpace {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keySpace"))
	return AVMetadataKeySpace(foundation.NSStringFromID(rv).String())
}

// The timestamp of the metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/time
func (m AVMetadataItem) Time() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("time"))
	return coremedia.CMTime(rv)
}

// The start date of the timed metadata.
//
// # Discussion
//
// The value is `nil` if the metadata item doesn’t provide a start date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/startDate
func (m AVMetadataItem) StartDate() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("startDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The duration of the metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/duration
func (m AVMetadataItem) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}

// The locale of the metadata item.
//
// # Discussion
//
// The locale may be `nil` if no locale information is available for the
// metadata item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/locale
func (m AVMetadataItem) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}

// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// # Discussion
//
// The value may be `nil` if no language tag information is available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/extendedLanguageTag
func (m AVMetadataItem) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}

// The metadata items an asset contains for common metadata identifiers that
// provide a value.
//
// See: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (m AVMetadataItem) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("commonMetadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (m AVMetadataItem) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setCommonMetadata:"), value)
}

// An array of metadata items for all metadata identifiers for which a value
// is available.
//
// See: https://developer.apple.com/documentation/avfoundation/avasset/metadata
func (m AVMetadataItem) Metadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (m AVMetadataItem) SetMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setMetadata:"), value)
}

// Protocol methods for AVAsynchronousKeyValueLoading

// MetadataItemWithPropertiesOfMetadataItemValueLoadingHandlerSync is a synchronous wrapper around [AVMetadataItem.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc AVMetadataItemClass) MetadataItemWithPropertiesOfMetadataItemValueLoadingHandlerSync(ctx context.Context, metadataItem IAVMetadataItem) (*AVMetadataItemValueRequest, error) {
	done := make(chan *AVMetadataItemValueRequest, 1)
	mc.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem, func(val *AVMetadataItemValueRequest) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
