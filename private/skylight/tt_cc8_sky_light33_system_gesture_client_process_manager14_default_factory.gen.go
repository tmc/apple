// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DefaultFactory] class.
var (
	_DefaultFactoryClass     DefaultFactoryClass
	_DefaultFactoryClassOnce sync.Once
)

func getDefaultFactoryClass() DefaultFactoryClass {
	_DefaultFactoryClassOnce.Do(func() {
		_DefaultFactoryClass = DefaultFactoryClass{class: objc.GetClass("_TtCC8SkyLight33SystemGestureClientProcessManager14DefaultFactory")}
	})
	return _DefaultFactoryClass
}

// GetDefaultFactoryClass returns the class object for _TtCC8SkyLight33SystemGestureClientProcessManager14DefaultFactory.
func GetDefaultFactoryClass() DefaultFactoryClass {
	return getDefaultFactoryClass()
}

type DefaultFactoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DefaultFactoryClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DefaultFactoryClass) Alloc() DefaultFactory {
	rv := objc.SendIfResponds[DefaultFactory](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DefaultFactory struct {
	objectivec.Object
}

// DefaultFactoryFromID constructs a [DefaultFactory] from an objc.ID.
func DefaultFactoryFromID(id objc.ID) DefaultFactory {
	return DefaultFactory{objectivec.Object{ID: id}}
}

// Ensure DefaultFactory implements IDefaultFactory.
var _ IDefaultFactory = DefaultFactory{}

// An interface definition for the [DefaultFactory] class.
type IDefaultFactory interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DefaultFactory) Init() DefaultFactory {
	rv := objc.SendIfResponds[DefaultFactory](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DefaultFactory) Autorelease() DefaultFactory {
	rv := objc.SendIfResponds[DefaultFactory](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDefaultFactory creates a new DefaultFactory instance.
func NewDefaultFactory() DefaultFactory {
	class := getDefaultFactoryClass()
	rv := objc.SendIfResponds[DefaultFactory](objc.ID(class.class), objc.Sel("new"))
	return rv
}
