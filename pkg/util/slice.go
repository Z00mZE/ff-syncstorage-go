package util

func SliceToMap[E any, K comparable, V any](collections []E, fn func(E) (K, V)) map[K]V {
	out := make(map[K]V, len(collections))
	if len(collections) == 0 {
		return out
	}
	for _, record := range collections {
		k, v := fn(record)
		out[k] = v
	}
	return out
}
