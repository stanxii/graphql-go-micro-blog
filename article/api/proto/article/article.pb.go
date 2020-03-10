// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/article/article.proto

package go_micro_srv_article

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

type TotalRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalRequest) Reset()         { *m = TotalRequest{} }
func (m *TotalRequest) String() string { return proto.CompactTextString(m) }
func (*TotalRequest) ProtoMessage()    {}
func (*TotalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{0}
}

func (m *TotalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalRequest.Unmarshal(m, b)
}
func (m *TotalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalRequest.Marshal(b, m, deterministic)
}
func (m *TotalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalRequest.Merge(m, src)
}
func (m *TotalRequest) XXX_Size() int {
	return xxx_messageInfo_TotalRequest.Size(m)
}
func (m *TotalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TotalRequest proto.InternalMessageInfo

type TotalResponse struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalResponse) Reset()         { *m = TotalResponse{} }
func (m *TotalResponse) String() string { return proto.CompactTextString(m) }
func (*TotalResponse) ProtoMessage()    {}
func (*TotalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{1}
}

func (m *TotalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalResponse.Unmarshal(m, b)
}
func (m *TotalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalResponse.Marshal(b, m, deterministic)
}
func (m *TotalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalResponse.Merge(m, src)
}
func (m *TotalResponse) XXX_Size() int {
	return xxx_messageInfo_TotalResponse.Size(m)
}
func (m *TotalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalResponse proto.InternalMessageInfo

func (m *TotalResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ArticlesRequest struct {
	Page                 int64           `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64           `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Filter               *ArticlesFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ArticlesRequest) Reset()         { *m = ArticlesRequest{} }
func (m *ArticlesRequest) String() string { return proto.CompactTextString(m) }
func (*ArticlesRequest) ProtoMessage()    {}
func (*ArticlesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{2}
}

func (m *ArticlesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlesRequest.Unmarshal(m, b)
}
func (m *ArticlesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlesRequest.Marshal(b, m, deterministic)
}
func (m *ArticlesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlesRequest.Merge(m, src)
}
func (m *ArticlesRequest) XXX_Size() int {
	return xxx_messageInfo_ArticlesRequest.Size(m)
}
func (m *ArticlesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlesRequest proto.InternalMessageInfo

func (m *ArticlesRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ArticlesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ArticlesRequest) GetFilter() *ArticlesFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ArticlesFilter struct {
	// 文章标题
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 文章是否删除
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// 分类id
	CategoryId string `protobuf:"bytes,3,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	// 排序
	Asc                  bool     `protobuf:"varint,4,opt,name=asc,proto3" json:"asc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticlesFilter) Reset()         { *m = ArticlesFilter{} }
func (m *ArticlesFilter) String() string { return proto.CompactTextString(m) }
func (*ArticlesFilter) ProtoMessage()    {}
func (*ArticlesFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{3}
}

func (m *ArticlesFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlesFilter.Unmarshal(m, b)
}
func (m *ArticlesFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlesFilter.Marshal(b, m, deterministic)
}
func (m *ArticlesFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlesFilter.Merge(m, src)
}
func (m *ArticlesFilter) XXX_Size() int {
	return xxx_messageInfo_ArticlesFilter.Size(m)
}
func (m *ArticlesFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlesFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlesFilter proto.InternalMessageInfo

func (m *ArticlesFilter) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticlesFilter) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *ArticlesFilter) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *ArticlesFilter) GetAsc() bool {
	if m != nil {
		return m.Asc
	}
	return false
}

type ArticlesResponse struct {
	Articles             []*ArticleResponse `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ArticlesResponse) Reset()         { *m = ArticlesResponse{} }
func (m *ArticlesResponse) String() string { return proto.CompactTextString(m) }
func (*ArticlesResponse) ProtoMessage()    {}
func (*ArticlesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{4}
}

func (m *ArticlesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlesResponse.Unmarshal(m, b)
}
func (m *ArticlesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlesResponse.Marshal(b, m, deterministic)
}
func (m *ArticlesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlesResponse.Merge(m, src)
}
func (m *ArticlesResponse) XXX_Size() int {
	return xxx_messageInfo_ArticlesResponse.Size(m)
}
func (m *ArticlesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlesResponse proto.InternalMessageInfo

func (m *ArticlesResponse) GetArticles() []*ArticleResponse {
	if m != nil {
		return m.Articles
	}
	return nil
}

type ArticleResponse struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Title                string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Excerpt              string               `protobuf:"bytes,4,opt,name=Excerpt,proto3" json:"Excerpt,omitempty"`
	CreateAt             string               `protobuf:"bytes,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt             string               `protobuf:"bytes,6,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
	UserId               int64                `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	Categorys            []*CategorysResponse `protobuf:"bytes,8,rep,name=categorys,proto3" json:"categorys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ArticleResponse) Reset()         { *m = ArticleResponse{} }
func (m *ArticleResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleResponse) ProtoMessage()    {}
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{5}
}

func (m *ArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleResponse.Unmarshal(m, b)
}
func (m *ArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleResponse.Marshal(b, m, deterministic)
}
func (m *ArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleResponse.Merge(m, src)
}
func (m *ArticleResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleResponse.Size(m)
}
func (m *ArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleResponse proto.InternalMessageInfo

func (m *ArticleResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ArticleResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleResponse) GetExcerpt() string {
	if m != nil {
		return m.Excerpt
	}
	return ""
}

func (m *ArticleResponse) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *ArticleResponse) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *ArticleResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ArticleResponse) GetCategorys() []*CategorysResponse {
	if m != nil {
		return m.Categorys
	}
	return nil
}

type CategorysResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategorysResponse) Reset()         { *m = CategorysResponse{} }
func (m *CategorysResponse) String() string { return proto.CompactTextString(m) }
func (*CategorysResponse) ProtoMessage()    {}
func (*CategorysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{6}
}

