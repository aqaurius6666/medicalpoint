// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: medipoint/super_admin.proto

package types

import (
	fmt "fmt"
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

type SuperAdmin struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SuperAdmin) Reset()         { *m = SuperAdmin{} }
func (m *SuperAdmin) String() string { return proto.CompactTextString(m) }
func (*SuperAdmin) ProtoMessage()    {}
func (*SuperAdmin) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d4f1bf43337551, []int{0}
}
func (m *SuperAdmin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperAdmin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperAdmin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperAdmin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperAdmin.Merge(m, src)
}
func (m *SuperAdmin) XXX_Size() int {
	return m.Size()
}
func (m *SuperAdmin) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperAdmin.DiscardUnknown(m)
}

var xxx_messageInfo_SuperAdmin proto.InternalMessageInfo

func (m *SuperAdmin) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SuperAdmin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SuperAdmin) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SuperAdmin)(nil), "sotanext.medipoint.medipoint.SuperAdmin")
}

func init() { proto.RegisterFile("medipoint/super_admin.proto", fileDescriptor_e4d4f1bf43337551) }

var fileDescriptor_e4d4f1bf43337551 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4d, 0x4d, 0xc9,
	0x2c, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2f, 0x2e, 0x2d, 0x48, 0x2d, 0x8a, 0x4f, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xce, 0x2f, 0x49, 0xcc, 0x4b, 0xad,
	0x28, 0xd1, 0x83, 0xab, 0x42, 0xb0, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x0a, 0xf5, 0x41,
	0x2c, 0x88, 0x1e, 0xa5, 0x00, 0x2e, 0xae, 0x60, 0x90, 0x41, 0x8e, 0x20, 0x73, 0x84, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x24, 0xb8,
	0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x60,
	0x5c, 0x90, 0x4c, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x33, 0x44, 0x06, 0xca, 0x75,
	0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xfd, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0x53, 0xf5, 0x11, 0x1e, 0xaa, 0x40, 0x62,
	0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x51, 0x12, 0xbb, 0x2f, 0xf6, 0x00, 0x00, 0x00,
}

func (m *SuperAdmin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperAdmin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperAdmin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSuperAdmin(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSuperAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSuperAdmin(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSuperAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovSuperAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuperAdmin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSuperAdmin(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSuperAdmin(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSuperAdmin(uint64(l))
	}
	return n
}

func sovSuperAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSuperAdmin(x uint64) (n int) {
	return sovSuperAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuperAdmin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperAdmin
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
			return fmt.Errorf("proto: SuperAdmin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperAdmin
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
				return ErrInvalidLengthSuperAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperAdmin
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
				return ErrInvalidLengthSuperAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperAdmin
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
func skipSuperAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSuperAdmin
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
					return 0, ErrIntOverflowSuperAdmin
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
					return 0, ErrIntOverflowSuperAdmin
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
				return 0, ErrInvalidLengthSuperAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSuperAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSuperAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSuperAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSuperAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSuperAdmin = fmt.Errorf("proto: unexpected end of group")
)
