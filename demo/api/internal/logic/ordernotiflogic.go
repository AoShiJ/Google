package logic

import (
	"context"
	"demo/api/pkg"
	"fmt"
	"net/url"
	"order/rpc/orderclient"

	"demo/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderNotifLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderNotifLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderNotifLogic {
	return &OrderNotifLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderNotifLogic) OrderNotif(req url.Values) error {
	fmt.Println("****++++++++++++++++++++++++++++++收到回调")
	msg, err := pkg.VerifySign(req)
	if err != nil {
		return err
	}
	fmt.Println("*************回调信息**********")
	fmt.Println(msg)

	_, err = l.svcCtx.OrderServer.UpdateOrderByGoodsNo(l.ctx, &orderclient.UpdateOrderByGoodsNoRequest{
		Info: &orderclient.OrderInfo{
			OrderNo: msg.OutTradeNo,
			State:   2,
		},
	})

	return err
}
