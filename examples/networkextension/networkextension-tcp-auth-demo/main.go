// Command networkextension-tcp-auth-demo exercises NetworkExtension TCP
// authentication delegate callbacks without creating a connection.
package main

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/tmc/apple/networkextension"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
	"github.com/tmc/apple/security"
)

type delegateCounts struct {
	shouldEvaluate uint64
	evaluate       uint64
	shouldProvide  uint64
}

func registerDelegateClass(name string, invoker *objcbridge.BlockInvoker, counts *delegateCounts) (objc.Class, error) {
	super := objc.GetClass("NSObject")
	if super == 0 {
		return 0, errors.New("lookup NSObject")
	}

	protocols := objcbridge.ProtocolsByName("NWTCPConnectionAuthenticationDelegate")
	cls, err := objc.RegisterClass(name, super, protocols, nil, nil)
	if err != nil {
		return 0, fmt.Errorf("register %s: %w", name, err)
	}

	methods := []objc.MethodDef{
		{
			Cmd: objc.RegisterName("shouldEvaluateTrustForConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connection objc.ID) bool {
				atomic.AddUint64(&counts.shouldEvaluate, 1)
				return true
			},
		},
		{
			Cmd: objc.RegisterName("evaluateTrustForConnection:peerCertificateChain:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connection objc.ID, chain objc.ID, completion objc.ID) {
				atomic.AddUint64(&counts.evaluate, 1)
				if err := invoker.Object(completion, 0); err != nil {
					panic(err)
				}
			},
		},
		{
			Cmd: objc.RegisterName("shouldProvideIdentityForConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connection objc.ID) bool {
				atomic.AddUint64(&counts.shouldProvide, 1)
				return false
			},
		},
	}
	if err := objcbridge.AddMethods(cls, name, methods); err != nil {
		return 0, err
	}
	return cls, nil
}

func runSmoke() error {
	var counts delegateCounts
	name := fmt.Sprintf("GoNWTCPAuthDelegateSmoke_%d_%d", os.Getpid(), time.Now().UnixNano())
	cls, err := registerDelegateClass(name, objcbridge.NewBlockInvoker(), &counts)
	if err != nil {
		return err
	}

	id := objc.Send[objc.ID](objc.ID(cls), objc.Sel("new"))
	if id == 0 {
		return fmt.Errorf("instantiate %s", name)
	}
	defer objc.Send[struct{}](id, objc.Sel("release"))

	if !objc.RespondsToSelector(id, objc.Sel("shouldEvaluateTrustForConnection:")) {
		return errors.New("delegate does not respond to trust decision selector")
	}
	if !objc.RespondsToSelector(id, objc.Sel("evaluateTrustForConnection:peerCertificateChain:completionHandler:")) {
		return errors.New("delegate does not respond to trust evaluation selector")
	}

	if ok := objc.Send[bool](id, objc.Sel("shouldEvaluateTrustForConnection:"), objc.ID(0)); !ok {
		return errors.New("trust decision returned false")
	}

	var completed uint64
	block, cleanup := networkextension.NewSecTrustRefBlock(func(trust security.SecTrustRef) {
		if trust != 0 {
			panic("unexpected trust object")
		}
		atomic.AddUint64(&completed, 1)
	})
	defer cleanup()
	objc.Send[struct{}](id, objc.Sel("evaluateTrustForConnection:peerCertificateChain:completionHandler:"), objc.ID(0), objc.ID(0), block)

	if ok := objc.Send[bool](id, objc.Sel("shouldProvideIdentityForConnection:"), objc.ID(0)); ok {
		return errors.New("identity decision returned true")
	}

	if got := atomic.LoadUint64(&counts.shouldEvaluate); got != 1 {
		return fmt.Errorf("trust decision callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.evaluate); got != 1 {
		return fmt.Errorf("trust evaluation callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&counts.shouldProvide); got != 1 {
		return fmt.Errorf("identity decision callbacks = %d, want 1", got)
	}
	if got := atomic.LoadUint64(&completed); got != 1 {
		return fmt.Errorf("completion callbacks = %d, want 1", got)
	}
	return nil
}

func main() {
	if err := runSmoke(); err != nil {
		fmt.Fprintf(os.Stderr, "networkextension-tcp-auth-demo: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("networkextension-tcp-auth-demo: NetworkExtension Objective-C helper smoke ok")
}
