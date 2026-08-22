// Command machportkevent demonstrates waiting on Mach ports with kqueue's
// EVFILT_MACHPORT filter.
//
// It builds a small request/response service: a worker goroutine owns a port
// set, registers it with a kqueue, and blocks in kevent64 until a message
// arrives on any member port. Clients send Mach messages to individual ports
// in that set. One wait covers every port, so adding a port costs nothing at
// the wait site.
//
// Run it:
//
//	go run ./examples/kernel/machportkevent
//	go run ./examples/kernel/machportkevent -ports=8 -messages=4
//
// # Why a dedicated thread
//
// This wait cannot be handed to the Go runtime's network poller. The poller
// only ever registers EVFILT_READ, EVFILT_WRITE and EVFILT_USER, and a kqueue
// descriptor has no read(2) interface for it to use, so a blocking kevent64
// would stall the P it runs on. runtime.LockOSThread gives the wait a thread
// of its own, which is what the netpoller would otherwise have avoided.
//
// The trade is real: a blocked thread costs a thread, but it costs no CPU
// while idle, unlike polling. Whether that is the right trade depends on how
// often messages arrive relative to how long the process can afford to spin.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// message is a Mach message with room for its body and the trailer the kernel
// appends on receive. A receive buffer that is too small fails the call with
// MACH_RCV_TOO_LARGE rather than truncating.
type message struct {
	hdr  kernel.MachMsgHdr
	body [64]byte
}

// receiveBuffer is sized for the message plus the largest trailer.
type receiveBuffer struct {
	hdr  kernel.MachMsgHdr
	body [256]byte
}

func main() {
	log.SetFlags(0)
	nports := flag.Int("ports", 4, "number of ports in the port set")
	nmessages := flag.Int("messages", 3, "messages to send to each port")
	flag.Parse()

	if *nports < 1 || *nmessages < 1 {
		log.Fatal("-ports and -messages must be at least 1")
	}

	self := kernel.Mach_task_self()
	if self == 0 {
		log.Fatal("mach_task_self returned 0; kernel bindings did not load")
	}

	// A port set is a receive-side grouping. Waiting on the set delivers a
	// message sent to any member, which is what makes one kevent64 wait cover
	// every port.
	var portSet uint32
	if kr := kernel.Mach_port_allocate(self, kernel.MachPortRightPortSet, &portSet); kr != 0 {
		log.Fatalf("allocate port set: kern_return_t %d", kr)
	}
	defer kernel.Mach_port_deallocate(self, portSet)

	ports := make([]uint32, *nports)
	for i := range ports {
		p, err := newServicePort(self, portSet)
		if err != nil {
			log.Fatalf("port %d: %v", i, err)
		}
		ports[i] = p
		defer kernel.Mach_port_mod_refs(self, p, kernel.MachPortRightReceive, -1)
	}
	fmt.Printf("allocated %d ports in port set %#x\n", len(ports), portSet)

	kq := kernel.Kqueue()
	if kq < 0 {
		log.Fatal("kqueue failed")
	}
	defer kernel.CloseFD(int(kq))

	// ext[0] and ext[1] name a buffer the kernel fills with the message during
	// event delivery, so the wait and the receive are a single call. The buffer
	// must outlive the registration, hence the KeepAlive in waitLoop.
	var inbox receiveBuffer
	change := kernel.Kevent64_s{
		Ident:  uint64(portSet),
		Filter: kernel.EVFILT_MACHPORT,
		Flags:  kernel.EV_ADD | kernel.EV_ENABLE,
		Fflags: kernel.MachRcvMsgFflag,
		Ext: [2]uint64{
			uint64(uintptr(unsafe.Pointer(&inbox))),
			uint64(unsafe.Sizeof(inbox)),
		},
	}
	if n := kernel.KeventWait(int(kq), []kernel.Kevent64_s{change}, nil, 0, &kernel.Timespec{}); n < 0 {
		log.Fatal("registering EVFILT_MACHPORT failed")
	}
	fmt.Printf("registered port set with EVFILT_MACHPORT on kqueue %d\n\n", kq)

	want := len(ports) * *nmessages
	var wg sync.WaitGroup
	wg.Add(1)
	received := make(chan int32, want)
	go func() {
		defer wg.Done()
		waitLoop(int(kq), want, &inbox, received)
	}()

	// Send after the waiter is registered. Messages that arrive first are not
	// lost — they queue on the port — but sending second keeps the output in
	// the order the demo describes.
	for round := 0; round < *nmessages; round++ {
		for i, p := range ports {
			id := int32(round*len(ports) + i)
			if err := send(p, id); err != nil {
				log.Fatalf("send to port %d: %v", i, err)
			}
		}
	}

	wg.Wait()
	close(received)

	seen := make(map[int32]bool, want)
	for id := range received {
		seen[id] = true
	}
	fmt.Printf("\nreceived %d/%d distinct messages\n", len(seen), want)
	if len(seen) != want {
		fmt.Fprintln(os.Stderr, "some messages were not delivered")
		os.Exit(1)
	}
}

