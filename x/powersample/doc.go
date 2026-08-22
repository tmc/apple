// Package powersample measures SoC subsystem energy on Apple silicon.
//
// A Meter measures for the duration of a region and integrates the CPU,
// GPU, and ANE energy into joules:
//
//	m, err := powersample.Start(500 * time.Millisecond)
//	if err != nil {
//		log.Fatal(err) // the error says how to grant access
//	}
//	work()
//	r, err := m.Stop()
//	fmt.Printf("GPU: %.2f J over %s\n", r.Energy.GPU, r.Duration)
//
// Start needs no privileges: it reads the SoC's cumulative energy
// counters (the "Energy Model" group in the private libIOReport.dylib —
// the same source powermetrics aggregates) directly. Where IOReport is
// unavailable it falls back to running powermetrics, which requires
// root: the fallback runs it directly when the process is root,
// otherwise through sudo -n, and when neither works the returned error
// contains the exact command that grants passwordless access to
// powermetrics alone. TestLivePowermetricsAgreement cross-checks the
// two backends against each other when run with root.
//
// The numbers are Apple's own estimates and are documented as unsuitable
// for comparing devices; they are for comparing configurations on one
// machine. Both backends are version-brittle in their own way — the
// IOReport channel names ("CPU Energy", "GPU Energy", "ANE") and the
// powermetrics text format, whose parser is pinned by hand-authored
// testdata matching the output shape observed on macOS 26.6.1 (25G76),
// not a captured live run. Report.Samples lets a caller detect a silent
// change in either (zero samples over a nonzero duration), and Stop
// returns an error when no energy was seen at all. TestLiveMeter
// exercises the IOReport path unprivileged.
package powersample
