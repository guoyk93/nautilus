package opts_redis

import "go.guoyk.net/env"

type Options struct {
	Host string
	Port int64
	Pass string
}

func Load(opts *Options) (err error) {
	if err = env.StringVar(&opts.Host, "REDIS_HOST", "127.0.0.1"); err != nil {
		return
	}
	if err = env.Int64Var(&opts.Port, "REDIS_PORT", 6379); err != nil {
		return
	}
	if err = env.StringVar(&opts.Pass, "REDIS_PASS", ""); err != nil {
		return
	}
	return
}
