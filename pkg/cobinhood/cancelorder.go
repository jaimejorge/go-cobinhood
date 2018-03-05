package cobinhood

import (
	"fmt"
)

func (c *CobinhoodClient) CancelOrder(orderid string) error {

	err := c.Delete(fmt.Sprintf("/v1/trading/orders/%s", orderid))

	if err != nil {
		return err
	}

	return nil
}
