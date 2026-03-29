// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorSpace] class.
var (
	_NSColorSpaceClass     NSColorSpaceClass
	_NSColorSpaceClassOnce sync.Once
)

func getNSColorSpaceClass() NSColorSpaceClass {
	_NSColorSpaceClassOnce.Do(func() {
		_NSColorSpaceClass = NSColorSpaceClass{class: objc.GetClass("NSColorSpace")}
	})
	return _NSColorSpaceClass
}

// GetNSColorSpaceClass returns the class object for NSColorSpace.
func GetNSColorSpaceClass() NSColorSpaceClass {
	return getNSColorSpaceClass()
}

type NSColorSpaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSColorSpaceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorSpaceClass) Alloc() NSColorSpace {
	rv := objc.Send[NSColorSpace](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a custom color space.
//
// # Overview
// 
// You can make custom color spaces from ColorSync profiles or from ICC
// profiles. [NSColorSpace] also has factory methods that return objects
// representing the system color spaces.
// 
// You can use the [ColorWithColorSpaceComponentsCount] method of the
// [NSColor] class to create color objects using custom [NSColorSpace]
// objects. You can also send the [ColorUsingColorSpace] message to an
// [NSColor] object to convert it between two color spaces, either of which
// may be a custom color space.
//
// # Initializing a Custom Color Space Object
//
//   - [NSColorSpace.InitWithCGColorSpace]: Initializes and returns a color space object initialized from a Core Graphics color-space object.
//   - [NSColorSpace.InitWithColorSyncProfile]: Initializes and returns a color space object from the specified ColorSync profile.
//   - [NSColorSpace.InitWithICCProfileData]: Initializes and returns a color space object from the specified ICC profile.
//
// # Accessing Color Space Data and Attributes
//
//   - [NSColorSpace.CGColorSpace]: The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//   - [NSColorSpace.ColorSpaceModel]: The model on which the color space is based.
//   - [NSColorSpace.ColorSyncProfile]: The ColorSync profile from which the color space was created.
//   - [NSColorSpace.ICCProfileData]: The ICC profile data from which the color space was created.
//   - [NSColorSpace.LocalizedName]: The localized name of the color space.
//   - [NSColorSpace.NumberOfColorComponents]: The number of components, excluding alpha, the color space supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace
type NSColorSpace struct {
	objectivec.Object
}

// NSColorSpaceFromID constructs a [NSColorSpace] from an objc.ID.
//
// An object that represents a custom color space.
func NSColorSpaceFromID(id objc.ID) NSColorSpace {
	return NSColorSpace{objectivec.Object{ID: id}}
}
// NOTE: NSColorSpace adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorSpace] class.
//
// # Initializing a Custom Color Space Object
//
//   - [INSColorSpace.InitWithCGColorSpace]: Initializes and returns a color space object initialized from a Core Graphics color-space object.
//   - [INSColorSpace.InitWithColorSyncProfile]: Initializes and returns a color space object from the specified ColorSync profile.
//   - [INSColorSpace.InitWithICCProfileData]: Initializes and returns a color space object from the specified ICC profile.
//
// # Accessing Color Space Data and Attributes
//
//   - [INSColorSpace.CGColorSpace]: The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//   - [INSColorSpace.ColorSpaceModel]: The model on which the color space is based.
//   - [INSColorSpace.ColorSyncProfile]: The ColorSync profile from which the color space was created.
//   - [INSColorSpace.ICCProfileData]: The ICC profile data from which the color space was created.
//   - [INSColorSpace.LocalizedName]: The localized name of the color space.
//   - [INSColorSpace.NumberOfColorComponents]: The number of components, excluding alpha, the color space supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace
type INSColorSpace interface {
	objectivec.IObject

	// Topic: Initializing a Custom Color Space Object

	// Initializes and returns a color space object initialized from a Core Graphics color-space object.
	InitWithCGColorSpace(cgColorSpace coregraphics.CGColorSpaceRef) NSColorSpace
	// Initializes and returns a color space object from the specified ColorSync profile.
	InitWithColorSyncProfile(prof unsafe.Pointer) NSColorSpace
	// Initializes and returns a color space object from the specified ICC profile.
	InitWithICCProfileData(iccData foundation.INSData) NSColorSpace

	// Topic: Accessing Color Space Data and Attributes

	// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
	CGColorSpace() coregraphics.CGColorSpaceRef
	// The model on which the color space is based.
	ColorSpaceModel() NSColorSpaceModel
	// The ColorSync profile from which the color space was created.
	ColorSyncProfile() unsafe.Pointer
	// The ICC profile data from which the color space was created.
	ICCProfileData() foundation.INSData
	// The localized name of the color space.
	LocalizedName() string
	// The number of components, excluding alpha, the color space supports.
	NumberOfColorComponents() int

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSColorSpace) Init() NSColorSpace {
	rv := objc.Send[NSColorSpace](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorSpace) Autorelease() NSColorSpace {
	rv := objc.Send[NSColorSpace](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorSpace creates a new NSColorSpace instance.
func NewNSColorSpace() NSColorSpace {
	class := getNSColorSpaceClass()
	rv := objc.Send[NSColorSpace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a color space object initialized from a Core
// Graphics color-space object.
//
// cgColorSpace: A reference to a Core Graphics color-space object ([CGColorSpace]).
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful, which might happen if the color space represented by the
// [CGColorSpace] object is not supported by [NSColorSpace].
//
// # Discussion
// 
// Because [NSColorSpace] might retain or copy the [CGColorSpace] object
// depending on circumstances, you should not assume pointer equality of the
// provided object with that returned by [CGColorSpace]. And even if the
// pointer equality is preserved during runtime, it may not be after the
// [NSColorSpace] object is archived and unarchived.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(cgColorSpace:)
func NewColorSpaceWithCGColorSpace(cgColorSpace coregraphics.CGColorSpaceRef) NSColorSpace {
	instance := getNSColorSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGColorSpace:"), cgColorSpace)
	return NSColorSpaceFromID(rv)
}

// Initializes and returns a color space object from the specified ColorSync
// profile.
//
// prof: The ColorSync profile to use when initializing the [NSColorSpace] object.
// This should be an object of opaque type CMProfileRef. See [ColorSync
// Manager] for further information on CMProfileRef.
// //
// [ColorSync Manager]: https://developer.apple.com/documentation/applicationservices/colorsync_manager
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(colorSyncProfile:)
func NewColorSpaceWithColorSyncProfile(prof unsafe.Pointer) NSColorSpace {
	instance := getNSColorSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColorSyncProfile:"), prof)
	return NSColorSpaceFromID(rv)
}

// Initializes and returns a color space object from the specified ICC
// profile.
//
// iccData: The ICC profile to use when initializing the [NSColorSpace] object. For
// information on ICC profiles, see the latest ICC specification at the
// [International Color Consortium website] website.
// //
// [International Color Consortium website]: http://www.color.org/icc_specs2.html
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(iccProfileData:)
func NewColorSpaceWithICCProfileData(iccData foundation.INSData) NSColorSpace {
	instance := getNSColorSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithICCProfileData:"), iccData)
	return NSColorSpaceFromID(rv)
}

// Initializes and returns a color space object initialized from a Core
// Graphics color-space object.
//
// cgColorSpace: A reference to a Core Graphics color-space object ([CGColorSpace]).
// //
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful, which might happen if the color space represented by the
// [CGColorSpace] object is not supported by [NSColorSpace].
//
// # Discussion
// 
// Because [NSColorSpace] might retain or copy the [CGColorSpace] object
// depending on circumstances, you should not assume pointer equality of the
// provided object with that returned by [CGColorSpace]. And even if the
// pointer equality is preserved during runtime, it may not be after the
// [NSColorSpace] object is archived and unarchived.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(cgColorSpace:)
func (c NSColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.CGColorSpaceRef) NSColorSpace {
	rv := objc.Send[NSColorSpace](c.ID, objc.Sel("initWithCGColorSpace:"), cgColorSpace)
	return rv
}
// Initializes and returns a color space object from the specified ColorSync
// profile.
//
// prof: The ColorSync profile to use when initializing the [NSColorSpace] object.
// This should be an object of opaque type CMProfileRef. See [ColorSync
// Manager] for further information on CMProfileRef.
// //
// [ColorSync Manager]: https://developer.apple.com/documentation/applicationservices/colorsync_manager
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(colorSyncProfile:)
func (c NSColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) NSColorSpace {
	rv := objc.Send[NSColorSpace](c.ID, objc.Sel("initWithColorSyncProfile:"), prof)
	return rv
}
// Initializes and returns a color space object from the specified ICC
// profile.
//
// iccData: The ICC profile to use when initializing the [NSColorSpace] object. For
// information on ICC profiles, see the latest ICC specification at the
// [International Color Consortium website] website.
// //
// [International Color Consortium website]: http://www.color.org/icc_specs2.html
//
// # Return Value
// 
// The initialized [NSColorSpace] object or `nil` if initialization was not
// successful.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(iccProfileData:)
func (c NSColorSpace) InitWithICCProfileData(iccData foundation.INSData) NSColorSpace {
	rv := objc.Send[NSColorSpace](c.ID, objc.Sel("initWithICCProfileData:"), iccData)
	return rv
}
func (c NSColorSpace) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the list of color spaces available on the system that are displayed
// in the color panel, in the order they are displayed in the color panel.
//
// model: The model to return the color spaces for.
//
// # Return Value
// 
// The list of color spaces, or an empty array if no color spaces are
// available for the specified model.
//
// # Discussion
// 
// This method doesn’t return color spaces created on the fly or spaces
// without user-displayable names. Pass [NSUnknownColorSpaceModel] as `model`
// to get all available color spaces.
//
// [NSUnknownColorSpaceModel]: https://developer.apple.com/documentation/AppKit/NSUnknownColorSpaceModel
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/availableColorSpaces(with:)
func (_NSColorSpaceClass NSColorSpaceClass) AvailableColorSpacesWithModel(model NSColorSpaceModel) []NSColorSpace {
	rv := objc.Send[[]objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("availableColorSpacesWithModel:"), model)
	return objc.ConvertSlice(rv, func(id objc.ID) NSColorSpace {
		return NSColorSpaceFromID(id)
	})
}

// The Core Graphics color-space object that represents a color space
// equivalent to the color space’s.
//
// # Discussion
// 
// The value of this property is a reference to an Core Graphics color-space
// object ([CGColorSpace]) or [NULL] if the type of color space represented by
// the receiver cannot be represented by a [CGColorSpace] object.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/cgColorSpace
func (c NSColorSpace) CGColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](c.ID, objc.Sel("CGColorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}
// The model on which the color space is based.
//
// # Discussion
// 
// See `Color Space Models` for a list of valid NSColorSpaceModel constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/colorSpaceModel
func (c NSColorSpace) ColorSpaceModel() NSColorSpaceModel {
	rv := objc.Send[NSColorSpaceModel](c.ID, objc.Sel("colorSpaceModel"))
	return NSColorSpaceModel(rv)
}
// The ColorSync profile from which the color space was created.
//
// # Discussion
// 
// The ColorSync profile on which the receiver is based. You need to cast this
// value to an object of opaque type CMProfileRef. Returns [NULL] if the
// receiver was created from a ICC-profile data instead. See [ColorSync
// Manager] for further information on CMProfileRef.
//
// [ColorSync Manager]: https://developer.apple.com/documentation/applicationservices/colorsync_manager
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/colorSyncProfile
func (c NSColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("colorSyncProfile"))
	return rv
}
// The ICC profile data from which the color space was created.
//
// # Discussion
// 
// The ICC profile from which the receiver was created. This method attempts
// to compute the profile data from a CMProfileRef object and returns `nil` if
// it is unable to.
// 
// For information on ICC profiles, see the latest ICC specification at the
// [International Color Consortium website].
//
// [International Color Consortium website]: http://www.color.org/icc_specs2.html
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/iccProfileData
func (c NSColorSpace) ICCProfileData() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ICCProfileData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The localized name of the color space.
//
// # Discussion
// 
// This property holds the name of the color space as a localized string or
// `nil` if no localized name exists.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/localizedName
func (c NSColorSpace) LocalizedName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
// The number of components, excluding alpha, the color space supports.
//
// # Discussion
// 
// This value is `0` if the color space isn’t based on `float` components.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/numberOfColorComponents
func (c NSColorSpace) NumberOfColorComponents() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfColorComponents"))
	return rv
}

// A color space object that represents a calibrated or device-dependent RGB
// color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color space has red, green, blue, and alpha
// components. Typical devices that use the color-additive RGB color space are
// displays and scanners. This object corresponds to the Cocoa color space
// name [NSDeviceRGBColorSpace].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/deviceRGB
func (_NSColorSpaceClass NSColorSpaceClass) DeviceRGBColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("deviceRGBColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a device-independent RGB color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color-additive color space has red, green,
// blue, and alpha components. This object corresponds to the Cocoa color
// space name [NSCalibratedRGBColorSpace].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericRGB
func (_NSColorSpaceClass NSColorSpaceClass) GenericRGBColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("genericRGBColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a calibrated or device-dependent CMYK
// color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color space has cyan, magenta, yellow,
// black, and alpha components. Typical devices that use the color-subtractive
// CMYK color space are color printers. This object corresponds to the Cocoa
// color space name [NSDeviceCMYKColorSpace].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/deviceCMYK
func (_NSColorSpaceClass NSColorSpaceClass) DeviceCMYKColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("deviceCMYKColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a device-independent CMYK color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color space has cyan, magenta, yellow,
// black and alpha component.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericCMYK
func (_NSColorSpaceClass NSColorSpaceClass) GenericCMYKColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("genericCMYKColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a calibrated or device-dependent gray
// color space.
//
// # Return Value
// 
// The [NSColorSpace] object. The color space also includes an alpha
// component. Typical devices that use this color space are grayscale printers
// and displays. This object corresponds to the Cocoa color space name
// [NSDeviceWhiteColorSpace].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/deviceGray
func (_NSColorSpaceClass NSColorSpaceClass) DeviceGrayColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("deviceGrayColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a device-independent gray color space.
//
// # Return Value
// 
// The [NSColorSpace] object. The color space also includes an alpha
// component. This object corresponds to the Cocoa color space name
// [NSCalibratedWhiteColorSpace].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericGray
func (_NSColorSpaceClass NSColorSpaceClass) GenericGrayColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("genericGrayColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents an sRGB color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color-additive color space has red, green,
// blue, and alpha components.
// 
// # Discussion
// 
// The sRGB color space is a standard color space for use on monitors,
// printers, and the Internet. For further information on sRGB, see
// [http://www.color.org/srgb.html].
//
// [http://www.color.org/srgb.html]: http://www.color.org/srgb.html
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/sRGB
func (_NSColorSpaceClass NSColorSpaceClass) SRGBColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("sRGBColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents an extended sRGB color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color-additive color space has red, green,
// blue, and alpha components.
// 
// # Discussion
// 
// This color space has the same colorimetry as sRGB, but component values
// below `0.0` and above `1.0` may be encoded in this color space. Negative
// values are encoded as the signed reflection of the original encoding
// function. `y(x) = sign(x)*f(abs(x))`
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedSRGB
func (_NSColorSpaceClass NSColorSpaceClass) ExtendedSRGBColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("extendedSRGBColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a P3 Display color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color-additive color space has red, green,
// blue, and alpha components.
// 
// # Discussion
// 
// The Display P3 color space, created by Apple Inc. This color space uses the
// DCI P3 primaries, a D65 white point, and the same gamma curve as the sRGB
// color space.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/displayP3
func (_NSColorSpaceClass NSColorSpaceClass) DisplayP3ColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("displayP3ColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents a gray color space with a gamma value
// of 2.2.
//
// # Return Value
// 
// The [NSColorSpace] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericGamma22Gray
func (_NSColorSpaceClass NSColorSpaceClass) GenericGamma22GrayColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("genericGamma22GrayColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents an extended gray color space with a
// gamma value of 2.2.
//
// # Return Value
// 
// The [NSColorSpace] object.
// 
// # Discussion
// 
// This color space has the same colorimetry as Generic Gray 2.2, but
// component values below `0.0` and above `1.0` may be encoded in this color
// space. Negative values are encoded as the signed reflection of the original
// encoding function. `y(x) = sign(x)*f(abs(x))`
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedGenericGamma22Gray
func (_NSColorSpaceClass NSColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// A color space object that represents an Adobe RGB (1998) color space.
//
// # Return Value
// 
// The [NSColorSpace] object. This color-additive color space has red, green,
// blue, and alpha components.
// 
// # Discussion
// 
// The Adobe RGB (1998) color space was designed to encompass most of the
// colors achievable on CMYK color printers, but by using RGB primary colors
// on a device such as the computer display. For more information on this
// color space, go to [http://www.adobe.com/digitalimag/adobergb.html].
//
// [http://www.adobe.com/digitalimag/adobergb.html]: http://www.adobe.com/digitalimag/adobergb.html
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpace/adobeRGB1998
func (_NSColorSpaceClass NSColorSpaceClass) AdobeRGB1998ColorSpace() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorSpaceClass.class), objc.Sel("adobeRGB1998ColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}

