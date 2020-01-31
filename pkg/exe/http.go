package exe

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunHTTPServer(s *http.Server) (err error) {
	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		chErr <- s.ListenAndServe()
	}()
	defer s.Shutdown(context.Background())

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Info().Str("signal", sig.String()).Msg("signaled")
	}
	return
}
