// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIFilter] class.
var (
	_CIFilterClass     CIFilterClass
	_CIFilterClassOnce sync.Once
)

func getCIFilterClass() CIFilterClass {
	_CIFilterClassOnce.Do(func() {
		_CIFilterClass = CIFilterClass{class: objc.GetClass("CIFilter")}
	})
	return _CIFilterClass
}

// GetCIFilterClass returns the class object for CIFilter.
func GetCIFilterClass() CIFilterClass {
	return getCIFilterClass()
}

type CIFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIFilterClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIFilterClass) Alloc() CIFilter {
	rv := objc.Send[CIFilter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An image processor that produces an image by manipulating one or more input
// images or by generating new image data.
//
// # Overview
//
// The [CIFilter] class produces a [CIImage] object as output. Typically, a
// filter takes one or more images as input. Some filters, however, generate
// an image based on other types of input parameters. The par[CIFilter]
// swift.class` object are set and retrieved through the use of key-value
// pairs.
//
// You use the [CIFilter] object in conjunction with other Core Image classes,
// such as [CIImage], [CIContext], and [CIColor], to take advantage of the
// built-in Core Image filters when processing images, creating filter
// generators, or writing custom filters.
//
// [CIFilter] objects are mutable, and thus cannot be shared safely among
// threads. Each thread must create its own [CIFilter] objects, but you can
// pass a filter’s immutable input and output [CIImage] objects between
// threads.
//
// To get a quick overview of how to set up and use Core Image filters, see
// [Core Image Programming Guide].
//
// # Create type-safe filters
//
// Core Image provides methods that create type-safe [CIFilter] instances. Use
// these filters to avoid run-time errors that can occur when relying on Core
// Image’s string-based API.
//
// To use the type-safe API, import `CoreImage.CIFilterBuiltins`:
//
// The type-safe approach returns a non-optional filter. Because the returned
// filter conforms to the relevant protocol—for example, [CIFalseColor] in
// the case of [CIFilter.FalseColorFilter]—the parameters are available as
// properties. The following creates and applies a false color filter:
//
// The false color filter maps luminance to a color ramp of two colors:
//
// [media-4336877]
//
// # Subclassing notes
//
// You can subclass [CIFilter] in order to create custom filter effects:
//
// - By chaining together two or more built-in Core Image filters - By using
// an image-processing kernel that you write
//
// Regardless of whether your subclass provides its effect by chaining filters
// or implementing its own kernel, you should:
//
// - Declare any input parameters as properties whose names are prefixed with
// `input`, such as `inputImage`. - Override the [CIFilter.SetDefaults] methods to
// provide default values for any input parameters you’ve declared. -
// Implement an `outputImage` method to create a new [CIImage] with your
// filter’s effect.
//
// The [CIFilter] class automatically manages input parameters when archiving,
// copying, and deallocating filters. For this reason, your subclass must obey
// the following guidelines to ensure proper behavior:
//
// - Store input parameters in instance variables whose names are prefixed
// with `input`.
//
// Don’t use auto-synthesized instance variables, because their names are
// automatically prefixed with an underscore. Instead, synthesize the property
// manually. For example:
//
// `@synthesize inputMyParameter;`
//
// - If using manual reference counting, don’t release input parameter
// instance variables in your [dealloc] method implementation. The [dealloc]
// implementation in the [CIFilter] class uses [Key-value coding] to
// automatically set the values of all input parameters to `nil`.
//
// # Getting filter parameters and attributes
//
//   - [CIFilter.Name]: A name associated with a filter.
//   - [CIFilter.SetName]
//   - [CIFilter.Enabled]: A Boolean value that determines whether the filter is enabled. Animatable.
//   - [CIFilter.SetEnabled]
//   - [CIFilter.Attributes]: A dictionary of key-value pairs that describe the filter.
//   - [CIFilter.InputKeys]: The names of all input parameters to the filter.
//   - [CIFilter.OutputKeys]: The names of all output parameters from the filter.
//   - [CIFilter.OutputImage]: Returns a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object that encapsulates the operations configured in the filter.
//
// # Setting default values
//
//   - [CIFilter.SetDefaults]: Sets all input values for a filter to default values.
//
// # Applying a filter
//
//   - [CIFilter.ApplyArgumentsOptions]: Produces a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object by applying arguments to a kernel function and using options to control how the kernel function is evaluated.
//
// # Creating a configuration view for a filter
//
//   - [CIFilter.ViewForUIConfigurationExcludedKeys]: Returns a filter view for the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
//
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
// [Key-value coding]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KeyValueCoding.html#//apple_ref/doc/uid/TP40008195-CH25
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
type CIFilter struct {
	objectivec.Object
}

// CIFilterFromID constructs a [CIFilter] from an objc.ID.
//
// An image processor that produces an image by manipulating one or more input
// images or by generating new image data.
func CIFilterFromID(id objc.ID) CIFilter {
	return CIFilter{objectivec.Object{ID: id}}
}

// NOTE: CIFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIFilter] class.
//
// # Getting filter parameters and attributes
//
//   - [ICIFilter.Name]: A name associated with a filter.
//   - [ICIFilter.SetName]
//   - [ICIFilter.Enabled]: A Boolean value that determines whether the filter is enabled. Animatable.
//   - [ICIFilter.SetEnabled]
//   - [ICIFilter.Attributes]: A dictionary of key-value pairs that describe the filter.
//   - [ICIFilter.InputKeys]: The names of all input parameters to the filter.
//   - [ICIFilter.OutputKeys]: The names of all output parameters from the filter.
//   - [ICIFilter.OutputImage]: Returns a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object that encapsulates the operations configured in the filter.
//
// # Setting default values
//
//   - [ICIFilter.SetDefaults]: Sets all input values for a filter to default values.
//
// # Applying a filter
//
//   - [ICIFilter.ApplyArgumentsOptions]: Produces a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object by applying arguments to a kernel function and using options to control how the kernel function is evaluated.
//
// # Creating a configuration view for a filter
//
//   - [ICIFilter.ViewForUIConfigurationExcludedKeys]: Returns a filter view for the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
type ICIFilter interface {
	objectivec.IObject

	// Topic: Getting filter parameters and attributes

	// A name associated with a filter.
	Name() string
	SetName(value string)
	// A Boolean value that determines whether the filter is enabled. Animatable.
	Enabled() bool
	SetEnabled(value bool)
	// A dictionary of key-value pairs that describe the filter.
	Attributes() foundation.INSDictionary
	// The names of all input parameters to the filter.
	InputKeys() []string
	// The names of all output parameters from the filter.
	OutputKeys() []string
	// Returns a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object that encapsulates the operations configured in the filter.
	OutputImage() ICIImage

	// Topic: Setting default values

	// Sets all input values for a filter to default values.
	SetDefaults()

	// Topic: Applying a filter

	// Produces a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object by applying arguments to a kernel function and using options to control how the kernel function is evaluated.
	ApplyArgumentsOptions(k ICIKernel, args foundation.INSArray, dict foundation.INSDictionary) ICIImage

	// Topic: Creating a configuration view for a filter

	// Returns a filter view for the filter.
	ViewForUIConfigurationExcludedKeys(inUIConfiguration foundation.INSDictionary, inKeys foundation.INSArray) objectivec.IObject

	// Produces a [CIImage](<doc://com.apple.coreimage/documentation/CoreImage/CIImage>) object by applying a kernel function.
	Apply(k ICIKernel) ICIImage
	EncodeWithCoder(coder foundation.INSCoder)
	// Sets the value of the property identified by the given key.
	SetValueForKey(value objectivec.IObject, key string)
	// Returns the value of the property identified by the given key.
	ValueForKey(key string) objectivec.IObject
}

// Init initializes the instance.
func (f CIFilter) Init() CIFilter {
	rv := objc.Send[CIFilter](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CIFilter) Autorelease() CIFilter {
	rv := objc.Send[CIFilter](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIFilter creates a new CIFilter instance.
func NewCIFilter() CIFilter {
	class := getCIFilterClass()
	rv := objc.Send[CIFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a [CIFilter] object for a specific kind of filter.
//
// name: The name of the filter. You must make sure the name is spelled correctly,
// otherwise your app will run but not produce any output images. For that
// reason, you should check for the existence of the filter after calling this
// method.
//
// # Return Value
//
// A [CIFilter] object whose input values are undefined.
//
// # Discussion
//
// In macOS, after creating a filter with this method you must call
// [SetDefaults] or set parameters individually by calling
// [setValue(_:forKey:)]. In iOS, the filter’s parameters are automatically
// set to default values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:)
//
// [setValue(_:forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKey:)
func NewFilterWithName(name string) CIFilter {
	rv := objc.Send[objc.ID](objc.ID(getCIFilterClass().class), objc.Sel("filterWithName:"), objc.String(name))
	return CIFilterFromID(rv)
}

// Creates a new filter of type ‘name’. The filter’s input parameters
// are set from the dictionary of key-value pairs. On OSX, any of the filter
// input parameters not specified in the dictionary will be undefined. On iOS,
// any of the filter input parameters not specified in the dictionary will be
// set to default values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:parameters:)
func NewFilterWithNameWithInputParameters(name string, params foundation.INSDictionary) CIFilter {
	rv := objc.Send[objc.ID](objc.ID(getCIFilterClass().class), objc.Sel("filterWithName:withInputParameters:"), objc.String(name), params)
	return CIFilterFromID(rv)
}

// Sets all input values for a filter to default values.
//
// # Discussion
//
// Input values whose default values are not defined are left unchanged.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/setDefaults()
func (f CIFilter) SetDefaults() {
	objc.Send[objc.ID](f.ID, objc.Sel("setDefaults"))
}

// Produces a [CIImage] object by applying arguments to a kernel function and
// using options to control how the kernel function is evaluated.
//
// k: A [CIKernel] object that contains a kernel function.
//
// args: The arguments that are type compatible with the function signature of the
// kernel function.
//
// dict: A dictionary that contains options (key-value pairs) to control how the
// kernel function is evaluated.
//
// # Return Value
//
// The [CIImage] object produced by a filter.
//
// # Discussion
//
// If you are implementing a custom filter, this method needs to be called
// from within the [OutputImage] method in order to apply your kernel function
// to the [CIImage] object. You can pass any of the keys defined in [Options
// for Applying a Filter], along with appropriate values, into the options
// dictionary.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/apply(_:arguments:options:)
//
// [Options for Applying a Filter]: https://developer.apple.com/documentation/CoreImage/options-for-applying-a-filter
func (f CIFilter) ApplyArgumentsOptions(k ICIKernel, args foundation.INSArray, dict foundation.INSDictionary) ICIImage {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("apply:arguments:options:"), k, args, dict)
	return CIImageFromID(rv)
}

// Returns a filter view for the filter.
//
// inUIConfiguration: A dictionary that contains values for the [IKUISizeFlavor] and
// [kCIUIParameterSet] keys. For allowed values for the [IKUISizeFlavor] key,
// see [User Interface Options]. For allowed values for the
// [kCIUIParameterSet] key, see [User Interface Control Options].
//
// inKeys: An array of the input keys for which you do want to provide a user
// interface. Pass `nil` if you want all input keys to be represented in the
// user interface.
//
// # Return Value
//
// An [IKFilterUIView] object.
//
// # Discussion
//
// Calling this method to receive a view for a filter causes the [CIFilter]
// class to invoke the [provideView(forUIConfiguration:excludedKeys:)] method.
// If you override [provideView(forUIConfiguration:excludedKeys:)] the user
// interface is created by your filter subclass. Otherwise, Core Image
// automatically generates the user interface based on the filter keys and
// attributes.
//
// Your app can retrieve a view whose control sizes complement the size of
// user interface elements already used in the application. It is also
// possible to choose which filter input parameters appear in the view.
// Consumer applications, for example, may want to show a small, basic set of
// input parameters whereas professional applications may want to provide
// access to all input parameters.
//
// When you request a user interface for a parameter set, all keys for that
// set and below are included. For example, the advanced set consists of all
// parameters in the basic, intermediate and advanced sets. The development
// set should contain parameters that are either experimental or for debugging
// purposes. You should use them only during the development of filters and
// client applications, and not in a shipping product.
//
// The controls in the view use bindings to set the values of the filter. See
// [Cocoa Bindings Programming Topics] if you are unfamiliar with bindings.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/view(forUIConfiguration:excludedKeys:)
//
// [IKUISizeFlavor]: https://developer.apple.com/documentation/Quartz/IKUISizeFlavor
// [User Interface Control Options]: https://developer.apple.com/documentation/CoreImage/user-interface-control-options
// [User Interface Options]: https://developer.apple.com/documentation/CoreImage/user-interface-options
// [kCIUIParameterSet]: https://developer.apple.com/documentation/CoreImage/kCIUIParameterSet
// [IKFilterUIView]: https://developer.apple.com/documentation/Quartz/IKFilterUIView
// [Cocoa Bindings Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaBindings/CocoaBindings.html#//apple_ref/doc/uid/10000167i
// [provideView(forUIConfiguration:excludedKeys:)]: https://developer.apple.com/documentation/Quartz/IKFilterCustomUIProvider/provideView(forUIConfiguration:excludedKeys:)
func (f CIFilter) ViewForUIConfigurationExcludedKeys(inUIConfiguration foundation.INSDictionary, inKeys foundation.INSArray) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("viewForUIConfiguration:excludedKeys:"), inUIConfiguration, inKeys)
	return objectivec.Object{ID: rv}
}

// Produces a [CIImage] object by applying a kernel function.
//
// k: A [CIKernel] object that contains a kernel function.
//
// # Discussion
//
// If you are implementing a custom filter, this method needs to be called
// from within the [OutputImage] method in order to apply your kernel function
// to the [CIImage] object. For example, if the kernel function has this
// signature:
//
// You would supply two arguments after the `k` argument to the `k, ...`
// method. In this case, the first argument must be a sampler and the second a
// floating-point value. For more information on kernels, see [Core Image
// Kernel Language Reference].
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/apply:
//
// [Core Image Kernel Language Reference]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CIKernelLangRef/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004397
func (f CIFilter) Apply(k ICIKernel) ICIImage {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("apply:"), k)
	return CIImageFromID(rv)
}
func (f CIFilter) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Sets the value of the property identified by the given key. [Full Topic]
func (f CIFilter) SetValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](f.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}

// Returns the value of the property identified by the given key. [Full Topic]
func (f CIFilter) ValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("valueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns an array of all published filter names that match all the specified
// categories.
//
// categories: One or more of the filter category keys defined in [Filter Category Keys].
// Pass `nil` to get all filters in all categories.
//
// # Return Value
//
// An array that contains all published filter names that match all the
// categories specified by the `categories` argument.
//
// # Discussion
//
// When you pass more than one filter category, this method returns the
// intersection of the filters in the categories. For example, if you pass the
// categories [kCICategoryBuiltIn] and [kCICategoryColorAdjustment], you
// obtain all the filters that are members of both the built-in and color
// adjustment categories. But if you pass in `kCICategoryGenerator` and
// [kCICategoryStylize], you will not get any filters returned to you because
// there are no filters that are members of both the generator and stylize
// categories. If you want to obtain all stylize and generator filters, you
// must call the “ method for each category separately and then merge the
// results.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterNames(inCategories:)
//
// [Filter Category Keys]: https://developer.apple.com/documentation/CoreImage/filter-category-keys
// [kCICategoryBuiltIn]: https://developer.apple.com/documentation/CoreImage/kCICategoryBuiltIn
// [kCICategoryColorAdjustment]: https://developer.apple.com/documentation/CoreImage/kCICategoryColorAdjustment
// [kCICategoryStylize]: https://developer.apple.com/documentation/CoreImage/kCICategoryStylize
func (_CIFilterClass CIFilterClass) FilterNamesInCategories(categories []string) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("filterNamesInCategories:"), objectivec.StringSliceToNSArray(categories))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array of all published filter names in the specified category.
//
// category: A string object that specifies one of the filter categories defined in
// [Filter Category Keys].
//
// # Return Value
//
// An array that contains all published names of the filter in a category.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterNames(inCategory:)
//
// [Filter Category Keys]: https://developer.apple.com/documentation/CoreImage/filter-category-keys
func (_CIFilterClass CIFilterClass) FilterNamesInCategory(category string) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("filterNamesInCategory:"), objc.String(category))
	return objc.ConvertSliceToStrings(rv)
}

// Publishes a custom filter that is not packaged as an image unit.
//
// name: A string object that specifies the name of the filter you want to publish.
//
// anObject: A constructor object that implements the `filterWithName` method.
//
// attributes: A dictionary that contains the class display name and filter categories
// attributes along with the appropriate value for each attributes. That is,
// the [kCIAttributeFilterDisplayName] attribute and a string that specifies
// the display name, and the [kCIAttributeFilterCategories] and an array that
// specifies the categories to which the filter belongs (such as
// [kCICategoryStillImage] and [kCICategoryDistortionEffect]). All other
// attributes for the filter should be returned by the custom `attributes`
// method implement by the filter.
//
// # Discussion
//
// In most cases you don’t need to use this method because the preferred way
// to register a custom filter that you write is to package it as an image
// unit. You do not need to use this method for a filter packaged as an image
// unit because you register your filter using the [CIPlugInRegistration]
// protocol. (See [Core Image Programming Guide] for additional details.)
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/registerName(_:constructor:classAttributes:)
//
// [kCIAttributeFilterCategories]: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterCategories
// [kCIAttributeFilterDisplayName]: https://developer.apple.com/documentation/CoreImage/kCIAttributeFilterDisplayName
// [kCICategoryDistortionEffect]: https://developer.apple.com/documentation/CoreImage/kCICategoryDistortionEffect
// [kCICategoryStillImage]: https://developer.apple.com/documentation/CoreImage/kCICategoryStillImage
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
func (_CIFilterClass CIFilterClass) RegisterFilterNameConstructorClassAttributes(name string, anObject CIFilterConstructor, attributes foundation.INSDictionary) {
	objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("registerFilterName:constructor:classAttributes:"), objc.String(name), anObject, attributes)
}

// Returns the localized name for the specified filter name.
//
// filterName: A filter name.
//
// # Return Value
//
// The localized name for the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedName(forFilterName:)
func (_CIFilterClass CIFilterClass) LocalizedNameForFilterName(filterName string) string {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("localizedNameForFilterName:"), objc.String(filterName))
	return foundation.NSStringFromID(rv).String()
}

// Returns the localized name for the specified filter category.
//
// category: A filter category.
//
// # Return Value
//
// The localized name for the filter category.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedName(forCategory:)
func (_CIFilterClass CIFilterClass) LocalizedNameForCategory(category string) string {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("localizedNameForCategory:"), objc.String(category))
	return foundation.NSStringFromID(rv).String()
}

// Returns the localized description of a filter for display in the user
// interface.
//
// filterName: The filter name.
//
// # Return Value
//
// The localized description of the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedDescription(forFilterName:)
func (_CIFilterClass CIFilterClass) LocalizedDescriptionForFilterName(filterName string) string {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("localizedDescriptionForFilterName:"), objc.String(filterName))
	return foundation.NSStringFromID(rv).String()
}

// Returns the location of the localized reference documentation that
// describes the filter.
//
// filterName: The filter name.
//
// # Return Value
//
// A URL that specifies the location of the localized documentation, or `nil`
// if the filter does not provide localized reference documentation.
//
// # Discussion
//
// The URL can be a local file or a remote document on a web server. Because
// filters created prior to OS X v10.5 could return `nil`, you should be make
// sure that your code handles this case gracefully.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/localizedReferenceDocumentation(forFilterName:)
func (_CIFilterClass CIFilterClass) LocalizedReferenceDocumentationForFilterName(filterName string) foundation.NSURL {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("localizedReferenceDocumentationForFilterName:"), objc.String(filterName))
	return foundation.NSURLFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAlphaWeightedHistogram()
func (_CIFilterClass CIFilterClass) AreaAlphaWeightedHistogramFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaAlphaWeightedHistogramFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaBoundsRed()
func (_CIFilterClass CIFilterClass) AreaBoundsRedFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaBoundsRedFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumScaleTransform()
func (_CIFilterClass CIFilterClass) MaximumScaleTransformFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("maximumScaleTransformFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/toneMapHeadroom()
func (_CIFilterClass CIFilterClass) ToneMapHeadroomFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("toneMapHeadroomFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAverageMaximumRed()
func (_CIFilterClass CIFilterClass) AreaAverageMaximumRedFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaAverageMaximumRedFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blurredRoundedRectangleGenerator()
func (_CIFilterClass CIFilterClass) BlurredRoundedRectangleGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blurredRoundedRectangleGeneratorFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/distanceGradientFromRedMask()
func (_CIFilterClass CIFilterClass) DistanceGradientFromRedMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("distanceGradientFromRedMaskFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedQRCodeGenerator()
func (_CIFilterClass CIFilterClass) RoundedQRCodeGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("roundedQRCodeGeneratorFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/signedDistanceGradientFromRedMask()
func (_CIFilterClass CIFilterClass) SignedDistanceGradientFromRedMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("signedDistanceGradientFromRedMaskFilter"))
	return CIFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/systemToneMap()
func (_CIFilterClass CIFilterClass) SystemToneMapFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("systemToneMapFilter"))
	return CIFilterFromID(rv)
}

// Transitions by folding and crossfading an image to reveal the target image.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the accordion fold transition filter to an image. The
// effect transitions from one image to another by unfolding and crossfading.
//
// The accordion fold transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `time`: A `float` representing
// the parametric time of the transition from start (at time 0) to end (at
// time 1) as an [NSNumber]. `numberOfFolds`: A `float` representing the
// number of accordion folds as a [NSNumber]. `foldShadowAmount`: A `float`
// representing the strength of the shadow as a [NSNumber].
//
// The following code creates a filter that produces folds in the input image
// and fades to the target image:
//
// [media-3616429]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/accordionFoldTransition()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) AccordionFoldTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("accordionFoldTransitionFilter"))
	return CIFilterFromID(rv)
}

