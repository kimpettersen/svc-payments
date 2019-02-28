package storage

import (
	"testing"

	pb "github.com/kimpettersen/svc-payments/proto"
	"github.com/stretchr/testify/require"
)

func TestStorePayment(t *testing.T) {
	s := Storage{}
	payment, err := s.StorePayment(&pb.Payment{
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	require.NoError(t, err)
	require.Equal(t, int64(69437818999), payment.Amount)
	require.Equal(t, "44941657987", payment.From)
	require.Equal(t, "12195650677", payment.To)
	require.Len(t, s.Payments, 1)
	payment, _ = s.StorePayment(&pb.Payment{
		Amount: 176525,
		From:   "156436",
		To:     "84998",
	})
	require.Len(t, s.Payments, 2)
}

func TestGetPayment(t *testing.T) {
	s := Storage{}
	p1, _ := s.StorePayment(&pb.Payment{

		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	p2, _ := s.StorePayment(&pb.Payment{

		Amount: 176525,
		From:   "156436",
		To:     "84998",
	})
	pr, _ := s.GetPaymentById(p1.Id)
	require.Equal(t, int64(69437818999), pr.Amount)
	pr, _ = s.GetPaymentById(p2.Id)
	require.Equal(t, "156436", pr.From)
}
