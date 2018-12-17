package usersrv

import (
	"github.com/chenminjian/go-arch/dao/user"
	"github.com/chenminjian/go-arch/model/entity"
)

type impl struct {
	UserDao userdao.Dao
}

func New(userDao userdao.Dao) Service {
	return &impl{
		UserDao: userDao,
	}
}

func (im *impl) GetByID(ID int64) (*entity.User, error) {
	return im.UserDao.GetByID(ID)
}
