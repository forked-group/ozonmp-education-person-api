// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: ozonmp/education_person_api/v1/education_person_api.proto

package education_person_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Sex int32

const (
	Sex_SEX_NONE   Sex = 0
	Sex_SEX_FEMALE Sex = 1
	Sex_SEX_MALE   Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "SEX_NONE",
		1: "SEX_FEMALE",
		2: "SEX_MALE",
	}
	Sex_value = map[string]int32{
		"SEX_NONE":   0,
		"SEX_FEMALE": 1,
		"SEX_MALE":   2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{0}
}

type Education int32

const (
	Education_EDUCATION_NONE Education = 0
	// Дошкольное
	Education_EDUCATION_PRESCHOOL Education = 1
	// Начальное общее — 1—4 классы
	Education_EDUCATION_PRIMARY_GENERAL Education = 2
	// Основное общее — 5—9 классы
	Education_EDUCATION_BASIC_GENERAL Education = 3
	// Среднее общее — 10—11 классы
	Education_EDUCATION_SECONDARY_GENERAL Education = 4
	// Среднее профессиональное
	Education_EDUCATION_SECONDARY_VOCATIONAL Education = 5
	// Высшее I степени — бакалавриат
	Education_EDUCATION_HIGHER_1 Education = 6
	// Высшее II степени — специалитет, магистратура
	Education_EDUCATION_HIGHER_2 Education = 7
	// Высшее III степени — подготовка кадров высшей квалификации
	Education_EDUCATION_HIGHER_3 Education = 8
)

// Enum value maps for Education.
var (
	Education_name = map[int32]string{
		0: "EDUCATION_NONE",
		1: "EDUCATION_PRESCHOOL",
		2: "EDUCATION_PRIMARY_GENERAL",
		3: "EDUCATION_BASIC_GENERAL",
		4: "EDUCATION_SECONDARY_GENERAL",
		5: "EDUCATION_SECONDARY_VOCATIONAL",
		6: "EDUCATION_HIGHER_1",
		7: "EDUCATION_HIGHER_2",
		8: "EDUCATION_HIGHER_3",
	}
	Education_value = map[string]int32{
		"EDUCATION_NONE":                 0,
		"EDUCATION_PRESCHOOL":            1,
		"EDUCATION_PRIMARY_GENERAL":      2,
		"EDUCATION_BASIC_GENERAL":        3,
		"EDUCATION_SECONDARY_GENERAL":    4,
		"EDUCATION_SECONDARY_VOCATIONAL": 5,
		"EDUCATION_HIGHER_1":             6,
		"EDUCATION_HIGHER_2":             7,
		"EDUCATION_HIGHER_3":             8,
	}
)

func (x Education) Enum() *Education {
	p := new(Education)
	*p = x
	return p
}

func (x Education) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Education) Descriptor() protoreflect.EnumDescriptor {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes[1].Descriptor()
}

func (Education) Type() protoreflect.EnumType {
	return &file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes[1]
}

func (x Education) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Education.Descriptor instead.
func (Education) EnumDescriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{1}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FistName   string                 `protobuf:"bytes,2,opt,name=fist_name,json=fistName,proto3" json:"fist_name,omitempty"`
	MiddleName string                 `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName   string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Sex        Sex                    `protobuf:"varint,6,opt,name=sex,proto3,enum=ozonmp.education_person_api.v1.Sex" json:"sex,omitempty"`
	Education  Education              `protobuf:"varint,7,opt,name=education,proto3,enum=ozonmp.education_person_api.v1.Education" json:"education,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetFistName() string {
	if x != nil {
		return x.FistName
	}
	return ""
}

func (x *Person) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *Person) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_SEX_NONE
}

func (x *Person) GetEducation() Education {
	if x != nil {
		return x.Education
	}
	return Education_EDUCATION_NONE
}

func (x *Person) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type CreatePersonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *CreatePersonV1Request) Reset() {
	*x = CreatePersonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonV1Request) ProtoMessage() {}

func (x *CreatePersonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonV1Request.ProtoReflect.Descriptor instead.
func (*CreatePersonV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePersonV1Request) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type CreatePersonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId uint64 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *CreatePersonV1Response) Reset() {
	*x = CreatePersonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonV1Response) ProtoMessage() {}

