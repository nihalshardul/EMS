package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type noticeRepository struct {
	conn *sql.DB
}

func NewNoticeRepository(conn *sql.DB) *noticeRepository {
	return &noticeRepository{conn: conn}
}

func (user1 *noticeRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Notice)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *noticeRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Notice)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *noticeRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Notice)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *noticeRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Notice{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *noticeRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Notice{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
