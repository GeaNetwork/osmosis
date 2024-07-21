// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/poolmanager/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// GenesisState defines the poolmanager module's genesis state.
type GenesisState struct {
	// the next_pool_id
	NextPoolId uint64 `protobuf:"varint,1,opt,name=next_pool_id,json=nextPoolId,proto3" json:"next_pool_id,omitempty"`
	// pool_routes is the container of the mappings from pool id to pool type.
	PoolRoutes []ModuleRoute `protobuf:"bytes,2,rep,name=pool_routes,json=poolRoutes,proto3" json:"pool_routes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2cd1c042c4d32a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetNextPoolId() uint64 {
	if m != nil {
		return m.NextPoolId
	}
	return 0
}

func (m *GenesisState) GetPoolRoutes() []ModuleRoute {
	if m != nil {
		return m.PoolRoutes
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dymensionxyz.dymension.poolmanager.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/poolmanager/v1beta1/genesis.proto", fileDescriptor_6d2cd1c042c4d32a)
}

var fileDescriptor_6d2cd1c042c4d32a = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0x93, 0xff, 0xbf, 0x62, 0x48, 0x3b, 0x55, 0x0c, 0xa5, 0x83, 0xa9, 0x98, 0x2a, 0x24,
	0x6c, 0x15, 0x84, 0xca, 0xc2, 0xd2, 0x05, 0x31, 0x20, 0x41, 0xd9, 0x18, 0x88, 0xec, 0xe6, 0x30,
	0x96, 0x12, 0x5f, 0x15, 0x3b, 0x55, 0xc3, 0x53, 0xf4, 0xb1, 0x3a, 0x76, 0x64, 0x42, 0xa8, 0x7d,
	0x11, 0x14, 0xbb, 0x8d, 0x0a, 0x13, 0x6c, 0xfe, 0xee, 0xbb, 0xdf, 0xf9, 0xee, 0x8b, 0xae, 0x92,
	0x32, 0x03, 0x6d, 0x14, 0xea, 0x79, 0xf9, 0xc6, 0x6a, 0xc1, 0xa6, 0x88, 0x69, 0xc6, 0x35, 0x97,
	0x90, 0xb3, 0xd9, 0x40, 0x80, 0xe5, 0x03, 0x26, 0x41, 0x83, 0x51, 0x86, 0x4e, 0x73, 0xb4, 0xd8,
	0x3e, 0xdd, 0x27, 0x69, 0x2d, 0xe8, 0x1e, 0x49, 0xb7, 0x64, 0xf7, 0x50, 0xa2, 0x44, 0x87, 0xb1,
	0xea, 0xe5, 0x27, 0x74, 0x8f, 0x24, 0xa2, 0x4c, 0x81, 0x39, 0x25, 0x8a, 0x17, 0xc6, 0x75, 0xb9,
	0xb3, 0x26, 0x68, 0x32, 0x34, 0xb1, 0x67, 0xbc, 0xd8, 0x5a, 0xe4, 0x27, 0x95, 0x14, 0x39, 0xb7,
	0xee, 0x67, 0xef, 0xfb, 0x6e, 0x26, 0xb8, 0x81, 0x7a, 0xf5, 0x09, 0xaa, 0x9d, 0x7f, 0xfd, 0x87,
	0x8b, 0x33, 0x4c, 0x8a, 0x14, 0xe2, 0x1c, 0x0b, 0x0b, 0x1e, 0x3f, 0x59, 0x84, 0x51, 0xeb, 0xc6,
	0x07, 0xf1, 0x68, 0xb9, 0x85, 0x76, 0x2f, 0x6a, 0x69, 0x98, 0xdb, 0xb8, 0xe2, 0x63, 0x95, 0x74,
	0xc2, 0x5e, 0xd8, 0x6f, 0x8c, 0xa3, 0xaa, 0x76, 0x8f, 0x98, 0xde, 0x26, 0xed, 0xe7, 0xa8, 0xe9,
	0x4c, 0x37, 0xc6, 0x74, 0xfe, 0xf5, 0xfe, 0xf7, 0x9b, 0xe7, 0x43, 0xfa, 0xfb, 0xfc, 0xe8, 0x9d,
	0xdb, 0x63, 0x5c, 0xf1, 0xa3, 0xc6, 0xf2, 0xe3, 0x38, 0x18, 0x47, 0x55, 0x9b, 0x2b, 0x98, 0xd1,
	0xc3, 0x72, 0x4d, 0xc2, 0xd5, 0x9a, 0x84, 0x9f, 0x6b, 0x12, 0x2e, 0x36, 0x24, 0x58, 0x6d, 0x48,
	0xf0, 0xbe, 0x21, 0xc1, 0xd3, 0x50, 0x2a, 0xfb, 0x5a, 0x08, 0x3a, 0xc1, 0x8c, 0xb9, 0x54, 0x94,
	0x39, 0x4b, 0xb9, 0x30, 0x3b, 0xc1, 0x66, 0x83, 0x4b, 0x36, 0xff, 0x76, 0xba, 0x2d, 0xa7, 0x60,
	0xc4, 0x81, 0x3b, 0xf6, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x25, 0x56, 0x8e, 0x1f, 0x02,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolRoutes) > 0 {
		for iNdEx := len(m.PoolRoutes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRoutes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.NextPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.NextPoolId))
	}
	if len(m.PoolRoutes) > 0 {
		for _, e := range m.PoolRoutes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPoolId", wireType)
			}
			m.NextPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRoutes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolRoutes = append(m.PoolRoutes, ModuleRoute{})
			if err := m.PoolRoutes[len(m.PoolRoutes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
