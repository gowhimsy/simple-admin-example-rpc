package teacher

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/ent"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherLogic {
	return &UpdateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTeacherLogic) UpdateTeacher(in *example.TeacherInfo) (*example.BaseResp, error) {
	err := l.svcCtx.DB.Teacher.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
		SetNotEmptyName(in.Name).
		SetNotEmptyAge(int(in.Age)).
		SetNotEmptyAgeInt32(in.AgeInt32).
		SetNotEmptyAgeInt64(in.AgeInt64).
		SetNotEmptyAgeUint(uint(in.AgeUint)).
		SetNotEmptyAgeUint32(in.AgeUint32).
		SetNotEmptyAgeUint64(in.AgeUint64).
		SetNotEmptyWeightFloat(in.WeightFloat).
		SetNotEmptyWeightFloat32(in.WeightFloat32).
		SetClassID(uuidx.ParseUUIDString(in.ClassId)).
		SetEnrollAt(time.Unix(in.EnrollAt, 0)).
		SetStatusBool(in.StatusBool).
		Exec(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &example.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
