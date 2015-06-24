package app

type Ping struct {
	Value string `json:"value"`
}

func (ping *Ping) Valid() bool {
	return len(ping.Value) > 0
}
