package rdma

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

var (
	// ErrNoDevice reports that no RDMA device matched the requested operation.
	ErrNoDevice = errors.New("rdma: no device")

	// ErrSymbolUnavailable reports that an RDMA provider symbol is unavailable.
	ErrSymbolUnavailable = errors.New("rdma: symbol unavailable")

	// ErrNilInput reports a nil handle or pointer passed before a provider call.
	ErrNilInput = errors.New("rdma: nil input")

	// ErrNilProviderResult reports a provider call that returned a nil handle.
	ErrNilProviderResult = errors.New("rdma: nil provider result")

	// ErrNegativeProviderReturn reports a provider call that returned a negative status.
	ErrNegativeProviderReturn = errors.New("rdma: negative provider return")

	// ErrProviderStatus reports a provider call that returned a non-zero status.
	ErrProviderStatus = errors.New("rdma: provider status")

	// ErrProviderTimeout reports a bounded provider subprocess that timed out.
	ErrProviderTimeout = errors.New("rdma: provider timeout")

	// ErrUnsupportedOperation reports an RDMA operation the provider does not support.
	ErrUnsupportedOperation = errors.New("rdma: unsupported operation")
)

// FailureClass identifies the provider-boundary failure class.
type FailureClass string

const (
	FailureNoDevice               FailureClass = "no_device"
	FailureSymbolUnavailable      FailureClass = "symbol_unavailable"
	FailureNilInput               FailureClass = "nil_input"
	FailureNilProviderResult      FailureClass = "nil_provider_result"
	FailureNegativeProviderReturn FailureClass = "negative_provider_return"
	FailureProviderStatus         FailureClass = "provider_status"
	FailureProviderTimeout        FailureClass = "provider_timeout"
	FailureUnsupportedOperation   FailureClass = "unsupported_operation"
)

// ProviderError describes an RDMA provider-boundary failure.
//
// Operation is the ibverbs operation. Device is set when the package can tie a
// context back to a named RDMA device. ContextOpen reports whether the failure
// happened after an RDMA context had been opened.
type ProviderError struct {
	Operation   string
	Device      string
	Context     RDMAContext
	ContextOpen bool
	Failure     FailureClass
	Input       string
	Result      string
	Return      int64
	ReturnSet   bool
	Errno       int
	ErrnoSet    bool
	Cause       error
}

