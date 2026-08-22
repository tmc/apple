// Command udpbatchbench measures the Darwin batched UDP datapath
// (sendmsg_x/recvmsg_x via x/udpbatch) against the one-syscall-per-datagram
// path every Go program uses today.
//
// The default mode sweeps batch sizes over loopback and prints a
// packets-per-second table:
//
//	go run . -sweep
//
// IMPORTANT: loopback numbers measure the syscall boundary, not the network.
// A batching win here is real but may not transfer to the wire at the same
// magnitude; run the send/recv modes across two machines for wire numbers
// (see design/high-performance-darwin.md, "Loopback is not the wire").
//
// Split modes for two-machine runs:
//
//	udpbatchbench -recv -addr :9999
//	udpbatchbench -send -addr host:9999 -batch 32
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/x/sockbuf"
	"github.com/tmc/apple/x/udpbatch"
)

var (
	sweep      = flag.Bool("sweep", false, "sweep batch sizes over loopback and print a pps table")
	send       = flag.Bool("send", false, "send datagrams to -addr")
	recv       = flag.Bool("recv", false, "receive datagrams on -addr")
	addr       = flag.String("addr", "127.0.0.1:9999", "address to send to / listen on")
	batch      = flag.Int("batch", udpbatch.BatchSize, "datagrams per syscall (1 = stdlib path)")
	maxBatch   = flag.Int("max-batch", 128, "largest batch size included in -sweep")
	sweepBatch = flag.Int("sweep-batch", 0, "measure only batch 1 and this batch in -sweep")
	size       = flag.Int("size", 1200, "datagram payload size in bytes (1200 ~ QUIC)")
	duration   = flag.Duration("duration", 3*time.Second, "how long to run each measurement")
	bufMiB     = flag.Int("buf", 4, "socket buffer request in MiB")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if !udpbatch.Available() {
		log.Printf("warning: batched datapath unavailable (%v); only -batch=1 will work", udpbatch.SelfTestError())
	}

	switch {
	case *sweep:
		runSweep()
	case *send:
		runSend()
	case *recv:
		runRecv()
	default:
		flag.Usage()
		os.Exit(2)
	}
}

// runSweep measures pps over loopback for a range of batch sizes, same
// payload size and duration each, and prints the table the design docs ask
// for: where does the curve flatten?
func runSweep() {
	fmt.Printf("loopback sweep (minimal connected headers): %d-byte payloads, %v per row (syscall-boundary measurement, not the wire)\n\n", *size, *duration)
	fmt.Printf("%8s %12s %12s %8s %10s\n", "batch", "send pps", "recv pps", "vs b=1", "fill/call")
	var base float64
	batches := []int{1, 2, 4, 8, 16, 32, 64, 128}
	if *sweepBatch > 1 {
		batches = []int{1, *sweepBatch}
	}
	for _, b := range batches {
		if b > *maxBatch {
			break
		}
		if b > 1 && !udpbatch.Available() {
			break
		}
		sendPPS, recvPPS, fill := measure(b)
		if b == 1 {
			base = sendPPS
		}
		fmt.Printf("%8d %12.0f %12.0f %7.2fx %10.1f\n", b, sendPPS, recvPPS, sendPPS/base, fill)
	}
}

