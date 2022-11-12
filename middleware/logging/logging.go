package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/thoohv5/common/errors"
	"github.com/thoohv5/common/log"
	"github.com/thoohv5/common/middleware"
	"github.com/thoohv5/common/transport"
)

// Server is an server logging middleware.
func Server(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			logger.Debugc(ctx, "middleware",
				func(fs log.IFields) {
					fs.Set("kind", "server")
					fs.Set("component", kind)
					fs.Set("operation", operation)
					fs.Set("args", extractArgs(req))
					fs.Set("code", code)
					fs.Set("reason", reason)
					fs.Set("stack", fmt.Sprintf("%+v", err))
					fs.Set("latency", time.Since(startTime).Seconds())
				},
			)
			return
		}
	}
}

// Client is an client logging middleware.
func Client(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			logger.Debugc(ctx, "middleware",
				func(fs log.IFields) {
					fs.Set("kind", "client")
					fs.Set("component", kind)
					fs.Set("operation", operation)
					fs.Set("args", extractArgs(req))
					fs.Set("code", code)
					fs.Set("reason", reason)
					fs.Set("stack", fmt.Sprintf("%+v", err))
					fs.Set("latency", time.Since(startTime).Seconds())
				},
			)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}
