package svc_user

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"go.guoyk.net/nrpc/v2"
	"nautilus/svc/svc_id"
	"nautilus/svc/svc_mp"
)

type GetQuery struct {
	MPOpenID string `json:"mp_open_id"`
}

type GetResp struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

type ServiceOptions struct {
	DB     *gorm.DB
	Client *nrpc.Client
}

type UserService struct {
	db       *gorm.DB
	client   *nrpc.Client
	idClient *svc_id.Client
}

func NewService(opts ServiceOptions) *UserService {
	return &UserService{
		db:       opts.DB,
		client:   opts.Client,
		idClient: svc_id.NewClient(opts.Client),
	}
}

func (s *UserService) Get(ctx context.Context, req *GetQuery) (out GetResp, err error) {
	if len(req.MPOpenID) == 0 {
		err = nrpc.Solid(errors.New("missing argument mp_open_id"))
		return
	}
	var mpinfo svc_mp.GetUserInfoResp
	if err = s.client.Query(
		"MPService.GetUserInfo",
		&svc_mp.GetUserInfoQuery{OpenID: req.MPOpenID},
		&mpinfo,
	).Do(ctx); err != nil {
		return
	}
	return
}
