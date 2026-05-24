package rdma

import (
	"errors"
	"fmt"
	"syscall"
)

var (
	ErrRTRUnsafe = errors.New("rdma rtr unsafe")
)

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
		return "errno 16 (EBUSY; may indicate AppleThunderboltRDMA resource exhaustion or contaminated IOKit state, reboot the affected node before retrying)"
	case 19:
		return "errno 19 (ENODEV)"
	case 22:
		return "errno 22 (EINVAL)"
	case 38:
		return "errno 38 (ENOSYS)"
	case 45:
		return "errno 45 (EOPNOTSUPP)"
	case 60:
		return "errno 60 (ETIMEDOUT)"
	case 95:
		return "errno 95 (ENOTSUP)"
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
	case 45:
		return "EOPNOTSUPP"
	case 60:
		return "ETIMEDOUT"
	case 95:
		return "ENOTSUP"
	case 96:
		return "EPROTONOSUPPORT"
	default:
		return fmt.Sprintf("errno %d", errno)
	}
}
