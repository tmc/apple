// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"fmt"
	"unsafe"

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
		return fmt.Sprintf("kernel: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("kernel: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("kernel: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("kernel: register symbol %s: %v", name, r)
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

var _mD5Final func(arg0 *byte, arg1 *MD5_CTX)
var _mD5FinalErr error

func tryMD5Final(arg0 []byte, arg1 *MD5_CTX) error {
	if _mD5Final == nil {
		return symbolCallError("MD5Final", "10.0", _mD5FinalErr)
	}
	_mD5Final(unsafe.SliceData(arg0), arg1)
	return nil
}

// MD5Final.
//
// See: https://developer.apple.com/documentation/kernel/1537348-md5final
func MD5Final(arg0 []byte, arg1 *MD5_CTX) {
	if callErr := tryMD5Final(arg0, arg1); callErr != nil {
		panic(callErr)
	}
}

var _acos func(arg0 float64) float64
var _acosErr error

func tryAcos(arg0 float64) (float64, error) {
	if _acos == nil {
		return 0.0, symbolCallError("acos", "10.10", _acosErr)
	}
	return _acos(arg0), nil
}

// Acos.
//
// See: https://developer.apple.com/documentation/kernel/1557251-acos
func Acos(arg0 float64) float64 {
	result, callErr := tryAcos(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _acosf func(arg0 float32) float32
var _acosfErr error

func tryAcosf(arg0 float32) (float32, error) {
	if _acosf == nil {
		return 0.0, symbolCallError("acosf", "10.10", _acosfErr)
	}
	return _acosf(arg0), nil
}

// Acosf.
//
// See: https://developer.apple.com/documentation/kernel/1557163-acosf
func Acosf(arg0 float32) float32 {
	result, callErr := tryAcosf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _acosh func(arg0 float64) float64
var _acoshErr error

func tryAcosh(arg0 float64) (float64, error) {
	if _acosh == nil {
		return 0.0, symbolCallError("acosh", "10.10", _acoshErr)
	}
	return _acosh(arg0), nil
}

// Acosh.
//
// See: https://developer.apple.com/documentation/kernel/1557170-acosh
func Acosh(arg0 float64) float64 {
	result, callErr := tryAcosh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _acoshf func(arg0 float32) float32
var _acoshfErr error

func tryAcoshf(arg0 float32) (float32, error) {
	if _acoshf == nil {
		return 0.0, symbolCallError("acoshf", "10.10", _acoshfErr)
	}
	return _acoshf(arg0), nil
}

// Acoshf.
//
// See: https://developer.apple.com/documentation/kernel/1557276-acoshf
func Acoshf(arg0 float32) float32 {
	result, callErr := tryAcoshf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _acoshl func(arg0 float64) float64
var _acoshlErr error

func tryAcoshl(arg0 float64) (float64, error) {
	if _acoshl == nil {
		return 0.0, symbolCallError("acoshl", "10.10", _acoshlErr)
	}
	return _acoshl(arg0), nil
}

// Acoshl.
//
// See: https://developer.apple.com/documentation/kernel/1557266-acoshl
func Acoshl(arg0 float64) float64 {
	result, callErr := tryAcoshl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _acosl func(arg0 float64) float64
var _acoslErr error

func tryAcosl(arg0 float64) (float64, error) {
	if _acosl == nil {
		return 0.0, symbolCallError("acosl", "10.10", _acoslErr)
	}
	return _acosl(arg0), nil
}

// Acosl.
//
// See: https://developer.apple.com/documentation/kernel/1557197-acosl
func Acosl(arg0 float64) float64 {
	result, callErr := tryAcosl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _act_get_state func(target_act Thread_read_t, flavor int32, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t
var _act_get_stateErr error

func tryAct_get_state(target_act Thread_read_t, flavor int32, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _act_get_state == nil {
		return *new(Kern_return_t), symbolCallError("act_get_state", "10.0", _act_get_stateErr)
	}
	return _act_get_state(target_act, flavor, old_state, old_stateCnt), nil
}

// Act_get_state.
//
// See: https://developer.apple.com/documentation/kernel/1418936-act_get_state
func Act_get_state(target_act Thread_read_t, flavor int32, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryAct_get_state(target_act, flavor, old_state, old_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _act_set_state func(target_act Thread_act_t, flavor int32, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t
var _act_set_stateErr error

func tryAct_set_state(target_act Thread_act_t, flavor int32, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _act_set_state == nil {
		return *new(Kern_return_t), symbolCallError("act_set_state", "10.0", _act_set_stateErr)
	}
	return _act_set_state(target_act, flavor, new_state, new_stateCnt), nil
}

// Act_set_state.
//
// See: https://developer.apple.com/documentation/kernel/1418961-act_set_state
func Act_set_state(target_act Thread_act_t, flavor int32, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryAct_set_state(target_act, flavor, new_state, new_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asin func(arg0 float64) float64
var _asinErr error

func tryAsin(arg0 float64) (float64, error) {
	if _asin == nil {
		return 0.0, symbolCallError("asin", "10.10", _asinErr)
	}
	return _asin(arg0), nil
}

// Asin.
//
// See: https://developer.apple.com/documentation/kernel/1557225-asin
func Asin(arg0 float64) float64 {
	result, callErr := tryAsin(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asinf func(arg0 float32) float32
var _asinfErr error

func tryAsinf(arg0 float32) (float32, error) {
	if _asinf == nil {
		return 0.0, symbolCallError("asinf", "10.10", _asinfErr)
	}
	return _asinf(arg0), nil
}

// Asinf.
//
// See: https://developer.apple.com/documentation/kernel/1557356-asinf
func Asinf(arg0 float32) float32 {
	result, callErr := tryAsinf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asinh func(arg0 float64) float64
var _asinhErr error

func tryAsinh(arg0 float64) (float64, error) {
	if _asinh == nil {
		return 0.0, symbolCallError("asinh", "10.10", _asinhErr)
	}
	return _asinh(arg0), nil
}

// Asinh.
//
// See: https://developer.apple.com/documentation/kernel/1557211-asinh
func Asinh(arg0 float64) float64 {
	result, callErr := tryAsinh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asinhf func(arg0 float32) float32
var _asinhfErr error

func tryAsinhf(arg0 float32) (float32, error) {
	if _asinhf == nil {
		return 0.0, symbolCallError("asinhf", "10.10", _asinhfErr)
	}
	return _asinhf(arg0), nil
}

// Asinhf.
//
// See: https://developer.apple.com/documentation/kernel/1557278-asinhf
func Asinhf(arg0 float32) float32 {
	result, callErr := tryAsinhf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asinhl func(arg0 float64) float64
var _asinhlErr error

func tryAsinhl(arg0 float64) (float64, error) {
	if _asinhl == nil {
		return 0.0, symbolCallError("asinhl", "10.10", _asinhlErr)
	}
	return _asinhl(arg0), nil
}

// Asinhl.
//
// See: https://developer.apple.com/documentation/kernel/1557157-asinhl
func Asinhl(arg0 float64) float64 {
	result, callErr := tryAsinhl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _asinl func(arg0 float64) float64
var _asinlErr error

func tryAsinl(arg0 float64) (float64, error) {
	if _asinl == nil {
		return 0.0, symbolCallError("asinl", "10.10", _asinlErr)
	}
	return _asinl(arg0), nil
}

// Asinl.
//
// See: https://developer.apple.com/documentation/kernel/1557222-asinl
func Asinl(arg0 float64) float64 {
	result, callErr := tryAsinl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atan func(arg0 float64) float64
var _atanErr error

func tryAtan(arg0 float64) (float64, error) {
	if _atan == nil {
		return 0.0, symbolCallError("atan", "10.10", _atanErr)
	}
	return _atan(arg0), nil
}

// Atan.
//
// See: https://developer.apple.com/documentation/kernel/1557165-atan
func Atan(arg0 float64) float64 {
	result, callErr := tryAtan(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atan2 func(arg0 float64, arg1 float64) float64
var _atan2Err error

func tryAtan2(arg0 float64, arg1 float64) (float64, error) {
	if _atan2 == nil {
		return 0.0, symbolCallError("atan2", "10.10", _atan2Err)
	}
	return _atan2(arg0, arg1), nil
}

// Atan2.
//
// See: https://developer.apple.com/documentation/kernel/1557368-atan2
func Atan2(arg0 float64, arg1 float64) float64 {
	result, callErr := tryAtan2(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atan2f func(arg0 float32, arg1 float32) float32
var _atan2fErr error

func tryAtan2f(arg0 float32, arg1 float32) (float32, error) {
	if _atan2f == nil {
		return 0.0, symbolCallError("atan2f", "10.10", _atan2fErr)
	}
	return _atan2f(arg0, arg1), nil
}

// Atan2f.
//
// See: https://developer.apple.com/documentation/kernel/1557144-atan2f
func Atan2f(arg0 float32, arg1 float32) float32 {
	result, callErr := tryAtan2f(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atan2l func(arg0 float64, arg1 float64) float64
var _atan2lErr error

func tryAtan2l(arg0 float64, arg1 float64) (float64, error) {
	if _atan2l == nil {
		return 0.0, symbolCallError("atan2l", "10.10", _atan2lErr)
	}
	return _atan2l(arg0, arg1), nil
}

// Atan2l.
//
// See: https://developer.apple.com/documentation/kernel/1557326-atan2l
func Atan2l(arg0 float64, arg1 float64) float64 {
	result, callErr := tryAtan2l(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atanf func(arg0 float32) float32
var _atanfErr error

func tryAtanf(arg0 float32) (float32, error) {
	if _atanf == nil {
		return 0.0, symbolCallError("atanf", "10.10", _atanfErr)
	}
	return _atanf(arg0), nil
}

// Atanf.
//
// See: https://developer.apple.com/documentation/kernel/1557247-atanf
func Atanf(arg0 float32) float32 {
	result, callErr := tryAtanf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atanh func(arg0 float64) float64
var _atanhErr error

func tryAtanh(arg0 float64) (float64, error) {
	if _atanh == nil {
		return 0.0, symbolCallError("atanh", "10.10", _atanhErr)
	}
	return _atanh(arg0), nil
}

// Atanh.
//
// See: https://developer.apple.com/documentation/kernel/1557372-atanh
func Atanh(arg0 float64) float64 {
	result, callErr := tryAtanh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atanhf func(arg0 float32) float32
var _atanhfErr error

func tryAtanhf(arg0 float32) (float32, error) {
	if _atanhf == nil {
		return 0.0, symbolCallError("atanhf", "10.10", _atanhfErr)
	}
	return _atanhf(arg0), nil
}

// Atanhf.
//
// See: https://developer.apple.com/documentation/kernel/1557262-atanhf
func Atanhf(arg0 float32) float32 {
	result, callErr := tryAtanhf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atanhl func(arg0 float64) float64
var _atanhlErr error

func tryAtanhl(arg0 float64) (float64, error) {
	if _atanhl == nil {
		return 0.0, symbolCallError("atanhl", "10.10", _atanhlErr)
	}
	return _atanhl(arg0), nil
}

// Atanhl.
//
// See: https://developer.apple.com/documentation/kernel/1557230-atanhl
func Atanhl(arg0 float64) float64 {
	result, callErr := tryAtanhl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _atanl func(arg0 float64) float64
var _atanlErr error

func tryAtanl(arg0 float64) (float64, error) {
	if _atanl == nil {
		return 0.0, symbolCallError("atanl", "10.10", _atanlErr)
	}
	return _atanl(arg0), nil
}

// Atanl.
//
// See: https://developer.apple.com/documentation/kernel/1557198-atanl
func Atanl(arg0 float64) float64 {
	result, callErr := tryAtanl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _backtrace func(bt *uintptr, btlen uint32, ctl *Backtrace_control, info_out *Backtrace_info_t) uint32
var _backtraceErr error

func tryBacktrace(bt *uintptr, btlen uint32, ctl *Backtrace_control, info_out *Backtrace_info_t) (uint32, error) {
	if _backtrace == nil {
		return 0, symbolCallError("backtrace", "10.12", _backtraceErr)
	}
	return _backtrace(bt, btlen, ctl, info_out), nil
}

// Backtrace.
//
// See: https://developer.apple.com/documentation/kernel/1644760-backtrace
func Backtrace(bt *uintptr, btlen uint32, ctl *Backtrace_control, info_out *Backtrace_info_t) uint32 {
	result, callErr := tryBacktrace(bt, btlen, ctl, info_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _bcmp func(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) int32
var _bcmpErr error

func tryBcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) (int32, error) {
	if _bcmp == nil {
		return 0, symbolCallError("bcmp", "10.0", _bcmpErr)
	}
	return _bcmp(s1, s2, n), nil
}

// Bcmp.
//
// See: https://developer.apple.com/documentation/kernel/1579330-bcmp
func Bcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) int32 {
	result, callErr := tryBcmp(s1, s2, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _bcopy func(src unsafe.Pointer, dst unsafe.Pointer, n uintptr)
var _bcopyErr error

func tryBcopy(src unsafe.Pointer, dst unsafe.Pointer, n uintptr) error {
	if _bcopy == nil {
		return symbolCallError("bcopy", "10.0", _bcopyErr)
	}
	_bcopy(src, dst, n)
	return nil
}

// Bcopy.
//
// See: https://developer.apple.com/documentation/kernel/1579339-bcopy
func Bcopy(src unsafe.Pointer, dst unsafe.Pointer, n uintptr) {
	if callErr := tryBcopy(src, dst, n); callErr != nil {
		panic(callErr)
	}
}

var _bzero func(s unsafe.Pointer, n uintptr)
var _bzeroErr error

func tryBzero(s unsafe.Pointer, n uintptr) error {
	if _bzero == nil {
		return symbolCallError("bzero", "10.0", _bzeroErr)
	}
	_bzero(s, n)
	return nil
}

// Bzero.
//
// See: https://developer.apple.com/documentation/kernel/1579350-bzero
func Bzero(s unsafe.Pointer, n uintptr) {
	if callErr := tryBzero(s, n); callErr != nil {
		panic(callErr)
	}
}

var _cbrt func(arg0 float64) float64
var _cbrtErr error

func tryCbrt(arg0 float64) (float64, error) {
	if _cbrt == nil {
		return 0.0, symbolCallError("cbrt", "10.10", _cbrtErr)
	}
	return _cbrt(arg0), nil
}

// Cbrt.
//
// See: https://developer.apple.com/documentation/kernel/1557257-cbrt
func Cbrt(arg0 float64) float64 {
	result, callErr := tryCbrt(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cbrtf func(arg0 float32) float32
var _cbrtfErr error

func tryCbrtf(arg0 float32) (float32, error) {
	if _cbrtf == nil {
		return 0.0, symbolCallError("cbrtf", "10.10", _cbrtfErr)
	}
	return _cbrtf(arg0), nil
}

// Cbrtf.
//
// See: https://developer.apple.com/documentation/kernel/1557327-cbrtf
func Cbrtf(arg0 float32) float32 {
	result, callErr := tryCbrtf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cbrtl func(arg0 float64) float64
var _cbrtlErr error

func tryCbrtl(arg0 float64) (float64, error) {
	if _cbrtl == nil {
		return 0.0, symbolCallError("cbrtl", "10.10", _cbrtlErr)
	}
	return _cbrtl(arg0), nil
}

// Cbrtl.
//
// See: https://developer.apple.com/documentation/kernel/1557373-cbrtl
func Cbrtl(arg0 float64) float64 {
	result, callErr := tryCbrtl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ceil func(arg0 float64) float64
var _ceilErr error

func tryCeil(arg0 float64) (float64, error) {
	if _ceil == nil {
		return 0.0, symbolCallError("ceil", "10.10", _ceilErr)
	}
	return _ceil(arg0), nil
}

// Ceil.
//
// See: https://developer.apple.com/documentation/kernel/1557272-ceil
func Ceil(arg0 float64) float64 {
	result, callErr := tryCeil(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ceilf func(arg0 float32) float32
var _ceilfErr error

func tryCeilf(arg0 float32) (float32, error) {
	if _ceilf == nil {
		return 0.0, symbolCallError("ceilf", "10.10", _ceilfErr)
	}
	return _ceilf(arg0), nil
}

// Ceilf.
//
// See: https://developer.apple.com/documentation/kernel/1557263-ceilf
func Ceilf(arg0 float32) float32 {
	result, callErr := tryCeilf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ceill func(arg0 float64) float64
var _ceillErr error

func tryCeill(arg0 float64) (float64, error) {
	if _ceill == nil {
		return 0.0, symbolCallError("ceill", "10.10", _ceillErr)
	}
	return _ceill(arg0), nil
}

// Ceill.
//
// See: https://developer.apple.com/documentation/kernel/1557207-ceill
func Ceill(arg0 float64) float64 {
	result, callErr := tryCeill(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_alarm func(clock_serv Clock_serv_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t, alarm_port Clock_reply_t) Kern_return_t
var _clock_alarmErr error

func tryClock_alarm(clock_serv Clock_serv_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t, alarm_port Clock_reply_t) (Kern_return_t, error) {
	if _clock_alarm == nil {
		return *new(Kern_return_t), symbolCallError("clock_alarm", "10.0", _clock_alarmErr)
	}
	return _clock_alarm(clock_serv, alarm_type, alarm_time, alarm_port), nil
}

// Clock_alarm.
//
// See: https://developer.apple.com/documentation/kernel/1420037-clock_alarm
func Clock_alarm(clock_serv Clock_serv_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t, alarm_port Clock_reply_t) Kern_return_t {
	result, callErr := tryClock_alarm(clock_serv, alarm_type, alarm_time, alarm_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_alarm_reply func(alarm_port Clock_reply_t, alarm_code Kern_return_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t) Kern_return_t
var _clock_alarm_replyErr error

func tryClock_alarm_reply(alarm_port Clock_reply_t, alarm_code Kern_return_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t) (Kern_return_t, error) {
	if _clock_alarm_reply == nil {
		return *new(Kern_return_t), symbolCallError("clock_alarm_reply", "10.0", _clock_alarm_replyErr)
	}
	return _clock_alarm_reply(alarm_port, alarm_code, alarm_type, alarm_time), nil
}

// Clock_alarm_reply.
//
// See: https://developer.apple.com/documentation/kernel/1390972-clock_alarm_reply
func Clock_alarm_reply(alarm_port Clock_reply_t, alarm_code Kern_return_t, alarm_type Alarm_type_t, alarm_time Mach_timespec_t) Kern_return_t {
	result, callErr := tryClock_alarm_reply(alarm_port, alarm_code, alarm_type, alarm_time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_get_attributes func(clock_serv Clock_serv_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt *Mach_msg_type_number_t) Kern_return_t
var _clock_get_attributesErr error

func tryClock_get_attributes(clock_serv Clock_serv_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _clock_get_attributes == nil {
		return *new(Kern_return_t), symbolCallError("clock_get_attributes", "10.0", _clock_get_attributesErr)
	}
	return _clock_get_attributes(clock_serv, flavor, clock_attr, clock_attrCnt), nil
}

// Clock_get_attributes.
//
// See: https://developer.apple.com/documentation/kernel/1420071-clock_get_attributes
func Clock_get_attributes(clock_serv Clock_serv_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryClock_get_attributes(clock_serv, flavor, clock_attr, clock_attrCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_get_time func(clock_serv Clock_serv_t, cur_time *Mach_timespec_t) Kern_return_t
var _clock_get_timeErr error

func tryClock_get_time(clock_serv Clock_serv_t, cur_time *Mach_timespec_t) (Kern_return_t, error) {
	if _clock_get_time == nil {
		return *new(Kern_return_t), symbolCallError("clock_get_time", "10.0", _clock_get_timeErr)
	}
	return _clock_get_time(clock_serv, cur_time), nil
}

// Clock_get_time.
//
// See: https://developer.apple.com/documentation/kernel/1420035-clock_get_time
func Clock_get_time(clock_serv Clock_serv_t, cur_time *Mach_timespec_t) Kern_return_t {
	result, callErr := tryClock_get_time(clock_serv, cur_time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_set_attributes func(clock_ctrl Clock_ctrl_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt Mach_msg_type_number_t) Kern_return_t
var _clock_set_attributesErr error

func tryClock_set_attributes(clock_ctrl Clock_ctrl_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _clock_set_attributes == nil {
		return *new(Kern_return_t), symbolCallError("clock_set_attributes", "10.0", _clock_set_attributesErr)
	}
	return _clock_set_attributes(clock_ctrl, flavor, clock_attr, clock_attrCnt), nil
}

// Clock_set_attributes.
//
// See: https://developer.apple.com/documentation/kernel/1551054-clock_set_attributes
func Clock_set_attributes(clock_ctrl Clock_ctrl_t, flavor Clock_flavor_t, clock_attr Clock_attr_t, clock_attrCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryClock_set_attributes(clock_ctrl, flavor, clock_attr, clock_attrCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _clock_set_time func(clock_ctrl Clock_ctrl_t, new_time Mach_timespec_t) Kern_return_t
var _clock_set_timeErr error

func tryClock_set_time(clock_ctrl Clock_ctrl_t, new_time Mach_timespec_t) (Kern_return_t, error) {
	if _clock_set_time == nil {
		return *new(Kern_return_t), symbolCallError("clock_set_time", "10.0", _clock_set_timeErr)
	}
	return _clock_set_time(clock_ctrl, new_time), nil
}

// Clock_set_time.
//
// See: https://developer.apple.com/documentation/kernel/1551049-clock_set_time
func Clock_set_time(clock_ctrl Clock_ctrl_t, new_time Mach_timespec_t) Kern_return_t {
	result, callErr := tryClock_set_time(clock_ctrl, new_time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copysign func(arg0 float64, arg1 float64) float64
var _copysignErr error

func tryCopysign(arg0 float64, arg1 float64) (float64, error) {
	if _copysign == nil {
		return 0.0, symbolCallError("copysign", "10.10", _copysignErr)
	}
	return _copysign(arg0, arg1), nil
}

// Copysign.
//
// See: https://developer.apple.com/documentation/kernel/1557306-copysign
func Copysign(arg0 float64, arg1 float64) float64 {
	result, callErr := tryCopysign(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copysignf func(arg0 float32, arg1 float32) float32
var _copysignfErr error

func tryCopysignf(arg0 float32, arg1 float32) (float32, error) {
	if _copysignf == nil {
		return 0.0, symbolCallError("copysignf", "10.10", _copysignfErr)
	}
	return _copysignf(arg0, arg1), nil
}

// Copysignf.
//
// See: https://developer.apple.com/documentation/kernel/1557234-copysignf
func Copysignf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryCopysignf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copysignl func(arg0 float64, arg1 float64) float64
var _copysignlErr error

func tryCopysignl(arg0 float64, arg1 float64) (float64, error) {
	if _copysignl == nil {
		return 0.0, symbolCallError("copysignl", "10.10", _copysignlErr)
	}
	return _copysignl(arg0, arg1), nil
}

// Copysignl.
//
// See: https://developer.apple.com/documentation/kernel/1557294-copysignl
func Copysignl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryCopysignl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cos func(arg0 float64) float64
var _cosErr error

func tryCos(arg0 float64) (float64, error) {
	if _cos == nil {
		return 0.0, symbolCallError("cos", "10.10", _cosErr)
	}
	return _cos(arg0), nil
}

// Cos.
//
// See: https://developer.apple.com/documentation/kernel/1557361-cos
func Cos(arg0 float64) float64 {
	result, callErr := tryCos(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cosf func(arg0 float32) float32
var _cosfErr error

func tryCosf(arg0 float32) (float32, error) {
	if _cosf == nil {
		return 0.0, symbolCallError("cosf", "10.10", _cosfErr)
	}
	return _cosf(arg0), nil
}

// Cosf.
//
// See: https://developer.apple.com/documentation/kernel/1532192-cosf
func Cosf(arg0 float32) float32 {
	result, callErr := tryCosf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cosh func(arg0 float64) float64
var _coshErr error

func tryCosh(arg0 float64) (float64, error) {
	if _cosh == nil {
		return 0.0, symbolCallError("cosh", "10.10", _coshErr)
	}
	return _cosh(arg0), nil
}

// Cosh.
//
// See: https://developer.apple.com/documentation/kernel/1557145-cosh
func Cosh(arg0 float64) float64 {
	result, callErr := tryCosh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _coshf func(arg0 float32) float32
var _coshfErr error

func tryCoshf(arg0 float32) (float32, error) {
	if _coshf == nil {
		return 0.0, symbolCallError("coshf", "10.10", _coshfErr)
	}
	return _coshf(arg0), nil
}

// Coshf.
//
// See: https://developer.apple.com/documentation/kernel/1557149-coshf
func Coshf(arg0 float32) float32 {
	result, callErr := tryCoshf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _coshl func(arg0 float64) float64
var _coshlErr error

func tryCoshl(arg0 float64) (float64, error) {
	if _coshl == nil {
		return 0.0, symbolCallError("coshl", "10.10", _coshlErr)
	}
	return _coshl(arg0), nil
}

// Coshl.
//
// See: https://developer.apple.com/documentation/kernel/1557214-coshl
func Coshl(arg0 float64) float64 {
	result, callErr := tryCoshl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cosl func(arg0 float64) float64
var _coslErr error

func tryCosl(arg0 float64) (float64, error) {
	if _cosl == nil {
		return 0.0, symbolCallError("cosl", "10.10", _coslErr)
	}
	return _cosl(arg0), nil
}

// Cosl.
//
// See: https://developer.apple.com/documentation/kernel/1557255-cosl
func Cosl(arg0 float64) float64 {
	result, callErr := tryCosl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erf func(arg0 float64) float64
var _erfErr error

func tryErf(arg0 float64) (float64, error) {
	if _erf == nil {
		return 0.0, symbolCallError("erf", "10.10", _erfErr)
	}
	return _erf(arg0), nil
}

// Erf.
//
// See: https://developer.apple.com/documentation/kernel/1557352-erf
func Erf(arg0 float64) float64 {
	result, callErr := tryErf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erfc func(arg0 float64) float64
var _erfcErr error

func tryErfc(arg0 float64) (float64, error) {
	if _erfc == nil {
		return 0.0, symbolCallError("erfc", "10.10", _erfcErr)
	}
	return _erfc(arg0), nil
}

// Erfc.
//
// See: https://developer.apple.com/documentation/kernel/1557322-erfc
func Erfc(arg0 float64) float64 {
	result, callErr := tryErfc(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erfcf func(arg0 float32) float32
var _erfcfErr error

func tryErfcf(arg0 float32) (float32, error) {
	if _erfcf == nil {
		return 0.0, symbolCallError("erfcf", "10.10", _erfcfErr)
	}
	return _erfcf(arg0), nil
}

// Erfcf.
//
// See: https://developer.apple.com/documentation/kernel/1557244-erfcf
func Erfcf(arg0 float32) float32 {
	result, callErr := tryErfcf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erfcl func(arg0 float64) float64
var _erfclErr error

func tryErfcl(arg0 float64) (float64, error) {
	if _erfcl == nil {
		return 0.0, symbolCallError("erfcl", "10.10", _erfclErr)
	}
	return _erfcl(arg0), nil
}

// Erfcl.
//
// See: https://developer.apple.com/documentation/kernel/1557164-erfcl
func Erfcl(arg0 float64) float64 {
	result, callErr := tryErfcl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erff func(arg0 float32) float32
var _erffErr error

func tryErff(arg0 float32) (float32, error) {
	if _erff == nil {
		return 0.0, symbolCallError("erff", "10.10", _erffErr)
	}
	return _erff(arg0), nil
}

// Erff.
//
// See: https://developer.apple.com/documentation/kernel/1557366-erff
func Erff(arg0 float32) float32 {
	result, callErr := tryErff(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _erfl func(arg0 float64) float64
var _erflErr error

func tryErfl(arg0 float64) (float64, error) {
	if _erfl == nil {
		return 0.0, symbolCallError("erfl", "10.10", _erflErr)
	}
	return _erfl(arg0), nil
}

// Erfl.
//
// See: https://developer.apple.com/documentation/kernel/1557285-erfl
func Erfl(arg0 float64) float64 {
	result, callErr := tryErfl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _etap_trace_thread func(target_act Thread_act_t, trace_status Boolean_t) Kern_return_t
var _etap_trace_threadErr error

func tryEtap_trace_thread(target_act Thread_act_t, trace_status Boolean_t) (Kern_return_t, error) {
	if _etap_trace_thread == nil {
		return *new(Kern_return_t), symbolCallError("etap_trace_thread", "10.0", _etap_trace_threadErr)
	}
	return _etap_trace_thread(target_act, trace_status), nil
}

// Etap_trace_thread.
//
// See: https://developer.apple.com/documentation/kernel/1418858-etap_trace_thread
func Etap_trace_thread(target_act Thread_act_t, trace_status Boolean_t) Kern_return_t {
	result, callErr := tryEtap_trace_thread(target_act, trace_status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exc_server func(InHeadP unsafe.Pointer, OutHeadP unsafe.Pointer) Boolean_t
var _exc_serverErr error

func tryExc_server(InHeadP unsafe.Pointer, OutHeadP unsafe.Pointer) (Boolean_t, error) {
	if _exc_server == nil {
		return *new(Boolean_t), symbolCallError("exc_server", "10.0", _exc_serverErr)
	}
	return _exc_server(InHeadP, OutHeadP), nil
}

// Exc_server.
//
// See: https://developer.apple.com/documentation/kernel/1537285-exc_server
func Exc_server(InHeadP unsafe.Pointer, OutHeadP unsafe.Pointer) Boolean_t {
	result, callErr := tryExc_server(InHeadP, OutHeadP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exc_server_routine func(InHeadP unsafe.Pointer) Mig_routine_t
var _exc_server_routineErr error

func tryExc_server_routine(InHeadP unsafe.Pointer) (Mig_routine_t, error) {
	if _exc_server_routine == nil {
		return *new(Mig_routine_t), symbolCallError("exc_server_routine", "10.0", _exc_server_routineErr)
	}
	return _exc_server_routine(InHeadP), nil
}

// Exc_server_routine.
//
// See: https://developer.apple.com/documentation/kernel/1537232-exc_server_routine
func Exc_server_routine(InHeadP unsafe.Pointer) Mig_routine_t {
	result, callErr := tryExc_server_routine(InHeadP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exp func(arg0 float64) float64
var _expErr error

func tryExp(arg0 float64) (float64, error) {
	if _exp == nil {
		return 0.0, symbolCallError("exp", "10.10", _expErr)
	}
	return _exp(arg0), nil
}

// Exp.
//
// See: https://developer.apple.com/documentation/kernel/1557217-exp
func Exp(arg0 float64) float64 {
	result, callErr := tryExp(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exp2 func(arg0 float64) float64
var _exp2Err error

func tryExp2(arg0 float64) (float64, error) {
	if _exp2 == nil {
		return 0.0, symbolCallError("exp2", "10.10", _exp2Err)
	}
	return _exp2(arg0), nil
}

// Exp2.
//
// See: https://developer.apple.com/documentation/kernel/1557304-exp2
func Exp2(arg0 float64) float64 {
	result, callErr := tryExp2(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exp2f func(arg0 float32) float32
var _exp2fErr error

func tryExp2f(arg0 float32) (float32, error) {
	if _exp2f == nil {
		return 0.0, symbolCallError("exp2f", "10.10", _exp2fErr)
	}
	return _exp2f(arg0), nil
}

// Exp2f.
//
// See: https://developer.apple.com/documentation/kernel/1557192-exp2f
func Exp2f(arg0 float32) float32 {
	result, callErr := tryExp2f(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _exp2l func(arg0 float64) float64
var _exp2lErr error

func tryExp2l(arg0 float64) (float64, error) {
	if _exp2l == nil {
		return 0.0, symbolCallError("exp2l", "10.10", _exp2lErr)
	}
	return _exp2l(arg0), nil
}

// Exp2l.
//
// See: https://developer.apple.com/documentation/kernel/1557194-exp2l
func Exp2l(arg0 float64) float64 {
	result, callErr := tryExp2l(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _expf func(arg0 float32) float32
var _expfErr error

func tryExpf(arg0 float32) (float32, error) {
	if _expf == nil {
		return 0.0, symbolCallError("expf", "10.9", _expfErr)
	}
	return _expf(arg0), nil
}

// Expf.
//
// See: https://developer.apple.com/documentation/kernel/1532210-expf
func Expf(arg0 float32) float32 {
	result, callErr := tryExpf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _expl func(arg0 float64) float64
var _explErr error

func tryExpl(arg0 float64) (float64, error) {
	if _expl == nil {
		return 0.0, symbolCallError("expl", "10.10", _explErr)
	}
	return _expl(arg0), nil
}

// Expl.
//
// See: https://developer.apple.com/documentation/kernel/1557224-expl
func Expl(arg0 float64) float64 {
	result, callErr := tryExpl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _expm1 func(arg0 float64) float64
var _expm1Err error

func tryExpm1(arg0 float64) (float64, error) {
	if _expm1 == nil {
		return 0.0, symbolCallError("expm1", "10.10", _expm1Err)
	}
	return _expm1(arg0), nil
}

// Expm1.
//
// See: https://developer.apple.com/documentation/kernel/1557227-expm1
func Expm1(arg0 float64) float64 {
	result, callErr := tryExpm1(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _expm1f func(arg0 float32) float32
var _expm1fErr error

func tryExpm1f(arg0 float32) (float32, error) {
	if _expm1f == nil {
		return 0.0, symbolCallError("expm1f", "10.10", _expm1fErr)
	}
	return _expm1f(arg0), nil
}

// Expm1f.
//
// See: https://developer.apple.com/documentation/kernel/1557179-expm1f
func Expm1f(arg0 float32) float32 {
	result, callErr := tryExpm1f(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _expm1l func(arg0 float64) float64
var _expm1lErr error

func tryExpm1l(arg0 float64) (float64, error) {
	if _expm1l == nil {
		return 0.0, symbolCallError("expm1l", "10.10", _expm1lErr)
	}
	return _expm1l(arg0), nil
}

// Expm1l.
//
// See: https://developer.apple.com/documentation/kernel/1557178-expm1l
func Expm1l(arg0 float64) float64 {
	result, callErr := tryExpm1l(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fabs func(arg0 float64) float64
var _fabsErr error

func tryFabs(arg0 float64) (float64, error) {
	if _fabs == nil {
		return 0.0, symbolCallError("fabs", "10.10", _fabsErr)
	}
	return _fabs(arg0), nil
}

// Fabs.
//
// See: https://developer.apple.com/documentation/kernel/1557277-fabs
func Fabs(arg0 float64) float64 {
	result, callErr := tryFabs(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fabsf func(arg0 float32) float32
var _fabsfErr error

func tryFabsf(arg0 float32) (float32, error) {
	if _fabsf == nil {
		return 0.0, symbolCallError("fabsf", "10.10", _fabsfErr)
	}
	return _fabsf(arg0), nil
}

// Fabsf.
//
// See: https://developer.apple.com/documentation/kernel/1557291-fabsf
func Fabsf(arg0 float32) float32 {
	result, callErr := tryFabsf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fabsl func(arg0 float64) float64
var _fabslErr error

func tryFabsl(arg0 float64) (float64, error) {
	if _fabsl == nil {
		return 0.0, symbolCallError("fabsl", "10.10", _fabslErr)
	}
	return _fabsl(arg0), nil
}

// Fabsl.
//
// See: https://developer.apple.com/documentation/kernel/1557341-fabsl
func Fabsl(arg0 float64) float64 {
	result, callErr := tryFabsl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fdim func(arg0 float64, arg1 float64) float64
var _fdimErr error

func tryFdim(arg0 float64, arg1 float64) (float64, error) {
	if _fdim == nil {
		return 0.0, symbolCallError("fdim", "10.10", _fdimErr)
	}
	return _fdim(arg0, arg1), nil
}

// Fdim.
//
// See: https://developer.apple.com/documentation/kernel/1557355-fdim
func Fdim(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFdim(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fdimf func(arg0 float32, arg1 float32) float32
var _fdimfErr error

func tryFdimf(arg0 float32, arg1 float32) (float32, error) {
	if _fdimf == nil {
		return 0.0, symbolCallError("fdimf", "10.10", _fdimfErr)
	}
	return _fdimf(arg0, arg1), nil
}

// Fdimf.
//
// See: https://developer.apple.com/documentation/kernel/1557210-fdimf
func Fdimf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryFdimf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fdiml func(arg0 float64, arg1 float64) float64
var _fdimlErr error

func tryFdiml(arg0 float64, arg1 float64) (float64, error) {
	if _fdiml == nil {
		return 0.0, symbolCallError("fdiml", "10.10", _fdimlErr)
	}
	return _fdiml(arg0, arg1), nil
}

// Fdiml.
//
// See: https://developer.apple.com/documentation/kernel/1557241-fdiml
func Fdiml(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFdiml(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ffs func(arg0 uint32) int32
var _ffsErr error

func tryFfs(arg0 uint32) (int32, error) {
	if _ffs == nil {
		return 0, symbolCallError("ffs", "10.0", _ffsErr)
	}
	return _ffs(arg0), nil
}

// Ffs.
//
// See: https://developer.apple.com/documentation/kernel/1441046-ffs
func Ffs(arg0 uint32) int32 {
	result, callErr := tryFfs(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ffsll func(arg0 uint64) int32
var _ffsllErr error

func tryFfsll(arg0 uint64) (int32, error) {
	if _ffsll == nil {
		return 0, symbolCallError("ffsll", "10.13", _ffsllErr)
	}
	return _ffsll(arg0), nil
}

// Ffsll.
//
// See: https://developer.apple.com/documentation/kernel/2869615-ffsll
func Ffsll(arg0 uint64) int32 {
	result, callErr := tryFfsll(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _floor func(arg0 float64) float64
var _floorErr error

func tryFloor(arg0 float64) (float64, error) {
	if _floor == nil {
		return 0.0, symbolCallError("floor", "10.10", _floorErr)
	}
	return _floor(arg0), nil
}

// Floor.
//
// See: https://developer.apple.com/documentation/kernel/1557338-floor
func Floor(arg0 float64) float64 {
	result, callErr := tryFloor(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _floorf func(arg0 float32) float32
var _floorfErr error

func tryFloorf(arg0 float32) (float32, error) {
	if _floorf == nil {
		return 0.0, symbolCallError("floorf", "10.10", _floorfErr)
	}
	return _floorf(arg0), nil
}

// Floorf.
//
// See: https://developer.apple.com/documentation/kernel/1557176-floorf
func Floorf(arg0 float32) float32 {
	result, callErr := tryFloorf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _floorl func(arg0 float64) float64
var _floorlErr error

func tryFloorl(arg0 float64) (float64, error) {
	if _floorl == nil {
		return 0.0, symbolCallError("floorl", "10.10", _floorlErr)
	}
	return _floorl(arg0), nil
}

// Floorl.
//
// See: https://developer.apple.com/documentation/kernel/1557330-floorl
func Floorl(arg0 float64) float64 {
	result, callErr := tryFloorl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fls func(arg0 uint32) int32
var _flsErr error

func tryFls(arg0 uint32) (int32, error) {
	if _fls == nil {
		return 0, symbolCallError("fls", "10.13", _flsErr)
	}
	return _fls(arg0), nil
}

// Fls.
//
// See: https://developer.apple.com/documentation/kernel/2869617-fls
func Fls(arg0 uint32) int32 {
	result, callErr := tryFls(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _flsll func(arg0 uint64) int32
var _flsllErr error

func tryFlsll(arg0 uint64) (int32, error) {
	if _flsll == nil {
		return 0, symbolCallError("flsll", "10.13", _flsllErr)
	}
	return _flsll(arg0), nil
}

// Flsll.
//
// See: https://developer.apple.com/documentation/kernel/2869614-flsll
func Flsll(arg0 uint64) int32 {
	result, callErr := tryFlsll(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fma func(arg0 float64, arg1 float64, arg2 float64) float64
var _fmaErr error

func tryFma(arg0 float64, arg1 float64, arg2 float64) (float64, error) {
	if _fma == nil {
		return 0.0, symbolCallError("fma", "10.10", _fmaErr)
	}
	return _fma(arg0, arg1, arg2), nil
}

// Fma.
//
// See: https://developer.apple.com/documentation/kernel/1557233-fma
func Fma(arg0 float64, arg1 float64, arg2 float64) float64 {
	result, callErr := tryFma(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmaf func(arg0 float32, arg1 float32, arg2 float32) float32
var _fmafErr error

func tryFmaf(arg0 float32, arg1 float32, arg2 float32) (float32, error) {
	if _fmaf == nil {
		return 0.0, symbolCallError("fmaf", "10.10", _fmafErr)
	}
	return _fmaf(arg0, arg1, arg2), nil
}

// Fmaf.
//
// See: https://developer.apple.com/documentation/kernel/1557358-fmaf
func Fmaf(arg0 float32, arg1 float32, arg2 float32) float32 {
	result, callErr := tryFmaf(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmal func(arg0 float64, arg1 float64, arg2 float64) float64
var _fmalErr error

func tryFmal(arg0 float64, arg1 float64, arg2 float64) (float64, error) {
	if _fmal == nil {
		return 0.0, symbolCallError("fmal", "10.10", _fmalErr)
	}
	return _fmal(arg0, arg1, arg2), nil
}

// Fmal.
//
// See: https://developer.apple.com/documentation/kernel/1557206-fmal
func Fmal(arg0 float64, arg1 float64, arg2 float64) float64 {
	result, callErr := tryFmal(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmax func(arg0 float64, arg1 float64) float64
var _fmaxErr error

func tryFmax(arg0 float64, arg1 float64) (float64, error) {
	if _fmax == nil {
		return 0.0, symbolCallError("fmax", "10.10", _fmaxErr)
	}
	return _fmax(arg0, arg1), nil
}

// Fmax.
//
// See: https://developer.apple.com/documentation/kernel/1557201-fmax
func Fmax(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFmax(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmaxf func(arg0 float32, arg1 float32) float32
var _fmaxfErr error

func tryFmaxf(arg0 float32, arg1 float32) (float32, error) {
	if _fmaxf == nil {
		return 0.0, symbolCallError("fmaxf", "10.10", _fmaxfErr)
	}
	return _fmaxf(arg0, arg1), nil
}

// Fmaxf.
//
// See: https://developer.apple.com/documentation/kernel/1557268-fmaxf
func Fmaxf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryFmaxf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmaxl func(arg0 float64, arg1 float64) float64
var _fmaxlErr error

func tryFmaxl(arg0 float64, arg1 float64) (float64, error) {
	if _fmaxl == nil {
		return 0.0, symbolCallError("fmaxl", "10.10", _fmaxlErr)
	}
	return _fmaxl(arg0, arg1), nil
}

// Fmaxl.
//
// See: https://developer.apple.com/documentation/kernel/1557200-fmaxl
func Fmaxl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFmaxl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmin func(arg0 float64, arg1 float64) float64
var _fminErr error

func tryFmin(arg0 float64, arg1 float64) (float64, error) {
	if _fmin == nil {
		return 0.0, symbolCallError("fmin", "10.10", _fminErr)
	}
	return _fmin(arg0, arg1), nil
}

// Fmin.
//
// See: https://developer.apple.com/documentation/kernel/1557189-fmin
func Fmin(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFmin(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fminf func(arg0 float32, arg1 float32) float32
var _fminfErr error

func tryFminf(arg0 float32, arg1 float32) (float32, error) {
	if _fminf == nil {
		return 0.0, symbolCallError("fminf", "10.10", _fminfErr)
	}
	return _fminf(arg0, arg1), nil
}

// Fminf.
//
// See: https://developer.apple.com/documentation/kernel/1557340-fminf
func Fminf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryFminf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fminl func(arg0 float64, arg1 float64) float64
var _fminlErr error

func tryFminl(arg0 float64, arg1 float64) (float64, error) {
	if _fminl == nil {
		return 0.0, symbolCallError("fminl", "10.10", _fminlErr)
	}
	return _fminl(arg0, arg1), nil
}

// Fminl.
//
// See: https://developer.apple.com/documentation/kernel/1557215-fminl
func Fminl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFminl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmod func(arg0 float64, arg1 float64) float64
var _fmodErr error

func tryFmod(arg0 float64, arg1 float64) (float64, error) {
	if _fmod == nil {
		return 0.0, symbolCallError("fmod", "10.10", _fmodErr)
	}
	return _fmod(arg0, arg1), nil
}

// Fmod.
//
// See: https://developer.apple.com/documentation/kernel/1557253-fmod
func Fmod(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFmod(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmodf func(arg0 float32, arg1 float32) float32
var _fmodfErr error

func tryFmodf(arg0 float32, arg1 float32) (float32, error) {
	if _fmodf == nil {
		return 0.0, symbolCallError("fmodf", "10.10", _fmodfErr)
	}
	return _fmodf(arg0, arg1), nil
}

// Fmodf.
//
// See: https://developer.apple.com/documentation/kernel/1557354-fmodf
func Fmodf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryFmodf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _fmodl func(arg0 float64, arg1 float64) float64
var _fmodlErr error

func tryFmodl(arg0 float64, arg1 float64) (float64, error) {
	if _fmodl == nil {
		return 0.0, symbolCallError("fmodl", "10.10", _fmodlErr)
	}
	return _fmodl(arg0, arg1), nil
}

// Fmodl.
//
// See: https://developer.apple.com/documentation/kernel/1557237-fmodl
func Fmodl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryFmodl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _frexp func(arg0 float64, arg1 *int32) float64
var _frexpErr error

func tryFrexp(arg0 float64, arg1 *int32) (float64, error) {
	if _frexp == nil {
		return 0.0, symbolCallError("frexp", "10.10", _frexpErr)
	}
	return _frexp(arg0, arg1), nil
}

// Frexp.
//
// See: https://developer.apple.com/documentation/kernel/1557221-frexp
func Frexp(arg0 float64, arg1 *int32) float64 {
	result, callErr := tryFrexp(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _frexpf func(arg0 float32, arg1 *int32) float32
var _frexpfErr error

func tryFrexpf(arg0 float32, arg1 *int32) (float32, error) {
	if _frexpf == nil {
		return 0.0, symbolCallError("frexpf", "10.10", _frexpfErr)
	}
	return _frexpf(arg0, arg1), nil
}

// Frexpf.
//
// See: https://developer.apple.com/documentation/kernel/1557321-frexpf
func Frexpf(arg0 float32, arg1 *int32) float32 {
	result, callErr := tryFrexpf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _frexpl func(arg0 float64, arg1 *int32) float64
var _frexplErr error

func tryFrexpl(arg0 float64, arg1 *int32) (float64, error) {
	if _frexpl == nil {
		return 0.0, symbolCallError("frexpl", "10.10", _frexplErr)
	}
	return _frexpl(arg0, arg1), nil
}

// Frexpl.
//
// See: https://developer.apple.com/documentation/kernel/1557175-frexpl
func Frexpl(arg0 float64, arg1 *int32) float64 {
	result, callErr := tryFrexpl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_create_mach_voucher func(host Host_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt Mach_msg_type_number_t, voucher *Ipc_voucher_t) Kern_return_t
var _host_create_mach_voucherErr error

func tryHost_create_mach_voucher(host Host_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt Mach_msg_type_number_t, voucher *Ipc_voucher_t) (Kern_return_t, error) {
	if _host_create_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("host_create_mach_voucher", "10.10", _host_create_mach_voucherErr)
	}
	return _host_create_mach_voucher(host, recipes, recipesCnt, voucher), nil
}

// Host_create_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1502476-host_create_mach_voucher
func Host_create_mach_voucher(host Host_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt Mach_msg_type_number_t, voucher *Ipc_voucher_t) Kern_return_t {
	result, callErr := tryHost_create_mach_voucher(host, recipes, recipesCnt, voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_default_memory_manager func(host_priv Host_priv_t, default_manager *Memory_object_default_t, cluster_size Memory_object_cluster_size_t) Kern_return_t
var _host_default_memory_managerErr error

func tryHost_default_memory_manager(host_priv Host_priv_t, default_manager *Memory_object_default_t, cluster_size Memory_object_cluster_size_t) (Kern_return_t, error) {
	if _host_default_memory_manager == nil {
		return *new(Kern_return_t), symbolCallError("host_default_memory_manager", "10.0", _host_default_memory_managerErr)
	}
	return _host_default_memory_manager(host_priv, default_manager, cluster_size), nil
}

// Host_default_memory_manager.
//
// See: https://developer.apple.com/documentation/kernel/1588899-host_default_memory_manager
func Host_default_memory_manager(host_priv Host_priv_t, default_manager *Memory_object_default_t, cluster_size Memory_object_cluster_size_t) Kern_return_t {
	result, callErr := tryHost_default_memory_manager(host_priv, default_manager, cluster_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_UNDServer func(host Host_priv_t, server *UNDServerRef) Kern_return_t
var _host_get_UNDServerErr error

func tryHost_get_UNDServer(host Host_priv_t, server *UNDServerRef) (Kern_return_t, error) {
	if _host_get_UNDServer == nil {
		return *new(Kern_return_t), symbolCallError("host_get_UNDServer", "10.0", _host_get_UNDServerErr)
	}
	return _host_get_UNDServer(host, server), nil
}

// Host_get_UNDServer.
//
// See: https://developer.apple.com/documentation/kernel/1588902-host_get_undserver
func Host_get_UNDServer(host Host_priv_t, server *UNDServerRef) Kern_return_t {
	result, callErr := tryHost_get_UNDServer(host, server)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_boot_info func(host_priv Host_priv_t, boot_info uintptr) Kern_return_t
var _host_get_boot_infoErr error

func tryHost_get_boot_info(host_priv Host_priv_t, boot_info uintptr) (Kern_return_t, error) {
	if _host_get_boot_info == nil {
		return *new(Kern_return_t), symbolCallError("host_get_boot_info", "10.0", _host_get_boot_infoErr)
	}
	return _host_get_boot_info(host_priv, boot_info), nil
}

// Host_get_boot_info.
//
// See: https://developer.apple.com/documentation/kernel/1588870-host_get_boot_info
func Host_get_boot_info(host_priv Host_priv_t, boot_info uintptr) Kern_return_t {
	result, callErr := tryHost_get_boot_info(host_priv, boot_info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_clock_control func(host_priv Host_priv_t, clock_id Clock_id_t, clock_ctrl *Clock_ctrl_t) Kern_return_t
var _host_get_clock_controlErr error

func tryHost_get_clock_control(host_priv Host_priv_t, clock_id Clock_id_t, clock_ctrl *Clock_ctrl_t) (Kern_return_t, error) {
	if _host_get_clock_control == nil {
		return *new(Kern_return_t), symbolCallError("host_get_clock_control", "10.0", _host_get_clock_controlErr)
	}
	return _host_get_clock_control(host_priv, clock_id, clock_ctrl), nil
}

// Host_get_clock_control.
//
// See: https://developer.apple.com/documentation/kernel/1588791-host_get_clock_control
func Host_get_clock_control(host_priv Host_priv_t, clock_id Clock_id_t, clock_ctrl *Clock_ctrl_t) Kern_return_t {
	result, callErr := tryHost_get_clock_control(host_priv, clock_id, clock_ctrl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_clock_service func(host Host_t, clock_id Clock_id_t, clock_serv *Clock_serv_t) Kern_return_t
var _host_get_clock_serviceErr error

func tryHost_get_clock_service(host Host_t, clock_id Clock_id_t, clock_serv *Clock_serv_t) (Kern_return_t, error) {
	if _host_get_clock_service == nil {
		return *new(Kern_return_t), symbolCallError("host_get_clock_service", "10.0", _host_get_clock_serviceErr)
	}
	return _host_get_clock_service(host, clock_id, clock_serv), nil
}

// Host_get_clock_service.
//
// See: https://developer.apple.com/documentation/kernel/1502796-host_get_clock_service
func Host_get_clock_service(host Host_t, clock_id Clock_id_t, clock_serv *Clock_serv_t) Kern_return_t {
	result, callErr := tryHost_get_clock_service(host, clock_id, clock_serv)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_exception_ports func(host_priv Host_priv_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _host_get_exception_portsErr error

func tryHost_get_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _host_get_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("host_get_exception_ports", "10.0", _host_get_exception_portsErr)
	}
	return _host_get_exception_ports(host_priv, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors), nil
}

// Host_get_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1588769-host_get_exception_ports
func Host_get_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryHost_get_exception_ports(host_priv, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_io_main func(host Host_t, io_main *Io_main_t) Kern_return_t
var _host_get_io_mainErr error

func tryHost_get_io_main(host Host_t, io_main *Io_main_t) (Kern_return_t, error) {
	if _host_get_io_main == nil {
		return *new(Kern_return_t), symbolCallError("host_get_io_main", "13.0", _host_get_io_mainErr)
	}
	return _host_get_io_main(host, io_main), nil
}

// Host_get_io_main.
//
// See: https://developer.apple.com/documentation/kernel/4013522-host_get_io_main
func Host_get_io_main(host Host_t, io_main *Io_main_t) Kern_return_t {
	result, callErr := tryHost_get_io_main(host, io_main)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_get_special_port func(host_priv Host_priv_t, node int32, which int32, port *uint32) Kern_return_t
var _host_get_special_portErr error

func tryHost_get_special_port(host_priv Host_priv_t, node int32, which int32, port *uint32) (Kern_return_t, error) {
	if _host_get_special_port == nil {
		return *new(Kern_return_t), symbolCallError("host_get_special_port", "10.0", _host_get_special_portErr)
	}
	return _host_get_special_port(host_priv, node, which, port), nil
}

// Host_get_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1589013-host_get_special_port
func Host_get_special_port(host_priv Host_priv_t, node int32, which int32, port *uint32) Kern_return_t {
	result, callErr := tryHost_get_special_port(host_priv, node, which, port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_info func(host Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _host_infoErr error

func tryHost_info(host Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_info == nil {
		return *new(Kern_return_t), symbolCallError("host_info", "10.0", _host_infoErr)
	}
	return _host_info(host, flavor, host_info_out, host_info_outCnt), nil
}

// Host_info.
//
// See: https://developer.apple.com/documentation/kernel/1502514-host_info
func Host_info(host Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_info(host, flavor, host_info_out, host_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_kernel_version func(host Host_t, kernel_version uintptr) Kern_return_t
var _host_kernel_versionErr error

func tryHost_kernel_version(host Host_t, kernel_version uintptr) (Kern_return_t, error) {
	if _host_kernel_version == nil {
		return *new(Kern_return_t), symbolCallError("host_kernel_version", "10.0", _host_kernel_versionErr)
	}
	return _host_kernel_version(host, kernel_version), nil
}

// Host_kernel_version.
//
// See: https://developer.apple.com/documentation/kernel/1502557-host_kernel_version
func Host_kernel_version(host Host_t, kernel_version uintptr) Kern_return_t {
	result, callErr := tryHost_kernel_version(host, kernel_version)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_lockgroup_info func(host Host_t, lockgroup_info *Lockgroup_info_array_t, lockgroup_infoCnt *Mach_msg_type_number_t) Kern_return_t
var _host_lockgroup_infoErr error

func tryHost_lockgroup_info(host Host_t, lockgroup_info *Lockgroup_info_array_t, lockgroup_infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_lockgroup_info == nil {
		return *new(Kern_return_t), symbolCallError("host_lockgroup_info", "10.4", _host_lockgroup_infoErr)
	}
	return _host_lockgroup_info(host, lockgroup_info, lockgroup_infoCnt), nil
}

// Host_lockgroup_info.
//
// See: https://developer.apple.com/documentation/kernel/1502519-host_lockgroup_info
func Host_lockgroup_info(host Host_t, lockgroup_info *Lockgroup_info_array_t, lockgroup_infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_lockgroup_info(host, lockgroup_info, lockgroup_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_page_size func(host Host_t, out_page_size *Vm_size_t) Kern_return_t
var _host_page_sizeErr error

func tryHost_page_size(host Host_t, out_page_size *Vm_size_t) (Kern_return_t, error) {
	if _host_page_size == nil {
		return *new(Kern_return_t), symbolCallError("host_page_size", "10.0", _host_page_sizeErr)
	}
	return _host_page_size(host, out_page_size), nil
}

// Host_page_size.
//
// See: https://developer.apple.com/documentation/kernel/1502512-host_page_size
func Host_page_size(host Host_t, out_page_size *Vm_size_t) Kern_return_t {
	result, callErr := tryHost_page_size(host, out_page_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_priv_statistics func(host_priv Host_priv_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _host_priv_statisticsErr error

func tryHost_priv_statistics(host_priv Host_priv_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_priv_statistics == nil {
		return *new(Kern_return_t), symbolCallError("host_priv_statistics", "10.0", _host_priv_statisticsErr)
	}
	return _host_priv_statistics(host_priv, flavor, host_info_out, host_info_outCnt), nil
}

// Host_priv_statistics.
//
// See: https://developer.apple.com/documentation/kernel/1588923-host_priv_statistics
func Host_priv_statistics(host_priv Host_priv_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_priv_statistics(host_priv, flavor, host_info_out, host_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_processor_info func(host Host_t, flavor Processor_flavor_t, out_processor_count *Natural_t, out_processor_info *Processor_info_array_t, out_processor_infoCnt *Mach_msg_type_number_t) Kern_return_t
var _host_processor_infoErr error

func tryHost_processor_info(host Host_t, flavor Processor_flavor_t, out_processor_count *Natural_t, out_processor_info *Processor_info_array_t, out_processor_infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_processor_info == nil {
		return *new(Kern_return_t), symbolCallError("host_processor_info", "10.0", _host_processor_infoErr)
	}
	return _host_processor_info(host, flavor, out_processor_count, out_processor_info, out_processor_infoCnt), nil
}

// Host_processor_info.
//
// See: https://developer.apple.com/documentation/kernel/1502854-host_processor_info
func Host_processor_info(host Host_t, flavor Processor_flavor_t, out_processor_count *Natural_t, out_processor_info *Processor_info_array_t, out_processor_infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_processor_info(host, flavor, out_processor_count, out_processor_info, out_processor_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_processor_set_priv func(host_priv Host_priv_t, set_name Processor_set_name_t, set *Processor_set_t) Kern_return_t
var _host_processor_set_privErr error

func tryHost_processor_set_priv(host_priv Host_priv_t, set_name Processor_set_name_t, set *Processor_set_t) (Kern_return_t, error) {
	if _host_processor_set_priv == nil {
		return *new(Kern_return_t), symbolCallError("host_processor_set_priv", "10.0", _host_processor_set_privErr)
	}
	return _host_processor_set_priv(host_priv, set_name, set), nil
}

// Host_processor_set_priv.
//
// See: https://developer.apple.com/documentation/kernel/1588828-host_processor_set_priv
func Host_processor_set_priv(host_priv Host_priv_t, set_name Processor_set_name_t, set *Processor_set_t) Kern_return_t {
	result, callErr := tryHost_processor_set_priv(host_priv, set_name, set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_processor_sets func(host_priv Host_priv_t, processor_sets *Processor_set_name_array_t, processor_setsCnt *Mach_msg_type_number_t) Kern_return_t
var _host_processor_setsErr error

func tryHost_processor_sets(host_priv Host_priv_t, processor_sets *Processor_set_name_array_t, processor_setsCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_processor_sets == nil {
		return *new(Kern_return_t), symbolCallError("host_processor_sets", "10.0", _host_processor_setsErr)
	}
	return _host_processor_sets(host_priv, processor_sets, processor_setsCnt), nil
}

// Host_processor_sets.
//
// See: https://developer.apple.com/documentation/kernel/1588954-host_processor_sets
func Host_processor_sets(host_priv Host_priv_t, processor_sets *Processor_set_name_array_t, processor_setsCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_processor_sets(host_priv, processor_sets, processor_setsCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_processors func(host_priv Host_priv_t, out_processor_list *Processor_array_t, out_processor_listCnt *Mach_msg_type_number_t) Kern_return_t
var _host_processorsErr error

func tryHost_processors(host_priv Host_priv_t, out_processor_list *Processor_array_t, out_processor_listCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_processors == nil {
		return *new(Kern_return_t), symbolCallError("host_processors", "10.0", _host_processorsErr)
	}
	return _host_processors(host_priv, out_processor_list, out_processor_listCnt), nil
}

// Host_processors.
//
// See: https://developer.apple.com/documentation/kernel/1588774-host_processors
func Host_processors(host_priv Host_priv_t, out_processor_list *Processor_array_t, out_processor_listCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_processors(host_priv, out_processor_list, out_processor_listCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_reboot func(host_priv Host_priv_t, options int32) Kern_return_t
var _host_rebootErr error

func tryHost_reboot(host_priv Host_priv_t, options int32) (Kern_return_t, error) {
	if _host_reboot == nil {
		return *new(Kern_return_t), symbolCallError("host_reboot", "10.0", _host_rebootErr)
	}
	return _host_reboot(host_priv, options), nil
}

// Host_reboot.
//
// See: https://developer.apple.com/documentation/kernel/1588900-host_reboot
func Host_reboot(host_priv Host_priv_t, options int32) Kern_return_t {
	result, callErr := tryHost_reboot(host_priv, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_register_mach_voucher_attr_manager func(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, new_key *Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) Kern_return_t
var _host_register_mach_voucher_attr_managerErr error

func tryHost_register_mach_voucher_attr_manager(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, new_key *Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) (Kern_return_t, error) {
	if _host_register_mach_voucher_attr_manager == nil {
		return *new(Kern_return_t), symbolCallError("host_register_mach_voucher_attr_manager", "10.10", _host_register_mach_voucher_attr_managerErr)
	}
	return _host_register_mach_voucher_attr_manager(host, attr_manager, default_value, new_key, new_attr_control), nil
}

// Host_register_mach_voucher_attr_manager.
//
// See: https://developer.apple.com/documentation/kernel/1502592-host_register_mach_voucher_attr_
func Host_register_mach_voucher_attr_manager(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, new_key *Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) Kern_return_t {
	result, callErr := tryHost_register_mach_voucher_attr_manager(host, attr_manager, default_value, new_key, new_attr_control)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_register_well_known_mach_voucher_attr_manager func(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, key Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) Kern_return_t
var _host_register_well_known_mach_voucher_attr_managerErr error

func tryHost_register_well_known_mach_voucher_attr_manager(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, key Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) (Kern_return_t, error) {
	if _host_register_well_known_mach_voucher_attr_manager == nil {
		return *new(Kern_return_t), symbolCallError("host_register_well_known_mach_voucher_attr_manager", "10.10", _host_register_well_known_mach_voucher_attr_managerErr)
	}
	return _host_register_well_known_mach_voucher_attr_manager(host, attr_manager, default_value, key, new_attr_control), nil
}

// Host_register_well_known_mach_voucher_attr_manager.
//
// See: https://developer.apple.com/documentation/kernel/1502856-host_register_well_known_mach_vo
func Host_register_well_known_mach_voucher_attr_manager(host Host_t, attr_manager Mach_voucher_attr_manager_t, default_value Mach_voucher_attr_value_handle_t, key Mach_voucher_attr_key_t, new_attr_control *Ipc_voucher_attr_control_t) Kern_return_t {
	result, callErr := tryHost_register_well_known_mach_voucher_attr_manager(host, attr_manager, default_value, key, new_attr_control)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_request_notification func(host Host_t, notify_type Host_flavor_t, notify_port uint32) Kern_return_t
var _host_request_notificationErr error

func tryHost_request_notification(host Host_t, notify_type Host_flavor_t, notify_port uint32) (Kern_return_t, error) {
	if _host_request_notification == nil {
		return *new(Kern_return_t), symbolCallError("host_request_notification", "10.3", _host_request_notificationErr)
	}
	return _host_request_notification(host, notify_type, notify_port), nil
}

// Host_request_notification.
//
// See: https://developer.apple.com/documentation/kernel/1502649-host_request_notification
func Host_request_notification(host Host_t, notify_type Host_flavor_t, notify_port uint32) Kern_return_t {
	result, callErr := tryHost_request_notification(host, notify_type, notify_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_security_create_task_token func(host_security Host_security_t, parent_task Task_t, sec_token Security_token_t, audit_token *[32]byte, host Host_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) Kern_return_t
var _host_security_create_task_tokenErr error

func tryHost_security_create_task_token(host_security Host_security_t, parent_task Task_t, sec_token Security_token_t, audit_token [32]byte, host Host_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) (Kern_return_t, error) {
	if _host_security_create_task_token == nil {
		return *new(Kern_return_t), symbolCallError("host_security_create_task_token", "10.0", _host_security_create_task_tokenErr)
	}
	return _host_security_create_task_token(host_security, parent_task, sec_token, &audit_token, host, ledgers, ledgersCnt, inherit_memory, child_task), nil
}

// Host_security_create_task_token.
//
// See: https://developer.apple.com/documentation/kernel/1437156-host_security_create_task_token
func Host_security_create_task_token(host_security Host_security_t, parent_task Task_t, sec_token Security_token_t, audit_token [32]byte, host Host_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) Kern_return_t {
	result, callErr := tryHost_security_create_task_token(host_security, parent_task, sec_token, audit_token, host, ledgers, ledgersCnt, inherit_memory, child_task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_security_set_task_token func(host_security Host_security_t, target_task Task_t, sec_token Security_token_t, audit_token *[32]byte, host Host_t) Kern_return_t
var _host_security_set_task_tokenErr error

func tryHost_security_set_task_token(host_security Host_security_t, target_task Task_t, sec_token Security_token_t, audit_token [32]byte, host Host_t) (Kern_return_t, error) {
	if _host_security_set_task_token == nil {
		return *new(Kern_return_t), symbolCallError("host_security_set_task_token", "10.0", _host_security_set_task_tokenErr)
	}
	return _host_security_set_task_token(host_security, target_task, sec_token, &audit_token, host), nil
}

// Host_security_set_task_token.
//
// See: https://developer.apple.com/documentation/kernel/1437135-host_security_set_task_token
func Host_security_set_task_token(host_security Host_security_t, target_task Task_t, sec_token Security_token_t, audit_token [32]byte, host Host_t) Kern_return_t {
	result, callErr := tryHost_security_set_task_token(host_security, target_task, sec_token, audit_token, host)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_self func() Host_t
var _host_selfErr error

func tryHost_self() (Host_t, error) {
	if _host_self == nil {
		return *new(Host_t), symbolCallError("host_self", "10.0", _host_selfErr)
	}
	return _host_self(), nil
}

// Host_self.
//
// See: https://developer.apple.com/documentation/kernel/1580828-host_self
func Host_self() Host_t {
	result, callErr := tryHost_self()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_set_UNDServer func(host Host_priv_t, server UNDServerRef) Kern_return_t
var _host_set_UNDServerErr error

func tryHost_set_UNDServer(host Host_priv_t, server UNDServerRef) (Kern_return_t, error) {
	if _host_set_UNDServer == nil {
		return *new(Kern_return_t), symbolCallError("host_set_UNDServer", "10.0", _host_set_UNDServerErr)
	}
	return _host_set_UNDServer(host, server), nil
}

// Host_set_UNDServer.
//
// See: https://developer.apple.com/documentation/kernel/1588841-host_set_undserver
func Host_set_UNDServer(host Host_priv_t, server UNDServerRef) Kern_return_t {
	result, callErr := tryHost_set_UNDServer(host, server)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_set_atm_diagnostic_flag func(host Host_t, diagnostic_flag uint32) Kern_return_t
var _host_set_atm_diagnostic_flagErr error

func tryHost_set_atm_diagnostic_flag(host Host_t, diagnostic_flag uint32) (Kern_return_t, error) {
	if _host_set_atm_diagnostic_flag == nil {
		return *new(Kern_return_t), symbolCallError("host_set_atm_diagnostic_flag", "10.11", _host_set_atm_diagnostic_flagErr)
	}
	return _host_set_atm_diagnostic_flag(host, diagnostic_flag), nil
}

// Host_set_atm_diagnostic_flag.
//
// See: https://developer.apple.com/documentation/kernel/1502446-host_set_atm_diagnostic_flag
func Host_set_atm_diagnostic_flag(host Host_t, diagnostic_flag uint32) Kern_return_t {
	result, callErr := tryHost_set_atm_diagnostic_flag(host, diagnostic_flag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_set_exception_ports func(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t
var _host_set_exception_portsErr error

func tryHost_set_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) (Kern_return_t, error) {
	if _host_set_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("host_set_exception_ports", "10.0", _host_set_exception_portsErr)
	}
	return _host_set_exception_ports(host_priv, exception_mask, new_port, behavior, new_flavor), nil
}

// Host_set_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1588834-host_set_exception_ports
func Host_set_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t {
	result, callErr := tryHost_set_exception_ports(host_priv, exception_mask, new_port, behavior, new_flavor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_set_multiuser_config_flags func(host_priv Host_priv_t, multiuser_flags uint32) Kern_return_t
var _host_set_multiuser_config_flagsErr error

func tryHost_set_multiuser_config_flags(host_priv Host_priv_t, multiuser_flags uint32) (Kern_return_t, error) {
	if _host_set_multiuser_config_flags == nil {
		return *new(Kern_return_t), symbolCallError("host_set_multiuser_config_flags", "10.11.4", _host_set_multiuser_config_flagsErr)
	}
	return _host_set_multiuser_config_flags(host_priv, multiuser_flags), nil
}

// Host_set_multiuser_config_flags.
//
// See: https://developer.apple.com/documentation/kernel/1502470-host_set_multiuser_config_flags
func Host_set_multiuser_config_flags(host_priv Host_priv_t, multiuser_flags uint32) Kern_return_t {
	result, callErr := tryHost_set_multiuser_config_flags(host_priv, multiuser_flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_set_special_port func(host_priv Host_priv_t, which int32, port uint32) Kern_return_t
var _host_set_special_portErr error

func tryHost_set_special_port(host_priv Host_priv_t, which int32, port uint32) (Kern_return_t, error) {
	if _host_set_special_port == nil {
		return *new(Kern_return_t), symbolCallError("host_set_special_port", "10.0", _host_set_special_portErr)
	}
	return _host_set_special_port(host_priv, which, port), nil
}

// Host_set_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1588941-host_set_special_port
func Host_set_special_port(host_priv Host_priv_t, which int32, port uint32) Kern_return_t {
	result, callErr := tryHost_set_special_port(host_priv, which, port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_statistics func(host_priv Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _host_statisticsErr error

func tryHost_statistics(host_priv Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_statistics == nil {
		return *new(Kern_return_t), symbolCallError("host_statistics", "10.0", _host_statisticsErr)
	}
	return _host_statistics(host_priv, flavor, host_info_out, host_info_outCnt), nil
}

// Host_statistics.
//
// See: https://developer.apple.com/documentation/kernel/1502546-host_statistics
func Host_statistics(host_priv Host_t, flavor Host_flavor_t, host_info_out Host_info_t, host_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_statistics(host_priv, flavor, host_info_out, host_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_statistics64 func(host_priv Host_t, flavor Host_flavor_t, host_info64_out Host_info64_t, host_info64_outCnt *Mach_msg_type_number_t) Kern_return_t
var _host_statistics64Err error

func tryHost_statistics64(host_priv Host_t, flavor Host_flavor_t, host_info64_out Host_info64_t, host_info64_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_statistics64 == nil {
		return *new(Kern_return_t), symbolCallError("host_statistics64", "10.6", _host_statistics64Err)
	}
	return _host_statistics64(host_priv, flavor, host_info64_out, host_info64_outCnt), nil
}

// Host_statistics64.
//
// See: https://developer.apple.com/documentation/kernel/1502863-host_statistics64
func Host_statistics64(host_priv Host_t, flavor Host_flavor_t, host_info64_out Host_info64_t, host_info64_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_statistics64(host_priv, flavor, host_info64_out, host_info64_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_swap_exception_ports func(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlerss Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _host_swap_exception_portsErr error

func tryHost_swap_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlerss Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _host_swap_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("host_swap_exception_ports", "10.0", _host_swap_exception_portsErr)
	}
	return _host_swap_exception_ports(host_priv, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlerss, old_behaviors, old_flavors), nil
}

// Host_swap_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1588836-host_swap_exception_ports
func Host_swap_exception_ports(host_priv Host_priv_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlerss Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryHost_swap_exception_ports(host_priv, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlerss, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _host_virtual_physical_table_info func(host Host_t, info *Hash_info_bucket_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _host_virtual_physical_table_infoErr error

func tryHost_virtual_physical_table_info(host Host_t, info *Hash_info_bucket_array_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _host_virtual_physical_table_info == nil {
		return *new(Kern_return_t), symbolCallError("host_virtual_physical_table_info", "10.0", _host_virtual_physical_table_infoErr)
	}
	return _host_virtual_physical_table_info(host, info, infoCnt), nil
}

// Host_virtual_physical_table_info.
//
// See: https://developer.apple.com/documentation/kernel/1502774-host_virtual_physical_table_info
func Host_virtual_physical_table_info(host Host_t, info *Hash_info_bucket_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryHost_virtual_physical_table_info(host, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hypot func(arg0 float64, arg1 float64) float64
var _hypotErr error

func tryHypot(arg0 float64, arg1 float64) (float64, error) {
	if _hypot == nil {
		return 0.0, symbolCallError("hypot", "10.10", _hypotErr)
	}
	return _hypot(arg0, arg1), nil
}

// Hypot.
//
// See: https://developer.apple.com/documentation/kernel/1557147-hypot
func Hypot(arg0 float64, arg1 float64) float64 {
	result, callErr := tryHypot(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hypotf func(arg0 float32, arg1 float32) float32
var _hypotfErr error

func tryHypotf(arg0 float32, arg1 float32) (float32, error) {
	if _hypotf == nil {
		return 0.0, symbolCallError("hypotf", "10.10", _hypotfErr)
	}
	return _hypotf(arg0, arg1), nil
}

// Hypotf.
//
// See: https://developer.apple.com/documentation/kernel/1557254-hypotf
func Hypotf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryHypotf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hypotl func(arg0 float64, arg1 float64) float64
var _hypotlErr error

func tryHypotl(arg0 float64, arg1 float64) (float64, error) {
	if _hypotl == nil {
		return 0.0, symbolCallError("hypotl", "10.10", _hypotlErr)
	}
	return _hypotl(arg0, arg1), nil
}

// Hypotl.
//
// See: https://developer.apple.com/documentation/kernel/1557299-hypotl
func Hypotl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryHypotl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ilogb func(arg0 float64) int32
var _ilogbErr error

func tryIlogb(arg0 float64) (int32, error) {
	if _ilogb == nil {
		return 0, symbolCallError("ilogb", "10.10", _ilogbErr)
	}
	return _ilogb(arg0), nil
}

// Ilogb.
//
// See: https://developer.apple.com/documentation/kernel/1557252-ilogb
func Ilogb(arg0 float64) int32 {
	result, callErr := tryIlogb(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ilogbf func(arg0 float32) int32
var _ilogbfErr error

func tryIlogbf(arg0 float32) (int32, error) {
	if _ilogbf == nil {
		return 0, symbolCallError("ilogbf", "10.10", _ilogbfErr)
	}
	return _ilogbf(arg0), nil
}

// Ilogbf.
//
// See: https://developer.apple.com/documentation/kernel/1557293-ilogbf
func Ilogbf(arg0 float32) int32 {
	result, callErr := tryIlogbf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ilogbl func(arg0 float64) int32
var _ilogblErr error

func tryIlogbl(arg0 float64) (int32, error) {
	if _ilogbl == nil {
		return 0, symbolCallError("ilogbl", "10.10", _ilogblErr)
	}
	return _ilogbl(arg0), nil
}

// Ilogbl.
//
// See: https://developer.apple.com/documentation/kernel/1557245-ilogbl
func Ilogbl(arg0 float64) int32 {
	result, callErr := tryIlogbl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _inet_aton func(arg0 string, arg1 *In_addr) int32
var _inet_atonErr error

func tryInet_aton(arg0 string, arg1 *In_addr) (int32, error) {
	if _inet_aton == nil {
		return 0, symbolCallError("inet_aton", "10.9", _inet_atonErr)
	}
	return _inet_aton(arg0, arg1), nil
}

// Inet_aton.
//
// See: https://developer.apple.com/documentation/kernel/1475455-inet_aton
func Inet_aton(arg0 string, arg1 *In_addr) int32 {
	result, callErr := tryInet_aton(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _inet_ntop func(arg0 int32, arg1 unsafe.Pointer, arg2 string, arg3 uint32) *byte
var _inet_ntopErr error

func tryInet_ntop(arg0 int32, arg1 unsafe.Pointer, arg2 string, arg3 uint32) (*byte, error) {
	if _inet_ntop == nil {
		return nil, symbolCallError("inet_ntop", "10.4", _inet_ntopErr)
	}
	return _inet_ntop(arg0, arg1, arg2, arg3), nil
}

// Inet_ntop.
//
// See: https://developer.apple.com/documentation/kernel/1475706-inet_ntop
func Inet_ntop(arg0 int32, arg1 unsafe.Pointer, arg2 string, arg3 uint32) *byte {
	result, callErr := tryInet_ntop(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _insque func(entry unsafe.Pointer, pred unsafe.Pointer)
var _insqueErr error

func tryInsque(entry unsafe.Pointer, pred unsafe.Pointer) error {
	if _insque == nil {
		return symbolCallError("insque", "10.0", _insqueErr)
	}
	_insque(entry, pred)
	return nil
}

// Insque.
//
// See: https://developer.apple.com/documentation/kernel/1567113-insque
func Insque(entry unsafe.Pointer, pred unsafe.Pointer) {
	if callErr := tryInsque(entry, pred); callErr != nil {
		panic(callErr)
	}
}

var _j0 func(arg0 float64) float64
var _j0Err error

func tryJ0(arg0 float64) (float64, error) {
	if _j0 == nil {
		return 0.0, symbolCallError("j0", "10.0", _j0Err)
	}
	return _j0(arg0), nil
}

// J0.
//
// See: https://developer.apple.com/documentation/kernel/1557261-j0
func J0(arg0 float64) float64 {
	result, callErr := tryJ0(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _j1 func(arg0 float64) float64
var _j1Err error

func tryJ1(arg0 float64) (float64, error) {
	if _j1 == nil {
		return 0.0, symbolCallError("j1", "10.0", _j1Err)
	}
	return _j1(arg0), nil
}

// J1.
//
// See: https://developer.apple.com/documentation/kernel/1557280-j1
func J1(arg0 float64) float64 {
	result, callErr := tryJ1(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jn func(arg0 int32, arg1 float64) float64
var _jnErr error

func tryJn(arg0 int32, arg1 float64) (float64, error) {
	if _jn == nil {
		return 0.0, symbolCallError("jn", "10.0", _jnErr)
	}
	return _jn(arg0, arg1), nil
}

// Jn.
//
// See: https://developer.apple.com/documentation/kernel/1557154-jn
func Jn(arg0 int32, arg1 float64) float64 {
	result, callErr := tryJn(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kdebug_timestamp func() uint64
var _kdebug_timestampErr error

func tryKdebug_timestamp() (uint64, error) {
	if _kdebug_timestamp == nil {
		return 0, symbolCallError("kdebug_timestamp", "12.0", _kdebug_timestampErr)
	}
	return _kdebug_timestamp(), nil
}

// Kdebug_timestamp.
//
// See: https://developer.apple.com/documentation/kernel/3755391-kdebug_timestamp
func Kdebug_timestamp() uint64 {
	result, callErr := tryKdebug_timestamp()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kdebug_timestamp_from_absolute func(abstime uint64) uint64
var _kdebug_timestamp_from_absoluteErr error

func tryKdebug_timestamp_from_absolute(abstime uint64) (uint64, error) {
	if _kdebug_timestamp_from_absolute == nil {
		return 0, symbolCallError("kdebug_timestamp_from_absolute", "12.0", _kdebug_timestamp_from_absoluteErr)
	}
	return _kdebug_timestamp_from_absolute(abstime), nil
}

// Kdebug_timestamp_from_absolute.
//
// See: https://developer.apple.com/documentation/kernel/3755392-kdebug_timestamp_from_absolute
func Kdebug_timestamp_from_absolute(abstime uint64) uint64 {
	result, callErr := tryKdebug_timestamp_from_absolute(abstime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kdebug_timestamp_from_continuous func(conttime uint64) uint64
var _kdebug_timestamp_from_continuousErr error

func tryKdebug_timestamp_from_continuous(conttime uint64) (uint64, error) {
	if _kdebug_timestamp_from_continuous == nil {
		return 0, symbolCallError("kdebug_timestamp_from_continuous", "12.0", _kdebug_timestamp_from_continuousErr)
	}
	return _kdebug_timestamp_from_continuous(conttime), nil
}

// Kdebug_timestamp_from_continuous.
//
// See: https://developer.apple.com/documentation/kernel/3755393-kdebug_timestamp_from_continuous
func Kdebug_timestamp_from_continuous(conttime uint64) uint64 {
	result, callErr := tryKdebug_timestamp_from_continuous(conttime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kdebug_using_continuous_time func() bool
var _kdebug_using_continuous_timeErr error

func tryKdebug_using_continuous_time() (bool, error) {
	if _kdebug_using_continuous_time == nil {
		return false, symbolCallError("kdebug_using_continuous_time", "10.15", _kdebug_using_continuous_timeErr)
	}
	return _kdebug_using_continuous_time(), nil
}

// Kdebug_using_continuous_time.
//
// See: https://developer.apple.com/documentation/kernel/3242843-kdebug_using_continuous_time
func Kdebug_using_continuous_time() bool {
	result, callErr := tryKdebug_using_continuous_time()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kevent64 func(kq int32, changelist uintptr, nchanges int32, eventlist uintptr, nevents int32, flags uint32, timeout uintptr) int32
var _kevent64Err error

func tryKevent64(kq int32, changelist uintptr, nchanges int32, eventlist uintptr, nevents int32, flags uint32, timeout uintptr) (int32, error) {
	if _kevent64 == nil {
		return 0, symbolCallError("kevent64", "", _kevent64Err)
	}
	return _kevent64(kq, changelist, nchanges, eventlist, nevents, flags, timeout), nil
}

// Kevent64 registers changes to a kernel event queue and waits for pending events.
func Kevent64(kq int32, changelist uintptr, nchanges int32, eventlist uintptr, nevents int32, flags uint32, timeout uintptr) int32 {
	result, callErr := tryKevent64(kq, changelist, nchanges, eventlist, nevents, flags, timeout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kext_request func(host_priv Host_priv_t, user_log_flags uint32, request_data Vm_offset_t, request_dataCnt Mach_msg_type_number_t, response_data *Vm_offset_t, response_dataCnt *Mach_msg_type_number_t, log_data *Vm_offset_t, log_dataCnt *Mach_msg_type_number_t, op_result *Kern_return_t) Kern_return_t
var _kext_requestErr error

func tryKext_request(host_priv Host_priv_t, user_log_flags uint32, request_data Vm_offset_t, request_dataCnt Mach_msg_type_number_t, response_data *Vm_offset_t, response_dataCnt *Mach_msg_type_number_t, log_data *Vm_offset_t, log_dataCnt *Mach_msg_type_number_t, op_result *Kern_return_t) (Kern_return_t, error) {
	if _kext_request == nil {
		return *new(Kern_return_t), symbolCallError("kext_request", "10.6", _kext_requestErr)
	}
	return _kext_request(host_priv, user_log_flags, request_data, request_dataCnt, response_data, response_dataCnt, log_data, log_dataCnt, op_result), nil
}

// Kext_request.
//
// See: https://developer.apple.com/documentation/kernel/1588829-kext_request
func Kext_request(host_priv Host_priv_t, user_log_flags uint32, request_data Vm_offset_t, request_dataCnt Mach_msg_type_number_t, response_data *Vm_offset_t, response_dataCnt *Mach_msg_type_number_t, log_data *Vm_offset_t, log_dataCnt *Mach_msg_type_number_t, op_result *Kern_return_t) Kern_return_t {
	result, callErr := tryKext_request(host_priv, user_log_flags, request_data, request_dataCnt, response_data, response_dataCnt, log_data, log_dataCnt, op_result)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kmod_control func(host_priv Host_priv_t, module Kmod_t, flavor Kmod_control_flavor_t, data *Kmod_args_t, dataCnt *Mach_msg_type_number_t) Kern_return_t
var _kmod_controlErr error

func tryKmod_control(host_priv Host_priv_t, module Kmod_t, flavor Kmod_control_flavor_t, data *Kmod_args_t, dataCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _kmod_control == nil {
		return *new(Kern_return_t), symbolCallError("kmod_control", "10.0", _kmod_controlErr)
	}
	return _kmod_control(host_priv, module, flavor, data, dataCnt), nil
}

// Kmod_control.
//
// See: https://developer.apple.com/documentation/kernel/1588743-kmod_control
func Kmod_control(host_priv Host_priv_t, module Kmod_t, flavor Kmod_control_flavor_t, data *Kmod_args_t, dataCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryKmod_control(host_priv, module, flavor, data, dataCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kmod_create func(host_priv Host_priv_t, info Vm_address_t, module *Kmod_t) Kern_return_t
var _kmod_createErr error

func tryKmod_create(host_priv Host_priv_t, info Vm_address_t, module *Kmod_t) (Kern_return_t, error) {
	if _kmod_create == nil {
		return *new(Kern_return_t), symbolCallError("kmod_create", "10.0", _kmod_createErr)
	}
	return _kmod_create(host_priv, info, module), nil
}

// Kmod_create.
//
// See: https://developer.apple.com/documentation/kernel/1588848-kmod_create
func Kmod_create(host_priv Host_priv_t, info Vm_address_t, module *Kmod_t) Kern_return_t {
	result, callErr := tryKmod_create(host_priv, info, module)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kmod_destroy func(host_priv Host_priv_t, module Kmod_t) Kern_return_t
var _kmod_destroyErr error

func tryKmod_destroy(host_priv Host_priv_t, module Kmod_t) (Kern_return_t, error) {
	if _kmod_destroy == nil {
		return *new(Kern_return_t), symbolCallError("kmod_destroy", "10.0", _kmod_destroyErr)
	}
	return _kmod_destroy(host_priv, module), nil
}

// Kmod_destroy.
//
// See: https://developer.apple.com/documentation/kernel/1588961-kmod_destroy
func Kmod_destroy(host_priv Host_priv_t, module Kmod_t) Kern_return_t {
	result, callErr := tryKmod_destroy(host_priv, module)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kmod_get_info func(host Host_t, modules *Kmod_args_t, modulesCnt *Mach_msg_type_number_t) Kern_return_t
var _kmod_get_infoErr error

func tryKmod_get_info(host Host_t, modules *Kmod_args_t, modulesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _kmod_get_info == nil {
		return *new(Kern_return_t), symbolCallError("kmod_get_info", "10.0", _kmod_get_infoErr)
	}
	return _kmod_get_info(host, modules, modulesCnt), nil
}

// Kmod_get_info.
//
// See: https://developer.apple.com/documentation/kernel/1502847-kmod_get_info
func Kmod_get_info(host Host_t, modules *Kmod_args_t, modulesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryKmod_get_info(host, modules, modulesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _kqueue func() int32
var _kqueueErr error

func tryKqueue() (int32, error) {
	if _kqueue == nil {
		return 0, symbolCallError("kqueue", "", _kqueueErr)
	}
	return _kqueue(), nil
}

// Kqueue creates a kernel event queue and returns a descriptor for it.
func Kqueue() int32 {
	result, callErr := tryKqueue()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ldexp func(arg0 float64, arg1 int32) float64
var _ldexpErr error

func tryLdexp(arg0 float64, arg1 int32) (float64, error) {
	if _ldexp == nil {
		return 0.0, symbolCallError("ldexp", "10.10", _ldexpErr)
	}
	return _ldexp(arg0, arg1), nil
}

// Ldexp.
//
// See: https://developer.apple.com/documentation/kernel/1557152-ldexp
func Ldexp(arg0 float64, arg1 int32) float64 {
	result, callErr := tryLdexp(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ldexpf func(arg0 float32, arg1 int32) float32
var _ldexpfErr error

func tryLdexpf(arg0 float32, arg1 int32) (float32, error) {
	if _ldexpf == nil {
		return 0.0, symbolCallError("ldexpf", "10.10", _ldexpfErr)
	}
	return _ldexpf(arg0, arg1), nil
}

// Ldexpf.
//
// See: https://developer.apple.com/documentation/kernel/1557190-ldexpf
func Ldexpf(arg0 float32, arg1 int32) float32 {
	result, callErr := tryLdexpf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ldexpl func(arg0 float64, arg1 int32) float64
var _ldexplErr error

func tryLdexpl(arg0 float64, arg1 int32) (float64, error) {
	if _ldexpl == nil {
		return 0.0, symbolCallError("ldexpl", "10.10", _ldexplErr)
	}
	return _ldexpl(arg0, arg1), nil
}

// Ldexpl.
//
// See: https://developer.apple.com/documentation/kernel/1557365-ldexpl
func Ldexpl(arg0 float64, arg1 int32) float64 {
	result, callErr := tryLdexpl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lgamma func(arg0 float64) float64
var _lgammaErr error

func tryLgamma(arg0 float64) (float64, error) {
	if _lgamma == nil {
		return 0.0, symbolCallError("lgamma", "10.10", _lgammaErr)
	}
	return _lgamma(arg0), nil
}

// Lgamma.
//
// See: https://developer.apple.com/documentation/kernel/1557344-lgamma
func Lgamma(arg0 float64) float64 {
	result, callErr := tryLgamma(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lgammaf func(arg0 float32) float32
var _lgammafErr error

func tryLgammaf(arg0 float32) (float32, error) {
	if _lgammaf == nil {
		return 0.0, symbolCallError("lgammaf", "10.10", _lgammafErr)
	}
	return _lgammaf(arg0), nil
}

// Lgammaf.
//
// See: https://developer.apple.com/documentation/kernel/1557160-lgammaf
func Lgammaf(arg0 float32) float32 {
	result, callErr := tryLgammaf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lgammal func(arg0 float64) float64
var _lgammalErr error

func tryLgammal(arg0 float64) (float64, error) {
	if _lgammal == nil {
		return 0.0, symbolCallError("lgammal", "10.10", _lgammalErr)
	}
	return _lgammal(arg0), nil
}

// Lgammal.
//
// See: https://developer.apple.com/documentation/kernel/1557282-lgammal
func Lgammal(arg0 float64) float64 {
	result, callErr := tryLgammal(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llrint func(arg0 float64) int64
var _llrintErr error

func tryLlrint(arg0 float64) (int64, error) {
	if _llrint == nil {
		return 0, symbolCallError("llrint", "10.10", _llrintErr)
	}
	return _llrint(arg0), nil
}

// Llrint.
//
// See: https://developer.apple.com/documentation/kernel/1557360-llrint
func Llrint(arg0 float64) int64 {
	result, callErr := tryLlrint(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llrintf func(arg0 float32) int64
var _llrintfErr error

func tryLlrintf(arg0 float32) (int64, error) {
	if _llrintf == nil {
		return 0, symbolCallError("llrintf", "10.10", _llrintfErr)
	}
	return _llrintf(arg0), nil
}

// Llrintf.
//
// See: https://developer.apple.com/documentation/kernel/1557166-llrintf
func Llrintf(arg0 float32) int64 {
	result, callErr := tryLlrintf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llrintl func(arg0 float64) int64
var _llrintlErr error

func tryLlrintl(arg0 float64) (int64, error) {
	if _llrintl == nil {
		return 0, symbolCallError("llrintl", "10.10", _llrintlErr)
	}
	return _llrintl(arg0), nil
}

// Llrintl.
//
// See: https://developer.apple.com/documentation/kernel/1557320-llrintl
func Llrintl(arg0 float64) int64 {
	result, callErr := tryLlrintl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llround func(arg0 float64) int64
var _llroundErr error

func tryLlround(arg0 float64) (int64, error) {
	if _llround == nil {
		return 0, symbolCallError("llround", "10.10", _llroundErr)
	}
	return _llround(arg0), nil
}

// Llround.
//
// See: https://developer.apple.com/documentation/kernel/1557265-llround
func Llround(arg0 float64) int64 {
	result, callErr := tryLlround(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llroundf func(arg0 float32) int64
var _llroundfErr error

func tryLlroundf(arg0 float32) (int64, error) {
	if _llroundf == nil {
		return 0, symbolCallError("llroundf", "10.10", _llroundfErr)
	}
	return _llroundf(arg0), nil
}

// Llroundf.
//
// See: https://developer.apple.com/documentation/kernel/1557367-llroundf
func Llroundf(arg0 float32) int64 {
	result, callErr := tryLlroundf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _llroundl func(arg0 float64) int64
var _llroundlErr error

func tryLlroundl(arg0 float64) (int64, error) {
	if _llroundl == nil {
		return 0, symbolCallError("llroundl", "10.10", _llroundlErr)
	}
	return _llroundl(arg0), nil
}

// Llroundl.
//
// See: https://developer.apple.com/documentation/kernel/1557335-llroundl
func Llroundl(arg0 float64) int64 {
	result, callErr := tryLlroundl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lock_set_create func(task Task_t, new_lock_set *Lock_set_t, n_ulocks int32, policy int32) Kern_return_t
var _lock_set_createErr error

func tryLock_set_create(task Task_t, new_lock_set *Lock_set_t, n_ulocks int32, policy int32) (Kern_return_t, error) {
	if _lock_set_create == nil {
		return *new(Kern_return_t), symbolCallError("lock_set_create", "10.0", _lock_set_createErr)
	}
	return _lock_set_create(task, new_lock_set, n_ulocks, policy), nil
}

// Lock_set_create.
//
// See: https://developer.apple.com/documentation/kernel/1537690-lock_set_create
func Lock_set_create(task Task_t, new_lock_set *Lock_set_t, n_ulocks int32, policy int32) Kern_return_t {
	result, callErr := tryLock_set_create(task, new_lock_set, n_ulocks, policy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lock_set_destroy func(task Task_t, lock_set Lock_set_t) Kern_return_t
var _lock_set_destroyErr error

func tryLock_set_destroy(task Task_t, lock_set Lock_set_t) (Kern_return_t, error) {
	if _lock_set_destroy == nil {
		return *new(Kern_return_t), symbolCallError("lock_set_destroy", "10.0", _lock_set_destroyErr)
	}
	return _lock_set_destroy(task, lock_set), nil
}

// Lock_set_destroy.
//
// See: https://developer.apple.com/documentation/kernel/1537941-lock_set_destroy
func Lock_set_destroy(task Task_t, lock_set Lock_set_t) Kern_return_t {
	result, callErr := tryLock_set_destroy(task, lock_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log func(arg0 int32, arg1 string)
var _logErr error

func tryLog(arg0 int32, arg1 string) error {
	if _log == nil {
		return symbolCallError("log", "10.0", _logErr)
	}
	_log(arg0, arg1)
	return nil
}

// Log.
//
// See: https://developer.apple.com/documentation/kernel/1516025-log
func Log(arg0 int32, arg1 string) {
	if callErr := tryLog(arg0, arg1); callErr != nil {
		panic(callErr)
	}
}

var _log10 func(arg0 float64) float64
var _log10Err error

func tryLog10(arg0 float64) (float64, error) {
	if _log10 == nil {
		return 0.0, symbolCallError("log10", "10.10", _log10Err)
	}
	return _log10(arg0), nil
}

// Log10.
//
// See: https://developer.apple.com/documentation/kernel/1557202-log10
func Log10(arg0 float64) float64 {
	result, callErr := tryLog10(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log10f func(arg0 float32) float32
var _log10fErr error

func tryLog10f(arg0 float32) (float32, error) {
	if _log10f == nil {
		return 0.0, symbolCallError("log10f", "10.9", _log10fErr)
	}
	return _log10f(arg0), nil
}

// Log10f.
//
// See: https://developer.apple.com/documentation/kernel/1532188-log10f
func Log10f(arg0 float32) float32 {
	result, callErr := tryLog10f(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log10l func(arg0 float64) float64
var _log10lErr error

func tryLog10l(arg0 float64) (float64, error) {
	if _log10l == nil {
		return 0.0, symbolCallError("log10l", "10.10", _log10lErr)
	}
	return _log10l(arg0), nil
}

// Log10l.
//
// See: https://developer.apple.com/documentation/kernel/1557363-log10l
func Log10l(arg0 float64) float64 {
	result, callErr := tryLog10l(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log1p func(arg0 float64) float64
var _log1pErr error

func tryLog1p(arg0 float64) (float64, error) {
	if _log1p == nil {
		return 0.0, symbolCallError("log1p", "10.10", _log1pErr)
	}
	return _log1p(arg0), nil
}

// Log1p.
//
// See: https://developer.apple.com/documentation/kernel/1557187-log1p
func Log1p(arg0 float64) float64 {
	result, callErr := tryLog1p(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log1pf func(arg0 float32) float32
var _log1pfErr error

func tryLog1pf(arg0 float32) (float32, error) {
	if _log1pf == nil {
		return 0.0, symbolCallError("log1pf", "10.10", _log1pfErr)
	}
	return _log1pf(arg0), nil
}

// Log1pf.
//
// See: https://developer.apple.com/documentation/kernel/1557337-log1pf
func Log1pf(arg0 float32) float32 {
	result, callErr := tryLog1pf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log1pl func(arg0 float64) float64
var _log1plErr error

func tryLog1pl(arg0 float64) (float64, error) {
	if _log1pl == nil {
		return 0.0, symbolCallError("log1pl", "10.10", _log1plErr)
	}
	return _log1pl(arg0), nil
}

// Log1pl.
//
// See: https://developer.apple.com/documentation/kernel/1557274-log1pl
func Log1pl(arg0 float64) float64 {
	result, callErr := tryLog1pl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log2 func(arg0 float64) float64
var _log2Err error

func tryLog2(arg0 float64) (float64, error) {
	if _log2 == nil {
		return 0.0, symbolCallError("log2", "10.10", _log2Err)
	}
	return _log2(arg0), nil
}

// Log2.
//
// See: https://developer.apple.com/documentation/kernel/1557246-log2
func Log2(arg0 float64) float64 {
	result, callErr := tryLog2(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log2f func(arg0 float32) float32
var _log2fErr error

func tryLog2f(arg0 float32) (float32, error) {
	if _log2f == nil {
		return 0.0, symbolCallError("log2f", "10.10", _log2fErr)
	}
	return _log2f(arg0), nil
}

// Log2f.
//
// See: https://developer.apple.com/documentation/kernel/1557332-log2f
func Log2f(arg0 float32) float32 {
	result, callErr := tryLog2f(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _log2l func(arg0 float64) float64
var _log2lErr error

func tryLog2l(arg0 float64) (float64, error) {
	if _log2l == nil {
		return 0.0, symbolCallError("log2l", "10.10", _log2lErr)
	}
	return _log2l(arg0), nil
}

// Log2l.
//
// See: https://developer.apple.com/documentation/kernel/1557169-log2l
func Log2l(arg0 float64) float64 {
	result, callErr := tryLog2l(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _logb func(arg0 float64) float64
var _logbErr error

func tryLogb(arg0 float64) (float64, error) {
	if _logb == nil {
		return 0.0, symbolCallError("logb", "10.10", _logbErr)
	}
	return _logb(arg0), nil
}

// Logb.
//
// See: https://developer.apple.com/documentation/kernel/1557235-logb
func Logb(arg0 float64) float64 {
	result, callErr := tryLogb(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _logbf func(arg0 float32) float32
var _logbfErr error

func tryLogbf(arg0 float32) (float32, error) {
	if _logbf == nil {
		return 0.0, symbolCallError("logbf", "10.10", _logbfErr)
	}
	return _logbf(arg0), nil
}

// Logbf.
//
// See: https://developer.apple.com/documentation/kernel/1557288-logbf
func Logbf(arg0 float32) float32 {
	result, callErr := tryLogbf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _logbl func(arg0 float64) float64
var _logblErr error

func tryLogbl(arg0 float64) (float64, error) {
	if _logbl == nil {
		return 0.0, symbolCallError("logbl", "10.10", _logblErr)
	}
	return _logbl(arg0), nil
}

// Logbl.
//
// See: https://developer.apple.com/documentation/kernel/1557168-logbl
func Logbl(arg0 float64) float64 {
	result, callErr := tryLogbl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _logf func(arg0 float32) float32
var _logfErr error

func tryLogf(arg0 float32) (float32, error) {
	if _logf == nil {
		return 0.0, symbolCallError("logf", "10.9", _logfErr)
	}
	return _logf(arg0), nil
}

// Logf.
//
// See: https://developer.apple.com/documentation/kernel/1532186-logf
func Logf(arg0 float32) float32 {
	result, callErr := tryLogf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _logl func(arg0 float64) float64
var _loglErr error

func tryLogl(arg0 float64) (float64, error) {
	if _logl == nil {
		return 0.0, symbolCallError("logl", "10.10", _loglErr)
	}
	return _logl(arg0), nil
}

// Logl.
//
// See: https://developer.apple.com/documentation/kernel/1557184-logl
func Logl(arg0 float64) float64 {
	result, callErr := tryLogl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lrint func(arg0 float64) int
var _lrintErr error

func tryLrint(arg0 float64) (int, error) {
	if _lrint == nil {
		return 0, symbolCallError("lrint", "10.10", _lrintErr)
	}
	return _lrint(arg0), nil
}

// Lrint.
//
// See: https://developer.apple.com/documentation/kernel/1557305-lrint
func Lrint(arg0 float64) int {
	result, callErr := tryLrint(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lrintf func(arg0 float32) int
var _lrintfErr error

func tryLrintf(arg0 float32) (int, error) {
	if _lrintf == nil {
		return 0, symbolCallError("lrintf", "10.10", _lrintfErr)
	}
	return _lrintf(arg0), nil
}

// Lrintf.
//
// See: https://developer.apple.com/documentation/kernel/1557142-lrintf
func Lrintf(arg0 float32) int {
	result, callErr := tryLrintf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lrintl func(arg0 float64) int
var _lrintlErr error

func tryLrintl(arg0 float64) (int, error) {
	if _lrintl == nil {
		return 0, symbolCallError("lrintl", "10.10", _lrintlErr)
	}
	return _lrintl(arg0), nil
}

// Lrintl.
//
// See: https://developer.apple.com/documentation/kernel/1557180-lrintl
func Lrintl(arg0 float64) int {
	result, callErr := tryLrintl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lround func(arg0 float64) int
var _lroundErr error

func tryLround(arg0 float64) (int, error) {
	if _lround == nil {
		return 0, symbolCallError("lround", "10.10", _lroundErr)
	}
	return _lround(arg0), nil
}

// Lround.
//
// See: https://developer.apple.com/documentation/kernel/1557329-lround
func Lround(arg0 float64) int {
	result, callErr := tryLround(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lroundf func(arg0 float32) int
var _lroundfErr error

func tryLroundf(arg0 float32) (int, error) {
	if _lroundf == nil {
		return 0, symbolCallError("lroundf", "10.10", _lroundfErr)
	}
	return _lroundf(arg0), nil
}

// Lroundf.
//
// See: https://developer.apple.com/documentation/kernel/1557186-lroundf
func Lroundf(arg0 float32) int {
	result, callErr := tryLroundf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _lroundl func(arg0 float64) int
var _lroundlErr error

func tryLroundl(arg0 float64) (int, error) {
	if _lroundl == nil {
		return 0, symbolCallError("lroundl", "10.10", _lroundlErr)
	}
	return _lroundl(arg0), nil
}

// Lroundl.
//
// See: https://developer.apple.com/documentation/kernel/1557342-lroundl
func Lroundl(arg0 float64) int {
	result, callErr := tryLroundl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_absolute_time func() uint64
var _mach_absolute_timeErr error

func tryMach_absolute_time() (uint64, error) {
	if _mach_absolute_time == nil {
		return 0, symbolCallError("mach_absolute_time", "10.0", _mach_absolute_timeErr)
	}
	return _mach_absolute_time(), nil
}

// Mach_absolute_time returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep.
//
// See: https://developer.apple.com/documentation/kernel/1462446-mach_absolute_time
func Mach_absolute_time() uint64 {
	result, callErr := tryMach_absolute_time()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_approximate_time func() uint64
var _mach_approximate_timeErr error

func tryMach_approximate_time() (uint64, error) {
	if _mach_approximate_time == nil {
		return 0, symbolCallError("mach_approximate_time", "10.10", _mach_approximate_timeErr)
	}
	return _mach_approximate_time(), nil
}

// Mach_approximate_time.
//
// See: https://developer.apple.com/documentation/kernel/1462443-mach_approximate_time
func Mach_approximate_time() uint64 {
	result, callErr := tryMach_approximate_time()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_continuous_approximate_time func() uint64
var _mach_continuous_approximate_timeErr error

func tryMach_continuous_approximate_time() (uint64, error) {
	if _mach_continuous_approximate_time == nil {
		return 0, symbolCallError("mach_continuous_approximate_time", "10.12", _mach_continuous_approximate_timeErr)
	}
	return _mach_continuous_approximate_time(), nil
}

// Mach_continuous_approximate_time.
//
// See: https://developer.apple.com/documentation/kernel/1646198-mach_continuous_approximate_time
func Mach_continuous_approximate_time() uint64 {
	result, callErr := tryMach_continuous_approximate_time()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_continuous_time func() uint64
var _mach_continuous_timeErr error

func tryMach_continuous_time() (uint64, error) {
	if _mach_continuous_time == nil {
		return 0, symbolCallError("mach_continuous_time", "10.12", _mach_continuous_timeErr)
	}
	return _mach_continuous_time(), nil
}

// Mach_continuous_time returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), including while the system is asleep.
//
// See: https://developer.apple.com/documentation/kernel/1646199-mach_continuous_time
func Mach_continuous_time() uint64 {
	result, callErr := tryMach_continuous_time()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_make_memory_entry func(target_task uint32, size *Vm_size_t, offset Vm_offset_t, permission Vm_prot_t, object_handle *Mem_entry_name_port_t, parent_entry Mem_entry_name_port_t) Kern_return_t
var _mach_make_memory_entryErr error

func tryMach_make_memory_entry(target_task uint32, size *Vm_size_t, offset Vm_offset_t, permission Vm_prot_t, object_handle *Mem_entry_name_port_t, parent_entry Mem_entry_name_port_t) (Kern_return_t, error) {
	if _mach_make_memory_entry == nil {
		return *new(Kern_return_t), symbolCallError("mach_make_memory_entry", "10.0", _mach_make_memory_entryErr)
	}
	return _mach_make_memory_entry(target_task, size, offset, permission, object_handle, parent_entry), nil
}

// Mach_make_memory_entry.
//
// See: https://developer.apple.com/documentation/kernel/1585446-mach_make_memory_entry
func Mach_make_memory_entry(target_task uint32, size *Vm_size_t, offset Vm_offset_t, permission Vm_prot_t, object_handle *Mem_entry_name_port_t, parent_entry Mem_entry_name_port_t) Kern_return_t {
	result, callErr := tryMach_make_memory_entry(target_task, size, offset, permission, object_handle, parent_entry)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_make_memory_entry_64 func(target_task uint32, size *Memory_object_size_t, offset Memory_object_offset_t, permission Vm_prot_t, object_handle *uint32, parent_entry Mem_entry_name_port_t) Kern_return_t
var _mach_make_memory_entry_64Err error

func tryMach_make_memory_entry_64(target_task uint32, size *Memory_object_size_t, offset Memory_object_offset_t, permission Vm_prot_t, object_handle *uint32, parent_entry Mem_entry_name_port_t) (Kern_return_t, error) {
	if _mach_make_memory_entry_64 == nil {
		return *new(Kern_return_t), symbolCallError("mach_make_memory_entry_64", "10.0", _mach_make_memory_entry_64Err)
	}
	return _mach_make_memory_entry_64(target_task, size, offset, permission, object_handle, parent_entry), nil
}

// Mach_make_memory_entry_64.
//
// See: https://developer.apple.com/documentation/kernel/1585405-mach_make_memory_entry_64
func Mach_make_memory_entry_64(target_task uint32, size *Memory_object_size_t, offset Memory_object_offset_t, permission Vm_prot_t, object_handle *uint32, parent_entry Mem_entry_name_port_t) Kern_return_t {
	result, callErr := tryMach_make_memory_entry_64(target_task, size, offset, permission, object_handle, parent_entry)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_entry_access_tracking func(mem_entry Mem_entry_name_port_t, access_tracking *int32, access_tracking_reads *uint32, access_tracking_writes *uint32) Kern_return_t
var _mach_memory_entry_access_trackingErr error

func tryMach_memory_entry_access_tracking(mem_entry Mem_entry_name_port_t, access_tracking *int32, access_tracking_reads *uint32, access_tracking_writes *uint32) (Kern_return_t, error) {
	if _mach_memory_entry_access_tracking == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_entry_access_tracking", "10.14", _mach_memory_entry_access_trackingErr)
	}
	return _mach_memory_entry_access_tracking(mem_entry, access_tracking, access_tracking_reads, access_tracking_writes), nil
}

// Mach_memory_entry_access_tracking.
//
// See: https://developer.apple.com/documentation/kernel/2967371-mach_memory_entry_access_trackin
func Mach_memory_entry_access_tracking(mem_entry Mem_entry_name_port_t, access_tracking *int32, access_tracking_reads *uint32, access_tracking_writes *uint32) Kern_return_t {
	result, callErr := tryMach_memory_entry_access_tracking(mem_entry, access_tracking, access_tracking_reads, access_tracking_writes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_entry_ownership func(mem_entry Mem_entry_name_port_t, owner Task_t, ledger_tag int32, ledger_flags int32) Kern_return_t
var _mach_memory_entry_ownershipErr error

func tryMach_memory_entry_ownership(mem_entry Mem_entry_name_port_t, owner Task_t, ledger_tag int32, ledger_flags int32) (Kern_return_t, error) {
	if _mach_memory_entry_ownership == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_entry_ownership", "10.15", _mach_memory_entry_ownershipErr)
	}
	return _mach_memory_entry_ownership(mem_entry, owner, ledger_tag, ledger_flags), nil
}

// Mach_memory_entry_ownership.
//
// See: https://developer.apple.com/documentation/kernel/3143277-mach_memory_entry_ownership
func Mach_memory_entry_ownership(mem_entry Mem_entry_name_port_t, owner Task_t, ledger_tag int32, ledger_flags int32) Kern_return_t {
	result, callErr := tryMach_memory_entry_ownership(mem_entry, owner, ledger_tag, ledger_flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_entry_purgable_control func(mem_entry Mem_entry_name_port_t, control Vm_purgable_t, state *int32) Kern_return_t
var _mach_memory_entry_purgable_controlErr error

func tryMach_memory_entry_purgable_control(mem_entry Mem_entry_name_port_t, control Vm_purgable_t, state *int32) (Kern_return_t, error) {
	if _mach_memory_entry_purgable_control == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_entry_purgable_control", "10.14", _mach_memory_entry_purgable_controlErr)
	}
	return _mach_memory_entry_purgable_control(mem_entry, control, state), nil
}

// Mach_memory_entry_purgable_control.
//
// See: https://developer.apple.com/documentation/kernel/2967372-mach_memory_entry_purgable_contr
func Mach_memory_entry_purgable_control(mem_entry Mem_entry_name_port_t, control Vm_purgable_t, state *int32) Kern_return_t {
	result, callErr := tryMach_memory_entry_purgable_control(mem_entry, control, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_info func(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t, memory_info *Mach_memory_info_array_t, memory_infoCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_memory_infoErr error

func tryMach_memory_info(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t, memory_info *Mach_memory_info_array_t, memory_infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_memory_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_info", "10.11", _mach_memory_infoErr)
	}
	return _mach_memory_info(host, names, namesCnt, info, infoCnt, memory_info, memory_infoCnt), nil
}

// Mach_memory_info.
//
// See: https://developer.apple.com/documentation/kernel/1502832-mach_memory_info
func Mach_memory_info(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t, memory_info *Mach_memory_info_array_t, memory_infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_memory_info(host, names, namesCnt, info, infoCnt, memory_info, memory_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_object_memory_entry func(host Host_t, internal Boolean_t, size Vm_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) Kern_return_t
var _mach_memory_object_memory_entryErr error

func tryMach_memory_object_memory_entry(host Host_t, internal Boolean_t, size Vm_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) (Kern_return_t, error) {
	if _mach_memory_object_memory_entry == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_object_memory_entry", "10.0", _mach_memory_object_memory_entryErr)
	}
	return _mach_memory_object_memory_entry(host, internal, size, permission, pager, entry_handle), nil
}

// Mach_memory_object_memory_entry.
//
// See: https://developer.apple.com/documentation/kernel/1502680-mach_memory_object_memory_entry
func Mach_memory_object_memory_entry(host Host_t, internal Boolean_t, size Vm_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) Kern_return_t {
	result, callErr := tryMach_memory_object_memory_entry(host, internal, size, permission, pager, entry_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_memory_object_memory_entry_64 func(host Host_t, internal Boolean_t, size Memory_object_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) Kern_return_t
var _mach_memory_object_memory_entry_64Err error

func tryMach_memory_object_memory_entry_64(host Host_t, internal Boolean_t, size Memory_object_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) (Kern_return_t, error) {
	if _mach_memory_object_memory_entry_64 == nil {
		return *new(Kern_return_t), symbolCallError("mach_memory_object_memory_entry_64", "10.0", _mach_memory_object_memory_entry_64Err)
	}
	return _mach_memory_object_memory_entry_64(host, internal, size, permission, pager, entry_handle), nil
}

// Mach_memory_object_memory_entry_64.
//
// See: https://developer.apple.com/documentation/kernel/1502560-mach_memory_object_memory_entry_
func Mach_memory_object_memory_entry_64(host Host_t, internal Boolean_t, size Memory_object_size_t, permission Vm_prot_t, pager Memory_object_t, entry_handle *uint32) Kern_return_t {
	result, callErr := tryMach_memory_object_memory_entry_64(host, internal, size, permission, pager, entry_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_msg func(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t) Mach_msg_return_t
var _mach_msgErr error

func tryMach_msg(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t) (Mach_msg_return_t, error) {
	if _mach_msg == nil {
		return *new(Mach_msg_return_t), symbolCallError("mach_msg", "", _mach_msgErr)
	}
	return _mach_msg(msg, option, send_size, rcv_size, rcv_name, timeout, notify), nil
}

// Mach_msg sends and receives a Mach message, according to the option bits.
func Mach_msg(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t) Mach_msg_return_t {
	result, callErr := tryMach_msg(msg, option, send_size, rcv_size, rcv_name, timeout, notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_msg_overwrite func(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t, rcv_msg unsafe.Pointer, rcv_limit Mach_msg_size_t) Mach_msg_return_t
var _mach_msg_overwriteErr error

func tryMach_msg_overwrite(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t, rcv_msg unsafe.Pointer, rcv_limit Mach_msg_size_t) (Mach_msg_return_t, error) {
	if _mach_msg_overwrite == nil {
		return *new(Mach_msg_return_t), symbolCallError("mach_msg_overwrite", "10.0", _mach_msg_overwriteErr)
	}
	return _mach_msg_overwrite(msg, option, send_size, rcv_size, rcv_name, timeout, notify, rcv_msg, rcv_limit), nil
}

// Mach_msg_overwrite.
//
// See: https://developer.apple.com/documentation/kernel/1528967-mach_msg_overwrite
func Mach_msg_overwrite(msg unsafe.Pointer, option Mach_msg_option_t, send_size Mach_msg_size_t, rcv_size Mach_msg_size_t, rcv_name Mach_port_name_t, timeout Mach_msg_timeout_t, notify Mach_port_name_t, rcv_msg unsafe.Pointer, rcv_limit Mach_msg_size_t) Mach_msg_return_t {
	result, callErr := tryMach_msg_overwrite(msg, option, send_size, rcv_size, rcv_name, timeout, notify, rcv_msg, rcv_limit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_allocate func(task Ipc_space_t, right Mach_port_right_t, name *Mach_port_name_t) Kern_return_t
var _mach_port_allocateErr error

func tryMach_port_allocate(task Ipc_space_t, right Mach_port_right_t, name *Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_allocate == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_allocate", "10.0", _mach_port_allocateErr)
	}
	return _mach_port_allocate(task, right, name), nil
}

// Mach_port_allocate.
//
// See: https://developer.apple.com/documentation/kernel/1578704-mach_port_allocate
func Mach_port_allocate(task Ipc_space_t, right Mach_port_right_t, name *Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_allocate(task, right, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_allocate_full func(task Ipc_space_t, right Mach_port_right_t, proto uint32, qos *Mach_port_qos_t, name *Mach_port_name_t) Kern_return_t
var _mach_port_allocate_fullErr error

func tryMach_port_allocate_full(task Ipc_space_t, right Mach_port_right_t, proto uint32, qos *Mach_port_qos_t, name *Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_allocate_full == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_allocate_full", "10.0", _mach_port_allocate_fullErr)
	}
	return _mach_port_allocate_full(task, right, proto, qos, name), nil
}

// Mach_port_allocate_full.
//
// See: https://developer.apple.com/documentation/kernel/1578763-mach_port_allocate_full
func Mach_port_allocate_full(task Ipc_space_t, right Mach_port_right_t, proto uint32, qos *Mach_port_qos_t, name *Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_allocate_full(task, right, proto, qos, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_allocate_name func(task Ipc_space_t, right Mach_port_right_t, name Mach_port_name_t) Kern_return_t
var _mach_port_allocate_nameErr error

func tryMach_port_allocate_name(task Ipc_space_t, right Mach_port_right_t, name Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_allocate_name == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_allocate_name", "10.0", _mach_port_allocate_nameErr)
	}
	return _mach_port_allocate_name(task, right, name), nil
}

// Mach_port_allocate_name.
//
// See: https://developer.apple.com/documentation/kernel/1578657-mach_port_allocate_name
func Mach_port_allocate_name(task Ipc_space_t, right Mach_port_right_t, name Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_allocate_name(task, right, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_allocate_qos func(task Ipc_space_t, right Mach_port_right_t, qos *Mach_port_qos_t, name *Mach_port_name_t) Kern_return_t
var _mach_port_allocate_qosErr error

func tryMach_port_allocate_qos(task Ipc_space_t, right Mach_port_right_t, qos *Mach_port_qos_t, name *Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_allocate_qos == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_allocate_qos", "10.0", _mach_port_allocate_qosErr)
	}
	return _mach_port_allocate_qos(task, right, qos, name), nil
}

// Mach_port_allocate_qos.
//
// See: https://developer.apple.com/documentation/kernel/1578746-mach_port_allocate_qos
func Mach_port_allocate_qos(task Ipc_space_t, right Mach_port_right_t, qos *Mach_port_qos_t, name *Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_allocate_qos(task, right, qos, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_assert_attributes func(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, info Mach_port_info_t, infoCnt Mach_msg_type_number_t) Kern_return_t
var _mach_port_assert_attributesErr error

func tryMach_port_assert_attributes(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, info Mach_port_info_t, infoCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_assert_attributes == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_assert_attributes", "12.0", _mach_port_assert_attributesErr)
	}
	return _mach_port_assert_attributes(task, name, flavor, info, infoCnt), nil
}

// Mach_port_assert_attributes.
//
// See: https://developer.apple.com/documentation/kernel/3786107-mach_port_assert_attributes
func Mach_port_assert_attributes(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, info Mach_port_info_t, infoCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_assert_attributes(task, name, flavor, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_construct func(task Ipc_space_t, options Mach_port_options_ptr_t, context uint64, name *Mach_port_name_t) Kern_return_t
var _mach_port_constructErr error

func tryMach_port_construct(task Ipc_space_t, options Mach_port_options_ptr_t, context uint64, name *Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_construct == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_construct", "10.9", _mach_port_constructErr)
	}
	return _mach_port_construct(task, options, context, name), nil
}

// Mach_port_construct.
//
// See: https://developer.apple.com/documentation/kernel/1578687-mach_port_construct
func Mach_port_construct(task Ipc_space_t, options Mach_port_options_ptr_t, context uint64, name *Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_construct(task, options, context, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_deallocate func(task Ipc_space_t, name Mach_port_name_t) Kern_return_t
var _mach_port_deallocateErr error

func tryMach_port_deallocate(task Ipc_space_t, name Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_deallocate == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_deallocate", "10.0", _mach_port_deallocateErr)
	}
	return _mach_port_deallocate(task, name), nil
}

// Mach_port_deallocate.
//
// See: https://developer.apple.com/documentation/kernel/1578777-mach_port_deallocate
func Mach_port_deallocate(task Ipc_space_t, name Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_deallocate(task, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_destroy func(task Ipc_space_t, name Mach_port_name_t) Kern_return_t
var _mach_port_destroyErr error

func tryMach_port_destroy(task Ipc_space_t, name Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_destroy == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_destroy", "10.0", _mach_port_destroyErr)
	}
	return _mach_port_destroy(task, name), nil
}

// Mach_port_destroy.
//
// See: https://developer.apple.com/documentation/kernel/1578817-mach_port_destroy
func Mach_port_destroy(task Ipc_space_t, name Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_destroy(task, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_destruct func(task Ipc_space_t, name Mach_port_name_t, srdelta Mach_port_delta_t, guard uint64) Kern_return_t
var _mach_port_destructErr error

func tryMach_port_destruct(task Ipc_space_t, name Mach_port_name_t, srdelta Mach_port_delta_t, guard uint64) (Kern_return_t, error) {
	if _mach_port_destruct == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_destruct", "10.9", _mach_port_destructErr)
	}
	return _mach_port_destruct(task, name, srdelta, guard), nil
}

// Mach_port_destruct.
//
// See: https://developer.apple.com/documentation/kernel/1578881-mach_port_destruct
func Mach_port_destruct(task Ipc_space_t, name Mach_port_name_t, srdelta Mach_port_delta_t, guard uint64) Kern_return_t {
	result, callErr := tryMach_port_destruct(task, name, srdelta, guard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_dnrequest_info func(task Ipc_space_t, name Mach_port_name_t, dnr_total *uint32, dnr_used *uint32) Kern_return_t
var _mach_port_dnrequest_infoErr error

func tryMach_port_dnrequest_info(task Ipc_space_t, name Mach_port_name_t, dnr_total *uint32, dnr_used *uint32) (Kern_return_t, error) {
	if _mach_port_dnrequest_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_dnrequest_info", "10.0", _mach_port_dnrequest_infoErr)
	}
	return _mach_port_dnrequest_info(task, name, dnr_total, dnr_used), nil
}

// Mach_port_dnrequest_info.
//
// See: https://developer.apple.com/documentation/kernel/1578613-mach_port_dnrequest_info
func Mach_port_dnrequest_info(task Ipc_space_t, name Mach_port_name_t, dnr_total *uint32, dnr_used *uint32) Kern_return_t {
	result, callErr := tryMach_port_dnrequest_info(task, name, dnr_total, dnr_used)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_extract_member func(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) Kern_return_t
var _mach_port_extract_memberErr error

func tryMach_port_extract_member(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_extract_member == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_extract_member", "10.0", _mach_port_extract_memberErr)
	}
	return _mach_port_extract_member(task, name, pset), nil
}

// Mach_port_extract_member.
//
// See: https://developer.apple.com/documentation/kernel/1578633-mach_port_extract_member
func Mach_port_extract_member(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_extract_member(task, name, pset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_extract_right func(task Ipc_space_t, name Mach_port_name_t, msgt_name Mach_msg_type_name_t, poly *uint32, polyPoly *Mach_msg_type_name_t) Kern_return_t
var _mach_port_extract_rightErr error

func tryMach_port_extract_right(task Ipc_space_t, name Mach_port_name_t, msgt_name Mach_msg_type_name_t, poly *uint32, polyPoly *Mach_msg_type_name_t) (Kern_return_t, error) {
	if _mach_port_extract_right == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_extract_right", "10.0", _mach_port_extract_rightErr)
	}
	return _mach_port_extract_right(task, name, msgt_name, poly, polyPoly), nil
}

// Mach_port_extract_right.
//
// See: https://developer.apple.com/documentation/kernel/1578688-mach_port_extract_right
func Mach_port_extract_right(task Ipc_space_t, name Mach_port_name_t, msgt_name Mach_msg_type_name_t, poly *uint32, polyPoly *Mach_msg_type_name_t) Kern_return_t {
	result, callErr := tryMach_port_extract_right(task, name, msgt_name, poly, polyPoly)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_attributes func(task Ipc_space_read_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info_out Mach_port_info_t, port_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_port_get_attributesErr error

func tryMach_port_get_attributes(task Ipc_space_read_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info_out Mach_port_info_t, port_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_get_attributes == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_attributes", "10.0", _mach_port_get_attributesErr)
	}
	return _mach_port_get_attributes(task, name, flavor, port_info_out, port_info_outCnt), nil
}

// Mach_port_get_attributes.
//
// See: https://developer.apple.com/documentation/kernel/1578800-mach_port_get_attributes
func Mach_port_get_attributes(task Ipc_space_read_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info_out Mach_port_info_t, port_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_get_attributes(task, name, flavor, port_info_out, port_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_context func(task Ipc_space_read_t, name Mach_port_name_t, context *Mach_vm_address_t) Kern_return_t
var _mach_port_get_contextErr error

func tryMach_port_get_context(task Ipc_space_read_t, name Mach_port_name_t, context *Mach_vm_address_t) (Kern_return_t, error) {
	if _mach_port_get_context == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_context", "10.6", _mach_port_get_contextErr)
	}
	return _mach_port_get_context(task, name, context), nil
}

// Mach_port_get_context.
//
// See: https://developer.apple.com/documentation/kernel/1578930-mach_port_get_context
func Mach_port_get_context(task Ipc_space_read_t, name Mach_port_name_t, context *Mach_vm_address_t) Kern_return_t {
	result, callErr := tryMach_port_get_context(task, name, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_refs func(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, refs *Mach_port_urefs_t) Kern_return_t
var _mach_port_get_refsErr error

func tryMach_port_get_refs(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, refs *Mach_port_urefs_t) (Kern_return_t, error) {
	if _mach_port_get_refs == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_refs", "10.0", _mach_port_get_refsErr)
	}
	return _mach_port_get_refs(task, name, right, refs), nil
}

// Mach_port_get_refs.
//
// See: https://developer.apple.com/documentation/kernel/1578946-mach_port_get_refs
func Mach_port_get_refs(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, refs *Mach_port_urefs_t) Kern_return_t {
	result, callErr := tryMach_port_get_refs(task, name, right, refs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_service_port_info func(task Ipc_space_read_t, name Mach_port_name_t, sp_info_out *Mach_service_port_info_data_t) Kern_return_t
var _mach_port_get_service_port_infoErr error

func tryMach_port_get_service_port_info(task Ipc_space_read_t, name Mach_port_name_t, sp_info_out *Mach_service_port_info_data_t) (Kern_return_t, error) {
	if _mach_port_get_service_port_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_service_port_info", "12.0", _mach_port_get_service_port_infoErr)
	}
	return _mach_port_get_service_port_info(task, name, sp_info_out), nil
}

// Mach_port_get_service_port_info.
//
// See: https://developer.apple.com/documentation/kernel/3753657-mach_port_get_service_port_info
func Mach_port_get_service_port_info(task Ipc_space_read_t, name Mach_port_name_t, sp_info_out *Mach_service_port_info_data_t) Kern_return_t {
	result, callErr := tryMach_port_get_service_port_info(task, name, sp_info_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_set_status func(task Ipc_space_read_t, name Mach_port_name_t, members *Mach_port_name_array_t, membersCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_port_get_set_statusErr error

func tryMach_port_get_set_status(task Ipc_space_read_t, name Mach_port_name_t, members *Mach_port_name_array_t, membersCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_get_set_status == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_set_status", "10.0", _mach_port_get_set_statusErr)
	}
	return _mach_port_get_set_status(task, name, members, membersCnt), nil
}

// Mach_port_get_set_status.
//
// See: https://developer.apple.com/documentation/kernel/1578936-mach_port_get_set_status
func Mach_port_get_set_status(task Ipc_space_read_t, name Mach_port_name_t, members *Mach_port_name_array_t, membersCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_get_set_status(task, name, members, membersCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_get_srights func(task Ipc_space_t, name Mach_port_name_t, srights *Mach_port_rights_t) Kern_return_t
var _mach_port_get_srightsErr error

func tryMach_port_get_srights(task Ipc_space_t, name Mach_port_name_t, srights *Mach_port_rights_t) (Kern_return_t, error) {
	if _mach_port_get_srights == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_get_srights", "10.0", _mach_port_get_srightsErr)
	}
	return _mach_port_get_srights(task, name, srights), nil
}

// Mach_port_get_srights.
//
// See: https://developer.apple.com/documentation/kernel/1578818-mach_port_get_srights
func Mach_port_get_srights(task Ipc_space_t, name Mach_port_name_t, srights *Mach_port_rights_t) Kern_return_t {
	result, callErr := tryMach_port_get_srights(task, name, srights)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_guard func(task Ipc_space_t, name Mach_port_name_t, guard uint64, strict Boolean_t) Kern_return_t
var _mach_port_guardErr error

func tryMach_port_guard(task Ipc_space_t, name Mach_port_name_t, guard uint64, strict Boolean_t) (Kern_return_t, error) {
	if _mach_port_guard == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_guard", "10.9", _mach_port_guardErr)
	}
	return _mach_port_guard(task, name, guard, strict), nil
}

// Mach_port_guard.
//
// See: https://developer.apple.com/documentation/kernel/1578772-mach_port_guard
func Mach_port_guard(task Ipc_space_t, name Mach_port_name_t, guard uint64, strict Boolean_t) Kern_return_t {
	result, callErr := tryMach_port_guard(task, name, guard, strict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_guard_with_flags func(task Ipc_space_t, name Mach_port_name_t, guard uint64, flags uint64) Kern_return_t
var _mach_port_guard_with_flagsErr error

func tryMach_port_guard_with_flags(task Ipc_space_t, name Mach_port_name_t, guard uint64, flags uint64) (Kern_return_t, error) {
	if _mach_port_guard_with_flags == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_guard_with_flags", "10.15", _mach_port_guard_with_flagsErr)
	}
	return _mach_port_guard_with_flags(task, name, guard, flags), nil
}

// Mach_port_guard_with_flags.
//
// See: https://developer.apple.com/documentation/kernel/3181823-mach_port_guard_with_flags
func Mach_port_guard_with_flags(task Ipc_space_t, name Mach_port_name_t, guard uint64, flags uint64) Kern_return_t {
	result, callErr := tryMach_port_guard_with_flags(task, name, guard, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_insert_member func(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) Kern_return_t
var _mach_port_insert_memberErr error

func tryMach_port_insert_member(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_insert_member == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_insert_member", "10.0", _mach_port_insert_memberErr)
	}
	return _mach_port_insert_member(task, name, pset), nil
}

// Mach_port_insert_member.
//
// See: https://developer.apple.com/documentation/kernel/1578885-mach_port_insert_member
func Mach_port_insert_member(task Ipc_space_t, name Mach_port_name_t, pset Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_insert_member(task, name, pset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_insert_right func(task Ipc_space_t, name Mach_port_name_t, poly uint32, polyPoly Mach_msg_type_name_t) Kern_return_t
var _mach_port_insert_rightErr error

func tryMach_port_insert_right(task Ipc_space_t, name Mach_port_name_t, poly uint32, polyPoly Mach_msg_type_name_t) (Kern_return_t, error) {
	if _mach_port_insert_right == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_insert_right", "10.0", _mach_port_insert_rightErr)
	}
	return _mach_port_insert_right(task, name, poly, polyPoly), nil
}

// Mach_port_insert_right.
//
// See: https://developer.apple.com/documentation/kernel/1578739-mach_port_insert_right
func Mach_port_insert_right(task Ipc_space_t, name Mach_port_name_t, poly uint32, polyPoly Mach_msg_type_name_t) Kern_return_t {
	result, callErr := tryMach_port_insert_right(task, name, poly, polyPoly)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_is_connection_for_service func(task Ipc_space_t, connection_port Mach_port_name_t, service_port Mach_port_name_t, filter_policy_id *uint64) Kern_return_t
var _mach_port_is_connection_for_serviceErr error

func tryMach_port_is_connection_for_service(task Ipc_space_t, connection_port Mach_port_name_t, service_port Mach_port_name_t, filter_policy_id *uint64) (Kern_return_t, error) {
	if _mach_port_is_connection_for_service == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_is_connection_for_service", "12.0", _mach_port_is_connection_for_serviceErr)
	}
	return _mach_port_is_connection_for_service(task, connection_port, service_port, filter_policy_id), nil
}

// Mach_port_is_connection_for_service.
//
// See: https://developer.apple.com/documentation/kernel/3753658-mach_port_is_connection_for_serv
func Mach_port_is_connection_for_service(task Ipc_space_t, connection_port Mach_port_name_t, service_port Mach_port_name_t, filter_policy_id *uint64) Kern_return_t {
	result, callErr := tryMach_port_is_connection_for_service(task, connection_port, service_port, filter_policy_id)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_kernel_object func(task Ipc_space_read_t, name Mach_port_name_t, object_type *uint32, object_addr *uint32) Kern_return_t
var _mach_port_kernel_objectErr error

func tryMach_port_kernel_object(task Ipc_space_read_t, name Mach_port_name_t, object_type *uint32, object_addr *uint32) (Kern_return_t, error) {
	if _mach_port_kernel_object == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_kernel_object", "10.0", _mach_port_kernel_objectErr)
	}
	return _mach_port_kernel_object(task, name, object_type, object_addr), nil
}

// Mach_port_kernel_object.
//
// See: https://developer.apple.com/documentation/kernel/1578723-mach_port_kernel_object
func Mach_port_kernel_object(task Ipc_space_read_t, name Mach_port_name_t, object_type *uint32, object_addr *uint32) Kern_return_t {
	result, callErr := tryMach_port_kernel_object(task, name, object_type, object_addr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_kobject func(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t) Kern_return_t
var _mach_port_kobjectErr error

func tryMach_port_kobject(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t) (Kern_return_t, error) {
	if _mach_port_kobject == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_kobject", "10.6", _mach_port_kobjectErr)
	}
	return _mach_port_kobject(task, name, object_type, object_addr), nil
}

// Mach_port_kobject.
//
// See: https://developer.apple.com/documentation/kernel/1578702-mach_port_kobject
func Mach_port_kobject(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t) Kern_return_t {
	result, callErr := tryMach_port_kobject(task, name, object_type, object_addr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_kobject_description func(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t, description uintptr) Kern_return_t
var _mach_port_kobject_descriptionErr error

func tryMach_port_kobject_description(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t, description uintptr) (Kern_return_t, error) {
	if _mach_port_kobject_description == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_kobject_description", "10.15.4", _mach_port_kobject_descriptionErr)
	}
	return _mach_port_kobject_description(task, name, object_type, object_addr, description), nil
}

// Mach_port_kobject_description.
//
// See: https://developer.apple.com/documentation/kernel/3516847-mach_port_kobject_description
func Mach_port_kobject_description(task Ipc_space_read_t, name Mach_port_name_t, object_type *Natural_t, object_addr *Mach_vm_address_t, description uintptr) Kern_return_t {
	result, callErr := tryMach_port_kobject_description(task, name, object_type, object_addr, description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_mod_refs func(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, delta Mach_port_delta_t) Kern_return_t
var _mach_port_mod_refsErr error

func tryMach_port_mod_refs(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, delta Mach_port_delta_t) (Kern_return_t, error) {
	if _mach_port_mod_refs == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_mod_refs", "10.0", _mach_port_mod_refsErr)
	}
	return _mach_port_mod_refs(task, name, right, delta), nil
}

// Mach_port_mod_refs.
//
// See: https://developer.apple.com/documentation/kernel/1578894-mach_port_mod_refs
func Mach_port_mod_refs(task Ipc_space_t, name Mach_port_name_t, right Mach_port_right_t, delta Mach_port_delta_t) Kern_return_t {
	result, callErr := tryMach_port_mod_refs(task, name, right, delta)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_move_member func(task Ipc_space_t, member Mach_port_name_t, after Mach_port_name_t) Kern_return_t
var _mach_port_move_memberErr error

func tryMach_port_move_member(task Ipc_space_t, member Mach_port_name_t, after Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_move_member == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_move_member", "10.0", _mach_port_move_memberErr)
	}
	return _mach_port_move_member(task, member, after), nil
}

// Mach_port_move_member.
//
// See: https://developer.apple.com/documentation/kernel/1578673-mach_port_move_member
func Mach_port_move_member(task Ipc_space_t, member Mach_port_name_t, after Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_move_member(task, member, after)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_names func(task Ipc_space_t, names *Mach_port_name_array_t, namesCnt *Mach_msg_type_number_t, types *Mach_port_type_array_t, typesCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_port_namesErr error

func tryMach_port_names(task Ipc_space_t, names *Mach_port_name_array_t, namesCnt *Mach_msg_type_number_t, types *Mach_port_type_array_t, typesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_names == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_names", "10.0", _mach_port_namesErr)
	}
	return _mach_port_names(task, names, namesCnt, types, typesCnt), nil
}

// Mach_port_names.
//
// See: https://developer.apple.com/documentation/kernel/1578814-mach_port_names
func Mach_port_names(task Ipc_space_t, names *Mach_port_name_array_t, namesCnt *Mach_msg_type_number_t, types *Mach_port_type_array_t, typesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_names(task, names, namesCnt, types, typesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_peek func(task Ipc_space_t, name Mach_port_name_t, trailer_type Mach_msg_trailer_type_t, request_seqnop *Mach_port_seqno_t, msg_sizep *Mach_msg_size_t, msg_idp *Mach_msg_id_t, trailer_infop Mach_msg_trailer_info_t, trailer_infopCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_port_peekErr error

func tryMach_port_peek(task Ipc_space_t, name Mach_port_name_t, trailer_type Mach_msg_trailer_type_t, request_seqnop *Mach_port_seqno_t, msg_sizep *Mach_msg_size_t, msg_idp *Mach_msg_id_t, trailer_infop Mach_msg_trailer_info_t, trailer_infopCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_peek == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_peek", "10.9", _mach_port_peekErr)
	}
	return _mach_port_peek(task, name, trailer_type, request_seqnop, msg_sizep, msg_idp, trailer_infop, trailer_infopCnt), nil
}

// Mach_port_peek.
//
// See: https://developer.apple.com/documentation/kernel/1578839-mach_port_peek
func Mach_port_peek(task Ipc_space_t, name Mach_port_name_t, trailer_type Mach_msg_trailer_type_t, request_seqnop *Mach_port_seqno_t, msg_sizep *Mach_msg_size_t, msg_idp *Mach_msg_id_t, trailer_infop Mach_msg_trailer_info_t, trailer_infopCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_peek(task, name, trailer_type, request_seqnop, msg_sizep, msg_idp, trailer_infop, trailer_infopCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_rename func(task Ipc_space_t, old_name Mach_port_name_t, new_name Mach_port_name_t) Kern_return_t
var _mach_port_renameErr error

func tryMach_port_rename(task Ipc_space_t, old_name Mach_port_name_t, new_name Mach_port_name_t) (Kern_return_t, error) {
	if _mach_port_rename == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_rename", "10.0", _mach_port_renameErr)
	}
	return _mach_port_rename(task, old_name, new_name), nil
}

// Mach_port_rename.
//
// See: https://developer.apple.com/documentation/kernel/1578909-mach_port_rename
func Mach_port_rename(task Ipc_space_t, old_name Mach_port_name_t, new_name Mach_port_name_t) Kern_return_t {
	result, callErr := tryMach_port_rename(task, old_name, new_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_request_notification func(task Ipc_space_t, name Mach_port_name_t, msgid Mach_msg_id_t, sync Mach_port_mscount_t, notify uint32, notifyPoly Mach_msg_type_name_t, previous *uint32) Kern_return_t
var _mach_port_request_notificationErr error

func tryMach_port_request_notification(task Ipc_space_t, name Mach_port_name_t, msgid Mach_msg_id_t, sync Mach_port_mscount_t, notify uint32, notifyPoly Mach_msg_type_name_t, previous *uint32) (Kern_return_t, error) {
	if _mach_port_request_notification == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_request_notification", "10.0", _mach_port_request_notificationErr)
	}
	return _mach_port_request_notification(task, name, msgid, sync, notify, notifyPoly, previous), nil
}

// Mach_port_request_notification.
//
// See: https://developer.apple.com/documentation/kernel/1578734-mach_port_request_notification
func Mach_port_request_notification(task Ipc_space_t, name Mach_port_name_t, msgid Mach_msg_id_t, sync Mach_port_mscount_t, notify uint32, notifyPoly Mach_msg_type_name_t, previous *uint32) Kern_return_t {
	result, callErr := tryMach_port_request_notification(task, name, msgid, sync, notify, notifyPoly, previous)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_set_attributes func(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info Mach_port_info_t, port_infoCnt Mach_msg_type_number_t) Kern_return_t
var _mach_port_set_attributesErr error

func tryMach_port_set_attributes(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info Mach_port_info_t, port_infoCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_set_attributes == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_set_attributes", "10.0", _mach_port_set_attributesErr)
	}
	return _mach_port_set_attributes(task, name, flavor, port_info, port_infoCnt), nil
}

// Mach_port_set_attributes.
//
// See: https://developer.apple.com/documentation/kernel/1578964-mach_port_set_attributes
func Mach_port_set_attributes(task Ipc_space_t, name Mach_port_name_t, flavor Mach_port_flavor_t, port_info Mach_port_info_t, port_infoCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_set_attributes(task, name, flavor, port_info, port_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_set_context func(task Ipc_space_t, name Mach_port_name_t, context Mach_vm_address_t) Kern_return_t
var _mach_port_set_contextErr error

func tryMach_port_set_context(task Ipc_space_t, name Mach_port_name_t, context Mach_vm_address_t) (Kern_return_t, error) {
	if _mach_port_set_context == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_set_context", "10.6", _mach_port_set_contextErr)
	}
	return _mach_port_set_context(task, name, context), nil
}

// Mach_port_set_context.
//
// See: https://developer.apple.com/documentation/kernel/1578733-mach_port_set_context
func Mach_port_set_context(task Ipc_space_t, name Mach_port_name_t, context Mach_vm_address_t) Kern_return_t {
	result, callErr := tryMach_port_set_context(task, name, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_set_mscount func(task Ipc_space_t, name Mach_port_name_t, mscount Mach_port_mscount_t) Kern_return_t
var _mach_port_set_mscountErr error

func tryMach_port_set_mscount(task Ipc_space_t, name Mach_port_name_t, mscount Mach_port_mscount_t) (Kern_return_t, error) {
	if _mach_port_set_mscount == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_set_mscount", "10.0", _mach_port_set_mscountErr)
	}
	return _mach_port_set_mscount(task, name, mscount), nil
}

// Mach_port_set_mscount.
//
// See: https://developer.apple.com/documentation/kernel/1578719-mach_port_set_mscount
func Mach_port_set_mscount(task Ipc_space_t, name Mach_port_name_t, mscount Mach_port_mscount_t) Kern_return_t {
	result, callErr := tryMach_port_set_mscount(task, name, mscount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_set_seqno func(task Ipc_space_t, name Mach_port_name_t, seqno Mach_port_seqno_t) Kern_return_t
var _mach_port_set_seqnoErr error

func tryMach_port_set_seqno(task Ipc_space_t, name Mach_port_name_t, seqno Mach_port_seqno_t) (Kern_return_t, error) {
	if _mach_port_set_seqno == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_set_seqno", "10.0", _mach_port_set_seqnoErr)
	}
	return _mach_port_set_seqno(task, name, seqno), nil
}

// Mach_port_set_seqno.
//
// See: https://developer.apple.com/documentation/kernel/1578744-mach_port_set_seqno
func Mach_port_set_seqno(task Ipc_space_t, name Mach_port_name_t, seqno Mach_port_seqno_t) Kern_return_t {
	result, callErr := tryMach_port_set_seqno(task, name, seqno)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_space_basic_info func(task Ipc_space_inspect_t, basic_info *Ipc_info_space_basic_t) Kern_return_t
var _mach_port_space_basic_infoErr error

func tryMach_port_space_basic_info(task Ipc_space_inspect_t, basic_info *Ipc_info_space_basic_t) (Kern_return_t, error) {
	if _mach_port_space_basic_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_space_basic_info", "10.10", _mach_port_space_basic_infoErr)
	}
	return _mach_port_space_basic_info(task, basic_info), nil
}

// Mach_port_space_basic_info.
//
// See: https://developer.apple.com/documentation/kernel/1578841-mach_port_space_basic_info
func Mach_port_space_basic_info(task Ipc_space_inspect_t, basic_info *Ipc_info_space_basic_t) Kern_return_t {
	result, callErr := tryMach_port_space_basic_info(task, basic_info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_space_info func(space Ipc_space_read_t, space_info *Ipc_info_space_t, table_info *Ipc_info_name_array_t, table_infoCnt *Mach_msg_type_number_t, tree_info *Ipc_info_tree_name_array_t, tree_infoCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_port_space_infoErr error

func tryMach_port_space_info(space Ipc_space_read_t, space_info *Ipc_info_space_t, table_info *Ipc_info_name_array_t, table_infoCnt *Mach_msg_type_number_t, tree_info *Ipc_info_tree_name_array_t, tree_infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_port_space_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_space_info", "10.0", _mach_port_space_infoErr)
	}
	return _mach_port_space_info(space, space_info, table_info, table_infoCnt, tree_info, tree_infoCnt), nil
}

// Mach_port_space_info.
//
// See: https://developer.apple.com/documentation/kernel/1578884-mach_port_space_info
func Mach_port_space_info(space Ipc_space_read_t, space_info *Ipc_info_space_t, table_info *Ipc_info_name_array_t, table_infoCnt *Mach_msg_type_number_t, tree_info *Ipc_info_tree_name_array_t, tree_infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_port_space_info(space, space_info, table_info, table_infoCnt, tree_info, tree_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_swap_guard func(task Ipc_space_t, name Mach_port_name_t, old_guard uint64, new_guard uint64) Kern_return_t
var _mach_port_swap_guardErr error

func tryMach_port_swap_guard(task Ipc_space_t, name Mach_port_name_t, old_guard uint64, new_guard uint64) (Kern_return_t, error) {
	if _mach_port_swap_guard == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_swap_guard", "10.15", _mach_port_swap_guardErr)
	}
	return _mach_port_swap_guard(task, name, old_guard, new_guard), nil
}

// Mach_port_swap_guard.
//
// See: https://developer.apple.com/documentation/kernel/3181824-mach_port_swap_guard
func Mach_port_swap_guard(task Ipc_space_t, name Mach_port_name_t, old_guard uint64, new_guard uint64) Kern_return_t {
	result, callErr := tryMach_port_swap_guard(task, name, old_guard, new_guard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_type func(task Ipc_space_t, name Mach_port_name_t, ptype *Mach_port_type_t) Kern_return_t
var _mach_port_typeErr error

func tryMach_port_type(task Ipc_space_t, name Mach_port_name_t, ptype *Mach_port_type_t) (Kern_return_t, error) {
	if _mach_port_type == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_type", "10.0", _mach_port_typeErr)
	}
	return _mach_port_type(task, name, ptype), nil
}

// Mach_port_type.
//
// See: https://developer.apple.com/documentation/kernel/1578714-mach_port_type
func Mach_port_type(task Ipc_space_t, name Mach_port_name_t, ptype *Mach_port_type_t) Kern_return_t {
	result, callErr := tryMach_port_type(task, name, ptype)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_port_unguard func(task Ipc_space_t, name Mach_port_name_t, guard uint64) Kern_return_t
var _mach_port_unguardErr error

func tryMach_port_unguard(task Ipc_space_t, name Mach_port_name_t, guard uint64) (Kern_return_t, error) {
	if _mach_port_unguard == nil {
		return *new(Kern_return_t), symbolCallError("mach_port_unguard", "10.9", _mach_port_unguardErr)
	}
	return _mach_port_unguard(task, name, guard), nil
}

// Mach_port_unguard.
//
// See: https://developer.apple.com/documentation/kernel/1578951-mach_port_unguard
func Mach_port_unguard(task Ipc_space_t, name Mach_port_name_t, guard uint64) Kern_return_t {
	result, callErr := tryMach_port_unguard(task, name, guard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_task_is_self func(task Task_name_t, is_self *Boolean_t) Kern_return_t
var _mach_task_is_selfErr error

func tryMach_task_is_self(task Task_name_t, is_self *Boolean_t) (Kern_return_t, error) {
	if _mach_task_is_self == nil {
		return *new(Kern_return_t), symbolCallError("mach_task_is_self", "11.3", _mach_task_is_selfErr)
	}
	return _mach_task_is_self(task, is_self), nil
}

// Mach_task_is_self.
//
// See: https://developer.apple.com/documentation/kernel/3727993-mach_task_is_self
func Mach_task_is_self(task Task_name_t, is_self *Boolean_t) Kern_return_t {
	result, callErr := tryMach_task_is_self(task, is_self)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_task_self func() uint32
var _mach_task_selfErr error

func tryMach_task_self() (uint32, error) {
	if _mach_task_self == nil {
		return 0, symbolCallError("mach_task_self", "", _mach_task_selfErr)
	}
	return _mach_task_self(), nil
}

// Mach_task_self returns a send right to the calling task's kernel port.
func Mach_task_self() uint32 {
	result, callErr := tryMach_task_self()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_allocate func(target uint32, address *Mach_vm_address_t, size Mach_vm_size_t, flags int32) Kern_return_t
var _mach_vm_allocateErr error

func tryMach_vm_allocate(target uint32, address *Mach_vm_address_t, size Mach_vm_size_t, flags int32) (Kern_return_t, error) {
	if _mach_vm_allocate == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_allocate", "10.4", _mach_vm_allocateErr)
	}
	return _mach_vm_allocate(target, address, size, flags), nil
}

// Mach_vm_allocate.
//
// See: https://developer.apple.com/documentation/kernel/1402376-mach_vm_allocate
func Mach_vm_allocate(target uint32, address *Mach_vm_address_t, size Mach_vm_size_t, flags int32) Kern_return_t {
	result, callErr := tryMach_vm_allocate(target, address, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_behavior_set func(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_behavior Vm_behavior_t) Kern_return_t
var _mach_vm_behavior_setErr error

func tryMach_vm_behavior_set(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_behavior Vm_behavior_t) (Kern_return_t, error) {
	if _mach_vm_behavior_set == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_behavior_set", "10.4", _mach_vm_behavior_setErr)
	}
	return _mach_vm_behavior_set(target_task, address, size, new_behavior), nil
}

// Mach_vm_behavior_set.
//
// See: https://developer.apple.com/documentation/kernel/1402468-mach_vm_behavior_set
func Mach_vm_behavior_set(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_behavior Vm_behavior_t) Kern_return_t {
	result, callErr := tryMach_vm_behavior_set(target_task, address, size, new_behavior)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_copy func(target_task uint32, source_address Mach_vm_address_t, size Mach_vm_size_t, dest_address Mach_vm_address_t) Kern_return_t
var _mach_vm_copyErr error

func tryMach_vm_copy(target_task uint32, source_address Mach_vm_address_t, size Mach_vm_size_t, dest_address Mach_vm_address_t) (Kern_return_t, error) {
	if _mach_vm_copy == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_copy", "10.4", _mach_vm_copyErr)
	}
	return _mach_vm_copy(target_task, source_address, size, dest_address), nil
}

// Mach_vm_copy.
//
// See: https://developer.apple.com/documentation/kernel/1402342-mach_vm_copy
func Mach_vm_copy(target_task uint32, source_address Mach_vm_address_t, size Mach_vm_size_t, dest_address Mach_vm_address_t) Kern_return_t {
	result, callErr := tryMach_vm_copy(target_task, source_address, size, dest_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_deallocate func(target uint32, address Mach_vm_address_t, size Mach_vm_size_t) Kern_return_t
var _mach_vm_deallocateErr error

func tryMach_vm_deallocate(target uint32, address Mach_vm_address_t, size Mach_vm_size_t) (Kern_return_t, error) {
	if _mach_vm_deallocate == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_deallocate", "10.4", _mach_vm_deallocateErr)
	}
	return _mach_vm_deallocate(target, address, size), nil
}

// Mach_vm_deallocate.
//
// See: https://developer.apple.com/documentation/kernel/1402285-mach_vm_deallocate
func Mach_vm_deallocate(target uint32, address Mach_vm_address_t, size Mach_vm_size_t) Kern_return_t {
	result, callErr := tryMach_vm_deallocate(target, address, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_deferred_reclamation_buffer_allocate func(target_task Task_t, address *Mach_vm_address_t, len_ uint32, max_len uint32) Kern_return_t
var _mach_vm_deferred_reclamation_buffer_allocateErr error

func tryMach_vm_deferred_reclamation_buffer_allocate(target_task Task_t, address *Mach_vm_address_t, len_ uint32, max_len uint32) (Kern_return_t, error) {
	if _mach_vm_deferred_reclamation_buffer_allocate == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_deferred_reclamation_buffer_allocate", "15.4", _mach_vm_deferred_reclamation_buffer_allocateErr)
	}
	return _mach_vm_deferred_reclamation_buffer_allocate(target_task, address, len_, max_len), nil
}

// Mach_vm_deferred_reclamation_buffer_allocate.
//
// See: https://developer.apple.com/documentation/kernel/4540746-mach_vm_deferred_reclamation_buf
func Mach_vm_deferred_reclamation_buffer_allocate(target_task Task_t, address *Mach_vm_address_t, len_ uint32, max_len uint32) Kern_return_t {
	result, callErr := tryMach_vm_deferred_reclamation_buffer_allocate(target_task, address, len_, max_len)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_deferred_reclamation_buffer_flush func(target_task Task_t, num_entries_to_reclaim uint32) Kern_return_t
var _mach_vm_deferred_reclamation_buffer_flushErr error

func tryMach_vm_deferred_reclamation_buffer_flush(target_task Task_t, num_entries_to_reclaim uint32) (Kern_return_t, error) {
	if _mach_vm_deferred_reclamation_buffer_flush == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_deferred_reclamation_buffer_flush", "15.4", _mach_vm_deferred_reclamation_buffer_flushErr)
	}
	return _mach_vm_deferred_reclamation_buffer_flush(target_task, num_entries_to_reclaim), nil
}

// Mach_vm_deferred_reclamation_buffer_flush.
//
// See: https://developer.apple.com/documentation/kernel/4540747-mach_vm_deferred_reclamation_buf
func Mach_vm_deferred_reclamation_buffer_flush(target_task Task_t, num_entries_to_reclaim uint32) Kern_return_t {
	result, callErr := tryMach_vm_deferred_reclamation_buffer_flush(target_task, num_entries_to_reclaim)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_deferred_reclamation_buffer_resize func(target_task Task_t, size uint32) Kern_return_t
var _mach_vm_deferred_reclamation_buffer_resizeErr error

func tryMach_vm_deferred_reclamation_buffer_resize(target_task Task_t, size uint32) (Kern_return_t, error) {
	if _mach_vm_deferred_reclamation_buffer_resize == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_deferred_reclamation_buffer_resize", "15.4", _mach_vm_deferred_reclamation_buffer_resizeErr)
	}
	return _mach_vm_deferred_reclamation_buffer_resize(target_task, size), nil
}

// Mach_vm_deferred_reclamation_buffer_resize.
//
// See: https://developer.apple.com/documentation/kernel/4540748-mach_vm_deferred_reclamation_buf
func Mach_vm_deferred_reclamation_buffer_resize(target_task Task_t, size uint32) Kern_return_t {
	result, callErr := tryMach_vm_deferred_reclamation_buffer_resize(target_task, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_inherit func(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_inheritance Vm_inherit_t) Kern_return_t
var _mach_vm_inheritErr error

func tryMach_vm_inherit(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _mach_vm_inherit == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_inherit", "10.4", _mach_vm_inheritErr)
	}
	return _mach_vm_inherit(target_task, address, size, new_inheritance), nil
}

// Mach_vm_inherit.
//
// See: https://developer.apple.com/documentation/kernel/1402141-mach_vm_inherit
func Mach_vm_inherit(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, new_inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryMach_vm_inherit(target_task, address, size, new_inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_machine_attribute func(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) Kern_return_t
var _mach_vm_machine_attributeErr error

func tryMach_vm_machine_attribute(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) (Kern_return_t, error) {
	if _mach_vm_machine_attribute == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_machine_attribute", "10.4", _mach_vm_machine_attributeErr)
	}
	return _mach_vm_machine_attribute(target_task, address, size, attribute, value), nil
}

// Mach_vm_machine_attribute.
//
// See: https://developer.apple.com/documentation/kernel/1402429-mach_vm_machine_attribute
func Mach_vm_machine_attribute(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) Kern_return_t {
	result, callErr := tryMach_vm_machine_attribute(target_task, address, size, attribute, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_map func(target_task uint32, address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, object Mem_entry_name_port_t, offset Memory_object_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _mach_vm_mapErr error

func tryMach_vm_map(target_task uint32, address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, object Mem_entry_name_port_t, offset Memory_object_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _mach_vm_map == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_map", "10.4", _mach_vm_mapErr)
	}
	return _mach_vm_map(target_task, address, size, mask, flags, object, offset, copy_, cur_protection, max_protection, inheritance), nil
}

// Mach_vm_map.
//
// See: https://developer.apple.com/documentation/kernel/1402481-mach_vm_map
func Mach_vm_map(target_task uint32, address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, object Mem_entry_name_port_t, offset Memory_object_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryMach_vm_map(target_task, address, size, mask, flags, object, offset, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_msync func(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, sync_flags Vm_sync_t) Kern_return_t
var _mach_vm_msyncErr error

func tryMach_vm_msync(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, sync_flags Vm_sync_t) (Kern_return_t, error) {
	if _mach_vm_msync == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_msync", "10.4", _mach_vm_msyncErr)
	}
	return _mach_vm_msync(target_task, address, size, sync_flags), nil
}

// Mach_vm_msync.
//
// See: https://developer.apple.com/documentation/kernel/1402328-mach_vm_msync
func Mach_vm_msync(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, sync_flags Vm_sync_t) Kern_return_t {
	result, callErr := tryMach_vm_msync(target_task, address, size, sync_flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_page_info func(target_task Vm_map_read_t, address Mach_vm_address_t, flavor Vm_page_info_flavor_t, info Vm_page_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_vm_page_infoErr error

func tryMach_vm_page_info(target_task Vm_map_read_t, address Mach_vm_address_t, flavor Vm_page_info_flavor_t, info Vm_page_info_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_vm_page_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_page_info", "10.6", _mach_vm_page_infoErr)
	}
	return _mach_vm_page_info(target_task, address, flavor, info, infoCnt), nil
}

// Mach_vm_page_info.
//
// See: https://developer.apple.com/documentation/kernel/1402504-mach_vm_page_info
func Mach_vm_page_info(target_task Vm_map_read_t, address Mach_vm_address_t, flavor Vm_page_info_flavor_t, info Vm_page_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_vm_page_info(target_task, address, flavor, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_page_query func(target_map Vm_map_read_t, offset Mach_vm_offset_t, disposition *Integer_t, ref_count *Integer_t) Kern_return_t
var _mach_vm_page_queryErr error

func tryMach_vm_page_query(target_map Vm_map_read_t, offset Mach_vm_offset_t, disposition *Integer_t, ref_count *Integer_t) (Kern_return_t, error) {
	if _mach_vm_page_query == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_page_query", "10.4", _mach_vm_page_queryErr)
	}
	return _mach_vm_page_query(target_map, offset, disposition, ref_count), nil
}

// Mach_vm_page_query.
//
// See: https://developer.apple.com/documentation/kernel/1402261-mach_vm_page_query
func Mach_vm_page_query(target_map Vm_map_read_t, offset Mach_vm_offset_t, disposition *Integer_t, ref_count *Integer_t) Kern_return_t {
	result, callErr := tryMach_vm_page_query(target_map, offset, disposition, ref_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_page_range_query func(target_map Vm_map_read_t, address Mach_vm_offset_t, size Mach_vm_size_t, dispositions Mach_vm_address_t, dispositions_count *Mach_vm_size_t) Kern_return_t
var _mach_vm_page_range_queryErr error

func tryMach_vm_page_range_query(target_map Vm_map_read_t, address Mach_vm_offset_t, size Mach_vm_size_t, dispositions Mach_vm_address_t, dispositions_count *Mach_vm_size_t) (Kern_return_t, error) {
	if _mach_vm_page_range_query == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_page_range_query", "10.13", _mach_vm_page_range_queryErr)
	}
	return _mach_vm_page_range_query(target_map, address, size, dispositions, dispositions_count), nil
}

// Mach_vm_page_range_query.
//
// See: https://developer.apple.com/documentation/kernel/2890784-mach_vm_page_range_query
func Mach_vm_page_range_query(target_map Vm_map_read_t, address Mach_vm_offset_t, size Mach_vm_size_t, dispositions Mach_vm_address_t, dispositions_count *Mach_vm_size_t) Kern_return_t {
	result, callErr := tryMach_vm_page_range_query(target_map, address, size, dispositions, dispositions_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_protect func(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) Kern_return_t
var _mach_vm_protectErr error

func tryMach_vm_protect(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) (Kern_return_t, error) {
	if _mach_vm_protect == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_protect", "10.4", _mach_vm_protectErr)
	}
	return _mach_vm_protect(target_task, address, size, set_maximum, new_protection), nil
}

// Mach_vm_protect.
//
// See: https://developer.apple.com/documentation/kernel/1402291-mach_vm_protect
func Mach_vm_protect(target_task uint32, address Mach_vm_address_t, size Mach_vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) Kern_return_t {
	result, callErr := tryMach_vm_protect(target_task, address, size, set_maximum, new_protection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_purgable_control func(target_task uint32, address Mach_vm_address_t, control Vm_purgable_t, state *int32) Kern_return_t
var _mach_vm_purgable_controlErr error

func tryMach_vm_purgable_control(target_task uint32, address Mach_vm_address_t, control Vm_purgable_t, state *int32) (Kern_return_t, error) {
	if _mach_vm_purgable_control == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_purgable_control", "10.5", _mach_vm_purgable_controlErr)
	}
	return _mach_vm_purgable_control(target_task, address, control, state), nil
}

// Mach_vm_purgable_control.
//
// See: https://developer.apple.com/documentation/kernel/1402224-mach_vm_purgable_control
func Mach_vm_purgable_control(target_task uint32, address Mach_vm_address_t, control Vm_purgable_t, state *int32) Kern_return_t {
	result, callErr := tryMach_vm_purgable_control(target_task, address, control, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_range_create func(target_task uint32, flavor Mach_vm_range_flavor_t, recipes Mach_vm_range_recipes_raw_t, recipesCnt Mach_msg_type_number_t) Kern_return_t
var _mach_vm_range_createErr error

func tryMach_vm_range_create(target_task uint32, flavor Mach_vm_range_flavor_t, recipes Mach_vm_range_recipes_raw_t, recipesCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_vm_range_create == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_range_create", "14.0", _mach_vm_range_createErr)
	}
	return _mach_vm_range_create(target_task, flavor, recipes, recipesCnt), nil
}

// Mach_vm_range_create.
//
// See: https://developer.apple.com/documentation/kernel/4168415-mach_vm_range_create
func Mach_vm_range_create(target_task uint32, flavor Mach_vm_range_flavor_t, recipes Mach_vm_range_recipes_raw_t, recipesCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_vm_range_create(target_task, flavor, recipes, recipesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_read func(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_vm_readErr error

func tryMach_vm_read(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_vm_read == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_read", "10.4", _mach_vm_readErr)
	}
	return _mach_vm_read(target_task, address, size, data, dataCnt), nil
}

// Mach_vm_read.
//
// See: https://developer.apple.com/documentation/kernel/1402405-mach_vm_read
func Mach_vm_read(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_vm_read(target_task, address, size, data, dataCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_read_list func(target_task Vm_map_read_t, data_list Mach_vm_read_entry_t, count Natural_t) Kern_return_t
var _mach_vm_read_listErr error

func tryMach_vm_read_list(target_task Vm_map_read_t, data_list Mach_vm_read_entry_t, count Natural_t) (Kern_return_t, error) {
	if _mach_vm_read_list == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_read_list", "10.4", _mach_vm_read_listErr)
	}
	return _mach_vm_read_list(target_task, data_list, count), nil
}

// Mach_vm_read_list.
//
// See: https://developer.apple.com/documentation/kernel/1402084-mach_vm_read_list
func Mach_vm_read_list(target_task Vm_map_read_t, data_list Mach_vm_read_entry_t, count Natural_t) Kern_return_t {
	result, callErr := tryMach_vm_read_list(target_task, data_list, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_read_overwrite func(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data Mach_vm_address_t, outsize *Mach_vm_size_t) Kern_return_t
var _mach_vm_read_overwriteErr error

func tryMach_vm_read_overwrite(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data Mach_vm_address_t, outsize *Mach_vm_size_t) (Kern_return_t, error) {
	if _mach_vm_read_overwrite == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_read_overwrite", "10.4", _mach_vm_read_overwriteErr)
	}
	return _mach_vm_read_overwrite(target_task, address, size, data, outsize), nil
}

// Mach_vm_read_overwrite.
//
// See: https://developer.apple.com/documentation/kernel/1402127-mach_vm_read_overwrite
func Mach_vm_read_overwrite(target_task Vm_map_read_t, address Mach_vm_address_t, size Mach_vm_size_t, data Mach_vm_address_t, outsize *Mach_vm_size_t) Kern_return_t {
	result, callErr := tryMach_vm_read_overwrite(target_task, address, size, data, outsize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_region func(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) Kern_return_t
var _mach_vm_regionErr error

func tryMach_vm_region(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) (Kern_return_t, error) {
	if _mach_vm_region == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_region", "10.4", _mach_vm_regionErr)
	}
	return _mach_vm_region(target_task, address, size, flavor, info, infoCnt, object_name), nil
}

// Mach_vm_region.
//
// See: https://developer.apple.com/documentation/kernel/1402149-mach_vm_region
func Mach_vm_region(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) Kern_return_t {
	result, callErr := tryMach_vm_region(target_task, address, size, flavor, info, infoCnt, object_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_region_recurse func(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_vm_region_recurseErr error

func tryMach_vm_region_recurse(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_vm_region_recurse == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_region_recurse", "10.4", _mach_vm_region_recurseErr)
	}
	return _mach_vm_region_recurse(target_task, address, size, nesting_depth, info, infoCnt), nil
}

// Mach_vm_region_recurse.
//
// See: https://developer.apple.com/documentation/kernel/1402114-mach_vm_region_recurse
func Mach_vm_region_recurse(target_task Vm_map_read_t, address *Mach_vm_address_t, size *Mach_vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_vm_region_recurse(target_task, address, size, nesting_depth, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_remap func(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task uint32, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _mach_vm_remapErr error

func tryMach_vm_remap(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task uint32, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _mach_vm_remap == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_remap", "10.4", _mach_vm_remapErr)
	}
	return _mach_vm_remap(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance), nil
}

// Mach_vm_remap.
//
// See: https://developer.apple.com/documentation/kernel/1402218-mach_vm_remap
func Mach_vm_remap(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task uint32, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryMach_vm_remap(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_remap_new func(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task Vm_map_read_t, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _mach_vm_remap_newErr error

func tryMach_vm_remap_new(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task Vm_map_read_t, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _mach_vm_remap_new == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_remap_new", "11.3", _mach_vm_remap_newErr)
	}
	return _mach_vm_remap_new(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance), nil
}

// Mach_vm_remap_new.
//
// See: https://developer.apple.com/documentation/kernel/3727986-mach_vm_remap_new
func Mach_vm_remap_new(target_task uint32, target_address *Mach_vm_address_t, size Mach_vm_size_t, mask Mach_vm_offset_t, flags int32, src_task Vm_map_read_t, src_address Mach_vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryMach_vm_remap_new(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_wire func(host_priv Host_priv_t, task uint32, address Mach_vm_address_t, size Mach_vm_size_t, desired_access Vm_prot_t) Kern_return_t
var _mach_vm_wireErr error

func tryMach_vm_wire(host_priv Host_priv_t, task uint32, address Mach_vm_address_t, size Mach_vm_size_t, desired_access Vm_prot_t) (Kern_return_t, error) {
	if _mach_vm_wire == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_wire", "10.4", _mach_vm_wireErr)
	}
	return _mach_vm_wire(host_priv, task, address, size, desired_access), nil
}

// Mach_vm_wire.
//
// See: https://developer.apple.com/documentation/kernel/1588962-mach_vm_wire
func Mach_vm_wire(host_priv Host_priv_t, task uint32, address Mach_vm_address_t, size Mach_vm_size_t, desired_access Vm_prot_t) Kern_return_t {
	result, callErr := tryMach_vm_wire(host_priv, task, address, size, desired_access)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_vm_write func(target_task uint32, address Mach_vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) Kern_return_t
var _mach_vm_writeErr error

func tryMach_vm_write(target_task uint32, address Mach_vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_vm_write == nil {
		return *new(Kern_return_t), symbolCallError("mach_vm_write", "10.4", _mach_vm_writeErr)
	}
	return _mach_vm_write(target_task, address, data, dataCnt), nil
}

// Mach_vm_write.
//
// See: https://developer.apple.com/documentation/kernel/1402070-mach_vm_write
func Mach_vm_write(target_task uint32, address Mach_vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_vm_write(target_task, address, data, dataCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_voucher_attr_command func(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, command Mach_voucher_attr_command_t, in_content Mach_voucher_attr_content_t, in_contentCnt Mach_msg_type_number_t, out_content Mach_voucher_attr_content_t, out_contentCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_voucher_attr_commandErr error

func tryMach_voucher_attr_command(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, command Mach_voucher_attr_command_t, in_content Mach_voucher_attr_content_t, in_contentCnt Mach_msg_type_number_t, out_content Mach_voucher_attr_content_t, out_contentCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_voucher_attr_command == nil {
		return *new(Kern_return_t), symbolCallError("mach_voucher_attr_command", "10.10", _mach_voucher_attr_commandErr)
	}
	return _mach_voucher_attr_command(voucher, key, command, in_content, in_contentCnt, out_content, out_contentCnt), nil
}

// Mach_voucher_attr_command.
//
// See: https://developer.apple.com/documentation/kernel/1410145-mach_voucher_attr_command
func Mach_voucher_attr_command(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, command Mach_voucher_attr_command_t, in_content Mach_voucher_attr_content_t, in_contentCnt Mach_msg_type_number_t, out_content Mach_voucher_attr_content_t, out_contentCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_voucher_attr_command(voucher, key, command, in_content, in_contentCnt, out_content, out_contentCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_voucher_debug_info func(task Ipc_space_read_t, voucher_name Mach_port_name_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_voucher_debug_infoErr error

func tryMach_voucher_debug_info(task Ipc_space_read_t, voucher_name Mach_port_name_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_voucher_debug_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_voucher_debug_info", "10.10", _mach_voucher_debug_infoErr)
	}
	return _mach_voucher_debug_info(task, voucher_name, recipes, recipesCnt), nil
}

// Mach_voucher_debug_info.
//
// See: https://developer.apple.com/documentation/kernel/1410149-mach_voucher_debug_info
func Mach_voucher_debug_info(task Ipc_space_read_t, voucher_name Mach_port_name_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_voucher_debug_info(task, voucher_name, recipes, recipesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_voucher_extract_all_attr_recipes func(voucher Ipc_voucher_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_voucher_extract_all_attr_recipesErr error

func tryMach_voucher_extract_all_attr_recipes(voucher Ipc_voucher_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_voucher_extract_all_attr_recipes == nil {
		return *new(Kern_return_t), symbolCallError("mach_voucher_extract_all_attr_recipes", "10.10", _mach_voucher_extract_all_attr_recipesErr)
	}
	return _mach_voucher_extract_all_attr_recipes(voucher, recipes, recipesCnt), nil
}

// Mach_voucher_extract_all_attr_recipes.
//
// See: https://developer.apple.com/documentation/kernel/1410119-mach_voucher_extract_all_attr_re
func Mach_voucher_extract_all_attr_recipes(voucher Ipc_voucher_t, recipes Mach_voucher_attr_raw_recipe_array_t, recipesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_voucher_extract_all_attr_recipes(voucher, recipes, recipesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_voucher_extract_attr_content func(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, content Mach_voucher_attr_content_t, contentCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_voucher_extract_attr_contentErr error

func tryMach_voucher_extract_attr_content(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, content Mach_voucher_attr_content_t, contentCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_voucher_extract_attr_content == nil {
		return *new(Kern_return_t), symbolCallError("mach_voucher_extract_attr_content", "10.10", _mach_voucher_extract_attr_contentErr)
	}
	return _mach_voucher_extract_attr_content(voucher, key, content, contentCnt), nil
}

// Mach_voucher_extract_attr_content.
//
// See: https://developer.apple.com/documentation/kernel/1410080-mach_voucher_extract_attr_conten
func Mach_voucher_extract_attr_content(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, content Mach_voucher_attr_content_t, contentCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_voucher_extract_attr_content(voucher, key, content, contentCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_voucher_extract_attr_recipe func(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, recipe Mach_voucher_attr_raw_recipe_t, recipeCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_voucher_extract_attr_recipeErr error

func tryMach_voucher_extract_attr_recipe(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, recipe Mach_voucher_attr_raw_recipe_t, recipeCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_voucher_extract_attr_recipe == nil {
		return *new(Kern_return_t), symbolCallError("mach_voucher_extract_attr_recipe", "10.10", _mach_voucher_extract_attr_recipeErr)
	}
	return _mach_voucher_extract_attr_recipe(voucher, key, recipe, recipeCnt), nil
}

// Mach_voucher_extract_attr_recipe.
//
// See: https://developer.apple.com/documentation/kernel/1410137-mach_voucher_extract_attr_recipe
func Mach_voucher_extract_attr_recipe(voucher Ipc_voucher_t, key Mach_voucher_attr_key_t, recipe Mach_voucher_attr_raw_recipe_t, recipeCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_voucher_extract_attr_recipe(voucher, key, recipe, recipeCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_force_gc func(host Host_t) Kern_return_t
var _mach_zone_force_gcErr error

func tryMach_zone_force_gc(host Host_t) (Kern_return_t, error) {
	if _mach_zone_force_gc == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_force_gc", "10.8", _mach_zone_force_gcErr)
	}
	return _mach_zone_force_gc(host), nil
}

// Mach_zone_force_gc.
//
// See: https://developer.apple.com/documentation/kernel/1502550-mach_zone_force_gc
func Mach_zone_force_gc(host Host_t) Kern_return_t {
	result, callErr := tryMach_zone_force_gc(host)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_get_btlog_records func(host Host_priv_t, name Mach_zone_name_t, recs *Zone_btrecord_array_t, recsCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_zone_get_btlog_recordsErr error

func tryMach_zone_get_btlog_records(host Host_priv_t, name Mach_zone_name_t, recs *Zone_btrecord_array_t, recsCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_zone_get_btlog_records == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_get_btlog_records", "10.14", _mach_zone_get_btlog_recordsErr)
	}
	return _mach_zone_get_btlog_records(host, name, recs, recsCnt), nil
}

// Mach_zone_get_btlog_records.
//
// See: https://developer.apple.com/documentation/kernel/2977296-mach_zone_get_btlog_records
func Mach_zone_get_btlog_records(host Host_priv_t, name Mach_zone_name_t, recs *Zone_btrecord_array_t, recsCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_zone_get_btlog_records(host, name, recs, recsCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_get_zlog_zones func(host Host_priv_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_zone_get_zlog_zonesErr error

func tryMach_zone_get_zlog_zones(host Host_priv_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_zone_get_zlog_zones == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_get_zlog_zones", "10.14", _mach_zone_get_zlog_zonesErr)
	}
	return _mach_zone_get_zlog_zones(host, names, namesCnt), nil
}

// Mach_zone_get_zlog_zones.
//
// See: https://developer.apple.com/documentation/kernel/2977297-mach_zone_get_zlog_zones
func Mach_zone_get_zlog_zones(host Host_priv_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_zone_get_zlog_zones(host, names, namesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_info func(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _mach_zone_infoErr error

func tryMach_zone_info(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _mach_zone_info == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_info", "10.7", _mach_zone_infoErr)
	}
	return _mach_zone_info(host, names, namesCnt, info, infoCnt), nil
}

// Mach_zone_info.
//
// See: https://developer.apple.com/documentation/kernel/1502472-mach_zone_info
func Mach_zone_info(host uint32, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Mach_zone_info_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryMach_zone_info(host, names, namesCnt, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_info_for_largest_zone func(host Host_priv_t, name *Mach_zone_name_t, info *Mach_zone_info_t) Kern_return_t
var _mach_zone_info_for_largest_zoneErr error

func tryMach_zone_info_for_largest_zone(host Host_priv_t, name *Mach_zone_name_t, info *Mach_zone_info_t) (Kern_return_t, error) {
	if _mach_zone_info_for_largest_zone == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_info_for_largest_zone", "10.13.4", _mach_zone_info_for_largest_zoneErr)
	}
	return _mach_zone_info_for_largest_zone(host, name, info), nil
}

// Mach_zone_info_for_largest_zone.
//
// See: https://developer.apple.com/documentation/kernel/2937923-mach_zone_info_for_largest_zone
func Mach_zone_info_for_largest_zone(host Host_priv_t, name *Mach_zone_name_t, info *Mach_zone_info_t) Kern_return_t {
	result, callErr := tryMach_zone_info_for_largest_zone(host, name, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mach_zone_info_for_zone func(host Host_priv_t, name Mach_zone_name_t, info *Mach_zone_info_t) Kern_return_t
var _mach_zone_info_for_zoneErr error

func tryMach_zone_info_for_zone(host Host_priv_t, name Mach_zone_name_t, info *Mach_zone_info_t) (Kern_return_t, error) {
	if _mach_zone_info_for_zone == nil {
		return *new(Kern_return_t), symbolCallError("mach_zone_info_for_zone", "10.13.4", _mach_zone_info_for_zoneErr)
	}
	return _mach_zone_info_for_zone(host, name, info), nil
}

// Mach_zone_info_for_zone.
//
// See: https://developer.apple.com/documentation/kernel/2937922-mach_zone_info_for_zone
func Mach_zone_info_for_zone(host Host_priv_t, name Mach_zone_name_t, info *Mach_zone_info_t) Kern_return_t {
	result, callErr := tryMach_zone_info_for_zone(host, name, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memchr func(arg0 unsafe.Pointer, arg1 int32, arg2 uintptr) unsafe.Pointer
var _memchrErr error

func tryMemchr(arg0 unsafe.Pointer, arg1 int32, arg2 uintptr) (unsafe.Pointer, error) {
	if _memchr == nil {
		return nil, symbolCallError("memchr", "10.9", _memchrErr)
	}
	return _memchr(arg0, arg1, arg2), nil
}

// Memchr.
//
// See: https://developer.apple.com/documentation/kernel/1441105-memchr
func Memchr(arg0 unsafe.Pointer, arg1 int32, arg2 uintptr) unsafe.Pointer {
	result, callErr := tryMemchr(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memcmp func(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) int32
var _memcmpErr error

func tryMemcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) (int32, error) {
	if _memcmp == nil {
		return 0, symbolCallError("memcmp", "10.0", _memcmpErr)
	}
	return _memcmp(s1, s2, n), nil
}

// Memcmp.
//
// See: https://developer.apple.com/documentation/kernel/1579327-memcmp
func Memcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, n uintptr) int32 {
	result, callErr := tryMemcmp(s1, s2, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memcpy func(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) unsafe.Pointer
var _memcpyErr error

func tryMemcpy(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) (unsafe.Pointer, error) {
	if _memcpy == nil {
		return nil, symbolCallError("memcpy", "10.0", _memcpyErr)
	}
	return _memcpy(dst, src, n), nil
}

// Memcpy.
//
// See: https://developer.apple.com/documentation/kernel/1579338-memcpy
func Memcpy(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) unsafe.Pointer {
	result, callErr := tryMemcpy(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memmove func(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) unsafe.Pointer
var _memmoveErr error

func tryMemmove(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) (unsafe.Pointer, error) {
	if _memmove == nil {
		return nil, symbolCallError("memmove", "10.0", _memmoveErr)
	}
	return _memmove(dst, src, n), nil
}

// Memmove.
//
// See: https://developer.apple.com/documentation/kernel/1579336-memmove
func Memmove(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) unsafe.Pointer {
	result, callErr := tryMemmove(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memset func(s unsafe.Pointer, c int32, n uintptr) unsafe.Pointer
var _memsetErr error

func tryMemset(s unsafe.Pointer, c int32, n uintptr) (unsafe.Pointer, error) {
	if _memset == nil {
		return nil, symbolCallError("memset", "10.0", _memsetErr)
	}
	return _memset(s, c, n), nil
}

// Memset.
//
// See: https://developer.apple.com/documentation/kernel/1579332-memset
func Memset(s unsafe.Pointer, c int32, n uintptr) unsafe.Pointer {
	result, callErr := tryMemset(s, c, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _memset_s func(s unsafe.Pointer, smax uintptr, c int32, n uintptr) int32
var _memset_sErr error

func tryMemset_s(s unsafe.Pointer, smax uintptr, c int32, n uintptr) (int32, error) {
	if _memset_s == nil {
		return 0, symbolCallError("memset_s", "10.13", _memset_sErr)
	}
	return _memset_s(s, smax, c, n), nil
}

// Memset_s.
//
// See: https://developer.apple.com/documentation/kernel/2876438-memset_s
func Memset_s(s unsafe.Pointer, smax uintptr, c int32, n uintptr) int32 {
	result, callErr := tryMemset_s(s, smax, c, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mig_allocate func(arg0 *Vm_address_t, arg1 Vm_size_t)
var _mig_allocateErr error

func tryMig_allocate(arg0 *Vm_address_t, arg1 Vm_size_t) error {
	if _mig_allocate == nil {
		return symbolCallError("mig_allocate", "10.4", _mig_allocateErr)
	}
	_mig_allocate(arg0, arg1)
	return nil
}

// Mig_allocate.
//
// See: https://developer.apple.com/documentation/kernel/1558649-mig_allocate
func Mig_allocate(arg0 *Vm_address_t, arg1 Vm_size_t) {
	if callErr := tryMig_allocate(arg0, arg1); callErr != nil {
		panic(callErr)
	}
}

var _mig_dealloc_reply_port func(reply_port uint32)
var _mig_dealloc_reply_portErr error

func tryMig_dealloc_reply_port(reply_port uint32) error {
	if _mig_dealloc_reply_port == nil {
		return symbolCallError("mig_dealloc_reply_port", "10.0", _mig_dealloc_reply_portErr)
	}
	_mig_dealloc_reply_port(reply_port)
	return nil
}

// Mig_dealloc_reply_port.
//
// See: https://developer.apple.com/documentation/kernel/1558668-mig_dealloc_reply_port
func Mig_dealloc_reply_port(reply_port uint32) {
	if callErr := tryMig_dealloc_reply_port(reply_port); callErr != nil {
		panic(callErr)
	}
}

var _mig_deallocate func(arg0 Vm_address_t, arg1 Vm_size_t)
var _mig_deallocateErr error

func tryMig_deallocate(arg0 Vm_address_t, arg1 Vm_size_t) error {
	if _mig_deallocate == nil {
		return symbolCallError("mig_deallocate", "10.4", _mig_deallocateErr)
	}
	_mig_deallocate(arg0, arg1)
	return nil
}

// Mig_deallocate.
//
// See: https://developer.apple.com/documentation/kernel/1558654-mig_deallocate
func Mig_deallocate(arg0 Vm_address_t, arg1 Vm_size_t) {
	if callErr := tryMig_deallocate(arg0, arg1); callErr != nil {
		panic(callErr)
	}
}

var _mig_get_reply_port func() uint32
var _mig_get_reply_portErr error

func tryMig_get_reply_port() (uint32, error) {
	if _mig_get_reply_port == nil {
		return 0, symbolCallError("mig_get_reply_port", "10.0", _mig_get_reply_portErr)
	}
	return _mig_get_reply_port(), nil
}

// Mig_get_reply_port.
//
// See: https://developer.apple.com/documentation/kernel/1558662-mig_get_reply_port
func Mig_get_reply_port() uint32 {
	result, callErr := tryMig_get_reply_port()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mig_put_reply_port func(reply_port uint32)
var _mig_put_reply_portErr error

func tryMig_put_reply_port(reply_port uint32) error {
	if _mig_put_reply_port == nil {
		return symbolCallError("mig_put_reply_port", "10.0", _mig_put_reply_portErr)
	}
	_mig_put_reply_port(reply_port)
	return nil
}

// Mig_put_reply_port.
//
// See: https://developer.apple.com/documentation/kernel/1558665-mig_put_reply_port
func Mig_put_reply_port(reply_port uint32) {
	if callErr := tryMig_put_reply_port(reply_port); callErr != nil {
		panic(callErr)
	}
}

var _mig_strncpy func(dest *byte, src string, len_ int32) int32
var _mig_strncpyErr error

func tryMig_strncpy(dest *byte, src string, len_ int32) (int32, error) {
	if _mig_strncpy == nil {
		return 0, symbolCallError("mig_strncpy", "10.0", _mig_strncpyErr)
	}
	return _mig_strncpy(dest, src, len_), nil
}

// Mig_strncpy.
//
// See: https://developer.apple.com/documentation/kernel/1558639-mig_strncpy
func Mig_strncpy(dest *byte, src string, len_ int32) int32 {
	result, callErr := tryMig_strncpy(dest, src, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _mig_strncpy_zerofill func(dest *byte, src string, len_ int32) int32
var _mig_strncpy_zerofillErr error

func tryMig_strncpy_zerofill(dest *byte, src string, len_ int32) (int32, error) {
	if _mig_strncpy_zerofill == nil {
		return 0, symbolCallError("mig_strncpy_zerofill", "10.12", _mig_strncpy_zerofillErr)
	}
	return _mig_strncpy_zerofill(dest, src, len_), nil
}

// Mig_strncpy_zerofill.
//
// See: https://developer.apple.com/documentation/kernel/1645207-mig_strncpy_zerofill
func Mig_strncpy_zerofill(dest *byte, src string, len_ int32) int32 {
	result, callErr := tryMig_strncpy_zerofill(dest, src, len_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _modf func(arg0 float64, arg1 *float64) float64
var _modfErr error

func tryModf(arg0 float64, arg1 []float64) (float64, error) {
	if _modf == nil {
		return 0.0, symbolCallError("modf", "10.10", _modfErr)
	}
	return _modf(arg0, unsafe.SliceData(arg1)), nil
}

// Modf.
//
// See: https://developer.apple.com/documentation/kernel/1557173-modf
func Modf(arg0 float64, arg1 []float64) float64 {
	result, callErr := tryModf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _modff func(arg0 float32, arg1 *float32) float32
var _modffErr error

func tryModff(arg0 float32, arg1 []float32) (float32, error) {
	if _modff == nil {
		return 0.0, symbolCallError("modff", "10.10", _modffErr)
	}
	return _modff(arg0, unsafe.SliceData(arg1)), nil
}

// Modff.
//
// See: https://developer.apple.com/documentation/kernel/1557317-modff
func Modff(arg0 float32, arg1 []float32) float32 {
	result, callErr := tryModff(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _modfl func(arg0 float64, arg1 *float64) float64
var _modflErr error

func tryModfl(arg0 float64, arg1 []float64) (float64, error) {
	if _modfl == nil {
		return 0.0, symbolCallError("modfl", "10.10", _modflErr)
	}
	return _modfl(arg0, unsafe.SliceData(arg1)), nil
}

// Modfl.
//
// See: https://developer.apple.com/documentation/kernel/1557161-modfl
func Modfl(arg0 float64, arg1 []float64) float64 {
	result, callErr := tryModfl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nan func(arg0 string) float64
var _nanErr error

func tryNan(arg0 string) (float64, error) {
	if _nan == nil {
		return 0.0, symbolCallError("nan", "10.10", _nanErr)
	}
	return _nan(arg0), nil
}

// Nan.
//
// See: https://developer.apple.com/documentation/kernel/1557310-nan
func Nan(arg0 string) float64 {
	result, callErr := tryNan(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nanf func(arg0 string) float32
var _nanfErr error

func tryNanf(arg0 string) (float32, error) {
	if _nanf == nil {
		return 0.0, symbolCallError("nanf", "10.10", _nanfErr)
	}
	return _nanf(arg0), nil
}

// Nanf.
//
// See: https://developer.apple.com/documentation/kernel/1557309-nanf
func Nanf(arg0 string) float32 {
	result, callErr := tryNanf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nanl func(arg0 string) float64
var _nanlErr error

func tryNanl(arg0 string) (float64, error) {
	if _nanl == nil {
		return 0.0, symbolCallError("nanl", "10.10", _nanlErr)
	}
	return _nanl(arg0), nil
}

// Nanl.
//
// See: https://developer.apple.com/documentation/kernel/1557311-nanl
func Nanl(arg0 string) float64 {
	result, callErr := tryNanl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nearbyint func(arg0 float64) float64
var _nearbyintErr error

func tryNearbyint(arg0 float64) (float64, error) {
	if _nearbyint == nil {
		return 0.0, symbolCallError("nearbyint", "10.10", _nearbyintErr)
	}
	return _nearbyint(arg0), nil
}

// Nearbyint.
//
// See: https://developer.apple.com/documentation/kernel/1557212-nearbyint
func Nearbyint(arg0 float64) float64 {
	result, callErr := tryNearbyint(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nearbyintf func(arg0 float32) float32
var _nearbyintfErr error

func tryNearbyintf(arg0 float32) (float32, error) {
	if _nearbyintf == nil {
		return 0.0, symbolCallError("nearbyintf", "10.10", _nearbyintfErr)
	}
	return _nearbyintf(arg0), nil
}

// Nearbyintf.
//
// See: https://developer.apple.com/documentation/kernel/1557346-nearbyintf
func Nearbyintf(arg0 float32) float32 {
	result, callErr := tryNearbyintf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nearbyintl func(arg0 float64) float64
var _nearbyintlErr error

func tryNearbyintl(arg0 float64) (float64, error) {
	if _nearbyintl == nil {
		return 0.0, symbolCallError("nearbyintl", "10.10", _nearbyintlErr)
	}
	return _nearbyintl(arg0), nil
}

// Nearbyintl.
//
// See: https://developer.apple.com/documentation/kernel/1557159-nearbyintl
func Nearbyintl(arg0 float64) float64 {
	result, callErr := tryNearbyintl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nextafter func(arg0 float64, arg1 float64) float64
var _nextafterErr error

func tryNextafter(arg0 float64, arg1 float64) (float64, error) {
	if _nextafter == nil {
		return 0.0, symbolCallError("nextafter", "10.10", _nextafterErr)
	}
	return _nextafter(arg0, arg1), nil
}

// Nextafter.
//
// See: https://developer.apple.com/documentation/kernel/1557351-nextafter
func Nextafter(arg0 float64, arg1 float64) float64 {
	result, callErr := tryNextafter(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nextafterf func(arg0 float32, arg1 float32) float32
var _nextafterfErr error

func tryNextafterf(arg0 float32, arg1 float32) (float32, error) {
	if _nextafterf == nil {
		return 0.0, symbolCallError("nextafterf", "10.10", _nextafterfErr)
	}
	return _nextafterf(arg0, arg1), nil
}

// Nextafterf.
//
// See: https://developer.apple.com/documentation/kernel/1557315-nextafterf
func Nextafterf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryNextafterf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nextafterl func(arg0 float64, arg1 float64) float64
var _nextafterlErr error

func tryNextafterl(arg0 float64, arg1 float64) (float64, error) {
	if _nextafterl == nil {
		return 0.0, symbolCallError("nextafterl", "10.10", _nextafterlErr)
	}
	return _nextafterl(arg0, arg1), nil
}

// Nextafterl.
//
// See: https://developer.apple.com/documentation/kernel/1557308-nextafterl
func Nextafterl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryNextafterl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nexttoward func(arg0 float64, arg1 float64) float64
var _nexttowardErr error

func tryNexttoward(arg0 float64, arg1 float64) (float64, error) {
	if _nexttoward == nil {
		return 0.0, symbolCallError("nexttoward", "10.10", _nexttowardErr)
	}
	return _nexttoward(arg0, arg1), nil
}

// Nexttoward.
//
// See: https://developer.apple.com/documentation/kernel/1557273-nexttoward
func Nexttoward(arg0 float64, arg1 float64) float64 {
	result, callErr := tryNexttoward(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nexttowardf func(arg0 float32, arg1 float64) float32
var _nexttowardfErr error

func tryNexttowardf(arg0 float32, arg1 float64) (float32, error) {
	if _nexttowardf == nil {
		return 0.0, symbolCallError("nexttowardf", "10.10", _nexttowardfErr)
	}
	return _nexttowardf(arg0, arg1), nil
}

// Nexttowardf.
//
// See: https://developer.apple.com/documentation/kernel/1557290-nexttowardf
func Nexttowardf(arg0 float32, arg1 float64) float32 {
	result, callErr := tryNexttowardf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nexttowardl func(arg0 float64, arg1 float64) float64
var _nexttowardlErr error

func tryNexttowardl(arg0 float64, arg1 float64) (float64, error) {
	if _nexttowardl == nil {
		return 0.0, symbolCallError("nexttowardl", "10.10", _nexttowardlErr)
	}
	return _nexttowardl(arg0, arg1), nil
}

// Nexttowardl.
//
// See: https://developer.apple.com/documentation/kernel/1557238-nexttowardl
func Nexttowardl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryNexttowardl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nfsclnt func(request uint, argstructp unsafe.Pointer) int32
var _nfsclntErr error

func tryNfsclnt(request uint, argstructp unsafe.Pointer) (int32, error) {
	if _nfsclnt == nil {
		return 0, symbolCallError("nfsclnt", "13.0", _nfsclntErr)
	}
	return _nfsclnt(request, argstructp), nil
}

// Nfsclnt.
//
// See: https://developer.apple.com/documentation/kernel/3964948-nfsclnt
func Nfsclnt(request uint, argstructp unsafe.Pointer) int32 {
	result, callErr := tryNfsclnt(request, argstructp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_log_create func(subsystem string, category string) Os_log_t
var _os_log_createErr error

func tryOs_log_create(subsystem string, category string) (Os_log_t, error) {
	if _os_log_create == nil {
		return *new(Os_log_t), symbolCallError("os_log_create", "10.12", _os_log_createErr)
	}
	return _os_log_create(subsystem, category), nil
}

// Os_log_create creates a custom log object, to be passed to logging functions for sending messages to the logging system.
//
// See: https://developer.apple.com/documentation/kernel/1643798-os_log_create
func Os_log_create(subsystem string, category string) Os_log_t {
	result, callErr := tryOs_log_create(subsystem, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _os_release func(object unsafe.Pointer)
var _os_releaseErr error

func tryOs_release(object unsafe.Pointer) error {
	if _os_release == nil {
		return symbolCallError("os_release", "10.12", _os_releaseErr)
	}
	_os_release(object)
	return nil
}

// Os_release.
//
// See: https://developer.apple.com/documentation/kernel/1646596-os_release
func Os_release(object unsafe.Pointer) {
	if callErr := tryOs_release(object); callErr != nil {
		panic(callErr)
	}
}

var _os_retain func(object unsafe.Pointer) unsafe.Pointer
var _os_retainErr error

func tryOs_retain(object unsafe.Pointer) (unsafe.Pointer, error) {
	if _os_retain == nil {
		return nil, symbolCallError("os_retain", "10.12", _os_retainErr)
	}
	return _os_retain(object), nil
}

// Os_retain.
//
// See: https://developer.apple.com/documentation/kernel/1646602-os_retain
func Os_retain(object unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryOs_retain(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _panic func(string_ string)
var _panicErr error

func tryPanic(string_ string) error {
	if _panic == nil {
		return symbolCallError("panic", "10.0", _panicErr)
	}
	_panic(string_)
	return nil
}

// Panic.
//
// See: https://developer.apple.com/documentation/kernel/1551635-panic
func Panic(string_ string) {
	if callErr := tryPanic(string_); callErr != nil {
		panic(callErr)
	}
}

var _pow func(arg0 float64, arg1 float64) float64
var _powErr error

func tryPow(arg0 float64, arg1 float64) (float64, error) {
	if _pow == nil {
		return 0.0, symbolCallError("pow", "10.10", _powErr)
	}
	return _pow(arg0, arg1), nil
}

// Pow.
//
// See: https://developer.apple.com/documentation/kernel/1557302-pow
func Pow(arg0 float64, arg1 float64) float64 {
	result, callErr := tryPow(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _powf func(arg0 float32, arg1 float32) float32
var _powfErr error

func tryPowf(arg0 float32, arg1 float32) (float32, error) {
	if _powf == nil {
		return 0.0, symbolCallError("powf", "10.10", _powfErr)
	}
	return _powf(arg0, arg1), nil
}

// Powf.
//
// See: https://developer.apple.com/documentation/kernel/1557297-powf
func Powf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryPowf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _powl func(arg0 float64, arg1 float64) float64
var _powlErr error

func tryPowl(arg0 float64, arg1 float64) (float64, error) {
	if _powl == nil {
		return 0.0, symbolCallError("powl", "10.10", _powlErr)
	}
	return _powl(arg0, arg1), nil
}

// Powl.
//
// See: https://developer.apple.com/documentation/kernel/1557343-powl
func Powl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryPowl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _printf func(arg0 string) int32
var _printfErr error

func tryPrintf(arg0 string) (int32, error) {
	if _printf == nil {
		return 0, symbolCallError("printf", "10.0", _printfErr)
	}
	return _printf(arg0), nil
}

// Printf.
//
// See: https://developer.apple.com/documentation/kernel/1441098-printf
func Printf(arg0 string) int32 {
	result, callErr := tryPrintf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _proc_name func(pid int32, buf *byte, size int32)
var _proc_nameErr error

func tryProc_name(pid int32, buf *byte, size int32) error {
	if _proc_name == nil {
		return symbolCallError("proc_name", "10.4", _proc_nameErr)
	}
	_proc_name(pid, buf, size)
	return nil
}

// Proc_name.
//
// See: https://developer.apple.com/documentation/kernel/1488959-proc_name
func Proc_name(pid int32, buf *byte, size int32) {
	if callErr := tryProc_name(pid, buf, size); callErr != nil {
		panic(callErr)
	}
}

var _processor_assign func(processor Processor_t, new_set Processor_set_t, wait Boolean_t) Kern_return_t
var _processor_assignErr error

func tryProcessor_assign(processor Processor_t, new_set Processor_set_t, wait Boolean_t) (Kern_return_t, error) {
	if _processor_assign == nil {
		return *new(Kern_return_t), symbolCallError("processor_assign", "10.0", _processor_assignErr)
	}
	return _processor_assign(processor, new_set, wait), nil
}

// Processor_assign.
//
// See: https://developer.apple.com/documentation/kernel/1409351-processor_assign
func Processor_assign(processor Processor_t, new_set Processor_set_t, wait Boolean_t) Kern_return_t {
	result, callErr := tryProcessor_assign(processor, new_set, wait)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_control func(processor Processor_t, processor_cmd Processor_info_t, processor_cmdCnt Mach_msg_type_number_t) Kern_return_t
var _processor_controlErr error

func tryProcessor_control(processor Processor_t, processor_cmd Processor_info_t, processor_cmdCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_control == nil {
		return *new(Kern_return_t), symbolCallError("processor_control", "10.0", _processor_controlErr)
	}
	return _processor_control(processor, processor_cmd, processor_cmdCnt), nil
}

// Processor_control.
//
// See: https://developer.apple.com/documentation/kernel/1409386-processor_control
func Processor_control(processor Processor_t, processor_cmd Processor_info_t, processor_cmdCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_control(processor, processor_cmd, processor_cmdCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_exit func(processor Processor_t) Kern_return_t
var _processor_exitErr error

func tryProcessor_exit(processor Processor_t) (Kern_return_t, error) {
	if _processor_exit == nil {
		return *new(Kern_return_t), symbolCallError("processor_exit", "10.0", _processor_exitErr)
	}
	return _processor_exit(processor), nil
}

// Processor_exit.
//
// See: https://developer.apple.com/documentation/kernel/1409353-processor_exit
func Processor_exit(processor Processor_t) Kern_return_t {
	result, callErr := tryProcessor_exit(processor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_get_assignment func(processor Processor_t, assigned_set *Processor_set_name_t) Kern_return_t
var _processor_get_assignmentErr error

func tryProcessor_get_assignment(processor Processor_t, assigned_set *Processor_set_name_t) (Kern_return_t, error) {
	if _processor_get_assignment == nil {
		return *new(Kern_return_t), symbolCallError("processor_get_assignment", "10.0", _processor_get_assignmentErr)
	}
	return _processor_get_assignment(processor, assigned_set), nil
}

// Processor_get_assignment.
//
// See: https://developer.apple.com/documentation/kernel/1409365-processor_get_assignment
func Processor_get_assignment(processor Processor_t, assigned_set *Processor_set_name_t) Kern_return_t {
	result, callErr := tryProcessor_get_assignment(processor, assigned_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_info func(processor Processor_t, flavor Processor_flavor_t, host *Host_t, processor_info_out Processor_info_t, processor_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_infoErr error

func tryProcessor_info(processor Processor_t, flavor Processor_flavor_t, host *Host_t, processor_info_out Processor_info_t, processor_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_info == nil {
		return *new(Kern_return_t), symbolCallError("processor_info", "10.0", _processor_infoErr)
	}
	return _processor_info(processor, flavor, host, processor_info_out, processor_info_outCnt), nil
}

// Processor_info.
//
// See: https://developer.apple.com/documentation/kernel/1409385-processor_info
func Processor_info(processor Processor_t, flavor Processor_flavor_t, host *Host_t, processor_info_out Processor_info_t, processor_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_info(processor, flavor, host, processor_info_out, processor_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_create func(host Host_t, new_set *Processor_set_t, new_name *Processor_set_name_t) Kern_return_t
var _processor_set_createErr error

func tryProcessor_set_create(host Host_t, new_set *Processor_set_t, new_name *Processor_set_name_t) (Kern_return_t, error) {
	if _processor_set_create == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_create", "10.0", _processor_set_createErr)
	}
	return _processor_set_create(host, new_set, new_name), nil
}

// Processor_set_create.
//
// See: https://developer.apple.com/documentation/kernel/1502809-processor_set_create
func Processor_set_create(host Host_t, new_set *Processor_set_t, new_name *Processor_set_name_t) Kern_return_t {
	result, callErr := tryProcessor_set_create(host, new_set, new_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_default func(host Host_t, default_set *Processor_set_name_t) Kern_return_t
var _processor_set_defaultErr error

func tryProcessor_set_default(host Host_t, default_set *Processor_set_name_t) (Kern_return_t, error) {
	if _processor_set_default == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_default", "10.0", _processor_set_defaultErr)
	}
	return _processor_set_default(host, default_set), nil
}

// Processor_set_default.
//
// See: https://developer.apple.com/documentation/kernel/1502766-processor_set_default
func Processor_set_default(host Host_t, default_set *Processor_set_name_t) Kern_return_t {
	result, callErr := tryProcessor_set_default(host, default_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_destroy func(set Processor_set_t) Kern_return_t
var _processor_set_destroyErr error

func tryProcessor_set_destroy(set Processor_set_t) (Kern_return_t, error) {
	if _processor_set_destroy == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_destroy", "10.0", _processor_set_destroyErr)
	}
	return _processor_set_destroy(set), nil
}

// Processor_set_destroy.
//
// See: https://developer.apple.com/documentation/kernel/1503645-processor_set_destroy
func Processor_set_destroy(set Processor_set_t) Kern_return_t {
	result, callErr := tryProcessor_set_destroy(set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_info func(set_name Processor_set_name_t, flavor int32, host *Host_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_set_infoErr error

func tryProcessor_set_info(set_name Processor_set_name_t, flavor int32, host *Host_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_set_info == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_info", "10.0", _processor_set_infoErr)
	}
	return _processor_set_info(set_name, flavor, host, info_out, info_outCnt), nil
}

// Processor_set_info.
//
// See: https://developer.apple.com/documentation/kernel/1503679-processor_set_info
func Processor_set_info(set_name Processor_set_name_t, flavor int32, host *Host_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_set_info(set_name, flavor, host, info_out, info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_max_priority func(processor_set Processor_set_t, max_priority int32, change_threads Boolean_t) Kern_return_t
var _processor_set_max_priorityErr error

func tryProcessor_set_max_priority(processor_set Processor_set_t, max_priority int32, change_threads Boolean_t) (Kern_return_t, error) {
	if _processor_set_max_priority == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_max_priority", "10.0", _processor_set_max_priorityErr)
	}
	return _processor_set_max_priority(processor_set, max_priority, change_threads), nil
}

// Processor_set_max_priority.
//
// See: https://developer.apple.com/documentation/kernel/1503640-processor_set_max_priority
func Processor_set_max_priority(processor_set Processor_set_t, max_priority int32, change_threads Boolean_t) Kern_return_t {
	result, callErr := tryProcessor_set_max_priority(processor_set, max_priority, change_threads)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_policy_control func(pset Processor_set_t, flavor Processor_set_flavor_t, policy_info Processor_set_info_t, policy_infoCnt Mach_msg_type_number_t, change Boolean_t) Kern_return_t
var _processor_set_policy_controlErr error

func tryProcessor_set_policy_control(pset Processor_set_t, flavor Processor_set_flavor_t, policy_info Processor_set_info_t, policy_infoCnt Mach_msg_type_number_t, change Boolean_t) (Kern_return_t, error) {
	if _processor_set_policy_control == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_policy_control", "10.0", _processor_set_policy_controlErr)
	}
	return _processor_set_policy_control(pset, flavor, policy_info, policy_infoCnt, change), nil
}

// Processor_set_policy_control.
//
// See: https://developer.apple.com/documentation/kernel/1503688-processor_set_policy_control
func Processor_set_policy_control(pset Processor_set_t, flavor Processor_set_flavor_t, policy_info Processor_set_info_t, policy_infoCnt Mach_msg_type_number_t, change Boolean_t) Kern_return_t {
	result, callErr := tryProcessor_set_policy_control(pset, flavor, policy_info, policy_infoCnt, change)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_policy_disable func(processor_set Processor_set_t, policy int32, change_threads Boolean_t) Kern_return_t
var _processor_set_policy_disableErr error

func tryProcessor_set_policy_disable(processor_set Processor_set_t, policy int32, change_threads Boolean_t) (Kern_return_t, error) {
	if _processor_set_policy_disable == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_policy_disable", "10.0", _processor_set_policy_disableErr)
	}
	return _processor_set_policy_disable(processor_set, policy, change_threads), nil
}

// Processor_set_policy_disable.
//
// See: https://developer.apple.com/documentation/kernel/1503665-processor_set_policy_disable
func Processor_set_policy_disable(processor_set Processor_set_t, policy int32, change_threads Boolean_t) Kern_return_t {
	result, callErr := tryProcessor_set_policy_disable(processor_set, policy, change_threads)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_policy_enable func(processor_set Processor_set_t, policy int32) Kern_return_t
var _processor_set_policy_enableErr error

func tryProcessor_set_policy_enable(processor_set Processor_set_t, policy int32) (Kern_return_t, error) {
	if _processor_set_policy_enable == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_policy_enable", "10.0", _processor_set_policy_enableErr)
	}
	return _processor_set_policy_enable(processor_set, policy), nil
}

// Processor_set_policy_enable.
//
// See: https://developer.apple.com/documentation/kernel/1503657-processor_set_policy_enable
func Processor_set_policy_enable(processor_set Processor_set_t, policy int32) Kern_return_t {
	result, callErr := tryProcessor_set_policy_enable(processor_set, policy)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_stack_usage func(pset Processor_set_t, ltotal *uint32, space *Vm_size_t, resident *Vm_size_t, maxusage *Vm_size_t, maxstack *Vm_offset_t) Kern_return_t
var _processor_set_stack_usageErr error

func tryProcessor_set_stack_usage(pset Processor_set_t, ltotal *uint32, space *Vm_size_t, resident *Vm_size_t, maxusage *Vm_size_t, maxstack *Vm_offset_t) (Kern_return_t, error) {
	if _processor_set_stack_usage == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_stack_usage", "10.0", _processor_set_stack_usageErr)
	}
	return _processor_set_stack_usage(pset, ltotal, space, resident, maxusage, maxstack), nil
}

// Processor_set_stack_usage.
//
// See: https://developer.apple.com/documentation/kernel/1503624-processor_set_stack_usage
func Processor_set_stack_usage(pset Processor_set_t, ltotal *uint32, space *Vm_size_t, resident *Vm_size_t, maxusage *Vm_size_t, maxstack *Vm_offset_t) Kern_return_t {
	result, callErr := tryProcessor_set_stack_usage(pset, ltotal, space, resident, maxusage, maxstack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_statistics func(pset Processor_set_name_t, flavor Processor_set_flavor_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_set_statisticsErr error

func tryProcessor_set_statistics(pset Processor_set_name_t, flavor Processor_set_flavor_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_set_statistics == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_statistics", "10.0", _processor_set_statisticsErr)
	}
	return _processor_set_statistics(pset, flavor, info_out, info_outCnt), nil
}

// Processor_set_statistics.
//
// See: https://developer.apple.com/documentation/kernel/1503608-processor_set_statistics
func Processor_set_statistics(pset Processor_set_name_t, flavor Processor_set_flavor_t, info_out Processor_set_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_set_statistics(pset, flavor, info_out, info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_tasks func(processor_set Processor_set_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_set_tasksErr error

func tryProcessor_set_tasks(processor_set Processor_set_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_set_tasks == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_tasks", "10.0", _processor_set_tasksErr)
	}
	return _processor_set_tasks(processor_set, task_list, task_listCnt), nil
}

// Processor_set_tasks.
//
// See: https://developer.apple.com/documentation/kernel/1503628-processor_set_tasks
func Processor_set_tasks(processor_set Processor_set_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_set_tasks(processor_set, task_list, task_listCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_tasks_with_flavor func(processor_set Processor_set_t, flavor Mach_task_flavor_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_set_tasks_with_flavorErr error

func tryProcessor_set_tasks_with_flavor(processor_set Processor_set_t, flavor Mach_task_flavor_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_set_tasks_with_flavor == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_tasks_with_flavor", "11.0", _processor_set_tasks_with_flavorErr)
	}
	return _processor_set_tasks_with_flavor(processor_set, flavor, task_list, task_listCnt), nil
}

// Processor_set_tasks_with_flavor.
//
// See: https://developer.apple.com/documentation/kernel/3553717-processor_set_tasks_with_flavor
func Processor_set_tasks_with_flavor(processor_set Processor_set_t, flavor Mach_task_flavor_t, task_list *Task_array_t, task_listCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_set_tasks_with_flavor(processor_set, flavor, task_list, task_listCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_set_threads func(processor_set Processor_set_t, thread_list *Thread_act_array_t, thread_listCnt *Mach_msg_type_number_t) Kern_return_t
var _processor_set_threadsErr error

func tryProcessor_set_threads(processor_set Processor_set_t, thread_list *Thread_act_array_t, thread_listCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _processor_set_threads == nil {
		return *new(Kern_return_t), symbolCallError("processor_set_threads", "10.0", _processor_set_threadsErr)
	}
	return _processor_set_threads(processor_set, thread_list, thread_listCnt), nil
}

// Processor_set_threads.
//
// See: https://developer.apple.com/documentation/kernel/1503576-processor_set_threads
func Processor_set_threads(processor_set Processor_set_t, thread_list *Thread_act_array_t, thread_listCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryProcessor_set_threads(processor_set, thread_list, thread_listCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processor_start func(processor Processor_t) Kern_return_t
var _processor_startErr error

func tryProcessor_start(processor Processor_t) (Kern_return_t, error) {
	if _processor_start == nil {
		return *new(Kern_return_t), symbolCallError("processor_start", "10.0", _processor_startErr)
	}
	return _processor_start(processor), nil
}

// Processor_start.
//
// See: https://developer.apple.com/documentation/kernel/1409392-processor_start
func Processor_start(processor Processor_t) Kern_return_t {
	result, callErr := tryProcessor_start(processor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _random func() U_int32_t
var _randomErr error

func tryRandom() (U_int32_t, error) {
	if _random == nil {
		return *new(U_int32_t), symbolCallError("random", "10.0", _randomErr)
	}
	return _random(), nil
}

// Random.
//
// See: https://developer.apple.com/documentation/kernel/1441069-random
func Random() U_int32_t {
	result, callErr := tryRandom()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remainder func(arg0 float64, arg1 float64) float64
var _remainderErr error

func tryRemainder(arg0 float64, arg1 float64) (float64, error) {
	if _remainder == nil {
		return 0.0, symbolCallError("remainder", "10.10", _remainderErr)
	}
	return _remainder(arg0, arg1), nil
}

// Remainder.
//
// See: https://developer.apple.com/documentation/kernel/1557314-remainder
func Remainder(arg0 float64, arg1 float64) float64 {
	result, callErr := tryRemainder(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remainderf func(arg0 float32, arg1 float32) float32
var _remainderfErr error

func tryRemainderf(arg0 float32, arg1 float32) (float32, error) {
	if _remainderf == nil {
		return 0.0, symbolCallError("remainderf", "10.10", _remainderfErr)
	}
	return _remainderf(arg0, arg1), nil
}

// Remainderf.
//
// See: https://developer.apple.com/documentation/kernel/1557334-remainderf
func Remainderf(arg0 float32, arg1 float32) float32 {
	result, callErr := tryRemainderf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remainderl func(arg0 float64, arg1 float64) float64
var _remainderlErr error

func tryRemainderl(arg0 float64, arg1 float64) (float64, error) {
	if _remainderl == nil {
		return 0.0, symbolCallError("remainderl", "10.10", _remainderlErr)
	}
	return _remainderl(arg0, arg1), nil
}

// Remainderl.
//
// See: https://developer.apple.com/documentation/kernel/1557193-remainderl
func Remainderl(arg0 float64, arg1 float64) float64 {
	result, callErr := tryRemainderl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remque func(elt unsafe.Pointer)
var _remqueErr error

func tryRemque(elt unsafe.Pointer) error {
	if _remque == nil {
		return symbolCallError("remque", "10.0", _remqueErr)
	}
	_remque(elt)
	return nil
}

// Remque.
//
// See: https://developer.apple.com/documentation/kernel/1567112-remque
func Remque(elt unsafe.Pointer) {
	if callErr := tryRemque(elt); callErr != nil {
		panic(callErr)
	}
}

var _remquo func(arg0 float64, arg1 float64, arg2 *int32) float64
var _remquoErr error

func tryRemquo(arg0 float64, arg1 float64, arg2 *int32) (float64, error) {
	if _remquo == nil {
		return 0.0, symbolCallError("remquo", "10.10", _remquoErr)
	}
	return _remquo(arg0, arg1, arg2), nil
}

// Remquo.
//
// See: https://developer.apple.com/documentation/kernel/1557171-remquo
func Remquo(arg0 float64, arg1 float64, arg2 *int32) float64 {
	result, callErr := tryRemquo(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remquof func(arg0 float32, arg1 float32, arg2 *int32) float32
var _remquofErr error

func tryRemquof(arg0 float32, arg1 float32, arg2 *int32) (float32, error) {
	if _remquof == nil {
		return 0.0, symbolCallError("remquof", "10.10", _remquofErr)
	}
	return _remquof(arg0, arg1, arg2), nil
}

// Remquof.
//
// See: https://developer.apple.com/documentation/kernel/1557226-remquof
func Remquof(arg0 float32, arg1 float32, arg2 *int32) float32 {
	result, callErr := tryRemquof(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _remquol func(arg0 float64, arg1 float64, arg2 *int32) float64
var _remquolErr error

func tryRemquol(arg0 float64, arg1 float64, arg2 *int32) (float64, error) {
	if _remquol == nil {
		return 0.0, symbolCallError("remquol", "10.10", _remquolErr)
	}
	return _remquol(arg0, arg1, arg2), nil
}

// Remquol.
//
// See: https://developer.apple.com/documentation/kernel/1557307-remquol
func Remquol(arg0 float64, arg1 float64, arg2 *int32) float64 {
	result, callErr := tryRemquol(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _rint func(arg0 float64) float64
var _rintErr error

func tryRint(arg0 float64) (float64, error) {
	if _rint == nil {
		return 0.0, symbolCallError("rint", "10.10", _rintErr)
	}
	return _rint(arg0), nil
}

// Rint.
//
// See: https://developer.apple.com/documentation/kernel/1557284-rint
func Rint(arg0 float64) float64 {
	result, callErr := tryRint(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _rintf func(arg0 float32) float32
var _rintfErr error

func tryRintf(arg0 float32) (float32, error) {
	if _rintf == nil {
		return 0.0, symbolCallError("rintf", "10.10", _rintfErr)
	}
	return _rintf(arg0), nil
}

// Rintf.
//
// See: https://developer.apple.com/documentation/kernel/1557348-rintf
func Rintf(arg0 float32) float32 {
	result, callErr := tryRintf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _rintl func(arg0 float64) float64
var _rintlErr error

func tryRintl(arg0 float64) (float64, error) {
	if _rintl == nil {
		return 0.0, symbolCallError("rintl", "10.10", _rintlErr)
	}
	return _rintl(arg0), nil
}

// Rintl.
//
// See: https://developer.apple.com/documentation/kernel/1557349-rintl
func Rintl(arg0 float64) float64 {
	result, callErr := tryRintl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _round func(arg0 float64) float64
var _roundErr error

func tryRound(arg0 float64) (float64, error) {
	if _round == nil {
		return 0.0, symbolCallError("round", "10.10", _roundErr)
	}
	return _round(arg0), nil
}

// Round.
//
// See: https://developer.apple.com/documentation/kernel/1557369-round
func Round(arg0 float64) float64 {
	result, callErr := tryRound(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _roundf func(arg0 float32) float32
var _roundfErr error

func tryRoundf(arg0 float32) (float32, error) {
	if _roundf == nil {
		return 0.0, symbolCallError("roundf", "10.10", _roundfErr)
	}
	return _roundf(arg0), nil
}

// Roundf.
//
// See: https://developer.apple.com/documentation/kernel/1557316-roundf
func Roundf(arg0 float32) float32 {
	result, callErr := tryRoundf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _roundl func(arg0 float64) float64
var _roundlErr error

func tryRoundl(arg0 float64) (float64, error) {
	if _roundl == nil {
		return 0.0, symbolCallError("roundl", "10.10", _roundlErr)
	}
	return _roundl(arg0), nil
}

// Roundl.
//
// See: https://developer.apple.com/documentation/kernel/1557143-roundl
func Roundl(arg0 float64) float64 {
	result, callErr := tryRoundl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalb func(arg0 float64, arg1 float64) float64
var _scalbErr error

func tryScalb(arg0 float64, arg1 float64) (float64, error) {
	if _scalb == nil {
		return 0.0, symbolCallError("scalb", "10.10", _scalbErr)
	}
	return _scalb(arg0, arg1), nil
}

// Scalb.
//
// See: https://developer.apple.com/documentation/kernel/1557195-scalb
func Scalb(arg0 float64, arg1 float64) float64 {
	result, callErr := tryScalb(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalbln func(arg0 float64, arg1 int) float64
var _scalblnErr error

func tryScalbln(arg0 float64, arg1 int) (float64, error) {
	if _scalbln == nil {
		return 0.0, symbolCallError("scalbln", "10.10", _scalblnErr)
	}
	return _scalbln(arg0, arg1), nil
}

// Scalbln.
//
// See: https://developer.apple.com/documentation/kernel/1557236-scalbln
func Scalbln(arg0 float64, arg1 int) float64 {
	result, callErr := tryScalbln(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalblnf func(arg0 float32, arg1 int) float32
var _scalblnfErr error

func tryScalblnf(arg0 float32, arg1 int) (float32, error) {
	if _scalblnf == nil {
		return 0.0, symbolCallError("scalblnf", "10.10", _scalblnfErr)
	}
	return _scalblnf(arg0, arg1), nil
}

// Scalblnf.
//
// See: https://developer.apple.com/documentation/kernel/1557182-scalblnf
func Scalblnf(arg0 float32, arg1 int) float32 {
	result, callErr := tryScalblnf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalblnl func(arg0 float64, arg1 int) float64
var _scalblnlErr error

func tryScalblnl(arg0 float64, arg1 int) (float64, error) {
	if _scalblnl == nil {
		return 0.0, symbolCallError("scalblnl", "10.10", _scalblnlErr)
	}
	return _scalblnl(arg0, arg1), nil
}

// Scalblnl.
//
// See: https://developer.apple.com/documentation/kernel/1557371-scalblnl
func Scalblnl(arg0 float64, arg1 int) float64 {
	result, callErr := tryScalblnl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalbn func(arg0 float64, arg1 int32) float64
var _scalbnErr error

func tryScalbn(arg0 float64, arg1 int32) (float64, error) {
	if _scalbn == nil {
		return 0.0, symbolCallError("scalbn", "10.10", _scalbnErr)
	}
	return _scalbn(arg0, arg1), nil
}

// Scalbn.
//
// See: https://developer.apple.com/documentation/kernel/1557301-scalbn
func Scalbn(arg0 float64, arg1 int32) float64 {
	result, callErr := tryScalbn(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalbnf func(arg0 float32, arg1 int32) float32
var _scalbnfErr error

func tryScalbnf(arg0 float32, arg1 int32) (float32, error) {
	if _scalbnf == nil {
		return 0.0, symbolCallError("scalbnf", "10.10", _scalbnfErr)
	}
	return _scalbnf(arg0, arg1), nil
}

// Scalbnf.
//
// See: https://developer.apple.com/documentation/kernel/1557209-scalbnf
func Scalbnf(arg0 float32, arg1 int32) float32 {
	result, callErr := tryScalbnf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _scalbnl func(arg0 float64, arg1 int32) float64
var _scalbnlErr error

func tryScalbnl(arg0 float64, arg1 int32) (float64, error) {
	if _scalbnl == nil {
		return 0.0, symbolCallError("scalbnl", "10.10", _scalbnlErr)
	}
	return _scalbnl(arg0, arg1), nil
}

// Scalbnl.
//
// See: https://developer.apple.com/documentation/kernel/1557350-scalbnl
func Scalbnl(arg0 float64, arg1 int32) float64 {
	result, callErr := tryScalbnl(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _selectFunc func(arg0 int32, arg1 *Fd_set, arg2 *Fd_set, arg3 *Fd_set, arg4 *Timeval) int32
var _selectFuncErr error

func trySelectFunc(arg0 int32, arg1 *Fd_set, arg2 *Fd_set, arg3 *Fd_set, arg4 *Timeval) (int32, error) {
	if _selectFunc == nil {
		return 0, symbolCallError("select", "11.0", _selectFuncErr)
	}
	return _selectFunc(arg0, arg1, arg2, arg3, arg4), nil
}

// SelectFunc.
//
// See: https://developer.apple.com/documentation/kernel/3589483-select
func SelectFunc(arg0 int32, arg1 *Fd_set, arg2 *Fd_set, arg3 *Fd_set, arg4 *Timeval) int32 {
	result, callErr := trySelectFunc(arg0, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _semaphore_create func(task Task_t, semaphore *Semaphore_t, policy int32, value int32) Kern_return_t
var _semaphore_createErr error

func trySemaphore_create(task Task_t, semaphore *Semaphore_t, policy int32, value int32) (Kern_return_t, error) {
	if _semaphore_create == nil {
		return *new(Kern_return_t), symbolCallError("semaphore_create", "10.0", _semaphore_createErr)
	}
	return _semaphore_create(task, semaphore, policy, value), nil
}

// Semaphore_create.
//
// See: https://developer.apple.com/documentation/kernel/1538205-semaphore_create
func Semaphore_create(task Task_t, semaphore *Semaphore_t, policy int32, value int32) Kern_return_t {
	result, callErr := trySemaphore_create(task, semaphore, policy, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _semaphore_destroy func(task Task_t, semaphore Semaphore_t) Kern_return_t
var _semaphore_destroyErr error

func trySemaphore_destroy(task Task_t, semaphore Semaphore_t) (Kern_return_t, error) {
	if _semaphore_destroy == nil {
		return *new(Kern_return_t), symbolCallError("semaphore_destroy", "10.0", _semaphore_destroyErr)
	}
	return _semaphore_destroy(task, semaphore), nil
}

// Semaphore_destroy.
//
// See: https://developer.apple.com/documentation/kernel/1537823-semaphore_destroy
func Semaphore_destroy(task Task_t, semaphore Semaphore_t) Kern_return_t {
	result, callErr := trySemaphore_destroy(task, semaphore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _semaphore_signal func(semaphore Semaphore_t) Kern_return_t
var _semaphore_signalErr error

func trySemaphore_signal(semaphore Semaphore_t) (Kern_return_t, error) {
	if _semaphore_signal == nil {
		return *new(Kern_return_t), symbolCallError("semaphore_signal", "10.0", _semaphore_signalErr)
	}
	return _semaphore_signal(semaphore), nil
}

// Semaphore_signal.
//
// See: https://developer.apple.com/documentation/kernel/1585827-semaphore_signal
func Semaphore_signal(semaphore Semaphore_t) Kern_return_t {
	result, callErr := trySemaphore_signal(semaphore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _semaphore_signal_all func(semaphore Semaphore_t) Kern_return_t
var _semaphore_signal_allErr error

func trySemaphore_signal_all(semaphore Semaphore_t) (Kern_return_t, error) {
	if _semaphore_signal_all == nil {
		return *new(Kern_return_t), symbolCallError("semaphore_signal_all", "10.0", _semaphore_signal_allErr)
	}
	return _semaphore_signal_all(semaphore), nil
}

// Semaphore_signal_all.
//
// See: https://developer.apple.com/documentation/kernel/1585830-semaphore_signal_all
func Semaphore_signal_all(semaphore Semaphore_t) Kern_return_t {
	result, callErr := trySemaphore_signal_all(semaphore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _semaphore_wait func(semaphore Semaphore_t) Kern_return_t
var _semaphore_waitErr error

func trySemaphore_wait(semaphore Semaphore_t) (Kern_return_t, error) {
	if _semaphore_wait == nil {
		return *new(Kern_return_t), symbolCallError("semaphore_wait", "10.0", _semaphore_waitErr)
	}
	return _semaphore_wait(semaphore), nil
}

// Semaphore_wait.
//
// See: https://developer.apple.com/documentation/kernel/1585828-semaphore_wait
func Semaphore_wait(semaphore Semaphore_t) Kern_return_t {
	result, callErr := trySemaphore_wait(semaphore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _signal func(arg0 int32, arg1 unsafe.Pointer) func(unsafe.Pointer)
var _signalErr error

func trySignal(arg0 int32, arg1 unsafe.Pointer) (func(unsafe.Pointer), error) {
	if _signal == nil {
		return nil, symbolCallError("signal", "10.0", _signalErr)
	}
	return _signal(arg0, arg1), nil
}

// Signal.
//
// See: https://developer.apple.com/documentation/kernel/1591562-signal
func Signal(arg0 int32, arg1 unsafe.Pointer) func(unsafe.Pointer) {
	result, callErr := trySignal(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sin func(arg0 float64) float64
var _sinErr error

func trySin(arg0 float64) (float64, error) {
	if _sin == nil {
		return 0.0, symbolCallError("sin", "10.10", _sinErr)
	}
	return _sin(arg0), nil
}

// Sin.
//
// See: https://developer.apple.com/documentation/kernel/1557267-sin
func Sin(arg0 float64) float64 {
	result, callErr := trySin(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sinf func(arg0 float32) float32
var _sinfErr error

func trySinf(arg0 float32) (float32, error) {
	if _sinf == nil {
		return 0.0, symbolCallError("sinf", "10.10", _sinfErr)
	}
	return _sinf(arg0), nil
}

// Sinf.
//
// See: https://developer.apple.com/documentation/kernel/1532196-sinf
func Sinf(arg0 float32) float32 {
	result, callErr := trySinf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sinh func(arg0 float64) float64
var _sinhErr error

func trySinh(arg0 float64) (float64, error) {
	if _sinh == nil {
		return 0.0, symbolCallError("sinh", "10.10", _sinhErr)
	}
	return _sinh(arg0), nil
}

// Sinh.
//
// See: https://developer.apple.com/documentation/kernel/1557279-sinh
func Sinh(arg0 float64) float64 {
	result, callErr := trySinh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sinhf func(arg0 float32) float32
var _sinhfErr error

func trySinhf(arg0 float32) (float32, error) {
	if _sinhf == nil {
		return 0.0, symbolCallError("sinhf", "10.10", _sinhfErr)
	}
	return _sinhf(arg0), nil
}

// Sinhf.
//
// See: https://developer.apple.com/documentation/kernel/1557250-sinhf
func Sinhf(arg0 float32) float32 {
	result, callErr := trySinhf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sinhl func(arg0 float64) float64
var _sinhlErr error

func trySinhl(arg0 float64) (float64, error) {
	if _sinhl == nil {
		return 0.0, symbolCallError("sinhl", "10.10", _sinhlErr)
	}
	return _sinhl(arg0), nil
}

// Sinhl.
//
// See: https://developer.apple.com/documentation/kernel/1557240-sinhl
func Sinhl(arg0 float64) float64 {
	result, callErr := trySinhl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sinl func(arg0 float64) float64
var _sinlErr error

func trySinl(arg0 float64) (float64, error) {
	if _sinl == nil {
		return 0.0, symbolCallError("sinl", "10.10", _sinlErr)
	}
	return _sinl(arg0), nil
}

// Sinl.
//
// See: https://developer.apple.com/documentation/kernel/1557325-sinl
func Sinl(arg0 float64) float64 {
	result, callErr := trySinl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _snprintf func(arg0 string, count uintptr, arg2 string) int32
var _snprintfErr error

func trySnprintf(arg0 string, count uintptr, arg2 string) (int32, error) {
	if _snprintf == nil {
		return 0, symbolCallError("snprintf", "10.0", _snprintfErr)
	}
	return _snprintf(arg0, count, arg2), nil
}

// Snprintf.
//
// See: https://developer.apple.com/documentation/kernel/1441052-snprintf
func Snprintf(arg0 string, count uintptr, arg2 string) int32 {
	result, callErr := trySnprintf(arg0, count, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sprintf func(bufp string, arg1 string) int32
var _sprintfErr error

func trySprintf(bufp string, arg1 string) (int32, error) {
	if _sprintf == nil {
		return 0, symbolCallError("sprintf", "10.12", _sprintfErr)
	}
	return _sprintf(bufp, arg1), nil
}

// Sprintf.
//
// Deprecated: Deprecated since macOS 10.13.1.
//
// See: https://developer.apple.com/documentation/kernel/1441083-sprintf
func Sprintf(bufp string, arg1 string) int32 {
	result, callErr := trySprintf(bufp, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sqrt func(arg0 float64) float64
var _sqrtErr error

func trySqrt(arg0 float64) (float64, error) {
	if _sqrt == nil {
		return 0.0, symbolCallError("sqrt", "10.10", _sqrtErr)
	}
	return _sqrt(arg0), nil
}

// Sqrt.
//
// See: https://developer.apple.com/documentation/kernel/1557357-sqrt
func Sqrt(arg0 float64) float64 {
	result, callErr := trySqrt(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sqrtf func(arg0 float32) float32
var _sqrtfErr error

func trySqrtf(arg0 float32) (float32, error) {
	if _sqrtf == nil {
		return 0.0, symbolCallError("sqrtf", "10.9", _sqrtfErr)
	}
	return _sqrtf(arg0), nil
}

// Sqrtf.
//
// See: https://developer.apple.com/documentation/kernel/1532170-sqrtf
func Sqrtf(arg0 float32) float32 {
	result, callErr := trySqrtf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sqrtl func(arg0 float64) float64
var _sqrtlErr error

func trySqrtl(arg0 float64) (float64, error) {
	if _sqrtl == nil {
		return 0.0, symbolCallError("sqrtl", "10.10", _sqrtlErr)
	}
	return _sqrtl(arg0), nil
}

// Sqrtl.
//
// See: https://developer.apple.com/documentation/kernel/1557199-sqrtl
func Sqrtl(arg0 float64) float64 {
	result, callErr := trySqrtl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sscanf func(arg0 string, arg1 string) int32
var _sscanfErr error

func trySscanf(arg0 string, arg1 string) (int32, error) {
	if _sscanf == nil {
		return 0, symbolCallError("sscanf", "10.0", _sscanfErr)
	}
	return _sscanf(arg0, arg1), nil
}

// Sscanf.
//
// See: https://developer.apple.com/documentation/kernel/1441040-sscanf
func Sscanf(arg0 string, arg1 string) int32 {
	result, callErr := trySscanf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strcasecmp func(s1 string, s2 string) int32
var _strcasecmpErr error

func tryStrcasecmp(s1 string, s2 string) (int32, error) {
	if _strcasecmp == nil {
		return 0, symbolCallError("strcasecmp", "10.4", _strcasecmpErr)
	}
	return _strcasecmp(s1, s2), nil
}

// Strcasecmp.
//
// See: https://developer.apple.com/documentation/kernel/1579333-strcasecmp
func Strcasecmp(s1 string, s2 string) int32 {
	result, callErr := tryStrcasecmp(s1, s2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strcat func(dst *byte, src string) *byte
var _strcatErr error

func tryStrcat(dst *byte, src string) (*byte, error) {
	if _strcat == nil {
		return nil, symbolCallError("strcat", "10.0", _strcatErr)
	}
	return _strcat(dst, src), nil
}

// Strcat.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/kernel/1579328-strcat
func Strcat(dst *byte, src string) *byte {
	result, callErr := tryStrcat(dst, src)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strchr func(s string, c int32) *byte
var _strchrErr error

func tryStrchr(s string, c int32) (*byte, error) {
	if _strchr == nil {
		return nil, symbolCallError("strchr", "10.0", _strchrErr)
	}
	return _strchr(s, c), nil
}

// Strchr.
//
// See: https://developer.apple.com/documentation/kernel/1579345-strchr
func Strchr(s string, c int32) *byte {
	result, callErr := tryStrchr(s, c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strcmp func(s1 string, s2 string) int32
var _strcmpErr error

func tryStrcmp(s1 string, s2 string) (int32, error) {
	if _strcmp == nil {
		return 0, symbolCallError("strcmp", "10.0", _strcmpErr)
	}
	return _strcmp(s1, s2), nil
}

// Strcmp.
//
// See: https://developer.apple.com/documentation/kernel/1579329-strcmp
func Strcmp(s1 string, s2 string) int32 {
	result, callErr := tryStrcmp(s1, s2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strcpy func(arg0 string, arg1 string) *byte
var _strcpyErr error

func tryStrcpy(arg0 string, arg1 string) (*byte, error) {
	if _strcpy == nil {
		return nil, symbolCallError("strcpy", "10.0", _strcpyErr)
	}
	return _strcpy(arg0, arg1), nil
}

// Strcpy.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/kernel/1579337-strcpy
func Strcpy(arg0 string, arg1 string) *byte {
	result, callErr := tryStrcpy(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strlcat func(dst *byte, src string, n uintptr) uintptr
var _strlcatErr error

func tryStrlcat(dst *byte, src string, n uintptr) (uintptr, error) {
	if _strlcat == nil {
		return 0, symbolCallError("strlcat", "10.5", _strlcatErr)
	}
	return _strlcat(dst, src, n), nil
}

// Strlcat.
//
// See: https://developer.apple.com/documentation/kernel/1579344-strlcat
func Strlcat(dst *byte, src string, n uintptr) uintptr {
	result, callErr := tryStrlcat(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strlcpy func(dst *byte, src string, n uintptr) uintptr
var _strlcpyErr error

func tryStrlcpy(dst *byte, src string, n uintptr) (uintptr, error) {
	if _strlcpy == nil {
		return 0, symbolCallError("strlcpy", "10.5", _strlcpyErr)
	}
	return _strlcpy(dst, src, n), nil
}

// Strlcpy.
//
// See: https://developer.apple.com/documentation/kernel/1579349-strlcpy
func Strlcpy(dst *byte, src string, n uintptr) uintptr {
	result, callErr := tryStrlcpy(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strlen func(s string) uintptr
var _strlenErr error

func tryStrlen(s string) (uintptr, error) {
	if _strlen == nil {
		return 0, symbolCallError("strlen", "10.0", _strlenErr)
	}
	return _strlen(s), nil
}

// Strlen.
//
// See: https://developer.apple.com/documentation/kernel/1579342-strlen
func Strlen(s string) uintptr {
	result, callErr := tryStrlen(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strncasecmp func(s1 string, s2 string, n uintptr) int32
var _strncasecmpErr error

func tryStrncasecmp(s1 string, s2 string, n uintptr) (int32, error) {
	if _strncasecmp == nil {
		return 0, symbolCallError("strncasecmp", "10.4", _strncasecmpErr)
	}
	return _strncasecmp(s1, s2, n), nil
}

// Strncasecmp.
//
// See: https://developer.apple.com/documentation/kernel/1579347-strncasecmp
func Strncasecmp(s1 string, s2 string, n uintptr) int32 {
	result, callErr := tryStrncasecmp(s1, s2, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strncat func(dst *byte, src string, n uintptr) *byte
var _strncatErr error

func tryStrncat(dst *byte, src string, n uintptr) (*byte, error) {
	if _strncat == nil {
		return nil, symbolCallError("strncat", "10.0", _strncatErr)
	}
	return _strncat(dst, src, n), nil
}

// Strncat.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/kernel/1579343-strncat
func Strncat(dst *byte, src string, n uintptr) *byte {
	result, callErr := tryStrncat(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strncmp func(s1 string, s2 string, n uintptr) int32
var _strncmpErr error

func tryStrncmp(s1 string, s2 string, n uintptr) (int32, error) {
	if _strncmp == nil {
		return 0, symbolCallError("strncmp", "10.0", _strncmpErr)
	}
	return _strncmp(s1, s2, n), nil
}

// Strncmp.
//
// See: https://developer.apple.com/documentation/kernel/1579335-strncmp
func Strncmp(s1 string, s2 string, n uintptr) int32 {
	result, callErr := tryStrncmp(s1, s2, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strncpy func(dst *byte, src string, n uintptr) *byte
var _strncpyErr error

func tryStrncpy(dst *byte, src string, n uintptr) (*byte, error) {
	if _strncpy == nil {
		return nil, symbolCallError("strncpy", "10.0", _strncpyErr)
	}
	return _strncpy(dst, src, n), nil
}

// Strncpy.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/kernel/1579331-strncpy
func Strncpy(dst *byte, src string, n uintptr) *byte {
	result, callErr := tryStrncpy(dst, src, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strnlen func(s string, n uintptr) uintptr
var _strnlenErr error

func tryStrnlen(s string, n uintptr) (uintptr, error) {
	if _strnlen == nil {
		return 0, symbolCallError("strnlen", "10.5", _strnlenErr)
	}
	return _strnlen(s, n), nil
}

// Strnlen.
//
// See: https://developer.apple.com/documentation/kernel/1579340-strnlen
func Strnlen(s string, n uintptr) uintptr {
	result, callErr := tryStrnlen(s, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strnstr func(s string, find string, slen uintptr) *byte
var _strnstrErr error

func tryStrnstr(s string, find string, slen uintptr) (*byte, error) {
	if _strnstr == nil {
		return nil, symbolCallError("strnstr", "10.9", _strnstrErr)
	}
	return _strnstr(s, find, slen), nil
}

// Strnstr.
//
// See: https://developer.apple.com/documentation/kernel/1579346-strnstr
func Strnstr(s string, find string, slen uintptr) *byte {
	result, callErr := tryStrnstr(s, find, slen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strsep func(arg0 string, arg1 string) *byte
var _strsepErr error

func tryStrsep(arg0 string, arg1 string) (*byte, error) {
	if _strsep == nil {
		return nil, symbolCallError("strsep", "10.5", _strsepErr)
	}
	return _strsep(arg0, arg1), nil
}

// Strsep.
//
// See: https://developer.apple.com/documentation/kernel/1441093-strsep
func Strsep(arg0 string, arg1 string) *byte {
	result, callErr := tryStrsep(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strtol func(arg0 string, arg1 string, arg2 int32) int
var _strtolErr error

func tryStrtol(arg0 string, arg1 string, arg2 int32) (int, error) {
	if _strtol == nil {
		return 0, symbolCallError("strtol", "10.0", _strtolErr)
	}
	return _strtol(arg0, arg1, arg2), nil
}

// Strtol.
//
// See: https://developer.apple.com/documentation/kernel/1441045-strtol
func Strtol(arg0 string, arg1 string, arg2 int32) int {
	result, callErr := tryStrtol(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strtoq func(arg0 string, arg1 string, arg2 int32) Quad_t
var _strtoqErr error

func tryStrtoq(arg0 string, arg1 string, arg2 int32) (Quad_t, error) {
	if _strtoq == nil {
		return *new(Quad_t), symbolCallError("strtoq", "10.1", _strtoqErr)
	}
	return _strtoq(arg0, arg1, arg2), nil
}

// Strtoq.
//
// See: https://developer.apple.com/documentation/kernel/1441090-strtoq
func Strtoq(arg0 string, arg1 string, arg2 int32) Quad_t {
	result, callErr := tryStrtoq(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strtoul func(arg0 string, arg1 string, arg2 int32) U_long
var _strtoulErr error

func tryStrtoul(arg0 string, arg1 string, arg2 int32) (U_long, error) {
	if _strtoul == nil {
		return *new(U_long), symbolCallError("strtoul", "10.0", _strtoulErr)
	}
	return _strtoul(arg0, arg1, arg2), nil
}

// Strtoul.
//
// See: https://developer.apple.com/documentation/kernel/1441100-strtoul
func Strtoul(arg0 string, arg1 string, arg2 int32) U_long {
	result, callErr := tryStrtoul(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _strtouq func(arg0 string, arg1 string, arg2 int32) U_quad_t
var _strtouqErr error

func tryStrtouq(arg0 string, arg1 string, arg2 int32) (U_quad_t, error) {
	if _strtouq == nil {
		return *new(U_quad_t), symbolCallError("strtouq", "10.1", _strtouqErr)
	}
	return _strtouq(arg0, arg1, arg2), nil
}

// Strtouq.
//
// See: https://developer.apple.com/documentation/kernel/1441084-strtouq
func Strtouq(arg0 string, arg1 string, arg2 int32) U_quad_t {
	result, callErr := tryStrtouq(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sysctlbyname func(arg0 string, arg1 unsafe.Pointer, arg2 *uintptr, arg3 unsafe.Pointer, arg4 uintptr) int32
var _sysctlbynameErr error

func trySysctlbyname(arg0 string, arg1 unsafe.Pointer, arg2 *uintptr, arg3 unsafe.Pointer, arg4 uintptr) (int32, error) {
	if _sysctlbyname == nil {
		return 0, symbolCallError("sysctlbyname", "10.0", _sysctlbynameErr)
	}
	return _sysctlbyname(arg0, arg1, arg2, arg3, arg4), nil
}

// Sysctlbyname gets or sets information about the system and environment.
//
// See: https://developer.apple.com/documentation/kernel/1387446-sysctlbyname
func Sysctlbyname(arg0 string, arg1 unsafe.Pointer, arg2 *uintptr, arg3 unsafe.Pointer, arg4 uintptr) int32 {
	result, callErr := trySysctlbyname(arg0, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tan func(arg0 float64) float64
var _tanErr error

func tryTan(arg0 float64) (float64, error) {
	if _tan == nil {
		return 0.0, symbolCallError("tan", "10.10", _tanErr)
	}
	return _tan(arg0), nil
}

// Tan.
//
// See: https://developer.apple.com/documentation/kernel/1557150-tan
func Tan(arg0 float64) float64 {
	result, callErr := tryTan(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tanf func(arg0 float32) float32
var _tanfErr error

func tryTanf(arg0 float32) (float32, error) {
	if _tanf == nil {
		return 0.0, symbolCallError("tanf", "10.10", _tanfErr)
	}
	return _tanf(arg0), nil
}

// Tanf.
//
// See: https://developer.apple.com/documentation/kernel/1557258-tanf
func Tanf(arg0 float32) float32 {
	result, callErr := tryTanf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tanh func(arg0 float64) float64
var _tanhErr error

func tryTanh(arg0 float64) (float64, error) {
	if _tanh == nil {
		return 0.0, symbolCallError("tanh", "10.10", _tanhErr)
	}
	return _tanh(arg0), nil
}

// Tanh.
//
// See: https://developer.apple.com/documentation/kernel/1557370-tanh
func Tanh(arg0 float64) float64 {
	result, callErr := tryTanh(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tanhf func(arg0 float32) float32
var _tanhfErr error

func tryTanhf(arg0 float32) (float32, error) {
	if _tanhf == nil {
		return 0.0, symbolCallError("tanhf", "10.10", _tanhfErr)
	}
	return _tanhf(arg0), nil
}

// Tanhf.
//
// See: https://developer.apple.com/documentation/kernel/1557286-tanhf
func Tanhf(arg0 float32) float32 {
	result, callErr := tryTanhf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tanhl func(arg0 float64) float64
var _tanhlErr error

func tryTanhl(arg0 float64) (float64, error) {
	if _tanhl == nil {
		return 0.0, symbolCallError("tanhl", "10.10", _tanhlErr)
	}
	return _tanhl(arg0), nil
}

// Tanhl.
//
// See: https://developer.apple.com/documentation/kernel/1557188-tanhl
func Tanhl(arg0 float64) float64 {
	result, callErr := tryTanhl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tanl func(arg0 float64) float64
var _tanlErr error

func tryTanl(arg0 float64) (float64, error) {
	if _tanl == nil {
		return 0.0, symbolCallError("tanl", "10.10", _tanlErr)
	}
	return _tanl(arg0), nil
}

// Tanl.
//
// See: https://developer.apple.com/documentation/kernel/1557239-tanl
func Tanl(arg0 float64) float64 {
	result, callErr := tryTanl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_assign func(task Task_t, new_set Processor_set_t, assign_threads Boolean_t) Kern_return_t
var _task_assignErr error

func tryTask_assign(task Task_t, new_set Processor_set_t, assign_threads Boolean_t) (Kern_return_t, error) {
	if _task_assign == nil {
		return *new(Kern_return_t), symbolCallError("task_assign", "10.0", _task_assignErr)
	}
	return _task_assign(task, new_set, assign_threads), nil
}

// Task_assign.
//
// See: https://developer.apple.com/documentation/kernel/1537803-task_assign
func Task_assign(task Task_t, new_set Processor_set_t, assign_threads Boolean_t) Kern_return_t {
	result, callErr := tryTask_assign(task, new_set, assign_threads)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_assign_default func(task Task_t, assign_threads Boolean_t) Kern_return_t
var _task_assign_defaultErr error

func tryTask_assign_default(task Task_t, assign_threads Boolean_t) (Kern_return_t, error) {
	if _task_assign_default == nil {
		return *new(Kern_return_t), symbolCallError("task_assign_default", "10.0", _task_assign_defaultErr)
	}
	return _task_assign_default(task, assign_threads), nil
}

// Task_assign_default.
//
// See: https://developer.apple.com/documentation/kernel/1537916-task_assign_default
func Task_assign_default(task Task_t, assign_threads Boolean_t) Kern_return_t {
	result, callErr := tryTask_assign_default(task, assign_threads)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_create func(target_task Task_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) Kern_return_t
var _task_createErr error

func tryTask_create(target_task Task_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) (Kern_return_t, error) {
	if _task_create == nil {
		return *new(Kern_return_t), symbolCallError("task_create", "10.0", _task_createErr)
	}
	return _task_create(target_task, ledgers, ledgersCnt, inherit_memory, child_task), nil
}

// Task_create.
//
// See: https://developer.apple.com/documentation/kernel/1538088-task_create
func Task_create(target_task Task_t, ledgers Ledger_array_t, ledgersCnt Mach_msg_type_number_t, inherit_memory Boolean_t, child_task *Task_t) Kern_return_t {
	result, callErr := tryTask_create(target_task, ledgers, ledgersCnt, inherit_memory, child_task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_create_identity_token func(task Task_t, token *Task_id_token_t) Kern_return_t
var _task_create_identity_tokenErr error

func tryTask_create_identity_token(task Task_t, token *Task_id_token_t) (Kern_return_t, error) {
	if _task_create_identity_token == nil {
		return *new(Kern_return_t), symbolCallError("task_create_identity_token", "11.3", _task_create_identity_tokenErr)
	}
	return _task_create_identity_token(task, token), nil
}

// Task_create_identity_token.
//
// See: https://developer.apple.com/documentation/kernel/3727994-task_create_identity_token
func Task_create_identity_token(task Task_t, token *Task_id_token_t) Kern_return_t {
	result, callErr := tryTask_create_identity_token(task, token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_dyld_process_info_notify_deregister func(target_task Task_read_t, notify Mach_port_name_t) Kern_return_t
var _task_dyld_process_info_notify_deregisterErr error

func tryTask_dyld_process_info_notify_deregister(target_task Task_read_t, notify Mach_port_name_t) (Kern_return_t, error) {
	if _task_dyld_process_info_notify_deregister == nil {
		return *new(Kern_return_t), symbolCallError("task_dyld_process_info_notify_deregister", "11.3", _task_dyld_process_info_notify_deregisterErr)
	}
	return _task_dyld_process_info_notify_deregister(target_task, notify), nil
}

// Task_dyld_process_info_notify_deregister.
//
// See: https://developer.apple.com/documentation/kernel/3727995-task_dyld_process_info_notify_de
func Task_dyld_process_info_notify_deregister(target_task Task_read_t, notify Mach_port_name_t) Kern_return_t {
	result, callErr := tryTask_dyld_process_info_notify_deregister(target_task, notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_dyld_process_info_notify_register func(target_task Task_read_t, notify uint32) Kern_return_t
var _task_dyld_process_info_notify_registerErr error

func tryTask_dyld_process_info_notify_register(target_task Task_read_t, notify uint32) (Kern_return_t, error) {
	if _task_dyld_process_info_notify_register == nil {
		return *new(Kern_return_t), symbolCallError("task_dyld_process_info_notify_register", "11.3", _task_dyld_process_info_notify_registerErr)
	}
	return _task_dyld_process_info_notify_register(target_task, notify), nil
}

// Task_dyld_process_info_notify_register.
//
// See: https://developer.apple.com/documentation/kernel/3727996-task_dyld_process_info_notify_re
func Task_dyld_process_info_notify_register(target_task Task_read_t, notify uint32) Kern_return_t {
	result, callErr := tryTask_dyld_process_info_notify_register(target_task, notify)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_generate_corpse func(task Task_read_t, corpse_task_port *uint32) Kern_return_t
var _task_generate_corpseErr error

func tryTask_generate_corpse(task Task_read_t, corpse_task_port *uint32) (Kern_return_t, error) {
	if _task_generate_corpse == nil {
		return *new(Kern_return_t), symbolCallError("task_generate_corpse", "10.12", _task_generate_corpseErr)
	}
	return _task_generate_corpse(task, corpse_task_port), nil
}

// Task_generate_corpse.
//
// See: https://developer.apple.com/documentation/kernel/1646547-task_generate_corpse
func Task_generate_corpse(task Task_read_t, corpse_task_port *uint32) Kern_return_t {
	result, callErr := tryTask_generate_corpse(task, corpse_task_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_assignment func(task Task_inspect_t, assigned_set *Processor_set_name_t) Kern_return_t
var _task_get_assignmentErr error

func tryTask_get_assignment(task Task_inspect_t, assigned_set *Processor_set_name_t) (Kern_return_t, error) {
	if _task_get_assignment == nil {
		return *new(Kern_return_t), symbolCallError("task_get_assignment", "10.0", _task_get_assignmentErr)
	}
	return _task_get_assignment(task, assigned_set), nil
}

// Task_get_assignment.
//
// See: https://developer.apple.com/documentation/kernel/1537882-task_get_assignment
func Task_get_assignment(task Task_inspect_t, assigned_set *Processor_set_name_t) Kern_return_t {
	result, callErr := tryTask_get_assignment(task, assigned_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_dyld_image_infos func(task Task_read_t, dyld_images *Dyld_kernel_image_info_array_t, dyld_imagesCnt *Mach_msg_type_number_t) Kern_return_t
var _task_get_dyld_image_infosErr error

func tryTask_get_dyld_image_infos(task Task_read_t, dyld_images *Dyld_kernel_image_info_array_t, dyld_imagesCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_get_dyld_image_infos == nil {
		return *new(Kern_return_t), symbolCallError("task_get_dyld_image_infos", "10.12", _task_get_dyld_image_infosErr)
	}
	return _task_get_dyld_image_infos(task, dyld_images, dyld_imagesCnt), nil
}

// Task_get_dyld_image_infos.
//
// See: https://developer.apple.com/documentation/kernel/1646553-task_get_dyld_image_infos
func Task_get_dyld_image_infos(task Task_read_t, dyld_images *Dyld_kernel_image_info_array_t, dyld_imagesCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_get_dyld_image_infos(task, dyld_images, dyld_imagesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_emulation_vector func(task Task_t, vector_start *int32, emulation_vector *Emulation_vector_t, emulation_vectorCnt *Mach_msg_type_number_t) Kern_return_t
var _task_get_emulation_vectorErr error

func tryTask_get_emulation_vector(task Task_t, vector_start *int32, emulation_vector *Emulation_vector_t, emulation_vectorCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_get_emulation_vector == nil {
		return *new(Kern_return_t), symbolCallError("task_get_emulation_vector", "10.0", _task_get_emulation_vectorErr)
	}
	return _task_get_emulation_vector(task, vector_start, emulation_vector, emulation_vectorCnt), nil
}

// Task_get_emulation_vector.
//
// See: https://developer.apple.com/documentation/kernel/1537831-task_get_emulation_vector
func Task_get_emulation_vector(task Task_t, vector_start *int32, emulation_vector *Emulation_vector_t, emulation_vectorCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_get_emulation_vector(task, vector_start, emulation_vector, emulation_vectorCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_exc_guard_behavior func(task Task_inspect_t, behavior *Task_exc_guard_behavior_t) Kern_return_t
var _task_get_exc_guard_behaviorErr error

func tryTask_get_exc_guard_behavior(task Task_inspect_t, behavior *Task_exc_guard_behavior_t) (Kern_return_t, error) {
	if _task_get_exc_guard_behavior == nil {
		return *new(Kern_return_t), symbolCallError("task_get_exc_guard_behavior", "10.15", _task_get_exc_guard_behaviorErr)
	}
	return _task_get_exc_guard_behavior(task, behavior), nil
}

// Task_get_exc_guard_behavior.
//
// See: https://developer.apple.com/documentation/kernel/3197704-task_get_exc_guard_behavior
func Task_get_exc_guard_behavior(task Task_inspect_t, behavior *Task_exc_guard_behavior_t) Kern_return_t {
	result, callErr := tryTask_get_exc_guard_behavior(task, behavior)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_exception_ports func(task Task_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _task_get_exception_portsErr error

func tryTask_get_exception_ports(task Task_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _task_get_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("task_get_exception_ports", "10.0", _task_get_exception_portsErr)
	}
	return _task_get_exception_ports(task, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors), nil
}

// Task_get_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1537860-task_get_exception_ports
func Task_get_exception_ports(task Task_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryTask_get_exception_ports(task, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_exception_ports_info func(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _task_get_exception_ports_infoErr error

func tryTask_get_exception_ports_info(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _task_get_exception_ports_info == nil {
		return *new(Kern_return_t), symbolCallError("task_get_exception_ports_info", "11.3", _task_get_exception_ports_infoErr)
	}
	return _task_get_exception_ports_info(port, exception_mask, masks, masksCnt, old_handlers_info, old_behaviors, old_flavors), nil
}

// Task_get_exception_ports_info.
//
// See: https://developer.apple.com/documentation/kernel/3727997-task_get_exception_ports_info
func Task_get_exception_ports_info(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryTask_get_exception_ports_info(port, exception_mask, masks, masksCnt, old_handlers_info, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_mach_voucher func(task Task_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) Kern_return_t
var _task_get_mach_voucherErr error

func tryTask_get_mach_voucher(task Task_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) (Kern_return_t, error) {
	if _task_get_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("task_get_mach_voucher", "10.10", _task_get_mach_voucherErr)
	}
	return _task_get_mach_voucher(task, which, voucher), nil
}

// Task_get_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1537867-task_get_mach_voucher
func Task_get_mach_voucher(task Task_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) Kern_return_t {
	result, callErr := tryTask_get_mach_voucher(task, which, voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_special_port func(task Task_inspect_t, which_port int32, special_port *uint32) Kern_return_t
var _task_get_special_portErr error

func tryTask_get_special_port(task Task_inspect_t, which_port int32, special_port *uint32) (Kern_return_t, error) {
	if _task_get_special_port == nil {
		return *new(Kern_return_t), symbolCallError("task_get_special_port", "10.0", _task_get_special_portErr)
	}
	return _task_get_special_port(task, which_port, special_port), nil
}

// Task_get_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1537682-task_get_special_port
func Task_get_special_port(task Task_inspect_t, which_port int32, special_port *uint32) Kern_return_t {
	result, callErr := tryTask_get_special_port(task, which_port, special_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_get_state func(task Task_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t
var _task_get_stateErr error

func tryTask_get_state(task Task_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_get_state == nil {
		return *new(Kern_return_t), symbolCallError("task_get_state", "10.6", _task_get_stateErr)
	}
	return _task_get_state(task, flavor, old_state, old_stateCnt), nil
}

// Task_get_state.
//
// See: https://developer.apple.com/documentation/kernel/1537707-task_get_state
func Task_get_state(task Task_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_get_state(task, flavor, old_state, old_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_identity_token_get_task_port func(token Task_id_token_t, flavor Task_flavor_t, task_port *uint32) Kern_return_t
var _task_identity_token_get_task_portErr error

func tryTask_identity_token_get_task_port(token Task_id_token_t, flavor Task_flavor_t, task_port *uint32) (Kern_return_t, error) {
	if _task_identity_token_get_task_port == nil {
		return *new(Kern_return_t), symbolCallError("task_identity_token_get_task_port", "11.3", _task_identity_token_get_task_portErr)
	}
	return _task_identity_token_get_task_port(token, flavor, task_port), nil
}

// Task_identity_token_get_task_port.
//
// See: https://developer.apple.com/documentation/kernel/3727998-task_identity_token_get_task_por
func Task_identity_token_get_task_port(token Task_id_token_t, flavor Task_flavor_t, task_port *uint32) Kern_return_t {
	result, callErr := tryTask_identity_token_get_task_port(token, flavor, task_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_info func(target_task Task_name_t, flavor Task_flavor_t, task_info_out Task_info_t, task_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _task_infoErr error

func tryTask_info(target_task Task_name_t, flavor Task_flavor_t, task_info_out Task_info_t, task_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_info == nil {
		return *new(Kern_return_t), symbolCallError("task_info", "10.0", _task_infoErr)
	}
	return _task_info(target_task, flavor, task_info_out, task_info_outCnt), nil
}

// Task_info.
//
// See: https://developer.apple.com/documentation/kernel/1537934-task_info
func Task_info(target_task Task_name_t, flavor Task_flavor_t, task_info_out Task_info_t, task_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_info(target_task, flavor, task_info_out, task_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_inspect func(task Task_inspect_t, flavor Task_inspect_flavor_t, info_out Task_inspect_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _task_inspectErr error

func tryTask_inspect(task Task_inspect_t, flavor Task_inspect_flavor_t, info_out Task_inspect_info_t, info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_inspect == nil {
		return *new(Kern_return_t), symbolCallError("task_inspect", "10.13", _task_inspectErr)
	}
	return _task_inspect(task, flavor, info_out, info_outCnt), nil
}

// Task_inspect.
//
// See: https://developer.apple.com/documentation/kernel/2876476-task_inspect
func Task_inspect(task Task_inspect_t, flavor Task_inspect_flavor_t, info_out Task_inspect_info_t, info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_inspect(task, flavor, info_out, info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_map_corpse_info func(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Vm_address_t, kcd_size *uint32) Kern_return_t
var _task_map_corpse_infoErr error

func tryTask_map_corpse_info(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Vm_address_t, kcd_size *uint32) (Kern_return_t, error) {
	if _task_map_corpse_info == nil {
		return *new(Kern_return_t), symbolCallError("task_map_corpse_info", "10.12", _task_map_corpse_infoErr)
	}
	return _task_map_corpse_info(task, corspe_task, kcd_addr_begin, kcd_size), nil
}

// Task_map_corpse_info.
//
// See: https://developer.apple.com/documentation/kernel/1646533-task_map_corpse_info
func Task_map_corpse_info(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Vm_address_t, kcd_size *uint32) Kern_return_t {
	result, callErr := tryTask_map_corpse_info(task, corspe_task, kcd_addr_begin, kcd_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_map_corpse_info_64 func(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) Kern_return_t
var _task_map_corpse_info_64Err error

func tryTask_map_corpse_info_64(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) (Kern_return_t, error) {
	if _task_map_corpse_info_64 == nil {
		return *new(Kern_return_t), symbolCallError("task_map_corpse_info_64", "10.12", _task_map_corpse_info_64Err)
	}
	return _task_map_corpse_info_64(task, corspe_task, kcd_addr_begin, kcd_size), nil
}

// Task_map_corpse_info_64.
//
// See: https://developer.apple.com/documentation/kernel/1911639-task_map_corpse_info_64
func Task_map_corpse_info_64(task Task_t, corspe_task Task_read_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) Kern_return_t {
	result, callErr := tryTask_map_corpse_info_64(task, corspe_task, kcd_addr_begin, kcd_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_map_kcdata_object_64 func(task Task_t, kcdata_object Kcdata_object_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) Kern_return_t
var _task_map_kcdata_object_64Err error

func tryTask_map_kcdata_object_64(task Task_t, kcdata_object Kcdata_object_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) (Kern_return_t, error) {
	if _task_map_kcdata_object_64 == nil {
		return *new(Kern_return_t), symbolCallError("task_map_kcdata_object_64", "13.0", _task_map_kcdata_object_64Err)
	}
	return _task_map_kcdata_object_64(task, kcdata_object, kcd_addr_begin, kcd_size), nil
}

// Task_map_kcdata_object_64.
//
// See: https://developer.apple.com/documentation/kernel/3943647-task_map_kcdata_object_64
func Task_map_kcdata_object_64(task Task_t, kcdata_object Kcdata_object_t, kcd_addr_begin *Mach_vm_address_t, kcd_size *Mach_vm_size_t) Kern_return_t {
	result, callErr := tryTask_map_kcdata_object_64(task, kcdata_object, kcd_addr_begin, kcd_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_policy func(task Task_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t, change Boolean_t) Kern_return_t
var _task_policyErr error

func tryTask_policy(task Task_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t, change Boolean_t) (Kern_return_t, error) {
	if _task_policy == nil {
		return *new(Kern_return_t), symbolCallError("task_policy", "10.0", _task_policyErr)
	}
	return _task_policy(task, policy, base, baseCnt, set_limit, change), nil
}

// Task_policy.
//
// See: https://developer.apple.com/documentation/kernel/1537990-task_policy
func Task_policy(task Task_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t, change Boolean_t) Kern_return_t {
	result, callErr := tryTask_policy(task, policy, base, baseCnt, set_limit, change)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_policy_get func(task Task_policy_get_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) Kern_return_t
var _task_policy_getErr error

func tryTask_policy_get(task Task_policy_get_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) (Kern_return_t, error) {
	if _task_policy_get == nil {
		return *new(Kern_return_t), symbolCallError("task_policy_get", "10.0", _task_policy_getErr)
	}
	return _task_policy_get(task, flavor, policy_info, policy_infoCnt, get_default), nil
}

// Task_policy_get.
//
// See: https://developer.apple.com/documentation/kernel/1537778-task_policy_get
func Task_policy_get(task Task_policy_get_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) Kern_return_t {
	result, callErr := tryTask_policy_get(task, flavor, policy_info, policy_infoCnt, get_default)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_policy_set func(task Task_policy_set_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt Mach_msg_type_number_t) Kern_return_t
var _task_policy_setErr error

func tryTask_policy_set(task Task_policy_set_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_policy_set == nil {
		return *new(Kern_return_t), symbolCallError("task_policy_set", "10.0", _task_policy_setErr)
	}
	return _task_policy_set(task, flavor, policy_info, policy_infoCnt), nil
}

// Task_policy_set.
//
// See: https://developer.apple.com/documentation/kernel/1537790-task_policy_set
func Task_policy_set(task Task_policy_set_t, flavor Task_policy_flavor_t, policy_info Task_policy_t, policy_infoCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_policy_set(task, flavor, policy_info, policy_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_purgable_info func(task Task_inspect_t, stats *Task_purgable_info_t) Kern_return_t
var _task_purgable_infoErr error

func tryTask_purgable_info(task Task_inspect_t, stats *Task_purgable_info_t) (Kern_return_t, error) {
	if _task_purgable_info == nil {
		return *new(Kern_return_t), symbolCallError("task_purgable_info", "10.9", _task_purgable_infoErr)
	}
	return _task_purgable_info(task, stats), nil
}

// Task_purgable_info.
//
// See: https://developer.apple.com/documentation/kernel/1538155-task_purgable_info
func Task_purgable_info(task Task_inspect_t, stats *Task_purgable_info_t) Kern_return_t {
	result, callErr := tryTask_purgable_info(task, stats)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_register_dyld_get_process_state func(task Task_t, dyld_process_state *Dyld_kernel_process_info_t) Kern_return_t
var _task_register_dyld_get_process_stateErr error

func tryTask_register_dyld_get_process_state(task Task_t, dyld_process_state *Dyld_kernel_process_info_t) (Kern_return_t, error) {
	if _task_register_dyld_get_process_state == nil {
		return *new(Kern_return_t), symbolCallError("task_register_dyld_get_process_state", "10.12", _task_register_dyld_get_process_stateErr)
	}
	return _task_register_dyld_get_process_state(task, dyld_process_state), nil
}

// Task_register_dyld_get_process_state.
//
// See: https://developer.apple.com/documentation/kernel/1646549-task_register_dyld_get_process_s
func Task_register_dyld_get_process_state(task Task_t, dyld_process_state *Dyld_kernel_process_info_t) Kern_return_t {
	result, callErr := tryTask_register_dyld_get_process_state(task, dyld_process_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_register_dyld_image_infos func(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) Kern_return_t
var _task_register_dyld_image_infosErr error

func tryTask_register_dyld_image_infos(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_register_dyld_image_infos == nil {
		return *new(Kern_return_t), symbolCallError("task_register_dyld_image_infos", "10.12", _task_register_dyld_image_infosErr)
	}
	return _task_register_dyld_image_infos(task, dyld_images, dyld_imagesCnt), nil
}

// Task_register_dyld_image_infos.
//
// See: https://developer.apple.com/documentation/kernel/1646550-task_register_dyld_image_infos
func Task_register_dyld_image_infos(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_register_dyld_image_infos(task, dyld_images, dyld_imagesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_register_dyld_set_dyld_state func(task Task_t, dyld_state uint8) Kern_return_t
var _task_register_dyld_set_dyld_stateErr error

func tryTask_register_dyld_set_dyld_state(task Task_t, dyld_state uint8) (Kern_return_t, error) {
	if _task_register_dyld_set_dyld_state == nil {
		return *new(Kern_return_t), symbolCallError("task_register_dyld_set_dyld_state", "10.12", _task_register_dyld_set_dyld_stateErr)
	}
	return _task_register_dyld_set_dyld_state(task, dyld_state), nil
}

// Task_register_dyld_set_dyld_state.
//
// See: https://developer.apple.com/documentation/kernel/1646572-task_register_dyld_set_dyld_stat
func Task_register_dyld_set_dyld_state(task Task_t, dyld_state uint8) Kern_return_t {
	result, callErr := tryTask_register_dyld_set_dyld_state(task, dyld_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_register_dyld_shared_cache_image_info func(task Task_t, dyld_cache_image Dyld_kernel_image_info_t, no_cache Boolean_t, private_cache Boolean_t) Kern_return_t
var _task_register_dyld_shared_cache_image_infoErr error

func tryTask_register_dyld_shared_cache_image_info(task Task_t, dyld_cache_image Dyld_kernel_image_info_t, no_cache Boolean_t, private_cache Boolean_t) (Kern_return_t, error) {
	if _task_register_dyld_shared_cache_image_info == nil {
		return *new(Kern_return_t), symbolCallError("task_register_dyld_shared_cache_image_info", "10.12", _task_register_dyld_shared_cache_image_infoErr)
	}
	return _task_register_dyld_shared_cache_image_info(task, dyld_cache_image, no_cache, private_cache), nil
}

// Task_register_dyld_shared_cache_image_info.
//
// See: https://developer.apple.com/documentation/kernel/1646581-task_register_dyld_shared_cache_
func Task_register_dyld_shared_cache_image_info(task Task_t, dyld_cache_image Dyld_kernel_image_info_t, no_cache Boolean_t, private_cache Boolean_t) Kern_return_t {
	result, callErr := tryTask_register_dyld_shared_cache_image_info(task, dyld_cache_image, no_cache, private_cache)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_register_hardened_exception_handler func(task Task_t, signed_pc_key uint32, exceptions_allowed Exception_mask_t, behaviors_allowed Exception_behavior_t, flavors_allowed Thread_state_flavor_t, new_exception_port uint32) Kern_return_t
var _task_register_hardened_exception_handlerErr error

func tryTask_register_hardened_exception_handler(task Task_t, signed_pc_key uint32, exceptions_allowed Exception_mask_t, behaviors_allowed Exception_behavior_t, flavors_allowed Thread_state_flavor_t, new_exception_port uint32) (Kern_return_t, error) {
	if _task_register_hardened_exception_handler == nil {
		return *new(Kern_return_t), symbolCallError("task_register_hardened_exception_handler", "15.0", _task_register_hardened_exception_handlerErr)
	}
	return _task_register_hardened_exception_handler(task, signed_pc_key, exceptions_allowed, behaviors_allowed, flavors_allowed, new_exception_port), nil
}

// Task_register_hardened_exception_handler.
//
// See: https://developer.apple.com/documentation/kernel/4360028-task_register_hardened_exception
func Task_register_hardened_exception_handler(task Task_t, signed_pc_key uint32, exceptions_allowed Exception_mask_t, behaviors_allowed Exception_behavior_t, flavors_allowed Thread_state_flavor_t, new_exception_port uint32) Kern_return_t {
	result, callErr := tryTask_register_hardened_exception_handler(task, signed_pc_key, exceptions_allowed, behaviors_allowed, flavors_allowed, new_exception_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_restartable_ranges_register func(task Task_t, ranges Task_restartable_range_array_t, count Mach_msg_type_number_t) Kern_return_t
var _task_restartable_ranges_registerErr error

func tryTask_restartable_ranges_register(task Task_t, ranges Task_restartable_range_array_t, count Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_restartable_ranges_register == nil {
		return *new(Kern_return_t), symbolCallError("task_restartable_ranges_register", "10.15", _task_restartable_ranges_registerErr)
	}
	return _task_restartable_ranges_register(task, ranges, count), nil
}

// Task_restartable_ranges_register.
//
// See: https://developer.apple.com/documentation/kernel/3143268-task_restartable_ranges_register
func Task_restartable_ranges_register(task Task_t, ranges Task_restartable_range_array_t, count Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_restartable_ranges_register(task, ranges, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_restartable_ranges_synchronize func(task Task_t) Kern_return_t
var _task_restartable_ranges_synchronizeErr error

func tryTask_restartable_ranges_synchronize(task Task_t) (Kern_return_t, error) {
	if _task_restartable_ranges_synchronize == nil {
		return *new(Kern_return_t), symbolCallError("task_restartable_ranges_synchronize", "10.15", _task_restartable_ranges_synchronizeErr)
	}
	return _task_restartable_ranges_synchronize(task), nil
}

// Task_restartable_ranges_synchronize.
//
// See: https://developer.apple.com/documentation/kernel/3143269-task_restartable_ranges_synchron
func Task_restartable_ranges_synchronize(task Task_t) Kern_return_t {
	result, callErr := tryTask_restartable_ranges_synchronize(task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_resume func(target_task Task_read_t) Kern_return_t
var _task_resumeErr error

func tryTask_resume(target_task Task_read_t) (Kern_return_t, error) {
	if _task_resume == nil {
		return *new(Kern_return_t), symbolCallError("task_resume", "10.0", _task_resumeErr)
	}
	return _task_resume(target_task), nil
}

// Task_resume.
//
// See: https://developer.apple.com/documentation/kernel/1537977-task_resume
func Task_resume(target_task Task_read_t) Kern_return_t {
	result, callErr := tryTask_resume(target_task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_resume2 func(suspend_token Task_suspension_token_t) Kern_return_t
var _task_resume2Err error

func tryTask_resume2(suspend_token Task_suspension_token_t) (Kern_return_t, error) {
	if _task_resume2 == nil {
		return *new(Kern_return_t), symbolCallError("task_resume2", "10.9", _task_resume2Err)
	}
	return _task_resume2(suspend_token), nil
}

// Task_resume2.
//
// See: https://developer.apple.com/documentation/kernel/1537653-task_resume2
func Task_resume2(suspend_token Task_suspension_token_t) Kern_return_t {
	result, callErr := tryTask_resume2(suspend_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_sample func(task Task_t, reply uint32) Kern_return_t
var _task_sampleErr error

func tryTask_sample(task Task_t, reply uint32) (Kern_return_t, error) {
	if _task_sample == nil {
		return *new(Kern_return_t), symbolCallError("task_sample", "10.0", _task_sampleErr)
	}
	return _task_sample(task, reply), nil
}

// Task_sample.
//
// See: https://developer.apple.com/documentation/kernel/1537655-task_sample
func Task_sample(task Task_t, reply uint32) Kern_return_t {
	result, callErr := tryTask_sample(task, reply)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_corpse_forking_behavior func(task Task_t, behavior Task_corpse_forking_behavior_t) Kern_return_t
var _task_set_corpse_forking_behaviorErr error

func tryTask_set_corpse_forking_behavior(task Task_t, behavior Task_corpse_forking_behavior_t) (Kern_return_t, error) {
	if _task_set_corpse_forking_behavior == nil {
		return *new(Kern_return_t), symbolCallError("task_set_corpse_forking_behavior", "12.0", _task_set_corpse_forking_behaviorErr)
	}
	return _task_set_corpse_forking_behavior(task, behavior), nil
}

// Task_set_corpse_forking_behavior.
//
// See: https://developer.apple.com/documentation/kernel/3852602-task_set_corpse_forking_behavior
func Task_set_corpse_forking_behavior(task Task_t, behavior Task_corpse_forking_behavior_t) Kern_return_t {
	result, callErr := tryTask_set_corpse_forking_behavior(task, behavior)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_emulation func(target_port Task_t, routine_entry_pt Vm_address_t, routine_number int32) Kern_return_t
var _task_set_emulationErr error

func tryTask_set_emulation(target_port Task_t, routine_entry_pt Vm_address_t, routine_number int32) (Kern_return_t, error) {
	if _task_set_emulation == nil {
		return *new(Kern_return_t), symbolCallError("task_set_emulation", "10.0", _task_set_emulationErr)
	}
	return _task_set_emulation(target_port, routine_entry_pt, routine_number), nil
}

// Task_set_emulation.
//
// See: https://developer.apple.com/documentation/kernel/1537794-task_set_emulation
func Task_set_emulation(target_port Task_t, routine_entry_pt Vm_address_t, routine_number int32) Kern_return_t {
	result, callErr := tryTask_set_emulation(target_port, routine_entry_pt, routine_number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_emulation_vector func(task Task_t, vector_start int32, emulation_vector Emulation_vector_t, emulation_vectorCnt Mach_msg_type_number_t) Kern_return_t
var _task_set_emulation_vectorErr error

func tryTask_set_emulation_vector(task Task_t, vector_start int32, emulation_vector Emulation_vector_t, emulation_vectorCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_set_emulation_vector == nil {
		return *new(Kern_return_t), symbolCallError("task_set_emulation_vector", "10.0", _task_set_emulation_vectorErr)
	}
	return _task_set_emulation_vector(task, vector_start, emulation_vector, emulation_vectorCnt), nil
}

// Task_set_emulation_vector.
//
// See: https://developer.apple.com/documentation/kernel/1537907-task_set_emulation_vector
func Task_set_emulation_vector(task Task_t, vector_start int32, emulation_vector Emulation_vector_t, emulation_vectorCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_set_emulation_vector(task, vector_start, emulation_vector, emulation_vectorCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_exc_guard_behavior func(task Task_t, behavior Task_exc_guard_behavior_t) Kern_return_t
var _task_set_exc_guard_behaviorErr error

func tryTask_set_exc_guard_behavior(task Task_t, behavior Task_exc_guard_behavior_t) (Kern_return_t, error) {
	if _task_set_exc_guard_behavior == nil {
		return *new(Kern_return_t), symbolCallError("task_set_exc_guard_behavior", "10.15", _task_set_exc_guard_behaviorErr)
	}
	return _task_set_exc_guard_behavior(task, behavior), nil
}

// Task_set_exc_guard_behavior.
//
// See: https://developer.apple.com/documentation/kernel/3197705-task_set_exc_guard_behavior
func Task_set_exc_guard_behavior(task Task_t, behavior Task_exc_guard_behavior_t) Kern_return_t {
	result, callErr := tryTask_set_exc_guard_behavior(task, behavior)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_exception_ports func(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t
var _task_set_exception_portsErr error

func tryTask_set_exception_ports(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) (Kern_return_t, error) {
	if _task_set_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("task_set_exception_ports", "10.0", _task_set_exception_portsErr)
	}
	return _task_set_exception_ports(task, exception_mask, new_port, behavior, new_flavor), nil
}

// Task_set_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1538049-task_set_exception_ports
func Task_set_exception_ports(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t {
	result, callErr := tryTask_set_exception_ports(task, exception_mask, new_port, behavior, new_flavor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_info func(target_task Task_t, flavor Task_flavor_t, task_info_in Task_info_t, task_info_inCnt Mach_msg_type_number_t) Kern_return_t
var _task_set_infoErr error

func tryTask_set_info(target_task Task_t, flavor Task_flavor_t, task_info_in Task_info_t, task_info_inCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_set_info == nil {
		return *new(Kern_return_t), symbolCallError("task_set_info", "10.0", _task_set_infoErr)
	}
	return _task_set_info(target_task, flavor, task_info_in, task_info_inCnt), nil
}

// Task_set_info.
//
// See: https://developer.apple.com/documentation/kernel/1537704-task_set_info
func Task_set_info(target_task Task_t, flavor Task_flavor_t, task_info_in Task_info_t, task_info_inCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_set_info(target_task, flavor, task_info_in, task_info_inCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_mach_voucher func(task Task_t, voucher Ipc_voucher_t) Kern_return_t
var _task_set_mach_voucherErr error

func tryTask_set_mach_voucher(task Task_t, voucher Ipc_voucher_t) (Kern_return_t, error) {
	if _task_set_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("task_set_mach_voucher", "10.10", _task_set_mach_voucherErr)
	}
	return _task_set_mach_voucher(task, voucher), nil
}

// Task_set_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1538121-task_set_mach_voucher
func Task_set_mach_voucher(task Task_t, voucher Ipc_voucher_t) Kern_return_t {
	result, callErr := tryTask_set_mach_voucher(task, voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_phys_footprint_limit func(task Task_t, new_limit int32, old_limit *int32) Kern_return_t
var _task_set_phys_footprint_limitErr error

func tryTask_set_phys_footprint_limit(task Task_t, new_limit int32, old_limit *int32) (Kern_return_t, error) {
	if _task_set_phys_footprint_limit == nil {
		return *new(Kern_return_t), symbolCallError("task_set_phys_footprint_limit", "10.9", _task_set_phys_footprint_limitErr)
	}
	return _task_set_phys_footprint_limit(task, new_limit, old_limit), nil
}

// Task_set_phys_footprint_limit.
//
// See: https://developer.apple.com/documentation/kernel/1538131-task_set_phys_footprint_limit
func Task_set_phys_footprint_limit(task Task_t, new_limit int32, old_limit *int32) Kern_return_t {
	result, callErr := tryTask_set_phys_footprint_limit(task, new_limit, old_limit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_policy func(task Task_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t, change Boolean_t) Kern_return_t
var _task_set_policyErr error

func tryTask_set_policy(task Task_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t, change Boolean_t) (Kern_return_t, error) {
	if _task_set_policy == nil {
		return *new(Kern_return_t), symbolCallError("task_set_policy", "10.0", _task_set_policyErr)
	}
	return _task_set_policy(task, pset, policy, base, baseCnt, limit, limitCnt, change), nil
}

// Task_set_policy.
//
// See: https://developer.apple.com/documentation/kernel/1537717-task_set_policy
func Task_set_policy(task Task_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t, change Boolean_t) Kern_return_t {
	result, callErr := tryTask_set_policy(task, pset, policy, base, baseCnt, limit, limitCnt, change)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_port_space func(task Ipc_space_t, table_entries int32) Kern_return_t
var _task_set_port_spaceErr error

func tryTask_set_port_space(task Ipc_space_t, table_entries int32) (Kern_return_t, error) {
	if _task_set_port_space == nil {
		return *new(Kern_return_t), symbolCallError("task_set_port_space", "10.0", _task_set_port_spaceErr)
	}
	return _task_set_port_space(task, table_entries), nil
}

// Task_set_port_space.
//
// See: https://developer.apple.com/documentation/kernel/1578836-task_set_port_space
func Task_set_port_space(task Ipc_space_t, table_entries int32) Kern_return_t {
	result, callErr := tryTask_set_port_space(task, table_entries)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_ras_pc func(target_task Task_t, basepc Vm_address_t, boundspc Vm_address_t) Kern_return_t
var _task_set_ras_pcErr error

func tryTask_set_ras_pc(target_task Task_t, basepc Vm_address_t, boundspc Vm_address_t) (Kern_return_t, error) {
	if _task_set_ras_pc == nil {
		return *new(Kern_return_t), symbolCallError("task_set_ras_pc", "10.0", _task_set_ras_pcErr)
	}
	return _task_set_ras_pc(target_task, basepc, boundspc), nil
}

// Task_set_ras_pc.
//
// See: https://developer.apple.com/documentation/kernel/1537742-task_set_ras_pc
func Task_set_ras_pc(target_task Task_t, basepc Vm_address_t, boundspc Vm_address_t) Kern_return_t {
	result, callErr := tryTask_set_ras_pc(target_task, basepc, boundspc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_special_port func(task Task_t, which_port int32, special_port uint32) Kern_return_t
var _task_set_special_portErr error

func tryTask_set_special_port(task Task_t, which_port int32, special_port uint32) (Kern_return_t, error) {
	if _task_set_special_port == nil {
		return *new(Kern_return_t), symbolCallError("task_set_special_port", "10.0", _task_set_special_portErr)
	}
	return _task_set_special_port(task, which_port, special_port), nil
}

// Task_set_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1537676-task_set_special_port
func Task_set_special_port(task Task_t, which_port int32, special_port uint32) Kern_return_t {
	result, callErr := tryTask_set_special_port(task, which_port, special_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_set_state func(task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t
var _task_set_stateErr error

func tryTask_set_state(task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_set_state == nil {
		return *new(Kern_return_t), symbolCallError("task_set_state", "10.6", _task_set_stateErr)
	}
	return _task_set_state(task, flavor, new_state, new_stateCnt), nil
}

// Task_set_state.
//
// See: https://developer.apple.com/documentation/kernel/1537962-task_set_state
func Task_set_state(task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_set_state(task, flavor, new_state, new_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_suspend func(target_task Task_read_t) Kern_return_t
var _task_suspendErr error

func tryTask_suspend(target_task Task_read_t) (Kern_return_t, error) {
	if _task_suspend == nil {
		return *new(Kern_return_t), symbolCallError("task_suspend", "10.0", _task_suspendErr)
	}
	return _task_suspend(target_task), nil
}

// Task_suspend.
//
// See: https://developer.apple.com/documentation/kernel/1538100-task_suspend
func Task_suspend(target_task Task_read_t) Kern_return_t {
	result, callErr := tryTask_suspend(target_task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_suspend2 func(target_task Task_read_t, suspend_token *Task_suspension_token_t) Kern_return_t
var _task_suspend2Err error

func tryTask_suspend2(target_task Task_read_t, suspend_token *Task_suspension_token_t) (Kern_return_t, error) {
	if _task_suspend2 == nil {
		return *new(Kern_return_t), symbolCallError("task_suspend2", "10.9", _task_suspend2Err)
	}
	return _task_suspend2(target_task, suspend_token), nil
}

// Task_suspend2.
//
// See: https://developer.apple.com/documentation/kernel/1538207-task_suspend2
func Task_suspend2(target_task Task_read_t, suspend_token *Task_suspension_token_t) Kern_return_t {
	result, callErr := tryTask_suspend2(target_task, suspend_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_swap_exception_ports func(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _task_swap_exception_portsErr error

func tryTask_swap_exception_ports(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _task_swap_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("task_swap_exception_ports", "10.0", _task_swap_exception_portsErr)
	}
	return _task_swap_exception_ports(task, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlers, old_behaviors, old_flavors), nil
}

// Task_swap_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1537854-task_swap_exception_ports
func Task_swap_exception_ports(task Task_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryTask_swap_exception_ports(task, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlers, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_swap_mach_voucher func(task Task_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) Kern_return_t
var _task_swap_mach_voucherErr error

func tryTask_swap_mach_voucher(task Task_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) (Kern_return_t, error) {
	if _task_swap_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("task_swap_mach_voucher", "10.10", _task_swap_mach_voucherErr)
	}
	return _task_swap_mach_voucher(task, new_voucher, old_voucher), nil
}

// Task_swap_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1537722-task_swap_mach_voucher
func Task_swap_mach_voucher(task Task_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) Kern_return_t {
	result, callErr := tryTask_swap_mach_voucher(task, new_voucher, old_voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_terminate func(target_task Task_t) Kern_return_t
var _task_terminateErr error

func tryTask_terminate(target_task Task_t) (Kern_return_t, error) {
	if _task_terminate == nil {
		return *new(Kern_return_t), symbolCallError("task_terminate", "10.0", _task_terminateErr)
	}
	return _task_terminate(target_task), nil
}

// Task_terminate.
//
// See: https://developer.apple.com/documentation/kernel/1537817-task_terminate
func Task_terminate(target_task Task_t) Kern_return_t {
	result, callErr := tryTask_terminate(target_task)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_test_async_upcall_propagation func(task Task_t, port uint32, qos int32, iotier int32) Kern_return_t
var _task_test_async_upcall_propagationErr error

func tryTask_test_async_upcall_propagation(task Task_t, port uint32, qos int32, iotier int32) (Kern_return_t, error) {
	if _task_test_async_upcall_propagation == nil {
		return *new(Kern_return_t), symbolCallError("task_test_async_upcall_propagation", "12.3", _task_test_async_upcall_propagationErr)
	}
	return _task_test_async_upcall_propagation(task, port, qos, iotier), nil
}

// Task_test_async_upcall_propagation.
//
// See: https://developer.apple.com/documentation/kernel/3917694-task_test_async_upcall_propagati
func Task_test_async_upcall_propagation(task Task_t, port uint32, qos int32, iotier int32) Kern_return_t {
	result, callErr := tryTask_test_async_upcall_propagation(task, port, qos, iotier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_test_sync_upcall func(task Task_t, port uint32) Kern_return_t
var _task_test_sync_upcallErr error

func tryTask_test_sync_upcall(task Task_t, port uint32) (Kern_return_t, error) {
	if _task_test_sync_upcall == nil {
		return *new(Kern_return_t), symbolCallError("task_test_sync_upcall", "12.0", _task_test_sync_upcallErr)
	}
	return _task_test_sync_upcall(task, port), nil
}

// Task_test_sync_upcall.
//
// See: https://developer.apple.com/documentation/kernel/3753679-task_test_sync_upcall
func Task_test_sync_upcall(task Task_t, port uint32) Kern_return_t {
	result, callErr := tryTask_test_sync_upcall(task, port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_threads func(target_task Task_inspect_t, act_list *Thread_act_array_t, act_listCnt *Mach_msg_type_number_t) Kern_return_t
var _task_threadsErr error

func tryTask_threads(target_task Task_inspect_t, act_list *Thread_act_array_t, act_listCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_threads == nil {
		return *new(Kern_return_t), symbolCallError("task_threads", "10.0", _task_threadsErr)
	}
	return _task_threads(target_task, act_list, act_listCnt), nil
}

// Task_threads.
//
// See: https://developer.apple.com/documentation/kernel/1537751-task_threads
func Task_threads(target_task Task_inspect_t, act_list *Thread_act_array_t, act_listCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_threads(target_task, act_list, act_listCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_unregister_dyld_image_infos func(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) Kern_return_t
var _task_unregister_dyld_image_infosErr error

func tryTask_unregister_dyld_image_infos(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_unregister_dyld_image_infos == nil {
		return *new(Kern_return_t), symbolCallError("task_unregister_dyld_image_infos", "10.12", _task_unregister_dyld_image_infosErr)
	}
	return _task_unregister_dyld_image_infos(task, dyld_images, dyld_imagesCnt), nil
}

// Task_unregister_dyld_image_infos.
//
// See: https://developer.apple.com/documentation/kernel/1646587-task_unregister_dyld_image_infos
func Task_unregister_dyld_image_infos(task Task_t, dyld_images Dyld_kernel_image_info_array_t, dyld_imagesCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_unregister_dyld_image_infos(task, dyld_images, dyld_imagesCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _task_zone_info func(target_task Task_inspect_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Task_zone_info_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _task_zone_infoErr error

func tryTask_zone_info(target_task Task_inspect_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Task_zone_info_array_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _task_zone_info == nil {
		return *new(Kern_return_t), symbolCallError("task_zone_info", "10.7", _task_zone_infoErr)
	}
	return _task_zone_info(target_task, names, namesCnt, info, infoCnt), nil
}

// Task_zone_info.
//
// See: https://developer.apple.com/documentation/kernel/1537645-task_zone_info
func Task_zone_info(target_task Task_inspect_t, names *Mach_zone_name_array_t, namesCnt *Mach_msg_type_number_t, info *Task_zone_info_array_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryTask_zone_info(target_task, names, namesCnt, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tgamma func(arg0 float64) float64
var _tgammaErr error

func tryTgamma(arg0 float64) (float64, error) {
	if _tgamma == nil {
		return 0.0, symbolCallError("tgamma", "10.10", _tgammaErr)
	}
	return _tgamma(arg0), nil
}

// Tgamma.
//
// See: https://developer.apple.com/documentation/kernel/1557229-tgamma
func Tgamma(arg0 float64) float64 {
	result, callErr := tryTgamma(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tgammaf func(arg0 float32) float32
var _tgammafErr error

func tryTgammaf(arg0 float32) (float32, error) {
	if _tgammaf == nil {
		return 0.0, symbolCallError("tgammaf", "10.10", _tgammafErr)
	}
	return _tgammaf(arg0), nil
}

// Tgammaf.
//
// See: https://developer.apple.com/documentation/kernel/1557191-tgammaf
func Tgammaf(arg0 float32) float32 {
	result, callErr := tryTgammaf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _tgammal func(arg0 float64) float64
var _tgammalErr error

func tryTgammal(arg0 float64) (float64, error) {
	if _tgammal == nil {
		return 0.0, symbolCallError("tgammal", "10.10", _tgammalErr)
	}
	return _tgammal(arg0), nil
}

// Tgammal.
//
// See: https://developer.apple.com/documentation/kernel/1557205-tgammal
func Tgammal(arg0 float64) float64 {
	result, callErr := tryTgammal(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_abort func(target_act Thread_act_t) Kern_return_t
var _thread_abortErr error

func tryThread_abort(target_act Thread_act_t) (Kern_return_t, error) {
	if _thread_abort == nil {
		return *new(Kern_return_t), symbolCallError("thread_abort", "10.0", _thread_abortErr)
	}
	return _thread_abort(target_act), nil
}

// Thread_abort.
//
// See: https://developer.apple.com/documentation/kernel/1418578-thread_abort
func Thread_abort(target_act Thread_act_t) Kern_return_t {
	result, callErr := tryThread_abort(target_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_abort_safely func(target_act Thread_act_t) Kern_return_t
var _thread_abort_safelyErr error

func tryThread_abort_safely(target_act Thread_act_t) (Kern_return_t, error) {
	if _thread_abort_safely == nil {
		return *new(Kern_return_t), symbolCallError("thread_abort_safely", "10.0", _thread_abort_safelyErr)
	}
	return _thread_abort_safely(target_act), nil
}

// Thread_abort_safely.
//
// See: https://developer.apple.com/documentation/kernel/1418913-thread_abort_safely
func Thread_abort_safely(target_act Thread_act_t) Kern_return_t {
	result, callErr := tryThread_abort_safely(target_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_adopt_exception_handler func(thread Thread_t, exc_port uint32, exc_mask Exception_mask_t, behavior_mask Exception_behavior_t, flavor_mask Thread_state_flavor_t) Kern_return_t
var _thread_adopt_exception_handlerErr error

func tryThread_adopt_exception_handler(thread Thread_t, exc_port uint32, exc_mask Exception_mask_t, behavior_mask Exception_behavior_t, flavor_mask Thread_state_flavor_t) (Kern_return_t, error) {
	if _thread_adopt_exception_handler == nil {
		return *new(Kern_return_t), symbolCallError("thread_adopt_exception_handler", "15.0", _thread_adopt_exception_handlerErr)
	}
	return _thread_adopt_exception_handler(thread, exc_port, exc_mask, behavior_mask, flavor_mask), nil
}

// Thread_adopt_exception_handler.
//
// See: https://developer.apple.com/documentation/kernel/4360032-thread_adopt_exception_handler
func Thread_adopt_exception_handler(thread Thread_t, exc_port uint32, exc_mask Exception_mask_t, behavior_mask Exception_behavior_t, flavor_mask Thread_state_flavor_t) Kern_return_t {
	result, callErr := tryThread_adopt_exception_handler(thread, exc_port, exc_mask, behavior_mask, flavor_mask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_assign func(thread Thread_act_t, new_set Processor_set_t) Kern_return_t
var _thread_assignErr error

func tryThread_assign(thread Thread_act_t, new_set Processor_set_t) (Kern_return_t, error) {
	if _thread_assign == nil {
		return *new(Kern_return_t), symbolCallError("thread_assign", "10.0", _thread_assignErr)
	}
	return _thread_assign(thread, new_set), nil
}

// Thread_assign.
//
// See: https://developer.apple.com/documentation/kernel/1418581-thread_assign
func Thread_assign(thread Thread_act_t, new_set Processor_set_t) Kern_return_t {
	result, callErr := tryThread_assign(thread, new_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_assign_default func(thread Thread_act_t) Kern_return_t
var _thread_assign_defaultErr error

func tryThread_assign_default(thread Thread_act_t) (Kern_return_t, error) {
	if _thread_assign_default == nil {
		return *new(Kern_return_t), symbolCallError("thread_assign_default", "10.0", _thread_assign_defaultErr)
	}
	return _thread_assign_default(thread), nil
}

// Thread_assign_default.
//
// See: https://developer.apple.com/documentation/kernel/1418700-thread_assign_default
func Thread_assign_default(thread Thread_act_t) Kern_return_t {
	result, callErr := tryThread_assign_default(thread)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_convert_thread_state func(thread Thread_act_t, direction int32, flavor Thread_state_flavor_t, in_state Thread_state_t, in_stateCnt Mach_msg_type_number_t, out_state Thread_state_t, out_stateCnt *Mach_msg_type_number_t) Kern_return_t
var _thread_convert_thread_stateErr error

func tryThread_convert_thread_state(thread Thread_act_t, direction int32, flavor Thread_state_flavor_t, in_state Thread_state_t, in_stateCnt Mach_msg_type_number_t, out_state Thread_state_t, out_stateCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_convert_thread_state == nil {
		return *new(Kern_return_t), symbolCallError("thread_convert_thread_state", "11.0", _thread_convert_thread_stateErr)
	}
	return _thread_convert_thread_state(thread, direction, flavor, in_state, in_stateCnt, out_state, out_stateCnt), nil
}

// Thread_convert_thread_state.
//
// See: https://developer.apple.com/documentation/kernel/3553733-thread_convert_thread_state
func Thread_convert_thread_state(thread Thread_act_t, direction int32, flavor Thread_state_flavor_t, in_state Thread_state_t, in_stateCnt Mach_msg_type_number_t, out_state Thread_state_t, out_stateCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_convert_thread_state(thread, direction, flavor, in_state, in_stateCnt, out_state, out_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_create func(parent_task Task_t, child_act *Thread_act_t) Kern_return_t
var _thread_createErr error

func tryThread_create(parent_task Task_t, child_act *Thread_act_t) (Kern_return_t, error) {
	if _thread_create == nil {
		return *new(Kern_return_t), symbolCallError("thread_create", "10.0", _thread_createErr)
	}
	return _thread_create(parent_task, child_act), nil
}

// Thread_create.
//
// See: https://developer.apple.com/documentation/kernel/1538152-thread_create
func Thread_create(parent_task Task_t, child_act *Thread_act_t) Kern_return_t {
	result, callErr := tryThread_create(parent_task, child_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_create_running func(parent_task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t, child_act *Thread_act_t) Kern_return_t
var _thread_create_runningErr error

func tryThread_create_running(parent_task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t, child_act *Thread_act_t) (Kern_return_t, error) {
	if _thread_create_running == nil {
		return *new(Kern_return_t), symbolCallError("thread_create_running", "10.0", _thread_create_runningErr)
	}
	return _thread_create_running(parent_task, flavor, new_state, new_stateCnt, child_act), nil
}

// Thread_create_running.
//
// See: https://developer.apple.com/documentation/kernel/1537886-thread_create_running
func Thread_create_running(parent_task Task_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t, child_act *Thread_act_t) Kern_return_t {
	result, callErr := tryThread_create_running(parent_task, flavor, new_state, new_stateCnt, child_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_depress_abort func(thread Thread_act_t) Kern_return_t
var _thread_depress_abortErr error

func tryThread_depress_abort(thread Thread_act_t) (Kern_return_t, error) {
	if _thread_depress_abort == nil {
		return *new(Kern_return_t), symbolCallError("thread_depress_abort", "10.0", _thread_depress_abortErr)
	}
	return _thread_depress_abort(thread), nil
}

// Thread_depress_abort.
//
// See: https://developer.apple.com/documentation/kernel/1418712-thread_depress_abort
func Thread_depress_abort(thread Thread_act_t) Kern_return_t {
	result, callErr := tryThread_depress_abort(thread)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_assignment func(thread Thread_inspect_t, assigned_set *Processor_set_name_t) Kern_return_t
var _thread_get_assignmentErr error

func tryThread_get_assignment(thread Thread_inspect_t, assigned_set *Processor_set_name_t) (Kern_return_t, error) {
	if _thread_get_assignment == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_assignment", "10.0", _thread_get_assignmentErr)
	}
	return _thread_get_assignment(thread, assigned_set), nil
}

// Thread_get_assignment.
//
// See: https://developer.apple.com/documentation/kernel/1418698-thread_get_assignment
func Thread_get_assignment(thread Thread_inspect_t, assigned_set *Processor_set_name_t) Kern_return_t {
	result, callErr := tryThread_get_assignment(thread, assigned_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_exception_ports func(thread Thread_act_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _thread_get_exception_portsErr error

func tryThread_get_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _thread_get_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_exception_ports", "10.0", _thread_get_exception_portsErr)
	}
	return _thread_get_exception_ports(thread, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors), nil
}

// Thread_get_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1418945-thread_get_exception_ports
func Thread_get_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryThread_get_exception_ports(thread, exception_mask, masks, masksCnt, old_handlers, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_exception_ports_info func(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _thread_get_exception_ports_infoErr error

func tryThread_get_exception_ports_info(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _thread_get_exception_ports_info == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_exception_ports_info", "11.3", _thread_get_exception_ports_infoErr)
	}
	return _thread_get_exception_ports_info(port, exception_mask, masks, masksCnt, old_handlers_info, old_behaviors, old_flavors), nil
}

// Thread_get_exception_ports_info.
//
// See: https://developer.apple.com/documentation/kernel/3727999-thread_get_exception_ports_info
func Thread_get_exception_ports_info(port uint32, exception_mask Exception_mask_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers_info Exception_handler_info_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryThread_get_exception_ports_info(port, exception_mask, masks, masksCnt, old_handlers_info, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_mach_voucher func(thr_act Thread_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) Kern_return_t
var _thread_get_mach_voucherErr error

func tryThread_get_mach_voucher(thr_act Thread_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) (Kern_return_t, error) {
	if _thread_get_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_mach_voucher", "10.10", _thread_get_mach_voucherErr)
	}
	return _thread_get_mach_voucher(thr_act, which, voucher), nil
}

// Thread_get_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1418540-thread_get_mach_voucher
func Thread_get_mach_voucher(thr_act Thread_read_t, which Mach_voucher_selector_t, voucher *Ipc_voucher_t) Kern_return_t {
	result, callErr := tryThread_get_mach_voucher(thr_act, which, voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_special_port func(thr_act Thread_inspect_t, which_port int32, special_port *uint32) Kern_return_t
var _thread_get_special_portErr error

func tryThread_get_special_port(thr_act Thread_inspect_t, which_port int32, special_port *uint32) (Kern_return_t, error) {
	if _thread_get_special_port == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_special_port", "10.0", _thread_get_special_portErr)
	}
	return _thread_get_special_port(thr_act, which_port, special_port), nil
}

// Thread_get_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1418728-thread_get_special_port
func Thread_get_special_port(thr_act Thread_inspect_t, which_port int32, special_port *uint32) Kern_return_t {
	result, callErr := tryThread_get_special_port(thr_act, which_port, special_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_get_state func(target_act Thread_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t
var _thread_get_stateErr error

func tryThread_get_state(target_act Thread_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_get_state == nil {
		return *new(Kern_return_t), symbolCallError("thread_get_state", "10.0", _thread_get_stateErr)
	}
	return _thread_get_state(target_act, flavor, old_state, old_stateCnt), nil
}

// Thread_get_state.
//
// See: https://developer.apple.com/documentation/kernel/1418576-thread_get_state
func Thread_get_state(target_act Thread_read_t, flavor Thread_state_flavor_t, old_state Thread_state_t, old_stateCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_get_state(target_act, flavor, old_state, old_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_info func(target_act Thread_inspect_t, flavor Thread_flavor_t, thread_info_out Thread_info_t, thread_info_outCnt *Mach_msg_type_number_t) Kern_return_t
var _thread_infoErr error

func tryThread_info(target_act Thread_inspect_t, flavor Thread_flavor_t, thread_info_out Thread_info_t, thread_info_outCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_info == nil {
		return *new(Kern_return_t), symbolCallError("thread_info", "10.0", _thread_infoErr)
	}
	return _thread_info(target_act, flavor, thread_info_out, thread_info_outCnt), nil
}

// Thread_info.
//
// See: https://developer.apple.com/documentation/kernel/1418630-thread_info
func Thread_info(target_act Thread_inspect_t, flavor Thread_flavor_t, thread_info_out Thread_info_t, thread_info_outCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_info(target_act, flavor, thread_info_out, thread_info_outCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_policy func(thr_act Thread_act_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t) Kern_return_t
var _thread_policyErr error

func tryThread_policy(thr_act Thread_act_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t) (Kern_return_t, error) {
	if _thread_policy == nil {
		return *new(Kern_return_t), symbolCallError("thread_policy", "10.0", _thread_policyErr)
	}
	return _thread_policy(thr_act, policy, base, baseCnt, set_limit), nil
}

// Thread_policy.
//
// See: https://developer.apple.com/documentation/kernel/1418640-thread_policy
func Thread_policy(thr_act Thread_act_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, set_limit Boolean_t) Kern_return_t {
	result, callErr := tryThread_policy(thr_act, policy, base, baseCnt, set_limit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_policy_get func(thread Thread_inspect_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) Kern_return_t
var _thread_policy_getErr error

func tryThread_policy_get(thread Thread_inspect_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) (Kern_return_t, error) {
	if _thread_policy_get == nil {
		return *new(Kern_return_t), symbolCallError("thread_policy_get", "10.0", _thread_policy_getErr)
	}
	return _thread_policy_get(thread, flavor, policy_info, policy_infoCnt, get_default), nil
}

// Thread_policy_get.
//
// See: https://developer.apple.com/documentation/kernel/1418518-thread_policy_get
func Thread_policy_get(thread Thread_inspect_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt *Mach_msg_type_number_t, get_default *Boolean_t) Kern_return_t {
	result, callErr := tryThread_policy_get(thread, flavor, policy_info, policy_infoCnt, get_default)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_policy_set func(thread Thread_act_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt Mach_msg_type_number_t) Kern_return_t
var _thread_policy_setErr error

func tryThread_policy_set(thread Thread_act_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_policy_set == nil {
		return *new(Kern_return_t), symbolCallError("thread_policy_set", "10.0", _thread_policy_setErr)
	}
	return _thread_policy_set(thread, flavor, policy_info, policy_infoCnt), nil
}

// Thread_policy_set.
//
// See: https://developer.apple.com/documentation/kernel/1418892-thread_policy_set
func Thread_policy_set(thread Thread_act_t, flavor Thread_policy_flavor_t, policy_info Thread_policy_t, policy_infoCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_policy_set(thread, flavor, policy_info, policy_infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_resume func(target_act Thread_read_t) Kern_return_t
var _thread_resumeErr error

func tryThread_resume(target_act Thread_read_t) (Kern_return_t, error) {
	if _thread_resume == nil {
		return *new(Kern_return_t), symbolCallError("thread_resume", "10.0", _thread_resumeErr)
	}
	return _thread_resume(target_act), nil
}

// Thread_resume.
//
// See: https://developer.apple.com/documentation/kernel/1418926-thread_resume
func Thread_resume(target_act Thread_read_t) Kern_return_t {
	result, callErr := tryThread_resume(target_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_sample func(thread Thread_act_t, reply uint32) Kern_return_t
var _thread_sampleErr error

func tryThread_sample(thread Thread_act_t, reply uint32) (Kern_return_t, error) {
	if _thread_sample == nil {
		return *new(Kern_return_t), symbolCallError("thread_sample", "10.0", _thread_sampleErr)
	}
	return _thread_sample(thread, reply), nil
}

// Thread_sample.
//
// See: https://developer.apple.com/documentation/kernel/1418814-thread_sample
func Thread_sample(thread Thread_act_t, reply uint32) Kern_return_t {
	result, callErr := tryThread_sample(thread, reply)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_set_exception_ports func(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t
var _thread_set_exception_portsErr error

func tryThread_set_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) (Kern_return_t, error) {
	if _thread_set_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("thread_set_exception_ports", "10.0", _thread_set_exception_portsErr)
	}
	return _thread_set_exception_ports(thread, exception_mask, new_port, behavior, new_flavor), nil
}

// Thread_set_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1418619-thread_set_exception_ports
func Thread_set_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t) Kern_return_t {
	result, callErr := tryThread_set_exception_ports(thread, exception_mask, new_port, behavior, new_flavor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_set_mach_voucher func(thr_act Thread_act_t, voucher Ipc_voucher_t) Kern_return_t
var _thread_set_mach_voucherErr error

func tryThread_set_mach_voucher(thr_act Thread_act_t, voucher Ipc_voucher_t) (Kern_return_t, error) {
	if _thread_set_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("thread_set_mach_voucher", "10.10", _thread_set_mach_voucherErr)
	}
	return _thread_set_mach_voucher(thr_act, voucher), nil
}

// Thread_set_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1418834-thread_set_mach_voucher
func Thread_set_mach_voucher(thr_act Thread_act_t, voucher Ipc_voucher_t) Kern_return_t {
	result, callErr := tryThread_set_mach_voucher(thr_act, voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_set_policy func(thr_act Thread_act_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t) Kern_return_t
var _thread_set_policyErr error

func tryThread_set_policy(thr_act Thread_act_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_set_policy == nil {
		return *new(Kern_return_t), symbolCallError("thread_set_policy", "10.0", _thread_set_policyErr)
	}
	return _thread_set_policy(thr_act, pset, policy, base, baseCnt, limit, limitCnt), nil
}

// Thread_set_policy.
//
// See: https://developer.apple.com/documentation/kernel/1418608-thread_set_policy
func Thread_set_policy(thr_act Thread_act_t, pset Processor_set_t, policy Policy_t, base Policy_base_t, baseCnt Mach_msg_type_number_t, limit Policy_limit_t, limitCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_set_policy(thr_act, pset, policy, base, baseCnt, limit, limitCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_set_special_port func(thr_act Thread_act_t, which_port int32, special_port uint32) Kern_return_t
var _thread_set_special_portErr error

func tryThread_set_special_port(thr_act Thread_act_t, which_port int32, special_port uint32) (Kern_return_t, error) {
	if _thread_set_special_port == nil {
		return *new(Kern_return_t), symbolCallError("thread_set_special_port", "10.0", _thread_set_special_portErr)
	}
	return _thread_set_special_port(thr_act, which_port, special_port), nil
}

// Thread_set_special_port.
//
// See: https://developer.apple.com/documentation/kernel/1418995-thread_set_special_port
func Thread_set_special_port(thr_act Thread_act_t, which_port int32, special_port uint32) Kern_return_t {
	result, callErr := tryThread_set_special_port(thr_act, which_port, special_port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_set_state func(target_act Thread_act_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t
var _thread_set_stateErr error

func tryThread_set_state(target_act Thread_act_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _thread_set_state == nil {
		return *new(Kern_return_t), symbolCallError("thread_set_state", "10.0", _thread_set_stateErr)
	}
	return _thread_set_state(target_act, flavor, new_state, new_stateCnt), nil
}

// Thread_set_state.
//
// See: https://developer.apple.com/documentation/kernel/1418827-thread_set_state
func Thread_set_state(target_act Thread_act_t, flavor Thread_state_flavor_t, new_state Thread_state_t, new_stateCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryThread_set_state(target_act, flavor, new_state, new_stateCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_suspend func(target_act Thread_read_t) Kern_return_t
var _thread_suspendErr error

func tryThread_suspend(target_act Thread_read_t) (Kern_return_t, error) {
	if _thread_suspend == nil {
		return *new(Kern_return_t), symbolCallError("thread_suspend", "10.0", _thread_suspendErr)
	}
	return _thread_suspend(target_act), nil
}

// Thread_suspend.
//
// See: https://developer.apple.com/documentation/kernel/1418833-thread_suspend
func Thread_suspend(target_act Thread_read_t) Kern_return_t {
	result, callErr := tryThread_suspend(target_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_swap_exception_ports func(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t
var _thread_swap_exception_portsErr error

func tryThread_swap_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) (Kern_return_t, error) {
	if _thread_swap_exception_ports == nil {
		return *new(Kern_return_t), symbolCallError("thread_swap_exception_ports", "10.0", _thread_swap_exception_portsErr)
	}
	return _thread_swap_exception_ports(thread, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlers, old_behaviors, old_flavors), nil
}

// Thread_swap_exception_ports.
//
// See: https://developer.apple.com/documentation/kernel/1418969-thread_swap_exception_ports
func Thread_swap_exception_ports(thread Thread_act_t, exception_mask Exception_mask_t, new_port uint32, behavior Exception_behavior_t, new_flavor Thread_state_flavor_t, masks Exception_mask_array_t, masksCnt *Mach_msg_type_number_t, old_handlers Exception_handler_array_t, old_behaviors Exception_behavior_array_t, old_flavors Exception_flavor_array_t) Kern_return_t {
	result, callErr := tryThread_swap_exception_ports(thread, exception_mask, new_port, behavior, new_flavor, masks, masksCnt, old_handlers, old_behaviors, old_flavors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_swap_mach_voucher func(thr_act Thread_act_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) Kern_return_t
var _thread_swap_mach_voucherErr error

func tryThread_swap_mach_voucher(thr_act Thread_act_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) (Kern_return_t, error) {
	if _thread_swap_mach_voucher == nil {
		return *new(Kern_return_t), symbolCallError("thread_swap_mach_voucher", "10.10", _thread_swap_mach_voucherErr)
	}
	return _thread_swap_mach_voucher(thr_act, new_voucher, old_voucher), nil
}

// Thread_swap_mach_voucher.
//
// See: https://developer.apple.com/documentation/kernel/1418678-thread_swap_mach_voucher
func Thread_swap_mach_voucher(thr_act Thread_act_t, new_voucher Ipc_voucher_t, old_voucher *Ipc_voucher_t) Kern_return_t {
	result, callErr := tryThread_swap_mach_voucher(thr_act, new_voucher, old_voucher)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_terminate func(target_act Thread_act_t) Kern_return_t
var _thread_terminateErr error

func tryThread_terminate(target_act Thread_act_t) (Kern_return_t, error) {
	if _thread_terminate == nil {
		return *new(Kern_return_t), symbolCallError("thread_terminate", "10.0", _thread_terminateErr)
	}
	return _thread_terminate(target_act), nil
}

// Thread_terminate.
//
// See: https://developer.apple.com/documentation/kernel/1418708-thread_terminate
func Thread_terminate(target_act Thread_act_t) Kern_return_t {
	result, callErr := tryThread_terminate(target_act)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _thread_wire func(host_priv Host_priv_t, thread Thread_act_t, wired Boolean_t) Kern_return_t
var _thread_wireErr error

func tryThread_wire(host_priv Host_priv_t, thread Thread_act_t, wired Boolean_t) (Kern_return_t, error) {
	if _thread_wire == nil {
		return *new(Kern_return_t), symbolCallError("thread_wire", "10.0", _thread_wireErr)
	}
	return _thread_wire(host_priv, thread, wired), nil
}

// Thread_wire.
//
// See: https://developer.apple.com/documentation/kernel/1588756-thread_wire
func Thread_wire(host_priv Host_priv_t, thread Thread_act_t, wired Boolean_t) Kern_return_t {
	result, callErr := tryThread_wire(host_priv, thread, wired)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _timingsafe_bcmp func(b1 unsafe.Pointer, b2 unsafe.Pointer, n uintptr) int32
var _timingsafe_bcmpErr error

func tryTimingsafe_bcmp(b1 unsafe.Pointer, b2 unsafe.Pointer, n uintptr) (int32, error) {
	if _timingsafe_bcmp == nil {
		return 0, symbolCallError("timingsafe_bcmp", "10.15", _timingsafe_bcmpErr)
	}
	return _timingsafe_bcmp(b1, b2, n), nil
}

// Timingsafe_bcmp.
//
// See: https://developer.apple.com/documentation/kernel/3197718-timingsafe_bcmp
func Timingsafe_bcmp(b1 unsafe.Pointer, b2 unsafe.Pointer, n uintptr) int32 {
	result, callErr := tryTimingsafe_bcmp(b1, b2, n)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _trunc func(arg0 float64) float64
var _truncErr error

func tryTrunc(arg0 float64) (float64, error) {
	if _trunc == nil {
		return 0.0, symbolCallError("trunc", "10.10", _truncErr)
	}
	return _trunc(arg0), nil
}

// Trunc.
//
// See: https://developer.apple.com/documentation/kernel/1557333-trunc
func Trunc(arg0 float64) float64 {
	result, callErr := tryTrunc(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _truncf func(arg0 float32) float32
var _truncfErr error

func tryTruncf(arg0 float32) (float32, error) {
	if _truncf == nil {
		return 0.0, symbolCallError("truncf", "10.10", _truncfErr)
	}
	return _truncf(arg0), nil
}

// Truncf.
//
// See: https://developer.apple.com/documentation/kernel/1557223-truncf
func Truncf(arg0 float32) float32 {
	result, callErr := tryTruncf(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _truncl func(arg0 float64) float64
var _trunclErr error

func tryTruncl(arg0 float64) (float64, error) {
	if _truncl == nil {
		return 0.0, symbolCallError("truncl", "10.10", _trunclErr)
	}
	return _truncl(arg0), nil
}

// Truncl.
//
// See: https://developer.apple.com/documentation/kernel/1557153-truncl
func Truncl(arg0 float64) float64 {
	result, callErr := tryTruncl(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uuid_clear func(uu *[16]byte)
var _uuid_clearErr error

func tryUuid_clear(uu [16]byte) error {
	if _uuid_clear == nil {
		return symbolCallError("uuid_clear", "10.4", _uuid_clearErr)
	}
	_uuid_clear(&uu)
	return nil
}

// Uuid_clear.
//
// See: https://developer.apple.com/documentation/kernel/1470606-uuid_clear
func Uuid_clear(uu [16]byte) {
	if callErr := tryUuid_clear(uu); callErr != nil {
		panic(callErr)
	}
}

var _uuid_compare func(uu1 *[16]byte, uu2 *[16]byte) int32
var _uuid_compareErr error

func tryUuid_compare(uu1 [16]byte, uu2 [16]byte) (int32, error) {
	if _uuid_compare == nil {
		return 0, symbolCallError("uuid_compare", "10.4", _uuid_compareErr)
	}
	return _uuid_compare(&uu1, &uu2), nil
}

// Uuid_compare.
//
// See: https://developer.apple.com/documentation/kernel/1470610-uuid_compare
func Uuid_compare(uu1 [16]byte, uu2 [16]byte) int32 {
	result, callErr := tryUuid_compare(uu1, uu2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uuid_copy func(dst *[16]byte, src *[16]byte)
var _uuid_copyErr error

func tryUuid_copy(dst [16]byte, src [16]byte) error {
	if _uuid_copy == nil {
		return symbolCallError("uuid_copy", "10.4", _uuid_copyErr)
	}
	_uuid_copy(&dst, &src)
	return nil
}

// Uuid_copy.
//
// See: https://developer.apple.com/documentation/kernel/1470608-uuid_copy
func Uuid_copy(dst [16]byte, src [16]byte) {
	if callErr := tryUuid_copy(dst, src); callErr != nil {
		panic(callErr)
	}
}

var _uuid_generate func(out *[16]byte)
var _uuid_generateErr error

func tryUuid_generate(out [16]byte) error {
	if _uuid_generate == nil {
		return symbolCallError("uuid_generate", "10.4", _uuid_generateErr)
	}
	_uuid_generate(&out)
	return nil
}

// Uuid_generate.
//
// See: https://developer.apple.com/documentation/kernel/1470614-uuid_generate
func Uuid_generate(out [16]byte) {
	if callErr := tryUuid_generate(out); callErr != nil {
		panic(callErr)
	}
}

var _uuid_generate_random func(out *[16]byte)
var _uuid_generate_randomErr error

func tryUuid_generate_random(out [16]byte) error {
	if _uuid_generate_random == nil {
		return symbolCallError("uuid_generate_random", "10.4", _uuid_generate_randomErr)
	}
	_uuid_generate_random(&out)
	return nil
}

// Uuid_generate_random.
//
// See: https://developer.apple.com/documentation/kernel/1470612-uuid_generate_random
func Uuid_generate_random(out [16]byte) {
	if callErr := tryUuid_generate_random(out); callErr != nil {
		panic(callErr)
	}
}

var _uuid_generate_time func(out *[16]byte)
var _uuid_generate_timeErr error

func tryUuid_generate_time(out [16]byte) error {
	if _uuid_generate_time == nil {
		return symbolCallError("uuid_generate_time", "10.4", _uuid_generate_timeErr)
	}
	_uuid_generate_time(&out)
	return nil
}

// Uuid_generate_time.
//
// See: https://developer.apple.com/documentation/kernel/1470625-uuid_generate_time
func Uuid_generate_time(out [16]byte) {
	if callErr := tryUuid_generate_time(out); callErr != nil {
		panic(callErr)
	}
}

var _uuid_is_null func(uu *[16]byte) int32
var _uuid_is_nullErr error

func tryUuid_is_null(uu [16]byte) (int32, error) {
	if _uuid_is_null == nil {
		return 0, symbolCallError("uuid_is_null", "10.4", _uuid_is_nullErr)
	}
	return _uuid_is_null(&uu), nil
}

// Uuid_is_null.
//
// See: https://developer.apple.com/documentation/kernel/1470616-uuid_is_null
func Uuid_is_null(uu [16]byte) int32 {
	result, callErr := tryUuid_is_null(uu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uuid_parse func(in Uuid_string_t, uu *[16]byte) int32
var _uuid_parseErr error

func tryUuid_parse(in Uuid_string_t, uu [16]byte) (int32, error) {
	if _uuid_parse == nil {
		return 0, symbolCallError("uuid_parse", "10.4", _uuid_parseErr)
	}
	return _uuid_parse(in, &uu), nil
}

// Uuid_parse.
//
// See: https://developer.apple.com/documentation/kernel/1470624-uuid_parse
func Uuid_parse(in Uuid_string_t, uu [16]byte) int32 {
	result, callErr := tryUuid_parse(in, uu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uuid_unparse func(uu *[16]byte, out Uuid_string_t)
var _uuid_unparseErr error

func tryUuid_unparse(uu [16]byte, out Uuid_string_t) error {
	if _uuid_unparse == nil {
		return symbolCallError("uuid_unparse", "10.4", _uuid_unparseErr)
	}
	_uuid_unparse(&uu, out)
	return nil
}

// Uuid_unparse.
//
// See: https://developer.apple.com/documentation/kernel/1470620-uuid_unparse
func Uuid_unparse(uu [16]byte, out Uuid_string_t) {
	if callErr := tryUuid_unparse(uu, out); callErr != nil {
		panic(callErr)
	}
}

var _uuid_unparse_lower func(uu *[16]byte, out Uuid_string_t)
var _uuid_unparse_lowerErr error

func tryUuid_unparse_lower(uu [16]byte, out Uuid_string_t) error {
	if _uuid_unparse_lower == nil {
		return symbolCallError("uuid_unparse_lower", "10.4", _uuid_unparse_lowerErr)
	}
	_uuid_unparse_lower(&uu, out)
	return nil
}

// Uuid_unparse_lower.
//
// See: https://developer.apple.com/documentation/kernel/1470622-uuid_unparse_lower
func Uuid_unparse_lower(uu [16]byte, out Uuid_string_t) {
	if callErr := tryUuid_unparse_lower(uu, out); callErr != nil {
		panic(callErr)
	}
}

var _uuid_unparse_upper func(uu *[16]byte, out Uuid_string_t)
var _uuid_unparse_upperErr error

func tryUuid_unparse_upper(uu [16]byte, out Uuid_string_t) error {
	if _uuid_unparse_upper == nil {
		return symbolCallError("uuid_unparse_upper", "10.4", _uuid_unparse_upperErr)
	}
	_uuid_unparse_upper(&uu, out)
	return nil
}

// Uuid_unparse_upper.
//
// See: https://developer.apple.com/documentation/kernel/1470618-uuid_unparse_upper
func Uuid_unparse_upper(uu [16]byte, out Uuid_string_t) {
	if callErr := tryUuid_unparse_upper(uu, out); callErr != nil {
		panic(callErr)
	}
}

var _vm_allocate func(target_task uint32, address *Vm_address_t, size Vm_size_t, flags int32) Kern_return_t
var _vm_allocateErr error

func tryVm_allocate(target_task uint32, address *Vm_address_t, size Vm_size_t, flags int32) (Kern_return_t, error) {
	if _vm_allocate == nil {
		return *new(Kern_return_t), symbolCallError("vm_allocate", "10.0", _vm_allocateErr)
	}
	return _vm_allocate(target_task, address, size, flags), nil
}

// Vm_allocate.
//
// See: https://developer.apple.com/documentation/kernel/1585381-vm_allocate
func Vm_allocate(target_task uint32, address *Vm_address_t, size Vm_size_t, flags int32) Kern_return_t {
	result, callErr := tryVm_allocate(target_task, address, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_allocate_cpm func(host_priv Host_priv_t, task uint32, address *Vm_address_t, size Vm_size_t, flags int32) Kern_return_t
var _vm_allocate_cpmErr error

func tryVm_allocate_cpm(host_priv Host_priv_t, task uint32, address *Vm_address_t, size Vm_size_t, flags int32) (Kern_return_t, error) {
	if _vm_allocate_cpm == nil {
		return *new(Kern_return_t), symbolCallError("vm_allocate_cpm", "10.0", _vm_allocate_cpmErr)
	}
	return _vm_allocate_cpm(host_priv, task, address, size, flags), nil
}

// Vm_allocate_cpm.
//
// See: https://developer.apple.com/documentation/kernel/1588863-vm_allocate_cpm
func Vm_allocate_cpm(host_priv Host_priv_t, task uint32, address *Vm_address_t, size Vm_size_t, flags int32) Kern_return_t {
	result, callErr := tryVm_allocate_cpm(host_priv, task, address, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_behavior_set func(target_task uint32, address Vm_address_t, size Vm_size_t, new_behavior Vm_behavior_t) Kern_return_t
var _vm_behavior_setErr error

func tryVm_behavior_set(target_task uint32, address Vm_address_t, size Vm_size_t, new_behavior Vm_behavior_t) (Kern_return_t, error) {
	if _vm_behavior_set == nil {
		return *new(Kern_return_t), symbolCallError("vm_behavior_set", "10.0", _vm_behavior_setErr)
	}
	return _vm_behavior_set(target_task, address, size, new_behavior), nil
}

// Vm_behavior_set.
//
// See: https://developer.apple.com/documentation/kernel/1585236-vm_behavior_set
func Vm_behavior_set(target_task uint32, address Vm_address_t, size Vm_size_t, new_behavior Vm_behavior_t) Kern_return_t {
	result, callErr := tryVm_behavior_set(target_task, address, size, new_behavior)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_copy func(target_task uint32, source_address Vm_address_t, size Vm_size_t, dest_address Vm_address_t) Kern_return_t
var _vm_copyErr error

func tryVm_copy(target_task uint32, source_address Vm_address_t, size Vm_size_t, dest_address Vm_address_t) (Kern_return_t, error) {
	if _vm_copy == nil {
		return *new(Kern_return_t), symbolCallError("vm_copy", "10.0", _vm_copyErr)
	}
	return _vm_copy(target_task, source_address, size, dest_address), nil
}

// Vm_copy.
//
// See: https://developer.apple.com/documentation/kernel/1585277-vm_copy
func Vm_copy(target_task uint32, source_address Vm_address_t, size Vm_size_t, dest_address Vm_address_t) Kern_return_t {
	result, callErr := tryVm_copy(target_task, source_address, size, dest_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_deallocate func(target_task uint32, address Vm_address_t, size Vm_size_t) Kern_return_t
var _vm_deallocateErr error

func tryVm_deallocate(target_task uint32, address Vm_address_t, size Vm_size_t) (Kern_return_t, error) {
	if _vm_deallocate == nil {
		return *new(Kern_return_t), symbolCallError("vm_deallocate", "10.0", _vm_deallocateErr)
	}
	return _vm_deallocate(target_task, address, size), nil
}

// Vm_deallocate.
//
// See: https://developer.apple.com/documentation/kernel/1585284-vm_deallocate
func Vm_deallocate(target_task uint32, address Vm_address_t, size Vm_size_t) Kern_return_t {
	result, callErr := tryVm_deallocate(target_task, address, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_inherit func(target_task uint32, address Vm_address_t, size Vm_size_t, new_inheritance Vm_inherit_t) Kern_return_t
var _vm_inheritErr error

func tryVm_inherit(target_task uint32, address Vm_address_t, size Vm_size_t, new_inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _vm_inherit == nil {
		return *new(Kern_return_t), symbolCallError("vm_inherit", "10.0", _vm_inheritErr)
	}
	return _vm_inherit(target_task, address, size, new_inheritance), nil
}

// Vm_inherit.
//
// See: https://developer.apple.com/documentation/kernel/1585275-vm_inherit
func Vm_inherit(target_task uint32, address Vm_address_t, size Vm_size_t, new_inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryVm_inherit(target_task, address, size, new_inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_machine_attribute func(target_task uint32, address Vm_address_t, size Vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) Kern_return_t
var _vm_machine_attributeErr error

func tryVm_machine_attribute(target_task uint32, address Vm_address_t, size Vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) (Kern_return_t, error) {
	if _vm_machine_attribute == nil {
		return *new(Kern_return_t), symbolCallError("vm_machine_attribute", "10.0", _vm_machine_attributeErr)
	}
	return _vm_machine_attribute(target_task, address, size, attribute, value), nil
}

// Vm_machine_attribute.
//
// See: https://developer.apple.com/documentation/kernel/1585354-vm_machine_attribute
func Vm_machine_attribute(target_task uint32, address Vm_address_t, size Vm_size_t, attribute Vm_machine_attribute_t, value *Vm_machine_attribute_val_t) Kern_return_t {
	result, callErr := tryVm_machine_attribute(target_task, address, size, attribute, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_map func(target_task uint32, address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, object Mem_entry_name_port_t, offset Vm_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _vm_mapErr error

func tryVm_map(target_task uint32, address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, object Mem_entry_name_port_t, offset Vm_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _vm_map == nil {
		return *new(Kern_return_t), symbolCallError("vm_map", "10.0", _vm_mapErr)
	}
	return _vm_map(target_task, address, size, mask, flags, object, offset, copy_, cur_protection, max_protection, inheritance), nil
}

// Vm_map.
//
// See: https://developer.apple.com/documentation/kernel/1585510-vm_map
func Vm_map(target_task uint32, address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, object Mem_entry_name_port_t, offset Vm_offset_t, copy_ Boolean_t, cur_protection Vm_prot_t, max_protection Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryVm_map(target_task, address, size, mask, flags, object, offset, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_map_page_query func(target_map Vm_map_read_t, offset Vm_offset_t, disposition *Integer_t, ref_count *Integer_t) Kern_return_t
var _vm_map_page_queryErr error

func tryVm_map_page_query(target_map Vm_map_read_t, offset Vm_offset_t, disposition *Integer_t, ref_count *Integer_t) (Kern_return_t, error) {
	if _vm_map_page_query == nil {
		return *new(Kern_return_t), symbolCallError("vm_map_page_query", "10.0", _vm_map_page_queryErr)
	}
	return _vm_map_page_query(target_map, offset, disposition, ref_count), nil
}

// Vm_map_page_query.
//
// See: https://developer.apple.com/documentation/kernel/1585356-vm_map_page_query
func Vm_map_page_query(target_map Vm_map_read_t, offset Vm_offset_t, disposition *Integer_t, ref_count *Integer_t) Kern_return_t {
	result, callErr := tryVm_map_page_query(target_map, offset, disposition, ref_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_msync func(target_task uint32, address Vm_address_t, size Vm_size_t, sync_flags Vm_sync_t) Kern_return_t
var _vm_msyncErr error

func tryVm_msync(target_task uint32, address Vm_address_t, size Vm_size_t, sync_flags Vm_sync_t) (Kern_return_t, error) {
	if _vm_msync == nil {
		return *new(Kern_return_t), symbolCallError("vm_msync", "10.0", _vm_msyncErr)
	}
	return _vm_msync(target_task, address, size, sync_flags), nil
}

// Vm_msync.
//
// See: https://developer.apple.com/documentation/kernel/1585201-vm_msync
func Vm_msync(target_task uint32, address Vm_address_t, size Vm_size_t, sync_flags Vm_sync_t) Kern_return_t {
	result, callErr := tryVm_msync(target_task, address, size, sync_flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_protect func(target_task uint32, address Vm_address_t, size Vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) Kern_return_t
var _vm_protectErr error

func tryVm_protect(target_task uint32, address Vm_address_t, size Vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) (Kern_return_t, error) {
	if _vm_protect == nil {
		return *new(Kern_return_t), symbolCallError("vm_protect", "10.0", _vm_protectErr)
	}
	return _vm_protect(target_task, address, size, set_maximum, new_protection), nil
}

// Vm_protect.
//
// See: https://developer.apple.com/documentation/kernel/1585294-vm_protect
func Vm_protect(target_task uint32, address Vm_address_t, size Vm_size_t, set_maximum Boolean_t, new_protection Vm_prot_t) Kern_return_t {
	result, callErr := tryVm_protect(target_task, address, size, set_maximum, new_protection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_purgable_control func(target_task uint32, address Vm_address_t, control Vm_purgable_t, state *int32) Kern_return_t
var _vm_purgable_controlErr error

func tryVm_purgable_control(target_task uint32, address Vm_address_t, control Vm_purgable_t, state *int32) (Kern_return_t, error) {
	if _vm_purgable_control == nil {
		return *new(Kern_return_t), symbolCallError("vm_purgable_control", "10.4", _vm_purgable_controlErr)
	}
	return _vm_purgable_control(target_task, address, control, state), nil
}

// Vm_purgable_control.
//
// See: https://developer.apple.com/documentation/kernel/1585267-vm_purgable_control
func Vm_purgable_control(target_task uint32, address Vm_address_t, control Vm_purgable_t, state *int32) Kern_return_t {
	result, callErr := tryVm_purgable_control(target_task, address, control, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_read func(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) Kern_return_t
var _vm_readErr error

func tryVm_read(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _vm_read == nil {
		return *new(Kern_return_t), symbolCallError("vm_read", "10.0", _vm_readErr)
	}
	return _vm_read(target_task, address, size, data, dataCnt), nil
}

// Vm_read.
//
// See: https://developer.apple.com/documentation/kernel/1585350-vm_read
func Vm_read(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data *Vm_offset_t, dataCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryVm_read(target_task, address, size, data, dataCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_read_list func(target_task Vm_map_read_t, data_list Vm_read_entry_t, count Natural_t) Kern_return_t
var _vm_read_listErr error

func tryVm_read_list(target_task Vm_map_read_t, data_list Vm_read_entry_t, count Natural_t) (Kern_return_t, error) {
	if _vm_read_list == nil {
		return *new(Kern_return_t), symbolCallError("vm_read_list", "10.0", _vm_read_listErr)
	}
	return _vm_read_list(target_task, data_list, count), nil
}

// Vm_read_list.
//
// See: https://developer.apple.com/documentation/kernel/1585516-vm_read_list
func Vm_read_list(target_task Vm_map_read_t, data_list Vm_read_entry_t, count Natural_t) Kern_return_t {
	result, callErr := tryVm_read_list(target_task, data_list, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_read_overwrite func(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data Vm_address_t, outsize *Vm_size_t) Kern_return_t
var _vm_read_overwriteErr error

func tryVm_read_overwrite(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data Vm_address_t, outsize *Vm_size_t) (Kern_return_t, error) {
	if _vm_read_overwrite == nil {
		return *new(Kern_return_t), symbolCallError("vm_read_overwrite", "10.0", _vm_read_overwriteErr)
	}
	return _vm_read_overwrite(target_task, address, size, data, outsize), nil
}

// Vm_read_overwrite.
//
// See: https://developer.apple.com/documentation/kernel/1585371-vm_read_overwrite
func Vm_read_overwrite(target_task Vm_map_read_t, address Vm_address_t, size Vm_size_t, data Vm_address_t, outsize *Vm_size_t) Kern_return_t {
	result, callErr := tryVm_read_overwrite(target_task, address, size, data, outsize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_region_64 func(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) Kern_return_t
var _vm_region_64Err error

func tryVm_region_64(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) (Kern_return_t, error) {
	if _vm_region_64 == nil {
		return *new(Kern_return_t), symbolCallError("vm_region_64", "10.0", _vm_region_64Err)
	}
	return _vm_region_64(target_task, address, size, flavor, info, infoCnt, object_name), nil
}

// Vm_region_64.
//
// See: https://developer.apple.com/documentation/kernel/1585386-vm_region_64
func Vm_region_64(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, flavor Vm_region_flavor_t, info Vm_region_info_t, infoCnt *Mach_msg_type_number_t, object_name *uint32) Kern_return_t {
	result, callErr := tryVm_region_64(target_task, address, size, flavor, info, infoCnt, object_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_region_recurse_64 func(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t
var _vm_region_recurse_64Err error

func tryVm_region_recurse_64(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) (Kern_return_t, error) {
	if _vm_region_recurse_64 == nil {
		return *new(Kern_return_t), symbolCallError("vm_region_recurse_64", "10.0", _vm_region_recurse_64Err)
	}
	return _vm_region_recurse_64(target_task, address, size, nesting_depth, info, infoCnt), nil
}

// Vm_region_recurse_64.
//
// See: https://developer.apple.com/documentation/kernel/1585424-vm_region_recurse_64
func Vm_region_recurse_64(target_task Vm_map_read_t, address *Vm_address_t, size *Vm_size_t, nesting_depth *Natural_t, info Vm_region_recurse_info_t, infoCnt *Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryVm_region_recurse_64(target_task, address, size, nesting_depth, info, infoCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_remap func(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task uint32, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _vm_remapErr error

func tryVm_remap(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task uint32, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _vm_remap == nil {
		return *new(Kern_return_t), symbolCallError("vm_remap", "10.0", _vm_remapErr)
	}
	return _vm_remap(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance), nil
}

// Vm_remap.
//
// See: https://developer.apple.com/documentation/kernel/1585336-vm_remap
func Vm_remap(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task uint32, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryVm_remap(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_remap_new func(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task Vm_map_read_t, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t
var _vm_remap_newErr error

func tryVm_remap_new(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task Vm_map_read_t, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) (Kern_return_t, error) {
	if _vm_remap_new == nil {
		return *new(Kern_return_t), symbolCallError("vm_remap_new", "11.3", _vm_remap_newErr)
	}
	return _vm_remap_new(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance), nil
}

// Vm_remap_new.
//
// See: https://developer.apple.com/documentation/kernel/3728001-vm_remap_new
func Vm_remap_new(target_task uint32, target_address *Vm_address_t, size Vm_size_t, mask Vm_address_t, flags int32, src_task Vm_map_read_t, src_address Vm_address_t, copy_ Boolean_t, cur_protection *Vm_prot_t, max_protection *Vm_prot_t, inheritance Vm_inherit_t) Kern_return_t {
	result, callErr := tryVm_remap_new(target_task, target_address, size, mask, flags, src_task, src_address, copy_, cur_protection, max_protection, inheritance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_wire func(host_priv Host_priv_t, task uint32, address Vm_address_t, size Vm_size_t, desired_access Vm_prot_t) Kern_return_t
var _vm_wireErr error

func tryVm_wire(host_priv Host_priv_t, task uint32, address Vm_address_t, size Vm_size_t, desired_access Vm_prot_t) (Kern_return_t, error) {
	if _vm_wire == nil {
		return *new(Kern_return_t), symbolCallError("vm_wire", "10.0", _vm_wireErr)
	}
	return _vm_wire(host_priv, task, address, size, desired_access), nil
}

// Vm_wire.
//
// See: https://developer.apple.com/documentation/kernel/1588985-vm_wire
func Vm_wire(host_priv Host_priv_t, task uint32, address Vm_address_t, size Vm_size_t, desired_access Vm_prot_t) Kern_return_t {
	result, callErr := tryVm_wire(host_priv, task, address, size, desired_access)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vm_write func(target_task uint32, address Vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) Kern_return_t
var _vm_writeErr error

func tryVm_write(target_task uint32, address Vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) (Kern_return_t, error) {
	if _vm_write == nil {
		return *new(Kern_return_t), symbolCallError("vm_write", "10.0", _vm_writeErr)
	}
	return _vm_write(target_task, address, data, dataCnt), nil
}

// Vm_write.
//
// See: https://developer.apple.com/documentation/kernel/1585462-vm_write
func Vm_write(target_task uint32, address Vm_address_t, data Vm_offset_t, dataCnt Mach_msg_type_number_t) Kern_return_t {
	result, callErr := tryVm_write(target_task, address, data, dataCnt)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vprintf func(arg0 string, arg1 uintptr) int32
var _vprintfErr error

func tryVprintf(arg0 string, arg1 uintptr) (int32, error) {
	if _vprintf == nil {
		return 0, symbolCallError("vprintf", "10.5", _vprintfErr)
	}
	return _vprintf(arg0, arg1), nil
}

// Vprintf.
//
// See: https://developer.apple.com/documentation/kernel/1441075-vprintf
func Vprintf(arg0 string, arg1 uintptr) int32 {
	result, callErr := tryVprintf(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vsnprintf func(arg0 string, arg1 uintptr, arg2 string, arg3 uintptr) int32
var _vsnprintfErr error

func tryVsnprintf(arg0 string, arg1 uintptr, arg2 string, arg3 uintptr) (int32, error) {
	if _vsnprintf == nil {
		return 0, symbolCallError("vsnprintf", "10.0", _vsnprintfErr)
	}
	return _vsnprintf(arg0, arg1, arg2, arg3), nil
}

// Vsnprintf.
//
// See: https://developer.apple.com/documentation/kernel/1441056-vsnprintf
func Vsnprintf(arg0 string, arg1 uintptr, arg2 string, arg3 uintptr) int32 {
	result, callErr := tryVsnprintf(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vsprintf func(bufp string, arg1 string, arg2 uintptr) int32
var _vsprintfErr error

func tryVsprintf(bufp string, arg1 string, arg2 uintptr) (int32, error) {
	if _vsprintf == nil {
		return 0, symbolCallError("vsprintf", "10.0", _vsprintfErr)
	}
	return _vsprintf(bufp, arg1, arg2), nil
}

// Vsprintf.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/kernel/1441062-vsprintf
func Vsprintf(bufp string, arg1 string, arg2 uintptr) int32 {
	result, callErr := tryVsprintf(bufp, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vsscanf func(arg0 string, arg1 string, arg2 uintptr) int32
var _vsscanfErr error

func tryVsscanf(arg0 string, arg1 string, arg2 uintptr) (int32, error) {
	if _vsscanf == nil {
		return 0, symbolCallError("vsscanf", "10.4", _vsscanfErr)
	}
	return _vsscanf(arg0, arg1, arg2), nil
}

// Vsscanf.
//
// See: https://developer.apple.com/documentation/kernel/1441086-vsscanf
func Vsscanf(arg0 string, arg1 string, arg2 uintptr) int32 {
	result, callErr := tryVsscanf(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _y0 func(arg0 float64) float64
var _y0Err error

func tryY0(arg0 float64) (float64, error) {
	if _y0 == nil {
		return 0.0, symbolCallError("y0", "10.0", _y0Err)
	}
	return _y0(arg0), nil
}

// Y0.
//
// See: https://developer.apple.com/documentation/kernel/1557220-y0
func Y0(arg0 float64) float64 {
	result, callErr := tryY0(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _y1 func(arg0 float64) float64
var _y1Err error

func tryY1(arg0 float64) (float64, error) {
	if _y1 == nil {
		return 0.0, symbolCallError("y1", "10.0", _y1Err)
	}
	return _y1(arg0), nil
}

// Y1.
//
// See: https://developer.apple.com/documentation/kernel/1557331-y1
func Y1(arg0 float64) float64 {
	result, callErr := tryY1(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _yn func(arg0 int32, arg1 float64) float64
var _ynErr error

func tryYn(arg0 int32, arg1 float64) (float64, error) {
	if _yn == nil {
		return 0.0, symbolCallError("yn", "10.0", _ynErr)
	}
	return _yn(arg0, arg1), nil
}

// Yn.
//
// See: https://developer.apple.com/documentation/kernel/1557339-yn
func Yn(arg0 int32, arg1 float64) float64 {
	result, callErr := tryYn(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_mD5Final, &_mD5FinalErr, frameworkHandle, "MD5Final", "10.0")
	registerFunc(&_acos, &_acosErr, frameworkHandle, "acos", "10.10")
	registerFunc(&_acosf, &_acosfErr, frameworkHandle, "acosf", "10.10")
	registerFunc(&_acosh, &_acoshErr, frameworkHandle, "acosh", "10.10")
	registerFunc(&_acoshf, &_acoshfErr, frameworkHandle, "acoshf", "10.10")
	registerFunc(&_acoshl, &_acoshlErr, frameworkHandle, "acoshl", "10.10")
	registerFunc(&_acosl, &_acoslErr, frameworkHandle, "acosl", "10.10")
	registerFunc(&_act_get_state, &_act_get_stateErr, frameworkHandle, "act_get_state", "10.0")
	registerFunc(&_act_set_state, &_act_set_stateErr, frameworkHandle, "act_set_state", "10.0")
	registerFunc(&_asin, &_asinErr, frameworkHandle, "asin", "10.10")
	registerFunc(&_asinf, &_asinfErr, frameworkHandle, "asinf", "10.10")
	registerFunc(&_asinh, &_asinhErr, frameworkHandle, "asinh", "10.10")
	registerFunc(&_asinhf, &_asinhfErr, frameworkHandle, "asinhf", "10.10")
	registerFunc(&_asinhl, &_asinhlErr, frameworkHandle, "asinhl", "10.10")
	registerFunc(&_asinl, &_asinlErr, frameworkHandle, "asinl", "10.10")
	registerFunc(&_atan, &_atanErr, frameworkHandle, "atan", "10.10")
	registerFunc(&_atan2, &_atan2Err, frameworkHandle, "atan2", "10.10")
	registerFunc(&_atan2f, &_atan2fErr, frameworkHandle, "atan2f", "10.10")
	registerFunc(&_atan2l, &_atan2lErr, frameworkHandle, "atan2l", "10.10")
	registerFunc(&_atanf, &_atanfErr, frameworkHandle, "atanf", "10.10")
	registerFunc(&_atanh, &_atanhErr, frameworkHandle, "atanh", "10.10")
	registerFunc(&_atanhf, &_atanhfErr, frameworkHandle, "atanhf", "10.10")
	registerFunc(&_atanhl, &_atanhlErr, frameworkHandle, "atanhl", "10.10")
	registerFunc(&_atanl, &_atanlErr, frameworkHandle, "atanl", "10.10")
	registerFunc(&_backtrace, &_backtraceErr, frameworkHandle, "backtrace", "10.12")
	registerFunc(&_bcmp, &_bcmpErr, frameworkHandle, "bcmp", "10.0")
	registerFunc(&_bcopy, &_bcopyErr, frameworkHandle, "bcopy", "10.0")
	registerFunc(&_bzero, &_bzeroErr, frameworkHandle, "bzero", "10.0")
	registerFunc(&_cbrt, &_cbrtErr, frameworkHandle, "cbrt", "10.10")
	registerFunc(&_cbrtf, &_cbrtfErr, frameworkHandle, "cbrtf", "10.10")
	registerFunc(&_cbrtl, &_cbrtlErr, frameworkHandle, "cbrtl", "10.10")
	registerFunc(&_ceil, &_ceilErr, frameworkHandle, "ceil", "10.10")
	registerFunc(&_ceilf, &_ceilfErr, frameworkHandle, "ceilf", "10.10")
	registerFunc(&_ceill, &_ceillErr, frameworkHandle, "ceill", "10.10")
	registerFunc(&_clock_alarm, &_clock_alarmErr, frameworkHandle, "clock_alarm", "10.0")
	registerFunc(&_clock_alarm_reply, &_clock_alarm_replyErr, frameworkHandle, "clock_alarm_reply", "10.0")
	registerFunc(&_clock_get_attributes, &_clock_get_attributesErr, frameworkHandle, "clock_get_attributes", "10.0")
	registerFunc(&_clock_get_time, &_clock_get_timeErr, frameworkHandle, "clock_get_time", "10.0")
	registerFunc(&_clock_set_attributes, &_clock_set_attributesErr, frameworkHandle, "clock_set_attributes", "10.0")
	registerFunc(&_clock_set_time, &_clock_set_timeErr, frameworkHandle, "clock_set_time", "10.0")
	registerFunc(&_copysign, &_copysignErr, frameworkHandle, "copysign", "10.10")
	registerFunc(&_copysignf, &_copysignfErr, frameworkHandle, "copysignf", "10.10")
	registerFunc(&_copysignl, &_copysignlErr, frameworkHandle, "copysignl", "10.10")
	registerFunc(&_cos, &_cosErr, frameworkHandle, "cos", "10.10")
	registerFunc(&_cosf, &_cosfErr, frameworkHandle, "cosf", "10.10")
	registerFunc(&_cosh, &_coshErr, frameworkHandle, "cosh", "10.10")
	registerFunc(&_coshf, &_coshfErr, frameworkHandle, "coshf", "10.10")
	registerFunc(&_coshl, &_coshlErr, frameworkHandle, "coshl", "10.10")
	registerFunc(&_cosl, &_coslErr, frameworkHandle, "cosl", "10.10")
	registerFunc(&_erf, &_erfErr, frameworkHandle, "erf", "10.10")
	registerFunc(&_erfc, &_erfcErr, frameworkHandle, "erfc", "10.10")
	registerFunc(&_erfcf, &_erfcfErr, frameworkHandle, "erfcf", "10.10")
	registerFunc(&_erfcl, &_erfclErr, frameworkHandle, "erfcl", "10.10")
	registerFunc(&_erff, &_erffErr, frameworkHandle, "erff", "10.10")
	registerFunc(&_erfl, &_erflErr, frameworkHandle, "erfl", "10.10")
	registerFunc(&_etap_trace_thread, &_etap_trace_threadErr, frameworkHandle, "etap_trace_thread", "10.0")
	registerFunc(&_exc_server, &_exc_serverErr, frameworkHandle, "exc_server", "10.0")
	registerFunc(&_exc_server_routine, &_exc_server_routineErr, frameworkHandle, "exc_server_routine", "10.0")
	registerFunc(&_exp, &_expErr, frameworkHandle, "exp", "10.10")
	registerFunc(&_exp2, &_exp2Err, frameworkHandle, "exp2", "10.10")
	registerFunc(&_exp2f, &_exp2fErr, frameworkHandle, "exp2f", "10.10")
	registerFunc(&_exp2l, &_exp2lErr, frameworkHandle, "exp2l", "10.10")
	registerFunc(&_expf, &_expfErr, frameworkHandle, "expf", "10.9")
	registerFunc(&_expl, &_explErr, frameworkHandle, "expl", "10.10")
	registerFunc(&_expm1, &_expm1Err, frameworkHandle, "expm1", "10.10")
	registerFunc(&_expm1f, &_expm1fErr, frameworkHandle, "expm1f", "10.10")
	registerFunc(&_expm1l, &_expm1lErr, frameworkHandle, "expm1l", "10.10")
	registerFunc(&_fabs, &_fabsErr, frameworkHandle, "fabs", "10.10")
	registerFunc(&_fabsf, &_fabsfErr, frameworkHandle, "fabsf", "10.10")
	registerFunc(&_fabsl, &_fabslErr, frameworkHandle, "fabsl", "10.10")
	registerFunc(&_fdim, &_fdimErr, frameworkHandle, "fdim", "10.10")
	registerFunc(&_fdimf, &_fdimfErr, frameworkHandle, "fdimf", "10.10")
	registerFunc(&_fdiml, &_fdimlErr, frameworkHandle, "fdiml", "10.10")
	registerFunc(&_ffs, &_ffsErr, frameworkHandle, "ffs", "10.0")
	registerFunc(&_ffsll, &_ffsllErr, frameworkHandle, "ffsll", "10.13")
	registerFunc(&_floor, &_floorErr, frameworkHandle, "floor", "10.10")
	registerFunc(&_floorf, &_floorfErr, frameworkHandle, "floorf", "10.10")
	registerFunc(&_floorl, &_floorlErr, frameworkHandle, "floorl", "10.10")
	registerFunc(&_fls, &_flsErr, frameworkHandle, "fls", "10.13")
	registerFunc(&_flsll, &_flsllErr, frameworkHandle, "flsll", "10.13")
	registerFunc(&_fma, &_fmaErr, frameworkHandle, "fma", "10.10")
	registerFunc(&_fmaf, &_fmafErr, frameworkHandle, "fmaf", "10.10")
	registerFunc(&_fmal, &_fmalErr, frameworkHandle, "fmal", "10.10")
	registerFunc(&_fmax, &_fmaxErr, frameworkHandle, "fmax", "10.10")
	registerFunc(&_fmaxf, &_fmaxfErr, frameworkHandle, "fmaxf", "10.10")
	registerFunc(&_fmaxl, &_fmaxlErr, frameworkHandle, "fmaxl", "10.10")
	registerFunc(&_fmin, &_fminErr, frameworkHandle, "fmin", "10.10")
	registerFunc(&_fminf, &_fminfErr, frameworkHandle, "fminf", "10.10")
	registerFunc(&_fminl, &_fminlErr, frameworkHandle, "fminl", "10.10")
	registerFunc(&_fmod, &_fmodErr, frameworkHandle, "fmod", "10.10")
	registerFunc(&_fmodf, &_fmodfErr, frameworkHandle, "fmodf", "10.10")
	registerFunc(&_fmodl, &_fmodlErr, frameworkHandle, "fmodl", "10.10")
	registerFunc(&_frexp, &_frexpErr, frameworkHandle, "frexp", "10.10")
	registerFunc(&_frexpf, &_frexpfErr, frameworkHandle, "frexpf", "10.10")
	registerFunc(&_frexpl, &_frexplErr, frameworkHandle, "frexpl", "10.10")
	registerFunc(&_host_create_mach_voucher, &_host_create_mach_voucherErr, frameworkHandle, "host_create_mach_voucher", "10.10")
	registerFunc(&_host_default_memory_manager, &_host_default_memory_managerErr, frameworkHandle, "host_default_memory_manager", "10.0")
	registerFunc(&_host_get_UNDServer, &_host_get_UNDServerErr, frameworkHandle, "host_get_UNDServer", "10.0")
	registerFunc(&_host_get_boot_info, &_host_get_boot_infoErr, frameworkHandle, "host_get_boot_info", "10.0")
	registerFunc(&_host_get_clock_control, &_host_get_clock_controlErr, frameworkHandle, "host_get_clock_control", "10.0")
	registerFunc(&_host_get_clock_service, &_host_get_clock_serviceErr, frameworkHandle, "host_get_clock_service", "10.0")
	registerFunc(&_host_get_exception_ports, &_host_get_exception_portsErr, frameworkHandle, "host_get_exception_ports", "10.0")
	registerFunc(&_host_get_io_main, &_host_get_io_mainErr, frameworkHandle, "host_get_io_main", "13.0")
	registerFunc(&_host_get_special_port, &_host_get_special_portErr, frameworkHandle, "host_get_special_port", "10.0")
	registerFunc(&_host_info, &_host_infoErr, frameworkHandle, "host_info", "10.0")
	registerFunc(&_host_kernel_version, &_host_kernel_versionErr, frameworkHandle, "host_kernel_version", "10.0")
	registerFunc(&_host_lockgroup_info, &_host_lockgroup_infoErr, frameworkHandle, "host_lockgroup_info", "10.4")
	registerFunc(&_host_page_size, &_host_page_sizeErr, frameworkHandle, "host_page_size", "10.0")
	registerFunc(&_host_priv_statistics, &_host_priv_statisticsErr, frameworkHandle, "host_priv_statistics", "10.0")
	registerFunc(&_host_processor_info, &_host_processor_infoErr, frameworkHandle, "host_processor_info", "10.0")
	registerFunc(&_host_processor_set_priv, &_host_processor_set_privErr, frameworkHandle, "host_processor_set_priv", "10.0")
	registerFunc(&_host_processor_sets, &_host_processor_setsErr, frameworkHandle, "host_processor_sets", "10.0")
	registerFunc(&_host_processors, &_host_processorsErr, frameworkHandle, "host_processors", "10.0")
	registerFunc(&_host_reboot, &_host_rebootErr, frameworkHandle, "host_reboot", "10.0")
	registerFunc(&_host_register_mach_voucher_attr_manager, &_host_register_mach_voucher_attr_managerErr, frameworkHandle, "host_register_mach_voucher_attr_manager", "10.10")
	registerFunc(&_host_register_well_known_mach_voucher_attr_manager, &_host_register_well_known_mach_voucher_attr_managerErr, frameworkHandle, "host_register_well_known_mach_voucher_attr_manager", "10.10")
	registerFunc(&_host_request_notification, &_host_request_notificationErr, frameworkHandle, "host_request_notification", "10.3")
	registerFunc(&_host_security_create_task_token, &_host_security_create_task_tokenErr, frameworkHandle, "host_security_create_task_token", "10.0")
	registerFunc(&_host_security_set_task_token, &_host_security_set_task_tokenErr, frameworkHandle, "host_security_set_task_token", "10.0")
	registerFunc(&_host_self, &_host_selfErr, frameworkHandle, "host_self", "10.0")
	registerFunc(&_host_set_UNDServer, &_host_set_UNDServerErr, frameworkHandle, "host_set_UNDServer", "10.0")
	registerFunc(&_host_set_atm_diagnostic_flag, &_host_set_atm_diagnostic_flagErr, frameworkHandle, "host_set_atm_diagnostic_flag", "10.11")
	registerFunc(&_host_set_exception_ports, &_host_set_exception_portsErr, frameworkHandle, "host_set_exception_ports", "10.0")
	registerFunc(&_host_set_multiuser_config_flags, &_host_set_multiuser_config_flagsErr, frameworkHandle, "host_set_multiuser_config_flags", "10.11.4")
	registerFunc(&_host_set_special_port, &_host_set_special_portErr, frameworkHandle, "host_set_special_port", "10.0")
	registerFunc(&_host_statistics, &_host_statisticsErr, frameworkHandle, "host_statistics", "10.0")
	registerFunc(&_host_statistics64, &_host_statistics64Err, frameworkHandle, "host_statistics64", "10.6")
	registerFunc(&_host_swap_exception_ports, &_host_swap_exception_portsErr, frameworkHandle, "host_swap_exception_ports", "10.0")
	registerFunc(&_host_virtual_physical_table_info, &_host_virtual_physical_table_infoErr, frameworkHandle, "host_virtual_physical_table_info", "10.0")
	registerFunc(&_hypot, &_hypotErr, frameworkHandle, "hypot", "10.10")
	registerFunc(&_hypotf, &_hypotfErr, frameworkHandle, "hypotf", "10.10")
	registerFunc(&_hypotl, &_hypotlErr, frameworkHandle, "hypotl", "10.10")
	registerFunc(&_ilogb, &_ilogbErr, frameworkHandle, "ilogb", "10.10")
	registerFunc(&_ilogbf, &_ilogbfErr, frameworkHandle, "ilogbf", "10.10")
	registerFunc(&_ilogbl, &_ilogblErr, frameworkHandle, "ilogbl", "10.10")
	registerFunc(&_inet_aton, &_inet_atonErr, frameworkHandle, "inet_aton", "10.9")
	registerFunc(&_inet_ntop, &_inet_ntopErr, frameworkHandle, "inet_ntop", "10.4")
	registerFunc(&_insque, &_insqueErr, frameworkHandle, "insque", "10.0")
	registerFunc(&_j0, &_j0Err, frameworkHandle, "j0", "10.0")
	registerFunc(&_j1, &_j1Err, frameworkHandle, "j1", "10.0")
	registerFunc(&_jn, &_jnErr, frameworkHandle, "jn", "10.0")
	registerFunc(&_kdebug_timestamp, &_kdebug_timestampErr, frameworkHandle, "kdebug_timestamp", "12.0")
	registerFunc(&_kdebug_timestamp_from_absolute, &_kdebug_timestamp_from_absoluteErr, frameworkHandle, "kdebug_timestamp_from_absolute", "12.0")
	registerFunc(&_kdebug_timestamp_from_continuous, &_kdebug_timestamp_from_continuousErr, frameworkHandle, "kdebug_timestamp_from_continuous", "12.0")
	registerFunc(&_kdebug_using_continuous_time, &_kdebug_using_continuous_timeErr, frameworkHandle, "kdebug_using_continuous_time", "10.15")
	registerFunc(&_kevent64, &_kevent64Err, frameworkHandle, "kevent64", "")
	registerFunc(&_kext_request, &_kext_requestErr, frameworkHandle, "kext_request", "10.6")
	registerFunc(&_kmod_control, &_kmod_controlErr, frameworkHandle, "kmod_control", "10.0")
	registerFunc(&_kmod_create, &_kmod_createErr, frameworkHandle, "kmod_create", "10.0")
	registerFunc(&_kmod_destroy, &_kmod_destroyErr, frameworkHandle, "kmod_destroy", "10.0")
	registerFunc(&_kmod_get_info, &_kmod_get_infoErr, frameworkHandle, "kmod_get_info", "10.0")
	registerFunc(&_kqueue, &_kqueueErr, frameworkHandle, "kqueue", "")
	registerFunc(&_ldexp, &_ldexpErr, frameworkHandle, "ldexp", "10.10")
	registerFunc(&_ldexpf, &_ldexpfErr, frameworkHandle, "ldexpf", "10.10")
	registerFunc(&_ldexpl, &_ldexplErr, frameworkHandle, "ldexpl", "10.10")
	registerFunc(&_lgamma, &_lgammaErr, frameworkHandle, "lgamma", "10.10")
	registerFunc(&_lgammaf, &_lgammafErr, frameworkHandle, "lgammaf", "10.10")
	registerFunc(&_lgammal, &_lgammalErr, frameworkHandle, "lgammal", "10.10")
	registerFunc(&_llrint, &_llrintErr, frameworkHandle, "llrint", "10.10")
	registerFunc(&_llrintf, &_llrintfErr, frameworkHandle, "llrintf", "10.10")
	registerFunc(&_llrintl, &_llrintlErr, frameworkHandle, "llrintl", "10.10")
	registerFunc(&_llround, &_llroundErr, frameworkHandle, "llround", "10.10")
	registerFunc(&_llroundf, &_llroundfErr, frameworkHandle, "llroundf", "10.10")
	registerFunc(&_llroundl, &_llroundlErr, frameworkHandle, "llroundl", "10.10")
	registerFunc(&_lock_set_create, &_lock_set_createErr, frameworkHandle, "lock_set_create", "10.0")
	registerFunc(&_lock_set_destroy, &_lock_set_destroyErr, frameworkHandle, "lock_set_destroy", "10.0")
	registerFunc(&_log, &_logErr, frameworkHandle, "log", "10.0")
	registerFunc(&_log10, &_log10Err, frameworkHandle, "log10", "10.10")
	registerFunc(&_log10f, &_log10fErr, frameworkHandle, "log10f", "10.9")
	registerFunc(&_log10l, &_log10lErr, frameworkHandle, "log10l", "10.10")
	registerFunc(&_log1p, &_log1pErr, frameworkHandle, "log1p", "10.10")
	registerFunc(&_log1pf, &_log1pfErr, frameworkHandle, "log1pf", "10.10")
	registerFunc(&_log1pl, &_log1plErr, frameworkHandle, "log1pl", "10.10")
	registerFunc(&_log2, &_log2Err, frameworkHandle, "log2", "10.10")
	registerFunc(&_log2f, &_log2fErr, frameworkHandle, "log2f", "10.10")
	registerFunc(&_log2l, &_log2lErr, frameworkHandle, "log2l", "10.10")
	registerFunc(&_logb, &_logbErr, frameworkHandle, "logb", "10.10")
	registerFunc(&_logbf, &_logbfErr, frameworkHandle, "logbf", "10.10")
	registerFunc(&_logbl, &_logblErr, frameworkHandle, "logbl", "10.10")
	registerFunc(&_logf, &_logfErr, frameworkHandle, "logf", "10.9")
	registerFunc(&_logl, &_loglErr, frameworkHandle, "logl", "10.10")
	registerFunc(&_lrint, &_lrintErr, frameworkHandle, "lrint", "10.10")
	registerFunc(&_lrintf, &_lrintfErr, frameworkHandle, "lrintf", "10.10")
	registerFunc(&_lrintl, &_lrintlErr, frameworkHandle, "lrintl", "10.10")
	registerFunc(&_lround, &_lroundErr, frameworkHandle, "lround", "10.10")
	registerFunc(&_lroundf, &_lroundfErr, frameworkHandle, "lroundf", "10.10")
	registerFunc(&_lroundl, &_lroundlErr, frameworkHandle, "lroundl", "10.10")
	registerFunc(&_mach_absolute_time, &_mach_absolute_timeErr, frameworkHandle, "mach_absolute_time", "10.0")
	registerFunc(&_mach_approximate_time, &_mach_approximate_timeErr, frameworkHandle, "mach_approximate_time", "10.10")
	registerFunc(&_mach_continuous_approximate_time, &_mach_continuous_approximate_timeErr, frameworkHandle, "mach_continuous_approximate_time", "10.12")
	registerFunc(&_mach_continuous_time, &_mach_continuous_timeErr, frameworkHandle, "mach_continuous_time", "10.12")
	registerFunc(&_mach_make_memory_entry, &_mach_make_memory_entryErr, frameworkHandle, "mach_make_memory_entry", "10.0")
	registerFunc(&_mach_make_memory_entry_64, &_mach_make_memory_entry_64Err, frameworkHandle, "mach_make_memory_entry_64", "10.0")
	registerFunc(&_mach_memory_entry_access_tracking, &_mach_memory_entry_access_trackingErr, frameworkHandle, "mach_memory_entry_access_tracking", "10.14")
	registerFunc(&_mach_memory_entry_ownership, &_mach_memory_entry_ownershipErr, frameworkHandle, "mach_memory_entry_ownership", "10.15")
	registerFunc(&_mach_memory_entry_purgable_control, &_mach_memory_entry_purgable_controlErr, frameworkHandle, "mach_memory_entry_purgable_control", "10.14")
	registerFunc(&_mach_memory_info, &_mach_memory_infoErr, frameworkHandle, "mach_memory_info", "10.11")
	registerFunc(&_mach_memory_object_memory_entry, &_mach_memory_object_memory_entryErr, frameworkHandle, "mach_memory_object_memory_entry", "10.0")
	registerFunc(&_mach_memory_object_memory_entry_64, &_mach_memory_object_memory_entry_64Err, frameworkHandle, "mach_memory_object_memory_entry_64", "10.0")
	registerFunc(&_mach_msg, &_mach_msgErr, frameworkHandle, "mach_msg", "")
	registerFunc(&_mach_msg_overwrite, &_mach_msg_overwriteErr, frameworkHandle, "mach_msg_overwrite", "10.0")
	registerFunc(&_mach_port_allocate, &_mach_port_allocateErr, frameworkHandle, "mach_port_allocate", "10.0")
	registerFunc(&_mach_port_allocate_full, &_mach_port_allocate_fullErr, frameworkHandle, "mach_port_allocate_full", "10.0")
	registerFunc(&_mach_port_allocate_name, &_mach_port_allocate_nameErr, frameworkHandle, "mach_port_allocate_name", "10.0")
	registerFunc(&_mach_port_allocate_qos, &_mach_port_allocate_qosErr, frameworkHandle, "mach_port_allocate_qos", "10.0")
	registerFunc(&_mach_port_assert_attributes, &_mach_port_assert_attributesErr, frameworkHandle, "mach_port_assert_attributes", "12.0")
	registerFunc(&_mach_port_construct, &_mach_port_constructErr, frameworkHandle, "mach_port_construct", "10.9")
	registerFunc(&_mach_port_deallocate, &_mach_port_deallocateErr, frameworkHandle, "mach_port_deallocate", "10.0")
	registerFunc(&_mach_port_destroy, &_mach_port_destroyErr, frameworkHandle, "mach_port_destroy", "10.0")
	registerFunc(&_mach_port_destruct, &_mach_port_destructErr, frameworkHandle, "mach_port_destruct", "10.9")
	registerFunc(&_mach_port_dnrequest_info, &_mach_port_dnrequest_infoErr, frameworkHandle, "mach_port_dnrequest_info", "10.0")
	registerFunc(&_mach_port_extract_member, &_mach_port_extract_memberErr, frameworkHandle, "mach_port_extract_member", "10.0")
	registerFunc(&_mach_port_extract_right, &_mach_port_extract_rightErr, frameworkHandle, "mach_port_extract_right", "10.0")
	registerFunc(&_mach_port_get_attributes, &_mach_port_get_attributesErr, frameworkHandle, "mach_port_get_attributes", "10.0")
	registerFunc(&_mach_port_get_context, &_mach_port_get_contextErr, frameworkHandle, "mach_port_get_context", "10.6")
	registerFunc(&_mach_port_get_refs, &_mach_port_get_refsErr, frameworkHandle, "mach_port_get_refs", "10.0")
	registerFunc(&_mach_port_get_service_port_info, &_mach_port_get_service_port_infoErr, frameworkHandle, "mach_port_get_service_port_info", "12.0")
	registerFunc(&_mach_port_get_set_status, &_mach_port_get_set_statusErr, frameworkHandle, "mach_port_get_set_status", "10.0")
	registerFunc(&_mach_port_get_srights, &_mach_port_get_srightsErr, frameworkHandle, "mach_port_get_srights", "10.0")
	registerFunc(&_mach_port_guard, &_mach_port_guardErr, frameworkHandle, "mach_port_guard", "10.9")
	registerFunc(&_mach_port_guard_with_flags, &_mach_port_guard_with_flagsErr, frameworkHandle, "mach_port_guard_with_flags", "10.15")
	registerFunc(&_mach_port_insert_member, &_mach_port_insert_memberErr, frameworkHandle, "mach_port_insert_member", "10.0")
	registerFunc(&_mach_port_insert_right, &_mach_port_insert_rightErr, frameworkHandle, "mach_port_insert_right", "10.0")
	registerFunc(&_mach_port_is_connection_for_service, &_mach_port_is_connection_for_serviceErr, frameworkHandle, "mach_port_is_connection_for_service", "12.0")
	registerFunc(&_mach_port_kernel_object, &_mach_port_kernel_objectErr, frameworkHandle, "mach_port_kernel_object", "10.0")
	registerFunc(&_mach_port_kobject, &_mach_port_kobjectErr, frameworkHandle, "mach_port_kobject", "10.6")
	registerFunc(&_mach_port_kobject_description, &_mach_port_kobject_descriptionErr, frameworkHandle, "mach_port_kobject_description", "10.15.4")
	registerFunc(&_mach_port_mod_refs, &_mach_port_mod_refsErr, frameworkHandle, "mach_port_mod_refs", "10.0")
	registerFunc(&_mach_port_move_member, &_mach_port_move_memberErr, frameworkHandle, "mach_port_move_member", "10.0")
	registerFunc(&_mach_port_names, &_mach_port_namesErr, frameworkHandle, "mach_port_names", "10.0")
	registerFunc(&_mach_port_peek, &_mach_port_peekErr, frameworkHandle, "mach_port_peek", "10.9")
	registerFunc(&_mach_port_rename, &_mach_port_renameErr, frameworkHandle, "mach_port_rename", "10.0")
	registerFunc(&_mach_port_request_notification, &_mach_port_request_notificationErr, frameworkHandle, "mach_port_request_notification", "10.0")
	registerFunc(&_mach_port_set_attributes, &_mach_port_set_attributesErr, frameworkHandle, "mach_port_set_attributes", "10.0")
	registerFunc(&_mach_port_set_context, &_mach_port_set_contextErr, frameworkHandle, "mach_port_set_context", "10.6")
	registerFunc(&_mach_port_set_mscount, &_mach_port_set_mscountErr, frameworkHandle, "mach_port_set_mscount", "10.0")
	registerFunc(&_mach_port_set_seqno, &_mach_port_set_seqnoErr, frameworkHandle, "mach_port_set_seqno", "10.0")
	registerFunc(&_mach_port_space_basic_info, &_mach_port_space_basic_infoErr, frameworkHandle, "mach_port_space_basic_info", "10.10")
	registerFunc(&_mach_port_space_info, &_mach_port_space_infoErr, frameworkHandle, "mach_port_space_info", "10.0")
	registerFunc(&_mach_port_swap_guard, &_mach_port_swap_guardErr, frameworkHandle, "mach_port_swap_guard", "10.15")
	registerFunc(&_mach_port_type, &_mach_port_typeErr, frameworkHandle, "mach_port_type", "10.0")
	registerFunc(&_mach_port_unguard, &_mach_port_unguardErr, frameworkHandle, "mach_port_unguard", "10.9")
	registerFunc(&_mach_task_is_self, &_mach_task_is_selfErr, frameworkHandle, "mach_task_is_self", "11.3")
	registerFunc(&_mach_task_self, &_mach_task_selfErr, frameworkHandle, "mach_task_self", "")
	registerFunc(&_mach_vm_allocate, &_mach_vm_allocateErr, frameworkHandle, "mach_vm_allocate", "10.4")
	registerFunc(&_mach_vm_behavior_set, &_mach_vm_behavior_setErr, frameworkHandle, "mach_vm_behavior_set", "10.4")
	registerFunc(&_mach_vm_copy, &_mach_vm_copyErr, frameworkHandle, "mach_vm_copy", "10.4")
	registerFunc(&_mach_vm_deallocate, &_mach_vm_deallocateErr, frameworkHandle, "mach_vm_deallocate", "10.4")
	registerFunc(&_mach_vm_deferred_reclamation_buffer_allocate, &_mach_vm_deferred_reclamation_buffer_allocateErr, frameworkHandle, "mach_vm_deferred_reclamation_buffer_allocate", "15.4")
	registerFunc(&_mach_vm_deferred_reclamation_buffer_flush, &_mach_vm_deferred_reclamation_buffer_flushErr, frameworkHandle, "mach_vm_deferred_reclamation_buffer_flush", "15.4")
	registerFunc(&_mach_vm_deferred_reclamation_buffer_resize, &_mach_vm_deferred_reclamation_buffer_resizeErr, frameworkHandle, "mach_vm_deferred_reclamation_buffer_resize", "15.4")
	registerFunc(&_mach_vm_inherit, &_mach_vm_inheritErr, frameworkHandle, "mach_vm_inherit", "10.4")
	registerFunc(&_mach_vm_machine_attribute, &_mach_vm_machine_attributeErr, frameworkHandle, "mach_vm_machine_attribute", "10.4")
	registerFunc(&_mach_vm_map, &_mach_vm_mapErr, frameworkHandle, "mach_vm_map", "10.4")
	registerFunc(&_mach_vm_msync, &_mach_vm_msyncErr, frameworkHandle, "mach_vm_msync", "10.4")
	registerFunc(&_mach_vm_page_info, &_mach_vm_page_infoErr, frameworkHandle, "mach_vm_page_info", "10.6")
	registerFunc(&_mach_vm_page_query, &_mach_vm_page_queryErr, frameworkHandle, "mach_vm_page_query", "10.4")
	registerFunc(&_mach_vm_page_range_query, &_mach_vm_page_range_queryErr, frameworkHandle, "mach_vm_page_range_query", "10.13")
	registerFunc(&_mach_vm_protect, &_mach_vm_protectErr, frameworkHandle, "mach_vm_protect", "10.4")
	registerFunc(&_mach_vm_purgable_control, &_mach_vm_purgable_controlErr, frameworkHandle, "mach_vm_purgable_control", "10.5")
	registerFunc(&_mach_vm_range_create, &_mach_vm_range_createErr, frameworkHandle, "mach_vm_range_create", "14.0")
	registerFunc(&_mach_vm_read, &_mach_vm_readErr, frameworkHandle, "mach_vm_read", "10.4")
	registerFunc(&_mach_vm_read_list, &_mach_vm_read_listErr, frameworkHandle, "mach_vm_read_list", "10.4")
	registerFunc(&_mach_vm_read_overwrite, &_mach_vm_read_overwriteErr, frameworkHandle, "mach_vm_read_overwrite", "10.4")
	registerFunc(&_mach_vm_region, &_mach_vm_regionErr, frameworkHandle, "mach_vm_region", "10.4")
	registerFunc(&_mach_vm_region_recurse, &_mach_vm_region_recurseErr, frameworkHandle, "mach_vm_region_recurse", "10.4")
	registerFunc(&_mach_vm_remap, &_mach_vm_remapErr, frameworkHandle, "mach_vm_remap", "10.4")
	registerFunc(&_mach_vm_remap_new, &_mach_vm_remap_newErr, frameworkHandle, "mach_vm_remap_new", "11.3")
	registerFunc(&_mach_vm_wire, &_mach_vm_wireErr, frameworkHandle, "mach_vm_wire", "10.4")
	registerFunc(&_mach_vm_write, &_mach_vm_writeErr, frameworkHandle, "mach_vm_write", "10.4")
	registerFunc(&_mach_voucher_attr_command, &_mach_voucher_attr_commandErr, frameworkHandle, "mach_voucher_attr_command", "10.10")
	registerFunc(&_mach_voucher_debug_info, &_mach_voucher_debug_infoErr, frameworkHandle, "mach_voucher_debug_info", "10.10")
	registerFunc(&_mach_voucher_extract_all_attr_recipes, &_mach_voucher_extract_all_attr_recipesErr, frameworkHandle, "mach_voucher_extract_all_attr_recipes", "10.10")
	registerFunc(&_mach_voucher_extract_attr_content, &_mach_voucher_extract_attr_contentErr, frameworkHandle, "mach_voucher_extract_attr_content", "10.10")
	registerFunc(&_mach_voucher_extract_attr_recipe, &_mach_voucher_extract_attr_recipeErr, frameworkHandle, "mach_voucher_extract_attr_recipe", "10.10")
	registerFunc(&_mach_zone_force_gc, &_mach_zone_force_gcErr, frameworkHandle, "mach_zone_force_gc", "10.8")
	registerFunc(&_mach_zone_get_btlog_records, &_mach_zone_get_btlog_recordsErr, frameworkHandle, "mach_zone_get_btlog_records", "10.14")
	registerFunc(&_mach_zone_get_zlog_zones, &_mach_zone_get_zlog_zonesErr, frameworkHandle, "mach_zone_get_zlog_zones", "10.14")
	registerFunc(&_mach_zone_info, &_mach_zone_infoErr, frameworkHandle, "mach_zone_info", "10.7")
	registerFunc(&_mach_zone_info_for_largest_zone, &_mach_zone_info_for_largest_zoneErr, frameworkHandle, "mach_zone_info_for_largest_zone", "10.13.4")
	registerFunc(&_mach_zone_info_for_zone, &_mach_zone_info_for_zoneErr, frameworkHandle, "mach_zone_info_for_zone", "10.13.4")
	registerFunc(&_memchr, &_memchrErr, frameworkHandle, "memchr", "10.9")
	registerFunc(&_memcmp, &_memcmpErr, frameworkHandle, "memcmp", "10.0")
	registerFunc(&_memcpy, &_memcpyErr, frameworkHandle, "memcpy", "10.0")
	registerFunc(&_memmove, &_memmoveErr, frameworkHandle, "memmove", "10.0")
	registerFunc(&_memset, &_memsetErr, frameworkHandle, "memset", "10.0")
	registerFunc(&_memset_s, &_memset_sErr, frameworkHandle, "memset_s", "10.13")
	registerFunc(&_mig_allocate, &_mig_allocateErr, frameworkHandle, "mig_allocate", "10.4")
	registerFunc(&_mig_dealloc_reply_port, &_mig_dealloc_reply_portErr, frameworkHandle, "mig_dealloc_reply_port", "10.0")
	registerFunc(&_mig_deallocate, &_mig_deallocateErr, frameworkHandle, "mig_deallocate", "10.4")
	registerFunc(&_mig_get_reply_port, &_mig_get_reply_portErr, frameworkHandle, "mig_get_reply_port", "10.0")
	registerFunc(&_mig_put_reply_port, &_mig_put_reply_portErr, frameworkHandle, "mig_put_reply_port", "10.0")
	registerFunc(&_mig_strncpy, &_mig_strncpyErr, frameworkHandle, "mig_strncpy", "10.0")
	registerFunc(&_mig_strncpy_zerofill, &_mig_strncpy_zerofillErr, frameworkHandle, "mig_strncpy_zerofill", "10.12")
	registerFunc(&_modf, &_modfErr, frameworkHandle, "modf", "10.10")
	registerFunc(&_modff, &_modffErr, frameworkHandle, "modff", "10.10")
	registerFunc(&_modfl, &_modflErr, frameworkHandle, "modfl", "10.10")
	registerFunc(&_nan, &_nanErr, frameworkHandle, "nan", "10.10")
	registerFunc(&_nanf, &_nanfErr, frameworkHandle, "nanf", "10.10")
	registerFunc(&_nanl, &_nanlErr, frameworkHandle, "nanl", "10.10")
	registerFunc(&_nearbyint, &_nearbyintErr, frameworkHandle, "nearbyint", "10.10")
	registerFunc(&_nearbyintf, &_nearbyintfErr, frameworkHandle, "nearbyintf", "10.10")
	registerFunc(&_nearbyintl, &_nearbyintlErr, frameworkHandle, "nearbyintl", "10.10")
	registerFunc(&_nextafter, &_nextafterErr, frameworkHandle, "nextafter", "10.10")
	registerFunc(&_nextafterf, &_nextafterfErr, frameworkHandle, "nextafterf", "10.10")
	registerFunc(&_nextafterl, &_nextafterlErr, frameworkHandle, "nextafterl", "10.10")
	registerFunc(&_nexttoward, &_nexttowardErr, frameworkHandle, "nexttoward", "10.10")
	registerFunc(&_nexttowardf, &_nexttowardfErr, frameworkHandle, "nexttowardf", "10.10")
	registerFunc(&_nexttowardl, &_nexttowardlErr, frameworkHandle, "nexttowardl", "10.10")
	registerFunc(&_nfsclnt, &_nfsclntErr, frameworkHandle, "nfsclnt", "13.0")
	registerFunc(&_os_log_create, &_os_log_createErr, frameworkHandle, "os_log_create", "10.12")
	registerFunc(&_os_release, &_os_releaseErr, frameworkHandle, "os_release", "10.12")
	registerFunc(&_os_retain, &_os_retainErr, frameworkHandle, "os_retain", "10.12")
	registerFunc(&_panic, &_panicErr, frameworkHandle, "panic", "10.0")
	registerFunc(&_pow, &_powErr, frameworkHandle, "pow", "10.10")
	registerFunc(&_powf, &_powfErr, frameworkHandle, "powf", "10.10")
	registerFunc(&_powl, &_powlErr, frameworkHandle, "powl", "10.10")
	registerFunc(&_printf, &_printfErr, frameworkHandle, "printf", "10.0")
	registerFunc(&_proc_name, &_proc_nameErr, frameworkHandle, "proc_name", "10.4")
	registerFunc(&_processor_assign, &_processor_assignErr, frameworkHandle, "processor_assign", "10.0")
	registerFunc(&_processor_control, &_processor_controlErr, frameworkHandle, "processor_control", "10.0")
	registerFunc(&_processor_exit, &_processor_exitErr, frameworkHandle, "processor_exit", "10.0")
	registerFunc(&_processor_get_assignment, &_processor_get_assignmentErr, frameworkHandle, "processor_get_assignment", "10.0")
	registerFunc(&_processor_info, &_processor_infoErr, frameworkHandle, "processor_info", "10.0")
	registerFunc(&_processor_set_create, &_processor_set_createErr, frameworkHandle, "processor_set_create", "10.0")
	registerFunc(&_processor_set_default, &_processor_set_defaultErr, frameworkHandle, "processor_set_default", "10.0")
	registerFunc(&_processor_set_destroy, &_processor_set_destroyErr, frameworkHandle, "processor_set_destroy", "10.0")
	registerFunc(&_processor_set_info, &_processor_set_infoErr, frameworkHandle, "processor_set_info", "10.0")
	registerFunc(&_processor_set_max_priority, &_processor_set_max_priorityErr, frameworkHandle, "processor_set_max_priority", "10.0")
	registerFunc(&_processor_set_policy_control, &_processor_set_policy_controlErr, frameworkHandle, "processor_set_policy_control", "10.0")
	registerFunc(&_processor_set_policy_disable, &_processor_set_policy_disableErr, frameworkHandle, "processor_set_policy_disable", "10.0")
	registerFunc(&_processor_set_policy_enable, &_processor_set_policy_enableErr, frameworkHandle, "processor_set_policy_enable", "10.0")
	registerFunc(&_processor_set_stack_usage, &_processor_set_stack_usageErr, frameworkHandle, "processor_set_stack_usage", "10.0")
	registerFunc(&_processor_set_statistics, &_processor_set_statisticsErr, frameworkHandle, "processor_set_statistics", "10.0")
	registerFunc(&_processor_set_tasks, &_processor_set_tasksErr, frameworkHandle, "processor_set_tasks", "10.0")
	registerFunc(&_processor_set_tasks_with_flavor, &_processor_set_tasks_with_flavorErr, frameworkHandle, "processor_set_tasks_with_flavor", "11.0")
	registerFunc(&_processor_set_threads, &_processor_set_threadsErr, frameworkHandle, "processor_set_threads", "10.0")
	registerFunc(&_processor_start, &_processor_startErr, frameworkHandle, "processor_start", "10.0")
	registerFunc(&_random, &_randomErr, frameworkHandle, "random", "10.0")
	registerFunc(&_remainder, &_remainderErr, frameworkHandle, "remainder", "10.10")
	registerFunc(&_remainderf, &_remainderfErr, frameworkHandle, "remainderf", "10.10")
	registerFunc(&_remainderl, &_remainderlErr, frameworkHandle, "remainderl", "10.10")
	registerFunc(&_remque, &_remqueErr, frameworkHandle, "remque", "10.0")
	registerFunc(&_remquo, &_remquoErr, frameworkHandle, "remquo", "10.10")
	registerFunc(&_remquof, &_remquofErr, frameworkHandle, "remquof", "10.10")
	registerFunc(&_remquol, &_remquolErr, frameworkHandle, "remquol", "10.10")
	registerFunc(&_rint, &_rintErr, frameworkHandle, "rint", "10.10")
	registerFunc(&_rintf, &_rintfErr, frameworkHandle, "rintf", "10.10")
	registerFunc(&_rintl, &_rintlErr, frameworkHandle, "rintl", "10.10")
	registerFunc(&_round, &_roundErr, frameworkHandle, "round", "10.10")
	registerFunc(&_roundf, &_roundfErr, frameworkHandle, "roundf", "10.10")
	registerFunc(&_roundl, &_roundlErr, frameworkHandle, "roundl", "10.10")
	registerFunc(&_scalb, &_scalbErr, frameworkHandle, "scalb", "10.10")
	registerFunc(&_scalbln, &_scalblnErr, frameworkHandle, "scalbln", "10.10")
	registerFunc(&_scalblnf, &_scalblnfErr, frameworkHandle, "scalblnf", "10.10")
	registerFunc(&_scalblnl, &_scalblnlErr, frameworkHandle, "scalblnl", "10.10")
	registerFunc(&_scalbn, &_scalbnErr, frameworkHandle, "scalbn", "10.10")
	registerFunc(&_scalbnf, &_scalbnfErr, frameworkHandle, "scalbnf", "10.10")
	registerFunc(&_scalbnl, &_scalbnlErr, frameworkHandle, "scalbnl", "10.10")
	registerFunc(&_selectFunc, &_selectFuncErr, frameworkHandle, "select", "11.0")
	registerFunc(&_semaphore_create, &_semaphore_createErr, frameworkHandle, "semaphore_create", "10.0")
	registerFunc(&_semaphore_destroy, &_semaphore_destroyErr, frameworkHandle, "semaphore_destroy", "10.0")
	registerFunc(&_semaphore_signal, &_semaphore_signalErr, frameworkHandle, "semaphore_signal", "10.0")
	registerFunc(&_semaphore_signal_all, &_semaphore_signal_allErr, frameworkHandle, "semaphore_signal_all", "10.0")
	registerFunc(&_semaphore_wait, &_semaphore_waitErr, frameworkHandle, "semaphore_wait", "10.0")
	registerFunc(&_signal, &_signalErr, frameworkHandle, "signal", "10.0")
	registerFunc(&_sin, &_sinErr, frameworkHandle, "sin", "10.10")
	registerFunc(&_sinf, &_sinfErr, frameworkHandle, "sinf", "10.10")
	registerFunc(&_sinh, &_sinhErr, frameworkHandle, "sinh", "10.10")
	registerFunc(&_sinhf, &_sinhfErr, frameworkHandle, "sinhf", "10.10")
	registerFunc(&_sinhl, &_sinhlErr, frameworkHandle, "sinhl", "10.10")
	registerFunc(&_sinl, &_sinlErr, frameworkHandle, "sinl", "10.10")
	registerFunc(&_snprintf, &_snprintfErr, frameworkHandle, "snprintf", "10.0")
	registerFunc(&_sprintf, &_sprintfErr, frameworkHandle, "sprintf", "10.12")
	registerFunc(&_sqrt, &_sqrtErr, frameworkHandle, "sqrt", "10.10")
	registerFunc(&_sqrtf, &_sqrtfErr, frameworkHandle, "sqrtf", "10.9")
	registerFunc(&_sqrtl, &_sqrtlErr, frameworkHandle, "sqrtl", "10.10")
	registerFunc(&_sscanf, &_sscanfErr, frameworkHandle, "sscanf", "10.0")
	registerFunc(&_strcasecmp, &_strcasecmpErr, frameworkHandle, "strcasecmp", "10.4")
	registerFunc(&_strcat, &_strcatErr, frameworkHandle, "strcat", "10.0")
	registerFunc(&_strchr, &_strchrErr, frameworkHandle, "strchr", "10.0")
	registerFunc(&_strcmp, &_strcmpErr, frameworkHandle, "strcmp", "10.0")
	registerFunc(&_strcpy, &_strcpyErr, frameworkHandle, "strcpy", "10.0")
	registerFunc(&_strlcat, &_strlcatErr, frameworkHandle, "strlcat", "10.5")
	registerFunc(&_strlcpy, &_strlcpyErr, frameworkHandle, "strlcpy", "10.5")
	registerFunc(&_strlen, &_strlenErr, frameworkHandle, "strlen", "10.0")
	registerFunc(&_strncasecmp, &_strncasecmpErr, frameworkHandle, "strncasecmp", "10.4")
	registerFunc(&_strncat, &_strncatErr, frameworkHandle, "strncat", "10.0")
	registerFunc(&_strncmp, &_strncmpErr, frameworkHandle, "strncmp", "10.0")
	registerFunc(&_strncpy, &_strncpyErr, frameworkHandle, "strncpy", "10.0")
	registerFunc(&_strnlen, &_strnlenErr, frameworkHandle, "strnlen", "10.5")
	registerFunc(&_strnstr, &_strnstrErr, frameworkHandle, "strnstr", "10.9")
	registerFunc(&_strsep, &_strsepErr, frameworkHandle, "strsep", "10.5")
	registerFunc(&_strtol, &_strtolErr, frameworkHandle, "strtol", "10.0")
	registerFunc(&_strtoq, &_strtoqErr, frameworkHandle, "strtoq", "10.1")
	registerFunc(&_strtoul, &_strtoulErr, frameworkHandle, "strtoul", "10.0")
	registerFunc(&_strtouq, &_strtouqErr, frameworkHandle, "strtouq", "10.1")
	registerFunc(&_sysctlbyname, &_sysctlbynameErr, frameworkHandle, "sysctlbyname", "10.0")
	registerFunc(&_tan, &_tanErr, frameworkHandle, "tan", "10.10")
	registerFunc(&_tanf, &_tanfErr, frameworkHandle, "tanf", "10.10")
	registerFunc(&_tanh, &_tanhErr, frameworkHandle, "tanh", "10.10")
	registerFunc(&_tanhf, &_tanhfErr, frameworkHandle, "tanhf", "10.10")
	registerFunc(&_tanhl, &_tanhlErr, frameworkHandle, "tanhl", "10.10")
	registerFunc(&_tanl, &_tanlErr, frameworkHandle, "tanl", "10.10")
	registerFunc(&_task_assign, &_task_assignErr, frameworkHandle, "task_assign", "10.0")
	registerFunc(&_task_assign_default, &_task_assign_defaultErr, frameworkHandle, "task_assign_default", "10.0")
	registerFunc(&_task_create, &_task_createErr, frameworkHandle, "task_create", "10.0")
	registerFunc(&_task_create_identity_token, &_task_create_identity_tokenErr, frameworkHandle, "task_create_identity_token", "11.3")
	registerFunc(&_task_dyld_process_info_notify_deregister, &_task_dyld_process_info_notify_deregisterErr, frameworkHandle, "task_dyld_process_info_notify_deregister", "11.3")
	registerFunc(&_task_dyld_process_info_notify_register, &_task_dyld_process_info_notify_registerErr, frameworkHandle, "task_dyld_process_info_notify_register", "11.3")
	registerFunc(&_task_generate_corpse, &_task_generate_corpseErr, frameworkHandle, "task_generate_corpse", "10.12")
	registerFunc(&_task_get_assignment, &_task_get_assignmentErr, frameworkHandle, "task_get_assignment", "10.0")
	registerFunc(&_task_get_dyld_image_infos, &_task_get_dyld_image_infosErr, frameworkHandle, "task_get_dyld_image_infos", "10.12")
	registerFunc(&_task_get_emulation_vector, &_task_get_emulation_vectorErr, frameworkHandle, "task_get_emulation_vector", "10.0")
	registerFunc(&_task_get_exc_guard_behavior, &_task_get_exc_guard_behaviorErr, frameworkHandle, "task_get_exc_guard_behavior", "10.15")
	registerFunc(&_task_get_exception_ports, &_task_get_exception_portsErr, frameworkHandle, "task_get_exception_ports", "10.0")
	registerFunc(&_task_get_exception_ports_info, &_task_get_exception_ports_infoErr, frameworkHandle, "task_get_exception_ports_info", "11.3")
	registerFunc(&_task_get_mach_voucher, &_task_get_mach_voucherErr, frameworkHandle, "task_get_mach_voucher", "10.10")
	registerFunc(&_task_get_special_port, &_task_get_special_portErr, frameworkHandle, "task_get_special_port", "10.0")
	registerFunc(&_task_get_state, &_task_get_stateErr, frameworkHandle, "task_get_state", "10.6")
	registerFunc(&_task_identity_token_get_task_port, &_task_identity_token_get_task_portErr, frameworkHandle, "task_identity_token_get_task_port", "11.3")
	registerFunc(&_task_info, &_task_infoErr, frameworkHandle, "task_info", "10.0")
	registerFunc(&_task_inspect, &_task_inspectErr, frameworkHandle, "task_inspect", "10.13")
	registerFunc(&_task_map_corpse_info, &_task_map_corpse_infoErr, frameworkHandle, "task_map_corpse_info", "10.12")
	registerFunc(&_task_map_corpse_info_64, &_task_map_corpse_info_64Err, frameworkHandle, "task_map_corpse_info_64", "10.12")
	registerFunc(&_task_map_kcdata_object_64, &_task_map_kcdata_object_64Err, frameworkHandle, "task_map_kcdata_object_64", "13.0")
	registerFunc(&_task_policy, &_task_policyErr, frameworkHandle, "task_policy", "10.0")
	registerFunc(&_task_policy_get, &_task_policy_getErr, frameworkHandle, "task_policy_get", "10.0")
	registerFunc(&_task_policy_set, &_task_policy_setErr, frameworkHandle, "task_policy_set", "10.0")
	registerFunc(&_task_purgable_info, &_task_purgable_infoErr, frameworkHandle, "task_purgable_info", "10.9")
	registerFunc(&_task_register_dyld_get_process_state, &_task_register_dyld_get_process_stateErr, frameworkHandle, "task_register_dyld_get_process_state", "10.12")
	registerFunc(&_task_register_dyld_image_infos, &_task_register_dyld_image_infosErr, frameworkHandle, "task_register_dyld_image_infos", "10.12")
	registerFunc(&_task_register_dyld_set_dyld_state, &_task_register_dyld_set_dyld_stateErr, frameworkHandle, "task_register_dyld_set_dyld_state", "10.12")
	registerFunc(&_task_register_dyld_shared_cache_image_info, &_task_register_dyld_shared_cache_image_infoErr, frameworkHandle, "task_register_dyld_shared_cache_image_info", "10.12")
	registerFunc(&_task_register_hardened_exception_handler, &_task_register_hardened_exception_handlerErr, frameworkHandle, "task_register_hardened_exception_handler", "15.0")
	registerFunc(&_task_restartable_ranges_register, &_task_restartable_ranges_registerErr, frameworkHandle, "task_restartable_ranges_register", "10.15")
	registerFunc(&_task_restartable_ranges_synchronize, &_task_restartable_ranges_synchronizeErr, frameworkHandle, "task_restartable_ranges_synchronize", "10.15")
	registerFunc(&_task_resume, &_task_resumeErr, frameworkHandle, "task_resume", "10.0")
	registerFunc(&_task_resume2, &_task_resume2Err, frameworkHandle, "task_resume2", "10.9")
	registerFunc(&_task_sample, &_task_sampleErr, frameworkHandle, "task_sample", "10.0")
	registerFunc(&_task_set_corpse_forking_behavior, &_task_set_corpse_forking_behaviorErr, frameworkHandle, "task_set_corpse_forking_behavior", "12.0")
	registerFunc(&_task_set_emulation, &_task_set_emulationErr, frameworkHandle, "task_set_emulation", "10.0")
	registerFunc(&_task_set_emulation_vector, &_task_set_emulation_vectorErr, frameworkHandle, "task_set_emulation_vector", "10.0")
	registerFunc(&_task_set_exc_guard_behavior, &_task_set_exc_guard_behaviorErr, frameworkHandle, "task_set_exc_guard_behavior", "10.15")
	registerFunc(&_task_set_exception_ports, &_task_set_exception_portsErr, frameworkHandle, "task_set_exception_ports", "10.0")
	registerFunc(&_task_set_info, &_task_set_infoErr, frameworkHandle, "task_set_info", "10.0")
	registerFunc(&_task_set_mach_voucher, &_task_set_mach_voucherErr, frameworkHandle, "task_set_mach_voucher", "10.10")
	registerFunc(&_task_set_phys_footprint_limit, &_task_set_phys_footprint_limitErr, frameworkHandle, "task_set_phys_footprint_limit", "10.9")
	registerFunc(&_task_set_policy, &_task_set_policyErr, frameworkHandle, "task_set_policy", "10.0")
	registerFunc(&_task_set_port_space, &_task_set_port_spaceErr, frameworkHandle, "task_set_port_space", "10.0")
	registerFunc(&_task_set_ras_pc, &_task_set_ras_pcErr, frameworkHandle, "task_set_ras_pc", "10.0")
	registerFunc(&_task_set_special_port, &_task_set_special_portErr, frameworkHandle, "task_set_special_port", "10.0")
	registerFunc(&_task_set_state, &_task_set_stateErr, frameworkHandle, "task_set_state", "10.6")
	registerFunc(&_task_suspend, &_task_suspendErr, frameworkHandle, "task_suspend", "10.0")
	registerFunc(&_task_suspend2, &_task_suspend2Err, frameworkHandle, "task_suspend2", "10.9")
	registerFunc(&_task_swap_exception_ports, &_task_swap_exception_portsErr, frameworkHandle, "task_swap_exception_ports", "10.0")
	registerFunc(&_task_swap_mach_voucher, &_task_swap_mach_voucherErr, frameworkHandle, "task_swap_mach_voucher", "10.10")
	registerFunc(&_task_terminate, &_task_terminateErr, frameworkHandle, "task_terminate", "10.0")
	registerFunc(&_task_test_async_upcall_propagation, &_task_test_async_upcall_propagationErr, frameworkHandle, "task_test_async_upcall_propagation", "12.3")
	registerFunc(&_task_test_sync_upcall, &_task_test_sync_upcallErr, frameworkHandle, "task_test_sync_upcall", "12.0")
	registerFunc(&_task_threads, &_task_threadsErr, frameworkHandle, "task_threads", "10.0")
	registerFunc(&_task_unregister_dyld_image_infos, &_task_unregister_dyld_image_infosErr, frameworkHandle, "task_unregister_dyld_image_infos", "10.12")
	registerFunc(&_task_zone_info, &_task_zone_infoErr, frameworkHandle, "task_zone_info", "10.7")
	registerFunc(&_tgamma, &_tgammaErr, frameworkHandle, "tgamma", "10.10")
	registerFunc(&_tgammaf, &_tgammafErr, frameworkHandle, "tgammaf", "10.10")
	registerFunc(&_tgammal, &_tgammalErr, frameworkHandle, "tgammal", "10.10")
	registerFunc(&_thread_abort, &_thread_abortErr, frameworkHandle, "thread_abort", "10.0")
	registerFunc(&_thread_abort_safely, &_thread_abort_safelyErr, frameworkHandle, "thread_abort_safely", "10.0")
	registerFunc(&_thread_adopt_exception_handler, &_thread_adopt_exception_handlerErr, frameworkHandle, "thread_adopt_exception_handler", "15.0")
	registerFunc(&_thread_assign, &_thread_assignErr, frameworkHandle, "thread_assign", "10.0")
	registerFunc(&_thread_assign_default, &_thread_assign_defaultErr, frameworkHandle, "thread_assign_default", "10.0")
	registerFunc(&_thread_convert_thread_state, &_thread_convert_thread_stateErr, frameworkHandle, "thread_convert_thread_state", "11.0")
	registerFunc(&_thread_create, &_thread_createErr, frameworkHandle, "thread_create", "10.0")
	registerFunc(&_thread_create_running, &_thread_create_runningErr, frameworkHandle, "thread_create_running", "10.0")
	registerFunc(&_thread_depress_abort, &_thread_depress_abortErr, frameworkHandle, "thread_depress_abort", "10.0")
	registerFunc(&_thread_get_assignment, &_thread_get_assignmentErr, frameworkHandle, "thread_get_assignment", "10.0")
	registerFunc(&_thread_get_exception_ports, &_thread_get_exception_portsErr, frameworkHandle, "thread_get_exception_ports", "10.0")
	registerFunc(&_thread_get_exception_ports_info, &_thread_get_exception_ports_infoErr, frameworkHandle, "thread_get_exception_ports_info", "11.3")
	registerFunc(&_thread_get_mach_voucher, &_thread_get_mach_voucherErr, frameworkHandle, "thread_get_mach_voucher", "10.10")
	registerFunc(&_thread_get_special_port, &_thread_get_special_portErr, frameworkHandle, "thread_get_special_port", "10.0")
	registerFunc(&_thread_get_state, &_thread_get_stateErr, frameworkHandle, "thread_get_state", "10.0")
	registerFunc(&_thread_info, &_thread_infoErr, frameworkHandle, "thread_info", "10.0")
	registerFunc(&_thread_policy, &_thread_policyErr, frameworkHandle, "thread_policy", "10.0")
	registerFunc(&_thread_policy_get, &_thread_policy_getErr, frameworkHandle, "thread_policy_get", "10.0")
	registerFunc(&_thread_policy_set, &_thread_policy_setErr, frameworkHandle, "thread_policy_set", "10.0")
	registerFunc(&_thread_resume, &_thread_resumeErr, frameworkHandle, "thread_resume", "10.0")
	registerFunc(&_thread_sample, &_thread_sampleErr, frameworkHandle, "thread_sample", "10.0")
	registerFunc(&_thread_set_exception_ports, &_thread_set_exception_portsErr, frameworkHandle, "thread_set_exception_ports", "10.0")
	registerFunc(&_thread_set_mach_voucher, &_thread_set_mach_voucherErr, frameworkHandle, "thread_set_mach_voucher", "10.10")
	registerFunc(&_thread_set_policy, &_thread_set_policyErr, frameworkHandle, "thread_set_policy", "10.0")
	registerFunc(&_thread_set_special_port, &_thread_set_special_portErr, frameworkHandle, "thread_set_special_port", "10.0")
	registerFunc(&_thread_set_state, &_thread_set_stateErr, frameworkHandle, "thread_set_state", "10.0")
	registerFunc(&_thread_suspend, &_thread_suspendErr, frameworkHandle, "thread_suspend", "10.0")
	registerFunc(&_thread_swap_exception_ports, &_thread_swap_exception_portsErr, frameworkHandle, "thread_swap_exception_ports", "10.0")
	registerFunc(&_thread_swap_mach_voucher, &_thread_swap_mach_voucherErr, frameworkHandle, "thread_swap_mach_voucher", "10.10")
	registerFunc(&_thread_terminate, &_thread_terminateErr, frameworkHandle, "thread_terminate", "10.0")
	registerFunc(&_thread_wire, &_thread_wireErr, frameworkHandle, "thread_wire", "10.0")
	registerFunc(&_timingsafe_bcmp, &_timingsafe_bcmpErr, frameworkHandle, "timingsafe_bcmp", "10.15")
	registerFunc(&_trunc, &_truncErr, frameworkHandle, "trunc", "10.10")
	registerFunc(&_truncf, &_truncfErr, frameworkHandle, "truncf", "10.10")
	registerFunc(&_truncl, &_trunclErr, frameworkHandle, "truncl", "10.10")
	registerFunc(&_uuid_clear, &_uuid_clearErr, frameworkHandle, "uuid_clear", "10.4")
	registerFunc(&_uuid_compare, &_uuid_compareErr, frameworkHandle, "uuid_compare", "10.4")
	registerFunc(&_uuid_copy, &_uuid_copyErr, frameworkHandle, "uuid_copy", "10.4")
	registerFunc(&_uuid_generate, &_uuid_generateErr, frameworkHandle, "uuid_generate", "10.4")
	registerFunc(&_uuid_generate_random, &_uuid_generate_randomErr, frameworkHandle, "uuid_generate_random", "10.4")
	registerFunc(&_uuid_generate_time, &_uuid_generate_timeErr, frameworkHandle, "uuid_generate_time", "10.4")
	registerFunc(&_uuid_is_null, &_uuid_is_nullErr, frameworkHandle, "uuid_is_null", "10.4")
	registerFunc(&_uuid_parse, &_uuid_parseErr, frameworkHandle, "uuid_parse", "10.4")
	registerFunc(&_uuid_unparse, &_uuid_unparseErr, frameworkHandle, "uuid_unparse", "10.4")
	registerFunc(&_uuid_unparse_lower, &_uuid_unparse_lowerErr, frameworkHandle, "uuid_unparse_lower", "10.4")
	registerFunc(&_uuid_unparse_upper, &_uuid_unparse_upperErr, frameworkHandle, "uuid_unparse_upper", "10.4")
	registerFunc(&_vm_allocate, &_vm_allocateErr, frameworkHandle, "vm_allocate", "10.0")
	registerFunc(&_vm_allocate_cpm, &_vm_allocate_cpmErr, frameworkHandle, "vm_allocate_cpm", "10.0")
	registerFunc(&_vm_behavior_set, &_vm_behavior_setErr, frameworkHandle, "vm_behavior_set", "10.0")
	registerFunc(&_vm_copy, &_vm_copyErr, frameworkHandle, "vm_copy", "10.0")
	registerFunc(&_vm_deallocate, &_vm_deallocateErr, frameworkHandle, "vm_deallocate", "10.0")
	registerFunc(&_vm_inherit, &_vm_inheritErr, frameworkHandle, "vm_inherit", "10.0")
	registerFunc(&_vm_machine_attribute, &_vm_machine_attributeErr, frameworkHandle, "vm_machine_attribute", "10.0")
	registerFunc(&_vm_map, &_vm_mapErr, frameworkHandle, "vm_map", "10.0")
	registerFunc(&_vm_map_page_query, &_vm_map_page_queryErr, frameworkHandle, "vm_map_page_query", "10.0")
	registerFunc(&_vm_msync, &_vm_msyncErr, frameworkHandle, "vm_msync", "10.0")
	registerFunc(&_vm_protect, &_vm_protectErr, frameworkHandle, "vm_protect", "10.0")
	registerFunc(&_vm_purgable_control, &_vm_purgable_controlErr, frameworkHandle, "vm_purgable_control", "10.4")
	registerFunc(&_vm_read, &_vm_readErr, frameworkHandle, "vm_read", "10.0")
	registerFunc(&_vm_read_list, &_vm_read_listErr, frameworkHandle, "vm_read_list", "10.0")
	registerFunc(&_vm_read_overwrite, &_vm_read_overwriteErr, frameworkHandle, "vm_read_overwrite", "10.0")
	registerFunc(&_vm_region_64, &_vm_region_64Err, frameworkHandle, "vm_region_64", "10.0")
	registerFunc(&_vm_region_recurse_64, &_vm_region_recurse_64Err, frameworkHandle, "vm_region_recurse_64", "10.0")
	registerFunc(&_vm_remap, &_vm_remapErr, frameworkHandle, "vm_remap", "10.0")
	registerFunc(&_vm_remap_new, &_vm_remap_newErr, frameworkHandle, "vm_remap_new", "11.3")
	registerFunc(&_vm_wire, &_vm_wireErr, frameworkHandle, "vm_wire", "10.0")
	registerFunc(&_vm_write, &_vm_writeErr, frameworkHandle, "vm_write", "10.0")
	registerFunc(&_vprintf, &_vprintfErr, frameworkHandle, "vprintf", "10.5")
	registerFunc(&_vsnprintf, &_vsnprintfErr, frameworkHandle, "vsnprintf", "10.0")
	registerFunc(&_vsprintf, &_vsprintfErr, frameworkHandle, "vsprintf", "10.0")
	registerFunc(&_vsscanf, &_vsscanfErr, frameworkHandle, "vsscanf", "10.4")
	registerFunc(&_y0, &_y0Err, frameworkHandle, "y0", "10.0")
	registerFunc(&_y1, &_y1Err, frameworkHandle, "y1", "10.0")
	registerFunc(&_yn, &_ynErr, frameworkHandle, "yn", "10.0")
}