// Blends colors from two images by addition.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// The filter calculates the sum of color components in the two input images
// to produce a brightening effect. People typically use this filter to add
// highlights and lens flares.
//
// The addition compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// that the filter applies the effect on with the type [CIImage].
//
// The following code creates a filter that combines the images’ colors to
// produce one image:
//
// [media-3546408]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/additionCompositing()
func (_CIFilterClass CIFilterClass) AdditionCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("additionCompositingFilter"))
	return CIFilterFromID(rv)
}

// Performs a transform on the image and extends the image edges to infinity.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the affine clamp filter to an image. This effect
// performs similarly to the affine transform filter except that it produces
// an infinite image. You can use this filter when you need to blur an image
// but you want to avoid a soft, black fringe along the edges.
//
// The affine clamp filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `transform`: A
// [CGAffineTransform] to constrain to the image.
//
// The following code creates a filter that produces a cropped image with
// colored edges to fill the rest of the image:
//
// [media-3624745]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/affineClamp()
//
// [CGAffineTransform]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
func (_CIFilterClass CIFilterClass) AffineClampFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("affineClampFilter"))
	return CIFilterFromID(rv)
}

// Performs a transform on the image and tiles the result.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the affine tile filter to an image. This effect
// performs an [CGAffineTransform] and then tiles the transformed image.
//
// The affine tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `transform`: A
// [CGAffineTransform] to apply to the image.
//
// The following code creates a filter that results in the image becoming
// tiled:
//
// [media-3624744]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/affineTile()
//
// [CGAffineTransform]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
func (_CIFilterClass CIFilterClass) AffineTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("affineTileFilter"))
	return CIFilterFromID(rv)
}

// Returns a 1 x 1 pixel image that contains the average color for the region
// of interest.
//
// # Return Value
//
// A 1 x 1 pixel image containing the average color for the region of
// interest.
//
// # Discussion
//
// This filter calculates the average color of the area defined by `extent`
// and creates a 1 x 1 pixel image with the result. The filter processes each
// color component (red, green, blue, alpha) of the input image independently.
//
// The area average filter uses the following properties:
//
// `inputImage`: The [CIImage] containing the image you want to process.
// `extent`: A [CGRect] that specifies the region of the image that you want
// to process.
//
// The following code creates a filter that calculates the average color of a
// 500 x 500 set of pixels from the center of the image:
//
// [media-4331783]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaAverage()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaAverageFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaAverageFilter"))
	return CIFilterFromID(rv)
}

// Returns a histogram of a specified area of the image.
//
// # Return Value
//
// A 1 pixel high image containing the calculated histogram`.`
//
// # Discussion
//
// This filter calculates histograms of the red, green, blue, and alpha colors
// in the region defined by `extent`. The `count` property controls the number
// of bins (or width) of the histogram. The filter scales the histogram so
// that the total of all the counts in the bins equals `scale`.
//
// The area histogram filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process. `scale`: The
// scale value to use for the histogram values. If the scale is 1, then the
// total of all the counts in the histogram equals 1. `count`: The number of
// bins for the histogram. This value determines the width of the output
// image. Minimum value 1, maximum value 2048.
//
// The following code creates a filter that results in an image that has a
// height of 1 pixel and a width of 256 pixels. The pixel color components
// contain the histogram values.
//
// To display the histogram, you can use the [HistogramDisplayFilter] filter:
//
// [media-4332392]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaHistogram()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaHistogramFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaHistogramFilter"))
	return CIFilterFromID(rv)
}

// Returns a logarithmic histogram of a specified area of the image.
//
// # Return Value
//
// A 1-pixel-high image containing the calculated histogram`.`
//
// # Discussion
//
// This filter calculates histograms of the `red,“green,“blue,` and `alpha`
// colors for the specified area of an image. A base two-logarithm function is
// applied to the values before binning. The `count` property controls the
// number of bins (or width) of the histogram. The histogram is scaled so that
// all the values sum to `scale`.
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image you want to process. `scale`: The
// scale value for the histogram values. If the scale is `1`, then the bins in
// the resulting image sum to `1`. `count`: The number of bins for the
// histogram. This value determines the width of the output image. Minimum
// value `1`, and maximum value `2048`. `minimumStop`: The minimum of the
// range of color channel values in the logarithmic histogram image. Defaults
// to `-10`. `maximumStop`: The maximum of the range of color channel values
// in the logarithmic histogram image. Defaults to 4.
//
// The following code creates a filter that results in a 1-pixel-tall image
// with a width of 256. The pixel color components contain the logarithmic
// histogram values:
//
// Use the [HistogramDisplayFilter] filter to display the histogram:
//
// [media-4407281]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaLogarithmicHistogram()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaLogarithmicHistogramFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaLogarithmicHistogramFilter"))
	return CIFilterFromID(rv)
}

// Finds the pixel with the highest alpha value.
//
// # Return Value
//
// A 1 x 1 size image containing the pixel with the maximum alpha value.
//
// # Discussion
//
// This filter returns the pixel with highest alpha value in the region
// defined by `extent`.
//
// The area maximum alpha filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that results in a single pixel image
// containing the pixel with the highest alpha value:
//
// [media-4332167]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMaximumAlpha()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMaximumAlphaFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMaximumAlphaFilter"))
	return CIFilterFromID(rv)
}

// Calculates the maximum color components of a specified area of the image.
//
// # Return Value
//
// A 1 x 1 size image containing the maximum color components.
//
// # Discussion
//
// This filter returns the maximum color components in the region defined by
// `extent`.
//
// The area maximum filter uses the following properties:.
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that calculates the maximum color
// components of a 500 x 500 set of pixels from the center of the image:
//
// [media-4331784]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMaximum()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMaximumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMaximumFilter"))
	return CIFilterFromID(rv)
}

// Calculates minimum and maximum color components for a specified area of the
// image.
//
// # Return Value
//
// A 2 x 1 pixel image containing the minimum and maximum color components.
//
// # Discussion
//
// This filter returns the maximum and minimum color components in the region
// defined by `extent`. The result is a 2 x 1 pixel image with the left pixel
// containing the minimum components and the right pixel containing the
// maximum components.
//
// The area min max filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that results in a 2 x 1 image where the
// minimum components are in the left pixel and the maximum components are in
// the right pixel.
//
// [media-4331786]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinMax()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMinMaxFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMinMaxFilter"))
	return CIFilterFromID(rv)
}

// Calculates the minimum and maximum red component value.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method applies the area-minimum-maximum-red filter to an image. This
// effect calculates the darkest and lightest red color value in the region
// defined by `extent`. The red and green components of the 1 x 1 pixel output
// image contain the result.
//
// The area-minimum-maximum-red filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that results in a 1 x 1 pixel image
// with the red and green color components populated:
//
// [media-4331787]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinMaxRed()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMinMaxRedFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMinMaxRedFilter"))
	return CIFilterFromID(rv)
}

// Calculates the pixel within a specified area that has the smallest alpha
// value.
//
// # Return Value
//
// A 1 x 1 pixel image containing the color with the smallest alpha value.
//
// # Discussion
//
// This method applies the area minimum alpha filter to an image. This effect
// finds and returns the pixel with the lowest alpha value in the region
// defined by `extent`.
//
// The area minimum alpha filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that results in a 1 x 1 pixel image
// containing the color with the lowest alpha value:
//
// [media-4332169]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinimumAlpha()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMinimumAlphaFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMinimumAlphaFilter"))
	return CIFilterFromID(rv)
}

// Calculates the minimum color component values for a specified area of the
// image.
//
// # Return Value
//
// A 1 x 1 size image containing the minimum color component values.
//
// # Discussion
//
// This filter returns the minimum color components in the region defined by
// `extent`.
//
// The area minimum filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that calculates the minimum color
// components of a 500 x 500 set of pixels from the center of the image:
//
// [media-4331785]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/areaMinimum()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) AreaMinimumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("areaMinimumFilter"))
	return CIFilterFromID(rv)
}

// Generates an attributed-text image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates an attributed-text image. The effect takes the input
// string property and the scale factor to scale up the text. You commonly
// combine this filter with other filters to create a watermark on images.
//
// The attributed-text image generator filter uses the following properties:
//
// `text`: An [NSAttributedString]. `scaleFactor`: A `float` representing the
// scale of the font to use for the generated text. padding: A `float`
// representing the value for an additional number of pixels to pad around the
// text’s bounding box.
//
// The following code creates a filter that generates an attributed-text
// image:
//
// [media-3546317]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/attributedTextImageGenerator()
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
func (_CIFilterClass CIFilterClass) AttributedTextImageGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("attributedTextImageGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Generates a low-density barcode.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates an Aztec code as an image. The Aztec barcode code is
// a low-density barcode defined in the ISO/IEC 24778:2008 standard. This code
// is commonly used for transport ticketing or boarding passes.
//
// The Aztec code generator filter uses the following properties:
//
// `message`: The data to be encoded as an Aztec code. An [NSData] object
// whose display name is Message. `correctionLevel`: A `float` representing
// the percentage of redundancy to add to the message data. A higher
// correction level allows the barcode to be correctly read even when
// partially damaged. `layers`: A `float` representing the number of
// concentric squares encoding the barcode data. When the value is zero, Core
// Image automatically determines the appropriate number of layers to encode
// the message data at the specified correction level. `compactStyle`: A
// `float` determining the use of compact or full-size Aztec barcode format.
// The compact format can store up to 44 bytes of message data in up to 4
// layers. The full-size format can store up to 1914 bytes of message data in
// up to 32 layers. Leave unset, or set to 0 for automatic.
//
// The following code creates a filter that generates an Aztec barcode:
//
// [media-3546310]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/aztecCodeGenerator()
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
func (_CIFilterClass CIFilterClass) AztecCodeGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("aztecCodeGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Generates a barcode as an image from the descriptor.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a custom barcode as an image. The effect uses barcode
// descriptors to specify properties of the generated barcode.
//
// The barcode generator uses the following property:
//
// `barcodeDescriptor`: An instance of [CIBarcodeDescriptor] with the input
// parameters supplied.
//
// The following code creates a filter that generates a QR code containing the
// text
//
// [media-4327881]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/barcodeGenerator()
func (_CIFilterClass CIFilterClass) BarcodeGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("barcodeGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Transitions between two images by removing rectangular portions of an
// image.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the bar swipe transition filter to an image. The effect
// transitions from one image to another by a series of moving bars passing
// over the target image.
//
// The bar swipe transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `time`: A `float` representing
// the parametric time of the transition from start (at time 0) to end (at
// time 1) as an [NSNumber]. `angle`: A `float` representing the angle of the
// motion as an [NSNumber]. `width`: A `float` representing the width of the
// bars in pixels as an [NSNumber]. `barOffset`: A `float` representing the
// offset of one bar in relation to others as a [NSNumber].
//
// The following code creates a filter that produces falling bars from the
// input image to transition to the target image:
//
// [media-3616431]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/barsSwipeTransition()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BarsSwipeTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("barsSwipeTransitionFilter"))
	return CIFilterFromID(rv)
}

// Produces a high-quality scaled version of an image.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the bicubic scale transform filter to an image. The
// effect produces a high-quality, scaled version of the input image. The
// parameters of [B] and [C] determine the sharpness and softness of the
// resampling.
//
// The bicubic scale transform filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `aspectRatio`: A `float`
// representing the aspect ratio as an [NSNumber]. `parameterB`: A `float`
// representing the value of B used for cubic resampling as an [NSNumber].
// `parameterC`: A `float` representing the value of C used for cubic
// resampling as an [NSNumber].
//
// The following code creates a filter that results in the image becoming
// square:
//
// [media-3582224]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bicubicScaleTransform()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BicubicScaleTransformFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("bicubicScaleTransformFilter"))
	return CIFilterFromID(rv)
}

// Blends two images by using an alpha mask image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the blend with alpha mask filter to an image. The
// effect uses values from the grayscale mask image to interpolate between the
// input and background images. The mask image consists of shades of gray that
// define the strength of the interpolation from zero (where the mask image is
// black) to the specified `radius` (where the mask image is white).
//
// The blend with alpha mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `maskImage`: An image that
// masks an area on the background image with the type [CIImage].
// `backgroundImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the replacement of
// white in the mask image with the detail of the input image:
//
// [media-3624591]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithAlphaMask()
func (_CIFilterClass CIFilterClass) BlendWithAlphaMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blendWithAlphaMaskFilter"))
	return CIFilterFromID(rv)
}

// Blends two images by using a blue mask image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the blend with blue mask filter to an image. The effect
// uses values from the blue mask image to interpolate between the input and
// background images. The mask image is made of shades of blue that define the
// strength of the interpolation from zero (where the mask image is black) to
// the specified `radius` (where the mask image is blue).
//
// The blend with blue mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `maskImage`: An image that
// masks an area on the background image with the type [CIImage].
// `backgroundImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the replacement of blue
// in the mask image with the detail of the input image:
//
// [media-3624592]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithBlueMask()
func (_CIFilterClass CIFilterClass) BlendWithBlueMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blendWithBlueMaskFilter"))
	return CIFilterFromID(rv)
}

// Blends two images by using a mask image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the blend with mask filter to an image. The effect uses
// values from the green mask image to interpolate between the input and
// background images. The mask image consists of shades of green that define
// the strength of the interpolation from zero (where the mask image is black)
// to the specified `radius` (where the mask image is green).
//
// The blend with mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `maskImage`: An image that
// masks an area on the background image with the type [CIImage].
// `backgroundImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the replacement of
// green in the mask image with the detail of the input image:
//
// [media-3624594]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithMask()
func (_CIFilterClass CIFilterClass) BlendWithMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blendWithMaskFilter"))
	return CIFilterFromID(rv)
}

// Blends two images by using a red mask image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the blend with red mask filter to an image. The effect
// uses values from the red mask image to interpolate between the input and
// background images. The mask image consists of shades of red that define the
// strength of the interpolation from zero (where the mask image is black) to
// the specified `radius` (where the mask image is red).
//
// The blend with red mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `maskImage`: An image that
// masks an area on the background image with the type [CIImage].
// `backgroundImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the replacement of
// green in the mask image with the detail of the input image:
//
// [media-3624593]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blendWithRedMask()
func (_CIFilterClass CIFilterClass) BlendWithRedMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blendWithRedMaskFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s colors by applying a blur effect.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the bloom filter to an image. The effect softens edges
// and adds a slight blur effect.
//
// The bloom filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber]. `intensity`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in a hazy effect on the
// image:
//
// [media-3599995]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bloom()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BloomFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("bloomFilter"))
	return CIFilterFromID(rv)
}

// Generates a blurred rectangle.
//
// # Return Value
//
// A [CIImage] containing a blurred rectangle.
//
// # Discussion
//
// Creates a [CIImage] containing a blurred rectangle. The resulting image
// size is the extent of the rectangle plus any additional space required for
// the blur effect.
//
// The blurred rectangle filter uses the following properties:
//
// `extent`: A [CGRect] that defines the extent of the effect. `color`: A
// [CIColor] specifying the color of the rectangle. `sigma`: A `float`
// specifying the sigma for the Gaussian blur.
//
// The following code creates a filter that generates a blurred red rectangle
// with a width of 200 x 100 pixels.
//
// [media-4407302]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/blurredRectangleGenerator()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) BlurredRectangleGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("blurredRectangleGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Applies a bokeh effect to an image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the bokeh blur filter to an image. The effect targets a
// circular area of pixels defined by the `radius` and blurs the area. The
// filter adds smaller intense blur rings.
//
// The bokeh blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `ringSize`: A `float` representing the ring size of the bokeh as an
// [NSNumber]. `ringAmount`: A `float` representing the emphasis at the ring
// of the bokeh as an [NSNumber]. `softness`: A `float` representing the
// softness of the bokeh effect as an [NSNumber]. `inputImage`: A [CIImage]
// representing the input image to apply the filter to.
//
// The following code creates a filter that adds a softer blur to the input
// image:
//
// [media-3544959]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bokehBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BokehBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("bokehBlurFilter"))
	return CIFilterFromID(rv)
}

// Applies a square-shaped blur to an area of an image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the box blur filter to an image. The effect targets a
// square area and calculates the median color value of the pixels to create
// the output image. The `radius` is the width of a square area with a larger
// area, resulting in a stronger blur effect on the output image.
//
// The box blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as a [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that results in less detail in the
// input image:
//
// [media-3544956]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/boxBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BoxBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("boxBlurFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image with a concave or convex bump.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the bump distortion filter to an image. This effect
// creates a concave or convex bump defined by the `scale`. A value of 0.0 has
// no effect, while a positive value creates an outward curvature and a
// negative value creates an inward curvature.
//
// The bump distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the amount of pixels the filter uses to create the distortion
// as an [NSNumber]. `center`: A [CGPoint] representing the center of the
// effect. `scale`: A `float` representing the curvature of the bump effect as
// an [NSNumber].
//
// The following code creates a filter that results in a concave bump
// distorting the image:
//
// [media-4407303]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bumpDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BumpDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("bumpDistortionFilter"))
	return CIFilterFromID(rv)
}

// Linearly distorts an image with a concave or convex bump.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the bump distortion linear filter to an image. This
// effect creates a concave or convex bump to a linear portion of the image.
// The curvature of the bump is defined by the `scale` property. A value of
// 0.0 has no effect, while a positive value creates an outward curvature and
// a negative value creates an inward curvature.
//
// The bump distortion linear filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the amount of pixels the filter uses to create the distortion
// as an [NSNumber]. `center`: A set of coordinates marking the center of the
// image as a [CGPoint]. `scale`: A `float` representing the curvature of the
// bump effect as an [NSNumber]. `angle`: A `float` representing the angle of
// the distortion, in radians, as an [NSNumber].
//
// The following code creates a filter that results in a vertical bump
// distorting the image:
//
// [media-4407320]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/bumpDistortionLinear()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) BumpDistortionLinearFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("bumpDistortionLinearFilter"))
	return CIFilterFromID(rv)
}

// Adds a series of colorful dots to an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a CMYK halftone filter to an image. The effect
// generates an image containing a series of dots. The dots contain only cyan,
// magenta, yellow, and black colors. Halftone effect is a set of lines, dots,
// or circles that contain detail. When viewing the image from a distance, the
// markings blend together creating the illusion of continuous lines and
// shapes. Print media commonly uses this effect.
//
// The CMYK halftone filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `angle`: A `float`
// representing the angle of the pattern as an [NSNumber]. `width`: A `float`
// representing the distance between dots in the pattern as an [NSNumber].
// `sharpness`: A `float` representing the sharpness of the pattern as an
// [NSNumber]. `center`: A set of coordinates marking the center of the image
// as a [CGPoint]. `grayComponentReplacement`: A `float` representing the grey
// component to be replaced as an [NSNumber]. `underColorRemoval`: A `float`
// representing the under-color removal value as an [NSNumber].
//
// The following code produces an image with visible dots and less color:
//
// [media-3595920]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/cmykHalftone()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CMYKHalftone() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("CMYKHalftone"))
	return CIFilterFromID(rv)
}

// Applies the Canny edge-detection algorithm to an image.
//
// # Return Value
//
// A [CIImage] with the detected edges.
//
// # Discussion
//
// This filter performs a Canny edge-detection on the input image, producing a
// black-and-white image with the detected edges. White pixels indicate an
// edge, and black pixels indicate no edge.
//
// The Canny edge-detection filter uses the following properties:
//
// `inputImage`: The [CIImage] to use as an input for the effect.
// `gaussianSigma`: A `float` specifying the sigma of the Gaussian blur to
// apply, reducing high-frequency noise. Defaults to `1.6`. `perceptual`: A
// [Boolean] specifying whether to use a perceptual color space to compute the
// edge thresholds. Defaults to `false`. `thresholdLow`: A `float` specifying
// the threshold for weak edges. Defaults to `0.02`. `thresholdHigh`: A
// `float` specifying the threshold for strong edges. Defaults to `0.05`.
// `hysteresisPasses`: The number of hysteresis passes to apply to promote
// weak edge pixels. Minimum value is `0`, maximum value is `20`, and defaults
// to `1`.
//
// The following code applies Canny edge-detection to an image:
//
// [media-4407284]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/cannyEdgeDetector()
func (_CIFilterClass CIFilterClass) CannyEdgeDetectorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("cannyEdgeDetectorFilter"))
	return CIFilterFromID(rv)
}

