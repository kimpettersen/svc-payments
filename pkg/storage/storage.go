package storage

import (
	"errors"
	"time"

	pb "github.com/kimpettersen/svc-payments/proto"
)

type Storage struct {
	Payments []pb.Payment
}

func (s *Storage) StorePayment(paymentRequest *pb.Payment) (*pb.Payment, error) {
	time.Sleep(5 * time.Second)
	paymentRequest.Status = pb.Status_COMPLETE
	s.Payments = append(s.Payments, *paymentRequest)
	return paymentRequest, nil
}

func (s *Storage) GetPaymentById(id string) (*pb.Payment, error) {

	for _, payment := range s.Payments {
		if payment.Id == id {
			return &payment, nil
		}
	}
	return nil, errors.New("Not found")
}
