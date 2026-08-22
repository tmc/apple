// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [CSSearchableItemAttributeSet] class.
var (
	_CSSearchableItemAttributeSetClass     CSSearchableItemAttributeSetClass
	_CSSearchableItemAttributeSetClassOnce sync.Once
)

func getCSSearchableItemAttributeSetClass() CSSearchableItemAttributeSetClass {
	_CSSearchableItemAttributeSetClassOnce.Do(func() {
		_CSSearchableItemAttributeSetClass = CSSearchableItemAttributeSetClass{class: objc.GetClass("CSSearchableItemAttributeSet")}
	})
	return _CSSearchableItemAttributeSetClass
}

// GetCSSearchableItemAttributeSetClass returns the class object for CSSearchableItemAttributeSet.
func GetCSSearchableItemAttributeSetClass() CSSearchableItemAttributeSetClass {
	return getCSSearchableItemAttributeSetClass()
}

type CSSearchableItemAttributeSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSearchableItemAttributeSetClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSearchableItemAttributeSetClass) Alloc() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The detailed metadata for a searchable item.
//
// # Overview
//
// A [CSSearchableItemAttributeSet] contains an extensive set of attributes
// that describe your app’s content. Attributes include information such as
// its title and a brief description. They can also refer to who created the
// item, what kind of data it represents, when someone created it, and more.
// During the indexing process, you create [CSSearchableItem] objects and use
// a [CSSearchableItemAttributeSet] to fill in the attributes for that item.
// During a search, you can query the index for items with attributes that
// match specific values.
//
// When creating a [CSSearchableItem], it’s important to fill out as much
// information in the accompanying [CSSearchableItemAttributeSet] object as
// possible. You don’t have to provide values for every attribute. Instead,
// choose attributes that match the domain of your content. This type divides
// attributes into groups such as media, documents, events, places, music,
// images, and more. You can also add custom attributes to describe new types
// of content. When defining custom attributes, be as specific as possible in
// your definition, and provide a value for the
// [CSSearchableItemAttributeSet.ContentTypeTree] property so your custom
// attribute inherits from a known type.
//
// # Creating an attribute set
//
//   - [CSSearchableItemAttributeSet.InitWithContentType]: Creates an attribute set for the specified content type.
//   - [CSSearchableItemAttributeSet.InitWithCoder]
//
// # Accessing custom attributes
//
//   - [CSSearchableItemAttributeSet.SetValueForCustomKey]: Sets the value for a custom attribute key.
//   - [CSSearchableItemAttributeSet.ValueForCustomKey]: Returns the value associated with the specified custom attribute key.
//
// # Handling Apple Intelligence prioritization and summarization
//
//   - [CSSearchableItemAttributeSet.IsPriority]: A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//   - [CSSearchableItemAttributeSet.TextContentSummary]: A string that presents the Apple Intelligence summarization of the item.
//   - [CSSearchableItemAttributeSet.TranscribedTextContent]: A string that represents the text the system transcribed.
//   - [CSSearchableItemAttributeSet.SetTranscribedTextContent]
//
// # Providing item representations
//
//   - [CSSearchableItemAttributeSet.ProviderDataTypeIdentifiers]: An array of identifiers that corresponds to data representations the delegate provides.
//   - [CSSearchableItemAttributeSet.SetProviderDataTypeIdentifiers]
//   - [CSSearchableItemAttributeSet.ProviderFileTypeIdentifiers]: An array of identifiers that corresponds to file representations the delegate provides.
//   - [CSSearchableItemAttributeSet.SetProviderFileTypeIdentifiers]
//   - [CSSearchableItemAttributeSet.ProviderInPlaceFileTypeIdentifiers]: An array of identifiers that corresponds to in-place file representations the delegate provides.
//   - [CSSearchableItemAttributeSet.SetProviderInPlaceFileTypeIdentifiers]
//
// # Describing general attributes
//
//   - [CSSearchableItemAttributeSet.AlternateNames]: An array of localized strings that represent alternate display names for the item.
//   - [CSSearchableItemAttributeSet.SetAlternateNames]
//   - [CSSearchableItemAttributeSet.ContentType]: The uniform type identifier (UTI) of the item.
//   - [CSSearchableItemAttributeSet.SetContentType]
//   - [CSSearchableItemAttributeSet.ContentTypeTree]: An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//   - [CSSearchableItemAttributeSet.SetContentTypeTree]
//   - [CSSearchableItemAttributeSet.ContentURL]: The file URL of the content to index.
//   - [CSSearchableItemAttributeSet.SetContentURL]
//   - [CSSearchableItemAttributeSet.DarkThumbnailURL]: The local file URL of the thumbnail image for the item when Dark Mode is active.
//   - [CSSearchableItemAttributeSet.SetDarkThumbnailURL]
//   - [CSSearchableItemAttributeSet.DisplayName]: A localized string that contains the name of the item, suitable to display in the user interface.
//   - [CSSearchableItemAttributeSet.SetDisplayName]
//   - [CSSearchableItemAttributeSet.Keywords]: An array of keywords associated with the item, such as work, birthday, important, and so on.
//   - [CSSearchableItemAttributeSet.SetKeywords]
//   - [CSSearchableItemAttributeSet.MetadataModificationDate]: The date on which the last metadata attribute was changed.
//   - [CSSearchableItemAttributeSet.SetMetadataModificationDate]
//   - [CSSearchableItemAttributeSet.Path]: The complete path to the item.
//   - [CSSearchableItemAttributeSet.SetPath]
//   - [CSSearchableItemAttributeSet.RankingHint]: A number that indicates the relative importance of the item among other items from the app.
//   - [CSSearchableItemAttributeSet.SetRankingHint]
//   - [CSSearchableItemAttributeSet.RelatedUniqueIdentifier]: The unique identifier for the item to which the activity is related.
//   - [CSSearchableItemAttributeSet.SetRelatedUniqueIdentifier]
//   - [CSSearchableItemAttributeSet.ThumbnailData]: Image data that represents the thumbnail of the item.
//   - [CSSearchableItemAttributeSet.SetThumbnailData]
//   - [CSSearchableItemAttributeSet.ThumbnailURL]: The local file URL of the thumbnail image for the item.
//   - [CSSearchableItemAttributeSet.SetThumbnailURL]
//   - [CSSearchableItemAttributeSet.Title]: The title of the item.
//   - [CSSearchableItemAttributeSet.SetTitle]
//   - [CSSearchableItemAttributeSet.DomainIdentifier]: An identifier that represents the domain or owner of the item.
//   - [CSSearchableItemAttributeSet.SetDomainIdentifier]
//   - [CSSearchableItemAttributeSet.WeakRelatedUniqueIdentifier]: The unique identifier for the item to which the activity is related, but not linked.
//   - [CSSearchableItemAttributeSet.SetWeakRelatedUniqueIdentifier]
//
// # Describing document content
//
//   - [CSSearchableItemAttributeSet.Audiences]: A class of entity for which the item is intended or useful.
//   - [CSSearchableItemAttributeSet.SetAudiences]
//   - [CSSearchableItemAttributeSet.ContentDescription]: A description of the item’s content.
//   - [CSSearchableItemAttributeSet.SetContentDescription]
//   - [CSSearchableItemAttributeSet.Creator]: The name of the app that created the content.
//   - [CSSearchableItemAttributeSet.SetCreator]
//   - [CSSearchableItemAttributeSet.EncodingApplications]: The name of the apps that converted the original content into a PDF stream.
//   - [CSSearchableItemAttributeSet.SetEncodingApplications]
//   - [CSSearchableItemAttributeSet.FileSize]: The size of the document file.
//   - [CSSearchableItemAttributeSet.SetFileSize]
//   - [CSSearchableItemAttributeSet.FontNames]: An array of font names the document uses.
//   - [CSSearchableItemAttributeSet.SetFontNames]
//   - [CSSearchableItemAttributeSet.Identifier]: A formal identifier that references the document the item represents.
//   - [CSSearchableItemAttributeSet.SetIdentifier]
//   - [CSSearchableItemAttributeSet.Kind]: A description of the kind of document the item represents.
//   - [CSSearchableItemAttributeSet.SetKind]
//   - [CSSearchableItemAttributeSet.PageCount]: The number of pages in the document.
//   - [CSSearchableItemAttributeSet.SetPageCount]
//   - [CSSearchableItemAttributeSet.PageHeight]: The height of the document page, in points (72 points per inch).
//   - [CSSearchableItemAttributeSet.SetPageHeight]
//   - [CSSearchableItemAttributeSet.PageWidth]: The width of the document page, in points (72 points per inch).
//   - [CSSearchableItemAttributeSet.SetPageWidth]
//   - [CSSearchableItemAttributeSet.SecurityMethod]: The security method (a type of encryption) that protects the document file.
//   - [CSSearchableItemAttributeSet.SetSecurityMethod]
//   - [CSSearchableItemAttributeSet.Subject]: The subject of the document.
//   - [CSSearchableItemAttributeSet.SetSubject]
//   - [CSSearchableItemAttributeSet.Theme]: The theme of the document.
//   - [CSSearchableItemAttributeSet.SetTheme]
//
// # Describing user involvement
//
//   - [CSSearchableItemAttributeSet.UserCreated]: A value that indicates the user created the item.
//   - [CSSearchableItemAttributeSet.SetUserCreated]
//   - [CSSearchableItemAttributeSet.UserCurated]: A value that indicates the user selected the item.
//   - [CSSearchableItemAttributeSet.SetUserCurated]
//   - [CSSearchableItemAttributeSet.UserOwned]: A value that indicates the user purchased or owns the item.
//   - [CSSearchableItemAttributeSet.SetUserOwned]
//
// # Describing events
//
//   - [CSSearchableItemAttributeSet.AllDay]: A value that indicates if the event covers an entire day.
//   - [CSSearchableItemAttributeSet.SetAllDay]
//   - [CSSearchableItemAttributeSet.CompletionDate]: The date on which the item was completed.
//   - [CSSearchableItemAttributeSet.SetCompletionDate]
//   - [CSSearchableItemAttributeSet.DueDate]: The date on which the item is due.
//   - [CSSearchableItemAttributeSet.SetDueDate]
//   - [CSSearchableItemAttributeSet.EndDate]: The end date for the item.
//   - [CSSearchableItemAttributeSet.SetEndDate]
//   - [CSSearchableItemAttributeSet.ImportantDates]: An array of important dates associated with the item.
//   - [CSSearchableItemAttributeSet.SetImportantDates]
//   - [CSSearchableItemAttributeSet.StartDate]: The start date for the item.
//   - [CSSearchableItemAttributeSet.SetStartDate]
//
// # Describing places
//
//   - [CSSearchableItemAttributeSet.Altitude]: The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//   - [CSSearchableItemAttributeSet.SetAltitude]
//   - [CSSearchableItemAttributeSet.City]: The city of the item’s origin according to guidelines that the provider establishes.
//   - [CSSearchableItemAttributeSet.SetCity]
//   - [CSSearchableItemAttributeSet.Country]: The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//   - [CSSearchableItemAttributeSet.SetCountry]
//   - [CSSearchableItemAttributeSet.GPSAreaInformation]: Information about the GPS area.
//   - [CSSearchableItemAttributeSet.SetGPSAreaInformation]
//   - [CSSearchableItemAttributeSet.GPSDOP]: The GPS dilution of precision value.
//   - [CSSearchableItemAttributeSet.SetGPSDOP]
//   - [CSSearchableItemAttributeSet.GPSDateStamp]: The date and time related to the GPS value.
//   - [CSSearchableItemAttributeSet.SetGPSDateStamp]
//   - [CSSearchableItemAttributeSet.GPSDestBearing]: The bearing to the destination point.
//   - [CSSearchableItemAttributeSet.SetGPSDestBearing]
//   - [CSSearchableItemAttributeSet.GPSDestDistance]: The distance to the destination point.
//   - [CSSearchableItemAttributeSet.SetGPSDestDistance]
//   - [CSSearchableItemAttributeSet.GPSDestLatitude]: The latitude of the destination point.
//   - [CSSearchableItemAttributeSet.SetGPSDestLatitude]
//   - [CSSearchableItemAttributeSet.GPSDestLongitude]: The longitude of the destination point.
//   - [CSSearchableItemAttributeSet.SetGPSDestLongitude]
//   - [CSSearchableItemAttributeSet.GPSDifferental]: The differential correction applied to the GPS receiver.
//   - [CSSearchableItemAttributeSet.SetGPSDifferental]
//   - [CSSearchableItemAttributeSet.GPSMapDatum]: The geodetic data that the GPS receiver uses.
//   - [CSSearchableItemAttributeSet.SetGPSMapDatum]
//   - [CSSearchableItemAttributeSet.GPSMeasureMode]: The measurement precision mode in use by the GPS receiver.
//   - [CSSearchableItemAttributeSet.SetGPSMeasureMode]
//   - [CSSearchableItemAttributeSet.GPSProcessingMethod]: The location finding method that the GPS receiver uses.
//   - [CSSearchableItemAttributeSet.SetGPSProcessingMethod]
//   - [CSSearchableItemAttributeSet.GPSStatus]: The status of the GPS receiver.
//   - [CSSearchableItemAttributeSet.SetGPSStatus]
//   - [CSSearchableItemAttributeSet.GPSTrack]: The direction of travel of the item in degrees from true north.
//   - [CSSearchableItemAttributeSet.SetGPSTrack]
//   - [CSSearchableItemAttributeSet.Headline]: A publishable string that provides a synopsis of the contents of the item.
//   - [CSSearchableItemAttributeSet.SetHeadline]
//   - [CSSearchableItemAttributeSet.ImageDirection]: The direction of the item’s image in degrees from true north.
//   - [CSSearchableItemAttributeSet.SetImageDirection]
//   - [CSSearchableItemAttributeSet.Instructions]: Instructions that concern the use of the item, such as an embargo or warning.
//   - [CSSearchableItemAttributeSet.SetInstructions]
//   - [CSSearchableItemAttributeSet.Latitude]: The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//   - [CSSearchableItemAttributeSet.SetLatitude]
//   - [CSSearchableItemAttributeSet.Longitude]: The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//   - [CSSearchableItemAttributeSet.SetLongitude]
//   - [CSSearchableItemAttributeSet.NamedLocation]: The name of the location or point of interest associated with the item.
//   - [CSSearchableItemAttributeSet.SetNamedLocation]
//   - [CSSearchableItemAttributeSet.Speed]: The speed of the item, in kilometers per hour.
//   - [CSSearchableItemAttributeSet.SetSpeed]
//   - [CSSearchableItemAttributeSet.StateOrProvince]: The province or state of origin according to guidelines the provider establishes.
//   - [CSSearchableItemAttributeSet.SetStateOrProvince]
//   - [CSSearchableItemAttributeSet.Timestamp]: The timestamp on the item.
//   - [CSSearchableItemAttributeSet.SetTimestamp]
//   - [CSSearchableItemAttributeSet.FullyFormattedAddress]: The fully formatted address of the item, received from MapKit.
//   - [CSSearchableItemAttributeSet.SetFullyFormattedAddress]
//   - [CSSearchableItemAttributeSet.PostalCode]: The postal code for the item according to guidelines the provider establishes.
//   - [CSSearchableItemAttributeSet.SetPostalCode]
//   - [CSSearchableItemAttributeSet.SubThoroughfare]: The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//   - [CSSearchableItemAttributeSet.SetSubThoroughfare]
//   - [CSSearchableItemAttributeSet.Thoroughfare]: The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//   - [CSSearchableItemAttributeSet.SetThoroughfare]
//
// # Describing media
//
//   - [CSSearchableItemAttributeSet.Comment]: A comment related to the media file.
//   - [CSSearchableItemAttributeSet.SetComment]
//   - [CSSearchableItemAttributeSet.ContentCreationDate]: The creation date of an edited or optimized version of the song or composition.
//   - [CSSearchableItemAttributeSet.SetContentCreationDate]
//   - [CSSearchableItemAttributeSet.ContentModificationDate]: The date on which the contents of the file was last modified.
//   - [CSSearchableItemAttributeSet.SetContentModificationDate]
//   - [CSSearchableItemAttributeSet.ContentSources]: An array of sources from which the media was obtained.
//   - [CSSearchableItemAttributeSet.SetContentSources]
//   - [CSSearchableItemAttributeSet.Copyright]: The copyright date of the content.
//   - [CSSearchableItemAttributeSet.SetCopyright]
//   - [CSSearchableItemAttributeSet.DownloadedDate]: The most recent date on which the file was downloaded or received.
//   - [CSSearchableItemAttributeSet.SetDownloadedDate]
//   - [CSSearchableItemAttributeSet.Editors]: A list of editors who have worked on the file.
//   - [CSSearchableItemAttributeSet.SetEditors]
//   - [CSSearchableItemAttributeSet.LastUsedDate]: The date on which the file was last used.
//   - [CSSearchableItemAttributeSet.SetLastUsedDate]
//   - [CSSearchableItemAttributeSet.Participants]: A list of people who are visible in an image or movie or written about in a document.
//   - [CSSearchableItemAttributeSet.SetParticipants]
//   - [CSSearchableItemAttributeSet.Projects]: A list of projects of which this file is a part.
//   - [CSSearchableItemAttributeSet.SetProjects]
//   - [CSSearchableItemAttributeSet.AddedDate]: The date on which the item was moved into its current location.
//   - [CSSearchableItemAttributeSet.SetAddedDate]
//   - [CSSearchableItemAttributeSet.Codecs]: The codecs used to encode/decode the media.
//   - [CSSearchableItemAttributeSet.SetCodecs]
//   - [CSSearchableItemAttributeSet.ContactKeywords]: A list of contacts who are associated with the content in some way, not including the author.
//   - [CSSearchableItemAttributeSet.SetContactKeywords]
//   - [CSSearchableItemAttributeSet.DeliveryType]: The delivery type of the file.
//   - [CSSearchableItemAttributeSet.SetDeliveryType]
//   - [CSSearchableItemAttributeSet.Duration]: The duration (if appropriate) of the content of the file, in seconds.
//   - [CSSearchableItemAttributeSet.SetDuration]
//   - [CSSearchableItemAttributeSet.MediaTypes]: The media types present in the content.
//   - [CSSearchableItemAttributeSet.SetMediaTypes]
//   - [CSSearchableItemAttributeSet.Organizations]: A list of companies or organizations that created the content.
//   - [CSSearchableItemAttributeSet.SetOrganizations]
//   - [CSSearchableItemAttributeSet.Streamable]: A value that indicates if the content is prepared for streaming.
//   - [CSSearchableItemAttributeSet.SetStreamable]
//   - [CSSearchableItemAttributeSet.TotalBitRate]: The total bit rate of the media, combining audio and video.
//   - [CSSearchableItemAttributeSet.SetTotalBitRate]
//   - [CSSearchableItemAttributeSet.AudioBitRate]: The audio bit rate of the media.
//   - [CSSearchableItemAttributeSet.SetAudioBitRate]
//   - [CSSearchableItemAttributeSet.Version]: A version string associated with the file.
//   - [CSSearchableItemAttributeSet.SetVersion]
//   - [CSSearchableItemAttributeSet.VideoBitRate]: The video bit rate of the media.
//   - [CSSearchableItemAttributeSet.SetVideoBitRate]
//   - [CSSearchableItemAttributeSet.Contributors]: A list of people, organizations, or services that made contributions to the media content.
//   - [CSSearchableItemAttributeSet.SetContributors]
//   - [CSSearchableItemAttributeSet.Languages]: A list of the included languages for the intellectual content of the media.
//   - [CSSearchableItemAttributeSet.SetLanguages]
//   - [CSSearchableItemAttributeSet.Publishers]: A list of people, organizations, services, or other entities responsible for making the media available.
//   - [CSSearchableItemAttributeSet.SetPublishers]
//   - [CSSearchableItemAttributeSet.Rights]: A link to information about the rights held in and over the media.
//   - [CSSearchableItemAttributeSet.SetRights]
//   - [CSSearchableItemAttributeSet.Role]: Indicates the role of the content creator.
//   - [CSSearchableItemAttributeSet.SetRole]
//   - [CSSearchableItemAttributeSet.ContentRating]: A value that indicates if the media contains explicit content.
//   - [CSSearchableItemAttributeSet.SetContentRating]
//   - [CSSearchableItemAttributeSet.Coverage]: A list of descriptors that specify the extent or scope of the media.
//   - [CSSearchableItemAttributeSet.SetCoverage]
//   - [CSSearchableItemAttributeSet.Director]: The name of the director of the media (for example, a movie director).
//   - [CSSearchableItemAttributeSet.SetDirector]
//   - [CSSearchableItemAttributeSet.Genre]: The genre of the media.
//   - [CSSearchableItemAttributeSet.SetGenre]
//   - [CSSearchableItemAttributeSet.Information]: Information about the media.
//   - [CSSearchableItemAttributeSet.SetInformation]
//   - [CSSearchableItemAttributeSet.Local]: A value that indicates if the media is local.
//   - [CSSearchableItemAttributeSet.SetLocal]
//   - [CSSearchableItemAttributeSet.OriginalFormat]: The original format of the media.
//   - [CSSearchableItemAttributeSet.SetOriginalFormat]
//   - [CSSearchableItemAttributeSet.OriginalSource]: The original source of the media.
//   - [CSSearchableItemAttributeSet.SetOriginalSource]
//   - [CSSearchableItemAttributeSet.Performers]: A list of performers in the media.
//   - [CSSearchableItemAttributeSet.SetPerformers]
//   - [CSSearchableItemAttributeSet.PlayCount]: A user-supplied play count for the media.
//   - [CSSearchableItemAttributeSet.SetPlayCount]
//   - [CSSearchableItemAttributeSet.Producer]: The producer of the content.
//   - [CSSearchableItemAttributeSet.SetProducer]
//   - [CSSearchableItemAttributeSet.Rating]: The user-supplied rating of the media.
//   - [CSSearchableItemAttributeSet.SetRating]
//   - [CSSearchableItemAttributeSet.RatingDescription]: A description of the rating.
//   - [CSSearchableItemAttributeSet.SetRatingDescription]
//   - [CSSearchableItemAttributeSet.URL]: The URL associated with the media.
//   - [CSSearchableItemAttributeSet.SetURL]
//
// # Describing music
//
//   - [CSSearchableItemAttributeSet.Album]: The title for a collection of audio media.
//   - [CSSearchableItemAttributeSet.SetAlbum]
//   - [CSSearchableItemAttributeSet.Artist]: The artist associated with the media.
//   - [CSSearchableItemAttributeSet.SetArtist]
//   - [CSSearchableItemAttributeSet.AudioChannelCount]: The number of channels in the audio data that the file contains.
//   - [CSSearchableItemAttributeSet.SetAudioChannelCount]
//   - [CSSearchableItemAttributeSet.AudioEncodingApplication]: The name of the application that encoded the data the audio file contains.
//   - [CSSearchableItemAttributeSet.SetAudioEncodingApplication]
//   - [CSSearchableItemAttributeSet.AudioSampleRate]: The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//   - [CSSearchableItemAttributeSet.SetAudioSampleRate]
//   - [CSSearchableItemAttributeSet.AudioTrackNumber]: The track number of a song or audio composition when part of an album.
//   - [CSSearchableItemAttributeSet.SetAudioTrackNumber]
//   - [CSSearchableItemAttributeSet.Composer]: The composer of the song or audio composition that the audio file contains.
//   - [CSSearchableItemAttributeSet.SetComposer]
//   - [CSSearchableItemAttributeSet.KeySignature]: The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//   - [CSSearchableItemAttributeSet.SetKeySignature]
//   - [CSSearchableItemAttributeSet.Lyricist]: The lyricist or text writer for the song or audio composition that the file contains.
//   - [CSSearchableItemAttributeSet.SetLyricist]
//   - [CSSearchableItemAttributeSet.MusicalGenre]: The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//   - [CSSearchableItemAttributeSet.SetMusicalGenre]
//   - [CSSearchableItemAttributeSet.RecordingDate]: The recording date of the song or composition.
//   - [CSSearchableItemAttributeSet.SetRecordingDate]
//   - [CSSearchableItemAttributeSet.Tempo]: The tempo of the music that the audio file contains, in beats per minute.
//   - [CSSearchableItemAttributeSet.SetTempo]
//   - [CSSearchableItemAttributeSet.TimeSignature]: The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//   - [CSSearchableItemAttributeSet.SetTimeSignature]
//   - [CSSearchableItemAttributeSet.GeneralMIDISequence]: A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//   - [CSSearchableItemAttributeSet.SetGeneralMIDISequence]
//   - [CSSearchableItemAttributeSet.MusicalInstrumentCategory]: The category of the instrument associated with the audio file.
//   - [CSSearchableItemAttributeSet.SetMusicalInstrumentCategory]
//   - [CSSearchableItemAttributeSet.MusicalInstrumentName]: The name of an instrument within the context of an instrument category.
//   - [CSSearchableItemAttributeSet.SetMusicalInstrumentName]
//
// # Describing images
//
//   - [CSSearchableItemAttributeSet.ISOSpeed]: The ISO speed setting at the time the camera captured the image.
//   - [CSSearchableItemAttributeSet.SetISOSpeed]
//   - [CSSearchableItemAttributeSet.AcquisitionMake]: The manufacturer of the device that captured the image.
//   - [CSSearchableItemAttributeSet.SetAcquisitionMake]
//   - [CSSearchableItemAttributeSet.AcquisitionModel]: The model of the device that captured the image.
//   - [CSSearchableItemAttributeSet.SetAcquisitionModel]
//   - [CSSearchableItemAttributeSet.Aperture]: The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//   - [CSSearchableItemAttributeSet.SetAperture]
//   - [CSSearchableItemAttributeSet.BitsPerSample]: The number of bits per sample.
//   - [CSSearchableItemAttributeSet.SetBitsPerSample]
//   - [CSSearchableItemAttributeSet.CameraOwner]: The owner of the camera that captured the image.
//   - [CSSearchableItemAttributeSet.SetCameraOwner]
//   - [CSSearchableItemAttributeSet.ColorSpace]: The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//   - [CSSearchableItemAttributeSet.SetColorSpace]
//   - [CSSearchableItemAttributeSet.FlashOn]: A value that indicates if the camera used a flash to capture the image.
//   - [CSSearchableItemAttributeSet.SetFlashOn]
//   - [CSSearchableItemAttributeSet.FocalLength]: The actual focal length of the lens, in millimeters.
//   - [CSSearchableItemAttributeSet.SetFocalLength]
//   - [CSSearchableItemAttributeSet.FocalLength35mm]: A value that indicates if the focal length is 35mm.
//   - [CSSearchableItemAttributeSet.SetFocalLength35mm]
//   - [CSSearchableItemAttributeSet.LayerNames]: An array that contains the names of the various layers in the file.
//   - [CSSearchableItemAttributeSet.SetLayerNames]
//   - [CSSearchableItemAttributeSet.LensModel]: The model of the lens that captured the image.
//   - [CSSearchableItemAttributeSet.SetLensModel]
//   - [CSSearchableItemAttributeSet.Orientation]: The orientation of the data.
//   - [CSSearchableItemAttributeSet.SetOrientation]
//   - [CSSearchableItemAttributeSet.PixelCount]: The total number of pixels in the image.
//   - [CSSearchableItemAttributeSet.SetPixelCount]
//   - [CSSearchableItemAttributeSet.PixelHeight]: The height of the item, such as image or video frame height, in pixels.
//   - [CSSearchableItemAttributeSet.SetPixelHeight]
//   - [CSSearchableItemAttributeSet.PixelWidth]: The width of the item, such as image or video frame width, in pixels.
//   - [CSSearchableItemAttributeSet.SetPixelWidth]
//   - [CSSearchableItemAttributeSet.WhiteBalance]: The white balance setting when the camera captured the image.
//   - [CSSearchableItemAttributeSet.SetWhiteBalance]
//   - [CSSearchableItemAttributeSet.EXIFGPSVersion]: The version of GPS Info IFD header that was used to generate the metadata for the image.
//   - [CSSearchableItemAttributeSet.SetEXIFGPSVersion]
//   - [CSSearchableItemAttributeSet.EXIFVersion]: The version of the EXIF header that was used to generate the metadata for the image.
//   - [CSSearchableItemAttributeSet.SetEXIFVersion]
//   - [CSSearchableItemAttributeSet.ExposureMode]: The mode the camera used for the exposure of the image.
//   - [CSSearchableItemAttributeSet.SetExposureMode]
//   - [CSSearchableItemAttributeSet.ExposureProgram]: The class of the program the camera used to set exposure when capturing the image.
//   - [CSSearchableItemAttributeSet.SetExposureProgram]
//   - [CSSearchableItemAttributeSet.ExposureTime]: The time that the lens was open during exposure, in seconds.
//   - [CSSearchableItemAttributeSet.SetExposureTime]
//   - [CSSearchableItemAttributeSet.ExposureTimeString]: The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//   - [CSSearchableItemAttributeSet.SetExposureTimeString]
//   - [CSSearchableItemAttributeSet.FNumber]: The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//   - [CSSearchableItemAttributeSet.SetFNumber]
//   - [CSSearchableItemAttributeSet.HasAlphaChannel]: Indicates if the image file has an alpha channel.
//   - [CSSearchableItemAttributeSet.SetHasAlphaChannel]
//   - [CSSearchableItemAttributeSet.MaxAperture]: The smallest F number of the lens.
//   - [CSSearchableItemAttributeSet.SetMaxAperture]
//   - [CSSearchableItemAttributeSet.MeteringMode]: The metering mode.
//   - [CSSearchableItemAttributeSet.SetMeteringMode]
//   - [CSSearchableItemAttributeSet.ProfileName]: The name of the color profile the camera used for the image.
//   - [CSSearchableItemAttributeSet.SetProfileName]
//   - [CSSearchableItemAttributeSet.RedEyeOn]: A value that indicates if the camera used red-eye reduction when capturing the image.
//   - [CSSearchableItemAttributeSet.SetRedEyeOn]
//   - [CSSearchableItemAttributeSet.ResolutionHeightDPI]: The resolution height of the image, in DPI.
//   - [CSSearchableItemAttributeSet.SetResolutionHeightDPI]
//   - [CSSearchableItemAttributeSet.ResolutionWidthDPI]: The resolution width of the image, in DPI.
//   - [CSSearchableItemAttributeSet.SetResolutionWidthDPI]
//
// # Describing messages
//
//   - [CSSearchableItemAttributeSet.HTMLContentData]: The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//   - [CSSearchableItemAttributeSet.SetHTMLContentData]
//   - [CSSearchableItemAttributeSet.AccountHandles]: An array of the canonical handles for the account with which the message is associated.
//   - [CSSearchableItemAttributeSet.SetAccountHandles]
//   - [CSSearchableItemAttributeSet.AccountIdentifier]: The unique identifier for the account with which the message is associated, if any.
//   - [CSSearchableItemAttributeSet.SetAccountIdentifier]
//   - [CSSearchableItemAttributeSet.AdditionalRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Cc: field in an email message.
//   - [CSSearchableItemAttributeSet.SetAdditionalRecipients]
//   - [CSSearchableItemAttributeSet.AuthorAddresses]: An array of addresses associated with the author of the message.
//   - [CSSearchableItemAttributeSet.SetAuthorAddresses]
//   - [CSSearchableItemAttributeSet.AuthorEmailAddresses]: An array of email addresses associated with the author of the message.
//   - [CSSearchableItemAttributeSet.SetAuthorEmailAddresses]
//   - [CSSearchableItemAttributeSet.AuthorNames]: An array of names representing the authors who have worked on the message.
//   - [CSSearchableItemAttributeSet.SetAuthorNames]
//   - [CSSearchableItemAttributeSet.Authors]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the From: field in an item.
//   - [CSSearchableItemAttributeSet.SetAuthors]
//   - [CSSearchableItemAttributeSet.EmailAddresses]: An array of email addresses associated with the message.
//   - [CSSearchableItemAttributeSet.SetEmailAddresses]
//   - [CSSearchableItemAttributeSet.EmailHeaders]: A dictionary that contains all the headers of the message.
//   - [CSSearchableItemAttributeSet.SetEmailHeaders]
//   - [CSSearchableItemAttributeSet.HiddenAdditionalRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Bcc: field in an email message.
//   - [CSSearchableItemAttributeSet.SetHiddenAdditionalRecipients]
//   - [CSSearchableItemAttributeSet.InstantMessageAddresses]: An array of instant message addresses for the message.
//   - [CSSearchableItemAttributeSet.SetInstantMessageAddresses]
//   - [CSSearchableItemAttributeSet.LikelyJunk]: A value that indicates if the message is likely to be considered junk.
//   - [CSSearchableItemAttributeSet.SetLikelyJunk]
//   - [CSSearchableItemAttributeSet.MailboxIdentifiers]: An array of mailbox identifiers associated with the message.
//   - [CSSearchableItemAttributeSet.SetMailboxIdentifiers]
//   - [CSSearchableItemAttributeSet.PhoneNumbers]: An array of phone numbers associated with the message.
//   - [CSSearchableItemAttributeSet.SetPhoneNumbers]
//   - [CSSearchableItemAttributeSet.PrimaryRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the To: field in an email message.
//   - [CSSearchableItemAttributeSet.SetPrimaryRecipients]
//   - [CSSearchableItemAttributeSet.RecipientAddresses]: An array of addresses associated with the recipients of the message.
//   - [CSSearchableItemAttributeSet.SetRecipientAddresses]
//   - [CSSearchableItemAttributeSet.RecipientEmailAddresses]: An array of email addresses associated with the recipient.
//   - [CSSearchableItemAttributeSet.SetRecipientEmailAddresses]
//   - [CSSearchableItemAttributeSet.RecipientNames]: An array of names representing the recipients of this message.
//   - [CSSearchableItemAttributeSet.SetRecipientNames]
//   - [CSSearchableItemAttributeSet.TextContent]: The textual content of the message.
//   - [CSSearchableItemAttributeSet.SetTextContent]
//
// # Describing containment
//
//   - [CSSearchableItemAttributeSet.ContainerDisplayName]: A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//   - [CSSearchableItemAttributeSet.SetContainerDisplayName]
//   - [CSSearchableItemAttributeSet.ContainerIdentifier]: The identifier of the container to which the item belongs.
//   - [CSSearchableItemAttributeSet.SetContainerIdentifier]
//   - [CSSearchableItemAttributeSet.ContainerOrder]: The order of the item within the container.
//   - [CSSearchableItemAttributeSet.SetContainerOrder]
//   - [CSSearchableItemAttributeSet.ContainerTitle]: The title of the container to which the item belongs.
//   - [CSSearchableItemAttributeSet.SetContainerTitle]
//
// # Describing supporting actions
//
//   - [CSSearchableItemAttributeSet.SupportsNavigation]: A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//   - [CSSearchableItemAttributeSet.SetSupportsNavigation]
//   - [CSSearchableItemAttributeSet.SupportsPhoneCall]: A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//   - [CSSearchableItemAttributeSet.SetSupportsPhoneCall]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet
type CSSearchableItemAttributeSet struct {
	objectivec.Object
}

