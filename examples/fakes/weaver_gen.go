// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package fakes

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/fakes/Clock",
		Iface: reflect.TypeOf((*Clock)(nil)).Elem(),
		Impl:  reflect.TypeOf(clock{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return clock_local_stub{impl: impl.(Clock), tracer: tracer, unixMicroMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/fakes/Clock", Method: "UnixMicro", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return clock_client_stub{stub: stub, unixMicroMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/fakes/Clock", Method: "UnixMicro", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return clock_server_stub{impl: impl.(Clock), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return clock_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Clock] = (*clock)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*clock)(nil)

// Local stub implementations.

type clock_local_stub struct {
	impl             Clock
	tracer           trace.Tracer
	unixMicroMetrics *codegen.MethodMetrics
}

// Check that clock_local_stub implements the Clock interface.
var _ Clock = (*clock_local_stub)(nil)

func (s clock_local_stub) UnixMicro(ctx context.Context) (r0 int64, err error) {
	// Update metrics.
	begin := s.unixMicroMetrics.Begin()
	defer func() { s.unixMicroMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "fakes.Clock.UnixMicro", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.UnixMicro(ctx)
}

// Client stub implementations.

type clock_client_stub struct {
	stub             codegen.Stub
	unixMicroMetrics *codegen.MethodMetrics
}

// Check that clock_client_stub implements the Clock interface.
var _ Clock = (*clock_client_stub)(nil)

func (s clock_client_stub) UnixMicro(ctx context.Context) (r0 int64, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.unixMicroMetrics.Begin()
	defer func() { s.unixMicroMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "fakes.Clock.UnixMicro", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int64()
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type clock_server_stub struct {
	impl    Clock
	addLoad func(key uint64, load float64)
}

// Check that clock_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*clock_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s clock_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "UnixMicro":
		return s.unixMicro
	default:
		return nil
	}
}

func (s clock_server_stub) unixMicro(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.UnixMicro(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type clock_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that clock_reflect_stub implements the Clock interface.
var _ Clock = (*clock_reflect_stub)(nil)

func (s clock_reflect_stub) UnixMicro(ctx context.Context) (r0 int64, err error) {
	err = s.caller("UnixMicro", ctx, []any{}, []any{&r0})
	return
}