// Generates a checkerboard image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a checkerboard pattern as an image. The effect
// requires the size, sharpness, and color properties to create the pattern.
//
// The checkerboard generator filter uses the following properties:
//
// `center`: A `vector` representing the center of the image as a [CIVector].
// `color0`: A [CIColor] representing the first color of the pattern.
// `color1`: A [CIColor] representing the second color of the pattern.
// `sharpness`: A `float` representing the sharpness of the pattern as an
// [NSNumber]. `width`: A `float` representing the width of the checkerboard
// squares as an [NSNumber].
//
// The following code creates a filter that generates a black-and-white
// checkered pattern:
//
// [media-3590970]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/checkerboardGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CheckerboardGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("checkerboardGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image with radiating circles to the periphery of the image.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the circle splash distortion filter to an image. This
// effect distorts the pixels starting at the circumference of a circle and
// emanating outward.
//
// The circle splash distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the amount in pixels the filter uses to create the distortion
// as an [NSNumber]. `center`: A [CGPoint] representing the center of the
// image.
//
// The following code creates a filter that results in a ripple effect applied
// to the image:
//
// [media-4407306]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circleSplashDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CircleSplashDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("circleSplashDistortionFilter"))
	return CIFilterFromID(rv)
}

// Adds a circular overlay to an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a circular screen filter to an image. The effect
// generates a monochrome image containing a series of circular rings. The
// halftone effect is a set of lines, dots, or circles that contain detail.
// When viewing the image from a distance, the markings blend together,
// creating the illusion of continuous lines and shapes. Print media commonly
// uses this effect.
//
// The circular screen filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `width`: A
// `float` representing the distance between each circle in the pattern as an
// [NSNumber]. `sharpness`: A `float` representing the sharpness of the
// circles in the pattern as an [NSNumber].
//
// The following code creates a filter that results in a monochrome image with
// a large circular pattern overlaying the image:
//
// [media-3595914]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circularScreen()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CircularScreenFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("circularScreenFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by increasing the distance of the center of the image.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the circular wrap filter to an image. This effect wraps
// an image around a transparent circle. The distortion of the image increases
// with the distance from the center of the circle.
//
// The circular wrap filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `angle`: A `float`
// representing the angle of the wrap, in radians, as an [NSNumber]. `radius`:
// A `float` representing the amount of pixels the filter uses to create the
// distortion as an [NSNumber]. `center`: A set of coordinates marking the
// center of the image as a [CGPoint].
//
// The following code creates a filter that results in a circular image
// generated from the input image:
//
// [media-4407319]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/circularWrap()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CircularWrapFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("circularWrapFilter"))
	return CIFilterFromID(rv)
}

// Generates a high-density, linear barcode.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a Code 128 barcode as an image. Code 128 is a
// high-density linear barcode defined in the ISO/IEC 15417:2007 standard. Use
// this filter to generate alphanumeric or numeric-only barcodes. The barcode
// can contain any of the 128 ASCII characters.
//
// The Code 128 barcode filter uses the following properties:
//
// `message`: [NSData] containing the message to encode in the Code 128
// barcode. `quietSpace`: [NSNumber] containing the number of empty white
// pixels that should surround the barcode. `barcodeHeight`: [NSNumber]
// containing the height of the generated barcode in pixels.
//
// The following code creates a filter that generates a Code 128 barcode:
//
// [media-3546314]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/code128BarcodeGenerator()
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Code128BarcodeGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("code128BarcodeGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Calculates the absolute difference between each color component in the
// input images.
//
// # Return Value
//
// An image containing the absolute color difference between the two input
// images.
//
// # Discussion
//
// This method applies the color absolute difference filter to an image. This
// filter calculates the absolute color difference of the red, green, and blue
// values between the two input images. The alpha channel is the product of
// the alpha channels from the two input images.
//
// The absolute difference filter uses the following properties:
//
// `inputImage`: The first [CIImage] for differencing. `inputImage2`: The
// second [CIImage] for differencing.
//
// The following code creates a filter that results in the color difference
// between two images:
//
// [media-4332166]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorAbsoluteDifference()
func (_CIFilterClass CIFilterClass) ColorAbsoluteDifferenceFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorAbsoluteDifferenceFilter"))
	return CIFilterFromID(rv)
}

// Blends color from two images using the luminance values from the background
// image and the hue and saturation values from the input image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color-blend mode filter to an image. The effect
// creates the output image by using the luminance values of the background
// image, while using the hue and saturation values of the input image. The
// effect preserves gray levels from the background image.
//
// The color-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that replaces the colors the image:
//
// [media-3546407]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorBlendMode()
func (_CIFilterClass CIFilterClass) ColorBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Blends color from two images while darkening the image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color-burn blend mode filter to an image. The
// effect darkens the background image to reflect the input image’s color.
//
// The color-burn blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that darkens the background image:
//
// [media-3546399]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorBurnBlendMode()
func (_CIFilterClass CIFilterClass) ColorBurnBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorBurnBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Alters the colors in an image based on color components.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color clamp filter to an image. The effect
// calculates each pixel’s color component value. Using this calculation,
// the effect adjusts the values that are outside the range of the
// `minComponents` or `maxComponents` properties and clamps them within the
// range.
//
// The color clamp filter uses the following properties:
//
// `minComponents`: [RGBA] values for the lower end of the range as a
// [CIVector]. `maxComponents`: [RGBA] values for the upper end of the range
// as a [CIVector]. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a red hue to the input image:
//
// [media-3545003]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorClamp()
func (_CIFilterClass CIFilterClass) ColorClampFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorClampFilter"))
	return CIFilterFromID(rv)
}

// Alters the brightness, contrast, and saturation of an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color controls filter to an image. The effect
// calculates saturation by linearly interpolating between a grayscale image
// with a saturation of `0.0` and the original image saturation of `1.0.`
//
// The color controls filter uses the following properties:
//
// `brightness`: A `float` representing the amount of brightness applied as a
// [NSNumber]. `contrast`: A `float` `r`epresenting the amount of contrast
// applied as a [NSNumber]. `saturation`: A float representing the amount of
// saturation applied as a [NSNumber]. `inputImage`: An image with the type
// [CIImage].
//
// The following code creates a filter that results in a darker image:
//
// [media-3545002]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorControls()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ColorControlsFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorControlsFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s color by applying polynomial cross-products.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color cross polynomial filter to an image. The
// effect targets each pixel individually and calculates the coefficients for
// the r`ed`, `green`, and `blue` channels according to the polynomial cross
// product.
//
// The color cross-polynomial filter uses the following properties:
//
// `redCoefficients`: A [CIVector] representing the polynomial coefficients
// for the red channel. `blueCoefficients`: A [CIVector] representing the
// polynomial coefficients for the blue channel. `greenCoefficients`: A
// [CIVector] representing polynomial coefficients for the green channel.
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a green hue to the input
// image:
//
// [media-3545029]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCrossPolynomial()
func (_CIFilterClass CIFilterClass) ColorCrossPolynomialFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorCrossPolynomialFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s pixels using a three-dimensional color table.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color cube filter to an image. The effect maps
// color values in the input image to new color values using a
// three-dimensional color look-up table, also called a color cube. For each
// [RGBA] pixel in the input image, the filter uses the pixel’s `red`,
// `green`, and `blue` component values to identify a location in the table.
// The [RGBA] value at that location becomes the [RGBA] value of the output
// pixel.
//
// The color cube filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `cubeData`: Data containing
// a 3-dimensional color table of floating-point premultiplied RGBA values.
// The cells are organized in a standard ordering. The columns and rows of the
// data are indexed by red and green, respectively. Each data plane is
// followed by the next higher plane in the data, with planes indexed by blue.
// `extrapolate`: If `true`, then the filter extrapolates the color cube for
// any RGB component values outside the range 0.0 to 1.0. `cubeDimension`: The
// dimension of the color cube.
//
// The following code creates a filter that adds a blue hue to the input
// image:
//
// [media-3545026]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCube()
func (_CIFilterClass CIFilterClass) ColorCubeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorCubeFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s pixels using a three-dimensional color table in
// specified color space.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color cube with color space filter to an image. The
// effect applies a mapping from `rgb` space within the color space defined to
// color values the cubeData defines. For each pixel, it matches the data and
// adjusts the color on the output image.
//
// The color cube with color space filter uses the following properties:
//
// `cubeData`: Data containing a 3-dimensional color table of floating-point
// premultiplied RGBA values. The cells are organized in a standard ordering.
// The columns and rows of the data are indexed by red and green,
// respectively. Each data plane is followed by the next higher plane in the
// data, with planes indexed by blue. `colorSpace`: A [CGColorSpace]
// representing the color space for the color cube. `cubeDimension`: The
// dimension of the color cube. `inputImage`: An image with the type
// [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545016]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCubeWithColorSpace()
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
func (_CIFilterClass CIFilterClass) ColorCubeWithColorSpaceFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorCubeWithColorSpaceFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s pixels using a three-dimensional color tables and a
// mask image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color cubes mixed with mask filter to an image. The
// effect uses two color cube tables to modify the input image. The filter
// uses the mask image to interpolate between the two color cubes.
//
// The color cubes mixed with mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `maskImage`: A mask image
// with the type [CIImage]. `cube0Data`: Data containing a 3-dimensional color
// table of floating-point premultiplied RGBA values. The cells are organized
// in a standard ordering. The columns and rows of the data are indexed by red
// and green, respectively. Each data plane is followed by the next higher
// plane in the data, with planes indexed by blue. `cube1Data`: Data
// containing a 3-dimensional color table of floating-point premultiplied RGBA
// values. The cells are organized in a standard ordering. The columns and
// rows of the data are indexed by red and green, respectively. Each data
// plane is followed by the next higher plane in the data, with planes indexed
// by blue. `colorSpace`: A [CGColorSpace] representing the color space for
// the color cubes. `cubeDimension`: The dimension of the color cubes
//
// The following code creates a filter that adds colors from the mask image
// and brightness to the input image:
//
// [media-3546428]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCubesMixedWithMask()
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
func (_CIFilterClass CIFilterClass) ColorCubesMixedWithMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorCubesMixedWithMaskFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s color curves.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color curves filter to an image. The effect uses a
// three-channel one-dimensional color table to transform the source image
// pixels. The color table must be comprised of floating-point RGB value.
//
// The color curves filter uses the following properties:
//
// `colorSpace`: A [CGColorSpace] representing the color space for the color
// curve. `curvesData`: Data containing a color table of floating-point RGB
// values as [NSData]. `curvesDomain`: A two-element vector that defines the
// minimum and maximum values of the curve data as a [CIVector]. `inputImage`:
// An image with the type [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545017]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorCurves()
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
func (_CIFilterClass CIFilterClass) ColorCurvesFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorCurvesFilter"))
	return CIFilterFromID(rv)
}

// Blends color from two images using dodging.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color-dodge blend mode filter to an image. The
// effect divides the background image color values by the inverted color
// values of the input image. The effect then decreases contrast between the
// input image and the background image.
//
// The color-dodge mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that darkens the background image:
//
// [media-3546418]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorDodgeBlendMode()
func (_CIFilterClass CIFilterClass) ColorDodgeBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorDodgeBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Inverts an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that invert the colors
// of an image.
//
// The color-invert filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that inverts the colors of the input
// image:
//
// [media-3545010]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorInvert()
func (_CIFilterClass CIFilterClass) ColorInvertFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorInvertFilter"))
	return CIFilterFromID(rv)
}

// Performs a transformation of the input image colors to colors from a
// gradient image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a color map filter to an image. The effect transforms
// source color values by converting the unpremultiplied RGB values to luma
// using the weighting `(0.2125, 0.7154, 0.0721)`. The luma value is then used
// to look up the new color from the gradient image.
//
// The color map filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `gradientImage`: An image
// representing the gradient of colors to be mapped to the input image colors
// with the type [CIImage].
//
// The following code creates a filter that adds the gradient image colors to
// the input image:
//
// [media-3558775]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMap()
func (_CIFilterClass CIFilterClass) ColorMapFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorMapFilter"))
	return CIFilterFromID(rv)
}

// Alters the colors in an image based on vectors provided.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color matrix filter to an image. The effect
// calculates the color matrix by multiplying the vector properties with the
// color values from the input image.
//
// The color matrix filter uses the following properties:
//
// `rVector`: A [CIVector] representing the amount of red to multiply the
// source color values by. `gVector`: A [CIVector] representing the amount of
// green to multiply the source color values by. `bVector`: A [CIVector]
// representing the amount of blue to multiply the source color values by.
// `aVector`: A [CIVector] representing the amount of alpha to multiply the
// source color values by. `biasVector`: A [CIVector] representing the amount
// of each vector that’s added to each color component. `inputImage`: An
// image with the type [CIImage].
//
// The following code creates a filter that adds a green hue to the input
// image:
//
// [media-3544998]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMatrix()
func (_CIFilterClass CIFilterClass) ColorMatrixFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorMatrixFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s colors to shades of a single color.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color monochrome filter to an image. The effect
// remaps the colors of the image to shades of the specified color.
//
// The color monochrome filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `color`: The color to map
// the input image colors to, as a [CIColor]. `intensity`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in the colors of the image
// becoming shades of red:
//
// [media-3545012]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorMonochrome()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ColorMonochromeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorMonochromeFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color polynomial filter to an image. The effect
// calculates the sum of each pixel’s color component value and the
// coefficient properties together to produce the output image.
//
// The color polynomial filter uses the following properties:
//
// `redCoefficients`: A vector representing the polynomial coefficients for
// the red channel as a [CIVector]. `greenCoefficients`: A vector representing
// the polynomial coefficients for the green channel as a [CIVector].
// `blueCoefficients`: A vector representing the polynomial coefficients for
// the blue channel as a [CIVector]. `alphaCoefficients`: A vector
// representing the polynomial coefficients for the alpha channel as a
// [CIVector]. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a lighter contrast to the
// input image:
//
// [media-3545009]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorPolynomial()
func (_CIFilterClass CIFilterClass) ColorPolynomialFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorPolynomialFilter"))
	return CIFilterFromID(rv)
}

// Flattens an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the color posterize filter to an image. The effect
// remaps red, green, and blue color components to a specified brightness
// value. The effect mimics the look of a silk-screened poster.
//
// The color posterize filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. levels: A `float`
// representing the brightness level as an [NSNumber].
//
// The following code creates a filter that flattens the colors in the input
// image:
//
// [media-3545025]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorPosterize()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ColorPosterizeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorPosterizeFilter"))
	return CIFilterFromID(rv)
}

// Compares the red, green, and blue components of the input image to a
// threshold and sets them to 1 or 0.
//
// # Return Value
//
// An image containing pixels with color components that are either 1 or 0.
//
// # Discussion
//
// This method applies the color threshold filter to an image. The filter
// compares the value of each color component (red, green, and blue) in the
// image against the threshold value. Any component higher than the threshold
// becomes 1 and any component lower becomes 0. The alpha component remains
// unchanged.
//
// The color threshold filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `threshold`: A `float`
// representing the threshold of color values as an [NSNumber].
//
// The following code creates a filter that results in an image where each
// color component is either 1 or 0.
//
// [media-4331780]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorThreshold()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ColorThresholdFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorThresholdFilter"))
	return CIFilterFromID(rv)
}

// Compares the red, green, and blue components of the input image against a
// threshold calculated using Otsu’s algorithm.
//
// # Return Value
//
// An image containing pixels with color components that are either 1 or 0.
//
// # Discussion
//
// The filter applies Otsu’s algorithm to the reg, green, and blue color
// components. The filter uses these thresholds to set the color to components
// to 1 or 0. The alpha component remains unchanged.
//
// The color threshold Otsu filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in an image where each
// color component is either 1 or 0:
//
// [media-4407336]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/colorThresholdOtsu()
func (_CIFilterClass CIFilterClass) ColorThresholdOtsuFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("colorThresholdOtsuFilter"))
	return CIFilterFromID(rv)
}

// Calculates the average color for a specified column of an image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method applies the column average filter to an image. This effect
// calculates the average color for a vertical column over a region defined by
// `extent`. The width of the resulting image is set by the width of the
// `extent`. The height of the resulting image is always 1 pixel.
//
// The column average filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates an image containing the average values in the
// columns from the middle of the image:
//
// [media-4331782]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/columnAverage()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) ColumnAverageFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("columnAverageFilter"))
	return CIFilterFromID(rv)
}

// Creates an image with a comic book effect.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the comic effect filter to an image. The effect
// simulates a comic book drawing by outlining edges and applying a color
// halftone effect.
//
// The comic effect filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the image appearing
// image in comic book style.
//
// [media-3601088]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/comicEffect()
func (_CIFilterClass CIFilterClass) ComicEffectFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("comicEffectFilter"))
	return CIFilterFromID(rv)
}

// Converts an image from CIELAB to RGB color space.
//
// # Return Value
//
// The converted [CIImage].
//
// # Discussion
//
// This filter converts an image from CIELAB color space to RGB. The CIELAB
// color space expresses color as three values: L* for the perceptual
// lightness, and a*b* for the colors red, green, blue, and yellow. The RGB
// color space expresses colors using the intensities of the three primary
// colors: red, green, and blue.
//
// `inputImage`: A [CIImage] containing the [RGB] image. `normalize`: If true,
// the three input channels are in the range 0 to 1. If false, the L* channel
// is in the range 0 to 100, and the a*b* channels are in the range -128 to
// 128.
//
// The following code applies the `convertLabToRGBFilter` to an image with the
// `normalize` flag set to the `true`:
//
// [media-4407298]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convertLabToRGB()
func (_CIFilterClass CIFilterClass) ConvertLabToRGBFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convertLabToRGBFilter"))
	return CIFilterFromID(rv)
}

// Converts an image from RGB to CIELAB color space.
//
// # Return Value
//
// The converted [CIImage].
//
// # Discussion
//
// This filter converts an image from RGB to CIELAB color space. The CIELAB
// color space expresses color as three values: L* for the perceptual
// lightness, and a*b* for the colors red, green, blue, and yellow. The RGB
// color space expresses colors using the intensities of the three primary
// colors: red, green, and blue.
//
// `inputImage`: A [CIImage] containing the [RGB] image. `normalize`: If true,
// the three output channels are in the range 0 to 1. If false, the L* channel
// is in the range 0 to 100 and the a*b* channels are in the range -128 to
// 128.
//
// The following code applies the `convertRGBToLabFilter` to an image with the
// `normalize` flag set to the `true`:
//
// [media-4407334]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convertRGBtoLab()
func (_CIFilterClass CIFilterClass) ConvertRGBtoLabFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convertRGBtoLabFilter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 3 x 3 filter to the [RGBA] components of an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a 3 x 3 convolution to the [RGBA] components of an
// image. The effect uses a 3 x 3 area surrounding an input pixel, the pixel
// itself, and those within a distance of 1 pixel horizontally and vertically.
// The effect repeats this for every pixel within the image. The work area is
// then combined with the weight property vector to produce the processed
// image. This filter differs from the [ConvolutionRGB3X3Filter], which only
// processes the [RGB] color components.
//
// The convolution 3 x 3 filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel as a [NSNumber]. `weights`: A [CIVector] representing the convolution
// kernel. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that sharpens the input image:
//
// [media-4334869]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution3X3()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Convolution3X3Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolution3X3Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 5 x 5 filter to the [RGBA] components image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a 5 x 5 convolution to the [RGBA] components of an
// image. The effect uses a 5 x 5 area surrounding an input pixel, the pixel
// itself, and those within a distance of 2 pixels horizontally and
// vertically. The effect repeats this for every pixel within the image. The
// work area is then combined with the weight property vector to produce the
// processed image. This filter differs from the [ConvolutionRGB5X5Filter]
// filter, which only processes the RGB components.
//
// The convolution 5 x 5 filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel as a [NSNumber]. `weights`: A [CIVector] representing the convolution
// kernel. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that blurs the input image:
//
// [media-4334867]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution5X5()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Convolution5X5Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolution5X5Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 7 x 7 filter to the [RGBA] color components of an
// image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the convolution 7 x 7 filter to an image. The effect
// uses a 7 x 7 area surrounding an input pixel, the pixel itself, and those
// within a distance of 3 pixels horizontally and vertically. The effect
// repeats this for every pixel within the image. The work area is then
// combined with the weight property vector to produce the processed image.
// This filter differs from the [ConvolutionRGB7X7Filter] filter, which only
// processes the RGB components.
//
// The convolution 7 x 7 filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel as a [NSNumber]. `weights`: A [CIVector] representing the convolution
// kernel. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that highlights edges in the input
// image:
//
// [media-4334871]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution7X7()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Convolution7X7Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolution7X7Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution-9 horizontal filter to the [RGBA] components of an
// image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a 9 x 1 convolution to the [RGBA] components of an
// image. The effect uses a 9 x 1 area surrounding an input pixel, the pixel
// itself, and those within a distance of 4 pixels horizontally. The effect
// repeats this for every pixel within the image. Unlike the convolution
// filters, which use square matrices, this filter can only produce effects
// along a horizontal axis. You can combine this filter with the
// [Convolution9VerticalFilter] to apply separable 9 x 9 convolutions.
//
// The convolution 9-horizontal filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel as a [NSNumber]. `weights`: A [CIVector] representing the convolution
// kernel. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that detects edges in the input image:
//
// [media-4334870]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution9Horizontal()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Convolution9HorizontalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolution9HorizontalFilter"))
	return CIFilterFromID(rv)
}

