package service

import (
	"context"
	"github.com/A1sca/Douyin-Mall-Go/app/checkout/infra/rpc"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/cart"
	checkout "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/checkout"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/payment"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
}

// NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		return nil, kerrors.NewBizStatusError(5004001, "cart is empty")
	}

	var total float32

	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32((cartItem.Quantity))
		total += cost
	}
	var orderId string

	u, _ := uuid.NewUUID()
	orderId = u.String()

	payReq := &payment.ChargeReq{
		Amount:  total,
		UserId:  req.UserId,
		OrderId: orderId,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})

	if err != nil {
		klog.Error(err.Error)
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return
}
