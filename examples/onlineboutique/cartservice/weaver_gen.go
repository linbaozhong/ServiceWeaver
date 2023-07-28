// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package cartservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

var _ codegen.LatestVersion = codegen.Version[[0][17]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.17.0 (codegen
version v0.17.0). The generated code is incompatible with the version of the
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

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, addItemMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "AddItem", Remote: false}), emptyCartMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "EmptyCart", Remote: false}), getCartMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "GetCart", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, addItemMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "AddItem", Remote: true}), emptyCartMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "EmptyCart", Remote: true}), getCartMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T", Method: "GetCart", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		RefData: "⟦e78910e9:wEaVeReDgE:github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/T→github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:   "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache",
		Iface:  reflect.TypeOf((*cartCache)(nil)).Elem(),
		Impl:   reflect.TypeOf(cartCacheImpl{}),
		Routed: true,
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return cartCache_local_stub{impl: impl.(cartCache), tracer: tracer, addMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Add", Remote: false}), getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Get", Remote: false}), removeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Remove", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return cartCache_client_stub{stub: stub, addMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Add", Remote: true}), getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Get", Remote: true}), removeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/cartservice/cartCache", Method: "Remove", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return cartCache_server_stub{impl: impl.(cartCache), addLoad: addLoad}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)
var _ weaver.InstanceOf[cartCache] = (*cartCacheImpl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)
var _ weaver.RoutedBy[cartCacheRouter] = (*cartCacheImpl)(nil)

// Component "cartCacheImpl", router "cartCacheRouter" checks.
var _ func(_ context.Context, key string, value []CartItem) string = (&cartCacheRouter{}).Add // routed
var _ func(_ context.Context, key string) string = (&cartCacheRouter{}).Get                   // routed
var _ func(_ context.Context, key string) string = (&cartCacheRouter{}).Remove                // routed

// Local stub implementations.

