package client

import (
	"fmt"
	"weather/models"

	"github.com/xbonlinenet/goup/frame/xrpc"
)

type Client struct {
	BaseURL string
}

func NewClient(BaseURL string) *Client {
	return &Client{
		BaseURL: BaseURL,
	}
}

func (c *Client) PostOrder(orderRequest *models.Weather) (*models.WeatherBaseResponse, error) {

	res := models.WeatherBaseResponse{}

	extraHeader := map[string]string{
		"Content-Type": "application/json",
	}

	if err := xrpc.HttpPostWithJsonResp(c.BaseURL, orderRequest, &res, xrpc.WithHeaders(extraHeader)); err != nil {
		return nil, err
	}

	if res.Success == false {
		return nil, fmt.Errorf("Ghtk responded with error: %s", res.Message)
	}

	return &res, nil
}
