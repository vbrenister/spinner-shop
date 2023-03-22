package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	StatusId       int
	Amount         int
	Currency       string
	LastFour       string
	BankReturnCode string
}

func (card *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return card.CreatePaymentIntent(currency, amount)
}

func (card *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = card.Secret

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		var msg string
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}

		return nil, msg, err
	}

	return pi, "", nil
}

func cardErrorMessage(errorCode stripe.ErrorCode) string {
	switch errorCode {
	case stripe.ErrorCodeCardDeclined:
		return "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		return "Your card is expired"
	case stripe.ErrorCodeIncorrectCVC:
		return "Incorrect CVV code"
	case stripe.ErrorCodeIncorrectZip:
		return "Incorrect zip/postal code"
	case stripe.ErrorCodeAmountTooLarge:
		return "The amount is too large to be charged to your card"
	case stripe.ErrorCodeAmountTooSmall:
		return "The amount is too small to be charged to your card"
	case stripe.ErrorCodeBalanceInsufficient:
		return "Insufficient balance"
	case stripe.ErrorCodePostalCodeInvalid:
		return "Invalid postal code"
	default:
		return "Your card was declined"
	}

}
