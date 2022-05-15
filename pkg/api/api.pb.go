// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api/api.proto

package homework_2

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

type SubmitResult int32

const (
	SubmitResult_CORRECT SubmitResult = 0
	SubmitResult_WRONG   SubmitResult = 1
)

// Enum value maps for SubmitResult.
var (
	SubmitResult_name = map[int32]string{
		0: "CORRECT",
		1: "WRONG",
	}
	SubmitResult_value = map[string]int32{
		"CORRECT": 0,
		"WRONG":   1,
	}
)

func (x SubmitResult) Enum() *SubmitResult {
	p := new(SubmitResult)
	*p = x
	return p
}

func (x SubmitResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubmitResult) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_proto_enumTypes[0].Descriptor()
}

func (SubmitResult) Type() protoreflect.EnumType {
	return &file_api_api_proto_enumTypes[0]
}

func (x SubmitResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubmitResult.Descriptor instead.
func (SubmitResult) EnumDescriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId   int64    `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Image       string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Description []string `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Problem) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *Problem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Problem) GetDescription() []string {
	if x != nil {
		return x.Description
	}
	return nil
}

type CheckAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Answer string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *CheckAnswerRequest) Reset() {
	*x = CheckAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAnswerRequest) ProtoMessage() {}

func (x *CheckAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAnswerRequest.ProtoReflect.Descriptor instead.
func (*CheckAnswerRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAnswerRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *CheckAnswerRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type CheckAnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64        `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Answer    string       `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Result    SubmitResult `protobuf:"varint,3,opt,name=result,proto3,enum=api.SubmitResult" json:"result,omitempty"`
}

func (x *CheckAnswerResponse) Reset() {
	*x = CheckAnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAnswerResponse) ProtoMessage() {}

func (x *CheckAnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAnswerResponse.ProtoReflect.Descriptor instead.
func (*CheckAnswerResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *CheckAnswerResponse) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *CheckAnswerResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CheckAnswerResponse) GetResult() SubmitResult {
	if x != nil {
		return x.Result
	}
	return SubmitResult_CORRECT
}

type GetRandomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *GetRandomRequest) Reset() {
	*x = GetRandomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomRequest) ProtoMessage() {}

