// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMetadataItem] class.
var (
	_NSMetadataItemClass     NSMetadataItemClass
	_NSMetadataItemClassOnce sync.Once
)

func getNSMetadataItemClass() NSMetadataItemClass {
	_NSMetadataItemClassOnce.Do(func() {
		_NSMetadataItemClass = NSMetadataItemClass{class: objc.GetClass("NSMetadataItem")}
	})
	return _NSMetadataItemClass
}

// GetNSMetadataItemClass returns the class object for NSMetadataItem.
func GetNSMetadataItemClass() NSMetadataItemClass {
	return getNSMetadataItemClass()
}

type NSMetadataItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMetadataItemClass) Alloc() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The metadata associated with a file.
//
// # Overview
// 
// Metadata items provide a simple interface to retrieve the available
// attribute names and values.
//
// # Creating a Metadata Item
//
//   - [NSMetadataItem.InitWithURL]: Initializes a metadata item with a given URL.
//
// # Getting Item Attributes
//
//   - [NSMetadataItem.Attributes]: An array containing the attribute keys for the metadata item’s values.
//   - [NSMetadataItem.ValueForAttribute]: Returns the receiver’s metadata attribute name specified by a given key.
//   - [NSMetadataItem.ValuesForAttributes]: Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
//
// # Item Attribute Keys
//
//   - [NSMetadataItem.NSMetadataItemAudiencesKey]
//   - [NSMetadataItem.NSMetadataItemAudioBitRateKey]
//   - [NSMetadataItem.NSMetadataItemAudioChannelCountKey]
//   - [NSMetadataItem.NSMetadataItemAudioEncodingApplicationKey]
//   - [NSMetadataItem.NSMetadataItemAudioSampleRateKey]
//   - [NSMetadataItem.NSMetadataItemAudioTrackNumberKey]
//   - [NSMetadataItem.NSMetadataItemAuthorAddressesKey]
//   - [NSMetadataItem.NSMetadataItemAuthorEmailAddressesKey]
//   - [NSMetadataItem.NSMetadataItemAuthorsKey]
//   - [NSMetadataItem.NSMetadataItemAcquisitionMakeKey]
//   - [NSMetadataItem.NSMetadataItemAcquisitionModelKey]
//   - [NSMetadataItem.NSMetadataItemAlbumKey]
//   - [NSMetadataItem.NSMetadataItemAltitudeKey]
//   - [NSMetadataItem.NSMetadataItemApertureKey]
//   - [NSMetadataItem.NSMetadataItemAppleLoopDescriptorsKey]
//   - [NSMetadataItem.NSMetadataItemAppleLoopsKeyFilterTypeKey]
//   - [NSMetadataItem.NSMetadataItemAppleLoopsLoopModeKey]
//   - [NSMetadataItem.NSMetadataItemAppleLoopsRootKeyKey]
//   - [NSMetadataItem.NSMetadataItemApplicationCategoriesKey]
//   - [NSMetadataItem.NSMetadataItemAttributeChangeDateKey]
//   - [NSMetadataItem.NSMetadataItemFSNameKey]
//   - [NSMetadataItem.NSMetadataItemDisplayNameKey]
//   - [NSMetadataItem.NSMetadataItemURLKey]
//   - [NSMetadataItem.NSMetadataItemPathKey]
//   - [NSMetadataItem.NSMetadataItemFSSizeKey]
//   - [NSMetadataItem.NSMetadataItemFSCreationDateKey]
//   - [NSMetadataItem.NSMetadataItemFSContentChangeDateKey]
//   - [NSMetadataItem.NSMetadataItemBitsPerSampleKey]
//   - [NSMetadataItem.NSMetadataItemCFBundleIdentifierKey]
//   - [NSMetadataItem.NSMetadataItemCameraOwnerKey]
//   - [NSMetadataItem.NSMetadataItemCityKey]
//   - [NSMetadataItem.NSMetadataItemCodecsKey]
//   - [NSMetadataItem.NSMetadataItemColorSpaceKey]
//   - [NSMetadataItem.NSMetadataItemCommentKey]
//   - [NSMetadataItem.NSMetadataItemComposerKey]
//   - [NSMetadataItem.NSMetadataItemContactKeywordsKey]
//   - [NSMetadataItem.NSMetadataItemContentCreationDateKey]
//   - [NSMetadataItem.NSMetadataItemContentModificationDateKey]
//   - [NSMetadataItem.NSMetadataItemContentTypeKey]
//   - [NSMetadataItem.NSMetadataItemContentTypeTreeKey]
//   - [NSMetadataItem.NSMetadataItemContributorsKey]
//   - [NSMetadataItem.NSMetadataItemCopyrightKey]
//   - [NSMetadataItem.NSMetadataItemCountryKey]
//   - [NSMetadataItem.NSMetadataItemCoverageKey]
//   - [NSMetadataItem.NSMetadataItemCreatorKey]
//   - [NSMetadataItem.NSMetadataItemDateAddedKey]
//   - [NSMetadataItem.NSMetadataItemDeliveryTypeKey]
//   - [NSMetadataItem.NSMetadataItemDescriptionKey]
//   - [NSMetadataItem.NSMetadataItemDirectorKey]
//   - [NSMetadataItem.NSMetadataItemDownloadedDateKey]
//   - [NSMetadataItem.NSMetadataItemDueDateKey]
//   - [NSMetadataItem.NSMetadataItemDurationSecondsKey]
//   - [NSMetadataItem.NSMetadataItemEXIFGPSVersionKey]
//   - [NSMetadataItem.NSMetadataItemEXIFVersionKey]
//   - [NSMetadataItem.NSMetadataItemEditorsKey]
//   - [NSMetadataItem.NSMetadataItemEmailAddressesKey]
//   - [NSMetadataItem.NSMetadataItemEncodingApplicationsKey]
//   - [NSMetadataItem.NSMetadataItemExecutableArchitecturesKey]
//   - [NSMetadataItem.NSMetadataItemExecutablePlatformKey]
//   - [NSMetadataItem.NSMetadataItemExposureModeKey]
//   - [NSMetadataItem.NSMetadataItemExposureProgramKey]
//   - [NSMetadataItem.NSMetadataItemExposureTimeSecondsKey]
//   - [NSMetadataItem.NSMetadataItemExposureTimeStringKey]
//   - [NSMetadataItem.NSMetadataItemFNumberKey]
//   - [NSMetadataItem.NSMetadataItemFinderCommentKey]
//   - [NSMetadataItem.NSMetadataItemFlashOnOffKey]
//   - [NSMetadataItem.NSMetadataItemFocalLength35mmKey]
//   - [NSMetadataItem.NSMetadataItemFocalLengthKey]
//   - [NSMetadataItem.NSMetadataItemFontsKey]
//   - [NSMetadataItem.NSMetadataItemGPSAreaInformationKey]
//   - [NSMetadataItem.NSMetadataItemGPSDOPKey]
//   - [NSMetadataItem.NSMetadataItemGPSDateStampKey]
//   - [NSMetadataItem.NSMetadataItemGPSDestBearingKey]
//   - [NSMetadataItem.NSMetadataItemGPSDestDistanceKey]
//   - [NSMetadataItem.NSMetadataItemGPSDestLatitudeKey]
//   - [NSMetadataItem.NSMetadataItemGPSDestLongitudeKey]
//   - [NSMetadataItem.NSMetadataItemGPSDifferentalKey]
//   - [NSMetadataItem.NSMetadataItemGPSMapDatumKey]
//   - [NSMetadataItem.NSMetadataItemGPSMeasureModeKey]
//   - [NSMetadataItem.NSMetadataItemGPSProcessingMethodKey]
//   - [NSMetadataItem.NSMetadataItemGPSStatusKey]
//   - [NSMetadataItem.NSMetadataItemGPSTrackKey]
//   - [NSMetadataItem.NSMetadataItemGenreKey]
//   - [NSMetadataItem.NSMetadataItemHasAlphaChannelKey]
//   - [NSMetadataItem.NSMetadataItemHeadlineKey]
//   - [NSMetadataItem.NSMetadataItemISOSpeedKey]
//   - [NSMetadataItem.NSMetadataItemIdentifierKey]
//   - [NSMetadataItem.NSMetadataItemImageDirectionKey]
//   - [NSMetadataItem.NSMetadataItemInformationKey]
//   - [NSMetadataItem.NSMetadataItemInstantMessageAddressesKey]
//   - [NSMetadataItem.NSMetadataItemInstructionsKey]
//   - [NSMetadataItem.NSMetadataItemIsApplicationManagedKey]
//   - [NSMetadataItem.NSMetadataItemIsGeneralMIDISequenceKey]
//   - [NSMetadataItem.NSMetadataItemIsLikelyJunkKey]
//   - [NSMetadataItem.NSMetadataItemKeySignatureKey]
//   - [NSMetadataItem.NSMetadataItemKeywordsKey]
//   - [NSMetadataItem.NSMetadataItemKindKey]
//   - [NSMetadataItem.NSMetadataItemLanguagesKey]
//   - [NSMetadataItem.NSMetadataItemLastUsedDateKey]
//   - [NSMetadataItem.NSMetadataItemLatitudeKey]
//   - [NSMetadataItem.NSMetadataItemLayerNamesKey]
//   - [NSMetadataItem.NSMetadataItemLensModelKey]
//   - [NSMetadataItem.NSMetadataItemLongitudeKey]
//   - [NSMetadataItem.NSMetadataItemLyricistKey]
//   - [NSMetadataItem.NSMetadataItemMaxApertureKey]
//   - [NSMetadataItem.NSMetadataItemMediaTypesKey]
//   - [NSMetadataItem.NSMetadataItemMeteringModeKey]
//   - [NSMetadataItem.NSMetadataItemMusicalGenreKey]
//   - [NSMetadataItem.NSMetadataItemMusicalInstrumentCategoryKey]
//   - [NSMetadataItem.NSMetadataItemMusicalInstrumentNameKey]
//   - [NSMetadataItem.NSMetadataItemNamedLocationKey]
//   - [NSMetadataItem.NSMetadataItemNumberOfPagesKey]
//   - [NSMetadataItem.NSMetadataItemOrganizationsKey]
//   - [NSMetadataItem.NSMetadataItemOrientationKey]
//   - [NSMetadataItem.NSMetadataItemOriginalFormatKey]
//   - [NSMetadataItem.NSMetadataItemOriginalSourceKey]
//   - [NSMetadataItem.NSMetadataItemPageHeightKey]
//   - [NSMetadataItem.NSMetadataItemPageWidthKey]
//   - [NSMetadataItem.NSMetadataItemParticipantsKey]
//   - [NSMetadataItem.NSMetadataItemPerformersKey]
//   - [NSMetadataItem.NSMetadataItemPhoneNumbersKey]
//   - [NSMetadataItem.NSMetadataItemPixelCountKey]
//   - [NSMetadataItem.NSMetadataItemPixelHeightKey]
//   - [NSMetadataItem.NSMetadataItemPixelWidthKey]
//   - [NSMetadataItem.NSMetadataItemProducerKey]
//   - [NSMetadataItem.NSMetadataItemProfileNameKey]
//   - [NSMetadataItem.NSMetadataItemProjectsKey]
//   - [NSMetadataItem.NSMetadataItemPublishersKey]
//   - [NSMetadataItem.NSMetadataItemRecipientAddressesKey]
//   - [NSMetadataItem.NSMetadataItemRecipientEmailAddressesKey]
//   - [NSMetadataItem.NSMetadataItemRecipientsKey]
//   - [NSMetadataItem.NSMetadataItemRecordingDateKey]
//   - [NSMetadataItem.NSMetadataItemRecordingYearKey]
//   - [NSMetadataItem.NSMetadataItemRedEyeOnOffKey]
//   - [NSMetadataItem.NSMetadataItemResolutionHeightDPIKey]
//   - [NSMetadataItem.NSMetadataItemResolutionWidthDPIKey]
//   - [NSMetadataItem.NSMetadataItemRightsKey]
//   - [NSMetadataItem.NSMetadataItemSecurityMethodKey]
//   - [NSMetadataItem.NSMetadataItemSpeedKey]
//   - [NSMetadataItem.NSMetadataItemStarRatingKey]
//   - [NSMetadataItem.NSMetadataItemStateOrProvinceKey]
//   - [NSMetadataItem.NSMetadataItemStreamableKey]
//   - [NSMetadataItem.NSMetadataItemSubjectKey]
//   - [NSMetadataItem.NSMetadataItemTempoKey]
//   - [NSMetadataItem.NSMetadataItemTextContentKey]
//   - [NSMetadataItem.NSMetadataItemThemeKey]
//   - [NSMetadataItem.NSMetadataItemTimeSignatureKey]
//   - [NSMetadataItem.NSMetadataItemTimestampKey]
//   - [NSMetadataItem.NSMetadataItemTitleKey]
//   - [NSMetadataItem.NSMetadataItemTotalBitRateKey]
//   - [NSMetadataItem.NSMetadataItemVersionKey]
//   - [NSMetadataItem.NSMetadataItemVideoBitRateKey]
//   - [NSMetadataItem.NSMetadataItemWhereFromsKey]
//   - [NSMetadataItem.NSMetadataItemWhiteBalanceKey]
//
// # iCloud Keys
//
//   - [NSMetadataItem.NSMetadataItemIsUbiquitousKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemContainerDisplayNameKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadRequestedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsExternalDocumentKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemURLInLocalContainerKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemHasUnresolvedConflictsKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsDownloadedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsDownloadingKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsUploadedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsUploadingKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemPercentDownloadedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemPercentUploadedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadingErrorKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemUploadingErrorKey]
//   - [NSMetadataItem.NSMetadataUbiquitousItemIsSharedKey]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemCurrentUserRoleKey]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemOwnerNameComponentsKey]
//
// # iCloud Download Status Values
//
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusCurrent]
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusDownloaded]
//   - [NSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusNotDownloaded]
//
// # iCloud Sharing Permissions Values
//
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemPermissionsReadOnly]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemPermissionsReadWrite]
//
// # iCloud Sharing Role Values
//
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemRoleOwner]
//   - [NSMetadataItem.NSMetadataUbiquitousSharedItemRoleParticipant]
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem
type NSMetadataItem struct {
	objectivec.Object
}

