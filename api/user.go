package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chenminjian/go-arch/model/entity"
	"github.com/gin-gonic/gin"
)

func (api *Api) UserDetail(c *gin.Context) {
	execute := func() (*entity.User, error) {
		idStr := c.Query("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}

		user, err := api.userSrv.GetByID(id)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	user, err := execute()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    user,
	})
}

func (api *Api) UserAdd(c *gin.Context) {
	execute := func() error {
		var user entity.User
		err := c.BindJSON(&user)
		if err != nil {
			return err
		}

		if user.Username == "" {
			return errors.New("username is empty")
		}

		err = api.userSrv.Add(&entity.User{Username:user.Username})
		if err != nil {
			return err
		}

		return nil
	}

	err := execute()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
	})
}

func (api *Api) UserRemove(c *gin.Context) {
	execute := func() error {
		idStr := c.Query("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		err = api.userSrv.Remove(id)
		if err != nil {
			return err
		}

		return nil
	}

	err := execute()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
	})
}