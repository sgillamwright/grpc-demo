// Code generated by protoc-gen-go. DO NOT EDIT.
// source: books.proto

/*
Package Books is a generated protocol buffer package.

Package Section [Optional]

It is generated from these files:
	books.proto

It has these top-level messages:
	EmptyMessage
	BookIDRequest
	BookListResponse
	Book
	Genre
	Author
*/
package Books

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Book_Published int32

const (
	Book_UNPUBLISHED Book_Published = 0
	Book_PUBLISHED   Book_Published = 1
)

var Book_Published_name = map[int32]string{
	0: "UNPUBLISHED",
	1: "PUBLISHED",
}
var Book_Published_value = map[string]int32{
	"UNPUBLISHED": 0,
	"PUBLISHED":   1,
}

func (x Book_Published) String() string {
	return proto.EnumName(Book_Published_name, int32(x))
}
func (Book_Published) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// Message for handling an empty request
type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Message for handling only sending the ID
type BookIDRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *BookIDRequest) Reset()                    { *m = BookIDRequest{} }
func (m *BookIDRequest) String() string            { return proto.CompactTextString(m) }
func (*BookIDRequest) ProtoMessage()               {}
func (*BookIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BookIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Message for handling a Book list response
type BookListResponse struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *BookListResponse) Reset()                    { *m = BookListResponse{} }
func (m *BookListResponse) String() string            { return proto.CompactTextString(m) }
func (*BookListResponse) ProtoMessage()               {}
func (*BookListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BookListResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

// Message Type for handling a single Book
type Book struct {
	Id        string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Genres    []*Genre       `protobuf:"bytes,3,rep,name=genres" json:"genres,omitempty"`
	Author    *Author        `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Published Book_Published `protobuf:"varint,5,opt,name=published,enum=Books.Book_Published" json:"published,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetGenres() []*Genre {
	if m != nil {
		return m.Genres
	}
	return nil
}

func (m *Book) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Book) GetPublished() Book_Published {
	if m != nil {
		return m.Published
	}
	return Book_UNPUBLISHED
}

