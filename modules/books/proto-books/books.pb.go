// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: modules/books/proto-books/books.proto

package proto_books

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *int64  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name   *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	BookId *string `protobuf:"bytes,3,opt,name=bookId,proto3,oneof" json:"bookId,omitempty"`
}

func (x *BookItem) Reset() {
	*x = BookItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookItem) ProtoMessage() {}

func (x *BookItem) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookItem.ProtoReflect.Descriptor instead.
func (*BookItem) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{0}
}

func (x *BookItem) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BookItem) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BookItem) GetBookId() string {
	if x != nil && x.BookId != nil {
		return *x.BookId
	}
	return ""
}

type BookItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BookItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BookItems) Reset() {
	*x = BookItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookItems) ProtoMessage() {}

func (x *BookItems) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookItems.ProtoReflect.Descriptor instead.
func (*BookItems) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{1}
}

func (x *BookItems) GetData() []*BookItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type BookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *string    `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	StatusCode *int64     `protobuf:"varint,2,opt,name=statusCode,proto3,oneof" json:"statusCode,omitempty"`
	Message    *string    `protobuf:"bytes,3,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Data       *BookItems `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{2}
}

func (x *BookResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *BookResponse) GetStatusCode() int64 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *BookResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *BookResponse) GetData() *BookItems {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	BookId *string `protobuf:"bytes,2,opt,name=bookId,proto3,oneof" json:"bookId,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateBookRequest) GetBookId() string {
	if x != nil && x.BookId != nil {
		return *x.BookId
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookData *BookItem `protobuf:"bytes,1,opt,name=bookData,proto3,oneof" json:"bookData,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookRequest) GetBookData() *BookItem {
	if x != nil {
		return x.BookData
	}
	return nil
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId *int64 `protobuf:"varint,1,opt,name=bookId,proto3,oneof" json:"bookId,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBookRequest) GetBookId() int64 {
	if x != nil && x.BookId != nil {
		return *x.BookId
	}
	return 0
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_books_proto_books_books_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_books_proto_books_books_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_modules_books_proto_books_books_proto_rawDescGZIP(), []int{6}
}

var File_modules_books_proto_books_books_proto protoreflect.FileDescriptor

var file_modules_books_proto_books_books_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x09, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x48, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xcc, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x31,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_books_proto_books_books_proto_rawDescOnce sync.Once
	file_modules_books_proto_books_books_proto_rawDescData = file_modules_books_proto_books_books_proto_rawDesc
)

func file_modules_books_proto_books_books_proto_rawDescGZIP() []byte {
	file_modules_books_proto_books_books_proto_rawDescOnce.Do(func() {
		file_modules_books_proto_books_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_books_proto_books_books_proto_rawDescData)
	})
	return file_modules_books_proto_books_books_proto_rawDescData
}

var file_modules_books_proto_books_books_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_modules_books_proto_books_books_proto_goTypes = []interface{}{
	(*BookItem)(nil),          // 0: BookItem
	(*BookItems)(nil),         // 1: BookItems
	(*BookResponse)(nil),      // 2: BookResponse
	(*CreateBookRequest)(nil), // 3: CreateBookRequest
	(*UpdateBookRequest)(nil), // 4: UpdateBookRequest
	(*DeleteBookRequest)(nil), // 5: DeleteBookRequest
	(*EmptyRequest)(nil),      // 6: EmptyRequest
}
var file_modules_books_proto_books_books_proto_depIdxs = []int32{
	0, // 0: BookItems.data:type_name -> BookItem
	1, // 1: BookResponse.data:type_name -> BookItems
	0, // 2: UpdateBookRequest.bookData:type_name -> BookItem
	3, // 3: Book.CreateBook:input_type -> CreateBookRequest
	4, // 4: Book.UpdateBook:input_type -> UpdateBookRequest
	6, // 5: Book.ListBooks:input_type -> EmptyRequest
	5, // 6: Book.DeleteBook:input_type -> DeleteBookRequest
	2, // 7: Book.CreateBook:output_type -> BookResponse
	2, // 8: Book.UpdateBook:output_type -> BookResponse
	2, // 9: Book.ListBooks:output_type -> BookResponse
	2, // 10: Book.DeleteBook:output_type -> BookResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_modules_books_proto_books_books_proto_init() }
func file_modules_books_proto_books_books_proto_init() {
	if File_modules_books_proto_books_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_books_proto_books_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookItems); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_modules_books_proto_books_books_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_modules_books_proto_books_books_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_modules_books_proto_books_books_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_modules_books_proto_books_books_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_modules_books_proto_books_books_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_modules_books_proto_books_books_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_books_proto_books_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_books_proto_books_books_proto_goTypes,
		DependencyIndexes: file_modules_books_proto_books_books_proto_depIdxs,
		MessageInfos:      file_modules_books_proto_books_books_proto_msgTypes,
	}.Build()
	File_modules_books_proto_books_books_proto = out.File
	file_modules_books_proto_books_books_proto_rawDesc = nil
	file_modules_books_proto_books_books_proto_goTypes = nil
	file_modules_books_proto_books_books_proto_depIdxs = nil
}
