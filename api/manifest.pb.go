// Copyright 2019 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/manifest.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Compliance int32

const (
	Compliance_FIPS_1402 Compliance = 0
)

var Compliance_name = map[int32]string{
	0: "FIPS_1402",
}

var Compliance_value = map[string]int32{
	"FIPS_1402": 0,
}

func (x Compliance) String() string {
	return proto.EnumName(Compliance_name, int32(x))
}

func (Compliance) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{0}
}

type OperatingSystemName int32

const (
	OperatingSystemName_UBUNTU OperatingSystemName = 0
	OperatingSystemName_DEBIAN OperatingSystemName = 1
	OperatingSystemName_CENTOS OperatingSystemName = 2
	OperatingSystemName_RHEL   OperatingSystemName = 3
	OperatingSystemName_MACOS  OperatingSystemName = 4
)

var OperatingSystemName_name = map[int32]string{
	0: "UBUNTU",
	1: "DEBIAN",
	2: "CENTOS",
	3: "RHEL",
	4: "MACOS",
}

var OperatingSystemName_value = map[string]int32{
	"UBUNTU": 0,
	"DEBIAN": 1,
	"CENTOS": 2,
	"RHEL":   3,
	"MACOS":  4,
}

func (x OperatingSystemName) String() string {
	return proto.EnumName(OperatingSystemName_name, int32(x))
}

func (OperatingSystemName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{1}
}

// Builds must be uniquely addressable from the top level so that they can be used to look up the location of binaries.
// Format: filter_profile(-compliance_profile):envoy_version/operating_system-operating_system_version
// Examples:
//   - istio-fips1402:1.10.1/ubuntu-16.04
//   - standard:nightly/debian-8
type Manifest struct {
	ManifestVersion string `protobuf:"bytes,1,opt,name=manifest_version,json=manifestVersion,proto3" json:"manifest_version,omitempty"`
	// Key is composite key of the value's filter_profile and compliance_profile
	// Note: compliance_profile is optional
	// Format: filter_profile(-compliance_profile)
	Flavors              map[string]*Flavor `protobuf:"bytes,2,rep,name=flavors,proto3" json:"flavors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{0}
}

func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest.Unmarshal(m, b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
}
func (m *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(m, src)
}
func (m *Manifest) XXX_Size() int {
	return xxx_messageInfo_Manifest.Size(m)
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetManifestVersion() string {
	if m != nil {
		return m.ManifestVersion
	}
	return ""
}

func (m *Manifest) GetFlavors() map[string]*Flavor {
	if m != nil {
		return m.Flavors
	}
	return nil
}

type Flavor struct {
	// This is duplicated in order to make flavor easier to sort
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter profile is the name of the collection of filters
	// E.g. standard, minimal, istio
	FilterProfile string `protobuf:"bytes,2,opt,name=filter_profile,json=filterProfile,proto3" json:"filter_profile,omitempty"`
	// All filters available in this flavor
	Filters     []string     `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Compliances []Compliance `protobuf:"varint,4,rep,packed,name=compliances,proto3,enum=api.Compliance" json:"compliances,omitempty"`
	// Key is Envoy version
	Versions             map[string]*Version `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Flavor) Reset()         { *m = Flavor{} }
func (m *Flavor) String() string { return proto.CompactTextString(m) }
func (*Flavor) ProtoMessage()    {}
func (*Flavor) Descriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{1}
}

func (m *Flavor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flavor.Unmarshal(m, b)
}
func (m *Flavor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flavor.Marshal(b, m, deterministic)
}
func (m *Flavor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flavor.Merge(m, src)
}
func (m *Flavor) XXX_Size() int {
	return xxx_messageInfo_Flavor.Size(m)
}
func (m *Flavor) XXX_DiscardUnknown() {
	xxx_messageInfo_Flavor.DiscardUnknown(m)
}

var xxx_messageInfo_Flavor proto.InternalMessageInfo

func (m *Flavor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Flavor) GetFilterProfile() string {
	if m != nil {
		return m.FilterProfile
	}
	return ""
}

func (m *Flavor) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Flavor) GetCompliances() []Compliance {
	if m != nil {
		return m.Compliances
	}
	return nil
}

func (m *Flavor) GetVersions() map[string]*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type Version struct {
	// This duplicated in order to make version easier to sort
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Key is operating system
	// Note: Keys in proto maps cannot be enums, so this has to be a string.
	OperatingSystems     map[string]*OperatingSystem `protobuf:"bytes,2,rep,name=operating_systems,json=operatingSystems,proto3" json:"operating_systems,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{2}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetOperatingSystems() map[string]*OperatingSystem {
	if m != nil {
		return m.OperatingSystems
	}
	return nil
}

