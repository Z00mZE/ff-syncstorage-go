package storage

type BSO struct {
	Id        string `json:"id"`
	Sortindex int    `json:"sortindex"`
	Payload   string `json:"payload"`
	Ttl       uint64 `json:"ttl"`
}
type Request = []BSO
