// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/treasury.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FundCollectedEvent struct {
	JobID       string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	BlockHeight uint64 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *FundCollectedEvent) Reset()         { *m = FundCollectedEvent{} }
func (m *FundCollectedEvent) String() string { return proto.CompactTextString(m) }
func (*FundCollectedEvent) ProtoMessage()    {}
func (*FundCollectedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_30eac93b35d65e1b, []int{0}
}
func (m *FundCollectedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundCollectedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundCollectedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundCollectedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundCollectedEvent.Merge(m, src)
}
func (m *FundCollectedEvent) XXX_Size() int {
	return m.Size()
}
func (m *FundCollectedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_FundCollectedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_FundCollectedEvent proto.InternalMessageInfo

func (m *FundCollectedEvent) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *FundCollectedEvent) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundCollectedEvent) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *FundCollectedEvent) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type CollectFunds struct {
	FromBlockTime time.Time `protobuf:"bytes,1,opt,name=fromBlockTime,proto3,stdtime" json:"fromBlockTime"`
	ToBlockTime   time.Time `protobuf:"bytes,2,opt,name=toBlockTime,proto3,stdtime" json:"toBlockTime"`
}

func (m *CollectFunds) Reset()         { *m = CollectFunds{} }
func (m *CollectFunds) String() string { return proto.CompactTextString(m) }
func (*CollectFunds) ProtoMessage()    {}
func (*CollectFunds) Descriptor() ([]byte, []int) {
	return fileDescriptor_30eac93b35d65e1b, []int{1}
}
func (m *CollectFunds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectFunds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectFunds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectFunds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectFunds.Merge(m, src)
}
func (m *CollectFunds) XXX_Size() int {
	return m.Size()
}
func (m *CollectFunds) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectFunds.DiscardUnknown(m)
}

var xxx_messageInfo_CollectFunds proto.InternalMessageInfo

func (m *CollectFunds) GetFromBlockTime() time.Time {
	if m != nil {
		return m.FromBlockTime
	}
	return time.Time{}
}

func (m *CollectFunds) GetToBlockTime() time.Time {
	if m != nil {
		return m.ToBlockTime
	}
	return time.Time{}
}

type FundCollectedEvidence struct {
	Evidence []*FundCollectedEvidence_Data `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence,omitempty"`
}

func (m *FundCollectedEvidence) Reset()         { *m = FundCollectedEvidence{} }
func (m *FundCollectedEvidence) String() string { return proto.CompactTextString(m) }
func (*FundCollectedEvidence) ProtoMessage()    {}
func (*FundCollectedEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_30eac93b35d65e1b, []int{2}
}
func (m *FundCollectedEvidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundCollectedEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundCollectedEvidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundCollectedEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundCollectedEvidence.Merge(m, src)
}
func (m *FundCollectedEvidence) XXX_Size() int {
	return m.Size()
}
func (m *FundCollectedEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_FundCollectedEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_FundCollectedEvidence proto.InternalMessageInfo

func (m *FundCollectedEvidence) GetEvidence() []*FundCollectedEvidence_Data {
	if m != nil {
		return m.Evidence
	}
	return nil
}

type FundCollectedEvidence_Data struct {
	From          string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Amount        string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	JobID         string `protobuf:"bytes,3,opt,name=jobID,proto3" json:"jobID,omitempty"`
	AtBlockHeight uint64 `protobuf:"varint,4,opt,name=atBlockHeight,proto3" json:"atBlockHeight,omitempty"`
}

func (m *FundCollectedEvidence_Data) Reset()         { *m = FundCollectedEvidence_Data{} }
func (m *FundCollectedEvidence_Data) String() string { return proto.CompactTextString(m) }
func (*FundCollectedEvidence_Data) ProtoMessage()    {}
func (*FundCollectedEvidence_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_30eac93b35d65e1b, []int{2, 0}
}
func (m *FundCollectedEvidence_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundCollectedEvidence_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundCollectedEvidence_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundCollectedEvidence_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundCollectedEvidence_Data.Merge(m, src)
}
func (m *FundCollectedEvidence_Data) XXX_Size() int {
	return m.Size()
}
func (m *FundCollectedEvidence_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_FundCollectedEvidence_Data.DiscardUnknown(m)
}

var xxx_messageInfo_FundCollectedEvidence_Data proto.InternalMessageInfo

func (m *FundCollectedEvidence_Data) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *FundCollectedEvidence_Data) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundCollectedEvidence_Data) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *FundCollectedEvidence_Data) GetAtBlockHeight() uint64 {
	if m != nil {
		return m.AtBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*FundCollectedEvent)(nil), "palomachain.paloma.evm.FundCollectedEvent")
	proto.RegisterType((*CollectFunds)(nil), "palomachain.paloma.evm.CollectFunds")
	proto.RegisterType((*FundCollectedEvidence)(nil), "palomachain.paloma.evm.FundCollectedEvidence")
	proto.RegisterType((*FundCollectedEvidence_Data)(nil), "palomachain.paloma.evm.FundCollectedEvidence.Data")
}

