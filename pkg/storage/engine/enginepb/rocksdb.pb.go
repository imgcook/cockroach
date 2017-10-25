// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cockroach/pkg/storage/engine/enginepb/rocksdb.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SSTUserProperties contains the user-added properties of a single sstable.
type SSTUserProperties struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ts_min is the minimum mvcc timestamp present in this sstable.
	TsMin *cockroach_util_hlc.Timestamp `protobuf:"bytes,2,opt,name=ts_min,json=tsMin" json:"ts_min,omitempty"`
	// ts_max is the maximum mvcc timestamp present in this sstable.
	TsMax *cockroach_util_hlc.Timestamp `protobuf:"bytes,3,opt,name=ts_max,json=tsMax" json:"ts_max,omitempty"`
}

func (m *SSTUserProperties) Reset()                    { *m = SSTUserProperties{} }
func (m *SSTUserProperties) String() string            { return proto.CompactTextString(m) }
func (*SSTUserProperties) ProtoMessage()               {}
func (*SSTUserProperties) Descriptor() ([]byte, []int) { return fileDescriptorRocksdb, []int{0} }

// SSTUserPropertiesCollection contains the user-added properties of every
// sstable in a RocksDB instance.
type SSTUserPropertiesCollection struct {
	Sst   []SSTUserProperties `protobuf:"bytes,1,rep,name=sst" json:"sst"`
	Error string              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *SSTUserPropertiesCollection) Reset()         { *m = SSTUserPropertiesCollection{} }
func (m *SSTUserPropertiesCollection) String() string { return proto.CompactTextString(m) }
func (*SSTUserPropertiesCollection) ProtoMessage()    {}
func (*SSTUserPropertiesCollection) Descriptor() ([]byte, []int) {
	return fileDescriptorRocksdb, []int{1}
}

