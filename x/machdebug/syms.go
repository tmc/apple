package machdebug

import (
	"fmt"
	"sync"

	"github.com/ebitengine/purego"
)

// taskForPID is task_for_pid, which has no documentation page and so is not
// generated into the kernel package. Resolving it here keeps the one missing
// primitive next to the code that needs it.
var taskForPID func(target uint32, pid int32, task *uint32) int32

var (
	resolveOnce sync.Once
	resolveErr  error
)

// machSymbols are every foreign symbol this package calls. They are all
// checked up front by resolve: a debugger that has to dlopen or bind a lazy
// stub while its target is stopped can deadlock against that target, so
// nothing is left to resolve after Attach touches the task.
//
// The kernel package binds its own symbols at package initialization, so the
// only work left here is to prove they are present and to bind task_for_pid.
var machSymbols = []string{
	"task_for_pid",
	"task_threads",
	"task_suspend",
	"task_resume",
	"task_get_exception_ports",
	"task_set_exception_ports",
	"thread_get_state",
	"thread_set_state",
	"mach_vm_read_overwrite",
	"mach_vm_write",
	"mach_vm_protect",
	"mach_vm_region",
	"mach_vm_deallocate",
	"mach_msg",
	"mach_port_allocate",
	"mach_port_insert_right",
	"mach_port_deallocate",
	"mach_port_mod_refs",
}

// resolve binds task_for_pid and verifies every other symbol is present.
func resolve() error {
	resolveOnce.Do(func() {
		for _, name := range machSymbols {
			sym, err := purego.Dlsym(purego.RTLD_DEFAULT, name)
			if err != nil || sym == 0 {
				resolveErr = fmt.Errorf("machdebug: resolve %s: %w", name, err)
				return
			}
			if name == "task_for_pid" {
				purego.RegisterFunc(&taskForPID, sym)
			}
		}
	})
	return resolveErr
}
