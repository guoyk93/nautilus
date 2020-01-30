package exe

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/nrpc/v2"
	"os"
	"os/signal"
	"syscall"
)

func RunNRPCServer(s *nrpc.Server) (err error) {
	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)
	s.Start(chErr)
	defer s.Shutdown(context.Background())

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Info().Str("signal", sig.String()).Msg("signaled")
	}
	return
}