// Applies a convolution-9 vertical filter to the [RGBA] components of an
// image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a 1 x 9 convolution filter to the [RGBA] components of
// an image. The effect uses a 1 x 9 area surrounding an input pixel, the
// pixel itself, and those within a distance of 4 pixels vertically. The
// effect repeats this for every pixel within the image. Unlike the
// convolution filters, which use square matrices, this filter can only
// produce effects along a vertical axis. You can combine this filter with the
// [Convolution9HorizontalFilter] to apply separable 9 x 9 convolutions.
//
// The convolution-9-vertical filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel as a [NSNumber]. `weights`: A [CIVector] representing the convolution
// kernel. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that detects edges in the input image:
//
// [media-4334868]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolution9Vertical()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) Convolution9VerticalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolution9VerticalFilter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 3 x 3 filter to the [RGB] components of an image.
//
// # Return Value
//
// The convolved image.
//
// # Discussion
//
// This method applies a 3 x 3 convolution to the [RGB] components of an
// image. The effect uses a 3 x 3 area surrounding an input pixel, the pixel
// itself, and those within a distance of 1 pixel horizontally and vertically.
// This filter differs from the [Convolution3X3Filter] filter, which processes
// all of the color components including the alpha component.
//
// The convolution-RGB 3 x 3 filter uses the following properties:
//
// `bias`: A `float` representing the value that’s added to each output
// pixel. `weights`: A [CIVector] representing the convolution kernel.
// `inputImage`: A [CIImage] containing the image to process.
//
// The following code creates a filter that sharpens the input image:
//
// [media-4407321]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB3X3()
func (_CIFilterClass CIFilterClass) ConvolutionRGB3X3Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolutionRGB3X3Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 5 x 5 filter to the [RGB] components of an image.
//
// # Return Value
//
// This method applies a 5 x 5 convolution to the [RGB] components image. The
// effect uses a 5 x 5 area surrounding an input pixel, the pixel itself, and
// those within a distance of two pixels horizontally and vertically. The
// effect repeats this for every pixel within the image. The work area is then
// combined with the weight property vector to produce the processed image.
// This filter differs from the [Convolution5X5Filter] filter, which processes
// all of the color components including the alpha component.
//
// # Discussion
//
// The convolution-RGB 5 x 5 filter uses the following properties:
//
// `inputImage`: A [CIImage] containing the image to process. `weights`: A
// [CIVector] representing the convolution kernel. `bias`: A `float`
// representing the value that’s added to each output pixel.
//
// The following code creates a filter that applies an unsharp kernel to the
// input image:
//
// [media-4407325]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB5X5()
func (_CIFilterClass CIFilterClass) ConvolutionRGB5X5Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolutionRGB5X5Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 7 x 7 filter to the RGB components of an image.
//
// # Return Value
//
// The convolved image.
//
// # Discussion
//
// This method applies a 7 x 7 convolution to the [RGB] components image. The
// effect uses a 7 x 7 area surrounding an input pixel, the pixel itself, and
// those within a distance of 3 pixels horizontally and vertically. The effect
// repeats this for every pixel within the image. The work area is then
// combined with the weight property vector to produce the processed image.
// This filter differs from the [Convolution7X7Filter] filter, which processes
// all of the color components including the alpha component.
//
// The convolution-RGB 7 x 7 filter uses the following properties:
//
// `inputImage`: A [CIImage] containing the image to process. `weights`: A
// [CIVector] representing the convolution kernel. `bias`: A `float`
// representing the value that’s added to each output pixel.
//
// The following code creates a filter that highlights edges in the input
// image:
//
// [media-4407333]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB7X7()
func (_CIFilterClass CIFilterClass) ConvolutionRGB7X7Filter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolutionRGB7X7Filter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 9 x 1 filter to the RGB components of an image.
//
// # Return Value
//
// The convolved image.
//
// # Discussion
//
// This method applies a 9 x 1 convolution to the [RGB] components of an
// image. The effect uses a 9 x 1 area surrounding an input pixel, the pixel
// itself, and those within a distance of 4 pixels horizontally. The effect
// repeats this for every pixel within the image. Unlike the convolution
// filters, which use square matrices, this filter can only produce effects
// along a vertical axis. You can combine this filter with the
// [ConvolutionRGB9VerticalFilter] to apply separable 9 x 9 convolutions. This
// filter differs from the [Convolution9HorizontalFilter] filter, which
// processes all of the color components including the alpha component.
//
// The convolution-RGB-9-vertical filter uses the following properties:
//
// `inputImage`: A [CIImage] containing the image to process. `weights`: A
// [CIVector] representing the convolution kernel. `bias`: A `float`
// representing the value that’s added to each output pixel.
//
// The following code creates a filter that blurs the image in the horizontal
// direction:
//
// [media-4407285]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB9Horizontal()
func (_CIFilterClass CIFilterClass) ConvolutionRGB9HorizontalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolutionRGB9HorizontalFilter"))
	return CIFilterFromID(rv)
}

// Applies a convolution 1 x 9 filter to the RGB components of an image.
//
// # Return Value
//
// This method applies a 1 x 9 convolution to the [RGB] components of an
// image. The effect uses a 1 x 9 area surrounding an input pixel, the pixel
// itself, and those within a distance of 4 pixels vertically. The effect
// repeats this for every pixel within the image. Unlike the convolution
// filters, which use square matrices, this filter can only produce effects
// along a vertical axis. You can combine this filter with the
// [ConvolutionRGB9HorizontalFilter] to apply separable 9 x 9 convolutions.
// This filter differs from the [Convolution9VerticalFilter] filter, which
// processes all of the color components including the alpha component.
//
// # Discussion
//
// The convolution-RGB-9-vertical filter uses the following properties:
//
// `inputImage`: A [CIImage] containing the image to process. `weights`: A
// [CIVector] representing the convolution kernel. `bias`: A `float`
// representing the value that’s added to each output pixel.
//
// The following code creates a filter that blurs the image in the vertical
// direction:
//
// [media-4407305]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/convolutionRGB9Vertical()
func (_CIFilterClass CIFilterClass) ConvolutionRGB9VerticalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("convolutionRGB9VerticalFilter"))
	return CIFilterFromID(rv)
}

// Simulates the effect of a copy machine scanner light to transiton between
// two images.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the copy machine transition filter to an image. The
// effect transitions from one image to another by simulating the scanning
// light effect of a copy machine.
//
// The copy machine transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `time`: A `float` representing
// the parametric time of the transition from start (at time 0) to end (at
// time 1) as an [NSNumber]. `angle`: A `float` representing the angle of the
// copier light, in radians as an [NSNumber]. `width`: A `float` representing
// the width of the effect as a [NSNumber]. `extent`: A [CGRect] representing
// the area of the copy machine effect. `color`: A [CIColor] representing the
// color of the light. `opacity`: A `float` representing the transparency of
// the copier light as an [NSNumber].
//
// The following code creates a filter that produces a light bar that glides
// across the input image revealing the target image:
//
// [media-3616432]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/copyMachineTransition()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CopyMachineTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("copyMachineTransitionFilter"))
	return CIFilterFromID(rv)
}

// Filters an image with a Core ML model.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the Core ML model filter to an image. The effect
// filters the image using a trained Core ML model to produce the result.
// Specifying the head index allows you to produce a result from various
// components of a multiheaded coreML model.
//
// The Core ML model filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `headIndex`: A `float`
// representing which output of a multihead Core ML model should be used for
// applying the effect to an image. `softmaxNormalization`: A [Boolean] value
// representing the softmax normalization to be applied to the output image
// created by the model. `inputModel`: The Core ML model to be used for
// applying effect on the image.
//
// The following code creates a filter that results in the flowers appearing
// to be glass panes:
//
// [media-3600002]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/coreMLModel()
func (_CIFilterClass CIFilterClass) CoreMLModelFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("coreMLModelFilter"))
	return CIFilterFromID(rv)
}

// Creates an image made with a series of colorful polygons.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the crystallize filter to an image. The effect creates
// polygon-shaped color blocks by aggregating pixel-color values.
//
// The crystallize filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint].
//
// The following code creates a filter that results in an image made of small
// polygons:
//
// [media-3599996]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/crystallize()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) CrystallizeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("crystallizeFilter"))
	return CIFilterFromID(rv)
}

// Blends colors from two images while darkening lighter pixels.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the darken-blend mode filter to an image. The effect
// targets lighter colors in the background image and darkens them, while
// darker colors remain unchanged. The result is the sum of the adjusted color
// values from both images.
//
// The darken-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that darkens the image:
//
// [media-3546416]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/darkenBlendMode()
func (_CIFilterClass CIFilterClass) DarkenBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("darkenBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Simulates a depth of field effect.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the depth of field filter to an image. The effect
// simulates changing the focus of the camera before taking a photograph.
//
// The depth of field filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber]. `point0`: A set of
// coordinates marking the first point to be focused on as a [CGPoint].
// `point1`: A set of coordinates marking the second point to be focused on as
// a [CGPoint]. `unsharpMaskRadius`: A `float` representing the radius of the
// unsharpened mask effect applied to the in-focus area of effect as an
// [NSNumber]. `unsharpMaskIntensity`: A `float` representing the intensity of
// the unsharp mask effect as an [NSNumber].
//
// The following code creates a filter that results in the center cilantro
// being in focus while gradually blurring to the top and bottom of the image:
//
// [media-3599997]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/depthOfField()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DepthOfFieldFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("depthOfFieldFilter"))
	return CIFilterFromID(rv)
}

// Converts from an image containing depth data to an image containing
// disparity data.
//
// # Return Value
//
// An image containing the disparity data.
//
// # Discussion
//
// This method applies the depth-to-disparity filter. The filter takes depth
// data as an input and produces disparity data in the output image. You can
// use the output of this filter to create a stereo image.
//
// The depth-to-disparity filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that generates a depth map image:
//
// [media-3598060]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/depthToDisparity()
func (_CIFilterClass CIFilterClass) DepthToDisparityFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("depthToDisparityFilter"))
	return CIFilterFromID(rv)
}

// Subtracts color values to blend colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the difference-blend mode filter to an image. The
// effect calculates the brightness value for both images and subtracts the
// smaller value, resulting in a slightly darker image. This effect doesn’t
// modify images that are black.
//
// The difference-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that blends colors in the background
// image and input image:
//
// [media-3546397]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/differenceBlendMode()
func (_CIFilterClass CIFilterClass) DifferenceBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("differenceBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Applies a circle-shaped blur to an area of an image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the disc blur filter to an image. The effect targets
// the pixels within a circle defined by a `radius` and calculates the median
// color value to create to the output image.
//
// The disc blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that adds a strong blur to the input
// image:
//
// [media-3544964]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/discBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DiscBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("discBlurFilter"))
	return CIFilterFromID(rv)
}

// Transitions between two images using a mask image.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the disintegrate with mask transition filter to an
// image. The effect transitions from one image to another using the shapes
// defined by the mask image.
//
// The disintegrate with mask transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `maskImage`: An image with the
// type [CIImage]. `time`: A `float` representing the parametric time of the
// transition from start (at time 0) to end (at time 1) as an [NSNumber].
// `shadowRadius`: A `float` representing the size of the shadow as a
// [NSNumber]. `shadowDensity`: A `float` representing the strength of the
// shadow as a [NSNumber]. `shadowOffset`: A [CGPoint] representing the size
// of the shadow from the mask image.
//
// The following code creates a filter that produces a transition between the
// input and target images starting in the area’s outline in the mask image:
//
// [media-3616430]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/disintegrateWithMaskTransition()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DisintegrateWithMaskTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("disintegrateWithMaskTransitionFilter"))
	return CIFilterFromID(rv)
}

// Creates depth data from an image containing disparity data.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// The method generates the disparity-to-depth filter. The filter converts a
// depth data image to disparity data. You can combine with other filters to
// create more sophisticated images.
//
// The disparity-to-depth filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that generates a disparity depth map
// image:
//
// [media-3598059]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/disparityToDepth()
func (_CIFilterClass CIFilterClass) DisparityToDepthFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("disparityToDepthFilter"))
	return CIFilterFromID(rv)
}

// Applies the grayscale values of the second image to the first image.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the displacement distortion filter to an image. This
// effect distorts an image by applying the grayscale color values of the
// texture image.
//
// The displacement distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `displacementImage`: An
// image with the type [CIImage]. `scale`: A `float` representing the scaling
// the filter uses to apply the texture to the input image as an [NSNumber].
//
// The following code creates a filter that applies the grayscale values of
// the displacement image to the input image:
//
// [media-4407313]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/displacementDistortion()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DisplacementDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("displacementDistortionFilter"))
	return CIFilterFromID(rv)
}

// Transitions between two images with a fade effect.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the disintegrate transition filter to an image. The
// effect transitions from one image to another by using a fade effect.
//
// The dissolve transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `time`: A `float` representing
// the parametric time of the transition from start (at time 0) to end (at
// time 1) as an [NSNumber].
//
// The following code creates a filter that produces a fade transition from
// the input image and the target image:
//
// [media-3616426]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dissolveTransition()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DissolveTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("dissolveTransitionFilter"))
	return CIFilterFromID(rv)
}

// Applies randomized noise to produce a processed look.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// The effect applies a dithering effect to the input image. The effect
// applies randomized noise to the input image to produce a processed look.
//
// The dither filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `intensity`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that adds desaturation to the input
// image:
//
// [media-3545020]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dither()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DitherFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("ditherFilter"))
	return CIFilterFromID(rv)
}

// Divides color values to blend colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the divide-blend mode filter to an image. The effect
// divides the background image color values by the color values of the input
// image, resulting in the output image.
//
// The divide-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that brightens and inverts colors in
// the background image:
//
// [media-3546405]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/divideBlendMode()
func (_CIFilterClass CIFilterClass) DivideBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("divideBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s shadows and contrast.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the document enhancer filter to an image. The effect
// removes unwanted shadows while whitening the background and enhancing
// contrast. The filter is commonly used to enhance scanned documents.
//
// The document enhancer filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `amount`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545030]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/documentEnhancer()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DocumentEnhancerFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("documentEnhancerFilter"))
	return CIFilterFromID(rv)
}

// Creates a monochrome image with a series of dots to add detail.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a dot screen filter to an image. The effect generates a
// monochrome image containing a series of dots creating detail. The halftone
// effect is a set of lines, dots, or circles that contain detail. When
// viewing the image from a distance, the markings blend together, creating
// the illusion of continuous lines and shapes. Print media commonly uses this
// effect.
//
// The dot screen filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `width`: A
// `float` representing the distance between dots in the pattern as an
// [NSNumber]. `sharpness`: A `float` representing the sharpness of the
// pattern as an [NSNumber].
//
// The following code creates a filter that produces an image containing
// monochrome dots of detail on a black background:
//
// [media-3595912]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/dotScreen()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DotScreenFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("dotScreenFilter"))
	return CIFilterFromID(rv)
}

// Stylizes an image with the Droste effect.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the Droste filter to an image. This effect creates a
// Droste effect that distorts the image by repeating smaller versions of the
// same image within itself.
//
// The Droste filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `rotation`: A `float`
// representing the angle of the rotation, in radians, as an [NSNumber].
// `zoom`: A `float` representing the zoom of the effect as an [NSNumber].
// `periodicity`: A float representing the amount of intervals as an
// [NSNumber]. `inputInsetPoint1`: A [CGPoint] representing the x and y
// position that defines the first inset point. `inputInsetPoint0`: A
// [CGPoint] representing the x and y position that defines the second inset
// point. `inputStrands`: A float representing the amount of strands as an
// [NSNumber].
//
// The following code creates a filter that results in the image becoming a
// repeated, scaled pattern:
//
// [media-4407275]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/droste()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) DrosteFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("drosteFilter"))
	return CIFilterFromID(rv)
}

// Creates a high-quality upscaled image.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the edge preserve upsample filter to an image. The
// effect upsamples a small input image to be the size of the scale image
// using the luminance of the input image to preserve detail.
//
// The edge preserve upsample filter uses the following properties:
//
// `inputImage`: An image representing the image to upscale with the type
// [CIImage]. `scaleImage`: An image representing the reference for scaling
// the input image with the type [CIImage]. `spatialSigma`: A float
// representing the influence of the input image’s spatial information on
// the upsampling operation as an [NSNumber]. `lumaSimga`: A float
// representing influence of the input image’s luma information on the
// upsampling operation as an [NSNumber].
//
// The following code creates a filter that upscales the smaller image to the
// size of the scale image:
//
// [media-3582220]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edgePreserveUpsample()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) EdgePreserveUpsampleFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("edgePreserveUpsampleFilter"))
	return CIFilterFromID(rv)
}

// Produces a black-and-white image that looks similar to a woodblock print.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the edge work filter to an image. The effect creates a
// stylized black-and-white rendition of the image that looks similar to a
// woodblock print.
//
// The edge work filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber].
//
// The following code creates a filter that results in a monochrome image with
// the edges of objects highlighted:
//
// [media-3600010]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edgeWork()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) EdgeWorkFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("edgeWorkFilter"))
	return CIFilterFromID(rv)
}

// Hilghlights edges of objects found within an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the edges filter to an image. The effect uses the
// `intensity` to compute and highlight edges of items within the image.
//
// The edges filter uses the following property:
//
// `inputImage`: An image with the type [CIImage]. `intensity`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in a darker image with the
// edges of objects highlighted with the colors of the input image:
//
// [media-3600007]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/edges()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) EdgesFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("edgesFilter"))
	return CIFilterFromID(rv)
}

// Creates an eight-way reflected pattern.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the eight-fold reflected tile filter to an image. The
// effect creates an eight-way symmetry from the input image and tiles it to
// create the output image.
//
// The eight-fold reflected tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in small symmetrical
// repeated tiles:
//
// [media-4333630]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/eightfoldReflectedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) EightfoldReflectedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("eightfoldReflectedTileFilter"))
	return CIFilterFromID(rv)
}

// Subtracts color values to blend colors with less contrast.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the exclusion-blend mode filter to an image. The effect
// calculates the brightness value for both images and subtracts the smaller
// value, resulting in a darker image.
//
// The exclusion-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that mixes the colors and results in an
// output image that’s less saturated:
//
// [media-3546411]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/exclusionBlendMode()
func (_CIFilterClass CIFilterClass) ExclusionBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("exclusionBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s exposure.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the exposure-adjust filter to an image. The effect uses
// multiplication of color values to simulate the change of exposure within
// the photo.
//
// The exposure-adjust filter uses the following properties:
//
// `ev`: A `float` representing the amount to adjust the exposure as an
// [NSNumber]. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545001]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/exposureAdjust()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ExposureAdjustFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("exposureAdjustFilter"))
	return CIFilterFromID(rv)
}

// Replaces an image’s colors with specified colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the false color filter to an image. The effect maps the
// luminance to a color ramp from `color0` to `color1`. People use this effect
// to process astronomical and other scientific data, such as ultraviolet and
// X-ray images.
//
// The false color filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `color0`: A [CIColor]
// representing the first color to use for the color ramp. `color1`: A
// [CIColor] representing the second color to use for the color ramp.
//
// The following code creates a filter that replaces the colors of the input
// image resulting in blue and yellow colors:
//
// [media-3545024]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/falseColor()
func (_CIFilterClass CIFilterClass) FalseColorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("falseColorFilter"))
	return CIFilterFromID(rv)
}

// Creates a [CIFilter] object for a specific kind of filter and initializes
// the input values with a `nil`-terminated list of arguments.
//
// name: The name of the filter. You must make sure the name is spelled correctly,
// otherwise your app will run but not produce any output images. For that
// reason, you should check for the existence of the filter after calling this
// method.
//
// key0: A list of key-value pairs to set as input values to the filter. Each key is
// a constant that specifies the name of the input value to set, and must be
// followed by a value. You signal the end of the list by passing a `nil`
// value.
//
// # Return Value
//
// A [CIFilter] object whose input values are initialized.
//
// # Discussion
//
// As with all Objective-C methods that accept `nil`-terminated argument
// lists, to prevent unintended behavior you must take take care not to pass a
// `nil` value before the intended end of the argument list. You can avoid
// such issues by using the [init(name:withInputParameters:)] method to create
// a filter, expressing the parameter list as a dictionary literal.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/filterWithName:keysAndValues:
//
// [init(name:withInputParameters:)]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/init(name:withInputParameters:)
func (_CIFilterClass CIFilterClass) FilterWithNameKeysAndValues(name string, key0 objectivec.IObject) CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("filterWithName:keysAndValues:"), objc.String(name), key0)
	return CIFilterFromID(rv)
}

// Creates a flash of light to transition between two images.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the flash transition filter to an image. The effect
// transitions from the input image to the target image by creating a flash
// that fills the image and fades to the target image.
//
// The flash transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `center`: A set of coordinates
// marking the center of the image as a [CGPoint]. `extent`: A [CGRect]
// representing the size of the rounded rectangle. `color`: A [CIColor]
// representing the color of the flash effect. `time`: A `float` representing
// the parametric time of the transition from start (at time 0) to end (at
// time 1) as an [NSNumber]. `maxStiriationRadius`: A `float` representing the
// radius of the light rays emanating from the flash as a [NSNumber].
// `striationStrength`: A `float` representing the strength of the light rays
// emanating from the flash as a [NSNumber]. `striationContrast`: A `float`
// representing the contrast that’s added to each output pixel as a
// [NSNumber]. `fadeThreshold`: A `float` representing the amount of fade
// between the flash and the target image as a [NSNumber].
//
// The following code creates a filter that transitions from the input image
// with a large flash of light and fades to the target image.
//
// [media-3616427]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/flashTransition()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) FlashTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("flashTransitionFilter"))
	return CIFilterFromID(rv)
}

