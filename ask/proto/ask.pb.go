// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/ask.proto

package ask

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

// Ask a question and receive an instant answer
type QuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the question to ask
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QuestionRequest) Reset() {
	*x = QuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionRequest) ProtoMessage() {}

func (x *QuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionRequest.ProtoReflect.Descriptor instead.
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ask_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type QuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the answer to your question
	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	// a related url
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// any related image
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *QuestionResponse) Reset() {
	*x = QuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ask_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionResponse) ProtoMessage() {}

func (x *QuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ask_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionResponse.ProtoReflect.Descriptor instead.
func (*QuestionResponse) Descriptor() ([]byte, []int) {
	return file_proto_ask_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *QuestionResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *QuestionResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_proto_ask_proto protoreflect.FileDescriptor

var file_proto_ask_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x73, 0x6b, 0x22, 0x27, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x52, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x32, 0x40, 0x0a, 0x03, 0x41, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ask_proto_rawDescOnce sync.Once
	file_proto_ask_proto_rawDescData = file_proto_ask_proto_rawDesc
)

func file_proto_ask_proto_rawDescGZIP() []byte {
	file_proto_ask_proto_rawDescOnce.Do(func() {
		file_proto_ask_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ask_proto_rawDescData)
	})
	return file_proto_ask_proto_rawDescData
}

var file_proto_ask_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ask_proto_goTypes = []interface{}{
	(*QuestionRequest)(nil),  // 0: ask.QuestionRequest
	(*QuestionResponse)(nil), // 1: ask.QuestionResponse
}
var file_proto_ask_proto_depIdxs = []int32{
	0, // 0: ask.Ask.Question:input_type -> ask.QuestionRequest
	1, // 1: ask.Ask.Question:output_type -> ask.QuestionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ask_proto_init() }
func file_proto_ask_proto_init() {
	if File_proto_ask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionRequest); i {
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
		file_proto_ask_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionResponse); i {
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
			RawDescriptor: file_proto_ask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ask_proto_goTypes,
		DependencyIndexes: file_proto_ask_proto_depIdxs,
		MessageInfos:      file_proto_ask_proto_msgTypes,
	}.Build()
	File_proto_ask_proto = out.File
	file_proto_ask_proto_rawDesc = nil
	file_proto_ask_proto_goTypes = nil
	file_proto_ask_proto_depIdxs = nil
}