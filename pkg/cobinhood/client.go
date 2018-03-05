package cobinhood

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type CobinhoodClient struct {
	Host   string
	ApiKey string
}

func New(host, ApiKey string) *CobinhoodClient {
	return &CobinhoodClient{
		Host:   host,
		ApiKey: ApiKey,
	}
}

func (c *CobinhoodClient) Get(path string, out interface{}) error {
	req, err := c.request("GET", path, nil)

	if err != nil {
		return err
	}

	res, err := c.client().Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, out)

	if err != nil {
		return err
	}

	return nil
}

func (c *CobinhoodClient) PostBody(path string, body io.Reader, out interface{}) error {

	req, err := c.request("POST", path, body)
	if err != nil {
		return  err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.client().Do(req)

	if err != nil {
		return  err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return  err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return  err
	}

	err = json.Unmarshal(data, out)

	if err != nil {
		return  err
	}

	return  nil
}


func (c *CobinhoodClient) Delete(path string) error {

	req, err := c.request("DELETE", path, nil)

	if err != nil {
		return nil
	}

	res, err := c.client().Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := responseError(res); err != nil {
		return err
	}

	return nil
}



func (c *CobinhoodClient) client() *http.Client {
	client := &http.Client{}
	return client
}

func (c *CobinhoodClient) request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://%s%s", c.Host, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if c.ApiKey != "" {
		req.Header.Add("Authorization", c.ApiKey)
	}

	req.Header.Add("nonce", fmt.Sprintf("%v", int32(time.Now().Unix())))

	return req, nil
}


