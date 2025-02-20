package service

import (
	"context"
	"github.com/A1sca/Douyin-Mall-Go/app/payment/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/payment/biz/model"
	payment "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
}

// NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(4004001, err.Error())
	}

	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005001, err.Error())
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: translationId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005002, err.Error())
	}

	return &payment.ChargeResp{TransactionId: translationId.String()}, nil

}
