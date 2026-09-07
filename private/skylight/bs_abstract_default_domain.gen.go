// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BSAbstractDefaultDomain] class.
var (
	_BSAbstractDefaultDomainClass     BSAbstractDefaultDomainClass
	_BSAbstractDefaultDomainClassOnce sync.Once
)

func getBSAbstractDefaultDomainClass() BSAbstractDefaultDomainClass {
	_BSAbstractDefaultDomainClassOnce.Do(func() {
		_BSAbstractDefaultDomainClass = BSAbstractDefaultDomainClass{class: objc.GetClass("BSAbstractDefaultDomain")}
	})
	return _BSAbstractDefaultDomainClass
}

// GetBSAbstractDefaultDomainClass returns the class object for BSAbstractDefaultDomain.
func GetBSAbstractDefaultDomainClass() BSAbstractDefaultDomainClass {
	return getBSAbstractDefaultDomainClass()
}

type BSAbstractDefaultDomainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BSAbstractDefaultDomainClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BSAbstractDefaultDomainClass) Alloc() BSAbstractDefaultDomain {
	rv := objc.SendIfResponds[BSAbstractDefaultDomain](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other skylight classes. [Full Topic]
type BSAbstractDefaultDomain struct {
	objectivec.Object
}

// BSAbstractDefaultDomainFromID constructs a [BSAbstractDefaultDomain] from an objc.ID.
//
// A parent class referenced by other skylight classes.
func BSAbstractDefaultDomainFromID(id objc.ID) BSAbstractDefaultDomain {
	return BSAbstractDefaultDomain{objectivec.Object{ID: id}}
}

// Ensure BSAbstractDefaultDomain implements IBSAbstractDefaultDomain.
var _ IBSAbstractDefaultDomain = BSAbstractDefaultDomain{}

// An interface definition for the [BSAbstractDefaultDomain] class.
type IBSAbstractDefaultDomain interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BSAbstractDefaultDomain) Init() BSAbstractDefaultDomain {
	rv := objc.SendIfResponds[BSAbstractDefaultDomain](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BSAbstractDefaultDomain) Autorelease() BSAbstractDefaultDomain {
	rv := objc.SendIfResponds[BSAbstractDefaultDomain](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBSAbstractDefaultDomain creates a new BSAbstractDefaultDomain instance.
func NewBSAbstractDefaultDomain() BSAbstractDefaultDomain {
	class := getBSAbstractDefaultDomainClass()
	rv := objc.SendIfResponds[BSAbstractDefaultDomain](objc.ID(class.class), objc.Sel("new"))
	return rv
}
