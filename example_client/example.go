// Code generated by goctl. DO NOT EDIT.
// Source: example.proto

package example_client

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp      = example.BaseIDResp
	BaseResp        = example.BaseResp
	BaseUUIDResp    = example.BaseUUIDResp
	Empty           = example.Empty
	IDReq           = example.IDReq
	IDsReq          = example.IDsReq
	PageInfoReq     = example.PageInfoReq
	StudentInfo     = example.StudentInfo
	StudentListReq  = example.StudentListReq
	StudentListResp = example.StudentListResp
	TeacherInfo     = example.TeacherInfo
	TeacherListReq  = example.TeacherListReq
	TeacherListResp = example.TeacherListResp
	UUIDReq         = example.UUIDReq
	UUIDsReq        = example.UUIDsReq

	Example interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Student management
		CreateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetStudentList(ctx context.Context, in *StudentListReq, opts ...grpc.CallOption) (*StudentListResp, error)
		GetStudentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*StudentInfo, error)
		DeleteStudent(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Teacher management
		CreateTeacher(ctx context.Context, in *TeacherInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateTeacher(ctx context.Context, in *TeacherInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetTeacherList(ctx context.Context, in *TeacherListReq, opts ...grpc.CallOption) (*TeacherListResp, error)
		GetTeacherById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TeacherInfo, error)
		DeleteTeacher(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
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
func (m *defaultExample) CreateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.CreateStudent(ctx, in, opts...)
}

func (m *defaultExample) UpdateStudent(ctx context.Context, in *StudentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.UpdateStudent(ctx, in, opts...)
}

func (m *defaultExample) GetStudentList(ctx context.Context, in *StudentListReq, opts ...grpc.CallOption) (*StudentListResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.GetStudentList(ctx, in, opts...)
}

func (m *defaultExample) GetStudentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*StudentInfo, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.GetStudentById(ctx, in, opts...)
}

func (m *defaultExample) DeleteStudent(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.DeleteStudent(ctx, in, opts...)
}

// Teacher management
func (m *defaultExample) CreateTeacher(ctx context.Context, in *TeacherInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.CreateTeacher(ctx, in, opts...)
}

func (m *defaultExample) UpdateTeacher(ctx context.Context, in *TeacherInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.UpdateTeacher(ctx, in, opts...)
}

func (m *defaultExample) GetTeacherList(ctx context.Context, in *TeacherListReq, opts ...grpc.CallOption) (*TeacherListResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.GetTeacherList(ctx, in, opts...)
}

func (m *defaultExample) GetTeacherById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TeacherInfo, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.GetTeacherById(ctx, in, opts...)
}

func (m *defaultExample) DeleteTeacher(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.DeleteTeacher(ctx, in, opts...)
}