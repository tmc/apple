package main

import (
	"fmt"

	"github.com/ebitengine/purego"
)

// amfiFramework is the private framework that owns the path validator.
const amfiFramework = "/System/Library/PrivateFrameworks/AppleMobileFileIntegrity.framework/AppleMobileFileIntegrity"

// validatorClass is the macOS path-validator class. AMFIPathValidator_ios is a
// sibling with a different, smaller layout — never bind to it and never assume
// anything carries over.
const validatorClass = "AMFIPathValidator_macos"

// Ivar offsets on AMFIPathValidator_macos, and the selector this tool hooks.
// These were read off macOS 27 with a live objc_getClass + ivar_getOffset
// probe and are hard fact for that OS build. They are ABI-pinned for the build
// and absent from Apple's documentation set, so they are written here rather
// than generated; resolveAMFI re-reads them from the live runtime and refuses
// to run if the runtime disagrees. There is no _cdhash ivar on this class:
// cdhashAsData is a computed method over _code, and _cdhash exists only on the
// _ios sibling.
const (
	ivarCode                      = 8  // _code, SecStaticCodeRef
	ivarHasRestrictedEntitlements = 36 // _hasRestrictedEntitlements, BOOL
	ivarIsValid                   = 49 // _isValid, BOOL
	ivarIsApple                   = 51 // _isApple, BOOL
	ivarShouldUnrestrict          = 53 // _shouldUnrestrict, BOOL
	ivarSigningIdentifier         = 96 // _signingIdentifier, NSString *
)

// validateSelector is the method whose return this tool breaks at. It is what
// populates the result ivars above.
const validateSelector = "validateWithError:"

// amfi holds everything resolved from the AMFI framework in our own process,
// translated to the addresses a breakpoint and ivar writes will use in the
// target. Every field is resolved before attach.
type amfi struct {
	imp uint64 // slid IMP of -[AMFIPathValidator_macos validateWithError:]

	// Ivar offsets, re-read from the live runtime and asserted against the
	// pinned constants above.
	offIsValid                   int64
	offIsApple                   int64
	offShouldUnrestrict          int64
	offHasRestrictedEntitlements int64
	offSigningIdentifier         int64

	boot bootargState
}

// bootargState reports whether amfi_get_out_of_my_way is set on this boot,
// read directly rather than inferred from a failure.
type bootargState struct {
	resolved      bool
	raw           uint32
	getOutOfMyWay bool
}

// amfiGetOutOfMyWayBit is the amfi_get_out_of_my_way flag within the value
// amfi_interface_query_bootarg_state returns. ABI-pinned, undocumented.
const amfiGetOutOfMyWayBit = 1 << 0

// resolveAMFI dlopens the AMFI framework in our own process and resolves the
// validator IMP and result-ivar offsets from the Objective-C runtime. It does
// not parse Mach-O symbols: the runtime yields the same numbers with none of
// the fragility. It re-reads each pinned ivar offset and refuses to run if the
// runtime disagrees, because the whole patch is written blind to those offsets.
func resolveAMFI() (*amfi, error) {
	if _, err := purego.Dlopen(amfiFramework, purego.RTLD_NOW|purego.RTLD_GLOBAL); err != nil {
		return nil, fmt.Errorf("dlopen AMFI framework: %w", err)
	}

	var (
		getClass    func(string) uintptr
		getIvar     func(uintptr, string) uintptr
		ivarOffset  func(uintptr) int64
		getMethod   func(uintptr, uintptr) uintptr
		getIMP      func(uintptr) uintptr
		selRegister func(string) uintptr
	)
	for _, b := range []struct {
		ptr  any
		name string
	}{
		{&getClass, "objc_getClass"},
		{&getIvar, "class_getInstanceVariable"},
		{&ivarOffset, "ivar_getOffset"},
		{&getMethod, "class_getInstanceMethod"},
		{&getIMP, "method_getImplementation"},
		{&selRegister, "sel_registerName"},
	} {
		if err := registerFunc(b.ptr, b.name); err != nil {
			return nil, err
		}
	}

	cls := getClass(validatorClass)
	if cls == 0 {
		return nil, fmt.Errorf("objc_getClass(%q) returned nil", validatorClass)
	}

	method := getMethod(cls, selRegister(validateSelector))
	if method == 0 {
		return nil, fmt.Errorf("no -[%s %s]", validatorClass, validateSelector)
	}
	imp := getIMP(method)
	if imp == 0 {
		return nil, fmt.Errorf("nil IMP for -[%s %s]", validatorClass, validateSelector)
	}

	resolveOffset := func(name string, want int64) (int64, error) {
		iv := getIvar(cls, name)
		if iv == 0 {
			return 0, fmt.Errorf("ivar %s absent on %s", name, validatorClass)
		}
		got := ivarOffset(iv)
		if got != want {
			return 0, fmt.Errorf("ivar %s offset %d disagrees with pinned %d; layout changed, refusing to patch blind", name, got, want)
		}
		return got, nil
	}

	a := &amfi{imp: uint64(imp)}
	var err error
	if a.offIsValid, err = resolveOffset("_isValid", ivarIsValid); err != nil {
		return nil, err
	}
	if a.offIsApple, err = resolveOffset("_isApple", ivarIsApple); err != nil {
		return nil, err
	}
	if a.offShouldUnrestrict, err = resolveOffset("_shouldUnrestrict", ivarShouldUnrestrict); err != nil {
		return nil, err
	}
	if a.offHasRestrictedEntitlements, err = resolveOffset("_hasRestrictedEntitlements", ivarHasRestrictedEntitlements); err != nil {
		return nil, err
	}
	if a.offSigningIdentifier, err = resolveOffset("_signingIdentifier", ivarSigningIdentifier); err != nil {
		return nil, err
	}

	a.boot = queryBootargState()
	return a, nil
}

// queryBootargState binds and calls amfi_interface_query_bootarg_state, a
// preflight-only read. It changes nothing about the mechanism; it reports
// whether amfi_get_out_of_my_way is actually set on this boot instead of
// leaving the operator to infer it. A resolution failure is not fatal.
func queryBootargState() bootargState {
	var query func() uint32
	if err := registerFunc(&query, "amfi_interface_query_bootarg_state"); err != nil {
		return bootargState{}
	}
	// Bind amfi_interface_cdhash_in_trustcache too, purely to pre-fault it so
	// a consumer that wants to ask "is the target already trusted?" need not
	// resolve a symbol while amfid is stopped. It is not called here: answering
	// it needs the target's cdhash, which is off the one-shot critical path.
	var trust func(uintptr, uintptr) int32
	_ = registerFunc(&trust, "amfi_interface_cdhash_in_trustcache")

	raw := query()
	return bootargState{resolved: true, raw: raw, getOutOfMyWay: raw&amfiGetOutOfMyWayBit != 0}
}

func (a *amfi) bootargState() bootargState { return a.boot }

func (a *amfi) summary() string {
	return fmt.Sprintf("resolved %s: validateWithError: IMP %#x; ivars isValid=%d isApple=%d shouldUnrestrict=%d hasRestrictedEntitlements=%d signingIdentifier=%d",
		validatorClass, a.imp, a.offIsValid, a.offIsApple, a.offShouldUnrestrict, a.offHasRestrictedEntitlements, a.offSigningIdentifier)
}