func (e *ProviderError) Error() string {
	if e == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString("rdma")
	if e.Operation != "" {
		b.WriteString(": ")
		b.WriteString(e.Operation)
	}
	if e.Device != "" {
		b.WriteString(": device ")
		b.WriteString(e.Device)
	}
	if e.ContextOpen {
		b.WriteString(": after open context")
	} else {
		b.WriteString(": before open context")
	}
	switch e.Failure {
	case FailureNoDevice:
		b.WriteString(": no RDMA device")
	case FailureSymbolUnavailable:
		b.WriteString(": symbol unavailable")
	case FailureNilInput:
		b.WriteString(": nil ")
		b.WriteString(e.Input)
	case FailureNilProviderResult:
		b.WriteString(": provider returned nil ")
		b.WriteString(e.Result)
	case FailureNegativeProviderReturn:
		b.WriteString(": provider returned negative status")
	case FailureProviderStatus:
		b.WriteString(": provider returned status")
	case FailureProviderTimeout:
		b.WriteString(": provider timed out")
	case FailureUnsupportedOperation:
		b.WriteString(": unsupported operation")
	default:
		if e.Failure != "" {
			b.WriteString(": ")
			b.WriteString(string(e.Failure))
		}
	}
	if e.ReturnSet || e.ErrnoSet || e.Context != 0 {
		b.WriteString(" (")
		first := true
		add := func(s string) {
			if !first {
				b.WriteString(", ")
			}
			first = false
			b.WriteString(s)
		}
		if e.ReturnSet {
			add("return=" + strconv.FormatInt(e.Return, 10))
		}
		if e.ErrnoSet {
			add("errno=" + strconv.Itoa(e.Errno))
		}
		if e.Context != 0 {
			add(fmt.Sprintf("context=%#x", uintptr(e.Context)))
		}
		b.WriteString(")")
	}
	if e.Cause != nil && !errors.Is(e.Cause, providerSentinel(e.Failure)) {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	return b.String()
}

func (e *ProviderError) Unwrap() error {
	if e == nil {
		return nil
	}
	if e.Cause != nil {
		return e.Cause
	}
	return providerSentinel(e.Failure)
}

func (e *ProviderError) Is(target error) bool {
	if e == nil {
		return false
	}
	return target == providerSentinel(e.Failure)
}

func providerSentinel(f FailureClass) error {
	switch f {
	case FailureNoDevice:
		return ErrNoDevice
	case FailureSymbolUnavailable:
		return ErrSymbolUnavailable
	case FailureNilInput:
		return ErrNilInput
	case FailureNilProviderResult:
		return ErrNilProviderResult
	case FailureNegativeProviderReturn:
		return ErrNegativeProviderReturn
	case FailureProviderStatus:
		return ErrProviderStatus
	case FailureProviderTimeout:
		return ErrProviderTimeout
	case FailureUnsupportedOperation:
		return ErrUnsupportedOperation
	default:
		return nil
	}
}

func rdmaNilHandleError(op, name string) error {
	return &ProviderError{
		Operation: op,
		Failure:   FailureNilInput,
		Input:     name + " handle",
		Cause:     ErrNilInput,
	}
}

func rdmaNilPointerError(op, name string) error {
	return &ProviderError{
		Operation: op,
		Failure:   FailureNilInput,
		Input:     name + " pointer",
		Cause:     ErrNilInput,
	}
}

func rdmaNilProviderResultError(op, result string, ret int64, errno int, errnoSet bool, context RDMAContext, contextOpen bool) error {
	return &ProviderError{
		Operation:   op,
		Device:      rdmaContextDevice(context),
		Context:     context,
		ContextOpen: contextOpen,
		Failure:     FailureNilProviderResult,
		Result:      result,
		Return:      ret,
		ReturnSet:   true,
		Errno:       errno,
		ErrnoSet:    errnoSet,
		Cause:       ErrNilProviderResult,
	}
}

func rdmaNegativeProviderReturnError(op string, rc int, errno int, errnoSet bool, context RDMAContext, contextOpen bool) error {
	return &ProviderError{
		Operation:   op,
		Device:      rdmaContextDevice(context),
		Context:     context,
		ContextOpen: contextOpen,
		Failure:     FailureNegativeProviderReturn,
		Return:      int64(rc),
		ReturnSet:   true,
		Errno:       errno,
		ErrnoSet:    errnoSet,
		Cause:       ErrNegativeProviderReturn,
	}
}

func rdmaProviderStatusError(op string, rc int, context RDMAContext, contextOpen bool) error {
	return &ProviderError{
		Operation:   op,
		Device:      rdmaContextDevice(context),
		Context:     context,
		ContextOpen: contextOpen,
		Failure:     FailureProviderStatus,
		Return:      int64(rc),
		ReturnSet:   true,
		Errno:       rc,
		ErrnoSet:    rc != 0,
		Cause:       ErrProviderStatus,
	}
}

func rdmaWithDevice(err error, device string) error {
	if err == nil || device == "" {
		return err
	}
	var e *ProviderError
	if !errors.As(err, &e) {
		return err
	}
	copy := *e
	if copy.Device == "" {
		copy.Device = device
	}
	return &copy
}

// ErrnoText returns a compact name for common Apple RDMA errno values.
func ErrnoText(errno int) string {
	switch errno {
	case 1:
		return "errno 1 (EPERM)"
	case 2:
		return "errno 2 (ENOENT)"
	case 5:
		return "errno 5 (EIO)"
	case 6:
		return "errno 6 (ENXIO)"
	case int(syscall.ENOMEM):
		return "errno 12 (ENOMEM)"
	case 13:
		return "errno 13 (EACCES)"
	case int(syscall.EBUSY):
		return "errno 16 (EBUSY)"
	case 19:
		return "errno 19 (ENODEV)"
	case 22:
		return "errno 22 (EINVAL)"
	case 38:
		return "errno 38 (ENOSYS)"
	case int(syscall.ENOTSUP):
		return "errno 45 (ENOTSUP)"
	case 60:
		return "errno 60 (ETIMEDOUT)"
	case 96:
		return "errno 96 (EPROTONOSUPPORT)"
	default:
		return fmt.Sprintf("errno %d", errno)
	}
}

// ErrnoName returns the symbolic name for common Apple RDMA errno values.
func ErrnoName(errno int) string {
	switch errno {
	case 1:
		return "EPERM"
	case 2:
		return "ENOENT"
	case 5:
		return "EIO"
	case 6:
		return "ENXIO"
	case int(syscall.ENOMEM):
		return "ENOMEM"
	case 13:
		return "EACCES"
	case int(syscall.EBUSY):
		return "EBUSY"
	case 19:
		return "ENODEV"
	case 22:
		return "EINVAL"
	case 38:
		return "ENOSYS"
	case int(syscall.ENOTSUP):
		return "ENOTSUP"
	case 60:
		return "ETIMEDOUT"
	case 96:
		return "EPROTONOSUPPORT"
	default:
		return fmt.Sprintf("errno %d", errno)
	}
}
