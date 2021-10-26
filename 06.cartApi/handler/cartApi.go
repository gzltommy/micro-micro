package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gzltommy/cart/proto/cart"
	log "github.com/micro/go-micro/v2/logger"
	api "github.com/gzltommy/cartApi/proto/imports"

	"strconv"
)

type CartApi struct {
	CartService cart.CartService
}

func (e *CartApi) FindAll(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("接收到 /cartApi/FindAll 访问请求")

	//// extract the client from the context
	//cartApiClient, ok := client.CartApiFromContext(ctx)
	//if !ok {
	//	return errors.InternalServerError("go.micro.api.cartApi.cartApi.call", "cartApi client not found")
	//}
	//
	//// make request
	//response, err := cartApiClient.Call(ctx, &cartApi.Request{
	//	Name: extractValue(req.Post["name"]),
	//})
	//if err != nil {
	//	return errors.InternalServerError("go.micro.api.cartApi.cartApi.call", err.Error())
	//}
	//
	//b, _ := json.Marshal(response)
	//
	//rsp.StatusCode = 200
	//rsp.Body = string(b)
	v, ok := req.Get["user_id"]
	if !ok {
		//rsp.StatusCode = 500
		return errors.New("参数异常")
	}

	userIdString := v.Values[0]
	fmt.Println("======", userIdString)

	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}

	// 获取购物车所有商品
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	if err != nil {
		return err
	}
	
	// 数据类型转换
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
