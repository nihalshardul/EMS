package model

type Groupinfo struct {
	GroupId       int64  `json:"group_id,omitempty" key:"primary"  column:"groups_id"`
	GroupName    string `json:"group_name column:"group_name"`
	UpdatedBy     int64  `json:"updated_by" column:"updated_by"`
}

func (group *Groupinfo) Table() string {
	return "groups"
}

func (group *Groupinfo) String() string {
	return Stringify(group)
}
