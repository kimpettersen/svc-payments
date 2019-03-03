package storage

import (
	"testing"

	pb "github.com/kimpettersen/svc-payments/proto"
	"github.com/stretchr/testify/require"
)

func TestStorePayment(t *testing.T) {
	s := InMem{}
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
	s := InMem{}
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

func TestConfirmPayment(t *testing.T) {
	s := InMem{}
	payment, _ := s.StorePayment(&pb.Payment{
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	require.Equal(t, pb.Status_PENDING, payment.GetStatus())
	s.ConfirmPayment(payment.Id)
	require.Equal(t, pb.Status_COMPLETE, payment.GetStatus())
}

func TestGetAllPayments(t *testing.T) {
	s := InMem{}
	s.StorePayment(&pb.Payment{
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	s.StorePayment(&pb.Payment{
		Amount: 69437818999,
		From:   "44941657987",
		To:     "12195650677",
	})
	paymentList, err := s.GetAllPayments()
	require.NoError(t, err)
	require.Len(t, paymentList.Payments, 2)
	require.NotEqual(t, paymentList.Payments[0], paymentList.Payments[1])
}
