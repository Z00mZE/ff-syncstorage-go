package storage

type Response struct {
	Modified float64           `json:"modified"`
	Success  []string          `json:"success"`
	Failed   map[string]string `json:"failed"`
	Batch    string            `json:"batch"`
}
