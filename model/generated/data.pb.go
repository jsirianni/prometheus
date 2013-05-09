// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

package io_prometheus

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import google_protobuf "google/protobuf/descriptor.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LabelPair struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}

func (m *LabelPair) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type LabelName struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LabelName) Reset()         { *m = LabelName{} }
func (m *LabelName) String() string { return proto.CompactTextString(m) }
func (*LabelName) ProtoMessage()    {}

func (m *LabelName) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type Metric struct {
	LabelPair        []*LabelPair `protobuf:"bytes,1,rep,name=label_pair" json:"label_pair,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}

func (m *Metric) GetLabelPair() []*LabelPair {
	if m != nil {
		return m.LabelPair
	}
	return nil
}

type Fingerprint struct {
	Signature        *string `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Fingerprint) Reset()         { *m = Fingerprint{} }
func (m *Fingerprint) String() string { return proto.CompactTextString(m) }
func (*Fingerprint) ProtoMessage()    {}

func (m *Fingerprint) GetSignature() string {
	if m != nil && m.Signature != nil {
		return *m.Signature
	}
	return ""
}

type FingerprintCollection struct {
	Member           []*Fingerprint `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FingerprintCollection) Reset()         { *m = FingerprintCollection{} }
func (m *FingerprintCollection) String() string { return proto.CompactTextString(m) }
func (*FingerprintCollection) ProtoMessage()    {}

func (m *FingerprintCollection) GetMember() []*Fingerprint {
	if m != nil {
		return m.Member
	}
	return nil
}

type LabelSet struct {
	Member           []*LabelPair `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *LabelSet) Reset()         { *m = LabelSet{} }
func (m *LabelSet) String() string { return proto.CompactTextString(m) }
func (*LabelSet) ProtoMessage()    {}

func (m *LabelSet) GetMember() []*LabelPair {
	if m != nil {
		return m.Member
	}
	return nil
}

type SampleKey struct {
	Fingerprint      *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint" json:"fingerprint,omitempty"`
	Timestamp        []byte       `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	LastTimestamp    *int64       `protobuf:"fixed64,3,opt,name=last_timestamp" json:"last_timestamp,omitempty"`
	SampleCount      *uint32      `protobuf:"fixed32,4,opt,name=sample_count" json:"sample_count,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SampleKey) Reset()         { *m = SampleKey{} }
func (m *SampleKey) String() string { return proto.CompactTextString(m) }
func (*SampleKey) ProtoMessage()    {}

func (m *SampleKey) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *SampleKey) GetTimestamp() []byte {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SampleKey) GetLastTimestamp() int64 {
	if m != nil && m.LastTimestamp != nil {
		return *m.LastTimestamp
	}
	return 0
}

func (m *SampleKey) GetSampleCount() uint32 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

type SampleValueSeries struct {
	Value            []*SampleValueSeries_Value `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SampleValueSeries) Reset()         { *m = SampleValueSeries{} }
func (m *SampleValueSeries) String() string { return proto.CompactTextString(m) }
func (*SampleValueSeries) ProtoMessage()    {}

func (m *SampleValueSeries) GetValue() []*SampleValueSeries_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type SampleValueSeries_Value struct {
	Timestamp        *int64   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SampleValueSeries_Value) Reset()         { *m = SampleValueSeries_Value{} }
func (m *SampleValueSeries_Value) String() string { return proto.CompactTextString(m) }
func (*SampleValueSeries_Value) ProtoMessage()    {}

func (m *SampleValueSeries_Value) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *SampleValueSeries_Value) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type MembershipIndexValue struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MembershipIndexValue) Reset()         { *m = MembershipIndexValue{} }
func (m *MembershipIndexValue) String() string { return proto.CompactTextString(m) }
func (*MembershipIndexValue) ProtoMessage()    {}

type MetricHighWatermark struct {
	Timestamp        *int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MetricHighWatermark) Reset()         { *m = MetricHighWatermark{} }
func (m *MetricHighWatermark) String() string { return proto.CompactTextString(m) }
func (*MetricHighWatermark) ProtoMessage()    {}

func (m *MetricHighWatermark) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type CompactionProcessorDefinition struct {
	MinimumGroupSize *uint32 `protobuf:"varint,1,opt,name=minimum_group_size" json:"minimum_group_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CompactionProcessorDefinition) Reset()         { *m = CompactionProcessorDefinition{} }
func (m *CompactionProcessorDefinition) String() string { return proto.CompactTextString(m) }
func (*CompactionProcessorDefinition) ProtoMessage()    {}

func (m *CompactionProcessorDefinition) GetMinimumGroupSize() uint32 {
	if m != nil && m.MinimumGroupSize != nil {
		return *m.MinimumGroupSize
	}
	return 0
}

type CurationKey struct {
	Fingerprint              *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint" json:"fingerprint,omitempty"`
	ProcessorMessageTypeName *string      `protobuf:"bytes,2,opt,name=processor_message_type_name" json:"processor_message_type_name,omitempty"`
	ProcessorMessageRaw      []byte       `protobuf:"bytes,3,opt,name=processor_message_raw" json:"processor_message_raw,omitempty"`
	IgnoreYoungerThan        *int64       `protobuf:"varint,4,opt,name=ignore_younger_than" json:"ignore_younger_than,omitempty"`
	XXX_unrecognized         []byte       `json:"-"`
}

func (m *CurationKey) Reset()         { *m = CurationKey{} }
func (m *CurationKey) String() string { return proto.CompactTextString(m) }
func (*CurationKey) ProtoMessage()    {}

func (m *CurationKey) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *CurationKey) GetProcessorMessageTypeName() string {
	if m != nil && m.ProcessorMessageTypeName != nil {
		return *m.ProcessorMessageTypeName
	}
	return ""
}

func (m *CurationKey) GetProcessorMessageRaw() []byte {
	if m != nil {
		return m.ProcessorMessageRaw
	}
	return nil
}

func (m *CurationKey) GetIgnoreYoungerThan() int64 {
	if m != nil && m.IgnoreYoungerThan != nil {
		return *m.IgnoreYoungerThan
	}
	return 0
}

type CurationValue struct {
	LastCompletionTimestamp *int64 `protobuf:"varint,1,opt,name=last_completion_timestamp" json:"last_completion_timestamp,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *CurationValue) Reset()         { *m = CurationValue{} }
func (m *CurationValue) String() string { return proto.CompactTextString(m) }
func (*CurationValue) ProtoMessage()    {}

func (m *CurationValue) GetLastCompletionTimestamp() int64 {
	if m != nil && m.LastCompletionTimestamp != nil {
		return *m.LastCompletionTimestamp
	}
	return 0
}

func init() {
}