type OperatingSystem struct {
	// Matches the key, duplicated here to be explicit.
	Name                 OperatingSystemName `protobuf:"varint,1,opt,name=name,proto3,enum=api.OperatingSystemName" json:"name,omitempty"`
	Builds               []*Build            `protobuf:"bytes,2,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OperatingSystem) Reset()         { *m = OperatingSystem{} }
func (m *OperatingSystem) String() string { return proto.CompactTextString(m) }
func (*OperatingSystem) ProtoMessage()    {}
func (*OperatingSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{3}
}

func (m *OperatingSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatingSystem.Unmarshal(m, b)
}
func (m *OperatingSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatingSystem.Marshal(b, m, deterministic)
}
func (m *OperatingSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatingSystem.Merge(m, src)
}
func (m *OperatingSystem) XXX_Size() int {
	return xxx_messageInfo_OperatingSystem.Size(m)
}
func (m *OperatingSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatingSystem.DiscardUnknown(m)
}

var xxx_messageInfo_OperatingSystem proto.InternalMessageInfo

func (m *OperatingSystem) GetName() OperatingSystemName {
	if m != nil {
		return m.Name
	}
	return OperatingSystemName_UBUNTU
}

func (m *OperatingSystem) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	OperatingSystemVersion string   `protobuf:"bytes,1,opt,name=operating_system_version,json=operatingSystemVersion,proto3" json:"operating_system_version,omitempty"`
	DownloadLocationUrl    string   `protobuf:"bytes,2,opt,name=download_location_url,json=downloadLocationUrl,proto3" json:"download_location_url,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_346685e380de5b1f, []int{4}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetOperatingSystemVersion() string {
	if m != nil {
		return m.OperatingSystemVersion
	}
	return ""
}