// Creates a four-way reflected pattern.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the four-fold reflected tile filter to an image. The
// effect produces a four-way reflected tile image.
//
// The four-fold reflected tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber]. `acute angle`: A `float` representing the primary angle for
// the repeating parallelogram tile as an [NSNumber].
//
// The following code creates a filter that results in a four-fold pattern:
//
// [media-3599880]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldReflectedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) FourfoldReflectedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("fourfoldReflectedTileFilter"))
	return CIFilterFromID(rv)
}

// Creates a tiled image by rotating a tile in increments of 90 degrees.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the four-fold rotated tile filter to an image. The
// effect produces a tiled image by rotating a tile from the iput image in
// increments of 90 degrees.
//
// The four-fold rotated tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in the flowers in the
// input image becoming rotated by 90 degrees and tiled:
//
// [media-3599888]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldRotatedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) FourfoldRotatedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("fourfoldRotatedTileFilter"))
	return CIFilterFromID(rv)
}

// Creates a tiled image by applying four translation operations.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the four-fold translated tile filter to an image. The
// effect produces a four-way tiled image by applying four translation
// operations. Translation operations map the position of each element in the
// photo to a new position in the output image.
//
// The four-fold translated tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. This controls
// the source of the tile contents. `angle`: A `float` representing the
// direction of the tiled patten, in radians as an [NSNumber]. `width`: A
// `float` representing the set width of each tile as an [NSNumber].
// `acuteAngle`: A `float` representing the primary angle for the repeating
// translated tile as an [NSNumber].
//
// The following code creates a filter that performs a four-fold translated
// tile operation on the image:
//
// [media-4333629]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/fourfoldTranslatedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) FourfoldTranslatedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("fourfoldTranslatedTileFilter"))
	return CIFilterFromID(rv)
}

// Highlights textures in an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the Gabor gradients filter to an image. The effect
// targets the texture of objects within the frame, and is frequently used to
// find detail in photographs of fingerprints.
//
// The gabor gradients filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in a darker image with
// shades of green and red outlining the texture of objects:
//
// [media-3600000]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaborGradients()
func (_CIFilterClass CIFilterClass) GaborGradientsFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("gaborGradientsFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s transition between black and white.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the gamma-adjust filter to an image. The effect adjusts
// the image’s mid-tone brightness. This filter is typically used to
// compensate for distortion caused by displays.
//
// The gamma-adjust filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds darker colors to the input
// image:
//
// [media-3545006]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gammaAdjust()
func (_CIFilterClass CIFilterClass) GammaAdjustFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("gammaAdjustFilter"))
	return CIFilterFromID(rv)
}

// Blurs an image with a Gaussian distribution pattern.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies a Gaussian blur filter to an image. The effect targets
// the pixels within a circle defined by a `radius` and uses Gaussian
// ditribution to blur the image from the center out.
//
// The Gaussian blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that adds a heavy blur to the input
// image:
//
// [media-3544963]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaussianBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) GaussianBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("gaussianBlurFilter"))
	return CIFilterFromID(rv)
}

// Generates a gradient that varies from one color to another using a Gaussian
// distribution.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a Gaussian gradient image. The effect uses the
// Gaussian kernel to calculate the even dispersal of the first color in the
// center to the second color in the image’s periphery.
//
// The Gaussian gradient filter uses the following properties:
//
// `center`: A [CGPoint] representing the center of the effect as x and y
// coordinates. `color0`: A [CIColor] representing the first color to use in
// the gradient. `color1`: A [CIColor] representing the second color to use in
// the gradient. `radius`: A `float` representing the radius of the Gaussian
// distribution as an [NSNumber].
//
// The following code creates a filter that generates a gradient image:
//
// [media-3558795]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gaussianGradient()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) GaussianGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("gaussianGradientFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by applying a glass-like texture.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the glass distortion filter to an image. This effect
// distorts an image by applying a glass texture from the raised portions of
// the texture map image.
//
// The glass distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `texture`: An image with
// the type [CIImage]. `scale`: The amount of texturing to apply. Larger
// values increase the effect. Defaults to 200. `center`: A set of coordinates
// marking the center of the image as a [CGPoint].
//
// The following code creates a filter that results in a glass-like distortion
// applied to the image:
//
// [media-4407314]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glassDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) GlassDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("glassDistortionFilter"))
	return CIFilterFromID(rv)
}

// Creates a lozenge-shaped lens and distorts the image.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the glass lozenge filter to an image. This effect
// distorts an image by creating a lozenge shape placed over the input image.
//
// The absolute threshold filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the radius of the lozenge distortion as an [NSNumber].
// `refraction`: A `float` representing the refraction of the glass as an
// [NSNumber]. `inputPoint1`: A [CGPoint] representing the x and y positions
// that define the center of the circle at the first end of the lozenge.
// `inputPoint2`: A [CGPoint] representing the x and y positions that define
// the center of the circle at the second end of the lozenge.
//
// The following code creates a filter that results in a large glass lozenge
// distorting the image:
//
// [media-4407280]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glassLozenge()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) GlassLozengeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("glassLozengeFilter"))
	return CIFilterFromID(rv)
}

// Tiles an image by rotating and reflecting a tile from the image.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the glide reflected tile filter to an image. The effect
// produces a tiled image by rotating and reflecting a tile from the input
// image.
//
// The glide reflected tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in flipping the image and
// then tiling the result:
//
// [media-3599883]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/glideReflectedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) GlideReflectedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("glideReflectedTileFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s color by applying a gloom filter.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the gloom filter to an image. The effect reduces the
// highlights of the image resulting in the image looking dull.
//
// The gloom filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. radius: A `float`
// representing the area of effect as an [NSNumber]. intensity: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in a darker image with a
// slight blur:
//
// [media-3599998]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/gloom()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) GloomFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("gloomFilter"))
	return CIFilterFromID(rv)
}

// Blends colors of two images by screening and multiplying.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the hard-light blend mode filter to an image. The
// effect calculates the brightness of the colors in the background image.
// Colors that are lighter than 50 percent gray become lighter. If the
// brightness of the colors in the input image are darker then 50 percent
// gray, the effect darkens the colors. The filter then uses the calculated
// results to create the output image.
//
// The hard-light blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that mixes the colors and results in an
// output image that’s darker and more saturated:
//
// [media-3546419]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hardLightBlendMode()
func (_CIFilterClass CIFilterClass) HardLightBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hardLightBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Creates a monochrome image with a series of lines to add detail.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a hatched screen filter to an image. The effect
// generates a monochrome image containing a series of lines in hatched
// pattern to create detail. The halftone effect is a set of lines, dots, or
// circles that contain detail. When viewing the image from a distance, the
// markings blend together, creating the illusion of continuous lines and
// shapes. The effect is often used in print media for more efficient
// printing.
//
// The hatched screen filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the angle of the pattern as an [NSNumber]. `width`: A
// `float` representing the distance between lines in the pattern as an
// [NSNumber]. `sharpness`: A `float` representing the sharpness of the
// pattern as an [NSNumber].
//
// The following code creates a filter that produces a monochrome image
// containing lines of detail on a black background:
//
// [media-3595917]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hatchedScreen()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HatchedScreenFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hatchedScreenFilter"))
	return CIFilterFromID(rv)
}

// Creates a realistic shaded height-field image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the height-field from the mask filter to an image. The
// effect targets the white in the input image and creates realistic shading.
//
// The height field from mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber].
//
// The following code creates a filter that results in the text having a
// shading effect:
//
// [media-3600011]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/heightFieldFromMask()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HeightFieldFromMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("heightFieldFromMaskFilter"))
	return CIFilterFromID(rv)
}

// Creates an image made of a series of colorful hexagons.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the hexagonal pixelate filter to an image. The effect
// creates an image containing colored hexagons.
//
// The hexagonal pixelate filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `scale`: A
// `float` representing the scale of the hexagons as an [NSNumber].
//
// The following code creates a filter that results in an image made up of
// hexagons:
//
// [media-3600003]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hexagonalPixellate()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HexagonalPixellateFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hexagonalPixellateFilter"))
	return CIFilterFromID(rv)
}

// Adjusts the highlights of colors to reduce shadows.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the highlight-shadow adjust filter to an image. The
// effect adjusts shadows, while preserving spatial detail in the image.
//
// The highlight-shadow adjust filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `shadowAmount`: A `float`
// representing the amount of generated shadow as an [NSNumber]. `radius`: A
// `float` representing the radius of the shadow as an [NSNumber].
// `highlightAmount`: A `float` representing the strength of the shadow as an
// [NSNumber].
//
// The following code creates a filter that results in a brighter image with
// reduced shadows:
//
// [media-3600012]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/highlightShadowAdjust()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HighlightShadowAdjustFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("highlightShadowAdjustFilter"))
	return CIFilterFromID(rv)
}

// Generates a histogram map from the image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method applies the histogram display filter to the result of the
// output from the [AreaHistogramFilter] filter. This effect shows a graphical
// representation of the tonal distribution of colors in the image.
//
// The histogram display filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. Typically this is the
// output from the area histogram filter. `height`: A `float` representing the
// height of the generated histogram image as an [NSNumber]. `lowLimit`: A
// `float` representing the fraction of the left portion of the histogram
// image to make darker as an [NSNumber]. `hightLimit`: A `float` representing
// the fraction of the right portion of the histogram to make lighter as an
// [NSNumber].
//
// The following code creates a filter that results in a histogram diagram
// generated from the input image:
//
// [media-4332168]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/histogramDisplay()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HistogramDisplayFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("histogramDisplayFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image with a circular area that pushes the image outward.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the hole distortion filter to an image. This effect
// distorts the image by generating a circular area that displaces pixels in
// the image by pushing them outward from the hole defined by the radius.
//
// The absolute threshold filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `radius`: A
// `float` representing the amount of pixels the filter uses to create the
// distortion as an [NSNumber].
//
// The following code creates a filter that results in an image becoming
// distorted from the center outward:
//
// [media-4407309]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/holeDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HoleDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("holeDistortionFilter"))
	return CIFilterFromID(rv)
}

// Modifies an image’s hue.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the hue-adjust filter to an image. The effect changes
// the hue or color of the pixels by using the angle to modify the image’s
// color data.
//
// The hue-adjust filter uses the following properties:
//
// `angle`: A `float` representing the angle in radians to adjust the current
// hue of the image as an [NSNumber]. `inputImage`: An image with the type
// [CIImage].
//
// The following code creates a filter that shifts the hue of the image by 5
// radians.
//
// [media-3544999]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueAdjust()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HueAdjustFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hueAdjustFilter"))
	return CIFilterFromID(rv)
}

// Blends colors of two images by computing the sum of image color values.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the hue-blend mode filter to an image. The effect uses
// the values of the saturation and luminance from the background image with
// the hue of the input image to produce the output.
//
// The hue-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that applies the hue-blend mode filter.
//
// [media-3546412]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueBlendMode()
func (_CIFilterClass CIFilterClass) HueBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hueBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Generates a gradient representing a specified color space.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a hue-saturation-value gradient image. The filter
// creates a color wheel that shows the hues and saturations for a specified
// [CGColorSpace].
//
// The hue-saturation-value gradient uses the following properties:
//
// `colorSpace`: A [CGColorSpace] representing the color space for the
// generated color wheel. `dither`: A `boolean` value specifying whether the
// distort the generated output. `radius`: A `float` representing the distance
// from the center of the effect as an [NSNumber]. `softness`: A `float`
// representing the softness of the generated color wheel as an [NSNumber].
// `value`: A `float` representing the lightness of the hue-saturation
// gradient as an [NSNumber].
//
// The following code creates a filter that generates a color-space image:
//
// [media-3558798]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/hueSaturationValueGradient()
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) HueSaturationValueGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("hueSaturationValueGradientFilter"))
	return CIFilterFromID(rv)
}

// Applies the k-means algorithm to find the most common colors in an image.
//
// # Return Value
//
// A one-dimensional [CIImage] containing the colors.
//
// # Discussion
//
// This filter uses the k-means clustering algorithm to find the most common
// colors in an input image. The result is a [CIImage] with `count` x 1
// dimensions. Each [RGBA] pixel in the result image represents the center of
// a k-means cluster. The [RGB] components contain the color and the alpha
// component represents the weight of the color. You typically use the
// [KMeansFilter] filter in conjunction with the [PalettizeFilter] filter to
// produce an image with a reduced number of colors.
//
// `inputImage`: A [CIImage] to process. `extent`: A [CGRect] specifying the
// area of the image to analyze. `means`: An optional [CIImage] containing a
// set of colors to use as seeds for the k-means clustering. `count`: The
// number of k-means color clusters that should be created. Maximum is `128`,
// and default is `8`. `passes`: The number of k-means passes that should run.
// Maximum is `20`, and default is `5`. `perceptual`: Whether the k-means
// color palette should use a perceptual color space.
//
// The following code example uses the [KMeansFilter] filter followed by the
// [PalettizeFilter] filter to reduce the colors in the image to four:
//
// [media-4332587]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/kMeans()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) KMeansFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("KMeansFilter"))
	return CIFilterFromID(rv)
}

// Creates a 12-way kaleidoscopic image from an image.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the kaleidoscope tile filter to an image. The effect
// produces a complex 12-way symmetrical reflected pattern from the input
// image.
//
// The kaleidoscope tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `count`: A `float` representing the number of reflections in
// the pattern as an [NSNumber].
//
// The following code creates a filter that results in the creation of a
// kaleidoscope effect from the input image:
//
// [media-3599884]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/kaleidoscope()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) KaleidoscopeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("kaleidoscopeFilter"))
	return CIFilterFromID(rv)
}

// Adjusts the image vertically and horizontally to remove distortion.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the keystone correction combined filter to an image.
// The effect applies a combined set of horizontal and vertical guides to
// adjust the shape of the input image. This effect is commonly used when
// cropping an image to correct distortion. In the figure below, both vertical
// and horizontal adjustments are made, resulting in a trapezoid-shaped image.
//
// The keystone correction filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
// `focalLength`: A `float` representing the simulated focal length as an
// [NSNumber].
//
// The following code creates a filter that distorts the image:
//
// [media-3582223]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionCombined()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) KeystoneCorrectionCombinedFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("keystoneCorrectionCombinedFilter"))
	return CIFilterFromID(rv)
}

// Horizontally adjusts an image to remove distortion.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the keystone correction horizontal filter to an image.
// The effect applies a set of horizontal guides and simulated focal length to
// adjust the shape of the input image. This effect is commonly used when
// cropping an image to correct distortion. In the figure below, both vertical
// and horizontal adjustments are made, resulting in a trapezoid-shaped image.
//
// The keystone correction horizontal filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
// `focalLength`: A `float` representing the simulated focal length as an
// [NSNumber].
//
// The following code creates a filter that distorts the image:
//
// [media-3582230]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionHorizontal()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) KeystoneCorrectionHorizontalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("keystoneCorrectionHorizontalFilter"))
	return CIFilterFromID(rv)
}

// Vertically adjusts an image to remove distortion.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the keystone correction vertical. The effect performs
// vertical adjustment of the image to shape the image to be rectangular. This
// effect is commonly used with multimedia projectors to correct the
// distortion caused by the projector being lower or higher than the projected
// screen.
//
// The keystone vertical filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
// `focalLength`: A `float` representing the simulated focal length as an
// [NSNumber].
//
// The following code creates a filter that distorts the image:
//
// [media-3582222]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/keystoneCorrectionVertical()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) KeystoneCorrectionVerticalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("keystoneCorrectionVerticalFilter"))
	return CIFilterFromID(rv)
}

// Compares an image’s color values.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the Lab ΔE filter to an image. The effect creates an
// image based on the visual color differences between the two input images.
// The resulting image contains ΔE 1994 values between 0.0 and 100.0.
//
// The Lab ΔE filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `image2`: An image with the
// type [CIImage] the system uses for comparison.
//
// The following code creates a filter that removes the background from the
// input image:
//
// [media-3546476]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/labDeltaE()
func (_CIFilterClass CIFilterClass) LabDeltaE() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("LabDeltaE"))
	return CIFilterFromID(rv)
}

// Creates a high-quality, scaled version of a source image.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the Lanczos scale transform filter to an image. The
// effect creates the output image by scaling the input image based on the
// scale and aspect ratio properties provided.
//
// The Lanczos scale filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `scale`: A `float`
// representing the scaling factor used on the image as an [NSNumber]. Values
// less than `1.0` scale down the images. Values greater than `1.0` scale up
// the image. `aspectRatio`: A `float` representing the additional horizontal
// scaling factor used on the image as an [NSNumber].
//
// The following code creates a filter that results in a smaller scaled image
// with high quality:
//
// [media-3582221]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lanczosScaleTransform()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) LanczosScaleTransformFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lanczosScaleTransformFilter"))
	return CIFilterFromID(rv)
}

// Generates a lenticular halo image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a lenticular halo image. You commonly combine this
// effect with an image to simulate a halo generated by the spread of light on
// a lens.
//
// The lenticular halo generator filter uses the following properties:
//
// `center`: A `vector` representing the center of the lens flare as a
// [CIVector]. `color`: A [CIColor] controlling the proportion of red, green,
// and blue halos. `haloWidth`: A `float` representing the halo width as an
// [NSNumber]. `haloRadius`: A `float` representing the halo radius as an
// [NSNumber]. `haloOverlap`: A `float` representing the overlap of red,
// green, and blue halos as an [NSNumber]. A value of 1 results in a full
// overlap. `striationStrength`: A `float` representing the brightness of the
// rainbow-colored halo area as an [NSNumber]. `striationContrast`: A `float`
// representing the contrast of the rainbow-colored halo area as an
// [NSNumber]. `time`: A `float` representing the addition of brightness to
// the halo as an [NSNumber].
//
// The following code creates a filter that generates a lenticular halo image:
//
// [media-3592392]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lenticularHaloGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) LenticularHaloGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lenticularHaloGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by generating a light tunnel.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the light tunnel filter to an image. This effect
// distorts the input image by warping the image to cylinder shape.
//
// The light tunnel filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the light tunnel as a [CGPoint].
// `radius`: A `float` representing the amount of pixels the filter uses to
// create the light tunnel as an [NSNumber]. rotation: A `float` representing
// the rotation angle of the light tunnel as an [NSNumber].
//
// The following code creates a filter that generates a swirling pattern from
// the input image:
//
// [media-4407304]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lightTunnel()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) LightTunnelFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lightTunnelFilter"))
	return CIFilterFromID(rv)
}

// Blends colors from two images by brightening colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the lighten-blend mode filter to an image. The effect
// replaces any samples in the background image that are darker than input
// image.
//
// The lighten-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image that blends
// the input and background image colors:
//
// [media-3546398]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lightenBlendMode()
func (_CIFilterClass CIFilterClass) LightenBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lightenBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Creates an image that resembles a sketch of the outlines of objects.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the line overlay filter to an image. The effect creats
// a sketch that outlines the edges of the image in black, leaving the
// non-outlined portion of the image transparent.
//
// The line overlay filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `nrNoiseLevel`: A `float`
// representing the desired level of noise as an [NSNumber]. `nrSharpness`: A
// `float` representing the desired level of sharpness as an [NSNumber].
// `edgeIntensity`: A `float` representing the Sobel gradient information for
// edge tracing as an [NSNumber]. `threshold`: A `float` representing the
// threshold of edge visibilty as an [NSNumber]. `contrast`: A `float`
// representing the desired contrast as an [NSNumber].
//
// The following code creates a filter that results in a monochrome image with
// lines outlining the edges of objects:
//
// [media-3600001]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lineOverlay()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) LineOverlayFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lineOverlayFilter"))
	return CIFilterFromID(rv)
}

// Creates a monochrome image with a series of small lines to add detail.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a line screen filter to an image. The effect generates
// a monochrome image containing a series of lines creating detail. The
// halftone effect is a set of lines, dots, or circles that contain detail.
// When viewing the image from a distance, the markings blend together,
// creating the illusion of continuous lines and shapes. Print media commonly
// uses this effect.
//
// The line screen filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the angle of the pattern as an [NSNumber]. `width`: A
// `float` representing the distance between lines in the pattern as an
// [NSNumber]. `sharpness`: A `float` representing the sharpness of the
// pattern as an [NSNumber].
//
// The following code creates a filter that creates a monochrome image
// containing small lines of detail on a black background:
//
// [media-3595915]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/lineScreen()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) LineScreenFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("lineScreenFilter"))
	return CIFilterFromID(rv)
}

