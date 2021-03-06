// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

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

// service id
type ProtoCMD int32

const (
	ProtoCMD_CMD_DEFAULT   ProtoCMD = 0
	ProtoCMD_CMD_HEARTBEAT ProtoCMD = 1
	// 登陆
	ProtoCMD_CMD_LOGIN_REQ    ProtoCMD = 1001
	ProtoCMD_CMD_LOGIN_RES    ProtoCMD = 1002
	ProtoCMD_CMD_LOGINOUT_REQ ProtoCMD = 1003
	ProtoCMD_CMD_LOGINOUT_RES ProtoCMD = 1004
	ProtoCMD_CMD_CREATE_REQ   ProtoCMD = 1005
	ProtoCMD_CMD_CREATE_RES   ProtoCMD = 1006
)

var ProtoCMD_name = map[int32]string{
	0:    "CMD_DEFAULT",
	1:    "CMD_HEARTBEAT",
	1001: "CMD_LOGIN_REQ",
	1002: "CMD_LOGIN_RES",
	1003: "CMD_LOGINOUT_REQ",
	1004: "CMD_LOGINOUT_RES",
	1005: "CMD_CREATE_REQ",
	1006: "CMD_CREATE_RES",
}
var ProtoCMD_value = map[string]int32{
	"CMD_DEFAULT":      0,
	"CMD_HEARTBEAT":    1,
	"CMD_LOGIN_REQ":    1001,
	"CMD_LOGIN_RES":    1002,
	"CMD_LOGINOUT_REQ": 1003,
	"CMD_LOGINOUT_RES": 1004,
	"CMD_CREATE_REQ":   1005,
	"CMD_CREATE_RES":   1006,
}

func (x ProtoCMD) String() string {
	return proto.EnumName(ProtoCMD_name, int32(x))
}
func (ProtoCMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_958b6b66202194ab, []int{0}
}

// 登陆状态
type LoginStatus int32

const (
	LoginStatus_LOGIN_SUCESS    LoginStatus = 0
	LoginStatus_LOGIN_NO_USER   LoginStatus = 1
	LoginStatus_LOGIN_ERR_PASS  LoginStatus = 2
	LoginStatus_LOGIN_NO_SERVER LoginStatus = 3
)

var LoginStatus_name = map[int32]string{
	0: "LOGIN_SUCESS",
	1: "LOGIN_NO_USER",
	2: "LOGIN_ERR_PASS",
	3: "LOGIN_NO_SERVER",
}
var LoginStatus_value = map[string]int32{
	"LOGIN_SUCESS":    0,
	"LOGIN_NO_USER":   1,
	"LOGIN_ERR_PASS":  2,
	"LOGIN_NO_SERVER": 3,
}

func (x LoginStatus) String() string {
	return proto.EnumName(LoginStatus_name, int32(x))
}
func (LoginStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_958b6b66202194ab, []int{1}
}

// 创建玩家状态
type CreatePlayerStatus int32

const (
	CreatePlayerStatus_CREATE_SUCESS  CreatePlayerStatus = 0
	CreatePlayerStatus_CREATE_EXIST   CreatePlayerStatus = 1
	CreatePlayerStatus_CREATE_ILLEGAL CreatePlayerStatus = 2
)

var CreatePlayerStatus_name = map[int32]string{
	0: "CREATE_SUCESS",
	1: "CREATE_EXIST",
	2: "CREATE_ILLEGAL",
}
var CreatePlayerStatus_value = map[string]int32{
	"CREATE_SUCESS":  0,
	"CREATE_EXIST":   1,
	"CREATE_ILLEGAL": 2,
}

func (x CreatePlayerStatus) String() string {
	return proto.EnumName(CreatePlayerStatus_name, int32(x))
}
func (CreatePlayerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_958b6b66202194ab, []int{2}
}

// 玩家属性
type PlayerAttr struct {
	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Level    int32  `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Exp      int32  `protobuf:"varint,4,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Gold     int32  `protobuf:"varint,5,opt,name=Gold,proto3" json:"Gold,omitempty"`
	Diamond  int32  `protobuf:"varint,6,opt,name=Diamond,proto3" json:"Diamond,omitempty"`
}

