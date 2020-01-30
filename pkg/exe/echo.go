package exe

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func RunEcho(e *echo.Echo, addr string) (err error) {
	e.HidePort = true
	e.HideBanner = true

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		chErr <- e.Start(addr)
	}()
	defer e.Shutdown(context.Background())
	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Info().Str("signal", sig.String()).Msg("signaled")
	}
	return
}
