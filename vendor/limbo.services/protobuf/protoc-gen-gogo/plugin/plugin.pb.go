// Code generated by protoc-gen-gogo.
// source: plugin.proto
// DO NOT EDIT!

/*
Package plugin_go is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	CodeGeneratorRequest
	CodeGeneratorResponse
*/
package plugin_go

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

// An encoded CodeGeneratorRequest is written to the plugin's stdin.
type CodeGeneratorRequest struct {
	// The .proto files that were explicitly listed on the command-line.  The
	// code generator should generate code only for these files.  Each file's
	// descriptor will be included in proto_file, below.
	FileToGenerate []string `protobuf:"bytes,1,rep,name=file_to_generate,json=fileToGenerate" json:"file_to_generate,omitempty"`
	// The generator parameter passed on the command-line.
	Parameter *string `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	// FileDescriptorProtos for all files in files_to_generate and everything
	// they import.  The files will appear in topological order, so each file
	// appears before any file that imports it.
	//
	// protoc guarantees that all proto_files will be written after
	// the fields above, even though this is not technically guaranteed by the
	// protobuf wire format.  This theoretically could allow a plugin to stream
	// in the FileDescriptorProtos and handle them one by one rather than read
	// the entire set into memory at once.  However, as of this writing, this
	// is not similarly optimized on protoc's end -- it will store all fields in
	// memory at once before sending them to the plugin.
	ProtoFile        []*google_protobuf.FileDescriptorProto `protobuf:"bytes,15,rep,name=proto_file,json=protoFile" json:"proto_file,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *CodeGeneratorRequest) Reset()                    { *m = CodeGeneratorRequest{} }
func (m *CodeGeneratorRequest) String() string            { return proto.CompactTextString(m) }
func (*CodeGeneratorRequest) ProtoMessage()               {}
func (*CodeGeneratorRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{0} }

func (m *CodeGeneratorRequest) GetFileToGenerate() []string {
	if m != nil {
		return m.FileToGenerate
	}
	return nil
}

func (m *CodeGeneratorRequest) GetParameter() string {
	if m != nil && m.Parameter != nil {
		return *m.Parameter
	}
	return ""
}

func (m *CodeGeneratorRequest) GetProtoFile() []*google_protobuf.FileDescriptorProto {
	if m != nil {
		return m.ProtoFile
	}
	return nil
}

// The plugin writes an encoded CodeGeneratorResponse to stdout.
type CodeGeneratorResponse struct {
	// Error message.  If non-empty, code generation failed.  The plugin process
	// should exit with status code zero even if it reports an error in this way.
	//
	// This should be used to indicate errors in .proto files which prevent the
	// code generator from generating correct code.  Errors which indicate a
	// problem in protoc itself -- such as the input CodeGeneratorRequest being
	// unparseable -- should be reported by writing a message to stderr and
	// exiting with a non-zero status code.
	Error            *string                       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	File             []*CodeGeneratorResponse_File `protobuf:"bytes,15,rep,name=file" json:"file,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *CodeGeneratorResponse) Reset()                    { *m = CodeGeneratorResponse{} }
func (m *CodeGeneratorResponse) String() string            { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse) ProtoMessage()               {}
func (*CodeGeneratorResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{1} }

func (m *CodeGeneratorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *CodeGeneratorResponse) GetFile() []*CodeGeneratorResponse_File {
	if m != nil {
		return m.File
	}
	return nil
}

// Represents a single generated file.
type CodeGeneratorResponse_File struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If non-empty, indicates that the named file should already exist, and the
	// content here is to be inserted into that file at a defined insertion
	// point.  This feature allows a code generator to extend the output
	// produced by another code generator.  The original generator may provide
	// insertion points by placing special annotations in the file that look
	// like:
	//   @@protoc_insertion_point(NAME)
	// The annotation can have arbitrary text before and after it on the line,
	// which allows it to be placed in a comment.  NAME should be replaced with
	// an identifier naming the point -- this is what other generators will use
	// as the insertion_point.  Code inserted at this point will be placed
	// immediately above the line containing the insertion point (thus multiple
	// insertions to the same point will come out in the order they were added).
	// The double-@ is intended to make it unlikely that the generated code
	// could contain things that look like insertion points by accident.
	//
	// For example, the C++ code generator places the following line in the
	// .pb.h files that it generates:
	//   // @@protoc_insertion_point(namespace_scope)
	// This line appears within the scope of the file's package namespace, but
	// outside of any particular class.  Another plugin can then specify the
	// insertion_point "namespace_scope" to generate additional classes or
	// other declarations that should be placed in this scope.
	//
	// Note that if the line containing the insertion point begins with
	// whitespace, the same whitespace will be added to every line of the
	// inserted text.  This is useful for languages like Python, where
	// indentation matters.  In these languages, the insertion point comment
	// should be indented the same amount as any inserted code will need to be
	// in order to work correctly in that context.
	//
	// The code generator that generates the initial file and the one which
	// inserts into it must both run as part of a single invocation of protoc.
	// Code generators are executed in the order in which they appear on the
	// command line.
	//
	// If |insertion_point| is present, |name| must also be present.
	InsertionPoint *string `protobuf:"bytes,2,opt,name=insertion_point,json=insertionPoint" json:"insertion_point,omitempty"`
	// The file contents.
	Content          *string `protobuf:"bytes,15,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CodeGeneratorResponse_File) Reset()         { *m = CodeGeneratorResponse_File{} }
func (m *CodeGeneratorResponse_File) String() string { return proto.CompactTextString(m) }
func (*CodeGeneratorResponse_File) ProtoMessage()    {}
func (*CodeGeneratorResponse_File) Descriptor() ([]byte, []int) {
	return fileDescriptorPlugin, []int{1, 0}
}

func (m *CodeGeneratorResponse_File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetInsertionPoint() string {
	if m != nil && m.InsertionPoint != nil {
		return *m.InsertionPoint
	}
	return ""
}

func (m *CodeGeneratorResponse_File) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*CodeGeneratorRequest)(nil), "google.protobuf.compiler.CodeGeneratorRequest")
	proto.RegisterType((*CodeGeneratorResponse)(nil), "google.protobuf.compiler.CodeGeneratorResponse")
	proto.RegisterType((*CodeGeneratorResponse_File)(nil), "google.protobuf.compiler.CodeGeneratorResponse.File")
}