// Message Type for handling a single genre
type Genre struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Genre) Reset()                    { *m = Genre{} }
func (m *Genre) String() string            { return proto.CompactTextString(m) }
func (*Genre) ProtoMessage()               {}
func (*Genre) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Genre) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Genre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Message Type for handling a single author
type Author struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Website     string `protobuf:"bytes,3,opt,name=website" json:"website,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *Author) Reset()                    { *m = Author{} }
func (m *Author) String() string            { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()               {}
func (*Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Author) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Author) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "Books.EmptyMessage")
	proto.RegisterType((*BookIDRequest)(nil), "Books.BookIDRequest")
	proto.RegisterType((*BookListResponse)(nil), "Books.BookListResponse")
	proto.RegisterType((*Book)(nil), "Books.Book")
	proto.RegisterType((*Genre)(nil), "Books.Genre")
	proto.RegisterType((*Author)(nil), "Books.Author")
	proto.RegisterEnum("Books.Book_Published", Book_Published_name, Book_Published_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BooksService service

type BooksServiceClient interface {
	// Get All Books
	GetAllBooks(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*BookListResponse, error)
	// Stream of books
	GetAllBooksStream(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (BooksService_GetAllBooksStreamClient, error)
	// Get Single Book
	GetOneBook(ctx context.Context, in *BookIDRequest, opts ...grpc.CallOption) (*Book, error)
	// Add Book
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	// Remove Book
	RemoveBook(ctx context.Context, in *BookIDRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type booksServiceClient struct {
	cc *grpc.ClientConn
}

func NewBooksServiceClient(cc *grpc.ClientConn) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) GetAllBooks(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*BookListResponse, error) {
	out := new(BookListResponse)
	err := grpc.Invoke(ctx, "/Books.BooksService/getAllBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetAllBooksStream(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (BooksService_GetAllBooksStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BooksService_serviceDesc.Streams[0], c.cc, "/Books.BooksService/getAllBooksStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &booksServiceGetAllBooksStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BooksService_GetAllBooksStreamClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type booksServiceGetAllBooksStreamClient struct {
	grpc.ClientStream
}

func (x *booksServiceGetAllBooksStreamClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *booksServiceClient) GetOneBook(ctx context.Context, in *BookIDRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/Books.BooksService/getOneBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/Books.BooksService/addBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) RemoveBook(ctx context.Context, in *BookIDRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/Books.BooksService/removeBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BooksService service

type BooksServiceServer interface {
	// Get All Books
	GetAllBooks(context.Context, *EmptyMessage) (*BookListResponse, error)
	// Stream of books
	GetAllBooksStream(*EmptyMessage, BooksService_GetAllBooksStreamServer) error
	// Get Single Book
	GetOneBook(context.Context, *BookIDRequest) (*Book, error)
	// Add Book
	AddBook(context.Context, *Book) (*Book, error)
	// Remove Book
	RemoveBook(context.Context, *BookIDRequest) (*EmptyMessage, error)
}

func RegisterBooksServiceServer(s *grpc.Server, srv BooksServiceServer) {
	s.RegisterService(&_BooksService_serviceDesc, srv)
}

func _BooksService_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Books.BooksService/GetAllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetAllBooks(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetAllBooksStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BooksServiceServer).GetAllBooksStream(m, &booksServiceGetAllBooksStreamServer{stream})
}

type BooksService_GetAllBooksStreamServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type booksServiceGetAllBooksStreamServer struct {
	grpc.ServerStream
}

func (x *booksServiceGetAllBooksStreamServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

func _BooksService_GetOneBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetOneBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Books.BooksService/GetOneBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetOneBook(ctx, req.(*BookIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Books.BooksService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_RemoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).RemoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Books.BooksService/RemoveBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).RemoveBook(ctx, req.(*BookIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BooksService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Books.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAllBooks",
			Handler:    _BooksService_GetAllBooks_Handler,
		},
		{
			MethodName: "getOneBook",
			Handler:    _BooksService_GetOneBook_Handler,
		},
		{
			MethodName: "addBook",
			Handler:    _BooksService_AddBook_Handler,
		},
		{
			MethodName: "removeBook",
			Handler:    _BooksService_RemoveBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAllBooksStream",
			Handler:       _BooksService_GetAllBooksStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "books.proto",
}

func init() { proto.RegisterFile("books.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0xdd, 0xd9, 0x8f, 0x94, 0xdc, 0xec, 0xae, 0xf5, 0xaa, 0x38, 0xf8, 0x62, 0x1c, 0x2c, 0x04,
	0x0a, 0x8b, 0xa6, 0xf8, 0xd0, 0x07, 0x1f, 0x5a, 0x5a, 0xb4, 0x50, 0xb5, 0xcc, 0xd2, 0x1f, 0x90,
	0x34, 0x97, 0xec, 0xe0, 0x26, 0x13, 0x33, 0xb3, 0x15, 0x7f, 0x8c, 0x3f, 0xcc, 0x7f, 0x23, 0x99,
	0xa6, 0xed, 0xb0, 0x94, 0xd2, 0x97, 0x90, 0x7b, 0xce, 0x99, 0xc3, 0x9d, 0x73, 0xef, 0x40, 0x94,
	0x6b, 0xfd, 0xd3, 0x2c, 0x9a, 0x56, 0x5b, 0x8d, 0x93, 0xe3, 0xae, 0x10, 0x73, 0x98, 0x9e, 0x56,
	0x8d, 0xfd, 0xf3, 0x8d, 0x8c, 0xc9, 0x4a, 0x12, 0x6f, 0x61, 0xd6, 0x11, 0x67, 0x27, 0x92, 0x7e,
	0x6d, 0xc8, 0x58, 0x9c, 0xc3, 0x50, 0x15, 0x9c, 0xc5, 0x2c, 0x09, 0xe5, 0x50, 0x15, 0xe2, 0x13,
	0xec, 0x76, 0x82, 0x73, 0x65, 0xac, 0x24, 0xd3, 0xe8, 0xda, 0x10, 0xbe, 0x83, 0x89, 0xb3, 0xe6,
	0x2c, 0x1e, 0x25, 0x51, 0x1a, 0x2d, 0x9c, 0xb7, 0xfb, 0xca, 0x1b, 0x46, 0xfc, 0x63, 0x30, 0xee,
	0xea, 0x6d, 0x3f, 0x44, 0x18, 0xd7, 0x59, 0x45, 0x7c, 0xe8, 0x10, 0xf7, 0x8f, 0xef, 0x21, 0x28,
	0xa9, 0x6e, 0xc9, 0xf0, 0x91, 0x33, 0x9c, 0xf6, 0x86, 0x5f, 0x3a, 0x50, 0xf6, 0x1c, 0xee, 0x41,
	0x90, 0x6d, 0xec, 0x4a, 0xb7, 0x7c, 0x1c, 0xb3, 0x24, 0x4a, 0x67, 0xbd, 0xea, 0xc8, 0x81, 0xb2,
	0x27, 0xf1, 0x00, 0xc2, 0x66, 0x93, 0xaf, 0x95, 0x59, 0x51, 0xc1, 0x27, 0x31, 0x4b, 0xe6, 0xe9,
	0x2b, 0xaf, 0xc1, 0xc5, 0xc5, 0x2d, 0x29, 0xef, 0x75, 0x62, 0x1f, 0xc2, 0x3b, 0x1c, 0x9f, 0x41,
	0x74, 0xf9, 0xfd, 0xe2, 0xf2, 0xf8, 0xfc, 0x6c, 0xf9, 0xf5, 0xf4, 0x64, 0x77, 0x80, 0x33, 0x08,
	0xef, 0x4b, 0x26, 0xf6, 0x61, 0xe2, 0x3a, 0x7b, 0xca, 0xdd, 0xc4, 0x0a, 0x82, 0x9b, 0x06, 0x9f,
	0x94, 0x04, 0x87, 0x9d, 0xdf, 0x94, 0x1b, 0x65, 0x89, 0x8f, 0x1c, 0x7c, 0x5b, 0x62, 0x0c, 0x51,
	0x41, 0xe6, 0xaa, 0x55, 0x8d, 0x55, 0xba, 0x76, 0x11, 0x84, 0xd2, 0x87, 0xd2, 0xbf, 0x43, 0x98,
	0xba, 0x7b, 0x2e, 0xa9, 0xbd, 0x56, 0x57, 0x84, 0x9f, 0x21, 0x2a, 0xc9, 0x1e, 0xad, 0xd7, 0x0e,
	0xc5, 0x17, 0x7d, 0x0a, 0xfe, 0xfc, 0xdf, 0xbc, 0xf6, 0xa2, 0xf1, 0x67, 0x2c, 0x06, 0x78, 0x08,
	0xcf, 0xbd, 0xe3, 0x4b, 0xdb, 0x52, 0x56, 0x3d, 0x6c, 0xe2, 0x2f, 0x80, 0x18, 0x7c, 0x60, 0xf8,
	0x11, 0xa0, 0x24, 0xfb, 0xa3, 0x26, 0xb7, 0x02, 0x2f, 0x3d, 0xfa, 0x6e, 0xd1, 0xb6, 0x0e, 0xe1,
	0x1e, 0xec, 0x64, 0x45, 0xe1, 0xf4, 0x3e, 0xb3, 0x2d, 0x3b, 0x04, 0x68, 0xa9, 0xd2, 0xd7, 0x8f,
	0x39, 0x3f, 0xd4, 0xa3, 0x18, 0xe4, 0x81, 0x7b, 0x08, 0x07, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x85, 0xa8, 0x2b, 0xa7, 0x17, 0x03, 0x00, 0x00,
}