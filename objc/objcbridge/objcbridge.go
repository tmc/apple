// Package objcbridge contains small Objective-C runtime helpers.
package objcbridge

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

// AddMethods installs methods on cls.
//
// Each method's Fn must take (objc.ID, objc.SEL) as its first two arguments;
// objc.NewIMP panics otherwise. See its documentation for why a missing prefix
// corrupts the call rather than failing it.
func AddMethods(cls objc.Class, className string, methods []objc.MethodDef) error {
	for _, method := range methods {
		if !objc.AddMethod(cls, method.Cmd, objc.NewIMP(method.Fn), "") {
			return fmt.Errorf("add method %v to %s", method.Cmd, className)
		}
	}
	return nil
}

// ProtocolsByName resolves Objective-C protocols, skipping protocols not
// present on the running system.
func ProtocolsByName(names ...string) []*objc.Protocol {
	var protocols []*objc.Protocol
	for _, name := range names {
		if protocol := objc.GetProtocol(name); protocol != nil {
			protocols = append(protocols, protocol)
		}
	}
	return protocols
}

// RequiredProtocolsByName resolves Objective-C protocols and reports any
// missing protocol.
func RequiredProtocolsByName(names ...string) ([]*objc.Protocol, error) {
	protocols := make([]*objc.Protocol, 0, len(names))
	for _, name := range names {
		protocol := objc.GetProtocol(name)
		if protocol == nil {
			return nil, fmt.Errorf("lookup protocol %s", name)
		}
		protocols = append(protocols, protocol)
	}
	return protocols, nil
}

type blockLayout struct {
	isa        uintptr
	flags      uint32
	reserved   uint32
	invoke     uintptr
	descriptor uintptr
}

var (
	memcpyOnce sync.Once
	memcpyAddr uintptr
	memcpyErr  error
)

func memcpy() (uintptr, error) {
	memcpyOnce.Do(func() {
		memcpyAddr, memcpyErr = purego.Dlsym(purego.RTLD_DEFAULT, "memcpy")
		if memcpyErr != nil {
			memcpyErr = fmt.Errorf("resolve memcpy: %w", memcpyErr)
		}
	})
	return memcpyAddr, memcpyErr
}

// BlockInvoker invokes Objective-C blocks with known signatures.
//
// If a typed shim is linked into the process, BlockInvoker uses that bridge;
// otherwise it falls back to the block's invoke pointer. Its methods return
// an error when block is nil or when the fallback invoke pointer cannot be
// resolved.
type BlockInvoker struct {
	once              sync.Once
	error             func(objc.ID, objc.ID)
	object            func(objc.ID, objc.ID)
	objectObject      func(objc.ID, objc.ID, objc.ID)
	uint              func(objc.ID, uint)
	objectError       func(objc.ID, objc.ID, objc.ID)
	objectObjectError func(objc.ID, objc.ID, objc.ID, objc.ID)
	uint64Error       func(objc.ID, uint64, objc.ID)
	boolError         func(objc.ID, bool, objc.ID)
	uintptrError      func(objc.ID, uintptr, objc.ID)
	shims             BlockShims
}

// BlockShims names optional linked functions used for typed block invocation.
type BlockShims struct {
	Error             string
	Object            string
	ObjectObject      string
	Uint              string
	ObjectError       string
	ObjectObjectError string
	Uint64Error       string
	BoolError         string
	UintptrError      string
}

// NewBlockInvoker returns a block invoker with no linked shim names.
func NewBlockInvoker() *BlockInvoker {
	return NewBlockInvokerWithShims(BlockShims{})
}

// NewBlockInvokerWithShims returns a block invoker that prefers linked typed
// shims when they are present.
func NewBlockInvokerWithShims(shims BlockShims) *BlockInvoker {
	return &BlockInvoker{shims: shims}
}

