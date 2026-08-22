package main

import (
	"syscall"
	"time"
)

// cpuTime is the processor time a repetition consumed, split the way the
// kernel accounts for it.
//
// The measurement covers the whole process, not the calling goroutine, which
// is the point: the work being attributed happens on dispatch worker threads
// as much as on the thread running the benchmark loop, and a per-thread
// number would miss most of it. Nothing else runs in the process during a
// repetition, so process time is per-message work plus a small constant.
//
// Comparing this against wall time separates work from waiting. A transport
// whose CPU time per message is far below its wall time per message is
// spending the difference blocked — on the network, or on a handoff between
// threads — and no amount of making the work cheaper will recover it.
type cpuTime struct {
	User float64 `json:"user_us"`
	Sys  float64 `json:"sys_us"`
}

// readCPUTime reports the process's consumed processor time. It returns the
// zero value if the kernel refuses, which loses the measurement rather than
// the run.
func readCPUTime() cpuTime {
	var ru syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &ru); err != nil {
		return cpuTime{}
	}
	return cpuTime{User: micros(ru.Utime), Sys: micros(ru.Stime)}
}

func micros(t syscall.Timeval) float64 {
	return float64(t.Sec)*1e6 + float64(t.Usec)
}

// sub returns the time consumed between an earlier reading and this one.
func (c cpuTime) sub(earlier cpuTime) cpuTime {
	return cpuTime{User: c.User - earlier.User, Sys: c.Sys - earlier.Sys}
}

func (c cpuTime) total() float64 { return c.User + c.Sys }

// busy is the fraction of elapsed time the process spent on a CPU. Values
// well below 1 mean the transport is waiting; values above 1 mean it had
// work on more than one core at once.
func (c cpuTime) busy(elapsed time.Duration) float64 {
	if elapsed <= 0 {
		return 0
	}
	return c.total() / (float64(elapsed) / float64(time.Microsecond))
}
