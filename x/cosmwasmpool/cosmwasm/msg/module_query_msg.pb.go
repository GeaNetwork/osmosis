// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/module_query_msg.proto

package msg

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// ===================== CalcOutAmtGivenIn
type CalcOutAmtGivenIn struct {
	// token_in is the token to be sent to the pool.
	TokenIn types.Coin `protobuf:"bytes,1,opt,name=token_in,json=tokenIn,proto3" json:"token_in"`
	// token_out_denom is the token denom to be received from the pool.
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
	// swap_fee is the swap fee for this swap estimate.
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee"`
}

func (m *CalcOutAmtGivenIn) Reset()         { *m = CalcOutAmtGivenIn{} }
func (m *CalcOutAmtGivenIn) String() string { return proto.CompactTextString(m) }
func (*CalcOutAmtGivenIn) ProtoMessage()    {}
func (*CalcOutAmtGivenIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{0}
}
func (m *CalcOutAmtGivenIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcOutAmtGivenIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcOutAmtGivenIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcOutAmtGivenIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcOutAmtGivenIn.Merge(m, src)
}
func (m *CalcOutAmtGivenIn) XXX_Size() int {
	return m.Size()
}
func (m *CalcOutAmtGivenIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcOutAmtGivenIn.DiscardUnknown(m)
}

var xxx_messageInfo_CalcOutAmtGivenIn proto.InternalMessageInfo

func (m *CalcOutAmtGivenIn) GetTokenIn() types.Coin {
	if m != nil {
		return m.TokenIn
	}
	return types.Coin{}
}

func (m *CalcOutAmtGivenIn) GetTokenOutDenom() string {
	if m != nil {
		return m.TokenOutDenom
	}
	return ""
}

type CalcOutAmtGivenInRequest struct {
	// calc_out_amt_given_in is the structure containing all the request
	// information for this query.
	CalcOutAmtGivenIn CalcOutAmtGivenIn `protobuf:"bytes,1,opt,name=calc_out_amt_given_in,json=calcOutAmtGivenIn,proto3" json:"calc_out_amt_given_in"`
}

