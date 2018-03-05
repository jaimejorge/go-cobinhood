package cobinhood

type Currency struct {
	Currency       string `json:"currency"`
	Name           string `json:"name"`
	MinUnit       string `json:"min_unit"`
	DepositFee    string `json:"deposit_fee"`
	WithdrawalFee string `json:"withdrawal_fee"`
}

type Currencies []Currency

func (c *CobinhoodClient) Getcurrencies() (Currencies, error) {
	var Respond Respond

	err := c.Get("/v1/market/currencies", &Respond)

	if err != nil {
		return Respond.Result.Currencies, err
	}

	return Respond.Result.Currencies, nil
}
