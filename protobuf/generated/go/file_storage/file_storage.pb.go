// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: file_storage/file_storage.proto

package filestorage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename      string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	FileExtension string `protobuf:"bytes,2,opt,name=fileExtension,proto3" json:"fileExtension,omitempty"`
	Chunk         []byte `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadRequest) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *UploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Extension string `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	Chunk     []byte `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DownloadResponse) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *DownloadResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type ShowFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID    string                 `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Filename  string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Extension string                 `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ShowFileResponse) Reset() {
	*x = ShowFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowFileResponse) ProtoMessage() {}

func (x *ShowFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowFileResponse.ProtoReflect.Descriptor instead.
func (*ShowFileResponse) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{4}
}

func (x *ShowFileResponse) GetFileID() string {
	if x != nil {
		return x.FileID
	}
	return ""
}

func (x *ShowFileResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ShowFileResponse) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *ShowFileResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ShowFileResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ShowFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*ShowFileResponse `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ShowFilesResponse) Reset() {
	*x = ShowFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_storage_file_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowFilesResponse) ProtoMessage() {}

func (x *ShowFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_storage_file_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowFilesResponse.ProtoReflect.Descriptor instead.
func (*ShowFilesResponse) Descriptor() ([]byte, []int) {
	return file_file_storage_file_storage_proto_rawDescGZIP(), []int{5}
}

func (x *ShowFilesResponse) GetFiles() []*ShowFileResponse {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_file_storage_file_storage_proto protoreflect.FileDescriptor

var file_file_storage_file_storage_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x28, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x29, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x10, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22,
	0xd8, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x68,
	0x6f, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x68,
	0x6f, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xeb, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4b,
	0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x09, 0x53,
	0x68, 0x6f, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x68, 0x6f, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x4b, 0x68, 0x6f, 0x72, 0x6b, 0x6f, 0x76, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_file_storage_file_storage_proto_rawDescOnce sync.Once
	file_file_storage_file_storage_proto_rawDescData = file_file_storage_file_storage_proto_rawDesc
)

func file_file_storage_file_storage_proto_rawDescGZIP() []byte {
	file_file_storage_file_storage_proto_rawDescOnce.Do(func() {
		file_file_storage_file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_storage_file_storage_proto_rawDescData)
	})
	return file_file_storage_file_storage_proto_rawDescData
}

var file_file_storage_file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_file_storage_file_storage_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),         // 0: file_storage.UploadRequest
	(*UploadResponse)(nil),        // 1: file_storage.UploadResponse
	(*DownloadRequest)(nil),       // 2: file_storage.DownloadRequest
	(*DownloadResponse)(nil),      // 3: file_storage.DownloadResponse
	(*ShowFileResponse)(nil),      // 4: file_storage.ShowFileResponse
	(*ShowFilesResponse)(nil),     // 5: file_storage.ShowFilesResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_file_storage_file_storage_proto_depIdxs = []int32{
	6, // 0: file_storage.ShowFileResponse.createdAt:type_name -> google.protobuf.Timestamp
	6, // 1: file_storage.ShowFileResponse.updatedAt:type_name -> google.protobuf.Timestamp
	4, // 2: file_storage.ShowFilesResponse.files:type_name -> file_storage.ShowFileResponse
	0, // 3: file_storage.FileService.Upload:input_type -> file_storage.UploadRequest
	2, // 4: file_storage.FileService.Download:input_type -> file_storage.DownloadRequest
	7, // 5: file_storage.FileService.ShowFiles:input_type -> google.protobuf.Empty
	1, // 6: file_storage.FileService.Upload:output_type -> file_storage.UploadResponse
	3, // 7: file_storage.FileService.Download:output_type -> file_storage.DownloadResponse
	5, // 8: file_storage.FileService.ShowFiles:output_type -> file_storage.ShowFilesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_file_storage_file_storage_proto_init() }
func file_file_storage_file_storage_proto_init() {
	if File_file_storage_file_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_storage_file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_file_storage_file_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_file_storage_file_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_file_storage_file_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
		file_file_storage_file_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowFileResponse); i {
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
		file_file_storage_file_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowFilesResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_storage_file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_storage_file_storage_proto_goTypes,
		DependencyIndexes: file_file_storage_file_storage_proto_depIdxs,
		MessageInfos:      file_file_storage_file_storage_proto_msgTypes,
	}.Build()
	File_file_storage_file_storage_proto = out.File
	file_file_storage_file_storage_proto_rawDesc = nil
	file_file_storage_file_storage_proto_goTypes = nil
	file_file_storage_file_storage_proto_depIdxs = nil
}
