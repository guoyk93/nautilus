package trackid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type (
	trackIdKeyType int

	trackIdType struct {
		trackId string
	}
)

const trackIdKey trackIdKeyType = 0

func Set(ctx context.Context, trackId string) context.Context {
	if len(trackId) == 0 {
		var buf [8]byte
		_, _ = rand.Read(buf[:])
		trackId = hex.EncodeToString(buf[:])
	}
	tid, ok := ctx.Value(trackIdKey).(*trackIdType)
	if ok {
		tid.trackId = trackId
		return ctx
	}
	tid = &trackIdType{trackId: trackId}
	ctx = context.WithValue(ctx, trackIdKey, tid)
	return ctx
}

func Get(ctx context.Context) string {
	tid, ok := ctx.Value(trackIdKey).(*trackIdType)
	if ok {
		return tid.trackId
	}
	return ""
}

func Clear(ctx context.Context) {
	tid, ok := ctx.Value(trackIdKey).(*trackIdType)
	if ok {
		tid.trackId = ""
	}
	return
}
