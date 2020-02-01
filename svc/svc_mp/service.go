package svc_mp

import (
	"context"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/go-redis/redis"
	"nautilus/pkg/cache"
	"time"
)

type GetUserInfoQuery struct {
	OpenID string `json:"open_id" query:"open_id"`
}

type GetUserInfoResp struct {
	OpenID     string `json:"open_id"`
	Nickname   string `json:"nickname"`
	Subscribed bool   `json:"subscribed"`
}

type ServiceOptions struct {
	MP    *core.Client
	Redis *redis.Client
}

type MPService struct {
	mp    *core.Client
	redis *redis.Client
	cache cache.Cache
}

func NewService(opts ServiceOptions) *MPService {
	return &MPService{mp: opts.MP, redis: opts.Redis, cache: cache.New(opts.Redis)}
}

func (s *MPService) GetUserInfo(ctx context.Context, q *GetUserInfoQuery) (ret GetUserInfoResp, err error) {
	info := &user.UserInfo{}

	cacheKey := fmt.Sprintf("mp.user.oid.%s", q.OpenID)

	var found bool
	if found, err = s.cache.Get(cacheKey, info); err != nil {
		return
	}

	if !found {
		if info, err = user.Get(s.mp, q.OpenID, "zh_CN"); err != nil {
			return
		}
		if err = s.cache.Set(cacheKey, info, time.Minute*10); err != nil {
			return
		}
	}

	ret.OpenID = q.OpenID
	ret.Nickname = info.Nickname
	ret.Subscribed = info.IsSubscriber == 1

	return
}
