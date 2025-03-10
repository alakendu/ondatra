// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: testbed.proto

package proto

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

type Device_Vendor int32

const (
	Device_VENDOR_UNSPECIFIED Device_Vendor = 0
	Device_ARISTA             Device_Vendor = 1
	Device_CISCO              Device_Vendor = 2
	Device_IXIA               Device_Vendor = 3
	Device_JUNIPER            Device_Vendor = 4
	Device_CIENA              Device_Vendor = 5
	Device_PALOALTO           Device_Vendor = 6
	Device_NOKIA              Device_Vendor = 7
	Device_ZPE                Device_Vendor = 8
	Device_DELL               Device_Vendor = 9
)

// Enum value maps for Device_Vendor.
var (
	Device_Vendor_name = map[int32]string{
		0: "VENDOR_UNSPECIFIED",
		1: "ARISTA",
		2: "CISCO",
		3: "IXIA",
		4: "JUNIPER",
		5: "CIENA",
		6: "PALOALTO",
		7: "NOKIA",
		8: "ZPE",
		9: "DELL",
	}
	Device_Vendor_value = map[string]int32{
		"VENDOR_UNSPECIFIED": 0,
		"ARISTA":             1,
		"CISCO":              2,
		"IXIA":               3,
		"JUNIPER":            4,
		"CIENA":              5,
		"PALOALTO":           6,
		"NOKIA":              7,
		"ZPE":                8,
		"DELL":               9,
	}
)

func (x Device_Vendor) Enum() *Device_Vendor {
	p := new(Device_Vendor)
	*p = x
	return p
}

func (x Device_Vendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_Vendor) Descriptor() protoreflect.EnumDescriptor {
	return file_testbed_proto_enumTypes[0].Descriptor()
}

func (Device_Vendor) Type() protoreflect.EnumType {
	return &file_testbed_proto_enumTypes[0]
}

func (x Device_Vendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_Vendor.Descriptor instead.
func (Device_Vendor) EnumDescriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{1, 0}
}

// Speed enum values are the port speed in Gbps.
type Port_Speed int32

const (
	Port_SPEED_UNSPECIFIED Port_Speed = 0
	Port_S_1GB             Port_Speed = 1
	Port_S_5GB             Port_Speed = 5
	Port_S_10GB            Port_Speed = 10
	Port_S_100GB           Port_Speed = 100
	Port_S_400GB           Port_Speed = 400
)

// Enum value maps for Port_Speed.
var (
	Port_Speed_name = map[int32]string{
		0:   "SPEED_UNSPECIFIED",
		1:   "S_1GB",
		5:   "S_5GB",
		10:  "S_10GB",
		100: "S_100GB",
		400: "S_400GB",
	}
	Port_Speed_value = map[string]int32{
		"SPEED_UNSPECIFIED": 0,
		"S_1GB":             1,
		"S_5GB":             5,
		"S_10GB":            10,
		"S_100GB":           100,
		"S_400GB":           400,
	}
)

func (x Port_Speed) Enum() *Port_Speed {
	p := new(Port_Speed)
	*p = x
	return p
}

func (x Port_Speed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Port_Speed) Descriptor() protoreflect.EnumDescriptor {
	return file_testbed_proto_enumTypes[1].Descriptor()
}

func (Port_Speed) Type() protoreflect.EnumType {
	return &file_testbed_proto_enumTypes[1]
}

func (x Port_Speed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Port_Speed.Descriptor instead.
func (Port_Speed) EnumDescriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{2, 0}
}

// Pmd enum values are PMD of the port.
type Port_Pmd int32

const (
	Port_PMD_UNSPECIFIED Port_Pmd = 0
	Port_PMD_100G_FR     Port_Pmd = 1
	Port_PMD_100G_LR4    Port_Pmd = 2
	Port_PMD_400G_FR4    Port_Pmd = 3
)

// Enum value maps for Port_Pmd.
var (
	Port_Pmd_name = map[int32]string{
		0: "PMD_UNSPECIFIED",
		1: "PMD_100G_FR",
		2: "PMD_100G_LR4",
		3: "PMD_400G_FR4",
	}
	Port_Pmd_value = map[string]int32{
		"PMD_UNSPECIFIED": 0,
		"PMD_100G_FR":     1,
		"PMD_100G_LR4":    2,
		"PMD_400G_FR4":    3,
	}
)

func (x Port_Pmd) Enum() *Port_Pmd {
	p := new(Port_Pmd)
	*p = x
	return p
}

func (x Port_Pmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Port_Pmd) Descriptor() protoreflect.EnumDescriptor {
	return file_testbed_proto_enumTypes[2].Descriptor()
}

func (Port_Pmd) Type() protoreflect.EnumType {
	return &file_testbed_proto_enumTypes[2]
}

func (x Port_Pmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Port_Pmd.Descriptor instead.
func (Port_Pmd) EnumDescriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{2, 1}
}

