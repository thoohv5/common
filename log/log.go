package log

import "context"

type IFields interface {
	Set(key string, val interface{}) IFields
	Data() map[string]interface{}
}
type Field func(IFields)

// Logger is a logger interface.
type Logger interface {
	Debugc(ctx context.Context, msg string, fields ...Field)
	Infoc(ctx context.Context, msg string, fields ...Field)
	Warnc(ctx context.Context, msg string, fields ...Field)
	Errorc(ctx context.Context, msg string, fields ...Field)

	Debugf(ctx context.Context, msg string, values ...interface{})
	Infof(ctx context.Context, msg string, values ...interface{})
	Warnf(ctx context.Context, msg string, values ...interface{})
	Errorf(ctx context.Context, msg string, values ...interface{})
}
