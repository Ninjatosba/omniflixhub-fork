// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: omniflix/flixdrop/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the flixdrop module's genesis state.
type GenesisState struct {
	// balance of the flixdrop module's account
	ModuleAccountBalance types.Coin `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	// actions defines all actions to perform to claim flixdrop
	Actions []WeightedAction `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions" yaml:"actions"`
	// list of claim records, one for every flixdrop recipient
	ClaimRecords []ClaimRecord `protobuf:"bytes,4,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records" yaml:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1f06d92f70d950e, []int{0}
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

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetActions() []WeightedAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "omniflix.flixdrop.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("omniflix/flixdrop/v1beta1/genesis.proto", fileDescriptor_f1f06d92f70d950e)
}

var fileDescriptor_f1f06d92f70d950e = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0xee, 0xd2, 0x30,
	0x1c, 0xc7, 0x37, 0x31, 0x7f, 0x93, 0xfd, 0xc1, 0xc3, 0x82, 0x64, 0x10, 0x1d, 0xb8, 0x04, 0xc4,
	0x4b, 0x27, 0x78, 0xf3, 0xc6, 0x48, 0x34, 0xf1, 0x22, 0x99, 0x07, 0x13, 0x3d, 0x90, 0xae, 0xab,
	0xa3, 0xc9, 0xda, 0x2e, 0x6b, 0x47, 0xe4, 0x2d, 0x7c, 0x04, 0x1f, 0x87, 0x23, 0x47, 0x4f, 0xc4,
	0xc0, 0x1b, 0xf8, 0x04, 0x66, 0x6d, 0x17, 0x34, 0x11, 0x2e, 0xcb, 0xd2, 0x7e, 0x7e, 0x9f, 0x6f,
	0xfb, 0xeb, 0xcf, 0x79, 0xc1, 0x29, 0x23, 0x5f, 0x73, 0xf2, 0x2d, 0xac, 0x3f, 0x69, 0xc9, 0x8b,
	0x70, 0x3b, 0x4b, 0xb0, 0x84, 0xb3, 0x30, 0xc3, 0x0c, 0x0b, 0x22, 0x40, 0x51, 0x72, 0xc9, 0xdd,
	0x7e, 0x03, 0x82, 0x06, 0x04, 0x06, 0x1c, 0x74, 0x33, 0x9e, 0x71, 0x45, 0x85, 0xf5, 0x9f, 0x2e,
	0x18, 0x4c, 0xae, 0x9b, 0x0b, 0x58, 0x42, 0x6a, 0xc4, 0x83, 0xf1, 0x75, 0x0e, 0xe5, 0x90, 0x50,
	0x83, 0xf9, 0x88, 0x0b, 0xca, 0x45, 0x98, 0x40, 0x81, 0x2f, 0x00, 0x27, 0x4c, 0xef, 0x07, 0x3f,
	0x5a, 0x4e, 0xfb, 0x9d, 0x3e, 0xf1, 0x47, 0x09, 0x25, 0x76, 0xb7, 0x4e, 0x8f, 0xf2, 0xb4, 0xca,
	0xf1, 0x1a, 0x22, 0xc4, 0x2b, 0x26, 0xd7, 0x09, 0xcc, 0x21, 0x43, 0xd8, 0xb3, 0x47, 0xf6, 0xf4,
	0x7e, 0xde, 0x07, 0xda, 0x08, 0x6a, 0x63, 0x73, 0x17, 0xb0, 0xe4, 0x84, 0x45, 0xe3, 0xfd, 0x71,
	0x68, 0xfd, 0x3e, 0x0e, 0x9f, 0xed, 0x20, 0xcd, 0xdf, 0x04, 0xff, 0xd7, 0x04, 0x71, 0x57, 0x6f,
	0x2c, 0xf4, 0x7a, 0xa4, 0x97, 0xdd, 0x95, 0x73, 0xa7, 0xef, 0xe7, 0x3d, 0x50, 0x39, 0xcf, 0xc1,
	0xd5, 0xce, 0x81, 0x95, 0x02, 0xa3, 0x27, 0x26, 0xaf, 0xa3, 0xf3, 0x74, 0x79, 0x10, 0x1b, 0x8f,
	0xfb, 0xc5, 0x79, 0x04, 0x91, 0x24, 0x9c, 0x09, 0xaf, 0x35, 0x6a, 0x4d, 0xef, 0xe7, 0x2f, 0x6f,
	0x28, 0x3f, 0x61, 0x92, 0x6d, 0x24, 0x4e, 0x17, 0xaa, 0x22, 0xea, 0x19, 0xf5, 0x63, 0xad, 0x36,
	0x9e, 0x20, 0x6e, 0x8c, 0x2e, 0x71, 0x3a, 0xaa, 0xcd, 0xeb, 0x12, 0x23, 0x5e, 0xa6, 0xc2, 0x7b,
	0xa8, 0x22, 0x26, 0x37, 0x22, 0x96, 0x35, 0x1f, 0x2b, 0x3c, 0x7a, 0x6a, 0xfc, 0x5d, 0xed, 0xff,
	0x47, 0x15, 0xc4, 0x6d, 0x74, 0x41, 0x45, 0xf4, 0x7e, 0x7f, 0xf2, 0xed, 0xc3, 0xc9, 0xb7, 0x7f,
	0x9d, 0x7c, 0xfb, 0xfb, 0xd9, 0xb7, 0x0e, 0x67, 0xdf, 0xfa, 0x79, 0xf6, 0xad, 0xcf, 0xaf, 0x32,
	0x22, 0x37, 0x55, 0x02, 0x10, 0xa7, 0xe1, 0x07, 0xca, 0xc8, 0xdb, 0x7a, 0x1c, 0x9a, 0x03, 0x6c,
	0xaa, 0x24, 0xfc, 0x6b, 0x38, 0xe4, 0xae, 0xc0, 0x22, 0xb9, 0x53, 0xaf, 0xfe, 0xfa, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd6, 0xde, 0xad, 0x9a, 0xc0, 0x02, 0x00, 0x00,
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
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
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
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
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
			m.Actions = append(m.Actions, WeightedAction{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
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
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
