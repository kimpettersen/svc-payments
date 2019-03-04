package storage

import (
	"errors"

	"github.com/google/uuid"
	pb "github.com/kimpettersen/svc-payments/proto"
)

type InMem struct {
	Payments []*pb.Payment
}

func (i *InMem) StorePayment(paymentRequest *pb.PaymentRequest) (*pb.Payment, error) {
	id := uuid.New().String()
	payment := &pb.Payment{
		Id:     id,
		Status: pb.Status_PENDING,
		Amount: paymentRequest.Amount,
		From:   paymentRequest.From,
		To:     paymentRequest.To,
	}

	i.Payments = append(i.Payments, payment)
	return payment, nil
}

func (i *InMem) GetPaymentById(id string) (*pb.Payment, error) {
	for _, payment := range i.Payments {
		if payment.Id == id {
			return payment, nil
		}
	}
	return nil, errors.New("Not found")
}

func (i *InMem) ConfirmPayment(id string) (*pb.Payment, error) {
	payment, err := i.GetPaymentById(id)
	if err != nil {
		return nil, errors.New("Not found")
	}
	payment.Status = pb.Status_COMPLETE
	return payment, nil
}

func (i *InMem) GetAllPayments() (*pb.PaymentList, error) {
	return &pb.PaymentList{
		Payments: i.Payments,
	}, nil
}