// A testbed.
type Testbed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duts  []*Device `protobuf:"bytes,1,rep,name=duts,proto3" json:"duts,omitempty"`
	Ates  []*Device `protobuf:"bytes,2,rep,name=ates,proto3" json:"ates,omitempty"`
	Links []*Link   `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Testbed) Reset() {
	*x = Testbed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testbed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testbed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testbed) ProtoMessage() {}

func (x *Testbed) ProtoReflect() protoreflect.Message {
	mi := &file_testbed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testbed.ProtoReflect.Descriptor instead.
func (*Testbed) Descriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{0}
}

func (x *Testbed) GetDuts() []*Device {
	if x != nil {
		return x.Duts
	}
	return nil
}

func (x *Testbed) GetAtes() []*Device {
	if x != nil {
		return x.Ates
	}
	return nil
}

func (x *Testbed) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

// A device.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vendor Device_Vendor `protobuf:"varint,2,opt,name=vendor,proto3,enum=ondatra.Device_Vendor" json:"vendor,omitempty"`
	// Hardware model of the device. Optional.
	//
	// If the value starts with "regex:" then the suffix is interpreted as a RE2
	// regular expression. The model will be restricted to models matching the
	// regular expression.
	HardwareModel string `protobuf:"bytes,4,opt,name=hardware_model,json=hardwareModel,proto3" json:"hardware_model,omitempty"`
	// Software version of the device. Optional.
	//
	// If the value starts with "regex:" then the suffix is interpreted as a RE2
	// regular expression. The version will be restricted to versions matching the
	// regular expression.
	SoftwareVersion string  `protobuf:"bytes,5,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty"`
	Ports           []*Port `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	// A key-value map of additional device dimensions. Optional.
	//
	// In addition to the above fields, the extra dimensions field can be used to
	// further restrict matching devices. The set of dimension keys that are
	// supported is specific to the binding implementation. For example, if the
	// binding supports filtering devices by a dimension named "label," the
	// testbed could specify an extra dimensions map of
	// extra_dimensions {
	//   key: "label",
	//   value: "foo",
	// }
	ExtraDimensions map[string]string `protobuf:"bytes,6,rep,name=extra_dimensions,json=extraDimensions,proto3" json:"extra_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testbed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_testbed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetVendor() Device_Vendor {
	if x != nil {
		return x.Vendor
	}
	return Device_VENDOR_UNSPECIFIED
}

func (x *Device) GetHardwareModel() string {
	if x != nil {
		return x.HardwareModel
	}
	return ""
}

func (x *Device) GetSoftwareVersion() string {
	if x != nil {
		return x.SoftwareVersion
	}
	return ""
}

func (x *Device) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Device) GetExtraDimensions() map[string]string {
	if x != nil {
		return x.ExtraDimensions
	}
	return nil
}

// A port.
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Speed Port_Speed `protobuf:"varint,2,opt,name=speed,proto3,enum=ondatra.Port_Speed" json:"speed,omitempty"`
	// The model of the card which contains the physical port.
	CardModel string `protobuf:"bytes,3,opt,name=card_model,json=cardModel,proto3" json:"card_model,omitempty"`
	// Types that are assignable to PmdValue:
	//	*Port_Pmd_
	//	*Port_PmdRegex
	PmdValue isPort_PmdValue `protobuf_oneof:"pmd_value"`
	Group    string          `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testbed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_testbed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{2}
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Port) GetSpeed() Port_Speed {
	if x != nil {
		return x.Speed
	}
	return Port_SPEED_UNSPECIFIED
}

func (x *Port) GetCardModel() string {
	if x != nil {
		return x.CardModel
	}
	return ""
}

func (m *Port) GetPmdValue() isPort_PmdValue {
	if m != nil {
		return m.PmdValue
	}
	return nil
}

func (x *Port) GetPmd() Port_Pmd {
	if x, ok := x.GetPmdValue().(*Port_Pmd_); ok {
		return x.Pmd
	}
	return Port_PMD_UNSPECIFIED
}

func (x *Port) GetPmdRegex() string {
	if x, ok := x.GetPmdValue().(*Port_PmdRegex); ok {
		return x.PmdRegex
	}
	return ""
}

func (x *Port) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type isPort_PmdValue interface {
	isPort_PmdValue()
}

type Port_Pmd_ struct {
	Pmd Port_Pmd `protobuf:"varint,4,opt,name=pmd,proto3,enum=ondatra.Port_Pmd,oneof"`
}

type Port_PmdRegex struct {
	PmdRegex string `protobuf:"bytes,5,opt,name=pmd_regex,json=pmdRegex,proto3,oneof"`
}

func (*Port_Pmd_) isPort_PmdValue() {}

func (*Port_PmdRegex) isPort_PmdValue() {}

// A physical link between ports on DUTs or ATEs.
// The order does not matter: links are symmetrical.
// A given port may be specified in at most one link (typically in exactly one
// link, because un-connected ports are not very interesting for testing).
type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"` // First port in the format "<device-id>:<port-id>".
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"` // Second port in the format "<device-id>:<port-id>".
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testbed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_testbed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_testbed_proto_rawDescGZIP(), []int{3}
}