// Blends color from two images while increasing contrast.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the linear-burn blend mode filter to an image. The
// effect calculates the brightness value for the background image then
// darkens it to reflect the brightness of the input image. The effect then
// decreases the contrast in the output image.
//
// The linear-burn blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in an output image
// that’s much darker with very little visible detail:
//
// [media-3546415]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearBurnBlendMode()
func (_CIFilterClass CIFilterClass) LinearBurnBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("linearBurnBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Blends colors of two images with dodging.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the linear-dodge-blend mode filter to an image. The
// effect calculates the brightness value for the background image then
// brightens it to reflect the brightness of the input image. The effect then
// increases the contrast in the output image.
//
// The linear-dodge-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// brighter with both images’ colors:
//
// [media-3546417]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearDodgeBlendMode()
func (_CIFilterClass CIFilterClass) LinearDodgeBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("linearDodgeBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Generates a color gradient that varies along a linear axis between two
// defined endpoints.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a linear-gradient image. The effect creates a
// gradient that varies linearly between the two input properties of `point0`
// and `point1`.
//
// The linear-gradient filter uses the following properties:
//
// `point0`: A [CGPoint] representing the starting position of the gradient.
// `point1`: A [CGPoint] representing the ending position of the gradient.
// `color0`: A [CIColor] representing the first color to use in the gradient.
// `color1`: A [CIColor] representing the second color to use the gradient.
//
// The following code creates a filter that generates a gradient image:
//
// [media-3558797]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearGradient()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) LinearGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("linearGradientFilter"))
	return CIFilterFromID(rv)
}

// A combination of linear burn and linear dodge blend modes.
//
// # Return Value
//
// The blended image as a [CIImage].
//
// # Discussion
//
// The linear-light blend mode combines the linear-dodge and linear-burn blend
// modes (rescaled so that neutral colors become middle gray). If the input
// image’s values are lighter than middle gray, the filter uses dodge; for
// darker values, the filter uses burn.
//
// `inputImage`: A [CIImage] containing the input image `backgroundImage`: A
// [CIImage] containing the background image.
//
// The following code sample applies the linear-light blend mode filter to two
// images:
//
// [media-4407310]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearLightBlendMode()
func (_CIFilterClass CIFilterClass) LinearLightBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("linearLightBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s color intensity.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the linear-to-sRGB tone curve filter to an image. The
// effect converts an image in linear color space to sRGB.
//
// The linear-to-sRGB tone curve filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545000]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/linearToSRGBToneCurve()
func (_CIFilterClass CIFilterClass) LinearToSRGBToneCurveFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("linearToSRGBToneCurveFilter"))
	return CIFilterFromID(rv)
}

// Blends color from two images by calculating the color, hue, and saturation.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the luminosity-blend mode filter to an image. The
// effect creates the output image by using the hue and saturation values of
// the background image while using the luminance values of the input image.
//
// The luminosity-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image having
// accurate gray colors while other colors are added from the background
// image:
//
// [media-3546410]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/luminosityBlendMode()
func (_CIFilterClass CIFilterClass) LuminosityBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("luminosityBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Converts an image to a white image with an alpha component.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the mask-to-alpha filter to an image. The value of the
// alpha component is determined by the grayscale value of the input image.
// Black pixels become completely transparent, white pixels are completely
// solid.
//
// The mask-to-alpha filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that makes the input image’s
// background transparent:
//
// [media-3545058]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maskToAlpha()
func (_CIFilterClass CIFilterClass) MaskToAlphaFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("maskToAlphaFilter"))
	return CIFilterFromID(rv)
}

// Blurs a specified portion of an image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the masked variable blur to an image. The effect blurs
// the image in an area defined by the mask image. The mask image contains
// shades of grey that define the strength of the blur. Black colors in the
// mask cause no blurring, and white colors cause maximum blur.
//
// The masked variable blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `mask`: An image that masks an area on the input image with the type
// [CIImage]. `inputImage`: A [CIImage] representing the input image to apply
// the filter to.
//
// The following code creates a filter that adds a blur to the bottom of the
// input image:
//
// [media-4334872]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maskedVariableBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MaskedVariableBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("maskedVariableBlurFilter"))
	return CIFilterFromID(rv)
}

// Creates a maximum RGB grayscale image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the maximum component filter to an image. The effect
// applies a preconfigured set of effects that result in the input image
// becoming grayscale using the maximum RGB color components.
//
// The maximum component filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds brightness and makes the
// input image grayscale:
//
// [media-3545023]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumComponent()
func (_CIFilterClass CIFilterClass) MaximumComponentFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("maximumComponentFilter"))
	return CIFilterFromID(rv)
}

// Applies a maximum compositing filter to an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the maximum compositing filter to an image. The effect
// calculates the maximum value for each color component in the input and
// background images. The filter then uses the resulting color to create the
// output image.
//
// The maximum compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in an image with a mixture
// of images’ colors:
//
// [media-3546413]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/maximumCompositing()
func (_CIFilterClass CIFilterClass) MaximumCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("maximumCompositingFilter"))
	return CIFilterFromID(rv)
}

// Calculates the median of an image to refine detail.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the median filter to an image. The effect computes the
// median value of colors for a group of neighboring pixels and replaces each
// pixel with calculated data.
//
// The median filter uses the following properties:
//
// inputImage: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that refines the detail in the input
// image:
//
// [media-3544967]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/median()
func (_CIFilterClass CIFilterClass) MedianFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("medianFilter"))
	return CIFilterFromID(rv)
}

// Generates a pattern made from an array of line segments.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a mesh generator image. The effect uses an array of
// line segments to create the resulting image.
//
// The mesh generator filter uses the following properties:
//
// `inputMesh`: An `array` of line segments stored as an array of [CIVector],
// each containing a start point and end point. `color`: A [CIColor]
// representing the color used to make the mesh. `width`: A `float`
// representing the width of the line segments as an [NSNumber]
//
// The following code creates a filter that generates a green star made from
// mesh segments:
//
// [media-3590974]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/meshGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MeshGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("meshGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Creates a minimum RGB grayscale image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the minimum component filter to an image. The effect
// applies a preconfigured set of effects that result in the input image
// becoming grayscale using the minimum RGB color components.
//
// The minimum component filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds darkness and makes the input
// image grayscale:
//
// [media-3545011]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/minimumComponent()
func (_CIFilterClass CIFilterClass) MinimumComponentFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("minimumComponentFilter"))
	return CIFilterFromID(rv)
}

// Blends colors from two images by computing minimum values.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the minimum compositing filter to an image. The effect
// calculates the minimum value for each color component in the input and
// background images. The filter then uses the resulting color to create the
// output image.
//
// The minimum compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the output image
// becoming brighter with both images’ colors:
//
// [media-3546409]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/minimumCompositing()
func (_CIFilterClass CIFilterClass) MinimumCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("minimumCompositingFilter"))
	return CIFilterFromID(rv)
}

// Blends two images together.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the mix filter to an image. The effect uses the amount
// property to interpolate between the input image and the background image,
// resulting in both images visible in the output image.
//
// The mix filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// representing the background image with the type [CIImage]. `amount`: A
// `float` representing the strength of the effect as an [NSNumber].
//
// The following code creates a filter that combines the input and background
// images to create one image with both images visible:
//
// [media-3600009]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/mix()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MixFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("mixFilter"))
	return CIFilterFromID(rv)
}

// Transitions between two images by applying irregularly shaped holes.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the mod transition filter to an image. The effect
// transitions from the input image to the output image by revealing the
// target image through irregularly shaped holes.
//
// The mod transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `center`: A [CGPoint]
// representing the center of the image. `angle`: A `float` representing the
// angle of the effect as an [NSNumber]. `radius`: A `float` representing the
// size of the area of effect as an [NSNumber]. `compression`: A `float`
// representing the amount of stretching applied to the mod hole pattern as an
// [NSNumber]. `time`: A `float` representing the parametric time of the
// transition from start (at time 0) to end (at time 1) as an [NSNumber].
//
// The following code creates a filter that transitions from the input image
// to the target image by creating a series of irregular shaped holes.
//
// [media-3616428]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/modTransition()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ModTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("modTransitionFilter"))
	return CIFilterFromID(rv)
}

// Detects and highlights edges of objects.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the morphology gradient filter to an image. The effect
// uses the `radius` to compute and highlight edges of items within the image.
//
// The morphology gradient filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that adds darkness to the overall image
// while the edges in the input photo become brighter:
//
// [media-3544969]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyGradient()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MorphologyGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("morphologyGradientFilter"))
	return CIFilterFromID(rv)
}

// Blurs a circular area by enlarging contrasting pixels.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the morphology maximum filter to an image. The effect
// targets a circular section of the image, calculating the median color
// values to find colors that make up more than half the working area. Using
// this calculation, the effect enlarges the pixels with contrasting colors to
// take up more of the working area. The effect is then repeated throughout
// the image.
//
// The morphology maximum filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that adds an intense blur to the input
// image:
//
// [media-3544961]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyMaximum()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MorphologyMaximumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("morphologyMaximumFilter"))
	return CIFilterFromID(rv)
}

// Blurs a circular area by reducing contrasting pixels.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the morphology minimum filter to an image. The effect
// targets a circular section of the image, calculating the median color
// values to find colors that make up more than half the working area. Using
// this calculation, the effect reduces the pixels with contrasting colors to
// take up less of the working area. The effect is then repeated throughout
// the image.
//
// The morphology minimum filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `inputImage`: A [CIImage] representing the input image to apply the filter
// to.
//
// The following code creates a filter that adds a blur that adds darkness to
// the input image:
//
// [media-3544966]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyMinimum()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MorphologyMinimumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("morphologyMinimumFilter"))
	return CIFilterFromID(rv)
}

// Blurs a rectangular area by enlarging contrasting pixels.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the morphology rectangle maximum filter to an image.
// The effect targets a rectangular section of the image, calculating the
// median color values to find colors that make up more than half the working
// area. Using this calculation, the effect enlarges the pixels with
// contrasting colors to take up more of the working area. The effect is then
// repeated throughout the image.
//
// The morphology rectangle maximum filter uses the following properties:
//
// `width`: A `float` representing the width in pixels of the working area as
// an [NSNumber]. `height`: A `float` representing the height in pixels of the
// working area as an [NSNumber]. `inputImage`: A [CIImage] representing the
// input image to apply the filter to.
//
// The following code creates a filter that adds a blur to the input image
// while brighting the palm trees:
//
// [media-3544958]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyRectangleMaximum()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MorphologyRectangleMaximumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("morphologyRectangleMaximumFilter"))
	return CIFilterFromID(rv)
}

// Blurs a rectangular area by reducing contrasting pixels.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the morphology rectangle minimum filter to an image.
// The effect targets a rectangular section of the image, calculating the
// median color values to find colors that make up more than half the working
// area. Using this calculation, the effect reduces the pixels with
// contrasting colors to take up more of the less area. The effect is then
// repeated throughout the image.
//
// The morphology rectangle minimum filter uses the following properties:
//
// width: A `float` representing the width in pixels of the working area as an
// [NSNumber]. height: A `float` representing the height in pixels of the
// working area as an [NSNumber]. inputImage: A [CIImage] representing the
// input image to apply the filter to.
//
// The following code creates a filter that adds an intense blur to the palm
// trees input image:
//
// [media-3544957]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/morphologyRectangleMinimum()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MorphologyRectangleMinimumFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("morphologyRectangleMinimumFilter"))
	return CIFilterFromID(rv)
}

// Creates motion blur on an image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the motion blur filter to an image. The filter uses the
// angle of a single row of pixels to determine the direction of the motion
// effect.
//
// The motion blur filter uses the following properties:
//
// `radius`: A `float` representing the area of effect as an [NSNumber].
// `angle`: A `float` representing the angle of the motion, in radians, that
// determines which direction the blur smears as an [NSNumber]. `inputImage`:
// A [CIImage] representing the input image to apply the filter to.
//
// The following code creates a filter that adds a motion blur to the input
// image:
//
// [media-3544965]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/motionBlur()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) MotionBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("motionBlurFilter"))
	return CIFilterFromID(rv)
}

// Blends colors from two images by multiplying color components.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the multiply-blend mode filter to an image. The effect
// calculates the colors in the output image by multiplying the color values
// for the input and background images, resulting in a darker image.
//
// The multiply-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// darker with less saturation:
//
// [media-3546403]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/multiplyBlendMode()
func (_CIFilterClass CIFilterClass) MultiplyBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("multiplyBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Blurs the colors of two images by multiplying color components.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the multiply compositing filter to an image. The effect
// calculates the color value of the output image by multiplying the color
// values from the input image and the background image.
//
// The multiply compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// darker with more saturation:
//
// [media-3546400]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/multiplyCompositing()
func (_CIFilterClass CIFilterClass) MultiplyCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("multiplyCompositingFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by stretching it between two breakpoints.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the nine-part stretched filter to an image. This effect
// distorts an image by stretching an image to the breakpoint properties while
// distorting the image based on the grow amount.
//
// The nine-part stretched filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `growAmount`: A [CGPoint]
// representing the amount of stretching applied. `breakpoint0`: A [CGPoint]
// representing the lower-left corner of the image to retain before stretching
// begins. `breakpoint1`: A [CGPoint] representing the upper-right corner of
// the image to retain after stretching ends.
//
// The following code creates a filter that results in a significantly warped
// image:
//
// [media-4407340]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/ninePartStretched()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) NinePartStretchedFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("ninePartStretchedFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by tiling portions of it.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the nine-part tiled filter to an image. This effect
// distorts an image by tiling an image based on the breakpoints properties.
//
// The nine-part tiled filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `flipYTiles`: A [Boolean]
// value representing if the y-axis should be flipped. `growAmount`: A
// [CGPoint] representing the amount of stretching applied. `breakpoint1`: A
// [CGPoint] representing the upper-right corner of the image to retain after
// tiling ends. `breakpoint0`: A [CGPoint] representing the lower-left corner
// of image to retain before stretching begins.
//
// The following code creates a filter that results in distorted tiles of the
// image becoming flipped:
//
// [media-4407331]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/ninePartTiled()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) NinePartTiledFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("ninePartTiledFilter"))
	return CIFilterFromID(rv)
}

// Reduces noise by sharpening the edges of objects.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the noise reduction filter to an image. The effect
// calculates changes in luminance below the noise level and locally blurs the
// area. Values above the threshold are determined to be edges, and become
// sharpened.
//
// The morphology noise reduction filter uses the following properties:
//
// `noiseLevel`: A `float` representing the amount of noise reduction as an
// [NSNumber]. `sharpness`: A `float` representing the sharpness of the final
// image as an [NSNumber]. `inputImage`: A [CIImage] representing the input
// image to apply the filter to.
//
// The following code creates a filter that reduces noise in the input image:
//
// [media-3544968]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/noiseReduction()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) NoiseReductionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("noiseReductionFilter"))
	return CIFilterFromID(rv)
}

// Produces an effect that mimics a style of visual art that uses optical
// illusions.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This filter extracts a tile from the image, applies any specified scaling
// and rotation, and then assembles the image again to give an optical
// illusion effect.
//
// The optical illusion tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber]. `scale`: A `float` representing the scale of numbers of
// tiles in the output image as an [NSNumber].
//
// The following code creates a filter that results in a distorted image with
// less detail:
//
// [media-3599881]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/opTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) OpTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("opTileFilter"))
	return CIFilterFromID(rv)
}

// Blends colors by overlaying images.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the overlay-blend mode filter to an image. The effect
// creates the output image by overlapping the input image over the background
// image while preserving highlights and shadows of the background image.
//
// The overlay-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the background image
// becoming darker with the input image overlaid on top:
//
// [media-3546406]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/overlayBlendMode()
func (_CIFilterClass CIFilterClass) OverlayBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("overlayBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Generates a high-density linear barcode.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a PDF417 barcode as an image. PDF417 is a
// high-density stacked linear barcode format defined in the ISO 15438
// standard. Use this filter to generate alphanumeric or numeric-only
// barcodes. Commonly used on identification cards or inventory management
// because of the large amount of data the barcode can hold.
//
// The PDF417 barcode generator filter uses the following properties:
//
// `message`: An [NSData] object representing the data to be encoded as a
// barcode. `minWidth`: A `float` representing the minimum width of the
// barcode’s data area, in pixels as an [NSNumber]. `maxWidth`: A `float`
// representing the maximum width of the barcode’s data area, in pixels, as
// an [NSNumber]. `maxHeight`: A `float` representing the maximum height of
// the barcode’s data area, in pixels, as an [NSNumber]. `minHeight`: A
// `float` representing the minimum height of the barcode’s data area, in
// pixels, as an [NSNumber]. `dataColums`: A `float` representing the number
// of columns in the data area as an [NSNumber]. `rows`: A `float`
// representing the number of rows in the data area as an [NSNumber].
// `preferredAspectRatio`: A `float` representing the desired aspect ratio as
// an [NSNumber]. `compactionMode`: An option that determines which method the
// generator uses to compress data as an [NSNumber]. See the note below for
// the possible values. `compactStyle`: A [Boolean] value of `0` or `1` that
// determines the omission of redundant elements to make the generated barcode
// more compact as an [NSNumber]. `correctionLevel`: A `float` between 0 and 8
// that determines the amount of redundancy to include in the barcode’s data
// to prevent errors when the barcode is read. If left unspecified, the
// generator chooses a correction level based on the size of the message data.
// `alwaysSpecifyCompaction`: A [Boolean] value of `0` or `1` that determines
// the inclusion of information about the compaction mode in the barcode as an
// [NSNumber]. If a PDF417 barcode doesn’t contain compaction mode
// information, the reader assumes text-based compaction.
//
// The `compactionMode` property takes one of the following numeric values:
//
// [Table data omitted]
//
// Select either 1 or the appropriate valid value for your data that gives the
// most compact output.
//
// The following code creates a filter that generates a PDF417 barcode:
//
// [media-3546316]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pdf417BarcodeGenerator()
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PDF417BarcodeGenerator() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("PDF417BarcodeGenerator"))
	return CIFilterFromID(rv)
}

// Simulates the curl of a page, revealing the target image.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the page curl transition filter to an image. The effect
// transitions from one image to another by simulating a curling page,
// revealing the target image as the page curls.
//
// The page curl transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `backsideImage`: An image used as
// the backside of the curl with the type [CIImage]. `extent`: A [CGRect]
// representing the size of the effect. `time`: A `float` representing the
// parametric time of the transition from start (at time 0) to end (at time 1)
// as an [NSNumber]. `angle`: A `float` representing the angle of the motion
// of the curl as an [NSNumber]. `radius`: A `float` representing the radius
// of the curl as an [NSNumber].
//
// The following code creates a filter that produces a page curling back to
// reveal the target image.
//
// [media-3616422]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pageCurlTransition()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PageCurlTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pageCurlTransitionFilter"))
	return CIFilterFromID(rv)
}

// Simulates the curl of a page, revealing the target image with added shadow.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the page curl with shadow transition filter to an
// image. The effect transitions from one image to another by simulating a
// curling page, revealing the target image as the page curls with a shadow
// effect from the backside image.
//
// The page curl with shadow transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `backsideImage`: An image used as
// the backside of the curl with the type [CIImage]. `extent`: A [CIVector]
// representing the extent of the effect. `angle`: A `float` representing the
// angle of the motion, in radians as an [NSNumber]. `shadowAmount`: A `float`
// representing the strength of the shadow as an [NSNumber]. `shadowExtent`: A
// [CIVector] representing the rectangular portion of the input image that is
// used to create the shadow. `shadowSize`: A `float` representing the maximum
// amount of pixels to make up the shadow as an [NSNumber]. `time`: A `float`
// representing the parametric time of the transition from start (at time 0)
// to end (at time 1) as an [NSNumber].
//
// The following code creates a page curling back to reveal the target image
// with an added shadow.
//
// [media-3616423]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pageCurlWithShadowTransition()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PageCurlWithShadowTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pageCurlWithShadowTransitionFilter"))
	return CIFilterFromID(rv)
}

// Calculates the location of an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the palette centroid filter to an image. The filter
// locates colors in the input image that the palette image defines and
// `outputImage.Extent()` provides the location of the colors of the image.
// You can combine with other filters to create more sophisticated images.
//
// The palette centroid filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `paletteImage`: An image
// that has the dimensions of x 1 where represents the amount of colors in the
// image, with type [CIImage]. `perceptual`: A Boolean value that specifies if
// the filter applies the color palette in a perceptual color space.
//
// The following code creates a filter that calculates the extent of the
// palette color:
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/paletteCentroid()
func (_CIFilterClass CIFilterClass) PaletteCentroidFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("paletteCentroidFilter"))
	return CIFilterFromID(rv)
}

// Replaces colors with colors from a palette image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the palette filter to an image. The effect uses the
// palette image that is x 1 pixels in size containing a set of colors,
// replacing the image colors.
//
// The palettize filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `paletteImage`: An image
// with the dimensions of x 1 where represents the colors to add to the image,
// with type [CIImage]. `perceptual`: A Boolean value that specifies if the
// filter applies the color palette in a perceptual color space.
//
// The following code creates a filter that replaces the colors of the input
// image with the specified colors found in the palette image:
//
// [media-3558713]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/palettize()
func (_CIFilterClass CIFilterClass) PalettizeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("palettizeFilter"))
	return CIFilterFromID(rv)
}

// Warps the image to create a parallelogram and tiles the result.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the parallelogram tile filter to an image. The effect
// warps the input image to create a parallelogram and then tiles the result.
//
// The parallelogram tile filter uses the following properties:
//
// inputImage: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber]. `acuteAngle`: A `float` representing the primary angle for
// the repeating parallelogram tile.
//
// The following code creates a filter that results in the image being cropped
// to a parallelogram and then tiled:
//
// [media-3599887]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/parallelogramTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ParallelogramTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("parallelogramTileFilter"))
	return CIFilterFromID(rv)
}

