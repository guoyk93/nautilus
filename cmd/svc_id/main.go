package main

import (
	"errors"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"go.guoyk.net/snowflake"
	"nautilus/pkg/exe"
	"nautilus/pkg/in_k8s"
	"nautilus/svc/svc_id"
	"time"
)

var (
	optBind      string
	optClusterID uint64
	optWorkerID  uint64
)

var (
	startTime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
)

const (
	Uint5Mask = uint64(1<<5) - 1
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":3000"); err != nil {
		return
	}
	if err = env.Uint64Var(&optClusterID, "CLUSTER_ID", 0); err != nil {
		return
	}
	if err = env.Uint64Var(&optWorkerID, "WORKER_ID", 0); err != nil {
		return
	}

	if optClusterID == 0 {
		err = errors.New("missing env CLUSTER_ID")
		return
	}
	if optWorkerID == 0 {
		optWorkerID = in_k8s.GetStatefulSetSequenceID()
		if optWorkerID == 0 {
			err = errors.New("missing env WORKER_ID")
			return
		}
	}
	if optClusterID > Uint5Mask {
		err = errors.New("invalid env CLUSTER_ID")
		return
	}
	if optWorkerID > Uint5Mask {
		err = errors.New("invalid env WORKER_ID")
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "svc_id"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	instanceID := optClusterID<<5 | optWorkerID

	svc := svc_id.NewIDService(snowflake.New(startTime, instanceID))

	log.Info().Uint64("cluster_id", optClusterID).Uint64("worker_id", optWorkerID).Msg("started")

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: ":3000"})
	s.Register(svc)

	err = exe.RunNRPCServer(s)
}
