// Command blescan scans for nearby Bluetooth Low Energy peripherals and
// prints each advertisement as it arrives.
//
// Usage:
//
//	blescan [flags]
//
// Scanning requires a Bluetooth radio that is powered on, and macOS prompts
// for Bluetooth permission the first time a process scans. blescan stops after
// the scan duration elapses, or when interrupted.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/tmc/apple/corebluetooth"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// advert is one advertisement observed by the delegate, copied out of
// Objective-C land so the main goroutine can print it safely.
type advert struct {
	id        string
	name      string
	rssi      int
	services  []string
	connect   bool
	txPower   string
	timestamp time.Time
}

func main() {
	dur := flag.Duration("d", 10*time.Second, "how long to scan")
	nameFilter := flag.String("name", "", "only report peripherals whose name contains this substring")
	svcFilter := flag.String("service", "", "only scan for this service UUID (e.g. 180D for heart rate)")
	uniq := flag.Bool("uniq", true, "report each peripheral once instead of on every advertisement")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: blescan [flags]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if err := run(*dur, *nameFilter, *svcFilter, *uniq); err != nil {
		fmt.Fprintf(os.Stderr, "blescan: %v\n", err)
		os.Exit(1)
	}
}

func run(dur time.Duration, nameFilter, svcFilter string, uniq bool) error {
	var services []corebluetooth.CBUUID
	if svcFilter != "" {
		services = append(services, corebluetooth.NewCBUUIDWithString(svcFilter))
	}

	// Delegate callbacks are delivered on this serial queue, not on a Go
	// goroutine, so everything they produce is handed to main over channels.
	queue := dispatch.QueueCreate("blescan.delegate")

	adverts := make(chan advert, 64)
	states := make(chan corebluetooth.CBManagerState, 8)

	delegate := corebluetooth.NewCBCentralManagerDelegate(corebluetooth.CBCentralManagerDelegateConfig{
		CentralManagerDidUpdateState: func(c corebluetooth.CBCentralManager) {
			select {
			case states <- c.State():
			default:
			}
		},
		CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI: func(c corebluetooth.CBCentralManager, p corebluetooth.CBPeripheral, data foundation.INSDictionary, rssi foundation.NSNumber) {
			select {
			case adverts <- makeAdvert(p, data, rssi):
			default: // drop rather than block the delegate queue
			}
		},
	})
	central := corebluetooth.NewCBCentralManagerWithDelegateQueue(delegate, queue)

	// The first state callback tells us whether scanning is possible at all.
	select {
	case st := <-states:
		if st != corebluetooth.CBManagerStatePoweredOn {
			return fmt.Errorf("bluetooth unavailable: %s", stateString(st))
		}
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timed out waiting for bluetooth state; is Bluetooth permission granted to this binary?")
	}

	central.ScanForPeripheralsWithServicesOptions(services, nil)
	if !central.IsScanning() {
		return fmt.Errorf("central manager did not start scanning")
	}
	defer central.StopScan()

	fmt.Fprintf(os.Stderr, "scanning for %v...\n", dur)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)

	deadline := time.After(dur)
	seen := make(map[string]bool)
	count := 0

	for {
		select {
		case a := <-adverts:
			if nameFilter != "" && !strings.Contains(strings.ToLower(a.name), strings.ToLower(nameFilter)) {
				continue
			}
			if uniq {
				if seen[a.id] {
					continue
				}
				seen[a.id] = true
			}
			count++
			printAdvert(a)
		case st := <-states:
			if st != corebluetooth.CBManagerStatePoweredOn {
				return fmt.Errorf("bluetooth became unavailable: %s", stateString(st))
			}
		case <-sig:
			fmt.Fprintf(os.Stderr, "interrupted; %d advertisement(s)\n", count)
			return nil
		case <-deadline:
			if count == 0 {
				fmt.Fprintf(os.Stderr, "no peripherals found; move closer to a BLE device and re-run\n")
			} else {
				fmt.Fprintf(os.Stderr, "%d advertisement(s)\n", count)
			}
			return nil
		}
	}
}

func makeAdvert(p corebluetooth.CBPeripheral, data foundation.INSDictionary, rssi foundation.NSNumber) advert {
	a := advert{
		id:        p.Identifier().UUIDString(),
		name:      p.Name(),
		rssi:      int(rssi.IntegerValue()),
		timestamp: time.Now(),
	}
	if data == nil {
		return a
	}
	dict := foundation.NSDictionaryFromID(data.GetID())
	if a.name == "" {
		if v := lookup(dict, corebluetooth.CBAdvertisementDataLocalNameKey); v != 0 {
			a.name = foundation.NSStringFromID(v).UTF8String()
		}
	}
	if v := lookup(dict, corebluetooth.CBAdvertisementDataIsConnectable); v != 0 {
		a.connect = foundation.NSNumberFromID(v).BoolValue()
	}
	if v := lookup(dict, corebluetooth.CBAdvertisementDataTxPowerLevelKey); v != 0 {
		a.txPower = fmt.Sprintf("%d dBm", foundation.NSNumberFromID(v).IntegerValue())
	}
	if v := lookup(dict, corebluetooth.CBAdvertisementDataServiceUUIDsKey); v != 0 {
		arr := foundation.NSArrayFromID(v)
		for i := uint(0); i < arr.Count(); i++ {
			u := corebluetooth.CBUUIDFromID(arr.ObjectAtIndex(i).GetID())
			a.services = append(a.services, u.UUIDString())
		}
		sort.Strings(a.services)
	}
	return a
}

// lookup returns the object stored under key, or 0 if absent. An empty key
// means the framework global was not resolved, which is treated as absent.
func lookup(dict foundation.NSDictionary, key string) objc.ID {
	if key == "" {
		return 0
	}
	v := dict.ObjectForKey(foundation.NewStringWithString(key))
	return v.GetID()
}

func printAdvert(a advert) {
	name := a.name
	if name == "" {
		name = "(unnamed)"
	}
	fmt.Printf("%s  %-28s rssi=%4d", a.timestamp.Format("15:04:05"), name, a.rssi)
	if a.connect {
		fmt.Printf(" connectable")
	}
	if a.txPower != "" {
		fmt.Printf(" tx=%s", a.txPower)
	}
	if len(a.services) > 0 {
		fmt.Printf(" services=[%s]", strings.Join(a.services, " "))
	}
	fmt.Printf("\n    %s\n", a.id)
}

func stateString(st corebluetooth.CBManagerState) string {
	switch st {
	case corebluetooth.CBManagerStatePoweredOff:
		return "powered off; turn Bluetooth on and re-run"
	case corebluetooth.CBManagerStateUnauthorized:
		return "unauthorized; grant Bluetooth access in System Settings > Privacy & Security"
	case corebluetooth.CBManagerStateUnsupported:
		return "unsupported; this machine has no Bluetooth low energy radio"
	case corebluetooth.CBManagerStateResetting:
		return "resetting; try again in a moment"
	case corebluetooth.CBManagerStateUnknown:
		return "unknown"
	}
	return st.String()
}