func (m *CategorysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategorysResponse.Unmarshal(m, b)
}
func (m *CategorysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategorysResponse.Marshal(b, m, deterministic)
}
func (m *CategorysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategorysResponse.Merge(m, src)
}
func (m *CategorysResponse) XXX_Size() int {
	return xxx_messageInfo_CategorysResponse.Size(m)
}
func (m *CategorysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategorysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategorysResponse proto.InternalMessageInfo

func (m *CategorysResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategorysResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ArticleRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleRequest) Reset()         { *m = ArticleRequest{} }
func (m *ArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ArticleRequest) ProtoMessage()    {}
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{7}
}

func (m *ArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleRequest.Unmarshal(m, b)
}
func (m *ArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleRequest.Marshal(b, m, deterministic)
}
func (m *ArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleRequest.Merge(m, src)
}
func (m *ArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ArticleRequest.Size(m)
}
func (m *ArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleRequest proto.InternalMessageInfo

func (m *ArticleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateArticleRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Excerpt              string   `protobuf:"bytes,3,opt,name=Excerpt,proto3" json:"Excerpt,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	CategoryIds          []int64  `protobuf:"varint,5,rep,packed,name=categoryIds,proto3" json:"categoryIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateArticleRequest) Reset()         { *m = CreateArticleRequest{} }
func (m *CreateArticleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArticleRequest) ProtoMessage()    {}
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{8}
}

func (m *CreateArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArticleRequest.Unmarshal(m, b)
}
func (m *CreateArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArticleRequest.Marshal(b, m, deterministic)
}
func (m *CreateArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArticleRequest.Merge(m, src)
}
func (m *CreateArticleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArticleRequest.Size(m)
}
func (m *CreateArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArticleRequest proto.InternalMessageInfo

func (m *CreateArticleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateArticleRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateArticleRequest) GetExcerpt() string {
	if m != nil {
		return m.Excerpt
	}
	return ""
}

func (m *CreateArticleRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateArticleRequest) GetCategoryIds() []int64 {
	if m != nil {
		return m.CategoryIds
	}
	return nil
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{9}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{10}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{11}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bfa7f65df4b480f, []int{12}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*TotalRequest)(nil), "go.micro.srv.article.TotalRequest")
	proto.RegisterType((*TotalResponse)(nil), "go.micro.srv.article.TotalResponse")
	proto.RegisterType((*ArticlesRequest)(nil), "go.micro.srv.article.ArticlesRequest")
	proto.RegisterType((*ArticlesFilter)(nil), "go.micro.srv.article.ArticlesFilter")
	proto.RegisterType((*ArticlesResponse)(nil), "go.micro.srv.article.ArticlesResponse")
	proto.RegisterType((*ArticleResponse)(nil), "go.micro.srv.article.ArticleResponse")
	proto.RegisterType((*CategorysResponse)(nil), "go.micro.srv.article.CategorysResponse")
	proto.RegisterType((*ArticleRequest)(nil), "go.micro.srv.article.ArticleRequest")
	proto.RegisterType((*CreateArticleRequest)(nil), "go.micro.srv.article.CreateArticleRequest")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.article.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.article.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.article.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.article.Pong")
}

