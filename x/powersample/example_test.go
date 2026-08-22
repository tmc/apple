//go:build darwin

package powersample_test

import (
	"fmt"
	"log"
	"time"

	"github.com/tmc/apple/x/powersample"
)

func Example() {
	m, err := powersample.Start(500 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	work()
	r, err := m.Stop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gpu %.2f J, ane %.2f J over %s (%d samples)\n",
		r.Energy.GPU, r.Energy.ANE, r.Duration, r.Samples)
}

func work() { time.Sleep(10 * time.Millisecond) }
