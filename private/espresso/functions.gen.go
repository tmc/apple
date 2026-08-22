// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"fmt"

	"github.com/ebitengine/purego"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("espresso: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("espresso: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("espresso: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("espresso: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _espressoContextDestroy func(ctx EspressoContext)
var _espressoContextDestroyErr error

func tryEspressoContextDestroy(ctx EspressoContext) error {
	if _espressoContextDestroy == nil {
		return symbolCallError("espresso_context_destroy", "", _espressoContextDestroyErr)
	}
	_espressoContextDestroy(ctx)
	return nil
}

// EspressoContextDestroy.
func EspressoContextDestroy(ctx EspressoContext) error {
	return tryEspressoContextDestroy(ctx)
}

var _espressoCreateContext func(platform int32, options int32) EspressoContext
var _espressoCreateContextErr error

func tryEspressoCreateContext(platform int32, options int32) (EspressoContext, error) {
	if _espressoCreateContext == nil {
		return EspressoContext{}, symbolCallError("espresso_create_context", "", _espressoCreateContextErr)
	}
	return _espressoCreateContext(platform, options), nil
}

// EspressoCreateContext.
func EspressoCreateContext(platform int32, options int32) (EspressoContext, error) {
	return tryEspressoCreateContext(platform, options)
}

var _espressoCreatePlan func(ctx EspressoContext, platform int32) EspressoPlan
var _espressoCreatePlanErr error

func tryEspressoCreatePlan(ctx EspressoContext, platform int32) (EspressoPlan, error) {
	if _espressoCreatePlan == nil {
		return *new(EspressoPlan), symbolCallError("espresso_create_plan", "", _espressoCreatePlanErr)
	}
	return _espressoCreatePlan(ctx, platform), nil
}

// EspressoCreatePlan.
func EspressoCreatePlan(ctx EspressoContext, platform int32) (EspressoPlan, error) {
	return tryEspressoCreatePlan(ctx, platform)
}

var _espressoGetVersionString func() uintptr
var _espressoGetVersionStringErr error

func tryEspressoGetVersionString() (uintptr, error) {
	if _espressoGetVersionString == nil {
		return 0, symbolCallError("espresso_get_version_string", "", _espressoGetVersionStringErr)
	}
	return _espressoGetVersionString(), nil
}

// EspressoGetVersionString.
func EspressoGetVersionString() (uintptr, error) {
	return tryEspressoGetVersionString()
}

var _espressoPlanBuild func(plan EspressoPlan) int32
var _espressoPlanBuildErr error

func tryEspressoPlanBuild(plan EspressoPlan) (int32, error) {
	if _espressoPlanBuild == nil {
		return 0, symbolCallError("espresso_plan_build", "", _espressoPlanBuildErr)
	}
	return _espressoPlanBuild(plan), nil
}

// EspressoPlanBuild.
func EspressoPlanBuild(plan EspressoPlan) (int32, error) {
	return tryEspressoPlanBuild(plan)
}

var _espressoPlanDestroy func(plan EspressoPlan)
var _espressoPlanDestroyErr error

func tryEspressoPlanDestroy(plan EspressoPlan) error {
	if _espressoPlanDestroy == nil {
		return symbolCallError("espresso_plan_destroy", "", _espressoPlanDestroyErr)
	}
	_espressoPlanDestroy(plan)
	return nil
}

// EspressoPlanDestroy.
func EspressoPlanDestroy(plan EspressoPlan) error {
	return tryEspressoPlanDestroy(plan)
}

var _espressoPlanExecuteSync func(plan EspressoPlan) int32
var _espressoPlanExecuteSyncErr error

func tryEspressoPlanExecuteSync(plan EspressoPlan) (int32, error) {
	if _espressoPlanExecuteSync == nil {
		return 0, symbolCallError("espresso_plan_execute_sync", "", _espressoPlanExecuteSyncErr)
	}
	return _espressoPlanExecuteSync(plan), nil
}

// EspressoPlanExecuteSync.
func EspressoPlanExecuteSync(plan EspressoPlan) (int32, error) {
	return tryEspressoPlanExecuteSync(plan)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_espressoContextDestroy, &_espressoContextDestroyErr, frameworkHandle, "espresso_context_destroy", "")
	registerFunc(&_espressoCreateContext, &_espressoCreateContextErr, frameworkHandle, "espresso_create_context", "")
	registerFunc(&_espressoCreatePlan, &_espressoCreatePlanErr, frameworkHandle, "espresso_create_plan", "")
	registerFunc(&_espressoGetVersionString, &_espressoGetVersionStringErr, frameworkHandle, "espresso_get_version_string", "")
	registerFunc(&_espressoPlanBuild, &_espressoPlanBuildErr, frameworkHandle, "espresso_plan_build", "")
	registerFunc(&_espressoPlanDestroy, &_espressoPlanDestroyErr, frameworkHandle, "espresso_plan_destroy", "")
	registerFunc(&_espressoPlanExecuteSync, &_espressoPlanExecuteSyncErr, frameworkHandle, "espresso_plan_execute_sync", "")
}
