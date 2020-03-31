package schema

type Point struct {
	UUID   string `json:"uuid"`
	Alias  string `json:"alias"`
	Tenant string `json:"tenant"`
	Server string `json:"server"`
	Uptime int64  `json:"uptime"`
	State  string `json:"state"`
	Device string `json:"device"`
	Switch string `json:"switch"`
}
