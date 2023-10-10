// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/cwhooks/v1/cwhooks.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Contract is the proto definition of a contract that can be registered for the hooks
type Contract struct {
	// contract_address
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// register_address
	RegisterAddress string `protobuf:"bytes,2,opt,name=register_address,json=registerAddress,proto3" json:"register_address,omitempty"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab9a924dd50ee7b, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Contract) GetRegisterAddress() string {
	if m != nil {
		return m.RegisterAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Contract)(nil), "juno.cwhooks.v1.Contract")
}

func init() { proto.RegisterFile("juno/cwhooks/v1/cwhooks.proto", fileDescriptor_4ab9a924dd50ee7b) }

var fileDescriptor_4ab9a924dd50ee7b = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2a, 0xcd, 0xcb,
	0xd7, 0x4f, 0x2e, 0xcf, 0xc8, 0xcf, 0xcf, 0x2e, 0xd6, 0x2f, 0x33, 0x84, 0x31, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x41, 0xd2, 0x7a, 0x30, 0xb1, 0x32, 0x43, 0xa5, 0x04, 0x2e, 0x0e,
	0xe7, 0xfc, 0xbc, 0x92, 0xa2, 0xc4, 0xe4, 0x12, 0x21, 0x4d, 0x2e, 0x81, 0x64, 0x28, 0x3b, 0x3e,
	0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x1f, 0x26,
	0xee, 0x08, 0x11, 0x06, 0x29, 0x2d, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x82, 0x2b, 0x65,
	0x82, 0x28, 0x85, 0x89, 0x43, 0x95, 0x3a, 0x79, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x73,
	0x7e, 0x71, 0x6e, 0x7e, 0x31, 0xcc, 0x29, 0xc5, 0xfa, 0x60, 0x6f, 0x54, 0xe8, 0x27, 0x97, 0xeb,
	0x42, 0x7c, 0x52, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x85, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0xf1, 0xfd, 0x54, 0xe6, 0x00, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RegisterAddress) > 0 {
		i -= len(m.RegisterAddress)
		copy(dAtA[i:], m.RegisterAddress)
		i = encodeVarintCwhooks(dAtA, i, uint64(len(m.RegisterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintCwhooks(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCwhooks(dAtA []byte, offset int, v uint64) int {
	offset -= sovCwhooks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovCwhooks(uint64(l))
	}
	l = len(m.RegisterAddress)
	if l > 0 {
		n += 1 + l + sovCwhooks(uint64(l))
	}
	return n
}

func sovCwhooks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCwhooks(x uint64) (n int) {
	return sovCwhooks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCwhooks
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwhooks
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
				return ErrInvalidLengthCwhooks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwhooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCwhooks
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
				return ErrInvalidLengthCwhooks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCwhooks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCwhooks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCwhooks
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
func skipCwhooks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCwhooks
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
					return 0, ErrIntOverflowCwhooks
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
					return 0, ErrIntOverflowCwhooks
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
				return 0, ErrInvalidLengthCwhooks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCwhooks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCwhooks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCwhooks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCwhooks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCwhooks = fmt.Errorf("proto: unexpected end of group")
)
