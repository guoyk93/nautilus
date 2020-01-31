package main

import (
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"nautilus/pkg/exe"
	"nautilus/pkg/foxtrot"
	"nautilus/svc/svc_id"
	"nautilus/web/web_main"
)

var (
	optBind           string
	optAssetDir       string
	optServiceIDAddr  string
	optReloadTemplate bool
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":4000"); err != nil {
		return
	}
	if err = env.StringVar(&optAssetDir, "ASSET_DIR", "/assets/web_main"); err != nil {
		return
	}
	if err = env.StringVar(&optServiceIDAddr, "SERVICE_ID_ADDR", "svc-id:3000"); err != nil {
		return
	}
	if err = env.BoolVar(&optReloadTemplate, "RELOAD_TEMPLATE", false); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "web_main"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	f := foxtrot.New(foxtrot.Options{
		Addr:           optBind,
		AssetDir:       optAssetDir,
		ReloadTemplate: optReloadTemplate,
	})

	c := nrpc.NewClient(nrpc.ClientOptions{})
	c.Register("IDService", optServiceIDAddr)

	w := &web_main.Web{
		ClientID: svc_id.NewClient(c),
	}

	f.GET("/", w.Index)

	err = exe.RunFoxtrot(f)
}
