package model

type GroupTask struct {
	Id    int64 `json:"id,omitempty"  column:"id"`
	GId   int64 `json:"gid"   column:"gid"`
	TskId int64 `json:"tsk_id"   column:"tsk_id"`
}

func (grouptsk *GroupTask) Table() string {
	return "group_task"
}

func (grouptsk *GroupTask) String() string {
	return Stringify(grouptsk)
}
