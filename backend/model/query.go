package model

type Query struct {
	Id      int64  `json:"id,omitempty" key:"primary"  column:"id"`
	EmpId   int64  `json:"emp_id"  column:"emp_id"`
	Comment string `json:"comment" column:"comment"`
}

func (ntc *Query) Table() string {
	return "query"
}

func (ntc *Query) String() string {
	return Stringify(ntc)
}
