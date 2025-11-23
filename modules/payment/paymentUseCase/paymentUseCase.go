package paymentusecase

import paymentrepository "github.com/tarlrsk/shop/modules/payment/paymentRepository"

type (
	PaymentUseCaseService interface {
	}

	paymentUseCase struct {
		paymentRepository paymentrepository.PaymentRepositoryService
	}
)

func NewPaymentUseCase(paymentRepository paymentrepository.PaymentRepositoryService) PaymentUseCaseService {
	return &paymentUseCase{
		paymentRepository: paymentRepository,
	}
}
