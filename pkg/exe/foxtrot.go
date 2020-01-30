package exe

import (
	"context"
	"github.com/rs/zerolog/log"
	"nautilus/pkg/foxtrot"
	"os"
	"os/signal"
	"syscall"
)

func RunFoxtrot(f *foxtrot.Foxtrot) (err error) {
	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)
	f.Start(chErr)
	defer f.Shutdown(context.Background())

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Info().Str("signal", sig.String()).Msg("signaled")
	}
	return
}