func (x *CreatePersonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonV1Response.ProtoReflect.Descriptor instead.
func (*CreatePersonV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePersonV1Response) GetPersonId() uint64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

type DescribePersonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId uint64 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *DescribePersonV1Request) Reset() {
	*x = DescribePersonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePersonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePersonV1Request) ProtoMessage() {}

func (x *DescribePersonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePersonV1Request.ProtoReflect.Descriptor instead.
func (*DescribePersonV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePersonV1Request) GetPersonId() uint64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

type DescribePersonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *DescribePersonV1Response) Reset() {
	*x = DescribePersonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePersonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePersonV1Response) ProtoMessage() {}

func (x *DescribePersonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePersonV1Response.ProtoReflect.Descriptor instead.
func (*DescribePersonV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribePersonV1Response) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type ListPersonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor uint64 `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListPersonV1Request) Reset() {
	*x = ListPersonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonV1Request) ProtoMessage() {}

func (x *ListPersonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonV1Request.ProtoReflect.Descriptor instead.
func (*ListPersonV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListPersonV1Request) GetCursor() uint64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

type ListPersonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person []*Person `protobuf:"bytes,1,rep,name=person,proto3" json:"person,omitempty"`
}

func (x *ListPersonV1Response) Reset() {
	*x = ListPersonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonV1Response) ProtoMessage() {}

func (x *ListPersonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonV1Response.ProtoReflect.Descriptor instead.
func (*ListPersonV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListPersonV1Response) GetPerson() []*Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type RemovePersonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId uint64 `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *RemovePersonV1Request) Reset() {
	*x = RemovePersonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePersonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePersonV1Request) ProtoMessage() {}

func (x *RemovePersonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePersonV1Request.ProtoReflect.Descriptor instead.
func (*RemovePersonV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemovePersonV1Request) GetPersonId() uint64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

type RemovePersonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *RemovePersonV1Response) Reset() {
	*x = RemovePersonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePersonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePersonV1Response) ProtoMessage() {}

