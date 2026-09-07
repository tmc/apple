// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Subscription] class.
var (
	_SubscriptionClass     SubscriptionClass
	_SubscriptionClassOnce sync.Once
)

func getSubscriptionClass() SubscriptionClass {
	_SubscriptionClassOnce.Do(func() {
		_SubscriptionClass = SubscriptionClass{class: objc.GetClass("_TtCV8SkyLight10WSNotifier12Subscription")}
	})
	return _SubscriptionClass
}

// GetSubscriptionClass returns the class object for _TtCV8SkyLight10WSNotifier12Subscription.
func GetSubscriptionClass() SubscriptionClass {
	return getSubscriptionClass()
}

type SubscriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SubscriptionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SubscriptionClass) Alloc() Subscription {
	rv := objc.SendIfResponds[Subscription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type Subscription struct {
	objectivec.Object
}

// SubscriptionFromID constructs a [Subscription] from an objc.ID.
func SubscriptionFromID(id objc.ID) Subscription {
	return Subscription{objectivec.Object{ID: id}}
}

// Ensure Subscription implements ISubscription.
var _ ISubscription = Subscription{}

// An interface definition for the [Subscription] class.
type ISubscription interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s Subscription) Init() Subscription {
	rv := objc.SendIfResponds[Subscription](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Subscription) Autorelease() Subscription {
	rv := objc.SendIfResponds[Subscription](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubscription creates a new Subscription instance.
func NewSubscription() Subscription {
	class := getSubscriptionClass()
	rv := objc.SendIfResponds[Subscription](objc.ID(class.class), objc.Sel("new"))
	return rv
}
