package main

import (
	"context"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
	"io"
	"log"
	"os"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

var logger *slog.Logger
var appEnv = os.Getenv("APP_ENV")

func main() {
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	logger = slog.New(handler)
	slog.Info("Hello, world!")
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

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

	//b, err := json.MarshalIndent(fields, "", "  ")
	//if err != nil {
	//	return err
	//}

	timeStr := r.Time.Format("[01-02 15:05:05.000]")
	msg := color.CyanString(r.Message)

	//h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))
	h.l.Println(timeStr, level, msg)

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