// NSMetadataItemFromID constructs a [NSMetadataItem] from an objc.ID.
//
// The metadata associated with a file.
func NSMetadataItemFromID(id objc.ID) NSMetadataItem {
	return NSMetadataItem{objectivec.Object{ID: id}}
}
// NOTE: NSMetadataItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMetadataItem] class.
//
// # Creating a Metadata Item
//
//   - [INSMetadataItem.InitWithURL]: Initializes a metadata item with a given URL.
//
// # Getting Item Attributes
//
//   - [INSMetadataItem.Attributes]: An array containing the attribute keys for the metadata item’s values.
//   - [INSMetadataItem.ValueForAttribute]: Returns the receiver’s metadata attribute name specified by a given key.
//   - [INSMetadataItem.ValuesForAttributes]: Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
//
// # Item Attribute Keys
//
//   - [INSMetadataItem.NSMetadataItemAudiencesKey]
//   - [INSMetadataItem.NSMetadataItemAudioBitRateKey]
//   - [INSMetadataItem.NSMetadataItemAudioChannelCountKey]
//   - [INSMetadataItem.NSMetadataItemAudioEncodingApplicationKey]
//   - [INSMetadataItem.NSMetadataItemAudioSampleRateKey]
//   - [INSMetadataItem.NSMetadataItemAudioTrackNumberKey]
//   - [INSMetadataItem.NSMetadataItemAuthorAddressesKey]
//   - [INSMetadataItem.NSMetadataItemAuthorEmailAddressesKey]
//   - [INSMetadataItem.NSMetadataItemAuthorsKey]
//   - [INSMetadataItem.NSMetadataItemAcquisitionMakeKey]
//   - [INSMetadataItem.NSMetadataItemAcquisitionModelKey]
//   - [INSMetadataItem.NSMetadataItemAlbumKey]
//   - [INSMetadataItem.NSMetadataItemAltitudeKey]
//   - [INSMetadataItem.NSMetadataItemApertureKey]
//   - [INSMetadataItem.NSMetadataItemAppleLoopDescriptorsKey]
//   - [INSMetadataItem.NSMetadataItemAppleLoopsKeyFilterTypeKey]
//   - [INSMetadataItem.NSMetadataItemAppleLoopsLoopModeKey]
//   - [INSMetadataItem.NSMetadataItemAppleLoopsRootKeyKey]
//   - [INSMetadataItem.NSMetadataItemApplicationCategoriesKey]
//   - [INSMetadataItem.NSMetadataItemAttributeChangeDateKey]
//   - [INSMetadataItem.NSMetadataItemFSNameKey]
//   - [INSMetadataItem.NSMetadataItemDisplayNameKey]
//   - [INSMetadataItem.NSMetadataItemURLKey]
//   - [INSMetadataItem.NSMetadataItemPathKey]
//   - [INSMetadataItem.NSMetadataItemFSSizeKey]
//   - [INSMetadataItem.NSMetadataItemFSCreationDateKey]
//   - [INSMetadataItem.NSMetadataItemFSContentChangeDateKey]
//   - [INSMetadataItem.NSMetadataItemBitsPerSampleKey]
//   - [INSMetadataItem.NSMetadataItemCFBundleIdentifierKey]
//   - [INSMetadataItem.NSMetadataItemCameraOwnerKey]
//   - [INSMetadataItem.NSMetadataItemCityKey]
//   - [INSMetadataItem.NSMetadataItemCodecsKey]
//   - [INSMetadataItem.NSMetadataItemColorSpaceKey]
//   - [INSMetadataItem.NSMetadataItemCommentKey]
//   - [INSMetadataItem.NSMetadataItemComposerKey]
//   - [INSMetadataItem.NSMetadataItemContactKeywordsKey]
//   - [INSMetadataItem.NSMetadataItemContentCreationDateKey]
//   - [INSMetadataItem.NSMetadataItemContentModificationDateKey]
//   - [INSMetadataItem.NSMetadataItemContentTypeKey]
//   - [INSMetadataItem.NSMetadataItemContentTypeTreeKey]
//   - [INSMetadataItem.NSMetadataItemContributorsKey]
//   - [INSMetadataItem.NSMetadataItemCopyrightKey]
//   - [INSMetadataItem.NSMetadataItemCountryKey]
//   - [INSMetadataItem.NSMetadataItemCoverageKey]
//   - [INSMetadataItem.NSMetadataItemCreatorKey]
//   - [INSMetadataItem.NSMetadataItemDateAddedKey]
//   - [INSMetadataItem.NSMetadataItemDeliveryTypeKey]
//   - [INSMetadataItem.NSMetadataItemDescriptionKey]
//   - [INSMetadataItem.NSMetadataItemDirectorKey]
//   - [INSMetadataItem.NSMetadataItemDownloadedDateKey]
//   - [INSMetadataItem.NSMetadataItemDueDateKey]
//   - [INSMetadataItem.NSMetadataItemDurationSecondsKey]
//   - [INSMetadataItem.NSMetadataItemEXIFGPSVersionKey]
//   - [INSMetadataItem.NSMetadataItemEXIFVersionKey]
//   - [INSMetadataItem.NSMetadataItemEditorsKey]
//   - [INSMetadataItem.NSMetadataItemEmailAddressesKey]
//   - [INSMetadataItem.NSMetadataItemEncodingApplicationsKey]
//   - [INSMetadataItem.NSMetadataItemExecutableArchitecturesKey]
//   - [INSMetadataItem.NSMetadataItemExecutablePlatformKey]
//   - [INSMetadataItem.NSMetadataItemExposureModeKey]
//   - [INSMetadataItem.NSMetadataItemExposureProgramKey]
//   - [INSMetadataItem.NSMetadataItemExposureTimeSecondsKey]
//   - [INSMetadataItem.NSMetadataItemExposureTimeStringKey]
//   - [INSMetadataItem.NSMetadataItemFNumberKey]
//   - [INSMetadataItem.NSMetadataItemFinderCommentKey]
//   - [INSMetadataItem.NSMetadataItemFlashOnOffKey]
//   - [INSMetadataItem.NSMetadataItemFocalLength35mmKey]
//   - [INSMetadataItem.NSMetadataItemFocalLengthKey]
//   - [INSMetadataItem.NSMetadataItemFontsKey]
//   - [INSMetadataItem.NSMetadataItemGPSAreaInformationKey]
//   - [INSMetadataItem.NSMetadataItemGPSDOPKey]
//   - [INSMetadataItem.NSMetadataItemGPSDateStampKey]
//   - [INSMetadataItem.NSMetadataItemGPSDestBearingKey]
//   - [INSMetadataItem.NSMetadataItemGPSDestDistanceKey]
//   - [INSMetadataItem.NSMetadataItemGPSDestLatitudeKey]
//   - [INSMetadataItem.NSMetadataItemGPSDestLongitudeKey]
//   - [INSMetadataItem.NSMetadataItemGPSDifferentalKey]
//   - [INSMetadataItem.NSMetadataItemGPSMapDatumKey]
//   - [INSMetadataItem.NSMetadataItemGPSMeasureModeKey]
//   - [INSMetadataItem.NSMetadataItemGPSProcessingMethodKey]
//   - [INSMetadataItem.NSMetadataItemGPSStatusKey]
//   - [INSMetadataItem.NSMetadataItemGPSTrackKey]
//   - [INSMetadataItem.NSMetadataItemGenreKey]
//   - [INSMetadataItem.NSMetadataItemHasAlphaChannelKey]
//   - [INSMetadataItem.NSMetadataItemHeadlineKey]
//   - [INSMetadataItem.NSMetadataItemISOSpeedKey]
//   - [INSMetadataItem.NSMetadataItemIdentifierKey]
//   - [INSMetadataItem.NSMetadataItemImageDirectionKey]
//   - [INSMetadataItem.NSMetadataItemInformationKey]
//   - [INSMetadataItem.NSMetadataItemInstantMessageAddressesKey]
//   - [INSMetadataItem.NSMetadataItemInstructionsKey]
//   - [INSMetadataItem.NSMetadataItemIsApplicationManagedKey]
//   - [INSMetadataItem.NSMetadataItemIsGeneralMIDISequenceKey]
//   - [INSMetadataItem.NSMetadataItemIsLikelyJunkKey]
//   - [INSMetadataItem.NSMetadataItemKeySignatureKey]
//   - [INSMetadataItem.NSMetadataItemKeywordsKey]
//   - [INSMetadataItem.NSMetadataItemKindKey]
//   - [INSMetadataItem.NSMetadataItemLanguagesKey]
//   - [INSMetadataItem.NSMetadataItemLastUsedDateKey]
//   - [INSMetadataItem.NSMetadataItemLatitudeKey]
//   - [INSMetadataItem.NSMetadataItemLayerNamesKey]
//   - [INSMetadataItem.NSMetadataItemLensModelKey]
//   - [INSMetadataItem.NSMetadataItemLongitudeKey]
//   - [INSMetadataItem.NSMetadataItemLyricistKey]
//   - [INSMetadataItem.NSMetadataItemMaxApertureKey]
//   - [INSMetadataItem.NSMetadataItemMediaTypesKey]
//   - [INSMetadataItem.NSMetadataItemMeteringModeKey]
//   - [INSMetadataItem.NSMetadataItemMusicalGenreKey]
//   - [INSMetadataItem.NSMetadataItemMusicalInstrumentCategoryKey]
//   - [INSMetadataItem.NSMetadataItemMusicalInstrumentNameKey]
//   - [INSMetadataItem.NSMetadataItemNamedLocationKey]
//   - [INSMetadataItem.NSMetadataItemNumberOfPagesKey]
//   - [INSMetadataItem.NSMetadataItemOrganizationsKey]
//   - [INSMetadataItem.NSMetadataItemOrientationKey]
//   - [INSMetadataItem.NSMetadataItemOriginalFormatKey]
//   - [INSMetadataItem.NSMetadataItemOriginalSourceKey]
//   - [INSMetadataItem.NSMetadataItemPageHeightKey]
//   - [INSMetadataItem.NSMetadataItemPageWidthKey]
//   - [INSMetadataItem.NSMetadataItemParticipantsKey]
//   - [INSMetadataItem.NSMetadataItemPerformersKey]
//   - [INSMetadataItem.NSMetadataItemPhoneNumbersKey]
//   - [INSMetadataItem.NSMetadataItemPixelCountKey]
//   - [INSMetadataItem.NSMetadataItemPixelHeightKey]
//   - [INSMetadataItem.NSMetadataItemPixelWidthKey]
//   - [INSMetadataItem.NSMetadataItemProducerKey]
//   - [INSMetadataItem.NSMetadataItemProfileNameKey]
//   - [INSMetadataItem.NSMetadataItemProjectsKey]
//   - [INSMetadataItem.NSMetadataItemPublishersKey]
//   - [INSMetadataItem.NSMetadataItemRecipientAddressesKey]
//   - [INSMetadataItem.NSMetadataItemRecipientEmailAddressesKey]
//   - [INSMetadataItem.NSMetadataItemRecipientsKey]
//   - [INSMetadataItem.NSMetadataItemRecordingDateKey]
//   - [INSMetadataItem.NSMetadataItemRecordingYearKey]
//   - [INSMetadataItem.NSMetadataItemRedEyeOnOffKey]
//   - [INSMetadataItem.NSMetadataItemResolutionHeightDPIKey]
//   - [INSMetadataItem.NSMetadataItemResolutionWidthDPIKey]
//   - [INSMetadataItem.NSMetadataItemRightsKey]
//   - [INSMetadataItem.NSMetadataItemSecurityMethodKey]
//   - [INSMetadataItem.NSMetadataItemSpeedKey]
//   - [INSMetadataItem.NSMetadataItemStarRatingKey]
//   - [INSMetadataItem.NSMetadataItemStateOrProvinceKey]
//   - [INSMetadataItem.NSMetadataItemStreamableKey]
//   - [INSMetadataItem.NSMetadataItemSubjectKey]
//   - [INSMetadataItem.NSMetadataItemTempoKey]
//   - [INSMetadataItem.NSMetadataItemTextContentKey]
//   - [INSMetadataItem.NSMetadataItemThemeKey]
//   - [INSMetadataItem.NSMetadataItemTimeSignatureKey]
//   - [INSMetadataItem.NSMetadataItemTimestampKey]
//   - [INSMetadataItem.NSMetadataItemTitleKey]
//   - [INSMetadataItem.NSMetadataItemTotalBitRateKey]
//   - [INSMetadataItem.NSMetadataItemVersionKey]
//   - [INSMetadataItem.NSMetadataItemVideoBitRateKey]
//   - [INSMetadataItem.NSMetadataItemWhereFromsKey]
//   - [INSMetadataItem.NSMetadataItemWhiteBalanceKey]
//
// # iCloud Keys
//
//   - [INSMetadataItem.NSMetadataItemIsUbiquitousKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemContainerDisplayNameKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadRequestedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsExternalDocumentKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemURLInLocalContainerKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemHasUnresolvedConflictsKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsDownloadedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsDownloadingKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsUploadedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsUploadingKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemPercentDownloadedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemPercentUploadedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadingErrorKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemUploadingErrorKey]
//   - [INSMetadataItem.NSMetadataUbiquitousItemIsSharedKey]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemCurrentUserRoleKey]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemOwnerNameComponentsKey]
//
// # iCloud Download Status Values
//
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusCurrent]
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusDownloaded]
//   - [INSMetadataItem.NSMetadataUbiquitousItemDownloadingStatusNotDownloaded]
//
// # iCloud Sharing Permissions Values
//
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemPermissionsReadOnly]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemPermissionsReadWrite]
//
// # iCloud Sharing Role Values
//
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemRoleOwner]
//   - [INSMetadataItem.NSMetadataUbiquitousSharedItemRoleParticipant]
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem
type INSMetadataItem interface {
	objectivec.IObject

	// Topic: Creating a Metadata Item

	// Initializes a metadata item with a given URL.
	InitWithURL(url INSURL) NSMetadataItem

	// Topic: Getting Item Attributes

	// An array containing the attribute keys for the metadata item’s values.
	Attributes() []string
	// Returns the receiver’s metadata attribute name specified by a given key.
	ValueForAttribute(key string) objectivec.IObject
	// Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
	ValuesForAttributes(keys []string) INSDictionary

	// Topic: Item Attribute Keys

	NSMetadataItemAudiencesKey() string
	NSMetadataItemAudioBitRateKey() string
	NSMetadataItemAudioChannelCountKey() string
	NSMetadataItemAudioEncodingApplicationKey() string
	NSMetadataItemAudioSampleRateKey() string
	NSMetadataItemAudioTrackNumberKey() string
	NSMetadataItemAuthorAddressesKey() string
	NSMetadataItemAuthorEmailAddressesKey() string
	NSMetadataItemAuthorsKey() string
	NSMetadataItemAcquisitionMakeKey() string
	NSMetadataItemAcquisitionModelKey() string
	NSMetadataItemAlbumKey() string
	NSMetadataItemAltitudeKey() string
	NSMetadataItemApertureKey() string
	NSMetadataItemAppleLoopDescriptorsKey() string
	NSMetadataItemAppleLoopsKeyFilterTypeKey() string
	NSMetadataItemAppleLoopsLoopModeKey() string
	NSMetadataItemAppleLoopsRootKeyKey() string
	NSMetadataItemApplicationCategoriesKey() string
	NSMetadataItemAttributeChangeDateKey() string
	NSMetadataItemFSNameKey() string
	NSMetadataItemDisplayNameKey() string
	NSMetadataItemURLKey() string
	NSMetadataItemPathKey() string
	NSMetadataItemFSSizeKey() string
	NSMetadataItemFSCreationDateKey() string
	NSMetadataItemFSContentChangeDateKey() string
	NSMetadataItemBitsPerSampleKey() string
	NSMetadataItemCFBundleIdentifierKey() string
	NSMetadataItemCameraOwnerKey() string
	NSMetadataItemCityKey() string
	NSMetadataItemCodecsKey() string
	NSMetadataItemColorSpaceKey() string
	NSMetadataItemCommentKey() string
	NSMetadataItemComposerKey() string
	NSMetadataItemContactKeywordsKey() string
	NSMetadataItemContentCreationDateKey() string
	NSMetadataItemContentModificationDateKey() string
	NSMetadataItemContentTypeKey() string
	NSMetadataItemContentTypeTreeKey() string
	NSMetadataItemContributorsKey() string
	NSMetadataItemCopyrightKey() string
	NSMetadataItemCountryKey() string
	NSMetadataItemCoverageKey() string
	NSMetadataItemCreatorKey() string
	NSMetadataItemDateAddedKey() string
	NSMetadataItemDeliveryTypeKey() string
	NSMetadataItemDescriptionKey() string
	NSMetadataItemDirectorKey() string
	NSMetadataItemDownloadedDateKey() string
	NSMetadataItemDueDateKey() string
	NSMetadataItemDurationSecondsKey() string
	NSMetadataItemEXIFGPSVersionKey() string
	NSMetadataItemEXIFVersionKey() string
	NSMetadataItemEditorsKey() string
	NSMetadataItemEmailAddressesKey() string
	NSMetadataItemEncodingApplicationsKey() string
	NSMetadataItemExecutableArchitecturesKey() string
	NSMetadataItemExecutablePlatformKey() string
	NSMetadataItemExposureModeKey() string
	NSMetadataItemExposureProgramKey() string
	NSMetadataItemExposureTimeSecondsKey() string
	NSMetadataItemExposureTimeStringKey() string
	NSMetadataItemFNumberKey() string
	NSMetadataItemFinderCommentKey() string
	NSMetadataItemFlashOnOffKey() string
	NSMetadataItemFocalLength35mmKey() string
	NSMetadataItemFocalLengthKey() string
	NSMetadataItemFontsKey() string
	NSMetadataItemGPSAreaInformationKey() string
	NSMetadataItemGPSDOPKey() string
	NSMetadataItemGPSDateStampKey() string
	NSMetadataItemGPSDestBearingKey() string
	NSMetadataItemGPSDestDistanceKey() string
	NSMetadataItemGPSDestLatitudeKey() string
	NSMetadataItemGPSDestLongitudeKey() string
	NSMetadataItemGPSDifferentalKey() string
	NSMetadataItemGPSMapDatumKey() string
	NSMetadataItemGPSMeasureModeKey() string
	NSMetadataItemGPSProcessingMethodKey() string
	NSMetadataItemGPSStatusKey() string
	NSMetadataItemGPSTrackKey() string
	NSMetadataItemGenreKey() string
	NSMetadataItemHasAlphaChannelKey() string
	NSMetadataItemHeadlineKey() string
	NSMetadataItemISOSpeedKey() string
	NSMetadataItemIdentifierKey() string
	NSMetadataItemImageDirectionKey() string
	NSMetadataItemInformationKey() string
	NSMetadataItemInstantMessageAddressesKey() string
	NSMetadataItemInstructionsKey() string
	NSMetadataItemIsApplicationManagedKey() string
	NSMetadataItemIsGeneralMIDISequenceKey() string
	NSMetadataItemIsLikelyJunkKey() string
	NSMetadataItemKeySignatureKey() string
	NSMetadataItemKeywordsKey() string
	NSMetadataItemKindKey() string
	NSMetadataItemLanguagesKey() string
	NSMetadataItemLastUsedDateKey() string
	NSMetadataItemLatitudeKey() string
	NSMetadataItemLayerNamesKey() string
	NSMetadataItemLensModelKey() string
	NSMetadataItemLongitudeKey() string
	NSMetadataItemLyricistKey() string
	NSMetadataItemMaxApertureKey() string
	NSMetadataItemMediaTypesKey() string
	NSMetadataItemMeteringModeKey() string
	NSMetadataItemMusicalGenreKey() string
	NSMetadataItemMusicalInstrumentCategoryKey() string
	NSMetadataItemMusicalInstrumentNameKey() string
	NSMetadataItemNamedLocationKey() string
	NSMetadataItemNumberOfPagesKey() string
	NSMetadataItemOrganizationsKey() string
	NSMetadataItemOrientationKey() string
	NSMetadataItemOriginalFormatKey() string
	NSMetadataItemOriginalSourceKey() string
	NSMetadataItemPageHeightKey() string
	NSMetadataItemPageWidthKey() string
	NSMetadataItemParticipantsKey() string
	NSMetadataItemPerformersKey() string
	NSMetadataItemPhoneNumbersKey() string
	NSMetadataItemPixelCountKey() string
	NSMetadataItemPixelHeightKey() string
	NSMetadataItemPixelWidthKey() string
	NSMetadataItemProducerKey() string
	NSMetadataItemProfileNameKey() string
	NSMetadataItemProjectsKey() string
	NSMetadataItemPublishersKey() string
	NSMetadataItemRecipientAddressesKey() string
	NSMetadataItemRecipientEmailAddressesKey() string
	NSMetadataItemRecipientsKey() string
	NSMetadataItemRecordingDateKey() string
	NSMetadataItemRecordingYearKey() string
	NSMetadataItemRedEyeOnOffKey() string
	NSMetadataItemResolutionHeightDPIKey() string
	NSMetadataItemResolutionWidthDPIKey() string
	NSMetadataItemRightsKey() string
	NSMetadataItemSecurityMethodKey() string
	NSMetadataItemSpeedKey() string
	NSMetadataItemStarRatingKey() string
	NSMetadataItemStateOrProvinceKey() string
	NSMetadataItemStreamableKey() string
	NSMetadataItemSubjectKey() string
	NSMetadataItemTempoKey() string
	NSMetadataItemTextContentKey() string
	NSMetadataItemThemeKey() string
	NSMetadataItemTimeSignatureKey() string
	NSMetadataItemTimestampKey() string
	NSMetadataItemTitleKey() string
	NSMetadataItemTotalBitRateKey() string
	NSMetadataItemVersionKey() string
	NSMetadataItemVideoBitRateKey() string
	NSMetadataItemWhereFromsKey() string
	NSMetadataItemWhiteBalanceKey() string

	// Topic: iCloud Keys

	NSMetadataItemIsUbiquitousKey() string
	NSMetadataUbiquitousItemContainerDisplayNameKey() string
	NSMetadataUbiquitousItemDownloadRequestedKey() string
	NSMetadataUbiquitousItemIsExternalDocumentKey() string
	NSMetadataUbiquitousItemURLInLocalContainerKey() string
	NSMetadataUbiquitousItemHasUnresolvedConflictsKey() string
	NSMetadataUbiquitousItemIsDownloadedKey() string
	NSMetadataUbiquitousItemIsDownloadingKey() string
	NSMetadataUbiquitousItemIsUploadedKey() string
	NSMetadataUbiquitousItemIsUploadingKey() string
	NSMetadataUbiquitousItemPercentDownloadedKey() string
	NSMetadataUbiquitousItemPercentUploadedKey() string
	NSMetadataUbiquitousItemDownloadingStatusKey() string
	NSMetadataUbiquitousItemDownloadingErrorKey() string
	NSMetadataUbiquitousItemUploadingErrorKey() string
	NSMetadataUbiquitousItemIsSharedKey() string
	NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey() string
	NSMetadataUbiquitousSharedItemCurrentUserRoleKey() string
	NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey() string
	NSMetadataUbiquitousSharedItemOwnerNameComponentsKey() string

	// Topic: iCloud Download Status Values

	NSMetadataUbiquitousItemDownloadingStatusCurrent() string
	NSMetadataUbiquitousItemDownloadingStatusDownloaded() string
	NSMetadataUbiquitousItemDownloadingStatusNotDownloaded() string

	// Topic: iCloud Sharing Permissions Values

	NSMetadataUbiquitousSharedItemPermissionsReadOnly() string
	NSMetadataUbiquitousSharedItemPermissionsReadWrite() string

	// Topic: iCloud Sharing Role Values

	NSMetadataUbiquitousSharedItemRoleOwner() string
	NSMetadataUbiquitousSharedItemRoleParticipant() string
}

