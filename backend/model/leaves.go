package model

type Leaves struct {
	Id       int64  `json:"id,omitempty" key:"primary"  column:"id"`
	EmpID    int64  `json:"emp_id" column:"emp_id"`
	FromDate string `json:"from_date" column:"from_date"`
	ToDate   string `json:"to_date" column:"to_date"`
}

func (ntc *Leaves) Table() string {
	return "leaves"
}

func (ntc *Leaves) String() string {
	return Stringify(ntc)
}
