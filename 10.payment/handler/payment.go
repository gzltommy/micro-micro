package handler

import (
	"github.com/gzltommy/common"
	"github.com/gzltommy/payment/domain/model"
	"github.com/gzltommy/payment/domain/service"
	"github.com/gzltommy/payment/proto/payment"
	"context"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (e *Payment) AddPayment(ctx context.Context, request *payment.PaymentInfo, response *payment.PaymentID) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		common.Error(err)
	}
	paymentID, err := e.PaymentDataService.AddPayment(payment)
	if err != nil {
		common.Error(err)
	}
	response.PaymentId = paymentID
	return nil
}

func (e *Payment) UpdatePayment(ctx context.Context, request *payment.PaymentInfo, response *payment.Response) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		common.Error(err)
		return err
	}
	return e.PaymentDataService.UpdatePayment(payment)
}

func (e *Payment) DeletePaymentByID(ctx context.Context, request *payment.PaymentID, response *payment.Response) error {
	return e.PaymentDataService.DeletePayment(request.PaymentId)
}

func (e *Payment) FindPaymentByID(ctx context.Context, request *payment.PaymentID, response *payment.PaymentInfo) error {
	payment, err := e.PaymentDataService.FindPaymentByID(request.PaymentId)
	if err != nil {
		common.Error(err)
		return err
	}
	return common.SwapTo(payment, response)
}

func (e *Payment) FindAllPayment(ctx context.Context, request *payment.All, response *payment.PaymentAll) error {
	allPayment, err := e.PaymentDataService.FindAllPayment()
	if err != nil {
		common.Error(err)
		return err
	}

	for _, v := range allPayment {
		paymentInfo := &payment.PaymentInfo{}
		if err := common.SwapTo(v, paymentInfo); err != nil {
			common.Error(err)
			continue
		}
		response.PaymentInfo = append(response.PaymentInfo, paymentInfo)
	}
	return nil
}
