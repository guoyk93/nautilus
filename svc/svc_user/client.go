package svc_user

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

func (c *Client) GetByMPOpenID(ctx context.Context, openID string) (ret GetResp, err error) {
	err = c.client.Query("UserService.Get", &GetQuery{MPOpenID: openID}, &ret).Do(ctx)
	return
}
