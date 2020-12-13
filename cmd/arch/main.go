package main

import (
	"github.com/chenminjian/go-arch/api"
	"github.com/chenminjian/go-arch/dao/db"
	"github.com/chenminjian/go-arch/dao/user"
	"github.com/chenminjian/go-arch/service/user"
)

func main() {
	if err := execute(); err != nil {
		panic(err)
	}
}

func execute() error {
	db:= db.New(&db.Default)

	userDao := userdao.New(db)

	userSrv := usersrv.New(userDao)

	api := api.New(userSrv)

	if err := api.Start(); err != nil {
		return err
	}

	return nil
}