// Creates a mask where red pixels indicate areas of the image that are likely
// to contain a person.
//
// # Return Value
//
// A [CIImage] containing the mask.
//
// # Discussion
//
// The person-segmentation filter creates a mask that contains red pixels in
// the areas of the input image that are likely to contain people.
//
// The person-segmentation filter takes the following properties:
//
// `inputIImage`: A [CIImage] containing the image to segment. `qualityLevel`:
// The size and quality of the resulting segmentation mask. 0 is accurate, `1`
// is balanced, and `2` is fast.
//
// The following code applies the person-segmentation filter to an image:
//
// [media-4407311]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/personSegmentation()
func (_CIFilterClass CIFilterClass) PersonSegmentationFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("personSegmentationFilter"))
	return CIFilterFromID(rv)
}

// Transforms an image’s perspective.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the perspective correction filter to an image. The
// effect applies a perspective correction transforming nonrectangular area in
// the source image to a rectangular output image.
//
// The perspective correction filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
//
// The following code creates a filter that corrects the perspective to appear
// straight:
//
// [media-3582228]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveCorrection()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) PerspectiveCorrectionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("perspectiveCorrectionFilter"))
	return CIFilterFromID(rv)
}

// Rotates an image in a 3D space.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the perspective rotate filter to an image. The effect
// rotates the image in 3D space to simulate the observer changing viewing
// position.
//
// The perspective rotate filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `pitch`: A `float`
// representing the adjustment along the pitch axis in 3D space as an
// [NSNumber]. `yaw`: A `float` representing the adjustment along the vertical
// axis as an [NSNumber]. `roll`: A `float` representing the amount of
// horizontal axis in 3D space as an [NSNumber]. `focalLength`: A `float`
// representing the simulated focal length as an [NSNumber].
//
// The following code creates a filter that rotates the image:
//
// [media-3582225]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveRotate()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PerspectiveRotateFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("perspectiveRotateFilter"))
	return CIFilterFromID(rv)
}

// Tiles an image by adjusting the perspective of the image.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the perspective tile filter to an image. The effect
// adjusts the perspective of the image and then tiles the result.
//
// The perspective tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] of
// the input image mapped to the top-left corner of the tile. `topRight`: A
// [CGPoint] of the input image mapped to the top-right corner of the tile.
// `bottomLeft`: A [CGPoint] of the input image mapped to the bottom-left
// corner of the tile. `bottomRight`: A [CGPoint] of the input image mapped to
// the bottom-right corner of the tile.
//
// The following code creates a filter that tiles the image and adjusts the
// perspective to add depth:
//
// [media-3599879]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) PerspectiveTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("perspectiveTileFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s geometry to adjust the perspective.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the perspective transform filter to an image. The
// effect alters the geometry of an image to simulate the observer changing
// viewing position. You can use the perspective filter to skew an image.
//
// The perspective transform filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
//
// The following code creates a filter that changes the perspective of the
// input image:
//
// [media-3582227]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTransform()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) PerspectiveTransformFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("perspectiveTransformFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s geometry to adjust the perspective while applying
// constraints.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the perspective transform with extent filter to an
// image. The effect alters the geometry of an image to simulate the observer
// changing viewing position. The extent filter crops the image within the
// bounds specified. You can use the perspective filter to skew an image.
//
// The perspective transform with extent filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `topLeft`: A [CGPoint] in
// the input image mapped to the top-left corner of the output image.
// `topRight`: A [CGPoint] in the input image mapped to the top-right corner
// of the output image. `bottomLeft`: A [CGPoint] in the input image mapped to
// the bottom-left corner of the output image. `bottomRight`: A [CGPoint] in
// the input image mapped to the bottom-right corner of the output image.
// `extent`: A [CGRect] representing the dimensions of the output image.
//
// The following code creates a filter that changes the perspective of the
// input image:
//
// [media-3582226]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/perspectiveTransformWithExtent()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) PerspectiveTransformWithExtentFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("perspectiveTransformWithExtentFilter"))
	return CIFilterFromID(rv)
}

// Exaggerates an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate vintage
// photography film with higher contrast.
//
// The photo effect chrome filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in diminished color in the
// input image:
//
// [media-3545019]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectChrome()
func (_CIFilterClass CIFilterClass) PhotoEffectChromeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectChromeFilter"))
	return CIFilterFromID(rv)
}

// Diminishes an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate vintage
// photography film with diminished color.
//
// The photo effect fade filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in a desaturated image:
//
// [media-3545034]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectFade()
func (_CIFilterClass CIFilterClass) PhotoEffectFadeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectFadeFilter"))
	return CIFilterFromID(rv)
}

// Desaturates an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate vintage
// photography film with desaturated colors.
//
// The photo effect instant filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the input image
// becoming desaturated:
//
// [media-3545013]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectInstant()
func (_CIFilterClass CIFilterClass) PhotoEffectInstantFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectInstantFilter"))
	return CIFilterFromID(rv)
}

// Adjust an image’s colors to black and white.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate
// black-and-white photography film with low contrast.
//
// The photo effect mono filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that produces a black-and-white image:
//
// [media-3545035]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectMono()
func (_CIFilterClass CIFilterClass) PhotoEffectMonoFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectMonoFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s colors to black and white and intensifies the
// contrast.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate
// black-and-white photography film and intensifies the contrast.
//
// The photo effect noir filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in a black-and-white
// image:
//
// [media-3545032]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectNoir()
func (_CIFilterClass CIFilterClass) PhotoEffectNoirFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectNoirFilter"))
	return CIFilterFromID(rv)
}

// Lowers the contrast of the input image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate vintage
// photography film with emphasized cool colors.
//
// The photo effect process filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a lower contrast to the input
// image:
//
// [media-3545027]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectProcess()
func (_CIFilterClass CIFilterClass) PhotoEffectProcessFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectProcessFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s colors to black and white.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate
// black-and-white photography film without significantly altering contrast.
//
// The photo effect tonal filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that produces a grayscale image:
//
// [media-3545036]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectTonal()
func (_CIFilterClass CIFilterClass) PhotoEffectTonalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectTonalFilter"))
	return CIFilterFromID(rv)
}

// Brightens an image’s colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate vintage
// photography film and emphasize warm colors.
//
// The photo effect transfer filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545015]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/photoEffectTransfer()
func (_CIFilterClass CIFilterClass) PhotoEffectTransferFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("photoEffectTransferFilter"))
	return CIFilterFromID(rv)
}

// Blends colors of two images by replacing brighter colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the pin-light-blend mode filter to an image. The effect
// calculates the brightness of the colors in the background image. Colors
// that are lighter than 50 percent gray remain unchanged, while the filter
// replaces colors that are darker than 50 percent with the corresponding
// color values of the input image. The effect then uses the calculated result
// to create the output image.
//
// The pin-light-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// darker with more saturation and the colors of the gradient image:
//
// [media-3546414]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pinLightBlendMode()
func (_CIFilterClass CIFilterClass) PinLightBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pinLightBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by creating a pinch effect with stronger distortion in
// the center.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the pinch distortion filter to an image. This effect
// creates a rectangular area that pinches source pixels inward, distorting
// those pixels closest to the rectangle the most.
//
// The pinch distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `scale`: A
// float representing the amount of pinching effect as an [NSNumber].
// `radius`: A float representing the amount of pixels used to create the
// distortion as an [NSNumber].
//
// The following code creates a filter that results in a distorted image from
// the center of the photo:
//
// [media-4407323]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pinchDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PinchDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pinchDistortionFilter"))
	return CIFilterFromID(rv)
}

// Enlarges the colors of the pixels to create a blurred effect.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the pixelate filter to an image. The effect produces a
// result by mapping the image to colored squares with colors defined by the
// pixels of the input image.
//
// The pixelate filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `scale`: A
// `float` representing the size of the squares as an [NSNumber].
//
// The following code creates a filter that results in a distorted image made
// of squares:
//
// [media-3599999]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pixellate()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) PixellateFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pixellateFilter"))
	return CIFilterFromID(rv)
}

// Applies a pointillize effect to an image.
//
// # Return Value
//
// A [CIImage] containing the pointillized image.
//
// # Discussion
//
// This filter applies a pointillize effect to an image. The effect generates
// an output image made of small, single-color, circular points distributed on
// a randomly perturbed grid.
//
// The pointillize filter uses the following properties:
//
// `inputImage`: A [CIImage] containing the input image. `radius`: The radius
// in pixels of the circular points. center: Determines the origin of the
// grid.
//
// The following code applies the pointillize filter with a radius of 40
// pixels.
//
// [media-4333706]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/pointillize()
func (_CIFilterClass CIFilterClass) PointillizeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("pointillizeFilter"))
	return CIFilterFromID(rv)
}

// Generates a quick response (QR) code image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a QR code as an image. QR codes are a high-density
// matrix barcode format defined in the ISO/IEC 18004:2006 standard.
//
// The QR code generator filter uses the following properties:
//
// `message`: A `string` representing the data to be encoded as a QR Code as
// [NSData]. `correctionLevel`: A single letter `string` representing the
// error-correction format as an [NSString]. L is 7 precent correction, M is
// 15 precent correction, Q is 25 precent correction, and H is 30 precent
// correction.
//
// The following code creates a filter that generates a QR code:
//
// [media-3546313]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/qrCodeGenerator()
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (_CIFilterClass CIFilterClass) QRCodeGenerator() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("QRCodeGenerator"))
	return CIFilterFromID(rv)
}

// Generates a gradient that varies radially between two circles having the
// same center.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a radial-gradient image. The effect generates a color
// shift between the `radius0` and `radius1` properties.
//
// The radial-gradient filter uses the following properties:
//
// `center`: A [CGPoint] representing the center of the effect as x and y
// coordinates. `color0`: A [CIColor] representing the first color to use in
// the gradient. `color1`: A [CIColor] representing the second color to use in
// the gradient. `radius0`: A `float` representing the radius of the starting
// circle to use in the gradient as a [NSNumber]. `radius1`: A `float`
// representing the radius of the ending circle to use in the gradient as a
// [NSNumber].
//
// The following code creates a filter that generates a gradient image:
//
// [media-3558800]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/radialGradient()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) RadialGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("radialGradientFilter"))
	return CIFilterFromID(rv)
}

// Generates a random filter image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates an image with infinite extent. The image pixels
// values are from one of four independent, uniformly distributed random
// colors.
//
// The following code creates a filter that generates a random color image:
//
// [media-3590973]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/randomGenerator()
func (_CIFilterClass CIFilterClass) RandomGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("randomGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Simulates a ripple in a pond to transiton from one image to another.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the ripple transition filter to an image. The effect
// transitions from one image to another by creating a circular wave that
// expands from the center point, revealing the target image through the wave
// effect.
//
// The ripple transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `center`: A set of coordinates
// marking the center of the image as a [CGPoint]. `width`: A `float`
// representing the width of the ripple effect as an [NSNumber]. `extent`: A
// [CGRect] representing the size of the ripple effect. `scale`: A `float`
// representing the scale of the effect as an [NSNumber]. `time`: A `float`
// representing the parametric time of the transition from start (at time 0)
// to end (at time 1) as an [NSNumber].
//
// The following code creates a filter that transitions from the input image
// to the target image with a water-like ripple effect.
//
// [media-3616425]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/rippleTransition()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) RippleTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("rippleTransitionFilter"))
	return CIFilterFromID(rv)
}

// Generates a rounded rectangle image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a rounded rectangle image with the specified size,
// corner radius, and color properties.
//
// The rounded rectangle generator filter uses the following properties:
//
// `color`: A [CIColor] representing the color of the rounded rectangle.
// `extent`: A [CGRect] representing the size of the rounded rectangle.
// `radius`: A `float` representing the curve of the rectangle’s corners.
//
// The following code creates a filter that generates a light blue square with
// rounded corners:
//
// [media-3546319]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedRectangleGenerator()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) RoundedRectangleGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("roundedRectangleGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Creates an image containing the outline of a rounded rectangle.
//
// # Return Value
//
// A [CIImage] containing the stroked rectangle.
//
// # Discussion
//
// This filter creates an outline of a rounded rectangle.
//
// The filter takes the following properties:
//
// `extent`: A [CGRect] containing the position and size of the rectangle.
// `width`: The width of the stroke to draw. `radius`: The corner radius.
//
// The following code generates an image containing a stroked rounded
// rectangle:
//
// [media-4407287]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/roundedRectangleStrokeGenerator()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) RoundedRectangleStrokeGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("roundedRectangleStrokeGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Calculates the average color for the specified row of pixels in an image.
//
// # Return Value
//
// # Discussion
//
// This method applies the row average filter to an image. This effect
// calculates the average color for a horizontal row over a region defined by
// `extent`. The height of the extent determines the width of the resulting
// image. The height is always 1 pixel.
//
// The row average filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `extent`: A [CGRect] that
// specifies the subregion of the image that you want to process.
//
// The following code creates a filter that calculates the row average for the
// middle section of an image:
//
// [media-4331788]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/rowAverage()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (_CIFilterClass CIFilterClass) RowAverageFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("rowAverageFilter"))
	return CIFilterFromID(rv)
}

// Converts the colors in an image from sRGB to linear.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the sRGB-tone-curve-to-linear filter to an image. The
// effect converts an image in sRGB space to linear color space.
//
// The sRGB-tone-curve-to-linear filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that converts from sRGB to linear color
// space.
//
// [media-4333632]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sRGBToneCurveToLinear()
func (_CIFilterClass CIFilterClass) SRGBToneCurveToLinearFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sRGBToneCurveToLinearFilter"))
	return CIFilterFromID(rv)
}

// Creates a saliency map from an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the saliency map filter to an image. The effect
// generates a saliency map representation of the input image.
//
// The saliency map filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that produces an image that’s easier
// for computers to analyze:
//
// [media-3624695]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/saliencyMap()
func (_CIFilterClass CIFilterClass) SaliencyMapFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("saliencyMapFilter"))
	return CIFilterFromID(rv)
}

// Blends the colors and saturation values of two images.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the saturation-blend mode filter to an image. The
// effect uses the values of the hue and luminance from the background image
// with the saturation of the input image to produce the output.
//
// The saturation-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image taking on the
// colors of the background image with low saturated values becoming gray.
//
// [media-3546404]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/saturationBlendMode()
func (_CIFilterClass CIFilterClass) SaturationBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("saturationBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Blends colors of two images by multiplying colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the screen-blend mode filter to an image. The effect
// calculates the colors in the output image by multiplying the inverse color
// values for the input and background images, resulting in a brighter image.
//
// The screen-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// lighter with colors from the input and background image:
//
// [media-3546401]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/screenBlendMode()
func (_CIFilterClass CIFilterClass) ScreenBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("screenBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s colors to shades of brown.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the sepia tone filter to an image. The effect maps the
// colors of the `inputImage` to various shades of brown.
//
// The sepia tone filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that results in the input image
// transforming to a brown hue:
//
// [media-3545028]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sepiaTone()
func (_CIFilterClass CIFilterClass) SepiaToneFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sepiaToneFilter"))
	return CIFilterFromID(rv)
}

// Creates a shaded image from a height-field image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the shaded material filter to an image. The effect
// produces a shaded image from a height-field image. Areas of the height
// field image that have a darker shaded area produce a stronger effect. You
// can combine the filter with [CIHeightFieldFromMask] to produce quick
// shadings of masks, such as text.
//
// The shaded material filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `shadingImage`: An image
// representing the color shading effect with type [CIImage]. `scale`: A
// `float` representing the strength of effect as an [NSNumber].
//
// The following code creates a filter that results in an image containing
// glossy text by applying the shading image.
//
// [media-3600005]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/shadedMaterial()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ShadedMaterialFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("shadedMaterialFilter"))
	return CIFilterFromID(rv)
}

// Applies a sharpening effect to an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the sharpen luminance filter to an image. The effect
// increases image detail by adjusting the luminance of each pixel within the
// `radius` property. Sharpening the luminance doesn’t effect the chroma
// data of each pixel.
//
// The bicubic sharpen luminance filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber]. `sharpness`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in detail from the sign in
// the image to be more visible:
//
// [media-3595821]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sharpenLuminance()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SharpenLuminanceFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sharpenLuminanceFilter"))
	return CIFilterFromID(rv)
}

// Produces a tiled image from a source image by applying a six-way reflected
// symmetry.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This filter produces a tiled image from a source image by applying a
// six-way reflected symmetry.
//
// The six-fold reflected tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion , in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in a six-fold pattern
// repeated and angled for distortion:
//
// [media-3599882]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sixfoldReflectedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SixfoldReflectedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sixfoldReflectedTileFilter"))
	return CIFilterFromID(rv)
}

// Creates a tiled image by rotating in increments of 60 degrees.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the six-fold reflected tile filter to an image. The
// effect produces a tiled image by rotating the image in increments 60
// degrees.
//
// The six-fold rotated tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a CGPoint. `angle`: A
// `float` representing the direction of distortion , in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in flowers in the input
// image becoming rotated by 60 degrees and tiled to create the output:
//
// [media-3599892]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sixfoldRotatedTile()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SixfoldRotatedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sixfoldRotatedTileFilter"))
	return CIFilterFromID(rv)
}

// Generates a gradient that blends colors along a linear axis between two
// defined endpoints.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a smooth linear-gradient image. The effect creates a
// gradient by gradually blending colors between `point0` and `point1` using
// the sigmoid curve function.
//
// The smooth linear-gradient filter uses the following properties:
//
// `point0`: A [CGPoint] representing the starting position of the gradient.
// `point1`: A [CGPoint] representing the ending position of the gradient.
// `color0`: A [CIColor] representing the first color used in the gradient.
// `color1`: A [CIColor] representing the second color used in the gradient.
//
// The following code creates a filter that generates a gradient image:
//
// [media-3558802]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/smoothLinearGradient()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (_CIFilterClass CIFilterClass) SmoothLinearGradientFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("smoothLinearGradientFilter"))
	return CIFilterFromID(rv)
}

// Calculates the Sobel gradients for an image.
//
// # Return Value
//
// A [CIImage] containing the Sobel gradients.
//
// # Discussion
//
// This filter applies the Sobel operator to the color components of the input
// image. You would typically use the Sobel filter as part of an
// edge-detection algorithm for performing.
//
// `inputImage`: A [CIImage] containing the image to process.
//
// The following code applies the [SobelGradientsFilter] filter to an image.
//
// [media-4407283]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sobelGradients()
func (_CIFilterClass CIFilterClass) SobelGradientsFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sobelGradientsFilter"))
	return CIFilterFromID(rv)
}

// Blurs the colors of two images by calculating luminance.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the soft-light blend mode filter to an image. The
// effect calculates the brightness of the colors in the background image.
// Colors that are lighter than 50 percent gray become lighter, while the
// filter further darkens colors that are darker than 50 percent. The filter
// then uses the calculated result to create the output image.
//
// The soft-light blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// slightly darker with more saturation and the colors of the gradient image:
//
// [media-3546420]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/softLightBlendMode()
func (_CIFilterClass CIFilterClass) SoftLightBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("softLightBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Overlaps two images to create one cropped image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the source-atop compositing filter to an image. The
// effect creates the result by overlaying the input image over the background
// image. The filter then removes the area that doesn’t overlap with the
// background image.
//
// The source-atop compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in an output image that
// shows the background image with the portion of the input image that
// overlaps it:
//
// [media-3546396]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceAtopCompositing()
func (_CIFilterClass CIFilterClass) SourceAtopCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sourceAtopCompositingFilter"))
	return CIFilterFromID(rv)
}

// Subtracts non-overlapping areas of two images, resulting in one image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the source-in compositing filter to an image. The
// effect creates the result by overlaying the input image over the background
// image. The filter then removes the non-overlapping area of both images.
//
// The source-in compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in an output image that
// shows the portion of the background image that overlaps with the input:
//
// [media-3546393]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceInCompositing()
func (_CIFilterClass CIFilterClass) SourceInCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sourceInCompositingFilter"))
	return CIFilterFromID(rv)
}

// Subtracts overlapping area of two images to create the output image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the source-out compositing filter to an image. The
// effect creates the result by overlaying the input image over the background
// image. The filter then removes the overlapping area of the background image
// from the result.
//
// The source-out compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in an output image that
// shows the portion of the background image that doesn’t overlap with the
// input image:
//
// [media-3546394]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceOutCompositing()
func (_CIFilterClass CIFilterClass) SourceOutCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sourceOutCompositingFilter"))
	return CIFilterFromID(rv)
}

// Places one image over a second image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the source-over compositing filter to an image. The
// effect creates the result by overlaying the input image over the background
// image. Unlike the other source compositing filters, source-over doesn’t
// subtract parts of the image.
//
// The source-over compositing filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in both of the input
// images becoming visible with no subtraction:
//
// [media-3546395]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sourceOverCompositing()
func (_CIFilterClass CIFilterClass) SourceOverCompositingFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sourceOverCompositingFilter"))
	return CIFilterFromID(rv)
}

