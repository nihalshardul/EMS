package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type user1Repository struct {
	conn *sql.DB
}

func NewUser1Repository(conn *sql.DB) *user1Repository {
	return &user1Repository{conn: conn}
}

func (user1 *user1Repository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Emp)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *user1Repository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Emp)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *user1Repository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Emp)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *user1Repository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Emp{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *user1Repository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Emp{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