func (x *RemovePersonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePersonV1Response.ProtoReflect.Descriptor instead.
func (*RemovePersonV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemovePersonV1Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_ozonmp_education_person_api_v1_education_person_api_proto protoreflect.FileDescriptor

var file_ozonmp_education_person_api_v1_education_person_api_proto_rawDesc = []byte{
	0x0a, 0x39, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x12, 0x35, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x47, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x22, 0x35, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x56, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x2a, 0x31, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x45, 0x58, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x58,
	0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x58,
	0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x81, 0x02, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x44, 0x55,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x41, 0x52, 0x59, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x04, 0x12,
	0x22, 0x0a, 0x1e, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x43,
	0x4f, 0x4e, 0x44, 0x41, 0x52, 0x59, 0x5f, 0x56, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41,
	0x4c, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x48, 0x49, 0x47, 0x48, 0x45, 0x52, 0x5f, 0x31, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x45, 0x52, 0x5f,
	0x32, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x48, 0x49, 0x47, 0x48, 0x45, 0x52, 0x5f, 0x33, 0x10, 0x08, 0x32, 0xc3, 0x04, 0x0a, 0x19,
	0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa8, 0x01, 0x0a, 0x10, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x37,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x35, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e,
	0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x33, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7f, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x56, 0x31, 0x12, 0x35, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2e, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x61, 0x61, 0x32, 0x70, 0x70, 0x70, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2d, 0x65,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescOnce sync.Once
	file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescData = file_ozonmp_education_person_api_v1_education_person_api_proto_rawDesc
)

func file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescGZIP() []byte {
	file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescOnce.Do(func() {
		file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescData)
	})
	return file_ozonmp_education_person_api_v1_education_person_api_proto_rawDescData
}

var file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ozonmp_education_person_api_v1_education_person_api_proto_goTypes = []interface{}{
	(Sex)(0),                         // 0: ozonmp.education_person_api.v1.Sex
	(Education)(0),                   // 1: ozonmp.education_person_api.v1.Education
	(*Person)(nil),                   // 2: ozonmp.education_person_api.v1.Person
	(*CreatePersonV1Request)(nil),    // 3: ozonmp.education_person_api.v1.CreatePersonV1Request
	(*CreatePersonV1Response)(nil),   // 4: ozonmp.education_person_api.v1.CreatePersonV1Response
	(*DescribePersonV1Request)(nil),  // 5: ozonmp.education_person_api.v1.DescribePersonV1Request
	(*DescribePersonV1Response)(nil), // 6: ozonmp.education_person_api.v1.DescribePersonV1Response
	(*ListPersonV1Request)(nil),      // 7: ozonmp.education_person_api.v1.ListPersonV1Request
	(*ListPersonV1Response)(nil),     // 8: ozonmp.education_person_api.v1.ListPersonV1Response
	(*RemovePersonV1Request)(nil),    // 9: ozonmp.education_person_api.v1.RemovePersonV1Request
	(*RemovePersonV1Response)(nil),   // 10: ozonmp.education_person_api.v1.RemovePersonV1Response
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_ozonmp_education_person_api_v1_education_person_api_proto_depIdxs = []int32{
	11, // 0: ozonmp.education_person_api.v1.Person.birthday:type_name -> google.protobuf.Timestamp
	0,  // 1: ozonmp.education_person_api.v1.Person.sex:type_name -> ozonmp.education_person_api.v1.Sex
	1,  // 2: ozonmp.education_person_api.v1.Person.education:type_name -> ozonmp.education_person_api.v1.Education
	11, // 3: ozonmp.education_person_api.v1.Person.created:type_name -> google.protobuf.Timestamp
	2,  // 4: ozonmp.education_person_api.v1.CreatePersonV1Request.person:type_name -> ozonmp.education_person_api.v1.Person
	2,  // 5: ozonmp.education_person_api.v1.DescribePersonV1Response.person:type_name -> ozonmp.education_person_api.v1.Person
	2,  // 6: ozonmp.education_person_api.v1.ListPersonV1Response.person:type_name -> ozonmp.education_person_api.v1.Person
	5,  // 7: ozonmp.education_person_api.v1.EducationPersonApiService.DescribePersonV1:input_type -> ozonmp.education_person_api.v1.DescribePersonV1Request
	3,  // 8: ozonmp.education_person_api.v1.EducationPersonApiService.CreatePersonV1:input_type -> ozonmp.education_person_api.v1.CreatePersonV1Request
	7,  // 9: ozonmp.education_person_api.v1.EducationPersonApiService.ListPersonV1:input_type -> ozonmp.education_person_api.v1.ListPersonV1Request
	9,  // 10: ozonmp.education_person_api.v1.EducationPersonApiService.RemovePersonV1:input_type -> ozonmp.education_person_api.v1.RemovePersonV1Request
	6,  // 11: ozonmp.education_person_api.v1.EducationPersonApiService.DescribePersonV1:output_type -> ozonmp.education_person_api.v1.DescribePersonV1Response
	4,  // 12: ozonmp.education_person_api.v1.EducationPersonApiService.CreatePersonV1:output_type -> ozonmp.education_person_api.v1.CreatePersonV1Response
	8,  // 13: ozonmp.education_person_api.v1.EducationPersonApiService.ListPersonV1:output_type -> ozonmp.education_person_api.v1.ListPersonV1Response
	10, // 14: ozonmp.education_person_api.v1.EducationPersonApiService.RemovePersonV1:output_type -> ozonmp.education_person_api.v1.RemovePersonV1Response
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ozonmp_education_person_api_v1_education_person_api_proto_init() }
func file_ozonmp_education_person_api_v1_education_person_api_proto_init() {
	if File_ozonmp_education_person_api_v1_education_person_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonV1Request); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonV1Response); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePersonV1Request); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePersonV1Response); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonV1Request); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonV1Response); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePersonV1Request); i {
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
		file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePersonV1Response); i {
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
			RawDescriptor: file_ozonmp_education_person_api_v1_education_person_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ozonmp_education_person_api_v1_education_person_api_proto_goTypes,
		DependencyIndexes: file_ozonmp_education_person_api_v1_education_person_api_proto_depIdxs,
		EnumInfos:         file_ozonmp_education_person_api_v1_education_person_api_proto_enumTypes,
		MessageInfos:      file_ozonmp_education_person_api_v1_education_person_api_proto_msgTypes,
	}.Build()
	File_ozonmp_education_person_api_v1_education_person_api_proto = out.File
	file_ozonmp_education_person_api_v1_education_person_api_proto_rawDesc = nil
	file_ozonmp_education_person_api_v1_education_person_api_proto_goTypes = nil
	file_ozonmp_education_person_api_v1_education_person_api_proto_depIdxs = nil
}
