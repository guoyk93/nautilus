package opts_mysql

import (
	"fmt"
	"go.guoyk.net/env"
)

type Options struct {
	Host string
	Port int64
	User string
	Pass string
	DB   string
}

func Load(opts *Options, pfx string) (err error) {
	if err = env.StringVar(&opts.Host, pfx+"MYSQL_HOST", "127.0.0.1"); err != nil {
		return
	}
	if err = env.Int64Var(&opts.Port, pfx+"MYSQL_PORT", 3306); err != nil {
		return
	}
	if err = env.StringVar(&opts.User, pfx+"MYSQL_USER", "root"); err != nil {
		return
	}
	if err = env.StringVar(&opts.Pass, pfx+"MYSQL_PASS", ""); err != nil {
		return
	}
	if err = env.StringVar(&opts.DB, pfx+"MYSQL_DB", ""); err != nil {
		return
	}
	return
}

func (opts Options) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opts.User,
		opts.Pass,
		opts.Host,
		opts.Port,
		opts.DB,
	)
}
