package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type createChargeRequest struct {
	AmountInCents int64 `json:"amount_in_cents"`
}

func init() {
	key := envy.Get("STRIPE_API_KEY", "")
	stripe.Key = key
}

// CreateChargeHandler create a stripe charge
func CreateChargeHandler(c buffalo.Context) error {
	re := &createChargeRequest{}

	if err := c.Bind(re); err != nil {
		return err
	}

	c.Logger().Debug("amount_in_cents: ", re.AmountInCents)

	chargeParams := stripe.ChargeParams{
		Amount:      stripe.Int64(re.AmountInCents),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("test"),
		Customer:    stripe.String("cus_DHYhnS8HyBYLqp"),
	}

	ch, err := charge.New(&chargeParams)

	if err != nil {
		return err
	}

	return c.Render(201, r.JSON(ch))
}
