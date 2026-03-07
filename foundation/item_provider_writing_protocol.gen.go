// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for implementing a class to allow an item provider to retrieve data from an instance of the class.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting
type NSItemProviderWriting interface {
	objectivec.IObject

	// Loads data of a particular type, identified by the given UTI.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/loadData(withTypeIdentifier:forItemProviderCompletionHandler:)
	LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress

	// An array of UTI strings representing the types of data that can be loaded for an item provider.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.property
	WritableTypeIdentifiersForItemProvider() []string
}



// NSItemProviderWritingObject wraps an existing Objective-C object that conforms to the NSItemProviderWriting protocol.
type NSItemProviderWritingObject struct {
	objectivec.Object
}
func (o NSItemProviderWritingObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSItemProviderWritingObjectFromID constructs a [NSItemProviderWritingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSItemProviderWritingObjectFromID(id objc.ID) NSItemProviderWritingObject {
	return NSItemProviderWritingObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Loads data of a particular type, identified by the given UTI.
//
// typeIdentifier: The uniform type identifier (UTI) identifying the type of data to load.
//
// completionHandler: The handler that’s called after the data is loaded.
//
// # Return Value
// 
// The progress of the data load process.
//
// # Discussion
// 
// When the system calls this method, the `typeIdentifier` parameter is set to
// one of the elements in the `writableTypeIdentifiersForItemProvider` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/loadData(withTypeIdentifier:forItemProviderCompletionHandler:)

func (o NSItemProviderWritingObject) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), objc.String(typeIdentifier), completionHandler)
	return NSProgressFromID(rv)
	}

// An array of UTI strings representing the types of data that can be loaded
// for an item provider.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.property

func (o NSItemProviderWritingObject) WritableTypeIdentifiersForItemProvider() []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return objc.ConvertSliceToStrings(rv)
	}

// Asks the item provider for the representation visibility specification for
// the given UTI.
//
// typeIdentifier: A uniform type identifier (UTI).
//
// # Return Value
// 
// A representation visibility specification for the given UTI.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/itemProviderVisibilityForRepresentation(withTypeIdentifier:)-swift.method

func (o NSItemProviderWritingObject) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) NSItemProviderRepresentationVisibility {
	
	rv := objc.Send[NSItemProviderRepresentationVisibility](o.ID, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), objc.String(typeIdentifier))
	return rv
	}









