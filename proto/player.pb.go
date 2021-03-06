// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: player.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// cmd:1005
type CreatePlayerReq struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Length   int32  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Cnt      int32  `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
}

func (m *CreatePlayerReq) Reset()         { *m = CreatePlayerReq{} }
func (m *CreatePlayerReq) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerReq) ProtoMessage()    {}
func (*CreatePlayerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_6d1d20271b1cfe7c, []int{0}
}
func (m *CreatePlayerReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePlayerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePlayerReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CreatePlayerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerReq.Merge(dst, src)
}
func (m *CreatePlayerReq) XXX_Size() int {
	return m.Size()
}
func (m *CreatePlayerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerReq proto.InternalMessageInfo

func (m *CreatePlayerReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreatePlayerReq) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *CreatePlayerReq) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

// cmd:1006
type CreatePlayerRes struct {
	Status CreatePlayerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.CreatePlayerStatus" json:"status,omitempty"`
	Attr   *PlayerAttr        `protobuf:"bytes,2,opt,name=attr" json:"attr,omitempty"`
}

func (m *CreatePlayerRes) Reset()         { *m = CreatePlayerRes{} }
func (m *CreatePlayerRes) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerRes) ProtoMessage()    {}
func (*CreatePlayerRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_6d1d20271b1cfe7c, []int{1}
}
func (m *CreatePlayerRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePlayerRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePlayerRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CreatePlayerRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerRes.Merge(dst, src)
}
func (m *CreatePlayerRes) XXX_Size() int {
	return m.Size()
}
func (m *CreatePlayerRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerRes proto.InternalMessageInfo

func (m *CreatePlayerRes) GetStatus() CreatePlayerStatus {
	if m != nil {
		return m.Status
	}
	return CreatePlayerStatus_CREATE_SUCESS
}

func (m *CreatePlayerRes) GetAttr() *PlayerAttr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePlayerReq)(nil), "proto.CreatePlayerReq")
	proto.RegisterType((*CreatePlayerRes)(nil), "proto.CreatePlayerRes")
}
func (m *CreatePlayerReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePlayerReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlayer(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if m.Length != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPlayer(dAtA, i, uint64(m.Length))
	}
	if m.Cnt != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPlayer(dAtA, i, uint64(m.Cnt))
	}
	return i, nil
}

func (m *CreatePlayerRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePlayerRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPlayer(dAtA, i, uint64(m.Status))
	}
	if m.Attr != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlayer(dAtA, i, uint64(m.Attr.Size()))
		n1, err := m.Attr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintPlayer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreatePlayerReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovPlayer(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovPlayer(uint64(m.Length))
	}
	if m.Cnt != 0 {
		n += 1 + sovPlayer(uint64(m.Cnt))
	}
	return n
}

func (m *CreatePlayerRes) Size() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPlayer(uint64(m.Status))
	}
	if m.Attr != nil {
		l = m.Attr.Size()
		n += 1 + l + sovPlayer(uint64(l))
	}
	return n
}

func sovPlayer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlayer(x uint64) (n int) {
	return sovPlayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePlayerReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePlayerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePlayerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cnt", wireType)
			}
			m.Cnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cnt |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayer
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
func (m *CreatePlayerRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePlayerRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePlayerRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (CreatePlayerStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attr == nil {
				m.Attr = &PlayerAttr{}
			}
			if err := m.Attr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlayer
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
func skipPlayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayer
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
					return 0, ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlayer
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPlayer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlayer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPlayer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPlayer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("player.proto", fileDescriptor_player_6d1d20271b1cfe7c) }

var fileDescriptor_player_6d1d20271b1cfe7c = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x49, 0xac,
	0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x5c, 0x49, 0x89,
	0xc5, 0xa9, 0x10, 0x21, 0xa5, 0x08, 0x2e, 0x7e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0x00, 0xb0,
	0xc2, 0xa0, 0xd4, 0x42, 0x21, 0x69, 0x2e, 0xce, 0xd2, 0xe2, 0xd4, 0xa2, 0xf8, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e, 0xaa, 0x90,
	0x18, 0x17, 0x5b, 0x4e, 0x6a, 0x5e, 0x7a, 0x49, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10,
	0x94, 0x27, 0x24, 0xc0, 0xc5, 0x9c, 0x9c, 0x57, 0x22, 0xc1, 0x0c, 0x16, 0x04, 0x31, 0x95, 0xb2,
	0xd1, 0x4d, 0x2e, 0x16, 0x32, 0xe4, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x06, 0x1b, 0xcb,
	0x67, 0x24, 0x09, 0x71, 0x84, 0x1e, 0xb2, 0xba, 0x60, 0xb0, 0x82, 0x20, 0xa8, 0x42, 0x21, 0x55,
	0x2e, 0x96, 0xc4, 0x92, 0x92, 0x22, 0xb0, 0x6d, 0xdc, 0x46, 0x82, 0x50, 0x0d, 0x10, 0xa5, 0x8e,
	0x25, 0x25, 0x45, 0x41, 0x60, 0x69, 0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x48, 0x62, 0x03, 0xeb, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x64, 0x90, 0xcb, 0x2a,
	0x0a, 0x01, 0x00, 0x00,
}
