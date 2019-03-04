package payments

import (
	"context"
	"log"

	"github.com/kimpettersen/svc-payments/pkg/storage"

	pb "github.com/kimpettersen/svc-payments/proto"
)

type PaymentsService struct {
	Storage *storage.InMem
}

func (p PaymentsService) Pay(ctx context.Context, in *pb.PaymentRequest) (*pb.Payment, error) {
	log.Printf("Request to Pay: %v\n", in)
	payment, err := p.Storage.StorePayment(in)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (p PaymentsService) GetById(ctx context.Context, in *pb.PaymentByIdRequest) (*pb.Payment, error) {
	log.Printf("Request to GetById: %v\n", in)
	id := in.GetId()
	payment, err := p.Storage.GetPaymentById(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (p PaymentsService) Confirm(ctx context.Context, in *pb.PaymentByIdRequest) (*pb.Payment, error) {
	log.Printf("Request to Confirm: %v\n", in)
	payment, err := p.Storage.ConfirmPayment(in.Id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (p PaymentsService) GetAll(ctx context.Context, in *pb.AllPaymentsRequest) (*pb.PaymentList, error) {
	log.Printf("Request to GetAll: %v\n", in)
	payments, err := p.Storage.GetAllPayments()
	if err != nil {
		return nil, err
	}
	return payments, nil
}
