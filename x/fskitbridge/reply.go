package fskitbridge

import (
	"syscall"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

// ReplyBlocks invokes FSKit reply blocks.
type ReplyBlocks struct {
	blocks *objcbridge.BlockInvoker
}

// ReplyBlockShims names optional linked functions used for typed FSKit
// reply block invocation. Each field names a C function that invokes a
// block of the corresponding shape; an empty name means no shim is linked
// for that shape.
type ReplyBlockShims struct {
	Error    string
	Object   string
	ItemName string
	Verifier string
	Bool     string
	Size     string
}

// NewReplyBlocks returns a reply block invoker with no linked shim names.
func NewReplyBlocks() *ReplyBlocks {
	return NewReplyBlocksWithShims(ReplyBlockShims{})
}

// NewReplyBlocksWithShims returns a reply block invoker that prefers linked
// typed shims when they are present.
func NewReplyBlocksWithShims(shims ReplyBlockShims) *ReplyBlocks {
	return &ReplyBlocks{blocks: objcbridge.NewBlockInvokerWithShims(objcbridge.BlockShims{
		Error:             shims.Error,
		ObjectError:       shims.Object,
		ObjectObjectError: shims.ItemName,
		Uint64Error:       shims.Verifier,
		BoolError:         shims.Bool,
		UintptrError:      shims.Size,
	})}
}

// The reply methods return the error from invoking the block (for example a
// nil block or an unresolvable invoke pointer); a non-nil error means the
// FSKit caller did not receive its reply.

// Void invokes a reply block that takes no arguments.
func (r *ReplyBlocks) Void(block objc.ID) error {
	return r.blocks.Void(block)
}

// Error invokes a reply block that takes an NSError.
func (r *ReplyBlocks) Error(block objc.ID, err objc.ID) error {
	return r.blocks.Error(block, err)
}

// ObjectError invokes a reply block that takes an object and an NSError.
func (r *ReplyBlocks) ObjectError(block objc.ID, object objc.ID, err objc.ID) error {
	return r.blocks.ObjectError(block, object, err)
}

// ItemNameError invokes a reply block that takes an FSItem, an FSFileName,
// and an NSError.
func (r *ReplyBlocks) ItemNameError(block objc.ID, item objc.ID, name objc.ID, err objc.ID) error {
	return r.blocks.ObjectObjectError(block, item, name, err)
}

// VerifierError invokes a reply block that takes a directory verifier and
// an NSError.
func (r *ReplyBlocks) VerifierError(block objc.ID, verifier uint64, err objc.ID) error {
	return r.blocks.Uint64Error(block, verifier, err)
}

// BoolError invokes a reply block that takes a boolean and an NSError.
func (r *ReplyBlocks) BoolError(block objc.ID, value bool, err objc.ID) error {
	return r.blocks.BoolError(block, value, err)
}

// SizeError invokes a reply block that takes a byte count and an NSError.
func (r *ReplyBlocks) SizeError(block objc.ID, size uintptr, err objc.ID) error {
	return r.blocks.UintptrError(block, size, err)
}

// POSIXError returns a new NSError in the POSIX error domain for errno,
// suitable for passing to a reply block.
func POSIXError(errno syscall.Errno) objc.ID {
	return foundation.NewErrorWithDomainCodeUserInfo(foundation.POSIXErrorDomain, int(errno), nil).GetID()
}
