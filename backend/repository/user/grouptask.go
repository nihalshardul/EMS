package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type gruptskRepository struct {
	conn *sql.DB
}

func NewGroupTaskRepository(conn *sql.DB) *gruptskRepository {
	return &gruptskRepository{conn: conn}
}

func (usergroup *gruptskRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	obj := new(model.GroupTask)
	return driver.GetById(usergroup.conn, obj, id)
}

func (usergroup *gruptskRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	usr := obj.(model.GroupTask)
	result, err := driver.Create(usergroup.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (usergroup *gruptskRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object in repository/usergroup module")
	usr := obj.(model.GroupTask)

	err := driver.UpdateById(usergroup.conn, &usr)
	return obj, err
}

func (usergroup *gruptskRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object repository/usergroup module")
	obj := &model.GroupTask{Id: id}
	return driver.DeleteById(usergroup.conn, obj, id)
}

func (usergroup *gruptskRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.Usergroup object repository/usergroup module")
	obj := &model.GroupTask{}
	return driver.GetAll(usergroup.conn, obj, 0, 0)
}