// CSSearchableItemAttributeSetFromID constructs a [CSSearchableItemAttributeSet] from an objc.ID.
//
// The detailed metadata for a searchable item.
func CSSearchableItemAttributeSetFromID(id objc.ID) CSSearchableItemAttributeSet {
	return CSSearchableItemAttributeSet{objectivec.Object{ID: id}}
}

// NOTE: CSSearchableItemAttributeSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSearchableItemAttributeSet] class.
//
// # Creating an attribute set
//
//   - [ICSSearchableItemAttributeSet.InitWithContentType]: Creates an attribute set for the specified content type.
//   - [ICSSearchableItemAttributeSet.InitWithCoder]
//
// # Accessing custom attributes
//
//   - [ICSSearchableItemAttributeSet.SetValueForCustomKey]: Sets the value for a custom attribute key.
//   - [ICSSearchableItemAttributeSet.ValueForCustomKey]: Returns the value associated with the specified custom attribute key.
//
// # Handling Apple Intelligence prioritization and summarization
//
//   - [ICSSearchableItemAttributeSet.IsPriority]: A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//   - [ICSSearchableItemAttributeSet.TextContentSummary]: A string that presents the Apple Intelligence summarization of the item.
//   - [ICSSearchableItemAttributeSet.TranscribedTextContent]: A string that represents the text the system transcribed.
//   - [ICSSearchableItemAttributeSet.SetTranscribedTextContent]
//
// # Providing item representations
//
//   - [ICSSearchableItemAttributeSet.ProviderDataTypeIdentifiers]: An array of identifiers that corresponds to data representations the delegate provides.
//   - [ICSSearchableItemAttributeSet.SetProviderDataTypeIdentifiers]
//   - [ICSSearchableItemAttributeSet.ProviderFileTypeIdentifiers]: An array of identifiers that corresponds to file representations the delegate provides.
//   - [ICSSearchableItemAttributeSet.SetProviderFileTypeIdentifiers]
//   - [ICSSearchableItemAttributeSet.ProviderInPlaceFileTypeIdentifiers]: An array of identifiers that corresponds to in-place file representations the delegate provides.
//   - [ICSSearchableItemAttributeSet.SetProviderInPlaceFileTypeIdentifiers]
//
// # Describing general attributes
//
//   - [ICSSearchableItemAttributeSet.AlternateNames]: An array of localized strings that represent alternate display names for the item.
//   - [ICSSearchableItemAttributeSet.SetAlternateNames]
//   - [ICSSearchableItemAttributeSet.ContentType]: The uniform type identifier (UTI) of the item.
//   - [ICSSearchableItemAttributeSet.SetContentType]
//   - [ICSSearchableItemAttributeSet.ContentTypeTree]: An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//   - [ICSSearchableItemAttributeSet.SetContentTypeTree]
//   - [ICSSearchableItemAttributeSet.ContentURL]: The file URL of the content to index.
//   - [ICSSearchableItemAttributeSet.SetContentURL]
//   - [ICSSearchableItemAttributeSet.DarkThumbnailURL]: The local file URL of the thumbnail image for the item when Dark Mode is active.
//   - [ICSSearchableItemAttributeSet.SetDarkThumbnailURL]
//   - [ICSSearchableItemAttributeSet.DisplayName]: A localized string that contains the name of the item, suitable to display in the user interface.
//   - [ICSSearchableItemAttributeSet.SetDisplayName]
//   - [ICSSearchableItemAttributeSet.Keywords]: An array of keywords associated with the item, such as work, birthday, important, and so on.
//   - [ICSSearchableItemAttributeSet.SetKeywords]
//   - [ICSSearchableItemAttributeSet.MetadataModificationDate]: The date on which the last metadata attribute was changed.
//   - [ICSSearchableItemAttributeSet.SetMetadataModificationDate]
//   - [ICSSearchableItemAttributeSet.Path]: The complete path to the item.
//   - [ICSSearchableItemAttributeSet.SetPath]
//   - [ICSSearchableItemAttributeSet.RankingHint]: A number that indicates the relative importance of the item among other items from the app.
//   - [ICSSearchableItemAttributeSet.SetRankingHint]
//   - [ICSSearchableItemAttributeSet.RelatedUniqueIdentifier]: The unique identifier for the item to which the activity is related.
//   - [ICSSearchableItemAttributeSet.SetRelatedUniqueIdentifier]
//   - [ICSSearchableItemAttributeSet.ThumbnailData]: Image data that represents the thumbnail of the item.
//   - [ICSSearchableItemAttributeSet.SetThumbnailData]
//   - [ICSSearchableItemAttributeSet.ThumbnailURL]: The local file URL of the thumbnail image for the item.
//   - [ICSSearchableItemAttributeSet.SetThumbnailURL]
//   - [ICSSearchableItemAttributeSet.Title]: The title of the item.
//   - [ICSSearchableItemAttributeSet.SetTitle]
//   - [ICSSearchableItemAttributeSet.DomainIdentifier]: An identifier that represents the domain or owner of the item.
//   - [ICSSearchableItemAttributeSet.SetDomainIdentifier]
//   - [ICSSearchableItemAttributeSet.WeakRelatedUniqueIdentifier]: The unique identifier for the item to which the activity is related, but not linked.
//   - [ICSSearchableItemAttributeSet.SetWeakRelatedUniqueIdentifier]
//
// # Describing document content
//
//   - [ICSSearchableItemAttributeSet.Audiences]: A class of entity for which the item is intended or useful.
//   - [ICSSearchableItemAttributeSet.SetAudiences]
//   - [ICSSearchableItemAttributeSet.ContentDescription]: A description of the item’s content.
//   - [ICSSearchableItemAttributeSet.SetContentDescription]
//   - [ICSSearchableItemAttributeSet.Creator]: The name of the app that created the content.
//   - [ICSSearchableItemAttributeSet.SetCreator]
//   - [ICSSearchableItemAttributeSet.EncodingApplications]: The name of the apps that converted the original content into a PDF stream.
//   - [ICSSearchableItemAttributeSet.SetEncodingApplications]
//   - [ICSSearchableItemAttributeSet.FileSize]: The size of the document file.
//   - [ICSSearchableItemAttributeSet.SetFileSize]
//   - [ICSSearchableItemAttributeSet.FontNames]: An array of font names the document uses.
//   - [ICSSearchableItemAttributeSet.SetFontNames]
//   - [ICSSearchableItemAttributeSet.Identifier]: A formal identifier that references the document the item represents.
//   - [ICSSearchableItemAttributeSet.SetIdentifier]
//   - [ICSSearchableItemAttributeSet.Kind]: A description of the kind of document the item represents.
//   - [ICSSearchableItemAttributeSet.SetKind]
//   - [ICSSearchableItemAttributeSet.PageCount]: The number of pages in the document.
//   - [ICSSearchableItemAttributeSet.SetPageCount]
//   - [ICSSearchableItemAttributeSet.PageHeight]: The height of the document page, in points (72 points per inch).
//   - [ICSSearchableItemAttributeSet.SetPageHeight]
//   - [ICSSearchableItemAttributeSet.PageWidth]: The width of the document page, in points (72 points per inch).
//   - [ICSSearchableItemAttributeSet.SetPageWidth]
//   - [ICSSearchableItemAttributeSet.SecurityMethod]: The security method (a type of encryption) that protects the document file.
//   - [ICSSearchableItemAttributeSet.SetSecurityMethod]
//   - [ICSSearchableItemAttributeSet.Subject]: The subject of the document.
//   - [ICSSearchableItemAttributeSet.SetSubject]
//   - [ICSSearchableItemAttributeSet.Theme]: The theme of the document.
//   - [ICSSearchableItemAttributeSet.SetTheme]
//
// # Describing user involvement
//
//   - [ICSSearchableItemAttributeSet.UserCreated]: A value that indicates the user created the item.
//   - [ICSSearchableItemAttributeSet.SetUserCreated]
//   - [ICSSearchableItemAttributeSet.UserCurated]: A value that indicates the user selected the item.
//   - [ICSSearchableItemAttributeSet.SetUserCurated]
//   - [ICSSearchableItemAttributeSet.UserOwned]: A value that indicates the user purchased or owns the item.
//   - [ICSSearchableItemAttributeSet.SetUserOwned]
//
// # Describing events
//
//   - [ICSSearchableItemAttributeSet.AllDay]: A value that indicates if the event covers an entire day.
//   - [ICSSearchableItemAttributeSet.SetAllDay]
//   - [ICSSearchableItemAttributeSet.CompletionDate]: The date on which the item was completed.
//   - [ICSSearchableItemAttributeSet.SetCompletionDate]
//   - [ICSSearchableItemAttributeSet.DueDate]: The date on which the item is due.
//   - [ICSSearchableItemAttributeSet.SetDueDate]
//   - [ICSSearchableItemAttributeSet.EndDate]: The end date for the item.
//   - [ICSSearchableItemAttributeSet.SetEndDate]
//   - [ICSSearchableItemAttributeSet.ImportantDates]: An array of important dates associated with the item.
//   - [ICSSearchableItemAttributeSet.SetImportantDates]
//   - [ICSSearchableItemAttributeSet.StartDate]: The start date for the item.
//   - [ICSSearchableItemAttributeSet.SetStartDate]
//
// # Describing places
//
//   - [ICSSearchableItemAttributeSet.Altitude]: The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//   - [ICSSearchableItemAttributeSet.SetAltitude]
//   - [ICSSearchableItemAttributeSet.City]: The city of the item’s origin according to guidelines that the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetCity]
//   - [ICSSearchableItemAttributeSet.Country]: The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetCountry]
//   - [ICSSearchableItemAttributeSet.GPSAreaInformation]: Information about the GPS area.
//   - [ICSSearchableItemAttributeSet.SetGPSAreaInformation]
//   - [ICSSearchableItemAttributeSet.GPSDOP]: The GPS dilution of precision value.
//   - [ICSSearchableItemAttributeSet.SetGPSDOP]
//   - [ICSSearchableItemAttributeSet.GPSDateStamp]: The date and time related to the GPS value.
//   - [ICSSearchableItemAttributeSet.SetGPSDateStamp]
//   - [ICSSearchableItemAttributeSet.GPSDestBearing]: The bearing to the destination point.
//   - [ICSSearchableItemAttributeSet.SetGPSDestBearing]
//   - [ICSSearchableItemAttributeSet.GPSDestDistance]: The distance to the destination point.
//   - [ICSSearchableItemAttributeSet.SetGPSDestDistance]
//   - [ICSSearchableItemAttributeSet.GPSDestLatitude]: The latitude of the destination point.
//   - [ICSSearchableItemAttributeSet.SetGPSDestLatitude]
//   - [ICSSearchableItemAttributeSet.GPSDestLongitude]: The longitude of the destination point.
//   - [ICSSearchableItemAttributeSet.SetGPSDestLongitude]
//   - [ICSSearchableItemAttributeSet.GPSDifferental]: The differential correction applied to the GPS receiver.
//   - [ICSSearchableItemAttributeSet.SetGPSDifferental]
//   - [ICSSearchableItemAttributeSet.GPSMapDatum]: The geodetic data that the GPS receiver uses.
//   - [ICSSearchableItemAttributeSet.SetGPSMapDatum]
//   - [ICSSearchableItemAttributeSet.GPSMeasureMode]: The measurement precision mode in use by the GPS receiver.
//   - [ICSSearchableItemAttributeSet.SetGPSMeasureMode]
//   - [ICSSearchableItemAttributeSet.GPSProcessingMethod]: The location finding method that the GPS receiver uses.
//   - [ICSSearchableItemAttributeSet.SetGPSProcessingMethod]
//   - [ICSSearchableItemAttributeSet.GPSStatus]: The status of the GPS receiver.
//   - [ICSSearchableItemAttributeSet.SetGPSStatus]
//   - [ICSSearchableItemAttributeSet.GPSTrack]: The direction of travel of the item in degrees from true north.
//   - [ICSSearchableItemAttributeSet.SetGPSTrack]
//   - [ICSSearchableItemAttributeSet.Headline]: A publishable string that provides a synopsis of the contents of the item.
//   - [ICSSearchableItemAttributeSet.SetHeadline]
//   - [ICSSearchableItemAttributeSet.ImageDirection]: The direction of the item’s image in degrees from true north.
//   - [ICSSearchableItemAttributeSet.SetImageDirection]
//   - [ICSSearchableItemAttributeSet.Instructions]: Instructions that concern the use of the item, such as an embargo or warning.
//   - [ICSSearchableItemAttributeSet.SetInstructions]
//   - [ICSSearchableItemAttributeSet.Latitude]: The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//   - [ICSSearchableItemAttributeSet.SetLatitude]
//   - [ICSSearchableItemAttributeSet.Longitude]: The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//   - [ICSSearchableItemAttributeSet.SetLongitude]
//   - [ICSSearchableItemAttributeSet.NamedLocation]: The name of the location or point of interest associated with the item.
//   - [ICSSearchableItemAttributeSet.SetNamedLocation]
//   - [ICSSearchableItemAttributeSet.Speed]: The speed of the item, in kilometers per hour.
//   - [ICSSearchableItemAttributeSet.SetSpeed]
//   - [ICSSearchableItemAttributeSet.StateOrProvince]: The province or state of origin according to guidelines the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetStateOrProvince]
//   - [ICSSearchableItemAttributeSet.Timestamp]: The timestamp on the item.
//   - [ICSSearchableItemAttributeSet.SetTimestamp]
//   - [ICSSearchableItemAttributeSet.FullyFormattedAddress]: The fully formatted address of the item, received from MapKit.
//   - [ICSSearchableItemAttributeSet.SetFullyFormattedAddress]
//   - [ICSSearchableItemAttributeSet.PostalCode]: The postal code for the item according to guidelines the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetPostalCode]
//   - [ICSSearchableItemAttributeSet.SubThoroughfare]: The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetSubThoroughfare]
//   - [ICSSearchableItemAttributeSet.Thoroughfare]: The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//   - [ICSSearchableItemAttributeSet.SetThoroughfare]
//
// # Describing media
//
//   - [ICSSearchableItemAttributeSet.Comment]: A comment related to the media file.
//   - [ICSSearchableItemAttributeSet.SetComment]
//   - [ICSSearchableItemAttributeSet.ContentCreationDate]: The creation date of an edited or optimized version of the song or composition.
//   - [ICSSearchableItemAttributeSet.SetContentCreationDate]
//   - [ICSSearchableItemAttributeSet.ContentModificationDate]: The date on which the contents of the file was last modified.
//   - [ICSSearchableItemAttributeSet.SetContentModificationDate]
//   - [ICSSearchableItemAttributeSet.ContentSources]: An array of sources from which the media was obtained.
//   - [ICSSearchableItemAttributeSet.SetContentSources]
//   - [ICSSearchableItemAttributeSet.Copyright]: The copyright date of the content.
//   - [ICSSearchableItemAttributeSet.SetCopyright]
//   - [ICSSearchableItemAttributeSet.DownloadedDate]: The most recent date on which the file was downloaded or received.
//   - [ICSSearchableItemAttributeSet.SetDownloadedDate]
//   - [ICSSearchableItemAttributeSet.Editors]: A list of editors who have worked on the file.
//   - [ICSSearchableItemAttributeSet.SetEditors]
//   - [ICSSearchableItemAttributeSet.LastUsedDate]: The date on which the file was last used.
//   - [ICSSearchableItemAttributeSet.SetLastUsedDate]
//   - [ICSSearchableItemAttributeSet.Participants]: A list of people who are visible in an image or movie or written about in a document.
//   - [ICSSearchableItemAttributeSet.SetParticipants]
//   - [ICSSearchableItemAttributeSet.Projects]: A list of projects of which this file is a part.
//   - [ICSSearchableItemAttributeSet.SetProjects]
//   - [ICSSearchableItemAttributeSet.AddedDate]: The date on which the item was moved into its current location.
//   - [ICSSearchableItemAttributeSet.SetAddedDate]
//   - [ICSSearchableItemAttributeSet.Codecs]: The codecs used to encode/decode the media.
//   - [ICSSearchableItemAttributeSet.SetCodecs]
//   - [ICSSearchableItemAttributeSet.ContactKeywords]: A list of contacts who are associated with the content in some way, not including the author.
//   - [ICSSearchableItemAttributeSet.SetContactKeywords]
//   - [ICSSearchableItemAttributeSet.DeliveryType]: The delivery type of the file.
//   - [ICSSearchableItemAttributeSet.SetDeliveryType]
//   - [ICSSearchableItemAttributeSet.Duration]: The duration (if appropriate) of the content of the file, in seconds.
//   - [ICSSearchableItemAttributeSet.SetDuration]
//   - [ICSSearchableItemAttributeSet.MediaTypes]: The media types present in the content.
//   - [ICSSearchableItemAttributeSet.SetMediaTypes]
//   - [ICSSearchableItemAttributeSet.Organizations]: A list of companies or organizations that created the content.
//   - [ICSSearchableItemAttributeSet.SetOrganizations]
//   - [ICSSearchableItemAttributeSet.Streamable]: A value that indicates if the content is prepared for streaming.
//   - [ICSSearchableItemAttributeSet.SetStreamable]
//   - [ICSSearchableItemAttributeSet.TotalBitRate]: The total bit rate of the media, combining audio and video.
//   - [ICSSearchableItemAttributeSet.SetTotalBitRate]
//   - [ICSSearchableItemAttributeSet.AudioBitRate]: The audio bit rate of the media.
//   - [ICSSearchableItemAttributeSet.SetAudioBitRate]
//   - [ICSSearchableItemAttributeSet.Version]: A version string associated with the file.
//   - [ICSSearchableItemAttributeSet.SetVersion]
//   - [ICSSearchableItemAttributeSet.VideoBitRate]: The video bit rate of the media.
//   - [ICSSearchableItemAttributeSet.SetVideoBitRate]
//   - [ICSSearchableItemAttributeSet.Contributors]: A list of people, organizations, or services that made contributions to the media content.
//   - [ICSSearchableItemAttributeSet.SetContributors]
//   - [ICSSearchableItemAttributeSet.Languages]: A list of the included languages for the intellectual content of the media.
//   - [ICSSearchableItemAttributeSet.SetLanguages]
//   - [ICSSearchableItemAttributeSet.Publishers]: A list of people, organizations, services, or other entities responsible for making the media available.
//   - [ICSSearchableItemAttributeSet.SetPublishers]
//   - [ICSSearchableItemAttributeSet.Rights]: A link to information about the rights held in and over the media.
//   - [ICSSearchableItemAttributeSet.SetRights]
//   - [ICSSearchableItemAttributeSet.Role]: Indicates the role of the content creator.
//   - [ICSSearchableItemAttributeSet.SetRole]
//   - [ICSSearchableItemAttributeSet.ContentRating]: A value that indicates if the media contains explicit content.
//   - [ICSSearchableItemAttributeSet.SetContentRating]
//   - [ICSSearchableItemAttributeSet.Coverage]: A list of descriptors that specify the extent or scope of the media.
//   - [ICSSearchableItemAttributeSet.SetCoverage]
//   - [ICSSearchableItemAttributeSet.Director]: The name of the director of the media (for example, a movie director).
//   - [ICSSearchableItemAttributeSet.SetDirector]
//   - [ICSSearchableItemAttributeSet.Genre]: The genre of the media.
//   - [ICSSearchableItemAttributeSet.SetGenre]
//   - [ICSSearchableItemAttributeSet.Information]: Information about the media.
//   - [ICSSearchableItemAttributeSet.SetInformation]
//   - [ICSSearchableItemAttributeSet.Local]: A value that indicates if the media is local.
//   - [ICSSearchableItemAttributeSet.SetLocal]
//   - [ICSSearchableItemAttributeSet.OriginalFormat]: The original format of the media.
//   - [ICSSearchableItemAttributeSet.SetOriginalFormat]
//   - [ICSSearchableItemAttributeSet.OriginalSource]: The original source of the media.
//   - [ICSSearchableItemAttributeSet.SetOriginalSource]
//   - [ICSSearchableItemAttributeSet.Performers]: A list of performers in the media.
//   - [ICSSearchableItemAttributeSet.SetPerformers]
//   - [ICSSearchableItemAttributeSet.PlayCount]: A user-supplied play count for the media.
//   - [ICSSearchableItemAttributeSet.SetPlayCount]
//   - [ICSSearchableItemAttributeSet.Producer]: The producer of the content.
//   - [ICSSearchableItemAttributeSet.SetProducer]
//   - [ICSSearchableItemAttributeSet.Rating]: The user-supplied rating of the media.
//   - [ICSSearchableItemAttributeSet.SetRating]
//   - [ICSSearchableItemAttributeSet.RatingDescription]: A description of the rating.
//   - [ICSSearchableItemAttributeSet.SetRatingDescription]
//   - [ICSSearchableItemAttributeSet.URL]: The URL associated with the media.
//   - [ICSSearchableItemAttributeSet.SetURL]
//
// # Describing music
//
//   - [ICSSearchableItemAttributeSet.Album]: The title for a collection of audio media.
//   - [ICSSearchableItemAttributeSet.SetAlbum]
//   - [ICSSearchableItemAttributeSet.Artist]: The artist associated with the media.
//   - [ICSSearchableItemAttributeSet.SetArtist]
//   - [ICSSearchableItemAttributeSet.AudioChannelCount]: The number of channels in the audio data that the file contains.
//   - [ICSSearchableItemAttributeSet.SetAudioChannelCount]
//   - [ICSSearchableItemAttributeSet.AudioEncodingApplication]: The name of the application that encoded the data the audio file contains.
//   - [ICSSearchableItemAttributeSet.SetAudioEncodingApplication]
//   - [ICSSearchableItemAttributeSet.AudioSampleRate]: The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//   - [ICSSearchableItemAttributeSet.SetAudioSampleRate]
//   - [ICSSearchableItemAttributeSet.AudioTrackNumber]: The track number of a song or audio composition when part of an album.
//   - [ICSSearchableItemAttributeSet.SetAudioTrackNumber]
//   - [ICSSearchableItemAttributeSet.Composer]: The composer of the song or audio composition that the audio file contains.
//   - [ICSSearchableItemAttributeSet.SetComposer]
//   - [ICSSearchableItemAttributeSet.KeySignature]: The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//   - [ICSSearchableItemAttributeSet.SetKeySignature]
//   - [ICSSearchableItemAttributeSet.Lyricist]: The lyricist or text writer for the song or audio composition that the file contains.
//   - [ICSSearchableItemAttributeSet.SetLyricist]
//   - [ICSSearchableItemAttributeSet.MusicalGenre]: The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//   - [ICSSearchableItemAttributeSet.SetMusicalGenre]
//   - [ICSSearchableItemAttributeSet.RecordingDate]: The recording date of the song or composition.
//   - [ICSSearchableItemAttributeSet.SetRecordingDate]
//   - [ICSSearchableItemAttributeSet.Tempo]: The tempo of the music that the audio file contains, in beats per minute.
//   - [ICSSearchableItemAttributeSet.SetTempo]
//   - [ICSSearchableItemAttributeSet.TimeSignature]: The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//   - [ICSSearchableItemAttributeSet.SetTimeSignature]
//   - [ICSSearchableItemAttributeSet.GeneralMIDISequence]: A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//   - [ICSSearchableItemAttributeSet.SetGeneralMIDISequence]
//   - [ICSSearchableItemAttributeSet.MusicalInstrumentCategory]: The category of the instrument associated with the audio file.
//   - [ICSSearchableItemAttributeSet.SetMusicalInstrumentCategory]
//   - [ICSSearchableItemAttributeSet.MusicalInstrumentName]: The name of an instrument within the context of an instrument category.
//   - [ICSSearchableItemAttributeSet.SetMusicalInstrumentName]
//
// # Describing images
//
//   - [ICSSearchableItemAttributeSet.ISOSpeed]: The ISO speed setting at the time the camera captured the image.
//   - [ICSSearchableItemAttributeSet.SetISOSpeed]
//   - [ICSSearchableItemAttributeSet.AcquisitionMake]: The manufacturer of the device that captured the image.
//   - [ICSSearchableItemAttributeSet.SetAcquisitionMake]
//   - [ICSSearchableItemAttributeSet.AcquisitionModel]: The model of the device that captured the image.
//   - [ICSSearchableItemAttributeSet.SetAcquisitionModel]
//   - [ICSSearchableItemAttributeSet.Aperture]: The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//   - [ICSSearchableItemAttributeSet.SetAperture]
//   - [ICSSearchableItemAttributeSet.BitsPerSample]: The number of bits per sample.
//   - [ICSSearchableItemAttributeSet.SetBitsPerSample]
//   - [ICSSearchableItemAttributeSet.CameraOwner]: The owner of the camera that captured the image.
//   - [ICSSearchableItemAttributeSet.SetCameraOwner]
//   - [ICSSearchableItemAttributeSet.ColorSpace]: The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//   - [ICSSearchableItemAttributeSet.SetColorSpace]
//   - [ICSSearchableItemAttributeSet.FlashOn]: A value that indicates if the camera used a flash to capture the image.
//   - [ICSSearchableItemAttributeSet.SetFlashOn]
//   - [ICSSearchableItemAttributeSet.FocalLength]: The actual focal length of the lens, in millimeters.
//   - [ICSSearchableItemAttributeSet.SetFocalLength]
//   - [ICSSearchableItemAttributeSet.FocalLength35mm]: A value that indicates if the focal length is 35mm.
//   - [ICSSearchableItemAttributeSet.SetFocalLength35mm]
//   - [ICSSearchableItemAttributeSet.LayerNames]: An array that contains the names of the various layers in the file.
//   - [ICSSearchableItemAttributeSet.SetLayerNames]
//   - [ICSSearchableItemAttributeSet.LensModel]: The model of the lens that captured the image.
//   - [ICSSearchableItemAttributeSet.SetLensModel]
//   - [ICSSearchableItemAttributeSet.Orientation]: The orientation of the data.
//   - [ICSSearchableItemAttributeSet.SetOrientation]
//   - [ICSSearchableItemAttributeSet.PixelCount]: The total number of pixels in the image.
//   - [ICSSearchableItemAttributeSet.SetPixelCount]
//   - [ICSSearchableItemAttributeSet.PixelHeight]: The height of the item, such as image or video frame height, in pixels.
//   - [ICSSearchableItemAttributeSet.SetPixelHeight]
//   - [ICSSearchableItemAttributeSet.PixelWidth]: The width of the item, such as image or video frame width, in pixels.
//   - [ICSSearchableItemAttributeSet.SetPixelWidth]
//   - [ICSSearchableItemAttributeSet.WhiteBalance]: The white balance setting when the camera captured the image.
//   - [ICSSearchableItemAttributeSet.SetWhiteBalance]
//   - [ICSSearchableItemAttributeSet.EXIFGPSVersion]: The version of GPS Info IFD header that was used to generate the metadata for the image.
//   - [ICSSearchableItemAttributeSet.SetEXIFGPSVersion]
//   - [ICSSearchableItemAttributeSet.EXIFVersion]: The version of the EXIF header that was used to generate the metadata for the image.
//   - [ICSSearchableItemAttributeSet.SetEXIFVersion]
//   - [ICSSearchableItemAttributeSet.ExposureMode]: The mode the camera used for the exposure of the image.
//   - [ICSSearchableItemAttributeSet.SetExposureMode]
//   - [ICSSearchableItemAttributeSet.ExposureProgram]: The class of the program the camera used to set exposure when capturing the image.
//   - [ICSSearchableItemAttributeSet.SetExposureProgram]
//   - [ICSSearchableItemAttributeSet.ExposureTime]: The time that the lens was open during exposure, in seconds.
//   - [ICSSearchableItemAttributeSet.SetExposureTime]
//   - [ICSSearchableItemAttributeSet.ExposureTimeString]: The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//   - [ICSSearchableItemAttributeSet.SetExposureTimeString]
//   - [ICSSearchableItemAttributeSet.FNumber]: The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//   - [ICSSearchableItemAttributeSet.SetFNumber]
//   - [ICSSearchableItemAttributeSet.HasAlphaChannel]: Indicates if the image file has an alpha channel.
//   - [ICSSearchableItemAttributeSet.SetHasAlphaChannel]
//   - [ICSSearchableItemAttributeSet.MaxAperture]: The smallest F number of the lens.
//   - [ICSSearchableItemAttributeSet.SetMaxAperture]
//   - [ICSSearchableItemAttributeSet.MeteringMode]: The metering mode.
//   - [ICSSearchableItemAttributeSet.SetMeteringMode]
//   - [ICSSearchableItemAttributeSet.ProfileName]: The name of the color profile the camera used for the image.
//   - [ICSSearchableItemAttributeSet.SetProfileName]
//   - [ICSSearchableItemAttributeSet.RedEyeOn]: A value that indicates if the camera used red-eye reduction when capturing the image.
//   - [ICSSearchableItemAttributeSet.SetRedEyeOn]
//   - [ICSSearchableItemAttributeSet.ResolutionHeightDPI]: The resolution height of the image, in DPI.
//   - [ICSSearchableItemAttributeSet.SetResolutionHeightDPI]
//   - [ICSSearchableItemAttributeSet.ResolutionWidthDPI]: The resolution width of the image, in DPI.
//   - [ICSSearchableItemAttributeSet.SetResolutionWidthDPI]
//
// # Describing messages
//
//   - [ICSSearchableItemAttributeSet.HTMLContentData]: The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//   - [ICSSearchableItemAttributeSet.SetHTMLContentData]
//   - [ICSSearchableItemAttributeSet.AccountHandles]: An array of the canonical handles for the account with which the message is associated.
//   - [ICSSearchableItemAttributeSet.SetAccountHandles]
//   - [ICSSearchableItemAttributeSet.AccountIdentifier]: The unique identifier for the account with which the message is associated, if any.
//   - [ICSSearchableItemAttributeSet.SetAccountIdentifier]
//   - [ICSSearchableItemAttributeSet.AdditionalRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Cc: field in an email message.
//   - [ICSSearchableItemAttributeSet.SetAdditionalRecipients]
//   - [ICSSearchableItemAttributeSet.AuthorAddresses]: An array of addresses associated with the author of the message.
//   - [ICSSearchableItemAttributeSet.SetAuthorAddresses]
//   - [ICSSearchableItemAttributeSet.AuthorEmailAddresses]: An array of email addresses associated with the author of the message.
//   - [ICSSearchableItemAttributeSet.SetAuthorEmailAddresses]
//   - [ICSSearchableItemAttributeSet.AuthorNames]: An array of names representing the authors who have worked on the message.
//   - [ICSSearchableItemAttributeSet.SetAuthorNames]
//   - [ICSSearchableItemAttributeSet.Authors]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the From: field in an item.
//   - [ICSSearchableItemAttributeSet.SetAuthors]
//   - [ICSSearchableItemAttributeSet.EmailAddresses]: An array of email addresses associated with the message.
//   - [ICSSearchableItemAttributeSet.SetEmailAddresses]
//   - [ICSSearchableItemAttributeSet.EmailHeaders]: A dictionary that contains all the headers of the message.
//   - [ICSSearchableItemAttributeSet.SetEmailHeaders]
//   - [ICSSearchableItemAttributeSet.HiddenAdditionalRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Bcc: field in an email message.
//   - [ICSSearchableItemAttributeSet.SetHiddenAdditionalRecipients]
//   - [ICSSearchableItemAttributeSet.InstantMessageAddresses]: An array of instant message addresses for the message.
//   - [ICSSearchableItemAttributeSet.SetInstantMessageAddresses]
//   - [ICSSearchableItemAttributeSet.LikelyJunk]: A value that indicates if the message is likely to be considered junk.
//   - [ICSSearchableItemAttributeSet.SetLikelyJunk]
//   - [ICSSearchableItemAttributeSet.MailboxIdentifiers]: An array of mailbox identifiers associated with the message.
//   - [ICSSearchableItemAttributeSet.SetMailboxIdentifiers]
//   - [ICSSearchableItemAttributeSet.PhoneNumbers]: An array of phone numbers associated with the message.
//   - [ICSSearchableItemAttributeSet.SetPhoneNumbers]
//   - [ICSSearchableItemAttributeSet.PrimaryRecipients]: An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the To: field in an email message.
//   - [ICSSearchableItemAttributeSet.SetPrimaryRecipients]
//   - [ICSSearchableItemAttributeSet.RecipientAddresses]: An array of addresses associated with the recipients of the message.
//   - [ICSSearchableItemAttributeSet.SetRecipientAddresses]
//   - [ICSSearchableItemAttributeSet.RecipientEmailAddresses]: An array of email addresses associated with the recipient.
//   - [ICSSearchableItemAttributeSet.SetRecipientEmailAddresses]
//   - [ICSSearchableItemAttributeSet.RecipientNames]: An array of names representing the recipients of this message.
//   - [ICSSearchableItemAttributeSet.SetRecipientNames]
//   - [ICSSearchableItemAttributeSet.TextContent]: The textual content of the message.
//   - [ICSSearchableItemAttributeSet.SetTextContent]
//
// # Describing containment
//
//   - [ICSSearchableItemAttributeSet.ContainerDisplayName]: A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//   - [ICSSearchableItemAttributeSet.SetContainerDisplayName]
//   - [ICSSearchableItemAttributeSet.ContainerIdentifier]: The identifier of the container to which the item belongs.
//   - [ICSSearchableItemAttributeSet.SetContainerIdentifier]
//   - [ICSSearchableItemAttributeSet.ContainerOrder]: The order of the item within the container.
//   - [ICSSearchableItemAttributeSet.SetContainerOrder]
//   - [ICSSearchableItemAttributeSet.ContainerTitle]: The title of the container to which the item belongs.
//   - [ICSSearchableItemAttributeSet.SetContainerTitle]
//
// # Describing supporting actions
//
//   - [ICSSearchableItemAttributeSet.SupportsNavigation]: A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//   - [ICSSearchableItemAttributeSet.SetSupportsNavigation]
//   - [ICSSearchableItemAttributeSet.SupportsPhoneCall]: A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//   - [ICSSearchableItemAttributeSet.SetSupportsPhoneCall]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet
type ICSSearchableItemAttributeSet interface {
	objectivec.IObject

	// Topic: Creating an attribute set

	// Creates an attribute set for the specified content type.
	InitWithContentType(contentType uniformtypeidentifiers.UTType) CSSearchableItemAttributeSet
	InitWithCoder(coder foundation.INSCoder) CSSearchableItemAttributeSet

	// Topic: Accessing custom attributes

	// Sets the value for a custom attribute key.
	SetValueForCustomKey(value foundation.NSSecureCoding, key ICSCustomAttributeKey)
	// Returns the value associated with the specified custom attribute key.
	ValueForCustomKey(key ICSCustomAttributeKey) foundation.NSSecureCoding

	// Topic: Handling Apple Intelligence prioritization and summarization

	// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
	IsPriority() foundation.NSNumber
	// A string that presents the Apple Intelligence summarization of the item.
	TextContentSummary() string
	// A string that represents the text the system transcribed.
	TranscribedTextContent() string
	SetTranscribedTextContent(value string)

	// Topic: Providing item representations

	// An array of identifiers that corresponds to data representations the delegate provides.
	ProviderDataTypeIdentifiers() []string
	SetProviderDataTypeIdentifiers(value []string)
	// An array of identifiers that corresponds to file representations the delegate provides.
	ProviderFileTypeIdentifiers() []string
	SetProviderFileTypeIdentifiers(value []string)
	// An array of identifiers that corresponds to in-place file representations the delegate provides.
	ProviderInPlaceFileTypeIdentifiers() []string
	SetProviderInPlaceFileTypeIdentifiers(value []string)

	// Topic: Describing general attributes

	// An array of localized strings that represent alternate display names for the item.
	AlternateNames() []string
	SetAlternateNames(value []string)
	// The uniform type identifier (UTI) of the item.
	ContentType() string
	SetContentType(value string)
	// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
	ContentTypeTree() []string
	SetContentTypeTree(value []string)
	// The file URL of the content to index.
	ContentURL() foundation.NSURL
	SetContentURL(value foundation.NSURL)
	// The local file URL of the thumbnail image for the item when Dark Mode is active.
	DarkThumbnailURL() foundation.NSURL
	SetDarkThumbnailURL(value foundation.NSURL)
	// A localized string that contains the name of the item, suitable to display in the user interface.
	DisplayName() string
	SetDisplayName(value string)
	// An array of keywords associated with the item, such as work, birthday, important, and so on.
	Keywords() []string
	SetKeywords(value []string)
	// The date on which the last metadata attribute was changed.
	MetadataModificationDate() foundation.NSDate
	SetMetadataModificationDate(value foundation.NSDate)
	// The complete path to the item.
	Path() string
	SetPath(value string)
	// A number that indicates the relative importance of the item among other items from the app.
	RankingHint() foundation.NSNumber
	SetRankingHint(value foundation.NSNumber)
	// The unique identifier for the item to which the activity is related.
	RelatedUniqueIdentifier() string
	SetRelatedUniqueIdentifier(value string)
	// Image data that represents the thumbnail of the item.
	ThumbnailData() foundation.NSData
	SetThumbnailData(value foundation.NSData)
	// The local file URL of the thumbnail image for the item.
	ThumbnailURL() foundation.NSURL
	SetThumbnailURL(value foundation.NSURL)
	// The title of the item.
	Title() string
	SetTitle(value string)
	// An identifier that represents the domain or owner of the item.
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	// The unique identifier for the item to which the activity is related, but not linked.
	WeakRelatedUniqueIdentifier() string
	SetWeakRelatedUniqueIdentifier(value string)

	// Topic: Describing document content

	// A class of entity for which the item is intended or useful.
	Audiences() []string
	SetAudiences(value []string)
	// A description of the item’s content.
	ContentDescription() string
	SetContentDescription(value string)
	// The name of the app that created the content.
	Creator() string
	SetCreator(value string)
	// The name of the apps that converted the original content into a PDF stream.
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	// The size of the document file.
	FileSize() foundation.NSNumber
	SetFileSize(value foundation.NSNumber)
	// An array of font names the document uses.
	FontNames() []string
	SetFontNames(value []string)
	// A formal identifier that references the document the item represents.
	Identifier() string
	SetIdentifier(value string)
	// A description of the kind of document the item represents.
	Kind() string
	SetKind(value string)
	// The number of pages in the document.
	PageCount() foundation.NSNumber
	SetPageCount(value foundation.NSNumber)
	// The height of the document page, in points (72 points per inch).
	PageHeight() foundation.NSNumber
	SetPageHeight(value foundation.NSNumber)
	// The width of the document page, in points (72 points per inch).
	PageWidth() foundation.NSNumber
	SetPageWidth(value foundation.NSNumber)
	// The security method (a type of encryption) that protects the document file.
	SecurityMethod() string
	SetSecurityMethod(value string)
	// The subject of the document.
	Subject() string
	SetSubject(value string)
	// The theme of the document.
	Theme() string
	SetTheme(value string)

	// Topic: Describing user involvement

	// A value that indicates the user created the item.
	UserCreated() foundation.NSNumber
	SetUserCreated(value foundation.NSNumber)
	// A value that indicates the user selected the item.
	UserCurated() foundation.NSNumber
	SetUserCurated(value foundation.NSNumber)
	// A value that indicates the user purchased or owns the item.
	UserOwned() foundation.NSNumber
	SetUserOwned(value foundation.NSNumber)

	// Topic: Describing events

	// A value that indicates if the event covers an entire day.
	AllDay() foundation.NSNumber
	SetAllDay(value foundation.NSNumber)
	// The date on which the item was completed.
	CompletionDate() foundation.NSDate
	SetCompletionDate(value foundation.NSDate)
	// The date on which the item is due.
	DueDate() foundation.NSDate
	SetDueDate(value foundation.NSDate)
	// The end date for the item.
	EndDate() foundation.NSDate
	SetEndDate(value foundation.NSDate)
	// An array of important dates associated with the item.
	ImportantDates() []foundation.NSDate
	SetImportantDates(value []foundation.NSDate)
	// The start date for the item.
	StartDate() foundation.NSDate
	SetStartDate(value foundation.NSDate)

	// Topic: Describing places

	// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
	Altitude() foundation.NSNumber
	SetAltitude(value foundation.NSNumber)
	// The city of the item’s origin according to guidelines that the provider establishes.
	City() string
	SetCity(value string)
	// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
	Country() string
	SetCountry(value string)
	// Information about the GPS area.
	GPSAreaInformation() string
	SetGPSAreaInformation(value string)
	// The GPS dilution of precision value.
	GPSDOP() foundation.NSNumber
	SetGPSDOP(value foundation.NSNumber)
	// The date and time related to the GPS value.
	GPSDateStamp() foundation.NSDate
	SetGPSDateStamp(value foundation.NSDate)
	// The bearing to the destination point.
	GPSDestBearing() foundation.NSNumber
	SetGPSDestBearing(value foundation.NSNumber)
	// The distance to the destination point.
	GPSDestDistance() foundation.NSNumber
	SetGPSDestDistance(value foundation.NSNumber)
	// The latitude of the destination point.
	GPSDestLatitude() foundation.NSNumber
	SetGPSDestLatitude(value foundation.NSNumber)
	// The longitude of the destination point.
	GPSDestLongitude() foundation.NSNumber
	SetGPSDestLongitude(value foundation.NSNumber)
	// The differential correction applied to the GPS receiver.
	GPSDifferental() foundation.NSNumber
	SetGPSDifferental(value foundation.NSNumber)
	// The geodetic data that the GPS receiver uses.
	GPSMapDatum() string
	SetGPSMapDatum(value string)
	// The measurement precision mode in use by the GPS receiver.
	GPSMeasureMode() string
	SetGPSMeasureMode(value string)
	// The location finding method that the GPS receiver uses.
	GPSProcessingMethod() string
	SetGPSProcessingMethod(value string)
	// The status of the GPS receiver.
	GPSStatus() string
	SetGPSStatus(value string)
	// The direction of travel of the item in degrees from true north.
	GPSTrack() foundation.NSNumber
	SetGPSTrack(value foundation.NSNumber)
	// A publishable string that provides a synopsis of the contents of the item.
	Headline() string
	SetHeadline(value string)
	// The direction of the item’s image in degrees from true north.
	ImageDirection() foundation.NSNumber
	SetImageDirection(value foundation.NSNumber)
	// Instructions that concern the use of the item, such as an embargo or warning.
	Instructions() string
	SetInstructions(value string)
	// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
	Latitude() foundation.NSNumber
	SetLatitude(value foundation.NSNumber)
	// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
	Longitude() foundation.NSNumber
	SetLongitude(value foundation.NSNumber)
	// The name of the location or point of interest associated with the item.
	NamedLocation() string
	SetNamedLocation(value string)
	// The speed of the item, in kilometers per hour.
	Speed() foundation.NSNumber
	SetSpeed(value foundation.NSNumber)
	// The province or state of origin according to guidelines the provider establishes.
	StateOrProvince() string
	SetStateOrProvince(value string)
	// The timestamp on the item.
	Timestamp() foundation.NSDate
	SetTimestamp(value foundation.NSDate)
	// The fully formatted address of the item, received from MapKit.
	FullyFormattedAddress() string
	SetFullyFormattedAddress(value string)
	// The postal code for the item according to guidelines the provider establishes.
	PostalCode() string
	SetPostalCode(value string)
	// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
	SubThoroughfare() string
	SetSubThoroughfare(value string)
	// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
	Thoroughfare() string
	SetThoroughfare(value string)

	// Topic: Describing media

	// A comment related to the media file.
	Comment() string
	SetComment(value string)
	// The creation date of an edited or optimized version of the song or composition.
	ContentCreationDate() foundation.NSDate
	SetContentCreationDate(value foundation.NSDate)
	// The date on which the contents of the file was last modified.
	ContentModificationDate() foundation.NSDate
	SetContentModificationDate(value foundation.NSDate)
	// An array of sources from which the media was obtained.
	ContentSources() []string
	SetContentSources(value []string)
	// The copyright date of the content.
	Copyright() string
	SetCopyright(value string)
	// The most recent date on which the file was downloaded or received.
	DownloadedDate() foundation.NSDate
	SetDownloadedDate(value foundation.NSDate)
	// A list of editors who have worked on the file.
	Editors() []string
	SetEditors(value []string)
	// The date on which the file was last used.
	LastUsedDate() foundation.NSDate
	SetLastUsedDate(value foundation.NSDate)
	// A list of people who are visible in an image or movie or written about in a document.
	Participants() []string
	SetParticipants(value []string)
	// A list of projects of which this file is a part.
	Projects() []string
	SetProjects(value []string)
	// The date on which the item was moved into its current location.
	AddedDate() foundation.NSDate
	SetAddedDate(value foundation.NSDate)
	// The codecs used to encode/decode the media.
	Codecs() []string
	SetCodecs(value []string)
	// A list of contacts who are associated with the content in some way, not including the author.
	ContactKeywords() []string
	SetContactKeywords(value []string)
	// The delivery type of the file.
	DeliveryType() foundation.NSNumber
	SetDeliveryType(value foundation.NSNumber)
	// The duration (if appropriate) of the content of the file, in seconds.
	Duration() foundation.NSNumber
	SetDuration(value foundation.NSNumber)
	// The media types present in the content.
	MediaTypes() []string
	SetMediaTypes(value []string)
	// A list of companies or organizations that created the content.
	Organizations() []string
	SetOrganizations(value []string)
	// A value that indicates if the content is prepared for streaming.
	Streamable() foundation.NSNumber
	SetStreamable(value foundation.NSNumber)
	// The total bit rate of the media, combining audio and video.
	TotalBitRate() foundation.NSNumber
	SetTotalBitRate(value foundation.NSNumber)
	// The audio bit rate of the media.
	AudioBitRate() foundation.NSNumber
	SetAudioBitRate(value foundation.NSNumber)
	// A version string associated with the file.
	Version() string
	SetVersion(value string)
	// The video bit rate of the media.
	VideoBitRate() foundation.NSNumber
	SetVideoBitRate(value foundation.NSNumber)
	// A list of people, organizations, or services that made contributions to the media content.
	Contributors() []string
	SetContributors(value []string)
	// A list of the included languages for the intellectual content of the media.
	Languages() []string
	SetLanguages(value []string)
	// A list of people, organizations, services, or other entities responsible for making the media available.
	Publishers() []string
	SetPublishers(value []string)
	// A link to information about the rights held in and over the media.
	Rights() string
	SetRights(value string)
	// Indicates the role of the content creator.
	Role() string
	SetRole(value string)
	// A value that indicates if the media contains explicit content.
	ContentRating() foundation.NSNumber
	SetContentRating(value foundation.NSNumber)
	// A list of descriptors that specify the extent or scope of the media.
	Coverage() []string
	SetCoverage(value []string)
	// The name of the director of the media (for example, a movie director).
	Director() string
	SetDirector(value string)
	// The genre of the media.
	Genre() string
	SetGenre(value string)
	// Information about the media.
	Information() string
	SetInformation(value string)
	// A value that indicates if the media is local.
	Local() foundation.NSNumber
	SetLocal(value foundation.NSNumber)
	// The original format of the media.
	OriginalFormat() string
	SetOriginalFormat(value string)
	// The original source of the media.
	OriginalSource() string
	SetOriginalSource(value string)
	// A list of performers in the media.
	Performers() []string
	SetPerformers(value []string)
	// A user-supplied play count for the media.
	PlayCount() foundation.NSNumber
	SetPlayCount(value foundation.NSNumber)
	// The producer of the content.
	Producer() string
	SetProducer(value string)
	// The user-supplied rating of the media.
	Rating() foundation.NSNumber
	SetRating(value foundation.NSNumber)
	// A description of the rating.
	RatingDescription() string
	SetRatingDescription(value string)
	// The URL associated with the media.
	URL() foundation.NSURL
	SetURL(value foundation.NSURL)

	// Topic: Describing music

	// The title for a collection of audio media.
	Album() string
	SetAlbum(value string)
	// The artist associated with the media.
	Artist() string
	SetArtist(value string)
	// The number of channels in the audio data that the file contains.
	AudioChannelCount() foundation.NSNumber
	SetAudioChannelCount(value foundation.NSNumber)
	// The name of the application that encoded the data the audio file contains.
	AudioEncodingApplication() string
	SetAudioEncodingApplication(value string)
	// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
	AudioSampleRate() foundation.NSNumber
	SetAudioSampleRate(value foundation.NSNumber)
	// The track number of a song or audio composition when part of an album.
	AudioTrackNumber() foundation.NSNumber
	SetAudioTrackNumber(value foundation.NSNumber)
	// The composer of the song or audio composition that the audio file contains.
	Composer() string
	SetComposer(value string)
	// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
	KeySignature() string
	SetKeySignature(value string)
	// The lyricist or text writer for the song or audio composition that the file contains.
	Lyricist() string
	SetLyricist(value string)
	// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
	MusicalGenre() string
	SetMusicalGenre(value string)
	// The recording date of the song or composition.
	RecordingDate() foundation.NSDate
	SetRecordingDate(value foundation.NSDate)
	// The tempo of the music that the audio file contains, in beats per minute.
	Tempo() foundation.NSNumber
	SetTempo(value foundation.NSNumber)
	// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
	TimeSignature() string
	SetTimeSignature(value string)
	// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
	GeneralMIDISequence() foundation.NSNumber
	SetGeneralMIDISequence(value foundation.NSNumber)
	// The category of the instrument associated with the audio file.
	MusicalInstrumentCategory() string
	SetMusicalInstrumentCategory(value string)
	// The name of an instrument within the context of an instrument category.
	MusicalInstrumentName() string
	SetMusicalInstrumentName(value string)

	// Topic: Describing images

	// The ISO speed setting at the time the camera captured the image.
	ISOSpeed() foundation.NSNumber
	SetISOSpeed(value foundation.NSNumber)
	// The manufacturer of the device that captured the image.
	AcquisitionMake() string
	SetAcquisitionMake(value string)
	// The model of the device that captured the image.
	AcquisitionModel() string
	SetAcquisitionModel(value string)
	// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
	Aperture() foundation.NSNumber
	SetAperture(value foundation.NSNumber)
	// The number of bits per sample.
	BitsPerSample() foundation.NSNumber
	SetBitsPerSample(value foundation.NSNumber)
	// The owner of the camera that captured the image.
	CameraOwner() string
	SetCameraOwner(value string)
	// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
	ColorSpace() string
	SetColorSpace(value string)
	// A value that indicates if the camera used a flash to capture the image.
	FlashOn() foundation.NSNumber
	SetFlashOn(value foundation.NSNumber)
	// The actual focal length of the lens, in millimeters.
	FocalLength() foundation.NSNumber
	SetFocalLength(value foundation.NSNumber)
	// A value that indicates if the focal length is 35mm.
	FocalLength35mm() foundation.NSNumber
	SetFocalLength35mm(value foundation.NSNumber)
	// An array that contains the names of the various layers in the file.
	LayerNames() []string
	SetLayerNames(value []string)
	// The model of the lens that captured the image.
	LensModel() string
	SetLensModel(value string)
	// The orientation of the data.
	Orientation() foundation.NSNumber
	SetOrientation(value foundation.NSNumber)
	// The total number of pixels in the image.
	PixelCount() foundation.NSNumber
	SetPixelCount(value foundation.NSNumber)
	// The height of the item, such as image or video frame height, in pixels.
	PixelHeight() foundation.NSNumber
	SetPixelHeight(value foundation.NSNumber)
	// The width of the item, such as image or video frame width, in pixels.
	PixelWidth() foundation.NSNumber
	SetPixelWidth(value foundation.NSNumber)
	// The white balance setting when the camera captured the image.
	WhiteBalance() foundation.NSNumber
	SetWhiteBalance(value foundation.NSNumber)
	// The version of GPS Info IFD header that was used to generate the metadata for the image.
	EXIFGPSVersion() string
	SetEXIFGPSVersion(value string)
	// The version of the EXIF header that was used to generate the metadata for the image.
	EXIFVersion() string
	SetEXIFVersion(value string)
	// The mode the camera used for the exposure of the image.
	ExposureMode() foundation.NSNumber
	SetExposureMode(value foundation.NSNumber)
	// The class of the program the camera used to set exposure when capturing the image.
	ExposureProgram() string
	SetExposureProgram(value string)
	// The time that the lens was open during exposure, in seconds.
	ExposureTime() foundation.NSNumber
	SetExposureTime(value foundation.NSNumber)
	// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
	ExposureTimeString() string
	SetExposureTimeString(value string)
	// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
	FNumber() foundation.NSNumber
	SetFNumber(value foundation.NSNumber)
	// Indicates if the image file has an alpha channel.
	HasAlphaChannel() foundation.NSNumber
	SetHasAlphaChannel(value foundation.NSNumber)
	// The smallest F number of the lens.
	MaxAperture() foundation.NSNumber
	SetMaxAperture(value foundation.NSNumber)
	// The metering mode.
	MeteringMode() string
	SetMeteringMode(value string)
	// The name of the color profile the camera used for the image.
	ProfileName() string
	SetProfileName(value string)
	// A value that indicates if the camera used red-eye reduction when capturing the image.
	RedEyeOn() foundation.NSNumber
	SetRedEyeOn(value foundation.NSNumber)
	// The resolution height of the image, in DPI.
	ResolutionHeightDPI() foundation.NSNumber
	SetResolutionHeightDPI(value foundation.NSNumber)
	// The resolution width of the image, in DPI.
	ResolutionWidthDPI() foundation.NSNumber
	SetResolutionWidthDPI(value foundation.NSNumber)

	// Topic: Describing messages

	// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
	HTMLContentData() foundation.NSData
	SetHTMLContentData(value foundation.NSData)
	// An array of the canonical handles for the account with which the message is associated.
	AccountHandles() []string
	SetAccountHandles(value []string)
	// The unique identifier for the account with which the message is associated, if any.
	AccountIdentifier() string
	SetAccountIdentifier(value string)
	// An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Cc: field in an email message.
	AdditionalRecipients() []CSPerson
	SetAdditionalRecipients(value []CSPerson)
	// An array of addresses associated with the author of the message.
	AuthorAddresses() []string
	SetAuthorAddresses(value []string)
	// An array of email addresses associated with the author of the message.
	AuthorEmailAddresses() []string
	SetAuthorEmailAddresses(value []string)
	// An array of names representing the authors who have worked on the message.
	AuthorNames() []string
	SetAuthorNames(value []string)
	// An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the From: field in an item.
	Authors() []CSPerson
	SetAuthors(value []CSPerson)
	// An array of email addresses associated with the message.
	EmailAddresses() []string
	SetEmailAddresses(value []string)
	// A dictionary that contains all the headers of the message.
	EmailHeaders() foundation.INSDictionary
	SetEmailHeaders(value foundation.INSDictionary)
	// An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the Bcc: field in an email message.
	HiddenAdditionalRecipients() []CSPerson
	SetHiddenAdditionalRecipients(value []CSPerson)
	// An array of instant message addresses for the message.
	InstantMessageAddresses() []string
	SetInstantMessageAddresses(value []string)
	// A value that indicates if the message is likely to be considered junk.
	LikelyJunk() foundation.NSNumber
	SetLikelyJunk(value foundation.NSNumber)
	// An array of mailbox identifiers associated with the message.
	MailboxIdentifiers() []string
	SetMailboxIdentifiers(value []string)
	// An array of phone numbers associated with the message.
	PhoneNumbers() []string
	SetPhoneNumbers(value []string)
	// An array of [CSPerson](<https://developer.apple.com/documentation/CoreSpotlight/CSPerson>) objects representing the content of the To: field in an email message.
	PrimaryRecipients() []CSPerson
	SetPrimaryRecipients(value []CSPerson)
	// An array of addresses associated with the recipients of the message.
	RecipientAddresses() []string
	SetRecipientAddresses(value []string)
	// An array of email addresses associated with the recipient.
	RecipientEmailAddresses() []string
	SetRecipientEmailAddresses(value []string)
	// An array of names representing the recipients of this message.
	RecipientNames() []string
	SetRecipientNames(value []string)
	// The textual content of the message.
	TextContent() string
	SetTextContent(value string)

	// Topic: Describing containment

	// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
	ContainerDisplayName() string
	SetContainerDisplayName(value string)
	// The identifier of the container to which the item belongs.
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	// The order of the item within the container.
	ContainerOrder() foundation.NSNumber
	SetContainerOrder(value foundation.NSNumber)
	// The title of the container to which the item belongs.
	ContainerTitle() string
	SetContainerTitle(value string)

	// Topic: Describing supporting actions

	// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
	SupportsNavigation() foundation.NSNumber
	SetSupportsNavigation(value foundation.NSNumber)
	// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
	SupportsPhoneCall() foundation.NSNumber
	SetSupportsPhoneCall(value foundation.NSNumber)

	MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSSearchableItemAttributeSet) Init() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSearchableItemAttributeSet) Autorelease() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableItemAttributeSet creates a new CSSearchableItemAttributeSet instance.
