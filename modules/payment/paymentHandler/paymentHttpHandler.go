package paymenthandler

import (
	"github.com/tarlrsk/shop/config"
	paymentusecase "github.com/tarlrsk/shop/modules/payment/paymentUseCase"
)

type (
	PaymentHttpHandlerService interface {
	}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUseCase paymentusecase.PaymentUseCaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUseCase paymentusecase.PaymentUseCaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUseCase: paymentUseCase,
	}
}
