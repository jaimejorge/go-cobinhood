package cobinhood

type Info struct {
	Phase    string  `json:"phase"`
	Revision string `json:"revision"`
}

func (c *CobinhoodClient) Getinfo() (Info, error) {
	var Respond Respond

	err := c.Get("/v1/system/info", &Respond)

	if err != nil {
		return Respond.Result.Info, err
	}

	return Respond.Result.Info, nil
}