func (x *GetRandomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomRequest.ProtoReflect.Descriptor instead.
func (*GetRandomRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetRandomRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type GetRandomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskNumber int64    `protobuf:"varint,1,opt,name=task_number,json=taskNumber,proto3" json:"task_number,omitempty"`
	Problem    *Problem `protobuf:"bytes,2,opt,name=problem,proto3" json:"problem,omitempty"`
}

func (x *GetRandomResponse) Reset() {
	*x = GetRandomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomResponse) ProtoMessage() {}

func (x *GetRandomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomResponse.ProtoReflect.Descriptor instead.
func (*GetRandomResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetRandomResponse) GetTaskNumber() int64 {
	if x != nil {
		return x.TaskNumber
	}
	return 0
}

func (x *GetRandomResponse) GetProblem() *Problem {
	if x != nil {
		return x.Problem
	}
	return nil
}

type GetProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId     int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	TaskNumber int64 `protobuf:"varint,2,opt,name=task_number,json=taskNumber,proto3" json:"task_number,omitempty"`
}

func (x *GetProblemRequest) Reset() {
	*x = GetProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemRequest) ProtoMessage() {}

func (x *GetProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemRequest.ProtoReflect.Descriptor instead.
func (*GetProblemRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetProblemRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetProblemRequest) GetTaskNumber() int64 {
	if x != nil {
		return x.TaskNumber
	}
	return 0
}

type GetProblemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problem *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
}

func (x *GetProblemResponse) Reset() {
	*x = GetProblemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProblemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProblemResponse) ProtoMessage() {}

func (x *GetProblemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProblemResponse.ProtoReflect.Descriptor instead.
func (*GetProblemResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetProblemResponse) GetProblem() *Problem {
	if x != nil {
		return x.Problem
	}
	return nil
}

type GetStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *GetStatRequest) Reset() {
	*x = GetStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatRequest) ProtoMessage() {}

func (x *GetStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatRequest.ProtoReflect.Descriptor instead.
func (*GetStatRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetStatRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type TaskStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskNumber int64 `protobuf:"varint,1,opt,name=task_number,json=taskNumber,proto3" json:"task_number,omitempty"`
	Correct    int64 `protobuf:"varint,2,opt,name=correct,proto3" json:"correct,omitempty"`
	All        int64 `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *TaskStat) Reset() {
	*x = TaskStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStat) ProtoMessage() {}

func (x *TaskStat) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStat.ProtoReflect.Descriptor instead.
func (*TaskStat) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *TaskStat) GetTaskNumber() int64 {
	if x != nil {
		return x.TaskNumber
	}
	return 0
}

func (x *TaskStat) GetCorrect() int64 {
	if x != nil {
		return x.Correct
	}
	return 0
}

func (x *TaskStat) GetAll() int64 {
	if x != nil {
		return x.All
	}
	return 0
}

type GetStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat []*TaskStat `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
}

func (x *GetStatResponse) Reset() {
	*x = GetStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatResponse) ProtoMessage() {}

func (x *GetStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatResponse.ProtoReflect.Descriptor instead.
func (*GetStatResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetStatResponse) GetStat() []*TaskStat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type GetRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *GetRatingRequest) Reset() {
	*x = GetRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingRequest) ProtoMessage() {}

func (x *GetRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingRequest.ProtoReflect.Descriptor instead.
func (*GetRatingRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{10}
}

func (x *GetRatingRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type GetRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int64 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	All      int64 `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *GetRatingResponse) Reset() {
	*x = GetRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingResponse) ProtoMessage() {}

func (x *GetRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingResponse.ProtoReflect.Descriptor instead.
func (*GetRatingResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{11}
}

func (x *GetRatingResponse) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *GetRatingResponse) GetAll() int64 {
	if x != nil {
		return x.All
	}
	return 0
}

type ParseProblemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *ParseProblemsRequest) Reset() {
	*x = ParseProblemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseProblemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseProblemsRequest) ProtoMessage() {}

func (x *ParseProblemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseProblemsRequest.ProtoReflect.Descriptor instead.
func (*ParseProblemsRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{12}
}

func (x *ParseProblemsRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type ParseProblemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ParseProblemsResponse) Reset() {
	*x = ParseProblemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseProblemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseProblemsResponse) ProtoMessage() {}

func (x *ParseProblemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseProblemsResponse.ProtoReflect.Descriptor instead.
func (*ParseProblemsResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{13}
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId int64 `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{14}
}

func (x *GetImageRequest) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{15}
}

func (x *GetImageResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x60, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x77, 0x0a,
	0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x61, 0x73, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x22, 0x29,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x08, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61,
	0x6c, 0x6c, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x22, 0x2f, 0x0a, 0x14, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2a, 0x26, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x52, 0x52, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x4f, 0x4e, 0x47,
	0x10, 0x01, 0x32, 0xf4, 0x02, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x68, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x76, 0x6a,
	0x6f, 0x6e, 0x6f, 0x6b, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_api_proto_goTypes = []interface{}{
	(SubmitResult)(0),             // 0: api.SubmitResult
	(*Problem)(nil),               // 1: api.Problem
	(*CheckAnswerRequest)(nil),    // 2: api.CheckAnswerRequest
	(*CheckAnswerResponse)(nil),   // 3: api.CheckAnswerResponse
	(*GetRandomRequest)(nil),      // 4: api.GetRandomRequest
	(*GetRandomResponse)(nil),     // 5: api.GetRandomResponse
	(*GetProblemRequest)(nil),     // 6: api.GetProblemRequest
	(*GetProblemResponse)(nil),    // 7: api.GetProblemResponse
	(*GetStatRequest)(nil),        // 8: api.GetStatRequest
	(*TaskStat)(nil),              // 9: api.TaskStat
	(*GetStatResponse)(nil),       // 10: api.GetStatResponse
	(*GetRatingRequest)(nil),      // 11: api.GetRatingRequest
	(*GetRatingResponse)(nil),     // 12: api.GetRatingResponse
	(*ParseProblemsRequest)(nil),  // 13: api.ParseProblemsRequest
	(*ParseProblemsResponse)(nil), // 14: api.ParseProblemsResponse
	(*GetImageRequest)(nil),       // 15: api.GetImageRequest
	(*GetImageResponse)(nil),      // 16: api.GetImageResponse
}
var file_api_api_proto_depIdxs = []int32{
	0,  // 0: api.CheckAnswerResponse.result:type_name -> api.SubmitResult
	1,  // 1: api.GetRandomResponse.problem:type_name -> api.Problem
	1,  // 2: api.GetProblemResponse.problem:type_name -> api.Problem
	9,  // 3: api.GetStatResponse.stat:type_name -> api.TaskStat
	11, // 4: api.MathHelper.GetRating:input_type -> api.GetRatingRequest
	8,  // 5: api.MathHelper.GetStat:input_type -> api.GetStatRequest
	2,  // 6: api.MathHelper.CheckAnswer:input_type -> api.CheckAnswerRequest
	4,  // 7: api.MathHelper.GetRandom:input_type -> api.GetRandomRequest
	6,  // 8: api.MathHelper.GetProblem:input_type -> api.GetProblemRequest
	15, // 9: api.MathHelper.GetImage:input_type -> api.GetImageRequest
	12, // 10: api.MathHelper.GetRating:output_type -> api.GetRatingResponse
	10, // 11: api.MathHelper.GetStat:output_type -> api.GetStatResponse
	3,  // 12: api.MathHelper.CheckAnswer:output_type -> api.CheckAnswerResponse
	5,  // 13: api.MathHelper.GetRandom:output_type -> api.GetRandomResponse
	7,  // 14: api.MathHelper.GetProblem:output_type -> api.GetProblemResponse
	16, // 15: api.MathHelper.GetImage:output_type -> api.GetImageResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAnswerRequest); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAnswerResponse); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomRequest); i {
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
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomResponse); i {
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
		file_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProblemRequest); i {
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
		file_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProblemResponse); i {
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
		file_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatRequest); i {
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
		file_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskStat); i {
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
		file_api_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatResponse); i {
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
		file_api_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingRequest); i {
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
		file_api_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatingResponse); i {
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
		file_api_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseProblemsRequest); i {
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
		file_api_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseProblemsResponse); i {
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
		file_api_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_api_api_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageResponse); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		EnumInfos:         file_api_api_proto_enumTypes,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
