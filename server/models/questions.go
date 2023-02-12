package models

import "encoding/json"

type Question struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`

	Answers []Answer `json:"answers"`
}

func (q *Question) ToJson() string {
	var data, err = json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(data)
}
