// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BKIOHIDService] class.
var (
	_BKIOHIDServiceClass     BKIOHIDServiceClass
	_BKIOHIDServiceClassOnce sync.Once
)

func getBKIOHIDServiceClass() BKIOHIDServiceClass {
	_BKIOHIDServiceClassOnce.Do(func() {
		_BKIOHIDServiceClass = BKIOHIDServiceClass{class: objc.GetClass("BKIOHIDService")}
	})
	return _BKIOHIDServiceClass
}

// GetBKIOHIDServiceClass returns the class object for BKIOHIDService.
func GetBKIOHIDServiceClass() BKIOHIDServiceClass {
	return getBKIOHIDServiceClass()
}

type BKIOHIDServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BKIOHIDServiceClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BKIOHIDServiceClass) Alloc() BKIOHIDService {
	rv := objc.SendIfResponds[BKIOHIDService](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other skylight classes. [Full Topic]
type BKIOHIDService struct {
	objectivec.Object
}

// BKIOHIDServiceFromID constructs a [BKIOHIDService] from an objc.ID.
//
// A parent class referenced by other skylight classes.
func BKIOHIDServiceFromID(id objc.ID) BKIOHIDService {
	return BKIOHIDService{objectivec.Object{ID: id}}
}

// Ensure BKIOHIDService implements IBKIOHIDService.
var _ IBKIOHIDService = BKIOHIDService{}

// An interface definition for the [BKIOHIDService] class.
type IBKIOHIDService interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BKIOHIDService) Init() BKIOHIDService {
	rv := objc.SendIfResponds[BKIOHIDService](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BKIOHIDService) Autorelease() BKIOHIDService {
	rv := objc.SendIfResponds[BKIOHIDService](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBKIOHIDService creates a new BKIOHIDService instance.
func NewBKIOHIDService() BKIOHIDService {
	class := getBKIOHIDServiceClass()
	rv := objc.SendIfResponds[BKIOHIDService](objc.ID(class.class), objc.Sel("new"))
	return rv
}
