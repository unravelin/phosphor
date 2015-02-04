// Code generated by protoc-gen-go.
// source: github.com/mattheath/phosphor/proto/trace.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/mattheath/phosphor/proto/trace.proto

It has these top-level messages:
	TraceFrame
	KeyValue
*/
package proto

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FrameType int32

const (
	FrameType_UNKNOWN    FrameType = 0
	FrameType_REQ        FrameType = 1
	FrameType_REP        FrameType = 2
	FrameType_IN         FrameType = 3
	FrameType_OUT        FrameType = 4
	FrameType_TIMEOUT    FrameType = 5
	FrameType_ANNOTATION FrameType = 6
)

var FrameType_name = map[int32]string{
	0: "UNKNOWN",
	1: "REQ",
	2: "REP",
	3: "IN",
	4: "OUT",
	5: "TIMEOUT",
	6: "ANNOTATION",
}
var FrameType_value = map[string]int32{
	"UNKNOWN":    0,
	"REQ":        1,
	"REP":        2,
	"IN":         3,
	"OUT":        4,
	"TIMEOUT":    5,
	"ANNOTATION": 6,
}

func (x FrameType) Enum() *FrameType {
	p := new(FrameType)
	*p = x
	return p
}
func (x FrameType) String() string {
	return proto.EnumName(FrameType_name, int32(x))
}
func (x *FrameType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FrameType_value, data, "FrameType")
	if err != nil {
		return err
	}
	*x = FrameType(value)
	return nil
}

type TraceFrame struct {
	// The ID of the trace this frame is a component of
	TraceId *string `protobuf:"bytes,1,req,name=trace_id" json:"trace_id,omitempty"`
	// The span this trace corresponds to, in the case this
	// is representing a service (REQ/REP) call
	SpanId *string `protobuf:"bytes,2,opt,name=span_id" json:"span_id,omitempty"`
	// The parent span this trace corresponds to, allowing us
	// to correlate trace frames and reconstruct the request
	ParentId *string `protobuf:"bytes,3,opt,name=parent_id" json:"parent_id,omitempty"`
	// The type of trace frame we're capturing
	Type *FrameType `protobuf:"varint,4,req,name=type,enum=proto.FrameType" json:"type,omitempty"`
	// Time since the epoch in nanoseconds
	//
	// 64bit time in nanoseconds with the default unix epoch gives us the following range:
	// Min time: 1677-09-21 00:12:43.145224192 +0000 UTC
	// Max time: 2262-04-11 23:47:16.854775807 +0000 UTC
	Timestamp *int64 `protobuf:"varint,5,req,name=timestamp" json:"timestamp,omitempty"`
	// Duration in nanoseconds
	// This should only be used to measure time on the same node
	// eg. the duration of service / rpc calls
	Duration *int64 `protobuf:"varint,6,opt,name=duration" json:"duration,omitempty"`
	// Machine hostname, container name etc
	Hostname *string `protobuf:"bytes,7,opt,name=hostname" json:"hostname,omitempty"`
	// Origin of this frame, likely a service
	Origin *string `protobuf:"bytes,8,opt,name=origin" json:"origin,omitempty"`
	// Destination of this frame's action
	// eg. the service which a request was destined for
	// likely not set for annotations
	Destination *string `protobuf:"bytes,9,opt,name=destination" json:"destination,omitempty"`
	// Payload as a string - eg. JSON encoded
	Payload *string `protobuf:"bytes,10,opt,name=payload" json:"payload,omitempty"`
	// Repeated series of key value fields for arbitrary data
	KeyValue         []*KeyValue `protobuf:"bytes,11,rep,name=key_value" json:"key_value,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TraceFrame) Reset()         { *m = TraceFrame{} }
func (m *TraceFrame) String() string { return proto.CompactTextString(m) }
func (*TraceFrame) ProtoMessage()    {}

func (m *TraceFrame) GetTraceId() string {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return ""
}

func (m *TraceFrame) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

func (m *TraceFrame) GetParentId() string {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return ""
}

func (m *TraceFrame) GetType() FrameType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return FrameType_UNKNOWN
}

func (m *TraceFrame) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TraceFrame) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *TraceFrame) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *TraceFrame) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *TraceFrame) GetDestination() string {
	if m != nil && m.Destination != nil {
		return *m.Destination
	}
	return ""
}

func (m *TraceFrame) GetPayload() string {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return ""
}

func (m *TraceFrame) GetKeyValue() []*KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type KeyValue struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.FrameType", FrameType_name, FrameType_value)
}
