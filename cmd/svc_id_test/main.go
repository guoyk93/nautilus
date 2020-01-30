package main

import (
	"context"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"go.guoyk.net/trackid"
	"log"
	"nautilus/svc/svc_id"
	"time"
)

var (
	optServiceIDAddr string
)

func main() {
	if err := env.StringVar(&optServiceIDAddr, "SERVICE_ID_ADDR", "svc-id:3000"); err != nil {
		panic(err)
	}

	c := nrpc.NewClient(nrpc.ClientOptions{})
	c.Register("IDService", optServiceIDAddr)

	ic := svc_id.NewClient(c)

	ctx := trackid.Set(context.Background(), "111111111111")

	for {
		if id, err := ic.NewID(ctx); err != nil {
			panic(err)
		} else {
			log.Printf("NewID: %s", id)
		}

		if ids, err := ic.NewIDs(ctx, 10); err != nil {
			panic(err)
		} else {
			log.Printf("NewIDs: %v", ids)
		}

		time.Sleep(time.Second * 10)
	}
}