func init() { proto.RegisterFile("evm/treasury.proto", fileDescriptor_30eac93b35d65e1b) }

var fileDescriptor_30eac93b35d65e1b = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xda, 0x40,
	0x14, 0xf4, 0x82, 0x8b, 0xe8, 0xba, 0x5c, 0x56, 0x14, 0x59, 0x3e, 0x18, 0x0b, 0xf5, 0x40, 0x2f,
	0x6b, 0x89, 0xfe, 0x81, 0xa1, 0xa8, 0xed, 0xa1, 0x07, 0x8b, 0x53, 0x6f, 0x6b, 0x7b, 0x31, 0x6e,
	0xbd, 0x5e, 0xcb, 0x5e, 0xa3, 0x92, 0xaf, 0xe0, 0x1b, 0xf2, 0x35, 0x1c, 0xb9, 0x25, 0xa7, 0x24,
	0x82, 0x1f, 0x89, 0xbc, 0x36, 0x89, 0x49, 0x50, 0xa4, 0xdc, 0xe6, 0x8d, 0x66, 0xe7, 0x79, 0xe6,
	0x19, 0x22, 0xba, 0x66, 0xb6, 0xc8, 0x28, 0xc9, 0x8b, 0x6c, 0x83, 0xd3, 0x8c, 0x0b, 0x8e, 0x06,
	0x29, 0x89, 0x39, 0x23, 0xfe, 0x8a, 0x44, 0x09, 0xae, 0x30, 0xa6, 0x6b, 0x66, 0xf4, 0x43, 0x1e,
	0x72, 0x29, 0xb1, 0x4b, 0x54, 0xa9, 0x8d, 0x61, 0xc8, 0x79, 0x18, 0x53, 0x5b, 0x4e, 0x5e, 0xb1,
	0xb4, 0x45, 0xc4, 0x68, 0x2e, 0x08, 0x4b, 0x2b, 0xc1, 0xe8, 0x0a, 0xa2, 0x79, 0x91, 0x04, 0x53,
	0x1e, 0xc7, 0xd4, 0x17, 0x34, 0xf8, 0xbe, 0xa6, 0x89, 0x40, 0x7d, 0xf8, 0xe1, 0x2f, 0xf7, 0x7e,
	0xce, 0x74, 0x60, 0x81, 0xf1, 0x47, 0xb7, 0x1a, 0xd0, 0x00, 0x76, 0x08, 0xe3, 0x45, 0x22, 0xf4,
	0x96, 0xa4, 0xeb, 0xa9, 0x54, 0x07, 0x34, 0xe1, 0x4c, 0x6f, 0x57, 0x6a, 0x39, 0x20, 0x0b, 0x6a,
	0x5e, 0xcc, 0xfd, 0x7f, 0x3f, 0x68, 0x14, 0xae, 0x84, 0xae, 0x5a, 0x60, 0xac, 0xba, 0x4d, 0x6a,
	0x74, 0x0d, 0xe0, 0xa7, 0x7a, 0x71, 0xf9, 0x0d, 0x39, 0xfa, 0x05, 0x7b, 0xcb, 0x8c, 0x33, 0xa7,
	0xd4, 0x2c, 0x22, 0x46, 0xe5, 0x7a, 0x6d, 0x62, 0xe0, 0x2a, 0x05, 0x3e, 0xa5, 0xc0, 0x8b, 0x53,
	0x0a, 0xa7, 0xbb, 0xbb, 0x1b, 0x2a, 0xdb, 0xfb, 0x21, 0x70, 0xcf, 0x9f, 0xa2, 0x39, 0xd4, 0x04,
	0x7f, 0x76, 0x6a, 0xbd, 0xc3, 0xa9, 0xf9, 0x70, 0x74, 0x03, 0xe0, 0xe7, 0x17, 0x0d, 0x45, 0x01,
	0x4d, 0x7c, 0x8a, 0x7e, 0xc3, 0x2e, 0xad, 0xb1, 0x0e, 0xac, 0xf6, 0x58, 0x9b, 0x4c, 0xf0, 0xe5,
	0xe3, 0xe0, 0x8b, 0x06, 0x78, 0x46, 0x04, 0x71, 0x9f, 0x3c, 0x8c, 0x04, 0xaa, 0x25, 0x83, 0x10,
	0x54, 0xcb, 0x28, 0x75, 0xf7, 0x12, 0xbf, 0x55, 0x7d, 0x75, 0xa8, 0x76, 0xf3, 0x50, 0x5f, 0x60,
	0x8f, 0x08, 0xe7, 0x55, 0xf9, 0xe7, 0xa4, 0x33, 0xdd, 0x1d, 0x4c, 0xb0, 0x3f, 0x98, 0xe0, 0xe1,
	0x60, 0x82, 0xed, 0xd1, 0x54, 0xf6, 0x47, 0x53, 0xb9, 0x3d, 0x9a, 0xca, 0x9f, 0xaf, 0x61, 0x24,
	0x56, 0x85, 0x87, 0x7d, 0xce, 0xec, 0x46, 0xa2, 0x1a, 0xdb, 0xff, 0x6d, 0xf9, 0x5f, 0x6e, 0x52,
	0x9a, 0x7b, 0x1d, 0xd9, 0xe4, 0xb7, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x5b, 0xe3, 0x20,
	0xab, 0x02, 0x00, 0x00,
}