// Init initializes the instance.
func (m NSMetadataItem) Init() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMetadataItem) Autorelease() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMetadataItem creates a new NSMetadataItem instance.
func NewNSMetadataItem() NSMetadataItem {
	class := getNSMetadataItemClass()
	rv := objc.Send[NSMetadataItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a metadata item with a given URL.
//
// url: The URL for the metadata item.
//
// # Return Value
// 
// A metadata item for the file identified by `url`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/init(url:)
func NewMetadataItemWithURL(url INSURL) NSMetadataItem {
	instance := getNSMetadataItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return NSMetadataItemFromID(rv)
}

// Initializes a metadata item with a given URL.
//
// url: The URL for the metadata item.
//
// # Return Value
// 
// A metadata item for the file identified by `url`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/init(url:)
func (m NSMetadataItem) InitWithURL(url INSURL) NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("initWithURL:"), url)
	return rv
}
// Returns the receiver’s metadata attribute name specified by a given key.
//
// key: The name of a metadata attribute. See the “Constants” section for a
// list of possible keys.
//
// # Return Value
// 
// The receiver’s metadata attribute name specified by `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/value(forAttribute:)
func (m NSMetadataItem) ValueForAttribute(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("valueForAttribute:"), objc.String(key))
	return objectivec.Object{ID: rv}
}
// Returns a dictionary containing the key-value pairs for the attribute names
// specified by a given array of keys.
//
// keys: An array containing [NSString] objects that specify the names of a metadata
// attributes. See the “Constants” section for a list of possible keys.
//
// # Return Value
// 
// A dictionary containing the key-value pairs for the attribute names
// specified by `keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/values(forAttributes:)
func (m NSMetadataItem) ValuesForAttributes(keys []string) INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("valuesForAttributes:"), objectivec.StringSliceToNSArray(keys))
	return NSDictionaryFromID(rv)
}

