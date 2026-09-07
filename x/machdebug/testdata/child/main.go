// Command child is the debug target used by the x/machdebug tests.
//
// It prints the addresses of the state the tests act on, then serves a
// line protocol on stdin so a test can ask it what it observes:
//
//	counter <hex>   address of the counter global
//	bump <hex>      address of the bump function
//	ready
//
// Commands, one per line: "get" prints the counter, "set <hex>" assigns it,
// "bump <hex>" prints bump(n), "loop <n>" prints the sum of bump(0..n-1),
// and "exit" quits.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

// counter is a global the debugger reads and writes.
var counter uint64 = 0x1111222233334444

// bump is the function the debugger breaks on. It returns its argument
// plus one, so a debugger that rewrites x0 at entry changes the answer.
//
//go:noinline
func bump(x uint64) uint64 {
	return x + 1
}

// funcAddr returns the code address of a func value.
func funcAddr(f any) uintptr {
	// A non-nil func value is a pointer to a funcval whose first word is
	// the entry PC.
	type iface struct{ typ, data unsafe.Pointer }
	i := (*iface)(unsafe.Pointer(&f))
	return **(**uintptr)(unsafe.Pointer(&i.data))
}

func main() {
	fmt.Printf("counter %#x\n", uintptr(unsafe.Pointer(&counter)))
	fmt.Printf("bump %#x\n", funcAddr(bump))
	fmt.Println("ready")
	os.Stdout.Sync()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		f := strings.Fields(in.Text())
		if len(f) == 0 {
			continue
		}
		switch f[0] {
		case "get":
			fmt.Printf("counter %#x\n", counter)
		case "bump":
			n, err := strconv.ParseUint(strings.TrimPrefix(f[1], "0x"), 16, 64)
			if err != nil {
				fmt.Println("err", err)
				continue
			}
			fmt.Printf("bump %#x\n", bump(n))
		case "loop":
			n, err := strconv.Atoi(f[1])
			if err != nil {
				fmt.Println("err", err)
				continue
			}
			var sum uint64
			for i := range uint64(n) {
				sum += bump(i)
			}
			fmt.Printf("loop %#x\n", sum)
		case "set":
			v, err := strconv.ParseUint(strings.TrimPrefix(f[1], "0x"), 16, 64)
			if err != nil {
				fmt.Println("err", err)
				continue
			}
			counter = v
			fmt.Printf("set %#x\n", counter)
		case "exit":
			return
		default:
			fmt.Println("err unknown command")
		}
	}
}
