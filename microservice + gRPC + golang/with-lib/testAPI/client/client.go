package client

import (
	"fmt"
	"ghtk/models"
	"net/http"

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

func (c *Client) CallGhtkApi(params interface{}, httpMethod string) (*models.BaseResponse, error) {
	res := models.BaseResponse{}

	extraHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if httpMethod == http.MethodGet {
		paramMap := map[string]interface{}{}
		if params != nil {
			paramMap = params.(map[string]interface{})
		}
		if err := xrpc.HttpGetWithJsonResp(c.BaseURL, paramMap, &res, xrpc.WithHeaders(extraHeader)); err != nil {
			return nil, err
		}
	} else if httpMethod == http.MethodPost {
		if err := xrpc.HttpPostWithJsonResp(c.BaseURL, params, &res, xrpc.WithHeaders(extraHeader)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Server responded with error: No matching http method ")

	}

	if res.Success == false {
		return nil, fmt.Errorf("Ghtk responded with error: %s", res.Success)
	}

	return &res, nil
}
