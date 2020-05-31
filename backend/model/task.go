package model

type Task struct {
	Id        int64  `json:"id,omitempty" key:"primary"  column:"id"`
	TaskName  string `json:"task_name" column:"task_name"`
	Comment   string `json:"comment" column:"comment"`
	Pending   int    `json:"pending" column:"pending"`
	Complete  int    `json:"complete" column:"complete"`
	UpdatedBy int64  `json:"updated_by" column:"updated_by"`
}

func (resources *Task) Table() string {
	return "task"
}

func (resources *Task) String() string {
	return Stringify(resources)
}
