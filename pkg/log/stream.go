package log

import (
	"context"
	"fmt"
	"io"
)

// Stream implements Logger and just pass the attributes to the fmt.Fprint function
// at each level, using the defined io.Writer.
type Stream struct {
	Fields map[string]interface{}
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

func (s Stream) WithField(f string, v interface{}) Logger {
	logger := s
	logger.Fields[f] = v

	return logger
}

func (s Stream) WithFields(fields map[string]interface{}) Logger {
	logger := s
	for f, v := range fields {
		logger.Fields[f] = v
	}

	return logger
}

func formatMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