func NewCSSearchableItemAttributeSet() CSSearchableItemAttributeSet {
	class := getCSSearchableItemAttributeSetClass()
	rv := objc.Send[CSSearchableItemAttributeSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(coder:)
func NewCSSearchableItemAttributeSetWithCoder(coder foundation.INSCoder) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSSearchableItemAttributeSetFromID(rv)
}

// Creates an attribute set for the specified content type.
//
// contentType: The type of the content. For example, [png] or [movie].
//
// # Return Value
//
// An attribute set that represents an item of the specified content type.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(contentType:)
//
// [movie]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/movie
// [png]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/png
func NewCSSearchableItemAttributeSetWithContentType(contentType uniformtypeidentifiers.UTType) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentType:"), contentType)
	return CSSearchableItemAttributeSetFromID(rv)
}

// Creates an attribute set for the specified content type.
//
// contentType: The type of the content. For example, [png] or [movie].
//
// # Return Value
//
// An attribute set that represents an item of the specified content type.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(contentType:)
//
// [movie]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/movie
// [png]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/png
func (c CSSearchableItemAttributeSet) InitWithContentType(contentType uniformtypeidentifiers.UTType) CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c.ID, objc.Sel("initWithContentType:"), contentType)
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(coder:)
func (c CSSearchableItemAttributeSet) InitWithCoder(coder foundation.INSCoder) CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Sets the value for a custom attribute key.
//
// value: The value of the custom attribute. Values must be common property list
// types, such as [NSString], [NSNumber], [NSNull], [NSData], or [NSDate], or
// an array of property list types.
//
// key: The custom attribute key.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/setValue(_:forCustomKey:)
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (c CSSearchableItemAttributeSet) SetValueForCustomKey(value foundation.NSSecureCoding, key ICSCustomAttributeKey) {
	objc.Send[objc.ID](c.ID, objc.Sel("setValue:forCustomKey:"), value, key)
}

