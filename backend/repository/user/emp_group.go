package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type empgroupRepository struct {
	conn *sql.DB
}

func NewEmpGroupRepository(conn *sql.DB) *empgroupRepository {
	return &empgroupRepository{conn: conn}
}

func (usergroup *empgroupRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	obj := new(model.EmpGroup)
	return driver.GetById(usergroup.conn, obj, id)
}

func (usergroup *empgroupRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	usr := obj.(model.EmpGroup)
	result, err := driver.Create(usergroup.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (usergroup *empgroupRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	usr := obj.(model.EmpGroup)

	err := driver.UpdateById(usergroup.conn, &usr)
	return obj, err
}

func (usergroup *empgroupRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object repository/usergroup module")
	obj := &model.EmpGroup{Id: id}
	return driver.DeleteById(usergroup.conn, obj, id)
}

func (usergroup *empgroupRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object repository/usergroup module")
	obj := &model.EmpGroup{}
	return driver.GetAll(usergroup.conn, obj, 0, 0)
}
