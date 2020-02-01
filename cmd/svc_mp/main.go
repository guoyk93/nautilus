package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/go-redis/redis"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"nautilus/opts/opts_redis"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_mp"
	"nautilus/svc/svc_mp_token"
)

var (
	optBind               string
	optServiceMPTokenAddr string
	optsRedis             opts_redis.Options
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":3000"); err != nil {
		return
	}
	if err = env.StringVar(&optServiceMPTokenAddr, "SERVICE_MP_TOKEN_ADDR", "svc-mp-token:3000"); err != nil {
		return
	}
	if err = opts_redis.Load(&optsRedis); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "svc-mp"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	nc := nrpc.NewClient(nrpc.ClientOptions{})
	nc.Register("MPTokenService", optServiceMPTokenAddr)

	tc := svc_mp_token.NewClient(nc).MPAccessTokenServer()
	mc := core.NewClient(tc, nil)

	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", optsRedis.Host, optsRedis.Port),
		Password: optsRedis.Pass,
	})

	ms := svc_mp.NewService(svc_mp.ServiceOptions{MP: mc, Redis: rc})

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: optBind})
	s.Register(ms)
	err = exe.RunNRPCServer(s)
}
