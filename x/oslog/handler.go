package oslog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

// Handler is a [slog.Handler] that writes records to the unified logging system
// through a [Logger]. Levels map to os_log types; the record message and its
// attributes are rendered into the log message text, since os_log has no named
// fields. The subsystem and category come from the Logger.
//
// Use it with slog:
//
//	log := slog.New(oslog.NewHandler(oslog.New("com.example.app", "worker"), nil))
//	log.Info("job done", "id", 42, "dur", "1.2s")
type Handler struct {
	logger *Logger
	opts   HandlerOptions
	groups []string // groups currently open, for qualifying Record attrs
	// preformatted holds attributes added via WithAttrs, already rendered with
	// the group prefix that was active when they were added. Storing the text
	// (rather than the raw attrs) preserves slog's point-in-time group nesting.
	preformatted string
}

// HandlerOptions configure a [Handler].
type HandlerOptions struct {
	// Level reports the minimum record level to log. If nil, slog.LevelInfo.
	Level slog.Leveler
	// Private redacts attribute values in the system log (they appear as
	// <private> unless the reader is entitled to see them). By default a
	// Handler logs its message and attributes publicly, since slog attributes
	// are developer-chosen structured fields meant to be read back. Set Private
	// for a logger that may carry sensitive data.
	Private bool
}

// NewHandler returns a [Handler] writing to logger. A nil opts is treated as
// the zero HandlerOptions.
func NewHandler(logger *Logger, opts *HandlerOptions) *Handler {
	h := &Handler{logger: logger}
	if opts != nil {
		h.opts = *opts
	}
	return h
}

// Enabled reports whether a record at level would be logged.
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	min := slog.LevelInfo
	if h.opts.Level != nil {
		min = h.opts.Level.Level()
	}
	if level < min {
		return false
	}
	return h.logger.Enabled(osType(level))
}

// Handle renders the record and writes it at the mapped os_log type.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	var b strings.Builder
	b.WriteString(r.Message)
	b.WriteString(h.preformatted)
	r.Attrs(func(a slog.Attr) bool {
		writeAttr(&b, h.groups, a)
		return true
	})

	// The rendered line is a single %s argument so no os_log format specifier in
	// the user's data is interpreted. It is public by default, redacted when the
	// handler is configured Private.
	if h.opts.Private {
		h.logger.Log(osType(r.Level), "%{private}s", b.String())
	} else {
		h.logger.Log(osType(r.Level), "%{public}s", b.String())
	}
	return nil
}

// WithAttrs returns a Handler whose records carry the given attributes. The
// attributes are rendered now, under the currently open groups, so later
// WithGroup calls do not retroactively re-qualify them.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	nh := h.clone()
	var b strings.Builder
	b.WriteString(nh.preformatted)
	for _, a := range attrs {
		writeAttr(&b, nh.groups, a)
	}
	nh.preformatted = b.String()
	return nh
}

// WithGroup returns a Handler that qualifies subsequent attribute keys with name.
func (h *Handler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	nh := h.clone()
	nh.groups = append(nh.groups, name)
	return nh
}

func (h *Handler) clone() *Handler {
	return &Handler{
		logger:       h.logger,
		opts:         h.opts,
		groups:       append([]string(nil), h.groups...),
		preformatted: h.preformatted,
	}
}

func writeAttr(b *strings.Builder, groups []string, a slog.Attr) {
	a.Value = a.Value.Resolve()
	if a.Value.Kind() == slog.KindGroup {
		g := a.Value.Group()
		if len(g) == 0 {
			return
		}
		ng := groups
		if a.Key != "" {
			ng = append(append([]string(nil), groups...), a.Key)
		}
		for _, ga := range g {
			writeAttr(b, ng, ga)
		}
		return
	}
	b.WriteByte(' ')
	for _, g := range groups {
		b.WriteString(g)
		b.WriteByte('.')
	}
	b.WriteString(a.Key)
	b.WriteByte('=')
	fmt.Fprint(b, a.Value.Any())
}

// osType maps an slog level to the closest os_log type.
func osType(level slog.Level) Type {
	switch {
	case level >= slog.LevelError:
		return TypeError
	case level >= slog.LevelWarn:
		return TypeDefault // os_log has no warning; Default is the nearest
	case level >= slog.LevelInfo:
		return TypeInfo
	default:
		return TypeDebug
	}
}
