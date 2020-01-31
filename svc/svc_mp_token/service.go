package svc_mp_token

import (
	"context"
	"github.com/chanxuehong/wechat/mp/core"
)

type MPTokenService struct {
	ats core.AccessTokenServer
}

func (s *MPTokenService) Get(ctx context.Context) (out GetResp, err error) {
	out.AccessToken, err = s.ats.Token()
	return
}

func (s *MPTokenService) Refresh(ctx context.Context, in *RefreshCommand) (out RefreshResp, err error) {
	out.AccessToken, err = s.ats.RefreshToken(in.CurrentToken)
	return
}

func NewService(ats core.AccessTokenServer) *MPTokenService {
	return &MPTokenService{ats: ats}
}