func (m *CalcOutAmtGivenInRequest) Reset()         { *m = CalcOutAmtGivenInRequest{} }
func (m *CalcOutAmtGivenInRequest) String() string { return proto.CompactTextString(m) }
func (*CalcOutAmtGivenInRequest) ProtoMessage()    {}
func (*CalcOutAmtGivenInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{1}
}
func (m *CalcOutAmtGivenInRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcOutAmtGivenInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcOutAmtGivenInRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcOutAmtGivenInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcOutAmtGivenInRequest.Merge(m, src)
}
func (m *CalcOutAmtGivenInRequest) XXX_Size() int {
	return m.Size()
}
func (m *CalcOutAmtGivenInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcOutAmtGivenInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcOutAmtGivenInRequest proto.InternalMessageInfo

func (m *CalcOutAmtGivenInRequest) GetCalcOutAmtGivenIn() CalcOutAmtGivenIn {
	if m != nil {
		return m.CalcOutAmtGivenIn
	}
	return CalcOutAmtGivenIn{}
}

type CalcOutAmtGivenInResponse struct {
	// token_out is the token out computed from this swap estimate call.
	TokenOut types.Coin `protobuf:"bytes,1,opt,name=token_out,json=tokenOut,proto3" json:"token_out"`
}

func (m *CalcOutAmtGivenInResponse) Reset()         { *m = CalcOutAmtGivenInResponse{} }
func (m *CalcOutAmtGivenInResponse) String() string { return proto.CompactTextString(m) }
func (*CalcOutAmtGivenInResponse) ProtoMessage()    {}
func (*CalcOutAmtGivenInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{2}
}
func (m *CalcOutAmtGivenInResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcOutAmtGivenInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcOutAmtGivenInResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcOutAmtGivenInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcOutAmtGivenInResponse.Merge(m, src)
}
func (m *CalcOutAmtGivenInResponse) XXX_Size() int {
	return m.Size()
}
func (m *CalcOutAmtGivenInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcOutAmtGivenInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcOutAmtGivenInResponse proto.InternalMessageInfo

func (m *CalcOutAmtGivenInResponse) GetTokenOut() types.Coin {
	if m != nil {
		return m.TokenOut
	}
	return types.Coin{}
}

// ===================== CalcInAmtGivenOut
type CalcInAmtGivenOut struct {
	// token_out is the token out to be receoved from the pool.
	TokenOut types.Coin `protobuf:"bytes,1,opt,name=token_out,json=tokenOut,proto3" json:"token_out"`
	// token_in_denom is the token denom to be sentt to the pool.
	TokenInDenom string `protobuf:"bytes,2,opt,name=token_in_denom,json=tokenInDenom,proto3" json:"token_in_denom,omitempty"`
	// swap_fee is the swap fee for this swap estimate.
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee"`
}

func (m *CalcInAmtGivenOut) Reset()         { *m = CalcInAmtGivenOut{} }
func (m *CalcInAmtGivenOut) String() string { return proto.CompactTextString(m) }
func (*CalcInAmtGivenOut) ProtoMessage()    {}
func (*CalcInAmtGivenOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{3}
}
func (m *CalcInAmtGivenOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcInAmtGivenOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcInAmtGivenOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcInAmtGivenOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcInAmtGivenOut.Merge(m, src)
}
func (m *CalcInAmtGivenOut) XXX_Size() int {
	return m.Size()
}
func (m *CalcInAmtGivenOut) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcInAmtGivenOut.DiscardUnknown(m)
}

var xxx_messageInfo_CalcInAmtGivenOut proto.InternalMessageInfo

func (m *CalcInAmtGivenOut) GetTokenOut() types.Coin {
	if m != nil {
		return m.TokenOut
	}
	return types.Coin{}
}

func (m *CalcInAmtGivenOut) GetTokenInDenom() string {
	if m != nil {
		return m.TokenInDenom
	}
	return ""
}

type CalcInAmtGivenOutRequest struct {
	// calc_in_amt_given_out is the structure containing all the request
	// information for this query.
	CalcInAmtGivenOut CalcInAmtGivenOut `protobuf:"bytes,1,opt,name=calc_in_amt_given_out,json=calcInAmtGivenOut,proto3" json:"calc_in_amt_given_out"`
}

func (m *CalcInAmtGivenOutRequest) Reset()         { *m = CalcInAmtGivenOutRequest{} }
func (m *CalcInAmtGivenOutRequest) String() string { return proto.CompactTextString(m) }
func (*CalcInAmtGivenOutRequest) ProtoMessage()    {}
func (*CalcInAmtGivenOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{4}
}
func (m *CalcInAmtGivenOutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcInAmtGivenOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcInAmtGivenOutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcInAmtGivenOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcInAmtGivenOutRequest.Merge(m, src)
}
func (m *CalcInAmtGivenOutRequest) XXX_Size() int {
	return m.Size()
}
func (m *CalcInAmtGivenOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcInAmtGivenOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcInAmtGivenOutRequest proto.InternalMessageInfo

func (m *CalcInAmtGivenOutRequest) GetCalcInAmtGivenOut() CalcInAmtGivenOut {
	if m != nil {
		return m.CalcInAmtGivenOut
	}
	return CalcInAmtGivenOut{}
}

type CalcInAmtGivenOutResponse struct {
	// token_in is the token in computed from this swap estimate call.
	TokenIn types.Coin `protobuf:"bytes,1,opt,name=token_in,json=tokenIn,proto3" json:"token_in"`
}

func (m *CalcInAmtGivenOutResponse) Reset()         { *m = CalcInAmtGivenOutResponse{} }
func (m *CalcInAmtGivenOutResponse) String() string { return proto.CompactTextString(m) }
func (*CalcInAmtGivenOutResponse) ProtoMessage()    {}
func (*CalcInAmtGivenOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e43e2b40a371ec3, []int{5}
}
func (m *CalcInAmtGivenOutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalcInAmtGivenOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalcInAmtGivenOutResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalcInAmtGivenOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcInAmtGivenOutResponse.Merge(m, src)
}
func (m *CalcInAmtGivenOutResponse) XXX_Size() int {
	return m.Size()
}
func (m *CalcInAmtGivenOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcInAmtGivenOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcInAmtGivenOutResponse proto.InternalMessageInfo

func (m *CalcInAmtGivenOutResponse) GetTokenIn() types.Coin {
	if m != nil {
		return m.TokenIn
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*CalcOutAmtGivenIn)(nil), "osmosis.cosmwasmpool.v1beta1.CalcOutAmtGivenIn")
	proto.RegisterType((*CalcOutAmtGivenInRequest)(nil), "osmosis.cosmwasmpool.v1beta1.CalcOutAmtGivenInRequest")
	proto.RegisterType((*CalcOutAmtGivenInResponse)(nil), "osmosis.cosmwasmpool.v1beta1.CalcOutAmtGivenInResponse")
	proto.RegisterType((*CalcInAmtGivenOut)(nil), "osmosis.cosmwasmpool.v1beta1.CalcInAmtGivenOut")
	proto.RegisterType((*CalcInAmtGivenOutRequest)(nil), "osmosis.cosmwasmpool.v1beta1.CalcInAmtGivenOutRequest")
	proto.RegisterType((*CalcInAmtGivenOutResponse)(nil), "osmosis.cosmwasmpool.v1beta1.CalcInAmtGivenOutResponse")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/module_query_msg.proto", fileDescriptor_4e43e2b40a371ec3)
}

var fileDescriptor_4e43e2b40a371ec3 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x33, 0x2a, 0x36, 0x1d, 0xff, 0xd1, 0x45, 0x21, 0x2d, 0xb2, 0x2d, 0x41, 0x4a, 0x2f,
	0x9d, 0x21, 0x0a, 0x82, 0x22, 0x88, 0x69, 0x51, 0x72, 0x0a, 0xe4, 0x22, 0xf5, 0xb2, 0xcc, 0x4e,
	0x5e, 0xd7, 0xa5, 0x3b, 0xf3, 0x6e, 0x3b, 0x33, 0xa9, 0x3d, 0xfb, 0x05, 0xfc, 0x4c, 0x9e, 0x7a,
	0xec, 0x51, 0x3c, 0x14, 0x49, 0xbe, 0x88, 0xcc, 0xfe, 0x49, 0xb3, 0xa9, 0x48, 0xa1, 0x78, 0x49,
	0x32, 0xf3, 0xce, 0xf3, 0xcc, 0xfb, 0xbc, 0xbf, 0xec, 0xd2, 0x57, 0x68, 0x14, 0x9a, 0xd4, 0x70,
	0x89, 0x46, 0x9d, 0x08, 0xa3, 0x72, 0xc4, 0x8c, 0x4f, 0x7a, 0x31, 0x58, 0xd1, 0xe3, 0x0a, 0xc7,
	0x90, 0xf9, 0x4f, 0x97, 0x41, 0x74, 0xe4, 0xe0, 0xf8, 0x34, 0x52, 0x26, 0x61, 0xf9, 0x31, 0x5a,
	0x0c, 0x9e, 0x56, 0x52, 0xb6, 0x28, 0x65, 0x95, 0x74, 0xe3, 0x71, 0x82, 0x09, 0x16, 0x07, 0xb9,
	0xff, 0x55, 0x6a, 0x36, 0x42, 0x59, 0x88, 0x78, 0x2c, 0x0c, 0xcc, 0x6f, 0x91, 0x98, 0xea, 0xb2,
	0xde, 0xfd, 0x41, 0xe8, 0xda, 0x9e, 0xc8, 0xe4, 0xd0, 0xd9, 0x77, 0xca, 0x7e, 0x48, 0x27, 0xa0,
	0x07, 0x3a, 0x78, 0x4d, 0xdb, 0x16, 0x0f, 0x41, 0x47, 0xa9, 0xee, 0x90, 0x2d, 0xb2, 0x73, 0xef,
	0xf9, 0x3a, 0x2b, 0x8d, 0x98, 0x37, 0xaa, 0xef, 0x64, 0x7b, 0x98, 0xea, 0xfe, 0x9d, 0xb3, 0x8b,
	0xcd, 0xd6, 0x68, 0xa5, 0x10, 0x0c, 0x74, 0xb0, 0x4d, 0x1f, 0x95, 0x5a, 0x74, 0x36, 0x1a, 0x83,
	0x46, 0xd5, 0xb9, 0xb5, 0x45, 0x76, 0x56, 0x47, 0x0f, 0x8a, 0xed, 0xa1, 0xb3, 0xfb, 0x7e, 0x33,
	0x18, 0xd0, 0xb6, 0x39, 0x11, 0x79, 0xf4, 0x19, 0xa0, 0x73, 0xdb, 0x1f, 0xe8, 0x33, 0x6f, 0xf4,
	0xeb, 0x62, 0x73, 0x3b, 0x49, 0xed, 0x17, 0x17, 0x33, 0x89, 0x8a, 0x57, 0xed, 0x97, 0x5f, 0xbb,
	0x66, 0x7c, 0xc8, 0xed, 0x69, 0x0e, 0x86, 0xed, 0x83, 0x1c, 0xad, 0x78, 0xfd, 0x7b, 0x80, 0xee,
	0x37, 0x42, 0x3b, 0x57, 0x42, 0x8c, 0xe0, 0xc8, 0x81, 0xb1, 0x41, 0x42, 0x9f, 0x48, 0x91, 0xc9,
	0xa2, 0x1d, 0xa1, 0x6c, 0x94, 0xf8, 0xf2, 0x65, 0x30, 0xce, 0xfe, 0x35, 0x55, 0x76, 0xc5, 0xb6,
	0x8a, 0xbb, 0x26, 0x97, 0x0b, 0xdd, 0x03, 0xba, 0xfe, 0x97, 0x26, 0x4c, 0x8e, 0xda, 0x40, 0xf0,
	0x86, 0xae, 0xce, 0xa7, 0x72, 0xdd, 0x91, 0xb6, 0xeb, 0x81, 0xcd, 0x29, 0x0d, 0x74, 0x6d, 0x3d,
	0x74, 0xf6, 0x66, 0x9e, 0xc1, 0x33, 0xfa, 0xb0, 0x66, 0xdc, 0xc0, 0x74, 0xbf, 0x02, 0xf9, 0xdf,
	0x28, 0x35, 0x42, 0x2c, 0x53, 0x4a, 0xf5, 0x02, 0xa4, 0xcb, 0x5c, 0xd7, 0xa0, 0xd4, 0xb0, 0x5d,
	0xa4, 0xd4, 0x28, 0x74, 0x3f, 0x96, 0x94, 0x96, 0x9a, 0xa8, 0x28, 0xdd, 0xe0, 0x7f, 0xdf, 0x3f,
	0x38, 0x9b, 0x86, 0xe4, 0x7c, 0x1a, 0x92, 0xdf, 0xd3, 0x90, 0x7c, 0x9f, 0x85, 0xad, 0xf3, 0x59,
	0xd8, 0xfa, 0x39, 0x0b, 0x5b, 0x9f, 0xde, 0x2e, 0x4c, 0xaa, 0x8a, 0xb1, 0x9b, 0x89, 0xd8, 0xd4,
	0x0b, 0x3e, 0xe9, 0xbd, 0xe4, 0x5f, 0x9b, 0x2f, 0x84, 0x7a, 0xc1, 0x95, 0x49, 0xe2, 0xbb, 0xc5,
	0xb3, 0xfa, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x58, 0x25, 0xeb, 0x3c, 0x04, 0x00,
	0x00,
}

func (m *CalcOutAmtGivenIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcOutAmtGivenIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcOutAmtGivenIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.TokenOutDenom) > 0 {
		i -= len(m.TokenOutDenom)
		copy(dAtA[i:], m.TokenOutDenom)
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(len(m.TokenOutDenom)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.TokenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CalcOutAmtGivenInRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcOutAmtGivenInRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcOutAmtGivenInRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CalcOutAmtGivenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CalcOutAmtGivenInResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcOutAmtGivenInResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcOutAmtGivenInResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CalcInAmtGivenOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcInAmtGivenOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcInAmtGivenOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.TokenInDenom) > 0 {
		i -= len(m.TokenInDenom)
		copy(dAtA[i:], m.TokenInDenom)
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(len(m.TokenInDenom)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.TokenOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CalcInAmtGivenOutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcInAmtGivenOutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcInAmtGivenOutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CalcInAmtGivenOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CalcInAmtGivenOutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalcInAmtGivenOutResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalcInAmtGivenOutResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModuleQueryMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintModuleQueryMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovModuleQueryMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CalcOutAmtGivenIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenIn.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	l = len(m.TokenOutDenom)
	if l > 0 {
		n += 1 + l + sovModuleQueryMsg(uint64(l))
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func (m *CalcOutAmtGivenInRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CalcOutAmtGivenIn.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func (m *CalcOutAmtGivenInResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenOut.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func (m *CalcInAmtGivenOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenOut.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	l = len(m.TokenInDenom)
	if l > 0 {
		n += 1 + l + sovModuleQueryMsg(uint64(l))
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func (m *CalcInAmtGivenOutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CalcInAmtGivenOut.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func (m *CalcInAmtGivenOutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenIn.Size()
	n += 1 + l + sovModuleQueryMsg(uint64(l))
	return n
}

func sovModuleQueryMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModuleQueryMsg(x uint64) (n int) {
	return sovModuleQueryMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CalcOutAmtGivenIn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcOutAmtGivenIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcOutAmtGivenIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOutDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func (m *CalcOutAmtGivenInRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcOutAmtGivenInRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcOutAmtGivenInRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CalcOutAmtGivenIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CalcOutAmtGivenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func (m *CalcOutAmtGivenInResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcOutAmtGivenInResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcOutAmtGivenInResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func (m *CalcInAmtGivenOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcInAmtGivenOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcInAmtGivenOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenInDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func (m *CalcInAmtGivenOutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcInAmtGivenOutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcInAmtGivenOutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CalcInAmtGivenOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CalcInAmtGivenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func (m *CalcInAmtGivenOutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleQueryMsg
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
			return fmt.Errorf("proto: CalcInAmtGivenOutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalcInAmtGivenOutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleQueryMsg
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
				return ErrInvalidLengthModuleQueryMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModuleQueryMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModuleQueryMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleQueryMsg
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
func skipModuleQueryMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModuleQueryMsg
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
					return 0, ErrIntOverflowModuleQueryMsg
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
					return 0, ErrIntOverflowModuleQueryMsg
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
				return 0, ErrInvalidLengthModuleQueryMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModuleQueryMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModuleQueryMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModuleQueryMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModuleQueryMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModuleQueryMsg = fmt.Errorf("proto: unexpected end of group")
)