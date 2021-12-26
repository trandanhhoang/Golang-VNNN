package client

import (
	"fmt"
	"ghtk/models"
	"net/http"

	"github.com/xbonlinenet/goup/frame/xrpc"
)

type Client struct {
	BaseURL string
	Token   string
}

func NewClient(BaseURL string, Token string) *Client {
	return &Client{
		BaseURL: BaseURL,
		Token:   Token,
	}
}

func (c *Client) CallGhtkApi(params interface{}, httpMethod string) (*models.GhtkBaseResponse, error) {
	res := models.GhtkBaseResponse{}

	extraHeader := map[string]string{
		"Token":        c.Token,
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
		return nil, fmt.Errorf("Ghtk responded with error: %s", res.Message)
	}

	return &res, nil
}

func (c *Client) CreateGHTKOrder(orderParams interface{}, orderDetailResp *models.GhtkOrderDetail) error {
	//create url

	// CallGhtkApi

	//extract field from base response
	return nil
}
