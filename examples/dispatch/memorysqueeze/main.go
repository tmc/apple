// Command memorysqueeze simulates an inference server holding a large
// cache and shows what DISPATCH_SOURCE_TYPE_MEMORYPRESSURE buys it.
//
// The process "serves tokens" while growing a KV-cache stand-in (mmap'd
// chunks, every page touched so they count against phys_footprint). With
// -shed, a dispatch memory-pressure source drops half the cache on WARN
// and all but the newest chunk on CRITICAL, and the process keeps
// serving. Without -shed, pressure is ignored and the footprint only
// grows — under a real squeeze that twin is the jetsam target.
//
// Run two terminals, then squeeze:
//
//	go run ./examples/dispatch/memorysqueeze -shed
//	go run ./examples/dispatch/memorysqueeze
//	sudo memory_pressure -S -l warn     # simulate WARN (no real pressure)
//	sudo memory_pressure -S -l critical # simulate CRITICAL
//
// Every line the demo prints carries its own evidence: the phys_footprint
// it reads from task_info(TASK_VM_INFO), which is the number jetsam uses.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
)

const (
	chunkBytes = 64 << 20 // one cache chunk
	pageBytes  = 16 << 10 // Apple silicon page size

	// task_info flavor and field offset, verified against
	// <mach/task_info.h>: TASK_VM_INFO=22, sizeof(task_vm_info_data_t)=372
	// (count 93), offsetof(phys_footprint)=144.
	taskVMInfo       = 22
	physFootprintOff = 144
	taskVMInfoCount  = 93
)

// physFootprint returns the kernel's ledgered footprint for this process,
// the same number Activity Monitor's "Memory" column and jetsam use.
func physFootprint() uint64 {
	var info kernel.Task_vm_info_data_t
	count := kernel.Mach_msg_type_number_t(taskVMInfoCount)
	kr := kernel.Task_info(kernel.Mach_task_self(), taskVMInfo,
		kernel.Task_info_t(unsafe.Pointer(&info)), &count)
	if kr != kernel.KERN_SUCCESS {
		return 0
	}
	raw := (*[4 * taskVMInfoCount]byte)(unsafe.Pointer(&info))
	return binary.LittleEndian.Uint64(raw[physFootprintOff : physFootprintOff+8])
}

// cache is the KV-cache stand-in: mmap'd chunks with every page touched.
type cache struct {
	mu     sync.Mutex
	chunks [][]byte
}

func (c *cache) grow() {
	b, err := syscall.Mmap(-1, 0, chunkBytes,
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		log.Fatalf("mmap: %v", err)
	}
	for i := 0; i < len(b); i += pageBytes {
		b[i] = 1 // touch, so the page is resident and ledgered
	}
	c.mu.Lock()
	c.chunks = append(c.chunks, b)
	c.mu.Unlock()
}

// shed unmaps all but keep chunks (newest kept) and returns bytes freed.
// Munmap returns pages to the kernel immediately, so the footprint drop
// is visible on the very next heartbeat.
func (c *cache) shed(keep int) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.chunks) <= keep {
		return 0
	}
	drop := c.chunks[:len(c.chunks)-keep]
	c.chunks = append([][]byte(nil), c.chunks[len(c.chunks)-keep:]...)
	var freed uint64
	for _, b := range drop {
		freed += uint64(len(b))
		if err := syscall.Munmap(b); err != nil {
			log.Fatalf("munmap: %v", err)
		}
	}
	return freed
}

func (c *cache) size() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return uint64(len(c.chunks)) * chunkBytes
}

func mib(b uint64) string { return fmt.Sprintf("%d MiB", b>>20) }

func main() {
	shed := flag.Bool("shed", false, "shed cache on memory-pressure WARN/CRITICAL")
	target := flag.Uint64("cache", 2<<30, "cache size to grow toward, in bytes")
	flag.Parse()

	who := "stock"
	if *shed {
		who = "shed"
	}
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("[%s pid=%d] ", who, os.Getpid()))

	kv := new(cache)

	if *shed {
		queue := dispatch.QueueCreate("memorysqueeze.pressure")
		src := dispatch.NewMemoryPressureSource(
			dispatch.MemPressureNormal|dispatch.MemPressureWarn|dispatch.MemPressureCritical,
			queue,
			func(events dispatch.MemoryPressureFlags) {
				before := physFootprint()
				switch {
				case events&dispatch.MemPressureCritical != 0:
					freed := kv.shed(1)
					log.Printf("SHED level=critical dropped=%s footprint %s -> %s",
						mib(freed), mib(before), mib(physFootprint()))
				case events&dispatch.MemPressureWarn != 0:
					kv.mu.Lock()
					n := len(kv.chunks)
					kv.mu.Unlock()
					freed := kv.shed(n / 2)
					log.Printf("SHED level=warn dropped=%s footprint %s -> %s",
						mib(freed), mib(before), mib(physFootprint()))
				case events&dispatch.MemPressureNormal != 0:
					log.Printf("pressure normal again, footprint=%s", mib(before))
				}
			})
		defer src.Cancel()
		log.Printf("memory-pressure source armed (warn: drop half, critical: drop all but newest chunk)")
	} else {
		log.Printf("no pressure source: cache only grows")
	}

	// Serve: a token every 100ms; the cache grows a chunk per second
	// until it reaches the target (refill after a shed, like a KV cache
	// repopulating on the next requests).
	tokens := 0
	tick := time.NewTicker(100 * time.Millisecond)
	heartbeat := time.NewTicker(1 * time.Second)
	growth := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			tokens++
		case <-growth.C:
			if kv.size() < *target {
				kv.grow()
			}
		case <-heartbeat.C:
			log.Printf("serving tokens=%d cache=%s footprint=%s",
				tokens, mib(kv.size()), mib(physFootprint()))
		}
	}
}
