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
	in, out := &NewIDQuery{Size: size}, &NewIDResp{}

	if err = c.client.Query("IDService.NewID", in, out).Do(ctx); err != nil {
		return
	}
	if len(out.IDs) != size {
		err = errors.New("invalid number of IDService.NewID replies")
		return
	}
	ids = out.IDs
	return
}

func (c *Client) NewIDInt64(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = c.NewIDsInt64(ctx, 1); err != nil {
		return
	}
	id = ids[0]
	return
}

func (c *Client) NewIDsInt64(ctx context.Context, size int) (ids []int64, err error) {
	var idsStr []string
	if idsStr, err = c.NewIDs(ctx, size); err != nil {
		return
	}
	for _, idStr := range idsStr {
		var id int64
		if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
			return
		}
		ids = append(ids, id)
	}
	return
}
