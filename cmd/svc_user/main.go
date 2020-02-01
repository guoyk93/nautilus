package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.guoyk.net/env"
	"go.guoyk.net/nrpc/v2"
	"nautilus/opts/opts_mysql"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_user"
)

var (
	optBind          string
	optsMySQL        opts_mysql.Options
	optServiceIDAddr string
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":3000"); err != nil {
		return
	}
	if err = opts_mysql.Load(&optsMySQL, "CORE_USER_"); err != nil {
		return
	}
	if err = env.StringVar(&optServiceIDAddr, "SERVICE_ID_ADDR", "svc-id:3000"); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "svc-user"
	exe.Setup()
	if err = setup(); err != nil {
		return
	}

	var db *gorm.DB
	if db, err = gorm.Open("mysql", optsMySQL.DSN()); err != nil {
		return
	}
	defer db.Close()

	nc := nrpc.NewClient(nrpc.ClientOptions{})
	nc.Register("IDService", optServiceIDAddr)

	us := svc_user.NewService(svc_user.ServiceOptions{
		DB:     db,
		Client: nc,
	})

	s := nrpc.NewServer(nrpc.ServerOptions{Addr: optBind})
	s.Register(us)

	err = exe.RunNRPCServer(s)
}
