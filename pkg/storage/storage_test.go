package storage

import (
	"testing"

	pb "github.com/kimpettersen/svc-payments/proto"
	"github.com/stretchr/testify/require"
)

func TestStorePayment(t *testing.T) {
	s := Storage{}
	payment, err := s.StorePayment(&pb.Payment{
		Id:     "64449399248",
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	require.NoError(t, err)
	require.Equal(t, "64449399248", payment.Id)
	require.Equal(t, int64(69437818999), payment.Amount)
	require.Equal(t, "44941657987", payment.From)
	require.Equal(t, "12195650677", payment.To)
	require.Len(t, s.Payments, 1)
	payment, _ = s.StorePayment(&pb.Payment{
		Id:     "97595",
		Amount: 176525,
		From:   "156436",
		To:     "84998",
	})
	require.Len(t, s.Payments, 2)
}

func TestGetPayment(t *testing.T) {
	s := Storage{}
	s.StorePayment(&pb.Payment{
		Id:     "64449399248",
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	s.StorePayment(&pb.Payment{
		Id:     "97595",
		Amount: 176525,
		From:   "156436",
		To:     "84998",
	})
	pr, _ := s.GetPaymentById("64449399248")
	require.Equal(t, int64(69437818999), pr.Amount)
	pr, _ = s.GetPaymentById("97595")
	require.Equal(t, "156436", pr.From)
}
