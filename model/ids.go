// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package model

import (
	modelv1 "github.com/jaegertracing/jaeger-idl/model/v1"
)

const (
	// traceIDShortBytesLen indicates length of 64bit traceID when represented as list of bytes
	traceIDShortBytesLen = 8
)

// TraceID is a random 128bit identifier for a trace
type TraceID = modelv1.TraceID

// SpanID is a random 64bit identifier for a span
type SpanID = modelv1.SpanID

// ------- TraceID -------

// NewTraceID creates a new TraceID from two 64bit unsigned ints.
var NewTraceID = modelv1.NewTraceID

// TraceIDFromString creates a TraceID from a hexadecimal string
func TraceIDFromString(s string) (TraceID, error) {
	return modelv1.TraceIDFromString(s)
}

// TraceIDFromBytes creates a TraceID from list of bytes
func TraceIDFromBytes(data []byte) (TraceID, error) {
	return modelv1.TraceIDFromBytes(data)
}

// ------- SpanID -------

// NewSpanID creates a new SpanID from a 64bit unsigned int.
var NewSpanID = modelv1.NewSpanID

// SpanIDFromString creates a SpanID from a hexadecimal string
func SpanIDFromString(s string) (SpanID, error) {
	return modelv1.SpanIDFromString(s)
}

// SpanIDFromBytes creates a SpandID from list of bytes
func SpanIDFromBytes(data []byte) (SpanID, error) {
	return modelv1.SpanIDFromBytes(data)
}
