package userdao

import (
	"github.com/chenminjian/go-arch/dao"
	"github.com/chenminjian/go-arch/model/entity"
)

type Dao interface {
	dao.Dao

	GetByID(ID int64) (*entity.User, error)

	Add(user *entity.User) error
}
