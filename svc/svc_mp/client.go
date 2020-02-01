package svc_mp

import (
	"context"
	"go.guoyk.net/nrpc/v2"
)

type Client struct {
	client *nrpc.Client
}

func NewClient(client *nrpc.Client) *Client {
	return &Client{client: client}
}

func (c *Client) GetUserInfo(ctx context.Context, openID string) (ret GetUserInfoResp, err error) {
	err = c.client.Query("MPService.GetUserInfo", &GetUserInfoQuery{OpenID: openID}, &ret).Do(ctx)
	return
}