func (x *Link) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Link) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_testbed_proto protoreflect.FileDescriptor

var file_testbed_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x62, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x22, 0x78, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74,
	0x62, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x04, 0x64, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x61, 0x74, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f,
	0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0xdc, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x52, 0x49, 0x53, 0x54, 0x41, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x49, 0x53, 0x43,
	0x4f, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x58, 0x49, 0x41, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x4a, 0x55, 0x4e, 0x49, 0x50, 0x45, 0x52, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x49,
	0x45, 0x4e, 0x41, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x4c, 0x4f, 0x41, 0x4c, 0x54,
	0x4f, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x4b, 0x49, 0x41, 0x10, 0x07, 0x12, 0x07,
	0x0a, 0x03, 0x5a, 0x50, 0x45, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4c, 0x4c, 0x10,
	0x09, 0x22, 0xf7, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6f, 0x6e, 0x64, 0x61,
	0x74, 0x72, 0x61, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x03, 0x70, 0x6d, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x2e, 0x50, 0x6d, 0x64, 0x48, 0x00, 0x52, 0x03, 0x70, 0x6d, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x70,
	0x6d, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x6d, 0x64, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x5b, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x50, 0x45,
	0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x5f, 0x31, 0x47, 0x42, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x5f, 0x35, 0x47, 0x42, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x5f, 0x31, 0x30, 0x47, 0x42,
	0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x5f, 0x31, 0x30, 0x30, 0x47, 0x42, 0x10, 0x64, 0x12,
	0x0c, 0x0a, 0x07, 0x53, 0x5f, 0x34, 0x30, 0x30, 0x47, 0x42, 0x10, 0x90, 0x03, 0x22, 0x4f, 0x0a,
	0x03, 0x50, 0x6d, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4d, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4d, 0x44,
	0x5f, 0x31, 0x30, 0x30, 0x47, 0x5f, 0x46, 0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4d,
	0x44, 0x5f, 0x31, 0x30, 0x30, 0x47, 0x5f, 0x4c, 0x52, 0x34, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x4d, 0x44, 0x5f, 0x34, 0x30, 0x30, 0x47, 0x5f, 0x46, 0x52, 0x34, 0x10, 0x03, 0x42, 0x0b,
	0x0a, 0x09, 0x70, 0x6d, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x62, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x72, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testbed_proto_rawDescOnce sync.Once
	file_testbed_proto_rawDescData = file_testbed_proto_rawDesc
)

func file_testbed_proto_rawDescGZIP() []byte {
	file_testbed_proto_rawDescOnce.Do(func() {
		file_testbed_proto_rawDescData = protoimpl.X.CompressGZIP(file_testbed_proto_rawDescData)
	})
	return file_testbed_proto_rawDescData
}

var file_testbed_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_testbed_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_testbed_proto_goTypes = []interface{}{
	(Device_Vendor)(0), // 0: ondatra.Device.Vendor
	(Port_Speed)(0),    // 1: ondatra.Port.Speed
	(Port_Pmd)(0),      // 2: ondatra.Port.Pmd
	(*Testbed)(nil),    // 3: ondatra.Testbed
	(*Device)(nil),     // 4: ondatra.Device
	(*Port)(nil),       // 5: ondatra.Port
	(*Link)(nil),       // 6: ondatra.Link
	nil,                // 7: ondatra.Device.ExtraDimensionsEntry
}
var file_testbed_proto_depIdxs = []int32{
	4, // 0: ondatra.Testbed.duts:type_name -> ondatra.Device
	4, // 1: ondatra.Testbed.ates:type_name -> ondatra.Device
	6, // 2: ondatra.Testbed.links:type_name -> ondatra.Link
	0, // 3: ondatra.Device.vendor:type_name -> ondatra.Device.Vendor
	5, // 4: ondatra.Device.ports:type_name -> ondatra.Port
	7, // 5: ondatra.Device.extra_dimensions:type_name -> ondatra.Device.ExtraDimensionsEntry
	1, // 6: ondatra.Port.speed:type_name -> ondatra.Port.Speed
	2, // 7: ondatra.Port.pmd:type_name -> ondatra.Port.Pmd
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_testbed_proto_init() }
func file_testbed_proto_init() {
	if File_testbed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testbed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testbed); i {
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
		file_testbed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_testbed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_testbed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
	file_testbed_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Port_Pmd_)(nil),
		(*Port_PmdRegex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testbed_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testbed_proto_goTypes,
		DependencyIndexes: file_testbed_proto_depIdxs,
		EnumInfos:         file_testbed_proto_enumTypes,
		MessageInfos:      file_testbed_proto_msgTypes,
	}.Build()
	File_testbed_proto = out.File
	file_testbed_proto_rawDesc = nil
	file_testbed_proto_goTypes = nil
	file_testbed_proto_depIdxs = nil
}
