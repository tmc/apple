// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
)

// The preferred style for presenting the item provider’s data.
//
// # Discussion
// 
// The default preferred presentation style is `unspecified`.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationStyle-swift.property
func (i NSItemProvider) PreferredPresentationStyle() UIPreferredPresentationStyle {
rv := objc.Send[UIPreferredPresentationStyle](i.ID, objc.Sel("preferredPresentationStyle"))
		return UIPreferredPresentationStyle(rv)
}
func (i NSItemProvider) SetPreferredPresentationStyle(value UIPreferredPresentationStyle) {
objc.Send[struct{}](i.ID, objc.Sel("setPreferredPresentationStyle:"), value)
}



// The collection of data an app uses to hold private team information during
// drag and drop.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/teamData
func (i NSItemProvider) TeamData() INSData {
rv := objc.Send[objc.ID](i.ID, objc.Sel("teamData"))
return NSDataFromID(rv)
}
func (i NSItemProvider) SetTeamData(value INSData) {
objc.Send[struct{}](i.ID, objc.Sel("setTeamData:"), value)
}








