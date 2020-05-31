package model

type AnsToQuery struct {
	Id      int64  `json:"id,omitempty" key:"primary"  column:"id"`
	QId     int64  `json:"q_id"  column:"q_id"`
	Comment string `json:"comment" column:"comment"`
}

func (ntc *AnsToQuery) Table() string {
	return "ans_to_query"
}

func (ntc *AnsToQuery) String() string {
	return Stringify(ntc)
}
