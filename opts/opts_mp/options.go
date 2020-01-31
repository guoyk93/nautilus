package opts_mp

import "go.guoyk.net/env"

type Options struct {
	Org       string
	AppID     string
	AppSecret string
	AppToken  string
	AppAESKey string
}

func Load() (opts Options, err error) {
	if err = env.StringVar(&opts.Org, "MP_ORG", ""); err != nil {
		return
	}
	if err = env.StringVar(&opts.AppID, "MP_APP_ID", ""); err != nil {
		return
	}
	if err = env.StringVar(&opts.AppSecret, "MP_APP_SECRET", ""); err != nil {
		return
	}
	if err = env.StringVar(&opts.AppToken, "MP_APP_TOKEN", ""); err != nil {
		return
	}
	if err = env.StringVar(&opts.AppAESKey, "MP_APP_AES_KEY", ""); err != nil {
		return
	}
	return
}
