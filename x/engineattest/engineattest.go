//go:build darwin

package engineattest

import "errors"

// ErrDidNotRun reports that the attested function completed without the
// engine's counter advancing: the claimed hardware never executed.
var ErrDidNotRun = errors.New("engine did not execute during the attested region")

// ErrUnattestable reports that the counter cannot move for this model or
// queue, so a zero reading proves nothing either way. For the ANE this
// means the sensitivity canary evaluation produced no hardware execution
// time (for example a shared in-memory MIL model, which exposes no
// hardware counters).
var ErrUnattestable = errors.New("hardware counter cannot move; claim is unattestable")
