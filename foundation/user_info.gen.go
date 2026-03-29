// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UserInfo] class.
var (
	_UserInfoClass     UserInfoClass
	_UserInfoClassOnce sync.Once
)

func getUserInfoClass() UserInfoClass {
	_UserInfoClassOnce.Do(func() {
		_UserInfoClass = UserInfoClass{class: objc.GetClass("userInfo")}
	})
	return _UserInfoClass
}

// GetUserInfoClass returns the class object for userInfo.
func GetUserInfoClass() UserInfoClass {
	return getUserInfoClass()
}

type UserInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UserInfoClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UserInfoClass) Alloc() UserInfo {
	rv := objc.Send[UserInfo](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSException/userInfo-c.ivar
type UserInfo struct {
	objectivec.Object
}

// UserInfoFromID constructs a [UserInfo] from an objc.ID.
func UserInfoFromID(id objc.ID) UserInfo {
	return UserInfo{objectivec.Object{ID: id}}
}
// Ensure UserInfo implements IUserInfo.
var _ IUserInfo = UserInfo{}

// An interface definition for the [UserInfo] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/userInfo-c.ivar
type IUserInfo interface {
	objectivec.IObject
}

// Init initializes the instance.
func (u UserInfo) Init() UserInfo {
	rv := objc.Send[UserInfo](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UserInfo) Autorelease() UserInfo {
	rv := objc.Send[UserInfo](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserInfo creates a new UserInfo instance.
func NewUserInfo() UserInfo {
	class := getUserInfoClass()
	rv := objc.Send[UserInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

