// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol used to provide information about the layout’s container and environment traits, such as size classes and display scale factor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEnvironment
type NSCollectionLayoutEnvironment interface {
	objectivec.IObject

	// Information about the layout’s container, such as its size and content insets.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEnvironment/container
	Container() NSCollectionLayoutContainer
}



// NSCollectionLayoutEnvironmentObject wraps an existing Objective-C object that conforms to the NSCollectionLayoutEnvironment protocol.
type NSCollectionLayoutEnvironmentObject struct {
	objectivec.Object
}
func (o NSCollectionLayoutEnvironmentObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSCollectionLayoutEnvironmentObjectFromID constructs a [NSCollectionLayoutEnvironmentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionLayoutEnvironmentObjectFromID(id objc.ID) NSCollectionLayoutEnvironmentObject {
	return NSCollectionLayoutEnvironmentObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Information about the layout’s container, such as its size and content
// insets.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEnvironment/container

func (o NSCollectionLayoutEnvironmentObject) Container() NSCollectionLayoutContainer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("container"))
	return NSCollectionLayoutContainerObjectFromID(rv)
	}