// newServicePort allocates a receive right, gives the task a send right to it,
// and joins it to set so one wait covers it.
func newServicePort(task, set uint32) (uint32, error) {
	var port uint32
	if kr := kernel.Mach_port_allocate(task, kernel.MachPortRightReceive, &port); kr != 0 {
		return 0, fmt.Errorf("allocate receive right: kern_return_t %d", kr)
	}
	if kr := kernel.Mach_port_insert_right(task, port, port, kernel.MachMsgTypeMakeSend); kr != 0 {
		return 0, fmt.Errorf("insert send right: kern_return_t %d", kr)
	}
	if kr := kernel.Mach_port_insert_member(task, port, set); kr != 0 {
		return 0, fmt.Errorf("join port set: kern_return_t %d", kr)
	}
	return port, nil
}

// send delivers one message to port. The send never blocks for long: the
// timeout bounds it in case the port's queue is full.
func send(port uint32, id int32) error {
	msg := message{hdr: kernel.MachMsgHdr{
		Bits:       kernel.MachMsgTypeCopySend,
		Size:       uint32(unsafe.Sizeof(message{})),
		RemotePort: port,
		ID:         id,
	}}
	kr := kernel.Mach_msg(
		unsafe.Pointer(&msg.hdr),
		kernel.MachSendMsg|kernel.MachSendTimeout,
		uint32(unsafe.Sizeof(message{})),
		0, 0, 1000, 0,
	)
	runtime.KeepAlive(&msg)
	if kr != 0 {
		return fmt.Errorf("mach_msg send: kern_return_t %#x", kr)
	}
	return nil
}

// waitLoop blocks in kevent64 until want messages have arrived, reporting each
// one's id. It runs on a locked thread because a blocking EVFILT_MACHPORT wait
// cannot be handed to the Go network poller.
func waitLoop(kq int, want int, inbox *receiveBuffer, out chan<- int32) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	events := make([]kernel.Kevent64_s, 1)
	timeout := &kernel.Timespec{Tv_sec: 5}
	start := time.Now()

	for got := 0; got < want; {
		n := kernel.KeventWait(kq, nil, events, 0, timeout)
		switch {
		case n < 0:
			log.Print("kevent64 failed")
			return
		case n == 0:
			log.Printf("timed out after %d/%d messages", got, want)
			return
		}
		if events[0].Flags&kernel.EV_ERROR != 0 {
			log.Printf("event reported errno %d", events[0].Data)
			return
		}

		// The kernel delivered the message into the registered buffer, so the
		// header is readable without a separate mach_msg receive.
		id := inbox.hdr.ID
		fmt.Printf("  woke after %-8s port %#x  msgh_id %d\n",
			time.Since(start).Round(time.Microsecond), inbox.hdr.LocalPort, id)
		out <- id
		got++
		runtime.KeepAlive(inbox)
	}
}
