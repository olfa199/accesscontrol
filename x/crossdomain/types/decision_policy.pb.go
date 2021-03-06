// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crossdomain/decision_policy.proto

package types

import (
	fmt "fmt"
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

type DecisionPolicy struct {
	Depth        uint64    `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	Cost         uint64    `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	LocationList []string  `protobuf:"bytes,3,rep,name=locationList,proto3" json:"locationList,omitempty"`
	InterestList []string  `protobuf:"bytes,4,rep,name=interestList,proto3" json:"interestList,omitempty"`
	Validity     *Validity `protobuf:"bytes,5,opt,name=validity,proto3" json:"validity,omitempty"`
	LastUpdate   string    `protobuf:"bytes,6,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	Creator      string    `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *DecisionPolicy) Reset()         { *m = DecisionPolicy{} }
func (m *DecisionPolicy) String() string { return proto.CompactTextString(m) }
func (*DecisionPolicy) ProtoMessage()    {}
func (*DecisionPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8042ce219f13373, []int{0}
}
func (m *DecisionPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DecisionPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DecisionPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DecisionPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionPolicy.Merge(m, src)
}
func (m *DecisionPolicy) XXX_Size() int {
	return m.Size()
}
func (m *DecisionPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionPolicy proto.InternalMessageInfo

func (m *DecisionPolicy) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *DecisionPolicy) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *DecisionPolicy) GetLocationList() []string {
	if m != nil {
		return m.LocationList
	}
	return nil
}

func (m *DecisionPolicy) GetInterestList() []string {
	if m != nil {
		return m.InterestList
	}
	return nil
}

func (m *DecisionPolicy) GetValidity() *Validity {
	if m != nil {
		return m.Validity
	}
	return nil
}

func (m *DecisionPolicy) GetLastUpdate() string {
	if m != nil {
		return m.LastUpdate
	}
	return ""
}

func (m *DecisionPolicy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*DecisionPolicy)(nil), "crossdomain.crossdomain.DecisionPolicy")
}

func init() { proto.RegisterFile("crossdomain/decision_policy.proto", fileDescriptor_e8042ce219f13373) }

var fileDescriptor_e8042ce219f13373 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xbf, 0x69, 0xfb, 0xd7, 0x20, 0x06, 0x0b, 0x09, 0xab, 0x83, 0x49, 0x3b, 0x65,
	0x4a, 0x25, 0x98, 0x18, 0x58, 0x10, 0x23, 0x03, 0x8a, 0x04, 0x03, 0x0b, 0x32, 0xb6, 0x25, 0x2c,
	0x85, 0x5c, 0xcb, 0xbe, 0x42, 0xf4, 0x2d, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x0c, 0x16,
	0x84, 0x9b, 0x22, 0x33, 0xb0, 0xdd, 0x73, 0xfc, 0x9d, 0x2b, 0xdf, 0x43, 0x17, 0xca, 0x43, 0x08,
	0x1a, 0x9e, 0xa4, 0x6d, 0x57, 0xda, 0x28, 0x1b, 0x2c, 0xb4, 0xf7, 0x0e, 0x1a, 0xab, 0xd6, 0x95,
	0xf3, 0x80, 0xc0, 0x8e, 0x12, 0xa4, 0x4a, 0xe6, 0xf9, 0x3c, 0xcd, 0x3e, 0xcb, 0xc6, 0x6a, 0x8b,
	0x43, 0x68, 0xf9, 0x49, 0xe8, 0xc1, 0xe5, 0xb0, 0xee, 0x3a, 0x6e, 0x63, 0x87, 0x74, 0xac, 0x8d,
	0xc3, 0x47, 0x4e, 0x0a, 0x52, 0xe6, 0xf5, 0x56, 0x30, 0x46, 0x73, 0x05, 0x01, 0xf9, 0xbf, 0x68,
	0xc6, 0x99, 0x2d, 0xe9, 0x7e, 0x03, 0x4a, 0xa2, 0x85, 0xf6, 0xca, 0x06, 0xe4, 0xa3, 0x62, 0x54,
	0xce, 0xea, 0x5f, 0xde, 0x37, 0x63, 0x5b, 0x34, 0xde, 0x04, 0x8c, 0x4c, 0xbe, 0x65, 0x52, 0x8f,
	0x9d, 0xd3, 0xff, 0xbb, 0x6f, 0xf1, 0x71, 0x41, 0xca, 0xbd, 0x93, 0x45, 0xf5, 0xc7, 0x31, 0xd5,
	0xed, 0x00, 0xd6, 0x3f, 0x11, 0x26, 0x28, 0x6d, 0x64, 0xc0, 0x1b, 0xa7, 0x25, 0x1a, 0x3e, 0x29,
	0x48, 0x39, 0xab, 0x13, 0x87, 0x71, 0x3a, 0x55, 0xde, 0x48, 0x04, 0xcf, 0xa7, 0xf1, 0x71, 0x27,
	0x2f, 0xce, 0xde, 0x3a, 0x41, 0x36, 0x9d, 0x20, 0x1f, 0x9d, 0x20, 0xaf, 0xbd, 0xc8, 0x36, 0xbd,
	0xc8, 0xde, 0x7b, 0x91, 0xdd, 0x1d, 0xa7, 0x9d, 0xbd, 0xac, 0x52, 0x85, 0x6b, 0x67, 0xc2, 0xc3,
	0x24, 0xf6, 0x77, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x47, 0xae, 0x0b, 0xd5, 0x99, 0x01, 0x00,
	0x00,
}

