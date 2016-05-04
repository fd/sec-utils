// Code generated by protoc-gen-gogo.
// source: limbo.services/core/runtime/limbo/annotations.proto
// DO NOT EDIT!

/*
Package limbo is a generated protocol buffer package.

It is generated from these files:
	limbo.services/core/runtime/limbo/annotations.proto
	limbo.services/core/runtime/limbo/authn.proto
	limbo.services/core/runtime/limbo/authz.proto
	limbo.services/core/runtime/limbo/http.proto
	limbo.services/core/runtime/limbo/sql.proto

It has these top-level messages:
	AuthnRule
	AuthzRule
	HttpRule
	ModelDescriptor
	ColumnDescriptor
	ScannerDescriptor
	JoinDescriptor
*/
package limbo

import proto "limbo.services/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "limbo.services/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

var E_Model = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*ModelDescriptor)(nil),
	Field:         51200,
	Name:          "limbo.model",
	Tag:           "bytes,51200,opt,name=model",
}

var E_Gosqlvaluer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51201,
	Name:          "limbo.gosqlvaluer",
	Tag:           "varint,51201,opt,name=gosqlvaluer",
}

var E_Column = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*ColumnDescriptor)(nil),
	Field:         51300,
	Name:          "limbo.column",
	Tag:           "bytes,51300,opt,name=column",
}

var E_Join = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*JoinDescriptor)(nil),
	Field:         51301,
	Name:          "limbo.join",
	Tag:           "bytes,51301,opt,name=join",
}

var E_HideInSwagger = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51302,
	Name:          "limbo.hideInSwagger",
	Tag:           "varint,51302,opt,name=hideInSwagger",
}

var E_Format = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51303,
	Name:          "limbo.format",
	Tag:           "bytes,51303,opt,name=format",
}

var E_Required = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51304,
	Name:          "limbo.required",
	Tag:           "varint,51304,opt,name=required",
}

var E_Pattern = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51305,
	Name:          "limbo.pattern",
	Tag:           "bytes,51305,opt,name=pattern",
}

var E_MinLength = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         51306,
	Name:          "limbo.minLength",
	Tag:           "varint,51306,opt,name=minLength",
}

var E_MaxLength = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         51307,
	Name:          "limbo.maxLength",
	Tag:           "varint,51307,opt,name=maxLength",
}

var E_MinItems = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         51308,
	Name:          "limbo.minItems",
	Tag:           "varint,51308,opt,name=minItems",
}

var E_MaxItems = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         51309,
	Name:          "limbo.maxItems",
	Tag:           "varint,51309,opt,name=maxItems",
}

var E_DefaultAuthn = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*AuthnRule)(nil),
	Field:         51400,
	Name:          "limbo.default_authn",
	Tag:           "bytes,51400,opt,name=default_authn,json=defaultAuthn",
}

var E_DefaultAuthz = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*AuthzRule)(nil),
	Field:         51401,
	Name:          "limbo.default_authz",
	Tag:           "bytes,51401,opt,name=default_authz,json=defaultAuthz",
}

var E_Authn = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*AuthnRule)(nil),
	Field:         51500,
	Name:          "limbo.authn",
	Tag:           "bytes,51500,opt,name=authn",
}

var E_Authz = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*AuthzRule)(nil),
	Field:         51501,
	Name:          "limbo.authz",
	Tag:           "bytes,51501,opt,name=authz",
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         51502,
	Name:          "limbo.http",
	Tag:           "bytes,51502,opt,name=http",
}

func init() {
	proto.RegisterExtension(E_Model)
	proto.RegisterExtension(E_Gosqlvaluer)
	proto.RegisterExtension(E_Column)
	proto.RegisterExtension(E_Join)
	proto.RegisterExtension(E_HideInSwagger)
	proto.RegisterExtension(E_Format)
	proto.RegisterExtension(E_Required)
	proto.RegisterExtension(E_Pattern)
	proto.RegisterExtension(E_MinLength)
	proto.RegisterExtension(E_MaxLength)
	proto.RegisterExtension(E_MinItems)
	proto.RegisterExtension(E_MaxItems)
	proto.RegisterExtension(E_DefaultAuthn)
	proto.RegisterExtension(E_DefaultAuthz)
	proto.RegisterExtension(E_Authn)
	proto.RegisterExtension(E_Authz)
	proto.RegisterExtension(E_Http)
}

