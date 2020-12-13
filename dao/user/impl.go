package userdao

import (
	"github.com/chenminjian/go-arch/model/entity"
	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) Dao {
	return &impl{
		db: db,
	}
}

func (im *impl) TableName() string {
	return "user"
}

func (im *impl) GetByID(id int64) (*entity.User, error) {
	var user entity.User
	if err := im.db.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (im *impl) Add(user *entity.User) error {
	if err := im.db.Table("user").Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (im *impl) Remove(id int64) error {
	if err := im.db.Table("user").Delete(&entity.User{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func (im *impl) List() ([]*entity.User, error) {
	items := make([]*entity.User, 0)
	if err := im.db.Table("user").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
