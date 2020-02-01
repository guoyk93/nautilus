package svc_mp_token

import (
	"context"
	"github.com/chanxuehong/wechat/mp/core"
)

type GetResp struct {
	AccessToken string `json:"access_token"`
}

type RefreshCommand struct {
	CurrentToken string `json:"current_token"`
}

type RefreshResp struct {
	AccessToken string `json:"access_token"`
}

type ServiceOptions struct {
	AccessTokenServer core.AccessTokenServer
}

type MPTokenService struct {
	ats core.AccessTokenServer
}

func NewService(opts ServiceOptions) *MPTokenService {
	return &MPTokenService{ats: opts.AccessTokenServer}
}

func (s *MPTokenService) HealthCheck(ctx context.Context) (err error) {
	_, err = s.Get(ctx)
	return
}

func (s *MPTokenService) Get(ctx context.Context) (out GetResp, err error) {
	out.AccessToken, err = s.ats.Token()
	return
}

func (s *MPTokenService) Refresh(ctx context.Context, in *RefreshCommand) (out RefreshResp, err error) {
	out.AccessToken, err = s.ats.RefreshToken(in.CurrentToken)
	return
}
