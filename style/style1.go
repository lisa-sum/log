package style

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
	"io"
	"log"
)

type PrettyHandlerOptions1 struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler1 struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle1(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	// 颜色
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[22:15:5:5.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler2(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
