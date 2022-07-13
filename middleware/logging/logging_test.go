package logging

import (
	"bytes"
	"context"
	"errors"
	slog "log"
	"os"
	"testing"

	"github.com/thoohv5/common/log"
	"github.com/thoohv5/common/middleware"
	"github.com/thoohv5/common/transport"
)

var _ transport.Transporter = &Transport{}

type Transport struct {
	kind      transport.Kind
	endpoint  string
	operation string
}

func (tr *Transport) Kind() transport.Kind {
	return tr.kind
}

func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

func (tr *Transport) Operation() string {
	return tr.operation
}

func (tr *Transport) RequestHeader() transport.Header {
	return nil
}

func (tr *Transport) ReplyHeader() transport.Header {
	return nil
}

func TestHTTP(t *testing.T) {
	err := errors.New("reply.error")
	bf := bytes.NewBuffer(nil)
	logger := NewDefaultLogger()

	tests := []struct {
		name string
		kind func(logger log.Logger) middleware.Middleware
		err  error
		ctx  context.Context
	}{
		{
			"http-server@fail",
			Server,
			err,
			func() context.Context {
				return transport.NewServerContext(context.Background(), &Transport{kind: transport.KindHTTP, endpoint: "endpoint", operation: "/package.service/method"})
			}(),
		},
		{
			"http-server@succ",
			Server,
			nil,
			func() context.Context {
				return transport.NewServerContext(context.Background(), &Transport{kind: transport.KindHTTP, endpoint: "endpoint", operation: "/package.service/method"})
			}(),
		},
		{
			"http-client@succ",
			Client,
			nil,
			func() context.Context {
				return transport.NewClientContext(context.Background(), &Transport{kind: transport.KindHTTP, endpoint: "endpoint", operation: "/package.service/method"})
			}(),
		},
		{
			"http-client@fail",
			Client,
			err,
			func() context.Context {
				return transport.NewClientContext(context.Background(), &Transport{kind: transport.KindHTTP, endpoint: "endpoint", operation: "/package.service/method"})
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bf.Reset()
			next := func(ctx context.Context, req interface{}) (interface{}, error) {
				return "reply", test.err
			}
			next = test.kind(logger)(next)
			v, e := next(test.ctx, "req.args")
			t.Logf("[%s]reply: %v, error: %v", test.name, v, e)
			t.Logf("[%s]log:%s", test.name, bf.String())
		})
	}
}

type defaultLogger struct {
	*slog.Logger
}

func NewDefaultLogger() log.Logger {
	return &defaultLogger{
		Logger: slog.New(os.Stdout, "config", slog.Lshortfile|slog.Lmicroseconds|slog.Ldate),
	}
}

func (d *defaultLogger) Debugf(msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Infof(msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Warnf(msg string, values ...interface{}) {
	d.Printf(msg, values...)
}

func (d *defaultLogger) Errorf(msg string, values ...interface{}) {
	d.Printf(msg, values...)
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
