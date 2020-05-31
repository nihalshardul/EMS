package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type taskRepository struct {
	conn *sql.DB
}

func NewTaskRepository(conn *sql.DB) *taskRepository {
	return &taskRepository{conn: conn}
}

func (user1 *taskRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Task)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *taskRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Task)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *taskRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Task)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *taskRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Task{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *taskRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Task{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
