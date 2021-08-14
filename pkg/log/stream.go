package log

import (
	"context"
	"fmt"
	"io"
)

// Stream implements Logger and just pass the attributes to the fmt.Fprint function
// at each level, using the defined io.Writer.
type Stream struct {
	Writer io.Writer
}

func (s Stream) Debug(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Debugf(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func (s Stream) Error(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Errorf(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func (s Stream) Fatal(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Fatalf(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func (s Stream) Info(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Infof(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func (s Stream) Panic(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Panicf(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func (s Stream) Warn(_ context.Context, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, args...)
}

func (s Stream) Warnf(_ context.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprint(s.Writer, formatMessage(format, args...))
}

func formatMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
