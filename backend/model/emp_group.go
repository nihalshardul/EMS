package model

type EmpGroup struct {
	Id    int64 `json:"id,omitempty"  column:"id"`
	GId   int64 `json:"gid"   column:"gid"`
	EmpId int64 `json:"emp_id"   column:"emp_id"`
}

func (empgroup *EmpGroup) Table() string {
	return "emp_group"
}

func (empgroup *EmpGroup) String() string {
	return Stringify(empgroup)
}
