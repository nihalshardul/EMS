package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type loginRepository struct {
	conn *sql.DB
}

func NewLoginRepository(conn *sql.DB) *loginRepository {
	return &loginRepository{conn: conn}
}

func (user1 *loginRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Login)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *loginRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Login)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *loginRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Login)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *loginRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Login{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *loginRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Login{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
