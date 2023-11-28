//go:build notracing

package internal

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func StartSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return noop.NewTracerProvider().Tracer("go-libp2p-kad-dht").Start(ctx, fmt.Sprintf("KademliaDHT.%s", name), opts...)
}

// KeyAsAttribute format a DHT key into a suitable tracing attribute.
// DHT keys can be either valid utf-8 or binary, when they are derived from, for example, a multihash.
// Tracing (and notably OpenTelemetry+grpc exporter) requires valid utf-8 for string attributes.
func KeyAsAttribute(name string, key string) attribute.KeyValue { return attribute.KeyValue{} }
