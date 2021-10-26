package handler

import (
	"github.com/gzltommy/common"
	"github.com/gzltommy/order/domain/model"
	"github.com/gzltommy/order/domain/service"
	"github.com/gzltommy/order/proto/order"
	"context"
)

type Order struct {
	OrderDataService service.IOrderDataService
}

//根据订单ID查询订单
func (o *Order) GetOrderByID(ctx context.Context, request *order.OrderID, response *order.OrderInfo) error {
	order, err := o.OrderDataService.FindOrderByID(request.OrderId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(order, response); err != nil {
		return err
	}
	return nil
}

//查找所有订单
func (o *Order) GetAllOrder(ctx context.Context, request *order.AllOrderRequest, response *order.AllOrder) error {
	orderAll, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}

	for _, v := range orderAll {
		order := &order.OrderInfo{}
		if err := common.SwapTo(v, order); err != nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, order)
	}
	return nil
}

//创建订单
func (o *Order) CreateOrder(ctx context.Context, request *order.OrderInfo, response *order.OrderID) error {
	orderAdd := &model.Order{}
	if err := common.SwapTo(request, orderAdd); err != nil {
		return err
	}
	orderID, err := o.OrderDataService.AddOrder(orderAdd)
	if err != nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

//删除订单
func (o *Order) DeleteOrderByID(ctx context.Context, request *order.OrderID, response *order.Response) error {
	if err := o.OrderDataService.DeleteOrder(request.OrderId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

//更新订单支付状态
func (o *Order) UpdateOrderPayStatus(ctx context.Context, request *order.PayStatus, response *order.Response) error {
	if err := o.OrderDataService.UpdatePayStatus(request.OrderId, request.PayStatus); err != nil {
		return err
	}
	response.Msg = "支付状态更新成功"
	return nil
}

//更新发货状态
func (o *Order) UpdateOrderShipStatus(ctx context.Context, request *order.ShipStatus, response *order.Response) error {
	if err := o.OrderDataService.UpdateShipStatus(request.OrderId, request.ShipStatus); err != nil {
		return err
	}
	response.Msg = "发货状态更新成功"
	return nil
}

//更新订单状态
func (o *Order) UpdateOrder(ctx context.Context, request *order.OrderInfo, response *order.Response) error {
	order := &model.Order{}
	if err := common.SwapTo(request, order); err != nil {
		return err
	}
	if err := o.OrderDataService.UpdateOrder(order); err != nil {
		return err
	}
	response.Msg = "订单更新成功"
	return nil
}