func (m *DecisionPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DecisionPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DecisionPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDecisionPolicy(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastUpdate) > 0 {
		i -= len(m.LastUpdate)
		copy(dAtA[i:], m.LastUpdate)
		i = encodeVarintDecisionPolicy(dAtA, i, uint64(len(m.LastUpdate)))
		i--
		dAtA[i] = 0x32
	}
	if m.Validity != nil {
		{
			size, err := m.Validity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDecisionPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.InterestList) > 0 {
		for iNdEx := len(m.InterestList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InterestList[iNdEx])
			copy(dAtA[i:], m.InterestList[iNdEx])
			i = encodeVarintDecisionPolicy(dAtA, i, uint64(len(m.InterestList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LocationList) > 0 {
		for iNdEx := len(m.LocationList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LocationList[iNdEx])
			copy(dAtA[i:], m.LocationList[iNdEx])
			i = encodeVarintDecisionPolicy(dAtA, i, uint64(len(m.LocationList[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Cost != 0 {
		i = encodeVarintDecisionPolicy(dAtA, i, uint64(m.Cost))
		i--
		dAtA[i] = 0x10
	}
	if m.Depth != 0 {
		i = encodeVarintDecisionPolicy(dAtA, i, uint64(m.Depth))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDecisionPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovDecisionPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DecisionPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Depth != 0 {
		n += 1 + sovDecisionPolicy(uint64(m.Depth))
	}
	if m.Cost != 0 {
		n += 1 + sovDecisionPolicy(uint64(m.Cost))
	}
	if len(m.LocationList) > 0 {
		for _, s := range m.LocationList {
			l = len(s)
			n += 1 + l + sovDecisionPolicy(uint64(l))
		}
	}
	if len(m.InterestList) > 0 {
		for _, s := range m.InterestList {
			l = len(s)
			n += 1 + l + sovDecisionPolicy(uint64(l))
		}
	}
	if m.Validity != nil {
		l = m.Validity.Size()
		n += 1 + l + sovDecisionPolicy(uint64(l))
	}
	l = len(m.LastUpdate)
	if l > 0 {
		n += 1 + l + sovDecisionPolicy(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDecisionPolicy(uint64(l))
	}
	return n
}

func sovDecisionPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDecisionPolicy(x uint64) (n int) {
	return sovDecisionPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DecisionPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDecisionPolicy
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
			return fmt.Errorf("proto: DecisionPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DecisionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depth", wireType)
			}
			m.Depth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Depth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cost", wireType)
			}
			m.Cost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocationList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
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
				return ErrInvalidLengthDecisionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDecisionPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocationList = append(m.LocationList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
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
				return ErrInvalidLengthDecisionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDecisionPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterestList = append(m.InterestList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
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
				return ErrInvalidLengthDecisionPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDecisionPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validity == nil {
				m.Validity = &Validity{}
			}
			if err := m.Validity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
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
				return ErrInvalidLengthDecisionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDecisionPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDecisionPolicy
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
				return ErrInvalidLengthDecisionPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDecisionPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDecisionPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDecisionPolicy
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
func skipDecisionPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDecisionPolicy
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
					return 0, ErrIntOverflowDecisionPolicy
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
					return 0, ErrIntOverflowDecisionPolicy
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
				return 0, ErrInvalidLengthDecisionPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDecisionPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDecisionPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDecisionPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDecisionPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDecisionPolicy = fmt.Errorf("proto: unexpected end of group")
)
