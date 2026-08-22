// Command midimon lists the MIDI sources on the system and prints the
// messages they send, decoding MIDI 1.0 channel voice messages.
//
// Usage: midimon [-l] [-d duration]
//
// With -l it lists the sources and exits. Otherwise it connects to every
// source and monitors them until the duration elapses or it is interrupted.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremidi"
)

// MIDIProtocolID values. The generated MIDIProtocolID enum has no constants.
const (
	protocol1_0 coremidi.MIDIProtocolID = 1
)

func main() {
	list := flag.Bool("l", false, "list sources and exit")
	dur := flag.Duration("d", time.Minute, "how long to monitor")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: midimon [-l] [-d duration]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if err := run(*list, *dur); err != nil {
		fmt.Fprintf(os.Stderr, "midimon: %v\n", err)
		os.Exit(1)
	}
}

func run(listOnly bool, dur time.Duration) error {
	n := coremidi.MIDIGetNumberOfSources()
	fmt.Printf("%d MIDI source(s), %d destination(s)\n", n, coremidi.MIDIGetNumberOfDestinations())
	for i := uint(0); i < n; i++ {
		src := coremidi.MIDIGetSource(i)
		fmt.Printf("  source %d: %s\n", i, endpointName(src))
	}
	if listOnly {
		return nil
	}
	if n == 0 {
		return fmt.Errorf("no MIDI sources found; connect a MIDI device or start a virtual source and re-run")
	}

	name := cfString("midimon")
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(name)))

	var client coremidi.MIDIClientRef
	if st := coremidi.MIDIClientCreateWithBlock(name, &client, func(*coremidi.MIDINotification) {}); st != 0 {
		return fmt.Errorf("MIDIClientCreateWithBlock: OSStatus %d", st)
	}
	defer coremidi.MIDIClientDispose(client)

	var port coremidi.MIDIPortRef
	st := coremidi.MIDIInputPortCreateWithProtocol(client, name, protocol1_0, &port, receive)
	if st != 0 {
		return fmt.Errorf("MIDIInputPortCreateWithProtocol: OSStatus %d", st)
	}
	defer coremidi.MIDIPortDispose(port)

	connected := 0
	for i := uint(0); i < n; i++ {
		src := coremidi.MIDIGetSource(i)
		if st := coremidi.MIDIPortConnectSource(port, src, nil); st != 0 {
			fmt.Fprintf(os.Stderr, "midimon: connect source %d: OSStatus %d\n", i, st)
			continue
		}
		connected++
	}
	if connected == 0 {
		return fmt.Errorf("could not connect to any source")
	}

	fmt.Printf("monitoring %d source(s) for %v; press ^C to stop\n", connected, dur)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer signal.Stop(interrupt)
	select {
	case <-interrupt:
	case <-time.After(dur):
	}
	return nil
}

// receive is the input port's MIDIReceiveBlock. CoreMIDI calls it on its own
// high-priority thread.
func receive(evtList *coremidi.MIDIEventList, _ unsafe.Pointer) {
	if evtList == nil {
		return
	}
	pkt := &evtList.Packet[0]
	for i := uint32(0); i < evtList.NumPackets; i++ {
		words := pkt.Words()
		nw := pkt.WordCount()
		if nw > uint32(len(words)) {
			nw = uint32(len(words))
		}
		for j := uint32(0); j < nw; j++ {
			if s := describe(words[j]); s != "" {
				fmt.Println(s)
			}
		}
		pkt = nextPacket(pkt, nw)
	}
}

// nextPacket returns the packet following pkt, which holds nw words. Event
// packets are variable length: an 8-byte timestamp, a 4-byte word count, then
// the words themselves.
func nextPacket(pkt *coremidi.MIDIEventPacket, nw uint32) *coremidi.MIDIEventPacket {
	off := 12 + 4*uintptr(nw)
	return (*coremidi.MIDIEventPacket)(unsafe.Add(unsafe.Pointer(pkt), off))
}

// describe renders one universal MIDI packet word. It decodes MIDI 1.0
// channel voice messages (message type 2) and reports other message types by
// their raw word.
func describe(w uint32) string {
	mt := byte(w >> 28)
	group := byte(w>>24) & 0x0f
	if mt != 2 {
		return fmt.Sprintf("group %d  type %d  raw %08x", group, mt, w)
	}
	status := byte(w>>20) & 0x0f
	ch := byte(w>>16) & 0x0f
	d1 := byte(w>>8) & 0x7f
	d2 := byte(w) & 0x7f
	switch status {
	case 0x8:
		return fmt.Sprintf("group %d ch %2d  note off  %-4s vel %3d", group, ch+1, noteName(d1), d2)
	case 0x9:
		if d2 == 0 {
			return fmt.Sprintf("group %d ch %2d  note off  %-4s vel %3d", group, ch+1, noteName(d1), d2)
		}
		return fmt.Sprintf("group %d ch %2d  note on   %-4s vel %3d", group, ch+1, noteName(d1), d2)
	case 0xa:
		return fmt.Sprintf("group %d ch %2d  aftertouch %-4s %3d", group, ch+1, noteName(d1), d2)
	case 0xb:
		return fmt.Sprintf("group %d ch %2d  control   cc %3d = %3d", group, ch+1, d1, d2)
	case 0xc:
		return fmt.Sprintf("group %d ch %2d  program   %3d", group, ch+1, d1)
	case 0xd:
		return fmt.Sprintf("group %d ch %2d  pressure  %3d", group, ch+1, d1)
	case 0xe:
		return fmt.Sprintf("group %d ch %2d  pitch bend %5d", group, ch+1, (int(d2)<<7|int(d1))-8192)
	}
	return fmt.Sprintf("group %d ch %2d  status %x  %3d %3d", group, ch+1, status, d1, d2)
}

var noteNames = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

func noteName(n byte) string {
	return fmt.Sprintf("%s%d", noteNames[n%12], int(n)/12-1)
}

// endpointName returns an endpoint's user-visible name, or a placeholder.
func endpointName(endpoint coremidi.MIDIEndpointRef) string {
	if coremidi.KMIDIPropertyDisplayName == "" {
		return "(unnamed)"
	}
	key := cfString(coremidi.KMIDIPropertyDisplayName)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(key)))
	var val corefoundation.CFStringRef
	if st := coremidi.MIDIObjectGetStringProperty(endpoint, key, &val); st != 0 || val == 0 {
		return "(unnamed)"
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(uintptr(val)))
	return goString(val)
}

// cfString creates a CFString from a Go string. The caller owns the result.
func cfString(s string) corefoundation.CFStringRef {
	return corefoundation.CFStringCreateWithCString(0, s, uint32(corefoundation.KCFStringEncodingUTF8))
}

// goString converts a CFString to a Go string.
func goString(s corefoundation.CFStringRef) string {
	size := corefoundation.CFStringGetMaximumSizeForEncoding(corefoundation.CFStringGetLength(s), uint32(corefoundation.KCFStringEncodingUTF8)) + 1
	buf := make([]byte, size)
	if !corefoundation.CFStringGetCString(s, &buf[0], size, uint32(corefoundation.KCFStringEncodingUTF8)) {
		return ""
	}
	return string(buf[:clen(buf)])
}

func clen(b []byte) int {
	for i, c := range b {
		if c == 0 {
			return i
		}
	}
	return len(b)
}
