package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"go.guoyk.net/trackid"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_id"
	"time"
)

var (
	optServiceIDAddr string
)

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "svc_id_test"
	exe.Setup()

	if err = env.StringVar(&optServiceIDAddr, "SERVICE_ID_ADDR", "svc-id:3000"); err != nil {
		return
	}

	c := nrpc.NewClient(nrpc.ClientOptions{})
	c.Register("IDService", optServiceIDAddr)

	ic := svc_id.NewClient(c)

	ctx := trackid.Set(context.Background(), "111111111111")

	for {
		var ids []string
		if ids, err = ic.NextStr(ctx, 10); err != nil {
			return
		} else {
			log.Info().Strs("ids", ids).Msg("invoked")
		}

		time.Sleep(time.Second * 10)
	}
}
