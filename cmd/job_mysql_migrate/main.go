package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"nautilus/opts/opts_mysql"
	"nautilus/pkg/exe"
	"nautilus/svc/svc_user"
)

var (
	optsCoreUserMySQL opts_mysql.Options
)

func setup() (err error) {
	if err = opts_mysql.Load(&optsCoreUserMySQL, "CORE_USER_"); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "job-mysql-migrate"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	var dbCoreUser *gorm.DB

	if dbCoreUser, err = gorm.Open("mysql", optsCoreUserMySQL.DSN()); err != nil {
		return
	}

	dbCoreUser.LogMode(true)

	if err = dbCoreUser.AutoMigrate(&svc_user.User{}, &svc_user.Auth{}).Error; err != nil {
		return
	}
}