func (m *PlayerAttr) Reset()         { *m = PlayerAttr{} }
func (m *PlayerAttr) String() string { return proto.CompactTextString(m) }
func (*PlayerAttr) ProtoMessage()    {}
func (*PlayerAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_958b6b66202194ab, []int{0}
}
func (m *PlayerAttr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerAttr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PlayerAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAttr.Merge(dst, src)
}
func (m *PlayerAttr) XXX_Size() int {
	return m.Size()
}
func (m *PlayerAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAttr.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAttr proto.InternalMessageInfo

func (m *PlayerAttr) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PlayerAttr) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *PlayerAttr) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *PlayerAttr) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *PlayerAttr) GetGold() int32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *PlayerAttr) GetDiamond() int32 {
	if m != nil {
		return m.Diamond
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayerAttr)(nil), "proto.PlayerAttr")
	proto.RegisterEnum("proto.ProtoCMD", ProtoCMD_name, ProtoCMD_value)
	proto.RegisterEnum("proto.LoginStatus", LoginStatus_name, LoginStatus_value)
	proto.RegisterEnum("proto.CreatePlayerStatus", CreatePlayerStatus_name, CreatePlayerStatus_value)
}
func (m *PlayerAttr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerAttr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.UserId))
	}
	if len(m.NickName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBase(dAtA, i, uint64(len(m.NickName)))
		i += copy(dAtA[i:], m.NickName)
	}
	if m.Level != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Level))
	}
	if m.Exp != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Exp))
	}
	if m.Gold != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Gold))
	}
	if m.Diamond != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintBase(dAtA, i, uint64(m.Diamond))
	}
	return i, nil
}

func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PlayerAttr) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovBase(uint64(m.UserId))
	}
	l = len(m.NickName)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if m.Level != 0 {
		n += 1 + sovBase(uint64(m.Level))
	}
	if m.Exp != 0 {
		n += 1 + sovBase(uint64(m.Exp))
	}
	if m.Gold != 0 {
		n += 1 + sovBase(uint64(m.Gold))
	}
	if m.Diamond != 0 {
		n += 1 + sovBase(uint64(m.Diamond))
	}
	return n
}

func sovBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerAttr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: PlayerAttr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerAttr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NickName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NickName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exp", wireType)
			}
			m.Exp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gold", wireType)
			}
			m.Gold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gold |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Diamond", wireType)
			}
			m.Diamond = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Diamond |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
				return 0, ErrInvalidLengthBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBase
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
				next, err := skipBase(dAtA[start:])
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
	ErrInvalidLengthBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("base.proto", fileDescriptor_base_958b6b66202194ab) }

