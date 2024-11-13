package model

type Allow struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AllowAction struct {
	Goal    string `json:"goal"`
	Content string `json:"content"`
}
