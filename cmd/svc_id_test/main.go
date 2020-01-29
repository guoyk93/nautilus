package main

import (
	"context"
	"go.guoyk.net/nrpc/v2"
	"go.guoyk.net/trackid"
	"log"
	"nautilus/svc/svc_id"
)

func main() {
	c := nrpc.NewClient(nrpc.ClientOptions{})
	c.Register("IDService", "127.0.0.1:3000")

	ic := svc_id.NewClient(c)

	ctx := trackid.Set(context.Background(), "111222333")

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
}
