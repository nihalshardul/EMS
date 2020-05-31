package model

type Attendance struct {
	Id          int64  `json:"id,omitempty" key:"primary"  column:"id"`
	EmpID       int64  `json:"emp_id" column:"emp_id"`
	Attend_date string `json:"attend_date" column:"attend_date"`
}

func (ntc *Attendance) Table() string {
	return "attendance"
}

func (ntc *Attendance) String() string {
	return Stringify(ntc)
}
