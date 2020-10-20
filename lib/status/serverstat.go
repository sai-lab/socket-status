package status

import "encoding/json"

type ServerStat struct {
	// Host
	HostName             string `json:"hostname"`
	HostID               string `json:"hostid"`
	VirtualizationSystem string `json:"virtualizationSystem"`
	// Socket
	Socket int `json:"socket"`
	// Time
	Time string `json:"time"`
	// Error
	ErrorInfo []error `json:"errorInfo"`
}

func (d ServerStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
