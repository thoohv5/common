package recovery

import (
	"context"
	"runtime"

	"github.com/thoohv5/common/errors"
	"github.com/thoohv5/common/log"
	"github.com/thoohv5/common/middleware"
)

// ErrUnknownRequest is unknown request error.
var ErrUnknownRequest = errors.InternalServer("UNKNOWN", "unknown request error")

// HandlerFunc is recovery handler func.
type HandlerFunc func(ctx context.Context, req, err interface{}) error

// Option is recovery option.
type Option func(*options)

type options struct {
	handler HandlerFunc
	logger  log.Logger
}

// WithHandler with recovery handler.
func WithHandler(h HandlerFunc) Option {
	return func(o *options) {
		o.handler = h
	}
}

// WithLogger with recovery logger.

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Recovery is a server middleware that recovers from any panics.
func Recovery(opts ...Option) middleware.Middleware {
	op := options{
		handler: func(ctx context.Context, req, err interface{}) error {
			return ErrUnknownRequest
		},
	}
	for _, o := range opts {
		o(&op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
				if rerr := recover(); rerr != nil {
					buf := make([]byte, 64<<10) //nolint:gomnd
					n := runtime.Stack(buf, false)
					buf = buf[:n]
					op.logger.Errorc(ctx, "异常崩溃", func(fs log.IFields) {
						fs.Set("错误", rerr)
						fs.Set("请求\n", req)
						fs.Set("堆栈\n", buf)
					})

					err = op.handler(ctx, req, rerr)
				}
			}()
			return handler(ctx, req)
		}
	}
}
