package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

type CashPayment struct{}
type DebitCardPayment struct{}
type CreditCardPayment struct{}

func (c *CashPayment) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payment done using cash\n", amount)
}

func (c *DebitCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payment done using debit card\n", amount)
}

func (c *CreditCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payment done using credit card\n", amount)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPayment), nil
	case DebitCard:
		return new(DebitCardPayment), nil
	case CreditCard:
		return new(CreditCardPayment), nil
	default:
		return nil, fmt.Errorf("payment method %d not valid", m)
	}
}
