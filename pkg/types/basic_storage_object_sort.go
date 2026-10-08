package types

import "time"

type BatchBasicStorageObject struct {
	UserID       uint64
	CollectionID uint64
	BatchID      uint64
	BatchBsoID   uint64
	SortIndex    int64
	Payload      string
	Ttl          time.Duration
}
