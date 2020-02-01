package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"nautilus/opts/opts_mp"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_user"
	"net/http"
)

var (
	optBind            string
	optsMP             opts_mp.Options
	optServiceUserAddr string
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":4000"); err != nil {
		return
	}
	if err = opts_mp.Load(&optsMP); err != nil {
		return
	}
	if err = env.StringVar(&optServiceUserAddr, "SERVICE_USER_ADDR", "svc-user:3000"); err != nil {
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

	nc := nrpc.NewClient(nrpc.ClientOptions{})
	nc.Register("UserService", optServiceUserAddr)

	uc := svc_user.NewClient(nc)

	mmux := core.NewServeMux()
	mmux.DefaultEventHandleFunc(func(c *core.Context) {
		_ = c.NoneResponse()
	})
	mmux.DefaultMsgHandleFunc(func(c *core.Context) {
		_ = c.NoneResponse()
	})
	mmux.MsgHandleFunc(request.MsgTypeText, func(c *core.Context) {
		msg := request.GetText(c.MixedMsg)
		u, err := uc.GetByMPOpenID(c.Request.Context(), msg.FromUserName)
		if err != nil {
			resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, "Err: "+err.Error())
			_ = c.AESResponse(resp, 0, "", nil)
		} else {
			resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, fmt.Sprintf("来自 %s:\n%s", u.Nickname, msg.Content))
			_ = c.AESResponse(resp, 0, "", nil)
		}
		log.Info().Str("from", msg.FromUserName).Str("text", msg.Content).Msg("received")
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
