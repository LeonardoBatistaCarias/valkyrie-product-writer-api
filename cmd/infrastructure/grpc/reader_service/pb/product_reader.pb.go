// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: reader_service/proto/product_reader.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_reader_service_proto_product_reader_proto protoreflect.FileDescriptor

var file_reader_service_proto_product_reader_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x32, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9,
	0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x15, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_reader_service_proto_product_reader_proto_goTypes = []interface{}{
	(*CreateProductReq)(nil),     // 0: CreateProductReq
	(*DeleteProductByIDReq)(nil), // 1: DeleteProductByIDReq
	(*UpdateProductByIDReq)(nil), // 2: UpdateProductByIDReq
	(*CreateProductRes)(nil),     // 3: CreateProductRes
	(*DeleteProductByIDRes)(nil), // 4: DeleteProductByIDRes
	(*UpdateProductByIDRes)(nil), // 5: UpdateProductByIDRes
}
var file_reader_service_proto_product_reader_proto_depIdxs = []int32{
	0, // 0: productReader.ProductReaderService.CreateProduct:input_type -> CreateProductReq
	1, // 1: productReader.ProductReaderService.DeleteProductByID:input_type -> DeleteProductByIDReq
	2, // 2: productReader.ProductReaderService.UpdateProductByID:input_type -> UpdateProductByIDReq
	3, // 3: productReader.ProductReaderService.CreateProduct:output_type -> CreateProductRes
	4, // 4: productReader.ProductReaderService.DeleteProductByID:output_type -> DeleteProductByIDRes
	5, // 5: productReader.ProductReaderService.UpdateProductByID:output_type -> UpdateProductByIDRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reader_service_proto_product_reader_proto_init() }
func file_reader_service_proto_product_reader_proto_init() {
	if File_reader_service_proto_product_reader_proto != nil {
		return
	}
	file_reader_service_proto_product_reader_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reader_service_proto_product_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reader_service_proto_product_reader_proto_goTypes,
		DependencyIndexes: file_reader_service_proto_product_reader_proto_depIdxs,
	}.Build()
	File_reader_service_proto_product_reader_proto = out.File
	file_reader_service_proto_product_reader_proto_rawDesc = nil
	file_reader_service_proto_product_reader_proto_goTypes = nil
	file_reader_service_proto_product_reader_proto_depIdxs = nil
}
