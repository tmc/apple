// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a factory to create a new format reader with a byte source.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderExtension
type MEFormatReaderExtension interface {
	objectivec.IObject

	// Creates a new format reader with the byte source and options that you specify.
	//
	// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderExtension/formatReader(with:options:)
	FormatReaderWithByteSourceOptionsError(primaryByteSource IMEByteSource, options IMEFormatReaderInstantiationOptions) (MEFormatReader, error)
}

// MEFormatReaderExtensionObject wraps an existing Objective-C object that conforms to the MEFormatReaderExtension protocol.
type MEFormatReaderExtensionObject struct {
	objectivec.Object
}

func (o MEFormatReaderExtensionObject) BaseObject() objectivec.Object {
	return o.Object
}

// MEFormatReaderExtensionObjectFromID constructs a [MEFormatReaderExtensionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MEFormatReaderExtensionObjectFromID(id objc.ID) MEFormatReaderExtensionObject {
	return MEFormatReaderExtensionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates a new format reader with the byte source and options that you
// specify.
//
// primaryByteSource: The primary byte source for the format reader.
//
// options: The reader instantiation options.
//
// # Return Value
//
// A new format reader.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderExtension/formatReader(with:options:)
func (o MEFormatReaderExtensionObject) FormatReaderWithByteSourceOptionsError(primaryByteSource IMEByteSource, options IMEFormatReaderInstantiationOptions) (MEFormatReader, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("formatReaderWithByteSource:options:error:"), primaryByteSource, options)
	if err != nil {
		return nil, err
	}
	return MEFormatReaderObjectFromID(rv), nil
}
