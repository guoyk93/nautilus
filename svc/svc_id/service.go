package svc_id

import (
	"context"
	"go.guoyk.net/snowflake"
	"strconv"
)

type NextQuery struct {
	Size int `json:"size" query:"size" default:"1"`
}

type NextResp struct {
	IDs []string `json:"ids"`
}

type ServiceOptions struct {
	Snowflake snowflake.Snowflake
}

type IDService struct {
	sf snowflake.Snowflake
}

func NewService(opts ServiceOptions) *IDService {
	return &IDService{
		sf: opts.Snowflake,
	}
}

func (s *IDService) Next(ctx context.Context, req *NextQuery) (resp NextResp, err error) {
	for i := 0; i < req.Size; i++ {
		resp.IDs = append(resp.IDs, strconv.FormatUint(s.sf.NewID(), 10))
	}
	return
}

// TODO: deprecated
func (s *IDService) NewID(ctx context.Context, req *NextQuery) (NextResp, error) {
	return s.Next(ctx, req)
}
