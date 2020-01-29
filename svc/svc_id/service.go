package svc_id

import (
	"context"
	"go.guoyk.net/snowflake"
	"strconv"
)

type IDService struct {
	sf snowflake.Snowflake
}

func NewIDService(sf snowflake.Snowflake) *IDService {
	return &IDService{
		sf: sf,
	}
}

func (s *IDService) HealthCheck(ctx context.Context) error {
	return nil
}

func (s *IDService) NewID(ctx context.Context, req *NewIDQuery) (resp NewIDResp, err error) {
	for i := 0; i < req.Size; i++ {
		resp.IDs = append(resp.IDs, strconv.FormatUint(s.sf.NewID(), 10))
	}
	return
}
