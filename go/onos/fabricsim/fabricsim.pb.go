// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/fabricsim/fabricsim.proto

package fabricsim

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetIOStatsRequest struct {
}

func (m *GetIOStatsRequest) Reset()         { *m = GetIOStatsRequest{} }
func (m *GetIOStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetIOStatsRequest) ProtoMessage()    {}
func (*GetIOStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f82323040cd66729, []int{0}
}
func (m *GetIOStatsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetIOStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetIOStatsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetIOStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIOStatsRequest.Merge(m, src)
}
func (m *GetIOStatsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetIOStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIOStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIOStatsRequest proto.InternalMessageInfo

type GetIOStatsResponse struct {
	Stats []*IOStats `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (m *GetIOStatsResponse) Reset()         { *m = GetIOStatsResponse{} }
func (m *GetIOStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetIOStatsResponse) ProtoMessage()    {}
func (*GetIOStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f82323040cd66729, []int{1}
}
func (m *GetIOStatsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetIOStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetIOStatsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetIOStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIOStatsResponse.Merge(m, src)
}
func (m *GetIOStatsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetIOStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIOStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIOStatsResponse proto.InternalMessageInfo

func (m *GetIOStatsResponse) GetStats() []*IOStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*GetIOStatsRequest)(nil), "onos.fabricsim.GetIOStatsRequest")
	proto.RegisterType((*GetIOStatsResponse)(nil), "onos.fabricsim.GetIOStatsResponse")
}

func init() { proto.RegisterFile("onos/fabricsim/fabricsim.proto", fileDescriptor_f82323040cd66729) }

var fileDescriptor_f82323040cd66729 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcf, 0xcb, 0x2f,
	0xd6, 0x4f, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0x2e, 0xce, 0xcc, 0x45, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xf8, 0x40, 0xf2, 0x7a, 0x70, 0x51, 0x29, 0x19, 0x34, 0xf5, 0x29, 0xa9, 0x65,
	0x99, 0xc9, 0xa9, 0xc5, 0x10, 0xd5, 0x4a, 0xc2, 0x5c, 0x82, 0xee, 0xa9, 0x25, 0x9e, 0xfe, 0xc1,
	0x25, 0x89, 0x25, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xce, 0x5c, 0x42, 0xc8,
	0x82, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xba, 0x5c, 0xac, 0xc5, 0x20, 0x01, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x71, 0x3d, 0x54, 0x8b, 0xf4, 0x60, 0xea, 0x21, 0xaa, 0x8c, 0xd2,
	0xb8, 0xf8, 0xdd, 0xc0, 0x72, 0xc1, 0x99, 0xb9, 0xa5, 0x39, 0x89, 0x25, 0xf9, 0x45, 0x42, 0xc1,
	0x5c, 0x5c, 0x08, 0x73, 0x85, 0x14, 0xd1, 0x0d, 0xc0, 0x70, 0x88, 0x94, 0x12, 0x3e, 0x25, 0x10,
	0x67, 0x39, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8,
	0x8b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xce, 0x2f, 0x9b, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabricSimulatorClient is the client API for FabricSimulator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabricSimulatorClient interface {
	// GetIOStats returns list of aggregate I/O stats.
	GetIOStats(ctx context.Context, in *GetIOStatsRequest, opts ...grpc.CallOption) (*GetIOStatsResponse, error)
}

type fabricSimulatorClient struct {
	cc *grpc.ClientConn
}

func NewFabricSimulatorClient(cc *grpc.ClientConn) FabricSimulatorClient {
	return &fabricSimulatorClient{cc}
}

func (c *fabricSimulatorClient) GetIOStats(ctx context.Context, in *GetIOStatsRequest, opts ...grpc.CallOption) (*GetIOStatsResponse, error) {
	out := new(GetIOStatsResponse)
	err := c.cc.Invoke(ctx, "/onos.fabricsim.FabricSimulator/GetIOStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FabricSimulatorServer is the server API for FabricSimulator service.
type FabricSimulatorServer interface {
	// GetIOStats returns list of aggregate I/O stats.
	GetIOStats(context.Context, *GetIOStatsRequest) (*GetIOStatsResponse, error)
}

// UnimplementedFabricSimulatorServer can be embedded to have forward compatible implementations.
type UnimplementedFabricSimulatorServer struct {
}

func (*UnimplementedFabricSimulatorServer) GetIOStats(ctx context.Context, req *GetIOStatsRequest) (*GetIOStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIOStats not implemented")
}

func RegisterFabricSimulatorServer(s *grpc.Server, srv FabricSimulatorServer) {
	s.RegisterService(&_FabricSimulator_serviceDesc, srv)
}

func _FabricSimulator_GetIOStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIOStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricSimulatorServer).GetIOStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.fabricsim.FabricSimulator/GetIOStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricSimulatorServer).GetIOStats(ctx, req.(*GetIOStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FabricSimulator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.fabricsim.FabricSimulator",
	HandlerType: (*FabricSimulatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIOStats",
			Handler:    _FabricSimulator_GetIOStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onos/fabricsim/fabricsim.proto",
}

func (m *GetIOStatsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetIOStatsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetIOStatsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetIOStatsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetIOStatsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetIOStatsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for iNdEx := len(m.Stats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFabricsim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintFabricsim(dAtA []byte, offset int, v uint64) int {
	offset -= sovFabricsim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetIOStatsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetIOStatsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovFabricsim(uint64(l))
		}
	}
	return n
}

func sovFabricsim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFabricsim(x uint64) (n int) {
	return sovFabricsim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetIOStatsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFabricsim
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
			return fmt.Errorf("proto: GetIOStatsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetIOStatsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFabricsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFabricsim
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
func (m *GetIOStatsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFabricsim
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
			return fmt.Errorf("proto: GetIOStatsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetIOStatsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFabricsim
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
				return ErrInvalidLengthFabricsim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFabricsim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, &IOStats{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFabricsim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFabricsim
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
func skipFabricsim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFabricsim
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
					return 0, ErrIntOverflowFabricsim
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
					return 0, ErrIntOverflowFabricsim
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
				return 0, ErrInvalidLengthFabricsim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFabricsim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFabricsim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFabricsim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFabricsim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFabricsim = fmt.Errorf("proto: unexpected end of group")
)