var fileDescriptor_base_958b6b66202194ab = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0x07, 0x70, 0x6f, 0x72, 0x8e, 0xef, 0xe6, 0x8e, 0xdc, 0x32, 0x07, 0x62, 0x25, 0x24, 0xeb,
	0x44, 0x75, 0x4a, 0x41, 0xc3, 0x13, 0x38, 0xf6, 0x12, 0x2c, 0x39, 0x1f, 0xec, 0xda, 0x08, 0x89,
	0xc2, 0x72, 0xf0, 0x0a, 0x59, 0x24, 0x76, 0xe4, 0x38, 0x08, 0xde, 0x82, 0x92, 0xa7, 0xe0, 0x39,
	0x28, 0x53, 0x52, 0xa2, 0xa4, 0xe2, 0xf3, 0x19, 0x90, 0x77, 0x43, 0x0a, 0xa0, 0xf2, 0xcc, 0x4f,
	0x63, 0xcd, 0x5f, 0x63, 0x03, 0xcc, 0xb3, 0xb5, 0x7a, 0xb8, 0xaa, 0xab, 0xa6, 0x42, 0x5b, 0x3f,
	0x1e, 0x7c, 0x20, 0x00, 0xb3, 0x45, 0xf6, 0x4e, 0xd5, 0x5e, 0xd3, 0xd4, 0x78, 0x0f, 0x9c, 0xcd,
	0x5a, 0xd5, 0x69, 0x91, 0x33, 0x72, 0x4d, 0x6e, 0x6c, 0xd1, 0x6b, 0xdb, 0x30, 0xc7, 0xfb, 0x70,
	0x56, 0x16, 0x2f, 0x5f, 0xa7, 0x65, 0xb6, 0x54, 0xac, 0x73, 0x4d, 0x6e, 0xce, 0xc4, 0x69, 0x0b,
	0x93, 0x6c, 0xa9, 0xf0, 0x0e, 0xd8, 0x91, 0x7a, 0xa3, 0x16, 0xac, 0xab, 0xdf, 0x31, 0x0d, 0x52,
	0xe8, 0xf2, 0xb7, 0x2b, 0x76, 0xa2, 0xad, 0x2d, 0x11, 0xe1, 0x64, 0x54, 0x2d, 0x72, 0x66, 0x6b,
	0xd2, 0x35, 0x32, 0x70, 0x82, 0x22, 0x5b, 0x56, 0x65, 0xce, 0x7a, 0x9a, 0xff, 0xb4, 0x83, 0x8f,
	0x04, 0x4e, 0x67, 0x6d, 0x48, 0x7f, 0x1c, 0xe0, 0x25, 0x9c, 0xfb, 0xe3, 0x20, 0x0d, 0xf8, 0x63,
	0x2f, 0x89, 0x62, 0x6a, 0xe1, 0x6d, 0xb8, 0xd5, 0xc2, 0x13, 0xee, 0x89, 0x78, 0xc8, 0xbd, 0x98,
	0x12, 0x44, 0x43, 0xd1, 0x74, 0x14, 0x4e, 0x52, 0xc1, 0x9f, 0xd2, 0xaf, 0xce, 0xdf, 0x26, 0xe9,
	0x37, 0x07, 0xef, 0x02, 0x3d, 0xda, 0x34, 0x89, 0xf5, 0xe8, 0xf7, 0xff, 0xb1, 0xa4, 0x3f, 0x1c,
	0xbc, 0x82, 0x7e, 0xcb, 0xbe, 0xe0, 0x5e, 0xcc, 0xf5, 0xec, 0xcf, 0x7f, 0x51, 0xd2, 0x5f, 0xce,
	0xe0, 0x05, 0x9c, 0x47, 0xd5, 0xab, 0xa2, 0x94, 0x4d, 0xd6, 0x6c, 0xd6, 0x48, 0xe1, 0xc2, 0xac,
	0x95, 0x89, 0xcf, 0xa5, 0x34, 0x99, 0x8d, 0x4c, 0xa6, 0x69, 0x22, 0xb9, 0xd0, 0x99, 0xfb, 0x86,
	0xb8, 0x10, 0xe9, 0xcc, 0x93, 0x92, 0x76, 0xf0, 0x0a, 0x2e, 0x8f, 0x63, 0x92, 0x8b, 0x67, 0x5c,
	0xd0, 0xee, 0x60, 0x0c, 0xe8, 0xd7, 0x2a, 0x6b, 0x94, 0xf9, 0x5a, 0x87, 0x1d, 0xed, 0x15, 0x4c,
	0x86, 0xe3, 0x12, 0x0a, 0x17, 0x07, 0xe2, 0xcf, 0x43, 0x69, 0xee, 0xd2, 0x3f, 0x48, 0x18, 0x45,
	0x7c, 0xe4, 0x45, 0xb4, 0x33, 0x64, 0x9f, 0x76, 0x2e, 0xd9, 0xee, 0x5c, 0xf2, 0x65, 0xe7, 0x92,
	0xf7, 0x7b, 0xd7, 0xda, 0xee, 0x5d, 0xeb, 0xf3, 0xde, 0xb5, 0xe6, 0x3d, 0xfd, 0x63, 0x3c, 0xfa,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x90, 0xcc, 0xa9, 0xd7, 0x2d, 0x02, 0x00, 0x00,
}