// Returns the value associated with the specified custom attribute key.
//
// key: The custom attribute key.
//
// # Return Value
//
// The value associated with the custom attribute key.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/value(forCustomKey:)
func (c CSSearchableItemAttributeSet) ValueForCustomKey(key ICSCustomAttributeKey) foundation.NSSecureCoding {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("valueForCustomKey:"), key)
	return foundation.NSSecureCodingObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/move(from:)
func (c CSSearchableItemAttributeSet) MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("moveFrom:"), sourceAttributeSet)
}
func (c CSSearchableItemAttributeSet) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether the mail or messages content
// represents a prioritized item.
//
// # Discussion
//
// When the value of this property is `1`, Apple Intelligence identified this
// email or message content as needing priority classification.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isPriority
func (c CSSearchableItemAttributeSet) IsPriority() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isPriority"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A string that presents the Apple Intelligence summarization of the item.
//
// # Discussion
//
// This property represents the summary of the text that Apple Intelligence
// generated.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContentSummary
func (c CSSearchableItemAttributeSet) TextContentSummary() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("textContentSummary"))
	return foundation.NSStringFromID(rv).String()
}

// A string that represents the text the system transcribed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent
func (c CSSearchableItemAttributeSet) TranscribedTextContent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("transcribedTextContent"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetTranscribedTextContent(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTranscribedTextContent:"), objc.String(value))
}

// An array of identifiers that corresponds to data representations the
// delegate provides.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers
func (c CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("providerDataTypeIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProviderDataTypeIdentifiers:"), objectivec.StringSliceToNSArray(value))
}

