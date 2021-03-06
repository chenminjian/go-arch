package usersrv

import "github.com/chenminjian/go-arch/model/entity"

type Service interface {
	GetByID(ID int64) (*entity.User, error)

	Add(user *entity.User) error

	Remove(id int64) error

	List()([]*entity.User, error)
}
