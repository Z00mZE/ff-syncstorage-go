package types

import "time"

// StorageObject (BSO) is the generic JSON wrapper around all items passed into and out of the SyncStorage server.
// Like all JSON documents, BSOs are composed of unicode character data rather than raw bytes and must be encoded for
// transmission over the network. The SyncStorage service always encodes BSOs in UTF8.
type StorageObject struct {
	//	An identifying string. For a user, the id must be unique for a BSO within a collection, though objects in
	//	different collections may have the same ID.
	//
	//	BSO ids must only contain printable ASCII characters. They should be exactly 12 base64-urlsafe characters;
	//	while this isn’t enforced by the server, the Firefox client expects it in most cases.
	ID string
	//The timestamp at which this object was last modified, in seconds since UNIX epoch (1970-01-01 00:00:00 UTC).
	//This is set automatically by the server according to its own clock; any client-supplied value for this field is ignored.
	Modified time.Time
	// An integer indicating the relative importance of this item in the collection.
	SortIndex int
	// A string containing the data of the record. The structure of this string is defined separately for each BSO type.
	// This spec makes no requirements for its format. In practice, JSONObjects are common.
	//
	//Servers must support payloads up to 256KiB in size. They may accept larger payloads, and advertise their maximum
	// payload size via dynamic configuration.
	Payload string
	// The number of seconds to keep this record. After that time this item will no longer be returned in response to
	// any request, and it may be pruned from the database. If not specified or null, the record will not expire.
	//
	// This field may be set on write, but is not returned by the server.
	TTL time.Duration
}