// An array of identifiers that corresponds to file representations the
// delegate provides.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers
func (c CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("providerFileTypeIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProviderFileTypeIdentifiers:"), objectivec.StringSliceToNSArray(value))
}

// An array of identifiers that corresponds to in-place file representations
// the delegate provides.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers
func (c CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), objectivec.StringSliceToNSArray(value))
}

// An array of localized strings that represent alternate display names for
// the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames
func (c CSSearchableItemAttributeSet) AlternateNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("alternateNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAlternateNames(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlternateNames:"), objectivec.StringSliceToNSArray(value))
}

// The uniform type identifier (UTI) of the item.
//
// # Discussion
//
// To learn more about UTIs, see [UTType].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType
func (c CSSearchableItemAttributeSet) ContentType() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentType"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetContentType(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentType:"), objc.String(value))
}

// An attribute type that identifies a custom hierarchy of types to describe
// the attributes of your item.
//
// # Discussion
//
// For example, the [CSSearchableItemAttributeSet.ContentTypeTree] for an item
// whose [CSSearchableItemAttributeSet.ContentType] is `public.M3u()-playlist`
// should include `public.M3u()-playlist` and `public.Playlist()`.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree
func (c CSSearchableItemAttributeSet) ContentTypeTree() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contentTypeTree"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetContentTypeTree(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentTypeTree:"), objectivec.StringSliceToNSArray(value))
}