func (m *FundCollectedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundCollectedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundCollectedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CollectFunds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectFunds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectFunds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ToBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ToBlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTreasury(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.FromBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.FromBlockTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTreasury(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FundCollectedEvidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundCollectedEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundCollectedEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTreasury(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FundCollectedEvidence_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundCollectedEvidence_Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundCollectedEvidence_Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AtBlockHeight != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.AtBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreasury(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FundCollectedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTreasury(uint64(m.BlockHeight))
	}
	return n
}

func (m *CollectFunds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.FromBlockTime)
	n += 1 + l + sovTreasury(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ToBlockTime)
	n += 1 + l + sovTreasury(uint64(l))
	return n
}

func (m *FundCollectedEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
			l = e.Size()
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func (m *FundCollectedEvidence_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	if m.AtBlockHeight != 0 {
		n += 1 + sovTreasury(uint64(m.AtBlockHeight))
	}
	return n
}

func sovTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreasury(x uint64) (n int) {
	return sovTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundCollectedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundCollectedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundCollectedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CollectFunds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectFunds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectFunds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.FromBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ToBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FundCollectedEvidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundCollectedEvidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundCollectedEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidence = append(m.Evidence, &FundCollectedEvidence_Data{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FundCollectedEvidence_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AtBlockHeight", wireType)
			}
			m.AtBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AtBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreasury
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreasury = fmt.Errorf("proto: unexpected end of group")
)
