package types

type BasicStorageObjectSortOrder = uint8

const (
	BasicStorageObjectSortOrderUndefined BasicStorageObjectSortOrder = iota
	BasicStorageObjectSortOrderNewest
	BasicStorageObjectSortOrderOldest
	BasicStorageObjectSortOrderIndex
)
