package service

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/infra/rpc"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	godotenv.Load("../../.env")
	rpc.InitClient()
	mysql.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value
	item := &cart.CartItem{ProductId: 2, Quantity: 1}
	req := &cart.AddItemReq{
		UserId: 2,
		Item:   item,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
