//go:build darwin

package powersample

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const tool = "/usr/bin/powermetrics"

// Report is the integrated result of one metering run.
type Report struct {
	Duration time.Duration // sum of sample windows actually measured
	Samples  int           // sample windows parsed; 0 means the format changed or the run was too short
	Energy   Power         // joules per rail
	Average  Power         // watts per rail (Energy / Duration)
}

// Meter is a running energy measurement.
type Meter struct {
	ior *iorMeter
	pm  *pmMeter
}

// Start begins measuring. It reads the SoC energy counters through
// IOReport (private API, no privileges needed); where that is
// unavailable it falls back to running powermetrics, which requires
// root — the fallback's error message contains the exact grant command
// when the process is not root and passwordless sudo is unavailable.
func Start(interval time.Duration) (*Meter, error) {
	if m, err := startIOReport(interval); err == nil {
		return &Meter{ior: m}, nil
	}
	pm, err := startPowermetrics(interval)
	if err != nil {
		return nil, err
	}
	return &Meter{pm: pm}, nil
}

// Stop ends the measurement and integrates the samples. The Meter
// cannot be reused.
func (m *Meter) Stop() (*Report, error) {
	if m.ior != nil {
		return m.ior.stop()
	}
	return m.pm.stop()
}

// pmMeter is a running powermetrics process accumulating samples.
type pmMeter struct {
	cmd     *exec.Cmd
	stderr  *bytes.Buffer
	done    chan struct{}
	samples []Sample
}

func startPowermetrics(interval time.Duration) (*pmMeter, error) {
	args := []string{
		"--samplers", "cpu_power,gpu_power,ane_power",
		"-i", fmt.Sprint(interval.Milliseconds()),
		"-a", "0", // no poweravg summary lines
		"-b", "1", // line-buffered output
	}
	var cmd *exec.Cmd
	if os.Geteuid() == 0 {
		cmd = exec.Command(tool, args...)
	} else {
		if err := exec.Command("sudo", "-n", "-v").Run(); err != nil {
			return nil, fmt.Errorf("powersample: powermetrics needs root and sudo -n failed; run the program under sudo, or grant it alone: echo \"$USER ALL=(root) NOPASSWD: %s\" | sudo tee /etc/sudoers.d/powermetrics", tool)
		}
		cmd = exec.Command("sudo", append([]string{"-n", tool}, args...)...)
	}
	m := &pmMeter{cmd: cmd, stderr: new(bytes.Buffer), done: make(chan struct{})}
	cmd.Stderr = m.stderr
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("powersample: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("powersample: start powermetrics: %w", err)
	}
	go func() {
		defer close(m.done)
		var p parser
		sc := bufio.NewScanner(out)
		sc.Buffer(make([]byte, 64<<10), 1<<20)
		for sc.Scan() {
			if s, ok := p.line(sc.Text()); ok {
				m.samples = append(m.samples, s)
			}
		}
		if s, ok := p.flush(); ok {
			m.samples = append(m.samples, s)
		}
	}()
	return m, nil
}

// stop terminates powermetrics, waits for its output to drain, and
// integrates the samples.
func (m *pmMeter) stop() (*Report, error) {
	// powermetrics exits cleanly on SIGINT; sudo relays the signal.
	m.cmd.Process.Signal(syscall.SIGINT)
	err := m.cmd.Wait()
	<-m.done
	if err != nil && len(m.samples) == 0 {
		return nil, fmt.Errorf("powersample: powermetrics: %w: %s", err, bytes.TrimSpace(m.stderr.Bytes()))
	}
	var r Report
	for _, s := range m.samples {
		sec := s.Elapsed.Seconds()
		r.Duration += s.Elapsed
		r.Energy.CPU += s.Power.CPU * sec
		r.Energy.GPU += s.Power.GPU * sec
		r.Energy.ANE += s.Power.ANE * sec
		r.Energy.Combined += s.Power.Combined * sec
	}
	r.Samples = len(m.samples)
	if sec := r.Duration.Seconds(); sec > 0 {
		r.Average = Power{
			CPU:      r.Energy.CPU / sec,
			GPU:      r.Energy.GPU / sec,
			ANE:      r.Energy.ANE / sec,
			Combined: r.Energy.Combined / sec,
		}
	}
	return &r, nil
}