// measure runs a sender and receiver over loopback for the flag duration at
// the given batch size and reports send and receive packets per second.
func measure(batchSize int) (sendPPS, recvPPS, fillPerCall float64) {
	rc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	sc, err := net.DialUDP("udp4", nil, rc.LocalAddr().(*net.UDPAddr))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	sizeBuffers(rc, sc)

	var sent, received, recvCalls atomic.Int64
	stop := make(chan struct{})

	go func() {
		msgs := make([]udpbatch.Message, batchSize)
		for i := range msgs {
			msgs[i].Payload = make([]byte, *size+64)
		}
		rraw, _ := rc.SyscallConn()
		buf := make([]byte, *size+64)
		// One deadline for the whole run: it exists only so the loop can
		// observe stop; resetting it per call would dominate the measurement.
		rc.SetReadDeadline(time.Now().Add(*duration + time.Second))
		for {
			select {
			case <-stop:
				return
			default:
			}
			if batchSize == 1 {
				if _, err := rc.Read(buf); err != nil {
					continue
				}
				received.Add(1)
			} else {
				n, err := udpbatch.Recv(rraw, msgs)
				if err != nil {
					continue
				}
				received.Add(int64(n))
				recvCalls.Add(1)
			}
		}
	}()

	payloads := make([][]byte, batchSize)
	for i := range payloads {
		payloads[i] = make([]byte, *size)
	}
	sraw, _ := sc.SyscallConn()
	deadline := time.Now().Add(*duration)
	start := time.Now()
	for time.Now().Before(deadline) {
		if batchSize == 1 {
			if _, err := sc.Write(payloads[0]); err != nil {
				continue // ENOBUFS on loopback overrun is normal
			}
			sent.Add(1)
		} else {
			n, err := udpbatch.Send(sraw, payloads)
			if err != nil {
				continue
			}
			sent.Add(int64(n))
		}
	}
	elapsed := time.Since(start)
	time.Sleep(100 * time.Millisecond) // drain
	close(stop)

	fill := 1.0
	if c := recvCalls.Load(); c > 0 {
		fill = float64(received.Load()) / float64(c)
	}
	return float64(sent.Load()) / elapsed.Seconds(), float64(received.Load()) / elapsed.Seconds(), fill
}

func runSend() {
	raddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	sc, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	sizeBuffers(sc)

	payloads := make([][]byte, *batch)
	for i := range payloads {
		payloads[i] = make([]byte, *size)
	}
	sraw, _ := sc.SyscallConn()
	var sent int64
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	last, lastT := int64(0), time.Now()
	for {
		select {
		case <-tick.C:
			now := time.Now()
			log.Printf("send: %8.0f pps (batch=%d, %dB)", float64(sent-last)/now.Sub(lastT).Seconds(), *batch, *size)
			last, lastT = sent, now
		default:
		}
		if *batch == 1 {
			if _, err := sc.Write(payloads[0]); err == nil {
				sent++
			}
		} else {
			if n, err := udpbatch.Send(sraw, payloads); err == nil {
				sent += int64(n)
			} else {
				log.Fatalf("Send: %v", err)
			}
		}
	}
}

func runRecv() {
	laddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	rc, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()
	sizeBuffers(rc)

	msgs := make([]udpbatch.Message, *batch)
	for i := range msgs {
		msgs[i].Payload = make([]byte, *size+64)
	}
	rraw, _ := rc.SyscallConn()
	buf := make([]byte, *size+64)
	var received int64
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	last, lastT := int64(0), time.Now()
	for {
		select {
		case <-tick.C:
			now := time.Now()
			log.Printf("recv: %8.0f pps (batch=%d)", float64(received-last)/now.Sub(lastT).Seconds(), *batch)
			last, lastT = received, now
		default:
		}
		if *batch == 1 {
			rc.SetReadDeadline(time.Now().Add(time.Second))
			if _, err := rc.Read(buf); err == nil {
				received++
			}
		} else {
			n, err := udpbatch.Recv(rraw, msgs)
			if err != nil {
				log.Fatalf("Recv: %v", err)
			}
			received += int64(n)
		}
	}
}

// sizeBuffers applies -buf to each conn and reports any clamp the kernel
// applied — the x/sockbuf read-back contract in action.
func sizeBuffers(conns ...*net.UDPConn) {
	want := *bufMiB << 20
	for _, c := range conns {
		raw, err := c.SyscallConn()
		if err != nil {
			continue
		}
		raw.Control(func(fd uintptr) {
			if actual, err := sockbuf.SetRecvBuffer(int(fd), want); err == nil && actual < want {
				log.Printf("note: kernel clamped SO_RCVBUF to %d (asked %d; ceiling: kern.ipc.maxsockbuf)", actual, want)
			}
			if actual, err := sockbuf.SetSendBuffer(int(fd), want); err == nil && actual < want {
				log.Printf("note: kernel clamped SO_SNDBUF to %d (asked %d)", actual, want)
			}
		})
	}
}
