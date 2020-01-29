package svc_id

import (
	"context"
	"errors"
	"go.guoyk.net/nrpc/v2"
)

type Client struct {
	client *nrpc.Client
}

func NewClient(c *nrpc.Client) *Client {
	return &Client{client: c}
}

func (c *Client) NewID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = c.NewIDs(ctx, 1); err != nil {
		return
	}
	id = ids[0]
	return
}

func (c *Client) NewIDs(ctx context.Context, size int) (ids []string, err error) {
	in := &NewIDQuery{Size: size}
	out := &NewIDResp{}

	if err = c.client.Query("IDService", "NewID").In(in).Out(out).Do(ctx); err != nil {
		return
	}
	if len(out.IDs) != size {
		err = errors.New("invalid number of IDService replies")
		return
	}
	ids = out.IDs
	return
}
