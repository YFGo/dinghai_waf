package types

import "time"

type NormalHttpInfo struct {
	ID            string    `json:"id"`
	IP            string    `json:"ip"`
	RequestURI    string    `json:"request_uri"`
	RequestTime   time.Time `json:"request_time"`
	RequestMethod string    `json:"request_method"`
	Protocol      string    `json:"protocol"`
	RequestBody   string    `json:"request_body"`
}
