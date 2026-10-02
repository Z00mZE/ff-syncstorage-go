package types

type Configuration struct {
	MaxPostBytes          uint64 `json:"max_post_bytes"`
	MaxPostRecords        uint64 `json:"max_post_records"`
	MaxRecordPayloadBytes uint64 `json:"max_record_payload_bytes"`
	MaxRequestBytes       uint64 `json:"max_request_bytes"`
	MaxTotalBytes         uint64 `json:"max_total_bytes"`
	MaxTotalRecords       uint64 `json:"max_total_records"`
	MaxQuotaLimit         uint64 `json:"max_quota_limit"`
}