type t_local_stub struct {
	impl             T
	tracer           trace.Tracer
	addItemMetrics   *codegen.MethodMetrics
	emptyCartMetrics *codegen.MethodMetrics
	getCartMetrics   *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) AddItem(ctx context.Context, a0 string, a1 CartItem) (err error) {
	// Update metrics.
	begin := s.addItemMetrics.Begin()
	defer func() { s.addItemMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.T.AddItem", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AddItem(ctx, a0, a1)
}

func (s t_local_stub) EmptyCart(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	begin := s.emptyCartMetrics.Begin()
	defer func() { s.emptyCartMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.T.EmptyCart", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.EmptyCart(ctx, a0)
}

func (s t_local_stub) GetCart(ctx context.Context, a0 string) (r0 []CartItem, err error) {
	// Update metrics.
	begin := s.getCartMetrics.Begin()
	defer func() { s.getCartMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.T.GetCart", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetCart(ctx, a0)
}

type cartCache_local_stub struct {
	impl          cartCache
	tracer        trace.Tracer
	addMetrics    *codegen.MethodMetrics
	getMetrics    *codegen.MethodMetrics
	removeMetrics *codegen.MethodMetrics
}

// Check that cartCache_local_stub implements the cartCache interface.
var _ cartCache = (*cartCache_local_stub)(nil)

func (s cartCache_local_stub) Add(ctx context.Context, a0 string, a1 []CartItem) (err error) {
	// Update metrics.
	begin := s.addMetrics.Begin()
	defer func() { s.addMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.cartCache.Add", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Add(ctx, a0, a1)
}

func (s cartCache_local_stub) Get(ctx context.Context, a0 string) (r0 []CartItem, err error) {
	// Update metrics.
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.cartCache.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx, a0)
}

func (s cartCache_local_stub) Remove(ctx context.Context, a0 string) (r0 bool, err error) {
	// Update metrics.
	begin := s.removeMetrics.Begin()
	defer func() { s.removeMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "cartservice.cartCache.Remove", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Remove(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub             codegen.Stub
	addItemMetrics   *codegen.MethodMetrics
	emptyCartMetrics *codegen.MethodMetrics
	getCartMetrics   *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) AddItem(ctx context.Context, a0 string, a1 CartItem) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.addItemMetrics.Begin()
	defer func() { s.addItemMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.T.AddItem", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += serviceweaver_size_CartItem_e3591e56(&a1)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	(a1).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) EmptyCart(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.emptyCartMetrics.Begin()
	defer func() { s.emptyCartMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.T.EmptyCart", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s t_client_stub) GetCart(ctx context.Context, a0 string) (r0 []CartItem, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getCartMetrics.Begin()
	defer func() { s.getCartMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.T.GetCart", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_CartItem_7a7ff11c(dec)
	err = dec.Error()
	return
}

type cartCache_client_stub struct {
	stub          codegen.Stub
	addMetrics    *codegen.MethodMetrics
	getMetrics    *codegen.MethodMetrics
	removeMetrics *codegen.MethodMetrics
}

// Check that cartCache_client_stub implements the cartCache interface.
var _ cartCache = (*cartCache_client_stub)(nil)

func (s cartCache_client_stub) Add(ctx context.Context, a0 string, a1 []CartItem) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.addMetrics.Begin()
	defer func() { s.addMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.cartCache.Add", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.String(a0)
	serviceweaver_enc_slice_CartItem_7a7ff11c(enc, a1)

	// Set the shardKey.
	var r cartCacheRouter
	shardKey := _hashCartCache(r.Add(ctx, a0, a1))

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s cartCache_client_stub) Get(ctx context.Context, a0 string) (r0 []CartItem, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.cartCache.Get", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)

	// Set the shardKey.
	var r cartCacheRouter
	shardKey := _hashCartCache(r.Get(ctx, a0))

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_CartItem_7a7ff11c(dec)
	err = dec.Error()
	return
}

func (s cartCache_client_stub) Remove(ctx context.Context, a0 string) (r0 bool, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.removeMetrics.Begin()
	defer func() { s.removeMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "cartservice.cartCache.Remove", trace.WithSpanKind(trace.SpanKindClient))
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

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)

	// Set the shardKey.
	var r cartCacheRouter
	shardKey := _hashCartCache(r.Remove(ctx, a0))

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "AddItem":
		return s.addItem
	case "EmptyCart":
		return s.emptyCart
	case "GetCart":
		return s.getCart
	default:
		return nil
	}
}

func (s t_server_stub) addItem(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 CartItem
	(&a1).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.AddItem(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) emptyCart(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.EmptyCart(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) getCart(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetCart(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_CartItem_7a7ff11c(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type cartCache_server_stub struct {
	impl    cartCache
	addLoad func(key uint64, load float64)
}

// Check that cartCache_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*cartCache_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s cartCache_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Add":
		return s.add
	case "Get":
		return s.get
	case "Remove":
		return s.remove
	default:
		return nil
	}
}

func (s cartCache_server_stub) add(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 []CartItem
	a1 = serviceweaver_dec_slice_CartItem_7a7ff11c(dec)
	var r cartCacheRouter
	s.addLoad(_hashCartCache(r.Add(ctx, a0, a1)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Add(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s cartCache_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var r cartCacheRouter
	s.addLoad(_hashCartCache(r.Get(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_CartItem_7a7ff11c(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s cartCache_server_stub) remove(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var r cartCacheRouter
	s.addLoad(_hashCartCache(r.Remove(ctx, a0)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Remove(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*CartItem)(nil)

type __is_CartItem[T ~struct {
	weaver.AutoMarshal
	ProductID string
	Quantity  int32
}] struct{}

var _ __is_CartItem[CartItem]

func (x *CartItem) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("CartItem.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ProductID)
	enc.Int32(x.Quantity)
}

func (x *CartItem) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("CartItem.WeaverUnmarshal: nil receiver"))
	}
	x.ProductID = dec.String()
	x.Quantity = dec.Int32()
}

// Router methods.

// _hashCartCache returns a 64 bit hash of the provided value.
func _hashCartCache(r string) uint64 {
	var h codegen.Hasher
	h.WriteString(string(r))
	return h.Sum64()
}

// _orderedCodeCartCache returns an order-preserving serialization of the provided value.
func _orderedCodeCartCache(r string) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteString(string(r))
	return enc.Encode()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_CartItem_7a7ff11c(enc *codegen.Encoder, arg []CartItem) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_CartItem_7a7ff11c(dec *codegen.Decoder) []CartItem {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]CartItem, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

// Size implementations.

// serviceweaver_size_CartItem_e3591e56 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_CartItem_e3591e56(x *CartItem) int {
	size := 0
	size += 0
	size += (4 + len(x.ProductID))
	size += 4
	return size
}
