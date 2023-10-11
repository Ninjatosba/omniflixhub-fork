// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: OmniFlix/globalfee/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState - initial state of module
type GenesisState struct {
	// Params of this module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede90870a7dc1a24, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Params defines the set of module parameters.
type Params struct {
	// minimum_gas_prices stores the minimum gas price(s) for all TX on the chain.
	// When multiple coins are defined then they are accepted alternatively.
	// The list must be sorted by denoms asc. No duplicate denoms or zero amount
	// values allowed. For more information see
	// https://docs.cosmos.network/main/modules/auth#concepts
	MinimumGasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=minimum_gas_prices,json=minimumGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum_gas_prices,omitempty" yaml:"minimum_gas_prices"`
	// bypass_min_fee_msg_types defines a list of message type urls
	// that are free of fee charge.
	BypassMinFeeMsgTypes []string `protobuf:"bytes,2,rep,name=bypass_min_fee_msg_types,json=bypassMinFeeMsgTypes,proto3" json:"bypass_min_fee_msg_types,omitempty" yaml:"bypass_min_fee_msg_types"`
	// max_total_bypass_min_fee_msg_gas_usage defines the total maximum gas usage
	// allowed for a transaction containing only messages of types in bypass_min_fee_msg_types
	// to bypass fee charge.
	MaxTotalBypassMinFeeMsgGasUsage uint64 `protobuf:"varint,3,opt,name=max_total_bypass_min_fee_msg_gas_usage,json=maxTotalBypassMinFeeMsgGasUsage,proto3" json:"max_total_bypass_min_fee_msg_gas_usage,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ede90870a7dc1a24, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMinimumGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinimumGasPrices
	}
	return nil
}

func (m *Params) GetBypassMinFeeMsgTypes() []string {
	if m != nil {
		return m.BypassMinFeeMsgTypes
	}
	return nil
}

func (m *Params) GetMaxTotalBypassMinFeeMsgGasUsage() uint64 {
	if m != nil {
		return m.MaxTotalBypassMinFeeMsgGasUsage
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "OmniFlix.globalfee.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "OmniFlix.globalfee.v1beta1.Params")
}

func init() {
	proto.RegisterFile("OmniFlix/globalfee/v1beta1/genesis.proto", fileDescriptor_ede90870a7dc1a24)
}

var fileDescriptor_ede90870a7dc1a24 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x6e, 0xd3, 0x40,
	0x1c, 0xc7, 0x73, 0x04, 0x45, 0xc2, 0x65, 0xa8, 0xac, 0x0e, 0x26, 0xaa, 0xec, 0xc8, 0x03, 0xb2,
	0x04, 0xdc, 0xa9, 0x81, 0x89, 0xd1, 0xa0, 0x46, 0x0c, 0xa5, 0x55, 0x08, 0x0b, 0xcb, 0xe9, 0xec,
	0x5e, 0xae, 0x27, 0x7c, 0x77, 0x56, 0x7e, 0x97, 0x2a, 0x19, 0x79, 0x03, 0x1e, 0x80, 0x27, 0xe0,
	0x19, 0x78, 0x80, 0x8e, 0x1d, 0x99, 0x0c, 0x4a, 0xb6, 0x8e, 0x3c, 0x01, 0xb2, 0xcf, 0x4d, 0x5b,
	0xb5, 0x99, 0x6c, 0xf9, 0xf7, 0xf9, 0xfe, 0xf1, 0x4f, 0x3f, 0x2f, 0x39, 0x56, 0x5a, 0x1e, 0x16,
	0x72, 0x41, 0x44, 0x61, 0x32, 0x56, 0x4c, 0x39, 0x27, 0xe7, 0x07, 0x19, 0xb7, 0xec, 0x80, 0x08,
	0xae, 0x39, 0x48, 0xc0, 0xe5, 0xcc, 0x58, 0xe3, 0xf7, 0xaf, 0x49, 0xbc, 0x21, 0x71, 0x4b, 0xf6,
	0xf7, 0x84, 0x11, 0xa6, 0xc1, 0x48, 0xfd, 0xe6, 0x14, 0xfd, 0x30, 0x37, 0xa0, 0x0c, 0x90, 0x8c,
	0xc1, 0x8d, 0x69, 0x6e, 0xa4, 0x76, 0xf3, 0xf8, 0xd4, 0x7b, 0x3a, 0x72, 0x11, 0x9f, 0x2c, 0xb3,
	0xdc, 0x9f, 0x78, 0xbd, 0x92, 0xcd, 0x98, 0x82, 0x00, 0x0d, 0x50, 0xb2, 0x33, 0x8c, 0xf1, 0xf6,
	0x48, 0x7c, 0xd2, 0x90, 0x69, 0x70, 0x51, 0x45, 0x9d, 0xab, 0x2a, 0xda, 0x75, 0xca, 0x97, 0x46,
	0x49, 0xcb, 0x55, 0x69, 0x97, 0xe3, 0xd6, 0x2b, 0xfe, 0xd1, 0xf5, 0x7a, 0x0e, 0xf6, 0x7f, 0x21,
	0xcf, 0x57, 0x52, 0x4b, 0x35, 0x57, 0x54, 0x30, 0xa0, 0xe5, 0x4c, 0xe6, 0xbc, 0x4e, 0xeb, 0x26,
	0x3b, 0xc3, 0x7d, 0xec, 0xea, 0xe2, 0xba, 0xee, 0x26, 0xe6, 0x3d, 0xcf, 0xdf, 0x19, 0xa9, 0xd3,
	0xb2, 0xcd, 0xd9, 0xbf, 0xaf, 0xbf, 0xc9, 0xfc, 0x57, 0x45, 0xcf, 0x96, 0x4c, 0x15, 0x6f, 0xe3,
	0xfb, 0x54, 0xfc, 0xf3, 0x4f, 0xf4, 0x42, 0x48, 0x7b, 0x36, 0xcf, 0x70, 0x6e, 0x14, 0x69, 0x77,
	0xe3, 0x1e, 0xaf, 0xe0, 0xf4, 0x2b, 0xb1, 0xcb, 0x92, 0xc3, 0x75, 0x20, 0x8c, 0x77, 0x5b, 0x8f,
	0x11, 0x83, 0x93, 0xc6, 0xc1, 0xff, 0x86, 0xbc, 0x20, 0x5b, 0x96, 0x0c, 0x80, 0x2a, 0xa9, 0xe9,
	0x94, 0x73, 0xaa, 0x40, 0xd0, 0x46, 0x17, 0x3c, 0x1a, 0x74, 0x93, 0x27, 0xe9, 0x87, 0xab, 0x2a,
	0x8a, 0xb7, 0x31, 0x77, 0x8a, 0x46, 0xae, 0xe8, 0x36, 0x36, 0x1e, 0xef, 0xb9, 0xd1, 0x91, 0xd4,
	0x87, 0x9c, 0x1f, 0x81, 0x98, 0xd4, 0x9f, 0xfd, 0x63, 0xef, 0xb9, 0x62, 0x0b, 0x6a, 0x8d, 0x65,
	0x05, 0x7d, 0x40, 0x5c, 0xff, 0xf0, 0x1c, 0x98, 0xe0, 0x41, 0x77, 0x80, 0x92, 0xc7, 0xe3, 0x48,
	0xb1, 0xc5, 0xa4, 0x86, 0xd3, 0xbb, 0x6e, 0x23, 0x06, 0x9f, 0x6b, 0x2c, 0xfd, 0x78, 0xb1, 0x0a,
	0xd1, 0xe5, 0x2a, 0x44, 0x7f, 0x57, 0x21, 0xfa, 0xbe, 0x0e, 0x3b, 0x97, 0xeb, 0xb0, 0xf3, 0x7b,
	0x1d, 0x76, 0xbe, 0xbc, 0xb9, 0xb5, 0xad, 0xcd, 0x95, 0x1a, 0xa5, 0xe5, 0xb4, 0x90, 0x8b, 0xb3,
	0x79, 0x46, 0xce, 0x87, 0xe4, 0xf6, 0xd9, 0x36, 0xbd, 0xb3, 0x5e, 0x73, 0x5b, 0xaf, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0x6f, 0x4d, 0x08, 0xd9, 0x02, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxTotalBypassMinFeeMsgGasUsage != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxTotalBypassMinFeeMsgGasUsage))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BypassMinFeeMsgTypes) > 0 {
		for iNdEx := len(m.BypassMinFeeMsgTypes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BypassMinFeeMsgTypes[iNdEx])
			copy(dAtA[i:], m.BypassMinFeeMsgTypes[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BypassMinFeeMsgTypes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MinimumGasPrices) > 0 {
		for iNdEx := len(m.MinimumGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinimumGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for _, e := range m.MinimumGasPrices {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BypassMinFeeMsgTypes) > 0 {
		for _, s := range m.BypassMinFeeMsgTypes {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MaxTotalBypassMinFeeMsgGasUsage != 0 {
		n += 1 + sovGenesis(uint64(m.MaxTotalBypassMinFeeMsgGasUsage))
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrices", wireType)
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
			m.MinimumGasPrices = append(m.MinimumGasPrices, types.DecCoin{})
			if err := m.MinimumGasPrices[len(m.MinimumGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BypassMinFeeMsgTypes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BypassMinFeeMsgTypes = append(m.BypassMinFeeMsgTypes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTotalBypassMinFeeMsgGasUsage", wireType)
			}
			m.MaxTotalBypassMinFeeMsgGasUsage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTotalBypassMinFeeMsgGasUsage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