var fileDescriptorAnnotations = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x85, 0xe8, 0xf4, 0xe2, 0x32, 0x02, 0x8d, 0xc4, 0x45, 0x95, 0x80, 0x2e, 0x91, 0x80,
	0x44, 0x82, 0x05, 0xc2, 0x88, 0x05, 0x14, 0x50, 0x5b, 0x51, 0x15, 0xa5, 0x1b, 0xc4, 0x06, 0xb9,
	0x93, 0x33, 0x89, 0x91, 0x63, 0xa7, 0x8e, 0x53, 0xaa, 0xac, 0xe0, 0x0d, 0xa0, 0xcf, 0x01, 0x3c,
	0x07, 0xbc, 0x07, 0xf7, 0xcb, 0x3b, 0x90, 0xd8, 0x27, 0x61, 0xa0, 0x33, 0xf2, 0xb0, 0xb5, 0xff,
	0xef, 0xf3, 0x39, 0x96, 0x8f, 0xc9, 0x75, 0xc1, 0xb3, 0x5d, 0x15, 0x14, 0xa0, 0xf7, 0xf9, 0x10,
	0x8a, 0x70, 0xa8, 0x34, 0x84, 0xba, 0x94, 0x86, 0x67, 0x10, 0xda, 0xbd, 0x90, 0x49, 0xa9, 0x0c,
	0x33, 0x5c, 0xc9, 0x22, 0xc8, 0xb5, 0x32, 0x6a, 0xd0, 0xb3, 0x1b, 0x2b, 0xab, 0x89, 0x52, 0x89,
	0x80, 0xd0, 0x2e, 0xee, 0x96, 0xa3, 0x30, 0x86, 0x62, 0xa8, 0x79, 0x6e, 0x94, 0x76, 0xc1, 0x95,
	0xab, 0x33, 0xd8, 0x4b, 0x93, 0xca, 0xff, 0x8b, 0x57, 0x18, 0xbf, 0xe2, 0x8f, 0xa7, 0xc6, 0xe4,
	0x98, 0xbe, 0xec, 0x4f, 0x17, 0x7b, 0xc2, 0x85, 0xe9, 0x36, 0xe9, 0x65, 0x2a, 0x06, 0x31, 0xb8,
	0x18, 0xb8, 0x26, 0x83, 0xb6, 0xc9, 0x60, 0x0b, 0x8a, 0x82, 0x25, 0xb0, 0x9d, 0xdb, 0x1b, 0x39,
	0xf7, 0xe2, 0xd5, 0xf1, 0xd5, 0x63, 0x97, 0x96, 0xaf, 0x9d, 0x09, 0x9c, 0x7f, 0xab, 0xc1, 0xee,
	0x75, 0x17, 0x11, 0x39, 0x0f, 0x5d, 0x23, 0xcb, 0x89, 0xaa, 0xfd, 0xfb, 0x4c, 0x94, 0xa0, 0xfd,
	0xda, 0x97, 0x56, 0xbb, 0x18, 0x8d, 0x53, 0xf4, 0x11, 0x99, 0x1f, 0x2a, 0x51, 0x66, 0x72, 0x70,
	0xfe, 0x08, 0xff, 0x80, 0x83, 0x88, 0x5b, 0xfa, 0x23, 0x16, 0x75, 0x16, 0x8b, 0x5a, 0xb3, 0xd4,
	0x58, 0x55, 0xe8, 0xa1, 0x9b, 0x64, 0xee, 0x99, 0xe2, 0x5e, 0xdf, 0x27, 0xf4, 0x9d, 0x46, 0xdf,
	0x66, 0xcd, 0x8c, 0xd9, 0xac, 0x83, 0xde, 0x27, 0xfd, 0x94, 0xc7, 0xb0, 0x21, 0x77, 0x9e, 0xb3,
	0x24, 0xa9, 0x9b, 0xf4, 0x48, 0x3f, 0x63, 0x8b, 0x7f, 0x53, 0xf4, 0x06, 0x99, 0x1f, 0x29, 0x9d,
	0x31, 0xe3, 0xe3, 0xbf, 0x58, 0x7e, 0x29, 0xc2, 0x38, 0xbd, 0x45, 0x16, 0x35, 0xec, 0x95, 0x5c,
	0x43, 0xec, 0x43, 0xbf, 0xe2, 0xd1, 0x1d, 0x40, 0x6f, 0x92, 0x85, 0x9c, 0x19, 0x03, 0xda, 0x7b,
	0x17, 0xdf, 0xf0, 0xd8, 0x36, 0x4f, 0x6f, 0x93, 0xa5, 0x8c, 0xcb, 0x87, 0x20, 0x13, 0x93, 0xfa,
	0xe0, 0xef, 0x16, 0xee, 0x47, 0x7f, 0x08, 0x8b, 0xb3, 0x83, 0xd9, 0xf0, 0x1f, 0x1d, 0xde, 0x12,
	0x4d, 0xd7, 0xb5, 0x6b, 0xc3, 0x40, 0x56, 0xf8, 0xe8, 0x9f, 0x48, 0x77, 0x80, 0x85, 0xd9, 0xc1,
	0x4c, 0xf0, 0xaf, 0x0e, 0x46, 0x80, 0x3e, 0x26, 0xfd, 0x18, 0x46, 0xac, 0x14, 0xe6, 0xa9, 0x1d,
	0xe2, 0x09, 0x8f, 0x7a, 0xc7, 0x0d, 0x5b, 0xeb, 0x78, 0xff, 0xda, 0x3d, 0xa3, 0x53, 0xf8, 0x8c,
	0xee, 0x34, 0x58, 0x54, 0x0a, 0x88, 0x4e, 0xa0, 0xc9, 0xae, 0xfc, 0x6b, 0xae, 0xfc, 0xe6, 0x0f,
	0x13, 0xcc, 0xd5, 0x11, 0x73, 0x45, 0xd7, 0x49, 0xcf, 0xd5, 0x7a, 0x61, 0xc2, 0x00, 0x9a, 0x54,
	0x75, 0xed, 0xbe, 0x39, 0x9c, 0x56, 0xaa, 0x13, 0xb4, 0xa6, 0xca, 0x6b, 0x7a, 0x7b, 0x38, 0xad,
	0x34, 0x27, 0xa8, 0xe7, 0x66, 0xae, 0xf9, 0xa6, 0xbc, 0xa2, 0x77, 0x28, 0x3a, 0x89, 0xa2, 0xf5,
	0x1a, 0xb2, 0x1e, 0x8b, 0xdf, 0x5d, 0x78, 0xe2, 0xbe, 0xe5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x65, 0xd8, 0xe9, 0x9e, 0xd3, 0x05, 0x00, 0x00,
}
