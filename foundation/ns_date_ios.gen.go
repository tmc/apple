// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/corefoundation"
)

//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(SRAbsoluteTime:)-886t8
func (d NSDate) InitWithSRAbsoluteTime(time corefoundation.CFTimeInterval) NSDate {
rv := objc.Send[NSDate](d.ID, objc.Sel("initWithSRAbsoluteTime:"), time)
return rv
}



//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(SRAbsoluteTime:)-9wpl1
func (_NSDateClass NSDateClass) DateWithSRAbsoluteTime(time corefoundation.CFTimeInterval) NSDate {
rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("dateWithSRAbsoluteTime:"), time)
return NSDateFromID(rv)
}






// See: https://developer.apple.com/documentation/Foundation/NSDate/srAbsoluteTime
func (d NSDate) SrAbsoluteTime() corefoundation.CFTimeInterval {
rv := objc.Send[corefoundation.CFTimeInterval](d.ID, objc.Sel("srAbsoluteTime"))
		return corefoundation.CFTimeInterval(rv)
}








