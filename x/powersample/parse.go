package powersample

import (
	"strconv"
	"strings"
	"time"
)

// Power holds one estimate per SoC rail, in the caller's unit of the
// moment: watts in a Sample, joules in Report.Energy.
type Power struct {
	CPU      float64
	GPU      float64
	ANE      float64
	Combined float64
}

// Sample is one powermetrics sample window.
type Sample struct {
	Elapsed time.Duration
	Power   Power // watts
}

// parser consumes powermetrics text output line by line and emits one
// Sample per "*** Sampled system activity ... (Nms elapsed) ***" block.
// Rail lines can repeat within a block (the gpu_power and ane_power
// samplers restate what cpu_power already printed); the first occurrence
// wins so nothing is double-counted.
type parser struct {
	open    bool
	elapsed time.Duration
	seen    map[string]bool
	cur     Power
}

func (p *parser) line(s string) (Sample, bool) {
	if strings.HasPrefix(s, "*** Sampled system activity") {
		out, ok := p.flush()
		p.open = true
		p.elapsed = parseElapsed(s)
		p.seen = map[string]bool{}
		p.cur = Power{}
		return out, ok
	}
	if !p.open {
		return Sample{}, false
	}
	for _, r := range []struct {
		prefix string
		field  *float64
	}{
		{"CPU Power:", &p.cur.CPU},
		{"GPU Power:", &p.cur.GPU},
		{"ANE Power:", &p.cur.ANE},
		{"Combined Power (CPU + GPU + ANE):", &p.cur.Combined},
	} {
		if rest, ok := strings.CutPrefix(s, r.prefix); ok && !p.seen[r.prefix] {
			if w, ok := parseWatts(rest); ok {
				p.seen[r.prefix] = true
				*r.field = w
			}
			return Sample{}, false
		}
	}
	return Sample{}, false
}

// flush closes the current block, returning its sample if one was open
// and carried a usable elapsed time.
func (p *parser) flush() (Sample, bool) {
	if !p.open || p.elapsed <= 0 {
		p.open = false
		return Sample{}, false
	}
	p.open = false
	return Sample{Elapsed: p.elapsed, Power: p.cur}, true
}

// parseElapsed extracts the sample duration from a header like
// "*** Sampled system activity (Tue Aug 12 ...) (502.45ms elapsed) ***".
func parseElapsed(s string) time.Duration {
	i := strings.LastIndex(s, "ms elapsed)")
	if i < 0 {
		return 0
	}
	j := strings.LastIndex(s[:i], "(")
	if j < 0 {
		return 0
	}
	ms, err := strconv.ParseFloat(strings.TrimSpace(s[j+1:i]), 64)
	if err != nil {
		return 0
	}
	return time.Duration(ms * float64(time.Millisecond))
}

// parseWatts parses " 458 mW" or " 4.5 W" into watts.
func parseWatts(s string) (float64, bool) {
	fields := strings.Fields(s)
	if len(fields) != 2 {
		return 0, false
	}
	v, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, false
	}
	switch fields[1] {
	case "mW":
		return v / 1000, true
	case "W":
		return v, true
	}
	return 0, false
}
