// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uber/cadence/shared/v1/cluster.proto

package sharedv1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HostInfo struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostInfo) Reset()         { *m = HostInfo{} }
func (m *HostInfo) String() string { return proto.CompactTextString(m) }
func (*HostInfo) ProtoMessage()    {}
func (*HostInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e4c5c03631024e, []int{0}
}
func (m *HostInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostInfo.Merge(m, src)
}
func (m *HostInfo) XXX_Size() int {
	return m.Size()
}
func (m *HostInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HostInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HostInfo proto.InternalMessageInfo

func (m *HostInfo) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type RingInfo struct {
	Role                 string      `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	MemberCount          int32       `protobuf:"varint,2,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	Members              []*HostInfo `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RingInfo) Reset()         { *m = RingInfo{} }
func (m *RingInfo) String() string { return proto.CompactTextString(m) }
func (*RingInfo) ProtoMessage()    {}
func (*RingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e4c5c03631024e, []int{1}
}
func (m *RingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RingInfo.Merge(m, src)
}
func (m *RingInfo) XXX_Size() int {
	return m.Size()
}
func (m *RingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RingInfo proto.InternalMessageInfo

func (m *RingInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RingInfo) GetMemberCount() int32 {
	if m != nil {
		return m.MemberCount
	}
	return 0
}

func (m *RingInfo) GetMembers() []*HostInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type MembershipInfo struct {
	CurrentHost          *HostInfo   `protobuf:"bytes,1,opt,name=current_host,json=currentHost,proto3" json:"current_host,omitempty"`
	ReachableMembers     []string    `protobuf:"bytes,2,rep,name=reachable_members,json=reachableMembers,proto3" json:"reachable_members,omitempty"`
	Rings                []*RingInfo `protobuf:"bytes,3,rep,name=rings,proto3" json:"rings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MembershipInfo) Reset()         { *m = MembershipInfo{} }
func (m *MembershipInfo) String() string { return proto.CompactTextString(m) }
func (*MembershipInfo) ProtoMessage()    {}
func (*MembershipInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e4c5c03631024e, []int{2}
}
func (m *MembershipInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MembershipInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MembershipInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MembershipInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipInfo.Merge(m, src)
}
func (m *MembershipInfo) XXX_Size() int {
	return m.Size()
}
func (m *MembershipInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipInfo proto.InternalMessageInfo

func (m *MembershipInfo) GetCurrentHost() *HostInfo {
	if m != nil {
		return m.CurrentHost
	}
	return nil
}

func (m *MembershipInfo) GetReachableMembers() []string {
	if m != nil {
		return m.ReachableMembers
	}
	return nil
}

func (m *MembershipInfo) GetRings() []*RingInfo {
	if m != nil {
		return m.Rings
	}
	return nil
}

type DomainCacheInfo struct {
	NumOfItemsInCacheById   int64    `protobuf:"varint,1,opt,name=num_of_items_in_cache_by_id,json=numOfItemsInCacheById,proto3" json:"num_of_items_in_cache_by_id,omitempty"`
	NumOfItemsInCacheByName int64    `protobuf:"varint,2,opt,name=num_of_items_in_cache_by_name,json=numOfItemsInCacheByName,proto3" json:"num_of_items_in_cache_by_name,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *DomainCacheInfo) Reset()         { *m = DomainCacheInfo{} }
func (m *DomainCacheInfo) String() string { return proto.CompactTextString(m) }
func (*DomainCacheInfo) ProtoMessage()    {}
func (*DomainCacheInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e4c5c03631024e, []int{3}
}
func (m *DomainCacheInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DomainCacheInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DomainCacheInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DomainCacheInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainCacheInfo.Merge(m, src)
}
func (m *DomainCacheInfo) XXX_Size() int {
	return m.Size()
}
func (m *DomainCacheInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainCacheInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DomainCacheInfo proto.InternalMessageInfo

func (m *DomainCacheInfo) GetNumOfItemsInCacheById() int64 {
	if m != nil {
		return m.NumOfItemsInCacheById
	}
	return 0
}

func (m *DomainCacheInfo) GetNumOfItemsInCacheByName() int64 {
	if m != nil {
		return m.NumOfItemsInCacheByName
	}
	return 0
}

func init() {
	proto.RegisterType((*HostInfo)(nil), "uber.cadence.shared.v1.HostInfo")
	proto.RegisterType((*RingInfo)(nil), "uber.cadence.shared.v1.RingInfo")
	proto.RegisterType((*MembershipInfo)(nil), "uber.cadence.shared.v1.MembershipInfo")
	proto.RegisterType((*DomainCacheInfo)(nil), "uber.cadence.shared.v1.DomainCacheInfo")
}

func init() {
	proto.RegisterFile("uber/cadence/shared/v1/cluster.proto", fileDescriptor_31e4c5c03631024e)
}

var fileDescriptor_31e4c5c03631024e = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xaa, 0xd4, 0x30,
	0x14, 0x87, 0xc9, 0xad, 0x57, 0xe7, 0x66, 0x2e, 0xfe, 0x09, 0xa8, 0x83, 0xe2, 0x50, 0x8b, 0x48,
	0x41, 0x48, 0x19, 0x45, 0x17, 0x57, 0x70, 0x71, 0xeb, 0xc2, 0x2e, 0x54, 0xc8, 0xd2, 0x4d, 0x48,
	0xd3, 0x33, 0x6d, 0x60, 0x92, 0x5c, 0xd2, 0x74, 0x60, 0x16, 0x3e, 0x82, 0x2f, 0xe4, 0x13, 0xb8,
	0xf4, 0x11, 0x64, 0x9e, 0x44, 0x9a, 0xb6, 0x23, 0xc2, 0x88, 0x77, 0x77, 0x72, 0xce, 0xf7, 0x6b,
	0xbf, 0x84, 0x83, 0x9f, 0x75, 0x25, 0xb8, 0x4c, 0x8a, 0x0a, 0x8c, 0x84, 0xac, 0x6d, 0x84, 0x83,
	0x2a, 0xdb, 0xae, 0x32, 0xb9, 0xe9, 0x5a, 0x0f, 0x8e, 0x5e, 0x39, 0xeb, 0x2d, 0x79, 0xd0, 0x53,
	0x74, 0xa4, 0xe8, 0x40, 0xd1, 0xed, 0x2a, 0x79, 0x8e, 0x67, 0x1f, 0x6c, 0xeb, 0x0b, 0xb3, 0xb6,
	0xe4, 0x11, 0x9e, 0xa9, 0x0a, 0x8c, 0x57, 0x7e, 0xb7, 0x40, 0x31, 0x4a, 0xcf, 0xd8, 0xe1, 0x9c,
	0x7c, 0xc5, 0x33, 0xa6, 0x4c, 0x1d, 0x38, 0x82, 0x6f, 0x38, 0xbb, 0x81, 0x91, 0x09, 0x35, 0x79,
	0x8a, 0xcf, 0x35, 0xe8, 0x12, 0x1c, 0x97, 0xb6, 0x33, 0x7e, 0x71, 0x12, 0xa3, 0xf4, 0x94, 0xcd,
	0x87, 0x5e, 0xde, 0xb7, 0xc8, 0x05, 0xbe, 0x35, 0x1c, 0xdb, 0x45, 0x14, 0x47, 0xe9, 0xfc, 0x65,
	0x4c, 0x8f, 0x4b, 0xd1, 0xc9, 0x88, 0x4d, 0x81, 0xe4, 0x3b, 0xc2, 0xb7, 0x3f, 0x0e, 0x75, 0xa3,
	0xae, 0x82, 0x45, 0x8e, 0xcf, 0x65, 0xe7, 0x1c, 0x18, 0xcf, 0x1b, 0xdb, 0xfa, 0x60, 0x73, 0x9d,
	0x6f, 0xce, 0xc7, 0x54, 0xdf, 0x20, 0x2f, 0xf0, 0x3d, 0x07, 0x42, 0x36, 0xa2, 0xdc, 0x00, 0x9f,
	0xec, 0x4e, 0xe2, 0x28, 0x3d, 0x63, 0x77, 0x0f, 0x83, 0xf1, 0xc7, 0xe4, 0x0d, 0x3e, 0x75, 0xca,
	0xd4, 0xff, 0xd5, 0x9f, 0x1e, 0x8a, 0x0d, 0x78, 0xf2, 0x0d, 0xe1, 0x3b, 0xef, 0xad, 0x16, 0xca,
	0xe4, 0x42, 0x36, 0x10, 0xec, 0x2f, 0xf0, 0x63, 0xd3, 0x69, 0x6e, 0xd7, 0x5c, 0x79, 0xd0, 0x2d,
	0x57, 0x86, 0xcb, 0x7e, 0xc8, 0xcb, 0x1d, 0x57, 0x55, 0xb8, 0x4c, 0xc4, 0xee, 0x9b, 0x4e, 0x7f,
	0x5e, 0x17, 0x3d, 0x50, 0x0c, 0xd9, 0xcb, 0x5d, 0x51, 0x91, 0x77, 0xf8, 0xc9, 0x3f, 0xb3, 0x46,
	0x68, 0x08, 0x8f, 0x1f, 0xb1, 0x87, 0x47, 0xd2, 0x9f, 0x84, 0x86, 0xcb, 0xfc, 0xc7, 0x7e, 0x89,
	0x7e, 0xee, 0x97, 0xe8, 0xd7, 0x7e, 0x89, 0xbe, 0xbc, 0xae, 0x95, 0x6f, 0xba, 0x92, 0x4a, 0xab,
	0xb3, 0xbf, 0x56, 0x89, 0xd6, 0x60, 0xb2, 0xb0, 0x3d, 0x7f, 0xb6, 0xea, 0xed, 0x50, 0x6d, 0x57,
	0xe5, 0xcd, 0x30, 0x79, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xd5, 0x2a, 0xe3, 0x7f, 0x02,
	0x00, 0x00,
}

func (m *HostInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCluster(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.MemberCount != 0 {
		i = encodeVarintCluster(dAtA, i, uint64(m.MemberCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MembershipInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MembershipInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MembershipInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Rings) > 0 {
		for iNdEx := len(m.Rings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCluster(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ReachableMembers) > 0 {
		for iNdEx := len(m.ReachableMembers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReachableMembers[iNdEx])
			copy(dAtA[i:], m.ReachableMembers[iNdEx])
			i = encodeVarintCluster(dAtA, i, uint64(len(m.ReachableMembers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CurrentHost != nil {
		{
			size, err := m.CurrentHost.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DomainCacheInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DomainCacheInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DomainCacheInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NumOfItemsInCacheByName != 0 {
		i = encodeVarintCluster(dAtA, i, uint64(m.NumOfItemsInCacheByName))
		i--
		dAtA[i] = 0x10
	}
	if m.NumOfItemsInCacheById != 0 {
		i = encodeVarintCluster(dAtA, i, uint64(m.NumOfItemsInCacheById))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	if m.MemberCount != 0 {
		n += 1 + sovCluster(uint64(m.MemberCount))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MembershipInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentHost != nil {
		l = m.CurrentHost.Size()
		n += 1 + l + sovCluster(uint64(l))
	}
	if len(m.ReachableMembers) > 0 {
		for _, s := range m.ReachableMembers {
			l = len(s)
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	if len(m.Rings) > 0 {
		for _, e := range m.Rings {
			l = e.Size()
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DomainCacheInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumOfItemsInCacheById != 0 {
		n += 1 + sovCluster(uint64(m.NumOfItemsInCacheById))
	}
	if m.NumOfItemsInCacheByName != 0 {
		n += 1 + sovCluster(uint64(m.NumOfItemsInCacheByName))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: HostInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: RingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberCount", wireType)
			}
			m.MemberCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &HostInfo{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MembershipInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: MembershipInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MembershipInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHost", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrentHost == nil {
				m.CurrentHost = &HostInfo{}
			}
			if err := m.CurrentHost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReachableMembers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReachableMembers = append(m.ReachableMembers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
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
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rings = append(m.Rings, &RingInfo{})
			if err := m.Rings[len(m.Rings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DomainCacheInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: DomainCacheInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DomainCacheInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfItemsInCacheById", wireType)
			}
			m.NumOfItemsInCacheById = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfItemsInCacheById |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfItemsInCacheByName", wireType)
			}
			m.NumOfItemsInCacheByName = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfItemsInCacheByName |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
				return 0, ErrInvalidLengthCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCluster = fmt.Errorf("proto: unexpected end of group")
)
