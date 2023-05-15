// go:build !ignoreWeaverGen

package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)
var _ codegen.LatestVersion = codegen.Version[[0][10]struct{}]("You used 'weaver generate' codegen version 0.10.0, but you built your code with an incompatible weaver module version. Try upgrading 'weaver generate' and re-running it.")

func init() {
	codegen.Register(codegen.Registration{
		Name:   "github.com/ServiceWeaver/weaver/examples/factors/Factorer",
		Iface:  reflect.TypeOf((*Factorer)(nil)).Elem(),
		Impl:   reflect.TypeOf(factorer{}),
		Routed: true,
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return factorer_local_stub{impl: impl.(Factorer), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return factorer_client_stub{stub: stub, factorsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/factors/Factorer", Method: "Factors"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return factorer_server_stub{impl: impl.(Factorer), addLoad: addLoad}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/Main",
		Iface: reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:  reflect.TypeOf(server{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		RefData: "⟦4724da9b:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/examples/factors/Factorer⟧\n",
	})
}

// weaver.Instance checks.
var _ weaver.InstanceOf[Factorer] = &factorer{}
var _ weaver.InstanceOf[weaver.Main] = &server{}

// Local stub implementations.

type factorer_local_stub struct {
	impl   Factorer
	tracer trace.Tracer
}

// Check that factorer_local_stub implements the Factorer interface.
var _ Factorer = &factorer_local_stub{}

func (s factorer_local_stub) Factors(ctx context.Context, a0 int) (r0 []int, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Factorer.Factors", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Factors(ctx, a0)
}

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = &main_local_stub{}

// Client stub implementations.

type factorer_client_stub struct {
	stub           codegen.Stub
	factorsMetrics *codegen.MethodMetrics
}

// Check that factorer_client_stub implements the Factorer interface.
var _ Factorer = &factorer_client_stub{}

func (s factorer_client_stub) Factors(ctx context.Context, a0 int) (r0 []int, err error) {
	// Update metrics.
	start := time.Now()
	s.factorsMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Factorer.Factors", trace.WithSpanKind(trace.SpanKindClient))
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
			s.factorsMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.factorsMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)

	// Set the shardKey.
	var r router
	shardKey := _hashFactorer(r.Factors(ctx, a0))

	// Call the remote method.
	s.factorsMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.factorsMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_int_7c8c8866(dec)
	err = dec.Error()
	return
}

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = &main_client_stub{}

// Server stub implementations.

type factorer_server_stub struct {
	impl    Factorer
	addLoad func(key uint64, load float64)
}

// Check that factorer_server_stub implements the codegen.Server interface.
var _ codegen.Server = &factorer_server_stub{}

// GetStubFn implements the codegen.Server interface.
func (s factorer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Factors":
		return s.factors
	default:
		return nil
	}
}

func (s factorer_server_stub) factors(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()
	var r router
	s.addLoad(_hashFactorer(r.Factors(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Factors(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_int_7c8c8866(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = &main_server_stub{}

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Router methods.

// _hashFactorer returns a 64 bit hash of the provided value.
func _hashFactorer(r int) uint64 {
	var h codegen.Hasher
	h.WriteInt(int(r))
	return h.Sum64()
}

// _orderedCodeFactorer returns an order-preserving serialization of the provided value.
func _orderedCodeFactorer(r int) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteInt(int(r))
	return enc.Encode()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_int_7c8c8866(enc *codegen.Encoder, arg []int) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Int(arg[i])
	}
}

func serviceweaver_dec_slice_int_7c8c8866(dec *codegen.Decoder) []int {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Int()
	}
	return res
}
