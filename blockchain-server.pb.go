// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockchain-server.proto

// The package name determines the name of the directories that truss creates

package blockchain

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/metaverse/truss/deftree/googlethirdparty"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	In string `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6c9083788a8f1d, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return m.Size()
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

type EchoResponse struct {
	Out string `protobuf:"bytes,1,opt,name=Out,proto3" json:"Out,omitempty"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6c9083788a8f1d, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return m.Size()
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "blockchain.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "blockchain.EchoResponse")
}

func init() { proto.RegisterFile("blockchain-server.proto", fileDescriptor_6d6c9083788a8f1d) }

var fileDescriptor_6d6c9083788a8f1d = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0xa8, 0x45, 0x57, 0x11, 0xdd, 0x4b, 0x42, 0xc1, 0xa5, 0xe4, 0xe4, 0xc5, 0x2c,
	0xe8, 0x1b, 0x14, 0x7a, 0xe8, 0x49, 0xc8, 0xd1, 0x93, 0x9b, 0xed, 0x98, 0x0d, 0xb6, 0x3b, 0x71,
	0x77, 0x52, 0xf0, 0xea, 0x13, 0x08, 0xbe, 0x94, 0xc7, 0x82, 0x17, 0x8f, 0x92, 0xf8, 0x20, 0xd2,
	0x46, 0xac, 0xd0, 0xdb, 0x0f, 0xdf, 0xcf, 0x3f, 0xdf, 0xb0, 0xb8, 0x98, 0xa3, 0x7e, 0xd4, 0x46,
	0x55, 0xf6, 0xca, 0x83, 0x5b, 0x82, 0xcb, 0x6a, 0x87, 0x84, 0x9c, 0x6d, 0xc1, 0x70, 0x52, 0x56,
	0x64, 0x9a, 0x22, 0xd3, 0xb8, 0x90, 0x0b, 0x20, 0xb5, 0x04, 0xe7, 0x41, 0x92, 0x6b, 0xbc, 0x97,
	0x33, 0x78, 0x20, 0x07, 0x20, 0x4b, 0xc4, 0x72, 0x0e, 0x64, 0x2a, 0x37, 0xab, 0x95, 0xa3, 0x67,
	0xa9, 0xac, 0x45, 0x52, 0x54, 0xa1, 0xf5, 0xfd, 0x64, 0x7a, 0xc1, 0x8e, 0x27, 0xda, 0x60, 0x0e,
	0x4f, 0x0d, 0x78, 0xe2, 0xa7, 0x2c, 0x9a, 0xda, 0x24, 0x1c, 0x85, 0x97, 0x47, 0x79, 0x34, 0xb5,
	0xe9, 0x88, 0x9d, 0xf4, 0xd8, 0xd7, 0x68, 0x3d, 0xf0, 0x33, 0xb6, 0x77, 0xdb, 0xd0, 0x6f, 0x61,
	0x1d, 0xaf, 0xef, 0x19, 0x1b, 0xff, 0x59, 0xf1, 0x9c, 0xed, 0xaf, 0xfb, 0x3c, 0xce, 0xb6, 0xaa,
	0xd9, 0xbf, 0x03, 0xc3, 0x64, 0x17, 0xf4, 0xd3, 0x69, 0xfc, 0xf2, 0xf1, 0xfd, 0x16, 0x9d, 0xf3,
	0x03, 0x09, 0xda, 0xe0, 0xdd, 0x21, 0x1f, 0x6c, 0x82, 0x1c, 0x27, 0xef, 0xad, 0x08, 0x57, 0xad,
	0x08, 0xbf, 0x5a, 0x11, 0xbe, 0x76, 0x22, 0x58, 0x75, 0x22, 0xf8, 0xec, 0x44, 0x50, 0x0c, 0x36,
	0x3f, 0xdc, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x58, 0xef, 0x07, 0x1f, 0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainClient interface {
	// Echo "echos" the incoming string
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/blockchain.Blockchain/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
type BlockchainServer interface {
	// Echo "echos" the incoming string
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedBlockchainServer can be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (*UnimplementedBlockchainServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.Blockchain/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blockchain.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Blockchain_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain-server.proto",
}

func (m *EchoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.In) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlockchainServer(dAtA, i, uint64(len(m.In)))
		i += copy(dAtA[i:], m.In)
	}
	return i, nil
}

func (m *EchoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Out) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlockchainServer(dAtA, i, uint64(len(m.Out)))
		i += copy(dAtA[i:], m.Out)
	}
	return i, nil
}

func encodeVarintBlockchainServer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EchoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.In)
	if l > 0 {
		n += 1 + l + sovBlockchainServer(uint64(l))
	}
	return n
}

func (m *EchoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Out)
	if l > 0 {
		n += 1 + l + sovBlockchainServer(uint64(l))
	}
	return n
}

func sovBlockchainServer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockchainServer(x uint64) (n int) {
	return sovBlockchainServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EchoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockchainServer
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
			return fmt.Errorf("proto: EchoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field In", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockchainServer
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
				return ErrInvalidLengthBlockchainServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockchainServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.In = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockchainServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockchainServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlockchainServer
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
func (m *EchoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockchainServer
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
			return fmt.Errorf("proto: EchoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Out", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockchainServer
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
				return ErrInvalidLengthBlockchainServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockchainServer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Out = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockchainServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockchainServer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlockchainServer
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
func skipBlockchainServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockchainServer
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
					return 0, ErrIntOverflowBlockchainServer
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
					return 0, ErrIntOverflowBlockchainServer
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
				return 0, ErrInvalidLengthBlockchainServer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBlockchainServer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlockchainServer
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
				next, err := skipBlockchainServer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBlockchainServer
				}
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
	ErrInvalidLengthBlockchainServer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockchainServer   = fmt.Errorf("proto: integer overflow")
)
