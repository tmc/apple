// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"fmt"
)

type ObjC uint

const (
	OBJC_COLLECT_IF_NEEDED ObjC = 8
	OBJC_EXHAUSTIVE_COLLECTION ObjC = 3
	OBJC_FULL_COLLECTION ObjC = 2
	OBJC_GENERATIONAL_COLLECTION ObjC = 1
	OBJC_RATIO_COLLECTION ObjC = 0
	OBJC_WAIT_UNTIL_DONE ObjC = 16
)


func (e ObjC) String() string {
	switch e {
	case OBJC_COLLECT_IF_NEEDED:
		return "OBJC_COLLECT_IF_NEEDED"
	case OBJC_EXHAUSTIVE_COLLECTION:
		return "OBJC_EXHAUSTIVE_COLLECTION"
	case OBJC_FULL_COLLECTION:
		return "OBJC_FULL_COLLECTION"
	case OBJC_GENERATIONAL_COLLECTION:
		return "OBJC_GENERATIONAL_COLLECTION"
	case OBJC_RATIO_COLLECTION:
		return "OBJC_RATIO_COLLECTION"
	case OBJC_WAIT_UNTIL_DONE:
		return "OBJC_WAIT_UNTIL_DONE"
	default:
		return fmt.Sprintf("ObjC(%d)", e)
	}
}




type ObjCClearResidentStack uint

const (
	OBJC_CLEAR_RESIDENT_STACK ObjCClearResidentStack = 1
)


func (e ObjCClearResidentStack) String() string {
	switch e {
	case OBJC_CLEAR_RESIDENT_STACK:
		return "OBJC_CLEAR_RESIDENT_STACK"
	default:
		return fmt.Sprintf("ObjCClearResidentStack(%d)", e)
	}
}




type ObjCSync uint

const (
	OBJC_SYNC_NOT_OWNING_THREAD_ERROR ObjCSync = 0
	OBJC_SYNC_SUCCESS ObjCSync = 0
)


func (e ObjCSync) String() string {
	switch e {
	case OBJC_SYNC_NOT_OWNING_THREAD_ERROR:
		return "OBJC_SYNC_NOT_OWNING_THREAD_ERROR"
	default:
		return fmt.Sprintf("ObjCSync(%d)", e)
	}
}



// See: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy
type Objc_AssociationPolicy int

const (
	// OBJC_ASSOCIATION_ASSIGN: Specifies an unsafe unretained reference to the associated object.
	OBJC_ASSOCIATION_ASSIGN Objc_AssociationPolicy = 0
	// OBJC_ASSOCIATION_COPY: Specifies that the associated object is copied, and that the association is made atomically.
	OBJC_ASSOCIATION_COPY Objc_AssociationPolicy = 771
	// OBJC_ASSOCIATION_COPY_NONATOMIC: Specifies that the associated object is copied, and that the association is not made atomically.
	OBJC_ASSOCIATION_COPY_NONATOMIC Objc_AssociationPolicy = 3
	// OBJC_ASSOCIATION_RETAIN: Specifies a strong reference to the associated object, and that the association is made atomically.
	OBJC_ASSOCIATION_RETAIN Objc_AssociationPolicy = 769
	// OBJC_ASSOCIATION_RETAIN_NONATOMIC: Specifies a strong reference to the associated object, and that the association is not made atomically.
	OBJC_ASSOCIATION_RETAIN_NONATOMIC Objc_AssociationPolicy = 1
)


func (e Objc_AssociationPolicy) String() string {
	switch e {
	case OBJC_ASSOCIATION_ASSIGN:
		return "OBJC_ASSOCIATION_ASSIGN"
	case OBJC_ASSOCIATION_COPY:
		return "OBJC_ASSOCIATION_COPY"
	case OBJC_ASSOCIATION_COPY_NONATOMIC:
		return "OBJC_ASSOCIATION_COPY_NONATOMIC"
	case OBJC_ASSOCIATION_RETAIN:
		return "OBJC_ASSOCIATION_RETAIN"
	case OBJC_ASSOCIATION_RETAIN_NONATOMIC:
		return "OBJC_ASSOCIATION_RETAIN_NONATOMIC"
	default:
		return fmt.Sprintf("Objc_AssociationPolicy(%d)", e)
	}
}





