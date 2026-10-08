package types

import "time"

type UserCollection struct {
	UserID       uint64
	CollectionID uint64
	Modified     time.Time
	Count        uint64
	TotalBytes   uint64
}
