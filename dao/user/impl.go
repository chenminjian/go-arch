package userdao

import (
	"fmt"

	"github.com/chenminjian/go-arch/dao/db"
	"github.com/chenminjian/go-arch/model/entity"
)

type impl struct {
	db *db.DB
}

func New(db *db.DB) Dao {
	return &impl{
		db: db,
	}
}

func (im *impl) TableName() string {
	return "user"
}

func (im *impl) GetByID(ID int64) (*entity.User, error) {

	var user entity.User

	sqlstr := fmt.Sprintf("SELECT id, username FROM %s WHERE id = ?", im.TableName())
	stmt, err := im.db.Prepare(sqlstr)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(ID).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
