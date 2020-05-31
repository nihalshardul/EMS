package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type queryRepository struct {
	conn *sql.DB
}

func NewQueryRepository(conn *sql.DB) *queryRepository {
	return &queryRepository{conn: conn}
}

func (user1 *queryRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Query)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *queryRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Query)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *queryRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Query)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *queryRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Query{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *queryRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Query{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
