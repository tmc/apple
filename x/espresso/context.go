//go:build darwin

package espresso

import (
	"runtime"
	"sync"

	"github.com/tmc/apple/private/espresso"
)

// Context manages execution state for Espresso networks.
type Context struct {
	ctx espresso.EspressoContext

	mu     sync.Mutex
	closed bool
}

type contextConfig struct {
	platform                     Platform
	gpuPriority                  uint32
	lowPriorityMaxMsPerCmdBuffer float32
	setPriority                  bool
}

// ContextOption configures a Context.
type ContextOption func(*contextConfig)

// WithPlatform selects the execution backend.
func WithPlatform(p Platform) ContextOption {
	return func(c *contextConfig) {
		c.platform = p
	}
}

// WithGPUPriority sets the GPU command buffer priority.
func WithGPUPriority(priority uint32) ContextOption {
	return func(c *contextConfig) {
		c.gpuPriority = priority
		c.setPriority = true
	}
}

// WithLowPriority sets the maximum milliseconds per command buffer for low-priority execution.
func WithLowPriority(maxMsPerCommandBuffer float32) ContextOption {
	return func(c *contextConfig) {
		c.lowPriorityMaxMsPerCmdBuffer = maxMsPerCommandBuffer
		c.setPriority = true
	}
}

// Open creates a new Espresso execution context.
func Open(opts ...ContextOption) (*Context, error) {
	var cfg contextConfig
	for _, o := range opts {
		o(&cfg)
	}

	var ctx espresso.EspressoContext
	switch cfg.platform {
	case PlatformAuto:
		ctx = espresso.NewEspressoContext()
	default:
		ctx = espresso.NewEspressoContextWithPlatform(int(cfg.platform))
	}

	if ctx.GetID() == 0 {
		return nil, ErrFramework
	}

	if cfg.setPriority {
		ctx.Set_priorityLow_priority_max_ms_per_command_bufferGpu_priority(
			true,
			cfg.lowPriorityMaxMsPerCmdBuffer,
			cfg.gpuPriority,
		)
	}

	c := &Context{ctx: ctx}
	runtime.SetFinalizer(c, (*Context).Close)
	return c, nil
}

// Close releases context resources.
func (c *Context) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	runtime.SetFinalizer(c, nil)
	return nil
}

func (c *Context) isClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

// Platform returns the platform this context was created with.
func (c *Context) Platform() int {
	return c.ctx.Platform()
}
