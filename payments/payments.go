package payments

import (
	"context"

	"github.com/kimpettersen/svc-payments/pkg/storage"
	"github.com/pkg/errors"

	pb "github.com/kimpettersen/svc-payments/proto"
)

type PaymentsService struct{}

func (p PaymentsService) Pay(ctx context.Context, in *pb.Payment) (*pb.Payment, error) {
	s := storage.Storage{}
	payment, err := s.StorePayment(in)
	if err != nil {
		errors.Wrap(err, "failed to create payment")
		return nil, err
	}
	return payment, nil
}
