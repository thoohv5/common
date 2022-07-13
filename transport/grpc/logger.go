package grpc

import (
	"context"
	"log"
	"os"
)

// ILogger is a logger interface.
type ILogger interface {
	Debugc(ctx context.Context, msg string, values ...interface{})
	Infoc(ctx context.Context, msg string, values ...interface{})
	Warnc(ctx context.Context, msg string, values ...interface{})
	Errorc(ctx context.Context, msg string, values ...interface{})
}

type defaultLogger struct {
	*log.Logger
}

func NewDefaultLogger() ILogger {
	return &defaultLogger{
		Logger: log.New(os.Stdout, "config", log.Lshortfile|log.Lmicroseconds|log.Ldate),
	}
}

func (d *defaultLogger) Debugc(ctx context.Context, msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Infoc(ctx context.Context, msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Warnc(ctx context.Context, msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Errorc(ctx context.Context, msg string, values ...interface{}) {
	d.Printf(msg, values...)
}
