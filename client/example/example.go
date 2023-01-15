// Code generated by goctl. DO NOT EDIT.
// Source: example.proto

package client

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResp          = example.BaseResp
	Empty             = example.Empty
	IDReq             = example.IDReq
	IDsReq            = example.IDsReq
	PageInfoReq       = example.PageInfoReq
	StatusCodeReq     = example.StatusCodeReq
	StatusCodeUUIDReq = example.StatusCodeUUIDReq
	StudentInfo       = example.StudentInfo
	StudentListResp   = example.StudentListResp
	StudentPageReq    = example.StudentPageReq
	TeacherInfo       = example.TeacherInfo
	TeacherListResp   = example.TeacherListResp
	TeacherPageReq    = example.TeacherPageReq
	UUIDReq           = example.UUIDReq
	UUIDsReq          = example.UUIDsReq

	Example interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Student management
		CreateOrUpdateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetStudentList(ctx context.Context, in *StudentPageReq, opts ...grpc.CallOption) (*StudentListResp, error)
		DeleteStudent(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		BatchDeleteStudent(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultExample struct {
		cli zrpc.Client
	}
)

func NewExample(cli zrpc.Client) Example {
	return &defaultExample{
		cli: cli,
	}
}

func (m *defaultExample) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Student management
func (m *defaultExample) CreateOrUpdateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.CreateOrUpdateStudent(ctx, in, opts...)
}

func (m *defaultExample) GetStudentList(ctx context.Context, in *StudentPageReq, opts ...grpc.CallOption) (*StudentListResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.GetStudentList(ctx, in, opts...)
}

func (m *defaultExample) DeleteStudent(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.DeleteStudent(ctx, in, opts...)
}

func (m *defaultExample) BatchDeleteStudent(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.BatchDeleteStudent(ctx, in, opts...)
}