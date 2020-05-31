package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type leavesRepository struct {
	conn *sql.DB
}

func NewLeavesRepository(conn *sql.DB) *leavesRepository {
	return &leavesRepository{conn: conn}
}

func (user1 *leavesRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Leaves)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *leavesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Leaves)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *leavesRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Leaves)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *leavesRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Leaves{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *leavesRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Leaves{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
