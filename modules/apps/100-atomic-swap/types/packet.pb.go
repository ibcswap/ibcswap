// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/atomic_swap/v1/packet.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
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

// Type defines a classification of swap messages
type SwapMessageType int32

const (
	// Default zero value enumeration
	UNSPECIFIED SwapMessageType = 0
	MAKE_SWAP   SwapMessageType = 1
	TAKE_SWAP   SwapMessageType = 2
	CANCEL_SWAP SwapMessageType = 3
)

var SwapMessageType_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_MSG_MAKE_SWAP",
	2: "TYPE_MSG_TAKE_SWAP",
	3: "TYPE_MSG_CANCEL_SWAP",
}

var SwapMessageType_value = map[string]int32{
	"TYPE_UNSPECIFIED":     0,
	"TYPE_MSG_MAKE_SWAP":   1,
	"TYPE_MSG_TAKE_SWAP":   2,
	"TYPE_MSG_CANCEL_SWAP": 3,
}

func (x SwapMessageType) String() string {
	return proto.EnumName(SwapMessageType_name, int32(x))
}

func (SwapMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e225b3c72fc646b1, []int{0}
}

// AtomicSwapPacketData is comprised of a raw transaction, type of transaction and optional memo field.
type AtomicSwapPacketData struct {
	Type SwapMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.applications.atomic_swap.v1.SwapMessageType" json:"type,omitempty"`
	Data []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Memo string          `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *AtomicSwapPacketData) Reset()         { *m = AtomicSwapPacketData{} }
func (m *AtomicSwapPacketData) String() string { return proto.CompactTextString(m) }
func (*AtomicSwapPacketData) ProtoMessage()    {}
func (*AtomicSwapPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e225b3c72fc646b1, []int{0}
}
func (m *AtomicSwapPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AtomicSwapPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AtomicSwapPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AtomicSwapPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtomicSwapPacketData.Merge(m, src)
}
func (m *AtomicSwapPacketData) XXX_Size() int {
	return m.Size()
}
func (m *AtomicSwapPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_AtomicSwapPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_AtomicSwapPacketData proto.InternalMessageInfo

func (m *AtomicSwapPacketData) GetType() SwapMessageType {
	if m != nil {
		return m.Type
	}
	return UNSPECIFIED
}

func (m *AtomicSwapPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AtomicSwapPacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterEnum("ibc.applications.atomic_swap.v1.SwapMessageType", SwapMessageType_name, SwapMessageType_value)
	proto.RegisterType((*AtomicSwapPacketData)(nil), "ibc.applications.atomic_swap.v1.AtomicSwapPacketData")
}

func init() {
	proto.RegisterFile("ibc/applications/atomic_swap/v1/packet.proto", fileDescriptor_e225b3c72fc646b1)
}

var fileDescriptor_e225b3c72fc646b1 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0x4c, 0x4a, 0xd6,
	0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0x2c, 0xc9,
	0xcf, 0xcd, 0x4c, 0x8e, 0x2f, 0x2e, 0x4f, 0x2c, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcf, 0x4c, 0x4a, 0xd6, 0x43, 0x56,
	0xad, 0x87, 0xa4, 0x5a, 0xaf, 0xcc, 0x50, 0x4a, 0x32, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x1f,
	0xac, 0x3c, 0xa9, 0x34, 0x4d, 0x3f, 0x31, 0xaf, 0x12, 0xa2, 0x57, 0x4a, 0x24, 0x3d, 0x3f, 0x3d,
	0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0x2d, 0x8c, 0x5c, 0x22, 0x8e, 0x60, 0x33, 0x82,
	0xcb, 0x13, 0x0b, 0x02, 0xc0, 0x96, 0xb9, 0x24, 0x96, 0x24, 0x0a, 0xb9, 0x70, 0xb1, 0x94, 0x54,
	0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x19, 0xe8, 0x11, 0xb0, 0x59, 0x0f, 0xa4,
	0xdd, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x35, 0xa4, 0xb2, 0x20, 0x35, 0x08, 0xac, 0x5b, 0x48,
	0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x06,
	0x89, 0xe5, 0xa6, 0xe6, 0xe6, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x5a, 0xdb,
	0x19, 0xb9, 0xf8, 0xd1, 0x4c, 0x10, 0x52, 0xe5, 0x12, 0x08, 0x89, 0x0c, 0x70, 0x8d, 0x0f, 0xf5,
	0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x90, 0xe2, 0xef, 0x9a, 0xab,
	0xc0, 0x8d, 0x24, 0x24, 0xa4, 0xca, 0x25, 0x04, 0x56, 0xe6, 0x1b, 0xec, 0x1e, 0xef, 0xeb, 0xe8,
	0xed, 0x1a, 0x1f, 0x1c, 0xee, 0x18, 0x20, 0xc0, 0x28, 0xc5, 0xdb, 0x35, 0x57, 0x81, 0x13, 0x2e,
	0x80, 0xa2, 0x2c, 0x04, 0xae, 0x8c, 0x09, 0xa2, 0x0c, 0x2e, 0x20, 0xa4, 0xc9, 0x25, 0x02, 0x57,
	0xe6, 0xec, 0xe8, 0xe7, 0xec, 0xea, 0x03, 0x51, 0xc8, 0x0c, 0xb1, 0x18, 0x49, 0x48, 0x8a, 0xa5,
	0x63, 0xb1, 0x1c, 0x83, 0x53, 0xd4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x39, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0x26, 0x25, 0x83,
	0x23, 0x14, 0x46, 0x97, 0x99, 0xe8, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0x83, 0x22, 0xbf,
	0x58, 0xdf, 0xd0, 0xc0, 0x40, 0x17, 0x39, 0xe2, 0x41, 0x81, 0x57, 0x9c, 0xc4, 0x06, 0x8e, 0x23,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x98, 0x70, 0x6b, 0x25, 0x02, 0x00, 0x00,
}

func (m *AtomicSwapPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AtomicSwapPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AtomicSwapPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AtomicSwapPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPacket(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AtomicSwapPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: AtomicSwapPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AtomicSwapPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SwapMessageType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
