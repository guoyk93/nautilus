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
	if optsMP, err = opts_mp.Load(); err != nil {
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

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: optBind})
	s.Register(svc_mp_token.NewService(ats))
	err = exe.RunNRPCServer(s)
}
