// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dapplink/incentives/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// active incentives
	Incentives []Incentive `protobuf:"bytes,2,rep,name=incentives,proto3" json:"incentives"`
	// active Gasmeters
	GasMeters []GasMeter `protobuf:"bytes,3,rep,name=gas_meters,json=gasMeters,proto3" json:"gas_meters"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cc649dfbad21d0, []int{0}
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

func (m *GenesisState) GetIncentives() []Incentive {
	if m != nil {
		return m.Incentives
	}
	return nil
}

func (m *GenesisState) GetGasMeters() []GasMeter {
	if m != nil {
		return m.GasMeters
	}
	return nil
}

// Params defines the incentives module params
type Params struct {
	// parameter to enable incentives
	EnableIncentives bool `protobuf:"varint,1,opt,name=enable_incentives,json=enableIncentives,proto3" json:"enable_incentives,omitempty"`
	// maximum percentage an incentive can allocate per denomination
	AllocationLimit github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=allocation_limit,json=allocationLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"allocation_limit"`
	// identifier for the epochs module hooks
	IncentivesEpochIdentifier string `protobuf:"bytes,3,opt,name=incentives_epoch_identifier,json=incentivesEpochIdentifier,proto3" json:"incentives_epoch_identifier,omitempty"`
	// scaling factor for capping rewards
	RewardScaler github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reward_scaler,json=rewardScaler,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_scaler"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cc649dfbad21d0, []int{1}
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

func (m *Params) GetEnableIncentives() bool {
	if m != nil {
		return m.EnableIncentives
	}
	return false
}

func (m *Params) GetIncentivesEpochIdentifier() string {
	if m != nil {
		return m.IncentivesEpochIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dapplink.incentives.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "dapplink.incentives.v1.Params")
}

func init() { proto.RegisterFile("dapplink/incentives/v1/genesis.proto", fileDescriptor_b5cc649dfbad21d0) }

