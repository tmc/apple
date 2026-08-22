// Command listen classifies the sounds in an audio file using the
// SoundAnalysis framework's built-in classifier, printing each result's
// time range and its top labels with confidence.
//
// Usage: listen [-n top] [-timeout d] <audio-file>
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/soundanalysis"
)

func main() {
	top := flag.Int("n", 3, "number of top classifications to print per result")
	timeout := flag.Duration("timeout", 2*time.Minute, "maximum time to wait for analysis")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: listen [-n top] [-timeout d] <audio-file>\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(flag.Arg(0), *top, *timeout); err != nil {
		fmt.Fprintf(os.Stderr, "listen: %v\n", err)
		os.Exit(1)
	}
}

func run(path string, top int, timeout time.Duration) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	if soundanalysis.SNClassifierIdentifierVersion1 == "" {
		return fmt.Errorf("SoundAnalysis framework unavailable: no built-in classifier identifier")
	}

	url := foundation.NewURLFileURLWithPath(path)
	analyzer, err := soundanalysis.NewAudioFileAnalyzerWithURLError(url)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	request, err := soundanalysis.NewClassifySoundRequestWithClassifierIdentifierError(soundanalysis.SNClassifierIdentifierVersion1)
	if err != nil {
		return fmt.Errorf("create classify request: %w", err)
	}

	failed := make(chan error, 1)
	results := 0
	observer := newObserver(top, failed, &results)
	if ok, err := analyzer.AddRequestWithObserverError(request, observer); err != nil {
		return fmt.Errorf("add request: %w", err)
	} else if !ok {
		return fmt.Errorf("add request: refused by analyzer")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ok, err := analyzer.AnalyzeSync(ctx)
	if err != nil {
		analyzer.CancelAnalysis()
		return fmt.Errorf("analyze: %w", err)
	}
	select {
	case err := <-failed:
		return err
	default:
	}
	if !ok {
		return fmt.Errorf("analysis did not complete")
	}
	if results == 0 {
		return fmt.Errorf("no classifications produced; the file may be shorter than the classifier's analysis window (about 3s)")
	}
	return nil
}

// newObserver registers an Objective-C class conforming to SNResultsObserving
// and returns an instance of it. The SoundAnalysis bindings generate no
// constructor for this protocol, so the class is built by hand here.
func newObserver(top int, failed chan<- error, results *int) soundanalysis.SNResultsObservingObject {
	methods := []objc.MethodDef{{
		Cmd: objc.RegisterName("request:didProduceResult:"),
		Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID, resultID objc.ID) {
			*results++
			printResult(soundanalysis.SNClassificationResultFromID(resultID), top)
		},
	}, {
		Cmd: objc.RegisterName("request:didFailWithError:"),
		Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID, errorID objc.ID) {
			err := fmt.Errorf("request failed: %s", objc.IDToString(objc.Send[objc.ID](errorID, objc.Sel("localizedDescription"))))
			select {
			case failed <- err:
			default:
			}
		},
	}, {
		Cmd: objc.RegisterName("requestDidComplete:"),
		Fn: func(self objc.ID, _cmd objc.SEL, requestID objc.ID) {},
	}}

	var protocols []*objc.Protocol
	if p := objc.GetProtocol("SNResultsObserving"); p != nil {
		protocols = append(protocols, p)
	}
	name := fmt.Sprintf("GoSNResultsObserver_%d", os.Getpid())
	cls, err := objc.RegisterClass(name, objc.GetClass("NSObject"), protocols, nil, methods)
	if err != nil {
		fmt.Fprintf(os.Stderr, "listen: register observer class: %v\n", err)
		os.Exit(1)
	}
	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return soundanalysis.SNResultsObservingObjectFromID(instance)
}

func printResult(result soundanalysis.SNClassificationResult, top int) {
	timeRange := result.TimeRange()
	start := coremedia.CMTimeGetSeconds(timeRange.Start())
	duration := coremedia.CMTimeGetSeconds(timeRange.Duration())

	classifications := result.Classifications()
	sort.SliceStable(classifications, func(i, j int) bool {
		return classifications[i].Confidence() > classifications[j].Confidence()
	})
	if len(classifications) > top {
		classifications = classifications[:top]
	}
	for i, c := range classifications {
		if i == 0 {
			fmt.Printf("%7.2fs +%.2fs  ", start, duration)
		} else {
			fmt.Printf("%18s", "")
		}
		fmt.Printf("%-24s %.3f\n", c.Identifier(), c.Confidence())
	}
}
