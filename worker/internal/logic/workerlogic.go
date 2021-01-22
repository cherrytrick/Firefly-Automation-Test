package logic

import (
	"context"

	"worker/internal/svc"
	"worker/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WorkerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) WorkerLogic {
	return WorkerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkerLogic) Worker(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