func init() { proto.RegisterFile("proto/article/article.proto", fileDescriptor_6bfa7f65df4b480f) }

var fileDescriptor_6bfa7f65df4b480f = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x63, 0x27, 0x75, 0xa6, 0x34, 0xa4, 0xab, 0x08, 0x59, 0x46, 0xaa, 0xac, 0xa5, 0xa5,
	0x86, 0x83, 0xa9, 0xc2, 0x81, 0x0b, 0x97, 0xa8, 0x2a, 0xa2, 0x37, 0xe4, 0x80, 0x84, 0x84, 0x38,
	0x18, 0x7b, 0xb1, 0x2c, 0x12, 0x6f, 0xd8, 0xdd, 0x54, 0xc0, 0x05, 0x89, 0xcf, 0xe0, 0x03, 0xf8,
	0x4e, 0xe4, 0xf5, 0xae, 0xed, 0xa4, 0x76, 0xd2, 0x53, 0xf6, 0xcd, 0x9b, 0x19, 0xcf, 0xbc, 0x9d,
	0xd9, 0xc0, 0xe3, 0x15, 0xa3, 0x82, 0xbe, 0x88, 0x98, 0xc8, 0xe2, 0x05, 0xd1, 0xbf, 0x81, 0xb4,
	0xa2, 0x49, 0x4a, 0x83, 0x65, 0x16, 0x33, 0x1a, 0x70, 0x76, 0x1b, 0x28, 0x0e, 0x8f, 0xe0, 0xc1,
	0x7b, 0x2a, 0xa2, 0x45, 0x48, 0xbe, 0xaf, 0x09, 0x17, 0xf8, 0x1c, 0x8e, 0x15, 0xe6, 0x2b, 0x9a,
	0x73, 0x82, 0x26, 0xd0, 0x17, 0x85, 0xc1, 0x31, 0x3c, 0xc3, 0x37, 0xc3, 0x12, 0xe0, 0xdf, 0xf0,
	0x70, 0x56, 0x66, 0xe0, 0x2a, 0x12, 0x21, 0xb0, 0x56, 0x51, 0x4a, 0x94, 0x9f, 0x3c, 0x23, 0x17,
	0xec, 0xe2, 0x77, 0x9e, 0xfd, 0x22, 0x4e, 0x4f, 0xda, 0x2b, 0x8c, 0x5e, 0xc3, 0xe0, 0x6b, 0xb6,
	0x10, 0x84, 0x39, 0xa6, 0x67, 0xf8, 0x47, 0xd3, 0xb3, 0xa0, 0xad, 0xc0, 0x40, 0x7f, 0xe6, 0x8d,
	0xf4, 0x0d, 0x55, 0x0c, 0x66, 0x30, 0xda, 0x64, 0x64, 0xa1, 0x99, 0x58, 0x94, 0x05, 0x0c, 0xc3,
	0x12, 0x20, 0x07, 0x0e, 0x13, 0xb2, 0x20, 0x82, 0x24, 0xb2, 0x00, 0x3b, 0xd4, 0x10, 0x9d, 0x02,
	0xc4, 0x91, 0x20, 0x29, 0x65, 0x3f, 0x6f, 0x12, 0x59, 0xc3, 0x30, 0x6c, 0x58, 0xd0, 0x18, 0xcc,
	0x88, 0xc7, 0x8e, 0x25, 0xa3, 0x8a, 0x23, 0xfe, 0x00, 0xe3, 0xba, 0x69, 0x25, 0xcf, 0x0c, 0x6c,
	0x55, 0x29, 0x77, 0x0c, 0xcf, 0xf4, 0x8f, 0xa6, 0xe7, 0x3b, 0xfb, 0xd0, 0x81, 0x61, 0x15, 0x86,
	0xff, 0xf4, 0x2a, 0x31, 0xab, 0xb4, 0x23, 0xe8, 0x65, 0x89, 0x92, 0xb2, 0x97, 0x25, 0x45, 0x1b,
	0x31, 0xcd, 0x05, 0xc9, 0x85, 0x6c, 0x63, 0x18, 0x6a, 0x58, 0xb7, 0x6d, 0x6e, 0xb5, 0x7d, 0xfd,
	0x23, 0x26, 0x6c, 0x25, 0x64, 0x03, 0xc3, 0x50, 0xc3, 0xe2, 0x4a, 0x62, 0x46, 0x22, 0x41, 0x66,
	0xc2, 0xe9, 0x4b, 0xaa, 0xc2, 0x05, 0xb7, 0x5e, 0x25, 0x25, 0x37, 0x28, 0x39, 0x8d, 0xd1, 0x23,
	0x18, 0xac, 0x39, 0x61, 0x37, 0x89, 0x73, 0x28, 0xab, 0x52, 0x08, 0x5d, 0xc3, 0x50, 0x8b, 0xc6,
	0x1d, 0x5b, 0x2a, 0x70, 0xd1, 0xae, 0xc0, 0x95, 0x76, 0xab, 0x34, 0xa8, 0x23, 0xf1, 0x2b, 0x38,
	0xb9, 0xc3, 0xdf, 0x51, 0x01, 0x81, 0x95, 0x47, 0x4b, 0xa2, 0x24, 0x90, 0x67, 0xec, 0x55, 0x83,
	0xa0, 0x07, 0x71, 0x2b, 0x0a, 0xff, 0x35, 0x60, 0x72, 0x55, 0xb6, 0xb8, 0xe9, 0xd8, 0x39, 0x31,
	0x1d, 0x52, 0x37, 0x44, 0x35, 0x37, 0x45, 0xad, 0xc5, 0xb1, 0x36, 0xc4, 0xf1, 0xe0, 0xa8, 0x9e,
	0x28, 0xee, 0xf4, 0x3d, 0xd3, 0x37, 0xc3, 0xa6, 0x09, 0xfb, 0x30, 0x9e, 0x0b, 0x46, 0xa2, 0x65,
	0x96, 0xa7, 0x8d, 0xba, 0x62, 0xba, 0xce, 0x85, 0x5e, 0x39, 0x09, 0xf0, 0x33, 0x38, 0x69, 0x78,
	0xd6, 0xdb, 0xd9, 0xe2, 0x7a, 0x0a, 0xd6, 0xbb, 0x2c, 0x4f, 0x8b, 0xb2, 0xb8, 0x60, 0xf4, 0x9b,
	0x5e, 0x4a, 0x85, 0x24, 0x4f, 0xbb, 0xf9, 0xe9, 0x3f, 0x0b, 0x40, 0x69, 0x35, 0x67, 0xb7, 0x28,
	0x84, 0xbe, 0x7c, 0x13, 0x10, 0x6e, 0xbf, 0xd8, 0xe6, 0x03, 0xe2, 0x3e, 0xd9, 0xe9, 0x53, 0x96,
	0x8d, 0x0f, 0xd0, 0x27, 0xb0, 0xf5, 0x2e, 0xa1, 0xdd, 0x1b, 0xa3, 0x1f, 0x18, 0xf7, 0xe9, 0x3e,
	0xb7, 0x2a, 0xf9, 0x47, 0x38, 0x54, 0x56, 0x74, 0xb6, 0x67, 0x1b, 0xcb, 0xd4, 0xf7, 0xdb, 0x59,
	0x7c, 0x80, 0x12, 0x38, 0xde, 0x18, 0x25, 0xf4, 0xbc, 0x63, 0xd6, 0x5b, 0xe6, 0xed, 0xfe, 0x5f,
	0xf9, 0x0c, 0x83, 0xf2, 0xaa, 0x51, 0x47, 0xcf, 0xdb, 0x23, 0xe3, 0x5e, 0xec, 0xf5, 0xd3, 0xc9,
	0x2f, 0x0d, 0xf4, 0x16, 0xec, 0x62, 0x3c, 0xe4, 0x08, 0xb8, 0xed, 0x81, 0x05, 0xef, 0x76, 0x71,
	0x34, 0x4f, 0xf1, 0x81, 0x6f, 0x5c, 0x1a, 0x5f, 0x06, 0xf2, 0xaf, 0xe5, 0xe5, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0xac, 0xc6, 0x8a, 0x79, 0x06, 0x00, 0x00,
}
