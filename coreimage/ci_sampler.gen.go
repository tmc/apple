// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CISampler] class.
var (
	_CISamplerClass     CISamplerClass
	_CISamplerClassOnce sync.Once
)

func getCISamplerClass() CISamplerClass {
	_CISamplerClassOnce.Do(func() {
		_CISamplerClass = CISamplerClass{class: objc.GetClass("CISampler")}
	})
	return _CISamplerClass
}

// GetCISamplerClass returns the class object for CISampler.
func GetCISamplerClass() CISamplerClass {
	return getCISamplerClass()
}

type CISamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CISamplerClass) Alloc() CISampler {
	rv := objc.Send[CISampler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that retrieves pixel samples for processing by a filter kernel.
//
// # Overview
// 
// The [CISampler] class retrieves samples of images for processing by a
// [CIKernel] object. A [CISampler] object defines a coordinate transform, and
// modes for interpolation and wrapping. You use [CISampler] objects in
// conjunction with other Core Image classes, such as [CIFilter], [CIKernel],
// and [CIFilterShape], to create custom filters.
//
// # Initializing a Sampler
//
//   - [CISampler.InitWithImage]: Initializes a sampler with an image object.
//   - [CISampler.InitWithImageOptions]: Initializes the sampler with an image object using options specified in a dictionary.
//
// # Getting Information About the Sampler Object
//
//   - [CISampler.Definition]: The domain of definition (DOD) of the sampler
//   - [CISampler.Extent]: The rectangle that specifies the extent of the sampler
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler
type CISampler struct {
	objectivec.Object
}

// CISamplerFromID constructs a [CISampler] from an objc.ID.
//
// An object that retrieves pixel samples for processing by a filter kernel.
func CISamplerFromID(id objc.ID) CISampler {
	return CISampler{objectivec.Object{ID: id}}
}
// NOTE: CISampler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CISampler] class.
//
// # Initializing a Sampler
//
//   - [ICISampler.InitWithImage]: Initializes a sampler with an image object.
//   - [ICISampler.InitWithImageOptions]: Initializes the sampler with an image object using options specified in a dictionary.
//
// # Getting Information About the Sampler Object
//
//   - [ICISampler.Definition]: The domain of definition (DOD) of the sampler
//   - [ICISampler.Extent]: The rectangle that specifies the extent of the sampler
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler
type ICISampler interface {
	objectivec.IObject

	// Topic: Initializing a Sampler

	// Initializes a sampler with an image object.
	InitWithImage(im ICIImage) CISampler
	// Initializes the sampler with an image object using options specified in a dictionary.
	InitWithImageOptions(im ICIImage, dict foundation.INSDictionary) CISampler

	// Topic: Getting Information About the Sampler Object

	// The domain of definition (DOD) of the sampler
	Definition() ICIFilterShape
	// The rectangle that specifies the extent of the sampler
	Extent() corefoundation.CGRect

	// Initializes the sampler with an image object using options specified as key-value pairs.
	InitWithImageKeysAndValues(im ICIImage, key0 objectivec.IObject) CISampler
}

// Init initializes the instance.
func (s CISampler) Init() CISampler {
	rv := objc.Send[CISampler](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s CISampler) Autorelease() CISampler {
	rv := objc.Send[CISampler](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewCISampler creates a new CISampler instance.
func NewCISampler() CISampler {
	class := getCISamplerClass()
	rv := objc.Send[CISampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a sampler with an image object.
//
// im: The image object to initialize the sampler with.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:)
func NewSamplerWithImage(im ICIImage) CISampler {
	instance := getCISamplerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:"), im)
	return CISamplerFromID(rv)
}

// Initializes the sampler with an image object using options specified as
// key-value pairs.
//
// im: The image object to initialize the sampler with.
//
// key0: A list of key-value pairs that represent options. Each key needs to be
// followed by that appropriate value. You can supply one or more key-value
// pairs. Use `nil` to specify the end of the key-value options. See [Sampler
// Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/initWithImage:keysAndValues:
func NewSamplerWithImageKeysAndValues(im ICIImage, key0 objectivec.IObject) CISampler {
	instance := getCISamplerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:keysAndValues:"), im, key0)
	return CISamplerFromID(rv)
}

// Initializes the sampler with an image object using options specified in a
// dictionary.
//
// im: The image to initialize the sampler with.
//
// dict: A dictionary that contains options specified as key-value pairs. See
// [Sampler Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:options:)
func NewSamplerWithImageOptions(im ICIImage, dict foundation.INSDictionary) CISampler {
	instance := getCISamplerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:options:"), im, dict)
	return CISamplerFromID(rv)
}

// Initializes a sampler with an image object.
//
// im: The image object to initialize the sampler with.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:)
func (s CISampler) InitWithImage(im ICIImage) CISampler {
	rv := objc.Send[CISampler](s.ID, objc.Sel("initWithImage:"), im)
	return rv
}
// Initializes the sampler with an image object using options specified in a
// dictionary.
//
// im: The image to initialize the sampler with.
//
// dict: A dictionary that contains options specified as key-value pairs. See
// [Sampler Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/init(image:options:)
func (s CISampler) InitWithImageOptions(im ICIImage, dict foundation.INSDictionary) CISampler {
	rv := objc.Send[CISampler](s.ID, objc.Sel("initWithImage:options:"), im, dict)
	return rv
}
// Initializes the sampler with an image object using options specified as
// key-value pairs.
//
// im: The image object to initialize the sampler with.
//
// key0: A list of key-value pairs that represent options. Each key needs to be
// followed by that appropriate value. You can supply one or more key-value
// pairs. Use `nil` to specify the end of the key-value options. See [Sampler
// Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/initWithImage:keysAndValues:
func (s CISampler) InitWithImageKeysAndValues(im ICIImage, key0 objectivec.IObject) CISampler {
	rv := objc.Send[CISampler](s.ID, objc.Sel("initWithImage:keysAndValues:"), im, key0)
	return rv
}

// Creates and returns a sampler that references an image.
//
// im: The image that you want the sampler to reference.
//
// # Return Value
// 
// A sampler object that references the image specified by the `im` argument.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:
func (_CISamplerClass CISamplerClass) SamplerWithImage(im ICIImage) CISampler {
	rv := objc.Send[objc.ID](objc.ID(_CISamplerClass.class), objc.Sel("samplerWithImage:"), im)
	return CISamplerFromID(rv)
}
// Creates and returns a sampler that references an image using options
// specified as key-value pairs.
//
// im: The image that you want the sampler to reference.
//
// key0: A list of key-value pairs that represent options. Each key needs to be
// followed by that appropriate value. You can supply one or more key-value
// pairs. Use `nil` to specify the end of the key-value options. See [Sampler
// Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// # Return Value
// 
// A sampler that references the image specified by the `im` argument and uses
// the specified options.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:keysAndValues:
func (_CISamplerClass CISamplerClass) SamplerWithImageKeysAndValues(im ICIImage, key0 objectivec.IObject) CISampler {
	rv := objc.Send[objc.ID](objc.ID(_CISamplerClass.class), objc.Sel("samplerWithImage:keysAndValues:"), im, key0)
	return CISamplerFromID(rv)
}
// Creates and returns a sampler that references an image using options
// specified in a dictionary.
//
// im: The image that you want the sampler to reference.
//
// dict: A dictionary that contains options specified as key-value pairs. See
// [Sampler Option Keys].
// //
// [Sampler Option Keys]: https://developer.apple.com/documentation/CoreImage/sampler-option-keys
//
// # Return Value
// 
// A sampler that references the image specified by the `im` argument and uses
// the options specified in the dictionary.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/samplerWithImage:options:
func (_CISamplerClass CISamplerClass) SamplerWithImageOptions(im ICIImage, dict foundation.INSDictionary) CISampler {
	rv := objc.Send[objc.ID](objc.ID(_CISamplerClass.class), objc.Sel("samplerWithImage:options:"), im, dict)
	return CISamplerFromID(rv)
}

// The domain of definition (DOD) of the sampler
//
// # Discussion
// 
// The DOD contains all nontransparent pixels produced by referencing the
// sampler.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/definition
func (s CISampler) Definition() ICIFilterShape {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("definition"))
	return CIFilterShapeFromID(objc.ID(rv))
}
// The rectangle that specifies the extent of the sampler
//
// # Discussion
// 
// Extent is the rectangle that specifies the area outside which the wrap mode
// set for the sampler is invoked.
//
// See: https://developer.apple.com/documentation/CoreImage/CISampler/extent
func (s CISampler) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("extent"))
	return corefoundation.CGRect(rv)
}