// The file URL of the content to index.
//
// # Discussion
//
// This is an optional property. An app that is also a client of iCloud Drive
// can set this property to allow Spotlight to deduplicate its searchable
// items against the iCloud Drive items. In this scenario, Spotlight does not
// display searchable items from iCloud Drive that have the same
// [CSSearchableItemAttributeSet.ContentURL] property.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL
func (c CSSearchableItemAttributeSet) ContentURL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetContentURL(value foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentURL:"), value)
}

// The local file URL of the thumbnail image for the item when Dark Mode is
// active.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL
func (c CSSearchableItemAttributeSet) DarkThumbnailURL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("darkThumbnailURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetDarkThumbnailURL(value foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setDarkThumbnailURL:"), value)
}

// A localized string that contains the name of the item, suitable to display
// in the user interface.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c CSSearchableItemAttributeSet) DisplayName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetDisplayName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDisplayName:"), objc.String(value))
}

// An array of keywords associated with the item, such as work, birthday,
// important, and so on.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords
func (c CSSearchableItemAttributeSet) Keywords() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("keywords"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetKeywords(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeywords:"), objectivec.StringSliceToNSArray(value))
}

// The date on which the last metadata attribute was changed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate
func (c CSSearchableItemAttributeSet) MetadataModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadataModificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetMetadataModificationDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadataModificationDate:"), value)
}

// The complete path to the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path
func (c CSSearchableItemAttributeSet) Path() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetPath(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPath:"), objc.String(value))
}

// A number that indicates the relative importance of the item among other
// items from the app.
//
// # Discussion
//
// Core Spotlight uses this value to distinguish between similar items from an
// app’s indexed content. Set this property to an integer value between zero
// and one hundred. Higher values indicate an item is more relevant, and
// Spotlight may display it more prominently.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint
func (c CSSearchableItemAttributeSet) RankingHint() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rankingHint"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetRankingHint(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setRankingHint:"), value)
}

// The unique identifier for the item to which the activity is related.
//
// # Discussion
//
// If you’re using both [NSUserActivity] and Core Spotlight APIs to index
// the same item, set this property in the activity to specify the unique
// identifier of the Core Spotlight item to which the activity is related, and
// to avoid displaying duplicate results in Spotlight.
//
// If the unique identifier to which the activity is related hasn’t already
// been indexed with Core Spotlight, the activity won’t be indexed. Note
// that when the item is deleted, the related activity is also deleted, unlike
// the behavior of [CSSearchableItemAttributeSet.WeakRelatedUniqueIdentifier].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier
//
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
func (c CSSearchableItemAttributeSet) RelatedUniqueIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("relatedUniqueIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRelatedUniqueIdentifier:"), objc.String(value))
}

// Image data that represents the thumbnail of the item.
//
// # Discussion
//
// This property is optional. For some guidance on creating a thumbnail image,
// see [Enhance Your Search Results].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData
//
// [Enhance Your Search Results]: https://developer.apple.com/library/archive/documentation/General/Conceptual/AppSearch/SearchUserExperience.html#//apple_ref/doc/uid/TP40016308-CH11
func (c CSSearchableItemAttributeSet) ThumbnailData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("thumbnailData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetThumbnailData(value foundation.NSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setThumbnailData:"), value)
}

// The local file URL of the thumbnail image for the item.
//
// # Discussion
//
// This property is optional. Note that the URL needs to point to a local
// file; it is not a web URL.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL
func (c CSSearchableItemAttributeSet) ThumbnailURL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("thumbnailURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetThumbnailURL(value foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setThumbnailURL:"), value)
}

// The title of the item.
//
// # Discussion
//
// An item title might be the title of a document or MP3 file or the subject
// of an email message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c CSSearchableItemAttributeSet) Title() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTitle:"), objc.String(value))
}

// An identifier that represents the domain or owner of the item.
//
// # Discussion
//
// Specify a domain identifier to group items together and make it easier to
// delete them from the index. For example, to delete a user activity, set
// this property on the [contentAttributeSet] property of the [NSUserActivity]
// object and then call
// [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]
// on your app’s index.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier
//
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [contentAttributeSet]: https://developer.apple.com/documentation/Foundation/NSUserActivity/contentAttributeSet
func (c CSSearchableItemAttributeSet) DomainIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("domainIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetDomainIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}

// The unique identifier for the item to which the activity is related, but
// not linked.
//
// # Discussion
//
// Unlike the similar [CSSearchableItemAttributeSet.RelatedUniqueIdentifier]
// property, this property does not link the lifetime of the item to the
// lifetime of the activity. In particular, deleting the item does not delete
// the activity.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier
func (c CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), objc.String(value))
}

// A class of entity for which the item is intended or useful.
//
// # Discussion
//
// A class of entity may be determined by the creator or the publisher of the
// item or by a third party.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences
func (c CSSearchableItemAttributeSet) Audiences() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("audiences"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAudiences(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudiences:"), objectivec.StringSliceToNSArray(value))
}

// A description of the item’s content.
//
// # Discussion
//
// A description may consist of an abstract, table of contents, reference to a
// graphical representation of content, a free-text account of the content, or
// other matter.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription
func (c CSSearchableItemAttributeSet) ContentDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetContentDescription(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentDescription:"), objc.String(value))
}

// The name of the app that created the content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator
func (c CSSearchableItemAttributeSet) Creator() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("creator"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetCreator(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCreator:"), objc.String(value))
}

// The name of the apps that converted the original content into a PDF stream.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c CSSearchableItemAttributeSet) EncodingApplications() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("encodingApplications"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetEncodingApplications(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEncodingApplications:"), objectivec.StringSliceToNSArray(value))
}

// The size of the document file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize
func (c CSSearchableItemAttributeSet) FileSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetFileSize(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFileSize:"), value)
}

// An array of font names the document uses.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames
func (c CSSearchableItemAttributeSet) FontNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("fontNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetFontNames(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setFontNames:"), objectivec.StringSliceToNSArray(value))
}

// A formal identifier that references the document the item represents.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier
func (c CSSearchableItemAttributeSet) Identifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// A description of the kind of document the item represents.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind
func (c CSSearchableItemAttributeSet) Kind() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("kind"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetKind(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setKind:"), objc.String(value))
}

// The number of pages in the document.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount
func (c CSSearchableItemAttributeSet) PageCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pageCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPageCount(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPageCount:"), value)
}

// The height of the document page, in points (72 points per inch).
//
// # Discussion
//
// For a PDF document, this property specifies the height of the first page
// only; other pages in a PDF document may have different heights.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight
func (c CSSearchableItemAttributeSet) PageHeight() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pageHeight"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPageHeight(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPageHeight:"), value)
}

// The width of the document page, in points (72 points per inch).
//
// # Discussion
//
// For a PDF document, this property specifies the width of the first page
// only; other pages in a PDF document may have different widths.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth
func (c CSSearchableItemAttributeSet) PageWidth() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pageWidth"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPageWidth(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPageWidth:"), value)
}

// The security method (a type of encryption) that protects the document file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod
func (c CSSearchableItemAttributeSet) SecurityMethod() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("securityMethod"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetSecurityMethod(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSecurityMethod:"), objc.String(value))
}

// The subject of the document.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject
func (c CSSearchableItemAttributeSet) Subject() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subject"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetSubject(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubject:"), objc.String(value))
}

// The theme of the document.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme
func (c CSSearchableItemAttributeSet) Theme() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("theme"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetTheme(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTheme:"), objc.String(value))
}

// A value that indicates the user created the item.
//
// # Discussion
//
// Examples of items to set this property on include notes and documents that
// the user creates or modifies.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated
func (c CSSearchableItemAttributeSet) UserCreated() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isUserCreated"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetUserCreated(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserCreated:"), value)
}

// A value that indicates the user selected the item.
//
// # Discussion
//
// Examples of items to set this property on include media content the user
// downloads and websites or news articles the user bookmarks.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated
func (c CSSearchableItemAttributeSet) UserCurated() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isUserCurated"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetUserCurated(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserCurated:"), value)
}

// A value that indicates the user purchased or owns the item.
//
// # Discussion
//
// Examples of items to set this property on include songs and movies the user
// purchases.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned
func (c CSSearchableItemAttributeSet) UserOwned() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isUserOwned"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetUserOwned(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserOwned:"), value)
}

// A value that indicates if the event covers an entire day.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay
func (c CSSearchableItemAttributeSet) AllDay() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("allDay"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAllDay(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllDay:"), value)
}

// The date on which the item was completed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate
func (c CSSearchableItemAttributeSet) CompletionDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("completionDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetCompletionDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompletionDate:"), value)
}

// The date on which the item is due.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate
func (c CSSearchableItemAttributeSet) DueDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dueDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetDueDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDueDate:"), value)
}

// The end date for the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c CSSearchableItemAttributeSet) EndDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("endDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetEndDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setEndDate:"), value)
}

// An array of important dates associated with the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates
func (c CSSearchableItemAttributeSet) ImportantDates() []foundation.NSDate {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("importantDates"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSDate {
		return foundation.NSDateFromID(id)
	})
}
func (c CSSearchableItemAttributeSet) SetImportantDates(value []foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setImportantDates:"), objectivec.IObjectSliceToNSArray(value))
}