var fileDescriptor_b5cc649dfbad21d0 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x6a, 0xdb, 0x40,
	0x18, 0xc7, 0x25, 0xdb, 0x98, 0xfa, 0xec, 0x52, 0x57, 0xed, 0xa0, 0xda, 0x54, 0x76, 0x4d, 0x29,
	0x86, 0xe2, 0x13, 0x76, 0xa7, 0x2e, 0x1d, 0x84, 0x8b, 0x31, 0xb4, 0x10, 0xe4, 0x29, 0x59, 0xc4,
	0x49, 0xba, 0xc8, 0x47, 0x24, 0x9d, 0xd0, 0x5d, 0x9c, 0xe4, 0x2d, 0xf2, 0x12, 0x79, 0x17, 0x8f,
	0x86, 0x2c, 0x21, 0x83, 0x09, 0xf6, 0x8b, 0x84, 0x3b, 0x29, 0x91, 0x06, 0x4d, 0x99, 0x74, 0x27,
	0xfd, 0xfe, 0xbf, 0xff, 0x27, 0xf8, 0xc0, 0x37, 0x97, 0x70, 0xca, 0x4c, 0x12, 0x7b, 0x38, 0xe6,
	0x64, 0x83, 0x99, 0xb9, 0x99, 0x9a, 0x01, 0x8e, 0x31, 0x23, 0x0c, 0x26, 0x29, 0xe5, 0x54, 0xfb,
	0x24, 0x11, 0x58, 0x20, 0x70, 0x33, 0xed, 0x7d, 0xaf, 0xca, 0x95, 0x10, 0x19, 0xed, 0x7d, 0x0e,
	0x68, 0x40, 0xe5, 0xd1, 0x14, 0xa7, 0xec, 0xed, 0xe8, 0x5e, 0x05, 0x9d, 0x45, 0x56, 0xb1, 0xe2,
	0x88, 0x63, 0xed, 0x37, 0x68, 0x26, 0x28, 0x45, 0x11, 0xd3, 0xd5, 0xa1, 0x3a, 0x6e, 0xcf, 0xfa,
	0xb0, 0xa2, 0x12, 0x9e, 0x48, 0xc4, 0x6a, 0x6c, 0xf7, 0x03, 0xc5, 0xce, 0x03, 0xda, 0x1c, 0x80,
	0x82, 0xd2, 0x6b, 0xc3, 0xfa, 0xb8, 0x3d, 0x33, 0x2a, 0xe3, 0xcb, 0x97, 0x5b, 0x6e, 0x28, 0xe5,
	0x34, 0x0b, 0x80, 0x00, 0x31, 0x27, 0xc2, 0x1c, 0xa7, 0x4c, 0xaf, 0x4b, 0xcb, 0xd7, 0x4a, 0xcb,
	0x02, 0xb1, 0xff, 0x82, 0xca, 0x25, 0xad, 0x20, 0xbf, 0xb3, 0xd1, 0x5d, 0x0d, 0x34, 0xb3, 0x11,
	0xb5, 0x9f, 0xe0, 0x23, 0x8e, 0x91, 0x1b, 0x62, 0xa7, 0x34, 0x9b, 0xf8, 0xb5, 0x77, 0x76, 0x37,
	0xfb, 0xb0, 0x2c, 0xba, 0x4f, 0x41, 0x17, 0x85, 0x21, 0xf5, 0x10, 0x27, 0x34, 0x76, 0x42, 0x12,
	0x11, 0xae, 0xd7, 0x86, 0xea, 0xb8, 0x65, 0x41, 0x51, 0xf1, 0xb8, 0x1f, 0xfc, 0x08, 0x08, 0x5f,
	0x5f, 0xba, 0xd0, 0xa3, 0x91, 0xe9, 0x51, 0x16, 0x51, 0x96, 0x3f, 0x26, 0xcc, 0xbf, 0x30, 0xf9,
	0x4d, 0x82, 0x19, 0x9c, 0x63, 0xcf, 0xfe, 0x50, 0x78, 0xfe, 0x09, 0x8d, 0xf6, 0x07, 0xf4, 0x8b,
	0x01, 0x1c, 0x9c, 0x50, 0x6f, 0xed, 0x10, 0x5f, 0xdc, 0xcf, 0x09, 0x4e, 0xf5, 0xba, 0x68, 0xb1,
	0xbf, 0x14, 0xc8, 0x5f, 0x41, 0x2c, 0x5f, 0x01, 0x6d, 0x05, 0xde, 0xa7, 0xf8, 0x0a, 0xa5, 0xbe,
	0xc3, 0x3c, 0x14, 0xe2, 0x54, 0x6f, 0xbc, 0x69, 0xae, 0x4e, 0x26, 0x59, 0x49, 0x87, 0xb5, 0xd8,
	0x1e, 0x0c, 0x75, 0x77, 0x30, 0xd4, 0xa7, 0x83, 0xa1, 0xde, 0x1e, 0x0d, 0x65, 0x77, 0x34, 0x94,
	0x87, 0xa3, 0xa1, 0x9c, 0x4d, 0x4a, 0x3e, 0x97, 0x70, 0x1f, 0xd1, 0x09, 0xa1, 0x66, 0xb6, 0x68,
	0xd7, 0xe5, 0x55, 0x93, 0x6a, 0xb7, 0x29, 0xb7, 0xe9, 0xd7, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xac, 0xa6, 0x0e, 0x7f, 0xc3, 0x02, 0x00, 0x00,
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
	if len(m.GasMeters) > 0 {
		for iNdEx := len(m.GasMeters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasMeters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Incentives) > 0 {
		for iNdEx := len(m.Incentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Incentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size := m.RewardScaler.Size()
		i -= size
		if _, err := m.RewardScaler.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.IncentivesEpochIdentifier) > 0 {
		i -= len(m.IncentivesEpochIdentifier)
		copy(dAtA[i:], m.IncentivesEpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IncentivesEpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.AllocationLimit.Size()
		i -= size
		if _, err := m.AllocationLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableIncentives {
		i--
		if m.EnableIncentives {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Incentives) > 0 {
		for _, e := range m.Incentives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasMeters) > 0 {
		for _, e := range m.GasMeters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableIncentives {
		n += 2
	}
	l = m.AllocationLimit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.IncentivesEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.RewardScaler.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incentives", wireType)
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
			m.Incentives = append(m.Incentives, Incentive{})
			if err := m.Incentives[len(m.Incentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasMeters", wireType)
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
			m.GasMeters = append(m.GasMeters, GasMeter{})
			if err := m.GasMeters[len(m.GasMeters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableIncentives", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableIncentives = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationLimit", wireType)
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
			if err := m.AllocationLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesEpochIdentifier", wireType)
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
			m.IncentivesEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardScaler", wireType)
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
			if err := m.RewardScaler.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
