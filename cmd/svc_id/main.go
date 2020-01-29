package main

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"go.guoyk.net/snowflake"
	"nautilus/pkg/in_k8s"
	"nautilus/svc/svc_id"
	"os"
	"os/signal"
	"syscall"
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

func exit(err *error) {
	if *err != nil {
		log.Error().Str("topic", "main").Str("service", "IDService").Err(*err).Msg("exited")
		os.Exit(1)
	}
}

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
	defer exit(&err)

	if err = setup(); err != nil {
		return
	}

	instanceID := optClusterID<<5 | optWorkerID

	svc := svc_id.NewIDService(snowflake.New(startTime, instanceID))

	log.Info().Str(
		"topic", "main",
	).Str(
		"service", "IDService",
	).Uint64(
		"cluster_id", optClusterID,
	).Uint64(
		"worker_id", optWorkerID,
	).Uint64(
		"instance_id", instanceID,
	).Msg("started")

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: ":3000"})
	s.Register(svc)
	s.Start(chErr)
	defer s.Shutdown(context.Background())

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Info().Str("topic", "main").Str("service", "IDService").Str("signal", sig.String()).Msg("signaled")
	}
}