// The start date for the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate
func (c CSSearchableItemAttributeSet) StartDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("startDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetStartDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setStartDate:"), value)
}

// The altitude of the item in meters above sea level, expressed using the
// WGS84 datum.
//
// # Discussion
//
// Negative values lie below sea level.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude
func (c CSSearchableItemAttributeSet) Altitude() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("altitude"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAltitude(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAltitude:"), value)
}

// The city of the item’s origin according to guidelines that the provider
// establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city
func (c CSSearchableItemAttributeSet) City() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("city"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetCity(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCity:"), objc.String(value))
}

// The full, publishable name of the country or region in which the
// intellectual property of the item was created, according to guidelines the
// provider establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country
func (c CSSearchableItemAttributeSet) Country() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("country"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetCountry(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCountry:"), objc.String(value))
}

// Information about the GPS area.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation
func (c CSSearchableItemAttributeSet) GPSAreaInformation() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSAreaInformation"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGPSAreaInformation(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSAreaInformation:"), objc.String(value))
}

// The GPS dilution of precision value.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c CSSearchableItemAttributeSet) GPSDOP() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDOP"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDOP(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDOP:"), value)
}

// The date and time related to the GPS value.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp
func (c CSSearchableItemAttributeSet) GPSDateStamp() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDateStamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDateStamp(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDateStamp:"), value)
}

// The bearing to the destination point.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing
func (c CSSearchableItemAttributeSet) GPSDestBearing() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDestBearing"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDestBearing(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDestBearing:"), value)
}

// The distance to the destination point.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance
func (c CSSearchableItemAttributeSet) GPSDestDistance() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDestDistance"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDestDistance(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDestDistance:"), value)
}

// The latitude of the destination point.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude
func (c CSSearchableItemAttributeSet) GPSDestLatitude() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDestLatitude"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDestLatitude(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDestLatitude:"), value)
}

// The longitude of the destination point.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude
func (c CSSearchableItemAttributeSet) GPSDestLongitude() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDestLongitude"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDestLongitude(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDestLongitude:"), value)
}

// The differential correction applied to the GPS receiver.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental
func (c CSSearchableItemAttributeSet) GPSDifferental() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSDifferental"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSDifferental(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSDifferental:"), value)
}

// The geodetic data that the GPS receiver uses.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum
func (c CSSearchableItemAttributeSet) GPSMapDatum() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSMapDatum"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGPSMapDatum(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSMapDatum:"), objc.String(value))
}

// The measurement precision mode in use by the GPS receiver.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode
func (c CSSearchableItemAttributeSet) GPSMeasureMode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSMeasureMode"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGPSMeasureMode(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSMeasureMode:"), objc.String(value))
}

// The location finding method that the GPS receiver uses.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod
func (c CSSearchableItemAttributeSet) GPSProcessingMethod() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSProcessingMethod"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGPSProcessingMethod(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSProcessingMethod:"), objc.String(value))
}

// The status of the GPS receiver.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus
func (c CSSearchableItemAttributeSet) GPSStatus() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSStatus"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGPSStatus(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSStatus:"), objc.String(value))
}

// The direction of travel of the item in degrees from true north.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack
func (c CSSearchableItemAttributeSet) GPSTrack() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("GPSTrack"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGPSTrack(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGPSTrack:"), value)
}

// A publishable string that provides a synopsis of the contents of the item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline
func (c CSSearchableItemAttributeSet) Headline() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("headline"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetHeadline(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setHeadline:"), objc.String(value))
}

// The direction of the item’s image in degrees from true north.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection
func (c CSSearchableItemAttributeSet) ImageDirection() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("imageDirection"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetImageDirection(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setImageDirection:"), value)
}

// Instructions that concern the use of the item, such as an embargo or
// warning.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions
func (c CSSearchableItemAttributeSet) Instructions() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("instructions"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetInstructions(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setInstructions:"), objc.String(value))
}

// The latitude of the item, in degrees north of the equator, expressed using
// the WGS84 datum.
//
// # Discussion
//
// Negative values lie south of the equator.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude
func (c CSSearchableItemAttributeSet) Latitude() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("latitude"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetLatitude(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setLatitude:"), value)
}

// The longitude of the item, in degrees east of the prime meridian, expressed
// using the WGS84 datum.
//
// # Discussion
//
// Negative values lie west of the prime meridian.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude
func (c CSSearchableItemAttributeSet) Longitude() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("longitude"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetLongitude(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setLongitude:"), value)
}

// The name of the location or point of interest associated with the item.
//
// # Discussion
//
// The name may be user-provided.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation
func (c CSSearchableItemAttributeSet) NamedLocation() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("namedLocation"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetNamedLocation(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setNamedLocation:"), objc.String(value))
}

// The speed of the item, in kilometers per hour.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed
func (c CSSearchableItemAttributeSet) Speed() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("speed"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetSpeed(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setSpeed:"), value)
}

// The province or state of origin according to guidelines the provider
// establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince
func (c CSSearchableItemAttributeSet) StateOrProvince() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stateOrProvince"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetStateOrProvince(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setStateOrProvince:"), objc.String(value))
}

// The timestamp on the item.
//
// # Discussion
//
// The value of this property is generally used to specify the time at which
// the event captured by the item took place.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp
func (c CSSearchableItemAttributeSet) Timestamp() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("timestamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetTimestamp(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimestamp:"), value)
}

// The fully formatted address of the item, received from MapKit.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress
func (c CSSearchableItemAttributeSet) FullyFormattedAddress() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fullyFormattedAddress"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetFullyFormattedAddress(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setFullyFormattedAddress:"), objc.String(value))
}

// The postal code for the item according to guidelines the provider
// establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode
func (c CSSearchableItemAttributeSet) PostalCode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("postalCode"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetPostalCode(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPostalCode:"), objc.String(value))
}

// The sublocation, such as a street number, for the item according to
// guidelines the provider establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare
func (c CSSearchableItemAttributeSet) SubThoroughfare() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subThoroughfare"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetSubThoroughfare(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubThoroughfare:"), objc.String(value))
}

// The thoroughfare, such as a street name, associated with the location for
// the item according to guidelines the provider establishes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare
func (c CSSearchableItemAttributeSet) Thoroughfare() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("thoroughfare"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetThoroughfare(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setThoroughfare:"), objc.String(value))
}

// A comment related to the media file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment
func (c CSSearchableItemAttributeSet) Comment() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("comment"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetComment(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setComment:"), objc.String(value))
}

// The creation date of an edited or optimized version of the song or
// composition.
//
// # Discussion
//
// This property is supplementary to
// [CSSearchableItemAttributeSet.RecordingDate], which indicates the original
// recording date of the song or composition.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate
func (c CSSearchableItemAttributeSet) ContentCreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentCreationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetContentCreationDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentCreationDate:"), value)
}

// The date on which the contents of the file was last modified.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c CSSearchableItemAttributeSet) ContentModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentModificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetContentModificationDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentModificationDate:"), value)
}

// An array of sources from which the media was obtained.
//
// # Discussion
//
// The string values in this property might include the URL of the website
// from which the file was downloaded or information that describes the email
// to which the file was attached.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources
func (c CSSearchableItemAttributeSet) ContentSources() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contentSources"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetContentSources(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentSources:"), objectivec.StringSliceToNSArray(value))
}

// The copyright date of the content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright
func (c CSSearchableItemAttributeSet) Copyright() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("copyright"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetCopyright(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCopyright:"), objc.String(value))
}

// The most recent date on which the file was downloaded or received.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate
func (c CSSearchableItemAttributeSet) DownloadedDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("downloadedDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetDownloadedDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDownloadedDate:"), value)
}

// A list of editors who have worked on the file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors
func (c CSSearchableItemAttributeSet) Editors() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("editors"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetEditors(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEditors:"), objectivec.StringSliceToNSArray(value))
}

// The date on which the file was last used.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate
func (c CSSearchableItemAttributeSet) LastUsedDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lastUsedDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetLastUsedDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setLastUsedDate:"), value)
}

// A list of people who are visible in an image or movie or written about in a
// document.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants
func (c CSSearchableItemAttributeSet) Participants() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("participants"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetParticipants(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setParticipants:"), objectivec.StringSliceToNSArray(value))
}

// A list of projects of which this file is a part.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects
func (c CSSearchableItemAttributeSet) Projects() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("projects"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetProjects(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProjects:"), objectivec.StringSliceToNSArray(value))
}

// The date on which the item was moved into its current location.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate
func (c CSSearchableItemAttributeSet) AddedDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("addedDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAddedDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setAddedDate:"), value)
}

// The codecs used to encode/decode the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs
func (c CSSearchableItemAttributeSet) Codecs() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("codecs"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetCodecs(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCodecs:"), objectivec.StringSliceToNSArray(value))
}

// A list of contacts who are associated with the content in some way, not
// including the author.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords
func (c CSSearchableItemAttributeSet) ContactKeywords() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contactKeywords"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetContactKeywords(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContactKeywords:"), objectivec.StringSliceToNSArray(value))
}

// The delivery type of the file.
//
// # Discussion
//
// The value of this property is 0 for fast start and 1 for RTSP.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType
func (c CSSearchableItemAttributeSet) DeliveryType() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("deliveryType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetDeliveryType(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setDeliveryType:"), value)
}

// The duration (if appropriate) of the content of the file, in seconds.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration
func (c CSSearchableItemAttributeSet) Duration() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("duration"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetDuration(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setDuration:"), value)
}

// The media types present in the content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes
func (c CSSearchableItemAttributeSet) MediaTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("mediaTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetMediaTypes(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMediaTypes:"), objectivec.StringSliceToNSArray(value))
}

// A list of companies or organizations that created the content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations
func (c CSSearchableItemAttributeSet) Organizations() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("organizations"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetOrganizations(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setOrganizations:"), objectivec.StringSliceToNSArray(value))
}

// A value that indicates if the content is prepared for streaming.
//
// # Discussion
//
// When the value of this property is 0, the content is not prepared for
// streaming; when the value is 1, it is prepared for streaming.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable
func (c CSSearchableItemAttributeSet) Streamable() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isStreamable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetStreamable(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setStreamable:"), value)
}

// The total bit rate of the media, combining audio and video.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate
func (c CSSearchableItemAttributeSet) TotalBitRate() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("totalBitRate"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetTotalBitRate(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setTotalBitRate:"), value)
}

// The audio bit rate of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate
func (c CSSearchableItemAttributeSet) AudioBitRate() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioBitRate"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAudioBitRate(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioBitRate:"), value)
}

// A version string associated with the file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version
func (c CSSearchableItemAttributeSet) Version() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setVersion:"), objc.String(value))
}

// The video bit rate of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate
func (c CSSearchableItemAttributeSet) VideoBitRate() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoBitRate"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetVideoBitRate(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoBitRate:"), value)
}

// A list of people, organizations, or services that made contributions to the
// media content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors
func (c CSSearchableItemAttributeSet) Contributors() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contributors"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetContributors(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContributors:"), objectivec.StringSliceToNSArray(value))
}

// A list of the included languages for the intellectual content of the media.
//
// # Discussion
//
// See BCP 47 for best practices in using language tags.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages
func (c CSSearchableItemAttributeSet) Languages() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("languages"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetLanguages(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLanguages:"), objectivec.StringSliceToNSArray(value))
}

// A list of people, organizations, services, or other entities responsible
// for making the media available.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers
func (c CSSearchableItemAttributeSet) Publishers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("publishers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetPublishers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPublishers:"), objectivec.StringSliceToNSArray(value))
}

// A link to information about the rights held in and over the media.
//
// # Discussion
//
// Typically, this property contains a rights management statement for the
// media, or references a service that provides such information. Rights
// information often encompasses Intellectual Property Rights (IPR),
// copyright, and various property rights. If the `rights` property is absent,
// no assumptions can be made about the status of these and other rights with
// respect to the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights
func (c CSSearchableItemAttributeSet) Rights() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rights"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetRights(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRights:"), objc.String(value))
}

// Indicates the role of the content creator.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role
func (c CSSearchableItemAttributeSet) Role() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("role"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetRole(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRole:"), objc.String(value))
}

// A value that indicates if the media contains explicit content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating
func (c CSSearchableItemAttributeSet) ContentRating() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentRating"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetContentRating(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentRating:"), value)
}

// A list of descriptors that specify the extent or scope of the media.
//
// # Discussion
//
// The string values in this property typically include a location (such as a
// place name or geographic coordinates), a temporal period (such as a period
// label, date, or date range), or a jurisdiction (such as a named
// administrative entity).
//
// It’s recommended that you select a value from a controlled vocabulary,
// and that when you need to specify a place or time period, you use a name
// instead of a numeric identifier, such as a set of coordinates or a date
// range.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage
func (c CSSearchableItemAttributeSet) Coverage() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("coverage"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetCoverage(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCoverage:"), objectivec.StringSliceToNSArray(value))
}

// The name of the director of the media (for example, a movie director).
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director
func (c CSSearchableItemAttributeSet) Director() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("director"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetDirector(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDirector:"), objc.String(value))
}

// The genre of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre
func (c CSSearchableItemAttributeSet) Genre() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("genre"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetGenre(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setGenre:"), objc.String(value))
}

// Information about the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information
func (c CSSearchableItemAttributeSet) Information() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("information"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetInformation(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setInformation:"), objc.String(value))
}

// A value that indicates if the media is local.
//
// # Discussion
//
// When the value of this property is 1, the media is local; when the value is
// 0, the media is not local.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local
func (c CSSearchableItemAttributeSet) Local() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isLocal"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetLocal(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocal:"), value)
}

// The original format of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat
func (c CSSearchableItemAttributeSet) OriginalFormat() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("originalFormat"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetOriginalFormat(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setOriginalFormat:"), objc.String(value))
}

// The original source of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource
func (c CSSearchableItemAttributeSet) OriginalSource() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("originalSource"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetOriginalSource(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setOriginalSource:"), objc.String(value))
}

// A list of performers in the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers
func (c CSSearchableItemAttributeSet) Performers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("performers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetPerformers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerformers:"), objectivec.StringSliceToNSArray(value))
}

// A user-supplied play count for the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount
func (c CSSearchableItemAttributeSet) PlayCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("playCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPlayCount(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPlayCount:"), value)
}

// The producer of the content.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer
func (c CSSearchableItemAttributeSet) Producer() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("producer"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetProducer(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProducer:"), objc.String(value))
}

// The user-supplied rating of the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating
func (c CSSearchableItemAttributeSet) Rating() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rating"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetRating(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setRating:"), value)
}

// A description of the rating.
//
// # Discussion
//
// For example, the description might include the number of reviewers who
// provided ratings.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription
func (c CSSearchableItemAttributeSet) RatingDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ratingDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetRatingDescription(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRatingDescription:"), objc.String(value))
}

// The URL associated with the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url
func (c CSSearchableItemAttributeSet) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetURL(value foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setURL:"), value)
}

// The title for a collection of audio media.
//
// # Discussion
//
// The value of this property is analogous to the title of a record album or
// photo album.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album
func (c CSSearchableItemAttributeSet) Album() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("album"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetAlbum(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlbum:"), objc.String(value))
}

// The artist associated with the media.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist
func (c CSSearchableItemAttributeSet) Artist() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("artist"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetArtist(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setArtist:"), objc.String(value))
}

// The number of channels in the audio data that the file contains.
//
// # Discussion
//
// The value of this property represents only the number of discreet channels
// of audio data found in the file. It does not indicate any audio data
// configuration related to a user’s speaker setup.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount
func (c CSSearchableItemAttributeSet) AudioChannelCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioChannelCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAudioChannelCount(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioChannelCount:"), value)
}

// The name of the application that encoded the data the audio file contains.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication
func (c CSSearchableItemAttributeSet) AudioEncodingApplication() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioEncodingApplication"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetAudioEncodingApplication(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioEncodingApplication:"), objc.String(value))
}

// The sample rate of the audio data the file contains, as a float value
// representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate
func (c CSSearchableItemAttributeSet) AudioSampleRate() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioSampleRate"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAudioSampleRate(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioSampleRate:"), value)
}

// The track number of a song or audio composition when part of an album.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber
func (c CSSearchableItemAttributeSet) AudioTrackNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioTrackNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAudioTrackNumber(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioTrackNumber:"), value)
}

// The composer of the song or audio composition that the audio file contains.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/composer
func (c CSSearchableItemAttributeSet) Composer() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("composer"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetComposer(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setComposer:"), objc.String(value))
}

// The musical key of the song or audio composition that the file contains,
// such as C, Dm, or F#m.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature
func (c CSSearchableItemAttributeSet) KeySignature() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keySignature"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetKeySignature(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setKeySignature:"), objc.String(value))
}

// The lyricist or text writer for the song or audio composition that the file
// contains.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist
func (c CSSearchableItemAttributeSet) Lyricist() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lyricist"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetLyricist(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLyricist:"), objc.String(value))
}

// The musical genre of the song or audio composition that the file contains,
// such as jazz, pop, rock, or classical.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre
func (c CSSearchableItemAttributeSet) MusicalGenre() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("musicalGenre"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetMusicalGenre(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMusicalGenre:"), objc.String(value))
}

// The recording date of the song or composition.
//
// # Discussion
//
// This property contains the original recording date of the song or
// composition and is supplementary to
// [CSSearchableItemAttributeSet.ContentCreationDate], which indicates the
// date of an edited or optimized version.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate
func (c CSSearchableItemAttributeSet) RecordingDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordingDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetRecordingDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordingDate:"), value)
}

