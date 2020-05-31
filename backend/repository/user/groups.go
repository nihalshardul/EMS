package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type groupsRepository struct {
	conn *sql.DB
}

func NewGroupsRepository(conn *sql.DB) *groupsRepository {
	return &groupsRepository{conn: conn}
}

func (groups *groupsRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.Groups object in repository/groups module")
	obj := new(model.Groups)
	return driver.GetById(groups.conn, obj, id)
}

func (groups *groupsRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Groups object in repository/groups module")
	usr := obj.(model.Groups)
	result, err := driver.Create(groups.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (groups *groupsRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Groups object in repository/groups module")
	usr := obj.(model.Groups)

	err := driver.UpdateById(groups.conn, &usr)
	return obj, err
}

func (groups *groupsRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.Groups object repository/groups module")
	obj := &model.Groups{Id: id}
	return driver.DeleteById(groups.conn, obj, id)
}

func (groups *groupsRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.Groups object repository/groups module")
	obj := &model.Groups{}
	return driver.GetAll(groups.conn, obj, 0, 0)
}
