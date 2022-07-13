package log

import "context"

// Logger is a logger interface.
type Logger interface {
	Debugc(ctx context.Context, msg string, values ...interface{})
	Infoc(ctx context.Context, msg string, values ...interface{})
	Warnc(ctx context.Context, msg string, values ...interface{})
	Errorc(ctx context.Context, msg string, values ...interface{})

	Debugf(msg string, values ...interface{})
	Infof(msg string, values ...interface{})
	Warnf(msg string, values ...interface{})
	Errorf(msg string, values ...interface{})
}
