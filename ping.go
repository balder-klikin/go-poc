package main

type Ping struct {
	Value string `json:"value"`
}

func (ping *Ping) valid() bool {
	return len(ping.Value) > 0
}