// The tempo of the music that the audio file contains, in beats per minute.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo
func (c CSSearchableItemAttributeSet) Tempo() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("tempo"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetTempo(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setTempo:"), value)
}

// The time signature of the musical composition that the audio or MIDI file
// contains, in a string, such as “4/4” or “7/8”.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature
func (c CSSearchableItemAttributeSet) TimeSignature() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("timeSignature"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetTimeSignature(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeSignature:"), objc.String(value))
}

// A value that indicates whether the MIDI sequence the file contains is set
// up for use with a general MIDI device.
//
// # Discussion
//
// When the value of this property is 1, the MIDI sequence in the file is set
// up for use with a general MIDI device; when the value is 0, it is not.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence
func (c CSSearchableItemAttributeSet) GeneralMIDISequence() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isGeneralMIDISequence"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetGeneralMIDISequence(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setGeneralMIDISequence:"), value)
}

// The category of the instrument associated with the audio file.
//
// # Discussion
//
// A file should specify an associated instrument (note that you can use
// “Other Instrument” to specify an unknown instrument). In some
// categories, you can use instrument names to provide a more detailed
// instrument definition. For example, the Keyboards category lets you include
// instrument names such as Piano or Organ.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory
func (c CSSearchableItemAttributeSet) MusicalInstrumentCategory() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("musicalInstrumentCategory"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMusicalInstrumentCategory:"), objc.String(value))
}

// The name of an instrument within the context of an instrument category.
//
// # Discussion
//
// For some instrument categories, such as Percussion and Keyboards, you can
// specify an instrument name.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName
func (c CSSearchableItemAttributeSet) MusicalInstrumentName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("musicalInstrumentName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetMusicalInstrumentName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMusicalInstrumentName:"), objc.String(value))
}

// The ISO speed setting at the time the camera captured the image.
//
// # Discussion
//
// Typical ISO speed values are 100, 200, 400, and so on.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed
func (c CSSearchableItemAttributeSet) ISOSpeed() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ISOSpeed"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetISOSpeed(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setISOSpeed:"), value)
}

// The manufacturer of the device that captured the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake
func (c CSSearchableItemAttributeSet) AcquisitionMake() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("acquisitionMake"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetAcquisitionMake(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAcquisitionMake:"), objc.String(value))
}

// The model of the device that captured the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel
func (c CSSearchableItemAttributeSet) AcquisitionModel() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("acquisitionModel"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetAcquisitionModel(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAcquisitionModel:"), objc.String(value))
}

// The size of the lens aperture at the time the camera captured the image, as
// a log-scale APEX value.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture
func (c CSSearchableItemAttributeSet) Aperture() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("aperture"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetAperture(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setAperture:"), value)
}

// The number of bits per sample.
//
// # Discussion
//
// The value of this property can represent the bit depth of an image (such as
// 8-bit or 16-bit) or the bit depth per audio sample of uncompressed audio
// data (such as 8, 16, 24, 32, 64, and so on).
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample
func (c CSSearchableItemAttributeSet) BitsPerSample() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("bitsPerSample"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetBitsPerSample(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setBitsPerSample:"), value)
}

// The owner of the camera that captured the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner
func (c CSSearchableItemAttributeSet) CameraOwner() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cameraOwner"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetCameraOwner(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCameraOwner:"), objc.String(value))
}

// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace
func (c CSSearchableItemAttributeSet) ColorSpace() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorSpace"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetColorSpace(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setColorSpace:"), objc.String(value))
}

// A value that indicates if the camera used a flash to capture the image.
//
// # Discussion
//
// The value of this property is 1 if the flash is on; otherwise, 0.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn
func (c CSSearchableItemAttributeSet) FlashOn() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isFlashOn"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetFlashOn(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlashOn:"), value)
}

// The actual focal length of the lens, in millimeters.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength
func (c CSSearchableItemAttributeSet) FocalLength() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("focalLength"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetFocalLength(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocalLength:"), value)
}

// A value that indicates if the focal length is 35mm.
//
// # Discussion
//
// The value of this property is 1 if the focal length is 35mm; otherwise, 0.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm
func (c CSSearchableItemAttributeSet) FocalLength35mm() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isFocalLength35mm"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetFocalLength35mm(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocalLength35mm:"), value)
}

// An array that contains the names of the various layers in the file.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames
func (c CSSearchableItemAttributeSet) LayerNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("layerNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetLayerNames(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLayerNames:"), objectivec.StringSliceToNSArray(value))
}

// The model of the lens that captured the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel
func (c CSSearchableItemAttributeSet) LensModel() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lensModel"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetLensModel(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLensModel:"), objc.String(value))
}

// The orientation of the data.
//
// # Discussion
//
// When the value of this property is 1, the orientation is portrait; when the
// value is 0, the orientation is landscape.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation
func (c CSSearchableItemAttributeSet) Orientation() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("orientation"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetOrientation(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setOrientation:"), value)
}

// The total number of pixels in the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount
func (c CSSearchableItemAttributeSet) PixelCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pixelCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPixelCount(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPixelCount:"), value)
}

// The height of the item, such as image or video frame height, in pixels.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight
func (c CSSearchableItemAttributeSet) PixelHeight() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pixelHeight"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPixelHeight(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPixelHeight:"), value)
}

// The width of the item, such as image or video frame width, in pixels.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth
func (c CSSearchableItemAttributeSet) PixelWidth() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pixelWidth"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetPixelWidth(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setPixelWidth:"), value)
}

// The white balance setting when the camera captured the image.
//
// # Discussion
//
// The value of this property is 0 for auto and 1 for manual.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance
func (c CSSearchableItemAttributeSet) WhiteBalance() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("whiteBalance"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetWhiteBalance(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setWhiteBalance:"), value)
}

// The version of GPS Info IFD header that was used to generate the metadata
// for the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion
func (c CSSearchableItemAttributeSet) EXIFGPSVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("EXIFGPSVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetEXIFGPSVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEXIFGPSVersion:"), objc.String(value))
}

// The version of the EXIF header that was used to generate the metadata for
// the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion
func (c CSSearchableItemAttributeSet) EXIFVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("EXIFVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetEXIFVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEXIFVersion:"), objc.String(value))
}

// The mode the camera used for the exposure of the image.
//
// # Discussion
//
// This property can have the following values:
//
// - 0 — auto exposure
// - 1 — manual
// - 2 — auto bracket
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode
func (c CSSearchableItemAttributeSet) ExposureMode() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exposureMode"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetExposureMode(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureMode:"), value)
}

// The class of the program the camera used to set exposure when capturing the
// image.
//
// # Discussion
//
// Can include manual, normal, aperture priority, and so on.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram
func (c CSSearchableItemAttributeSet) ExposureProgram() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exposureProgram"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetExposureProgram(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureProgram:"), objc.String(value))
}

// The time that the lens was open during exposure, in seconds.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime
func (c CSSearchableItemAttributeSet) ExposureTime() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exposureTime"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetExposureTime(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureTime:"), value)
}

// The time that the lens was open during exposure, in a string, such as
// “1/250 seconds”.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString
func (c CSSearchableItemAttributeSet) ExposureTimeString() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("exposureTimeString"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetExposureTimeString(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureTimeString:"), objc.String(value))
}

// The focal length of the lens, divided by the diameter of the aperture when
// the camera captured the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber
func (c CSSearchableItemAttributeSet) FNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fNumber"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetFNumber(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setFNumber:"), value)
}

// Indicates if the image file has an alpha channel.
//
// # Discussion
//
// When the value of this property is 0, no alpha channel is used; when the
// value is 1, an alpha channel is used.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel
func (c CSSearchableItemAttributeSet) HasAlphaChannel() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("hasAlphaChannel"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetHasAlphaChannel(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasAlphaChannel:"), value)
}

// The smallest F number of the lens.
//
// # Discussion
//
// The unit of the F number is the APEX value, which is typically in the range
// of 00.00 to 99.99.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture
func (c CSSearchableItemAttributeSet) MaxAperture() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("maxAperture"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetMaxAperture(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxAperture:"), value)
}

// The metering mode.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode
func (c CSSearchableItemAttributeSet) MeteringMode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("meteringMode"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetMeteringMode(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMeteringMode:"), objc.String(value))
}

// The name of the color profile the camera used for the image.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName
func (c CSSearchableItemAttributeSet) ProfileName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("profileName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetProfileName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setProfileName:"), objc.String(value))
}

// A value that indicates if the camera used red-eye reduction when capturing
// the image.
//
// # Discussion
//
// When the value of this property is 0, no red-eye reduction was used; when
// the value is 1, red-eye reduction was used.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn
func (c CSSearchableItemAttributeSet) RedEyeOn() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isRedEyeOn"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetRedEyeOn(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setRedEyeOn:"), value)
}

// The resolution height of the image, in DPI.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI
func (c CSSearchableItemAttributeSet) ResolutionHeightDPI() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resolutionHeightDPI"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetResolutionHeightDPI(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setResolutionHeightDPI:"), value)
}

// The resolution width of the image, in DPI.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI
func (c CSSearchableItemAttributeSet) ResolutionWidthDPI() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resolutionWidthDPI"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetResolutionWidthDPI(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setResolutionWidthDPI:"), value)
}

// The HTML content of the document encoded as an NSData object representing a
// UTF-8 encoded string.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData
func (c CSSearchableItemAttributeSet) HTMLContentData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("HTMLContentData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetHTMLContentData(value foundation.NSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setHTMLContentData:"), value)
}

// An array of the canonical handles for the account with which the message is
// associated.
//
// # Discussion
//
// Account handles can include chat handles, email addresses, phone numbers,
// and so on.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles
func (c CSSearchableItemAttributeSet) AccountHandles() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("accountHandles"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAccountHandles(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccountHandles:"), objectivec.StringSliceToNSArray(value))
}

// The unique identifier for the account with which the message is associated,
// if any.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier
func (c CSSearchableItemAttributeSet) AccountIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("accountIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetAccountIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccountIdentifier:"), objc.String(value))
}

// An array of [CSPerson] objects representing the content of the Cc: field in
// an email message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/additionalRecipients
func (c CSSearchableItemAttributeSet) AdditionalRecipients() []CSPerson {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("additionalRecipients"))
	return objc.ConvertSlice(rv, func(id objc.ID) CSPerson {
		return CSPersonFromID(id)
	})
}
func (c CSSearchableItemAttributeSet) SetAdditionalRecipients(value []CSPerson) {
	objc.Send[struct{}](c.ID, objc.Sel("setAdditionalRecipients:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of addresses associated with the author of the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses
func (c CSSearchableItemAttributeSet) AuthorAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("authorAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAuthorAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAuthorAddresses:"), objectivec.StringSliceToNSArray(value))
}

// An array of email addresses associated with the author of the message.
//
// # Discussion
//
// The contents of this property is not human-readable.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses
func (c CSSearchableItemAttributeSet) AuthorEmailAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("authorEmailAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAuthorEmailAddresses:"), objectivec.StringSliceToNSArray(value))
}

// An array of names representing the authors who have worked on the message.
//
// # Discussion
//
// A message may have zero or more authors. Although the array preserves the
// order of authors, it is not intended to identify the main author or
// represent the relative importance of authors.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames
func (c CSSearchableItemAttributeSet) AuthorNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("authorNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetAuthorNames(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAuthorNames:"), objectivec.StringSliceToNSArray(value))
}

// An array of [CSPerson] objects representing the content of the From: field
// in an item.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authors
func (c CSSearchableItemAttributeSet) Authors() []CSPerson {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("authors"))
	return objc.ConvertSlice(rv, func(id objc.ID) CSPerson {
		return CSPersonFromID(id)
	})
}
func (c CSSearchableItemAttributeSet) SetAuthors(value []CSPerson) {
	objc.Send[struct{}](c.ID, objc.Sel("setAuthors:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of email addresses associated with the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses
func (c CSSearchableItemAttributeSet) EmailAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("emailAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetEmailAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmailAddresses:"), objectivec.StringSliceToNSArray(value))
}

// A dictionary that contains all the headers of the message.
//
// # Discussion
//
// Dictionary keys are header names and the values are arrays of strings,
// because one header can occur more than once in an email message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders
func (c CSSearchableItemAttributeSet) EmailHeaders() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("emailHeaders"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetEmailHeaders(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmailHeaders:"), value)
}

// An array of [CSPerson] objects representing the content of the Bcc: field
// in an email message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hiddenAdditionalRecipients
func (c CSSearchableItemAttributeSet) HiddenAdditionalRecipients() []CSPerson {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("hiddenAdditionalRecipients"))
	return objc.ConvertSlice(rv, func(id objc.ID) CSPerson {
		return CSPersonFromID(id)
	})
}
func (c CSSearchableItemAttributeSet) SetHiddenAdditionalRecipients(value []CSPerson) {
	objc.Send[struct{}](c.ID, objc.Sel("setHiddenAdditionalRecipients:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of instant message addresses for the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses
func (c CSSearchableItemAttributeSet) InstantMessageAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("instantMessageAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetInstantMessageAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setInstantMessageAddresses:"), objectivec.StringSliceToNSArray(value))
}

// A value that indicates if the message is likely to be considered junk.
//
// # Discussion
//
// When the value of this property is 1, the message is likely to be
// considered junk; when the value is 0, the message is not likely to be
// considered junk.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk
func (c CSSearchableItemAttributeSet) LikelyJunk() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("isLikelyJunk"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetLikelyJunk(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setLikelyJunk:"), value)
}

// An array of mailbox identifiers associated with the message.
//
// # Discussion
//
// Can include [CSMailboxInbox], [CSMailboxDrafts], [CSMailboxSent], or a
// custom identifier.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers
func (c CSSearchableItemAttributeSet) MailboxIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("mailboxIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetMailboxIdentifiers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMailboxIdentifiers:"), objectivec.StringSliceToNSArray(value))
}

// An array of phone numbers associated with the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers
func (c CSSearchableItemAttributeSet) PhoneNumbers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("phoneNumbers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetPhoneNumbers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPhoneNumbers:"), objectivec.StringSliceToNSArray(value))
}

// An array of [CSPerson] objects representing the content of the To: field in
// an email message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/primaryRecipients
func (c CSSearchableItemAttributeSet) PrimaryRecipients() []CSPerson {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("primaryRecipients"))
	return objc.ConvertSlice(rv, func(id objc.ID) CSPerson {
		return CSPersonFromID(id)
	})
}
func (c CSSearchableItemAttributeSet) SetPrimaryRecipients(value []CSPerson) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryRecipients:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of addresses associated with the recipients of the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses
func (c CSSearchableItemAttributeSet) RecipientAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recipientAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetRecipientAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecipientAddresses:"), objectivec.StringSliceToNSArray(value))
}

// An array of email addresses associated with the recipient.
//
// # Discussion
//
// The contents of this property is not human readable.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses
func (c CSSearchableItemAttributeSet) RecipientEmailAddresses() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recipientEmailAddresses"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecipientEmailAddresses:"), objectivec.StringSliceToNSArray(value))
}

// An array of names representing the recipients of this message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames
func (c CSSearchableItemAttributeSet) RecipientNames() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recipientNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetRecipientNames(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecipientNames:"), objectivec.StringSliceToNSArray(value))
}

// The textual content of the message.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent
func (c CSSearchableItemAttributeSet) TextContent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("textContent"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetTextContent(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTextContent:"), objc.String(value))
}

// A localized string that specifies the name of a container to which the item
// belongs, suitable to display in the user interface.
//
// # Discussion
//
// For example, a container display name might be the title of a series of
// books. When you specify the containment properties, Spotlight can treat
// individual items as part of an ordered set.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName
func (c CSSearchableItemAttributeSet) ContainerDisplayName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerDisplayName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetContainerDisplayName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerDisplayName:"), objc.String(value))
}

// The identifier of the container to which the item belongs.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier
func (c CSSearchableItemAttributeSet) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetContainerIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}

// The order of the item within the container.
//
// # Discussion
//
// For example, if the container represents a series of books, this property
// specifies the order in which the books should be read.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder
func (c CSSearchableItemAttributeSet) ContainerOrder() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerOrder"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetContainerOrder(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerOrder:"), value)
}

// The title of the container to which the item belongs.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle
func (c CSSearchableItemAttributeSet) ContainerTitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItemAttributeSet) SetContainerTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainerTitle:"), objc.String(value))
}

// A value that indicates whether the item contains information sufficient to
// provide navigation to the location it represents.
//
// # Discussion
//
// When an item includes [CSSearchableItemAttributeSet.Latitude] and
// [CSSearchableItemAttributeSet.Longitude] properties, these properties can
// be used for navigation to the location represented by the item. For
// example, it makes sense to set
// [CSSearchableItemAttributeSet.SupportsNavigation] to `1` for an item that
// represents review for a specific restaurant, but not for an item that
// represents a photo of a person.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c CSSearchableItemAttributeSet) SupportsNavigation() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportsNavigation"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetSupportsNavigation(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportsNavigation:"), value)
}

// A value that indicates whether the item contains information sufficient to
// allow a phone call to a number associated with the item.
//
// # Discussion
//
// When an item includes the [CSSearchableItemAttributeSet.PhoneNumbers]
// property, the phone number can be used to initiate phone calls. You can use
// the [CSSearchableItemAttributeSet.SupportsPhoneCall] property to indicate
// when making a phone call is appropriate and likely to be a primary action
// for the user. For example, you might set
// [CSSearchableItemAttributeSet.SupportsPhoneCall] to `1` for an item that
// represents a business, but not for an item that represents an academic
// paper that lists the phone numbers of the authors or the institution.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall
func (c CSSearchableItemAttributeSet) SupportsPhoneCall() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportsPhoneCall"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c CSSearchableItemAttributeSet) SetSupportsPhoneCall(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupportsPhoneCall:"), value)
}