func (m *Build) GetDownloadLocationUrl() string {
	if m != nil {
		return m.DownloadLocationUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Compliance", Compliance_name, Compliance_value)
	proto.RegisterEnum("api.OperatingSystemName", OperatingSystemName_name, OperatingSystemName_value)
	proto.RegisterType((*Manifest)(nil), "api.Manifest")
	proto.RegisterMapType((map[string]*Flavor)(nil), "api.Manifest.FlavorsEntry")
	proto.RegisterType((*Flavor)(nil), "api.Flavor")
	proto.RegisterMapType((map[string]*Version)(nil), "api.Flavor.VersionsEntry")
	proto.RegisterType((*Version)(nil), "api.Version")
	proto.RegisterMapType((map[string]*OperatingSystem)(nil), "api.Version.OperatingSystemsEntry")
	proto.RegisterType((*OperatingSystem)(nil), "api.OperatingSystem")
	proto.RegisterType((*Build)(nil), "api.Build")
}

func init() { proto.RegisterFile("api/manifest.proto", fileDescriptor_346685e380de5b1f) }

var fileDescriptor_346685e380de5b1f = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdf, 0x6e, 0xd3, 0x3c,
	0x1c, 0x5d, 0x9a, 0xfe, 0xfd, 0x75, 0x6d, 0xf3, 0x79, 0xdf, 0x50, 0x28, 0x37, 0x25, 0x12, 0x52,
	0xa9, 0x50, 0x60, 0x61, 0x48, 0x13, 0x77, 0x6b, 0xe9, 0xa0, 0xd2, 0x96, 0x4e, 0xe9, 0x8a, 0xc4,
	0x55, 0xe4, 0xb5, 0x2e, 0xb2, 0x70, 0xe2, 0x28, 0x49, 0x8b, 0xfa, 0x1a, 0x3c, 0x0b, 0x2f, 0xc1,
	0x5b, 0xa1, 0xc6, 0xf6, 0x48, 0xa3, 0xdc, 0xfd, 0x72, 0xce, 0xf1, 0x89, 0xcf, 0xb1, 0x0d, 0x08,
	0x47, 0xf4, 0x6d, 0x80, 0x43, 0xba, 0x21, 0x49, 0x6a, 0x47, 0x31, 0x4f, 0x39, 0xd2, 0x71, 0x44,
	0xad, 0xdf, 0x1a, 0x34, 0xef, 0x24, 0x8e, 0x5e, 0x83, 0xa1, 0x34, 0xfe, 0x8e, 0xc4, 0x09, 0xe5,
	0xa1, 0xa9, 0x0d, 0xb4, 0x61, 0xcb, 0xeb, 0x29, 0xfc, 0xab, 0x80, 0xd1, 0x25, 0x34, 0x36, 0x0c,
	0xef, 0x78, 0x9c, 0x98, 0x95, 0x81, 0x3e, 0x6c, 0x3b, 0x7d, 0x1b, 0x47, 0xd4, 0x56, 0x56, 0xf6,
	0x8d, 0x20, 0xa7, 0x61, 0x1a, 0xef, 0x3d, 0x25, 0xed, 0x7f, 0x86, 0xd3, 0x3c, 0x81, 0x0c, 0xd0,
	0x7f, 0x90, 0xbd, 0xfc, 0xc7, 0x61, 0x44, 0x2f, 0xa1, 0xb6, 0xc3, 0x6c, 0x4b, 0xcc, 0xca, 0x40,
	0x1b, 0xb6, 0x9d, 0x76, 0xe6, 0x2a, 0xd6, 0x78, 0x82, 0xf9, 0x58, 0xb9, 0xd2, 0xac, 0x5f, 0x15,
	0xa8, 0x0b, 0x14, 0x21, 0xa8, 0x86, 0x38, 0x20, 0xd2, 0x24, 0x9b, 0xd1, 0x2b, 0xe8, 0x6e, 0x28,
	0x4b, 0x49, 0xec, 0x47, 0x31, 0xdf, 0x50, 0x26, 0xec, 0x5a, 0x5e, 0x47, 0xa0, 0xf7, 0x02, 0x44,
	0x26, 0x34, 0x04, 0x90, 0x98, 0xfa, 0x40, 0x1f, 0xb6, 0x3c, 0xf5, 0x89, 0x2e, 0xa0, 0xbd, 0xe2,
	0x41, 0xc4, 0x28, 0x0e, 0x57, 0x24, 0x31, 0xab, 0x03, 0x7d, 0xd8, 0x75, 0x7a, 0xd9, 0x66, 0x26,
	0x4f, 0xb8, 0x97, 0xd7, 0xa0, 0x0f, 0xd0, 0x94, 0x9d, 0x25, 0x66, 0x2d, 0xab, 0xe4, 0x79, 0x6e,
	0xf3, 0xb6, 0x2c, 0x4e, 0x36, 0xf2, 0x24, 0xed, 0xcf, 0xa0, 0x73, 0x44, 0x95, 0x74, 0x62, 0x1d,
	0x77, 0x72, 0x9a, 0xd9, 0xca, 0x45, 0xf9, 0x52, 0xfe, 0x68, 0xd0, 0x50, 0xe7, 0x53, 0xd6, 0xca,
	0x1c, 0xfe, 0xe3, 0x11, 0x89, 0x71, 0x4a, 0xc3, 0xef, 0x7e, 0xb2, 0x4f, 0x52, 0x12, 0xa8, 0xd3,
	0xb3, 0xf2, 0x9e, 0xf6, 0x5c, 0xa9, 0x16, 0x42, 0x24, 0xf6, 0x6c, 0xf0, 0x02, 0xdc, 0xff, 0x06,
	0xe7, 0xa5, 0xd2, 0x92, 0x0c, 0xa3, 0xe3, 0x0c, 0xff, 0x67, 0xff, 0x2b, 0x2c, 0xce, 0x67, 0x59,
	0x41, 0xaf, 0xc0, 0xa2, 0x37, 0xb9, 0x48, 0x5d, 0xc7, 0x2c, 0x73, 0x70, 0x71, 0x40, 0x64, 0x58,
	0x0b, 0xea, 0x8f, 0x5b, 0xca, 0xd6, 0x2a, 0x21, 0x64, 0xfa, 0xf1, 0x01, 0xf2, 0x24, 0x63, 0x6d,
	0xa1, 0x96, 0x01, 0xe8, 0x0a, 0xcc, 0x62, 0x33, 0x85, 0x07, 0xf0, 0xac, 0x10, 0x5e, 0xf5, 0xec,
	0xc0, 0xf9, 0x9a, 0xff, 0x0c, 0x19, 0xc7, 0x6b, 0x9f, 0xf1, 0x15, 0x4e, 0x29, 0x0f, 0xfd, 0x6d,
	0xcc, 0xe4, 0x85, 0x3b, 0x53, 0xe4, 0xad, 0xe4, 0x96, 0x31, 0x1b, 0xbd, 0x00, 0xf8, 0x77, 0x89,
	0x50, 0x07, 0x5a, 0x37, 0xb3, 0xfb, 0x85, 0x7f, 0x71, 0xf9, 0xce, 0x31, 0x4e, 0x46, 0x2e, 0x9c,
	0x95, 0x84, 0x42, 0x00, 0xf5, 0xe5, 0x78, 0xe9, 0x3e, 0x2c, 0x8d, 0x93, 0xc3, 0xfc, 0x69, 0x3a,
	0x9e, 0x5d, 0xbb, 0x86, 0x76, 0x98, 0x27, 0x53, 0xf7, 0x61, 0xbe, 0x30, 0x2a, 0xa8, 0x09, 0x55,
	0xef, 0xcb, 0xf4, 0xd6, 0xd0, 0x51, 0x0b, 0x6a, 0x77, 0xd7, 0x93, 0xf9, 0xc2, 0xa8, 0x3e, 0xd6,
	0xb3, 0xc7, 0xfe, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x32, 0x39, 0xbf, 0x02, 0x04,
	0x00, 0x00,
}
