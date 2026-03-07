// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDataAsset] class.
var (
	_NSDataAssetClass     NSDataAssetClass
	_NSDataAssetClassOnce sync.Once
)

func getNSDataAssetClass() NSDataAssetClass {
	_NSDataAssetClassOnce.Do(func() {
		_NSDataAssetClass = NSDataAssetClass{class: objc.GetClass("NSDataAsset")}
	})
	return _NSDataAssetClass
}

// GetNSDataAssetClass returns the class object for NSDataAsset.
func GetNSDataAssetClass() NSDataAssetClass {
	return getNSDataAssetClass()
}

type NSDataAssetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDataAssetClass) Alloc() NSDataAsset {
	rv := objc.Send[NSDataAsset](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object from a data set type stored in an asset catalog.
//
// # Overview
// 
// The object’s content is stored as a set of one or more files with
// associated device attributes. These sets can also be tagged for use as
// on-demand resources.
// 
// # Initialize data assets
// 
// Data assets are initialized from a named data set in an asset catalog. You
// create data sets during app development. Each data set contains one or more
// data files. Each file has associated attributes for features of the device,
// including the minimum amount of memory and the version of Metal. When you
// initialize the data asset, the system selects the data file that best
// matches the current device.
// 
// For more information on the data set type in an asset catalog, see [Data
// Set Type] in [Asset Catalog Format Reference]. For information on asset
// catalogs, see [Managing assets with asset catalogs].
// 
// # Access the data
// 
// You access the data file by using the [NSDataAsset.Data] property. Because the property
// is of type [NSData] it provides methods for accessing the raw data only as
// bytes and ranges of bytes.
// 
// To access structured data, convert the bytes into the appropriate format.
// The system can convert some data types for you. One example is XML data
// using the [init(data:)] method of [XMLParser]. Other data types require
// code for parsing and converting the raw data. You may need to convert
// larger data files incrementally.
//
// [Asset Catalog Format Reference]: https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/index.html#//apple_ref/doc/uid/TP40015170
// [Data Set Type]: https://developer.apple.com/library/archive/documentation/Xcode/Reference/xcode_ref-Asset_Catalog_Format/DataSetType.html#//apple_ref/doc/uid/TP40015170-CH23
// [Managing assets with asset catalogs]: https://developer.apple.com/documentation/Xcode/managing-assets-with-asset-catalogs
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [XMLParser]: https://developer.apple.com/documentation/Foundation/XMLParser
// [init(data:)]: https://developer.apple.com/documentation/Foundation/XMLParser/init(data:)
//
// # Initializing the data asset
//
//   - [NSDataAsset.InitWithName]: Initializes and returns an object with a reference to the named data asset in an asset catalog.
//   - [NSDataAsset.InitWithNameBundle]: Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
//
// # Accessing data
//
//   - [NSDataAsset.Data]: The raw data values in the data asset.
//
// # Getting data asset information
//
//   - [NSDataAsset.Name]: The name of the data set in the asset catalog.
//   - [NSDataAsset.TypeIdentifier]: The uniform type identifier for the data asset.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset
type NSDataAsset struct {
	objectivec.Object
}

// NSDataAssetFromID constructs a [NSDataAsset] from an objc.ID.
//
// An object from a data set type stored in an asset catalog.
func NSDataAssetFromID(id objc.ID) NSDataAsset {
	return NSDataAsset{objectivec.Object{id}}
}
// NOTE: NSDataAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDataAsset] class.
//
// # Initializing the data asset
//
//   - [INSDataAsset.InitWithName]: Initializes and returns an object with a reference to the named data asset in an asset catalog.
//   - [INSDataAsset.InitWithNameBundle]: Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
//
// # Accessing data
//
//   - [INSDataAsset.Data]: The raw data values in the data asset.
//
// # Getting data asset information
//
//   - [INSDataAsset.Name]: The name of the data set in the asset catalog.
//   - [INSDataAsset.TypeIdentifier]: The uniform type identifier for the data asset.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset
type INSDataAsset interface {
	objectivec.IObject

	// Topic: Initializing the data asset

	// Initializes and returns an object with a reference to the named data asset in an asset catalog.
	InitWithName(name NSDataAssetName) NSDataAsset
	// Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
	InitWithNameBundle(name NSDataAssetName, bundle *foundation.NSBundle) NSDataAsset

	// Topic: Accessing data

	// The raw data values in the data asset.
	Data() foundation.INSData

	// Topic: Getting data asset information

	// The name of the data set in the asset catalog.
	Name() NSDataAssetName
	// The uniform type identifier for the data asset.
	TypeIdentifier() string
}





// Init initializes the instance.
func (d NSDataAsset) Init() NSDataAsset {
	rv := objc.Send[NSDataAsset](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDataAsset) Autorelease() NSDataAsset {
	rv := objc.Send[NSDataAsset](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDataAsset creates a new NSDataAsset instance.
func NewNSDataAsset() NSDataAsset {
	class := getNSDataAssetClass()
	rv := objc.Send[NSDataAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns an object with a reference to the named data asset
// in an asset catalog.
//
// name: The name of the data set in the asset catalog.
//
// # Return Value
// 
// The data asset object for the named data set, or `nil` if the data set is
// not found.
//
// # Discussion
// 
// If there are multiple data files in the named data set, this method returns
// the one with attributes that most closely match the current device
// available screen space.
// 
// This method looks in the asset catalog, in the main bundle for the named
// data set.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:)
func NewDataAssetWithName(name NSDataAssetName) NSDataAsset {
	instance := getNSDataAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(string(name)))
	return NSDataAssetFromID(rv)
}


// Initializes and returns an object with a reference to the named data asset
// that’s in an asset catalog in the specified bundle.
//
// name: The name of the data set in the asset catalog.
//
// bundle: The bundle used to store the asset catalog. Pass `nil` for the main bundle.
// The bundle must be the same as the one used in the Xcode project.
//
// # Return Value
// 
// The data asset object for the named data set in the specified bundle, or
// `nil` if the data set is not found.
//
// # Discussion
// 
// If there are multiple data files in the named data set, this method returns
// the one with attributes that most closely match the current device
// available screen space.
// 
// This method looks in the asset catalog, in the bundle specified by the
// `bundle` parameter for the named data set.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:bundle:)
func NewDataAssetWithNameBundle(name NSDataAssetName, bundle *foundation.NSBundle) NSDataAsset {
	instance := getNSDataAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:bundle:"), objc.String(string(name)), bundle)
	return NSDataAssetFromID(rv)
}







// Initializes and returns an object with a reference to the named data asset
// in an asset catalog.
//
// name: The name of the data set in the asset catalog.
//
// # Return Value
// 
// The data asset object for the named data set, or `nil` if the data set is
// not found.
//
// # Discussion
// 
// If there are multiple data files in the named data set, this method returns
// the one with attributes that most closely match the current device
// available screen space.
// 
// This method looks in the asset catalog, in the main bundle for the named
// data set.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:)
func (d NSDataAsset) InitWithName(name NSDataAssetName) NSDataAsset {
	rv := objc.Send[NSDataAsset](d.ID, objc.Sel("initWithName:"), objc.String(string(name)))
	return rv
}

// Initializes and returns an object with a reference to the named data asset
// that’s in an asset catalog in the specified bundle.
//
// name: The name of the data set in the asset catalog.
//
// bundle: The bundle used to store the asset catalog. Pass `nil` for the main bundle.
// The bundle must be the same as the one used in the Xcode project.
//
// # Return Value
// 
// The data asset object for the named data set in the specified bundle, or
// `nil` if the data set is not found.
//
// # Discussion
// 
// If there are multiple data files in the named data set, this method returns
// the one with attributes that most closely match the current device
// available screen space.
// 
// This method looks in the asset catalog, in the bundle specified by the
// `bundle` parameter for the named data set.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:bundle:)
func (d NSDataAsset) InitWithNameBundle(name NSDataAssetName, bundle *foundation.NSBundle) NSDataAsset {
	rv := objc.Send[NSDataAsset](d.ID, objc.Sel("initWithName:bundle:"), objc.String(string(name)), bundle)
	return rv
}












// The raw data values in the data asset.
//
// # Discussion
// 
// For more information on accessing structured data, see [NSDataAsset].
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/data
func (d NSDataAsset) Data() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}



// The name of the data set in the asset catalog.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/name-swift.property
func (d NSDataAsset) Name() NSDataAssetName {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("name"))
	return NSDataAssetName(foundation.NSStringFromID(rv).String())
}



// The uniform type identifier for the data asset.
//
// # Discussion
// 
// A uniform type identifier is a string for identifying the type of data.
// This UTI is the same as the one specified in the asset catalog. For more
// information, see [Uniform Type Identifiers Overview].
//
// [Uniform Type Identifiers Overview]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/understanding_utis/understand_utis_intro/understand_utis_intro.html#//apple_ref/doc/uid/TP40001319
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/typeIdentifier
func (d NSDataAsset) TypeIdentifier() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("typeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
























