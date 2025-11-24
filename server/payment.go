package server

import (
	paymenthandler "github.com/tarlrsk/shop/modules/payment/paymentHandler"
	paymentrepository "github.com/tarlrsk/shop/modules/payment/paymentRepository"
	paymentusecase "github.com/tarlrsk/shop/modules/payment/paymentUseCase"
)

func (s *server) paymentService() {
	repo := paymentrepository.NewPaymentRepository(s.db)
	useCase := paymentusecase.NewPaymentUseCase(repo)
	httpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, useCase)

	_ = httpHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("", s.healthCheckService)
}