func init() {
	proto.RegisterType((*SSTUserProperties)(nil), "cockroach.storage.engine.enginepb.SSTUserProperties")
	proto.RegisterType((*SSTUserPropertiesCollection)(nil), "cockroach.storage.engine.enginepb.SSTUserPropertiesCollection")
}
func (m *SSTUserProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserProperties) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.TsMin != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMin.Size()))
		n1, err := m.TsMin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TsMax != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMax.Size()))
		n2, err := m.TsMax.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SSTUserPropertiesCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserPropertiesCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, msg := range m.Sst {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func encodeVarintRocksdb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SSTUserProperties) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMin != nil {
		l = m.TsMin.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMax != nil {
		l = m.TsMax.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func (m *SSTUserPropertiesCollection) Size() (n int) {
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, e := range m.Sst {
			l = e.Size()
			n += 1 + l + sovRocksdb(uint64(l))
		}
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func sovRocksdb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRocksdb(x uint64) (n int) {
	return sovRocksdb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SSTUserProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
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
			return fmt.Errorf("proto: SSTUserProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMin == nil {
				m.TsMin = &cockroach_util_hlc.Timestamp{}
			}
			if err := m.TsMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMax", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMax == nil {
				m.TsMax = &cockroach_util_hlc.Timestamp{}
			}
			if err := m.TsMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRocksdb
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
func (m *SSTUserPropertiesCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
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
			return fmt.Errorf("proto: SSTUserPropertiesCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserPropertiesCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sst", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sst = append(m.Sst, SSTUserProperties{})
			if err := m.Sst[len(m.Sst)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
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
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRocksdb
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
func skipRocksdb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRocksdb
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
					return 0, ErrIntOverflowRocksdb
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
					return 0, ErrIntOverflowRocksdb
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
				return 0, ErrInvalidLengthRocksdb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRocksdb
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
				next, err := skipRocksdb(dAtA[start:])
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
	ErrInvalidLengthRocksdb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRocksdb   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("cockroach/pkg/storage/engine/enginepb/rocksdb.proto", fileDescriptorRocksdb)
}

var fileDescriptorRocksdb = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbb, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xeb, 0xbf, 0x07, 0xfd, 0x75, 0x27, 0xac, 0x0e, 0x51, 0x11, 0xa6, 0x74, 0x40, 0x9d,
	0x6c, 0xa9, 0xed, 0x15, 0x94, 0x15, 0x24, 0x94, 0x96, 0x85, 0x05, 0xb9, 0xc6, 0x4a, 0xac, 0x26,
	0xf9, 0x2c, 0xdb, 0x48, 0x5d, 0xb9, 0x03, 0x24, 0x6e, 0x2a, 0x23, 0x23, 0x13, 0x82, 0x70, 0x23,
	0x28, 0x09, 0x01, 0xa1, 0x0e, 0x30, 0xf9, 0xb5, 0xe5, 0xe7, 0xf9, 0x0e, 0x78, 0x2e, 0x41, 0x6e,
	0x2d, 0x08, 0x19, 0x73, 0xb3, 0x8d, 0xb8, 0xf3, 0x60, 0x45, 0xa4, 0xb8, 0xca, 0x22, 0x9d, 0x35,
	0x87, 0xd9, 0x70, 0x0b, 0x72, 0xeb, 0x6e, 0x37, 0xcc, 0x58, 0xf0, 0x40, 0x4e, 0xbe, 0x20, 0xf6,
	0x09, 0xb0, 0xfa, 0x27, 0x6b, 0x80, 0xd1, 0xe9, 0x4f, 0xef, 0x9d, 0xd7, 0x09, 0x8f, 0x13, 0xc9,
	0xbd, 0x4e, 0x95, 0xf3, 0x22, 0x35, 0xb5, 0x6a, 0x34, 0x8c, 0x20, 0x82, 0x2a, 0xf2, 0x32, 0xd5,
	0xaf, 0x93, 0x47, 0x84, 0x0f, 0x56, 0xab, 0xf5, 0x95, 0x53, 0xf6, 0xd2, 0x82, 0x51, 0xd6, 0x6b,
	0xe5, 0x08, 0xc1, 0x1d, 0x23, 0x7c, 0x1c, 0xa0, 0x31, 0x9a, 0xf6, 0xc3, 0x2a, 0x93, 0x05, 0xee,
	0x79, 0x77, 0x93, 0xea, 0x2c, 0xf8, 0x37, 0x46, 0xd3, 0xc1, 0xec, 0x88, 0x7d, 0xf7, 0x56, 0x16,
	0x65, 0x71, 0x22, 0xd9, 0xba, 0x29, 0x1a, 0x76, 0xbd, 0xbb, 0xd0, 0x59, 0x43, 0x89, 0x5d, 0xd0,
	0xfe, 0x2b, 0x25, 0x76, 0x93, 0x7b, 0x84, 0x0f, 0xf7, 0xba, 0x3a, 0x83, 0x24, 0x51, 0xd2, 0x6b,
	0xc8, 0xc8, 0x39, 0x6e, 0x3b, 0xe7, 0x03, 0x34, 0x6e, 0x4f, 0x07, 0xb3, 0x05, 0xfb, 0x75, 0x49,
	0x6c, 0x4f, 0xb6, 0xec, 0xe4, 0x2f, 0xc7, 0xad, 0xb0, 0xd4, 0x90, 0x21, 0xee, 0x2a, 0x6b, 0xc1,
	0x56, 0x83, 0xf5, 0xc3, 0xfa, 0xb2, 0x9c, 0xe4, 0x6f, 0xb4, 0x95, 0x17, 0x14, 0x3d, 0x15, 0x14,
	0x3d, 0x17, 0x14, 0xbd, 0x16, 0x14, 0x3d, 0xbc, 0xd3, 0xd6, 0xf5, 0xff, 0x46, 0xbb, 0xe9, 0x55,
	0x4b, 0x9c, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20, 0x6d, 0x5c, 0x37, 0xdc, 0x01, 0x00, 0x00,
}