// An array containing the attribute keys for the metadata item’s values.
//
// # Discussion
// 
// This property contains an array of attribute keys, representing the values
// available from this metadata item. For a list of possible keys, see
// `Attribute Keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/attributes
func (m NSMetadataItem) Attributes() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("attributes"))
	return objc.ConvertSliceToStrings(rv)
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudienceskey
func (m NSMetadataItem) NSMetadataItemAudiencesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudiencesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiobitratekey
func (m NSMetadataItem) NSMetadataItemAudioBitRateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudioBitRateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiochannelcountkey
func (m NSMetadataItem) NSMetadataItemAudioChannelCountKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudioChannelCountKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudioencodingapplicationkey
func (m NSMetadataItem) NSMetadataItemAudioEncodingApplicationKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudioEncodingApplicationKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiosampleratekey
func (m NSMetadataItem) NSMetadataItemAudioSampleRateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudioSampleRateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiotracknumberkey
func (m NSMetadataItem) NSMetadataItemAudioTrackNumberKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAudioTrackNumberKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoraddresseskey
func (m NSMetadataItem) NSMetadataItemAuthorAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAuthorAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoremailaddresseskey
func (m NSMetadataItem) NSMetadataItemAuthorEmailAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAuthorEmailAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemauthorskey
func (m NSMetadataItem) NSMetadataItemAuthorsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAuthorsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmakekey
func (m NSMetadataItem) NSMetadataItemAcquisitionMakeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAcquisitionMakeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmodelkey
func (m NSMetadataItem) NSMetadataItemAcquisitionModelKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAcquisitionModelKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemalbumkey
func (m NSMetadataItem) NSMetadataItemAlbumKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAlbumKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaltitudekey
func (m NSMetadataItem) NSMetadataItemAltitudeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAltitudeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemaperturekey
func (m NSMetadataItem) NSMetadataItemApertureKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemApertureKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopdescriptorskey
func (m NSMetadataItem) NSMetadataItemAppleLoopDescriptorsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAppleLoopDescriptorsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopskeyfiltertypekey
func (m NSMetadataItem) NSMetadataItemAppleLoopsKeyFilterTypeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAppleLoopsKeyFilterTypeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsloopmodekey
func (m NSMetadataItem) NSMetadataItemAppleLoopsLoopModeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAppleLoopsLoopModeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsrootkeykey
func (m NSMetadataItem) NSMetadataItemAppleLoopsRootKeyKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAppleLoopsRootKeyKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemapplicationcategorieskey
func (m NSMetadataItem) NSMetadataItemApplicationCategoriesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemApplicationCategoriesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemattributechangedatekey
func (m NSMetadataItem) NSMetadataItemAttributeChangeDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemAttributeChangeDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfsnamekey
func (m NSMetadataItem) NSMetadataItemFSNameKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFSNameKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdisplaynamekey
func (m NSMetadataItem) NSMetadataItemDisplayNameKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDisplayNameKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemurlkey
func (m NSMetadataItem) NSMetadataItemURLKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemURLKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempathkey
func (m NSMetadataItem) NSMetadataItemPathKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPathKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfssizekey
func (m NSMetadataItem) NSMetadataItemFSSizeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFSSizeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfscreationdatekey
func (m NSMetadataItem) NSMetadataItemFSCreationDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFSCreationDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfscontentchangedatekey
func (m NSMetadataItem) NSMetadataItemFSContentChangeDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFSContentChangeDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitembitspersamplekey
func (m NSMetadataItem) NSMetadataItemBitsPerSampleKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemBitsPerSampleKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcfbundleidentifierkey
func (m NSMetadataItem) NSMetadataItemCFBundleIdentifierKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCFBundleIdentifierKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcameraownerkey
func (m NSMetadataItem) NSMetadataItemCameraOwnerKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCameraOwnerKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcitykey
func (m NSMetadataItem) NSMetadataItemCityKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCityKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcodecskey
func (m NSMetadataItem) NSMetadataItemCodecsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCodecsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcolorspacekey
func (m NSMetadataItem) NSMetadataItemColorSpaceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemColorSpaceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcommentkey
func (m NSMetadataItem) NSMetadataItemCommentKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCommentKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcomposerkey
func (m NSMetadataItem) NSMetadataItemComposerKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemComposerKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontactkeywordskey
func (m NSMetadataItem) NSMetadataItemContactKeywordsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContactKeywordsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentcreationdatekey
func (m NSMetadataItem) NSMetadataItemContentCreationDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContentCreationDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentmodificationdatekey
func (m NSMetadataItem) NSMetadataItemContentModificationDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContentModificationDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypekey
func (m NSMetadataItem) NSMetadataItemContentTypeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContentTypeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypetreekey
func (m NSMetadataItem) NSMetadataItemContentTypeTreeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContentTypeTreeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcontributorskey
func (m NSMetadataItem) NSMetadataItemContributorsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemContributorsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcopyrightkey
func (m NSMetadataItem) NSMetadataItemCopyrightKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCopyrightKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcountrykey
func (m NSMetadataItem) NSMetadataItemCountryKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCountryKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcoveragekey
func (m NSMetadataItem) NSMetadataItemCoverageKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCoverageKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemcreatorkey
func (m NSMetadataItem) NSMetadataItemCreatorKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemCreatorKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdateaddedkey
func (m NSMetadataItem) NSMetadataItemDateAddedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDateAddedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdeliverytypekey
func (m NSMetadataItem) NSMetadataItemDeliveryTypeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDeliveryTypeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdescriptionkey
func (m NSMetadataItem) NSMetadataItemDescriptionKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDescriptionKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdirectorkey
func (m NSMetadataItem) NSMetadataItemDirectorKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDirectorKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdownloadeddatekey
func (m NSMetadataItem) NSMetadataItemDownloadedDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDownloadedDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemduedatekey
func (m NSMetadataItem) NSMetadataItemDueDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDueDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemdurationsecondskey
func (m NSMetadataItem) NSMetadataItemDurationSecondsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemDurationSecondsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexifgpsversionkey
func (m NSMetadataItem) NSMetadataItemEXIFGPSVersionKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemEXIFGPSVersionKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexifversionkey
func (m NSMetadataItem) NSMetadataItemEXIFVersionKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemEXIFVersionKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemeditorskey
func (m NSMetadataItem) NSMetadataItemEditorsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemEditorsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitememailaddresseskey
func (m NSMetadataItem) NSMetadataItemEmailAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemEmailAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemencodingapplicationskey
func (m NSMetadataItem) NSMetadataItemEncodingApplicationsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemEncodingApplicationsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutablearchitectureskey
func (m NSMetadataItem) NSMetadataItemExecutableArchitecturesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExecutableArchitecturesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutableplatformkey
func (m NSMetadataItem) NSMetadataItemExecutablePlatformKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExecutablePlatformKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuremodekey
func (m NSMetadataItem) NSMetadataItemExposureModeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExposureModeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexposureprogramkey
func (m NSMetadataItem) NSMetadataItemExposureProgramKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExposureProgramKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimesecondskey
func (m NSMetadataItem) NSMetadataItemExposureTimeSecondsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExposureTimeSecondsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimestringkey
func (m NSMetadataItem) NSMetadataItemExposureTimeStringKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemExposureTimeStringKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfnumberkey
func (m NSMetadataItem) NSMetadataItemFNumberKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFNumberKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfindercommentkey
func (m NSMetadataItem) NSMetadataItemFinderCommentKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFinderCommentKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemflashonoffkey
func (m NSMetadataItem) NSMetadataItemFlashOnOffKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFlashOnOffKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallength35mmkey
func (m NSMetadataItem) NSMetadataItemFocalLength35mmKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFocalLength35mmKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallengthkey
func (m NSMetadataItem) NSMetadataItemFocalLengthKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFocalLengthKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemfontskey
func (m NSMetadataItem) NSMetadataItemFontsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemFontsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsareainformationkey
func (m NSMetadataItem) NSMetadataItemGPSAreaInformationKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSAreaInformationKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdopkey
func (m NSMetadataItem) NSMetadataItemGPSDOPKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDOPKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdatestampkey
func (m NSMetadataItem) NSMetadataItemGPSDateStampKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDateStampKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestbearingkey
func (m NSMetadataItem) NSMetadataItemGPSDestBearingKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDestBearingKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestdistancekey
func (m NSMetadataItem) NSMetadataItemGPSDestDistanceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDestDistanceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlatitudekey
func (m NSMetadataItem) NSMetadataItemGPSDestLatitudeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDestLatitudeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlongitudekey
func (m NSMetadataItem) NSMetadataItemGPSDestLongitudeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDestLongitudeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdifferentalkey
func (m NSMetadataItem) NSMetadataItemGPSDifferentalKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSDifferentalKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmapdatumkey
func (m NSMetadataItem) NSMetadataItemGPSMapDatumKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSMapDatumKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmeasuremodekey
func (m NSMetadataItem) NSMetadataItemGPSMeasureModeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSMeasureModeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsprocessingmethodkey
func (m NSMetadataItem) NSMetadataItemGPSProcessingMethodKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSProcessingMethodKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsstatuskey
func (m NSMetadataItem) NSMetadataItemGPSStatusKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSStatusKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgpstrackkey
func (m NSMetadataItem) NSMetadataItemGPSTrackKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGPSTrackKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemgenrekey
func (m NSMetadataItem) NSMetadataItemGenreKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemGenreKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemhasalphachannelkey
func (m NSMetadataItem) NSMetadataItemHasAlphaChannelKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemHasAlphaChannelKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemheadlinekey
func (m NSMetadataItem) NSMetadataItemHeadlineKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemHeadlineKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemisospeedkey
func (m NSMetadataItem) NSMetadataItemISOSpeedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemISOSpeedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemidentifierkey
func (m NSMetadataItem) NSMetadataItemIdentifierKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemIdentifierKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemimagedirectionkey
func (m NSMetadataItem) NSMetadataItemImageDirectionKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemImageDirectionKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataiteminformationkey
func (m NSMetadataItem) NSMetadataItemInformationKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemInformationKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataiteminstantmessageaddresseskey
func (m NSMetadataItem) NSMetadataItemInstantMessageAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemInstantMessageAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataiteminstructionskey
func (m NSMetadataItem) NSMetadataItemInstructionsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemInstructionsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemisapplicationmanagedkey
func (m NSMetadataItem) NSMetadataItemIsApplicationManagedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemIsApplicationManagedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemisgeneralmidisequencekey
func (m NSMetadataItem) NSMetadataItemIsGeneralMIDISequenceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemIsGeneralMIDISequenceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemislikelyjunkkey
func (m NSMetadataItem) NSMetadataItemIsLikelyJunkKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemIsLikelyJunkKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemkeysignaturekey
func (m NSMetadataItem) NSMetadataItemKeySignatureKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemKeySignatureKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemkeywordskey
func (m NSMetadataItem) NSMetadataItemKeywordsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemKeywordsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemkindkey
func (m NSMetadataItem) NSMetadataItemKindKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemKindKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlanguageskey
func (m NSMetadataItem) NSMetadataItemLanguagesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLanguagesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlastuseddatekey
func (m NSMetadataItem) NSMetadataItemLastUsedDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLastUsedDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlatitudekey
func (m NSMetadataItem) NSMetadataItemLatitudeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLatitudeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlayernameskey
func (m NSMetadataItem) NSMetadataItemLayerNamesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLayerNamesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlensmodelkey
func (m NSMetadataItem) NSMetadataItemLensModelKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLensModelKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlongitudekey
func (m NSMetadataItem) NSMetadataItemLongitudeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLongitudeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemlyricistkey
func (m NSMetadataItem) NSMetadataItemLyricistKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemLyricistKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmaxaperturekey
func (m NSMetadataItem) NSMetadataItemMaxApertureKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMaxApertureKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmediatypeskey
func (m NSMetadataItem) NSMetadataItemMediaTypesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMediaTypesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmeteringmodekey
func (m NSMetadataItem) NSMetadataItemMeteringModeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMeteringModeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalgenrekey
func (m NSMetadataItem) NSMetadataItemMusicalGenreKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMusicalGenreKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentcategorykey
func (m NSMetadataItem) NSMetadataItemMusicalInstrumentCategoryKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMusicalInstrumentCategoryKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentnamekey
func (m NSMetadataItem) NSMetadataItemMusicalInstrumentNameKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemMusicalInstrumentNameKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemnamedlocationkey
func (m NSMetadataItem) NSMetadataItemNamedLocationKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemNamedLocationKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemnumberofpageskey
func (m NSMetadataItem) NSMetadataItemNumberOfPagesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemNumberOfPagesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemorganizationskey
func (m NSMetadataItem) NSMetadataItemOrganizationsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemOrganizationsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemorientationkey
func (m NSMetadataItem) NSMetadataItemOrientationKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemOrientationKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalformatkey
func (m NSMetadataItem) NSMetadataItemOriginalFormatKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemOriginalFormatKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalsourcekey
func (m NSMetadataItem) NSMetadataItemOriginalSourceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemOriginalSourceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempageheightkey
func (m NSMetadataItem) NSMetadataItemPageHeightKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPageHeightKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempagewidthkey
func (m NSMetadataItem) NSMetadataItemPageWidthKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPageWidthKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemparticipantskey
func (m NSMetadataItem) NSMetadataItemParticipantsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemParticipantsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemperformerskey
func (m NSMetadataItem) NSMetadataItemPerformersKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPerformersKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemphonenumberskey
func (m NSMetadataItem) NSMetadataItemPhoneNumbersKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPhoneNumbersKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempixelcountkey
func (m NSMetadataItem) NSMetadataItemPixelCountKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPixelCountKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempixelheightkey
func (m NSMetadataItem) NSMetadataItemPixelHeightKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPixelHeightKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempixelwidthkey
func (m NSMetadataItem) NSMetadataItemPixelWidthKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPixelWidthKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemproducerkey
func (m NSMetadataItem) NSMetadataItemProducerKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemProducerKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemprofilenamekey
func (m NSMetadataItem) NSMetadataItemProfileNameKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemProfileNameKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemprojectskey
func (m NSMetadataItem) NSMetadataItemProjectsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemProjectsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitempublisherskey
func (m NSMetadataItem) NSMetadataItemPublishersKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemPublishersKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientaddresseskey
func (m NSMetadataItem) NSMetadataItemRecipientAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRecipientAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientemailaddresseskey
func (m NSMetadataItem) NSMetadataItemRecipientEmailAddressesKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRecipientEmailAddressesKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientskey
func (m NSMetadataItem) NSMetadataItemRecipientsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRecipientsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingdatekey
func (m NSMetadataItem) NSMetadataItemRecordingDateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRecordingDateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingyearkey
func (m NSMetadataItem) NSMetadataItemRecordingYearKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRecordingYearKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemredeyeonoffkey
func (m NSMetadataItem) NSMetadataItemRedEyeOnOffKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRedEyeOnOffKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionheightdpikey
func (m NSMetadataItem) NSMetadataItemResolutionHeightDPIKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemResolutionHeightDPIKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionwidthdpikey
func (m NSMetadataItem) NSMetadataItemResolutionWidthDPIKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemResolutionWidthDPIKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemrightskey
func (m NSMetadataItem) NSMetadataItemRightsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemRightsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemsecuritymethodkey
func (m NSMetadataItem) NSMetadataItemSecurityMethodKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemSecurityMethodKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemspeedkey
func (m NSMetadataItem) NSMetadataItemSpeedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemSpeedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemstarratingkey
func (m NSMetadataItem) NSMetadataItemStarRatingKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemStarRatingKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemstateorprovincekey
func (m NSMetadataItem) NSMetadataItemStateOrProvinceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemStateOrProvinceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemstreamablekey
func (m NSMetadataItem) NSMetadataItemStreamableKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemStreamableKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemsubjectkey
func (m NSMetadataItem) NSMetadataItemSubjectKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemSubjectKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtempokey
func (m NSMetadataItem) NSMetadataItemTempoKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTempoKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtextcontentkey
func (m NSMetadataItem) NSMetadataItemTextContentKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTextContentKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemthemekey
func (m NSMetadataItem) NSMetadataItemThemeKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemThemeKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtimesignaturekey
func (m NSMetadataItem) NSMetadataItemTimeSignatureKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTimeSignatureKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtimestampkey
func (m NSMetadataItem) NSMetadataItemTimestampKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTimestampKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtitlekey
func (m NSMetadataItem) NSMetadataItemTitleKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTitleKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemtotalbitratekey
func (m NSMetadataItem) NSMetadataItemTotalBitRateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemTotalBitRateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemversionkey
func (m NSMetadataItem) NSMetadataItemVersionKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemVersionKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemvideobitratekey
func (m NSMetadataItem) NSMetadataItemVideoBitRateKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemVideoBitRateKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemwherefromskey
func (m NSMetadataItem) NSMetadataItemWhereFromsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemWhereFromsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemwhitebalancekey
func (m NSMetadataItem) NSMetadataItemWhiteBalanceKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemWhiteBalanceKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataitemisubiquitouskey
func (m NSMetadataItem) NSMetadataItemIsUbiquitousKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataItemIsUbiquitousKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemcontainerdisplaynamekey
func (m NSMetadataItem) NSMetadataUbiquitousItemContainerDisplayNameKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemContainerDisplayNameKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadrequestedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadRequestedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadRequestedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisexternaldocumentkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsExternalDocumentKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsExternalDocumentKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemurlinlocalcontainerkey
func (m NSMetadataItem) NSMetadataUbiquitousItemURLInLocalContainerKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemURLInLocalContainerKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemhasunresolvedconflictskey
func (m NSMetadataItem) NSMetadataUbiquitousItemHasUnresolvedConflictsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemHasUnresolvedConflictsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsDownloadedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadingkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsDownloadingKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadingKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsUploadedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadingkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsUploadingKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadingKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentdownloadedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemPercentDownloadedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemPercentDownloadedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentuploadedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemPercentUploadedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemPercentUploadedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuskey
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadingStatusKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingerrorkey
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadingErrorKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingErrorKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemuploadingerrorkey
func (m NSMetadataItem) NSMetadataUbiquitousItemUploadingErrorKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemUploadingErrorKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemissharedkey
func (m NSMetadataItem) NSMetadataUbiquitousItemIsSharedKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemIsSharedKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserpermissionskey
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserrolekey
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemCurrentUserRoleKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserRoleKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemmostrecenteditornamecomponentskey
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemownernamecomponentskey
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemOwnerNameComponentsKey() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemOwnerNameComponentsKey"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuscurrent
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadingStatusCurrent() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusCurrent"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusdownloaded
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadingStatusDownloaded() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusDownloaded"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusnotdownloaded
func (m NSMetadataItem) NSMetadataUbiquitousItemDownloadingStatusNotDownloaded() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusNotDownloaded"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadonly
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadOnly() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadOnly"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadwrite
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadWrite() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadWrite"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleowner
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemRoleOwner() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleOwner"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleparticipant
func (m NSMetadataItem) NSMetadataUbiquitousSharedItemRoleParticipant() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleParticipant"))
	return NSStringFromID(rv).String()
}

