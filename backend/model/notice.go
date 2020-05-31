package model

type Notice struct {
	Id   int64  `json:"id,omitempty" key:"primary"  column:"id"`
	Info string `json:"info" column:"info"`
}

func (ntc *Notice) Table() string {
	return "notice"
}

func (ntc *Notice) String() string {
	return Stringify(ntc)
}
