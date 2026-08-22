// Command signpostweb serves HTTP and emits one os_signpost interval per
// request, making the Go server legible to Instruments.
//
// Run it, generate some load, and watch the intervals arrive:
//
//	go run . &
//	log stream --signpost --predicate 'subsystem == "com.github.tmc.apple.signpostweb"' &
//	hey -n 200 http://localhost:8077/work
//
// Or record a trace and open it in Instruments (os_signpost instrument):
//
//	xctrace record --template 'Blank' --instrument os_signpost \
//	    --launch -- ./signpostweb
//
// Signpost names are compile-time on Darwin: the log tools decode them from
// the emitting binary's __TEXT,__oslogstring section, which signpostnames
// (see go:generate below) fills by scanning this file for the literal names
// passed to handle. Names built at run time would pair but decode as
// "<missing name>", so each route is instrumented with a literal.
package main

//go:generate go run github.com/tmc/apple/x/signpost/cmd/signpostnames -o names_darwin.syso -funcs handle=1,stage=0

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/tmc/apple/x/signpost"
)

var sp = signpost.New("com.github.tmc.apple.signpostweb", "requests")

// handle registers h on mux wrapped in a signpost interval. name must be a
// string literal so signpostnames can pool it (see the package comment).
func handle(mux *http.ServeMux, name string, h http.HandlerFunc) {
	mux.HandleFunc(name, func(w http.ResponseWriter, r *http.Request) {
		id := sp.NewID()
		sp.IntervalBegin(id, name)
		defer sp.IntervalEnd(id, name)
		h(w, r)
	})
}

// stage runs fn as a named child interval nested inside the surrounding
// request interval. name must be a string literal (pooled via -funcs
// stage=0 in the go:generate line).
func stage(name string, fn func()) {
	id := sp.NewID()
	sp.IntervalBegin(id, name)
	fn()
	sp.IntervalEnd(id, name)
}

func main() {
	addr := flag.String("addr", "localhost:8077", "listen address")
	load := flag.Int("load", 0, "self-generated requests per second (0 = none)")
	flag.Parse()

	mux := http.NewServeMux()
	handle(mux, "GET /work", func(w http.ResponseWriter, r *http.Request) {
		// Simulated work with variable latency so the trace has shape.
		d := time.Duration(1+rand.IntN(20)) * time.Millisecond
		time.Sleep(d)
		fmt.Fprintf(w, "slept %v\n", d)
	})
	handle(mux, "GET /flaky", func(w http.ResponseWriter, r *http.Request) {
		// Nested stages, with a 1-in-5 slow "query" so the trace shows a
		// latency tail worth finding.
		stage("parse", func() { time.Sleep(200 * time.Microsecond) })
		stage("db", func() {
			d := 5 * time.Millisecond
			if rand.IntN(5) == 0 {
				d = 100 * time.Millisecond
				sp.Event(sp.NewID(), "slow-query")
			}
			time.Sleep(d)
		})
		stage("render", func() { fmt.Fprintln(w, "ok") })
	})
	handle(mux, "GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "signpostweb: try /work or /flaky")
	})

	if *load > 0 {
		go func() {
			paths := []string{"/work", "/work", "/flaky"}
			tick := time.NewTicker(time.Second / time.Duration(*load))
			defer tick.Stop()
			for range tick.C {
				go func(p string) {
					resp, err := http.Get("http://" + *addr + p)
					if err == nil {
						resp.Body.Close()
					}
				}(paths[rand.IntN(len(paths))])
			}
		}()
	}

	log.Printf("listening on http://%s (try /work or /flaky)", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
