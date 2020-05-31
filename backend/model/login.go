package model

type Login struct {
	Id        int64  `json:"id,omitempty" key:"primary"  column:"id"`
	EmpID     int64  `json:"emp_id" column:"emp_id"`
	LoginDate string `json:"login_date" column:"login_date"`
}

func (ntc *Login) Table() string {
	return "login"
}

func (ntc *Login) String() string {
	return Stringify(ntc)
}