// Replaces colors of an image with specifed colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the spot color filter to an image. The effect replaces
// one or more of the color ranges of the input image with properties.
//
// The spot color filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `centerColor1`: A [CIColor]
// representing the median value of the first color to be replaced.
// `centerColor2`: A [CIColor] representing the median value of the second
// color to be replaced. `centerColor3`: A [CIColor] representing the median
// value of the third color to be replaced. `replacementColor1`: A [CIColor]
// to replace the first color. `replacementColor2`: A [CIColor] to replace the
// second color. `replacementColor3`: A [CIColor] to replace the third color.
// `closeness1`: A `float` representing how closely the first center color
// must match before it’s replaced. `closeness2`: A `float` representing how
// closely the second center color must match before it’s replaced.
// `closeness3`: A `float` representing how closely the third center color
// must match before it’s replaced. `contrast1`: A `float` representing the
// contrast of the first replacement color as an [NSNumber]. `contrast2`: A
// `float` representing the contrast of the second replacement color as an
// [NSNumber]. `contrast3`: A `float` representing the contrast of the third
// replacement color as an [NSNumber].
//
// The following code creates a filter that replaces the colors of the input
// image with the specified colors:
//
// [media-3600008]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/spotColor()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SpotColorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("spotColorFilter"))
	return CIFilterFromID(rv)
}

// Highlights a definined area of the image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the spotlight filter to an image. The effect applies a
// directional spotlight effect to an image while creating a transparent area
// not highlighted by the spotlight.
//
// The spotlight filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `lightPointsAt`: A
// [CIVector] with the x and y positions that the spotlight points at.
// `brightness`: A `float` representing the brightness of the spotlight as an
// [NSNumber]. `lightPosition`: A [CIVector] containing the x and y position
// of the spotlight. `concentration`: A `float` representing the size of the
// spotlight in pixels as an [NSNumber]. `color`: A [CIColor] representing the
// spotlight color.
//
// The following code creates a filter that results in only the bottom left of
// the image becoming visible while the rest of the image gradually becomes
// transparent:
//
// [media-3600004]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/spotLight()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SpotLightFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("spotLightFilter"))
	return CIFilterFromID(rv)
}

// Generates a star-shine image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a star-shine image. The effect is similar to a
// supernova effect. You can use this filter to simulate a lens flare.
//
// The star-shine generator filter uses the following properties:
//
// `center`: A `vector` representing the center of the flare as a [CGPoint].
// `color`: A color representing the color of the flare as a [cgColor].
// `radius`: A `float` representing the radius of the flare as an [NSNumber].
// `crossScale`: A `float` representing the cross flare size relative to the
// round central flare as an [NSNumber]. `crossAngle`: A `float` representing
// the angle of the flare as an [NSNumber]. `crossOpacity`: A `float`
// representing the thickness of the cross opacity as an [NSNumber].
// `crossWidth`: A `float` representing the cross width as an [NSNumber].
// `epsilon`: A `float` representing the epsilon as an [NSNumber].
//
// The following code generates a star-shaped silhouette with a black
// background.
//
// [media-3590972]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/starShineGenerator()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [cgColor]: https://developer.apple.com/documentation/UIKit/UIColor/cgColor
func (_CIFilterClass CIFilterClass) StarShineGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("starShineGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Rotates and crops an image.
//
// # Return Value
//
// The adjusted image.
//
// # Discussion
//
// This method applies the straighten filter to an image. The effect rotates
// the image based on the `angle` property while cropping and scaling the
// image to remain the same size as the original image.
//
// The straighten filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `angle`: A `float`
// representing the angle to rotate the image as an [NSNumber].
//
// The following code creates a filter that rotates the image 135 degrees:
//
// [media-3582229]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/straighten()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) StraightenFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("straightenFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by stretching or cropping to fit a specified size.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the stretch crop filter to an image. This effect
// distorts an image by stretching an image and then applies the crop extent.
// If the crop value is 0, the filter only uses stretching. If the value is 1,
// then the filter only uses cropping.
//
// The stretch crop filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `centerStretchAmount`: A
// `float` representing the amount of stretching of the center of the image as
// an [NSNumber]. `size`: A [CGPoint] representing the desired size of the
// output image. `cropAmount`: A `float` representing the amount of cropping
// you apply to achieve the target size as an [NSNumber].
//
// The following code creates a filter that results in a smaller image
// that’s distorted and cropped to be the defined size:
//
// [media-4407279]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/stretchCrop()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) StretchCropFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("stretchCropFilter"))
	return CIFilterFromID(rv)
}

// Generates a line of stripes as an image
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a vertical stripped line pattern as an image.
//
// The stripes generator filter uses the following properties:
//
// `center`: A [CIVector] representing the center of the image. `color0`: A
// [CIColor] representing the stripes color. `color1`: A [CIColor]
// representing the background color. `width`: A `float` representing the
// width of the lines as an [NSNumber]. `sharpness`: A `float` representing
// the sharpness of the lines as an [NSNumber].
//
// The following code creates a filter that generates a black and white
// vertical striped image:
//
// [media-3590971]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/stripesGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) StripesGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("stripesGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Blends colors by subtracting color values from two images.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the subtract-blend mode filter to an image. The effect
// calculates the colors in the output image by subtracting the color values
// that differ between the background image and the input image.
//
// The subtract-blend mode filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `backgroundImage`: An image
// with the type [CIImage].
//
// The following code creates a filter that results in the image becoming
// darker with less detail:
//
// [media-3546402]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/subtractBlendMode()
func (_CIFilterClass CIFilterClass) SubtractBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("subtractBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Generates an image resembling the sun.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a sunbeam as an image. The effect generates a
// center-textured sun with striations. You can combine with other filters to
// create more sophisticated images.
//
// The sunbeams generator filter uses the following properties:
//
// `center`: A vector representing the center of the image as a [CIVector].
// `color`: A [CIColor] representing the color of the sun. `sunRadius`: A
// `float` representing the radius of the center sun as an [NSNumber].
// `maxStriationRadius`: A `float` representing the striation radius as an
// [NSNumber]. `striationStrength`: A `float` representing the striation
// strength as an [NSNumber]. `striationContrast`: A `float` representing the
// striation contrast as an [NSNumber]. `time`: A `float` representing the
// time as an [NSNumber].
//
// The following code creates a filter that generates an image that resembles
// a yellow sun with sunbeams:
//
// [media-3546315]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/sunbeamsGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SunbeamsGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("sunbeamsGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Gradually transitions from one image to another with a swiping motion.
//
// # Return Value
//
// The transition image.
//
// # Discussion
//
// This method applies the swipe transition filter to an image. The effect
// transitions from the input image to the target image by simulating a
// swiping motion.
//
// The swipe transition filter uses the following properties:
//
// `inputImage`: The starting image with the type [CIImage]. `targetImage`:
// The ending image with the type [CIImage]. `extent`: A [CGRect] representing
// the size of the rounded rectangle. `time`: A `float` representing the
// parametric time of the transition from start (at time 0) to end (at time 1)
// as an [NSNumber]. angle: A `float` representing the angle of the motion of
// the swipe as an [NSNumber]. width: A `float` representing the width of the
// swipe effect as an [NSNumber]. opacity: A `float` representing the
// transparency of the swipe as an [NSNumber].
//
// The following code creates a filter that transitions from the input image
// to the target image with a gradual fade from left to right.
//
// [media-3616424]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/swipeTransition()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) SwipeTransitionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("swipeTransitionFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s temperature and tint.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the image temperature and tint filter to an image. The
// effect adjusts the white balance of the input image to match the
// `targetNeutral` property, resulting in a cooler or warmer tone image.
//
// The temperature and tint filter uses the following properties:
//
// `neutral`: A `vector` containing the source white point as a [CIVector].
// `targetNeutral`: A vector containing the desired white point as a
// [CIVector]. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds an orange hue to the input
// image:
//
// [media-3545008]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/temperatureAndTint()
func (_CIFilterClass CIFilterClass) TemperatureAndTintFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("temperatureAndTintFilter"))
	return CIFilterFromID(rv)
}

// Generates a text image.
//
// # Return Value
//
// The generated image.
//
// # Discussion
//
// This method generates a text image. The effect takes the input string
// property and the scale factor to scale up the text. You commonly combine
// this filter with other filters to create a watermark on images.
//
// The text image generator filter uses the following properties:
//
// `text`: The `string` to render. The string can contain non-ASCII
// characters. `fontName`: A `string` representing the name of the font to be
// used to generate the image. `fontSize`: A `float` representing the size of
// the font as an [NSNumber]. `scaleFactor`: A `float` representing the scale
// of the font for the generated text as an [NSNumber].
//
// The following code creates a filter that generates a string of text as a
// grayscale image:
//
// [media-3546321]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/textImageGenerator()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TextImageGeneratorFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("textImageGeneratorFilter"))
	return CIFilterFromID(rv)
}

// Alters the image to make it look like it was taken by a thermal camera.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that make it appear like
// a thermal camera captured the image.
//
// The thermal filter uses the following property:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a thermal effect to the input
// image:
//
// [media-3545033]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/thermal()
func (_CIFilterClass CIFilterClass) ThermalFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("thermalFilter"))
	return CIFilterFromID(rv)
}

// Alters an image’s tone curve according to a series of data points.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the tone curve filter to an image. The effect
// calculates the adjustment of the tone curve by the sum of the red, green,
// and blue color values with the point value properties specified.
//
// The tone curve filter uses the following properties:
//
// `point0`: A v`ector` containing the position of the first point of the tone
// curve as a [CIVector]. `point1`: A v`ector` containing the position of the
// second point of the tone curve as a [CIVector]. `point2`: A `vector`
// containing the position of the third point of the tone curve as a
// [CIVector]. `point3`: A v`ector` containing the position of the fourth
// point of the tone curve as a [CIVector]. `point4`: A v`ector` containing
// the position of the fifth point of the tone curve as a [CIVector].
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds brightness to the input
// image:
//
// [media-3545005]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/toneCurve()
func (_CIFilterClass CIFilterClass) ToneCurveFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("toneCurveFilter"))
	return CIFilterFromID(rv)
}

// Creates a torus-shaped lens to distort the image.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the torus lens distortion filter to an image. This
// effect distorts an image by creating a torus-shaped object, placing it over
// the input image, and applying the refraction.
//
// The torus lens distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `radius`: A
// `float` representing the amount of pixels the filter uses in the tours as
// an [NSNumber]. `refraction`: A `float` representing the refraction of the
// glass as an [NSNumber]. `width`: A `float` representing the width of the
// torus ring as an [NSNumber].
//
// The following code creates a filter that results in a torus-shaped object
// placed over the image:
//
// [media-4407288]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/torusLensDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TorusLensDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("torusLensDistortionFilter"))
	return CIFilterFromID(rv)
}

// Create a triangular kaleidoscope effect and then tiles the result.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the triangle kaleidoscope filter to an image. The
// effect produces a complex tiled pattern from a triangular area input image.
//
// The triangle kaleidoscope tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `decay`: A `float`
// representing the intensity of the color fade from the center of the
// triangle as an [NSNumber]. `point`: A set of coordinates marking the center
// of the triangular area of the input image as a [CIVector]. `rotation`: A
// `float` representing the angle of rotation of the triangle as an
// [NSNumber]. `size`: A `float` representing the size in pixels of the
// triangle as an [NSNumber].
//
// The following code creates a filter that produces a triangle tile of the
// input image, creating an optical illusion:
//
// [media-3599889]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/triangleKaleidoscope()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TriangleKaleidoscopeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("triangleKaleidoscopeFilter"))
	return CIFilterFromID(rv)
}

// Tiles a triangular area of an image.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the triangle tile filter to an image. The effect
// creates a tiled pattern from a triangular area from the input image.
//
// The triangle tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that produces a triangle of the input
// image and tiles the result:
//
// [media-3599891]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/triangleTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TriangleTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("triangleTileFilter"))
	return CIFilterFromID(rv)
}

// Creates a tiled image by rotating in increments of 30 degrees.
//
// # Return Value
//
// The tiled image.
//
// # Discussion
//
// This method applies the 12-fold reflected tile filter to an image. The
// effect produces a 12-way reflected tile image.
//
// The 12-fold reflected tile filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `center`: A set of
// coordinates marking the center of the image as a [CGPoint]. `angle`: A
// `float` representing the direction of distortion, in radians as an
// [NSNumber]. `width`: A `float` representing the set width of each tile as
// an [NSNumber].
//
// The following code creates a filter that results in a 12-fold pattern
// angled at 30 degrees and then repeated:
//
// [media-3599890]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/twelvefoldReflectedTile()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TwelvefoldReflectedTileFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("twelvefoldReflectedTileFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by rotating pixels around a center point.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the twirl distortion filter to an image. This effect
// distorts an image by rotating pixels around the defined center to create a
// twirling effect. You can specify the number of rotations to control the
// strength of the effect.
//
// The twirl distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the amount of pixels the filter uses to create the distortion
// as an [NSNumber]. `center`: A set of coordinates marking the center of the
// image as a [CGPoint]. `angle`: A `float` representing the angle of the
// twirl, in radians, as an [NSNumber].
//
// The following code creates a filter that results in the center of the image
// becoming twirled:
//
// [media-4407322]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/twirlDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) TwirlDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("twirlDistortionFilter"))
	return CIFilterFromID(rv)
}

// Increases an image’s contrast between two colors.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the unsharp mask filter to an image. The effect
// increases the contrast of the edge between pixels of different colors
// within the defined radius property.
//
// The unsharp mask filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `radius`: A `float`
// representing the area of effect as an [NSNumber]. `intensity`: A `float`
// representing the desired strength of the effect as an [NSNumber].
//
// The following code creates a filter that results in the objects within the
// image becoming darker:
//
// [media-3595819]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/unsharpMask()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) UnsharpMaskFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("unsharpMaskFilter"))
	return CIFilterFromID(rv)
}

// Adjusts an image’s vibrancy.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the vibrance filter to an image. The effect adjusts the
// saturation of the image while preserving skin tone colors.
//
// The vibrance filter uses the following properties:
//
// `amount`: A `float` representing the amount to adjust the saturation with
// the type [NSNumber]. `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds darkness to the input image:
//
// [media-3544997]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vibrance()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) VibranceFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("vibranceFilter"))
	return CIFilterFromID(rv)
}

// Gradually darkens a specified area of an image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the vignette effect filter to an image. This effect
// reduces brightness of the image at the periphery of a specified region.
//
// The vignette effect filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `intensity`: A `float`
// representing the intensity of the vignette effect as an [NSNumber].
// `radius`: A `float` representing the radius of the effect as an [NSNumber].
// `falloff`: A `float` representing the fall off of brightness toward the
// edge of the image as an [NSNumber]. `center`: A [CGPoint] representing the
// center of the image.
//
// The following code creates a filter that darkens the edges of an area on
// the input image:
//
// [media-3600013]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vignetteEffect()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) VignetteEffectFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("vignetteEffectFilter"))
	return CIFilterFromID(rv)
}

// Gradually darkens an image’s edges.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the vignette filter to an image. This is a
// preconfigured effect that reduces brightness of the image at the periphery.
//
// The vignette filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `intensity`: A `float`
// representing the intensity of the vignette effect as an [NSNumber].
// `radius`: A `float` representing the radius of the effect as an [NSNumber].
//
// The following code creates a filter that darkens the edges of the input
// image:
//
// [media-3545014]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vignette()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) VignetteFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("vignetteFilter"))
	return CIFilterFromID(rv)
}

// A combination of color-burn and color-dodge blend modes.
//
// # Return Value
//
// The blended image as a [CIImage].
//
// # Discussion
//
// The vivid-light blend mode combines the color-dodge and color-burn blend
// modes (rescaled so that neutral colors become middle gray). If the input
// Images values are lighter than middle gray, the filter uses dodge; for
// darker values, the filter uses burn.
//
// `inputImage`: A [CIImage] containing the input image. `backgroundImage`: A
// [CIImage] containing the background image.
//
// The following code sample applies the vivid-light blend mode filter to two
// images:
//
// [media-4407307]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vividLightBlendMode()
func (_CIFilterClass CIFilterClass) VividLightBlendModeFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("vividLightBlendModeFilter"))
	return CIFilterFromID(rv)
}

// Distorts an image by using a vortex effect created by rotating pixels
// around a point.
//
// # Return Value
//
// The distorted image.
//
// # Discussion
//
// This method applies the vortex distortion filter to an image. This effect
// distorts an image by rotating pixels around the defined center to simulate
// a vortex. You can specify the number of rotations to control the strength
// of the effect.
//
// The vortex distortion filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage]. `angle`: A `float`
// representing the angle of the vortex, in radians, as an [NSNumber].
// `center`: A set of coordinates marking the center of the image as a
// [CGPoint]. `radius`: A `float` representing the amount of pixels the filter
// uses to create the distortion as an [NSNumber].
//
// The following code creates a filter that results in a small vortex effect:
//
// [media-4407301]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/vortexDistortion()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) VortexDistortionFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("vortexDistortionFilter"))
	return CIFilterFromID(rv)
}

// Adjusts the image’s white-point.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies the white-point adjust filter to an image. The effect
// adjusts the white-point of the input image by mapping all shades of gray to
// shades of the color property.
//
// The white-point adjust filter uses the following properties:
//
// `color`: The new white point color with the type of [CIColor].
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds a red hue to the input image:
//
// [media-3545004]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/whitePointAdjust()
func (_CIFilterClass CIFilterClass) WhitePointAdjustFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("whitePointAdjustFilter"))
	return CIFilterFromID(rv)
}

// Alters an image to make it look like an X-ray image.
//
// # Return Value
//
// The modified image.
//
// # Discussion
//
// This method applies a preconfigured set of effects that imitate a photo
// being scanned by an X-ray.
//
// The X-ray filter uses the following properties:
//
// `inputImage`: An image with the type [CIImage].
//
// The following code creates a filter that adds an X-ray effect to the input
// image:
//
// [media-3545022]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/xRay()
func (_CIFilterClass CIFilterClass) XRayFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("xRayFilter"))
	return CIFilterFromID(rv)
}

// Creates a zoom blur centered around a single point on the image.
//
// # Return Value
//
// The blurred image.
//
// # Discussion
//
// This method applies the zoom blur filter to an image. This effect mimics
// the zoom of a camera when capturing the image.
//
// The zoom blur filter uses the following properties:
//
// `amount`: A `float` representing the zoom-in amount as an [NSNumber].
// `center`: A set of coordinates marking the center of the image as a
// [CGPoint]. `inputImage`: A [CIImage] representing the input image to apply
// the filter to.
//
// The following code creates a filter that adds a zoom blur to the input
// image:
//
// [media-3544962]
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/zoomBlur()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (_CIFilterClass CIFilterClass) ZoomBlurFilter() CIFilter {
	rv := objc.Send[objc.ID](objc.ID(_CIFilterClass.class), objc.Sel("zoomBlurFilter"))
	return CIFilterFromID(rv)
}

// A name associated with a filter.
//
// # Discussion
//
// You use a filter’s name to construct key paths to its attributes when the
// filter is attached to a Core Animation layer. For example, if a [CALayer]
// object has an attached [CIFilter] instance whose name is
// `myExposureFilter`, you can refer to attributes of the filter using a key
// path such as `filters.MyExposureFilterXCUIElementTypeInputEV()`. Layer
// animations may also access filter attributes via these key paths.
//
// Core Animation can animate this property on a layer.
//
// The default value for this property is `nil`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/name
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
func (f CIFilter) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (f CIFilter) SetName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setName:"), objc.String(value))
}

// A Boolean value that determines whether the filter is enabled. Animatable.
//
// # Discussion
//
// The filter is applied to its input when this property is set to true (the
// default).
//
// Use this property in conjunction with the [Name] property when attaching
// filters to Core Animation layers and accessing or animating filter
// properties through key-value animations. Core Animation can animate this
// property on a layer.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/isEnabled
func (f CIFilter) Enabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEnabled"))
	return rv
}
func (f CIFilter) SetEnabled(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setEnabled:"), value)
}

// A dictionary of key-value pairs that describe the filter.
//
// # Discussion
//
// This dictionary contains two kinds of key-value pairs:
//
// - Keys listed in [Filter Attribute Keys] describe the filter, providing
// information such as a human-readable name and categories you can use to
// organize filters in your app’s UI. - Other keys, including those listed
// in [Filter Parameter Keys] and those starting with “input” or
// “output”, describe parameters that control the filter’s behavior. For
// each parameter key, the corresponding value is another dictionary
// describing possible values for that parameter, such as the value class and
// minimum/maximum values. You can use this information to build a UI to
// control the filter.
//
// For example, the attributes dictionary for the [CIColorControls] filter
// contains the following information:
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/attributes
//
// [Filter Attribute Keys]: https://developer.apple.com/documentation/CoreImage/filter-attribute-keys
// [Filter Parameter Keys]: https://developer.apple.com/documentation/CoreImage/filter-parameter-keys
func (f CIFilter) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The names of all input parameters to the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/inputKeys
func (f CIFilter) InputKeys() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("inputKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// The names of all output parameters from the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/outputKeys
func (f CIFilter) OutputKeys() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("outputKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns a [CIImage] object that encapsulates the operations configured in
// the filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class/outputImage
func (f CIFilter) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outputImage"))
	return CIImageFromID(objc.ID(rv))
}
