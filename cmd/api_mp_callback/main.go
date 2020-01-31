package main

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"go.guoyk.net/env"
	"nautilus/opts/opts_mp"
	"nautilus/pkg/exe"
	"net/http"
)

var (
	optBind string
	optsMP  opts_mp.Options
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":4000"); err != nil {
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

	exe.Project = "api-mp-callback"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	mmux := core.NewServeMux()
	mmux.DefaultEventHandleFunc(func(c *core.Context) {
		_ = c.NoneResponse()
	})
	mmux.DefaultMsgHandleFunc(func(c *core.Context) {
		_ = c.NoneResponse()
	})
	mmux.MsgHandleFunc(request.MsgTypeText, func(c *core.Context) {
		msg := request.GetText(c.MixedMsg)
		resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, "Re: "+msg.Content)
		_ = c.AESResponse(resp, 0, "", nil)
	})
	ms := core.NewServer(optsMP.Org, optsMP.AppID, optsMP.AppToken, optsMP.AppAESKey, mmux, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte("OK"))
	})
	mux.HandleFunc("/mp_callback/callback", func(rw http.ResponseWriter, req *http.Request) {
		ms.ServeHTTP(rw, req, nil)
	})

	s := &http.Server{Addr: optBind, Handler: mux}

	err = exe.RunHTTPServer(s)
}
