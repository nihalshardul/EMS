package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type anstoqueryRepository struct {
	conn *sql.DB
}

func NewAnsToQueryRepository(conn *sql.DB) *anstoqueryRepository {
	return &anstoqueryRepository{conn: conn}
}

func (user1 *anstoqueryRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.AnsToQuery)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *anstoqueryRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.AnsToQuery)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *anstoqueryRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.AnsToQuery)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *anstoqueryRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.AnsToQuery{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *anstoqueryRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.AnsToQuery{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