var fileDescriptorPlugin = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x24, 0xbf, 0x5f, 0x44, 0xf2, 0x2c, 0x8d, 0x2c, 0x15, 0x42, 0xe9, 0x21, 0x14, 0xc1, 0x9e,
	0x52, 0x10, 0xc1, 0x7b, 0x2b, 0xea, 0x31, 0x04, 0x4f, 0x82, 0x84, 0x98, 0xbe, 0x86, 0x85, 0x74,
	0xdf, 0xba, 0xd9, 0x7c, 0x22, 0xbf, 0x93, 0x9f, 0xc7, 0xfd, 0x93, 0x56, 0x29, 0xf6, 0x94, 0xbc,
	0x99, 0xd9, 0x99, 0xd9, 0x7d, 0x30, 0x92, 0x6d, 0xdf, 0x70, 0x91, 0x49, 0x45, 0x9a, 0x58, 0xd2,
	0x10, 0x35, 0x2d, 0xfa, 0xe9, 0xbd, 0xdf, 0x66, 0x35, 0xed, 0x24, 0x6f, 0x51, 0x4d, 0x53, 0xcf,
	0x2c, 0xf7, 0xcc, 0x72, 0x83, 0x5d, 0xad, 0xb8, 0xd4, 0xa4, 0xbc, 0x7a, 0xfe, 0x19, 0xc0, 0x64,
	0x4d, 0x1b, 0x7c, 0x42, 0x81, 0xaa, 0x32, 0x78, 0x81, 0x1f, 0x3d, 0x76, 0x9a, 0x2d, 0xe0, 0x72,
	0x6b, 0x3c, 0x4a, 0x4d, 0x65, 0xe3, 0x39, 0x4c, 0x82, 0xf4, 0xff, 0x22, 0x2a, 0xc6, 0x16, 0x7f,
	0xa1, 0xe1, 0x04, 0xb2, 0x19, 0x44, 0xb2, 0x52, 0xd5, 0x0e, 0x35, 0xaa, 0xe4, 0x5f, 0x1a, 0x18,
	0xc9, 0x0f, 0xc0, 0xd6, 0x00, 0x2e, 0xa9, 0xb4, 0xa7, 0x92, 0xd8, 0x38, 0x5c, 0xdc, 0x5e, 0x67,
	0xc7, 0x8d, 0x1f, 0x0d, 0xf9, 0x70, 0xe8, 0x96, 0x5b, 0xd8, 0x98, 0xd8, 0x8f, 0x65, 0xe6, 0x5f,
	0x01, 0x5c, 0x1d, 0xb5, 0xec, 0x24, 0x89, 0x0e, 0xd9, 0x04, 0xce, 0x50, 0x29, 0x52, 0xa6, 0x9b,
	0x0d, 0xf6, 0x03, 0x7b, 0x86, 0xf0, 0x57, 0xdc, 0x5d, 0x76, 0xea, 0x81, 0xb2, 0x3f, 0x4d, 0x5d,
	0x9b, 0xc2, 0x39, 0x4c, 0xdf, 0x20, 0xb4, 0x13, 0x63, 0x10, 0x0a, 0x73, 0xa3, 0x21, 0xc6, 0xfd,
	0xb3, 0x1b, 0x88, 0xb9, 0x91, 0x2b, 0xcd, 0x49, 0x94, 0x92, 0xb8, 0xd0, 0xc3, 0xf5, 0xc7, 0x07,
	0x38, 0xb7, 0x28, 0x4b, 0xe0, 0xbc, 0x26, 0xa1, 0xd1, 0x08, 0x62, 0x27, 0xd8, 0x8f, 0xab, 0x7b,
	0x98, 0x99, 0x2e, 0x27, 0xfb, 0xad, 0x46, 0xb9, 0x5b, 0xb4, 0x7b, 0x90, 0xee, 0x35, 0xf2, 0x6b,
	0x2f, 0x1b, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa2, 0xbd, 0x30, 0x02, 0x02, 0x00, 0x00,
}