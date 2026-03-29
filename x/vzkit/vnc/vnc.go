package vnc

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

// SecurityMode selects the security configuration for the VNC server.
type SecurityMode string

const (
	// SecurityNone disables authentication.
	SecurityNone SecurityMode = "none"
	// SecurityPassword enables password authentication.
	SecurityPassword SecurityMode = "password"
)

// Config controls VNC server creation.
type Config struct {
	Port        uint16
	ServiceName string
	Queue       *vm.Queue
	Mode        SecurityMode
	Password    string
}

// SecurityConfig wraps a private VNC security configuration.
type SecurityConfig struct {
	raw pvz.VZVNCSecurityConfiguration
}

// Object returns the Objective-C object for the security configuration.
func (c SecurityConfig) Object() objectivec.IObject {
	return c.raw
}

// NewNoSecurityConfig creates a configuration that allows unauthenticated access.
func NewNoSecurityConfig() SecurityConfig {
	cfg := pvz.NewVZVNCNoSecuritySecurityConfiguration()
	return SecurityConfig{raw: pvz.VZVNCSecurityConfigurationFromID(cfg.ID)}
}

// NewPasswordSecurityConfig creates a password-protected VNC security configuration.
func NewPasswordSecurityConfig(password string) (SecurityConfig, error) {
	if password == "" {
		return SecurityConfig{}, fmt.Errorf("password required")
	}
	pass := foundation.NewStringWithString(password)
	cfg := pvz.NewVZVNCAuthenticationSecurityConfigurationWithPassword(pass)
	if cfg.ID == 0 {
		return SecurityConfig{}, fmt.Errorf("create vnc password security configuration")
	}
	return SecurityConfig{raw: pvz.VZVNCSecurityConfigurationFromID(cfg.ID)}, nil
}

// NewSecurityConfig builds a security configuration from Config.
func NewSecurityConfig(cfg Config) (SecurityConfig, error) {
	switch cfg.Mode {
	case "", SecurityNone:
		return NewNoSecurityConfig(), nil
	case SecurityPassword:
		return NewPasswordSecurityConfig(cfg.Password)
	default:
		return SecurityConfig{}, fmt.Errorf("unknown security mode: %s", cfg.Mode)
	}
}

// Server wraps a private VNC server.
type Server struct {
	raw   pvz.VZVNCServer
	queue *vm.Queue
}

// Raw returns the underlying private VNC server.
func (s *Server) Raw() pvz.VZVNCServer {
	return s.raw
}

// New constructs a VNC server using the provided configuration.
func New(cfg Config) (*Server, error) {
	if cfg.Port == 0 {
		cfg.Port = 5900
	}

	var sec SecurityConfig
	var err error
	if cfg.Mode != "" && cfg.Mode != SecurityNone {
		sec, err = NewSecurityConfig(cfg)
		if err != nil {
			return nil, err
		}
	}

	var server pvz.VZVNCServer
	queue := cfg.Queue
	if queue != nil {
		q := queue.Queue()
		switch {
		case cfg.ServiceName != "" && sec.raw.GetID() != 0:
			server = pvz.NewVZVNCServerWithBonjourServiceNameQueueSecurityConfiguration(
				foundation.NewStringWithString(cfg.ServiceName), q, sec.Object())
		case cfg.ServiceName != "":
			server = pvz.NewVZVNCServerWithBonjourServiceNameQueue(
				foundation.NewStringWithString(cfg.ServiceName), q)
		case sec.raw.GetID() != 0:
			server = pvz.NewVZVNCServerWithPortQueueSecurityConfiguration(cfg.Port, q, sec.Object())
		default:
			server = pvz.NewVZVNCServerWithPortQueue(cfg.Port, q)
		}
	} else {
		switch {
		case cfg.ServiceName != "" && sec.raw.GetID() != 0:
			server = pvz.NewVZVNCServerWithBonjourServiceNameQueueSecurityConfiguration(
				foundation.NewStringWithString(cfg.ServiceName), objectivec.Object{}, sec.Object())
		case cfg.ServiceName != "":
			server = pvz.NewVZVNCServerWithBonjourServiceName(
				foundation.NewStringWithString(cfg.ServiceName))
		case sec.raw.GetID() != 0:
			server = pvz.NewVZVNCServerWithPortQueueSecurityConfiguration(cfg.Port, objectivec.Object{}, sec.Object())
		default:
			server = pvz.NewVZVNCServerWithPort(cfg.Port)
		}
	}
	if server.ID == 0 {
		return nil, fmt.Errorf("create vnc server")
	}
	return &Server{raw: server, queue: cfg.Queue}, nil
}

func (s *Server) sync(fn func()) {
	if s.queue == nil {
		fn()
		return
	}
	s.queue.Sync(fn)
}

// Start starts the VNC server.
func (s *Server) Start() {
	s.sync(func() { s.raw.Start() })
}

// Stop stops the VNC server.
func (s *Server) Stop() {
	s.sync(func() { s.raw.Stop() })
}

// SetVirtualMachine attaches a VM to the server.
func (s *Server) SetVirtualMachine(vm pvz.IVZVirtualMachine) {
	s.sync(func() { s.raw.SetVirtualMachine(vm) })
}

// SetGraphicsDisplay attaches a graphics display to the server.
func (s *Server) SetGraphicsDisplay(display pvz.IVZGraphicsDisplay) {
	s.sync(func() { s.raw.SetGraphicsDisplay(display) })
}

// Port returns the configured VNC port.
func (s *Server) Port() uint16 {
	var port uint16
	s.sync(func() { port = s.raw.Port() })
	return port
}

// State returns the VNC server state.
func (s *Server) State() int64 {
	var state int64
	s.sync(func() { state = s.raw.State() })
	return state
}

// SecurityConfiguration returns the private security configuration.
func (s *Server) SecurityConfiguration() SecurityConfig {
	var cfg SecurityConfig
	s.sync(func() {
		raw := s.raw.SecurityConfiguration()
		if raw != nil {
			cfg = SecurityConfig{raw: *raw}
		}
	})
	return cfg
}

// Description returns the Objective-C description string.
func (s *Server) Description() string {
	var desc string
	s.sync(func() { desc = s.raw.Description() })
	return desc
}

// DebugDescription returns the Objective-C debug description string.
func (s *Server) DebugDescription() string {
	var desc string
	s.sync(func() { desc = s.raw.DebugDescription() })
	return desc
}
