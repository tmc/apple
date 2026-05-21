// Command networkextension-provider-callback-demo exercises provider-style
// NetworkExtension completion callbacks without installing an extension.
package main

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

type providerCounts struct {
	start   uint64
	stop    uint64
	message uint64
	sleep   uint64
	wake    uint64
}

func registerProviderClass(name string, invoker *objcbridge.BlockInvoker, counts *providerCounts) (objc.Class, error) {
	super := objc.GetClass("NSObject")
	if super == 0 {
		return 0, errors.New("lookup NSObject")
	}

	cls, err := objc.RegisterClass(name, super, nil, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("register %s: %w", name, err)
	}

	methods := []objc.MethodDef{
		{
			Cmd: objc.RegisterName("startTunnelWithOptions:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, options objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.start, 1)
				if err := invoker.Error(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("stopTunnelWithReason:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, reason networkextension.NEProviderStopReason, completion objc.ID) {
				atomic.AddUint64(&counts.stop, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("handleAppMessage:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, message objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.message, 1)
				if err := invoker.Object(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("sleepWithCompletionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, completion objc.ID) {
				atomic.AddUint64(&counts.sleep, 1)
				if err := invoker.Void(completion); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("wake"),
			Fn: func(self objc.ID, _cmd objc.SEL) {
				atomic.AddUint64(&counts.wake, 1)
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func runSmoke() error {
	var counts providerCounts
	name := fmt.Sprintf("GoPacketTunnelProviderSmoke_%d_%d", os.Getpid(), time.Now().UnixNano())
	cls, err := registerProviderClass(name, objcbridge.NewBlockInvoker(), &counts)
	if err != nil {
		return err
	}

	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("new"))
	if id == 0 {
		return fmt.Errorf("instantiate %s", name)
	}
	defer objc.Send[struct{}](id, objc.Sel("release"))

	for _, sel := range []string{
		"startTunnelWithOptions:completionHandler:",
		"stopTunnelWithReason:completionHandler:",
		"handleAppMessage:completionHandler:",
		"sleepWithCompletionHandler:",
		"wake",
	} {
		if !objc.RespondsToSelector(id, objc.Sel(sel)) {
			return fmt.Errorf("provider does not respond to %s", sel)
		}
	}

	var startDone uint64
	startBlock, startCleanup := networkextension.NewErrorBlock(func(err error) {
		if err != nil {
			panic(err)
		}
		atomic.AddUint64(&startDone, 1)
	})
	defer startCleanup()
	objc.Send[struct{}](id, objc.Sel("startTunnelWithOptions:completionHandler:"), objc.ID(0), startBlock)

	var stopDone uint64
	stopBlock, stopCleanup := networkextension.NewVoidBlock(func() {
		atomic.AddUint64(&stopDone, 1)
	})
	defer stopCleanup()
	objc.Send[struct{}](id, objc.Sel("stopTunnelWithReason:completionHandler:"), networkextension.NEProviderStopReasonUserInitiated, stopBlock)

	var messageDone uint64
	messageBlock, messageCleanup := networkextension.NewDataBlock(func(data *foundation.NSData) {
		if data != nil {
			panic("unexpected response data")
		}
		atomic.AddUint64(&messageDone, 1)
	})
	defer messageCleanup()
	objc.Send[struct{}](id, objc.Sel("handleAppMessage:completionHandler:"), objc.ID(0), messageBlock)

	var sleepDone uint64
	sleepBlock, sleepCleanup := networkextension.NewVoidBlock(func() {
		atomic.AddUint64(&sleepDone, 1)
	})
	defer sleepCleanup()
	objc.Send[struct{}](id, objc.Sel("sleepWithCompletionHandler:"), sleepBlock)
	objc.Send[struct{}](id, objc.Sel("wake"))

	if got := atomic.LoadUint64(&counts.start); got != 1 {
		return fmt.Errorf("start callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&startDone); got != 1 {
		return fmt.Errorf("start completions = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.stop); got != 1 {
		return fmt.Errorf("stop callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&stopDone); got != 1 {
		return fmt.Errorf("stop completions = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.message); got != 1 {
		return fmt.Errorf("message callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&messageDone); got != 1 {
		return fmt.Errorf("message completions = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.sleep); got != 1 {
		return fmt.Errorf("sleep callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&sleepDone); got != 1 {
		return fmt.Errorf("sleep completions = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.wake); got != 1 {
		return fmt.Errorf("wake callbacks = %d, want 1", got)
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-provider-callback-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-provider-callback-demo: NetworkExtension provider callback smoke ok")
}
