package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend/driver"
	"github.com/backend/model"
)

type attendanceRepository struct {
	conn *sql.DB
}

func NewAttendanceRepository(conn *sql.DB) *attendanceRepository {
	return &attendanceRepository{conn: conn}
}

func (user1 *attendanceRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	obj := new(model.Attendance)
	return driver.GetById(user1.conn, obj, id)
}

func (user1 *attendanceRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Attendance)
	result, err := driver.Create(user1.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user1 *attendanceRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object in repository/user1 module")
	usr := obj.(model.Attendance)

	err := driver.UpdateById(user1.conn, &usr)
	return obj, err
}

func (user1 *attendanceRepository) Delete(cntx context.Context, id int64) (interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Attendance{Id: id}
	return driver.DeleteById(user1.conn, obj, id)
}

func (user1 *attendanceRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.User1 object repository/user1 module")
	obj := &model.Attendance{}
	return driver.GetAll(user1.conn, obj, 0, 0)
}
