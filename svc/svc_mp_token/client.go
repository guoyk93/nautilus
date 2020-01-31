package svc_mp_token

import (
	"context"
	"go.guoyk.net/nrpc/v2"
)

type MPTokenClient struct {
	client *nrpc.Client
}

func NewMPTokenClient(c *nrpc.Client) *MPTokenClient {
	return &MPTokenClient{client: c}
}

func (c *MPTokenClient) Token() (token string, err error) {
	out := GetResp{}
	if err = c.client.Query("MPTokenService.Get", nil, &out).Do(context.Background()); err != nil {
		return
	}
	token = out.AccessToken
	return
}

func (c *MPTokenClient) RefreshToken(currentToken string) (token string, err error) {
	in := RefreshCommand{CurrentToken: currentToken}
	out := RefreshResp{}
	if err = c.client.Query("MPTokenService.Refresh", &in, &out).Do(context.Background()); err != nil {
		return
	}
	token = out.AccessToken
	return
}

func (c *MPTokenClient) IID01332E16DF5011E5A9D5A4DB30FED8E1() {
}
