package model

type Emp struct {
	Id          int64  `json:"id,omitempty" key:"primary"  column:"id"`
	FirstName   string `json:"first_name" column:"first_name"`
	LastName    string `json:"last_name" column:"last_name"`
	Designation string `json:"designation" column:"designation"`
	Email       string `json:"email" column:"email"`
	Password    string `json:"password" column:"password"`
	Salary      int    `json:"salary" column:"salary"`
	Lvs         int    `json:"lvs" column:"lvs"`
	IsEmpRoot   int    `json:"is_emp_root" column:"is_emp_root"`
	UpdatedBy   int64  `json:"updated_by" column:"updated_by"`
}

func (user1 *Emp) Table() string {
	return "emp"
}

func (user1 *Emp) String() string {
	return Stringify(user1)
}
