package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chenminjian/go-arch/service/user"
	"github.com/gin-gonic/gin"
)

type Api struct {
	router  *gin.Engine
	userSrv usersrv.Service
}

func New(userSrv usersrv.Service) *Api {
	router := gin.New()
	router.Use(gin.Recovery())

	api := &Api{
		router:  router,
		userSrv: userSrv,
	}
	api.init()

	return api
}

func (api *Api) init() {
	api.router.POST("/user/add", api.UserAdd)
	api.router.GET("/user/detail", api.UserDetail)
	api.router.DELETE("/user/remove", api.UserRemove)
	api.router.GET("/user/list", api.UserList)
}

func (api *Api) Start() error {
	hs := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8001),
		Handler:        api.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := hs.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
