package main

import (
	"github.com/chanxuehong/wechat/mp/core"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"nautilus/opts/opts_mp"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_mp_token"
)

var (
	optBind string
	optsMP  opts_mp.Options
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":3000"); err != nil {
		return
	}
	if err = opts_mp.Load(&optsMP); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "svc-mp-token"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	ats := core.NewDefaultAccessTokenServer(optsMP.AppID, optsMP.AppSecret, nil)

	ts := svc_mp_token.NewService(svc_mp_token.ServiceOptions{AccessTokenServer: ats})

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: optBind})
	s.Register(ts)
	err = exe.RunNRPCServer(s)
}
