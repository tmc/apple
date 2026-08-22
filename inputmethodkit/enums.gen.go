// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"fmt"
)

type KIMKLocateCandidates uint32

const (
	// KIMKLocateCandidatesAboveHint: # Discussion
	KIMKLocateCandidatesAboveHint KIMKLocateCandidates = 1
	// KIMKLocateCandidatesBelowHint: # Discussion
	KIMKLocateCandidatesBelowHint KIMKLocateCandidates = 2
	// KIMKLocateCandidatesLeftHint: # Discussion
	KIMKLocateCandidatesLeftHint KIMKLocateCandidates = 3
	// KIMKLocateCandidatesRightHint: # Discussion
	KIMKLocateCandidatesRightHint KIMKLocateCandidates = 4
)

func (e KIMKLocateCandidates) String() string {
	switch e {
	case KIMKLocateCandidatesAboveHint:
		return "KIMKLocateCandidatesAboveHint"
	case KIMKLocateCandidatesBelowHint:
		return "KIMKLocateCandidatesBelowHint"
	case KIMKLocateCandidatesLeftHint:
		return "KIMKLocateCandidatesLeftHint"
	case KIMKLocateCandidatesRightHint:
		return "KIMKLocateCandidatesRightHint"
	default:
		return fmt.Sprintf("KIMKLocateCandidates(%d)", e)
	}
}

type KIMKMain uint32

const (
	KIMKAnnotation KIMKMain = 1
	KIMKMainValue  KIMKMain = 0
	KIMKSubList    KIMKMain = 2
)

func (e KIMKMain) String() string {
	switch e {
	case KIMKAnnotation:
		return "KIMKAnnotation"
	case KIMKMainValue:
		return "KIMKMainValue"
	case KIMKSubList:
		return "KIMKSubList"
	default:
		return fmt.Sprintf("KIMKMain(%d)", e)
	}
}

type KIMKSingleColumnScrollingCandidatePanel uint32

const (
	// KIMKScrollingGridCandidatePanel: # Discussion
	KIMKScrollingGridCandidatePanel KIMKSingleColumnScrollingCandidatePanel = 2
	// KIMKSingleColumnScrollingCandidatePanelValue: # Discussion
	KIMKSingleColumnScrollingCandidatePanelValue KIMKSingleColumnScrollingCandidatePanel = 1
	// KIMKSingleRowSteppingCandidatePanel: # Discussion
	KIMKSingleRowSteppingCandidatePanel KIMKSingleColumnScrollingCandidatePanel = 3
)

func (e KIMKSingleColumnScrollingCandidatePanel) String() string {
	switch e {
	case KIMKScrollingGridCandidatePanel:
		return "KIMKScrollingGridCandidatePanel"
	case KIMKSingleColumnScrollingCandidatePanelValue:
		return "KIMKSingleColumnScrollingCandidatePanelValue"
	case KIMKSingleRowSteppingCandidatePanel:
		return "KIMKSingleRowSteppingCandidatePanel"
	default:
		return fmt.Sprintf("KIMKSingleColumnScrollingCandidatePanel(%d)", e)
	}
}
