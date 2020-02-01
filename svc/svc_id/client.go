package svc_id

import (
	"context"
	"errors"
	"go.guoyk.net/nrpc/v2"
	"strconv"
)

type Client struct {
	client *nrpc.Client
}

func NewClient(client *nrpc.Client) *Client {
	return &Client{client: client}
}

func (c *Client) Next(ctx context.Context, size int) (ret []int64, err error) {
	var idStrs []string
	if idStrs, err = c.NextStr(ctx, size); err != nil {
		return
	}
	for _, idStr := range idStrs {
		var id int64
		if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
			return
		}
		ret = append(ret, id)
	}
	return
}

func (c *Client) NextStr(ctx context.Context, size int) (ret []string, err error) {
	in, out := &NextQuery{Size: size}, &NextResp{}

	if err = c.client.Query("IDService.Next", in, out).Do(ctx); err != nil {
		return
	}
	if len(out.IDs) != size {
		err = errors.New("invalid number of IDService.Next replies")
		return
	}
	ret = out.IDs
	return
}
