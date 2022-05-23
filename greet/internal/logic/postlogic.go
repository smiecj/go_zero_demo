package logic

import (
	"context"
	"fmt"

	"github.com/smiecj/go_zero_demo/greet/internal/svc"
	"github.com/smiecj/go_zero_demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLogic {
	return &PostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) Post(req *types.Query) (resp *types.Response, err error) {
	return &types.Response{
		Message: fmt.Sprintf("post id: %s", req.Id),
	}, nil
}
