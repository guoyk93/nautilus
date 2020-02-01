package svc_mp_token

import (
	"context"
	"github.com/chanxuehong/wechat/mp/core"
	"go.guoyk.net/nrpc/v2"
)

type Client struct {
	client *nrpc.Client
}

func NewClient(client *nrpc.Client) *Client {
	return &Client{client: client}
}

func (c *Client) Get(ctx context.Context) (ret string, err error) {
	out := GetResp{}
	if err = c.client.Query("MPTokenService.Get", nil, &out).Do(ctx); err != nil {
		return
	}
	ret = out.AccessToken
	return
}

func (c *Client) Refresh(ctx context.Context, current string) (ret string, err error) {
	in := RefreshCommand{CurrentToken: current}
	out := RefreshResp{}
	if err = c.client.Command("MPTokenService.Refresh", &in, &out).Do(ctx); err != nil {
		return
	}
	ret = out.AccessToken
	return
}

func (c *Client) MPAccessTokenServer() core.AccessTokenServer {
	return &accessTokenServer{client: c}
}

type accessTokenServer struct {
	client *Client
}

func (a *accessTokenServer) Token() (string, error) {
	return a.client.Get(context.Background())
}

func (a *accessTokenServer) RefreshToken(currentToken string) (string, error) {
	return a.client.Refresh(context.Background(), currentToken)
}

func (a *accessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {
}
