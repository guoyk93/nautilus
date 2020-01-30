package exe

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.guoyk.net/env"
	"os"
)

var (
	Project     = "noname"
	Instance, _ = os.Hostname()
	Debug       bool
	Trace       bool
)

func Setup() {
	_ = env.BoolVar(&Debug, "DEBUG", false)
	_ = env.BoolVar(&Trace, "TRACE", false)
	level := zerolog.InfoLevel
	if Trace {
		level = zerolog.TraceLevel
	} else if Debug {
		level = zerolog.DebugLevel
	}
	log.Logger = zerolog.New(os.Stdout).Level(level).With().Str("project", Project).Str("instance", Instance).Logger()
}