func (r *BlockInvoker) loadShims() {
	r.once.Do(func() {
		registerOptionalFunc(&r.error, r.shims.Error)
		registerOptionalFunc(&r.object, r.shims.Object)
		registerOptionalFunc(&r.objectObject, r.shims.ObjectObject)
		registerOptionalFunc(&r.uint, r.shims.Uint)
		registerOptionalFunc(&r.objectError, r.shims.ObjectError)
		registerOptionalFunc(&r.objectObjectError, r.shims.ObjectObjectError)
		registerOptionalFunc(&r.uint64Error, r.shims.Uint64Error)
		registerOptionalFunc(&r.boolError, r.shims.BoolError)
		registerOptionalFunc(&r.uintptrError, r.shims.UintptrError)
	})
}

func registerOptionalFunc(fptr any, name string) {
	if name == "" {
		return
	}
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
	if err != nil {
		return
	}
	purego.RegisterFunc(fptr, sym)
}

func blockInvoke(block objc.ID) (uintptr, error) {
	if block == 0 {
		return 0, errors.New("nil block")
	}
	copyMemory, err := memcpy()
	if err != nil {
		return 0, err
	}
	var invoke uintptr
	purego.SyscallN(
		copyMemory,
		uintptr(unsafe.Pointer(&invoke)),
		uintptr(block)+unsafe.Offsetof(blockLayout{}.invoke),
		unsafe.Sizeof(invoke),
	)
	if invoke == 0 {
		return 0, errors.New("block has nil invoke")
	}
	return invoke, nil
}

func checkBlock(block objc.ID) error {
	if block == 0 {
		return errors.New("nil block")
	}
	return nil
}

// Void invokes block().
func (r *BlockInvoker) Void(block objc.ID) error {
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block)
	return nil
}

// Object invokes block(object).
func (r *BlockInvoker) Object(block objc.ID, object objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.object != nil {
		r.object(block, object)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, object)
	return nil
}

// ObjectObject invokes block(a, b).
func (r *BlockInvoker) ObjectObject(block objc.ID, a objc.ID, b objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.objectObject != nil {
		r.objectObject(block, a, b)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, objc.ID, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, a, b)
	return nil
}

// Uint invokes block(value).
func (r *BlockInvoker) Uint(block objc.ID, value uint) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.uint != nil {
		r.uint(block, value)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, uint)
	purego.RegisterFunc(&fn, invoke)
	fn(block, value)
	return nil
}

// Error invokes block(err).
func (r *BlockInvoker) Error(block objc.ID, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.error != nil {
		r.error(block, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, errID)
	return nil
}

// ObjectError invokes block(object, err).
func (r *BlockInvoker) ObjectError(block objc.ID, object objc.ID, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.objectError != nil {
		r.objectError(block, object, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, objc.ID, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, object, errID)
	return nil
}

// ObjectObjectError invokes block(a, b, err).
func (r *BlockInvoker) ObjectObjectError(block objc.ID, a objc.ID, b objc.ID, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.objectObjectError != nil {
		r.objectObjectError(block, a, b, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, objc.ID, objc.ID, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, a, b, errID)
	return nil
}

// Uint64Error invokes block(value, err).
func (r *BlockInvoker) Uint64Error(block objc.ID, value uint64, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.uint64Error != nil {
		r.uint64Error(block, value, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, uint64, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, value, errID)
	return nil
}

// BoolError invokes block(value, err).
func (r *BlockInvoker) BoolError(block objc.ID, value bool, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.boolError != nil {
		r.boolError(block, value, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, bool, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, value, errID)
	return nil
}

// UintptrError invokes block(value, err).
func (r *BlockInvoker) UintptrError(block objc.ID, value uintptr, errID objc.ID) error {
	r.loadShims()
	if err := checkBlock(block); err != nil {
		return err
	}
	if r.uintptrError != nil {
		r.uintptrError(block, value, errID)
		return nil
	}
	invoke, err := blockInvoke(block)
	if err != nil {
		return err
	}
	var fn func(objc.ID, uintptr, objc.ID)
	purego.RegisterFunc(&fn, invoke)
	fn(block, value, errID)
	return nil
}
